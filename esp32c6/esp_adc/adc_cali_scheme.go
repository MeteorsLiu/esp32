package esp_adc

import _ "unsafe"

/*
---------------------------------------------------------------

	Curve Fitting Calibration Scheme

---------------------------------------------------------------
*/
type AdcCaliCurveFittingConfigT struct {
	UnitId   AdcUnitT
	Chan     AdcChannelT
	Atten    AdcAttenT
	Bitwidth AdcBitwidthT
}

/**
 * @brief Create a Curve Fitting calibration scheme
 *
 * After creating, you'll get a handle to this scheme. Then you can use the driver APIS in `esp_adc/adc_cali.h` to do the
 * ADC calibration via the handle you get.
 *
 * @param[in]  config  Initial configurations
 * @param[out] handle  ADC calibration handle
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_NO_MEM:        No enough memory
 *        - ESP_ERR_NOT_SUPPORTED: Scheme required eFuse bits not burnt
 */
// llgo:link (*AdcCaliCurveFittingConfigT).AdcCaliCreateSchemeCurveFitting C.adc_cali_create_scheme_curve_fitting
func (recv_ *AdcCaliCurveFittingConfigT) AdcCaliCreateSchemeCurveFitting(ret_handle *AdcCaliHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete the Curve Fitting calibration scheme handle
 *
 * @param[in] handle ADC calibration handle
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 */
//go:linkname AdcCaliDeleteSchemeCurveFitting C.adc_cali_delete_scheme_curve_fitting
func AdcCaliDeleteSchemeCurveFitting(handle AdcCaliHandleT) EspErrT
