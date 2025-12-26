package main

import "math"

var drivers = map[string]func(vars) int{
	"veth":               calcVethMTU,
	"tun":                calcTunMTU,
	"virtio_net":         calcVirtioNetMTU,
	"xen-netfront":       calcXenNetfrontMTU,
	"Amazon ENA":         calcEnaMTU,
	"Aquantia Atlantic":  calcAQMTU,
	"Broadcom BNXT":      calcBNXTMTU,
	"Cavium Thunder":     calcNICVFMTU,
	"Engelder":           calcTsnepMTU,
	"Freescale fec_enet": calcFecMTU,
	"Freescale DPAA":     calcDPAAMTU,
	"Freescale DPAA2":    calcDPAA2MTU,
	"Freescale ENETC":    calcENETCMTU,
	"Fungible Funeth":    calcFunMTU,
	"Google GVE":         calcGVEMTU,
	"Intel i40e":         calcI40EMTU,
	"Intel ice":          calcICEMTU,
	"Intel igb":          calcIGBMTU,
	"Intel igc":          calcIGCMTU,
	"Intel ixgbe":        calcIXGBE,
	"Intel ixgbevf":      calcIXGBEVF,
	"Intel idpf":         calcIDPFMTU,
	"Marvell mvneta":     calcMVNetaMTU,
	"Marvell mvpp2":      calcMVPP2MTU,
	"Marvell OcteonTX2":  calcOTX2MTU,
	"Mediatek MTK":       calcMtkMTU,
	"Mellanox mlx4":      calcMLX4MTU,
	"Mellanox mlx5":      calcMLX5MTU,
	"Microchip LAN966x":  calcLan966xMTU,
	"Microsoft Mana":     calcManaMTU,
	"Netronome NFP":      calcNFPMTU,
	"Pensando Ionic":     calcIonicMTU,
	"QLogic qede":        calcQedeMTU,
	"Solarflare SFP":     calcSFPMTU,
	"Socionext Netsec":   calcNetsecMTU,
	"STMicro stmmac":     calcStmmacMTU,
	"TI CPSW":            calcCPSWMTU,
	"Hyper-V Netvsc":     calcNetvscMTU,
	"VMware vmxnet3":     calcVmxnet3MTU,
	"Meta FBNIC":         calcFBNICMTUDefault,
}

func calcVethMTU(v vars) int {
	vethXDPHeadroom := v.XDPPacketheadroom + v.NetIPAlign()
	maxMTU := macroSKBWithOverhead(v.PageSize-vethXDPHeadroom, v) - v.PeerHardHeaderLen

	if !v.Frags {
		return maxMTU
	}

	return maxMTU + v.PageSize*v.MaxSKBFrags()
}

func calcTunMTU(v vars) int {
	return ETH_DATA_LEN
}

func calcVirtioNetMTU(v vars) int {
	const virtioXDPHeadroom = 256
	room := macroSKBDataAlign(virtioXDPHeadroom+v.SizeOfSKBSharedInfo(), v)
	maxSz := v.PageSize - room - ETH_HLEN

	if v.Frags {
		return math.MaxInt
	}

	return maxSz
}

func calcXenNetfrontMTU(v vars) int {
	const xenPageSize = 1 << 12
	maxMTU := xenPageSize - v.XDPPacketheadroom

	return maxMTU
}

func calcEnaMTU(v vars) int {
	const SZ_16K = 0x4000
	var enaPageSize int
	if v.PageSize > SZ_16K {
		enaPageSize = SZ_16K
	} else {
		enaPageSize = v.PageSize
	}

	return enaPageSize - ETH_HLEN - ETH_FCS_LEN - VLAN_HLEN - v.XDPPacketheadroom -
		macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v)
}

func calcAQMTU(v vars) int {
	if v.Frags {
		return math.MaxInt
	}

	const AQ_CFG_RX_FRAME_MAX = 2 * 1024
	return AQ_CFG_RX_FRAME_MAX
}

func calcBNXTMTU(v vars) int {
	if v.Frags {
		return math.MaxInt
	}

	BNXT_MAX_PAGE_MODE_MTU_SBUF := v.PageSize - VLAN_ETH_HLEN - v.NetIPAlign() - v.XDPPacketheadroom
	BNXT_MAX_PAGE_MODE_MTU := BNXT_MAX_PAGE_MODE_MTU_SBUF - macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v)

	return BNXT_MAX_PAGE_MODE_MTU
}

func calcNICVFMTU(v vars) int {
	const MAX_XDP_MTU = 1530 - ETH_HLEN - VLAN_HLEN*2
	return MAX_XDP_MTU
}

