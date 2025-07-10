package esp_driver_spi

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// / Descriptor of data to send/receive
type SpiSlaveHdDataT struct {
	Data     *c.Uint8T
	Len      c.SizeT
	TransLen c.SizeT
	Flags    c.Uint32T
	Arg      c.Pointer
}

// / Information of SPI Slave HD event
type SpiSlaveHdEventT struct {
	Event SpiEventT
	Trans *SpiSlaveHdDataT
}

// llgo:type C
type SlaveCbT func(c.Pointer, *SpiSlaveHdEventT, *BaseTypeT) bool
type SpiSlaveChanT c.Int

const (
	SPI_SLAVE_CHAN_TX SpiSlaveChanT = 0
	SPI_SLAVE_CHAN_RX SpiSlaveChanT = 1
)

// / Callback configuration structure for SPI Slave HD
type SpiSlaveHdCallbackConfigT struct {
	CbBufferTx     SlaveCbT
	CbBufferRx     SlaveCbT
	CbSendDmaReady SlaveCbT
	CbSent         SlaveCbT
	CbRecvDmaReady SlaveCbT
	CbRecv         SlaveCbT
	CbCmd9         SlaveCbT
	CbCmdA         SlaveCbT
	Arg            c.Pointer
}

// / Configuration structure for the SPI Slave HD driver
type SpiSlaveHdSlotConfigT struct {
	Mode        c.Uint8T
	SpicsIoNum  c.Uint32T
	Flags       c.Uint32T
	CommandBits c.Uint32T
	AddressBits c.Uint32T
	DummyBits   c.Uint32T
	QueueSize   c.Uint32T
	DmaChan     SpiDmaChanT
	CbConfig    SpiSlaveHdCallbackConfigT
}

/**
 * @brief Initialize the SPI Slave HD driver.
 *
 * @param host_id       The host to use
 * @param bus_config    Bus configuration for the bus used
 * @param config        Configuration for the SPI Slave HD driver
 * @return
 *  - ESP_OK:                on success
 *  - ESP_ERR_INVALID_ARG:   invalid argument given
 *  - ESP_ERR_INVALID_STATE: function called in invalid state, may be some resources are already in use
 *  - ESP_ERR_NOT_FOUND      if there is no available DMA channel
 *  - ESP_ERR_NO_MEM:        memory allocation failed
 *  - or other return value from `esp_intr_alloc`
 */
//go:linkname SpiSlaveHdInit C.spi_slave_hd_init
func SpiSlaveHdInit(host_id SpiHostDeviceT, bus_config *SpiBusConfigT, config *SpiSlaveHdSlotConfigT) EspErrT

/**
 * @brief Deinitialize the SPI Slave HD driver
 *
 * @param host_id The host to deinitialize the driver
 * @return
 *  - ESP_OK: on success
 *  - ESP_ERR_INVALID_ARG: if the host_id is not correct
 */
//go:linkname SpiSlaveHdDeinit C.spi_slave_hd_deinit
func SpiSlaveHdDeinit(host_id SpiHostDeviceT) EspErrT

/**
 * @brief Queue transactions (segment mode)
 *
 * @param host_id   Host to queue the transaction
 * @param chan      SPI_SLAVE_CHAN_TX or SPI_SLAVE_CHAN_RX
 * @param trans     Transaction descriptors
 * @param timeout   Timeout before the data is queued
 * @return
 *  - ESP_OK: on success
 *  - ESP_ERR_INVALID_ARG: The input argument is invalid. Can be the following reason:
 *      - The buffer given is not DMA capable
 *      - The length of data is invalid (not larger than 0, or exceed the max transfer length)
 *      - The transaction direction is invalid
 *  - ESP_ERR_TIMEOUT: Cannot queue the data before timeout. Master is still processing previous transaction.
 *  - ESP_ERR_INVALID_STATE: Function called in invalid state. This API should be called under segment mode.
 */
//go:linkname SpiSlaveHdQueueTrans C.spi_slave_hd_queue_trans
func SpiSlaveHdQueueTrans(host_id SpiHostDeviceT, chan_ SpiSlaveChanT, trans *SpiSlaveHdDataT, timeout TickTypeT) EspErrT

