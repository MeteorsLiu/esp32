package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief ISP color configurations
 */

type EspIspColorConfigT struct {
	ColorContrast   IspColorContrastT
	ColorSaturation IspColorSaturationT
	ColorHue        c.Uint32T
	ColorBrightness c.Int
}

/**
 * @brief ISP Color configuration
 *
 * @note After calling this API, Color doesn't take into effect until `esp_isp_color_enable` is called
 * @note API is ISR available
 *
 * @param[in] proc    Processor handle
 * @param[in] config  Color configurations, set NULL to de-configure the ISP Color
 *
 * @return
 *        - ESP_OK                 On success
 *        - ESP_ERR_INVALID_ARG    If the combination of arguments is invalid
 */
//go:linkname EspIspColorConfigure C.esp_isp_color_configure
func EspIspColorConfigure(proc IspProcHandleT, config *EspIspColorConfigT) EspErrT

/**
 * @brief Enable ISP color function
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspColorEnable C.esp_isp_color_enable
func EspIspColorEnable(proc IspProcHandleT) EspErrT

/**
 * @brief Disable ISP color function
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspColorDisable C.esp_isp_color_disable
func EspIspColorDisable(proc IspProcHandleT) EspErrT
