package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configuration register */
/** Type of bt_select register
 *  GPIO bit select register
 */

type GpioBtSelectRegT struct {
	Val c.Uint32T
}

/** Type of out register
 *  GPIO output register
 */

type GpioOutRegT struct {
	Val c.Uint32T
}

/** Type of out_w1ts register
 *  GPIO output set register
 */

type GpioOutW1tsRegT struct {
	Val c.Uint32T
}

/** Type of out_w1tc register
 *  GPIO output clear register
 */

type GpioOutW1tcRegT struct {
	Val c.Uint32T
}

/** Type of sdio_select register
 *  GPIO sdio select register
 */

type GpioSdioSelectRegT struct {
	Val c.Uint32T
}

/** Type of enable register
 *  GPIO output enable register
 */

type GpioEnableRegT struct {
	Val c.Uint32T
}

/** Type of enable_w1ts register
 *  GPIO output enable set register
 */

type GpioEnableW1tsRegT struct {
	Val c.Uint32T
}

/** Type of enable_w1tc register
 *  GPIO output enable clear register
 */

type GpioEnableW1tcRegT struct {
	Val c.Uint32T
}

/** Type of strap register
 *  pad strapping register
 */

type GpioStrapRegT struct {
	Val c.Uint32T
}

/** Type of in register
 *  GPIO input register
 */

type GpioInRegT struct {
	Val c.Uint32T
}

/** Type of status register
 *  GPIO interrupt status register
 */

type GpioStatusRegT struct {
	Val c.Uint32T
}

/** Type of status_w1ts register
 *  GPIO interrupt status set register
 */

type GpioStatusW1tsRegT struct {
	Val c.Uint32T
}

/** Type of status_w1tc register
 *  GPIO interrupt status clear register
 */

type GpioStatusW1tcRegT struct {
	Val c.Uint32T
}

/** Type of pcpu_int register
 *  GPIO PRO_CPU interrupt status register
 */

type GpioPcpuIntRegT struct {
	Val c.Uint32T
}

/** Type of pcpu_nmi_int register
 *  GPIO PRO_CPU(not shielded) interrupt status register
 */

type GpioPcpuNmiIntRegT struct {
	Val c.Uint32T
}

/** Type of cpusdio_int register
 *  GPIO CPUSDIO interrupt status register
 */

type GpioCpusdioIntRegT struct {
	Val c.Uint32T
}

/** Type of pin0 register
 *  GPIO pin configuration register
 */

type GpioPinRegT struct {
	Val c.Uint32T
}

/** Type of status_next register
 *  GPIO interrupt source register
 */

type GpioStatusNextRegT struct {
	Val c.Uint32T
}

/** Type of in_sel_cfg register
 *  GPIO input function configuration register
 */

type GpioFuncInSelCfgRegT struct {
	Val c.Uint32T
}

/** Type of out_sel_cfg register
 *  GPIO output function select register
 */

type GpioFuncOutSelCfgRegT struct {
	Val c.Uint32T
}

/** Type of clock_gate_reg register
 *  GPIO clock gate register
 */

type GpioClockGateRegRegT struct {
	Val c.Uint32T
}

/** Type of reg_date_reg register
 *  GPIO version register
 */

type GpioRegDateRegRegT struct {
	Val c.Uint32T
}

type GpioDevT struct {
	BtSelect      GpioBtSelectRegT
	Out           GpioOutRegT
	OutW1ts       GpioOutW1tsRegT
	OutW1tc       GpioOutW1tcRegT
	Reserved010   [3]c.Uint32T
	SdioSelect    GpioSdioSelectRegT
	Enable        GpioEnableRegT
	EnableW1ts    GpioEnableW1tsRegT
	EnableW1tc    GpioEnableW1tcRegT
	Reserved02c   [3]c.Uint32T
	Strap         GpioStrapRegT
	In            GpioInRegT
	Reserved040   c.Uint32T
	Status        GpioStatusRegT
	StatusW1ts    GpioStatusW1tsRegT
	StatusW1tc    GpioStatusW1tcRegT
	Reserved050   [3]c.Uint32T
	PcpuInt       GpioPcpuIntRegT
	PcpuNmiInt    GpioPcpuNmiIntRegT
	CpusdioInt    GpioCpusdioIntRegT
	Reserved068   [3]c.Uint32T
	Pin           [25]GpioPinRegT
	Reserved0d8   [29]c.Uint32T
	StatusNext    GpioStatusNextRegT
	Reserved150   c.Uint32T
	FuncInSelCfg  [128]GpioFuncInSelCfgRegT
	Reserved354   [128]c.Uint32T
	FuncOutSelCfg [25]GpioFuncOutSelCfgRegT
	Reserved5b8   [29]c.Uint32T
	ClockGateReg  GpioClockGateRegRegT
	Reserved630   [51]c.Uint32T
	RegDateReg    GpioRegDateRegRegT
}
