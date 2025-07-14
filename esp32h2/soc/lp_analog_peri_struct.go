package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configure_register */
/** Type of bod_mode0_cntl register
 *  need_des
 */

type LpAnaBodMode0CntlRegT struct {
	Val c.Uint32T
}

/** Type of bod_mode1_cntl register
 *  need_des
 */

type LpAnaBodMode1CntlRegT struct {
	Val c.Uint32T
}

/** Type of vdd_source_cntl register
 *  need_des
 */

type LpAnaVddSourceCntlRegT struct {
	Val c.Uint32T
}

/** Type of vddbat_bod_cntl register
 *  need_des
 */

type LpAnaVddbatBodCntlRegT struct {
	Val c.Uint32T
}

/** Type of vddbat_charge_cntl register
 *  need_des
 */

type LpAnaVddbatChargeCntlRegT struct {
	Val c.Uint32T
}

/** Type of ck_glitch_cntl register
 *  need_des
 */

type LpAnaCkGlitchCntlRegT struct {
	Val c.Uint32T
}

/** Type of pg_glitch_cntl register
 *  need_des
 */

type LpAnaPgGlitchCntlRegT struct {
	Val c.Uint32T
}

/** Type of fib_enable register
 *  need_des
 */

type LpAnaFibEnableRegT struct {
	Val c.Uint32T
}

/** Type of int_raw register
 *  need_des
 */

type LpAnaIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  need_des
 */

type LpAnaIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  need_des
 */

type LpAnaIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  need_des
 */

type LpAnaIntClrRegT struct {
	Val c.Uint32T
}

/** Type of lp_int_raw register
 *  need_des
 */

type LpAnaLpIntRawRegT struct {
	Val c.Uint32T
}

/** Type of lp_int_st register
 *  need_des
 */

type LpAnaLpIntStRegT struct {
	Val c.Uint32T
}

/** Type of lp_int_ena register
 *  need_des
 */

type LpAnaLpIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of lp_int_clr register
 *  need_des
 */

type LpAnaLpIntClrRegT struct {
	Val c.Uint32T
}

/** Type of date register
 *  need_des
 */

type LpAnaDateRegT struct {
	Val c.Uint32T
}

type LpAnaDevT struct {
	BodMode0Cntl     LpAnaBodMode0CntlRegT
	BodMode1Cntl     LpAnaBodMode1CntlRegT
	VddSourceCntl    LpAnaVddSourceCntlRegT
	VddbatBodCntl    LpAnaVddbatBodCntlRegT
	VddbatChargeCntl LpAnaVddbatChargeCntlRegT
	CkGlitchCntl     LpAnaCkGlitchCntlRegT
	PgGlitchCntl     LpAnaPgGlitchCntlRegT
	FibEnable        LpAnaFibEnableRegT
	IntRaw           LpAnaIntRawRegT
	IntSt            LpAnaIntStRegT
	IntEna           LpAnaIntEnaRegT
	IntClr           LpAnaIntClrRegT
	LpIntRaw         LpAnaLpIntRawRegT
	LpIntSt          LpAnaLpIntStRegT
	LpIntEna         LpAnaLpIntEnaRegT
	LpIntClr         LpAnaLpIntClrRegT
	Reserved040      [239]c.Uint32T
	Date             LpAnaDateRegT
}
