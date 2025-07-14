package esp_adc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspAdcCalValueT c.Int

const (
	ESP_ADC_CAL_VAL_EFUSE_VREF    EspAdcCalValueT = 0
	ESP_ADC_CAL_VAL_EFUSE_TP      EspAdcCalValueT = 1
	ESP_ADC_CAL_VAL_DEFAULT_VREF  EspAdcCalValueT = 2
	ESP_ADC_CAL_VAL_EFUSE_TP_FIT  EspAdcCalValueT = 3
	ESP_ADC_CAL_VAL_MAX           EspAdcCalValueT = 4
	ESP_ADC_CAL_VAL_NOT_SUPPORTED EspAdcCalValueT = 4
)

/**
 * @brief Structure storing characteristics of an ADC
 *
 * @note Call esp_adc_cal_characterize() to initialize the structure
 */

type EspAdcCalCharacteristicsT struct {
	AdcNum    AdcUnitT
	Atten     AdcAttenT
	BitWidth  AdcBitsWidthT
	CoeffA    c.Uint32T
	CoeffB    c.Uint32T
	Vref      c.Uint32T
	LowCurve  *c.Uint32T
	HighCurve *c.Uint32T
	Version   c.Uint8T
}
