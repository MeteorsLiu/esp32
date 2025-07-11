package esp_driver_mcpwm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Group of supported MCPWM timer event callbacks
 * @note The callbacks are all running under ISR environment
 */

type McpwmTimerEventCallbacksT struct {
	OnFull  McpwmTimerEventCbT
	OnEmpty McpwmTimerEventCbT
	OnStop  McpwmTimerEventCbT
}

/**
 * @brief MCPWM timer configuration
 */

type McpwmTimerConfigT struct {
	GroupId      c.Int
	ClkSrc       McpwmTimerClockSourceT
	ResolutionHz c.Uint32T
	CountMode    McpwmTimerCountModeT
	PeriodTicks  c.Uint32T
	IntrPriority c.Int
	Flags        struct {
		UpdatePeriodOnEmpty c.Uint32T
		UpdatePeriodOnSync  c.Uint32T
		AllowPd             c.Uint32T
	}
}

/**
 * @brief Create MCPWM timer
 *
 * @param[in] config MCPWM timer configuration
 * @param[out] ret_timer Returned MCPWM timer handle
 * @return
 *      - ESP_OK: Create MCPWM timer successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM timer failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM timer failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create MCPWM timer failed because all hardware timers are used up and no more free one
 *      - ESP_FAIL: Create MCPWM timer failed because of other error
 */
// llgo:link (*McpwmTimerConfigT).McpwmNewTimer C.mcpwm_new_timer
func (recv_ *McpwmTimerConfigT) McpwmNewTimer(ret_timer *McpwmTimerHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete MCPWM timer
 *
 * @param[in] timer MCPWM timer handle, allocated by `mcpwm_new_timer()`
 * @return
 *      - ESP_OK: Delete MCPWM timer successfully
 *      - ESP_ERR_INVALID_ARG: Delete MCPWM timer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Delete MCPWM timer failed because timer is not in init state
 *      - ESP_FAIL: Delete MCPWM timer failed because of other error
 */
//go:linkname McpwmDelTimer C.mcpwm_del_timer
func McpwmDelTimer(timer McpwmTimerHandleT) EspErrT

/**
 * @brief Set a new period for MCPWM timer
 *
 * @note If `mcpwm_timer_config_t::update_period_on_empty` and `mcpwm_timer_config_t::update_period_on_sync` are not set,
 *       the new period will take effect immediately.
 *       Otherwise, the new period will take effect when timer counts to zero or on sync event.
 * @note You may need to use `mcpwm_comparator_set_compare_value` to set a new compare value for MCPWM comparator
 *       in order to keep the same PWM duty cycle.
 *
 * @param[in] timer MCPWM timer handle, allocated by `mcpwm_new_timer`
 * @param[in] period_ticks New period in count ticks
 * @return
 *      - ESP_OK: Set new period for MCPWM timer successfully
 *      - ESP_ERR_INVALID_ARG: Set new period for MCPWM timer failed because of invalid argument
 *      - ESP_FAIL: Set new period for MCPWM timer failed because of other error
 */
//go:linkname McpwmTimerSetPeriod C.mcpwm_timer_set_period
func McpwmTimerSetPeriod(timer McpwmTimerHandleT, period_ticks c.Uint32T) EspErrT

/**
 * @brief Enable MCPWM timer
 *
 * @param[in] timer MCPWM timer handle, allocated by `mcpwm_new_timer()`
 * @return
 *      - ESP_OK: Enable MCPWM timer successfully
 *      - ESP_ERR_INVALID_ARG: Enable MCPWM timer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Enable MCPWM timer failed because timer is enabled already
 *      - ESP_FAIL: Enable MCPWM timer failed because of other error
 */
//go:linkname McpwmTimerEnable C.mcpwm_timer_enable
func McpwmTimerEnable(timer McpwmTimerHandleT) EspErrT

/**
 * @brief Disable MCPWM timer
 *
 * @param[in] timer MCPWM timer handle, allocated by `mcpwm_new_timer()`
 * @return
 *      - ESP_OK: Disable MCPWM timer successfully
 *      - ESP_ERR_INVALID_ARG: Disable MCPWM timer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Disable MCPWM timer failed because timer is disabled already
 *      - ESP_FAIL: Disable MCPWM timer failed because of other error
 */
//go:linkname McpwmTimerDisable C.mcpwm_timer_disable
func McpwmTimerDisable(timer McpwmTimerHandleT) EspErrT

/**
 * @brief Send specific start/stop commands to MCPWM timer
 *
 * @param[in] timer MCPWM timer handle, allocated by `mcpwm_new_timer()`
 * @param[in] command Supported command list for MCPWM timer
 * @return
 *      - ESP_OK: Start or stop MCPWM timer successfully
 *      - ESP_ERR_INVALID_ARG: Start or stop MCPWM timer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Start or stop MCPWM timer failed because timer is not enabled
 *      - ESP_FAIL: Start or stop MCPWM timer failed because of other error
 */
//go:linkname McpwmTimerStartStop C.mcpwm_timer_start_stop
func McpwmTimerStartStop(timer McpwmTimerHandleT, command McpwmTimerStartStopCmdT) EspErrT

/**
 * @brief Set event callbacks for MCPWM timer
 *
 * @note The first call to this function needs to be before the call to `mcpwm_timer_enable`
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 *
 * @param[in] timer MCPWM timer handle, allocated by `mcpwm_new_timer()`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Set event callbacks failed because timer is not in init state
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname McpwmTimerRegisterEventCallbacks C.mcpwm_timer_register_event_callbacks
func McpwmTimerRegisterEventCallbacks(timer McpwmTimerHandleT, cbs *McpwmTimerEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief MCPWM Timer sync phase configuration
 */

type McpwmTimerSyncPhaseConfigT struct {
	SyncSrc    McpwmSyncHandleT
	CountValue c.Uint32T
	Direction  McpwmTimerDirectionT
}

/**
 * @brief Set sync phase for MCPWM timer
 *
 * @param[in] timer MCPWM timer handle, allocated by `mcpwm_new_timer()`
 * @param[in] config MCPWM timer sync phase configuration
 * @return
 *      - ESP_OK: Set sync phase for MCPWM timer successfully
 *      - ESP_ERR_INVALID_ARG: Set sync phase for MCPWM timer failed because of invalid argument
 *      - ESP_FAIL: Set sync phase for MCPWM timer failed because of other error
 */
//go:linkname McpwmTimerSetPhaseOnSync C.mcpwm_timer_set_phase_on_sync
func McpwmTimerSetPhaseOnSync(timer McpwmTimerHandleT, config *McpwmTimerSyncPhaseConfigT) EspErrT
