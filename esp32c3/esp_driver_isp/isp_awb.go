package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief AWB controller config
 */

type EspIspAwbConfigT struct {
	SamplePoint IspAwbSamplePointT
	Window      IspWindowT
	WhitePatch  struct {
		Luminance      IspU32RangeT
		RedGreenRatio  IspFloatRangeT
		BlueGreenRatio IspFloatRangeT
	}
	IntrPriority c.Int
}

/**
 * @brief New an ISP AWB controller
 *
 * @param[in]  isp_proc   ISP Processor handle
 * @param[in]  awb_cfg    Pointer to AWB config. Refer to ``esp_isp_awb_config_t``.
 * @param[out] ret_hdl    AWB controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid
 *         - ESP_ERR_INVALID_STATE Invalid state
 *         - ESP_ERR_NOT_FOUND     No free interrupt found with the specified flags
 *         - ESP_ERR_NO_MEM        If out of memory
 */
//go:linkname EspIspNewAwbController C.esp_isp_new_awb_controller
func EspIspNewAwbController(isp_proc IspProcHandleT, awb_cfg *EspIspAwbConfigT, ret_hdl *IspAwbCtlrT) EspErrT

/**
 * @brief Delete an ISP AWB controller
 *
 * @param[in] awb_ctlr  AWB controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspDelAwbController C.esp_isp_del_awb_controller
func EspIspDelAwbController(awb_ctlr IspAwbCtlrT) EspErrT

/**
 * @brief Reconfigure the ISP AWB controller
 * @note  This function is allowed to be called no matter the awb controller is enabled or not.
 *
 * @param[in] awb_ctlr  AWB controller handle
 * @param[in]  awb_cfg    Pointer to AWB config. Refer to ``esp_isp_awb_config_t``
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid
 */
//go:linkname EspIspAwbControllerReconfig C.esp_isp_awb_controller_reconfig
func EspIspAwbControllerReconfig(awb_ctlr IspAwbCtlrT, awb_cfg *EspIspAwbConfigT) EspErrT

/**
 * @brief Enable an ISP AWB controller
 *
 * @param[in] awb_ctlr  AWB controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAwbControllerEnable C.esp_isp_awb_controller_enable
func EspIspAwbControllerEnable(awb_ctlr IspAwbCtlrT) EspErrT

/**
 * @brief Disable an ISP AWB controller
 *
 * @param[in] awb_ctlr  AWB controller handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAwbControllerDisable C.esp_isp_awb_controller_disable
func EspIspAwbControllerDisable(awb_ctlr IspAwbCtlrT) EspErrT

/**
 * @brief Trigger AWB white patch statistics for one time and get the result
 *
 * @param[in]  awb_ctlr  AWB controller handle
 * @param[in]  timeout_ms Timeout in millisecond
 *                            - timeout_ms < 0:   Won't return until finished
 *                            - timeout_ms = 0:   No timeout, trigger one time statistics and return immediately,
 *                                                in this case, the result won't be assigned in this function,
 *                                                but you can get the result in the callback `esp_isp_awb_cbs_t::on_statistics_done`
 *                            - timeout_ms > 0:   Wait for specified milliseconds, if not finished, then return timeout error
 * @param[out] out_res    AWB white patch statistics result
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_TIMEOUT       Wait for the result timeout
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAwbControllerGetOneshotStatistics C.esp_isp_awb_controller_get_oneshot_statistics
func EspIspAwbControllerGetOneshotStatistics(awb_ctlr IspAwbCtlrT, timeout_ms c.Int, out_res *IspAwbStatResultT) EspErrT

/**
 * @brief Start AWB continuous statistics of the white patch in the window
 * @note  This function is an asynchronous and non-block function,
 *        it will start the continuous statistics and return immediately.
 *        You have to register the AWB callback and get the result from the callback event data.
 *
 * @param[in]  awb_ctlr  AWB controller handle
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   Null pointer
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAwbControllerStartContinuousStatistics C.esp_isp_awb_controller_start_continuous_statistics
func EspIspAwbControllerStartContinuousStatistics(awb_ctlr IspAwbCtlrT) EspErrT

/**
 * @brief Stop AWB continuous statistics of the white patch in the window
 *
 * @param[in]  awb_ctlr  AWB controller handle
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   Null pointer
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspAwbControllerStopContinuousStatistics C.esp_isp_awb_controller_stop_continuous_statistics
func EspIspAwbControllerStopContinuousStatistics(awb_ctlr IspAwbCtlrT) EspErrT

/**
 * @brief Event data of callbacks
 */

type EspIspAwbEvtDataT struct {
	AwbResult IspAwbStatResultT
}

// llgo:type C
type EspIspAwbCallbackT func(IspAwbCtlrT, *EspIspAwbEvtDataT, c.Pointer) bool

/**
 * @brief Group of ISP AWB callbacks
 *
 * @note These callbacks are all running in an ISR environment.
 * @note When CONFIG_ISP_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables should be in internal RAM as well.
 */

type EspIspAwbCbsT struct {
	OnStatisticsDone EspIspAwbCallbackT
}

/**
 * @brief Register AWB event callbacks
 *
 * @note User can deregister a previously registered callback by calling this function and setting the to-be-deregistered callback member in
 *       the `cbs` structure to NULL.
 * @note When CONFIG_ISP_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables (including `user_data`) should be in internal RAM as well.
 *
 * @param[in] awb_ctlr         AWB controller handle
 * @param[in] cbs              Group of callback functions
 * @param[in] user_data        User data, which will be delivered to the callback functions directly
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_INVALID_STATE: Driver state is invalid, you shouldn't call this API at this moment
 */
//go:linkname EspIspAwbRegisterEventCallbacks C.esp_isp_awb_register_event_callbacks
func EspIspAwbRegisterEventCallbacks(awb_ctlr IspAwbCtlrT, cbs *EspIspAwbCbsT, user_data c.Pointer) EspErrT
