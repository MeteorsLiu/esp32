package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Registers */
/** Type of core0_cpu_int_enable register
 *  register description
 */

type IntpriCore0CpuIntEnableRegT struct {
	Val c.Uint32T
}

/** Type of core0_cpu_int_type register
 *  register description
 */

type IntpriCore0CpuIntTypeRegT struct {
	Val c.Uint32T
}

/** Type of core0_cpu_int_eip_status register
 *  register description
 */

type IntpriCore0CpuIntEipStatusRegT struct {
	Val c.Uint32T
}

/** Type of core0_cpu_int_pri_n register
 *  register description
 */

type IntpriCore0CpuIntPriNRegT struct {
	Val c.Uint32T
}

/** Type of core0_cpu_int_thresh register
 *  register description
 */

type IntpriCore0CpuIntThreshRegT struct {
	Val c.Uint32T
}

/** Type of clock_gate register
 *  register description
 */

type IntpriClockGateRegT struct {
	Val c.Uint32T
}

/** Type of core0_cpu_int_clear register
 *  register description
 */

type IntpriCore0CpuIntClearRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Registers */
/** Type of cpu_intr_from_cpu_0 register
 *  register description
 */

type IntpriCpuIntrFromCpu0RegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_1 register
 *  register description
 */

type IntpriCpuIntrFromCpu1RegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_2 register
 *  register description
 */

type IntpriCpuIntrFromCpu2RegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_3 register
 *  register description
 */

type IntpriCpuIntrFromCpu3RegT struct {
	Val c.Uint32T
}

/** Group: Version Registers */
/** Type of date register
 *  register description
 */

type IntpriDateRegT struct {
	Val c.Uint32T
}

/** Group: Redcy ECO Registers */
/** Type of rnd_eco register
 *  redcy eco register.
 */

type IntpriRndEcoRegT struct {
	Val c.Uint32T
}

/** Type of rnd_eco_low register
 *  redcy eco low register.
 */

type IntpriRndEcoLowRegT struct {
	Val c.Uint32T
}

/** Type of rnd_eco_high register
 *  redcy eco high register.
 */

type IntpriRndEcoHighRegT struct {
	Val c.Uint32T
}

type IntpriDevT struct {
	Core0CpuIntEnable    IntpriCore0CpuIntEnableRegT
	Core0CpuIntType      IntpriCore0CpuIntTypeRegT
	Core0CpuIntEipStatus IntpriCore0CpuIntEipStatusRegT
	Core0CpuIntPri       [32]IntpriCore0CpuIntPriNRegT
	Core0CpuIntThresh    IntpriCore0CpuIntThreshRegT
	CpuIntrFromCpu0      IntpriCpuIntrFromCpu0RegT
	CpuIntrFromCpu1      IntpriCpuIntrFromCpu1RegT
	CpuIntrFromCpu2      IntpriCpuIntrFromCpu2RegT
	CpuIntrFromCpu3      IntpriCpuIntrFromCpu3RegT
	Date                 IntpriDateRegT
	ClockGate            IntpriClockGateRegT
	Core0CpuIntClear     IntpriCore0CpuIntClearRegT
	RndEco               IntpriRndEcoRegT
	RndEcoLow            IntpriRndEcoLowRegT
	Reserved0b4          [210]c.Uint32T
	RndEcoHigh           IntpriRndEcoHighRegT
}
