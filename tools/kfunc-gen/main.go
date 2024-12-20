package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/cilium/ebpf/btf"
	"gopkg.in/yaml.v3"
)

var projectroot = flag.String("project-root", "", "Root of the project")

type kfuncs struct {
	Sets map[string]idSet `yaml:"sets"`
}

type idSet struct {
	Funcs        []kfunc  `yaml:"funcs"`
	ProgramTypes []string `yaml:"program_types"`
}

type kfunc struct {
	Name  string   `yaml:"name"`
	Flags []string `yaml:"flags"`
}

const (
	kfuncDefStart     = `<!-- [KFUNC_DEF] -->`
	kfuncDefEnd       = `<!-- [/KFUNC_DEF] -->`
	kfuncProgRefStart = `<!-- [KFUNC_PROG_REF] -->`
	kfuncProgRefEnd   = `<!-- [/KFUNC_PROG_REF] -->`

	progKfuncRefStart = `<!-- [PROG_KFUNC_REF] -->`
	progKfuncRefEnd   = `<!-- [/PROG_KFUNC_REF] -->`
)

// List of kfuncs which we purposefully ignore in the data file
var ignoreKfuncs = []string{
	// These are technically usable kfuncs, but they do not do anything useful.
	// They are just here for testing purposes. So we will not document them.
	"bpf_fentry_test1",
	"bpf_modify_return_test",
	"bpf_modify_return_test2",
	"bpf_modify_return_test_tp",
	"bpf_kfunc_call_memb_release",
	"bpf_kfunc_call_test_release",
}

// List of kfuncs which are known to have been removed
var removeKfuncs = []string{
	"hid_bpf_attach_prog",
}

