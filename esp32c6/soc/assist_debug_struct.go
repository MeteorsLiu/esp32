package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: monitor configuration registers */
/** Type of core_0_intr_ena register
 *  core0 monitor enable configuration register
 */

type AssistDebugCore0IntrEnaRegT struct {
	Val c.Uint32T
}

/** Type of core_0_area_dram0_0_min register
 *  core0 dram0 region0 addr configuration register
 */

type AssistDebugCore0AreaDram00MinRegT struct {
	Val c.Uint32T
}

/** Type of core_0_area_dram0_0_max register
 *  core0 dram0 region0 addr configuration register
 */

type AssistDebugCore0AreaDram00MaxRegT struct {
	Val c.Uint32T
}

/** Type of core_0_area_dram0_1_min register
 *  core0 dram0 region1 addr configuration register
 */

type AssistDebugCore0AreaDram01MinRegT struct {
	Val c.Uint32T
}

/** Type of core_0_area_dram0_1_max register
 *  core0 dram0 region1 addr configuration register
 */

type AssistDebugCore0AreaDram01MaxRegT struct {
	Val c.Uint32T
}

/** Type of core_0_area_pif_0_min register
 *  core0 PIF region0 addr configuration register
 */

type AssistDebugCore0AreaPif0MinRegT struct {
	Val c.Uint32T
}

/** Type of core_0_area_pif_0_max register
 *  core0 PIF region0 addr configuration register
 */

type AssistDebugCore0AreaPif0MaxRegT struct {
	Val c.Uint32T
}

/** Type of core_0_area_pif_1_min register
 *  core0 PIF region1 addr configuration register
 */

type AssistDebugCore0AreaPif1MinRegT struct {
	Val c.Uint32T
}

/** Type of core_0_area_pif_1_max register
 *  core0 PIF region1 addr configuration register
 */

type AssistDebugCore0AreaPif1MaxRegT struct {
	Val c.Uint32T
}

/** Type of core_0_area_pc register
 *  core0 area pc status register
 */

type AssistDebugCore0AreaPcRegT struct {
	Val c.Uint32T
}

/** Type of core_0_area_sp register
 *  core0 area sp status register
 */

type AssistDebugCore0AreaSpRegT struct {
	Val c.Uint32T
}

/** Type of core_0_sp_min register
 *  stack min value
 */

type AssistDebugCore0SpMinRegT struct {
	Val c.Uint32T
}

/** Type of core_0_sp_max register
 *  stack max value
 */

type AssistDebugCore0SpMaxRegT struct {
	Val c.Uint32T
}

/** Type of core_0_sp_pc register
 *  stack monitor pc status register
 */

type AssistDebugCore0SpPcRegT struct {
	Val c.Uint32T
}

/** Group: interrupt configuration register */
/** Type of core_0_intr_raw register
 *  core0 monitor interrupt status register
 */

type AssistDebugCore0IntrRawRegT struct {
	Val c.Uint32T
}

/** Type of core_0_intr_rls register
 *  core0 monitor interrupt enable register
 */

type AssistDebugCore0IntrRlsRegT struct {
	Val c.Uint32T
}

/** Type of core_0_intr_clr register
 *  core0 monitor interrupt clr register
 */

type AssistDebugCore0IntrClrRegT struct {
	Val c.Uint32T
}

/** Group: pc recording configuration register */
/** Type of core_0_rcd_en register
 *  record enable configuration register
 */

type AssistDebugCore0RcdEnRegT struct {
	Val c.Uint32T
}

/** Group: pc recording status register */
/** Type of core_0_rcd_pdebugpc register
 *  record status register
 */

type AssistDebugCore0RcdPdebugpcRegT struct {
	Val c.Uint32T
}

/** Type of core_0_rcd_pdebugsp register
 *  record status register
 */

type AssistDebugCore0RcdPdebugspRegT struct {
	Val c.Uint32T
}

/** Group: exception monitor register */
/** Type of core_0_iram0_exception_monitor_0 register
 *  exception monitor status register0
 */

type AssistDebugCore0Iram0ExceptionMonitor0RegT struct {
	Val c.Uint32T
}

/** Type of core_0_iram0_exception_monitor_1 register
 *  exception monitor status register1
 */

type AssistDebugCore0Iram0ExceptionMonitor1RegT struct {
	Val c.Uint32T
}

/** Type of core_0_dram0_exception_monitor_0 register
 *  exception monitor status register2
 */

type AssistDebugCore0Dram0ExceptionMonitor0RegT struct {
	Val c.Uint32T
}

/** Type of core_0_dram0_exception_monitor_1 register
 *  exception monitor status register3
 */

