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
 *  GPIO output register for GPIO0-31
 */

type GpioOutRegT struct {
	Val c.Uint32T
}

/** Type of out_w1ts register
 *  GPIO output set register for GPIO0-31
 */

type GpioOutW1tsRegT struct {
	Val c.Uint32T
}

/** Type of out_w1tc register
 *  GPIO output clear register for GPIO0-31
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
 *  GPIO output enable register for GPIO0-31
 */

type GpioEnableRegT struct {
	Val c.Uint32T
}

/** Type of enable_w1ts register
 *  GPIO output enable set register for GPIO0-31
 */

type GpioEnableW1tsRegT struct {
	Val c.Uint32T
}

/** Type of enable_w1tc register
 *  GPIO output enable clear register for GPIO0-31
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
 *  GPIO input register for GPIO0-31
 */

type GpioInRegT struct {
	Val c.Uint32T
}

/** Type of status register
 *  GPIO interrupt status register for GPIO0-31
 */

type GpioStatusRegT struct {
	Val c.Uint32T
}

/** Type of status_w1ts register
 *  GPIO interrupt status set register for GPIO0-31
 */

type GpioStatusW1tsRegT struct {
	Val c.Uint32T
}

/** Type of status_w1tc register
 *  GPIO interrupt status clear register for GPIO0-31
 */

type GpioStatusW1tcRegT struct {
	Val c.Uint32T
}

/** Type of pcpu_int register
 *  GPIO PRO_CPU interrupt status register for GPIO0-31
 */

type GpioPcpuIntRegT struct {
	Val c.Uint32T
}

/** Type of pcpu_nmi_int register
 *  GPIO PRO_CPU(not shielded) interrupt status register for GPIO0-31
 */

type GpioPcpuNmiIntRegT struct {
	Val c.Uint32T
}

/** Type of cpusdio_int register
 *  GPIO CPUSDIO interrupt status register for GPIO0-31
 */

type GpioCpusdioIntRegT struct {
	Val c.Uint32T
}

/** Type of pinn register
 *  GPIO pin configuration register
 */

type GpioPinRegT struct {
	Val c.Uint32T
}

/** Type of status_next register
 *  GPIO interrupt source register for GPIO0-31
 */

type GpioStatusNextRegT struct {
	Val c.Uint32T
}

/** Type of func0_in_sel_cfg register
 *  GPIO input function configuration register
 */

type GpioFuncInSelCfgRegT struct {
	Val c.Uint32T
}

/** Type of funcn_out_sel_cfg register
 *  GPIO output function select register
 */

type GpioFuncOutSelCfgRegT struct {
	Val c.Uint32T
}

/** Type of clock_gate register
 *  GPIO clock gate register
 */

type GpioClockGateRegT struct {
	Val c.Uint32T
}

/** Type of date register
 *  GPIO version register
 */

type GpioDateRegT struct {
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
	Pin           [32]GpioPinRegT
	Reserved0f4   [22]c.Uint32T
	StatusNext    GpioStatusNextRegT
	Reserved150   c.Uint32T
	FuncInSelCfg  [128]GpioFuncInSelCfgRegT
	Reserved34b   [128]c.Uint32T
	FuncOutSelCfg [32]GpioFuncOutSelCfgRegT
	Reserved5d4   [22]c.Uint32T
	ClockGate     GpioClockGateRegT
	Reserved630   [51]c.Uint32T
	Date          GpioDateRegT
}
