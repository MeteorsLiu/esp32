package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration and Control Register */
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
 *  Timer x current value, high 32 bits
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
 *  Timer x reload value, high 32 bits
 */

type TimgTxloadhiRegT struct {
	Val c.Uint32T
}

/** Type of txload register
 *  Write to reload timer from TIMG_T0_(LOADLOLOADHI)_REG
 */

type TimgTxloadRegT struct {
	Val c.Uint32T
}

/** Group: Configuration and Control Register for WDT */
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

/** Group: Configuration and Control Register for RTC CALI */
/** Type of rtccalicfg register
 *  RTC calibration configuration register
 */

type TimgRtccalicfgRegT struct {
	Val c.Uint32T
}

/** Type of rtccalicfg1 register
 *  RTC calibration configuration1 register
 */

type TimgRtccalicfg1RegT struct {
	Val c.Uint32T
}

/** Group: Configuration and Control Register for LACT */
/** Type of lactconfig register
 *  LACT configuration register
 */

type TimgLactconfigRegT struct {
	Val c.Uint32T
}

/** Type of lactrtc register
 *  LACT RTC register
 */

type TimgLactrtcRegT struct {
	Val c.Uint32T
}

/** Type of lactlo register
 *  LACT low register
 */

type TimgLactloRegT struct {
	Val c.Uint32T
}

/** Type of lacthi register
 *  LACT high register
 */

type TimgLacthiRegT struct {
	Val c.Uint32T
}

/** Type of lactupdate register
 *  LACT update register
 */

type TimgLactupdateRegT struct {
	Val c.Uint32T
}

/** Type of lactalarmlo register
 *  LACT alarm low register
 */

type TimgLactalarmloRegT struct {
	Val c.Uint32T
}

/** Type of lactalarmhi register
 *  LACT alarm high register
 */

type TimgLactalarmhiRegT struct {
	Val c.Uint32T
}

/** Type of lactloadlo register
 *  LACT load low register
 */

type TimgLactloadloRegT struct {
	Val c.Uint32T
}

/** Type of lactloadhi register
 *  Timer LACT load high register
 */

type TimgLactloadhiRegT struct {
	Val c.Uint32T
}

/** Type of lactload register
 *  Timer LACT load register
 */

type TimgLactloadRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Register */
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

/** Group: Version Register */
/** Type of timers_date register
 *  Version control register
 */

type TimgTimersDateRegT struct {
	Val c.Uint32T
}

/** Group: Configuration Register */
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
	HwTimer      [2]TimgHwtimerRegT
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
	Lactconfig   TimgLactconfigRegT
	Lactrtc      TimgLactrtcRegT
	Lactlo       TimgLactloRegT
	Lacthi       TimgLacthiRegT
	Lactupdate   TimgLactupdateRegT
	Lactalarmlo  TimgLactalarmloRegT
	Lactalarmhi  TimgLactalarmhiRegT
	Lactloadlo   TimgLactloadloRegT
	Lactloadhi   TimgLactloadhiRegT
	Lactload     TimgLactloadRegT
	IntEnaTimers TimgIntEnaTimersRegT
	IntRawTimers TimgIntRawTimersRegT
	IntStTimers  TimgIntStTimersRegT
	IntClrTimers TimgIntClrTimersRegT
	Reserved0ac  [20]c.Uint32T
	TimersDate   TimgTimersDateRegT
	Regclk       TimgRegclkRegT
}
