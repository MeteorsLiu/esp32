package esp_driver_gpio

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Determine if the specified GPIO is a valid RTC GPIO.
 *
 * @param gpio_num GPIO number
 * @return true if GPIO is valid for RTC GPIO use. false otherwise.
 */
//go:linkname RtcGpioIsValidGpio C.rtc_gpio_is_valid_gpio
func RtcGpioIsValidGpio(gpio_num GpioNumT) bool

/**
 * @brief Get RTC IO index number by gpio number.
 *
 * @param gpio_num GPIO number
 * @return
 *        >=0: Index of rtcio.
 *        -1 : The gpio is not rtcio.
 */
//go:linkname RtcIoNumberGet C.rtc_io_number_get
func RtcIoNumberGet(gpio_num GpioNumT) c.Int

/**
 * @brief Init a GPIO as RTC GPIO
 *
 * This function must be called when initializing a pad for an analog function.
 *
 * @param  gpio_num GPIO number (e.g. GPIO_NUM_12)
 *
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioInit C.rtc_gpio_init
func RtcGpioInit(gpio_num GpioNumT) EspErrT

/**
 * @brief Init a GPIO as digital GPIO
 *
 * @param  gpio_num GPIO number (e.g. GPIO_NUM_12)
 *
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioDeinit C.rtc_gpio_deinit
func RtcGpioDeinit(gpio_num GpioNumT) EspErrT

/**
 * @brief Enable hold function on an RTC IO pad
 *
 * Enabling HOLD function will cause the pad to latch current values of
 * input enable, output enable, output value, function, drive strength values.
 * This function is useful when going into light or deep sleep mode to prevent
 * the pin configuration from changing.
 *
 * @param gpio_num GPIO number (e.g. GPIO_NUM_12)
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioHoldEn C.rtc_gpio_hold_en
func RtcGpioHoldEn(gpio_num GpioNumT) EspErrT

/**
 * @brief Disable hold function on an RTC IO pad
 *
 * Disabling hold function will allow the pad receive the values of
 * input enable, output enable, output value, function, drive strength from
 * RTC_IO peripheral.
 *
 * @param gpio_num GPIO number (e.g. GPIO_NUM_12)
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioHoldDis C.rtc_gpio_hold_dis
func RtcGpioHoldDis(gpio_num GpioNumT) EspErrT

/**
 * @brief Enable force hold signal for all RTC IOs
 *
 * Each RTC pad has a "force hold" input signal from the RTC controller.
 * If this signal is set, pad latches current values of input enable,
 * function, output enable, and other signals which come from the RTC mux.
 * Force hold signal is enabled before going into deep sleep for pins which
 * are used for EXT1 wakeup.
 */
//go:linkname RtcGpioForceHoldEnAll C.rtc_gpio_force_hold_en_all
func RtcGpioForceHoldEnAll() EspErrT

/**
 * @brief Disable force hold signal for all RTC IOs
 */
//go:linkname RtcGpioForceHoldDisAll C.rtc_gpio_force_hold_dis_all
func RtcGpioForceHoldDisAll() EspErrT
