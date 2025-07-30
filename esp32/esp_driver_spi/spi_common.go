package esp_driver_spi

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SPICOMMON_BUSFLAG_SLAVE = 0

type SpiCommonDmaT c.Int

const (
	SPI_DMA_DISABLED SpiCommonDmaT = 0
	SPI_DMA_CH1      SpiCommonDmaT = 1
	SPI_DMA_CH2      SpiCommonDmaT = 2
	SPI_DMA_CH_AUTO  SpiCommonDmaT = 3
)

type SpiDmaChanT SpiCommonDmaT

/**
 * @brief This is a configuration structure for a SPI bus.
 *
 * You can use this structure to specify the GPIO pins of the bus. Normally, the driver will use the
 * GPIO matrix to route the signals. An exception is made when all signals either can be routed through
 * the IO_MUX or are -1. In that case, the IO_MUX is used. On ESP32, using GPIO matrix will bring about 25ns of input
 * delay, which may cause incorrect read for >40MHz speeds.
 *
 * @note Be advised that the slave driver does not use the quadwp/quadhd lines and fields in spi_bus_config_t referring to these lines will be ignored and can thus safely be left uninitialized.
 */

type SpiBusConfigT struct {
	SclkIoNum          c.Int
	Data4IoNum         c.Int
	Data5IoNum         c.Int
	Data6IoNum         c.Int
	Data7IoNum         c.Int
	DataIoDefaultLevel bool
	MaxTransferSz      c.Int
	Flags              c.Uint32T
	IsrCpuId           EspIntrCpuAffinityT
	IntrFlags          c.Int
}

/**
 * @brief Initialize a SPI bus
 *
 * @warning SPI0/1 is not supported
 *
 * @param host_id       SPI peripheral that controls this bus
 * @param bus_config    Pointer to a spi_bus_config_t struct specifying how the host should be initialized
 * @param dma_chan      - Selecting a DMA channel for an SPI bus allows transactions on the bus with size only limited by the amount of internal memory.
 *                      - Selecting SPI_DMA_DISABLED limits the size of transactions.
 *                      - Set to SPI_DMA_DISABLED if only the SPI flash uses this bus.
 *                      - Set to SPI_DMA_CH_AUTO to let the driver to allocate the DMA channel.
 *
 * @warning If a DMA channel is selected, any transmit and receive buffer used should be allocated in
 *          DMA-capable memory.
 *
 * @warning The ISR of SPI is always executed on the core which calls this
 *          function. Never starve the ISR on this core or the SPI transactions will not
 *          be handled.
 *
 * @return
 *         - ESP_ERR_INVALID_ARG   if configuration is invalid
 *         - ESP_ERR_INVALID_STATE if host already is in use
 *         - ESP_ERR_NOT_FOUND     if there is no available DMA channel
 *         - ESP_ERR_NO_MEM        if out of memory
 *         - ESP_OK                on success
 */
//go:linkname SpiBusInitialize C.spi_bus_initialize
func SpiBusInitialize(host_id SpiHostDeviceT, bus_config *SpiBusConfigT, dma_chan SpiDmaChanT) EspErrT

/**
 * @brief Free a SPI bus
 *
 * @warning In order for this to succeed, all devices have to be removed first.
 *
 * @param host_id SPI peripheral to free
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_ERR_INVALID_STATE if bus hasn't been initialized before, or not all devices on the bus are freed
 *         - ESP_OK                on success
 */
//go:linkname SpiBusFree C.spi_bus_free
func SpiBusFree(host_id SpiHostDeviceT) EspErrT

/**
 * @brief Helper function for malloc DMA capable memory for SPI driver
 *
 * @note This API will take care of the cache and hardware alignment internally.
 *       To free/release memory allocated by this helper function, simply calling `free()`
 *
 * @param[in]  host_id          SPI peripheral who will using the memory
 * @param[in]  size             Size in bytes, the amount of memory to allocate
 * @param[in]  extra_heap_caps  Extra heap caps based on MALLOC_CAP_DMA
 *
 * @return                      Pointer to the memory if allocated successfully
 */
//go:linkname SpiBusDmaMemoryAlloc C.spi_bus_dma_memory_alloc
func SpiBusDmaMemoryAlloc(host_id SpiHostDeviceT, size c.SizeT, extra_heap_caps c.Uint32T) c.Pointer
