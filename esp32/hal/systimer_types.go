package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SystimerAlarmModeT c.Int

const (
	SYSTIMER_ALARM_MODE_ONESHOT SystimerAlarmModeT = 0
	SYSTIMER_ALARM_MODE_PERIOD  SystimerAlarmModeT = 1
)