/**
 * @brief Get the result of a data transaction (segment mode)
 *
 * @note This API should be called successfully the same times as the ``spi_slave_hd_queue_trans``.
 *
 * @param host_id   Host to queue the transaction
 * @param chan      Channel to get the result, SPI_SLAVE_CHAN_TX or SPI_SLAVE_CHAN_RX
 * @param[out] out_trans Pointer to the transaction descriptor (``spi_slave_hd_data_t``) passed to the driver before. Hardware has finished this transaction. Member ``trans_len`` indicates the actual number of bytes of received data, it's meaningless for TX.
 * @param timeout   Timeout before the result is got
 * @return
 *  - ESP_OK: on success
 *  - ESP_ERR_INVALID_ARG: Function is not valid
 *  - ESP_ERR_TIMEOUT: There's no transaction done before timeout
 *  - ESP_ERR_INVALID_STATE: Function called in invalid state. This API should be called under segment mode.
 */
//go:linkname SpiSlaveHdGetTransRes C.spi_slave_hd_get_trans_res
func SpiSlaveHdGetTransRes(host_id SpiHostDeviceT, chan_ SpiSlaveChanT, out_trans **SpiSlaveHdDataT, timeout TickTypeT) EspErrT

/**
 * @brief Read the shared registers
 *
 * @param host_id   Host to read the shared registers
 * @param addr      Address of register to read, 0 to ``SOC_SPI_MAXIMUM_BUFFER_SIZE-1``
 * @param[out] out_data Output buffer to store the read data
 * @param len       Length to read, not larger than ``SOC_SPI_MAXIMUM_BUFFER_SIZE-addr``
 */
//go:linkname SpiSlaveHdReadBuffer C.spi_slave_hd_read_buffer
func SpiSlaveHdReadBuffer(host_id SpiHostDeviceT, addr c.Int, out_data *c.Uint8T, len c.SizeT)

/**
 * @brief Write the shared registers
 *
 * @param host_id   Host to write the shared registers
 * @param addr      Address of register to write, 0 to ``SOC_SPI_MAXIMUM_BUFFER_SIZE-1``
 * @param data      Buffer holding the data to write
 * @param len       Length to write, ``SOC_SPI_MAXIMUM_BUFFER_SIZE-addr``
 */
//go:linkname SpiSlaveHdWriteBuffer C.spi_slave_hd_write_buffer
func SpiSlaveHdWriteBuffer(host_id SpiHostDeviceT, addr c.Int, data *c.Uint8T, len c.SizeT)

/**
 * @brief Load transactions (append mode)
 *
 * @note In this mode, user transaction descriptors will be appended to the DMA and the DMA will keep processing the data without stopping
 *
 * @param host_id   Host to load transactions
 * @param chan      SPI_SLAVE_CHAN_TX or SPI_SLAVE_CHAN_RX
 * @param trans     Transaction descriptor
 * @param timeout   Timeout before the transaction is loaded
 * @return
 *  - ESP_OK: on success
 *  - ESP_ERR_INVALID_ARG: The input argument is invalid. Can be the following reason:
 *      - The buffer given is not DMA capable
 *      - The length of data is invalid (not larger than 0, or exceed the max transfer length)
 *      - The transaction direction is invalid
 *  - ESP_ERR_TIMEOUT: Master is still processing previous transaction. There is no available transaction for slave to load
 *  - ESP_ERR_INVALID_STATE: Function called in invalid state. This API should be called under append mode.
 */
//go:linkname SpiSlaveHdAppendTrans C.spi_slave_hd_append_trans
func SpiSlaveHdAppendTrans(host_id SpiHostDeviceT, chan_ SpiSlaveChanT, trans *SpiSlaveHdDataT, timeout TickTypeT) EspErrT

/**
 * @brief Get the result of a data transaction (append mode)
 *
 * @note This API should be called the same times as the ``spi_slave_hd_append_trans``
 *
 * @param host_id   Host to load the transaction
 * @param chan      SPI_SLAVE_CHAN_TX or SPI_SLAVE_CHAN_RX
 * @param[out] out_trans Pointer to the transaction descriptor (``spi_slave_hd_data_t``) passed to the driver before. Hardware has finished this transaction. Member ``trans_len`` indicates the actual number of bytes of received data, it's meaningless for TX.
 * @param timeout   Timeout before the result is got
 * @return
 *  - ESP_OK: on success
 *  - ESP_ERR_INVALID_ARG: Function is not valid
 *  - ESP_ERR_TIMEOUT: There's no transaction done before timeout
 *  - ESP_ERR_INVALID_STATE: Function called in invalid state. This API should be called under append mode.
 */
//go:linkname SpiSlaveHdGetAppendTransRes C.spi_slave_hd_get_append_trans_res
func SpiSlaveHdGetAppendTransRes(host_id SpiHostDeviceT, chan_ SpiSlaveChanT, out_trans **SpiSlaveHdDataT, timeout TickTypeT) EspErrT
