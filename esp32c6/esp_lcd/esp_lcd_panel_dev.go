package esp_lcd

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Configuration structure for panel device
 */

type EspLcdPanelDevConfigT struct {
	ResetGpioNum c.Int
	DataEndian   LcdRgbDataEndianT
	BitsPerPixel c.Uint32T
	Flags        struct {
		ResetActiveHigh c.Uint32T
	}
	VendorConfig c.Pointer
}
