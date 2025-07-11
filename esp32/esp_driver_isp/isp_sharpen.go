package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief ISP Sharpen configurations
 */

type EspIspSharpenConfigT struct {
	HFreqCoeff                     IspSharpenHFreqCoeffT
	MFreqCoeff                     IspSharpenMFreqCoeff
	HThresh                        c.Uint8T
	LThresh                        c.Uint8T
	PaddingMode                    IspSharpenEdgePaddingModeT
	PaddingData                    c.Uint8T
	SharpenTemplate                [0][0]c.Uint8T
	PaddingLineTailValidStartPixel c.Uint8T
	PaddingLineTailValidEndPixel   c.Uint8T
}

/**
 * @brief ISP Sharpen configuration
 *
 * @note After calling this API, sharpen doesn't take into effect until `esp_isp_sharpen_enable` is called
 *
 * @param[in] proc    Processor handle
 * @param[in] config  Sharpen configurations, set NULL to de-configure the ISP Sharpen
 *
 * @return
 *        - ESP_OK                 On success
 *        - ESP_ERR_INVALID_ARG    If the combination of arguments is invalid
 */
//go:linkname EspIspSharpenConfigure C.esp_isp_sharpen_configure
func EspIspSharpenConfigure(proc IspProcHandleT, config *EspIspSharpenConfigT) EspErrT

/**
 * @brief Enable ISP Sharpen function
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspSharpenEnable C.esp_isp_sharpen_enable
func EspIspSharpenEnable(proc IspProcHandleT) EspErrT

/**
 * @brief Disable ISP Sharpen function
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspSharpenDisable C.esp_isp_sharpen_disable
func EspIspSharpenDisable(proc IspProcHandleT) EspErrT
