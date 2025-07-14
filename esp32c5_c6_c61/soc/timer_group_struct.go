package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: T0 Control and configuration registers */
/** Type of txconfig register
 *  Timer x configuration register
 */

type TimgTxconfigRegT struct {
	Val c.Uint32T
}

/** Type of txlo register
 *  Timer x current value, low 32 bits
 */

type TimgTxloRegT struct {
	Val c.Uint32T
}

/** Type of txhi register
 *  Timer x current value, high 22 bits
 */

type TimgTxhiRegT struct {
	Val c.Uint32T
}

/** Type of txupdate register
 *  Write to copy current timer value to TIMGn_Tx_(LO/HI)_REG
 */

type TimgTxupdateRegT struct {
	Val c.Uint32T
}

/** Type of txalarmlo register
 *  Timer x alarm value, low 32 bits
 */

type TimgTxalarmloRegT struct {
	Val c.Uint32T
}

/** Type of txalarmhi register
 *  Timer x alarm value, high bits
 */

type TimgTxalarmhiRegT struct {
	Val c.Uint32T
}

/** Type of txloadlo register
 *  Timer x reload value, low 32 bits
 */

type TimgTxloadloRegT struct {
	Val c.Uint32T
}

/** Type of txloadhi register
 *  Timer x reload value, high 22 bits
 */

type TimgTxloadhiRegT struct {
	Val c.Uint32T
}

/** Type of txload register
 *  Write to reload timer from TIMG_Tx_(LOADLOLOADHI)_REG
 */

type TimgTxloadRegT struct {
	Val c.Uint32T
}

/** Group: WDT Control and configuration registers */
/** Type of wdtconfig0 register
 *  Watchdog timer configuration register
 */

type TimgWdtconfig0RegT struct {
	Val c.Uint32T
}

/** Type of wdtconfig1 register
 *  Watchdog timer prescaler register
 */

type TimgWdtconfig1RegT struct {
	Val c.Uint32T
}

/** Type of wdtconfig2 register
 *  Watchdog timer stage 0 timeout value
 */

type TimgWdtconfig2RegT struct {
	Val c.Uint32T
}

/** Type of wdtconfig3 register
 *  Watchdog timer stage 1 timeout value
 */

type TimgWdtconfig3RegT struct {
	Val c.Uint32T
}

/** Type of wdtconfig4 register
 *  Watchdog timer stage 2 timeout value
 */

type TimgWdtconfig4RegT struct {
	Val c.Uint32T
}

/** Type of wdtconfig5 register
 *  Watchdog timer stage 3 timeout value
 */

type TimgWdtconfig5RegT struct {
	Val c.Uint32T
}

/** Type of wdtfeed register
 *  Write to feed the watchdog timer
 */

type TimgWdtfeedRegT struct {
	Val c.Uint32T
}

/** Type of wdtwprotect register
 *  Watchdog write protect register
 */

type TimgWdtwprotectRegT struct {
	Val c.Uint32T
}

/** Group: RTC CALI Control and configuration registers */
/** Type of rtccalicfg register
 *  RTC calibration configure register
 */

type TimgRtccalicfgRegT struct {
	Val c.Uint32T
}

/** Type of rtccalicfg1 register
 *  RTC calibration configure1 register
 */

type TimgRtccalicfg1RegT struct {
	Val c.Uint32T
}

/** Type of rtccalicfg2 register
 *  Timer group calibration register
 */

type TimgRtccalicfg2RegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of int_ena_timers register
 *  Interrupt enable bits
 */

type TimgIntEnaTimersRegT struct {
	Val c.Uint32T
}

/** Type of int_raw_timers register
 *  Raw interrupt status
 */

type TimgIntRawTimersRegT struct {
	Val c.Uint32T
}

/** Type of int_st_timers register
 *  Masked interrupt status
 */

type TimgIntStTimersRegT struct {
	Val c.Uint32T
}

/** Type of int_clr_timers register
 *  Interrupt clear bits
 */

type TimgIntClrTimersRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of ntimers_date register
 *  Timer version control register
 */

type TimgNtimersDateRegT struct {
	Val c.Uint32T
}

/** Group: Clock configuration registers */
/** Type of regclk register
 *  Timer group clock gate register
 */

type TimgRegclkRegT struct {
	Val c.Uint32T
}

type TimgHwtimerRegT struct {
	Config  TimgTxconfigRegT
	Lo      TimgTxloRegT
	Hi      TimgTxhiRegT
	Update  TimgTxupdateRegT
	Alarmlo TimgTxalarmloRegT
	Alarmhi TimgTxalarmhiRegT
	Loadlo  TimgTxloadloRegT
	Loadhi  TimgTxloadhiRegT
	Load    TimgTxloadRegT
}

type TimgDevT struct {
	HwTimer      [1]TimgHwtimerRegT
	Reserved024  [9]c.Uint32T
	Wdtconfig0   TimgWdtconfig0RegT
	Wdtconfig1   TimgWdtconfig1RegT
	Wdtconfig2   TimgWdtconfig2RegT
	Wdtconfig3   TimgWdtconfig3RegT
	Wdtconfig4   TimgWdtconfig4RegT
	Wdtconfig5   TimgWdtconfig5RegT
	Wdtfeed      TimgWdtfeedRegT
	Wdtwprotect  TimgWdtwprotectRegT
	Rtccalicfg   TimgRtccalicfgRegT
	Rtccalicfg1  TimgRtccalicfg1RegT
	IntEnaTimers TimgIntEnaTimersRegT
	IntRawTimers TimgIntRawTimersRegT
	IntStTimers  TimgIntStTimersRegT
	IntClrTimers TimgIntClrTimersRegT
	Rtccalicfg2  TimgRtccalicfg2RegT
	Reserved084  [29]c.Uint32T
	NtimersDate  TimgNtimersDateRegT
	Regclk       TimgRegclkRegT
}
