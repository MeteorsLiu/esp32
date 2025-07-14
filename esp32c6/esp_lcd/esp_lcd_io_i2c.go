package esp_lcd

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspLcdI2cBusHandleT c.Uint32T

/**
 * @brief Panel IO configuration structure, for I2C interface
 *
 */

type EspLcdPanelIoI2cConfigT struct {
	DevAddr           c.Uint32T
	OnColorTransDone  EspLcdPanelIoColorTransDoneCbT
	UserCtx           c.Pointer
	ControlPhaseBytes c.SizeT
	DcBitOffset       c.Uint
	LcdCmdBits        c.Int
	LcdParamBits      c.Int
	Flags             struct {
		DcLowOnData         c.Uint
		DisableControlPhase c.Uint
	}
	SclSpeedHz c.Uint32T
}

/**
 * @brief Create LCD panel IO handle, for I2C interface in legacy implementation
 *
 * @param[in] bus I2C bus handle, (in uint32_t)
 * @param[in] io_config IO configuration, for I2C interface
 * @param[out] ret_io Returned IO handle
 *
 * @note Please don't call this function in your project directly. Please call `esp_lcd_new_panel_to_i2c` instead.
 *
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_ERR_NO_MEM        if out of memory
 *          - ESP_OK                on success
 */
//go:linkname EspLcdNewPanelIoI2cV1 C.esp_lcd_new_panel_io_i2c_v1
func EspLcdNewPanelIoI2cV1(bus c.Uint32T, io_config *EspLcdPanelIoI2cConfigT, ret_io *EspLcdPanelIoHandleT) EspErrT

/**
 * @brief Create LCD panel IO handle, for I2C interface in new implementation
 *
 * @param[in] bus I2C bus handle, (in i2c_master_dev_handle_t)
 * @param[in] io_config IO configuration, for I2C interface
 * @param[out] ret_io Returned IO handle
 *
 * @note Please don't call this function in your project directly. Please call `esp_lcd_new_panel_to_i2c` instead.
 *
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_ERR_NO_MEM        if out of memory
 *          - ESP_OK                on success
 */
//go:linkname EspLcdNewPanelIoI2cV2 C.esp_lcd_new_panel_io_i2c_v2
func EspLcdNewPanelIoI2cV2(bus I2cMasterBusHandleT, io_config *EspLcdPanelIoI2cConfigT, ret_io *EspLcdPanelIoHandleT) EspErrT
