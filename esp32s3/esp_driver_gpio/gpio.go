package esp_driver_gpio

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GpioIsrHandleT IntrHandleT

// llgo:type C
type GpioIsrT func(c.Pointer)

/**
 * @brief Configuration parameters of GPIO pad for gpio_config function
 */

type GpioConfigT struct {
	PinBitMask c.Uint64T
	Mode       GpioModeT
	PullUpEn   GpioPullupT
	PullDownEn GpioPulldownT
	IntrType   GpioIntTypeT
}

/**
 * @brief GPIO common configuration
 *
 *        Configure GPIO's Mode,pull-up,PullDown,IntrType
 *
 * @param  pGPIOConfig Pointer to GPIO configure struct
 *
 * @note This function always overwrite all the current IO configurations
 *
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *
 */
// llgo:link (*GpioConfigT).GpioConfig C.gpio_config
func (recv_ *GpioConfigT) GpioConfig() EspErrT {
	return 0
}

/**
 * @brief Reset an gpio to default state (select gpio function, enable pullup and disable input and output).
 *
 * @param gpio_num GPIO number.
 *
 * @note This function also configures the IOMUX for this pin to the GPIO
 *       function, and disconnects any other peripheral output configured via GPIO
 *       Matrix.
 *
 * @return Always return ESP_OK.
 */
//go:linkname GpioResetPin C.gpio_reset_pin
func GpioResetPin(gpio_num GpioNumT) EspErrT

/**
 * @brief  GPIO set interrupt trigger type
 *
 * @param  gpio_num GPIO number. If you want to set the trigger type of e.g. of GPIO16, gpio_num should be GPIO_NUM_16 (16);
 * @param  intr_type Interrupt type, select from gpio_int_type_t
 *
 * @return
 *     - ESP_OK  Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *
 */
//go:linkname GpioSetIntrType C.gpio_set_intr_type
func GpioSetIntrType(gpio_num GpioNumT, intr_type GpioIntTypeT) EspErrT

/**
 * @brief  Enable GPIO module interrupt signal
 *
 * @note ESP32: Please do not use the interrupt of GPIO36 and GPIO39 when using ADC or Wi-Fi and Bluetooth with sleep mode enabled.
 *       Please refer to the comments of `adc1_get_raw`.
 *       Please refer to Section 3.11 of <a href="https://espressif.com/sites/default/files/documentation/eco_and_workarounds_for_bugs_in_esp32_en.pdf">ESP32 ECO and Workarounds for Bugs</a> for the description of this issue.

 *
 * @param  gpio_num GPIO number. If you want to enable an interrupt on e.g. GPIO16, gpio_num should be GPIO_NUM_16 (16);
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *
 */
//go:linkname GpioIntrEnable C.gpio_intr_enable
func GpioIntrEnable(gpio_num GpioNumT) EspErrT

/**
 * @brief  Disable GPIO module interrupt signal
 *
 * @note This function is allowed to be executed when Cache is disabled within ISR context, by enabling `CONFIG_GPIO_CTRL_FUNC_IN_IRAM`
 *
 * @param  gpio_num GPIO number. If you want to disable the interrupt of e.g. GPIO16, gpio_num should be GPIO_NUM_16 (16);
 *
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *
 */
//go:linkname GpioIntrDisable C.gpio_intr_disable
func GpioIntrDisable(gpio_num GpioNumT) EspErrT

/**
 * @brief  GPIO set output level
 *
 * @note This function is allowed to be executed when Cache is disabled within ISR context, by enabling `CONFIG_GPIO_CTRL_FUNC_IN_IRAM`
 *
 * @param  gpio_num GPIO number. If you want to set the output level of e.g. GPIO16, gpio_num should be GPIO_NUM_16 (16);
 * @param  level Output level. 0: low ; 1: high
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO number error
 *
 */
//go:linkname GpioSetLevel C.gpio_set_level
func GpioSetLevel(gpio_num GpioNumT, level c.Uint32T) EspErrT

/**
 * @brief  GPIO get input level
 *
 * @warning If the pad is not configured for input (or input and output) the returned value is always 0.
 *
 * @param  gpio_num GPIO number. If you want to get the logic level of e.g. pin GPIO16, gpio_num should be GPIO_NUM_16 (16);
 *
 * @return
 *     - 0 the GPIO input level is 0
 *     - 1 the GPIO input level is 1
 *
 */
//go:linkname GpioGetLevel C.gpio_get_level
func GpioGetLevel(gpio_num GpioNumT) c.Int