func main() {
	flag.Parse()

	var kfuncsConfig kfuncs

	kfuncData, err := os.ReadFile(*projectroot + "/data/kfuncs.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(kfuncData, &kfuncsConfig)
	if err != nil {
		panic(err)
	}

	var allDataKfuncs []kfunc
	for _, set := range kfuncsConfig.Sets {
		allDataKfuncs = append(allDataKfuncs, set.Funcs...)
	}

	spec, err := btf.LoadSpec(*projectroot + "/tools/kfunc-gen/vmlinux")
	if err != nil {
		panic(err)
	}

	throw := false

	var allBTFKfuncs []*btf.Func
	iter := spec.Iterate()
	for iter.Next() {
		switch t := (iter.Type).(type) {
		case *btf.Func:
			if slices.Contains(t.Tags, "bpf_kfunc") {
				if slices.Contains(ignoreKfuncs, t.Name) {
					continue
				}

				allBTFKfuncs = append(allBTFKfuncs, t)
				if !slices.ContainsFunc(allDataKfuncs, func(k kfunc) bool {
					return k.Name == t.Name
				}) {
					fmt.Printf("Missing kfunc in data file: '%s', possibly newly added\n", t.Name)
					throw = true
				}
			}
		}
	}
	for _, k := range allDataKfuncs {
		if slices.Contains(removeKfuncs, k.Name) {
			continue
		}

		if !slices.ContainsFunc(allBTFKfuncs, func(f *btf.Func) bool {
			return f.Name == k.Name
		}) {
			fmt.Printf("Missing kfunc in BTF: '%s', possibly deleted from kernel\n", k.Name)
			throw = true
		}
	}

	if throw {
		os.Exit(1)
	}

	for _, set := range kfuncsConfig.Sets {
		for _, kfunc := range set.Funcs {
			if slices.Contains(removeKfuncs, kfunc.Name) {
				continue
			}

			file, err := os.OpenFile(*projectroot+"/docs/linux/kfuncs/"+kfunc.Name+".md", os.O_RDWR, 0644)
			if err != nil {
				panic(err)
			}

			fileContents, err := io.ReadAll(file)
			if err != nil {
				panic(err)
			}
			fileStr := string(fileContents)

			var fn *btf.Func
			err = spec.TypeByName(kfunc.Name, &fn)
			if err != nil {
				fmt.Printf("%s: not found\n", kfunc.Name)
				continue
			}

			sig := cFuncSignature(fn)

			startIdx := strings.Index(fileStr, kfuncDefStart)
			endIdx := strings.Index(fileStr, kfuncDefEnd)

			if startIdx == -1 || endIdx == -1 {
				continue
			}

			var newFile strings.Builder
			// Write everything before the marker
			newFile.WriteString(fileStr[:startIdx])
			newFile.WriteString(kfuncDefStart)

			newFile.WriteString(fmt.Sprintf("\n`#!c %s`\n", sig))

			for _, flag := range kfunc.Flags {
				switch flag {
				case "KF_ACQUIRE":
					newFile.WriteString(kfAcquireNotice)
				case "KF_RELEASE":
					newFile.WriteString(kfReleaseNotice)
				case "KF_RET_NULL":
					newFile.WriteString(kfRetNullNotice)
				case "KF_TRUSTED_ARGS":
				case "KF_SLEEPABLE":
					newFile.WriteString(kfSleepableNotice)
				case "KF_DESTRUCTIVE":
					newFile.WriteString(kfDestructiveNotice)
				case "KF_RCU":
				case "KF_ITER_NEW":
				case "KF_ITER_NEXT":
				case "KF_ITER_DESTROY":
				case "KF_RCU_PROTECTED":
					newFile.WriteString(kfRCUProtectedNotice)
				}
			}

			newFile.WriteString(kfuncDefEnd)
			newFile.WriteString(fileStr[endIdx+len(kfuncDefEnd):])

			_, err = file.Seek(0, 0)
			if err != nil {
				panic(err)
			}

			err = file.Truncate(0)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(file, strings.NewReader(newFile.String()))
			if err != nil {
				panic(err)
			}
			file.Close()
		}
	}

	progToKfunc := make(map[string][]string)

	for _, set := range kfuncsConfig.Sets {
		for _, kfunc := range set.Funcs {
			file, err := os.OpenFile(*projectroot+"/docs/linux/kfuncs/"+kfunc.Name+".md", os.O_RDWR, 0644)
			if err != nil {
				panic(err)
			}

			fileContents, err := io.ReadAll(file)
			if err != nil {
				panic(err)
			}

			var progTypes []string
			for _, progType := range set.ProgramTypes {
				if progType == "BPF_PROG_TYPE_UNSPEC" {
					progTypes = append(progTypes, kfuncProgramTypes...)
				} else {
					progTypes = append(progTypes, progType)
				}
			}

			sort.Strings(progTypes)
			progTypes = slices.Compact(progTypes)

			for _, progType := range progTypes {
				progToKfunc[progType] = append(progToKfunc[progType], kfunc.Name)
			}

			fileStr := string(fileContents)

			startIdx := strings.Index(fileStr, kfuncProgRefStart)
			endIdx := strings.Index(fileStr, kfuncProgRefEnd)

			if startIdx == -1 || endIdx == -1 {
				continue
			}

			var newFile strings.Builder
			// Write everything before the marker
			newFile.WriteString(fileStr[:startIdx])
			newFile.WriteString(kfuncProgRefStart)

			newFile.WriteString("\n")

			for _, progType := range progTypes {
				newFile.WriteString(fmt.Sprintf("- [`%s`](../program-type/%s.md)\n", progType, progType))
			}

			newFile.WriteString(kfuncProgRefEnd)
			newFile.WriteString(fileStr[endIdx+len(kfuncProgRefEnd):])

			_, err = file.Seek(0, 0)
			if err != nil {
				panic(err)
			}

			err = file.Truncate(0)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(file, strings.NewReader(newFile.String()))
			if err != nil {
				panic(err)
			}

			file.Close()
		}
	}

	progDirEntries, err := os.ReadDir(*projectroot + "/docs/linux/program-type")
	if err != nil {
		panic(err)
	}
	for _, dirEntry := range progDirEntries {
		if dirEntry.IsDir() {
			continue
		}

		fileName := dirEntry.Name()
		progName := strings.TrimSuffix(fileName, ".md")
		filePath := *projectroot + "/docs/linux/program-type/" + fileName
		file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}

		fileContents, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		fileStr := string(fileContents)

		startIdx := strings.Index(fileStr, progKfuncRefStart)
		endIdx := strings.Index(fileStr, progKfuncRefEnd)

		if startIdx == -1 || endIdx == -1 {
			continue
		}

		var newFile strings.Builder

		// Write everything before the marker
		newFile.WriteString(fileStr[:startIdx])
		newFile.WriteString(progKfuncRefStart)
		newFile.WriteString("\n")

		kfuncs, ok := progToKfunc[progName]
		if ok {
			newFile.WriteString("??? abstract \"Supported kfuncs\"\n")

			sort.Strings(kfuncs)

			for _, kfunc := range kfuncs {
				newFile.WriteString(fmt.Sprintf("    - [`%s`](../kfuncs/%s.md)\n", kfunc, kfunc))
			}
		} else {
			newFile.WriteString("There are currently no kfuncs supported for this program type\n")
		}

		newFile.WriteString(progKfuncRefEnd)
		newFile.WriteString(fileStr[endIdx+len(progKfuncRefEnd):])

		_, err = file.Seek(0, 0)
		if err != nil {
			panic(err)
		}

		err = file.Truncate(0)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(file, strings.NewReader(newFile.String()))
		if err != nil {
			panic(err)
		}

		file.Close()
	}
}

