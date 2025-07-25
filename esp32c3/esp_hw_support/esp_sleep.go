package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type EspDeepSleepCbT func()
type EspDeepsleepGpioWakeUpModeT c.Int

const (
	ESP_GPIO_WAKEUP_GPIO_LOW  EspDeepsleepGpioWakeUpModeT = 0
	ESP_GPIO_WAKEUP_GPIO_HIGH EspDeepsleepGpioWakeUpModeT = 1
)

type EspSleepPdDomainT c.Int

const (
	ESP_PD_DOMAIN_XTAL    EspSleepPdDomainT = 0
	ESP_PD_DOMAIN_RC_FAST EspSleepPdDomainT = 1
	ESP_PD_DOMAIN_CPU     EspSleepPdDomainT = 2
	ESP_PD_DOMAIN_VDDSDIO EspSleepPdDomainT = 3
	ESP_PD_DOMAIN_MAX     EspSleepPdDomainT = 4
)

type EspSleepPdOptionT c.Int

const (
	ESP_PD_OPTION_OFF  EspSleepPdOptionT = 0
	ESP_PD_OPTION_ON   EspSleepPdOptionT = 1
	ESP_PD_OPTION_AUTO EspSleepPdOptionT = 2
)

type EspSleepSourceT c.Int

const (
	ESP_SLEEP_WAKEUP_UNDEFINED       EspSleepSourceT = 0
	ESP_SLEEP_WAKEUP_ALL             EspSleepSourceT = 1
	ESP_SLEEP_WAKEUP_EXT0            EspSleepSourceT = 2
	ESP_SLEEP_WAKEUP_EXT1            EspSleepSourceT = 3
	ESP_SLEEP_WAKEUP_TIMER           EspSleepSourceT = 4
	ESP_SLEEP_WAKEUP_TOUCHPAD        EspSleepSourceT = 5
	ESP_SLEEP_WAKEUP_ULP             EspSleepSourceT = 6
	ESP_SLEEP_WAKEUP_GPIO            EspSleepSourceT = 7
	ESP_SLEEP_WAKEUP_UART            EspSleepSourceT = 8
	ESP_SLEEP_WAKEUP_WIFI            EspSleepSourceT = 9
	ESP_SLEEP_WAKEUP_COCPU           EspSleepSourceT = 10
	ESP_SLEEP_WAKEUP_COCPU_TRAP_TRIG EspSleepSourceT = 11
	ESP_SLEEP_WAKEUP_BT              EspSleepSourceT = 12
	ESP_SLEEP_WAKEUP_VAD             EspSleepSourceT = 13
)

type EspSleepModeT c.Int

const (
	ESP_SLEEP_MODE_LIGHT_SLEEP EspSleepModeT = 0
	ESP_SLEEP_MODE_DEEP_SLEEP  EspSleepModeT = 1
)

type EspSleepWakeupCauseT EspSleepSourceT

const (
	ESP_ERR_SLEEP_REJECT                   c.Int = 259
	ESP_ERR_SLEEP_TOO_SHORT_SLEEP_DURATION c.Int = 258
)

/**
 * @brief Disable wakeup source
 *
 * This function is used to deactivate wake up trigger for source
 * defined as parameter of the function.
 *
 * @note This function does not modify wake up configuration in RTC.
 *       It will be performed in esp_deep_sleep_start/esp_light_sleep_start function.
 *
 * See docs/sleep-modes.rst for details.
 *
 * @param source - number of source to disable of type esp_sleep_source_t
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_STATE if trigger was not active
 */
// llgo:link EspSleepSourceT.EspSleepDisableWakeupSource C.esp_sleep_disable_wakeup_source
func (recv_ EspSleepSourceT) EspSleepDisableWakeupSource() EspErrT {
	return 0
}

/**
 * @brief Enable wakeup by timer
 * @param time_in_us  time before wakeup, in microseconds
 * @note  The valid `time_in_us` value depends on the bit width of the lp_timer/rtc_timer counter and the
 *        current slow clock source selection (Refer RTC clock source configuration in menuconfig).
 *        Valid values should be positive values less than RTC slow clock period * (2 ^ RTC timer bitwidth).
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if value is out of range.
 */