func calcTsnepMTU(v vars) int {
	return math.MaxInt
}

func calcFecMTU(v vars) int {
	return math.MaxInt
}

func calcDPAAMTU(v vars) int {
	const (
		DPAA_A050385_ALIGN      = 256
		DPAA_TIME_STAMP_SIZE    = 8
		DPAA_HASH_RESULTS_SIZE  = 8
		DPAA_PARSE_RESULTS_SIZE = 18 // sizeof(struct fman_prs_result)
		DPAA_HWA_SIZE           = DPAA_PARSE_RESULTS_SIZE + DPAA_TIME_STAMP_SIZE + DPAA_HASH_RESULTS_SIZE
	)
	const bpRawSize = 4096

	dpaaBPSize := func(rawSize int) int {
		if v.ERRATUM_A050385 {
			return macroSKBWithOverhead(rawSize, v) & ^(DPAA_A050385_ALIGN - 1)
		}
		return macroSKBWithOverhead(rawSize, v)
	}
	dpaaGetHeadroom := func(priv_data_size int) int {
		headroom := priv_data_size + DPAA_HWA_SIZE
		if v.ERRATUM_A050385 {
			headroom = v.XDPPacketheadroom
		}

		DPAA_FD_RX_DATA_ALIGNMENT := 16
		if v.ERRATUM_A050385 {
			DPAA_FD_RX_DATA_ALIGNMENT = 64
		}

		return macroAlign(headroom, DPAA_FD_RX_DATA_ALIGNMENT)
	}

	bpSize := dpaaBPSize(bpRawSize)
	rxHeadroom := dpaaGetHeadroom(0000)
	maxConfigSize := bpSize - rxHeadroom

	return maxConfigSize - (VLAN_ETH_HLEN + ETH_FCS_LEN)
}

func calcDPAA2MTU(v vars) int {
	// TODO: figure out typical priv->tx_data_offset which is recieved from the NIC
	// without it we can't calculate the MTU
	return 0
}

func calcENETCMTU(v vars) int {
	return math.MaxInt
}

func calcFunMTU(v vars) int {
	const FUN_XDP_HEADROOM = 192
	FUN_RX_TAILROOM := macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v)
	return v.PageSize - FUN_XDP_HEADROOM - VLAN_ETH_HLEN - FUN_RX_TAILROOM
}

func calcGVEMTU(v vars) int {
	const GVE_DEFAULT_RX_BUFFER_SIZE = 2048
	const GVE_RX_PAD = 2
	return GVE_DEFAULT_RX_BUFFER_SIZE - ETH_HLEN - GVE_RX_PAD
}

func calcI40EMTU(v vars) int {
	const (
		I40E_RXBUFFER_2048          = 2048
		I40E_RXBUFFER_3072          = 3072
		I40E_MAX_CHAINED_RX_BUFFERS = 5
		I40E_MAX_RXBUFFER           = 9728
		I40E_PACKET_HDR_PAD         = ETH_HLEN + ETH_FCS_LEN + (VLAN_HLEN * 2)
	)

	var vsiRXBufLen int
	if v.I40ELegacyRXENA {
		vsiRXBufLen = macroSKBWithOverhead(I40E_RXBUFFER_2048, v)
	} else {
		if v.PageSize < 8192 {
			vsiRXBufLen = I40E_RXBUFFER_3072
		} else {
			vsiRXBufLen = I40E_RXBUFFER_2048
		}
	}

	chainLen := 1
	if v.Frags {
		chainLen = I40E_MAX_CHAINED_RX_BUFFERS
	}

	frameSize := min(vsiRXBufLen*chainLen, I40E_MAX_RXBUFFER)
	return frameSize - I40E_PACKET_HDR_PAD
}

func calcICEMTU(v vars) int {
	const (
		ICE_RXBUF_1664      = 1664
		ICE_RXBUF_3072      = 3072
		ICE_ETH_PKT_HDR_PAD = ETH_HLEN + ETH_FCS_LEN + (VLAN_HLEN * 2)
	)

	var maxFrameSize int
	if v.ICELegacyRX {
		maxFrameSize = ICE_RXBUF_1664
	} else {
		maxFrameSize = ICE_RXBUF_3072
	}

	return maxFrameSize - ICE_ETH_PKT_HDR_PAD
}

