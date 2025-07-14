package esp_driver_mcpwm

import _ "unsafe"

/**
 * @brief MCPWM event comparator ETM event configuration
 */

type McpwmCmprEtmEventConfigT struct {
	EventType McpwmComparatorEtmEventTypeT
}

/**
 * @brief Get the ETM event for MCPWM comparator
 *
 * @note The created ETM event object can be deleted later by calling `esp_etm_del_event`
 *
 * @param[in] cmpr MCPWM comparator, allocated by `mcpwm_new_comparator()` or `mcpwm_new_event_comparator()`
 * @param[in] config MCPWM ETM comparator event configuration
 * @param[out] out_event Returned ETM event handle
 * @return
 *      - ESP_OK: Get ETM event successfully
 *      - ESP_ERR_INVALID_ARG: Get ETM event failed because of invalid argument
 *      - ESP_FAIL: Get ETM event failed because of other error
 */
//go:linkname McpwmComparatorNewEtmEvent C.mcpwm_comparator_new_etm_event
func McpwmComparatorNewEtmEvent(cmpr McpwmCmprHandleT, config *McpwmCmprEtmEventConfigT, out_event *EspEtmEventHandleT) EspErrT
