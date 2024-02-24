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
)

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

	spec, err := btf.LoadSpec(*projectroot + "/tools/kfunc-gen/vmlinux")
	if err != nil {
		panic(err)
	}

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

			fileStr := string(fileContents)

			startIdx := strings.Index(fileStr, kfuncProgRefStart)
			endIdx := strings.Index(fileStr, kfuncProgRefEnd)

			var newFile strings.Builder
			// Write everything before the marker
			newFile.WriteString(fileStr[:startIdx])
			newFile.WriteString(kfuncProgRefStart)

			newFile.WriteString("\n")
			var progTypes []string
			for _, progType := range set.ProgramTypes {
				if progType == "BPF_PROG_TYPE_UNSPEC" {
					progTypes = append(progTypes, kfuncProgramTypes...)
				} else {
					progTypes = append(progTypes, progType)
				}
			}
			sort.Strings(progTypes)
			slices.Compact(progTypes)
			for _, progType := range progTypes {
				newFile.WriteString(fmt.Sprintf("- [%s](../program-type/%s.md)\n", progType, progType))
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
    This function may sleep, and therefore can only be used from [sleepable programs](../../syscall/BPF_PROG_LOAD/#bpf_f_sleepable).
`

	kfAcquireNotice = `
!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [` + "`bpf_kptr_xchg`" + `](../../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
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
	Due to this additional restrictions apply to these calls. At the moment they only require CAP_SYS_BOOT capability, 
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
