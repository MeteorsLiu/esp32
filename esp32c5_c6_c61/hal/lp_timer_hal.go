package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * @brief set alarm target value
 *
 * @param timer_id timer num of lp_timer, 0 or 1 for esp32c6 and esp32h2
 *
 * @param value when counter reaches alarm value, alarm event will be triggered
 */
//go:linkname LpTimerHalSetAlarmTarget C.lp_timer_hal_set_alarm_target
func LpTimerHalSetAlarmTarget(timer_id c.Uint8T, value c.Uint64T)

/**
 * @brief get current counter value
 */
//go:linkname LpTimerHalGetCycleCount C.lp_timer_hal_get_cycle_count
func LpTimerHalGetCycleCount() c.Uint64T

/**
 * @brief clear alarm interrupt status
 */
//go:linkname LpTimerHalClearAlarmIntrStatus C.lp_timer_hal_clear_alarm_intr_status
func LpTimerHalClearAlarmIntrStatus()

/**
 * @brief clear overflow interrupt status
 */
//go:linkname LpTimerHalClearOverflowIntrStatus C.lp_timer_hal_clear_overflow_intr_status
func LpTimerHalClearOverflowIntrStatus()
