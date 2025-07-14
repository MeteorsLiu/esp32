package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Register */
/** Type of chn_conf0 register
 *  Configuration register 0 for channel n
 */

type LedcChnConf0RegT struct {
	Val c.Uint32T
}

/** Type of chn_conf1 register
 *  Configuration register 1 for channel n
 */

type LedcChnConf1RegT struct {
	Val c.Uint32T
}

/** Type of evt_task_en0 register
 *  Ledc event task enable bit register0.
 */

type LedcEvtTaskEn0RegT struct {
	Val c.Uint32T
}

/** Type of evt_task_en1 register
 *  Ledc event task enable bit register1.
 */

type LedcEvtTaskEn1RegT struct {
	Val c.Uint32T
}

/** Type of evt_task_en2 register
 *  Ledc event task enable bit register2.
 */

type LedcEvtTaskEn2RegT struct {
	Val c.Uint32T
}

/** Type of timern_cmp register
 *  Ledc timern compare value register.
 */

type LedcTimernCmpRegT struct {
	Val c.Uint32T
}

/** Type of timern_cnt_cap register
 *  Ledc timern count value capture register.
 */

type LedcTimernCntCapRegT struct {
	Val c.Uint32T
}

/** Type of conf register
 *  Global ledc configuration register
 */

type LedcConfRegT struct {
	Val c.Uint32T
}

/** Group: Hpoint Register */
/** Type of chn_hpoint register
 *  High point register for channel n
 */

type LedcChnHpointRegT struct {
	Val c.Uint32T
}

/** Group: Duty Cycle Register */
/** Type of chn_duty register
 *  Initial duty cycle for channel n
 */

type LedcChnDutyRegT struct {
	Val c.Uint32T
}

/** Type of chn_duty_r register
 *  Current duty cycle for channel n
 */

type LedcChnDutyRRegT struct {
	Val c.Uint32T
}

/** Group: Timer Register */
/** Type of timern_conf register
 *  Timer n configuration
 */

type LedcTimerxConfRegT struct {
	Val c.Uint32T
}

/** Type of timern_value register
 *  Timer n current counter value
 */

type LedcTimerxValueRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Register */
/** Type of int_raw register
 *  Raw interrupt status
 */

type LedcIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Masked interrupt status
 */

type LedcIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Interrupt enable bits
 */

type LedcIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear bits
 */

type LedcIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Gamma RAM Register */
/** Type of chn_gamma_wr register
 *  Ledc chn gamma ram write register.
 */

type LedcChnGammaWrRegT struct {
	Val c.Uint32T
}

/** Type of chn_gamma_wr_addr register
 *  Ledc chn gamma ram write address register.
 */

type LedcChnGammaWrAddrRegT struct {
	Val c.Uint32T
}

/** Type of chn_gamma_rd_addr register
 *  Ledc chn gamma ram read address register.
 */

type LedcChnGammaRdAddrRegT struct {
	Val c.Uint32T
}

/** Type of chn_gamma_rd_data register
 *  Ledc chn gamma ram read data register.
 */

type LedcChnGammaRdDataRegT struct {
	Val c.Uint32T
}

/** Group: Gamma Config Register */
/** Type of chn_gamma_conf register
 *  Ledc chn gamma config register.
 */

type LedcChnGammaConfRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  Version control register
 */

type LedcDateRegT struct {
	Val c.Uint32T
}

type LedcChnRegT struct {
	Conf0  LedcChnConf0RegT
	Hpoint LedcChnHpointRegT
	Duty   LedcChnDutyRegT
	Conf1  LedcChnConf1RegT
	DutyRd LedcChnDutyRRegT
}

type LedcChGroupRegT struct {
	Channel [6]LedcChnRegT
}

type LedcTimerxRegT struct {
	Conf  LedcTimerxConfRegT
	Value LedcTimerxValueRegT
}

type LedcTimerGroupRegT struct {
	Timer [4]LedcTimerxRegT
}

type LedcChnGammaRegT struct {
	Wr     LedcChnGammaWrRegT
	WrAddr LedcChnGammaWrAddrRegT
	RdAddr LedcChnGammaRdAddrRegT
	RdData LedcChnGammaRdDataRegT
}

type LedcChGammaGroupRegT struct {
	Channel [6]LedcChnGammaRegT
}

type LedcChGammaConfGroupRegT struct {
	GammaConf [6]LedcChnGammaConfRegT
}

type LedcTimerCmpGroupRegT struct {
	Cmp [4]LedcTimernCmpRegT
}

type LedcTimerCntCapGroupRegT struct {
	CntCap [4]LedcTimernCntCapRegT
}

type LedcDevT struct {
	ChannelGroup          [1]LedcChGroupRegT
	Reserved078           [10]c.Uint32T
	TimerGroup            [1]LedcTimerGroupRegT
	IntRaw                LedcIntRawRegT
	IntSt                 LedcIntStRegT
	IntEna                LedcIntEnaRegT
	IntClr                LedcIntClrRegT
	Reserved0d0           [12]c.Uint32T
	ChannelGammaGroup     [1]LedcChGammaGroupRegT
	Reserved160           [8]c.Uint32T
	ChannelGammaConfGroup [1]LedcChGammaConfGroupRegT
	Reserved198           [2]c.Uint32T
	EvtTaskEn0            LedcEvtTaskEn0RegT
	EvtTaskEn1            LedcEvtTaskEn1RegT
	EvtTaskEn2            LedcEvtTaskEn2RegT
	Reserved1ac           c.Uint32T
	TimerCmpGroup         [1]LedcTimerCmpGroupRegT
	TimerCntCapGroup      [1]LedcTimerCntCapGroupRegT
	Reserved1d0           [8]c.Uint32T
	Conf                  LedcConfRegT
	Reserved1f4           [2]c.Uint32T
	Date                  LedcDateRegT
}
