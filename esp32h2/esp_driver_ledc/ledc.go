package esp_driver_ledc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LedcSleepModeT c.Int

const (
	LEDC_SLEEP_MODE_NO_ALIVE_NO_PD    LedcSleepModeT = 0
	LEDC_SLEEP_MODE_NO_ALIVE_ALLOW_PD LedcSleepModeT = 1
	LEDC_SLEEP_MODE_KEEP_ALIVE        LedcSleepModeT = 2
	LEDC_SLEEP_MODE_INVALID           LedcSleepModeT = 3
)

/**
 * @brief Configuration parameters of LEDC channel for ledc_channel_config function
 */

type LedcChannelConfigT struct {
	GpioNum   c.Int
	SpeedMode LedcModeT
	Channel   LedcChannelT
	IntrType  LedcIntrTypeT
	TimerSel  LedcTimerT
	Duty      c.Uint32T
	Hpoint    c.Int
	SleepMode LedcSleepModeT
	Flags     struct {
		OutputInvert c.Uint
	}
}

/**
 * @brief Configuration parameters of LEDC timer for ledc_timer_config function
 */

type LedcTimerConfigT struct {
	SpeedMode      LedcModeT
	DutyResolution LedcTimerBitT
	TimerNum       LedcTimerT
	FreqHz         c.Uint32T
	ClkCfg         LedcClkCfgT
	Deconfigure    bool
}
type LedcIsrHandleT IntrHandleT
type LedcCbEventT c.Int

const LEDC_FADE_END_EVT LedcCbEventT = 0

/**
 * @brief LEDC callback parameter
 */

type LedcCbParamT struct {
	Event     LedcCbEventT
	SpeedMode c.Uint32T
	Channel   c.Uint32T
	Duty      c.Uint32T
}

// llgo:type C
type LedcCbT func(*LedcCbParamT, c.Pointer) bool

/**
 * @brief Group of supported LEDC callbacks
 * @note The callbacks are all running under ISR environment
 */

type LedcCbsT struct {
	FadeCb LedcCbT
}

/**
 * @brief LEDC channel configuration
 *        Configure LEDC channel with the given channel/output gpio_num/interrupt/source timer/frequency(Hz)/LEDC duty
 *
 * @param ledc_conf Pointer of LEDC channel configure struct
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link (*LedcChannelConfigT).LedcChannelConfig C.ledc_channel_config
func (recv_ *LedcChannelConfigT) LedcChannelConfig() EspErrT {
	return 0
}

/**
 * @brief Helper function to find the maximum possible duty resolution in bits for ledc_timer_config()
 *
 * @param src_clk_freq LEDC timer source clock frequency (Hz) (See doxygen comments of `ledc_clk_cfg_t` or get from `esp_clk_tree_src_get_freq_hz`)
 * @param timer_freq Desired LEDC timer frequency (Hz)
 *
 * @return
 *     - 0 The timer frequency cannot be achieved
 *     - Others The largest duty resolution value to be set
 */
//go:linkname LedcFindSuitableDutyResolution C.ledc_find_suitable_duty_resolution
func LedcFindSuitableDutyResolution(src_clk_freq c.Uint32T, timer_freq c.Uint32T) c.Uint32T

/**
 * @brief LEDC timer configuration
 *        Configure LEDC timer with the given source timer/frequency(Hz)/duty_resolution
 *
 * @param  timer_conf Pointer of LEDC timer configure struct
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_FAIL Can not find a proper pre-divider number base on the given frequency and the current duty_resolution.
 *     - ESP_ERR_INVALID_STATE Timer cannot be de-configured because timer is not configured or is not paused
 */
// llgo:link (*LedcTimerConfigT).LedcTimerConfig C.ledc_timer_config
func (recv_ *LedcTimerConfigT) LedcTimerConfig() EspErrT {
	return 0
}

