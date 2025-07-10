package touch_element

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief   Button initialization configuration passed to touch_button_install
 */

type TouchButtonGlobalConfigT struct {
	ThresholdDivider c.Float
	DefaultLpTime    c.Uint32T
}

/**
 * @brief   Button configuration (for new instance) passed to touch_button_create()
 */

type TouchButtonConfigT struct {
	ChannelNum  TouchPadT
	ChannelSens c.Float
}
type TouchButtonEventT c.Int

const (
	TOUCH_BUTTON_EVT_ON_PRESS     TouchButtonEventT = 0
	TOUCH_BUTTON_EVT_ON_RELEASE   TouchButtonEventT = 1
	TOUCH_BUTTON_EVT_ON_LONGPRESS TouchButtonEventT = 2
	TOUCH_BUTTON_EVT_MAX          TouchButtonEventT = 3
)

/**
 * @brief   Button message type
 */

type TouchButtonMessageT struct {
	Event TouchButtonEventT
}
type TouchButtonHandleT TouchElemHandleT

// llgo:type C
type TouchButtonCallbackT func(TouchButtonHandleT, *TouchButtonMessageT, c.Pointer)

/**
 * @brief   Touch Button initialize
 *
 * This function initializes touch button global and acts on all
 * touch button instances.
 *
 * @param[in] global_config   Button object initialization configuration
 *
 * @return
 *      - ESP_OK: Successfully initialized touch button
 *      - ESP_ERR_INVALID_STATE: Touch element library was not initialized
 *      - ESP_ERR_INVALID_ARG: button_init is NULL
 *      - ESP_ERR_NO_MEM: Insufficient memory
 */
// llgo:link (*TouchButtonGlobalConfigT).TouchButtonInstall C.touch_button_install
func (recv_ *TouchButtonGlobalConfigT) TouchButtonInstall() EspErrT {
	return 0
}

/**
 * @brief   Release resources allocated using touch_button_install()
 */
//go:linkname TouchButtonUninstall C.touch_button_uninstall
func TouchButtonUninstall()

/**
 * @brief   Create a new touch button instance
 *
 * @param[in]  button_config    Button configuration
 * @param[out] button_handle    Button handle
 *
 * @note    The sensitivity has to be explored in experiments,
 *          Sensitivity = (Raw(touch) - Raw(release)) / Raw(release) * 100%
 *
 * @return
 *      - ESP_OK: Successfully create touch button
 *      - ESP_ERR_INVALID_STATE: Touch button driver was not initialized
 *      - ESP_ERR_NO_MEM: Insufficient memory
 *      - ESP_ERR_INVALID_ARG: Invalid configuration struct or arguments is NULL
 */
// llgo:link (*TouchButtonConfigT).TouchButtonCreate C.touch_button_create
func (recv_ *TouchButtonConfigT) TouchButtonCreate(button_handle *TouchButtonHandleT) EspErrT {
	return 0
}

/**
 * @brief Release resources allocated using touch_button_create()
 *
 * @param[in] button_handle   Button handle
 * @return
 *      - ESP_OK: Successfully released resources
 *      - ESP_ERR_INVALID_STATE: Touch button driver was not initialized
 *      - ESP_ERR_INVALID_ARG: button_handle is null
 *      - ESP_ERR_NOT_FOUND: Input handle is not a button handle
 */
//go:linkname TouchButtonDelete C.touch_button_delete
func TouchButtonDelete(button_handle TouchButtonHandleT) EspErrT

/**
 * @brief   Touch button subscribes event
 *
 * This function uses event mask to subscribe to touch button events, once one of
 * the subscribed events occurs, the event message could be retrieved by calling
 * touch_element_message_receive() or input callback routine.
 *
 * @param[in] button_handle     Button handle
 * @param[in] event_mask        Button subscription event mask
 * @param[in] arg               User input argument
 *
 * @note    Touch button only support three kind of event masks, they are
 *          TOUCH_ELEM_EVENT_ON_PRESS, TOUCH_ELEM_EVENT_ON_RELEASE, TOUCH_ELEM_EVENT_ON_LONGPRESS.
 *          You can use those event masks in any combination to achieve the desired effect.
 *
 * @return
 *      - ESP_OK: Successfully subscribed touch button event
 *      - ESP_ERR_INVALID_STATE: Touch button driver was not initialized
 *      - ESP_ERR_INVALID_ARG: button_handle is null or event is not supported
 */
//go:linkname TouchButtonSubscribeEvent C.touch_button_subscribe_event
func TouchButtonSubscribeEvent(button_handle TouchButtonHandleT, event_mask c.Uint32T, arg c.Pointer) EspErrT

/**
 * @brief   Touch button set dispatch method
 *
 * This function sets a dispatch method that the driver core will use
 * this method as the event notification method.
 *
 * @param[in] button_handle     Button handle
 * @param[in] dispatch_method   Dispatch method (By callback/event)
 *
 * @return
 *      - ESP_OK: Successfully set dispatch method
 *      - ESP_ERR_INVALID_STATE: Touch button driver was not initialized
 *      - ESP_ERR_INVALID_ARG: button_handle is null or dispatch_method is invalid
 */
//go:linkname TouchButtonSetDispatchMethod C.touch_button_set_dispatch_method
func TouchButtonSetDispatchMethod(button_handle TouchButtonHandleT, dispatch_method TouchElemDispatchT) EspErrT

/**
 * @brief   Touch button set callback
 *
 * This function sets a callback routine into touch element driver core,
 * when the subscribed events occur, the callback routine will be called.
 *
 * @param[in] button_handle     Button handle
 * @param[in] button_callback   User input callback
 *
 * @note        Button message will be passed from the callback function and it will be destroyed when the callback function return.
 *
 * @warning     Since this input callback routine runs on driver core (esp-timer callback routine),
 *              it should not do something that attempts to Block, such as calling vTaskDelay().
 *
 * @return
 *      - ESP_OK: Successfully set callback
 *      - ESP_ERR_INVALID_STATE: Touch button driver was not initialized
 *      - ESP_ERR_INVALID_ARG: button_handle or button_callback is null
 */
//go:linkname TouchButtonSetCallback C.touch_button_set_callback
func TouchButtonSetCallback(button_handle TouchButtonHandleT, button_callback TouchButtonCallbackT) EspErrT

/**
 * @brief   Touch button set long press trigger time
 *
 * This function sets the threshold time (ms) for a long press event. If a button is pressed
 * and held for a period of time that exceeds the threshold time, a long press event is triggered.
 *
 * @param[in] button_handle     Button handle
 * @param[in] threshold_time    Threshold time (ms) of long press event occur
 *
 * @return
 *      - ESP_OK: Successfully set the threshold time of long press event
 *      - ESP_ERR_INVALID_STATE: Touch button driver was not initialized
 *      - ESP_ERR_INVALID_ARG: button_handle is null or time (ms) is not lager than 0
 */
//go:linkname TouchButtonSetLongpress C.touch_button_set_longpress
func TouchButtonSetLongpress(button_handle TouchButtonHandleT, threshold_time c.Uint32T) EspErrT

/**
 * @brief   Touch button get message
 *
 * This function decodes the element message from touch_element_message_receive() and return
 * a button message pointer.
 *
 * @param[in] element_message   element message
 *
 * @return  Touch button message pointer
 */
// llgo:link (*TouchElemMessageT).TouchButtonGetMessage C.touch_button_get_message
func (recv_ *TouchElemMessageT) TouchButtonGetMessage() *TouchButtonMessageT {
	return nil
}
