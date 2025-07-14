package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: User-defined control registers */
/** Type of cmd register
 *  Command control register
 */

type SpiCmdRegT struct {
	Val c.Uint32T
}

/** Type of addr register
 *  Address value register
 */

type SpiAddrRegT struct {
	Val c.Uint32T
}

/** Type of user register
 *  SPI USER control register
 */

type SpiUserRegT struct {
	Val c.Uint32T
}

/** Type of user1 register
 *  SPI USER control register 1
 */

type SpiUser1RegT struct {
	Val c.Uint32T
}

/** Type of user2 register
 *  SPI USER control register 2
 */

type SpiUser2RegT struct {
	Val c.Uint32T
}

/** Group: Control and configuration registers */
/** Type of ctrl register
 *  SPI control register
 */

type SpiCtrlRegT struct {
	Val c.Uint32T
}

/** Type of ms_dlen register
 *  SPI data bit length control register
 */

type SpiMsDlenRegT struct {
	Val c.Uint32T
}

/** Type of misc register
 *  SPI misc register
 */

type SpiMiscRegT struct {
	Val c.Uint32T
}

/** Type of dma_conf register
 *  SPI DMA control register
 */

type SpiDmaConfRegT struct {
	Val c.Uint32T
}

/** Type of slave register
 *  SPI slave control register
 */

type SpiSlaveRegT struct {
	Val c.Uint32T
}

/** Type of slave1 register
 *  SPI slave control register 1
 */

type SpiSlave1RegT struct {
	Val c.Uint32T
}

/** Group: Clock control registers */
/** Type of clock register
 *  SPI clock control register
 */

type SpiClockRegT struct {
	Val c.Uint32T
}

/** Type of clk_gate register
 *  SPI module clock and register clock control
 */

type SpiClkGateRegT struct {
	Val c.Uint32T
}

/** Group: Timing registers */
/** Type of din_mode register
 *  SPI input delay mode configuration
 */

type SpiDinModeRegT struct {
	Val c.Uint32T
}

/** Type of din_num register
 *  SPI input delay number configuration
 */

type SpiDinNumRegT struct {
	Val c.Uint32T
}

/** Type of dout_mode register
 *  SPI output delay mode configuration
 */

type SpiDoutModeRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of dma_int_ena register
 *  SPI interrupt enable register
 */

type SpiDmaIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of dma_int_clr register
 *  SPI interrupt clear register
 */

type SpiDmaIntClrRegT struct {
	Val c.Uint32T
}

/** Type of dma_int_raw register
 *  SPI interrupt raw register
 */

type SpiDmaIntRawRegT struct {
	Val c.Uint32T
}

/** Type of dma_int_st register
 *  SPI interrupt status register
 */

type SpiDmaIntStRegT struct {
	Val c.Uint32T
}

/** Type of dma_int_set register
 *  SPI interrupt software set register
 */

type SpiDmaIntSetRegT struct {
	Val c.Uint32T
}

/** Group: CPU-controlled data buffer */
/** Type of w0 register
 *  SPI CPU-controlled buffer0
 */

type SpiWnRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version control
 */

type SpiDateRegT struct {
	Val c.Uint32T
}

type SpiDevT struct {
	Cmd         SpiCmdRegT
	Addr        SpiAddrRegT
	Ctrl        SpiCtrlRegT
	Clock       SpiClockRegT
	User        SpiUserRegT
	User1       SpiUser1RegT
	User2       SpiUser2RegT
	MsDlen      SpiMsDlenRegT
	Misc        SpiMiscRegT
	DinMode     SpiDinModeRegT
	DinNum      SpiDinNumRegT
	DoutMode    SpiDoutModeRegT
	DmaConf     SpiDmaConfRegT
	DmaIntEna   SpiDmaIntEnaRegT
	DmaIntClr   SpiDmaIntClrRegT
	DmaIntRaw   SpiDmaIntRawRegT
	DmaIntSt    SpiDmaIntStRegT
	DmaIntSet   SpiDmaIntSetRegT
	Reserved048 [20]c.Uint32T
	DataBuf     [16]SpiWnRegT
	Reserved0d8 [2]c.Uint32T
	Slave       SpiSlaveRegT
	Slave1      SpiSlave1RegT
	ClkGate     SpiClkGateRegT
	Reserved0ec c.Uint32T
	Date        SpiDateRegT
}