/**
 * @brief LEDC update channel parameters
 *
 * @note  Call this function to activate the LEDC updated parameters.
 *        After ledc_set_duty, we need to call this function to update the settings.
 *        And the new LEDC parameters don't take effect until the next PWM cycle.
 * @note  ledc_set_duty, ledc_set_duty_with_hpoint and ledc_update_duty are not thread-safe, do not call these functions to
 *        control one LEDC channel in different tasks at the same time.
 *        A thread-safe version of API is ledc_set_duty_and_update
 * @note  If `CONFIG_LEDC_CTRL_FUNC_IN_IRAM` is enabled, this function will be placed in the IRAM by linker,
 *        makes it possible to execute even when the Cache is disabled.
 * @note  This function is allowed to run within ISR context.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname LedcUpdateDuty C.ledc_update_duty
func LedcUpdateDuty(speed_mode LedcModeT, channel LedcChannelT) EspErrT

/**
 * @brief Set LEDC output gpio.
 *
 * @note This function only routes the LEDC signal to GPIO through matrix, other LEDC resources initialization are not involved.
 *       Please use `ledc_channel_config()` instead to fully configure a LEDC channel.
 *
 * @param  gpio_num The LEDC output gpio
 * @param  speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param  channel LEDC channel (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname LedcSetPin C.ledc_set_pin
func LedcSetPin(gpio_num c.Int, speed_mode LedcModeT, channel LedcChannelT) EspErrT

/**
 * @brief LEDC stop.
 *        Disable LEDC output, and set idle level
 *
 * @note  If `CONFIG_LEDC_CTRL_FUNC_IN_IRAM` is enabled, this function will be placed in the IRAM by linker,
 *        makes it possible to execute even when the Cache is disabled.
 * @note  This function is allowed to run within ISR context.
 *
 * @param  speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param  channel LEDC channel (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param  idle_level Set output idle level after LEDC stops.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname LedcStop C.ledc_stop
func LedcStop(speed_mode LedcModeT, channel LedcChannelT, idle_level c.Uint32T) EspErrT

/**
 * @brief LEDC set channel frequency (Hz)
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param  timer_num LEDC timer index (0-3), select from ledc_timer_t
 * @param  freq_hz Set the LEDC frequency
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_FAIL Can not find a proper pre-divider number base on the given frequency and the current duty_resolution.
 */
//go:linkname LedcSetFreq C.ledc_set_freq
func LedcSetFreq(speed_mode LedcModeT, timer_num LedcTimerT, freq_hz c.Uint32T) EspErrT

/**
 * @brief      LEDC get channel frequency (Hz)
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param timer_num LEDC timer index (0-3), select from ledc_timer_t
 *
 * @return
 *     - 0  error
 *     - Others Current LEDC frequency
 */
//go:linkname LedcGetFreq C.ledc_get_freq
func LedcGetFreq(speed_mode LedcModeT, timer_num LedcTimerT) c.Uint32T

/**
 * @brief LEDC set duty and hpoint value
 *        Only after calling ledc_update_duty will the duty update.
 *
 * @note  ledc_set_duty, ledc_set_duty_with_hpoint and ledc_update_duty are not thread-safe, do not call these functions to
 *        control one LEDC channel in different tasks at the same time.
 *        A thread-safe version of API is ledc_set_duty_and_update
 * @note  For ESP32, hardware does not support any duty change while a fade operation is running in progress on that channel.
 *        Other duty operations will have to wait until the fade operation has finished.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param duty Set the LEDC duty, the range of duty setting is [0, (2**duty_resolution)]
 * @param hpoint Set the LEDC hpoint value, the range is [0, (2**duty_resolution)-1]
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname LedcSetDutyWithHpoint C.ledc_set_duty_with_hpoint
func LedcSetDutyWithHpoint(speed_mode LedcModeT, channel LedcChannelT, duty c.Uint32T, hpoint c.Uint32T) EspErrT

/**
 * @brief LEDC get hpoint value, the counter value when the output is set high level.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 *
 * @return
 *     - LEDC_ERR_VAL if parameter error
 *     - Others Current hpoint value of LEDC channel
 */
