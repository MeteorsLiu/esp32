package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcI2sSourceT c.Int

const (
	ADC_I2S_DATA_SRC_IO_SIG AdcI2sSourceT = 0
	ADC_I2S_DATA_SRC_ADC    AdcI2sSourceT = 1
)

/*---------------------------------------------------------------
            ESP32 Deprecated API
---------------------------------------------------------------*/
/**
 * @brief Set I2S data source
 * @param src I2S DMA data source, I2S DMA can get data from digital signals or from ADC.
 * @return
 *     - ESP_OK success
 */
// llgo:link AdcI2sSourceT.AdcSetI2sDataSource C.adc_set_i2s_data_source
func (recv_ AdcI2sSourceT) AdcSetI2sDataSource() EspErrT {
	return 0
}

/**
 * @brief Initialize I2S ADC mode
 * @param adc_unit ADC unit index
 * @param channel ADC channel index
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname AdcI2sModeInit C.adc_i2s_mode_init
func AdcI2sModeInit(adc_unit AdcUnitT, channel AdcChannelT) EspErrT
