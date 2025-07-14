package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Read the counter value of hardware timer.
 *
 * @param group_num Timer group, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param timer_val Pointer to accept timer counter value.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerGetCounterValue C.timer_get_counter_value
func (recv_ TimerGroupT) TimerGetCounterValue(timer_num TimerIdxT, timer_val *c.Uint64T) EspErrT {
	return 0
}

/**
 * @brief Read the counter value of hardware timer, in unit of a given scale.
 *
 * @param group_num Timer group, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param time Pointer, type of double*, to accept timer counter value, in seconds.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerGetCounterTimeSec C.timer_get_counter_time_sec
func (recv_ TimerGroupT) TimerGetCounterTimeSec(timer_num TimerIdxT, time *c.Double) EspErrT {
	return 0
}

/**
 * @brief Set counter value to hardware timer.
 *
 * @param group_num Timer group, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param load_val Counter value to write to the hardware timer.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerSetCounterValue C.timer_set_counter_value
func (recv_ TimerGroupT) TimerSetCounterValue(timer_num TimerIdxT, load_val c.Uint64T) EspErrT {
	return 0
}

/**
 * @brief Start the counter of hardware timer.
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerStart C.timer_start
func (recv_ TimerGroupT) TimerStart(timer_num TimerIdxT) EspErrT {
	return 0
}

/**
 * @brief Pause the counter of hardware timer.
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerPause C.timer_pause
func (recv_ TimerGroupT) TimerPause(timer_num TimerIdxT) EspErrT {
	return 0
}

/**
 * @brief Set counting mode for hardware timer.
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param counter_dir Counting direction of timer, count-up or count-down
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerSetCounterMode C.timer_set_counter_mode
func (recv_ TimerGroupT) TimerSetCounterMode(timer_num TimerIdxT, counter_dir TimerCountDirT) EspErrT {
	return 0
}

/**
 * @brief Enable or disable counter reload function when alarm event occurs.
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param reload Counter reload mode.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerSetAutoReload C.timer_set_auto_reload
func (recv_ TimerGroupT) TimerSetAutoReload(timer_num TimerIdxT, reload TimerAutoreloadT) EspErrT {
	return 0
}

/**
 * @brief Set hardware divider of the source clock to the timer group.
 * By default, the source clock is APB clock running at 80 MHz.
 * For more information, please check Chapter Reset and Clock in Chip Technical Reference Manual.
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param divider Timer clock divider value. The divider's range is from from 2 to 65536.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerSetDivider C.timer_set_divider
func (recv_ TimerGroupT) TimerSetDivider(timer_num TimerIdxT, divider c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Set timer alarm value.
 *
 * @param group_num Timer group, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param alarm_value A 64-bit value to set the alarm value.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerSetAlarmValue C.timer_set_alarm_value
func (recv_ TimerGroupT) TimerSetAlarmValue(timer_num TimerIdxT, alarm_value c.Uint64T) EspErrT {
	return 0
}

/**
 * @brief Get timer alarm value.
 *
 * @param group_num Timer group, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param alarm_value Pointer of A 64-bit value to accept the alarm value.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerGetAlarmValue C.timer_get_alarm_value
func (recv_ TimerGroupT) TimerGetAlarmValue(timer_num TimerIdxT, alarm_value *c.Uint64T) EspErrT {
	return 0
}

/**
 * @brief Enable or disable generation of timer alarm events.
 *
 * @param group_num Timer group, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param alarm_en To enable or disable timer alarm function.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerSetAlarm C.timer_set_alarm
func (recv_ TimerGroupT) TimerSetAlarm(timer_num TimerIdxT, alarm_en TimerAlarmT) EspErrT {
	return 0
}

/**
 * @brief Add ISR handle callback for the corresponding timer.
 *
 * @param group_num Timer group number
 * @param timer_num Timer index of timer group
 * @param isr_handler Interrupt handler function, it is a callback function.
 * @param arg Parameter for handler function
 * @param intr_alloc_flags Flags used to allocate the interrupt. One or multiple (ORred)
 *        ESP_INTR_FLAG_* values. See esp_intr_alloc.h for more info.
 *
 * @note This ISR handler will be called from an ISR.
 *       This ISR handler do not need to handle interrupt status, and should be kept short.
 *       If you want to realize some specific applications or write the whole ISR, you can
 *       call timer_isr_register(...) to register ISR.
 *
 *       The callback should return a bool value to determine whether need to do YIELD at
 *       the end of the ISR.
 *
 *       If the intr_alloc_flags value ESP_INTR_FLAG_IRAM is set,
 *       the handler function must be declared with IRAM_ATTR attribute
 *       and can only call functions in IRAM or ROM. It cannot call other timer APIs.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerIsrCallbackAdd C.timer_isr_callback_add
func (recv_ TimerGroupT) TimerIsrCallbackAdd(timer_num TimerIdxT, isr_handler TimerIsrT, arg c.Pointer, intr_alloc_flags c.Int) EspErrT {
	return 0
}

/**
 * @brief Remove ISR handle callback for the corresponding timer.
 *
 * @param group_num Timer group number
 * @param timer_num Timer index of timer group
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerIsrCallbackRemove C.timer_isr_callback_remove
func (recv_ TimerGroupT) TimerIsrCallbackRemove(timer_num TimerIdxT) EspErrT {
	return 0
}

/**
 * @brief Register Timer interrupt handler, the handler is an ISR.
 *        The handler will be attached to the same CPU core that this function is running on.
 *
 * @param group_num Timer group number
 * @param timer_num Timer index of timer group
 * @param fn Interrupt handler function.
 * @param arg Parameter for handler function
 * @param intr_alloc_flags Flags used to allocate the interrupt. One or multiple (ORred)
 *        ESP_INTR_FLAG_* values. See esp_intr_alloc.h for more info.
 * @param handle Pointer to return handle. If non-NULL, a handle for the interrupt will
 *        be returned here.
 *
 * @note If use this function to register ISR, you need to write the whole ISR.
 *       In the interrupt handler, you need to call timer_spinlock_take(..) before
 *       your handling, and call timer_spinlock_give(...) after your handling.
 *
 *       If the intr_alloc_flags value ESP_INTR_FLAG_IRAM is set,
 *       the handler function must be declared with IRAM_ATTR attribute
 *       and can only call functions in IRAM or ROM. It cannot call other timer APIs.
 *       Use direct register access to configure timers from inside the ISR in this case.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerIsrRegister C.timer_isr_register
func (recv_ TimerGroupT) TimerIsrRegister(timer_num TimerIdxT, fn func(c.Pointer), arg c.Pointer, intr_alloc_flags c.Int, handle *TimerIsrHandleT) EspErrT {
	return 0
}

/** @brief Initializes and configure the timer.
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param config Pointer to timer initialization parameters.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerInit C.timer_init
func (recv_ TimerGroupT) TimerInit(timer_num TimerIdxT, config *TimerConfigT) EspErrT {
	return 0
}

/** @brief Deinitializes the timer.
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerDeinit C.timer_deinit
func (recv_ TimerGroupT) TimerDeinit(timer_num TimerIdxT) EspErrT {
	return 0
}

/** @brief Get timer configure value.
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index, 0 for hw_timer[0] & 1 for hw_timer[1]
 * @param config Pointer of struct to accept timer parameters.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerGetConfig C.timer_get_config
func (recv_ TimerGroupT) TimerGetConfig(timer_num TimerIdxT, config *TimerConfigT) EspErrT {
	return 0
}

/** @brief Enable timer group interrupt, by enable mask
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param intr_mask Timer interrupt enable mask.
 *          - TIMER_INTR_T0: t0 interrupt
 *          - TIMER_INTR_T1: t1 interrupt
 *          - TIMER_INTR_WDT: watchdog interrupt
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerGroupIntrEnable C.timer_group_intr_enable
func (recv_ TimerGroupT) TimerGroupIntrEnable(intr_mask TimerIntrT) EspErrT {
	return 0
}

/** @brief Disable timer group interrupt, by disable mask
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param intr_mask Timer interrupt disable mask.
 *          - TIMER_INTR_T0: t0 interrupt
 *          - TIMER_INTR_T1: t1 interrupt
 *          - TIMER_INTR_WDT: watchdog interrupt
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerGroupIntrDisable C.timer_group_intr_disable
func (recv_ TimerGroupT) TimerGroupIntrDisable(intr_mask TimerIntrT) EspErrT {
	return 0
}

/** @brief Enable timer interrupt
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerEnableIntr C.timer_enable_intr
func (recv_ TimerGroupT) TimerEnableIntr(timer_num TimerIdxT) EspErrT {
	return 0
}

/** @brief Disable timer interrupt
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link TimerGroupT.TimerDisableIntr C.timer_disable_intr
func (recv_ TimerGroupT) TimerDisableIntr(timer_num TimerIdxT) EspErrT {
	return 0
}

/** @brief Clear timer interrupt status, just used in ISR
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index.
 *
 */
