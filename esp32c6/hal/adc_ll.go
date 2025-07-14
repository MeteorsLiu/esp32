package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ADC_LL_CLKM_DIV_NUM_DEFAULT = 15
const ADC_LL_CLKM_DIV_B_DEFAULT = 1
const ADC_LL_CLKM_DIV_A_DEFAULT = 0
const ADC_LL_DEFAULT_CONV_LIMIT_EN = 0
const ADC_LL_DEFAULT_CONV_LIMIT_NUM = 10
const ADC_LL_POWER_MANAGE_SUPPORTED = 1

type AdcLlPowerT c.Int

const (
	ADC_LL_POWER_BY_FSM AdcLlPowerT = 0
	ADC_LL_POWER_SW_ON  AdcLlPowerT = 1
	ADC_LL_POWER_SW_OFF AdcLlPowerT = 2
)

type AdcLlControllerT c.Int

const ADC_LL_CTRL_DIG AdcLlControllerT = 0

type AdcLlDigiConvertModeT c.Int

const ADC_LL_DIGI_CONV_ONLY_ADC1 AdcLlDigiConvertModeT = 0

type AdcLlDigiPatternTableT struct {
	Unused [8]uint8
}
