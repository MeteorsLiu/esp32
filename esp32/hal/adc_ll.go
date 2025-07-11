package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ADC_LL_DEFAULT_CONV_LIMIT_EN = 1
const ADC_LL_DEFAULT_CONV_LIMIT_NUM = 10

type AdcLlControllerT c.Int

const (
	ADC_LL_CTRL_RTC   AdcLlControllerT = 0
	ADC_LL_CTRL_ULP   AdcLlControllerT = 1
	ADC_LL_CTRL_DIG   AdcLlControllerT = 2
	ADC_LL_CTRL_PWDET AdcLlControllerT = 3
)

type AdcLlDigiConvertModeT c.Int

const (
	ADC_LL_DIGI_CONV_ONLY_ADC1  AdcLlDigiConvertModeT = 0
	ADC_LL_DIGI_CONV_ONLY_ADC2  AdcLlDigiConvertModeT = 1
	ADC_LL_DIGI_CONV_BOTH_UNIT  AdcLlDigiConvertModeT = 2
	ADC_LL_DIGI_CONV_ALTER_UNIT AdcLlDigiConvertModeT = 3
)

// Need a unit test for bit_width
type AdcLlDigiPatternTableT struct {
	Unused [8]uint8
}
type AdcLlHallControllerT c.Int

const (
	ADC_HALL_CTRL_ULP AdcLlHallControllerT = 0
	ADC_HALL_CTRL_RTC AdcLlHallControllerT = 1
)