func calcIGBMTU(v vars) int {
	const (
		IGB_RXBUFFER_1536   = 1536
		IGB_RXBUFFER_2048   = 2048
		IGB_RXBUFFER_3072   = 3072
		IGB_ETH_PKT_HDR_PAD = ETH_HLEN + ETH_FCS_LEN + (VLAN_HLEN * 2)
	)
	IGB_MAX_FRAME_BUILD_SKB := IGB_RXBUFFER_1536 - v.NetIPAlign()

	var rxBufSize int
	if v.PageSize < 8192 {
		if v.IGBUseLargeRing {
			rxBufSize = IGB_RXBUFFER_3072
		} else if v.IGBUseSKB {
			rxBufSize = IGB_MAX_FRAME_BUILD_SKB
		} else {
			rxBufSize = IGB_RXBUFFER_2048
		}
	} else {
		rxBufSize = IGB_RXBUFFER_2048
	}

	return rxBufSize - IGB_ETH_PKT_HDR_PAD
}

func calcIGCMTU(v vars) int {
	return ETH_DATA_LEN
}

func calcIXGBE(v vars) int {
	const (
		IXGBE_RXBUFFER_1536 = 1536
		IXGBE_RXBUFFER_2K   = 2048
		IXGBE_RXBUFFER_3K   = 3072
	)
	IXGBE_MAX_2K_FRAME_BUILD_SKB := IXGBE_RXBUFFER_1536 - v.NetIPAlign()

	var rxBufSize int
	if v.IXGBRX3kBuffer {
		rxBufSize = IXGBE_RXBUFFER_3K
	} else {
		if v.PageSize < 8192 {
			if v.IXGBUseSKB {
				rxBufSize = IXGBE_MAX_2K_FRAME_BUILD_SKB
			} else {
				rxBufSize = IXGBE_RXBUFFER_2K
			}
		} else {
			rxBufSize = IXGBE_RXBUFFER_2K
		}

	}
	return rxBufSize - (ETH_HLEN + ETH_FCS_LEN + VLAN_HLEN)
}

func calcIXGBEVF(v vars) int {
	const (
		IXGBEVF_RXBUFFER_1536 = 1536
		IXGBEVF_RXBUFFER_2048 = 2048
		IXGBEVF_RXBUFFER_3072 = 3072
	)

	var IXGBEVF_MAX_FRAME_BUILD_SKB int
	if v.PageSize < 8192 {
		IXGBEVF_SKB_PAD := v.NetSKBPad() + v.NetIPAlign()
		IXGBEVF_MAX_FRAME_BUILD_SKB = macroSKBWithOverhead(IXGBEVF_RXBUFFER_2048, v) - IXGBEVF_SKB_PAD
	} else {
		IXGBEVF_MAX_FRAME_BUILD_SKB = IXGBEVF_RXBUFFER_2048
	}

	rxBufSize := func() int {
		if v.PageSize < 8192 {
			if v.IXGBEVLargeBuffer {
				return IXGBEVF_RXBUFFER_3072
			}

			if v.IXGBEVFUseSKB {
				return IXGBEVF_MAX_FRAME_BUILD_SKB
			}

		}
		return IXGBEVF_RXBUFFER_2048
	}()

	return rxBufSize - (ETH_HLEN + ETH_FCS_LEN + VLAN_HLEN)
}

func calcIDPFMTU(v vars) int {
	return math.MaxInt
}

func calcMVNetaMTU(v vars) int {
	if v.Frags {
		return math.MaxInt
	}

	MVNETA_SKB_HEADROOM := macroAlign(max(v.NetSKBPad(), v.XDPPacketheadroom), 8)
	MVNETA_SKB_PAD := macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v) + MVNETA_SKB_HEADROOM
	return v.PageSize - MVNETA_SKB_PAD
}

func calcMVPP2MTU(v vars) int {
	MVPP2_SKB_HEADROOM := min(max(v.XDPPacketheadroom, v.NetSKBPad()), 224)
	MVPP2_SKB_SHINFO_SIZE := macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v)
	return v.PageSize - MVPP2_SKB_SHINFO_SIZE - MVPP2_SKB_HEADROOM
}

func calcOTX2MTU(v vars) int {
	OTX2_ETH_HLEN := VLAN_ETH_HLEN + VLAN_HLEN
	return 1530 - OTX2_ETH_HLEN
}

func calcMtkMTU(v vars) int {
	MTK_PP_HEADROOM := v.XDPPacketheadroom
	MTK_PP_PAD := MTK_PP_HEADROOM + macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v)
	return v.PageSize - MTK_PP_PAD
}

func calcMLX4MTU(v vars) int {
	return v.PageSize - ETH_HLEN - (2 * VLAN_HLEN) - v.XDPPacketheadroom - macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v)
}

