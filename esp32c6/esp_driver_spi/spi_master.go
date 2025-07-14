package esp_driver_spi

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SpiTransactionT struct {
	Unused [8]uint8
}

// llgo:type C
type TransactionCbT func(*SpiTransactionT)

/**
 * @brief This is a configuration for a SPI slave device that is connected to one of the SPI buses.
 */

type SpiDeviceInterfaceConfigT struct {
	CommandBits    c.Uint8T
	AddressBits    c.Uint8T
	DummyBits      c.Uint8T
	Mode           c.Uint8T
	ClockSource    SpiClockSourceT
	DutyCyclePos   c.Uint16T
	CsEnaPretrans  c.Uint16T
	CsEnaPosttrans c.Uint8T
	ClockSpeedHz   c.Int
	InputDelayNs   c.Int
	SamplePoint    SpiSamplingPointT
	SpicsIoNum     c.Int
	Flags          c.Uint32T
	QueueSize      c.Int
	PreCb          TransactionCbT
	PostCb         TransactionCbT
}

/**
 * This struct is for SPI transactions which may change their address and command length.
 * Please do set the flags in base to ``SPI_TRANS_VARIABLE_CMD_ADR`` to use the bit length here.
 */

type SpiTransactionExtT struct {
	Base        SpiTransactionT
	CommandBits c.Uint8T
	AddressBits c.Uint8T
	DummyBits   c.Uint8T
}

type SpiDeviceT struct {
	Unused [8]uint8
}
type SpiDeviceHandleT *SpiDeviceT

/**
 * @brief Allocate a device on a SPI bus
 *
 * This initializes the internal structures for a device, plus allocates a CS pin on the indicated SPI master
 * peripheral and routes it to the indicated GPIO. All SPI master devices have three CS pins and can thus control
 * up to three devices.
 *
 * @note On ESP32, due to the delay of GPIO matrix, the maximum frequency SPI Master can correctly samples the slave's
 *       output is lower than the case using IOMUX. Typical maximum frequency communicating with an ideal slave
 *       without data output delay: 80MHz (IOMUX pins) and 26MHz (GPIO matrix pins). With the help of extra dummy
 *       cycles in half-duplex mode, the delay can be compensated by setting `input_delay_ns` in `dev_config` structure
 *       correctly.
 *
 *       There's no notable delay on chips other than ESP32.
 *
 * @param host_id SPI peripheral to allocate device on
 * @param dev_config SPI interface protocol config for the device
 * @param handle Pointer to variable to hold the device handle
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid or configuration combination is not supported (e.g.
 *                                 `dev_config->post_cb` isn't set while flag `SPI_DEVICE_NO_RETURN_RESULT` is enabled)
 *         - ESP_ERR_INVALID_STATE if selected clock source is unavailable or spi bus not initialized
 *         - ESP_ERR_NOT_FOUND     if host doesn't have any free CS slots
 *         - ESP_ERR_NO_MEM        if out of memory
 *         - ESP_OK                on success
 */
//go:linkname SpiBusAddDevice C.spi_bus_add_device
func SpiBusAddDevice(host_id SpiHostDeviceT, dev_config *SpiDeviceInterfaceConfigT, handle *SpiDeviceHandleT) EspErrT

/**
 * @brief Remove a device from the SPI bus
 *
 * @param handle Device handle to free
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_ERR_INVALID_STATE if device already is freed
 *         - ESP_OK                on success
 */
//go:linkname SpiBusRemoveDevice C.spi_bus_remove_device
func SpiBusRemoveDevice(handle SpiDeviceHandleT) EspErrT

/**
 * @brief Queue a SPI transaction for interrupt transaction execution. Get the result by ``spi_device_get_trans_result``.
 *
 * @note Normally a device cannot start (queue) polling and interrupt
 *      transactions simultaneously.
 *
 * @param handle Device handle obtained using spi_host_add_dev
 * @param trans_desc Description of transaction to execute
 * @param ticks_to_wait Ticks to wait until there's room in the queue; use portMAX_DELAY to
 *                      never time out.
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid. This can happen if SPI_TRANS_CS_KEEP_ACTIVE flag is specified while
 *                                 the bus was not acquired (`spi_device_acquire_bus()` should be called first)
 *                                 or set flag SPI_TRANS_DMA_BUFFER_ALIGN_MANUAL but tx or rx buffer not DMA-capable, or addr&len not align to cache line size
 *         - ESP_ERR_TIMEOUT       if there was no room in the queue before ticks_to_wait expired
 *         - ESP_ERR_NO_MEM        if allocating DMA-capable temporary buffer failed
 *         - ESP_ERR_INVALID_STATE if previous transactions are not finished
 *         - ESP_OK                on success
 */