type AssistDebugCore0Dram0ExceptionMonitor1RegT struct {
	Val c.Uint32T
}

/** Type of core_0_dram0_exception_monitor_2 register
 *  exception monitor status register4
 */

type AssistDebugCore0Dram0ExceptionMonitor2RegT struct {
	Val c.Uint32T
}

/** Type of core_0_dram0_exception_monitor_3 register
 *  exception monitor status register5
 */

type AssistDebugCore0Dram0ExceptionMonitor3RegT struct {
	Val c.Uint32T
}

/** Type of core_x_iram0_dram0_exception_monitor_0 register
 *  exception monitor status register6
 */

type AssistDebugCoreXIram0Dram0ExceptionMonitor0RegT struct {
	Val c.Uint32T
}

/** Type of core_x_iram0_dram0_exception_monitor_1 register
 *  exception monitor status register7
 */

type AssistDebugCoreXIram0Dram0ExceptionMonitor1RegT struct {
	Val c.Uint32T
}

/** Group: cpu status registers */
/** Type of core_0_lastpc_before_exception register
 *  cpu status register
 */

type AssistDebugCore0LastpcBeforeExceptionRegT struct {
	Val c.Uint32T
}

/** Type of core_0_debug_mode register
 *  cpu status register
 */

type AssistDebugCore0DebugModeRegT struct {
	Val c.Uint32T
}

/** Group: Configuration Registers */
/** Type of clock_gate register
 *  clock register
 */

type AssistDebugClockGateRegT struct {
	Val c.Uint32T
}

/** Type of date register
 *  version register
 */

type AssistDebugDateRegT struct {
	Val c.Uint32T
}

type AssistDebugDevT struct {
	Core0IntrEna                     AssistDebugCore0IntrEnaRegT
	Core0IntrRaw                     AssistDebugCore0IntrRawRegT
	Core0IntrRls                     AssistDebugCore0IntrRlsRegT
	Core0IntrClr                     AssistDebugCore0IntrClrRegT
	Core0AreaDram00Min               AssistDebugCore0AreaDram00MinRegT
	Core0AreaDram00Max               AssistDebugCore0AreaDram00MaxRegT
	Core0AreaDram01Min               AssistDebugCore0AreaDram01MinRegT
	Core0AreaDram01Max               AssistDebugCore0AreaDram01MaxRegT
	Core0AreaPif0Min                 AssistDebugCore0AreaPif0MinRegT
	Core0AreaPif0Max                 AssistDebugCore0AreaPif0MaxRegT
	Core0AreaPif1Min                 AssistDebugCore0AreaPif1MinRegT
	Core0AreaPif1Max                 AssistDebugCore0AreaPif1MaxRegT
	Core0AreaPc                      AssistDebugCore0AreaPcRegT
	Core0AreaSp                      AssistDebugCore0AreaSpRegT
	Core0SpMin                       AssistDebugCore0SpMinRegT
	Core0SpMax                       AssistDebugCore0SpMaxRegT
	Core0SpPc                        AssistDebugCore0SpPcRegT
	Core0RcdEn                       AssistDebugCore0RcdEnRegT
	Core0RcdPdebugpc                 AssistDebugCore0RcdPdebugpcRegT
	Core0RcdPdebugsp                 AssistDebugCore0RcdPdebugspRegT
	Core0Iram0ExceptionMonitor0      AssistDebugCore0Iram0ExceptionMonitor0RegT
	Core0Iram0ExceptionMonitor1      AssistDebugCore0Iram0ExceptionMonitor1RegT
	Core0Dram0ExceptionMonitor0      AssistDebugCore0Dram0ExceptionMonitor0RegT
	Core0Dram0ExceptionMonitor1      AssistDebugCore0Dram0ExceptionMonitor1RegT
	Core0Dram0ExceptionMonitor2      AssistDebugCore0Dram0ExceptionMonitor2RegT
	Core0Dram0ExceptionMonitor3      AssistDebugCore0Dram0ExceptionMonitor3RegT
	CoreXIram0Dram0ExceptionMonitor0 AssistDebugCoreXIram0Dram0ExceptionMonitor0RegT
	CoreXIram0Dram0ExceptionMonitor1 AssistDebugCoreXIram0Dram0ExceptionMonitor1RegT
	Core0LastpcBeforeException       AssistDebugCore0LastpcBeforeExceptionRegT
	Core0DebugMode                   AssistDebugCore0DebugModeRegT
	ClockGate                        AssistDebugClockGateRegT
	Reserved07c                      [224]c.Uint32T
	Date                             AssistDebugDateRegT
}