func calcMLX5MTU(v vars) int {
	if v.Frags {
		return math.MaxInt
	}

	headroom := 256
	hwmtu := v.PageSize - headroom
	MLX5E_ETH_HARD_MTU := (ETH_HLEN + VLAN_HLEN + ETH_FCS_LEN)
	return hwmtu - macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v) - MLX5E_ETH_HARD_MTU
}

func calcLan966xMTU(v vars) int {
	return math.MaxInt
}

func calcManaMTU(v vars) int {
	MANA_RXBUF_PAD := macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v) + ETH_HLEN
	return v.PageSize - MANA_RXBUF_PAD - v.XDPPacketheadroom
}

func calcNFPMTU(v vars) int {
	return v.PageSize
}

func calcIonicMTU(v vars) int {
	return v.PageSize - (VLAN_ETH_HLEN + v.XDPPacketheadroom + macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v))
}

func calcQedeMTU(v vars) int {
	return math.MaxInt
}

func calcSFPMTUModel(model SFCModel) func(v vars) int {
	return func(v vars) int {
		v.SFCModel = model
		return calcSFPMTU(v)
	}
}

func calcSFPMTU(v vars) int {
	var rxPrefixSize int
	switch v.SFCModel {
	case hunt:
		rxPrefixSize = 14
	case ef100:
		rxPrefixSize = 22
	case falcon_a1:
		rxPrefixSize = 0
	case falcon_b0:
		rxPrefixSize = 16
	case siena_a0:
		rxPrefixSize = 16
	}

	var rxBufferPadding int
	switch v.SFCModel {
	case falcon_a1:
		rxBufferPadding = 0x24
	default:
		rxBufferPadding = 0
	}

	var rxIPAlign int
	if v.NetIPAlign() != 0 {
		rxIPAlign = rxPrefixSize + v.NetIPAlign()
	}

	EFX_FRAME_PAD := 16
	maxFrameLen := func(mtu int) int {
		return (macroAlign(mtu+ETH_HLEN+VLAN_HLEN+ETH_FCS_LEN+EFX_FRAME_PAD, 8))
	}

	headroom := 128
	tailroom := macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v)
	sizeOfRXPageSize := v.SMPCacheBytes() // sizeof(struct efx_rx_page_state) (aligned to cache line)
	overhead := maxFrameLen(0) + sizeOfRXPageSize + rxPrefixSize + rxBufferPadding +
		rxIPAlign + tailroom + headroom

	return v.PageSize - overhead
}

func calcNetsecMTU(v vars) int {
	return ETH_DATA_LEN
}

func calcStmmacMTU(v vars) int {
	return ETH_DATA_LEN
}

func calcCPSWMTU(v vars) int {
	return math.MaxInt
}

func calcNetvscMTU(v vars) int {
	const NETVSC_XDP_HDRM = 256
	return v.PageSize - (ETH_HLEN + NETVSC_XDP_HDRM + macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v))
}

func calcVmxnet3MTU(v vars) int {
	VMXNET3_XDP_HEADROOM := v.XDPPacketheadroom + v.NetIPAlign()
	VMXNET3_XDP_RX_TAILROOM := macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v)
	VMXNET3_XDP_MAX_FRSIZE := v.PageSize - VMXNET3_XDP_HEADROOM - VMXNET3_XDP_RX_TAILROOM
	return VMXNET3_XDP_MAX_FRSIZE - ETH_HLEN - 2*VLAN_HLEN - ETH_FCS_LEN
}

func calcFBNICMTUDefault(v vars) int {
	FBNIC_RX_PAD := 0
	FBNIC_HDS_THRESH_DEFAULT := (1536 - FBNIC_RX_PAD)
	mtu := FBNIC_HDS_THRESH_DEFAULT - ETH_HLEN
	return mtu
}

func calcFBNICMTUMax(v vars) int {
	FBNIC_RX_PAD := 0
	FBNIC_RX_HROOM_PAD := 128
	FBNIC_RX_TROOM := macroSKBDataAlign(v.SizeOfSKBSharedInfo(), v)
	FBNIC_RX_HROOM := macroAlign(FBNIC_RX_TROOM+FBNIC_RX_HROOM_PAD, 128) - FBNIC_RX_TROOM
	FBNIC_HDS_THRESH_MAX := (4096 - FBNIC_RX_HROOM - FBNIC_RX_TROOM - FBNIC_RX_PAD)
	mtu := FBNIC_HDS_THRESH_MAX - ETH_HLEN
	return mtu
}
