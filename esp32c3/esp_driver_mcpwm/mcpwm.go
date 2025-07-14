package esp_driver_mcpwm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get MCPWM timer phase
 *
 * @param[in] timer MCPWM timer handle, allocated by `mcpwm_new_timer()`
 * @param[out] count_value Returned MCPWM timer phase
 * @param[out] direction Returned MCPWM timer counting direction
 * @return
 *      - ESP_OK: Get MCPWM timer status successfully
 *      - ESP_ERR_INVALID_ARG: Get MCPWM timer status failed because of invalid argument
 *      - ESP_FAIL: Get MCPWM timer status failed because of other error
 */
//go:linkname McpwmTimerGetPhase C.mcpwm_timer_get_phase
func McpwmTimerGetPhase(timer McpwmTimerHandleT, count_value *c.Uint32T, direction *McpwmTimerDirectionT) EspErrT