//go:linkname LedcGetHpoint C.ledc_get_hpoint
func LedcGetHpoint(speed_mode LedcModeT, channel LedcChannelT) c.Int

/**
 * @brief LEDC set duty
 *        This function do not change the hpoint value of this channel. if needed, please call ledc_set_duty_with_hpoint.
 *        only after calling ledc_update_duty will the duty update.
 *
 * @note  ledc_set_duty, ledc_set_duty_with_hpoint and ledc_update_duty are not thread-safe, do not call these functions to
 *        control one LEDC channel in different tasks at the same time.
 *        A thread-safe version of API is ledc_set_duty_and_update.
 * @note  For ESP32, hardware does not support any duty change while a fade operation is running in progress on that channel.
 *        Other duty operations will have to wait until the fade operation has finished.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param duty Set the LEDC duty, the range of duty setting is [0, (2**duty_resolution)]
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname LedcSetDuty C.ledc_set_duty
func LedcSetDuty(speed_mode LedcModeT, channel LedcChannelT, duty c.Uint32T) EspErrT

/**
 * @brief LEDC get duty
 *        This function returns the duty at the present PWM cycle.
 *        You shouldn't expect the function to return the new duty in the same cycle of calling ledc_update_duty,
 *        because duty update doesn't take effect until the next cycle.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 *
 * @return
 *     - LEDC_ERR_DUTY if parameter error
 *     - Others Current LEDC duty
 */
//go:linkname LedcGetDuty C.ledc_get_duty
func LedcGetDuty(speed_mode LedcModeT, channel LedcChannelT) c.Uint32T

/**
 * @brief LEDC set gradient
 *        Set LEDC gradient, After the function calls the ledc_update_duty function, the function can take effect.
 *
 * @note  For ESP32, hardware does not support any duty change while a fade operation is running in progress on that channel.
 *        Other duty operations will have to wait until the fade operation has finished.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param duty Set the start of the gradient duty, the range of duty setting is [0, (2**duty_resolution)]
 * @param fade_direction Set the direction of the gradient
 * @param step_num Set the number of the gradient
 * @param duty_cycle_num Set how many LEDC tick each time the gradient lasts
 * @param duty_scale Set gradient change amplitude
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname LedcSetFade C.ledc_set_fade
func LedcSetFade(speed_mode LedcModeT, channel LedcChannelT, duty c.Uint32T, fade_direction LedcDutyDirectionT, step_num c.Uint32T, duty_cycle_num c.Uint32T, duty_scale c.Uint32T) EspErrT

/**
 * @brief Register LEDC interrupt handler, the handler is an ISR.
 *        The handler will be attached to the same CPU core that this function is running on.
 *
 * @param fn Interrupt handler function.
 * @param arg User-supplied argument passed to the handler function.
 * @param intr_alloc_flags Flags used to allocate the interrupt. One or multiple (ORred)
 *        ESP_INTR_FLAG_* values. See esp_intr_alloc.h for more info.
 * @param handle Pointer to return handle. If non-NULL, a handle for the interrupt will
 *        be returned here.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_NOT_FOUND Failed to find available interrupt source
 */
//go:linkname LedcIsrRegister C.ledc_isr_register
func LedcIsrRegister(fn func(c.Pointer), arg c.Pointer, intr_alloc_flags c.Int, handle *LedcIsrHandleT) EspErrT

/**
 * @brief Configure LEDC settings
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param timer_sel  Timer index (0-3), there are 4 timers in LEDC module
 * @param clock_divider Timer clock divide value, the timer clock is divided from the selected clock source
 * @param duty_resolution Resolution of duty setting in number of bits. The range is [1, SOC_LEDC_TIMER_BIT_WIDTH]
 * @param clk_src Select LEDC source clock.
 *
 * @return
 *     - (-1) Parameter error
 *     - Other Current LEDC duty
 */
