package main

import (
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
)

func printTableMatrix(w io.Writer) {
	arches := []Arch{
		x86,
		arm,
		arm64,
		armV7,
		riscv,
	}

	for _, frags := range []bool{false, true} {
		vars := defaultVars()

		maxWidths := make([]int, len(arches))
		for _, row := range varTable {
			for _, subrow := range row.Subrow {
				for i, arch := range arches {
					vars.Arch = arch
					vars.Frags = frags
					var mtu string
					if subrow.Render == nil {
						mtu = RenderDefault(subrow, vars)
					} else {
						mtu = subrow.Render(subrow, vars)
					}
					width := len([]rune(mtu))
					if width > maxWidths[i] {
						maxWidths[i] = width
					}
				}
			}
		}

		if frags {
			fmt.Fprintln(w, "\n=== \"XDP with Fragments\"")
		} else {
			fmt.Fprintln(w, "\n=== \"Plain XDP\"")
		}
		fmt.Fprintln(w)
		fmt.Fprint(w, "    | Vendor           | Driver                |")
		for i, arch := range arches {
			fmt.Fprintf(w, " %-*s |", maxWidths[i], arch.String())
		}
		fmt.Fprintln(w)
		fmt.Fprint(w, "    | ---------------- | --------------------- |")
		for i := range arches {
			fmt.Fprint(w, " ", strings.Repeat("-", maxWidths[i]), " |")
		}
		fmt.Fprintln(w)

		for _, row := range varTable {
			for _, subrow := range row.Subrow {
				fmt.Fprintf(w, "    | %-16s | %-21s |",
					row.Vendor,
					subrow.Driver,
				)

				for i, arch := range arches {
					vars.Arch = arch
					vars.Frags = frags
					var mtu string
					if subrow.Render == nil {
						mtu = RenderDefault(subrow, vars)
					} else {
						mtu = subrow.Render(subrow, vars)
					}
					fmt.Fprintf(w, " %-*s |", maxWidths[i], mtu)
				}
				fmt.Fprintln(w)
			}
		}
		fmt.Fprintln(w)
	}
}

type footnote string

func (f footnote) String() string {
	return fmt.Sprintf("%s: %s", f.Ref(), string(f))
}

func (f footnote) Ref() string {
	return fmt.Sprintf("[^%d]", slices.Index(tableFootnotes, f)+1)
}

var (
	infinityFootnote    = footnote("Driver has no MTU limit")
	nicSpecificFootnote = footnote("MTU limit specified by firmware")
	tunFootnote         = footnote("Depends on slave device(s)")
	hdsFootnote         = footnote("MTU limit depends on HDS threshold")
	tableFootnotes      = []footnote{
		footnote("reserved"),
		infinityFootnote,
		nicSpecificFootnote,
		tunFootnote,
		hdsFootnote,
	}
)

