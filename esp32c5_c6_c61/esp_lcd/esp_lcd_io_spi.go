package esp_lcd

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspLcdSpiBusHandleT c.Int

/**
 * @brief Panel IO configuration structure, for SPI interface
 */

type EspLcdPanelIoSpiConfigT struct {
	CsGpioNum        c.Int
	DcGpioNum        c.Int
	SpiMode          c.Int
	PclkHz           c.Uint
	TransQueueDepth  c.SizeT
	OnColorTransDone EspLcdPanelIoColorTransDoneCbT
	UserCtx          c.Pointer
	LcdCmdBits       c.Int
	LcdParamBits     c.Int
	CsEnaPretrans    c.Uint8T
	CsEnaPosttrans   c.Uint8T
	Flags            struct {
		DcHighOnCmd  c.Uint
		DcLowOnData  c.Uint
		DcLowOnParam c.Uint
		OctalMode    c.Uint
		QuadMode     c.Uint
		SioMode      c.Uint
		LsbFirst     c.Uint
		CsHighActive c.Uint
	}
}

/**
 * @brief Create LCD panel IO handle, for SPI interface
 *
 * @param[in] bus SPI bus handle
 * @param[in] io_config IO configuration, for SPI interface
 * @param[out] ret_io Returned IO handle
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_ERR_NO_MEM        if out of memory
 *          - ESP_OK                on success
 */
// llgo:link EspLcdSpiBusHandleT.EspLcdNewPanelIoSpi C.esp_lcd_new_panel_io_spi
func (recv_ EspLcdSpiBusHandleT) EspLcdNewPanelIoSpi(io_config *EspLcdPanelIoSpiConfigT, ret_io *EspLcdPanelIoHandleT) EspErrT {
	return 0
}