//go:linkname LedcTimerSet C.ledc_timer_set
func LedcTimerSet(speed_mode LedcModeT, timer_sel LedcTimerT, clock_divider c.Uint32T, duty_resolution c.Uint32T, clk_src LedcClkSrcT) EspErrT

/**
 * @brief Reset LEDC timer
 *
 * @param  speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param  timer_sel LEDC timer index (0-3), select from ledc_timer_t
 *
 * @return
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_OK Success
 */
//go:linkname LedcTimerRst C.ledc_timer_rst
func LedcTimerRst(speed_mode LedcModeT, timer_sel LedcTimerT) EspErrT

/**
 * @brief Pause LEDC timer counter
 *
 * @param  speed_mode  Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param  timer_sel LEDC timer index (0-3), select from ledc_timer_t
 *
 * @return
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_OK Success
 */
//go:linkname LedcTimerPause C.ledc_timer_pause
func LedcTimerPause(speed_mode LedcModeT, timer_sel LedcTimerT) EspErrT

/**
 * @brief Resume LEDC timer
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param timer_sel LEDC timer index (0-3), select from ledc_timer_t
 *
 * @return
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_OK Success
 */
//go:linkname LedcTimerResume C.ledc_timer_resume
func LedcTimerResume(speed_mode LedcModeT, timer_sel LedcTimerT) EspErrT

/**
 * @brief Bind LEDC channel with the selected timer
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel index (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param timer_sel LEDC timer index (0-3), select from ledc_timer_t
 *
 * @return
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_OK Success
 */
//go:linkname LedcBindChannelTimer C.ledc_bind_channel_timer
func LedcBindChannelTimer(speed_mode LedcModeT, channel LedcChannelT, timer_sel LedcTimerT) EspErrT

/**
 * @brief Set LEDC fade function.
 *
 * @note  Call ledc_fade_func_install() once before calling this function.
 *        Call ledc_fade_start() after this to start fading.
 * @note  ledc_set_fade_with_step, ledc_set_fade_with_time and ledc_fade_start are not thread-safe, do not call these functions to
 *        control one LEDC channel in different tasks at the same time.
 *        A thread-safe version of API is ledc_set_fade_step_and_start
 * @note  For ESP32, hardware does not support any duty change while a fade operation is running in progress on that channel.
 *        Other duty operations will have to wait until the fade operation has finished.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel index (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param target_duty Target duty of fading [0, (2**duty_resolution)]
 * @param scale Controls the increase or decrease step scale.
 * @param cycle_num increase or decrease the duty every cycle_num cycles
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Channel not initialized
 *     - ESP_FAIL Fade function init error
 */
//go:linkname LedcSetFadeWithStep C.ledc_set_fade_with_step
func LedcSetFadeWithStep(speed_mode LedcModeT, channel LedcChannelT, target_duty c.Uint32T, scale c.Uint32T, cycle_num c.Uint32T) EspErrT

/**
 * @brief Set LEDC fade function, with a limited time.
 *
 * @note  Call ledc_fade_func_install() once before calling this function.
 *        Call ledc_fade_start() after this to start fading.
 * @note  ledc_set_fade_with_step, ledc_set_fade_with_time and ledc_fade_start are not thread-safe, do not call these functions to
 *        control one LEDC channel in different tasks at the same time.
 *        A thread-safe version of API is ledc_set_fade_step_and_start
 * @note  For ESP32, hardware does not support any duty change while a fade operation is running in progress on that channel.
 *        Other duty operations will have to wait until the fade operation has finished.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel index (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param target_duty Target duty of fading [0, (2**duty_resolution)]
 * @param desired_fade_time_ms The intended time of the fading ( ms ).
 *                             Note that the actual time it takes to complete the fade could vary by a factor of up to 2x shorter
 *                             or longer than the expected time due to internal rounding errors in calculations.
 *                             Specifically:
 *                             * The total number of cycles (total_cycle_num = desired_fade_time_ms * freq / 1000)
 *                             * The difference in duty cycle (duty_delta = |target_duty - current_duty|)
 *                             The fade may complete faster than expected if total_cycle_num larger than duty_delta. Conversely,
 *                             it may take longer than expected if total_cycle_num is less than duty_delta.
 *                             The closer the ratio of total_cycle_num/duty_delta (or its inverse) is to a whole number (the floor value),
 *                             the more accurately the actual fade duration will match the intended time.
 *                             If an exact fade time is expected, please consider to split the entire fade into several smaller linear fades.
 *                             The split should make each fade step has a divisible total_cycle_num/duty_delta (or its inverse) ratio.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Channel not initialized
 *     - ESP_FAIL Fade function init error
 */
