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

/** Type of out1 register
 *  GPIO output register for GPIO32-34
 */

type GpioOut1RegT struct {
	Val c.Uint32T
}

/** Type of out1_w1ts register
 *  GPIO output set register for GPIO32-34
 */

type GpioOut1W1tsRegT struct {
	Val c.Uint32T
}

/** Type of out1_w1tc register
 *  GPIO output clear register for GPIO32-34
 */

type GpioOut1W1tcRegT struct {
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

/** Type of enable1 register
 *  GPIO output enable register for GPIO32-34
 */

type GpioEnable1RegT struct {
	Val c.Uint32T
}

/** Type of enable1_w1ts register
 *  GPIO output enable set register for GPIO32-34
 */

type GpioEnable1W1tsRegT struct {
	Val c.Uint32T
}

/** Type of enable1_w1tc register
 *  GPIO output enable clear register for GPIO32-34
 */

type GpioEnable1W1tcRegT struct {
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

/** Type of in1 register
 *  GPIO input register for GPIO32-34
 */

type GpioIn1RegT struct {
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

/** Type of status1 register
 *  GPIO interrupt status register for GPIO32-34
 */

type GpioStatus1RegT struct {
	Val c.Uint32T
}

/** Type of status1_w1ts register
 *  GPIO interrupt status set register for GPIO32-34
 */

type GpioStatus1W1tsRegT struct {
	Val c.Uint32T
}

/** Type of status1_w1tc register
 *  GPIO interrupt status clear register for GPIO32-34
 */

type GpioStatus1W1tcRegT struct {
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

/** Type of pcpu_int1 register
 *  GPIO PRO_CPU interrupt status register for GPIO32-34
 */

type GpioPcpuInt1RegT struct {
	Val c.Uint32T
}

/** Type of pcpu_nmi_int1 register
 *  GPIO PRO_CPU(not shielded) interrupt status register for GPIO32-34
 */

type GpioPcpuNmiInt1RegT struct {
	Val c.Uint32T
}

/** Type of cpusdio_int1 register
 *  GPIO CPUSDIO interrupt status register for GPIO32-34
 */

type GpioCpusdioInt1RegT struct {
	Val c.Uint32T
}

/** Type of pin register
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

/** Type of status_next1 register
 *  GPIO interrupt source register for GPIO32-34
 */

type GpioStatusNext1RegT struct {
	Val c.Uint32T
}

/** Type of func_in_sel_cfg register
 *  GPIO input function configuration register
 */

type GpioFuncInSelCfgRegT struct {
	Val c.Uint32T
}

/** Type of func_out_sel_cfg register
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
	Out1          GpioOut1RegT
	Out1W1ts      GpioOut1W1tsRegT
	Out1W1tc      GpioOut1W1tcRegT
	SdioSelect    GpioSdioSelectRegT
	Enable        GpioEnableRegT
	EnableW1ts    GpioEnableW1tsRegT
	EnableW1tc    GpioEnableW1tcRegT
	Enable1       GpioEnable1RegT
	Enable1W1ts   GpioEnable1W1tsRegT
	Enable1W1tc   GpioEnable1W1tcRegT
	Strap         GpioStrapRegT
	In            GpioInRegT
	In1           GpioIn1RegT
	Status        GpioStatusRegT
	StatusW1ts    GpioStatusW1tsRegT
	StatusW1tc    GpioStatusW1tcRegT
	Status1       GpioStatus1RegT
	Status1W1ts   GpioStatus1W1tsRegT
	Status1W1tc   GpioStatus1W1tcRegT
	PcpuInt       GpioPcpuIntRegT
	PcpuNmiInt    GpioPcpuNmiIntRegT
	CpusdioInt    GpioCpusdioIntRegT
	PcpuInt1      GpioPcpuInt1RegT
	PcpuNmiInt1   GpioPcpuNmiInt1RegT
	CpusdioInt1   GpioCpusdioInt1RegT
	Pin           [35]GpioPinRegT
	Reserved100   [19]c.Uint32T
	StatusNext    GpioStatusNextRegT
	StatusNext1   GpioStatusNext1RegT
	FuncInSelCfg  [128]GpioFuncInSelCfgRegT
	Reserved34b   [128]c.Uint32T
	FuncOutSelCfg [35]GpioFuncOutSelCfgRegT
	Reserved5e0   [19]c.Uint32T
	ClockGate     GpioClockGateRegT
	Reserved630   [51]c.Uint32T
	Date          GpioDateRegT
}
