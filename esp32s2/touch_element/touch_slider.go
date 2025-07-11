package touch_element

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief   Slider initialization configuration passed to touch_slider_install
 */

type TouchSliderGlobalConfigT struct {
	QuantifyLowerThreshold c.Float
	ThresholdDivider       c.Float
	FilterResetTime        c.Uint16T
	BenchmarkUpdateTime    c.Uint16T
	PositionFilterSize     c.Uint8T
	PositionFilterFactor   c.Uint8T
	CalculateChannelCount  c.Uint8T
}

/**
 * @brief   Slider configuration (for new instance) passed to touch_slider_create()
 */

type TouchSliderConfigT struct {
	ChannelArray     *TouchPadT
	SensitivityArray *c.Float
	ChannelNum       c.Uint8T
	PositionRange    c.Uint8T
}
type TouchSliderEventT c.Int

const (
	TOUCH_SLIDER_EVT_ON_PRESS       TouchSliderEventT = 0
	TOUCH_SLIDER_EVT_ON_RELEASE     TouchSliderEventT = 1
	TOUCH_SLIDER_EVT_ON_CALCULATION TouchSliderEventT = 2
	TOUCH_SLIDER_EVT_MAX            TouchSliderEventT = 3
)

type TouchSliderPositionT c.Uint32T

/**
 * @brief   Slider message type
 */

type TouchSliderMessageT struct {
	Event    TouchSliderEventT
	Position TouchSliderPositionT
}
type TouchSliderHandleT TouchElemHandleT

// llgo:type C
type TouchSliderCallbackT func(TouchSliderHandleT, *TouchSliderMessageT, c.Pointer)

/**
 * @brief   Touch slider initialize
 *
 * This function initializes touch slider object and acts on all
 * touch slider instances.
 *
 * @param[in] global_config   Touch slider global initialization configuration
 *
 * @return
 *      - ESP_OK: Successfully initialized touch slider
 *      - ESP_ERR_INVALID_STATE: Touch element library was not initialized
 *      - ESP_ERR_INVALID_ARG: slider_init is NULL
 *      - ESP_ERR_NO_MEM: Insufficient memory
 */
// llgo:link (*TouchSliderGlobalConfigT).TouchSliderInstall C.touch_slider_install
func (recv_ *TouchSliderGlobalConfigT) TouchSliderInstall() EspErrT {
	return 0
}

/**
 * @brief   Release resources allocated using touch_slider_install()
 *
 */
//go:linkname TouchSliderUninstall C.touch_slider_uninstall
func TouchSliderUninstall()

/**
* @brief   Create a new touch slider instance
*
* @param[in]  slider_config     Slider configuration
* @param[out] slider_handle     Slider handle
*
* @note    The index of Channel array and sensitivity array must be one-one correspondence
*
* @return
*      - ESP_OK: Successfully create touch slider
*      - ESP_ERR_INVALID_STATE: Touch slider driver was not initialized
*      - ESP_ERR_INVALID_ARG: Invalid configuration struct or arguments is NULL
*      - ESP_ERR_NO_MEM: Insufficient memory
 */
// llgo:link (*TouchSliderConfigT).TouchSliderCreate C.touch_slider_create
func (recv_ *TouchSliderConfigT) TouchSliderCreate(slider_handle *TouchSliderHandleT) EspErrT {
	return 0
}

/**
 * @brief   Release resources allocated using touch_slider_create
 *
 * @param[in] slider_handle   Slider handle
 * @return
 *      - ESP_OK: Successfully released resources
 *      - ESP_ERR_INVALID_STATE: Touch slider driver was not initialized
 *      - ESP_ERR_INVALID_ARG: slider_handle is null
 *      - ESP_ERR_NOT_FOUND: Input handle is not a slider handle
 */
//go:linkname TouchSliderDelete C.touch_slider_delete
func TouchSliderDelete(slider_handle TouchSliderHandleT) EspErrT

/**
 * @brief   Touch slider subscribes event
 *
 * This function uses event mask to subscribe to touch slider events, once one of
 * the subscribed events occurs, the event message could be retrieved by calling
 * touch_element_message_receive() or input callback routine.
 *
 * @param[in] slider_handle     Slider handle
 * @param[in] event_mask        Slider subscription event mask
 * @param[in] arg               User input argument
 *
 * @note    Touch slider only support three kind of event masks, they are
 *          TOUCH_ELEM_EVENT_ON_PRESS, TOUCH_ELEM_EVENT_ON_RELEASE. You can use those event masks in any
 *          combination to achieve the desired effect.
 *
 * @return
 *      - ESP_OK: Successfully subscribed touch slider event
 *      - ESP_ERR_INVALID_STATE: Touch slider driver was not initialized
 *      - ESP_ERR_INVALID_ARG: slider_handle is null or event is not supported
 */
//go:linkname TouchSliderSubscribeEvent C.touch_slider_subscribe_event
func TouchSliderSubscribeEvent(slider_handle TouchSliderHandleT, event_mask c.Uint32T, arg c.Pointer) EspErrT

/**
 * @brief   Touch slider set dispatch method
 *
 * This function sets a dispatch method that the driver core will use
 * this method as the event notification method.
 *
 * @param[in] slider_handle     Slider handle
 * @param[in] dispatch_method   Dispatch method (By callback/event)
 *
 * @return
 *      - ESP_OK: Successfully set dispatch method
 *      - ESP_ERR_INVALID_STATE: Touch slider driver was not initialized
 *      - ESP_ERR_INVALID_ARG: slider_handle is null or dispatch_method is invalid
 */
//go:linkname TouchSliderSetDispatchMethod C.touch_slider_set_dispatch_method
func TouchSliderSetDispatchMethod(slider_handle TouchSliderHandleT, dispatch_method TouchElemDispatchT) EspErrT

/**
 * @brief   Touch slider set callback
 *
 * This function sets a callback routine into touch element driver core,
 * when the subscribed events occur, the callback routine will be called.
 *
 * @param[in] slider_handle     Slider handle
 * @param[in] slider_callback   User input callback
 *
 * @note        Slider message will be passed from the callback function and it will be destroyed when the callback function return.
 *
 * @warning     Since this input callback routine runs on driver core (esp-timer callback routine),
 *              it should not do something that attempts to Block, such as calling vTaskDelay().
 *
 * @return
 *      - ESP_OK: Successfully set callback
 *      - ESP_ERR_INVALID_STATE: Touch slider driver was not initialized
 *      - ESP_ERR_INVALID_ARG: slider_handle or slider_callback is null
 */
//go:linkname TouchSliderSetCallback C.touch_slider_set_callback
func TouchSliderSetCallback(slider_handle TouchSliderHandleT, slider_callback TouchSliderCallbackT) EspErrT

/**
 * @brief   Touch slider get message
 *
 * This function decodes the element message from touch_element_message_receive() and return
 * a slider message pointer.
 *
 * @param[in] element_message   element message
 *
 * @return  Touch slider message pointer
 */
// llgo:link (*TouchElemMessageT).TouchSliderGetMessage C.touch_slider_get_message
func (recv_ *TouchElemMessageT) TouchSliderGetMessage() *TouchSliderMessageT {
	return nil
}
