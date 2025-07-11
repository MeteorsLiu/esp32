package esp_driver_gpio

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GpioGlitchFilterT struct {
	Unused [8]uint8
}
type GpioGlitchFilterHandleT *GpioGlitchFilterT

/**
 * @brief Configuration of GPIO pin glitch filter
 */

type GpioPinGlitchFilterConfigT struct {
	ClkSrc  GlitchFilterClockSourceT
	GpioNum GpioNumT
}

/**
 * @brief Configuration of GPIO flex glitch filter
 */

type GpioFlexGlitchFilterConfigT struct {
	ClkSrc        GlitchFilterClockSourceT
	GpioNum       GpioNumT
	WindowWidthNs c.Uint32T
	WindowThresNs c.Uint32T
}

/**
 * @brief Delete a glitch filter
 *
 * @param[in] filter Glitch filter handle returned from `gpio_new_flex_glitch_filter` or `gpio_new_pin_glitch_filter`
 * @return
 *      - ESP_OK: Delete glitch filter successfully
 *      - ESP_ERR_INVALID_ARG: Delete glitch filter failed because of invalid arguments
 *      - ESP_ERR_INVALID_STATE: Delete glitch filter failed because the glitch filter is still in working
 *      - ESP_FAIL: Delete glitch filter failed because of other error
 */
//go:linkname GpioDelGlitchFilter C.gpio_del_glitch_filter
func GpioDelGlitchFilter(filter GpioGlitchFilterHandleT) EspErrT

/**
 * @brief Enable a glitch filter
 *
 * @param[in] filter Glitch filter handle returned from `gpio_new_flex_glitch_filter` or `gpio_new_pin_glitch_filter`
 * @return
 *      - ESP_OK: Enable glitch filter successfully
 *      - ESP_ERR_INVALID_ARG: Enable glitch filter failed because of invalid arguments
 *      - ESP_ERR_INVALID_STATE: Enable glitch filter failed because the glitch filter is already enabled
 *      - ESP_FAIL: Enable glitch filter failed because of other error
 */
//go:linkname GpioGlitchFilterEnable C.gpio_glitch_filter_enable
func GpioGlitchFilterEnable(filter GpioGlitchFilterHandleT) EspErrT

/**
 * @brief Disable a glitch filter
 *
 * @param[in] filter Glitch filter handle returned from `gpio_new_flex_glitch_filter` or `gpio_new_pin_glitch_filter`
 * @return
 *      - ESP_OK: Disable glitch filter successfully
 *      - ESP_ERR_INVALID_ARG: Disable glitch filter failed because of invalid arguments
 *      - ESP_ERR_INVALID_STATE: Disable glitch filter failed because the glitch filter is not enabled yet
 *      - ESP_FAIL: Disable glitch filter failed because of other error
 */
//go:linkname GpioGlitchFilterDisable C.gpio_glitch_filter_disable
func GpioGlitchFilterDisable(filter GpioGlitchFilterHandleT) EspErrT
