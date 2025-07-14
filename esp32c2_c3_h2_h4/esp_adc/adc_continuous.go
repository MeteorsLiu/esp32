package esp_adc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcContinuousCtxT struct {
	Unused [8]uint8
}
type AdcContinuousHandleT *AdcContinuousCtxT

/**
 * @brief ADC continuous mode driver initial configurations
 */

type AdcContinuousHandleCfgT struct {
	MaxStoreBufSize c.Uint32T
	ConvFrameSize   c.Uint32T
	Flags           struct {
		FlushPool c.Uint32T
	}
}

/**
 * @brief ADC continuous mode driver configurations
 */

type AdcContinuousConfigT struct {
	PatternNum   c.Uint32T
	AdcPattern   *AdcDigiPatternConfigT
	SampleFreqHz c.Uint32T
	ConvMode     AdcDigiConvertModeT
	Format       AdcDigiOutputFormatT
}

/**
 * @brief Event data structure
 * @note The `conv_frame_buffer` is maintained by the driver itself, so never free this piece of memory.
 */

type AdcContinuousEvtDataT struct {
	ConvFrameBuffer *c.Uint8T
	Size            c.Uint32T
}

// llgo:type C
type AdcContinuousCallbackT func(AdcContinuousHandleT, *AdcContinuousEvtDataT, c.Pointer) bool

/**
 * @brief Group of ADC continuous mode callbacks
 *
 * @note These callbacks are all running in an ISR environment.
 * @note When CONFIG_ADC_CONTINUOUS_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables should be in internal RAM as well.
 */

type AdcContinuousEvtCbsT struct {
	OnConvDone AdcContinuousCallbackT
	OnPoolOvf  AdcContinuousCallbackT
}

/**
 * @brief Initialize ADC continuous driver and get a handle to it
 *
 * @param[in]  hdl_config  Pointer to ADC initialization config. Refer to ``adc_continuous_handle_cfg_t``.
 * @param[out] ret_handle  ADC continuous mode driver handle
 *
 * @return
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_NOT_FOUND     No free interrupt found with the specified flags
 *         - ESP_ERR_NO_MEM        If out of memory
 *         - ESP_OK                On success
 */
// llgo:link (*AdcContinuousHandleCfgT).AdcContinuousNewHandle C.adc_continuous_new_handle
func (recv_ *AdcContinuousHandleCfgT) AdcContinuousNewHandle(ret_handle *AdcContinuousHandleT) EspErrT {
	return 0
}

/**
 * @brief Set ADC continuous mode required configurations
 *
 * @param[in] handle ADC continuous mode driver handle
 * @param[in] config Refer to ``adc_digi_config_t``.
 *
 * @return
 *      - ESP_ERR_INVALID_STATE: Driver state is invalid, you shouldn't call this API at this moment
 *      - ESP_ERR_INVALID_ARG:   If the combination of arguments is invalid.
 *      - ESP_OK:                On success
 */
//go:linkname AdcContinuousConfig C.adc_continuous_config
func AdcContinuousConfig(handle AdcContinuousHandleT, config *AdcContinuousConfigT) EspErrT

/**
 * @brief Register callbacks
 *
 * @note User can deregister a previously registered callback by calling this function and setting the to-be-deregistered callback member in
 *       the `cbs` structure to NULL.
 * @note When CONFIG_ADC_CONTINUOUS_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables (including `user_data`) should be in internal RAM as well.
 * @note You should only call this API when the ADC continuous mode driver isn't started. Check return value to know this.
 *
 * @param[in] handle    ADC continuous mode driver handle
 * @param[in] cbs       Group of callback functions
 * @param[in] user_data User data, which will be delivered to the callback functions directly
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_INVALID_STATE: Driver state is invalid, you shouldn't call this API at this moment
 */
//go:linkname AdcContinuousRegisterEventCallbacks C.adc_continuous_register_event_callbacks
func AdcContinuousRegisterEventCallbacks(handle AdcContinuousHandleT, cbs *AdcContinuousEvtCbsT, user_data c.Pointer) EspErrT