/**
 * @brief    GPIO set direction
 *
 * Configure GPIO mode,such as output_only,input_only,output_and_input
 *
 * @param  gpio_num  Configure GPIO pins number, it should be GPIO number. If you want to set direction of e.g. GPIO16, gpio_num should be GPIO_NUM_16 (16);
 * @param  mode GPIO direction
 *
 * @note This function always overwrite all the current modes that have applied on the IO pin
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO error
 *
 */
//go:linkname GpioSetDirection C.gpio_set_direction
func GpioSetDirection(gpio_num GpioNumT, mode GpioModeT) EspErrT

/**
 * @brief Enable input for an IO
 *
 * @param gpio_num GPIO number
 *
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_ARG GPIO number error
 */
//go:linkname GpioInputEnable C.gpio_input_enable
func GpioInputEnable(gpio_num GpioNumT) EspErrT

/**
 * @brief  Configure GPIO internal pull-up/pull-down resistors
 *
 * @note This function always overwrite the current pull-up/pull-down configurations
 * @note ESP32: Only pins that support both input & output have integrated pull-up and pull-down resistors. Input-only GPIOs 34-39 do not.
 *
 * @param  gpio_num GPIO number. If you want to set pull up or down mode for e.g. GPIO16, gpio_num should be GPIO_NUM_16 (16);
 * @param  pull GPIO pull up/down mode.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG : Parameter error
 *
 */
//go:linkname GpioSetPullMode C.gpio_set_pull_mode
func GpioSetPullMode(gpio_num GpioNumT, pull GpioPullModeT) EspErrT

/**
 * @brief Enable GPIO wake-up function.
 *
 * @param gpio_num GPIO number.
 *
 * @param intr_type GPIO wake-up type. Only GPIO_INTR_LOW_LEVEL or GPIO_INTR_HIGH_LEVEL can be used.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioWakeupEnable C.gpio_wakeup_enable
func GpioWakeupEnable(gpio_num GpioNumT, intr_type GpioIntTypeT) EspErrT

/**
 * @brief Disable GPIO wake-up function.
 *
 * @param gpio_num GPIO number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioWakeupDisable C.gpio_wakeup_disable
func GpioWakeupDisable(gpio_num GpioNumT) EspErrT

/**
 * @brief   Register GPIO interrupt handler, the handler is an ISR.
 *          The handler will be attached to the same CPU core that this function is running on.
 *
 * This ISR function is called whenever any GPIO interrupt occurs. See
 * the alternative gpio_install_isr_service() and
 * gpio_isr_handler_add() API in order to have the driver support
 * per-GPIO ISRs.
 *
 * @param  fn  Interrupt handler function.
 * @param  arg  Parameter for handler function
 * @param  intr_alloc_flags Flags used to allocate the interrupt. One or multiple (ORred)
 *            ESP_INTR_FLAG_* values. See esp_intr_alloc.h for more info.
 * @param  handle Pointer to return handle. If non-NULL, a handle for the interrupt will be returned here.
 *
 * \verbatim embed:rst:leading-asterisk
 * To disable or remove the ISR, pass the returned handle to the :doc:`interrupt allocation functions </api-reference/system/intr_alloc>`.
 * \endverbatim
 *
 * @return
 *     - ESP_OK Success ;
 *     - ESP_ERR_INVALID_ARG GPIO error
 *     - ESP_ERR_NOT_FOUND No free interrupt found with the specified flags
 */
//go:linkname GpioIsrRegister C.gpio_isr_register
func GpioIsrRegister(fn func(c.Pointer), arg c.Pointer, intr_alloc_flags c.Int, handle *GpioIsrHandleT) EspErrT