//go:linkname SpiDeviceQueueTrans C.spi_device_queue_trans
func SpiDeviceQueueTrans(handle SpiDeviceHandleT, trans_desc *SpiTransactionT, ticks_to_wait TickTypeT) EspErrT

/**
* @brief Get the result of a SPI transaction queued earlier by ``spi_device_queue_trans``.
*
* This routine will wait until a transaction to the given device
* successfully completed. It will then return the description of the
* completed transaction so software can inspect the result and e.g. free the memory or
* reuse the buffers.
*
* @param handle Device handle obtained using spi_host_add_dev
* @param trans_desc Pointer to variable able to contain a pointer to the description of the transaction
       that is executed. The descriptor should not be modified until the descriptor is returned by
       spi_device_get_trans_result.
* @param ticks_to_wait Ticks to wait until there's a returned item; use portMAX_DELAY to never time
                       out.
* @return
*         - ESP_ERR_INVALID_ARG   if parameter is invalid
*         - ESP_ERR_NOT_SUPPORTED if flag `SPI_DEVICE_NO_RETURN_RESULT` is set
*         - ESP_ERR_TIMEOUT       if there was no completed transaction before ticks_to_wait expired
*         - ESP_OK                on success
*/
//go:linkname SpiDeviceGetTransResult C.spi_device_get_trans_result
func SpiDeviceGetTransResult(handle SpiDeviceHandleT, trans_desc **SpiTransactionT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Send a SPI transaction, wait for it to complete, and return the result
 *
 * This function is the equivalent of calling spi_device_queue_trans() followed by spi_device_get_trans_result().
 * Do not use this when there is still a transaction separately queued (started) from spi_device_queue_trans() or polling_start/transmit that hasn't been finalized.
 *
 * @note This function is not thread safe when multiple tasks access the same SPI device.
 *      Normally a device cannot start (queue) polling and interrupt
 *      transactions simutanuously.
 *
 * @param handle Device handle obtained using spi_host_add_dev
 * @param trans_desc Description of transaction to execute
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_OK                on success
 */
//go:linkname SpiDeviceTransmit C.spi_device_transmit
func SpiDeviceTransmit(handle SpiDeviceHandleT, trans_desc *SpiTransactionT) EspErrT

/**
 * @brief Immediately start a polling transaction.
 *
 * @note Normally a device cannot start (queue) polling and interrupt
 *      transactions simutanuously. Moreover, a device cannot start a new polling
 *      transaction if another polling transaction is not finished.
 *
 * @param handle Device handle obtained using spi_host_add_dev
 * @param trans_desc Description of transaction to execute
 * @param ticks_to_wait Ticks to wait until there's room in the queue;
 *              currently only portMAX_DELAY is supported.
 *
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid. This can happen if SPI_TRANS_CS_KEEP_ACTIVE flag is specified while
 *                                 the bus was not acquired (`spi_device_acquire_bus()` should be called first)
 *                                 or set flag SPI_TRANS_DMA_BUFFER_ALIGN_MANUAL but tx or rx buffer not DMA-capable, or addr&len not align to cache line size
 *         - ESP_ERR_TIMEOUT       if the device cannot get control of the bus before ``ticks_to_wait`` expired
 *         - ESP_ERR_NO_MEM        if allocating DMA-capable temporary buffer failed
 *         - ESP_ERR_INVALID_STATE if previous transactions are not finished
 *         - ESP_OK                on success
 */
//go:linkname SpiDevicePollingStart C.spi_device_polling_start
func SpiDevicePollingStart(handle SpiDeviceHandleT, trans_desc *SpiTransactionT, ticks_to_wait TickTypeT) EspErrT

/**
* @brief Poll until the polling transaction ends.
*
* This routine will not return until the transaction to the given device has
* successfully completed. The task is not blocked, but actively busy-spins for
* the transaction to be completed.
*
* @param handle Device handle obtained using spi_host_add_dev
* @param ticks_to_wait Ticks to wait until there's a returned item; use portMAX_DELAY to never time
                       out.
* @return
*         - ESP_ERR_INVALID_ARG   if parameter is invalid
*         - ESP_ERR_TIMEOUT       if the transaction cannot finish before ticks_to_wait expired
*         - ESP_OK                on success
*/
//go:linkname SpiDevicePollingEnd C.spi_device_polling_end
func SpiDevicePollingEnd(handle SpiDeviceHandleT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Send a polling transaction, wait for it to complete, and return the result
 *
 * This function is the equivalent of calling spi_device_polling_start() followed by spi_device_polling_end().
 * Do not use this when there is still a transaction that hasn't been finalized.
 *
 * @note This function is not thread safe when multiple tasks access the same SPI device.
 *      Normally a device cannot start (queue) polling and interrupt
 *      transactions simutanuously.
 *
 * @param handle Device handle obtained using spi_host_add_dev
 * @param trans_desc Description of transaction to execute
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_ERR_TIMEOUT       if the device cannot get control of the bus
 *         - ESP_ERR_NO_MEM        if allocating DMA-capable temporary buffer failed
 *         - ESP_ERR_INVALID_STATE if previous transactions of same device are not finished
 *         - ESP_OK                on success
 */
//go:linkname SpiDevicePollingTransmit C.spi_device_polling_transmit
func SpiDevicePollingTransmit(handle SpiDeviceHandleT, trans_desc *SpiTransactionT) EspErrT

/**
 * @brief Occupy the SPI bus for a device to do continuous transactions.
 *
 * Transactions to all other devices will be put off until ``spi_device_release_bus`` is called.
 *
 * @note The function will wait until all the existing transactions have been sent.
 *
 * @param device The device to occupy the bus.
 * @param wait Time to wait before the the bus is occupied by the device. Currently MUST set to portMAX_DELAY.
 *
 * @return
 *      - ESP_ERR_INVALID_ARG : ``wait`` is not set to portMAX_DELAY.
 *      - ESP_OK : Success.
 */
//go:linkname SpiDeviceAcquireBus C.spi_device_acquire_bus
func SpiDeviceAcquireBus(device SpiDeviceHandleT, wait TickTypeT) EspErrT

/**
 * @brief Release the SPI bus occupied by the device. All other devices can start sending transactions.
 *
 * @param dev The device to release the bus.
 */
//go:linkname SpiDeviceReleaseBus C.spi_device_release_bus
func SpiDeviceReleaseBus(dev SpiDeviceHandleT)

/**
 * @brief Calculate working frequency for specific device
 *
 * @param handle SPI device handle
 * @param[out] freq_khz output parameter to hold calculated frequency in kHz
 *
 * @return
 *      - ESP_ERR_INVALID_ARG : ``handle`` or ``freq_khz`` parameter is NULL
 *      - ESP_OK : Success
 */
//go:linkname SpiDeviceGetActualFreq C.spi_device_get_actual_freq
func SpiDeviceGetActualFreq(handle SpiDeviceHandleT, freq_khz *c.Int) EspErrT

/**
 * @brief Calculate the working frequency that is most close to desired frequency.
 *
 * @param fapb The frequency of apb clock, should be ``APB_CLK_FREQ``.
 * @param hz Desired working frequency
 * @param duty_cycle Duty cycle of the spi clock
 *
 * @return Actual working frequency that most fit.
 */
//go:linkname SpiGetActualClock C.spi_get_actual_clock
func SpiGetActualClock(fapb c.Int, hz c.Int, duty_cycle c.Int) c.Int

/**
 * @brief Calculate the timing settings of specified frequency and settings.
 *
 * @param gpio_is_used True if using GPIO matrix, or False if iomux pins are used.
 * @param input_delay_ns Input delay from SCLK launch edge to MISO data valid.
 * @param eff_clk Effective clock frequency (in Hz) from `spi_get_actual_clock()`.
 * @param dummy_o Address of dummy bits used output. Set to NULL if not needed.
 * @param cycles_remain_o Address of cycles remaining (after dummy bits are used) output.
 *         - -1 If too many cycles remaining, suggest to compensate half a clock.
 *         - 0 If no remaining cycles or dummy bits are not used.
 *         - positive value: cycles suggest to compensate.
 *
 * @note If **dummy_o* is not zero, it means dummy bits should be applied in half duplex mode, and full duplex mode may not work.
 */
//go:linkname SpiGetTiming C.spi_get_timing
func SpiGetTiming(gpio_is_used bool, input_delay_ns c.Int, eff_clk c.Int, dummy_o *c.Int, cycles_remain_o *c.Int)

/**
 * @brief Get the frequency limit of current configurations.
 *         SPI master working at this limit is OK, while above the limit, full duplex mode and DMA will not work,
 *         and dummy bits will be applied in the half duplex mode.
 *
 * @param gpio_is_used True if using GPIO matrix, or False if native pins are used.
 * @param input_delay_ns Input delay from SCLK launch edge to MISO data valid.
 * @return Frequency limit of current configurations.
 */
//go:linkname SpiGetFreqLimit C.spi_get_freq_limit
func SpiGetFreqLimit(gpio_is_used bool, input_delay_ns c.Int) c.Int

/**
 * @brief Get max length (in bytes) of one transaction
 *
 * @param       host_id    SPI peripheral
 * @param[out]  max_bytes  Max length of one transaction, in bytes
 *
 * @return
 *        - ESP_OK:               On success
 *        - ESP_ERR_INVALID_ARG:  Invalid argument
 */
//go:linkname SpiBusGetMaxTransactionLen C.spi_bus_get_max_transaction_len
func SpiBusGetMaxTransactionLen(host_id SpiHostDeviceT, max_bytes *c.SizeT) EspErrT
