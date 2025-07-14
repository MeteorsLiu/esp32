package esp_driver_cam

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Enable ESP CAM controller
 *
 * @param[in] handle  ESP CAM controller handle
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname EspCamCtlrEnable C.esp_cam_ctlr_enable
func EspCamCtlrEnable(handle EspCamCtlrHandleT) EspErrT

/**
 * @brief Start ESP CAM controller
 *
 * @param[in] handle  ESP CAM controller handle
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname EspCamCtlrStart C.esp_cam_ctlr_start
func EspCamCtlrStart(handle EspCamCtlrHandleT) EspErrT

/**
 * @brief Stop ESP CAM controller
 *
 * @param[in] handle  ESP CAM controller handle
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname EspCamCtlrStop C.esp_cam_ctlr_stop
func EspCamCtlrStop(handle EspCamCtlrHandleT) EspErrT

/**
 * @brief Disable ESP CAM controller
 *
 * @param[in] handle  ESP CAM controller handle
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname EspCamCtlrDisable C.esp_cam_ctlr_disable
func EspCamCtlrDisable(handle EspCamCtlrHandleT) EspErrT

/**
 * @brief Receive data to the given transaction
 *
 * @param[in] handle      ESP CAM controller handle
 * @param[in] trans       ESP CAM controller transaction type
 * @param[in] timeout_ms  Timeout in ms
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname EspCamCtlrReceive C.esp_cam_ctlr_receive
func EspCamCtlrReceive(handle EspCamCtlrHandleT, trans *EspCamCtlrTransT, timeout_ms c.Uint32T) EspErrT

/**
 * @brief Delete ESP CAM controller handle
 *
 * @param[in] handle  ESP CAM controller handle
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname EspCamCtlrDel C.esp_cam_ctlr_del
func EspCamCtlrDel(handle EspCamCtlrHandleT) EspErrT

/**
 * @brief Register ESP CAM controller event callbacks
 *
 * @param[in] handle     ESP CAM controller handle
 * @param[in] cbs        ESP CAM controller event callbacks
 * @param[in] user_data  User data
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname EspCamCtlrRegisterEventCallbacks C.esp_cam_ctlr_register_event_callbacks
func EspCamCtlrRegisterEventCallbacks(handle EspCamCtlrHandleT, cbs *EspCamCtlrEvtCbsT, user_data c.Pointer) EspErrT

/**
 * @brief Get ESP CAM controller internal malloced backup buffer(s) addr
 *
 * @note Generally, data in internal buffer is ready when `on_trans_finished` event
 *
 * @param[in]   handle      ESP CAM controller handle
 * @param[in]   fb_num      Number of frame buffer(s) to get. This value must be the same as the number of the followed fbN parameters
 * @param[out]  fb0         Address of the frame buffer 0 (first frame buffer)
 * @param[out]  ...         List of other frame buffers if any
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid driver state
 */
//go:linkname EspCamCtlrGetFrameBuffer C.esp_cam_ctlr_get_frame_buffer
func EspCamCtlrGetFrameBuffer(handle EspCamCtlrHandleT, fb_num c.Uint32T, fb0 *c.Pointer, __llgo_va_list ...interface{}) EspErrT

/**
 * @brief Get ESP CAM controller internal backup buffer length
 *
 * @param[in]   handle      ESP CAM controller handle
 * @param[out]  ret_fb_len  Optional, The size of each frame buffer, in bytes.
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   NULL ptr
 *        - ESP_ERR_INVALID_STATE: Invalid driver state
 */
//go:linkname EspCamCtlrGetFrameBufferLen C.esp_cam_ctlr_get_frame_buffer_len
func EspCamCtlrGetFrameBufferLen(handle EspCamCtlrHandleT, ret_fb_len *c.SizeT) EspErrT