/**
 * @brief Enable pull-up on GPIO.
 *
 * @param gpio_num GPIO number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioPullupEn C.gpio_pullup_en
func GpioPullupEn(gpio_num GpioNumT) EspErrT

/**
 * @brief Disable pull-up on GPIO.
 *
 * @param gpio_num GPIO number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioPullupDis C.gpio_pullup_dis
func GpioPullupDis(gpio_num GpioNumT) EspErrT

/**
 * @brief Enable pull-down on GPIO.
 *
 * @param gpio_num GPIO number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioPulldownEn C.gpio_pulldown_en
func GpioPulldownEn(gpio_num GpioNumT) EspErrT

/**
 * @brief Disable pull-down on GPIO.
 *
 * @param gpio_num GPIO number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioPulldownDis C.gpio_pulldown_dis
func GpioPulldownDis(gpio_num GpioNumT) EspErrT

/**
 * @brief Install the GPIO driver's ETS_GPIO_INTR_SOURCE ISR handler service, which allows per-pin GPIO interrupt handlers.
 *
 * This function is incompatible with gpio_isr_register() - if that function is used, a single global ISR is registered for all GPIO interrupts. If this function is used, the ISR service provides a global GPIO ISR and individual pin handlers are registered via the gpio_isr_handler_add() function.
 *
 * @param intr_alloc_flags Flags used to allocate the interrupt. One or multiple (ORred)
 *            ESP_INTR_FLAG_* values. See esp_intr_alloc.h for more info.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_NO_MEM No memory to install this service
 *     - ESP_ERR_INVALID_STATE ISR service already installed.
 *     - ESP_ERR_NOT_FOUND No free interrupt found with the specified flags
 *     - ESP_ERR_INVALID_ARG GPIO error
 */
//go:linkname GpioInstallIsrService C.gpio_install_isr_service
func GpioInstallIsrService(intr_alloc_flags c.Int) EspErrT

/**
 * @brief Uninstall the driver's GPIO ISR service, freeing related resources.
 */
//go:linkname GpioUninstallIsrService C.gpio_uninstall_isr_service
func GpioUninstallIsrService()

/**
 * @brief Add ISR handler for the corresponding GPIO pin.
 *
 * Call this function after using gpio_install_isr_service() to
 * install the driver's GPIO ISR handler service.
 *
 * The pin ISR handlers no longer need to be declared with IRAM_ATTR,
 * unless you pass the ESP_INTR_FLAG_IRAM flag when allocating the
 * ISR in gpio_install_isr_service().
 *
 * This ISR handler will be called from an ISR. So there is a stack
 * size limit (configurable as "ISR stack size" in menuconfig). This
 * limit is smaller compared to a global GPIO interrupt handler due
 * to the additional level of indirection.
 *
 * @param gpio_num GPIO number
 * @param isr_handler ISR handler function for the corresponding GPIO number.
 * @param args parameter for ISR handler.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE Wrong state, the ISR service has not been initialized.
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioIsrHandlerAdd C.gpio_isr_handler_add
func GpioIsrHandlerAdd(gpio_num GpioNumT, isr_handler GpioIsrT, args c.Pointer) EspErrT

/**
 * @brief Remove ISR handler for the corresponding GPIO pin.
 *
 * @param gpio_num GPIO number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE Wrong state, the ISR service has not been initialized.
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioIsrHandlerRemove C.gpio_isr_handler_remove
func GpioIsrHandlerRemove(gpio_num GpioNumT) EspErrT

/**
 * @brief Set GPIO pad drive capability
 *
 * @param gpio_num GPIO number, only support output GPIOs
 * @param strength Drive capability of the pad
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioSetDriveCapability C.gpio_set_drive_capability
func GpioSetDriveCapability(gpio_num GpioNumT, strength GpioDriveCapT) EspErrT

/**
 * @brief Get GPIO pad drive capability
 *
 * @param gpio_num GPIO number, only support output GPIOs
 * @param strength Pointer to accept drive capability of the pad
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioGetDriveCapability C.gpio_get_drive_capability
func GpioGetDriveCapability(gpio_num GpioNumT, strength *GpioDriveCapT) EspErrT

/**
 * @brief Enable gpio pad hold function.
 *
 * When a GPIO is set to hold, its state is latched at that moment and will not change when the internal
 * signal or the IO MUX/GPIO configuration is modified (including input enable, output enable, output value,
 * function, and drive strength values). This function can be used to retain the state of GPIOs when the power
 * domain of where GPIO/IOMUX belongs to becomes off. For example, chip or system is reset (e.g. watchdog
 * time-out, deep-sleep events are triggered), or peripheral power-down in light-sleep.
 *
 * This function works in both input and output modes, and only applicable to output-capable GPIOs.
 * If this function is enabled:
 *   in output mode: the output level of the GPIO will be locked and can not be changed.
 *   in input mode: the input read value can still reflect the changes of the input signal.
 *
 * Please be aware that,
 *
 * On ESP32P4, the states of IOs can not be hold after waking up from Deep-sleep.
 *
 * Additionally, on ESP32/S2/C3/S3/C2, this function cannot be used to hold the state of a digital GPIO during Deep-sleep.
 * Even if this function is enabled, the digital GPIO will be reset to its default state when the chip wakes up from
 * Deep-sleep. If you want to hold the state of a digital GPIO during Deep-sleep, please call `gpio_deep_sleep_hold_en`.
 *
 * Power down or call `gpio_hold_dis` will disable this function.
 *
 * @param gpio_num GPIO number, only support output-capable GPIOs
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_NOT_SUPPORTED Not support pad hold function
 */
