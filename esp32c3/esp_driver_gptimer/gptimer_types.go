package esp_driver_gptimer

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GptimerT struct {
	Unused [8]uint8
}
type GptimerHandleT *GptimerT

/**
 * @brief GPTimer alarm event data
 */

type GptimerAlarmEventDataT struct {
	CountValue c.Uint64T
	AlarmValue c.Uint64T
}

// llgo:type C
type GptimerAlarmCbT func(GptimerHandleT, *GptimerAlarmEventDataT, c.Pointer) bool
