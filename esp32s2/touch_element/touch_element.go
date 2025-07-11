package touch_element

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* -------------------------------- Global hardware & software configuration struct --------------------------------- */
/**
 * @brief   Touch element software configuration
 */

type TouchElemSwConfigT struct {
	WaterproofThresholdDivider c.Float
	ProcessingPeriod           c.Uint8T
	IntrMessageSize            c.Uint8T
	EventMessageSize           c.Uint8T
}

/**
 * @brief   Touch element hardware configuration
 */

type TouchElemHwConfigT struct {
	UpperVoltage                  TouchHighVoltT
	VoltageAttenuation            TouchVoltAttenT
	LowerVoltage                  TouchLowVoltT
	SuspendChannelPolarity        TouchPadConnTypeT
	DenoiseLevel                  TouchPadDenoiseGradeT
	DenoiseEquivalentCap          TouchPadDenoiseCapT
	SmoothFilterMode              TouchSmoothModeT
	BenchmarkFilterMode           TouchFilterModeT
	SampleCount                   c.Uint16T
	SleepCycle                    c.Uint16T
	BenchmarkDebounceCount        c.Uint8T
	BenchmarkCalibrationThreshold c.Uint8T
	BenchmarkJitterStep           c.Uint8T
}

/**
 * @brief   Touch element global configuration passed to touch_element_install
 */

type TouchElemGlobalConfigT struct {
	Hardware TouchElemHwConfigT
	Software TouchElemSwConfigT
}

/**
 * @brief   Touch element waterproof configuration passed to touch_element_waterproof_install
 */

type TouchElemWaterproofConfigT struct {
	GuardChannel     TouchPadT
	GuardSensitivity c.Float
}

/**
 * @brief   Touch element sleep configuration passed to touch_element_enable_light_sleep or touch_element_enable_deep_sleep
 */
type TouchElemSleepConfigT struct {
	SampleCount c.Uint16T
	SleepCycle  c.Uint16T
}
type TouchElemHandleT c.Pointer
type TouchElemEventT c.Uint32T
type TouchElemTypeT c.Int

const (
	TOUCH_ELEM_TYPE_BUTTON TouchElemTypeT = 0
	TOUCH_ELEM_TYPE_SLIDER TouchElemTypeT = 1
	TOUCH_ELEM_TYPE_MATRIX TouchElemTypeT = 2
)

type TouchElemDispatchT c.Int

const (
	TOUCH_ELEM_DISP_EVENT    TouchElemDispatchT = 0
	TOUCH_ELEM_DISP_CALLBACK TouchElemDispatchT = 1
	TOUCH_ELEM_DISP_MAX      TouchElemDispatchT = 2
)

/**
 * @brief   Touch element event message type from touch_element_message_receive()
 */

type TouchElemMessageT struct {
	Handle      TouchElemHandleT
	ElementType TouchElemTypeT
	Arg         c.Pointer
	ChildMsg    [8]c.Uint8T
}

/**
 * @brief   Touch element processing initialization
 *
 * @param[in]   global_config   Global initialization configuration structure
 *
 * @note    To reinitialize the touch element object, call touch_element_uninstall() first
 *
 * @return
 *      - ESP_OK: Successfully initialized
 *      - ESP_ERR_INVALID_ARG: Invalid argument
 *      - ESP_ERR_NO_MEM: Insufficient memory
 *      - ESP_ERR_INVALID_STATE: Touch element is already initialized
 *      - Others: Unknown touch driver layer or lower layer error
 */
// llgo:link (*TouchElemGlobalConfigT).TouchElementInstall C.touch_element_install
func (recv_ *TouchElemGlobalConfigT) TouchElementInstall() EspErrT {
	return 0
}

/**
 * @brief   Touch element processing start
 *
 * This function starts the touch element processing system
 *
 * @note    This function must only be called after all the touch element instances finished creating
 *
 * @return
 *      - ESP_OK: Successfully started to process
 *      - Others: Unknown touch driver layer or lower layer error
 */
//go:linkname TouchElementStart C.touch_element_start
func TouchElementStart() EspErrT