//go:linkname LedcSetFadeWithTime C.ledc_set_fade_with_time
func LedcSetFadeWithTime(speed_mode LedcModeT, channel LedcChannelT, target_duty c.Uint32T, desired_fade_time_ms c.Int) EspErrT

/**
 * @brief Install LEDC fade function. This function will occupy interrupt of LEDC module.
 *
 * @param intr_alloc_flags Flags used to allocate the interrupt. One or multiple (ORred)
 *        ESP_INTR_FLAG_* values. See esp_intr_alloc.h for more info.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Intr flag error
 *     - ESP_ERR_NOT_FOUND Failed to find available interrupt source
 *     - ESP_ERR_INVALID_STATE Fade function already installed
 */
//go:linkname LedcFadeFuncInstall C.ledc_fade_func_install
func LedcFadeFuncInstall(intr_alloc_flags c.Int) EspErrT

/**
 * @brief Uninstall LEDC fade function.
 */
//go:linkname LedcFadeFuncUninstall C.ledc_fade_func_uninstall
func LedcFadeFuncUninstall()

/**
 * @brief Start LEDC fading.
 *
 * @note  Call ledc_fade_func_install() once before calling this function.
 *        Call this API right after ledc_set_fade_with_time or ledc_set_fade_with_step before to start fading.
 * @note  Starting fade operation with this API is not thread-safe, use with care.
 * @note  For ESP32, hardware does not support any duty change while a fade operation is running in progress on that channel.
 *        Other duty operations will have to wait until the fade operation has finished.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel number
 * @param fade_mode Whether to block until fading done. See ledc_types.h ledc_fade_mode_t for more info.
 *        Note that this function will not return until fading to the target duty if LEDC_FADE_WAIT_DONE mode is selected.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE Channel not initialized or fade function not installed.
 *     - ESP_ERR_INVALID_ARG Parameter error.
 */
//go:linkname LedcFadeStart C.ledc_fade_start
func LedcFadeStart(speed_mode LedcModeT, channel LedcChannelT, fade_mode LedcFadeModeT) EspErrT

/**
 * @brief Stop LEDC fading. The duty of the channel is guaranteed to be fixed at most one PWM cycle after the function returns.
 *
 * @note  This API can be called if a new fixed duty or a new fade want to be set while the last fade operation is still running in progress.
 * @note  Call this API will abort the fading operation only if it was started by calling ledc_fade_start with LEDC_FADE_NO_WAIT mode.
 * @note  If a fade was started with LEDC_FADE_WAIT_DONE mode, calling this API afterwards has no use in stopping the fade. Fade will continue until it reaches the target duty.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE Channel not initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_FAIL Fade function init error
 */
//go:linkname LedcFadeStop C.ledc_fade_stop
func LedcFadeStop(speed_mode LedcModeT, channel LedcChannelT) EspErrT

/**
 * @brief A thread-safe API to set duty for LEDC channel and return when duty updated.
 *
 * @note  For ESP32, hardware does not support any duty change while a fade operation is running in progress on that channel.
 *        Other duty operations will have to wait until the fade operation has finished.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param duty Set the LEDC duty, the range of duty setting is [0, (2**duty_resolution)]
 * @param hpoint Set the LEDC hpoint value, the range is [0, (2**duty_resolution)-1]
 *
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_STATE Channel not initialized
 *      - ESP_ERR_INVALID_ARG Parameter error
 *      - ESP_FAIL Fade function init error
 */