//go:linkname EspSleepEnableTimerWakeup C.esp_sleep_enable_timer_wakeup
func EspSleepEnableTimerWakeup(time_in_us c.Uint64T) EspErrT

/**
 * @brief Returns true if a GPIO number is valid for use as wakeup source.
 *
 * @note For SoCs with RTC IO capability, this can be any valid RTC IO input pin.
 *
 * @param gpio_num Number of the GPIO to test for wakeup source capability
 *
 * @return True if this GPIO number will be accepted as a sleep wakeup source.
 */
//go:linkname EspSleepIsValidWakeupGpio C.esp_sleep_is_valid_wakeup_gpio
func EspSleepIsValidWakeupGpio(gpio_num GpioNumT) bool

/**
 * @brief Enable wakeup using specific gpio pins
 *
 * This function enables an IO pin to wake up the chip from deep sleep.
 *
 * @note 1.This function does not modify pin configuration. The pins are configured
 *          inside `esp_deep_sleep_start`, immediately before entering sleep mode.
 *       2.This function is also applicable to waking up the lightsleep when the peripheral
 *         power domain is powered off, see PM_POWER_DOWN_PERIPHERAL_IN_LIGHT_SLEEP in menuconfig.
 *
 * @note You don't need to worry about pull-up or pull-down resistors before
 *       using this function because the ESP_SLEEP_GPIO_ENABLE_INTERNAL_RESISTORS
 *       option is enabled by default. It will automatically set pull-up or pull-down
 *       resistors internally in esp_deep_sleep_start based on the wakeup mode. However,
 *       when using external pull-up or pull-down resistors, please be sure to disable
 *       the ESP_SLEEP_GPIO_ENABLE_INTERNAL_RESISTORS option, as the combination of internal
 *       and external resistors may cause interference. BTW, when you use low level to wake up the
 *       chip, we strongly recommend you to add external resistors (pull-up).
 *
 * @param gpio_pin_mask  Bit mask of GPIO numbers which will cause wakeup. Only GPIOs
 *              which have RTC functionality (pads that powered by VDD3P3_RTC) can be used in this bit map.
 * @param mode Select logic function used to determine wakeup condition:
 *            - ESP_GPIO_WAKEUP_GPIO_LOW: wake up when the gpio turn to low.
 *            - ESP_GPIO_WAKEUP_GPIO_HIGH: wake up when the gpio turn to high.
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the mask contains any invalid deep sleep wakeup pin or wakeup mode is invalid
 */
//go:linkname EspDeepSleepEnableGpioWakeup C.esp_deep_sleep_enable_gpio_wakeup
func EspDeepSleepEnableGpioWakeup(gpio_pin_mask c.Uint64T, mode EspDeepsleepGpioWakeUpModeT) EspErrT

/**
 * @brief Enable wakeup from light sleep using GPIOs
 *
 * Each GPIO supports wakeup function, which can be triggered on either low level
 * or high level. Unlike EXT0 and EXT1 wakeup sources, this method can be used
 * both for all IOs: RTC IOs and digital IOs. It can only be used to wakeup from
 * light sleep though.
 *
 * To enable wakeup, first call gpio_wakeup_enable, specifying gpio number and
 * wakeup level, for each GPIO which is used for wakeup.
 * Then call this function to enable wakeup feature.
 *
 * @note 1. On ESP32, GPIO wakeup source can not be used together with touch or ULP wakeup sources.
 *       2. If PM_POWER_DOWN_PERIPHERAL_IN_LIGHT_SLEEP is enabled (if target supported),
 *          this API is unavailable since the GPIO module is powered down during sleep.
 *          You can use `esp_deep_sleep_enable_gpio_wakeup` instead, or use EXT1 wakeup source
 *          by `esp_sleep_enable_ext1_wakeup_io` to achieve the same function.
 *          (Only GPIOs which have RTC functionality can be used)
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_STATE if wakeup triggers conflict
 */
