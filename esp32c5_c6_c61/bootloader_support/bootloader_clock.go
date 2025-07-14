package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** @brief Configure clocks for early boot
 *
 * Called by bootloader, or by the app if the bootloader version is old (pre v2.1).
 */
//go:linkname BootloaderClockConfigure C.bootloader_clock_configure
func BootloaderClockConfigure()

/** @brief Return the rated maximum frequency of this chip
 */
//go:linkname BootloaderClockGetRatedFreqMhz C.bootloader_clock_get_rated_freq_mhz
func BootloaderClockGetRatedFreqMhz() c.Int
