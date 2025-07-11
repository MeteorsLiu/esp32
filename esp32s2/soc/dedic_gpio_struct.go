package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Configuration registers */
/** Type of gpio_out_drt register
 *  Dedicated GPIO Directive Output register
 */

type DedicGpioOutDrtRegT struct {
	Val c.Uint32T
}

/** Type of gpio_out_msk register
 *  Dedicated GPIO Mask Output register
 */

type DedicGpioOutMskRegT struct {
	Val c.Uint32T
}

/** Type of gpio_out_idv register
 *  Dedicated GPIO Individual Output register
 */

type DedicGpioOutIdvRegT struct {
	Val c.Uint32T
}

/** Type of gpio_out_cpu register
 *  Dedicated GPIO Output Mode Select register
 */

type DedicGpioOutCpuRegT struct {
	Val c.Uint32T
}

/** Type of gpio_in_dly register
 *  Dedicated GPIO Input Delay Configuration register
 */

type DedicGpioInDlyRegT struct {
	Val c.Uint32T
}

/** Type of gpio_intr_rcgn register
 *  Dedicated GPIO Interrupts Generate Mode register
 */

type DedicGpioIntrRcgnRegT struct {
	Val c.Uint32T
}

/** Status registers */
/** Type of gpio_out_scan register
 *  Dedicated GPIO Output Status register
 */

type DedicGpioOutScanRegT struct {
	Val c.Uint32T
}

/** Type of gpio_in_scan register
 *  Dedicated GPIO Input Status register
 */

type DedicGpioInScanRegT struct {
	Val c.Uint32T
}

/** Interrupt registers */
/** Type of gpio_intr_raw register
 *  Raw interrupt status
 */

type DedicGpioIntrRawRegT struct {
	Val c.Uint32T
}

/** Type of gpio_intr_rls register
 *  Interrupt enable bits
 */

type DedicGpioIntrRlsRegT struct {
	Val c.Uint32T
}

/** Type of gpio_intr_st register
 *  Masked interrupt status
 */

type DedicGpioIntrStRegT struct {
	Val c.Uint32T
}

/** Type of gpio_intr_clr register
 *  Interrupt clear bits
 */

type DedicGpioIntrClrRegT struct {
	Val c.Uint32T
}

type DedicDevT struct {
	GpioOutDrt   DedicGpioOutDrtRegT
	GpioOutMsk   DedicGpioOutMskRegT
	GpioOutIdv   DedicGpioOutIdvRegT
	GpioOutScan  DedicGpioOutScanRegT
	GpioOutCpu   DedicGpioOutCpuRegT
	GpioInDly    DedicGpioInDlyRegT
	GpioInScan   DedicGpioInScanRegT
	GpioIntrRcgn DedicGpioIntrRcgnRegT
	GpioIntrRaw  DedicGpioIntrRawRegT
	GpioIntrRls  DedicGpioIntrRlsRegT
	GpioIntrSt   DedicGpioIntrStRegT
	GpioIntrClr  DedicGpioIntrClrRegT
}
