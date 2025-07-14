package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configure_register */
/** Type of lp_clk_conf register
 *  need_des
 */

type LpClkrstLpClkConfRegT struct {
	Val c.Uint32T
}

/** Type of lp_clk_po_en register
 *  need_des
 */

type LpClkrstLpClkPoEnRegT struct {
	Val c.Uint32T
}

/** Type of lp_clk_en register
 *  need_des
 */

type LpClkrstLpClkEnRegT struct {
	Val c.Uint32T
}

/** Type of lp_rst_en register
 *  need_des
 */

type LpClkrstLpRstEnRegT struct {
	Val c.Uint32T
}

/** Type of reset_cause register
 *  need_des
 */

type LpClkrstResetCauseRegT struct {
	Val c.Uint32T
}

/** Type of cpu_reset register
 *  need_des
 */

type LpClkrstCpuResetRegT struct {
	Val c.Uint32T
}

/** Type of fosc_cntl register
 *  need_des
 */

type LpClkrstFoscCntlRegT struct {
	Val c.Uint32T
}

/** Type of rc32k_cntl register
 *  need_des
 */

type LpClkrstRc32kCntlRegT struct {
	Val c.Uint32T
}

/** Type of clk_to_hp register
 *  need_des
 */

type LpClkrstClkToHpRegT struct {
	Val c.Uint32T
}

/** Type of lpmem_force register
 *  need_des
 */

type LpClkrstLpmemForceRegT struct {
	Val c.Uint32T
}

/** Type of lpperi register
 *  need_des
 */

type LpClkrstLpperiRegT struct {
	Val c.Uint32T
}

/** Type of xtal32k register
 *  need_des
 */

type LpClkrstXtal32kRegT struct {
	Val c.Uint32T
}

/** Type of date register
 *  need_des
 */

type LpClkrstDateRegT struct {
	Val c.Uint32T
}

type LpClkrstDevT struct {
	LpClkConf   LpClkrstLpClkConfRegT
	LpClkPoEn   LpClkrstLpClkPoEnRegT
	LpClkEn     LpClkrstLpClkEnRegT
	LpRstEn     LpClkrstLpRstEnRegT
	ResetCause  LpClkrstResetCauseRegT
	CpuReset    LpClkrstCpuResetRegT
	FoscCntl    LpClkrstFoscCntlRegT
	Rc32kCntl   LpClkrstRc32kCntlRegT
	ClkToHp     LpClkrstClkToHpRegT
	LpmemForce  LpClkrstLpmemForceRegT
	Lpperi      LpClkrstLpperiRegT
	Xtal32k     LpClkrstXtal32kRegT
	Reserved030 [243]c.Uint32T
	Date        LpClkrstDateRegT
}
