package esp_lcd

import _ "unsafe"

/**
 * @brief Create LCD panel for model NT35510
 *
 * @param[in] io LCD panel IO handle
 * @param[in] panel_dev_config general panel device configuration
 * @param[out] ret_panel Returned LCD panel handle
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_ERR_NO_MEM        if out of memory
 *          - ESP_OK                on success
 */
//go:linkname EspLcdNewPanelNt35510 C.esp_lcd_new_panel_nt35510
func EspLcdNewPanelNt35510(io EspLcdPanelIoHandleT, panel_dev_config *EspLcdPanelDevConfigT, ret_panel *EspLcdPanelHandleT) EspErrT
