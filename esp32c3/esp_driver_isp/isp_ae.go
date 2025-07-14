package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief AE controller config
 */

type EspIspAeConfigT struct {
	SamplePoint  IspAeSamplePointT
	Window       IspWindowT
	IntrPriority c.Int
}

/**
 * @brief New an ISP AE controller
 *
 * @param[in]  isp_proc   ISP Processor handle
 * @param[in]  ae_config  Pointer to AE config. Refer to ``esp_isp_ae_config_t``.
 * @param[out] ret_hdl    AE controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid
 *         - ESP_ERR_INVALID_STATE Invalid state
 *         - ESP_ERR_NOT_FOUND     No free interrupt found with the specified flags
 *         - ESP_ERR_NO_MEM        If out of memory
 */
//go:linkname EspIspNewAeController C.esp_isp_new_ae_controller
func EspIspNewAeController(isp_proc IspProcHandleT, ae_config *EspIspAeConfigT, ret_hdl *IspAeCtlrT) EspErrT

/**
 * @brief Delete an ISP AE controller
 *
 * @param[in] ae_ctlr  AE controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspDelAeController C.esp_isp_del_ae_controller
func EspIspDelAeController(ae_ctlr IspAeCtlrT) EspErrT

/**
 * @brief Enable an ISP AE controller
 *
 * @param[in] ae_ctlr  AE controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAeControllerEnable C.esp_isp_ae_controller_enable
func EspIspAeControllerEnable(ae_ctlr IspAeCtlrT) EspErrT

/**
 * @brief Disable an ISP AE controller
 *
 * @param[in] ae_ctlr  AE controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAeControllerDisable C.esp_isp_ae_controller_disable
func EspIspAeControllerDisable(ae_ctlr IspAeCtlrT) EspErrT

/**
 * @brief Trigger AE luminance statistics for one time and get the result
 *
 * @param[in]  ae_ctlr   AE controller handle
 * @param[in]  timeout_ms Timeout in millisecond
 *                            - timeout_ms < 0:   Won't return until finished
 *                            - timeout_ms = 0:   No timeout, trigger one time statistics and return immediately,
 *                                                in this case, the result won't be assigned in this function,
 *                                                but you can get the result in the callback `esp_isp_ae_env_detector_evt_cbs_t::on_env_statistics_done`
 *                            - timeout_ms > 0:   Wait for specified milliseconds, if not finished, then return timeout error
 * @param[out] out_res    AE luminance statistics result, can be NULL if `timeout_ms = 0`
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_TIMEOUT       If the waiting time exceeds the specified timeout.
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAeControllerGetOneshotStatistics C.esp_isp_ae_controller_get_oneshot_statistics
func EspIspAeControllerGetOneshotStatistics(ae_ctlr IspAeCtlrT, timeout_ms c.Int, out_res *IspAeResultT) EspErrT

/**
 * @brief Start AE continuous statistics of the luminance in the windows
 * @note  This function is an asynchronous and non-block function,
 *        it will start the continuous statistics and return immediately.
 *        You have to register the AE callback and get the result from the callback event data.
 * @note  When using oneshot statistics, the AE Environment Detector will be temporarily disabled
 *        and will automatically recover once the oneshot is complete.
 * @param[in]  ae_ctlr  AE controller handle
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   Null pointer
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAeControllerStartContinuousStatistics C.esp_isp_ae_controller_start_continuous_statistics
func EspIspAeControllerStartContinuousStatistics(ae_ctlr IspAeCtlrT) EspErrT

/**
 * @brief Stop AE continuous statistics of the luminance in the windows
 *
 * @param[in]  ae_ctlr  AE controller handle
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   Null pointer
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAeControllerStopContinuousStatistics C.esp_isp_ae_controller_stop_continuous_statistics
func EspIspAeControllerStopContinuousStatistics(ae_ctlr IspAeCtlrT) EspErrT

/*---------------------------------------------
                AE env detector
----------------------------------------------*/
/**
 * @brief AE environment detector config
 */

type EspIspAeEnvConfigT struct {
	Interval c.Int
}

/**
 * @brief AE environment detector config
 */

type EspIspAeEnvThreshT struct {
	LowThresh  c.Int
	HighThresh c.Int
}

/**
 * @brief Set ISP AE environment detector
 *
 * @param[in] ae_ctlr    AE controller handle
 * @param[in] env_config  AE Env detector configuration
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAeControllerSetEnvDetector C.esp_isp_ae_controller_set_env_detector
func EspIspAeControllerSetEnvDetector(ae_ctlr IspAeCtlrT, env_config *EspIspAeEnvConfigT) EspErrT

/**
 * @brief Set ISP AE environment detector detecting threshold
 *
 * @param[in] ae_ctlr           AE controller handle
 * @param[in] env_thresh         Luminance thresholds for AE env detector
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAeControllerSetEnvDetectorThreshold C.esp_isp_ae_controller_set_env_detector_threshold
func EspIspAeControllerSetEnvDetectorThreshold(ae_ctlr IspAeCtlrT, env_thresh *EspIspAeEnvThreshT) EspErrT

/**
 * @brief Event data structure
 */

type EspIspAeEnvDetectorEvtDataT struct {
	AeResult IspAeResultT
}

// llgo:type C
type EspIspAeEnvDetectorCallbackT func(IspAeCtlrT, *EspIspAeEnvDetectorEvtDataT, c.Pointer) bool

/**
 * @brief Group of ISP AE env_detector
 * @note These callbacks are all running in an ISR environment.
 * @note When CONFIG_ISP_ISR_IRAM_SAEE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables should be in internal RAM as well.
 */

type EspIspAeEnvDetectorEvtCbsT struct {
	OnEnvStatisticsDone EspIspAeEnvDetectorCallbackT
	OnEnvChange         EspIspAeEnvDetectorCallbackT
}

/**
 * @brief Register AE Env detector event callbacks
 *
 * @note User can deregister a previously registered callback by calling this function and setting the to-be-deregistered callback member in
 *       the `cbs` structure to NULL.
 * @note When CONFIG_ISP_ISR_IRAM_SAEE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables (including `user_data`) should be in internal RAM as well.
 *
 * @param[in] ae_ctlr         AE controller handle
 * @param[in] cbs              Group of callback functions
 * @param[in] user_data        User data, which will be delivered to the callback functions directly
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_INVALID_STATE: Driver state is invalid, you shouldn't call this API at this moment
 */
//go:linkname EspIspAeEnvDetectorRegisterEventCallbacks C.esp_isp_ae_env_detector_register_event_callbacks
func EspIspAeEnvDetectorRegisterEventCallbacks(ae_ctlr IspAeCtlrT, cbs *EspIspAeEnvDetectorEvtCbsT, user_data c.Pointer) EspErrT
