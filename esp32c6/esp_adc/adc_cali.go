package esp_adc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcCaliHandleT *AdcCaliSchemeT
type AdcCaliSchemeVerT c.Int

const (
	ADC_CALI_SCHEME_VER_LINE_FITTING  AdcCaliSchemeVerT = 1
	ADC_CALI_SCHEME_VER_CURVE_FITTING AdcCaliSchemeVerT = 2
)

/**
 * @brief Check the supported ADC calibration scheme
 *
 * @param[out] scheme_mask    Supported ADC calibration scheme(s)
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_NOT_SUPPORTED: No supported calibration scheme
 */
// llgo:link (*AdcCaliSchemeVerT).AdcCaliCheckScheme C.adc_cali_check_scheme
func (recv_ *AdcCaliSchemeVerT) AdcCaliCheckScheme() EspErrT {
	return 0
}

/**
 * @brief Convert ADC raw data to calibrated voltage
 *
 * @param[in]  handle     ADC calibration handle
 * @param[in]  raw        ADC raw data
 * @param[out] voltage    Calibrated ADC voltage (in mV)
 *
 * @return
 *         - ESP_OK:                On success
 *         - ESP_ERR_INVALID_ARG:   Invalid argument
 *         - ESP_ERR_INVALID_STATE: Invalid state, scheme didn't registered
 */
//go:linkname AdcCaliRawToVoltage C.adc_cali_raw_to_voltage
func AdcCaliRawToVoltage(handle AdcCaliHandleT, raw c.Int, voltage *c.Int) EspErrT