// llgo:link TimerGroupT.TimerGroupClrIntrStatusInIsr C.timer_group_clr_intr_status_in_isr
func (recv_ TimerGroupT) TimerGroupClrIntrStatusInIsr(timer_num TimerIdxT) {
}

/** @brief Enable alarm interrupt, just used in ISR
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index.
 *
 */
// llgo:link TimerGroupT.TimerGroupEnableAlarmInIsr C.timer_group_enable_alarm_in_isr
func (recv_ TimerGroupT) TimerGroupEnableAlarmInIsr(timer_num TimerIdxT) {
}

/** @brief Get the current counter value, just used in ISR
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index.
 *
 * @return
 *     - Counter value
 */
// llgo:link TimerGroupT.TimerGroupGetCounterValueInIsr C.timer_group_get_counter_value_in_isr
func (recv_ TimerGroupT) TimerGroupGetCounterValueInIsr(timer_num TimerIdxT) c.Uint64T {
	return 0
}

/** @brief Set the alarm threshold for the timer, just used in ISR
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index.
 * @param alarm_val Alarm threshold.
 *
 */
// llgo:link TimerGroupT.TimerGroupSetAlarmValueInIsr C.timer_group_set_alarm_value_in_isr
func (recv_ TimerGroupT) TimerGroupSetAlarmValueInIsr(timer_num TimerIdxT, alarm_val c.Uint64T) {
}

