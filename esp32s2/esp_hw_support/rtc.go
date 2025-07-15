package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get current value of RTC counter in microseconds
 *
 * Note: this function may take up to 1 RTC_SLOW_CLK cycle to execute
 *
 * @return current value of RTC counter in microseconds
 */
//go:linkname EspRtcGetTimeUs C.esp_rtc_get_time_us
func EspRtcGetTimeUs() c.Uint64T
