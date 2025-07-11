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
