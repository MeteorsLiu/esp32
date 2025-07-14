package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configure_register */
/** Type of clk_en register
 *  need_des
 */

type LpperiRev00ClkEnRegT struct {
	Val c.Uint32T
}

/** Type of reset_en register
 *  need_des
 */

type LpperiRev00ResetEnRegT struct {
	Val c.Uint32T
}

/** Type of rng_data register
 *  need_des
 */

type LpperiRev00RngDataRegT struct {
	Val c.Uint32T
}

/** Type of cpu register
 *  need_des
 */

type LpperiRev00CpuRegT struct {
	Val c.Uint32T
}

/** Type of bus_timeout register
 *  need_des
 */

type LpperiRev00BusTimeoutRegT struct {
	Val c.Uint32T
}

/** Type of bus_timeout_addr register
 *  need_des
 */

type LpperiRev00BusTimeoutAddrRegT struct {
	Val c.Uint32T
}

/** Type of bus_timeout_uid register
 *  need_des
 */

type LpperiRev00BusTimeoutUidRegT struct {
	Val c.Uint32T
}

/** Type of mem_ctrl register
 *  need_des
 */

type LpperiRev00MemCtrlRegT struct {
	Val c.Uint32T
}

/** Type of interrupt_source register
 *  need_des
 */

type LpperiRev00InterruptSourceRegT struct {
	Val c.Uint32T
}

/** Type of debug_sel0 register
 *  need des
 */

type LpperiRev00DebugSel0RegT struct {
	Val c.Uint32T
}

/** Type of debug_sel1 register
 *  need des
 */

type LpperiRev00DebugSel1RegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  need_des
 */

type LpperiRev00DateRegT struct {
	Val c.Uint32T
}

type LpperiRev00DevT struct {
	ClkEn           LpperiRev00ClkEnRegT
	ResetEn         LpperiRev00ResetEnRegT
	RngData         LpperiRev00RngDataRegT
	Cpu             LpperiRev00CpuRegT
	BusTimeout      LpperiRev00BusTimeoutRegT
	BusTimeoutAddr  LpperiRev00BusTimeoutAddrRegT
	BusTimeoutUid   LpperiRev00BusTimeoutUidRegT
	MemCtrl         LpperiRev00MemCtrlRegT
	InterruptSource LpperiRev00InterruptSourceRegT
	DebugSel0       LpperiRev00DebugSel0RegT
	DebugSel1       LpperiRev00DebugSel1RegT
	Reserved02c     [244]c.Uint32T
	Date            LpperiRev00DateRegT
}