/**
 * @brief   Touch element processing stop
 *
 * This function stops the touch element processing system
 *
 * @note    This function must be called before changing the system (hardware, software) parameters
 *
 * @return
 *      - ESP_OK: Successfully stopped to process
 *      - Others: Unknown touch driver layer or lower layer error
 */
//go:linkname TouchElementStop C.touch_element_stop
func TouchElementStop() EspErrT

/**
 * @brief   Release resources allocated using touch_element_install
 *
 */
//go:linkname TouchElementUninstall C.touch_element_uninstall
func TouchElementUninstall()

/**
 * @brief   Get current event message of touch element instance
 *
 * This function will receive the touch element message (handle, event type, etc...)
 * from te_event_give(). It will block until a touch element event or a timeout occurs.
 *
 * @param[out]  element_message          Touch element event message structure
 * @param[in]   ticks_to_wait   Number of FreeRTOS ticks to block for waiting event
 * @return
 *      - ESP_OK: Successfully received touch element event
 *      - ESP_ERR_INVALID_STATE: Touch element library is not initialized
 *      - ESP_ERR_INVALID_ARG: element_message is null
 *      - ESP_ERR_TIMEOUT: Timed out waiting for event
 */
// llgo:link (*TouchElemMessageT).TouchElementMessageReceive C.touch_element_message_receive
func (recv_ *TouchElemMessageT) TouchElementMessageReceive(ticks_to_wait c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief   Touch element waterproof initialization
 *
 * This function enables the hardware waterproof, then touch element system uses Shield-Sensor
 * and Guard-Sensor to mitigate the influence of water-drop and water-stream.
 *
 * @param[in] waterproof_config     Waterproof configuration
 *
 * @note    If the waterproof function is used, Shield-Sensor can not be disabled and it will use channel 14 as
 *          it's internal channel. Hence, the user can not use channel 14 for another propose. And the Guard-Sensor
 *          is not necessary since it is optional.
 *
 * @note    Shield-Sensor: It always uses channel 14 as the shield channel, so user must connect
 *          the channel 14 and Shield-Layer in PCB since it will generate a synchronous signal automatically
 *
 * @note    Guard-Sensor: This function is optional. If used, the user must connect the guard channel and Guard-Ring
 *          in PCB. Any channels user wants to protect should be added into Guard-Ring in PCB.
 *
 * @return
 *      - ESP_OK: Successfully initialized
 *      - ESP_ERR_INVALID_STATE: Touch element library is not initialized
 *      - ESP_ERR_INVALID_ARG: waterproof_config is null or invalid Guard-Sensor channel
 *      - ESP_ERR_NO_MEM: Insufficient memory
 */
// llgo:link (*TouchElemWaterproofConfigT).TouchElementWaterproofInstall C.touch_element_waterproof_install
func (recv_ *TouchElemWaterproofConfigT) TouchElementWaterproofInstall() EspErrT {
	return 0
}

/**
 * @brief   Release resources allocated using touch_element_waterproof_install()
 */
//go:linkname TouchElementWaterproofUninstall C.touch_element_waterproof_uninstall
func TouchElementWaterproofUninstall()

/**
 * @brief   Add a masked handle to protect while Guard-Sensor has been triggered
 *
 * This function will add an application handle (button, slider, etc...) as a masked handle. While Guard-Sensor
 * has been triggered, waterproof function will start working and lock the application internal state. While the
 * influence of water is reduced, the application will be unlock and reset into IDLE state.
 *
 * @param[in] element_handle     Touch element instance handle
 *
 * @note    The waterproof protection logic must follow the real circuit in PCB, it means that all of the channels
 *          inside the input handle must be inside the Guard-Ring in real circuit.
 *
 * @return
 *      - ESP_OK: Successfully added a masked handle
 *      - ESP_ERR_INVALID_STATE: Waterproof is not initialized
 *      - ESP_ERR_INVALID_ARG: element_handle is null
 */
//go:linkname TouchElementWaterproofAdd C.touch_element_waterproof_add
func TouchElementWaterproofAdd(element_handle TouchElemHandleT) EspErrT

/**
 * @brief   Remove a masked handle to protect
 *
 * This function will remove an application handle from masked handle table.
 *
 * @param[in] element_handle     Touch element instance handle
 *
 * @return
 *      - ESP_OK: Successfully removed a masked handle
 *      - ESP_ERR_INVALID_STATE: Waterproof is not initialized
 *      - ESP_ERR_INVALID_ARG: element_handle is null
 *      - ESP_ERR_NOT_FOUND: Failed to search element_handle from waterproof mask_handle list
 */
//go:linkname TouchElementWaterproofRemove C.touch_element_waterproof_remove
func TouchElementWaterproofRemove(element_handle TouchElemHandleT) EspErrT

/**
 * @brief   Touch element light sleep initialization
 *
 * @note    It should be called after touch button element installed.
 *          Any of installed touch element can wake up from the light sleep
 *
 * @param[in] sleep_config Sleep configurations, set NULL to use default config
 * @return
 *      - ESP_OK: Successfully initialized touch sleep
 *      - ESP_ERR_INVALID_STATE: Touch element is not installed or touch sleep has been installed
 *      - ESP_ERR_INVALID_ARG: inputed argument is NULL
 *      - ESP_ERR_NO_MEM: no memory for touch sleep struct
 *      - ESP_ERR_NOT_SUPPORTED: inputed wakeup_elem_handle is not touch_button_handle_t type, currently only touch_button_handle_t supported
 */
// llgo:link (*TouchElemSleepConfigT).TouchElementEnableLightSleep C.touch_element_enable_light_sleep
func (recv_ *TouchElemSleepConfigT) TouchElementEnableLightSleep() EspErrT {
	return 0
}

/**
 * @brief   Release the resources that allocated by touch_element_enable_deep_sleep()
 *
 * This function will also disable the touch sensor to wake up the device
 *
 * @return
 *      - ESP_OK: uninstall success
 *      - ESP_ERR_INVALID_STATE: touch sleep has not been installed
 */
//go:linkname TouchElementDisableLightSleep C.touch_element_disable_light_sleep
func TouchElementDisableLightSleep() EspErrT

/**
 * @brief   Touch element deep sleep initialization
 *
 * This function will enable the device wake-up from deep sleep or light sleep by touch sensor
 *
 * @note    It should be called after touch button element installed.
 *          Only one touch button can be registered as the deep sleep wake-up button
 *
 * @param[in] wakeup_elem_handle    Touch element instance handle for waking up the device, only support button element
 * @param[in] sleep_config          Sleep configurations, set NULL to use default config
 *
 * @return
 *      - ESP_OK: Successfully initialized touch sleep
 *      - ESP_ERR_INVALID_STATE: Touch element is not installed or touch sleep has been installed
 *      - ESP_ERR_INVALID_ARG: inputed argument is NULL
 *      - ESP_ERR_NO_MEM: no memory for touch sleep struct
 *      - ESP_ERR_NOT_SUPPORTED: inputed wakeup_elem_handle is not touch_button_handle_t type, currently only touch_button_handle_t supported
 */
//go:linkname TouchElementEnableDeepSleep C.touch_element_enable_deep_sleep
func TouchElementEnableDeepSleep(wakeup_elem_handle TouchElemHandleT, sleep_config *TouchElemSleepConfigT) EspErrT

/**
 * @brief   Release the resources that allocated by touch_element_enable_deep_sleep()
 *
 * This function will also disable the touch sensor to wake up the device
 *
 * @return
 *      - ESP_OK: uninstall success
 *      - ESP_ERR_INVALID_STATE: touch sleep has not been installed
 */
//go:linkname TouchElementDisableDeepSleep C.touch_element_disable_deep_sleep
func TouchElementDisableDeepSleep() EspErrT

/**
 * @brief   Touch element wake up calibrations
 *
 * This function will also disable the touch sensor to wake up the device
 *
 * @return
 *      - ESP_OK: uninstall success
 *      - ESP_ERR_INVALID_STATE: touch sleep has not been installed
 */
//go:linkname TouchElementSleepEnableWakeupCalibration C.touch_element_sleep_enable_wakeup_calibration
func TouchElementSleepEnableWakeupCalibration(element_handle TouchElemHandleT, en bool) EspErrT
