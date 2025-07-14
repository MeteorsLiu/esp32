package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Register */
/** Type of regdma_conf register
 *  Peri backup control register
 */

type PauRegdmaConfRegT struct {
	Val c.Uint32T
}

/** Type of regdma_clk_conf register
 *  Clock control register
 */

type PauRegdmaClkConfRegT struct {
	Val c.Uint32T
}

/** Type of regdma_etm_ctrl register
 *  ETM start ctrl reg
 */

type PauRegdmaEtmCtrlRegT struct {
	Val c.Uint32T
}

/** Type of regdma_link_0_addr register
 *  link_0_addr
 */

type PauRegdmaLink0AddrRegT struct {
	Val c.Uint32T
}

/** Type of regdma_link_1_addr register
 *  Link_1_addr
 */

type PauRegdmaLink1AddrRegT struct {
	Val c.Uint32T
}

/** Type of regdma_link_2_addr register
 *  Link_2_addr
 */

type PauRegdmaLink2AddrRegT struct {
	Val c.Uint32T
}

/** Type of regdma_link_3_addr register
 *  Link_3_addr
 */

type PauRegdmaLink3AddrRegT struct {
	Val c.Uint32T
}

/** Type of regdma_link_mac_addr register
 *  Link_mac_addr
 */

type PauRegdmaLinkMacAddrRegT struct {
	Val c.Uint32T
}

/** Type of regdma_current_link_addr register
 *  current link addr
 */

type PauRegdmaCurrentLinkAddrRegT struct {
	Val c.Uint32T
}

/** Type of regdma_backup_addr register
 *  Backup addr
 */

type PauRegdmaBackupAddrRegT struct {
	Val c.Uint32T
}

/** Type of regdma_mem_addr register
 *  mem addr
 */

type PauRegdmaMemAddrRegT struct {
	Val c.Uint32T
}

/** Type of regdma_bkp_conf register
 *  backup config
 */

type PauRegdmaBkpConfRegT struct {
	Val c.Uint32T
}

/** Type of retention_link_base register
 *  retention dma link base
 */

type PauRetentionLinkBaseRegT struct {
	Val c.Uint32T
}

/** Type of retention_cfg register
 *  retention_cfg
 */

type PauRetentionCfgRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Read only register for error and done
 */

type PauIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_raw register
 *  Read only register for error and done
 */

type PauIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Read only register for error and done
 */

type PauIntClrRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Read only register for error and done
 */

type PauIntStRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  Date register.
 */

type PauDateRegT struct {
	Val c.Uint32T
}

type PauDevT struct {
	RegdmaConf            PauRegdmaConfRegT
	RegdmaClkConf         PauRegdmaClkConfRegT
	RegdmaEtmCtrl         PauRegdmaEtmCtrlRegT
	RegdmaLink0Addr       PauRegdmaLink0AddrRegT
	RegdmaLink1Addr       PauRegdmaLink1AddrRegT
	RegdmaLink2Addr       PauRegdmaLink2AddrRegT
	RegdmaLink3Addr       PauRegdmaLink3AddrRegT
	RegdmaLinkMacAddr     PauRegdmaLinkMacAddrRegT
	RegdmaCurrentLinkAddr PauRegdmaCurrentLinkAddrRegT
	RegdmaBackupAddr      PauRegdmaBackupAddrRegT
	RegdmaMemAddr         PauRegdmaMemAddrRegT
	RegdmaBkpConf         PauRegdmaBkpConfRegT
	RetentionLinkBase     PauRetentionLinkBaseRegT
	RetentionCfg          PauRetentionCfgRegT
	IntEna                PauIntEnaRegT
	IntRaw                PauIntRawRegT
	IntClr                PauIntClrRegT
	IntSt                 PauIntStRegT
	Reserved048           [237]c.Uint32T
	Date                  PauDateRegT
}
