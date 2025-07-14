package esp_driver_mcpwm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief MCPWM comparator configuration
 */

type McpwmComparatorConfigT struct {
	IntrPriority c.Int
	Flags        struct {
		UpdateCmpOnTez  c.Uint32T
		UpdateCmpOnTep  c.Uint32T
		UpdateCmpOnSync c.Uint32T
	}
}

/**
 * @brief Create MCPWM comparator
 *
 * @param[in] oper MCPWM operator, allocated by `mcpwm_new_operator()`, the new comparator will be allocated from this operator
 * @param[in] config MCPWM comparator configuration
 * @param[out] ret_cmpr Returned MCPWM comparator
 * @return
 *      - ESP_OK: Create MCPWM comparator successfully
 *      - ESP_ERR_INVALID_ARG: Create MCPWM comparator failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create MCPWM comparator failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create MCPWM comparator failed because can't find free resource
 *      - ESP_FAIL: Create MCPWM comparator failed because of other error
 */
//go:linkname McpwmNewComparator C.mcpwm_new_comparator
func McpwmNewComparator(oper McpwmOperHandleT, config *McpwmComparatorConfigT, ret_cmpr *McpwmCmprHandleT) EspErrT

/**
 * @brief Delete MCPWM comparator
 *
 * @param[in] cmpr MCPWM comparator handle, allocated by `mcpwm_new_comparator()`
 * @return
 *      - ESP_OK: Delete MCPWM comparator successfully
 *      - ESP_ERR_INVALID_ARG: Delete MCPWM comparator failed because of invalid argument
 *      - ESP_FAIL: Delete MCPWM comparator failed because of other error
 */
//go:linkname McpwmDelComparator C.mcpwm_del_comparator
func McpwmDelComparator(cmpr McpwmCmprHandleT) EspErrT

/**
 * @brief Group of supported MCPWM compare event callbacks
 * @note The callbacks are all running under ISR environment
 */

type McpwmComparatorEventCallbacksT struct {
	OnReach McpwmCompareEventCbT
}

/**
 * @brief Set event callbacks for MCPWM comparator
 *
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 *
 * @param[in] cmpr MCPWM comparator handle, allocated by `mcpwm_new_comparator()`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname McpwmComparatorRegisterEventCallbacks C.mcpwm_comparator_register_event_callbacks
func McpwmComparatorRegisterEventCallbacks(cmpr McpwmCmprHandleT, cbs *McpwmComparatorEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Set MCPWM comparator's compare value
 *
 * @param[in] cmpr MCPWM comparator handle, allocated by `mcpwm_new_comparator()`
 * @param[in] cmp_ticks The new compare value
 * @return
 *      - ESP_OK: Set MCPWM compare value successfully
 *      - ESP_ERR_INVALID_ARG: Set MCPWM compare value failed because of invalid argument (e.g. the cmp_ticks is out of range)
 *      - ESP_ERR_INVALID_STATE: Set MCPWM compare value failed because the operator doesn't have a timer connected
 *      - ESP_FAIL: Set MCPWM compare value failed because of other error
 */
//go:linkname McpwmComparatorSetCompareValue C.mcpwm_comparator_set_compare_value
func McpwmComparatorSetCompareValue(cmpr McpwmCmprHandleT, cmp_ticks c.Uint32T) EspErrT