//go:linkname GpioHoldEn C.gpio_hold_en
func GpioHoldEn(gpio_num GpioNumT) EspErrT

/**
 * @brief Disable gpio pad hold function.
 *
 * When the chip is woken up from peripheral power-down sleep, the gpio will be set to the default mode,
 * so, the gpio will output the default level if this function is called. If you don't want the level changes, the
 * gpio should be configured to a known state before this function is called.
 *  e.g.
 *     If you hold gpio18 high during Deep-sleep, after the chip is woken up and `gpio_hold_dis` is called,
 *     gpio18 will output low level(because gpio18 is input mode by default). If you don't want this behavior,
 *     you should configure gpio18 as output mode and set it to high level before calling `gpio_hold_dis`.
 *
 * @param gpio_num GPIO number, only support output-capable GPIOs
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_NOT_SUPPORTED Not support pad hold function
 */
//go:linkname GpioHoldDis C.gpio_hold_dis
func GpioHoldDis(gpio_num GpioNumT) EspErrT

/**
 * @brief Enable all digital gpio pads hold function during Deep-sleep.
 *
 * Enabling this feature makes all digital gpio pads be at the holding state during Deep-sleep. The state of each pad
 * holds is its active configuration (not pad's sleep configuration!).
 *
 * Note:
 *   1. For digital IO, this API takes effect only if the corresponding digital IO pad hold function has been enabled. You
 *      can enable the GPIO pad hold function by calling `gpio_hold_en`.
 *      has been enabled. You can call `gpio_hold_en` to enable the gpio pad hold function.
 *   2. Though this API targets all digital IOs, the pad hold feature only works when the chip is in Deep-sleep mode. When
 *      the chip is in active mode, the digital GPIO state can be changed freely even if you have called this function, except
 *      for IOs that are already held by `gpio_hold_en`.
 *
 * After this API is being called, the digital gpio Deep-sleep hold feature will work during every sleep process. You
 * should call `gpio_deep_sleep_hold_dis` to disable this feature.
 */
//go:linkname GpioDeepSleepHoldEn C.gpio_deep_sleep_hold_en
func GpioDeepSleepHoldEn()

/**
 * @brief Disable all digital gpio pads hold function during Deep-sleep.
 */
//go:linkname GpioDeepSleepHoldDis C.gpio_deep_sleep_hold_dis
func GpioDeepSleepHoldDis()

/**
 * @brief Set pad input to a peripheral signal through the IOMUX.
 * @param gpio_num GPIO number of the pad.
 * @param signal_idx Peripheral signal id to input. One of the ``*_IN_IDX`` signals in ``soc/gpio_sig_map.h``.
 */
//go:linkname GpioIomuxIn C.gpio_iomux_in
func GpioIomuxIn(gpio_num c.Uint32T, signal_idx c.Uint32T)

/**
 * @brief Set peripheral output to an GPIO pad through the IOMUX.
 * @param gpio_num gpio_num GPIO number of the pad.
 * @param func The function number of the peripheral pin to output pin.
 *        One of the ``FUNC_X_*`` of specified pin (X) in ``soc/io_mux_reg.h``.
 * @param out_en_inv True if the output enable needs to be inverted, otherwise False.
 */
//go:linkname GpioIomuxOut C.gpio_iomux_out
func GpioIomuxOut(gpio_num c.Uint8T, func_ c.Int, out_en_inv bool)

/**
 * @brief Force hold all digital and rtc gpio pads.
 *
 * GPIO force hold, no matter the chip in active mode or sleep modes.
 *
 * This function will immediately cause all pads to latch the current values of input enable, output enable,
 * output value, function, and drive strength values.
 *
 * @warning
 *   1. This function will hold flash and UART pins as well. Therefore, this function, and all code run afterwards
 *      (till calling `gpio_force_unhold_all` to disable this feature), MUST be placed in internal RAM as holding the flash
 *      pins will halt SPI flash operation, and holding the UART pins will halt any UART logging.
 *   2. The hold state of all pads will be cancelled during ROM boot, so it is not recommended to use this API to hold
 *      the pads state during deepsleep and reset.
 * */
