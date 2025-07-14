package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Region filter enable register */
/** Type of region_filter_en register
 *  Region filter enable register
 */

type LpApmRegionFilterEnRegT struct {
	Val c.Uint32T
}

/** Group: Region address register */
/** Type of region0_addr_start register
 *  Region address register
 */

type LpApmRegion0AddrStartRegT struct {
	Val c.Uint32T
}

/** Type of region0_addr_end register
 *  Region address register
 */

type LpApmRegion0AddrEndRegT struct {
	Val c.Uint32T
}

/** Type of region1_addr_start register
 *  Region address register
 */

type LpApmRegion1AddrStartRegT struct {
	Val c.Uint32T
}

/** Type of region1_addr_end register
 *  Region address register
 */

type LpApmRegion1AddrEndRegT struct {
	Val c.Uint32T
}

/** Group: Region access authority attribute register */
/** Type of region0_pms_attr register
 *  Region access authority attribute register
 */

type LpApmRegion0PmsAttrRegT struct {
	Val c.Uint32T
}

/** Type of region1_pms_attr register
 *  Region access authority attribute register
 */

type LpApmRegion1PmsAttrRegT struct {
	Val c.Uint32T
}

/** Group: PMS function control register */
/** Type of func_ctrl register
 *  PMS function control register
 */

type LpApmFuncCtrlRegT struct {
	Val c.Uint32T
}

/** Group: M0 status register */
/** Type of m0_status register
 *  M0 status register
 */

type LpApmM0StatusRegT struct {
	Val c.Uint32T
}

/** Group: M0 status clear register */
/** Type of m0_status_clr register
 *  M0 status clear register
 */

type LpApmM0StatusClrRegT struct {
	Val c.Uint32T
}

/** Group: M0 exception_info0 register */
/** Type of m0_exception_info0 register
 *  M0 exception_info0 register
 */

type LpApmM0ExceptionInfo0RegT struct {
	Val c.Uint32T
}

/** Group: M0 exception_info1 register */
/** Type of m0_exception_info1 register
 *  M0 exception_info1 register
 */

type LpApmM0ExceptionInfo1RegT struct {
	Val c.Uint32T
}

/** Group: APM interrupt enable register */
/** Type of int_en register
 *  APM interrupt enable register
 */

type LpApmIntEnRegT struct {
	Val c.Uint32T
}

/** Group: clock gating register */
/** Type of clock_gate register
 *  clock gating register
 */

type LpApmClockGateRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version register
 */

type LpApmDateRegT struct {
	Val c.Uint32T
}

type LpApmDevT struct {
	RegionFilterEn   LpApmRegionFilterEnRegT
	Region0AddrStart LpApmRegion0AddrStartRegT
	Region0AddrEnd   LpApmRegion0AddrEndRegT
	Region0PmsAttr   LpApmRegion0PmsAttrRegT
	Region1AddrStart LpApmRegion1AddrStartRegT
	Region1AddrEnd   LpApmRegion1AddrEndRegT
	Region1PmsAttr   LpApmRegion1PmsAttrRegT
	Reserved01c      [42]c.Uint32T
	FuncCtrl         LpApmFuncCtrlRegT
	M0Status         LpApmM0StatusRegT
	M0StatusClr      LpApmM0StatusClrRegT
	M0ExceptionInfo0 LpApmM0ExceptionInfo0RegT
	M0ExceptionInfo1 LpApmM0ExceptionInfo1RegT
	Reserved0d8      [4]c.Uint32T
	IntEn            LpApmIntEnRegT
	ClockGate        LpApmClockGateRegT
	Reserved0f0      [3]c.Uint32T
	Date             LpApmDateRegT
}
