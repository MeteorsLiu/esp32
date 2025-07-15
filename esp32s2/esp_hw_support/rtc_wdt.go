package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RtcWdtStageT c.Int

const (
	RTC_WDT_STAGE0 RtcWdtStageT = 0
	RTC_WDT_STAGE1 RtcWdtStageT = 1
	RTC_WDT_STAGE2 RtcWdtStageT = 2
	RTC_WDT_STAGE3 RtcWdtStageT = 3
)

type RtcWdtStageActionT c.Int

const (
	RTC_WDT_STAGE_ACTION_OFF          RtcWdtStageActionT = 0
	RTC_WDT_STAGE_ACTION_INTERRUPT    RtcWdtStageActionT = 1
	RTC_WDT_STAGE_ACTION_RESET_CPU    RtcWdtStageActionT = 2
	RTC_WDT_STAGE_ACTION_RESET_SYSTEM RtcWdtStageActionT = 3
	RTC_WDT_STAGE_ACTION_RESET_RTC    RtcWdtStageActionT = 4
)

type RtcWdtResetSigT c.Int

const (
	RTC_WDT_SYS_RESET_SIG RtcWdtResetSigT = 0
	RTC_WDT_CPU_RESET_SIG RtcWdtResetSigT = 1
)

type RtcWdtLengthSigT c.Int

const (
	RTC_WDT_LENGTH_100ns RtcWdtLengthSigT = 0
	RTC_WDT_LENGTH_200ns RtcWdtLengthSigT = 1
	RTC_WDT_LENGTH_300ns RtcWdtLengthSigT = 2
	RTC_WDT_LENGTH_400ns RtcWdtLengthSigT = 3
	RTC_WDT_LENGTH_500ns RtcWdtLengthSigT = 4
	RTC_WDT_LENGTH_800ns RtcWdtLengthSigT = 5
	RTC_WDT_LENGTH_1_6us RtcWdtLengthSigT = 6
	RTC_WDT_LENGTH_3_2us RtcWdtLengthSigT = 7
)

/**
 * @brief Get status of protect of rtc_wdt.
 *
 * @return
 *         - True if the protect of RTC_WDT is set
 */
//go:linkname RtcWdtGetProtectStatus C.rtc_wdt_get_protect_status
func RtcWdtGetProtectStatus() bool

/**
 * @brief Set protect of rtc_wdt.
 */
//go:linkname RtcWdtProtectOn C.rtc_wdt_protect_on
func RtcWdtProtectOn()

/**
 * @brief Reset protect of rtc_wdt.
 */
//go:linkname RtcWdtProtectOff C.rtc_wdt_protect_off
func RtcWdtProtectOff()

/**
 * @brief Enable rtc_wdt.
 */
//go:linkname RtcWdtEnable C.rtc_wdt_enable
func RtcWdtEnable()

/**
 * @brief Enable the flash boot protection procedure for WDT.
 *
 * Do not recommend to use it in the app.
 * This function was added to be compatibility with the old bootloaders.
 * This mode is disabled in bootloader or using rtc_wdt_disable() function.
 */
//go:linkname RtcWdtFlashbootModeEnable C.rtc_wdt_flashboot_mode_enable
func RtcWdtFlashbootModeEnable()

/**
 * @brief Disable rtc_wdt.
 */
//go:linkname RtcWdtDisable C.rtc_wdt_disable
func RtcWdtDisable()

/**
 * @brief Reset counter rtc_wdt.
 *
 * It returns to stage 0 and its expiry counter restarts from 0.
 */
//go:linkname RtcWdtFeed C.rtc_wdt_feed
func RtcWdtFeed()

/**
 * @brief Set time for required stage.
 *
 * @param[in] stage Stage of rtc_wdt.
 * @param[in] timeout_ms Timeout for this stage.
 *
 * @return
 *         - ESP_OK In case of success
 *         - ESP_ERR_INVALID_ARG If stage has invalid value
 */
// llgo:link RtcWdtStageT.RtcWdtSetTime C.rtc_wdt_set_time
func (recv_ RtcWdtStageT) RtcWdtSetTime(timeout_ms c.Uint) EspErrT {
	return 0
}

/**
 * @brief Get the timeout set for the required stage.
 *
 * @param[in]  stage Stage of rtc_wdt.
 * @param[out] timeout_ms Timeout set for this stage. (not elapsed time).
 *
 * @return
 *         - ESP_OK In case of success
 *         - ESP_ERR_INVALID_ARG If stage has invalid value
 */
// llgo:link RtcWdtStageT.RtcWdtGetTimeout C.rtc_wdt_get_timeout
func (recv_ RtcWdtStageT) RtcWdtGetTimeout(timeout_ms *c.Uint) EspErrT {
	return 0
}

/**
 * @brief Set an action for required stage.
 *
 * @param[in] stage Stage of rtc_wdt.
 * @param[in] stage_sel Action for this stage. When the time of stage expires this action will be triggered.
 *
 * @return
 *         - ESP_OK In case of success
 *         - ESP_ERR_INVALID_ARG If stage or stage_sel have invalid value
 */
// llgo:link RtcWdtStageT.RtcWdtSetStage C.rtc_wdt_set_stage
func (recv_ RtcWdtStageT) RtcWdtSetStage(stage_sel RtcWdtStageActionT) EspErrT {
	return 0
}

/**
 * @brief Set a length of reset signal.
 *
 * @param[in] reset_src Type of reset signal.
 * @param[in] reset_signal_length A length of reset signal.
 *
 * @return
 *         - ESP_OK In case of success
 *         - ESP_ERR_INVALID_ARG If reset_src  or reset_signal_length have invalid value
 */
// llgo:link RtcWdtResetSigT.RtcWdtSetLengthOfResetSignal C.rtc_wdt_set_length_of_reset_signal
func (recv_ RtcWdtResetSigT) RtcWdtSetLengthOfResetSignal(reset_signal_length RtcWdtLengthSigT) EspErrT {
	return 0
}

/**
 * @brief Return true if rtc_wdt is enabled.
 *
 * @return
 *         - True rtc_wdt is enabled
 */
//go:linkname RtcWdtIsOn C.rtc_wdt_is_on
func RtcWdtIsOn() bool
