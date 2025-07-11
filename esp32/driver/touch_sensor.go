package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Configure touch pad interrupt threshold.
 *
 * @note  If FSM mode is set to TOUCH_FSM_MODE_TIMER, this function will be blocked for one measurement cycle and wait for data to be valid.
 *
 * @param touch_num touch pad index
 * @param threshold interrupt threshold,
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG if argument wrong
 *     - ESP_FAIL if touch pad not initialized
 */
//go:linkname TouchPadConfig C.touch_pad_config
func TouchPadConfig(touch_num TouchPadT, threshold c.Uint16T) EspErrT

/**
 * @brief get touch sensor counter value.
 *        Each touch sensor has a counter to count the number of charge/discharge cycles.
 *        When the pad is not 'touched', we can get a number of the counter.
 *        When the pad is 'touched', the value in counter will get smaller because of the larger equivalent capacitance.
 *
 * @note This API requests hardware measurement once. If IIR filter mode is enabled,
 *       please use 'touch_pad_read_raw_data' interface instead.
 *
 * @param touch_num touch pad index
 * @param touch_value pointer to accept touch sensor value
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Touch pad parameter error
 *     - ESP_ERR_INVALID_STATE This touch pad hardware connection is error, the value of "touch_value" is 0.
 *     - ESP_FAIL Touch pad not initialized
 */
//go:linkname TouchPadRead C.touch_pad_read
func TouchPadRead(touch_num TouchPadT, touch_value *c.Uint16T) EspErrT

/**
 * @brief get filtered touch sensor counter value by IIR filter.
 *
 * @note touch_pad_filter_start has to be called before calling touch_pad_read_filtered.
 *       This function can be called from ISR
 *
 * @param touch_num touch pad index
 * @param touch_value pointer to accept touch sensor value
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Touch pad parameter error
 *     - ESP_ERR_INVALID_STATE This touch pad hardware connection is error, the value of "touch_value" is 0.
 *     - ESP_FAIL Touch pad not initialized
 */
//go:linkname TouchPadReadFiltered C.touch_pad_read_filtered
func TouchPadReadFiltered(touch_num TouchPadT, touch_value *c.Uint16T) EspErrT

/**
 * @brief get raw data (touch sensor counter value) from IIR filter process.
 *        Need not request hardware measurements.
 *
 * @note touch_pad_filter_start has to be called before calling touch_pad_read_raw_data.
 *       This function can be called from ISR
 *
 * @param touch_num touch pad index
 * @param touch_value pointer to accept touch sensor value
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Touch pad parameter error
 *     - ESP_ERR_INVALID_STATE This touch pad hardware connection is error, the value of "touch_value" is 0.
 *     - ESP_FAIL Touch pad not initialized
 */
//go:linkname TouchPadReadRawData C.touch_pad_read_raw_data
func TouchPadReadRawData(touch_num TouchPadT, touch_value *c.Uint16T) EspErrT

// llgo:type C
type FilterCbT func(*c.Uint16T, *c.Uint16T)

/**
 * @brief Register the callback function that is called after each IIR filter calculation.
 * @note The 'read_cb' callback is called in timer task in each filtering cycle.
 * @param read_cb  Pointer to filtered callback function.
 *                 If the argument passed in is NULL, the callback will stop.
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_ARG set error
 */
//go:linkname TouchPadSetFilterReadCb C.touch_pad_set_filter_read_cb
func TouchPadSetFilterReadCb(read_cb FilterCbT) EspErrT

/**
 * @brief   Register touch-pad ISR.
 *          The handler will be attached to the same CPU core that this function is running on.
 * @param fn  Pointer to ISR handler
 * @param arg  Parameter for ISR
 * @return
 *     - ESP_OK Success ;
 *     - ESP_ERR_INVALID_ARG GPIO error
 *     - ESP_ERR_NO_MEM No memory
 */
//go:linkname TouchPadIsrRegister C.touch_pad_isr_register
func TouchPadIsrRegister(fn IntrHandlerT, arg c.Pointer) EspErrT

