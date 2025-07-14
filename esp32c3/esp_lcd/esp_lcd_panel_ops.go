package esp_lcd

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Reset LCD panel
 *
 * @note Panel reset must be called before attempting to initialize the panel using `esp_lcd_panel_init()`.
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @return
 *          - ESP_OK on success
 */
//go:linkname EspLcdPanelReset C.esp_lcd_panel_reset
func EspLcdPanelReset(panel EspLcdPanelHandleT) EspErrT

/**
 * @brief Initialize LCD panel
 *
 * @note Before calling this function, make sure the LCD panel has finished the `reset` stage by `esp_lcd_panel_reset()`.
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @return
 *          - ESP_OK on success
 */
//go:linkname EspLcdPanelInit C.esp_lcd_panel_init
func EspLcdPanelInit(panel EspLcdPanelHandleT) EspErrT

/**
 * @brief Deinitialize the LCD panel
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @return
 *          - ESP_OK on success
 */
//go:linkname EspLcdPanelDel C.esp_lcd_panel_del
func EspLcdPanelDel(panel EspLcdPanelHandleT) EspErrT

/**
 * @brief Draw bitmap on LCD panel
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @param[in] x_start Start pixel index in the target frame buffer, on x-axis (x_start is included)
 * @param[in] y_start Start pixel index in the target frame buffer, on y-axis (y_start is included)
 * @param[in] x_end End pixel index in the target frame buffer, on x-axis (x_end is not included)
 * @param[in] y_end End pixel index in the target frame buffer, on y-axis (y_end is not included)
 * @param[in] color_data RGB color data that will be dumped to the specific window range
 * @return
 *          - ESP_OK on success
 */
//go:linkname EspLcdPanelDrawBitmap C.esp_lcd_panel_draw_bitmap
func EspLcdPanelDrawBitmap(panel EspLcdPanelHandleT, x_start c.Int, y_start c.Int, x_end c.Int, y_end c.Int, color_data c.Pointer) EspErrT

/**
 * @brief Mirror the LCD panel on specific axis
 *
 * @note Combined with `esp_lcd_panel_swap_xy()`, one can realize screen rotation
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @param[in] mirror_x Whether the panel will be mirrored about the x axis
 * @param[in] mirror_y Whether the panel will be mirrored about the y axis
 * @return
 *          - ESP_OK on success
 *          - ESP_ERR_NOT_SUPPORTED if this function is not supported by the panel
 */
//go:linkname EspLcdPanelMirror C.esp_lcd_panel_mirror
func EspLcdPanelMirror(panel EspLcdPanelHandleT, mirror_x bool, mirror_y bool) EspErrT

/**
 * @brief Swap/Exchange x and y axis
 *
 * @note Combined with `esp_lcd_panel_mirror()`, one can realize screen rotation
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @param[in] swap_axes Whether to swap the x and y axis
 * @return
 *          - ESP_OK on success
 *          - ESP_ERR_NOT_SUPPORTED if this function is not supported by the panel
 */
//go:linkname EspLcdPanelSwapXy C.esp_lcd_panel_swap_xy
func EspLcdPanelSwapXy(panel EspLcdPanelHandleT, swap_axes bool) EspErrT

/**
 * @brief Set extra gap in x and y axis
 *
 * The gap is the space (in pixels) between the left/top sides of the LCD panel and the first row/column respectively of the actual contents displayed.
 *
 * @note Setting a gap is useful when positioning or centering a frame that is smaller than the LCD.
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @param[in] x_gap Extra gap on x axis, in pixels
 * @param[in] y_gap Extra gap on y axis, in pixels
 * @return
 *          - ESP_OK on success
 */
//go:linkname EspLcdPanelSetGap C.esp_lcd_panel_set_gap
func EspLcdPanelSetGap(panel EspLcdPanelHandleT, x_gap c.Int, y_gap c.Int) EspErrT

/**
 * @brief Invert the color (bit-wise invert the color data line)
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @param[in] invert_color_data Whether to invert the color data
 * @return
 *          - ESP_OK on success
 */
//go:linkname EspLcdPanelInvertColor C.esp_lcd_panel_invert_color
func EspLcdPanelInvertColor(panel EspLcdPanelHandleT, invert_color_data bool) EspErrT

/**
 * @brief Turn on or off the display
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @param[in] on_off True to turns on display, False to turns off display
 * @return
 *          - ESP_OK on success
 *          - ESP_ERR_NOT_SUPPORTED if this function is not supported by the panel
 */
//go:linkname EspLcdPanelDispOnOff C.esp_lcd_panel_disp_on_off
func EspLcdPanelDispOnOff(panel EspLcdPanelHandleT, on_off bool) EspErrT

/**
 * @brief Turn off the display
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @param[in] off Whether to turn off the screen
 * @return
 *          - ESP_OK on success
 *          - ESP_ERR_NOT_SUPPORTED if this function is not supported by the panel
 */
//go:linkname EspLcdPanelDispOff C.esp_lcd_panel_disp_off
func EspLcdPanelDispOff(panel EspLcdPanelHandleT, off bool) EspErrT

/**
 * @brief Enter or exit sleep mode
 *
 * @param[in] panel LCD panel handle, which is created by other factory API like `esp_lcd_new_panel_st7789()`
 * @param[in] sleep True to enter sleep mode, False to wake up
 * @return
 *          - ESP_OK on success
 *          - ESP_ERR_NOT_SUPPORTED if this function is not supported by the panel
 */
//go:linkname EspLcdPanelDispSleep C.esp_lcd_panel_disp_sleep
func EspLcdPanelDispSleep(panel EspLcdPanelHandleT, sleep bool) EspErrT
