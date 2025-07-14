package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Timeout configurations for flash operations, all in us */

type FlashChipOpTimeoutT struct {
	IdleTimeout        c.Uint32T
	ChipEraseTimeout   c.Uint32T
	BlockEraseTimeout  c.Uint32T
	SectorEraseTimeout c.Uint32T
	PageProgramTimeout c.Uint32T
}
type SpiFlashRegisterT c.Int

const SPI_FLASH_REG_STATUS SpiFlashRegisterT = 1

type SpiFlashCapsT c.Int

const (
	SPI_FLASH_CHIP_CAP_SUSPEND      SpiFlashCapsT = 1
	SPI_FLASH_CHIP_CAP_32MB_SUPPORT SpiFlashCapsT = 2
	SPI_FLASH_CHIP_CAP_UNIQUE_ID    SpiFlashCapsT = 4
)
