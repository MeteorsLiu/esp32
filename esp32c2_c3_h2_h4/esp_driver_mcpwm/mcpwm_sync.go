package esp_driver_mcpwm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief MCPWM timer sync source configuration
 */

type McpwmTimerSyncSrcConfigT struct {
	TimerEvent McpwmTimerEventT
	Flags      struct {
		PropagateInputSync c.Uint32T
	}
}

/**
 * @brief Create MCPWM timer sync source
 *
 * @param[in] timer MCPWM timer handle, allocated by `mcpwm_new_timer()`
 * @param[in] config MCPWM timer sync source configuration
 * @param[out] ret_sync Returned MCPWM sync handle
 * @return
 *      - ESP_OK: Create MCPWM timer sync source successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM timer sync source failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM timer sync source failed because out of memory
 *      - ESP_ERR_INVALID_STATE: Create MCPWM timer sync source failed because the timer has created a sync source before
 *      - ESP_FAIL: Create MCPWM timer sync source failed because of other error
 */
//go:linkname McpwmNewTimerSyncSrc C.mcpwm_new_timer_sync_src
func McpwmNewTimerSyncSrc(timer McpwmTimerHandleT, config *McpwmTimerSyncSrcConfigT, ret_sync *McpwmSyncHandleT) EspErrT

/**
 * @brief MCPWM GPIO sync source configuration
 */

type McpwmGpioSyncSrcConfigT struct {
	GroupId c.Int
	GpioNum c.Int
	Flags   struct {
		ActiveNeg  c.Uint32T
		IoLoopBack c.Uint32T
		PullUp     c.Uint32T
		PullDown   c.Uint32T
	}
}

/**
 * @brief Create MCPWM GPIO sync source
 *
 * @param[in] config MCPWM GPIO sync source configuration
 * @param[out] ret_sync Returned MCPWM GPIO sync handle
 * @return
 *      - ESP_OK: Create MCPWM GPIO sync source successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM GPIO sync source failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM GPIO sync source failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create MCPWM GPIO sync source failed because can't find free resource
 *      - ESP_FAIL: Create MCPWM GPIO sync source failed because of other error
 */
// llgo:link (*McpwmGpioSyncSrcConfigT).McpwmNewGpioSyncSrc C.mcpwm_new_gpio_sync_src
func (recv_ *McpwmGpioSyncSrcConfigT) McpwmNewGpioSyncSrc(ret_sync *McpwmSyncHandleT) EspErrT {
	return 0
}

/**
 * @brief MCPWM software sync configuration structure
 */

type McpwmSoftSyncConfigT struct {
	Unused [8]uint8
}

/**
 * @brief Create MCPWM software sync source
 *
 * @param[in] config MCPWM software sync source configuration
 * @param[out] ret_sync Returned software sync handle
 * @return
 *      - ESP_OK: Create MCPWM software sync successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM software sync failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM software sync failed because out of memory
 *      - ESP_FAIL: Create MCPWM software sync failed because of other error
 */
// llgo:link (*McpwmSoftSyncConfigT).McpwmNewSoftSyncSrc C.mcpwm_new_soft_sync_src
func (recv_ *McpwmSoftSyncConfigT) McpwmNewSoftSyncSrc(ret_sync *McpwmSyncHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete MCPWM sync source
 *
 * @param[in] sync MCPWM sync handle, allocated by `mcpwm_new_timer_sync_src()` or `mcpwm_new_gpio_sync_src()` or `mcpwm_new_soft_sync_src()`
 * @return
 *      - ESP_OK: Delete MCPWM sync source successfully
 *      - ESP_ERR_INVALID_ARG: Delete MCPWM sync source failed because of invalid argument
 *      - ESP_FAIL: Delete MCPWM sync source failed because of other error
 */
//go:linkname McpwmDelSyncSrc C.mcpwm_del_sync_src
func McpwmDelSyncSrc(sync McpwmSyncHandleT) EspErrT

/**
 * @brief Activate the software sync, trigger the sync event for once
 *
 * @param[in] sync MCPWM soft sync handle, allocated by `mcpwm_new_soft_sync_src()`
 * @return
 *      - ESP_OK: Trigger MCPWM software sync event successfully
 *      - ESP_ERR_INVALID_ARG: Trigger MCPWM software sync event failed because of invalid argument
 *      - ESP_FAIL: Trigger MCPWM software sync event failed because of other error
 */
//go:linkname McpwmSoftSyncActivate C.mcpwm_soft_sync_activate
func McpwmSoftSyncActivate(sync McpwmSyncHandleT) EspErrT
