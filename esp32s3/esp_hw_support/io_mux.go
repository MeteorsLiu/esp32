package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

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

type RtcIoStatusT struct {
	RtcIoEnabledCnt [22]c.Uint8T
	RtcIoUsingMask  c.Uint32T
}

/**
 * @brief Enable/Disable LP_IO peripheral clock
 *
 * @param gpio_num GPIO number
 * @param enable   true to enable the clock / false to disable the clock
 */
//go:linkname IoMuxEnableLpIoClock C.io_mux_enable_lp_io_clock
func IoMuxEnableLpIoClock(gpio_num GpioNumT, enable bool)

/**
 * Force disable one LP_IO to clock dependency
 * @param gpio_num GPIO number
 */
//go:linkname IoMuxForceDisableLpIoClock C.io_mux_force_disable_lp_io_clock
func IoMuxForceDisableLpIoClock(gpio_num GpioNumT)
