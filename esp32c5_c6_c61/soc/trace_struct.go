package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Trace memory configuration registers */
/** Type of mem_start_addr register
 *  mem start addr
 */

type TraceMemStartAddrRegT struct {
	Val c.Uint32T
}

/** Type of mem_end_addr register
 *  mem end addr
 */

type TraceMemEndAddrRegT struct {
	Val c.Uint32T
}

/** Type of mem_current_addr register
 *  mem current addr
 */

type TraceMemCurrentAddrRegT struct {
	Val c.Uint32T
}

/** Type of mem_addr_update register
 *  mem addr update
 */

type TraceMemAddrUpdateRegT struct {
	Val c.Uint32T
}

/** Group: Trace fifo status register */
/** Type of fifo_status register
 *  fifo status register
 */

type TraceFifoStatusRegT struct {
	Val c.Uint32T
}

/** Group: Trace interrupt configuration registers */
/** Type of intr_ena register
 *  interrupt enable register
 */

type TraceIntrEnaRegT struct {
	Val c.Uint32T
}

/** Type of intr_raw register
 *  interrupt status register
 */

type TraceIntrRawRegT struct {
	Val c.Uint32T
}

/** Type of intr_clr register
 *  interrupt clear register
 */

type TraceIntrClrRegT struct {
	Val c.Uint32T
}

/** Group: Trace configuration register */
/** Type of trigger register
 *  trigger register
 */

type TraceTriggerRegT struct {
	Val c.Uint32T
}

/** Type of resync_prolonged register
 *  resync configuration register
 */

type TraceResyncProlongedRegT struct {
	Val c.Uint32T
}

/** Group: Clock Gate Control and configuration register */
/** Type of clock_gate register
 *  Clock gate control register
 */

type TraceClockGateRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version control register
 */

type TraceDateRegT struct {
	Val c.Uint32T
}

type TraceDevT struct {
	MemStartAddr    TraceMemStartAddrRegT
	MemEndAddr      TraceMemEndAddrRegT
	MemCurrentAddr  TraceMemCurrentAddrRegT
	MemAddrUpdate   TraceMemAddrUpdateRegT
	FifoStatus      TraceFifoStatusRegT
	IntrEna         TraceIntrEnaRegT
	IntrRaw         TraceIntrRawRegT
	IntrClr         TraceIntrClrRegT
	Trigger         TraceTriggerRegT
	ResyncProlonged TraceResyncProlongedRegT
	ClockGate       TraceClockGateRegT
	Reserved02c     [244]c.Uint32T
	Date            TraceDateRegT
}
