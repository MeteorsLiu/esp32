package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief This function initializes each gpio signal for MCPWM
 *
 * @note This function initializes one gpio at a time.
 *
 * @param mcpwm_num set MCPWM unit
 * @param io_signal set MCPWM signals, each MCPWM unit has 6 output(MCPWMXA, MCPWMXB) and 9 input(SYNC_X, FAULT_X, CAP_X)
 *                  'X' is timer_num(0-2)
 * @param gpio_num set this to configure gpio for MCPWM, if you want to use gpio16, gpio_num = 16
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmGpioInit C.mcpwm_gpio_init
func (recv_ McpwmUnitT) McpwmGpioInit(io_signal McpwmIoSignalsT, gpio_num c.Int) EspErrT {
	return 0
}

/**
 * @brief Initialize MCPWM gpio structure
 *
 * @note This function initialize a group of MCPWM GPIOs at a time.
 *
 * @param mcpwm_num set MCPWM unit
 * @param mcpwm_pin MCPWM pin structure
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmSetPin C.mcpwm_set_pin
func (recv_ McpwmUnitT) McpwmSetPin(mcpwm_pin *McpwmPinConfigT) EspErrT {
	return 0
}

/**
 * @brief Initialize MCPWM parameters
 * @note The default resolution configured for MCPWM timer is 1M, it can be changed by `mcpwm_timer_set_resolution`.
 * @note The default resolution configured for MCPWM group can be different on different esp targets (because of different clock source).
 *       You can change the group resolution by mcpwm_group_set_resolution()
 * @note If you want to change the preset resolution of MCPWM group and timer, please call them before this function.
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers.
 * @param mcpwm_conf configure structure mcpwm_config_t
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmInit C.mcpwm_init
func (recv_ McpwmUnitT) McpwmInit(timer_num McpwmTimerT, mcpwm_conf *McpwmConfigT) EspErrT {
	return 0
}

/**
 * @brief Set resolution of the MCPWM group
 * @note This will override default resolution of MCPWM group.
 * @note This WILL NOT automatically update PWM frequency and duty. Please call `mcpwm_set_frequency` and `mcpwm_set_duty` manually to reflect the change.
 * @note The group resolution must be an integral multiple of timer resolution.
 *
 * @param mcpwm_num set MCPWM unit
 * @param resolution set expected frequency resolution
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmGroupSetResolution C.mcpwm_group_set_resolution
func (recv_ McpwmUnitT) McpwmGroupSetResolution(resolution c.Ulong) EspErrT {
	return 0
}

/**
 * @brief Set resolution of each timer
 * @note This will override default resolution of timer(=1,000,000).
 * @note This WILL NOT automatically update PWM frequency and duty. Please call `mcpwm_set_frequency` and `mcpwm_set_duty` manually to reflect the change.
 * @note The group resolution must be an integral multiple of timer resolution.
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param resolution set expected frequency resolution
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmTimerSetResolution C.mcpwm_timer_set_resolution
func (recv_ McpwmUnitT) McpwmTimerSetResolution(timer_num McpwmTimerT, resolution c.Ulong) EspErrT {
	return 0
}

/**
 * @brief Set frequency(in Hz) of MCPWM timer
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param frequency set the frequency in Hz of each timer
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmSetFrequency C.mcpwm_set_frequency
func (recv_ McpwmUnitT) McpwmSetFrequency(timer_num McpwmTimerT, frequency c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Set duty cycle of each operator(MCPWMXA/MCPWMXB)
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param gen set the generator(MCPWMXA/MCPWMXB), 'X' is operator number selected
 * @param duty set duty cycle in %(i.e for 62.3% duty cycle, duty = 62.3) of each operator
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmSetDuty C.mcpwm_set_duty
func (recv_ McpwmUnitT) McpwmSetDuty(timer_num McpwmTimerT, gen McpwmGeneratorT, duty c.Float) EspErrT {
	return 0
}

/**
 * @brief Set duty cycle of each operator(MCPWMXA/MCPWMXB) in us
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param gen set the generator(MCPWMXA/MCPWMXB), 'x' is operator number selected
 * @param duty_in_us set duty value in microseconds of each operator
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmSetDutyInUs C.mcpwm_set_duty_in_us
func (recv_ McpwmUnitT) McpwmSetDutyInUs(timer_num McpwmTimerT, gen McpwmGeneratorT, duty_in_us c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Set duty either active high or active low(out of phase/inverted)
 * @note
 *        Call this function every time after mcpwm_set_signal_high or mcpwm_set_signal_low to resume with previously set duty cycle
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param gen set the generator(MCPWMXA/MCPWMXB), 'x' is operator number selected
 * @param duty_type set active low or active high duty type
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmSetDutyType C.mcpwm_set_duty_type
func (recv_ McpwmUnitT) McpwmSetDutyType(timer_num McpwmTimerT, gen McpwmGeneratorT, duty_type McpwmDutyTypeT) EspErrT {
	return 0
}

/**
 * @brief Get frequency of timer
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 *
 * @return
 *     - frequency of timer
 */
