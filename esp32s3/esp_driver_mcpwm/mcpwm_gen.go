package esp_driver_mcpwm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief MCPWM generator configuration
 */

type McpwmGeneratorConfigT struct {
	GenGpioNum c.Int
	Flags      struct {
		InvertPwm  c.Uint32T
		IoLoopBack c.Uint32T
		IoOdMode   c.Uint32T
		PullUp     c.Uint32T
		PullDown   c.Uint32T
	}
}

/**
 * @brief Allocate MCPWM generator from given operator
 *
 * @param[in] oper MCPWM operator, allocated by `mcpwm_new_operator()`
 * @param[in] config MCPWM generator configuration
 * @param[out] ret_gen Returned MCPWM generator
 * @return
 *      - ESP_OK: Create MCPWM generator successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM generator failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM generator failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create MCPWM generator failed because can't find free resource
 *      - ESP_FAIL: Create MCPWM generator failed because of other error
 */
//go:linkname McpwmNewGenerator C.mcpwm_new_generator
func McpwmNewGenerator(oper McpwmOperHandleT, config *McpwmGeneratorConfigT, ret_gen *McpwmGenHandleT) EspErrT

/**
 * @brief Delete MCPWM generator
 *
 * @param[in] gen MCPWM generator handle, allocated by `mcpwm_new_generator()`
 * @return
 *      - ESP_OK: Delete MCPWM generator successfully
 *      - ESP_ERR_INVALID_ARG: Delete MCPWM generator failed because of invalid argument
 *      - ESP_FAIL: Delete MCPWM generator failed because of other error
 */
//go:linkname McpwmDelGenerator C.mcpwm_del_generator
func McpwmDelGenerator(gen McpwmGenHandleT) EspErrT

/**
 * @brief Set force level for MCPWM generator
 *
 * @note The force level will be applied to the generator immediately, regardless any other events that would change the generator's behaviour.
 * @note If the `hold_on` is true, the force level will retain forever, until user removes the force level by setting the force level to `-1`.
 * @note If the `hold_on` is false, the force level can be overridden by the next event action.
 * @note The force level set by this function can be inverted by GPIO matrix or dead-time module. So the level set here doesn't equal to the final output level.
 *
 * @param[in] gen MCPWM generator handle, allocated by `mcpwm_new_generator()`
 * @param[in] level GPIO level to be applied to MCPWM generator, specially, -1 means to remove the force level
 * @param[in] hold_on Whether the forced PWM level should retain (i.e. will remain unchanged until manually remove the force level)
 * @return
 *      - ESP_OK: Set force level for MCPWM generator successfully
 *      - ESP_ERR_INVALID_ARG: Set force level for MCPWM generator failed because of invalid argument
 *      - ESP_FAIL: Set force level for MCPWM generator failed because of other error
 */
//go:linkname McpwmGeneratorSetForceLevel C.mcpwm_generator_set_force_level
func McpwmGeneratorSetForceLevel(gen McpwmGenHandleT, level c.Int, hold_on bool) EspErrT

/**
 * @brief Generator action on specific timer event
 */

type McpwmGenTimerEventActionT struct {
	Direction McpwmTimerDirectionT
	Event     McpwmTimerEventT
	Action    McpwmGeneratorActionT
}

/**
 * @brief Set generator action on MCPWM timer event
 *
 * @param[in] gen MCPWM generator handle, allocated by `mcpwm_new_generator()`
 * @param[in] ev_act MCPWM timer event action, can be constructed by `MCPWM_GEN_TIMER_EVENT_ACTION` helper macro
 * @return
 *      - ESP_OK: Set generator action successfully
 *      - ESP_ERR_INVALID_ARG: Set generator action failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Set generator action failed because of timer is not connected to operator
 *      - ESP_FAIL: Set generator action failed because of other error
 */
//go:linkname McpwmGeneratorSetActionOnTimerEvent C.mcpwm_generator_set_action_on_timer_event
func McpwmGeneratorSetActionOnTimerEvent(gen McpwmGenHandleT, ev_act McpwmGenTimerEventActionT) EspErrT

