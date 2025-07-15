package esp_hw_support

import _ "unsafe"

/**
 * @brief Set the clock source for IO MUX
 *
 * @note IO MUX clock is shared by submodules like SDM, Glitch Filter.
 *       The submodule drivers should call this function to detect if the user set the clock differently.
 *
 * @param clk_src The clock source for IO MUX
 * @return
 *      - ESP_OK: Success
 *      - ESP_ERR_INVALID_STATE: The IO MUX has been set to another clock source
 */
//go:linkname IoMuxSetClockSource C.io_mux_set_clock_source
func IoMuxSetClockSource(clk_src SocModuleClkT) EspErrT
