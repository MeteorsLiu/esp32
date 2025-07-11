package touch_element

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* ------------------------------------------------------------------------------------------------------------------ */
/**
 * @brief   Matrix button initialization configuration passed to touch_matrix_install
 */

type TouchMatrixGlobalConfigT struct {
	ThresholdDivider c.Float
	DefaultLpTime    c.Uint32T
}

/**
 * @brief   Matrix button configuration (for new instance) passed to touch_matrix_create()
 */

type TouchMatrixConfigT struct {
	XChannelArray     *TouchPadT
	YChannelArray     *TouchPadT
	XSensitivityArray *c.Float
	YSensitivityArray *c.Float
	XChannelNum       c.Uint8T
	YChannelNum       c.Uint8T
}
type TouchMatrixEventT c.Int

const (
	TOUCH_MATRIX_EVT_ON_PRESS     TouchMatrixEventT = 0
	TOUCH_MATRIX_EVT_ON_RELEASE   TouchMatrixEventT = 1
	TOUCH_MATRIX_EVT_ON_LONGPRESS TouchMatrixEventT = 2
	TOUCH_MATRIX_EVT_MAX          TouchMatrixEventT = 3
)

/**
 * @brief   Matrix button position data type
 */

type TouchMatrixPositionT struct {
	XAxis c.Uint8T
	YAxis c.Uint8T
	Index c.Uint8T
}

/**
 * @brief   Matrix message type
 */

type TouchMatrixMessageT struct {
	Event    TouchMatrixEventT
	Position TouchMatrixPositionT
}
type TouchMatrixHandleT TouchElemHandleT

// llgo:type C
type TouchMatrixCallbackT func(TouchMatrixHandleT, *TouchMatrixMessageT, c.Pointer)

/**
 * @brief   Touch matrix button initialize
 *
 * This function initializes touch matrix button object and acts on all
 * touch matrix button instances.
 *
 * @param[in] global_config   Touch matrix global initialization configuration
 *
 * @return
 *      - ESP_OK: Successfully initialized touch matrix button
 *      - ESP_ERR_INVALID_STATE: Touch element library was not initialized
 *      - ESP_ERR_INVALID_ARG: matrix_init is NULL
 *      - ESP_ERR_NO_MEM: Insufficient memory
 */
// llgo:link (*TouchMatrixGlobalConfigT).TouchMatrixInstall C.touch_matrix_install
func (recv_ *TouchMatrixGlobalConfigT) TouchMatrixInstall() EspErrT {
	return 0
}

/**
 * @brief   Release resources allocated using touch_matrix_install()
 *
 */
//go:linkname TouchMatrixUninstall C.touch_matrix_uninstall
func TouchMatrixUninstall()

/**
 * @brief   Create a new touch matrix button instance
 *
 * @param[in]  matrix_config    Matrix button configuration
 * @param[out] matrix_handle    Matrix button handle
 *
 * @note    Channel array and sensitivity array must be one-one correspondence in those array
 *
 * @note    Touch matrix button does not support Multi-Touch now
 *
 * @return
 *      - ESP_OK: Successfully create touch matrix button
 *      - ESP_ERR_INVALID_STATE: Touch matrix driver was not initialized
 *      - ESP_ERR_INVALID_ARG: Invalid configuration struct or arguments is NULL
 *      - ESP_ERR_NO_MEM: Insufficient memory
 */
// llgo:link (*TouchMatrixConfigT).TouchMatrixCreate C.touch_matrix_create
func (recv_ *TouchMatrixConfigT) TouchMatrixCreate(matrix_handle *TouchMatrixHandleT) EspErrT {
	return 0
}

/**
 * @brief   Release resources allocated using touch_matrix_create()
 *
 * @param[in] matrix_handle    Matrix handle
 * @return
 *      - ESP_OK: Successfully released resources
 *      - ESP_ERR_INVALID_STATE: Touch matrix driver was not initialized
 *      - ESP_ERR_INVALID_ARG: matrix_handle is null
 *      - ESP_ERR_NOT_FOUND: Input handle is not a matrix handle
 */
//go:linkname TouchMatrixDelete C.touch_matrix_delete
func TouchMatrixDelete(matrix_handle TouchMatrixHandleT) EspErrT