/**
 * @brief Set generator actions on multiple MCPWM timer events
 *
 * @note This is an aggregation version of `mcpwm_generator_set_action_on_timer_event`, which allows user to set multiple actions in one call.
 *
 * @param[in] gen MCPWM generator handle, allocated by `mcpwm_new_generator()`
 * @param[in] ev_act MCPWM timer event action list, must be terminated by `MCPWM_GEN_TIMER_EVENT_ACTION_END()`
 * @return
 *      - ESP_OK: Set generator actions successfully
 *      - ESP_ERR_INVALID_ARG: Set generator actions failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Set generator actions failed because of timer is not connected to operator
 *      - ESP_FAIL: Set generator actions failed because of other error
 */
//go:linkname McpwmGeneratorSetActionsOnTimerEvent C.mcpwm_generator_set_actions_on_timer_event
func McpwmGeneratorSetActionsOnTimerEvent(gen McpwmGenHandleT, ev_act McpwmGenTimerEventActionT, __llgo_va_list ...interface{}) EspErrT

/**
 * @brief Generator action on specific comparator event
 */

type McpwmGenCompareEventActionT struct {
	Direction  McpwmTimerDirectionT
	Comparator McpwmCmprHandleT
	Action     McpwmGeneratorActionT
}

/**
 * @brief Set generator action on MCPWM compare event
 *
 * @param[in] generator MCPWM generator handle, allocated by `mcpwm_new_generator()`
 * @param[in] ev_act MCPWM compare event action, can be constructed by `MCPWM_GEN_COMPARE_EVENT_ACTION` helper macro
 * @return
 *      - ESP_OK: Set generator action successfully
 *      - ESP_ERR_INVALID_ARG: Set generator action failed because of invalid argument
 *      - ESP_FAIL: Set generator action failed because of other error
 */
//go:linkname McpwmGeneratorSetActionOnCompareEvent C.mcpwm_generator_set_action_on_compare_event
func McpwmGeneratorSetActionOnCompareEvent(generator McpwmGenHandleT, ev_act McpwmGenCompareEventActionT) EspErrT

/**
 * @brief Set generator actions on multiple MCPWM compare events
 *
 * @note This is an aggregation version of `mcpwm_generator_set_action_on_compare_event`, which allows user to set multiple actions in one call.
 *
 * @param[in] generator MCPWM generator handle, allocated by `mcpwm_new_generator()`
 * @param[in] ev_act MCPWM compare event action list, must be terminated by `MCPWM_GEN_COMPARE_EVENT_ACTION_END()`
 * @return
 *      - ESP_OK: Set generator actions successfully
 *      - ESP_ERR_INVALID_ARG: Set generator actions failed because of invalid argument
 *      - ESP_FAIL: Set generator actions failed because of other error
 */
//go:linkname McpwmGeneratorSetActionsOnCompareEvent C.mcpwm_generator_set_actions_on_compare_event
func McpwmGeneratorSetActionsOnCompareEvent(generator McpwmGenHandleT, ev_act McpwmGenCompareEventActionT, __llgo_va_list ...interface{}) EspErrT

/**
 * @brief Generator action on specific brake event
 */

type McpwmGenBrakeEventActionT struct {
	Direction McpwmTimerDirectionT
	BrakeMode McpwmOperatorBrakeModeT
	Action    McpwmGeneratorActionT
}

/**
 * @brief Set generator action on MCPWM brake event
 *
 * @param[in] generator MCPWM generator handle, allocated by `mcpwm_new_generator()`
 * @param[in] ev_act MCPWM brake event action, can be constructed by `MCPWM_GEN_BRAKE_EVENT_ACTION` helper macro
 * @return
 *      - ESP_OK: Set generator action successfully
 *      - ESP_ERR_INVALID_ARG: Set generator action failed because of invalid argument
 *      - ESP_FAIL: Set generator action failed because of other error
 */
//go:linkname McpwmGeneratorSetActionOnBrakeEvent C.mcpwm_generator_set_action_on_brake_event
func McpwmGeneratorSetActionOnBrakeEvent(generator McpwmGenHandleT, ev_act McpwmGenBrakeEventActionT) EspErrT

/**
 * @brief Set generator actions on multiple MCPWM brake events
 *
 * @note This is an aggregation version of `mcpwm_generator_set_action_on_brake_event`, which allows user to set multiple actions in one call.
 *
 * @param[in] generator MCPWM generator handle, allocated by `mcpwm_new_generator()`
 * @param[in] ev_act MCPWM brake event action list, must be terminated by `MCPWM_GEN_BRAKE_EVENT_ACTION_END()`
 * @return
 *      - ESP_OK: Set generator actions successfully
 *      - ESP_ERR_INVALID_ARG: Set generator actions failed because of invalid argument
 *      - ESP_FAIL: Set generator actions failed because of other error
 */
