package esp_adc

import _ "unsafe"

type AdcIirFilterT struct {
	Unused [8]uint8
}
type AdcIirFilterHandleT *AdcIirFilterT

/**
 * @brief Filter Configuration
 */

type AdcContinuousIirFilterConfigT struct {
	Unit    AdcUnitT
	Channel AdcChannelT
	Coeff   AdcDigiIirFilterCoeffT
}

/**
 * @brief New an ADC continuous mode IIR filter
 *
 * @param[in]   handle   ADC continuous mode driver handle
 * @param[in]   config   Filter configuration
 * @param[out]  ret_hdl  Returned IIR filter handle
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:    Invalid argument
 *        - ESP_ERR_INVALID_STATE:  Invalid state
 *        - ESP_ERR_NO_MEM:         No memory
 */
//go:linkname AdcNewContinuousIirFilter C.adc_new_continuous_iir_filter
func AdcNewContinuousIirFilter(handle AdcContinuousHandleT, config *AdcContinuousIirFilterConfigT, ret_hdl *AdcIirFilterHandleT) EspErrT

/**
 * @brief Enable an IIR filter
 *
 * @param[in] filter_hdl  ADC IIR filter handle
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:    Invalid argument
 *        - ESP_ERR_INVALID_STATE:  Invalid state
 */
//go:linkname AdcContinuousIirFilterEnable C.adc_continuous_iir_filter_enable
func AdcContinuousIirFilterEnable(filter_hdl AdcIirFilterHandleT) EspErrT

/**
 * @brief Disable an IIR filter
 *
 * @param[in] filter_hdl  ADC IIR filter handle
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:    Invalid argument
 *        - ESP_ERR_INVALID_STATE:  Invalid state
 */
//go:linkname AdcContinuousIirFilterDisable C.adc_continuous_iir_filter_disable
func AdcContinuousIirFilterDisable(filter_hdl AdcIirFilterHandleT) EspErrT

/**
 * @brief Delete the IIR filter
 *
 * @param[in] filter_hdl  ADC IIR filter handle
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:    Invalid argument
 *        - ESP_ERR_INVALID_STATE:  Invalid state
 */
//go:linkname AdcDelContinuousIirFilter C.adc_del_continuous_iir_filter
func AdcDelContinuousIirFilter(filter_hdl AdcIirFilterHandleT) EspErrT
