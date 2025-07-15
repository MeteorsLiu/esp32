package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Type by which software timers are referenced.  For example, a call to
 * xTimerCreate() returns an TimerHandle_t variable that can then be used to
 * reference the subject timer in calls to other software timer API functions
 * (for example, xTimerStart(), xTimerReset(), etc.).
 */

type TmrTimerControl struct {
	Unused [8]uint8
}
type TimerHandleT *TmrTimerControl

// llgo:type C
type TimerCallbackFunctionT func(TimerHandleT)

// llgo:type C
type PendedFunctionT func(c.Pointer, c.Int)