//go:linkname McpwmGeneratorSetActionsOnBrakeEvent C.mcpwm_generator_set_actions_on_brake_event
func McpwmGeneratorSetActionsOnBrakeEvent(generator McpwmGenHandleT, ev_act McpwmGenBrakeEventActionT, __llgo_va_list ...interface{}) EspErrT

/**
 * @brief Generator action on specific fault event
 */

type McpwmGenFaultEventActionT struct {
	Direction McpwmTimerDirectionT
	Fault     McpwmFaultHandleT
	Action    McpwmGeneratorActionT
}

/**
 * @brief Set generator action on MCPWM Fault event
 *
 * @param[in] generator MCPWM generator handle, allocated by `mcpwm_new_generator()`
 * @param[in] ev_act MCPWM trigger event action, can be constructed by `MCPWM_GEN_FAULT_EVENT_ACTION` helper macro
 * @return
 *      - ESP_OK: Set generator action successfully
 *      - ESP_ERR_INVALID_ARG: Set generator action failed because of invalid argument
 *      - ESP_FAIL: Set generator action failed because of other error
 */
//go:linkname McpwmGeneratorSetActionOnFaultEvent C.mcpwm_generator_set_action_on_fault_event
func McpwmGeneratorSetActionOnFaultEvent(generator McpwmGenHandleT, ev_act McpwmGenFaultEventActionT) EspErrT

/**
 * @brief Generator action on specific sync event
 */

type McpwmGenSyncEventActionT struct {
	Direction McpwmTimerDirectionT
	Sync      McpwmSyncHandleT
	Action    McpwmGeneratorActionT
}

/**
 * @brief Set generator action on MCPWM Sync event
 *
 * @note The trigger only support one sync action, regardless of the kinds. Should not call this function more than once.
 *
 * @param[in] generator MCPWM generator handle, allocated by `mcpwm_new_generator()`
 * @param[in] ev_act MCPWM trigger event action, can be constructed by `MCPWM_GEN_SYNC_EVENT_ACTION` helper macro
 * @return
 *      - ESP_OK: Set generator action successfully
 *      - ESP_ERR_INVALID_ARG: Set generator action failed because of invalid argument
 *      - ESP_FAIL: Set generator action failed because of other error
 */
//go:linkname McpwmGeneratorSetActionOnSyncEvent C.mcpwm_generator_set_action_on_sync_event
func McpwmGeneratorSetActionOnSyncEvent(generator McpwmGenHandleT, ev_act McpwmGenSyncEventActionT) EspErrT

/**
 * @brief MCPWM dead time configuration structure
 */

type McpwmDeadTimeConfigT struct {
	PosedgeDelayTicks c.Uint32T
	NegedgeDelayTicks c.Uint32T
	Flags             struct {
		InvertOutput c.Uint32T
	}
}

/**
 * @brief Set dead time for MCPWM generator
 *
 * @note Due to a hardware limitation, you can't set rising edge delay for both MCPWM generator 0 and 1 at the same time,
 *       otherwise, there will be a conflict inside the dead time module. The same goes for the falling edge setting.
 *       But you can set both the rising edge and falling edge delay for the same MCPWM generator.
 *
 * @param[in] in_generator MCPWM generator, before adding the dead time
 * @param[in] out_generator MCPWM generator, after adding the dead time
 * @param[in] config MCPWM dead time configuration
 * @return
 *      - ESP_OK: Set dead time for MCPWM generator successfully
 *      - ESP_ERR_INVALID_ARG: Set dead time for MCPWM generator failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Set dead time for MCPWM generator failed because of invalid state (e.g. delay module is already in use by other generator)
 *      - ESP_FAIL: Set dead time for MCPWM generator failed because of other error
 */
//go:linkname McpwmGeneratorSetDeadTime C.mcpwm_generator_set_dead_time
func McpwmGeneratorSetDeadTime(in_generator McpwmGenHandleT, out_generator McpwmGenHandleT, config *McpwmDeadTimeConfigT) EspErrT
