package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: SDM Configure Registers */
/** Type of sigmadeltan register
 *  Duty Cycle Configure Register of SDMn
 */

type GpioSigmadeltaChnRegT struct {
	Val c.Uint32T
}

/** Type of sigmadelta_misc register
 *  MISC Register
 */

type GpioSigmadeltaMiscRegT struct {
	Val c.Uint32T
}

/** Group: Clock gate Register */
/** Type of clock_gate register
 *  Clock Gating Configure Register
 */

type GpioSigmadeltaClockGateRegT struct {
	Val c.Uint32T
}

/** Group: Configure Registers */
/** Type of pad_comp_config register
 *  PAD Compare configure Register
 */

type GpioPadCompConfigRegT struct {
	Val c.Uint32T
}

/** Type of pad_comp_filter register
 *  Zero Detect filter Register
 */

type GpioPadCompFilterRegT struct {
	Val c.Uint32T
}

/** Group: Glitch filter Configure Registers */
/** Type of glitch_filter_chn register
 *  Glitch Filter Configure Register of Channeln
 */

type GpioGlitchFilterChnRegT struct {
	Val c.Uint32T
}

/** Group: Etm Configure Registers */
/** Type of etm_event_chn_cfg register
 *  Etm Config register of Channeln
 */

type GpioEtmEventChnCfgRegT struct {
	Val c.Uint32T
}

/** Type of etm_task_p0_cfg register
 *  Etm Configure Register to decide which GPIO been chosen
 */

type GpioEtmTaskPnCfgRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Registers */
/** Type of int_raw register
 *  GPIOSD interrupt raw register
 */

type GpioExtIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  GPIOSD interrupt masked register
 */

type GpioExtIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  GPIOSD interrupt enable register
 */

type GpioExtIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  GPIOSD interrupt clear register
 */

type GpioExtIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of version register
 *  Version Control Register
 */

type GpioExtVersionRegT struct {
	Val c.Uint32T
}

type GpioSdDevT struct {
	Channel     [4]GpioSigmadeltaChnRegT
	Reserved010 [4]c.Uint32T
	ClockGate   GpioSigmadeltaClockGateRegT
	Misc        GpioSigmadeltaMiscRegT
}

type GpioGlitchFilterDevT struct {
	GlitchFilterChn [8]GpioGlitchFilterChnRegT
}

type GpioEtmDevT struct {
	EtmEventChnCfg [8]GpioEtmEventChnCfgRegT
	Reserved080    [8]c.Uint32T
	EtmTaskPnCfg   [7]GpioEtmTaskPnCfgRegT
}

type GpioExtDevT struct {
	SigmaDelta    GpioSdDevT
	PadCompConfig GpioPadCompConfigRegT
	PadCompFilter GpioPadCompFilterRegT
	GlitchFilter  GpioGlitchFilterDevT
	Reserved050   [4]c.Uint32T
	Etm           GpioEtmDevT
	Reserved0bc   [9]c.Uint32T
	IntRaw        GpioExtIntRawRegT
	IntSt         GpioExtIntStRegT
	IntEna        GpioExtIntEnaRegT
	IntClr        GpioExtIntClrRegT
	Reserved0f0   [3]c.Uint32T
	Version       GpioExtVersionRegT
}
