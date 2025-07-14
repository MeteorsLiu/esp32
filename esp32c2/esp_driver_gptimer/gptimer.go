package esp_driver_gptimer

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief General Purpose Timer configuration
 */

type GptimerConfigT struct {
	ClkSrc       GptimerClockSourceT
	Direction    GptimerCountDirectionT
	ResolutionHz c.Uint32T
	IntrPriority c.Int
	Flags        struct {
		IntrShared        c.Uint32T
		AllowPd           c.Uint32T
		BackupBeforeSleep c.Uint32T
	}
}

/**
 * @brief Create a new General Purpose Timer, and return the handle
 *
 * @note The newly created timer is put in the "init" state.
 *
 * @param[in] config GPTimer configuration
 * @param[out] ret_timer Returned timer handle
 * @return
 *      - ESP_OK: Create GPTimer successfully
 *      - ESP_ERR_INVALID_ARG: Create GPTimer failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create GPTimer failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create GPTimer failed because all hardware timers are used up and no more free one
 *      - ESP_FAIL: Create GPTimer failed because of other error
 */
// llgo:link (*GptimerConfigT).GptimerNewTimer C.gptimer_new_timer
func (recv_ *GptimerConfigT) GptimerNewTimer(ret_timer *GptimerHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete the GPTimer handle
 *
 * @note A timer must be in the "init" state before it can be deleted.
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @return
 *      - ESP_OK: Delete GPTimer successfully
 *      - ESP_ERR_INVALID_ARG: Delete GPTimer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Delete GPTimer failed because the timer is not in init state
 *      - ESP_FAIL: Delete GPTimer failed because of other error
 */
//go:linkname GptimerDelTimer C.gptimer_del_timer
func GptimerDelTimer(timer GptimerHandleT) EspErrT

/**
 * @brief Set GPTimer raw count value
 *
 * @note When updating the raw count of an active timer, the timer will immediately start counting from the new value.
 * @note This function is allowed to run within ISR context
 * @note If `CONFIG_GPTIMER_CTRL_FUNC_IN_IRAM` is enabled, this function will be placed in the IRAM by linker,
 *       makes it possible to execute even when the Flash Cache is disabled.
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @param[in] value Count value to be set
 * @return
 *      - ESP_OK: Set GPTimer raw count value successfully
 *      - ESP_ERR_INVALID_ARG: Set GPTimer raw count value failed because of invalid argument
 *      - ESP_FAIL: Set GPTimer raw count value failed because of other error
 */
//go:linkname GptimerSetRawCount C.gptimer_set_raw_count
func GptimerSetRawCount(timer GptimerHandleT, value c.Uint64T) EspErrT

/**
 * @brief Get GPTimer raw count value
 *
 * @note This function will trigger a software capture event and then return the captured count value.
 * @note With the raw count value and the resolution returned from `gptimer_get_resolution`, you can convert the count value into seconds.
 * @note This function is allowed to run within ISR context
 * @note If `CONFIG_GPTIMER_CTRL_FUNC_IN_IRAM` is enabled, this function will be placed in the IRAM by linker,
 *       makes it possible to execute even when the Flash Cache is disabled.
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @param[out] value Returned GPTimer count value
 * @return
 *      - ESP_OK: Get GPTimer raw count value successfully
 *      - ESP_ERR_INVALID_ARG: Get GPTimer raw count value failed because of invalid argument
 *      - ESP_FAIL: Get GPTimer raw count value failed because of other error
 */
//go:linkname GptimerGetRawCount C.gptimer_get_raw_count
func GptimerGetRawCount(timer GptimerHandleT, value *c.Uint64T) EspErrT

/**
 * @brief Return the real resolution of the timer
 *
 * @note usually the timer resolution is same as what you configured in the `gptimer_config_t::resolution_hz`,
 *       but some unstable clock source (e.g. RC_FAST) will do a calibration, the real resolution can be different from the configured one.
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @param[out] out_resolution Returned timer resolution, in Hz
 * @return
 *      - ESP_OK: Get GPTimer resolution successfully
 *      - ESP_ERR_INVALID_ARG: Get GPTimer resolution failed because of invalid argument
 *      - ESP_FAIL: Get GPTimer resolution failed because of other error
 */
//go:linkname GptimerGetResolution C.gptimer_get_resolution
func GptimerGetResolution(timer GptimerHandleT, out_resolution *c.Uint32T) EspErrT

/**
 * @brief Get GPTimer captured count value
 *
 * @note Different from `gptimer_get_raw_count`, this function won't trigger a software capture event. It just returns the last captured count value.
 *       It's especially useful when the capture has already been triggered by an external event and you want to read the captured value.
 * @note This function is allowed to run within ISR context
 * @note If `CONFIG_GPTIMER_CTRL_FUNC_IN_IRAM` is enabled, this function will be placed in the IRAM by linker,
 *       makes it possible to execute even when the Flash Cache is disabled.
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @param[out] value Returned captured count value
 * @return
 *      - ESP_OK: Get GPTimer captured count value successfully
 *      - ESP_ERR_INVALID_ARG: Get GPTimer captured count value failed because of invalid argument
 *      - ESP_FAIL: Get GPTimer captured count value failed because of other error
 */
//go:linkname GptimerGetCapturedCount C.gptimer_get_captured_count
func GptimerGetCapturedCount(timer GptimerHandleT, value *c.Uint64T) EspErrT

/**
 * @brief Group of supported GPTimer callbacks
 * @note The callbacks are all running under ISR environment
 * @note When CONFIG_GPTIMER_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 */

type GptimerEventCallbacksT struct {
	OnAlarm GptimerAlarmCbT
}

/**
 * @brief Set callbacks for GPTimer
 *
 * @note User registered callbacks are expected to be runnable within ISR context
 * @note The first call to this function needs to be before the call to `gptimer_enable`
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Set event callbacks failed because the timer is not in init state
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname GptimerRegisterEventCallbacks C.gptimer_register_event_callbacks
func GptimerRegisterEventCallbacks(timer GptimerHandleT, cbs *GptimerEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief General Purpose Timer alarm configuration
 */

type GptimerAlarmConfigT struct {
	AlarmCount  c.Uint64T
	ReloadCount c.Uint64T
	Flags       struct {
		AutoReloadOnAlarm c.Uint32T
	}
}

/**
 * @brief Set alarm event actions for GPTimer.
 *
 * @note This function is allowed to run within ISR context, so you can update new alarm action immediately in any ISR callback.
 * @note If `CONFIG_GPTIMER_CTRL_FUNC_IN_IRAM` is enabled, this function will be placed in the IRAM by linker,
 *       makes it possible to execute even when the Flash Cache is disabled.
 *       In this case, please also ensure the `gptimer_alarm_config_t` instance is placed in the static data section
 *       instead of in the read-only data section. e.g.: `static gptimer_alarm_config_t alarm_config = { ... };`
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @param[in] config Alarm configuration, especially, set config to NULL means disabling the alarm function
 * @return
 *      - ESP_OK: Set alarm action for GPTimer successfully
 *      - ESP_ERR_INVALID_ARG: Set alarm action for GPTimer failed because of invalid argument
 *      - ESP_FAIL: Set alarm action for GPTimer failed because of other error
 */
//go:linkname GptimerSetAlarmAction C.gptimer_set_alarm_action
func GptimerSetAlarmAction(timer GptimerHandleT, config *GptimerAlarmConfigT) EspErrT

/**
 * @brief Enable GPTimer
 *
 * @note This function will transit the timer state from "init" to "enable".
 * @note This function will enable the interrupt service, if it's lazy installed in `gptimer_register_event_callbacks`.
 * @note This function will acquire a PM lock, if a specific source clock (e.g. APB) is selected in the `gptimer_config_t`, while `CONFIG_PM_ENABLE` is enabled.
 * @note Enable a timer doesn't mean to start it. See also `gptimer_start` for how to make the timer start counting.
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @return
 *      - ESP_OK: Enable GPTimer successfully
 *      - ESP_ERR_INVALID_ARG: Enable GPTimer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Enable GPTimer failed because the timer is already enabled
 *      - ESP_FAIL: Enable GPTimer failed because of other error
 */
//go:linkname GptimerEnable C.gptimer_enable
func GptimerEnable(timer GptimerHandleT) EspErrT

/**
 * @brief Disable GPTimer
 *
 * @note This function will transit the timer state from "enable" to "init".
 * @note This function will disable the interrupt service if it's installed.
 * @note This function will release the PM lock if it's acquired in the `gptimer_enable`.
 * @note Disable a timer doesn't mean to stop it. See also `gptimer_stop` for how to make the timer stop counting.
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @return
 *      - ESP_OK: Disable GPTimer successfully
 *      - ESP_ERR_INVALID_ARG: Disable GPTimer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Disable GPTimer failed because the timer is not enabled yet
 *      - ESP_FAIL: Disable GPTimer failed because of other error
 */
//go:linkname GptimerDisable C.gptimer_disable
func GptimerDisable(timer GptimerHandleT) EspErrT

/**
 * @brief Start GPTimer (internal counter starts counting)
 *
 * @note This function will transit the timer state from "enable" to "run".
 * @note This function is allowed to run within ISR context
 * @note If `CONFIG_GPTIMER_CTRL_FUNC_IN_IRAM` is enabled, this function will be placed in the IRAM by linker,
 *       makes it possible to execute even when the Flash Cache is disabled.
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @return
 *      - ESP_OK: Start GPTimer successfully
 *      - ESP_ERR_INVALID_ARG: Start GPTimer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Start GPTimer failed because the timer is not enabled or is already in running
 *      - ESP_FAIL: Start GPTimer failed because of other error
 */
//go:linkname GptimerStart C.gptimer_start
func GptimerStart(timer GptimerHandleT) EspErrT

/**
 * @brief Stop GPTimer (internal counter stops counting)
 *
 * @note This function will transit the timer state from "run" to "enable".
 * @note This function is allowed to run within ISR context
 * @note If `CONFIG_GPTIMER_CTRL_FUNC_IN_IRAM` is enabled, this function will be placed in the IRAM by linker,
 *       makes it possible to execute even when the Flash Cache is disabled.
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @return
 *      - ESP_OK: Stop GPTimer successfully
 *      - ESP_ERR_INVALID_ARG: Stop GPTimer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Stop GPTimer failed because the timer is not in running.
 *      - ESP_FAIL: Stop GPTimer failed because of other error
 */
//go:linkname GptimerStop C.gptimer_stop
func GptimerStop(timer GptimerHandleT) EspErrT

/**
 * @brief Get GPTimer interrupt handle
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer()`
 * @param[out] ret_intr_handle Timer's internal interrupt handle
 * @return
 *      - ESP_OK: Get GPTimer interrupt handle successfully
 *      - ESP_ERR_INVALID_ARG: Get GPTimer interrupt handle failed because of invalid argument
 *      - ESP_FAIL: Get GPTimer interrupt handle failed because of other error
 */
//go:linkname GptimerGetIntrHandle C.gptimer_get_intr_handle
func GptimerGetIntrHandle(timer GptimerHandleT, ret_intr_handle *IntrHandleT) EspErrT

/**
 * @brief Get GPTimer power management lock
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer()`
 * @param[out] ret_pm_lock Timer's internal power management lock
 * @return
 *      - ESP_OK: Get GPTimer power management lock successfully
 *      - ESP_ERR_INVALID_ARG: Get GPTimer power management lock failed because of invalid argument
 *      - ESP_FAIL: Get GPTimer power management lock failed because of other error
 */
//go:linkname GptimerGetPmLock C.gptimer_get_pm_lock
func GptimerGetPmLock(timer GptimerHandleT, ret_pm_lock *EspPmLockHandleT) EspErrT

/**
 * @brief Get the group_id from the timer handle
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer()`
 * @return
 *      - ESP_OK: Get GPTimer group_id from handler successfully
 *      - ESP_ERR_INVALID_ARG: Get GPTimer group_id failed because of invalid argument
 */
//go:linkname GptimerGetGroupId C.gptimer_get_group_id
func GptimerGetGroupId(timer GptimerHandleT, group_id *c.Int) EspErrT
