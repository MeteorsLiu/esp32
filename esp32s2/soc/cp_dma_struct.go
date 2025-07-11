package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Interrupt Registers */
/** Type of dma_int_raw register
 *  Raw interrupt status
 */

type CpDmaIntRawRegT struct {
	Val c.Uint32T
}

/** Type of dma_int_st register
 *  Masked interrupt status
 */

type CpDmaIntStRegT struct {
	Val c.Uint32T
}

/** Type of dma_int_ena register
 *  Interrupt enable bits
 */

type CpDmaIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of dma_int_clr register
 *  Interrupt clear bits
 */

type CpDmaIntClrRegT struct {
	Val c.Uint32T
}

/** Configuration Registers */
/** Type of dma_out_link register
 *  Link descriptor address and control
 */

type CpDmaOutLinkRegT struct {
	Val c.Uint32T
}

/** Type of dma_in_link register
 *  Link descriptor address and control
 */

type CpDmaInLinkRegT struct {
	Val c.Uint32T
}

/** Type of dma_conf register
 *  Copy DMA configuration register
 */

type CpDmaConfRegT struct {
	Val c.Uint32T
}

/** Status Registers */
/** Type of dma_out_eof_des_addr register
 *  Outlink descriptor address when EOF occurs
 */

type CpDmaOutEofDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of dma_in_eof_des_addr register
 *  Inlink descriptor address when EOF occurs
 */

type CpDmaInEofDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of dma_out_eof_bfr_des_addr register
 *  Outlink descriptor address before the last outlink descriptor
 */

type CpDmaOutEofBfrDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of dma_inlink_dscr register
 *  Address of current inlink descriptor
 */

type CpDmaInlinkDscrRegT struct {
	Val c.Uint32T
}

/** Type of dma_inlink_dscr_bf0 register
 *  Address of last inlink descriptor
 */

type CpDmaInlinkDscrBf0RegT struct {
	Val c.Uint32T
}

/** Type of dma_inlink_dscr_bf1 register
 *  Address of the second-to-last inlink descriptor
 */

type CpDmaInlinkDscrBf1RegT struct {
	Val c.Uint32T
}

/** Type of dma_outlink_dscr register
 *  Address of current outlink descriptor
 */

type CpDmaOutlinkDscrRegT struct {
	Val c.Uint32T
}

/** Type of dma_outlink_dscr_bf0 register
 *  Address of last outlink descriptor
 */

type CpDmaOutlinkDscrBf0RegT struct {
	Val c.Uint32T
}

/** Type of dma_outlink_dscr_bf1 register
 *  Address of the second-to-last outlink descriptor
 */

type CpDmaOutlinkDscrBf1RegT struct {
	Val c.Uint32T
}

/** Type of dma_in_st register
 *  Status register of receiving data
 */

type CpDmaInStRegT struct {
	Val c.Uint32T
}

/** Type of dma_out_st register
 *  Status register of transmitting data
 */

type CpDmaOutStRegT struct {
	Val c.Uint32T
}

/** Type of dma_crc_out register
 *  CRC result register
 */

type CpDmaCrcOutRegT struct {
	Val c.Uint32T
}

/** Type of dma_date register
 *  Copy DMA version register
 */

type CpDmaDateRegT struct {
	Val c.Uint32T
}

type CpDmaDevT struct {
	DmaIntRaw           CpDmaIntRawRegT
	DmaIntSt            CpDmaIntStRegT
	DmaIntEna           CpDmaIntEnaRegT
	DmaIntClr           CpDmaIntClrRegT
	DmaOutLink          CpDmaOutLinkRegT
	DmaInLink           CpDmaInLinkRegT
	DmaOutEofDesAddr    CpDmaOutEofDesAddrRegT
	DmaInEofDesAddr     CpDmaInEofDesAddrRegT
	DmaOutEofBfrDesAddr CpDmaOutEofBfrDesAddrRegT
	DmaInlinkDscr       CpDmaInlinkDscrRegT
	DmaInlinkDscrBf0    CpDmaInlinkDscrBf0RegT
	DmaInlinkDscrBf1    CpDmaInlinkDscrBf1RegT
	DmaOutlinkDscr      CpDmaOutlinkDscrRegT
	DmaOutlinkDscrBf0   CpDmaOutlinkDscrBf0RegT
	DmaOutlinkDscrBf1   CpDmaOutlinkDscrBf1RegT
	DmaConf             CpDmaConfRegT
	DmaInSt             CpDmaInStRegT
	DmaOutSt            CpDmaOutStRegT
	DmaCrcOut           CpDmaCrcOutRegT
	Reserved04c         c.Uint32T
	Reserved050         c.Uint32T
	Reserved054         c.Uint32T
	Reserved058         c.Uint32T
	Reserved05c         c.Uint32T
	Reserved060         c.Uint32T
	Reserved064         c.Uint32T
	Reserved068         c.Uint32T
	Reserved06c         c.Uint32T
	Reserved070         c.Uint32T
	Reserved074         c.Uint32T
	Reserved078         c.Uint32T
	Reserved07c         c.Uint32T
	Reserved080         c.Uint32T
	Reserved084         c.Uint32T
	Reserved088         c.Uint32T
	Reserved08c         c.Uint32T
	Reserved090         c.Uint32T
	Reserved094         c.Uint32T
	Reserved098         c.Uint32T
	Reserved09c         c.Uint32T
	Reserved0a0         c.Uint32T
	Reserved0a4         c.Uint32T
	Reserved0a8         c.Uint32T
	Reserved0ac         c.Uint32T
	Reserved0b0         c.Uint32T
	Reserved0b4         c.Uint32T
	Reserved0b8         c.Uint32T
	Reserved0bc         c.Uint32T
	Reserved0c0         c.Uint32T
	Reserved0c4         c.Uint32T
	Reserved0c8         c.Uint32T
	Reserved0cc         c.Uint32T
	Reserved0d0         c.Uint32T
	Reserved0d4         c.Uint32T
	Reserved0d8         c.Uint32T
	Reserved0dc         c.Uint32T
	Reserved0e0         c.Uint32T
	Reserved0e4         c.Uint32T
	Reserved0e8         c.Uint32T
	Reserved0ec         c.Uint32T
	Reserved0f0         c.Uint32T
	Reserved0f4         c.Uint32T
	Reserved0f8         c.Uint32T
	DmaDate             CpDmaDateRegT
}
