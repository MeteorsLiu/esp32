package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Hist controller config
 */

type EspIspHistConfigT struct {
	Window           IspWindowT
	HistMode         IspHistSamplingModeT
	RgbCoefficient   IspHistRgbCoefficientT
	WindowWeight     [0]IspHistWeightT
	SegmentThreshold [0]c.Uint32T
}

/**
 * @brief New an ISP hist controller
 *
 * @param[in]  isp_proc   ISP Processor handle
 * @param[in]  hist_cfg    Pointer to hist config. Refer to ``esp_isp_hist_config_t``.
 * @param[out] ret_hdl    hist controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid
 *         - ESP_ERR_INVALID_STATE Invalid state
 *         - ESP_ERR_NOT_FOUND     No free interrupt found with the specified flags
 *         - ESP_ERR_NO_MEM        If out of memory
 */
//go:linkname EspIspNewHistController C.esp_isp_new_hist_controller
func EspIspNewHistController(isp_proc IspProcHandleT, hist_cfg *EspIspHistConfigT, ret_hdl *IspHistCtlrT) EspErrT

/**
 * @brief Delete an ISP hist controller
 *
 * @param[in] hist_ctlr  hist controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspDelHistController C.esp_isp_del_hist_controller
func EspIspDelHistController(hist_ctlr IspHistCtlrT) EspErrT

/**
 * @brief Enable an ISP hist controller
 *
 * @param[in] hist_ctlr  hist controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspHistControllerEnable C.esp_isp_hist_controller_enable
func EspIspHistControllerEnable(hist_ctlr IspHistCtlrT) EspErrT

/**
 * @brief Disable an ISP hist controller
 *
 * @param[in] hist_ctlr  hist controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspHistControllerDisable C.esp_isp_hist_controller_disable
func EspIspHistControllerDisable(hist_ctlr IspHistCtlrT) EspErrT

/**
 * @brief Trigger hist reference statistics for one time and get the result
 *
 * @param[in]  hist_ctlr  hist controller handle
 * @param[in]  timeout_ms Timeout in millisecond
 *                            - timeout_ms < 0:   Won't return until finished
 *                            - timeout_ms = 0:   No timeout, trigger one time statistics and return immediately,
 *                                                in this case, the result won't be assigned in this function,
 *                                                but you can get the result in the callback `esp_isp_hist_cbs_t::on_statistics_done`
 *                            - timeout_ms > 0:   Wait for specified milliseconds, if not finished, then return timeout error
 * @param[out] out_res    hist reference statistics result
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_TIMEOUT       Wait for the result timeout
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspHistControllerGetOneshotStatistics C.esp_isp_hist_controller_get_oneshot_statistics
func EspIspHistControllerGetOneshotStatistics(hist_ctlr IspHistCtlrT, timeout_ms c.Int, out_res *IspHistResultT) EspErrT

/**
 * @brief Start hist continuous statistics of the reference in the window
 * @note  This function is an asynchronous and non-block function,
 *        it will start the continuous statistics and return immediately.
 *        You have to register the hist callback and get the result from the callback event data.
 *
 * @param[in]  hist_ctlr  hist controller handle
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   Null pointer
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspHistControllerStartContinuousStatistics C.esp_isp_hist_controller_start_continuous_statistics
func EspIspHistControllerStartContinuousStatistics(hist_ctlr IspHistCtlrT) EspErrT

/**
 * @brief Stop hist continuous statistics of the reference in the window
 *
 * @param[in]  hist_ctlr  hist controller handle
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   Null pointer
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspHistControllerStopContinuousStatistics C.esp_isp_hist_controller_stop_continuous_statistics
func EspIspHistControllerStopContinuousStatistics(hist_ctlr IspHistCtlrT) EspErrT

/**
 * @brief Event data of callbacks
 */

type EspIspHistEvtDataT struct {
	HistResult IspHistResultT
}

// llgo:type C
type EspIspHistCallbackT func(IspHistCtlrT, *EspIspHistEvtDataT, c.Pointer) bool

/**
 * @brief Group of ISP hist callbacks
 *
 * @note These callbacks are all running in an ISR environment.
 * @note When CONFIG_ISP_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables should be in internal RAM as well.
 */

type EspIspHistCbsT struct {
	OnStatisticsDone EspIspHistCallbackT
}

/**
 * @brief Register hist event callbacks
 *
 * @note User can deregister a previously registered callback by calling this function and setting the to-be-deregistered callback member in
 *       the `cbs` structure to NULL.
 * @note When CONFIG_ISP_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables (including `user_data`) should be in internal RAM as well.
 *
 * @param[in] hist_ctlr         hist controller handle
 * @param[in] cbs              Group of callback functions
 * @param[in] user_data        User data, which will be delivered to the callback functions directly
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_INVALID_STATE: Driver state is invalid, you shouldn't call this API at this moment
 */
//go:linkname EspIspHistRegisterEventCallbacks C.esp_isp_hist_register_event_callbacks
func EspIspHistRegisterEventCallbacks(hist_ctlr IspHistCtlrT, cbs *EspIspHistCbsT, user_data c.Pointer) EspErrT
