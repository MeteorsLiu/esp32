package esp_driver_mcpwm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief MCPWM GPIO fault configuration structure
 */

type McpwmGpioFaultConfigT struct {
	GroupId      c.Int
	IntrPriority c.Int
	GpioNum      c.Int
	Flags        struct {
		ActiveLevel c.Uint32T
		IoLoopBack  c.Uint32T
		PullUp      c.Uint32T
		PullDown    c.Uint32T
	}
}

/**
 * @brief Create MCPWM GPIO fault
 *
 * @param[in] config MCPWM GPIO fault configuration
 * @param[out] ret_fault Returned GPIO fault handle
 * @return
 *      - ESP_OK: Create MCPWM GPIO fault successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM GPIO fault failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM GPIO fault failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create MCPWM GPIO fault failed because can't find free resource
 *      - ESP_FAIL: Create MCPWM GPIO fault failed because of other error
 */
// llgo:link (*McpwmGpioFaultConfigT).McpwmNewGpioFault C.mcpwm_new_gpio_fault
func (recv_ *McpwmGpioFaultConfigT) McpwmNewGpioFault(ret_fault *McpwmFaultHandleT) EspErrT {
	return 0
}

/**
 * @brief MCPWM software fault configuration structure
 */

type McpwmSoftFaultConfigT struct {
	Unused [8]uint8
}

/**
 * @brief Create MCPWM software fault
 *
 * @param[in] config MCPWM software fault configuration
 * @param[out] ret_fault Returned software fault handle
 * @return
 *      - ESP_OK: Create MCPWM software fault successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM software fault failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM software fault failed because out of memory
 *      - ESP_FAIL: Create MCPWM software fault failed because of other error
 */
// llgo:link (*McpwmSoftFaultConfigT).McpwmNewSoftFault C.mcpwm_new_soft_fault
func (recv_ *McpwmSoftFaultConfigT) McpwmNewSoftFault(ret_fault *McpwmFaultHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete MCPWM fault
 *
 * @param[in] fault MCPWM fault handle allocated by `mcpwm_new_gpio_fault()` or `mcpwm_new_soft_fault()`
 * @return
 *      - ESP_OK: Delete MCPWM fault successfully
 *      - ESP_ERR_INVALID_ARG: Delete MCPWM fault failed because of invalid argument
 *      - ESP_FAIL: Delete MCPWM fault failed because of other error
 */
//go:linkname McpwmDelFault C.mcpwm_del_fault
func McpwmDelFault(fault McpwmFaultHandleT) EspErrT

/**
 * @brief Activate the software fault, trigger the fault event for once
 *
 * @param[in] fault MCPWM soft fault, allocated by `mcpwm_new_soft_fault()`
 * @return
 *      - ESP_OK: Trigger MCPWM software fault event successfully
 *      - ESP_ERR_INVALID_ARG: Trigger MCPWM software fault event failed because of invalid argument
 *      - ESP_FAIL: Trigger MCPWM software fault event failed because of other error
 */
//go:linkname McpwmSoftFaultActivate C.mcpwm_soft_fault_activate
func McpwmSoftFaultActivate(fault McpwmFaultHandleT) EspErrT

/**
 * @brief Group of supported MCPWM fault event callbacks
 * @note The callbacks are all running under ISR environment
 */

type McpwmFaultEventCallbacksT struct {
	OnFaultEnter McpwmFaultEventCbT
	OnFaultExit  McpwmFaultEventCbT
}

/**
 * @brief Set event callbacks for MCPWM fault
 *
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 *
 * @param[in] fault MCPWM GPIO fault handle, allocated by `mcpwm_new_gpio_fault()`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname McpwmFaultRegisterEventCallbacks C.mcpwm_fault_register_event_callbacks
func McpwmFaultRegisterEventCallbacks(fault McpwmFaultHandleT, cbs *McpwmFaultEventCallbacksT, user_data c.Pointer) EspErrT
