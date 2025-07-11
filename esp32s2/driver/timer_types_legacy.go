package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TimerGroupT c.Int

const (
	TIMER_GROUP_0   TimerGroupT = 0
	TIMER_GROUP_1   TimerGroupT = 1
	TIMER_GROUP_MAX TimerGroupT = 2
)

type TimerIdxT c.Int

const (
	TIMER_0   TimerIdxT = 0
	TIMER_1   TimerIdxT = 1
	TIMER_MAX TimerIdxT = 2
)

type TimerIntrT c.Int

const (
	TIMER_INTR_T0   TimerIntrT = 1
	TIMER_INTR_T1   TimerIntrT = 2
	TIMER_INTR_WDT  TimerIntrT = 4
	TIMER_INTR_NONE TimerIntrT = 0
)

type TimerCountDirT c.Int

const (
	TIMER_COUNT_DOWN TimerCountDirT = 0
	TIMER_COUNT_UP   TimerCountDirT = 1
	TIMER_COUNT_MAX  TimerCountDirT = 2
)

type TimerStartT c.Int

const (
	TIMER_PAUSE TimerStartT = 0
	TIMER_START TimerStartT = 1
)

type TimerAlarmT c.Int

const (
	TIMER_ALARM_DIS TimerAlarmT = 0
	TIMER_ALARM_EN  TimerAlarmT = 1
	TIMER_ALARM_MAX TimerAlarmT = 2
)

type TimerIntrModeT c.Int

const (
	TIMER_INTR_LEVEL TimerIntrModeT = 0
	TIMER_INTR_MAX   TimerIntrModeT = 1
)

type TimerAutoreloadT c.Int

const (
	TIMER_AUTORELOAD_DIS TimerAutoreloadT = 0
	TIMER_AUTORELOAD_EN  TimerAutoreloadT = 1
	TIMER_AUTORELOAD_MAX TimerAutoreloadT = 2
)

type TimerSrcClkT SocPeriphTgClkSrcLegacyT

// llgo:type C
type TimerIsrT func(c.Pointer) bool
type TimerIsrHandleT IntrHandleT

/**
 * @brief Timer configurations
 */

type TimerConfigT struct {
	AlarmEn    TimerAlarmT
	CounterEn  TimerStartT
	IntrType   TimerIntrModeT
	CounterDir TimerCountDirT
	AutoReload TimerAutoreloadT
	ClkSrc     TimerSrcClkT
	Divider    c.Uint32T
}
