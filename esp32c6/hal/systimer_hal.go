package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SystimerSocHandleT *SystimerDevT

// llgo:type C
type TicksToUsFuncT func(c.Uint64T) c.Uint64T

// llgo:type C
type UsToTicksFuncT func(c.Uint64T) c.Uint64T

/**
 * @brief systimer HAL context structure
 */

type SystimerHalContextT struct {
	Dev       SystimerSocHandleT
	TicksToUs TicksToUsFuncT
	UsToTicks UsToTicksFuncT
}

/**
 * @brief systimer HAL configuration structure
 */

type SystimerHalTickRateOpsT struct {
	TicksToUs TicksToUsFuncT
	UsToTicks UsToTicksFuncT
}
type SystimerClockSourceT SocPeriphSystimerClkSrcT
