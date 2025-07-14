package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief AF controller config
 */

type EspIspAfConfigT struct {
	Window       [0]IspWindowT
	EdgeThresh   c.Int
	IntrPriority c.Int
}

/**
 * @brief New an ISP AF controller
 *
 * @param[in]  isp_proc   ISP Processor handle
 * @param[in]  af_config  Pointer to AF config. Refer to ``esp_isp_af_config_t``.
 * @param[out] ret_hdl    AF controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid
 *         - ESP_ERR_INVALID_STATE Invalid state
 *         - ESP_ERR_NOT_FOUND     No free interrupt found with the specified flags
 *         - ESP_ERR_NO_MEM        If out of memory
 */
//go:linkname EspIspNewAfController C.esp_isp_new_af_controller
func EspIspNewAfController(isp_proc IspProcHandleT, af_config *EspIspAfConfigT, ret_hdl *IspAfCtlrT) EspErrT

/**
 * @brief Delete an ISP AF controller
 *
 * @param[in] af_ctrlr  AF controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspDelAfController C.esp_isp_del_af_controller
func EspIspDelAfController(af_ctrlr IspAfCtlrT) EspErrT

/**
 * @brief Enable an ISP AF controller
 *
 * @param[in] af_ctrlr  AF controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAfControllerEnable C.esp_isp_af_controller_enable
func EspIspAfControllerEnable(af_ctrlr IspAfCtlrT) EspErrT

/**
 * @brief Disable an ISP AF controller
 *
 * @param[in] af_ctrlr  AF controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAfControllerDisable C.esp_isp_af_controller_disable
func EspIspAfControllerDisable(af_ctrlr IspAfCtlrT) EspErrT

/**
 * @brief Trigger AF luminance and definition statistics for one time and get the result
 *
 * @param[in]  af_ctrlr   AF controller handle
 * @param[in]  timeout_ms Timeout in millisecond
 *                            - timeout_ms < 0:   Won't return until finished
 *                            - timeout_ms = 0:   No timeout, trigger one time statistics and return immediately,
 *                                                in this case, the result won't be assigned in this function,
 *                                                but you can get the result in the callback `esp_isp_af_env_detector_evt_cbs_t::on_env_statistics_done`
 *                            - timeout_ms > 0:   Wait for specified milliseconds, if not finished, then return timeout error
 * @param[out] out_res    AF luminance and definition statistics result, can be NULL if `timeout_ms = 0`
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_TIMEOUT       If the waiting time exceeds the specified timeout.
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAfControllerGetOneshotStatistics C.esp_isp_af_controller_get_oneshot_statistics
func EspIspAfControllerGetOneshotStatistics(af_ctrlr IspAfCtlrT, timeout_ms c.Int, out_res *IspAfResultT) EspErrT

/**
 * @brief Start AF continuous statistics of the luminance and definition in the windows
 * @note  This function is an asynchronous and non-block function,
 *        it will start the continuous statistics and return immediately.
 *        You have to register the AF callback and get the result from the callback event data.
 * @note  When continuous mode start, AF environment detector will be invalid
 *
 * @param[in]  af_ctrlr  AF controller handle
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   Null pointer
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAfControllerStartContinuousStatistics C.esp_isp_af_controller_start_continuous_statistics
func EspIspAfControllerStartContinuousStatistics(af_ctrlr IspAfCtlrT) EspErrT

/**
 * @brief Stop AF continuous statistics of the luminance and definition in the windows
 *
 * @param[in]  af_ctrlr  AF controller handle
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   Null pointer
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAfControllerStopContinuousStatistics C.esp_isp_af_controller_stop_continuous_statistics
func EspIspAfControllerStopContinuousStatistics(af_ctrlr IspAfCtlrT) EspErrT

/*---------------------------------------------
                AF Env Detector
----------------------------------------------*/
/**
 * @brief AF environment detector config
 */

type EspIspAfEnvConfigT struct {
	Interval c.Int
}

/**
 * @brief Set ISP AF environment detector
 *
 * @param[in] af_ctrlr    AF controller handle
 * @param[in] env_config  AF Env detector configuration
 *
 * @note    When continuous mode start, AF environment detector will be invalid
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAfControllerSetEnvDetector C.esp_isp_af_controller_set_env_detector
func EspIspAfControllerSetEnvDetector(af_ctrlr IspAfCtlrT, env_config *EspIspAfEnvConfigT) EspErrT

/**
 * @brief Set ISP AF environment detector detecting threshold
 *
 * @param[in] af_ctrlr           AF controller handle
 * @param[in] definition_thresh  Threshold for definition
 * @param[in] luminance_thresh   Threshold for luminance
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAfControllerSetEnvDetectorThreshold C.esp_isp_af_controller_set_env_detector_threshold
func EspIspAfControllerSetEnvDetectorThreshold(af_ctrlr IspAfCtlrT, definition_thresh c.Int, luminance_thresh c.Int) EspErrT

/**
 * @brief Event data structure
 */

type EspIspAfEnvDetectorEvtDataT struct {
	AfResult IspAfResultT
}

// llgo:type C
type EspIspAfEnvDetectorCallbackT func(IspAfCtlrT, *EspIspAfEnvDetectorEvtDataT, c.Pointer) bool

/**
 * @brief Group of ISP AF Env detector callbacks
 *
 * @note These callbacks are all running in an ISR environment.
 * @note When CONFIG_ISP_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables should be in internal RAM as well.
 */

type EspIspAfEnvDetectorEvtCbsT struct {
	OnEnvStatisticsDone EspIspAfEnvDetectorCallbackT
	OnEnvChange         EspIspAfEnvDetectorCallbackT
}

/**
 * @brief Register AF environment detector event callbacks
 *
 * @note User can deregister a previously registered callback by calling this function and setting the to-be-deregistered callback member in
 *       the `cbs` structure to NULL.
 * @note When CONFIG_ISP_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables (including `user_data`) should be in internal RAM as well.
 *
 * @param[in] af_ctrlr         AF controller handle
 * @param[in] cbs              Group of callback functions
 * @param[in] user_data        User data, which will be delivered to the callback functions directly
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_INVALID_STATE: Driver state is invalid, you shouldn't call this API at this moment
 */
//go:linkname EspIspAfEnvDetectorRegisterEventCallbacks C.esp_isp_af_env_detector_register_event_callbacks
func EspIspAfEnvDetectorRegisterEventCallbacks(af_ctrlr IspAfCtlrT, cbs *EspIspAfEnvDetectorEvtCbsT, user_data c.Pointer) EspErrT
