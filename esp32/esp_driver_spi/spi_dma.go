package esp_driver_spi

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SpiDmaChanDirT c.Int

const (
	DMA_CHANNEL_DIRECTION_TX SpiDmaChanDirT = 0
	DMA_CHANNEL_DIRECTION_RX SpiDmaChanDirT = 1
)

type SpiDmaChanHandleT struct {
	HostId SpiHostDeviceT
	Dir    SpiDmaChanDirT
	ChanId c.Int
}

/**
 * Enable/Disable data/desc burst for spi_dma channel
 *
 * @param chan_handle   Context of the spi_dma channel.
 * @param data_burst    enable or disable data burst
 * @param desc_burst    enable or disable desc burst
 */
// llgo:link SpiDmaChanHandleT.SpiDmaEnableBurst C.spi_dma_enable_burst
func (recv_ SpiDmaChanHandleT) SpiDmaEnableBurst(data_burst bool, desc_burst bool) {
}

/**
 * Reset dma channel for spi_dma
 *
 * @param chan_handle   Context of the spi_dma channel.
 */
// llgo:link SpiDmaChanHandleT.SpiDmaReset C.spi_dma_reset
func (recv_ SpiDmaChanHandleT) SpiDmaReset() {
}

/**
 * Start dma channel for spi_dma
 *
 * @param chan_handle   Context of the spi_dma channel.
 * @param addr          Addr of linked dma descriptor to mount
 */
// llgo:link SpiDmaChanHandleT.SpiDmaStart C.spi_dma_start
func (recv_ SpiDmaChanHandleT) SpiDmaStart(addr c.Pointer) {
}
