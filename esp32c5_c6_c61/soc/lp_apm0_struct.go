package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Region filter enable register */
/** Type of region_filter_en register
 *  Region filter enable register
 */

type LpApm0RegionFilterEnRegT struct {
	Val c.Uint32T
}

/** Group: Region address register */
/** Type of region0_addr_start register
 *  Region address register
 */

type LpApm0Region0AddrStartRegT struct {
	Val c.Uint32T
}

/** Type of region0_addr_end register
 *  Region address register
 */

type LpApm0Region0AddrEndRegT struct {
	Val c.Uint32T
}

/** Type of region1_addr_start register
 *  Region address register
 */

type LpApm0Region1AddrStartRegT struct {
	Val c.Uint32T
}

/** Type of region1_addr_end register
 *  Region address register
 */

type LpApm0Region1AddrEndRegT struct {
	Val c.Uint32T
}

/** Type of region2_addr_start register
 *  Region address register
 */

type LpApm0Region2AddrStartRegT struct {
	Val c.Uint32T
}

/** Type of region2_addr_end register
 *  Region address register
 */

type LpApm0Region2AddrEndRegT struct {
	Val c.Uint32T
}

/** Type of region3_addr_start register
 *  Region address register
 */

type LpApm0Region3AddrStartRegT struct {
	Val c.Uint32T
}

/** Type of region3_addr_end register
 *  Region address register
 */

type LpApm0Region3AddrEndRegT struct {
	Val c.Uint32T
}

/** Group: Region access authority attribute register */
/** Type of region0_pms_attr register
 *  Region access authority attribute register
 */

type LpApm0Region0PmsAttrRegT struct {
	Val c.Uint32T
}

/** Type of region1_pms_attr register
 *  Region access authority attribute register
 */

type LpApm0Region1PmsAttrRegT struct {
	Val c.Uint32T
}

/** Type of region2_pms_attr register
 *  Region access authority attribute register
 */

type LpApm0Region2PmsAttrRegT struct {
	Val c.Uint32T
}

/** Type of region3_pms_attr register
 *  Region access authority attribute register
 */

type LpApm0Region3PmsAttrRegT struct {
	Val c.Uint32T
}

/** Group: PMS function control register */
/** Type of func_ctrl register
 *  PMS function control register
 */

type LpApm0FuncCtrlRegT struct {
	Val c.Uint32T
}

/** Group: M0 status register */
/** Type of m0_status register
 *  M0 status register
 */

type LpApm0M0StatusRegT struct {
	Val c.Uint32T
}

/** Group: M0 status clear register */
/** Type of m0_status_clr register
 *  M0 status clear register
 */

type LpApm0M0StatusClrRegT struct {
	Val c.Uint32T
}

/** Group: M0 exception_info0 register */
/** Type of m0_exception_info0 register
 *  M0 exception_info0 register
 */

type LpApm0M0ExceptionInfo0RegT struct {
	Val c.Uint32T
}

/** Group: M0 exception_info1 register */
/** Type of m0_exception_info1 register
 *  M0 exception_info1 register
 */

type LpApm0M0ExceptionInfo1RegT struct {
	Val c.Uint32T
}

/** Group: APM interrupt enable register */
/** Type of int_en register
 *  APM interrupt enable register
 */

type LpApm0IntEnRegT struct {
	Val c.Uint32T
}

/** Group: clock gating register */
/** Type of clock_gate register
 *  clock gating register
 */

type LpApm0ClockGateRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version register
 */

type LpApm0DateRegT struct {
	Val c.Uint32T
}

type LpApm0DevT struct {
	RegionFilterEn   LpApm0RegionFilterEnRegT
	Region0AddrStart LpApm0Region0AddrStartRegT
	Region0AddrEnd   LpApm0Region0AddrEndRegT
	Region0PmsAttr   LpApm0Region0PmsAttrRegT
	Region1AddrStart LpApm0Region1AddrStartRegT
	Region1AddrEnd   LpApm0Region1AddrEndRegT
	Region1PmsAttr   LpApm0Region1PmsAttrRegT
	Region2AddrStart LpApm0Region2AddrStartRegT
	Region2AddrEnd   LpApm0Region2AddrEndRegT
	Region2PmsAttr   LpApm0Region2PmsAttrRegT
	Region3AddrStart LpApm0Region3AddrStartRegT
	Region3AddrEnd   LpApm0Region3AddrEndRegT
	Region3PmsAttr   LpApm0Region3PmsAttrRegT
	Reserved034      [36]c.Uint32T
	FuncCtrl         LpApm0FuncCtrlRegT
	M0Status         LpApm0M0StatusRegT
	M0StatusClr      LpApm0M0StatusClrRegT
	M0ExceptionInfo0 LpApm0M0ExceptionInfo0RegT
	M0ExceptionInfo1 LpApm0M0ExceptionInfo1RegT
	IntEn            LpApm0IntEnRegT
	ClockGate        LpApm0ClockGateRegT
	Reserved0e0      [455]c.Uint32T
	Date             LpApm0DateRegT
}
