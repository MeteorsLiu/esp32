package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** SYSTEM TIMER REGISTER */
/** Type of conf register
 *  Configure system timer clock
 */

type SystimerConfRegT struct {
	Val c.Uint32T
}

/** Type of load register
 *  load value to system timer
 */

type SystimerLoadRegT struct {
	Val c.Uint32T
}

/** Type of load_hi register
 *  High 32-bit load to system timer
 */

type SystimerLoadHiRegT struct {
	Val c.Uint32T
}

/** Type of load_lo register
 *  Low 32-bit load to system timer
 */

type SystimerLoadLoRegT struct {
	Val c.Uint32T
}

/** Type of step register
 *  system timer accumulation step
 */

type SystimerStepRegT struct {
	Val c.Uint32T
}

/** Type of target_val register
 *  System timer target value
 */

type SystimerTargetValRegT struct {
	Hi struct {
		Val c.Uint32T
	}
	Lo struct {
		Val c.Uint32T
	}
}

/** Type of target_conf register
 *  Configure system timer target work mode
 */

type SystimerTargetConfRegT struct {
	Val c.Uint32T
}

/** Type of update register
 *  Read out system timer value
 */

type SystimerUpdateRegT struct {
	Val c.Uint32T
}

/** Type of value_hi register
 *  system timer high 32-bit
 */

type SystimerValueHiRegT struct {
	Val c.Uint32T
}

/** Type of value_lo register
 *  system timer low 32-bit
 */

type SystimerValueLoRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  system timer interrupt enable
 */

type SystimerIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_raw register
 *  system timer interrupt raw
 */

type SystimerIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  system timer interrupt clear
 */

type SystimerIntClrRegT struct {
	Val c.Uint32T
}

/** DATE */
/** Type of date register
 *  system timer register version
 */

type SystimerDateRegT struct {
	Val c.Uint32T
}

type SystimerDevT struct {
	Conf        SystimerConfRegT
	Load        SystimerLoadRegT
	LoadHi      SystimerLoadHiRegT
	LoadLo      SystimerLoadLoRegT
	Step        SystimerStepRegT
	TargetVal   [3]SystimerTargetValRegT
	TargetConf  [3]SystimerTargetConfRegT
	Update      SystimerUpdateRegT
	ValueHi     SystimerValueHiRegT
	ValueLo     SystimerValueLoRegT
	IntEna      SystimerIntEnaRegT
	IntRaw      SystimerIntRawRegT
	IntClr      SystimerIntClrRegT
	Reserved050 c.Uint32T
	Reserved054 c.Uint32T
	Reserved058 c.Uint32T
	Reserved05c c.Uint32T
	Reserved060 c.Uint32T
	Reserved064 c.Uint32T
	Reserved068 c.Uint32T
	Reserved06c c.Uint32T
	Reserved070 c.Uint32T
	Reserved074 c.Uint32T
	Reserved078 c.Uint32T
	Reserved07c c.Uint32T
	Reserved080 c.Uint32T
	Reserved084 c.Uint32T
	Reserved088 c.Uint32T
	Reserved08c c.Uint32T
	Reserved090 c.Uint32T
	Reserved094 c.Uint32T
	Reserved098 c.Uint32T
	Reserved09c c.Uint32T
	Reserved0a0 c.Uint32T
	Reserved0a4 c.Uint32T
	Reserved0a8 c.Uint32T
	Reserved0ac c.Uint32T
	Reserved0b0 c.Uint32T
	Reserved0b4 c.Uint32T
	Reserved0b8 c.Uint32T
	Reserved0bc c.Uint32T
	Reserved0c0 c.Uint32T
	Reserved0c4 c.Uint32T
	Reserved0c8 c.Uint32T
	Reserved0cc c.Uint32T
	Reserved0d0 c.Uint32T
	Reserved0d4 c.Uint32T
	Reserved0d8 c.Uint32T
	Reserved0dc c.Uint32T
	Reserved0e0 c.Uint32T
	Reserved0e4 c.Uint32T
	Reserved0e8 c.Uint32T
	Reserved0ec c.Uint32T
	Reserved0f0 c.Uint32T
	Reserved0f4 c.Uint32T
	Reserved0f8 c.Uint32T
	Date        SystimerDateRegT
}
