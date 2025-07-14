package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ModemClockHalContextT struct {
	SysconDev *ModemSysconDevT
	LpconDev  *ModemLpconDevT
}

// llgo:link (*ModemClockHalContextT).ModemClockHalEnableModemCommonFeClock C.modem_clock_hal_enable_modem_common_fe_clock
func (recv_ *ModemClockHalContextT) ModemClockHalEnableModemCommonFeClock(enable bool) {
}

// llgo:link (*ModemClockHalContextT).ModemClockHalEnableModemPrivateFeClock C.modem_clock_hal_enable_modem_private_fe_clock
func (recv_ *ModemClockHalContextT) ModemClockHalEnableModemPrivateFeClock(enable bool) {
}

// llgo:link (*ModemClockHalContextT).ModemClockHalSetBleRtcTimerDivisorValue C.modem_clock_hal_set_ble_rtc_timer_divisor_value
func (recv_ *ModemClockHalContextT) ModemClockHalSetBleRtcTimerDivisorValue(divider c.Uint32T) {
}

// llgo:link (*ModemClockHalContextT).ModemClockHalEnableBleRtcTimerClock C.modem_clock_hal_enable_ble_rtc_timer_clock
func (recv_ *ModemClockHalContextT) ModemClockHalEnableBleRtcTimerClock(enable bool) {
}

// llgo:link (*ModemClockHalContextT).ModemClockHalSelectBleRtcTimerLpclkSource C.modem_clock_hal_select_ble_rtc_timer_lpclk_source
func (recv_ *ModemClockHalContextT) ModemClockHalSelectBleRtcTimerLpclkSource(src ModemClockLpclkSrcT) {
}

// llgo:link (*ModemClockHalContextT).ModemClockHalDeselectAllBleRtcTimerLpclkSource C.modem_clock_hal_deselect_all_ble_rtc_timer_lpclk_source
func (recv_ *ModemClockHalContextT) ModemClockHalDeselectAllBleRtcTimerLpclkSource() {
}

// llgo:link (*ModemClockHalContextT).ModemClockHalSelectCoexLpclkSource C.modem_clock_hal_select_coex_lpclk_source
func (recv_ *ModemClockHalContextT) ModemClockHalSelectCoexLpclkSource(src ModemClockLpclkSrcT) {
}

// llgo:link (*ModemClockHalContextT).ModemClockHalDeselectAllCoexLpclkSource C.modem_clock_hal_deselect_all_coex_lpclk_source
func (recv_ *ModemClockHalContextT) ModemClockHalDeselectAllCoexLpclkSource() {
}
