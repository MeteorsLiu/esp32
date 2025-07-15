package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Print formatted string to console device
 * @note float and long long data are not supported!
 *
 * @param fmt Format string
 * @param ap List of arguments.
 * @return int: Total number of characters written on success; A negative number on failure.
 */
//go:linkname EspRomVprintf C.esp_rom_vprintf
func EspRomVprintf(fmt *c.Char, ap c.VaList) c.Int

/**
 * @brief esp_rom_printf can print message to different channels simultaneously.
 *        This function can help install the low level putc function for esp_rom_printf.
 *
 * @param channel Channel number (starting from 1)
 * @param putc Function pointer to the putc implementation. Set NULL can disconnect esp_rom_printf with putc.
 */
//go:linkname EspRomInstallChannelPutc C.esp_rom_install_channel_putc
func EspRomInstallChannelPutc(channel c.Int, putc func(c.Char))

/**
 * @brief It outputs a character to different channels simultaneously.
 *        This function is used by esp_rom_printf/esp_rom_vprintf.
 *
 * @param c Char to output.
 */
//go:linkname EspRomOutputToChannels C.esp_rom_output_to_channels
func EspRomOutputToChannels(c c.Char)

/**
 * @brief Set the real CPU tick rate
 *
 * @note Call this function when CPU frequency is changed, otherwise the `esp_rom_delay_us` can be inaccurate.
 *
 * @param ticks_per_us CPU ticks per us
 */
//go:linkname EspRomSetCpuTicksPerUs C.esp_rom_set_cpu_ticks_per_us
func EspRomSetCpuTicksPerUs(ticks_per_us c.Uint32T)
