package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configure_register */
/** Type of clk_en register
 *  need_des
 */

type LpperiClkEnRegT struct {
	Val c.Uint32T
}

/** Type of reset_en register
 *  need_des
 */

type LpperiResetEnRegT struct {
	Val c.Uint32T
}

/** Type of rng_data register
 *  need_des
 */

type LpperiRngDataRegT struct {
	Val c.Uint32T
}

/** Type of cpu register
 *  need_des
 */

type LpperiCpuRegT struct {
	Val c.Uint32T
}

/** Type of bus_timeout register
 *  need_des
 */

type LpperiBusTimeoutRegT struct {
	Val c.Uint32T
}

/** Type of bus_timeout_addr register
 *  need_des
 */

type LpperiBusTimeoutAddrRegT struct {
	Val c.Uint32T
}

/** Type of bus_timeout_uid register
 *  need_des
 */

type LpperiBusTimeoutUidRegT struct {
	Val c.Uint32T
}

/** Type of mem_ctrl register
 *  need_des
 */

type LpperiMemCtrlRegT struct {
	Val c.Uint32T
}

/** Type of interrupt_source register
 *  need_des
 */

type LpperiInterruptSourceRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  need_des
 */

type LpperiDateRegT struct {
	Val c.Uint32T
}

type LpperiDevT struct {
	ClkEn           LpperiClkEnRegT
	ResetEn         LpperiResetEnRegT
	RngData         LpperiRngDataRegT
	Cpu             LpperiCpuRegT
	BusTimeout      LpperiBusTimeoutRegT
	BusTimeoutAddr  LpperiBusTimeoutAddrRegT
	BusTimeoutUid   LpperiBusTimeoutUidRegT
	MemCtrl         LpperiMemCtrlRegT
	InterruptSource LpperiInterruptSourceRegT
	Reserved024     [246]c.Uint32T
	Date            LpperiDateRegT
}