var kfuncProgramTypes = []string{
	"BPF_PROG_TYPE_XDP",
	"BPF_PROG_TYPE_SCHED_CLS",
	"BPF_PROG_TYPE_STRUCT_OPS",
	"BPF_PROG_TYPE_TRACING",
	"BPF_PROG_TYPE_LSM",
	"BPF_PROG_TYPE_SYSCALL",
	"BPF_PROG_TYPE_CGROUP_SKB",
	"BPF_PROG_TYPE_CGROUP_SOCK_ADDR",
	"BPF_PROG_TYPE_SCHED_ACT",
	"BPF_PROG_TYPE_SK_SKB",
	"BPF_PROG_TYPE_SOCKET_FILTER",
	"BPF_PROG_TYPE_LWT_OUT",
	"BPF_PROG_TYPE_LWT_IN",
	"BPF_PROG_TYPE_LWT_XMIT",
	"BPF_PROG_TYPE_LWT_SEG6LOCAL",
	"BPF_PROG_TYPE_NETFILTER",
}

func cFuncSignature(fn *btf.Func) string {
	proto := fn.Type.(*btf.FuncProto)
	args := make([]string, len(proto.Params))
	for i, param := range proto.Params {
		if p, ok := param.Type.(*btf.Pointer); ok {
			if fp, ok := p.Target.(*btf.FuncProto); ok {
				params := make([]string, len(fp.Params))
				for i, param := range fp.Params {
					params[i] = fmt.Sprintf("%s %s", typeToC(param.Type), param.Name)
				}
				args[i] = fmt.Sprintf("%s (%s)(%s)", typeToC(fp.Return), param.Name, strings.Join(params, ", "))
			} else {
				args[i] = fmt.Sprintf("%s%s", typeToC(param.Type), param.Name)
			}
		} else {
			args[i] = fmt.Sprintf("%s %s", typeToC(param.Type), param.Name)
		}
	}

	if _, ok := proto.Return.(*btf.Pointer); ok {
		return fmt.Sprintf("%s%s(%s)", typeToC(proto.Return), fn.Name, strings.Join(args, ", "))
	} else {
		return fmt.Sprintf("%s %s(%s)", typeToC(proto.Return), fn.Name, strings.Join(args, ", "))
	}
}

func typeToC(t btf.Type) string {
	switch t := t.(type) {
	case *btf.Int:
		return t.Name
	case *btf.Struct:
		return "struct " + t.Name
	case *btf.Pointer:
		return typeToC(t.Target) + " *"
	case *btf.Array:
		return fmt.Sprintf("%s[%d]", typeToC(t.Type), t.Nelems)
	case *btf.Func:
		return t.Name
	case *btf.Enum:
		return t.Name
	case *btf.Union:
		return "union " + t.Name
	case *btf.Volatile:
		return "volatile " + typeToC(t.Type)
	case *btf.Const:
		return "const " + typeToC(t.Type)
	case *btf.Restrict:
		return "restrict " + typeToC(t.Type)
	case *btf.Void:
		return "void"
	case *btf.Typedef:
		return t.Name
	default:
		return fmt.Sprintf("unknown (%T)", t)
	}
}

const (
	kfSleepableNotice = `
!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
`

	kfAcquireNotice = `
!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [` + "`bpf_kptr_xchg`" + `](../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.
`

	kfReleaseNotice = `
!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.
`

	kfRetNullNotice = `
!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
`

	kfDestructiveNotice = `
!!! warning
	This kfunc is destructive to the system. For example such a call can result in system rebooting or panicking. 
	Due to this additional restrictions apply to these calls. At the moment they only require ` + "`" + "CAP_SYS_BOOT" + "`" + `capability, 
	but more can be added later.
`

	kfRCUProtectedNotice = `
!!! note
	This kfunc is RCU protected. This means that the kfunc can be called from RCU read-side critical section.
	If a program isn't called from RCU read-side critical section, such as sleepable programs, the 
	[` + "`" + `bpf_rcu_read_lock` + "`" + `](../kfuncs/bpf_rcu_read_lock.md) and 
	[` + "`" + `bpf_rcu_read_unlock` + "`" + `](../kfuncs/bpf_rcu_read_unlock.md) to protect the calls to such KFuncs.
`
)
