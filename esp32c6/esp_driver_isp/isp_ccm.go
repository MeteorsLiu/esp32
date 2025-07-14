package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Color Correction Matrix configurations
 *
 */

type EspIspCcmConfigT struct {
	Matrix     [0][0]c.Float
	Saturation bool
}

/**
 * @brief ISP Color Correction Matrix (CCM) configuration
 *
 * @note This function is allowed to be called before or after `esp_isp_ccm_enable`,
 *       but it only takes effect until `esp_isp_ccm_enable` is called
 *
 * @param[in] proc    Processor handle
 * @param[in] ccm_cfg CCM configurations
 *
 * @return
 *        - ESP_OK                 On success
 *        - ESP_ERR_INVALID_ARG    If the combination of arguments is invalid
 */
//go:linkname EspIspCcmConfigure C.esp_isp_ccm_configure
func EspIspCcmConfigure(proc IspProcHandleT, ccm_cfg *EspIspCcmConfigT) EspErrT

/**
 * @brief Enable ISP CCM function
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 */
//go:linkname EspIspCcmEnable C.esp_isp_ccm_enable
func EspIspCcmEnable(proc IspProcHandleT) EspErrT

/**
 * @brief Disable ISP CCM function
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 */
//go:linkname EspIspCcmDisable C.esp_isp_ccm_disable
func EspIspCcmDisable(proc IspProcHandleT) EspErrT
