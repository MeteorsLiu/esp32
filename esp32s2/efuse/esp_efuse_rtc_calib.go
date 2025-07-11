package efuse

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

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
