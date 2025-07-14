package efuse

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_EFUSE_ADC_CALIB_VER1 = 1

/**
 * @brief Get the RTC calibration efuse version
 *
 * @return Version of the stored efuse
 */
//go:linkname EspEfuseRtcCalibGetVer C.esp_efuse_rtc_calib_get_ver
func EspEfuseRtcCalibGetVer() c.Int

/**
 * @brief Get the init code in the efuse, for the corresponding attenuation.
 *
 * @param version   Version of the stored efuse
 * @param adc_unit  ADC unit. Not used, for compatibility. On esp32h2, for calibration v1, both ADC units use the same init code (calibrated by ADC1)
 * @param atten     Attenuation of the init code
 * @return The init code stored in efuse
 */
//go:linkname EspEfuseRtcCalibGetInitCode C.esp_efuse_rtc_calib_get_init_code
func EspEfuseRtcCalibGetInitCode(version c.Int, adc_unit c.Uint32T, atten c.Int) c.Uint32T

/**
 * @brief Get the channel specific calibration compensation
 *
 * @param version   Version of the stored efuse
 * @param adc_unit  ADC unit. Not used, for compatibility. ESP32H2 only supports one ADC unit
 * @param atten     Attenuation of the init code
 * @return          The channel calibration compensation value
 */
//go:linkname EspEfuseRtcCalibGetChanCompens C.esp_efuse_rtc_calib_get_chan_compens
func EspEfuseRtcCalibGetChanCompens(version c.Int, adc_unit c.Uint32T, adc_channel c.Uint32T, atten c.Int) c.Int

/**
 * @brief Get the calibration digits stored in the efuse, and the corresponding voltage.
 *
 * @param version Version of the stored efuse
 * @param adc_unit      ADC unit (not used on ESP32H2, for compatibility)
 * @param atten         Attenuation to use
 * @param out_digi      Output buffer of the digits
 * @param out_vol_mv    Output of the voltage, in mV
 * @return
 *      - ESP_ERR_INVALID_ARG: If efuse version or attenuation is invalid
 *      - ESP_OK: if success
 */
//go:linkname EspEfuseRtcCalibGetCalVoltage C.esp_efuse_rtc_calib_get_cal_voltage
func EspEfuseRtcCalibGetCalVoltage(version c.Int, adc_unit c.Uint32T, atten c.Int, out_digi *c.Uint32T, out_vol_mv *c.Uint32T) EspErrT

/**
 * @brief Get the temperature sensor calibration number delta_T stored in the efuse.
 *
 * @param tsens_cal Pointer of the specification of temperature sensor calibration number in efuse.
 *
 * @return ESP_OK if get the calibration value successfully.
 *         ESP_ERR_INVALID_ARG if can't get the calibration value.
 */
//go:linkname EspEfuseRtcCalibGetTsensVal C.esp_efuse_rtc_calib_get_tsens_val
func EspEfuseRtcCalibGetTsensVal(tsens_cal *c.Float) EspErrT
