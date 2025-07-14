package esp_lcd

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief SSD1306 configuration structure
 *
 * To be used as esp_lcd_panel_dev_config_t.vendor_config.
 * See esp_lcd_new_panel_ssd1306().
 */

type EspLcdPanelSsd1306ConfigT struct {
	Height c.Uint8T
}

/**
 * @brief Create LCD panel for model SSD1306
 *
 * @param[in] io LCD panel IO handle
 * @param[in] panel_dev_config general panel device configuration
 * @param[out] ret_panel Returned LCD panel handle
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_ERR_NO_MEM        if out of memory
 *          - ESP_OK                on success
 *
 * @note The default panel size is 128x64.
 * @note Use esp_lcd_panel_ssd1306_config_t to set the correct size.
 * Example usage:
 * @code {c}
 *
 * esp_lcd_panel_ssd1306_config_t ssd1306_config = {
 *     .height = 32
 * };
 * esp_lcd_panel_dev_config_t panel_config = {
 *     <...>
 *     .vendor_config = &ssd1306_config
 * };
 *
 * esp_lcd_panel_handle_t panel_handle = NULL;
 * esp_lcd_new_panel_ssd1306(io_handle, &panel_config, &panel_handle);
 * @endcode
 */
//go:linkname EspLcdNewPanelSsd1306 C.esp_lcd_new_panel_ssd1306
func EspLcdNewPanelSsd1306(io EspLcdPanelIoHandleT, panel_dev_config *EspLcdPanelDevConfigT, ret_panel *EspLcdPanelHandleT) EspErrT
