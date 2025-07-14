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
/** Type of timerx_conf register
 *  Timer x configuration
 */

type LedcTimerxConfRegT struct {
	Val c.Uint32T
}

/** Type of timerx_value register
 *  Timer x current counter value
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

type LedcDevT struct {
	ChannelGroup [1]LedcChGroupRegT
	Reserved078  [10]c.Uint32T
	TimerGroup   [1]LedcTimerGroupRegT
	IntRaw       LedcIntRawRegT
	IntSt        LedcIntStRegT
	IntEna       LedcIntEnaRegT
	IntClr       LedcIntClrRegT
	Conf         LedcConfRegT
	Reserved0d4  [10]c.Uint32T
	Date         LedcDateRegT
}
