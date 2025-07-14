package esp_lcd

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Timing parameters for the video data transmission
 */

type EspLcdVideoTimingT struct {
	HSize           c.Uint32T
	VSize           c.Uint32T
	HsyncPulseWidth c.Uint32T
	HsyncBackPorch  c.Uint32T
	HsyncFrontPorch c.Uint32T
	VsyncPulseWidth c.Uint32T
	VsyncBackPorch  c.Uint32T
	VsyncFrontPorch c.Uint32T
}

type EspLcdPanelIoT struct {
	Unused [8]uint8
}
type EspLcdPanelIoHandleT *EspLcdPanelIoT
type EspLcdPanelHandleT *EspLcdPanelT
type LcdRgbElementOrderT c.Int

const (
	LCD_RGB_ELEMENT_ORDER_RGB LcdRgbElementOrderT = 0
	LCD_RGB_ELEMENT_ORDER_BGR LcdRgbElementOrderT = 1
)

type LcdColorRgbEndianT LcdRgbElementOrderT
type EspLcdColorSpaceT LcdRgbElementOrderT

/**
 * @brief Type of LCD panel IO event data
 */

type EspLcdPanelIoEventDataT struct {
	Unused [8]uint8
}

// llgo:type C
type EspLcdPanelIoColorTransDoneCbT func(EspLcdPanelIoHandleT, *EspLcdPanelIoEventDataT, c.Pointer) bool

/**
 * @brief Type of LCD panel IO callbacks
 */

type EspLcdPanelIoCallbacksT struct {
	OnColorTransDone EspLcdPanelIoColorTransDoneCbT
}

/**
 * @brief Configuration of LCD color conversion
 */

type EspLcdColorConvConfigT struct {
	InColorRange  LcdColorRangeT
	OutColorRange LcdColorRangeT
	Spec          struct {
		Yuv struct {
			ConvStd LcdYuvConvStdT
			Yuv422  struct {
				InPackOrder LcdYuv422PackOrderT
			}
		}
	}
}
