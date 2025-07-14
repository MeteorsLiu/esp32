package esp_driver_gpio

import _ "unsafe"

/**
 * @brief Determine if the specified GPIO is a valid RTC GPIO.
 *
 * @param gpio_num GPIO number
 * @return true if GPIO is valid for RTC GPIO use. false otherwise.
 */
//go:linkname RtcGpioIsValidGpio C.rtc_gpio_is_valid_gpio
func RtcGpioIsValidGpio(gpio_num GpioNumT) bool
