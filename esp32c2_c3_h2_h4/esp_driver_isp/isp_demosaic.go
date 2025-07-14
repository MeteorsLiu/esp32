package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief ISP Demosaic configurations
 */

type EspIspDemosaicConfigT struct {
	GradRatio                      IspDemosaicGradRatioT
	PaddingMode                    IspDemosaicEdgePaddingModeT
	PaddingData                    c.Uint8T
	PaddingLineTailValidStartPixel c.Uint8T
	PaddingLineTailValidEndPixel   c.Uint8T
}

/**
 * @brief ISP Demosaic configuration
 *
 * @note After calling this API, Demosaic doesn't take into effect until `esp_isp_demosaic_enable` is called
 *
 * @param[in] proc    Processor handle
 * @param[in] config  Demosaic configurations, set NULL to de-configure the ISP Demosaic
 *
 * @return
 *        - ESP_OK                 On success
 *        - ESP_ERR_INVALID_ARG    If the combination of arguments is invalid
 */
//go:linkname EspIspDemosaicConfigure C.esp_isp_demosaic_configure
func EspIspDemosaicConfigure(proc IspProcHandleT, config *EspIspDemosaicConfigT) EspErrT

/**
 * @brief Enable ISP Demosaic function
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspDemosaicEnable C.esp_isp_demosaic_enable
func EspIspDemosaicEnable(proc IspProcHandleT) EspErrT

/**
 * @brief Disable ISP Demosaic function
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspDemosaicDisable C.esp_isp_demosaic_disable
func EspIspDemosaicDisable(proc IspProcHandleT) EspErrT
