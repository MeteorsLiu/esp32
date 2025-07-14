package esp_driver_pcnt

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PcntUnitT struct {
	Unused [8]uint8
}
type PcntUnitHandleT *PcntUnitT

type PcntChanT struct {
	Unused [8]uint8
}
type PcntChannelHandleT *PcntChanT

/**
 * @brief PCNT watch event data
 */

type PcntWatchEventDataT struct {
	WatchPointValue c.Int
	ZeroCrossMode   PcntUnitZeroCrossModeT
}

// llgo:type C
type PcntWatchCbT func(PcntUnitHandleT, *PcntWatchEventDataT, c.Pointer) bool

/**
 * @brief Group of supported PCNT callbacks
 * @note The callbacks are all running under ISR environment
 * @note When CONFIG_PCNT_ISR_IRAM_SAFE is enabled, the callback itself and functions callbed by it should be placed in IRAM.
 */

type PcntEventCallbacksT struct {
	OnReach PcntWatchCbT
}

/**
 * @brief PCNT unit configuration
 */

type PcntUnitConfigT struct {
	LowLimit     c.Int
	HighLimit    c.Int
	IntrPriority c.Int
	Flags        struct {
		AccumCount c.Uint32T
	}
}

/**
 * @brief PCNT channel configuration
 */

type PcntChanConfigT struct {
	EdgeGpioNum  c.Int
	LevelGpioNum c.Int
	Flags        struct {
		InvertEdgeInput  c.Uint32T
		InvertLevelInput c.Uint32T
		VirtEdgeIoLevel  c.Uint32T
		VirtLevelIoLevel c.Uint32T
		IoLoopBack       c.Uint32T
	}
}

/**
 * @brief PCNT glitch filter configuration
 */

type PcntGlitchFilterConfigT struct {
	MaxGlitchNs c.Uint32T
}

/**
 * @brief Create a new PCNT unit, and return the handle
 *
 * @note The newly created PCNT unit is put in the init state.
 *
 * @param[in] config PCNT unit configuration
 * @param[out] ret_unit Returned PCNT unit handle
 * @return
 *      - ESP_OK: Create PCNT unit successfully
 *      - ESP_ERR_INVALID_ARG: Create PCNT unit failed because of invalid argument (e.g. high/low limit value out of the range)
 *      - ESP_ERR_NO_MEM: Create PCNT unit failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create PCNT unit failed because all PCNT units are used up and no more free one
 *      - ESP_FAIL: Create PCNT unit failed because of other error
 */