/** @brief Enable/disable a counter, just used in ISR
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index.
 * @param counter_en Enable/disable.
 *
 */
// llgo:link TimerGroupT.TimerGroupSetCounterEnableInIsr C.timer_group_set_counter_enable_in_isr
func (recv_ TimerGroupT) TimerGroupSetCounterEnableInIsr(timer_num TimerIdxT, counter_en TimerStartT) {
}

/** @brief Get interrupt status, just used in ISR
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 *
 * @return
 *     - Interrupt status
 */
// llgo:link TimerGroupT.TimerGroupGetIntrStatusInIsr C.timer_group_get_intr_status_in_isr
func (recv_ TimerGroupT) TimerGroupGetIntrStatusInIsr() c.Uint32T {
	return 0
}

/** @brief Get auto reload enable status, just used in ISR
 *
 * @param group_num Timer group number, 0 for TIMERG0 or 1 for TIMERG1
 * @param timer_num Timer index
 *
 * @return
 *     - True Auto reload enabled
 *     - False Auto reload disabled
 */
// llgo:link TimerGroupT.TimerGroupGetAutoReloadInIsr C.timer_group_get_auto_reload_in_isr
func (recv_ TimerGroupT) TimerGroupGetAutoReloadInIsr(timer_num TimerIdxT) bool {
	return false
}