// llgo:link McpwmUnitT.McpwmGetFrequency C.mcpwm_get_frequency
func (recv_ McpwmUnitT) McpwmGetFrequency(timer_num McpwmTimerT) c.Uint32T {
	return 0
}

/**
 * @brief Get duty cycle of each operator
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param gen set the generator(MCPWMXA/MCPWMXB), 'x' is operator number selected
 *
 * @return
 *     - duty cycle in % of each operator(56.7 means duty is 56.7%)
 */
// llgo:link McpwmUnitT.McpwmGetDuty C.mcpwm_get_duty
func (recv_ McpwmUnitT) McpwmGetDuty(timer_num McpwmTimerT, gen McpwmOperatorT) c.Float {
	return 0
}

/**
 * @brief Get duty cycle of each operator in us
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param gen set the generator(MCPWMXA/MCPWMXB), 'x' is operator number selected
 *
 * @return
 *     - duty cycle in us of each operator
 */
// llgo:link McpwmUnitT.McpwmGetDutyInUs C.mcpwm_get_duty_in_us
func (recv_ McpwmUnitT) McpwmGetDutyInUs(timer_num McpwmTimerT, gen McpwmOperatorT) c.Uint32T {
	return 0
}

/**
 * @brief Use this function to set MCPWM signal high
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param gen set the operator(MCPWMXA/MCPWMXB), 'x' is timer number selected
 *
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmSetSignalHigh C.mcpwm_set_signal_high
func (recv_ McpwmUnitT) McpwmSetSignalHigh(timer_num McpwmTimerT, gen McpwmGeneratorT) EspErrT {
	return 0
}

/**
 * @brief Use this function to set MCPWM signal low
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param gen set the operator(MCPWMXA/MCPWMXB), 'x' is timer number selected
 *
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmSetSignalLow C.mcpwm_set_signal_low
func (recv_ McpwmUnitT) McpwmSetSignalLow(timer_num McpwmTimerT, gen McpwmGeneratorT) EspErrT {
	return 0
}

/**
 * @brief Start MCPWM signal on timer 'x'
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmStart C.mcpwm_start
func (recv_ McpwmUnitT) McpwmStart(timer_num McpwmTimerT) EspErrT {
	return 0
}

/**
 * @brief Start MCPWM signal on timer 'x'
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmStop C.mcpwm_stop
func (recv_ McpwmUnitT) McpwmStop(timer_num McpwmTimerT) EspErrT {
	return 0
}

/**
 * @brief  Initialize carrier configuration
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param carrier_conf configure structure mcpwm_carrier_config_t
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmCarrierInit C.mcpwm_carrier_init
func (recv_ McpwmUnitT) McpwmCarrierInit(timer_num McpwmTimerT, carrier_conf *McpwmCarrierConfigT) EspErrT {
	return 0
}

/**
 * @brief Enable MCPWM carrier submodule, for respective timer
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmCarrierEnable C.mcpwm_carrier_enable
func (recv_ McpwmUnitT) McpwmCarrierEnable(timer_num McpwmTimerT) EspErrT {
	return 0
}

/**
 * @brief Disable MCPWM carrier submodule, for respective timer
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmCarrierDisable C.mcpwm_carrier_disable
func (recv_ McpwmUnitT) McpwmCarrierDisable(timer_num McpwmTimerT) EspErrT {
	return 0
}

/**
 * @brief Set period of carrier
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param carrier_period set the carrier period of each timer, carrier period = (carrier_period + 1)*800ns
 *                    (carrier_period <= 15)
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmCarrierSetPeriod C.mcpwm_carrier_set_period
func (recv_ McpwmUnitT) McpwmCarrierSetPeriod(timer_num McpwmTimerT, carrier_period c.Uint8T) EspErrT {
	return 0
}

/**
 * @brief Set duty_cycle of carrier
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param carrier_duty set duty_cycle of carrier , carrier duty cycle = carrier_duty*12.5%
 *                  (chop_duty <= 7)
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmCarrierSetDutyCycle C.mcpwm_carrier_set_duty_cycle
func (recv_ McpwmUnitT) McpwmCarrierSetDutyCycle(timer_num McpwmTimerT, carrier_duty c.Uint8T) EspErrT {
	return 0
}

/**
 * @brief Enable and set width of first pulse in carrier oneshot mode
 *
 * @note The carrier oneshot pulse can't disabled.
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param pulse_width set pulse width of first pulse in oneshot mode, width = (carrier period)*(pulse_width +1)
 *                    (pulse_width <= 15)
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmCarrierOneshotModeEnable C.mcpwm_carrier_oneshot_mode_enable
func (recv_ McpwmUnitT) McpwmCarrierOneshotModeEnable(timer_num McpwmTimerT, pulse_width c.Uint8T) EspErrT {
	return 0
}

/**
 * @brief Enable or disable carrier output inversion
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param carrier_ivt_mode enable or disable carrier output inversion
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmCarrierOutputInvert C.mcpwm_carrier_output_invert
func (recv_ McpwmUnitT) McpwmCarrierOutputInvert(timer_num McpwmTimerT, carrier_ivt_mode McpwmCarrierOutIvtT) EspErrT {
	return 0
}

/**
 * @brief Enable and initialize deadtime for each MCPWM timer
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param dt_mode set deadtime mode
 * @param red set rising edge delay = (red + 1) * MCPWM Group Resolution (default to 100ns, can be changed by `mcpwm_group_set_resolution`)
 * @param fed set rising edge delay = (fed + 1) * MCPWM Group Resolution (default to 100ns, can be changed by `mcpwm_group_set_resolution`)
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmDeadtimeEnable C.mcpwm_deadtime_enable
func (recv_ McpwmUnitT) McpwmDeadtimeEnable(timer_num McpwmTimerT, dt_mode McpwmDeadtimeTypeT, red c.Uint32T, fed c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Disable deadtime on MCPWM timer
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmDeadtimeDisable C.mcpwm_deadtime_disable
func (recv_ McpwmUnitT) McpwmDeadtimeDisable(timer_num McpwmTimerT) EspErrT {
	return 0
}

/**
 * @brief Initialize fault submodule, currently low level triggering is not supported
 *
 * @param mcpwm_num set MCPWM unit
 * @param intput_level set fault signal level, which will cause fault to occur
 * @param fault_sig set the fault pin, which needs to be enabled
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmFaultInit C.mcpwm_fault_init
func (recv_ McpwmUnitT) McpwmFaultInit(intput_level McpwmFaultInputLevelT, fault_sig McpwmFaultSignalT) EspErrT {
	return 0
}

/**
 * @brief Set oneshot mode on fault detection, once fault occur in oneshot mode reset is required to resume MCPWM signals
 * @note
 *        currently low level triggering is not supported
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param fault_sig set the fault pin, which needs to be enabled for oneshot mode
 * @param action_on_pwmxa action to be taken on MCPWMXA when fault occurs, either no change or high or low or toggle
 * @param action_on_pwmxb action to be taken on MCPWMXB when fault occurs, either no change or high or low or toggle
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmFaultSetOneshotMode C.mcpwm_fault_set_oneshot_mode
func (recv_ McpwmUnitT) McpwmFaultSetOneshotMode(timer_num McpwmTimerT, fault_sig McpwmFaultSignalT, action_on_pwmxa McpwmOutputActionT, action_on_pwmxb McpwmOutputActionT) EspErrT {
	return 0
}

/**
 * @brief Set cycle-by-cycle mode on fault detection, once fault occur in cyc mode MCPWM signal resumes as soon as fault signal becomes inactive
 * @note
 *        currently low level triggering is not supported
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param fault_sig set the fault pin, which needs to be enabled for cyc mode
 * @param action_on_pwmxa action to be taken on MCPWMXA when fault occurs, either no change or high or low or toggle
 * @param action_on_pwmxb action to be taken on MCPWMXB when fault occurs, either no change or high or low or toggle
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmFaultSetCycMode C.mcpwm_fault_set_cyc_mode
func (recv_ McpwmUnitT) McpwmFaultSetCycMode(timer_num McpwmTimerT, fault_sig McpwmFaultSignalT, action_on_pwmxa McpwmOutputActionT, action_on_pwmxb McpwmOutputActionT) EspErrT {
	return 0
}

/**
 * @brief Disable fault signal
 *
 * @param mcpwm_num set MCPWM unit
 * @param fault_sig fault pin, which needs to be disabled
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmFaultDeinit C.mcpwm_fault_deinit
func (recv_ McpwmUnitT) McpwmFaultDeinit(fault_sig McpwmFaultSignalT) EspErrT {
	return 0
}

/**
 * @brief Enable capture channel
 *
 * @param mcpwm_num set MCPWM unit
 * @param cap_channel capture channel, which needs to be enabled
 * @param cap_conf capture channel configuration
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmCaptureEnableChannel C.mcpwm_capture_enable_channel
func (recv_ McpwmUnitT) McpwmCaptureEnableChannel(cap_channel McpwmCaptureChannelIdT, cap_conf *McpwmCaptureConfigT) EspErrT {
	return 0
}

/**
 * @brief Disable capture channel
 *
 * @param mcpwm_num set MCPWM unit
 * @param cap_channel capture channel, which needs to be disabled
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmCaptureDisableChannel C.mcpwm_capture_disable_channel
func (recv_ McpwmUnitT) McpwmCaptureDisableChannel(cap_channel McpwmCaptureChannelIdT) EspErrT {
	return 0
}

/**
 * @brief Get capture value
 *
 * @param mcpwm_num set MCPWM unit
 * @param cap_sig capture channel on which value is to be measured
 *
 * @return
 *     Captured value
 */