//go:linkname LedcSetDutyAndUpdate C.ledc_set_duty_and_update
func LedcSetDutyAndUpdate(speed_mode LedcModeT, channel LedcChannelT, duty c.Uint32T, hpoint c.Uint32T) EspErrT

/**
 * @brief A thread-safe API to set and start LEDC fade function, with a limited time.
 *
 * @note  Call ledc_fade_func_install() once, before calling this function.
 * @note  For ESP32, hardware does not support any duty change while a fade operation is running in progress on that channel.
 *        Other duty operations will have to wait until the fade operation has finished.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel index (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param target_duty Target duty of fading [0, (2**duty_resolution)]
 * @param desired_fade_time_ms The intended time of the fading ( ms ).
 *                             Note that the actual time it takes to complete the fade could vary by a factor of up to 2x shorter
 *                             or longer than the expected time due to internal rounding errors in calculations.
 *                             Specifically:
 *                             * The total number of cycles (total_cycle_num = desired_fade_time_ms * freq / 1000)
 *                             * The difference in duty cycle (duty_delta = |target_duty - current_duty|)
 *                             The fade may complete faster than expected if total_cycle_num larger than duty_delta. Conversely,
 *                             it may take longer than expected if total_cycle_num is less than duty_delta.
 *                             The closer the ratio of total_cycle_num/duty_delta (or its inverse) is to a whole number (the floor value),
 *                             the more accurately the actual fade duration will match the intended time.
 *                             If an exact fade time is expected, please consider to split the entire fade into several smaller linear fades.
 *                             The split should make each fade step has a divisible total_cycle_num/duty_delta (or its inverse) ratio.
 * @param fade_mode choose blocking or non-blocking mode
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Channel not initialized
 *     - ESP_FAIL Fade function init error
 */
//go:linkname LedcSetFadeTimeAndStart C.ledc_set_fade_time_and_start
func LedcSetFadeTimeAndStart(speed_mode LedcModeT, channel LedcChannelT, target_duty c.Uint32T, desired_fade_time_ms c.Uint32T, fade_mode LedcFadeModeT) EspErrT

/**
 * @brief A thread-safe API to set and start LEDC fade function.
 *
 * @note  Call ledc_fade_func_install() once before calling this function.
 * @note  For ESP32, hardware does not support any duty change while a fade operation is running in progress on that channel.
 *        Other duty operations will have to wait until the fade operation has finished.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel index (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param target_duty Target duty of fading [0, (2**duty_resolution)]
 * @param scale Controls the increase or decrease step scale.
 * @param cycle_num increase or decrease the duty every cycle_num cycles
 * @param fade_mode choose blocking or non-blocking mode
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Channel not initialized
 *     - ESP_FAIL Fade function init error
 */
//go:linkname LedcSetFadeStepAndStart C.ledc_set_fade_step_and_start
func LedcSetFadeStepAndStart(speed_mode LedcModeT, channel LedcChannelT, target_duty c.Uint32T, scale c.Uint32T, cycle_num c.Uint32T, fade_mode LedcFadeModeT) EspErrT

/**
 * @brief LEDC callback registration function
 *
 * @note  The callback is called from an ISR, it must never attempt to block, and any FreeRTOS API called must be ISR capable.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel index (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param cbs Group of LEDC callback functions
 * @param user_arg user registered data for the callback function
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Channel not initialized
 *     - ESP_FAIL Fade function init error
 */
//go:linkname LedcCbRegister C.ledc_cb_register
func LedcCbRegister(speed_mode LedcModeT, channel LedcChannelT, cbs *LedcCbsT, user_arg c.Pointer) EspErrT

