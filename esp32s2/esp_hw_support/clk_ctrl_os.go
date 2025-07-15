package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief This function is used to enable the digital RC_FAST clock,
 *        to support the peripherals.
 *
 * @note If this function is called a number of times, the `periph_rtc_dig_clk8m_disable`
 *       function needs to be called same times to disable.
 *
 * @return true: success for enable the RC_FAST clock, false: RC_FAST clock enable failed
 */
//go:linkname PeriphRtcDigClk8mEnable C.periph_rtc_dig_clk8m_enable
func PeriphRtcDigClk8mEnable() bool

/**
 * @brief This function is used to disable the digital RC_FAST clock, which should be called
 *        with the `periph_rtc_dig_clk8m_enable` pairedly
 *
 * @note If this function is called a number of times, the `periph_rtc_dig_clk8m_disable`
 *       function needs to be called same times to disable.
 */
//go:linkname PeriphRtcDigClk8mDisable C.periph_rtc_dig_clk8m_disable
func PeriphRtcDigClk8mDisable()

/**
 * @brief This function is used to get the real clock frequency value of RC_FAST clock
 *
 * @return The real clock value, in Hz
 */
//go:linkname PeriphRtcDigClk8mGetFreq C.periph_rtc_dig_clk8m_get_freq
func PeriphRtcDigClk8mGetFreq() c.Uint32T

/**
 * @brief Enable APLL power if it has not enabled
 */
//go:linkname PeriphRtcApllAcquire C.periph_rtc_apll_acquire
func PeriphRtcApllAcquire()

/**
 * @brief Shut down APLL power if no peripherals using APLL
 */
//go:linkname PeriphRtcApllRelease C.periph_rtc_apll_release
func PeriphRtcApllRelease()

/**
 * @brief Calculate and set APLL coefficients by given frequency
 * @note  Have to call 'periph_rtc_apll_acquire' to enable APLL power before setting frequency
 * @note  This calculation is based on the inequality:
 *        xtal_freq * (4 + sdm2 + sdm1/256 + sdm0/65536) >= CLK_LL_APLL_MULTIPLIER_MIN_HZ(350 MHz)
 *        It will always calculate the minimum coefficients that can satisfy the inequality above, instead of loop them one by one.
 *        which means more appropriate coefficients are likely to exist.
 *        But this algorithm can meet almost all the cases and the accuracy can be guaranteed as well.
 * @note  The APLL frequency is only allowed to set when there is only one peripheral refer to it.
 *        If APLL is already set by another peripheral, this function will return `ESP_ERR_INVALID_STATE`
 *        and output the current frequency by parameter `real_freq`.
 *
 * @param expt_freq Expected APLL frequency (unit: Hz)
 * @param real_freq APLL real working frequency [output] (unit: Hz)
 * @return
 *      - ESP_OK: APLL frequency set success
 *      - ESP_ERR_INVALID_ARG: The input expt_freq is out of APLL support range
 *      - ESP_ERR_INVALID_STATE: APLL is referred by more than one peripherals, not allowed to change its frequency now
 */
//go:linkname PeriphRtcApllFreqSet C.periph_rtc_apll_freq_set
func PeriphRtcApllFreqSet(expt_freq_hz c.Uint32T, real_freq_hz *c.Uint32T) EspErrT
