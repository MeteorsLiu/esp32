package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Tee mode control register */
/** Type of m0_mode_ctrl register
 *  Tee mode control register
 */

type TeeM0ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m1_mode_ctrl register
 *  Tee mode control register
 */

type TeeM1ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m2_mode_ctrl register
 *  Tee mode control register
 */

type TeeM2ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m3_mode_ctrl register
 *  Tee mode control register
 */

type TeeM3ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m4_mode_ctrl register
 *  Tee mode control register
 */

type TeeM4ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m5_mode_ctrl register
 *  Tee mode control register
 */

type TeeM5ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m6_mode_ctrl register
 *  Tee mode control register
 */

type TeeM6ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m7_mode_ctrl register
 *  Tee mode control register
 */

type TeeM7ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m8_mode_ctrl register
 *  Tee mode control register
 */

type TeeM8ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m9_mode_ctrl register
 *  Tee mode control register
 */

type TeeM9ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m10_mode_ctrl register
 *  Tee mode control register
 */

type TeeM10ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m11_mode_ctrl register
 *  Tee mode control register
 */

type TeeM11ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m12_mode_ctrl register
 *  Tee mode control register
 */

type TeeM12ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m13_mode_ctrl register
 *  Tee mode control register
 */

type TeeM13ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m14_mode_ctrl register
 *  Tee mode control register
 */

type TeeM14ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m15_mode_ctrl register
 *  Tee mode control register
 */

type TeeM15ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m16_mode_ctrl register
 *  Tee mode control register
 */

type TeeM16ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m17_mode_ctrl register
 *  Tee mode control register
 */

type TeeM17ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m18_mode_ctrl register
 *  Tee mode control register
 */

type TeeM18ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m19_mode_ctrl register
 *  Tee mode control register
 */

type TeeM19ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m20_mode_ctrl register
 *  Tee mode control register
 */

type TeeM20ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m21_mode_ctrl register
 *  Tee mode control register
 */

type TeeM21ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m22_mode_ctrl register
 *  Tee mode control register
 */

type TeeM22ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m23_mode_ctrl register
 *  Tee mode control register
 */

type TeeM23ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m24_mode_ctrl register
 *  Tee mode control register
 */

type TeeM24ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m25_mode_ctrl register
 *  Tee mode control register
 */

type TeeM25ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m26_mode_ctrl register
 *  Tee mode control register
 */

type TeeM26ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m27_mode_ctrl register
 *  Tee mode control register
 */

type TeeM27ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m28_mode_ctrl register
 *  Tee mode control register
 */

type TeeM28ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m29_mode_ctrl register
 *  Tee mode control register
 */

type TeeM29ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m30_mode_ctrl register
 *  Tee mode control register
 */

type TeeM30ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of m31_mode_ctrl register
 *  Tee mode control register
 */

type TeeM31ModeCtrlRegT struct {
	Val c.Uint32T
}

/** Group: clock gating register */
/** Type of clock_gate register
 *  Clock gating register
 */

type TeeClockGateRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version register
 */

type TeeDateRegT struct {
	Val c.Uint32T
}

type TeeDevT struct {
	M0ModeCtrl  TeeM0ModeCtrlRegT
	M1ModeCtrl  TeeM1ModeCtrlRegT
	M2ModeCtrl  TeeM2ModeCtrlRegT
	M3ModeCtrl  TeeM3ModeCtrlRegT
	M4ModeCtrl  TeeM4ModeCtrlRegT
	M5ModeCtrl  TeeM5ModeCtrlRegT
	M6ModeCtrl  TeeM6ModeCtrlRegT
	M7ModeCtrl  TeeM7ModeCtrlRegT
	M8ModeCtrl  TeeM8ModeCtrlRegT
	M9ModeCtrl  TeeM9ModeCtrlRegT
	M10ModeCtrl TeeM10ModeCtrlRegT
	M11ModeCtrl TeeM11ModeCtrlRegT
	M12ModeCtrl TeeM12ModeCtrlRegT
	M13ModeCtrl TeeM13ModeCtrlRegT
	M14ModeCtrl TeeM14ModeCtrlRegT
	M15ModeCtrl TeeM15ModeCtrlRegT
	M16ModeCtrl TeeM16ModeCtrlRegT
	M17ModeCtrl TeeM17ModeCtrlRegT
	M18ModeCtrl TeeM18ModeCtrlRegT
	M19ModeCtrl TeeM19ModeCtrlRegT
	M20ModeCtrl TeeM20ModeCtrlRegT
	M21ModeCtrl TeeM21ModeCtrlRegT
	M22ModeCtrl TeeM22ModeCtrlRegT
	M23ModeCtrl TeeM23ModeCtrlRegT
	M24ModeCtrl TeeM24ModeCtrlRegT
	M25ModeCtrl TeeM25ModeCtrlRegT
	M26ModeCtrl TeeM26ModeCtrlRegT
	M27ModeCtrl TeeM27ModeCtrlRegT
	M28ModeCtrl TeeM28ModeCtrlRegT
	M29ModeCtrl TeeM29ModeCtrlRegT
	M30ModeCtrl TeeM30ModeCtrlRegT
	M31ModeCtrl TeeM31ModeCtrlRegT
	ClockGate   TeeClockGateRegT
	Reserved084 [990]c.Uint32T
	Date        TeeDateRegT
}
