package esp_driver_sdspi

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SDSPI_IO_ACTIVE_LOW = 0

type SdspiDevHandleT c.Int

/**
 * Extra configuration for SD SPI device.
 */

type SdspiDeviceConfigT struct {
	HostId         SpiHostDeviceT
	GpioCs         GpioNumT
	GpioCd         GpioNumT
	GpioWp         GpioNumT
	GpioInt        GpioNumT
	GpioWpPolarity bool
	DutyCyclePos   c.Uint16T
}

/**
 * @brief Initialize SD SPI driver
 *
 * @note This function is not thread safe
 *
 * @return
 *      - ESP_OK on success
 *      - other error codes may be returned in future versions
 */
//go:linkname SdspiHostInit C.sdspi_host_init
func SdspiHostInit() EspErrT

/**
* @brief Attach and initialize an SD SPI device on the specific SPI bus
*
* @note This function is not thread safe
*
* @note Initialize the SPI bus by `spi_bus_initialize()` before calling this function.
*
* @note The SDIO over sdspi needs an extra interrupt line. Call ``gpio_install_isr_service()`` before this function.
*
* @param dev_config pointer to device configuration structure
* @param out_handle Output of the handle to the sdspi device.

* @return
*      - ESP_OK on success
*      - ESP_ERR_INVALID_ARG if sdspi_host_init_device has invalid arguments
*      - ESP_ERR_NO_MEM if memory can not be allocated
*      - other errors from the underlying spi_master and gpio drivers
 */
// llgo:link (*SdspiDeviceConfigT).SdspiHostInitDevice C.sdspi_host_init_device
func (recv_ *SdspiDeviceConfigT) SdspiHostInitDevice(out_handle *SdspiDevHandleT) EspErrT {
	return 0
}

/**
 * @brief Remove an SD SPI device
 *
 * @param handle Handle of the SD SPI device
 * @return Always ESP_OK
 */
// llgo:link SdspiDevHandleT.SdspiHostRemoveDevice C.sdspi_host_remove_device
func (recv_ SdspiDevHandleT) SdspiHostRemoveDevice() EspErrT {
	return 0
}

/**
 * @brief Send command to the card and get response
 *
 * This function returns when command is sent and response is received,
 * or data is transferred, or timeout occurs.
 *
 * @note This function is not thread safe w.r.t. init/deinit functions,
 *       and bus width/clock speed configuration functions. Multiple tasks
 *       can call sdspi_host_do_transaction as long as other sdspi_host_*
 *       functions are not called.
 *
 * @param handle    Handle of the sdspi device
 * @param cmdinfo   pointer to structure describing command and data to transfer
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_TIMEOUT if response or data transfer has timed out
 *      - ESP_ERR_INVALID_CRC if response or data transfer CRC check has failed
 *      - ESP_ERR_INVALID_RESPONSE if the card has sent an invalid response
 */
// llgo:link SdspiDevHandleT.SdspiHostDoTransaction C.sdspi_host_do_transaction
func (recv_ SdspiDevHandleT) SdspiHostDoTransaction(cmdinfo *SdmmcCommandT) EspErrT {
	return 0
}

/**
 * @brief Set card clock frequency
 *
 * Currently only integer fractions of 40MHz clock can be used.
 * For High Speed cards, 40MHz can be used.
 * For Default Speed cards, 20MHz can be used.
 *
 * @note This function is not thread safe
 *
 * @param host    Handle of the sdspi device
 * @param freq_khz  card clock frequency, in kHz
 * @return
 *      - ESP_OK on success
 *      - other error codes may be returned in the future
 */
// llgo:link SdspiDevHandleT.SdspiHostSetCardClk C.sdspi_host_set_card_clk
func (recv_ SdspiDevHandleT) SdspiHostSetCardClk(freq_khz c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Calculate working frequency for specific device
 *
 * @param handle SDSPI device handle
 * @param[out] real_freq_khz output parameter to hold the calculated frequency (in kHz)
 *
 * @return
 *      - ESP_ERR_INVALID_ARG : ``handle`` is NULL or invalid or ``real_freq_khz`` parameter is NULL
 *      - ESP_OK : Success
 */
// llgo:link SdspiDevHandleT.SdspiHostGetRealFreq C.sdspi_host_get_real_freq
func (recv_ SdspiDevHandleT) SdspiHostGetRealFreq(real_freq_khz *c.Int) EspErrT {
	return 0
}

/**
 * @brief Release resources allocated using sdspi_host_init
 *
 * @note This function is not thread safe
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_STATE if sdspi_host_init function has not been called
 */
//go:linkname SdspiHostDeinit C.sdspi_host_deinit
func SdspiHostDeinit() EspErrT

/**
 * @brief Enable SDIO interrupt.
 *
 * @param handle    Handle of the sdspi device
 *
 * @return
 *      - ESP_OK on success
 */
// llgo:link SdspiDevHandleT.SdspiHostIoIntEnable C.sdspi_host_io_int_enable
func (recv_ SdspiDevHandleT) SdspiHostIoIntEnable() EspErrT {
	return 0
}

/**
 * @brief Wait for SDIO interrupt until timeout.
 *
 * @param handle    Handle of the sdspi device
 * @param timeout_ticks Ticks to wait before timeout.
 *
 * @return
 *      - ESP_OK on success
 */
// llgo:link SdspiDevHandleT.SdspiHostIoIntWait C.sdspi_host_io_int_wait
func (recv_ SdspiDevHandleT) SdspiHostIoIntWait(timeout_ticks TickTypeT) EspErrT {
	return 0
}

/**
 * @brief Get the DMA memory information for the host driver
 *
 * @param[in]  slot          Not used
 * @param[out] dma_mem_info  DMA memory information structure
 * @return
 *        - ESP_OK:                ON success.
 *        - ESP_ERR_INVALID_ARG:   Invalid argument.
 */
//go:linkname SdspiHostGetDmaInfo C.sdspi_host_get_dma_info
func SdspiHostGetDmaInfo(slot c.Int, dma_mem_info *EspDmaMemInfoT) EspErrT