//go:linkname EspSleepEnableGpioWakeup C.esp_sleep_enable_gpio_wakeup
func EspSleepEnableGpioWakeup() EspErrT

/**
 * @brief Enable wakeup from light sleep using UART
 *
 * Use uart_set_wakeup_threshold function to configure UART wakeup threshold.
 *
 * Wakeup from light sleep takes some time, so not every character sent
 * to the UART can be received by the application.
 *
 * @note 1. ESP32 does not support wakeup from UART2.
 *       2. If PM_POWER_DOWN_PERIPHERAL_IN_LIGHT_SLEEP is enabled (if target supported),
 *          this API is unavailable since the UART module is powered down during sleep.
 *
 * @param uart_num  UART port to wake up from
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if wakeup from given UART is not supported
 */
//go:linkname EspSleepEnableUartWakeup C.esp_sleep_enable_uart_wakeup
func EspSleepEnableUartWakeup(uart_num c.Int) EspErrT

/**
 * @brief Enable wakeup by bluetooth
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NOT_SUPPORTED if wakeup from bluetooth is not supported
 */
//go:linkname EspSleepEnableBtWakeup C.esp_sleep_enable_bt_wakeup
func EspSleepEnableBtWakeup() EspErrT

/**
 * @brief Disable wakeup by bluetooth
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NOT_SUPPORTED if wakeup from bluetooth is not supported
 */
//go:linkname EspSleepDisableBtWakeup C.esp_sleep_disable_bt_wakeup
func EspSleepDisableBtWakeup() EspErrT

/**
 * @brief Enable wakeup by WiFi MAC
 * @return
 *      - ESP_OK on success
 */
//go:linkname EspSleepEnableWifiWakeup C.esp_sleep_enable_wifi_wakeup
func EspSleepEnableWifiWakeup() EspErrT

/**
 * @brief Disable wakeup by WiFi MAC
 * @return
 *      - ESP_OK on success
 */
//go:linkname EspSleepDisableWifiWakeup C.esp_sleep_disable_wifi_wakeup
func EspSleepDisableWifiWakeup() EspErrT

/**
 * @brief Enable beacon wakeup by WiFi MAC, it will wake up the system into modem state
 * @return
 *      - ESP_OK on success
 */
//go:linkname EspSleepEnableWifiBeaconWakeup C.esp_sleep_enable_wifi_beacon_wakeup
func EspSleepEnableWifiBeaconWakeup() EspErrT

/**
 * @brief Disable beacon wakeup by WiFi MAC
 * @return
 *      - ESP_OK on success
 */
//go:linkname EspSleepDisableWifiBeaconWakeup C.esp_sleep_disable_wifi_beacon_wakeup
func EspSleepDisableWifiBeaconWakeup() EspErrT

/**
 * @brief Get the bit mask of GPIOs which caused wakeup (gpio)
 *
 * If wakeup was caused by another source, this function will return 0.
 *
 * @return bit mask, if GPIOn caused wakeup, BIT(n) will be set
 */
//go:linkname EspSleepGetGpioWakeupStatus C.esp_sleep_get_gpio_wakeup_status
func EspSleepGetGpioWakeupStatus() c.Uint64T

/**
 * @brief Set power down mode for an RTC power domain in sleep mode
 *
 * If not set set using this API, all power domains default to ESP_PD_OPTION_AUTO.
 *
 * @param domain  power domain to configure
 * @param option  power down option (ESP_PD_OPTION_OFF, ESP_PD_OPTION_ON, or ESP_PD_OPTION_AUTO)
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if either of the arguments is out of range
 */
// llgo:link EspSleepPdDomainT.EspSleepPdConfig C.esp_sleep_pd_config
func (recv_ EspSleepPdDomainT) EspSleepPdConfig(option EspSleepPdOptionT) EspErrT {
	return 0
}