/**
 * @brief   Touch matrix button subscribes event
 *
 * This function uses event mask to subscribe to touch matrix events, once one of
 * the subscribed events occurs, the event message could be retrieved by calling
 * touch_element_message_receive() or input callback routine.
 *
 * @param[in] matrix_handle     Matrix handle
 * @param[in] event_mask        Matrix subscription event mask
 * @param[in] arg               User input argument
 *
 * @note    Touch matrix button only support three kind of event masks, they are
 *          TOUCH_ELEM_EVENT_ON_PRESS, TOUCH_ELEM_EVENT_ON_RELEASE, TOUCH_ELEM_EVENT_ON_LONGPRESS. You can use those event
 *          masks in any combination to achieve the desired effect.
 *
 * @return
 *      - ESP_OK: Successfully subscribed touch matrix event
 *      - ESP_ERR_INVALID_STATE: Touch matrix driver was not initialized
 *      - ESP_ERR_INVALID_ARG: matrix_handle is null or event is not supported
 */
//go:linkname TouchMatrixSubscribeEvent C.touch_matrix_subscribe_event
func TouchMatrixSubscribeEvent(matrix_handle TouchMatrixHandleT, event_mask c.Uint32T, arg c.Pointer) EspErrT

/**
 * @brief   Touch matrix button set dispatch method
 *
 * This function sets a dispatch method that the driver core will use
 * this method as the event notification method.
 *
 * @param[in] matrix_handle     Matrix button handle
 * @param[in] dispatch_method   Dispatch method (By callback/event)
 *
 * @return
 *      - ESP_OK: Successfully set dispatch method
 *      - ESP_ERR_INVALID_STATE: Touch matrix driver was not initialized
 *      - ESP_ERR_INVALID_ARG: matrix_handle is null or dispatch_method is invalid
 */
//go:linkname TouchMatrixSetDispatchMethod C.touch_matrix_set_dispatch_method
func TouchMatrixSetDispatchMethod(matrix_handle TouchMatrixHandleT, dispatch_method TouchElemDispatchT) EspErrT

/**
 * @brief   Touch matrix button set callback
 *
 * This function sets a callback routine into touch element driver core,
 * when the subscribed events occur, the callback routine will be called.
 *
 * @param[in] matrix_handle     Matrix button handle
 * @param[in] matrix_callback   User input callback
 *
 * @note        Matrix message will be passed from the callback function and it will be destroyed when the callback function return.
 *
 * @warning     Since this input callback routine runs on driver core (esp-timer callback routine),
 *              it should not do something that attempts to Block, such as calling vTaskDelay().
 *
 * @return
 *      - ESP_OK: Successfully set callback
 *      - ESP_ERR_INVALID_STATE: Touch matrix driver was not initialized
 *      - ESP_ERR_INVALID_ARG: matrix_handle or matrix_callback is null
 */
//go:linkname TouchMatrixSetCallback C.touch_matrix_set_callback
func TouchMatrixSetCallback(matrix_handle TouchMatrixHandleT, matrix_callback TouchMatrixCallbackT) EspErrT

/**
 * @brief   Touch matrix button set long press trigger time
 *
 * This function sets the threshold time (ms) for a long press event. If a matrix button is pressed
 * and held for a period of time that exceeds the threshold time, a long press event is triggered.
 *
 * @param[in] matrix_handle     Matrix button handle
 * @param[in] threshold_time    Threshold time (ms) of long press event occur
 *
 * @return
 *      - ESP_OK: Successfully set the time of long press event
 *      - ESP_ERR_INVALID_STATE: Touch matrix driver was not initialized
 *      - ESP_ERR_INVALID_ARG: matrix_handle is null or time (ms) is 0
 */
//go:linkname TouchMatrixSetLongpress C.touch_matrix_set_longpress
func TouchMatrixSetLongpress(matrix_handle TouchMatrixHandleT, threshold_time c.Uint32T) EspErrT

/**
 * @brief   Touch matrix get message
 *
 * This function decodes the element message from touch_element_message_receive() and return
 * a matrix message pointer.
 *
 * @param[in] element_message   element message
 *
 * @return  Touch matrix message pointer
 */
// llgo:link (*TouchElemMessageT).TouchMatrixGetMessage C.touch_matrix_get_message
func (recv_ *TouchElemMessageT) TouchMatrixGetMessage() *TouchMatrixMessageT {
	return nil
}
