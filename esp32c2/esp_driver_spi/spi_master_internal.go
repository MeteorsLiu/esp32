package esp_driver_spi

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * This struct is for SPI SCT (Segmented-Configure-Transfer) Mode.
 *
 * By default, length of each SPI Phase will not change per segment. Each segment will use the phase length you set when `spi_bus_add_device()`
 * However, you could force a segment to use its custom phase length. To achieve this, set the `SPI_SEG_TRANS_XX` flags, to customize phase length.
 */

type SpiMultiTransactionT struct {
	Base           SpiTransactionT
	CsEnaPretrans  c.Uint8T
	CsEnaPosttrans c.Uint8T
	CommandBits    c.Uint8T
	AddressBits    c.Uint8T
	DummyBits      c.Uint8T
	SctGapLen      c.Uint32T
	SegTransFlags  c.Uint32T
}

/**
 * @brief Enable/Disable Segmented-Configure-Transfer (SCT) mode
 *
 * Search for `@Backgrounds: `SCT Mode`` in this header file to know what is SCT mode
 *
 * @note This API isn't thread safe. Besides, after enabling this, current SPI host will be switched into SCT mode.
 *       Therefore, never call this API when in multiple threads, or when an SPI transaction is ongoing (on this SPI host).
 *
 * @param handle Device handle obtained using spi_host_add_dev
 * @param enable True: to enable SCT mode; False: to disable SCT mode
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_INVALID_STATE: Invalid states, e.g.: an SPI polling transaction is ongoing, SPI internal Queue isn't empty, etc.
 */
//go:linkname SpiBusMultiTransModeEnable C.spi_bus_multi_trans_mode_enable
func SpiBusMultiTransModeEnable(handle SpiDeviceHandleT, enable bool) EspErrT

/**
 * @brief Queue an SPI Segmented-Configure-Transaction (SCT) list for interrupt transaction execution.
 *
 * Search for `@Backgrounds: `SCT Mode`` in this header file to know what is SCT mode
 *
 * @note After calling this API, call `spi_device_get_multi_trans_result` to get the transaction results.
 *
 * @param handle         Device handle obtained using spi_host_add_dev
 * @param seg_trans_desc Pointer to the transaction segments list head (a one-segment-list is also acceptable)
 * @param trans_num      Transaction number in this segment
 * @param ticks_to_wait  Ticks to wait until there's room in the queue; use portMAX_DELAY to never time out.
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_INVALID_STATE: Invalid states, e.g.: an SPI polling transaction is ongoing, SCT mode isn't enabled, DMA descriptors not enough, etc.
 *        - ESP_ERR_TIMEOUT:       Timeout, this SCT transaction isn't queued successfully
 */
//go:linkname SpiDeviceQueueMultiTrans C.spi_device_queue_multi_trans
func SpiDeviceQueueMultiTrans(handle SpiDeviceHandleT, seg_trans_desc *SpiMultiTransactionT, trans_num c.Uint32T, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Get the result of an SPI Segmented-Configure-Transaction (SCT).
 *
 * Search for `@Backgrounds: `SCT Mode`` in this header file to know what is SCT mode
 *
 * @note Until this API returns (with `ESP_OK`), you can now recycle the memory used for this SCT list (pointed by `seg_trans_desc`).
 *       You must maintain the SCT list related memory before this API returns, otherwise the SCT transaction may fail
 *
 * @param handle              Device handle obtained using spi_host_add_dev
 * @param[out] seg_trans_desc Pointer to the completed SCT list head (then you can recycle this list of memory).
 * @param ticks_to_wait       Ticks to wait until there's a returned item; use portMAX_DELAY to never time out.
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_INVALID_STATE: Invalid states, e.g.: SCT mode isn't enabled, etc.
 *        - ESP_ERR_TIMEOUT:       Timeout, didn't get a completed SCT transaction
 */
//go:linkname SpiDeviceGetMultiTransResult C.spi_device_get_multi_trans_result
func SpiDeviceGetMultiTransResult(handle SpiDeviceHandleT, seg_trans_desc **SpiMultiTransactionT, ticks_to_wait TickTypeT) EspErrT
