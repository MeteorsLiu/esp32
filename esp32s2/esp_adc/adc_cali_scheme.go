package esp_adc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcCaliLineFittingEfuseValT c.Int

const (
	ADC_CALI_LINE_FITTING_EFUSE_VAL_EFUSE_VREF   AdcCaliLineFittingEfuseValT = 0
	ADC_CALI_LINE_FITTING_EFUSE_VAL_EFUSE_TP     AdcCaliLineFittingEfuseValT = 1
	ADC_CALI_LINE_FITTING_EFUSE_VAL_DEFAULT_VREF AdcCaliLineFittingEfuseValT = 2
)

type AdcCaliLineFittingConfigT struct {
	UnitId   AdcUnitT
	Atten    AdcAttenT
	Bitwidth AdcBitwidthT
}

/**
 * @brief Create a Line Fitting calibration scheme
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
// llgo:link (*AdcCaliLineFittingConfigT).AdcCaliCreateSchemeLineFitting C.adc_cali_create_scheme_line_fitting
func (recv_ *AdcCaliLineFittingConfigT) AdcCaliCreateSchemeLineFitting(ret_handle *AdcCaliHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete the Line Fitting calibration scheme handle
 *
 * @param[in] handle ADC calibration handle
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 */
//go:linkname AdcCaliDeleteSchemeLineFitting C.adc_cali_delete_scheme_line_fitting
func AdcCaliDeleteSchemeLineFitting(handle AdcCaliHandleT) EspErrT
