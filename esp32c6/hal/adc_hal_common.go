package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcHalWorkModeT c.Int

const (
	ADC_HAL_SINGLE_READ_MODE     AdcHalWorkModeT = 0
	ADC_HAL_CONTINUOUS_READ_MODE AdcHalWorkModeT = 1
	ADC_HAL_PWDET_MODE           AdcHalWorkModeT = 2
	ADC_HAL_LP_MODE              AdcHalWorkModeT = 3
)

/**
 * Set ADC work mode
 *
 * @param unit       ADC unit
 * @param work_mode  see `adc_hal_work_mode_t`
 */
// llgo:link AdcUnitT.AdcHalSetController C.adc_hal_set_controller
func (recv_ AdcUnitT) AdcHalSetController(work_mode AdcHalWorkModeT) {
}

/**
 * @brief Initialize default parameter for the calibration block.
 *
 * @param adc_n ADC index number
 */
// llgo:link AdcUnitT.AdcHalCalibrationInit C.adc_hal_calibration_init
func (recv_ AdcUnitT) AdcHalCalibrationInit() {
}

/**
 * Set the calibration result (initial data) to ADC.
 *
 * @note  Different ADC units and different attenuation options use different calibration data (initial data).
 *
 * @param adc_n ADC index number.
 * @param param the calibration parameter to configure
 */
// llgo:link AdcUnitT.AdcHalSetCalibrationParam C.adc_hal_set_calibration_param
func (recv_ AdcUnitT) AdcHalSetCalibrationParam(param c.Uint32T) {
}

/**
 * Calibrate the ADC using internal connections.
 *
 * @note  Different ADC units and different attenuation options use different calibration data (initial data).
 *
 * @param adc_n ADC index number.
 * @param atten ADC attenuation
 * @param internal_gnd true:  Disconnect from the IO port and use the internal GND as the calibration voltage.
 *                     false: Use IO external voltage as calibration voltage.
 *
 * @return
 *      - The calibration result (initial data) to ADC, use `adc_hal_set_calibration_param` to set.
 */
// llgo:link AdcUnitT.AdcHalSelfCalibration C.adc_hal_self_calibration
func (recv_ AdcUnitT) AdcHalSelfCalibration(atten AdcAttenT, internal_gnd bool) c.Uint32T {
	return 0
}
