package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Type of dma descriptor with appended members
 *        this structure inherits DMA descriptor, with a pointer to the transaction descriptor passed from users.
 */

type SpiSlaveHdHalDescAppendT struct {
	Desc *SpiDmaDescT
	Arg  c.Pointer
}

// / Configuration of the HAL
type SpiSlaveHdHalConfigT struct {
	HostId      c.Uint32T
	DmaEnabled  bool
	AppendMode  bool
	SpicsIoNum  c.Uint32T
	Mode        c.Uint8T
	CommandBits c.Uint32T
	AddressBits c.Uint32T
	DummyBits   c.Uint32T
}

// / Context of the HAL, initialized by :cpp:func:`spi_slave_hd_hal_init`.
type SpiSlaveHdHalContextT struct {
	DmadescTx         *SpiSlaveHdHalDescAppendT
	DmadescRx         *SpiSlaveHdHalDescAppendT
	Dev               *SpiDevT
	DmaEnabled        bool
	AppendMode        bool
	DmaDescNum        c.Uint32T
	CurrentEofAddr    c.Uint32T
	TxCurDesc         *SpiSlaveHdHalDescAppendT
	TxDmaHead         *SpiSlaveHdHalDescAppendT
	TxDmaTail         *SpiSlaveHdHalDescAppendT
	TxUsedDescCnt     c.Uint32T
	TxRecycledDescCnt c.Uint32T
	RxCurDesc         *SpiSlaveHdHalDescAppendT
	RxDmaHead         *SpiSlaveHdHalDescAppendT
	RxDmaTail         *SpiSlaveHdHalDescAppendT
	RxUsedDescCnt     c.Uint32T
	RxRecycledDescCnt c.Uint32T
	IntrNotTriggered  c.Uint32T
	TxDmaStarted      bool
	RxDmaStarted      bool
}
