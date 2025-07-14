package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief LSC Gain array
 */

type EspIspLscGainArrayT struct {
	GainR  *IspLscGainT
	GainGr *IspLscGainT
	GainGb *IspLscGainT
	GainB  *IspLscGainT
}

/**
 * @brief ISP LSC configurations
 */

type EspIspLscConfigT struct {
	GainArray *EspIspLscGainArrayT
}

/**
 * @brief Helper function to allocate gain array for LSC
 *
 * @param[in] isp_proc                     Processor handle
 * @param[in] gain_array                   Gain array to be allocated
 * @param[out] out_array_size_per_channel  Array size
 *
 * @return
 *        - ESP_OK                 On success
 *        - ESP_ERR_INVALID_STATE  Not allowed to be called under current state
 *        - ESP_ERR_INVALID_ARG    If the combination of arguments is invalid
 *        - ESP_ERR_NO_MEM         Out of memory
 */
//go:linkname EspIspLscAllocateGainArray C.esp_isp_lsc_allocate_gain_array
func EspIspLscAllocateGainArray(isp_proc IspProcHandleT, gain_array *EspIspLscGainArrayT, out_array_size_per_channel *c.SizeT) EspErrT

/**
 * @brief ISP LSC configuration
 *
 * @note After calling this API, LSC doesn't take into effect until `esp_isp_lsc_enable` is called
 *
 * @param[in] isp_proc    Processor handle
 * @param[in] config      LSC configurations
 *
 * @return
 *        - ESP_OK                 On success
 *        - ESP_ERR_INVALID_STATE  Not allowed to be called under current state
 *        - ESP_ERR_INVALID_ARG    If the combination of arguments is invalid
 *        - ESP_ERR_NOT_SUPPORTED  Not supported
 */
//go:linkname EspIspLscConfigure C.esp_isp_lsc_configure
func EspIspLscConfigure(isp_proc IspProcHandleT, config *EspIspLscConfigT) EspErrT

/**
 * @brief Enable ISP LSC function
 *
 * @param[in] isp_proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspLscEnable C.esp_isp_lsc_enable
func EspIspLscEnable(isp_proc IspProcHandleT) EspErrT

/**
 * @brief Disable ISP LSC function
 *
 * @param[in] isp_proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspLscDisable C.esp_isp_lsc_disable
func EspIspLscDisable(isp_proc IspProcHandleT) EspErrT
