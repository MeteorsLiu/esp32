package esp_hw_support

import _ "unsafe"

type EspClockOutputMapping struct {
	Unused [8]uint8
}
type EspClockOutputMappingHandleT *EspClockOutputMapping

/**
* @brief Start output specified clock signal to specified GPIO, will also
*        initialize the clkout_mapping_ret_hdl.
*
* @param[in]   clk_src  The clock signal source to be mapped to GPIOs
* @param[in]   gpio_num GPIO number to be mapped soc_root_clk signal source
* @param[out]  clkout_mapping_ret_hdl Clock output control handler
* @return
*      - ESP_OK: Output specified clock signal to specified GPIO successfully
*      - ESP_ERR_INVALID_ARG: Specified GPIO not supported to output internal clock
*                             or specified GPIO is already mapped to other internal clock source.
 *     - ESP_FAIL: There are no clock out signals that can be allocated.
*/
//go:linkname EspClockOutputStart C.esp_clock_output_start
func EspClockOutputStart(clk_sig SocClkoutSigIdT, gpio_num GpioNumT, clkout_mapping_ret_hdl *EspClockOutputMappingHandleT) EspErrT

/**
 * @brief Stop clock signal to GPIO outputting
 * @param[in]  clkout_mapping_hdl Clock output mapping control handle
 * @return
 *     - ESP_OK: Disable the clock output on GPIO successfully
 *     - ESP_ERR_INVALID_ARG  The clock mapping handle is not initialized yet
 *     - ESP_ERR_INVALID_STATE  The clock mapping handle is already in the disabled state
 */
//go:linkname EspClockOutputStop C.esp_clock_output_stop
func EspClockOutputStop(clkout_mapping_hdl EspClockOutputMappingHandleT) EspErrT