/**
 * @brief Set the clock cycles of each measurement
 * @note  This function will specify the clock cycles of each measurement
 *        and the clock is sourced from SOC_MOD_CLK_RTC_FAST, its default frequency is SOC_CLK_RC_FAST_FREQ_APPROX
 *        The touch sensor will record the charge and discharge times during these clock cycles as the final result (raw value)
 * @note  If clock cycles is too small, it may lead to inaccurate results.
 *
 * @param clock_cycle   The clock cycles of each measurement
 *                      measure_time = clock_cycle / SOC_CLK_RC_FAST_FREQ_APPROX, the maximum measure time is 0xffff / SOC_CLK_RC_FAST_FREQ_APPROX
 * @return
 *      - ESP_OK    Set the clock cycle success
 */
//go:linkname TouchPadSetMeasurementClockCycles C.touch_pad_set_measurement_clock_cycles
func TouchPadSetMeasurementClockCycles(clock_cycle c.Uint16T) EspErrT

/**
 * @brief Get the clock cycles of each measurement
 *
 * @param clock_cycle   The clock cycles of each measurement
 * @return
 *      - ESP_OK    Get the clock cycle success
 *      - ESP_ERR_INVALID_ARG The input parameter is NULL
 */
//go:linkname TouchPadGetMeasurementClockCycles C.touch_pad_get_measurement_clock_cycles
func TouchPadGetMeasurementClockCycles(clock_cycle *c.Uint16T) EspErrT

/**
 * @brief Set the interval between two measurements
 * @note  The touch sensor will sleep between two measurements
 *        This function is to set the interval cycle
 *        And the interval is clocked from SOC_MOD_CLK_RTC_SLOW, its default frequency is SOC_CLK_RC_SLOW_FREQ_APPROX
 *
 * @param interval_cycle    The interval between two measurements
 *                          sleep_time = interval_cycle / SOC_CLK_RC_SLOW_FREQ_APPROX.
 *                          The approximate frequency value of RTC_SLOW_CLK can be obtained using rtc_clk_slow_freq_get_hz function.
 * @return
 *      - ESP_OK    Set interval cycle success
 */
//go:linkname TouchPadSetMeasurementInterval C.touch_pad_set_measurement_interval
func TouchPadSetMeasurementInterval(interval_cycle c.Uint16T) EspErrT

/**
 * @brief Get the interval between two measurements
 *
 * @param interval_cycle    The interval between two measurements
 * @return
 *      - ESP_OK    Get interval cycle success
 *      - ESP_ERR_INVALID_ARG   The input parameter is NULL
 */
//go:linkname TouchPadGetMeasurementInterval C.touch_pad_get_measurement_interval
func TouchPadGetMeasurementInterval(interval_cycle *c.Uint16T) EspErrT

/**
 * @brief Set touch sensor measurement and sleep time.
 *        Excessive total time will slow down the touch response.
 *        Too small measurement time will not be sampled enough, resulting in inaccurate measurements.
 * @note The touch sensor will count the number of charge/discharge cycles over a fixed period of time (specified as the second parameter).
 *       That means the number of cycles (raw value) will decrease as the capacity of the touch pad is increasing.
 * @note The greater the duty cycle of the measurement time, the more system power is consumed.
 *
 * @param sleep_cycle  The touch sensor will sleep after each measurement.
 *                     sleep_cycle decide the interval between each measurement.
 *                     t_sleep = sleep_cycle / SOC_CLK_RC_SLOW_FREQ_APPROX.
 *                     The approximate frequency value of RTC_SLOW_CLK can be obtained using rtc_clk_slow_freq_get_hz function.
 * @param meas_cycle The duration of the touch sensor measurement.
 *                   t_meas = meas_cycle / SOC_CLK_RC_FAST_FREQ_APPROX, the maximum measure time is 0xffff / SOC_CLK_RC_FAST_FREQ_APPROX
 * @return
 *      - ESP_OK on success
 */
//go:linkname TouchPadSetMeasTime C.touch_pad_set_meas_time
func TouchPadSetMeasTime(sleep_cycle c.Uint16T, meas_cycle c.Uint16T) EspErrT