// llgo:link McpwmUnitT.McpwmCaptureSignalGetValue C.mcpwm_capture_signal_get_value
func (recv_ McpwmUnitT) McpwmCaptureSignalGetValue(cap_sig McpwmCaptureSignalT) c.Uint32T {
	return 0
}

/**
 * @brief Get capture timer's resolution
 *
 * @param mcpwm_num set MCPWM unit
 * @return Capture timer's resolution
 */
// llgo:link McpwmUnitT.McpwmCaptureGetResolution C.mcpwm_capture_get_resolution
func (recv_ McpwmUnitT) McpwmCaptureGetResolution() c.Uint32T {
	return 0
}

/**
 * @brief Get edge of capture signal
 *
 * @param mcpwm_num set MCPWM unit
 * @param cap_sig capture channel of whose edge is to be determined
 *
 * @return
 *     Capture signal edge: 1 - positive edge, 2 - negative edge, 0 - Invalid
 */
// llgo:link McpwmUnitT.McpwmCaptureSignalGetEdge C.mcpwm_capture_signal_get_edge
func (recv_ McpwmUnitT) McpwmCaptureSignalGetEdge(cap_sig McpwmCaptureSignalT) c.Uint32T {
	return 0
}

/**
 * @brief Initialize sync submodule and sets the signal that will cause the timer be loaded with pre-defined value
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param sync_conf sync configuration on this timer
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmSyncConfigure C.mcpwm_sync_configure
func (recv_ McpwmUnitT) McpwmSyncConfigure(timer_num McpwmTimerT, sync_conf *McpwmSyncConfigT) EspErrT {
	return 0
}

/**
 * @brief Disable sync submodule on given timer
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmSyncDisable C.mcpwm_sync_disable
func (recv_ McpwmUnitT) McpwmSyncDisable(timer_num McpwmTimerT) EspErrT {
	return 0
}

/**
 * @brief Set sync output on given timer
 *        Configures what event triggers MCPWM timer to output a sync signal.
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 * @param trigger set the trigger that will cause the timer to generate a software sync signal.
 *                Specifically, `MCPWM_SWSYNC_SOURCE_DISABLED` will disable the timer from generating sync signal.
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link McpwmUnitT.McpwmSetTimerSyncOutput C.mcpwm_set_timer_sync_output
func (recv_ McpwmUnitT) McpwmSetTimerSyncOutput(timer_num McpwmTimerT, trigger McpwmTimerSyncTriggerT) EspErrT {
	return 0
}

/**
 * @brief Trigger a software sync event and sends it to a specific timer.
 *
 * @param mcpwm_num set MCPWM unit
 * @param timer_num set timer number(0-2) of MCPWM, each MCPWM unit has 3 timers
 *
 * @note This software sync event will have the same effect as hw one, except that:
 *         - On esp32s3 the soft sync event can be routed to its output if `MCPWM_SWSYNC_SOURCE_SYNCIN` is selected via `mcpwm_set_timer_sync_output()`
 *         - On esp32 there is no such behavior and soft sync event will only take effect on this timer and can not be propagated to others.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Function pointer error.
 */
// llgo:link McpwmUnitT.McpwmTimerTriggerSoftSync C.mcpwm_timer_trigger_soft_sync
func (recv_ McpwmUnitT) McpwmTimerTriggerSoftSync(timer_num McpwmTimerT) EspErrT {
	return 0
}

/**
 * @brief Set external GPIO sync input inverter
 *
 * @param mcpwm_num set MCPWM unit
 * @param sync_sig set sync signal of MCPWM, only supports GPIO sync signal
 * @param invert whether GPIO sync source input is inverted (to get negative edge trigger)
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Function pointer error.
 */
// llgo:link McpwmUnitT.McpwmSyncInvertGpioSynchro C.mcpwm_sync_invert_gpio_synchro
func (recv_ McpwmUnitT) McpwmSyncInvertGpioSynchro(sync_sig McpwmSyncSignalT, invert bool) EspErrT {
	return 0
}