//go:linkname GpioForceHoldAll C.gpio_force_hold_all
func GpioForceHoldAll() EspErrT

/**
 * @brief Unhold all digital and rtc gpio pads.
 *
 * @note  The global hold signal and the hold signal of each IO act on the PAD through 'or' logic, so if a pad has already
 *        been configured to hold by `gpio_hold_en`, this API can't release its hold state.
 * */
//go:linkname GpioForceUnholdAll C.gpio_force_unhold_all
func GpioForceUnholdAll() EspErrT

/**
 * @brief Enable SLP_SEL to change GPIO status automantically in lightsleep.
 * @param gpio_num GPIO number of the pad.
 *
 * @return
 *     - ESP_OK Success
 *
 */
//go:linkname GpioSleepSelEn C.gpio_sleep_sel_en
func GpioSleepSelEn(gpio_num GpioNumT) EspErrT

/**
 * @brief Disable SLP_SEL to change GPIO status automantically in lightsleep.
 * @param gpio_num GPIO number of the pad.
 *
 * @return
 *     - ESP_OK Success
 */
//go:linkname GpioSleepSelDis C.gpio_sleep_sel_dis
func GpioSleepSelDis(gpio_num GpioNumT) EspErrT

/**
 * @brief    GPIO set direction at sleep
 *
 * Configure GPIO direction,such as output_only,input_only,output_and_input
 *
 * @param  gpio_num  Configure GPIO pins number, it should be GPIO number. If you want to set direction of e.g. GPIO16, gpio_num should be GPIO_NUM_16 (16);
 * @param  mode GPIO direction
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG GPIO error
 */
//go:linkname GpioSleepSetDirection C.gpio_sleep_set_direction
func GpioSleepSetDirection(gpio_num GpioNumT, mode GpioModeT) EspErrT

/**
 * @brief  Configure GPIO pull-up/pull-down resistors at sleep
 *
 * @note ESP32: Only pins that support both input & output have integrated pull-up and pull-down resistors. Input-only GPIOs 34-39 do not.
 *
 * @param  gpio_num GPIO number. If you want to set pull up or down mode for e.g. GPIO16, gpio_num should be GPIO_NUM_16 (16);
 * @param  pull GPIO pull up/down mode.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG : Parameter error
 */
//go:linkname GpioSleepSetPullMode C.gpio_sleep_set_pull_mode
func GpioSleepSetPullMode(gpio_num GpioNumT, pull GpioPullModeT) EspErrT

/**
 * @brief Dump IO configuration information to console
 *
 * @param out_stream IO stream (e.g. stdout)
 * @param io_bit_mask IO pin bit mask, each bit maps to an IO
 *
 * @return
 *    - ESP_OK Success
 *    - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname GpioDumpIoConfiguration C.gpio_dump_io_configuration
func GpioDumpIoConfiguration(out_stream *c.FILE, io_bit_mask c.Uint64T) EspErrT

/**
 * @brief Configure a pin to perform GPIO function or an IOMUX function
 *
 * @param gpio_num GPIO number.
 * @param func Function to assign to the pin. see "io_mux_reg.h"
 *
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_ARG GPIO number error
 */
//go:linkname GpioFuncSel C.gpio_func_sel
func GpioFuncSel(gpio_num GpioNumT, func_ c.Uint32T) EspErrT

/**
 * @brief Enable output for an IO
 *
 * @param gpio_num GPIO number
 *
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_ARG GPIO number error
 */
//go:linkname GpioOutputEnable C.gpio_output_enable
func GpioOutputEnable(gpio_num GpioNumT) EspErrT

/**
 * @brief Disable output for an IO
 *
 * @param gpio_num GPIO number
 *
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_ARG GPIO number error
 */
//go:linkname GpioOutputDisable C.gpio_output_disable
func GpioOutputDisable(gpio_num GpioNumT) EspErrT

/**
 * @brief Enable open-drain for an IO
 *
 * @param gpio_num GPIO number
 *
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_ARG GPIO number error
 */
//go:linkname GpioOdEnable C.gpio_od_enable
func GpioOdEnable(gpio_num GpioNumT) EspErrT

/**
 * @brief Disable open-drain for an IO
 *
 * @param gpio_num GPIO number
 *
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_ARG GPIO number error
 */
//go:linkname GpioOdDisable C.gpio_od_disable
func GpioOdDisable(gpio_num GpioNumT) EspErrT
