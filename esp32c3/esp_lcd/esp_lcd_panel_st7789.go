package esp_lcd

import _ "unsafe"

/**
 * @brief Create LCD panel for model ST7789
 *
 * @param[in] io LCD panel IO handle
 * @param[in] panel_dev_config general panel device configuration
 * @param[out] ret_panel Returned LCD panel handle
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_ERR_NO_MEM        if out of memory
 *          - ESP_OK                on success
 */
//go:linkname EspLcdNewPanelSt7789 C.esp_lcd_new_panel_st7789
func EspLcdNewPanelSt7789(io EspLcdPanelIoHandleT, panel_dev_config *EspLcdPanelDevConfigT, ret_panel *EspLcdPanelHandleT) EspErrT
