package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
Note: Most of esp_flash APIs in ROM are compatible with headers in ESP-IDF, this function

	just adds ROM-specific parts
*/
type SpiFlashChipT struct {
	Unused [8]uint8
}

type EspFlashT struct {
	Unused [8]uint8
}

/* Structure to wrap "global" data used by esp_flash in ROM */

type EspFlashRomGlobalDataT struct {
	DefaultChip *EspFlashT
	ApiFuncs    struct {
		Start     c.Pointer
		End       c.Pointer
		ChipCheck c.Pointer
	}
}
