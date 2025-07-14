package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: SYSTEM TIMER CLK CONTROL REGISTER */
/** Type of conf register
 *  Configure system timer clock
 */

type SystimerConfRegT struct {
	Val c.Uint32T
}

/** Group: SYSTEM TIMER UNIT CONTROL AND CONFIGURATION REGISTER */
/** Type of unit_op register
 *  system timer unit value update register
 */

type SystimerUnitOpRegT struct {
	Val c.Uint32T
}

/** Type of unit_load register
 *  system timer unit value high and low load register
 */

type SystimerUnitLoadValRegT struct {
	Hi struct {
		Val c.Uint32T
	}
	Lo struct {
		Val c.Uint32T
	}
}

/** Type of unit_value_hi register
 *  system timer unit value high and low register
 */

type SystimerUnitValueRegT struct {
	Hi struct {
		Val c.Uint32T
	}
	Lo struct {
		Val c.Uint32T
	}
}

/** Type of unit_load register
 *  system timer unit conf sync register
 */

type SystimerUnitLoadRegT struct {
	Val c.Uint32T
}

/** Group: SYSTEM TIMER COMP CONTROL AND CONFIGURATION REGISTER */
/** Type of target register
 *  system timer comp value high and low register
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
 *  system timer comp target mode register
 */

type SystimerTargetConfRegT struct {
	Val c.Uint32T
}

/** Type of comp_load register
 *  system timer comp conf sync register
 */

type SystimerCompLoadRegT struct {
	Val c.Uint32T
}

/** Group: SYSTEM TIMER INTERRUPT REGISTER */
/** Type of int_ena register
 *  systimer interrupt enable register
 */

type SystimerIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_raw register
 *  systimer interrupt raw register
 */

type SystimerIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  systimer interrupt clear register
 */

type SystimerIntClrRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  systimer interrupt status register
 */

type SystimerIntStRegT struct {
	Val c.Uint32T
}

/** Group: SYSTEM TIMER COMP STATUS REGISTER */
/** Type of real_target_hi/lo register
 *  system timer comp actual target value low register
 */

type SystimerRealTargetRegT struct {
	Lo struct {
		Val c.Uint32T
	}
	Hi struct {
		Val c.Uint32T
	}
}

/** Group: VERSION REGISTER */
/** Type of date register
 *  system timer version control register
 */

type SystimerDateRegT struct {
	Val c.Uint32T
}

type SystimerDevT struct {
	Conf        SystimerConfRegT
	UnitOp      [2]SystimerUnitOpRegT
	UnitLoadVal [2]SystimerUnitLoadValRegT
	TargetVal   [3]SystimerTargetValRegT
	TargetConf  [3]SystimerTargetConfRegT
	UnitVal     [2]SystimerUnitValueRegT
	CompLoad    [3]SystimerCompLoadRegT
	UnitLoad    [2]SystimerUnitLoadRegT
	IntEna      SystimerIntEnaRegT
	IntRaw      SystimerIntRawRegT
	IntClr      SystimerIntClrRegT
	IntSt       SystimerIntStRegT
	RealTarget  [3]SystimerRealTargetRegT
	Reserved08c [28]c.Uint32T
	Date        SystimerDateRegT
}