/**
 * @brief Enter deep sleep with the configured wakeup options
 *
 * @note In general, the function does not return, but if the sleep is rejected,
 * then it returns from it.
 *
 * The reason for the rejection can be such as a short sleep time.
 *
 * @return
 *  - No return - If the sleep is not rejected.
 *  - ESP_ERR_SLEEP_REJECT sleep request is rejected(wakeup source set before the sleep request)
 */
//go:linkname EspDeepSleepTryToStart C.esp_deep_sleep_try_to_start
func EspDeepSleepTryToStart() EspErrT

/**
 * @brief Enter deep sleep with the configured wakeup options
 *
 * @note The function does not do a return (no rejection). Even if wakeup source set before the sleep request
 * it goes to deep sleep anyway.
 */
//go:linkname EspDeepSleepStart C.esp_deep_sleep_start
func EspDeepSleepStart()

/**
 * @brief Enter light sleep with the configured wakeup options
 *
 * @return
 *  - ESP_OK on success (returned after wakeup)
 *  - ESP_ERR_SLEEP_REJECT sleep request is rejected(wakeup source set before the sleep request)
 *  - ESP_ERR_SLEEP_TOO_SHORT_SLEEP_DURATION after deducting the sleep flow overhead, the final sleep duration
 *                                           is too short to cover the minimum sleep duration of the chip, when
 *                                           rtc timer wakeup source enabled
 */
//go:linkname EspLightSleepStart C.esp_light_sleep_start
func EspLightSleepStart() EspErrT

/**
 * @brief Enter deep-sleep mode
 *
 * The device will automatically wake up after the deep-sleep time
 * Upon waking up, the device calls deep sleep wake stub, and then proceeds
 * to load application.
 *
 * Call to this function is equivalent to a call to esp_deep_sleep_enable_timer_wakeup
 * followed by a call to esp_deep_sleep_start.
 *
 * @note In general, the function does not return, but if the sleep is rejected,
 * then it returns from it.
 *
 * The reason for the rejection can be such as a short sleep time.
 *
 * @param time_in_us  deep-sleep time, unit: microsecond
 *
 * @return
 *  - No return - If the sleep is not rejected.
 *  - ESP_ERR_SLEEP_REJECT sleep request is rejected(wakeup source set before the sleep request)
 */
//go:linkname EspDeepSleepTry C.esp_deep_sleep_try
func EspDeepSleepTry(time_in_us c.Uint64T) EspErrT

/**
 * @brief Enter deep-sleep mode
 *
 * The device will automatically wake up after the deep-sleep time
 * Upon waking up, the device calls deep sleep wake stub, and then proceeds
 * to load application.
 *
 * Call to this function is equivalent to a call to esp_deep_sleep_enable_timer_wakeup
 * followed by a call to esp_deep_sleep_start.
 *
 * @note The function does not do a return (no rejection).. Even if wakeup source set before the sleep request
 * it goes to deep sleep anyway.
 *
 * @param time_in_us  deep-sleep time, unit: microsecond
 */
//go:linkname EspDeepSleep C.esp_deep_sleep
func EspDeepSleep(time_in_us c.Uint64T)

/**
 * @brief Register a callback to be called from the deep sleep prepare
 *
 * @warning deepsleep callbacks should without parameters, and MUST NOT,
 *          UNDER ANY CIRCUMSTANCES, CALL A FUNCTION THAT MIGHT BLOCK.
 *
 * @param new_dslp_cb     Callback to be called
 *
 * @return
 *     - ESP_OK:         Callback registered to the deepsleep misc_modules_sleep_prepare
 *     - ESP_ERR_NO_MEM: No more hook space for register the callback
 */
//go:linkname EspDeepSleepRegisterHook C.esp_deep_sleep_register_hook
func EspDeepSleepRegisterHook(new_dslp_cb EspDeepSleepCbT) EspErrT

/**
 * @brief Unregister an deepsleep callback
 *
 * @param old_dslp_cb     Callback to be unregistered
 */
//go:linkname EspDeepSleepDeregisterHook C.esp_deep_sleep_deregister_hook
func EspDeepSleepDeregisterHook(old_dslp_cb EspDeepSleepCbT)

