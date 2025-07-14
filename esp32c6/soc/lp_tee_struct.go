package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Tee mode control register */
/** Type of m0_mode_ctrl register
 *  Tee mode control register
 */

type LpTeeM0ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Group: clock gating register */
/** Type of clock_gate register
 *  Clock gating register
 */

type LpTeeClockGateRegT struct {
	Val c.Uint32T
}

/** Group: configure_register */
/** Type of force_acc_hp register
 *  need_des
 */

type LpTeeForceAccHpRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version register
 */

type LpTeeDateRegT struct {
	Val c.Uint32T
}

type LpTeeDevT struct {
	M0ModeCtrl  LpTeeM0ModeCtrlRegT
	ClockGate   LpTeeClockGateRegT
	Reserved008 [34]c.Uint32T
	ForceAccHp  LpTeeForceAccHpRegT
	Reserved094 [26]c.Uint32T
	Date        LpTeeDateRegT
}
