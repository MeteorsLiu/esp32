package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief ISP configurations
 */

type EspIspProcessorCfgT struct {
	ClkSrc              IspClkSrcT
	ClkHz               c.Uint32T
	InputDataSource     IspInputDataSourceT
	InputDataColorType  IspColorT
	OutputDataColorType IspColorT
	YuvRange            IspColorRangeT
	YuvStd              IspYuvConvStdT
	HasLineStartPacket  bool
	HasLineEndPacket    bool
	HRes                c.Uint32T
	VRes                c.Uint32T
	BayerOrder          ColorRawElementOrderT
	IntrPriority        c.Int
}

/**
 * @brief New an ISP processor
 *
 * @param[in]  proc_config  Pointer to ISP config. Refer to ``esp_isp_processor_cfg_t``.
 * @param[out] ret_proc     Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_NOT_FOUND     No free interrupt found with the specified flags
 *         - ESP_ERR_NOT_SUPPORTED Not supported mode
 *         - ESP_ERR_NO_MEM        If out of memory
 */
// llgo:link (*EspIspProcessorCfgT).EspIspNewProcessor C.esp_isp_new_processor
func (recv_ *EspIspProcessorCfgT) EspIspNewProcessor(ret_proc *IspProcHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete an ISP processor
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspDelProcessor C.esp_isp_del_processor
func EspIspDelProcessor(proc IspProcHandleT) EspErrT

/**
 * @brief Enable an ISP processor
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspEnable C.esp_isp_enable
func EspIspEnable(proc IspProcHandleT) EspErrT

/**
 * @brief Disable an ISP processor
 *
 * @param[in] proc  Processor handle
 *
 * @return
 *         - ESP_OK                On success
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 */
//go:linkname EspIspDisable C.esp_isp_disable
func EspIspDisable(proc IspProcHandleT) EspErrT

/*---------------------------------------------
                Event Callbacks
----------------------------------------------*/
/**
 * @brief Register ISP event callbacks
 *
 * @note User can deregister a previously registered callback by calling this function and setting the to-be-deregistered callback member in
 *       the `cbs` structure to NULL.
 * @note When CONFIG_ISP_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables (including `user_data`) should be in internal RAM as well.
 *
 * @param[in] proc             Processor handle
 * @param[in] cbs              Group of callback functions
 * @param[in] user_data        User data, which will be delivered to the callback functions directly
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_INVALID_STATE: Driver state is invalid, you shouldn't call this API at this moment
 */
//go:linkname EspIspRegisterEventCallbacks C.esp_isp_register_event_callbacks
func EspIspRegisterEventCallbacks(proc IspProcHandleT, cbs *EspIspEvtCbsT, user_data c.Pointer) EspErrT