/**
 * @brief Structure for the fade parameters for one hardware fade to be written to gamma wr register
 *
 * @verbatim
 *                                  duty ^                 ONE HW LINEAR FADE
 *                                       |
 *                                       |
 *                                       |
 *                                       |
 *     start_duty + scale * n = end_duty |. . . . . . . . . . . . . . . . . . . . . . . . . .+-
 *                                       |                                                   |
 *                                       |                                                   |
 *                                       |                                          +--------+
 *                                       |                                          |        .
 *                                       |                                          |        .
 *                                       |                                   -------+        .
 *                                       |                                  .                .
 *                                       |                                .                  .
 *                                       |                              .                    .
 *                                       |                            .                      .
 *                                 ^ --- |. . . . . . . . . .+--------                       .
 *                            scale|     |                   |                               .
 *                                 |     |                   |                               .
 *                                 v --- |. . . . .+---------+                               .
 *                                       |         |         .                               .
 *                                       |         |         .                               .
 *                            start_duty +---------+         .                               .
 *                                       |         .         .                               .
 *                                       |         .         .                               .
 *                                       +----------------------------------------------------------->
 *                                                                                                  PWM cycle
 *                                       |         |         |                               |
 *                                       | 1 step  | 1 step  |                               |
 *                                       |<------->|<------->|                               |
 *                                       | m cycles  m cycles                                |
 *                                       |                                                   |
 *                                       <--------------------------------------------------->
 *                                                           n total steps
 *                                                          cycles = m * n
 * @endverbatim
 *
 * @note Be aware of the maximum value available on each element
 */

type LedcFadeParamConfigT struct {
	Dir      c.Uint32T
	CycleNum c.Uint32T
	Scale    c.Uint32T
	StepNum  c.Uint32T
}

/**
 * @brief Set a LEDC multi-fade
 *
 * @note  Call `ledc_fade_func_install()` once before calling this function.
 *        Call `ledc_fade_start()` after this to start fading.
 * @note  This function is not thread-safe, do not call it to control one LEDC channel in different tasks at the same time.
 *        A thread-safe version of API is ledc_set_multi_fade_and_start
 * @note  This function does not prohibit from duty overflow. User should take care of this by themselves. If duty
 *        overflow happens, the PWM signal will suddenly change from 100% duty cycle to 0%, or the other way around.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel index (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param start_duty Set the start of the gradient duty, the range of duty setting is [0, (2**duty_resolution)]
 * @param fade_params_list Pointer to the array of fade parameters for a multi-fade
 * @param list_len Length of the fade_params_list, i.e. number of fade ranges for a multi-fade (1 - SOC_LEDC_GAMMA_CURVE_FADE_RANGE_MAX)
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Channel not initialized
 *     - ESP_FAIL Fade function init error
 */
//go:linkname LedcSetMultiFade C.ledc_set_multi_fade
func LedcSetMultiFade(speed_mode LedcModeT, channel LedcChannelT, start_duty c.Uint32T, fade_params_list *LedcFadeParamConfigT, list_len c.Uint32T) EspErrT

/**
 * @brief A thread-safe API to set and start LEDC multi-fade function
 *
 * @note  Call `ledc_fade_func_install()` once before calling this function.
 * @note  Fade will always begin from the current duty cycle. Make sure it is stable and synchronized to the desired
 *        initial value before calling this function. Otherwise, you may see unexpected duty change.
 * @note  This function does not prohibit from duty overflow. User should take care of this by themselves. If duty
 *        overflow happens, the PWM signal will suddenly change from 100% duty cycle to 0%, or the other way around.
 *
 * @param speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param channel LEDC channel index (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param start_duty Set the start of the gradient duty, the range of duty setting is [0, (2**duty_resolution)]
 * @param fade_params_list Pointer to the array of fade parameters for a multi-fade
 * @param list_len Length of the fade_params_list, i.e. number of fade ranges for a multi-fade (1 - SOC_LEDC_GAMMA_CURVE_FADE_RANGE_MAX)
 * @param fade_mode Choose blocking or non-blocking mode
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Channel not initialized
 *     - ESP_FAIL Fade function init error
 */
