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
 * @brief Get the RTC IO input level
 *
 * @param  gpio_num GPIO number (e.g. GPIO_NUM_12)
 *
 * @return
 *     - 1 High level
 *     - 0 Low level
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioGetLevel C.rtc_gpio_get_level
func RtcGpioGetLevel(gpio_num GpioNumT) c.Uint32T

/**
 * @brief Set the RTC IO output level
 *
 * @param  gpio_num GPIO number (e.g. GPIO_NUM_12)
 * @param  level output level
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioSetLevel C.rtc_gpio_set_level
func RtcGpioSetLevel(gpio_num GpioNumT, level c.Uint32T) EspErrT

/**
 * @brief    RTC GPIO set direction
 *
 * Configure RTC GPIO direction, such as output only, input only,
 * output and input.
 *
 * @param  gpio_num GPIO number (e.g. GPIO_NUM_12)
 * @param  mode GPIO direction
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioSetDirection C.rtc_gpio_set_direction
func RtcGpioSetDirection(gpio_num GpioNumT, mode RtcGpioModeT) EspErrT

/**
 * @brief RTC GPIO set direction in deep sleep mode or disable sleep status (default).
 *        In some application scenarios, IO needs to have another states during deep sleep.
 *
 * NOTE: ESP32 supports INPUT_ONLY mode.
 *       The rest targets support INPUT_ONLY, OUTPUT_ONLY, INPUT_OUTPUT mode.
 *
 * @param  gpio_num GPIO number (e.g. GPIO_NUM_12)
 * @param  mode GPIO direction
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioSetDirectionInSleep C.rtc_gpio_set_direction_in_sleep
func RtcGpioSetDirectionInSleep(gpio_num GpioNumT, mode RtcGpioModeT) EspErrT

/**
 * @brief  RTC GPIO pullup enable
 *
 * This function only works for RTC IOs. In general, call gpio_pullup_en,
 * which will work both for normal GPIOs and RTC IOs.
 *
 * @param  gpio_num GPIO number (e.g. GPIO_NUM_12)
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioPullupEn C.rtc_gpio_pullup_en
func RtcGpioPullupEn(gpio_num GpioNumT) EspErrT

/**
 * @brief  RTC GPIO pulldown enable
 *
 * This function only works for RTC IOs. In general, call gpio_pulldown_en,
 * which will work both for normal GPIOs and RTC IOs.
 *
 * @param  gpio_num GPIO number (e.g. GPIO_NUM_12)
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioPulldownEn C.rtc_gpio_pulldown_en
func RtcGpioPulldownEn(gpio_num GpioNumT) EspErrT

/**
 * @brief  RTC GPIO pullup disable
 *
 * This function only works for RTC IOs. In general, call gpio_pullup_dis,
 * which will work both for normal GPIOs and RTC IOs.
 *
 * @param  gpio_num GPIO number (e.g. GPIO_NUM_12)
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioPullupDis C.rtc_gpio_pullup_dis
func RtcGpioPullupDis(gpio_num GpioNumT) EspErrT

/**
 * @brief  RTC GPIO pulldown disable
 *
 * This function only works for RTC IOs. In general, call gpio_pulldown_dis,
 * which will work both for normal GPIOs and RTC IOs.
 *
 * @param  gpio_num GPIO number (e.g. GPIO_NUM_12)
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO is not an RTC IO
 */
//go:linkname RtcGpioPulldownDis C.rtc_gpio_pulldown_dis
func RtcGpioPulldownDis(gpio_num GpioNumT) EspErrT

/**
 * @brief Set RTC GPIO pad drive capability
 *
 * @param gpio_num GPIO number, only support output GPIOs
 * @param strength Drive capability of the pad
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname RtcGpioSetDriveCapability C.rtc_gpio_set_drive_capability
func RtcGpioSetDriveCapability(gpio_num GpioNumT, strength GpioDriveCapT) EspErrT

/**
 * @brief Get RTC GPIO pad drive capability
 *
 * @param gpio_num GPIO number, only support output GPIOs
 * @param strength Pointer to accept drive capability of the pad
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname RtcGpioGetDriveCapability C.rtc_gpio_get_drive_capability
func RtcGpioGetDriveCapability(gpio_num GpioNumT, strength *GpioDriveCapT) EspErrT

/**
 * @brief Select a RTC IOMUX function for the RTC IO
 *
 * @param gpio_num GPIO number
 * @param func Function to assign to the pin
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname RtcGpioIomuxFuncSel C.rtc_gpio_iomux_func_sel
func RtcGpioIomuxFuncSel(gpio_num GpioNumT, func_ c.Int) EspErrT

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

/**
 * @brief Helper function to disconnect internal circuits from an RTC IO
 * This function disables input, output, pullup, pulldown, and enables
 * hold feature for an RTC IO.
 * Use this function if an RTC IO needs to be disconnected from internal
 * circuits in deep sleep, to minimize leakage current.
 *
 * In particular, for ESP32-WROVER module, call
 * rtc_gpio_isolate(GPIO_NUM_12) before entering deep sleep, to reduce
 * deep sleep current.
 *
 * @param gpio_num GPIO number (e.g. GPIO_NUM_12).
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if GPIO is not an RTC IO
 */
//go:linkname RtcGpioIsolate C.rtc_gpio_isolate
func RtcGpioIsolate(gpio_num GpioNumT) EspErrT

/**
 * @brief Enable wakeup from sleep mode using specific GPIO
 * @param gpio_num  GPIO number
 * @param intr_type  Wakeup on high level (GPIO_INTR_HIGH_LEVEL) or low level
 *                   (GPIO_INTR_LOW_LEVEL)
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if gpio_num is not an RTC IO, or intr_type is not
 *        one of GPIO_INTR_HIGH_LEVEL, GPIO_INTR_LOW_LEVEL.
 */
//go:linkname RtcGpioWakeupEnable C.rtc_gpio_wakeup_enable
func RtcGpioWakeupEnable(gpio_num GpioNumT, intr_type GpioIntTypeT) EspErrT

/**
 * @brief Disable wakeup from sleep mode using specific GPIO
 * @param gpio_num  GPIO number
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if gpio_num is not an RTC IO
 */
//go:linkname RtcGpioWakeupDisable C.rtc_gpio_wakeup_disable
func RtcGpioWakeupDisable(gpio_num GpioNumT) EspErrT
