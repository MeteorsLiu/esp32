package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * ISSI SPI flash chip_drv, uses all the above functions for its operations. In
 * default autodetection, this is used as a catchall if a more specific chip_drv
 * is not found.
 */
// llgo:link (*EspFlashT).SpiFlashChipIssiProbe C.spi_flash_chip_issi_probe
func (recv_ *EspFlashT) SpiFlashChipIssiProbe(flash_id c.Uint32T) EspErrT {
	return 0
}

// llgo:link (*EspFlashT).SpiFlashChipIssiSetIoMode C.spi_flash_chip_issi_set_io_mode
func (recv_ *EspFlashT) SpiFlashChipIssiSetIoMode() EspErrT {
	return 0
}

// llgo:link (*EspFlashT).SpiFlashChipIssiGetIoMode C.spi_flash_chip_issi_get_io_mode
func (recv_ *EspFlashT) SpiFlashChipIssiGetIoMode(out_io_mode *EspFlashIoModeT) EspErrT {
	return 0
}
