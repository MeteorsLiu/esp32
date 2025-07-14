package esp_driver_spi

import _ "unsafe"

/**
 * @brief Reset the trans Queue of slave driver
 * @note
 * This API is used to reset SPI Slave transaction queue. After calling this function:
 * - The SPI Slave transaction queue will be reset.
 *
 * @note This API shouldn't be called when the corresponding SPI Master is doing an SPI transaction.
 * If this gets called when its corresponding SPI Master is doing an SPI transaction, the SPI Slave behaviour is undefined
 *
 * @param host SPI peripheral that is acting as a slave
 *
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_OK                on success
 */
//go:linkname SpiSlaveQueueReset C.spi_slave_queue_reset
func SpiSlaveQueueReset(host SpiHostDeviceT) EspErrT

/**
 * @brief Reset the trans Queue from within ISR of slave driver
 * @note
 * This API is used to reset SPI Slave transaction queue from within ISR. After calling this function:
 * - The SPI Slave transaction queue will be empty.
 *
 * @param host SPI peripheral that is acting as a slave
 *
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_OK                on success
 */
//go:linkname SpiSlaveQueueResetIsr C.spi_slave_queue_reset_isr
func SpiSlaveQueueResetIsr(host SpiHostDeviceT) EspErrT

/**
 * @brief Queue a SPI transaction in ISR
 * @note
 * Similar as ``spi_slave_queue_trans``, but can and can only called within an ISR, then get the transaction results
 * through the transaction descriptor passed in ``spi_slave_interface_config_t::post_trans_cb``. if use this API, you
 * should trigger a transaction by normal ``spi_slave_queue_trans`` once and only once to start isr
 *
 * If you use both ``spi_slave_queue_trans`` and ``spi_slave_queue_trans_isr`` simultaneously to transfer valid data,
 * you should deal with concurrency issues on your self risk
 *
 * @param host SPI peripheral that is acting as a slave
 * @param trans_desc Description of transaction to execute. Not const because we may want to write status back
 *                   into the transaction description.
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_ERR_NO_MEM        if trans_queue is full
 *         - ESP_OK                on success
 */
//go:linkname SpiSlaveQueueTransIsr C.spi_slave_queue_trans_isr
func SpiSlaveQueueTransIsr(host SpiHostDeviceT, trans_desc *SpiSlaveTransactionT) EspErrT
