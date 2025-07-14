package esp_driver_spi

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SpiSlaveTransactionT struct {
	Unused [8]uint8
}

// llgo:type C
type SlaveTransactionCbT func(*SpiSlaveTransactionT)

/**
 * @brief This is a configuration for a SPI host acting as a slave device.
 */

type SpiSlaveInterfaceConfigT struct {
	SpicsIoNum  c.Int
	Flags       c.Uint32T
	QueueSize   c.Int
	Mode        c.Uint8T
	PostSetupCb SlaveTransactionCbT
	PostTransCb SlaveTransactionCbT
}

/**
 * @brief Initialize a SPI bus as a slave interface
 *
 * @warning SPI0/1 is not supported
 *
 * @param host          SPI peripheral to use as a SPI slave interface
 * @param bus_config    Pointer to a spi_bus_config_t struct specifying how the host should be initialized
 * @param slave_config  Pointer to a spi_slave_interface_config_t struct specifying the details for the slave interface
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
//go:linkname SpiSlaveInitialize C.spi_slave_initialize
func SpiSlaveInitialize(host SpiHostDeviceT, bus_config *SpiBusConfigT, slave_config *SpiSlaveInterfaceConfigT, dma_chan SpiDmaChanT) EspErrT

/**
 * @brief Free a SPI bus claimed as a SPI slave interface
 *
 * @param host SPI peripheral to free
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_ERR_INVALID_STATE if not all devices on the bus are freed
 *         - ESP_OK                on success
 */
//go:linkname SpiSlaveFree C.spi_slave_free
func SpiSlaveFree(host SpiHostDeviceT) EspErrT

/**
 * @brief Queue a SPI transaction for execution
 *
 * @note On esp32, if trans length not WORD aligned, the rx buffer last word memory will still overwritten by DMA HW
 *
 * Queues a SPI transaction to be executed by this slave device. (The transaction queue size was specified when the slave
 * device was initialised via spi_slave_initialize.) This function may block if the queue is full (depending on the
 * ticks_to_wait parameter). No SPI operation is directly initiated by this function, the next queued transaction
 * will happen when the master initiates a SPI transaction by pulling down CS and sending out clock signals.
 *
 * This function hands over ownership of the buffers in ``trans_desc`` to the SPI slave driver; the application is
 * not to access this memory until ``spi_slave_queue_trans`` is called to hand ownership back to the application.
 *
 * @param host SPI peripheral that is acting as a slave
 * @param trans_desc Description of transaction to execute. Not const because we may want to write status back
 *                   into the transaction description.
 * @param ticks_to_wait Ticks to wait until there's room in the queue; use portMAX_DELAY to
 *                      never time out.
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_ERR_NO_MEM        if set flag `SPI_SLAVE_TRANS_DMA_BUFFER_ALIGN_AUTO` but there is no free memory
 *         - ESP_ERR_INVALID_STATE if sync data between Cache and memory failed
 *         - ESP_OK                on success
 */
//go:linkname SpiSlaveQueueTrans C.spi_slave_queue_trans
func SpiSlaveQueueTrans(host SpiHostDeviceT, trans_desc *SpiSlaveTransactionT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Get the result of a SPI transaction queued earlier
 *
 * This routine will wait until a transaction to the given device (queued earlier with
 * spi_slave_queue_trans) has successfully completed. It will then return the description of the
 * completed transaction so software can inspect the result and e.g. free the memory or
 * reuse the buffers.
 *
 * It is mandatory to eventually use this function for any transaction queued by ``spi_slave_queue_trans``.
 *
 * @param host SPI peripheral to that is acting as a slave
 * @param[out] trans_desc Pointer to variable able to contain a pointer to the description of the
 *                   transaction that is executed
 * @param ticks_to_wait Ticks to wait until there's a returned item; use portMAX_DELAY to never time
 *                      out.
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_ERR_NOT_SUPPORTED if flag `SPI_SLAVE_NO_RETURN_RESULT` is set
 *         - ESP_OK                on success
 */
//go:linkname SpiSlaveGetTransResult C.spi_slave_get_trans_result
func SpiSlaveGetTransResult(host SpiHostDeviceT, trans_desc **SpiSlaveTransactionT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Do a SPI transaction
 *
 * Essentially does the same as spi_slave_queue_trans followed by spi_slave_get_trans_result. Do
 * not use this when there is still a transaction queued that hasn't been finalized
 * using spi_slave_get_trans_result.
 *
 * @param host SPI peripheral to that is acting as a slave
 * @param trans_desc Pointer to variable able to contain a pointer to the description of the
 *                   transaction that is executed. Not const because we may want to write status back
 *                   into the transaction description.
 * @param ticks_to_wait Ticks to wait until there's a returned item; use portMAX_DELAY to never time
 *                      out.
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_OK                on success
 */
//go:linkname SpiSlaveTransmit C.spi_slave_transmit
func SpiSlaveTransmit(host SpiHostDeviceT, trans_desc *SpiSlaveTransactionT, ticks_to_wait TickTypeT) EspErrT