/**
 * @brief Get the wakeup source which caused wakeup from sleep
 *
 * @return cause of wake up from last sleep (deep sleep or light sleep)
 */
//go:linkname EspSleepGetWakeupCause C.esp_sleep_get_wakeup_cause
func EspSleepGetWakeupCause() EspSleepWakeupCauseT

/**
 * @brief Default stub to run on wake from deep sleep.
 *
 * Allows for executing code immediately on wake from sleep, before
 * the software bootloader or ESP-IDF app has started up.
 *
 * This function is weak-linked, so you can implement your own version
 * to run code immediately when the chip wakes from
 * sleep.
 *
 * See docs/deep-sleep-stub.rst for details.
 */
//go:linkname EspWakeDeepSleep C.esp_wake_deep_sleep
func EspWakeDeepSleep()

// llgo:type C
type EspDeepSleepWakeStubFnT func()

/**
 * @brief Install a new stub at runtime to run on wake from deep sleep
 *
 * If implementing esp_wake_deep_sleep() then it is not necessary to
 * call this function.
 *
 * However, it is possible to call this function to substitute a
 * different deep sleep stub. Any function used as a deep sleep stub
 * must be marked RTC_IRAM_ATTR, and must obey the same rules given
 * for esp_wake_deep_sleep().
 */
//go:linkname EspSetDeepSleepWakeStub C.esp_set_deep_sleep_wake_stub
func EspSetDeepSleepWakeStub(new_stub EspDeepSleepWakeStubFnT)

/**
 * @brief Get current wake from deep sleep stub
 * @return Return current wake from deep sleep stub, or NULL if
 *         no stub is installed.
 */
//go:linkname EspGetDeepSleepWakeStub C.esp_get_deep_sleep_wake_stub
func EspGetDeepSleepWakeStub() EspDeepSleepWakeStubFnT

/**
 *  @brief The default esp-idf-provided esp_wake_deep_sleep() stub.
 *
 *  See docs/deep-sleep-stub.rst for details.
 */
//go:linkname EspDefaultWakeDeepSleep C.esp_default_wake_deep_sleep
func EspDefaultWakeDeepSleep()

/**
 * @brief Disable logging from the ROM code after deep sleep.
 *
 * Using LSB of RTC_STORE4.
 */
//go:linkname EspDeepSleepDisableRomLogging C.esp_deep_sleep_disable_rom_logging
func EspDeepSleepDisableRomLogging()

/**
 * @brief CPU Power down low-level initialize, enable CPU power down during light sleep
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NO_MEM not enough retention memory
 */
//go:linkname EspSleepCpuPdLowInit C.esp_sleep_cpu_pd_low_init
func EspSleepCpuPdLowInit() EspErrT

/**
 * @brief CPU Power down low-level deinitialize, disable CPU power down during light sleep
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NO_MEM not enough retention memory
 */
//go:linkname EspSleepCpuPdLowDeinit C.esp_sleep_cpu_pd_low_deinit
func EspSleepCpuPdLowDeinit() EspErrT

/**
 * @brief CPU Power down initialize
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NO_MEM not enough retention memory
 */
//go:linkname EspSleepCpuRetentionInit C.esp_sleep_cpu_retention_init
func EspSleepCpuRetentionInit() EspErrT

/**
 * @brief CPU Power down de-initialize
 *
 * @return
 *      - ESP_OK on success
 *
 * Release system retention memory.
 */
//go:linkname EspSleepCpuRetentionDeinit C.esp_sleep_cpu_retention_deinit
func EspSleepCpuRetentionDeinit() EspErrT

/**
 * @brief Configure to isolate all GPIO pins in sleep state
 */
//go:linkname EspSleepConfigGpioIsolate C.esp_sleep_config_gpio_isolate
func EspSleepConfigGpioIsolate()

/**
 * @brief Enable or disable GPIO pins status switching between slept status and waked status.
 * @param enable decide whether to switch status or not
 */
//go:linkname EspSleepEnableGpioSwitch C.esp_sleep_enable_gpio_switch
func EspSleepEnableGpioSwitch(enable bool)