var varTable = []tableRow{
	{
		Vendor: "Kernel",
		Subrow: []tableSubrow{
			{
				Driver:       "Veth",
				VersionXDP:   KernelVersion{4, 19, 0},
				VersionFrags: KernelVersion{5, 5, 0},
				DriverFunc:   calcVethMTU,
			},
			{
				Driver:       "VirtIO",
				VersionXDP:   KernelVersion{4, 10, 0},
				VersionFrags: KernelVersion{6, 3, 0},
				DriverFunc:   calcVirtioNetMTU,
			},
			{
				Driver:     "Tun",
				VersionXDP: KernelVersion{4, 14, 0},
				DriverFunc: calcTunMTU,
			},
			{
				Driver:     "Bond",
				VersionXDP: KernelVersion{5, 15, 0},
				Render: func(subrow tableSubrow, vars vars) string {
					if vars.KernelVersion.Less(subrow.VersionXDP) {
						return ":material-close:"
					}

					return tunFootnote.Ref()
				},
			},
		},
	},
	{
		Vendor: "Xen",
		Subrow: []tableSubrow{
			{
				Driver:     "Netfront",
				VersionXDP: KernelVersion{5, 9, 0},
				DriverFunc: calcXenNetfrontMTU,
			},
		},
	},
	{
		Vendor: "Amazon",
		Subrow: []tableSubrow{
			{
				Driver:     "ENA",
				VersionXDP: KernelVersion{5, 6, 0},
				DriverFunc: calcEnaMTU,
			},
		},
	},
	{
		Vendor: "Aquantia/Marvell",
		Subrow: []tableSubrow{
			{
				Driver:       "AQtion",
				VersionXDP:   KernelVersion{5, 19, 0},
				VersionFrags: KernelVersion{5, 19, 0},
				DriverFunc:   calcAQMTU,
			},
		},
	},
	{
		Vendor: "Broadcom",
		Subrow: []tableSubrow{
			{
				Driver:       "BNXT",
				VersionXDP:   KernelVersion{4, 11, 0},
				VersionFrags: KernelVersion{5, 19, 0},
				DriverFunc:   calcBNXTMTU,
			},
		},
	},
	{
		Vendor: "Cavium",
		Subrow: []tableSubrow{
			{
				Driver:     "Thunder (nicvf)",
				VersionXDP: KernelVersion{4, 12, 0},
				DriverFunc: calcNICVFMTU,
			},
		},
	},
	{
		Vendor: "Engelder",
		Subrow: []tableSubrow{
			{
				Driver:     "TSN Endpoint",
				VersionXDP: KernelVersion{6, 3, 0},
				DriverFunc: calcTsnepMTU,
			},
		},
	},
	{
		Vendor: "Freescale",
		Subrow: []tableSubrow{
			{
				Driver:     "FEC",
				VersionXDP: KernelVersion{6, 2, 0},
				DriverFunc: calcFecMTU,
			},
			{
				Driver:     "DPAA",
				VersionXDP: KernelVersion{5, 11, 0},
				DriverFunc: calcDPAAMTU,
			},
			{
				Driver:     "DPAA2",
				VersionXDP: KernelVersion{5, 0, 0},
				Render: func(subrow tableSubrow, vars vars) string {
					if vars.KernelVersion.Less(subrow.VersionXDP) {
						return ":material-close:"
					}

					return "?" + nicSpecificFootnote.Ref()
				},
			},
			{
				Driver:     "ENETC",
				VersionXDP: KernelVersion{5, 13, 0},
				DriverFunc: calcENETCMTU,
			},
		},
	},
	{
		Vendor: "Fungible",
		Subrow: []tableSubrow{
			{
				Driver:     "Funeth",
				VersionXDP: KernelVersion{5, 18, 0},
				DriverFunc: calcFunMTU,
			},
		},
	},
	{
		Vendor: "Google",
		Subrow: []tableSubrow{
			{
				Driver:     "GVE",
				VersionXDP: KernelVersion{6, 4, 0},
				DriverFunc: calcGVEMTU,
			},
		},
	},
	{
		Vendor: "Intel",
		Subrow: []tableSubrow{
			{
				Driver:       "I40e",
				VersionXDP:   KernelVersion{4, 13, 0},
				VersionFrags: KernelVersion{6, 4, 0},
				DriverFunc:   calcI40EMTU,
			},
			{
				Driver:       "ICE",
				VersionXDP:   KernelVersion{5, 5, 0},
				VersionFrags: KernelVersion{6, 3, 0},
				DriverFunc:   calcICEMTU,
			},
			{
				Driver:     "IGB",
				VersionXDP: KernelVersion{5, 10, 0},
				DriverFunc: calcIGBMTU,
			},
			{
				Driver:     "IGC",
				VersionXDP: KernelVersion{5, 13, 0},
				DriverFunc: calcIGCMTU,
			},
			{
				Driver:     "IXGBE",
				VersionXDP: KernelVersion{4, 12, 0},
				DriverFunc: calcIXGBE,
			},
			{
				Driver:     "IXGBEVF",
				VersionXDP: KernelVersion{4, 17, 0},
				DriverFunc: calcIXGBEVF,
			},
			{
				Driver:     "IDPF",
				VersionXDP: KernelVersion{6, 18, 0},
				DriverFunc: calcIDPFMTU,
			},
		},
	},
	{
		Vendor: "Marvell",
		Subrow: []tableSubrow{
			{
				Driver:       "NETA",
				VersionXDP:   KernelVersion{5, 5, 0},
				VersionFrags: KernelVersion{5, 18, 0},
				DriverFunc:   calcMVNetaMTU,
			},
			{
				Driver:       "PPv2",
				VersionXDP:   KernelVersion{5, 9, 0},
				VersionFrags: KernelVersion{5, 18, 0},
				DriverFunc:   calcMVPP2MTU,
			},
			{
				Driver:     "Octeon TX2",
				VersionXDP: KernelVersion{5, 16, 0},
				DriverFunc: calcOTX2MTU,
			},
		},
	},
	{
		Vendor: "MediaTek",
		Subrow: []tableSubrow{
			{
				Driver:     "MTK",
				VersionXDP: KernelVersion{6, 0, 0},
				DriverFunc: calcMtkMTU,
			},
		},
	},
	{
		Vendor: "Mellanox",
		Subrow: []tableSubrow{
			{
				Driver:     "MLX4",
				VersionXDP: KernelVersion{4, 8, 0},
				DriverFunc: calcMLX4MTU,
			},
			{
				Driver:       "MLX5",
				VersionXDP:   KernelVersion{4, 9, 0},
				VersionFrags: KernelVersion{5, 18, 0},
				DriverFunc:   calcMLX5MTU,
			},
		},
	},
	{
		Vendor: "Microchip",
		Subrow: []tableSubrow{
			{
				Driver:     "LAN966x",
				VersionXDP: KernelVersion{6, 2, 0},
				DriverFunc: calcLan966xMTU,
			},
		},
	},
	{
		Vendor: "Microsoft",
		Subrow: []tableSubrow{
			{
				Driver:     "Mana",
				VersionXDP: KernelVersion{5, 17, 0},
				DriverFunc: calcManaMTU,
			},
			{
				Driver:     "Hyper-V",
				VersionXDP: KernelVersion{5, 6, 0},
				DriverFunc: calcNetvscMTU,
			},
		},
	},
	{
		Vendor: "Netronome",
		Subrow: []tableSubrow{
			{
				Driver:     "NFP",
				VersionXDP: KernelVersion{4, 10, 0},
				DriverFunc: calcNFPMTU,
			},
		},
	},
	{
		Vendor: "Pensando",
		Subrow: []tableSubrow{
			{
				Driver:     "Ionic",
				VersionXDP: KernelVersion{6, 9, 0},
				DriverFunc: calcIonicMTU,
			},
		},
	},
	{
		Vendor: "Qlogic",
		Subrow: []tableSubrow{
			{
				Driver:     "QEDE",
				VersionXDP: KernelVersion{4, 10, 0},
				DriverFunc: calcQedeMTU,
			},
		},
	},
	{
		Vendor: "Solarflare",
		Subrow: []tableSubrow{
			{
				Driver:     "SFP (SFC9xxx PF/VF)",
				VersionXDP: KernelVersion{4, 10, 0},
				DriverFunc: calcSFPMTUModel(hunt),
			},
			{
				Driver:     "SFP (Riverhead)",
				VersionXDP: KernelVersion{4, 10, 0},
				DriverFunc: calcSFPMTUModel(ef100),
			},
			{
				Driver:     "SFP (SFC4000A)",
				VersionXDP: KernelVersion{4, 10, 0},
				DriverFunc: calcSFPMTUModel(falcon_a1),
			},
			{
				Driver:     "SFP (SFC4000B)",
				VersionXDP: KernelVersion{4, 10, 0},
				DriverFunc: calcSFPMTUModel(falcon_b0),
			},
			{
				Driver:     "SFP (SFC9020/SFL9021)",
				VersionXDP: KernelVersion{4, 10, 0},
				DriverFunc: calcSFPMTUModel(siena_a0),
			},
		},
	},
	{
		Vendor: "Socionext",
		Subrow: []tableSubrow{
			{
				Driver:     "NetSec",
				VersionXDP: KernelVersion{5, 3, 0},
				DriverFunc: calcNetsecMTU,
			},
		},
	},
	{
		Vendor: "STMicro",
		Subrow: []tableSubrow{
			{
				Driver:     "ST MAC",
				VersionXDP: KernelVersion{5, 13, 0},
				DriverFunc: calcStmmacMTU,
			},
		},
	},
	{
		Vendor: "TI",
		Subrow: []tableSubrow{
			{
				Driver:     "CPSW",
				VersionXDP: KernelVersion{5, 3, 0},
				DriverFunc: calcCPSWMTU,
			},
			{
				Driver:     "ICSSG",
				VersionXDP: KernelVersion{6, 15, 0},
				DriverFunc: calcICSSGMTU,
			},
		},
	},
	{
		Vendor: "VMWare",
		Subrow: []tableSubrow{
			{
				Driver:     "VMXNET 3",
				VersionXDP: KernelVersion{6, 6, 0},
				DriverFunc: calcVmxnet3MTU,
			},
		},
	},
	{
		Vendor: "Meta",
		Subrow: []tableSubrow{
			{
				Driver:     "FBNIC",
				VersionXDP: KernelVersion{6, 18, 0},
				Render: func(subrow tableSubrow, vars vars) string {
					if vars.Frags {
						return "∞" + infinityFootnote.Ref()
					}

					return "128 ≤ " + strconv.Itoa(calcFBNICMTUDefault(vars)) + "(default) ≤ " + strconv.Itoa(calcFBNICMTUMax(vars)) + hdsFootnote.Ref()
				},
			},
		},
	},
}

type tableRow struct {
	Vendor string
	Subrow []tableSubrow
}

type tableSubrow struct {
	Driver       string
	VersionXDP   KernelVersion
	VersionFrags KernelVersion
	DriverFunc   func(vars vars) int
	Render       func(subrow tableSubrow, vars vars) string
}

func RenderDefault(subrow tableSubrow, vars vars) string {
	if vars.KernelVersion.Less(subrow.VersionXDP) {
		return ":material-close:"
	}

	mtu := subrow.DriverFunc(vars)

	if vars.Frags && subrow.VersionFrags == KernelVersionUnknown {
		return ":material-close:"
	}

	if mtu == math.MaxInt {
		return "∞" + infinityFootnote.Ref()
	}

	return fmt.Sprint(mtu)
}