//go:linkname LedcSetMultiFadeAndStart C.ledc_set_multi_fade_and_start
func LedcSetMultiFadeAndStart(speed_mode LedcModeT, channel LedcChannelT, start_duty c.Uint32T, fade_params_list *LedcFadeParamConfigT, list_len c.Uint32T, fade_mode LedcFadeModeT) EspErrT

/**
 * @brief Helper function to fill the fade params for a multi-fade. Useful if desires a gamma curve fading.
 *
 * @note  The fade params are calculated based on the given start_duty and end_duty. If the duty is not at
 *        the start duty (gamma-corrected) when the fade begins, you may see undesired brightness change.
 *        Therefore, please always remember thet when passing the fade_params to either `ledc_set_multi_fade` or
 *        `ledc_set_multi_fade_and start`, the start_duty argument has to be the gamma-corrected start_duty.
 *
 * @param[in] speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param[in] channel LEDC channel index (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param[in] start_duty Duty cycle [0, (2**duty_resolution)] where the multi-fade begins with. This value should be a non-gamma-corrected duty cycle.
 * @param[in] end_duty Duty cycle [0, (2**duty_resolution)] where the multi-fade ends with. This value should be a non-gamma-corrected duty cycle.
 * @param[in] linear_phase_num Number of linear fades to simulate a gamma curved fade (1 - SOC_LEDC_GAMMA_CURVE_FADE_RANGE_MAX)
 * @param[in] max_fade_time_ms The maximum time of the fading ( ms ).
 * @param[in] gamma_correction_operator User provided gamma correction function. The function argument should be able to
 *                                      take any value within [0, (2**duty_resolution)]. And returns the gamma-corrected duty cycle.
 * @param[in] fade_params_list_size The size of the fade_params_list user allocated (1 - SOC_LEDC_GAMMA_CURVE_FADE_RANGE_MAX)
 * @param[out] fade_params_list Pointer to the array of ledc_fade_param_config_t structure
 * @param[out] hw_fade_range_num Number of fade ranges for this multi-fade
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Channel not initialized
 *     - ESP_FAIL Required number of hardware ranges exceeds the size of the ledc_fade_param_config_t array user allocated
 */
//go:linkname LedcFillMultiFadeParamList C.ledc_fill_multi_fade_param_list
func LedcFillMultiFadeParamList(speed_mode LedcModeT, channel LedcChannelT, start_duty c.Uint32T, end_duty c.Uint32T, linear_phase_num c.Uint32T, max_fade_time_ms c.Uint32T, gamma_correction_operator func(c.Uint32T) c.Uint32T, fade_params_list_size c.Uint32T, fade_params_list *LedcFadeParamConfigT, hw_fade_range_num *c.Uint32T) EspErrT

/**
 * @brief Get the fade parameters that are stored in gamma ram for a certain fade range
 *
 * Gamma ram is where saves the fade parameters for each fade range. The fade parameters are written in during fade
 * configuration. When fade begins, the duty will change according to the parameters in gamma ram.
 *
 * @param[in] speed_mode Select the LEDC channel group with specified speed mode. Note that not all targets support high speed mode.
 * @param[in] channel LEDC channel index (0 - LEDC_CHANNEL_MAX-1), select from ledc_channel_t
 * @param[in] range Range index (0 - (SOC_LEDC_GAMMA_CURVE_FADE_RANGE_MAX-1)), it specifies to which range in gamma ram to read
 * @param[out] dir Pointer to accept fade direction value
 * @param[out] cycle Pointer to accept fade cycle value
 * @param[out] scale Pointer to accept fade scale value
 * @param[out] step Pointer to accept fade step value
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Channel not initialized
 */
//go:linkname LedcReadFadeParam C.ledc_read_fade_param
func LedcReadFadeParam(speed_mode LedcModeT, channel LedcChannelT, range_ c.Uint32T, dir *c.Uint32T, cycle *c.Uint32T, scale *c.Uint32T, step *c.Uint32T) EspErrT