/**
 * @brief Get touch sensor measurement and sleep time
 * @param sleep_cycle  Pointer to accept sleep cycle number
 * @param meas_cycle Pointer to accept measurement cycle count.
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG   The input parameter is NULL
 */
//go:linkname TouchPadGetMeasTime C.touch_pad_get_meas_time
func TouchPadGetMeasTime(sleep_cycle *c.Uint16T, meas_cycle *c.Uint16T) EspErrT

/**
 * @brief Trigger a touch sensor measurement, only support in SW mode of FSM
 * @return
 *      - ESP_OK on success
 */
//go:linkname TouchPadSwStart C.touch_pad_sw_start
func TouchPadSwStart() EspErrT

/**
 * @brief Set touch sensor interrupt threshold
 * @param touch_num touch pad index
 * @param threshold threshold of touchpad count, refer to touch_pad_set_trigger_mode to see how to set trigger mode.
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if argument is wrong
 */
//go:linkname TouchPadSetThresh C.touch_pad_set_thresh
func TouchPadSetThresh(touch_num TouchPadT, threshold c.Uint16T) EspErrT

/**
 * @brief Get touch sensor interrupt threshold
 * @param touch_num touch pad index
 * @param threshold pointer to accept threshold
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if argument is wrong
 */
//go:linkname TouchPadGetThresh C.touch_pad_get_thresh
func TouchPadGetThresh(touch_num TouchPadT, threshold *c.Uint16T) EspErrT

/**
 * @brief Set touch sensor interrupt trigger mode.
 *        Interrupt can be triggered either when counter result is less than
 *        threshold or when counter result is more than threshold.
 * @param mode touch sensor interrupt trigger mode
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if argument is wrong
 */
//go:linkname TouchPadSetTriggerMode C.touch_pad_set_trigger_mode
func TouchPadSetTriggerMode(mode TouchTriggerModeT) EspErrT

/**
 * @brief Get touch sensor interrupt trigger mode
 * @param mode pointer to accept touch sensor interrupt trigger mode
 * @return
 *      - ESP_OK on success
 */
//go:linkname TouchPadGetTriggerMode C.touch_pad_get_trigger_mode
func TouchPadGetTriggerMode(mode *TouchTriggerModeT) EspErrT

/**
 * @brief Set touch sensor interrupt trigger source. There are two sets of touch signals.
 *        Set1 and set2 can be mapped to several touch signals. Either set will be triggered
 *        if at least one of its touch signal is 'touched'. The interrupt can be configured to be generated
 *        if set1 is triggered, or only if both sets are triggered.
 * @param src touch sensor interrupt trigger source
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if argument is wrong
 */
//go:linkname TouchPadSetTriggerSource C.touch_pad_set_trigger_source
func TouchPadSetTriggerSource(src TouchTriggerSrcT) EspErrT

/**
 * @brief Get touch sensor interrupt trigger source
 * @param src pointer to accept touch sensor interrupt trigger source
 * @return
 *      - ESP_OK on success
 */
//go:linkname TouchPadGetTriggerSource C.touch_pad_get_trigger_source
func TouchPadGetTriggerSource(src *TouchTriggerSrcT) EspErrT

/**
 * @brief Set touch sensor group mask.
 *        Touch pad module has two sets of signals, 'Touched' signal is triggered only if
 *        at least one of touch pad in this group is "touched".
 *        This function will set the register bits according to the given bitmask.
 * @param set1_mask bitmask of touch sensor signal group1, it's a 10-bit value
 * @param set2_mask bitmask of touch sensor signal group2, it's a 10-bit value
 * @param en_mask bitmask of touch sensor work enable, it's a 10-bit value
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if argument is wrong
 */
//go:linkname TouchPadSetGroupMask C.touch_pad_set_group_mask
func TouchPadSetGroupMask(set1_mask c.Uint16T, set2_mask c.Uint16T, en_mask c.Uint16T) EspErrT

/**
 * @brief Get touch sensor group mask.
 * @param set1_mask pointer to accept bitmask of touch sensor signal group1, it's a 10-bit value
 * @param set2_mask pointer to accept bitmask of touch sensor signal group2, it's a 10-bit value
 * @param en_mask pointer to accept bitmask of touch sensor work enable, it's a 10-bit value
 * @return
 *      - ESP_OK on success
 */
