package esp_bootloader_format

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Bootloader description structure
 */

type EspBootloaderDescT struct {
	MagicByte c.Uint8T
	Reserved  [3]c.Uint8T
	Version   c.Uint32T
	IdfVer    [32]c.Char
	DateTime  [24]c.Char
	Reserved2 [16]c.Uint8T
}

/**
 * @brief   Return esp_bootloader_desc structure.
 *
 * Intended for use by the bootloader.
 * @return Pointer to esp_bootloader_desc structure.
 */
//go:linkname EspBootloaderGetDescription C.esp_bootloader_get_description
func EspBootloaderGetDescription() *EspBootloaderDescT
