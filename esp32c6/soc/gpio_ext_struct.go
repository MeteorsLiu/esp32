package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: SDM Configure Registers */
/** Type of sigmadelta_chn register
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
	EventChnCfg  [8]GpioEtmEventChnCfgRegT
	Reserved080  [8]c.Uint32T
	EtmTaskPnCfg [8]GpioEtmTaskPnCfgRegT
}

type GpioExtDevT struct {
	SigmaDelta   GpioSdDevT
	Reserved028  [2]c.Uint32T
	GlitchFilter GpioGlitchFilterDevT
	Reserved050  [4]c.Uint32T
	Etm          GpioEtmDevT
	Reserved0c0  [15]c.Uint32T
	Version      GpioExtVersionRegT
}
