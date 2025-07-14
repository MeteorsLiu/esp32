package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief ISP BF configurations
 */

type EspIspBfConfigT struct {
	PaddingMode                    IspBfEdgePaddingModeT
	PaddingData                    c.Uint8T
	BfTemplate                     [0][0]c.Uint8T
	DenoisingLevel                 c.Uint8T
	PaddingLineTailValidStartPixel c.Uint8T
	PaddingLineTailValidEndPixel   c.Uint8T
}

/**
 * @brief ISP BF configuration
 *
 * @note After calling this API, BF doesn't take into effect until `esp_isp_bf_enable` is called
 *
 * @param[in] proc    Processor handle
 * @param[in] config  BF configurations, set NULL to de-configure the ISP BF
 *
 * @return
 *        - ESP_OK                 On success
 *        - ESP_ERR_INVALID_STATE  Not allowed to be called under current state
 *        - ESP_ERR_INVALID_ARG    If the combination of arguments is invalid
 */
//go:linkname EspIspBfConfigure C.esp_isp_bf_configure
func EspIspBfConfigure(proc IspProcHandleT, config *EspIspBfConfigT) EspErrT

/**
 * @brief Enable ISP BF function
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspBfEnable C.esp_isp_bf_enable
func EspIspBfEnable(proc IspProcHandleT) EspErrT

/**
 * @brief Disable ISP BF function
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspBfDisable C.esp_isp_bf_disable
func EspIspBfDisable(proc IspProcHandleT) EspErrT
