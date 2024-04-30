package main

type Arch int

const (
	alpha Arch = iota
	alphaGeneric
	alphaEV6
	arc
	arm
	armV6
	armV7
	arm64
	csky610
	csky807
	csky810
	csky860
	hexagon
	loongarch
	m68k
	microblaze
	mips
	nios2
	openrisc
	parisc
	powerpc
	riscv
	s390
	sh
	sparc32
	sparc64
	um
	x86
	xtensa
)

func (a Arch) String() string {
	switch a {
	case alpha:
		return "alpha"
	case alphaGeneric:
		return "alpha-generic"
	case alphaEV6:
		return "alpha-ev6"
	case arc:
		return "arc"
	case arm:
		return "arm"
	case armV6:
		return "armv6"
	case armV7:
		return "armv7"
	case arm64:
		return "arm64"
	case csky610:
		return "csky610"
	case csky807:
		return "csky807"
	case csky810:
		return "csky810"
	case csky860:
		return "csky860"
	case hexagon:
		return "hexagon"
	case loongarch:
		return "loongarch"
	case m68k:
		return "m68k"
	case microblaze:
		return "microblaze"
	case mips:
		return "mips"
	case nios2:
		return "nios2"
	case openrisc:
		return "openrisc"
	case parisc:
		return "parisc"
	case powerpc:
		return "powerpc"
	case riscv:
		return "riscv"
	case s390:
		return "s390"
	case sh:
		return "sh"
	case sparc32:
		return "sparc32"
	case sparc64:
		return "sparc64"
	case um:
		return "um"
	case x86:
		return "x86"
	case xtensa:
		return "xtensa"
	}
	return "unknown"

}

type KernelVersion struct {
	Major int
	Minor int
	Patch int
}

func (v KernelVersion) Less(o KernelVersion) bool {
	if v.Major != o.Major {
		return v.Major < o.Major
	}
	if v.Minor != o.Minor {
		return v.Minor < o.Minor
	}
	return v.Patch < o.Patch
}

func (v KernelVersion) Greater(o KernelVersion) bool {
	if v.Major != o.Major {
		return v.Major > o.Major
	}
	if v.Minor != o.Minor {
		return v.Minor > o.Minor
	}
	return v.Patch > o.Patch
}

var (
	KernelVersionUnknown = KernelVersion{0, 0, 0}
	KernelVersion4_19    = KernelVersion{4, 19, 0}
	KernelVersion5_4     = KernelVersion{5, 4, 0}
	KernelVersion5_10    = KernelVersion{5, 10, 0}
	KernelVersion5_15    = KernelVersion{5, 15, 0}
	KernelVersion6_1     = KernelVersion{6, 1, 0}
	KernelVersion6_6     = KernelVersion{6, 6, 0}
	KernelVersion6_9     = KernelVersion{6, 9, 0}
)

type SFCModel int

const (
	hunt SFCModel = iota
	ef100
	falcon_a1
	falcon_b0
	siena_a0
)

type vars struct {
	Frags             bool
	Arch              Arch
	PageSize          int
	XDPPacketheadroom int
	KernelVersion     KernelVersion
	ConfigSKBFrags    *int

	// veth specific
	PeerHardHeaderLen int

	// DPAA specific
	ERRATUM_A050385 bool

	// I40e specific
	I40ELegacyRXENA bool

	// ICE specific
	ICELegacyRX bool

	// IGB specific
	IGBUseLargeRing bool
	IGBUseSKB       bool

	// IXGBE specific
	IXGBRX3kBuffer bool
	IXGBUseSKB     bool

	// IXGBEVF specific
	IXGBEVLargeBuffer bool
	IXGBEVFUseSKB     bool

	// SFP specific
	SFCModel SFCModel
}

func (v vars) NetIPAlign() int {
	switch v.Arch {
	case x86, arm64, powerpc:
		return 0
	}
	return 2
}

func (v vars) SMPCacheBytes() int {
	switch v.Arch {
	case alpha:
		return 32
	case alphaGeneric, alphaEV6:
		return 64
	case arc:
		return 128
	case microblaze:
		return 1 << 5
	case parisc:
		return 16
	case sparc32:
		return 1 << 5
	case sparc64:
		return 1 << 6
	case xtensa:
		return 32
	default:
		return v.L1CacheBytes()
	}
}

func (v vars) L1CacheBytes() int {
	switch v.Arch {
	case arc:
		return 1 << 6
	case arm:
		return 1 << 5
	case armV6:
		return 1 << 6
	case armV7:
		return 1 << 7
	case arm64:
		return 1 << 6
	case csky610:
		return 1 << 4
	case csky807, csky810:
		return 1 << 5
	case csky860:
		return 1 << 6
	case hexagon:
		return 1 << 5
	case loongarch:
		return 1 << 6
	case microblaze:
		return 1 << 5
	case mips:
		return 1 << 5
	case nios2:
		return 1 << 5
	case openrisc:
		return 16
	case parisc:
		return 16
	case riscv:
		return 1 << 6
	case s390:
		return 256
	case sh:
		return 1 << 5
	case sparc32:
		return 32
	case um:
		return 1 << 5
	case x86:
		return 1 << 6
	case xtensa:
		return 1 << 5
	}

	panic("missing L1 cache size for arch")
}

func (v vars) SizeOfSKBSharedInfo() int {
	// NOTE: manual checking of multiple distros and kernel versions all resulted in 320
	// This is likely not universal, but very complex to determine and model at this time.
	return 320
}

func (v vars) MaxSKBFrags() int {
	switch v.KernelVersion {
	case KernelVersion6_6, KernelVersion6_9:
		if v.ConfigSKBFrags != nil {
			return *v.ConfigSKBFrags
		}
		return 17

	default:
		if (65536/v.PageSize + 1) < 16 {
			return 16
		}
		return 65536/v.PageSize + 1
	}
}

func (v vars) NetSKBPad() int {
	return max(32, v.L1CacheBytes())
}

func defaultVars() vars {
	return vars{
		Frags:             false,
		Arch:              x86,
		PageSize:          4096,
		XDPPacketheadroom: 256,
		KernelVersion:     KernelVersion6_9,

		PeerHardHeaderLen: 0,

		IGBUseLargeRing: true,

		IXGBRX3kBuffer: true,

		IXGBEVLargeBuffer: true,

		SFCModel: hunt,
	}
}
