package esp_driver_mcpwm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief MCPWM capture timer configuration structure
 */

type McpwmCaptureTimerConfigT struct {
	GroupId      c.Int
	ClkSrc       McpwmCaptureClockSourceT
	ResolutionHz c.Uint32T
	Flags        struct {
		AllowPd c.Uint32T
	}
}

/**
 * @brief Create MCPWM capture timer
 *
 * @param[in] config MCPWM capture timer configuration
 * @param[out] ret_cap_timer Returned MCPWM capture timer handle
 * @return
 *      - ESP_OK: Create MCPWM capture timer successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM capture timer failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM capture timer failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create MCPWM capture timer failed because can't find free resource
 *      - ESP_FAIL: Create MCPWM capture timer failed because of other error
 */
// llgo:link (*McpwmCaptureTimerConfigT).McpwmNewCaptureTimer C.mcpwm_new_capture_timer
func (recv_ *McpwmCaptureTimerConfigT) McpwmNewCaptureTimer(ret_cap_timer *McpwmCapTimerHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete MCPWM capture timer
 *
 * @param[in] cap_timer MCPWM capture timer, allocated by `mcpwm_new_capture_timer()`
 * @return
 *      - ESP_OK: Delete MCPWM capture timer successfully
 *      - ESP_ERR_INVALID_ARG: Delete MCPWM capture timer failed because of invalid argument
 *      - ESP_FAIL: Delete MCPWM capture timer failed because of other error
 */
//go:linkname McpwmDelCaptureTimer C.mcpwm_del_capture_timer
func McpwmDelCaptureTimer(cap_timer McpwmCapTimerHandleT) EspErrT

/**
 * @brief Enable MCPWM capture timer
 *
 * @param[in] cap_timer MCPWM capture timer handle, allocated by `mcpwm_new_capture_timer()`
 * @return
 *      - ESP_OK: Enable MCPWM capture timer successfully
 *      - ESP_ERR_INVALID_ARG: Enable MCPWM capture timer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Enable MCPWM capture timer failed because timer is enabled already
 *      - ESP_FAIL: Enable MCPWM capture timer failed because of other error
 */
//go:linkname McpwmCaptureTimerEnable C.mcpwm_capture_timer_enable
func McpwmCaptureTimerEnable(cap_timer McpwmCapTimerHandleT) EspErrT

/**
 * @brief Disable MCPWM capture timer
 *
 * @param[in] cap_timer MCPWM capture timer handle, allocated by `mcpwm_new_capture_timer()`
 * @return
 *      - ESP_OK: Disable MCPWM capture timer successfully
 *      - ESP_ERR_INVALID_ARG: Disable MCPWM capture timer failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Disable MCPWM capture timer failed because timer is disabled already
 *      - ESP_FAIL: Disable MCPWM capture timer failed because of other error
 */
//go:linkname McpwmCaptureTimerDisable C.mcpwm_capture_timer_disable
func McpwmCaptureTimerDisable(cap_timer McpwmCapTimerHandleT) EspErrT

/**
 * @brief Start MCPWM capture timer
 *
 * @param[in] cap_timer MCPWM capture timer, allocated by `mcpwm_new_capture_timer()`
 * @return
 *      - ESP_OK: Start MCPWM capture timer successfully
 *      - ESP_ERR_INVALID_ARG: Start MCPWM capture timer failed because of invalid argument
 *      - ESP_FAIL: Start MCPWM capture timer failed because of other error
 */
//go:linkname McpwmCaptureTimerStart C.mcpwm_capture_timer_start
func McpwmCaptureTimerStart(cap_timer McpwmCapTimerHandleT) EspErrT

/**
 * @brief Stop MCPWM capture timer
 *
 * @param[in] cap_timer MCPWM capture timer, allocated by `mcpwm_new_capture_timer()`
 * @return
 *      - ESP_OK: Stop MCPWM capture timer successfully
 *      - ESP_ERR_INVALID_ARG: Stop MCPWM capture timer failed because of invalid argument
 *      - ESP_FAIL: Stop MCPWM capture timer failed because of other error
 */
//go:linkname McpwmCaptureTimerStop C.mcpwm_capture_timer_stop
func McpwmCaptureTimerStop(cap_timer McpwmCapTimerHandleT) EspErrT

/**
 * @brief Get MCPWM capture timer resolution, in Hz
 *
 * @param[in] cap_timer MCPWM capture timer, allocated by `mcpwm_new_capture_timer()`
 * @param[out] out_resolution Returned capture timer resolution, in Hz
 * @return
 *      - ESP_OK: Get capture timer resolution successfully
 *      - ESP_ERR_INVALID_ARG: Get capture timer resolution failed because of invalid argument
 *      - ESP_FAIL: Get capture timer resolution failed because of other error
 */
//go:linkname McpwmCaptureTimerGetResolution C.mcpwm_capture_timer_get_resolution
func McpwmCaptureTimerGetResolution(cap_timer McpwmCapTimerHandleT, out_resolution *c.Uint32T) EspErrT

/**
 * @brief MCPWM Capture timer sync phase configuration
 */

type McpwmCaptureTimerSyncPhaseConfigT struct {
	SyncSrc    McpwmSyncHandleT
	CountValue c.Uint32T
	Direction  McpwmTimerDirectionT
}

/**
 * @brief Set sync phase for MCPWM capture timer
 *
 * @param[in] cap_timer MCPWM capture timer, allocated by `mcpwm_new_capture_timer()`
 * @param[in] config MCPWM capture timer sync phase configuration
 * @return
 *      - ESP_OK: Set sync phase for MCPWM capture timer successfully
 *      - ESP_ERR_INVALID_ARG: Set sync phase for MCPWM capture timer failed because of invalid argument
 *      - ESP_FAIL: Set sync phase for MCPWM capture timer failed because of other error
 */
//go:linkname McpwmCaptureTimerSetPhaseOnSync C.mcpwm_capture_timer_set_phase_on_sync
func McpwmCaptureTimerSetPhaseOnSync(cap_timer McpwmCapTimerHandleT, config *McpwmCaptureTimerSyncPhaseConfigT) EspErrT

/**
 * @brief MCPWM capture channel configuration structure
 */

type McpwmCaptureChannelConfigT struct {
	GpioNum      c.Int
	IntrPriority c.Int
	Prescale     c.Uint32T
	Flags        ExtraCaptureChannelFlags
}

type ExtraCaptureChannelFlags struct {
	Unused [8]uint8
}

/**
 * @brief Create MCPWM capture channel
 *
 * @note The created capture channel won't be enabled until calling `mcpwm_capture_channel_enable`
 *
 * @param[in] cap_timer MCPWM capture timer, allocated by `mcpwm_new_capture_timer()`, will be connected to the new capture channel
 * @param[in] config MCPWM capture channel configuration
 * @param[out] ret_cap_channel Returned MCPWM capture channel
 * @return
 *      - ESP_OK: Create MCPWM capture channel successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM capture channel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM capture channel failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create MCPWM capture channel failed because can't find free resource
 *      - ESP_FAIL: Create MCPWM capture channel failed because of other error
 */
//go:linkname McpwmNewCaptureChannel C.mcpwm_new_capture_channel
func McpwmNewCaptureChannel(cap_timer McpwmCapTimerHandleT, config *McpwmCaptureChannelConfigT, ret_cap_channel *McpwmCapChannelHandleT) EspErrT

/**
 * @brief Delete MCPWM capture channel
 *
 * @param[in] cap_channel MCPWM capture channel handle, allocated by `mcpwm_new_capture_channel()`
 * @return
 *      - ESP_OK: Delete MCPWM capture channel successfully
 *      - ESP_ERR_INVALID_ARG: Delete MCPWM capture channel failed because of invalid argument
 *      - ESP_FAIL: Delete MCPWM capture channel failed because of other error
 */
//go:linkname McpwmDelCaptureChannel C.mcpwm_del_capture_channel
func McpwmDelCaptureChannel(cap_channel McpwmCapChannelHandleT) EspErrT

/**
 * @brief Enable MCPWM capture channel
 *
 * @note This function will transit the channel state from init to enable.
 * @note This function will enable the interrupt service, if it's lazy installed in `mcpwm_capture_channel_register_event_callbacks()`.
 *
 * @param[in] cap_channel MCPWM capture channel handle, allocated by `mcpwm_new_capture_channel()`
 * @return
 *      - ESP_OK: Enable MCPWM capture channel successfully
 *      - ESP_ERR_INVALID_ARG: Enable MCPWM capture channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Enable MCPWM capture channel failed because the channel is already enabled
 *      - ESP_FAIL: Enable MCPWM capture channel failed because of other error
 */
//go:linkname McpwmCaptureChannelEnable C.mcpwm_capture_channel_enable
func McpwmCaptureChannelEnable(cap_channel McpwmCapChannelHandleT) EspErrT

/**
 * @brief Disable MCPWM capture channel
 *
 * @param[in] cap_channel MCPWM capture channel handle, allocated by `mcpwm_new_capture_channel()`
 * @return
 *      - ESP_OK: Disable MCPWM capture channel successfully
 *      - ESP_ERR_INVALID_ARG: Disable MCPWM capture channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Disable MCPWM capture channel failed because the channel is not enabled yet
 *      - ESP_FAIL: Disable MCPWM capture channel failed because of other error
 */
//go:linkname McpwmCaptureChannelDisable C.mcpwm_capture_channel_disable
func McpwmCaptureChannelDisable(cap_channel McpwmCapChannelHandleT) EspErrT

/**
 * @brief Group of supported MCPWM capture event callbacks
 * @note The callbacks are all running under ISR environment
 */

type McpwmCaptureEventCallbacksT struct {
	OnCap McpwmCaptureEventCbT
}

/**
 * @brief Set event callbacks for MCPWM capture channel
 *
 * @note The first call to this function needs to be before the call to `mcpwm_capture_channel_enable`
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 *
 * @param[in] cap_channel MCPWM capture channel handle, allocated by `mcpwm_new_capture_channel()`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Set event callbacks failed because the channel is not in init state
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname McpwmCaptureChannelRegisterEventCallbacks C.mcpwm_capture_channel_register_event_callbacks
func McpwmCaptureChannelRegisterEventCallbacks(cap_channel McpwmCapChannelHandleT, cbs *McpwmCaptureEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Trigger a catch by software
 *
 * @param[in] cap_channel MCPWM capture channel handle, allocated by `mcpwm_new_capture_channel()`
 * @return
 *      - ESP_OK: Trigger software catch successfully
 *      - ESP_ERR_INVALID_ARG: Trigger software catch failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Trigger software catch failed because the channel is not enabled yet
 *      - ESP_FAIL: Trigger software catch failed because of other error
 */
//go:linkname McpwmCaptureChannelTriggerSoftCatch C.mcpwm_capture_channel_trigger_soft_catch
func McpwmCaptureChannelTriggerSoftCatch(cap_channel McpwmCapChannelHandleT) EspErrT
