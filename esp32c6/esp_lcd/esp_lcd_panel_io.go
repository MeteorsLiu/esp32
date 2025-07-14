package esp_lcd

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Transmit LCD command and receive corresponding parameters
 *
 * @note Commands sent by this function are short, so they are sent using polling transactions.
 *       The function does not return before the command transfer is completed.
 *       If any queued transactions sent by `esp_lcd_panel_io_tx_color()` are still pending when this function is called,
 *       this function will wait until they are finished and the queue is empty before sending the command(s).
 *
 * @param[in]  io LCD panel IO handle, which is created by other factory API like `esp_lcd_new_panel_io_spi()`
 * @param[in]  lcd_cmd The specific LCD command, set to -1 if no command needed
 * @param[out] param Buffer for the command data
 * @param[in]  param_size Size of `param` buffer
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_ERR_NOT_SUPPORTED if read is not supported by transport
 *          - ESP_OK                on success
 */
//go:linkname EspLcdPanelIoRxParam C.esp_lcd_panel_io_rx_param
func EspLcdPanelIoRxParam(io EspLcdPanelIoHandleT, lcd_cmd c.Int, param c.Pointer, param_size c.SizeT) EspErrT

/**
 * @brief Transmit LCD command and corresponding parameters
 *
 * @note Commands sent by this function are short, so they are sent using polling transactions.
 *       The function does not return before the command transfer is completed.
 *       If any queued transactions sent by `esp_lcd_panel_io_tx_color()` are still pending when this function is called,
 *       this function will wait until they are finished and the queue is empty before sending the command(s).
 *
 * @param[in] io LCD panel IO handle, which is created by other factory API like `esp_lcd_new_panel_io_spi()`
 * @param[in] lcd_cmd The specific LCD command, set to -1 if no command needed
 * @param[in] param Buffer that holds the command specific parameters, set to NULL if no parameter is needed for the command
 * @param[in] param_size Size of `param` in memory, in bytes, set to zero if no parameter is needed for the command
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_OK                on success
 */
//go:linkname EspLcdPanelIoTxParam C.esp_lcd_panel_io_tx_param
func EspLcdPanelIoTxParam(io EspLcdPanelIoHandleT, lcd_cmd c.Int, param c.Pointer, param_size c.SizeT) EspErrT

/**
 * @brief Transmit LCD RGB data
 *
 * @note This function will package the command and RGB data into a transaction, and push into a queue.
 *       The real transmission is performed in the background (DMA+interrupt).
 *       The caller should take care of the lifecycle of the `color` buffer.
 *       Recycling of color buffer should be done in the callback `on_color_trans_done()`.
 *
 * @param[in] io LCD panel IO handle, which is created by factory API like `esp_lcd_new_panel_io_spi()`
 * @param[in] lcd_cmd The specific LCD command, set to -1 if no command needed
 * @param[in] color Buffer that holds the RGB color data
 * @param[in] color_size Size of `color` in memory, in bytes
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_OK                on success
 */
//go:linkname EspLcdPanelIoTxColor C.esp_lcd_panel_io_tx_color
func EspLcdPanelIoTxColor(io EspLcdPanelIoHandleT, lcd_cmd c.Int, color c.Pointer, color_size c.SizeT) EspErrT

/**
 * @brief Destroy LCD panel IO handle (deinitialize panel and free all corresponding resource)
 *
 * @param[in] io LCD panel IO handle, which is created by factory API like `esp_lcd_new_panel_io_spi()`
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_OK                on success
 */
//go:linkname EspLcdPanelIoDel C.esp_lcd_panel_io_del
func EspLcdPanelIoDel(io EspLcdPanelIoHandleT) EspErrT

/**
 * @brief Register LCD panel IO callbacks
 *
 * @param[in] io LCD panel IO handle, which is created by factory API like `esp_lcd_new_panel_io_spi()`
 * @param[in] cbs structure with all LCD panel IO callbacks
 * @param[in] user_ctx User private data, passed directly to callback's user_ctx
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_OK                on success
 */
//go:linkname EspLcdPanelIoRegisterEventCallbacks C.esp_lcd_panel_io_register_event_callbacks
func EspLcdPanelIoRegisterEventCallbacks(io EspLcdPanelIoHandleT, cbs *EspLcdPanelIoCallbacksT, user_ctx c.Pointer) EspErrT
