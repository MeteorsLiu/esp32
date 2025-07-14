package esp_adc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Checks if ADC calibration values are burned into eFuse
 *
 * This function checks if ADC reference voltage or Two Point values have been
 * burned to the eFuse of the current ESP32
 *
 * @param   value_type  Type of calibration value (ESP_ADC_CAL_VAL_EFUSE_VREF or ESP_ADC_CAL_VAL_EFUSE_TP)
 * @note in ESP32S2, only ESP_ADC_CAL_VAL_EFUSE_TP is supported. Some old ESP32S2s do not support this, either.
 * In which case you have to calibrate it manually, possibly by performing your own two-point calibration on the chip.
 *
 * @return
 *      - ESP_OK: The calibration mode is supported in eFuse
 *      - ESP_ERR_NOT_SUPPORTED: Error, eFuse values are not burned
 *      - ESP_ERR_INVALID_ARG: Error, invalid argument (ESP_ADC_CAL_VAL_DEFAULT_VREF)
 */
// llgo:link EspAdcCalValueT.EspAdcCalCheckEfuse C.esp_adc_cal_check_efuse
func (recv_ EspAdcCalValueT) EspAdcCalCheckEfuse() EspErrT {
	return 0
}

/**
 * @brief Characterize an ADC at a particular attenuation
 *
 * This function will characterize the ADC at a particular attenuation and generate
 * the ADC-Voltage curve in the form of [y = coeff_a * x + coeff_b].
 * Characterization can be based on Two Point values, eFuse Vref, or default Vref
 * and the calibration values will be prioritized in that order.
 *
 * @note
 * For ESP32, Two Point values and eFuse Vref calibration can be enabled/disabled using menuconfig.
 * For ESP32s2, only Two Point values calibration and only ADC_WIDTH_BIT_13 is supported. The parameter default_vref is unused.
 *
 *
 * @param[in]   adc_num         ADC to characterize (ADC_UNIT_1 or ADC_UNIT_2)
 * @param[in]   atten           Attenuation to characterize
 * @param[in]   bit_width       Bit width configuration of ADC
 * @param[in]   default_vref    Default ADC reference voltage in mV (Only in ESP32, used if eFuse values is not available)
 * @param[out]  chars           Pointer to empty structure used to store ADC characteristics
 *
 * @return
 *      - ESP_ADC_CAL_VAL_EFUSE_VREF: eFuse Vref used for characterization
 *      - ESP_ADC_CAL_VAL_EFUSE_TP: Two Point value used for characterization (only in Linear Mode)
 *      - ESP_ADC_CAL_VAL_DEFAULT_VREF: Default Vref used for characterization
 */
//go:linkname EspAdcCalCharacterize C.esp_adc_cal_characterize
func EspAdcCalCharacterize(adc_num AdcUnitT, atten AdcAttenT, bit_width AdcBitsWidthT, default_vref c.Uint32T, chars *EspAdcCalCharacteristicsT) EspAdcCalValueT

/**
 * @brief   Convert an ADC reading to voltage in mV
 *
 * This function converts an ADC reading to a voltage in mV based on the ADC's
 * characteristics.
 *
 * @note    Characteristics structure must be initialized before this function
 *          is called (call esp_adc_cal_characterize())
 *
 * @param[in]   adc_reading     ADC reading
 * @param[in]   chars           Pointer to initialized structure containing ADC characteristics
 *
 * @return      Voltage in mV
 */
//go:linkname EspAdcCalRawToVoltage C.esp_adc_cal_raw_to_voltage
func EspAdcCalRawToVoltage(adc_reading c.Uint32T, chars *EspAdcCalCharacteristicsT) c.Uint32T

/**
 * @brief   Reads an ADC and converts the reading to a voltage in mV
 *
 * This function reads an ADC then converts the raw reading to a voltage in mV
 * based on the characteristics provided. The ADC that is read is also
 * determined by the characteristics.
 *
 * @note    The Characteristics structure must be initialized before this
 *          function is called (call esp_adc_cal_characterize())
 *
 * @param[in]   channel     ADC Channel to read
 * @param[in]   chars       Pointer to initialized ADC characteristics structure
 * @param[out]  voltage     Pointer to store converted voltage
 *
 * @return
 *      - ESP_OK: ADC read and converted to mV
 *      - ESP_ERR_INVALID_ARG: Error due to invalid arguments
 *      - ESP_ERR_INVALID_STATE: Reading result is invalid. Try to read again.
 */
//go:linkname EspAdcCalGetVoltage C.esp_adc_cal_get_voltage
func EspAdcCalGetVoltage(channel AdcChannelT, chars *EspAdcCalCharacteristicsT, voltage *c.Uint32T) EspErrT
