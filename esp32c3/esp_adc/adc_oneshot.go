package esp_adc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief ADC oneshot driver initial configurations
 */

type AdcOneshotUnitInitCfgT struct {
	UnitId  AdcUnitT
	ClkSrc  AdcOneshotClkSrcT
	UlpMode AdcUlpModeT
}

/**
 * @brief ADC channel configurations
 */

type AdcOneshotChanCfgT struct {
	Atten    AdcAttenT
	Bitwidth AdcBitwidthT
}

/**
 * @brief Create a handle to a specific ADC unit
 *
 * @note This API is thread-safe. For more details, see ADC programming guide
 *
 * @param[in]  init_config    Driver initial configurations
 * @param[out] ret_unit       ADC unit handle
 *
 * @return
 *        - ESP_OK:              On success
 *        - ESP_ERR_INVALID_ARG: Invalid arguments
 *        - ESP_ERR_NO_MEM:      No memory
 *        - ESP_ERR_NOT_FOUND:   The ADC peripheral to be claimed is already in use
 *        - ESP_FAIL:            Clock source isn't initialised correctly
 */
// llgo:link (*AdcOneshotUnitInitCfgT).AdcOneshotNewUnit C.adc_oneshot_new_unit
func (recv_ *AdcOneshotUnitInitCfgT) AdcOneshotNewUnit(ret_unit *AdcOneshotUnitHandleT) EspErrT {
	return 0
}

/**
 * @brief Set ADC oneshot mode required configurations
 *
 * @note This API is thread-safe. For more details, see ADC programming guide
 *
 * @param[in] handle    ADC handle
 * @param[in] channel   ADC channel to be configured
 * @param[in] config    ADC configurations
 *
 * @return
 *        - ESP_OK:              On success
 *        - ESP_ERR_INVALID_ARG: Invalid arguments
 */
//go:linkname AdcOneshotConfigChannel C.adc_oneshot_config_channel
func AdcOneshotConfigChannel(handle AdcOneshotUnitHandleT, channel AdcChannelT, config *AdcOneshotChanCfgT) EspErrT

/**
 * @brief Get one ADC conversion raw result
 *
 * @note This API is thread-safe. For more details, see ADC programming guide
 * @note This API should NOT be called in an ISR context
 *
 * @param[in] handle    ADC handle
 * @param[in] chan      ADC channel
 * @param[out] out_raw  ADC conversion raw result
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_TIMEOUT:       Timeout, the ADC result is invalid
 */
//go:linkname AdcOneshotRead C.adc_oneshot_read
func AdcOneshotRead(handle AdcOneshotUnitHandleT, chan_ AdcChannelT, out_raw *c.Int) EspErrT

/**
 * @brief Delete the ADC unit handle
 *
 * @note This API is thread-safe. For more details, see ADC programming guide
 *
 * @param[in] handle    ADC handle
 *
 * @return
 *        - ESP_OK:              On success
 *        - ESP_ERR_INVALID_ARG: Invalid arguments
 *        - ESP_ERR_NOT_FOUND:   The ADC peripheral to be disclaimed isn't in use
 */
//go:linkname AdcOneshotDelUnit C.adc_oneshot_del_unit
func AdcOneshotDelUnit(handle AdcOneshotUnitHandleT) EspErrT

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
//go:linkname AdcOneshotIoToChannel C.adc_oneshot_io_to_channel
func AdcOneshotIoToChannel(io_num c.Int, unit_id *AdcUnitT, channel *AdcChannelT) EspErrT

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
//go:linkname AdcOneshotChannelToIo C.adc_oneshot_channel_to_io
func AdcOneshotChannelToIo(unit_id AdcUnitT, channel AdcChannelT, io_num *c.Int) EspErrT

/**
 * @brief Convenience function to get ADC calibrated result
 *
 * This is an all-in-one function which does:
 * - oneshot read ADC raw result
 * - calibrate the raw result and convert it into calibrated result (in mV)
 *
 * @param[in]  handle       ADC oneshot handle, you should call adc_oneshot_new_unit() to get this handle
 * @param[in]  cali_handle  ADC calibration handle, you should call adc_cali_create_scheme_x() in adc_cali_scheme.h to create a handle
 * @param[in]  chan         ADC channel
 * @param[out] cali_result  Calibrated ADC result (in mV)
 *
 * @return
 *        - ESP_OK
 *        Other return errors from adc_oneshot_read() and adc_cali_raw_to_voltage()
 */
//go:linkname AdcOneshotGetCalibratedResult C.adc_oneshot_get_calibrated_result
func AdcOneshotGetCalibratedResult(handle AdcOneshotUnitHandleT, cali_handle AdcCaliHandleT, chan_ AdcChannelT, cali_result *c.Int) EspErrT
