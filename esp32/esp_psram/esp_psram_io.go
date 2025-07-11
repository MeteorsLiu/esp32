package esp_psram

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief get psram CS IO
 *
 * This interface should be called after PSRAM is enabled, otherwise it will
 * return an invalid value -1/0xff.
 *
 * @return psram CS IO or -1/0xff if psram not enabled
 */
//go:linkname EspPsramIoGetCsIo C.esp_psram_io_get_cs_io
func EspPsramIoGetCsIo() c.Uint8T
