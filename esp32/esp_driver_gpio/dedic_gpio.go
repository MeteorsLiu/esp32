package esp_driver_gpio

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DedicGpioBundleT struct {
	Unused [8]uint8
}
type DedicGpioBundleHandleT *DedicGpioBundleT

/**
 * @brief Type of Dedicated GPIO bundle configuration
 */

type DedicGpioBundleConfigT struct {
	GpioArray *c.Int
	ArraySize c.SizeT
	Flags     struct {
		InEn      c.Uint
		InInvert  c.Uint
		OutEn     c.Uint
		OutInvert c.Uint
	}
}