/**
 * @brief Start the ADC under continuous mode. After this, the hardware starts working.
 *
 * @param[in]  handle              ADC continuous mode driver handle
 *
 * @return
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 *         - ESP_OK                On success
 */
//go:linkname AdcContinuousStart C.adc_continuous_start
func AdcContinuousStart(handle AdcContinuousHandleT) EspErrT

/**
 * @brief Read bytes from ADC under continuous mode.
 *
 * @param[in]  handle              ADC continuous mode driver handle
 * @param[out] buf                 Conversion result buffer to read from ADC. Suggest convert to `adc_digi_output_data_t` for `ADC Conversion Results`.
 *                                 See the subsection `Driver Backgrounds` in this header file to learn about this concept.
 * @param[in]  length_max          Expected length of the Conversion Results read from the ADC, in bytes.
 * @param[out] out_length          Real length of the Conversion Results read from the ADC via this API, in bytes.
 * @param[in]  timeout_ms          Time to wait for data via this API, in millisecond.
 *
 * @return
 *         - ESP_ERR_INVALID_STATE Driver state is invalid. Usually it means the ADC sampling rate is faster than the task processing rate.
 *         - ESP_ERR_TIMEOUT       Operation timed out
 *         - ESP_OK                On success
 */
//go:linkname AdcContinuousRead C.adc_continuous_read
func AdcContinuousRead(handle AdcContinuousHandleT, buf *c.Uint8T, length_max c.Uint32T, out_length *c.Uint32T, timeout_ms c.Uint32T) EspErrT

/**
 * @brief Stop the ADC. After this, the hardware stops working.
 *
 * @param[in]  handle              ADC continuous mode driver handle
 *
 * @return
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 *         - ESP_OK                On success
 */
//go:linkname AdcContinuousStop C.adc_continuous_stop
func AdcContinuousStop(handle AdcContinuousHandleT) EspErrT

/**
 * @brief Deinitialize the ADC continuous driver.
 *
 * @param[in]  handle              ADC continuous mode driver handle
 *
 * @return
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 *         - ESP_OK                On success
 */
//go:linkname AdcContinuousDeinit C.adc_continuous_deinit
func AdcContinuousDeinit(handle AdcContinuousHandleT) EspErrT

/**
 * @brief Flush the driver internal pool
 *
 * @note This API is not supposed to be called in an ISR context
 *
 * @param[in]  handle              ADC continuous mode driver handle
 *
 * @return
 *         - ESP_ERR_INVALID_STATE Driver state is invalid, you should call this API when it's in init state
 *         - ESP_ERR_INVALID_ARG:  Invalid arguments
 *         - ESP_OK                On success
 */
//go:linkname AdcContinuousFlushPool C.adc_continuous_flush_pool
func AdcContinuousFlushPool(handle AdcContinuousHandleT) EspErrT

/**
 * @brief Get ADC channel from the given GPIO number
 *
 * @param[in]  io_num     GPIO number
 * @param[out] unit_id    ADC unit
 * @param[out] channel    ADC channel
 *
 * @return
 *        - ESP_OK:              On success
 *        - ESP_ERR_INVALID_ARG: Invalid argument
 *        - ESP_ERR_NOT_FOUND:   The IO is not a valid ADC pad
 */
//go:linkname AdcContinuousIoToChannel C.adc_continuous_io_to_channel
func AdcContinuousIoToChannel(io_num c.Int, unit_id *AdcUnitT, channel *AdcChannelT) EspErrT

/**
 * @brief Get GPIO number from the given ADC channel
 *
 * @param[in]  unit_id    ADC unit
 * @param[in]  channel    ADC channel
 * @param[out] io_num     GPIO number
 *
 * @param
 *       - ESP_OK:              On success
 *       - ESP_ERR_INVALID_ARG: Invalid argument
 */
//go:linkname AdcContinuousChannelToIo C.adc_continuous_channel_to_io
func AdcContinuousChannelToIo(unit_id AdcUnitT, channel AdcChannelT, io_num *c.Int) EspErrT
