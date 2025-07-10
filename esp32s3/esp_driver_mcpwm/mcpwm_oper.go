package esp_driver_mcpwm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief MCPWM operator configuration
 */

type McpwmOperatorConfigT struct {
	GroupId      c.Int
	IntrPriority c.Int
	Flags        struct {
		UpdateGenActionOnTez  c.Uint32T
		UpdateGenActionOnTep  c.Uint32T
		UpdateGenActionOnSync c.Uint32T
		UpdateDeadTimeOnTez   c.Uint32T
		UpdateDeadTimeOnTep   c.Uint32T
		UpdateDeadTimeOnSync  c.Uint32T
	}
}

/**
 * @brief Create MCPWM operator
 *
 * @param[in] config MCPWM operator configuration
 * @param[out] ret_oper Returned MCPWM operator handle
 * @return
 *      - ESP_OK: Create MCPWM operator successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM operator failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM operator failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create MCPWM operator failed because can't find free resource
 *      - ESP_FAIL: Create MCPWM operator failed because of other error
 */
// llgo:link (*McpwmOperatorConfigT).McpwmNewOperator C.mcpwm_new_operator
func (recv_ *McpwmOperatorConfigT) McpwmNewOperator(ret_oper *McpwmOperHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete MCPWM operator
 *
 * @param[in] oper MCPWM operator, allocated by `mcpwm_new_operator()`
 * @return
 *      - ESP_OK: Delete MCPWM operator successfully
 *      - ESP_ERR_INVALID_ARG: Delete MCPWM operator failed because of invalid argument
 *      - ESP_FAIL: Delete MCPWM operator failed because of other error
 */
//go:linkname McpwmDelOperator C.mcpwm_del_operator
func McpwmDelOperator(oper McpwmOperHandleT) EspErrT

/**
 * @brief Connect MCPWM operator and timer, so that the operator can be driven by the timer
 *
 * @param[in] oper MCPWM operator handle, allocated by `mcpwm_new_operator()`
 * @param[in] timer MCPWM timer handle, allocated by `mcpwm_new_timer()`
 * @return
 *      - ESP_OK: Connect MCPWM operator and timer successfully
 *      - ESP_ERR_INVALID_ARG: Connect MCPWM operator and timer failed because of invalid argument
 *      - ESP_FAIL: Connect MCPWM operator and timer failed because of other error
 */
//go:linkname McpwmOperatorConnectTimer C.mcpwm_operator_connect_timer
func McpwmOperatorConnectTimer(oper McpwmOperHandleT, timer McpwmTimerHandleT) EspErrT

/**
 * @brief MCPWM brake configuration structure
 */

type McpwmBrakeConfigT struct {
	Fault     McpwmFaultHandleT
	BrakeMode McpwmOperatorBrakeModeT
	Flags     struct {
		CbcRecoverOnTez c.Uint32T
		CbcRecoverOnTep c.Uint32T
	}
}

/**
 * @brief Set brake method for MCPWM operator
 *
 * @param[in] oper MCPWM operator, allocated by `mcpwm_new_operator()`
 * @param[in] config MCPWM brake configuration
 * @return
 *      - ESP_OK: Set trip for operator successfully
 *      - ESP_ERR_INVALID_ARG: Set trip for operator failed because of invalid argument
 *      - ESP_FAIL: Set trip for operator failed because of other error
 */
//go:linkname McpwmOperatorSetBrakeOnFault C.mcpwm_operator_set_brake_on_fault
func McpwmOperatorSetBrakeOnFault(oper McpwmOperHandleT, config *McpwmBrakeConfigT) EspErrT

/**
 * @brief Try to make the operator recover from fault
 *
 * @note To recover from fault or escape from trip, you make sure the fault signal has disappeared already.
 *       Otherwise the recovery can't succeed.
 *
 * @param[in] oper MCPWM operator, allocated by `mcpwm_new_operator()`
 * @param[in] fault MCPWM fault handle
 * @return
 *      - ESP_OK: Recover from fault successfully
 *      - ESP_ERR_INVALID_ARG: Recover from fault failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Recover from fault failed because the fault source is still active
 *      - ESP_FAIL: Recover from fault failed because of other error
 */
//go:linkname McpwmOperatorRecoverFromFault C.mcpwm_operator_recover_from_fault
func McpwmOperatorRecoverFromFault(oper McpwmOperHandleT, fault McpwmFaultHandleT) EspErrT

/**
 * @brief Group of supported MCPWM operator event callbacks
 * @note The callbacks are all running under ISR environment
 */

type McpwmOperatorEventCallbacksT struct {
	OnBrakeCbc McpwmBrakeEventCbT
	OnBrakeOst McpwmBrakeEventCbT
}

/**
 * @brief Set event callbacks for MCPWM operator
 *
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 *
 * @param[in] oper MCPWM operator handle, allocated by `mcpwm_new_operator()`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname McpwmOperatorRegisterEventCallbacks C.mcpwm_operator_register_event_callbacks
func McpwmOperatorRegisterEventCallbacks(oper McpwmOperHandleT, cbs *McpwmOperatorEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief MCPWM carrier configuration structure
 */

type McpwmCarrierConfigT struct {
	ClkSrc               McpwmCarrierClockSourceT
	FrequencyHz          c.Uint32T
	FirstPulseDurationUs c.Uint32T
	DutyCycle            c.Float
	Flags                struct {
		InvertBeforeModulate c.Uint32T
		InvertAfterModulate  c.Uint32T
	}
}

/**
 * @brief Apply carrier feature for MCPWM operator
 *
 * @param[in] oper MCPWM operator, allocated by `mcpwm_new_operator()`
 * @param[in] config MCPWM carrier specific configuration
 * @return
 *      - ESP_OK: Set carrier for operator successfully
 *      - ESP_ERR_INVALID_ARG: Set carrier for operator failed because of invalid argument
 *      - ESP_FAIL: Set carrier for operator failed because of other error
 */
//go:linkname McpwmOperatorApplyCarrier C.mcpwm_operator_apply_carrier
func McpwmOperatorApplyCarrier(oper McpwmOperHandleT, config *McpwmCarrierConfigT) EspErrT
