package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * GD (GigaDevice) SPI flash chip_drv, uses all the above functions for its operations. In
 * default autodetection, this is used as a catchall if a more specific chip_drv
 * is not found.
 *
 * Note that this is for GD chips with product ID 40H (GD25Q) and 60H (GD25LQ). The chip diver uses
 * different commands to write the SR2 register according to the chip ID. For GD25Q40 - GD25Q16
 * chips, and GD25LQ chips, WRSR (01H) command is used; while WRSR2 (31H) is used for GD25Q32 -
 * GD25Q127 chips.
 */
// llgo:link (*EspFlashT).SpiFlashChipGdProbe C.spi_flash_chip_gd_probe
func (recv_ *EspFlashT) SpiFlashChipGdProbe(flash_id c.Uint32T) EspErrT {
	return 0
}

// llgo:link (*EspFlashT).SpiFlashChipGdSetIoMode C.spi_flash_chip_gd_set_io_mode
func (recv_ *EspFlashT) SpiFlashChipGdSetIoMode() EspErrT {
	return 0
}

// llgo:link (*EspFlashT).SpiFlashChipGdGetIoMode C.spi_flash_chip_gd_get_io_mode
func (recv_ *EspFlashT) SpiFlashChipGdGetIoMode(out_io_mode *EspFlashIoModeT) EspErrT {
	return 0
}