//go:linkname TouchPadGetGroupMask C.touch_pad_get_group_mask
func TouchPadGetGroupMask(set1_mask *c.Uint16T, set2_mask *c.Uint16T, en_mask *c.Uint16T) EspErrT

/**
 * @brief Clear touch sensor group mask.
 *        Touch pad module has two sets of signals, Interrupt is triggered only if
 *        at least one of touch pad in this group is "touched".
 *        This function will clear the register bits according to the given bitmask.
 * @param set1_mask bitmask touch sensor signal group1, it's a 10-bit value
 * @param set2_mask bitmask touch sensor signal group2, it's a 10-bit value
 * @param en_mask bitmask of touch sensor work enable, it's a 10-bit value
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if argument is wrong
 */
//go:linkname TouchPadClearGroupMask C.touch_pad_clear_group_mask
func TouchPadClearGroupMask(set1_mask c.Uint16T, set2_mask c.Uint16T, en_mask c.Uint16T) EspErrT

/**
 * @brief To enable touch pad interrupt
 * @return
 *      - ESP_OK on success
 */
//go:linkname TouchPadIntrEnable C.touch_pad_intr_enable
func TouchPadIntrEnable() EspErrT

/**
 * @brief To disable touch pad interrupt
 * @return
 *      - ESP_OK on success
 */
//go:linkname TouchPadIntrDisable C.touch_pad_intr_disable
func TouchPadIntrDisable() EspErrT

/**
 * @brief To clear touch pad interrupt
 * @return
 *      - ESP_OK on success
 */
//go:linkname TouchPadIntrClear C.touch_pad_intr_clear
func TouchPadIntrClear() EspErrT

/**
 * @brief set touch pad filter calibration period, in ms.
 *        Need to call touch_pad_filter_start before all touch filter APIs
 * @param new_period_ms filter period, in ms
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_STATE driver state error
 *      - ESP_ERR_INVALID_ARG parameter error
 */
//go:linkname TouchPadSetFilterPeriod C.touch_pad_set_filter_period
func TouchPadSetFilterPeriod(new_period_ms c.Uint32T) EspErrT

/**
 * @brief get touch pad filter calibration period, in ms
 *        Need to call touch_pad_filter_start before all touch filter APIs
 * @param p_period_ms pointer to accept period
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_STATE driver state error
 *      - ESP_ERR_INVALID_ARG parameter error
 */
//go:linkname TouchPadGetFilterPeriod C.touch_pad_get_filter_period
func TouchPadGetFilterPeriod(p_period_ms *c.Uint32T) EspErrT

/**
 * @brief start touch pad filter function
 *      This API will start a filter to process the noise in order to prevent false triggering
 *      when detecting slight change of capacitance.
 *      Need to call touch_pad_filter_start before all touch filter APIs
 *
 * @note This filter uses FreeRTOS timer, which is dispatched from a task with
 *       priority 1 by default on CPU 0. So if some application task with higher priority
 *       takes a lot of CPU0 time, then the quality of data obtained from this filter will be affected.
 *       You can adjust FreeRTOS timer task priority in menuconfig.
 * @param filter_period_ms filter calibration period, in ms
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_ARG parameter error
 *      - ESP_ERR_NO_MEM No memory for driver
 *      - ESP_ERR_INVALID_STATE driver state error
 */
//go:linkname TouchPadFilterStart C.touch_pad_filter_start
func TouchPadFilterStart(filter_period_ms c.Uint32T) EspErrT

/**
 * @brief stop touch pad filter function
 *        Need to call touch_pad_filter_start before all touch filter APIs
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_STATE driver state error
 */
//go:linkname TouchPadFilterStop C.touch_pad_filter_stop
func TouchPadFilterStop() EspErrT

/**
 * @brief delete touch pad filter driver and release the memory
 *        Need to call touch_pad_filter_start before all touch filter APIs
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_STATE driver state error
 */
//go:linkname TouchPadFilterDelete C.touch_pad_filter_delete
func TouchPadFilterDelete() EspErrT
