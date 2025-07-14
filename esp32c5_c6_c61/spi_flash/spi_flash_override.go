package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Structure for flash dummy bits.
 *        For some flash chips, dummy bits are configurable under different conditions.
 */

type SpiFlashHpmDummyConfT struct {
	DioDummy    c.Uint8T
	DoutDummy   c.Uint8T
	QioDummy    c.Uint8T
	QoutDummy   c.Uint8T
	FastrdDummy c.Uint8T
}
type SpiFlashRequirementT c.Int

const (
	SPI_FLASH_HPM_CMD_NEEDED      SpiFlashRequirementT = 0
	SPI_FLASH_HPM_DUMMY_NEEDED    SpiFlashRequirementT = 1
	SPI_FLASH_HPM_WRITE_SR_NEEDED SpiFlashRequirementT = 2
	SPI_FLASH_HPM_UNNEEDED        SpiFlashRequirementT = 3
	SPI_FLASH_HPM_BEYOND_LIMIT    SpiFlashRequirementT = 4
)

// llgo:type C
type SpiFlashHpmEnableFnT func()

// llgo:type C
type SpiFlashHpfCheckFnT func() EspErrT

// llgo:type C
type SpiFlashGetChipDummyFnT func(*SpiFlashHpmDummyConfT)

// llgo:type C
type SpiFlashHpmProbeFnT func(c.Uint32T) EspErrT

// llgo:type C
type SpiFlashHpmChipRequirementCheckT func(c.Uint32T, c.Uint32T, c.Int, c.Int) SpiFlashRequirementT

type SpiFlashHpmInfoT struct {
	Method                  *c.Char
	Probe                   SpiFlashHpmProbeFnT
	ChipHpmRequirementCheck SpiFlashHpmChipRequirementCheckT
	FlashHpmEnable          SpiFlashHpmEnableFnT
	FlashHpfCheck           SpiFlashHpfCheckFnT
	FlashGetDummy           SpiFlashGetChipDummyFnT
}
type SpiFlashWrapSizeT c.Int

const (
	FLASH_WRAP_SIZE_8B  SpiFlashWrapSizeT = 8
	FLASH_WRAP_SIZE_16B SpiFlashWrapSizeT = 16
	FLASH_WRAP_SIZE_32B SpiFlashWrapSizeT = 32
	FLASH_WRAP_SIZE_64B SpiFlashWrapSizeT = 64
)

// llgo:type C
type SpiFlashWrapProbeFnT func(c.Uint32T) EspErrT

// llgo:type C
type SpiFlashWrapSetFnT func(SpiFlashWrapSizeT) EspErrT

// llgo:type C
type SpiFlashWrapClrFnT func() EspErrT

type SpiFlashWrapInfoT struct {
	Method      *c.Char
	Probe       SpiFlashWrapProbeFnT
	ChipWrapSet SpiFlashWrapSetFnT
	ChipWrapClr SpiFlashWrapClrFnT
}
