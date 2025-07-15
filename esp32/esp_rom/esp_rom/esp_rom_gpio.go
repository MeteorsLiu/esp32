package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Combine a peripheral signal which tagged as output attribute with a GPIO.
 *
 * @note There's no limitation on the number of signals that a GPIO can combine with.
 * @note Internally, the signal will be connected first, then output will be enabled on the pad.
 *
 * @param gpio_num GPIO number
 * @param signal_idx Peripheral signal index (tagged as output attribute). Particularly, `SIG_GPIO_OUT_IDX` means disconnect GPIO and other peripherals. Only the GPIO driver can control the output level.
 * @param out_inv Whether to signal to be inverted or not
 * @param oen_inv Whether the output enable control is inverted or not
 */
//go:linkname EspRomGpioConnectOutSignal C.esp_rom_gpio_connect_out_signal
func EspRomGpioConnectOutSignal(gpio_num c.Uint32T, signal_idx c.Uint32T, out_inv bool, oen_inv bool)