// llgo:link (*PcntUnitConfigT).PcntNewUnit C.pcnt_new_unit
func (recv_ *PcntUnitConfigT) PcntNewUnit(ret_unit *PcntUnitHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete the PCNT unit handle
 *
 * @note A PCNT unit can't be in the enable state when this function is invoked.
 *       See also `pcnt_unit_disable()` for how to disable a unit.
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @return
 *      - ESP_OK: Delete the PCNT unit successfully
 *      - ESP_ERR_INVALID_ARG: Delete the PCNT unit failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Delete the PCNT unit failed because the unit is not in init state or some PCNT channel is still in working
 *      - ESP_FAIL: Delete the PCNT unit failed because of other error
 */
//go:linkname PcntDelUnit C.pcnt_del_unit
func PcntDelUnit(unit PcntUnitHandleT) EspErrT

/**
 * @brief Set glitch filter for PCNT unit
 *
 * @note This function should be called when the PCNT unit is in the init state (i.e. before calling `pcnt_unit_enable()`)
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @param[in] config PCNT filter configuration, set config to NULL means disabling the filter function
 * @return
 *      - ESP_OK: Set glitch filter successfully
 *      - ESP_ERR_INVALID_ARG: Set glitch filter failed because of invalid argument (e.g. glitch width is too big)
 *      - ESP_ERR_INVALID_STATE: Set glitch filter failed because the unit is not in the init state
 *      - ESP_FAIL: Set glitch filter failed because of other error
 */
//go:linkname PcntUnitSetGlitchFilter C.pcnt_unit_set_glitch_filter
func PcntUnitSetGlitchFilter(unit PcntUnitHandleT, config *PcntGlitchFilterConfigT) EspErrT

/**
 * @brief Enable the PCNT unit
 *
 * @note This function will transit the unit state from init to enable.
 * @note This function will enable the interrupt service, if it's lazy installed in `pcnt_unit_register_event_callbacks()`.
 * @note This function will acquire the PM lock if it's lazy installed in `pcnt_unit_set_glitch_filter()`.
 * @note Enable a PCNT unit doesn't mean to start it. See also `pcnt_unit_start()` for how to start the PCNT counter.
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @return
 *      - ESP_OK: Enable PCNT unit successfully
 *      - ESP_ERR_INVALID_ARG: Enable PCNT unit failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Enable PCNT unit failed because the unit is already enabled
 *      - ESP_FAIL: Enable PCNT unit failed because of other error
 */
//go:linkname PcntUnitEnable C.pcnt_unit_enable
func PcntUnitEnable(unit PcntUnitHandleT) EspErrT

/**
 * @brief Disable the PCNT unit
 *
 * @note This function will do the opposite work to the `pcnt_unit_enable()`
 * @note Disable a PCNT unit doesn't mean to stop it. See also `pcnt_unit_stop()` for how to stop the PCNT counter.
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @return
 *      - ESP_OK: Disable PCNT unit successfully
 *      - ESP_ERR_INVALID_ARG: Disable PCNT unit failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Disable PCNT unit failed because the unit is not enabled yet
 *      - ESP_FAIL: Disable PCNT unit failed because of other error
 */
//go:linkname PcntUnitDisable C.pcnt_unit_disable
func PcntUnitDisable(unit PcntUnitHandleT) EspErrT

/**
 * @brief Start the PCNT unit, the counter will start to count according to the edge and/or level input signals
 *
 * @note This function should be called when the unit is in the enable state (i.e. after calling `pcnt_unit_enable()`)
 * @note This function is allowed to run within ISR context
 * @note This function will be placed into IRAM if `CONFIG_PCNT_CTRL_FUNC_IN_IRAM` is on, so that it's allowed to be executed when Cache is disabled
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @return
 *      - ESP_OK: Start PCNT unit successfully
 *      - ESP_ERR_INVALID_ARG: Start PCNT unit failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Start PCNT unit failed because the unit is not enabled yet
 *      - ESP_FAIL: Start PCNT unit failed because of other error
 */
//go:linkname PcntUnitStart C.pcnt_unit_start
func PcntUnitStart(unit PcntUnitHandleT) EspErrT

/**
 * @brief Stop PCNT from counting
 *
 * @note This function should be called when the unit is in the enable state (i.e. after calling `pcnt_unit_enable()`)
 * @note The stop operation won't clear the counter. Also see `pcnt_unit_clear_count()` for how to clear pulse count value.
 * @note This function is allowed to run within ISR context
 * @note This function will be placed into IRAM if `CONFIG_PCNT_CTRL_FUNC_IN_IRAM`, so that it is allowed to be executed when Cache is disabled
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @return
 *      - ESP_OK: Stop PCNT unit successfully
 *      - ESP_ERR_INVALID_ARG: Stop PCNT unit failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Stop PCNT unit failed because the unit is not enabled yet
 *      - ESP_FAIL: Stop PCNT unit failed because of other error
 */
//go:linkname PcntUnitStop C.pcnt_unit_stop
func PcntUnitStop(unit PcntUnitHandleT) EspErrT

/**
 * @brief Clear PCNT pulse count value to zero
 *
 * @note It's recommended to call this function after adding a watch point by `pcnt_unit_add_watch_point()`, so that the newly added watch point is effective immediately.
 * @note This function is allowed to run within ISR context
 * @note This function will be placed into IRAM if `CONFIG_PCNT_CTRL_FUNC_IN_IRAM`, so that it's allowed to be executed when Cache is disabled
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @return
 *      - ESP_OK: Clear PCNT pulse count successfully
 *      - ESP_ERR_INVALID_ARG: Clear PCNT pulse count failed because of invalid argument
 *      - ESP_FAIL: Clear PCNT pulse count failed because of other error
 */
//go:linkname PcntUnitClearCount C.pcnt_unit_clear_count
func PcntUnitClearCount(unit PcntUnitHandleT) EspErrT

/**
 * @brief Get PCNT count value
 *
 * @note This function is allowed to run within ISR context
 * @note This function will be placed into IRAM if `CONFIG_PCNT_CTRL_FUNC_IN_IRAM`, so that it's allowed to be executed when Cache is disabled
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @param[out] value Returned count value
 * @return
 *      - ESP_OK: Get PCNT pulse count successfully
 *      - ESP_ERR_INVALID_ARG: Get PCNT pulse count failed because of invalid argument
 *      - ESP_FAIL: Get PCNT pulse count failed because of other error
 */
//go:linkname PcntUnitGetCount C.pcnt_unit_get_count
func PcntUnitGetCount(unit PcntUnitHandleT, value *c.Int) EspErrT

/**
 * @brief Set event callbacks for PCNT unit
 *
 * @note User registered callbacks are expected to be runnable within ISR context
 * @note The first call to this function needs to be before the call to `pcnt_unit_enable`
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Set event callbacks failed because the unit is not in init state
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname PcntUnitRegisterEventCallbacks C.pcnt_unit_register_event_callbacks
func PcntUnitRegisterEventCallbacks(unit PcntUnitHandleT, cbs *PcntEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Add a watch point for PCNT unit, PCNT will generate an event when the counter value reaches the watch point value
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @param[in] watch_point Value to be watched
 * @return
 *      - ESP_OK: Add watch point successfully
 *      - ESP_ERR_INVALID_ARG: Add watch point failed because of invalid argument (e.g. the value to be watched is out of the limitation set in `pcnt_unit_config_t`)
 *      - ESP_ERR_INVALID_STATE: Add watch point failed because the same watch point has already been added
 *      - ESP_ERR_NOT_FOUND: Add watch point failed because no more hardware watch point can be configured
 *      - ESP_FAIL: Add watch point failed because of other error
 */
//go:linkname PcntUnitAddWatchPoint C.pcnt_unit_add_watch_point
func PcntUnitAddWatchPoint(unit PcntUnitHandleT, watch_point c.Int) EspErrT

/**
 * @brief Remove a watch point for PCNT unit
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @param[in] watch_point Watch point value
 * @return
 *      - ESP_OK: Remove watch point successfully
 *      - ESP_ERR_INVALID_ARG: Remove watch point failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Remove watch point failed because the watch point was not added by `pcnt_unit_add_watch_point()` yet
 *      - ESP_FAIL: Remove watch point failed because of other error
 */
//go:linkname PcntUnitRemoveWatchPoint C.pcnt_unit_remove_watch_point
func PcntUnitRemoveWatchPoint(unit PcntUnitHandleT, watch_point c.Int) EspErrT

/**
 * @brief Add a step notify for PCNT unit, PCNT will generate an event when the incremental(can be positive or negative) of counter value reaches the step interval
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @param[in] step_interval PCNT step notify interval value
 * @return
 *      - ESP_OK: Add step notify successfully
 *      - ESP_ERR_INVALID_ARG: Add step notify failed because of invalid argument (e.g. the value incremental to be watched is out of the limitation set in `pcnt_unit_config_t`)
 *      - ESP_ERR_INVALID_STATE: Add step notify failed because the step notify has already been added
 *      - ESP_FAIL: Add step notify failed because of other error
 */
//go:linkname PcntUnitAddWatchStep C.pcnt_unit_add_watch_step
func PcntUnitAddWatchStep(unit PcntUnitHandleT, step_interval c.Int) EspErrT

/**
 * @brief Remove a step notify for PCNT unit
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @return
 *      - ESP_OK: Remove step notify successfully
 *      - ESP_ERR_INVALID_ARG: Remove step notify failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Remove step notify failed because the step notify was not added by `pcnt_unit_add_watch_step()` yet
 *      - ESP_FAIL: Remove step notify failed because of other error
 */
//go:linkname PcntUnitRemoveWatchStep C.pcnt_unit_remove_watch_step
func PcntUnitRemoveWatchStep(unit PcntUnitHandleT) EspErrT

/**
 * @brief Create PCNT channel for specific unit, each PCNT has several channels associated with it
 *
 * @note This function should be called when the unit is in init state (i.e. before calling `pcnt_unit_enable()`)
 *
 * @param[in] unit PCNT unit handle created by `pcnt_new_unit()`
 * @param[in] config PCNT channel configuration
 * @param[out] ret_chan Returned channel handle
 * @return
 *      - ESP_OK: Create PCNT channel successfully
 *      - ESP_ERR_INVALID_ARG: Create PCNT channel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create PCNT channel failed because of insufficient memory
 *      - ESP_ERR_NOT_FOUND: Create PCNT channel failed because all PCNT channels are used up and no more free one
 *      - ESP_ERR_INVALID_STATE: Create PCNT channel failed because the unit is not in the init state
 *      - ESP_FAIL: Create PCNT channel failed because of other error
 */
//go:linkname PcntNewChannel C.pcnt_new_channel
func PcntNewChannel(unit PcntUnitHandleT, config *PcntChanConfigT, ret_chan *PcntChannelHandleT) EspErrT

/**
 * @brief Delete the PCNT channel
 *
 * @param[in] chan PCNT channel handle created by `pcnt_new_channel()`
 * @return
 *      - ESP_OK: Delete the PCNT channel successfully
 *      - ESP_ERR_INVALID_ARG: Delete the PCNT channel failed because of invalid argument
 *      - ESP_FAIL: Delete the PCNT channel failed because of other error
 */
//go:linkname PcntDelChannel C.pcnt_del_channel
func PcntDelChannel(chan_ PcntChannelHandleT) EspErrT

/**
 * @brief Set channel actions when edge signal changes (e.g. falling or rising edge occurred).
 *        The edge signal is input from the `edge_gpio_num` configured in `pcnt_chan_config_t`.
 *        We use these actions to control when and how to change the counter value.
 *
 * @param[in] chan PCNT channel handle created by `pcnt_new_channel()`
 * @param[in] pos_act Action on posedge signal
 * @param[in] neg_act Action on negedge signal
 * @return
 *      - ESP_OK: Set edge action for PCNT channel successfully
 *      - ESP_ERR_INVALID_ARG: Set edge action for PCNT channel failed because of invalid argument
 *      - ESP_FAIL: Set edge action for PCNT channel failed because of other error
 */
//go:linkname PcntChannelSetEdgeAction C.pcnt_channel_set_edge_action
func PcntChannelSetEdgeAction(chan_ PcntChannelHandleT, pos_act PcntChannelEdgeActionT, neg_act PcntChannelEdgeActionT) EspErrT

/**
 * @brief Set channel actions when level signal changes (e.g. signal level goes from high to low).
 *        The level signal is input from the `level_gpio_num` configured in `pcnt_chan_config_t`.
 *        We use these actions to control when and how to change the counting mode.
 *
 * @param[in] chan PCNT channel handle created by `pcnt_new_channel()`
 * @param[in] high_act Action on high level signal
 * @param[in] low_act Action on low level signal
 * @return
 *      - ESP_OK: Set level action for PCNT channel successfully
 *      - ESP_ERR_INVALID_ARG: Set level action for PCNT channel failed because of invalid argument
 *      - ESP_FAIL: Set level action for PCNT channel failed because of other error
 */
//go:linkname PcntChannelSetLevelAction C.pcnt_channel_set_level_action
func PcntChannelSetLevelAction(chan_ PcntChannelHandleT, high_act PcntChannelLevelActionT, low_act PcntChannelLevelActionT) EspErrT
