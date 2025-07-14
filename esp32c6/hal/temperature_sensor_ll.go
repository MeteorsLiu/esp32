package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TemperatureSensorLlWakeupModeT c.Int

const (
	TEMPERATURE_SENSOR_LL_WAKE_ABSOLUTE TemperatureSensorLlWakeupModeT = 0
	TEMPERATURE_SENSOR_LL_WAKE_DELTA    TemperatureSensorLlWakeupModeT = 1
)
