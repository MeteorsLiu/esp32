package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SPI0_R_QIO_DUMMY_CYCLELEN = 5
const SPI0_R_QIO_ADDR_BITSLEN = 23
const SPI0_R_FAST_DUMMY_CYCLELEN = 7
const SPI0_R_DIO_DUMMY_CYCLELEN = 3
const SPI0_R_FAST_ADDR_BITSLEN = 23
const SPI0_R_SIO_ADDR_BITSLEN = 23
const SPI1_R_QIO_DUMMY_CYCLELEN = 5
const SPI1_R_QIO_ADDR_BITSLEN = 23
const SPI1_R_FAST_DUMMY_CYCLELEN = 7
const SPI1_R_DIO_DUMMY_CYCLELEN = 3
const SPI1_R_DIO_ADDR_BITSLEN = 23
const SPI1_R_FAST_ADDR_BITSLEN = 23
const SPI1_R_SIO_ADDR_BITSLEN = 23
const ESP_ROM_SPIFLASH_W_SIO_ADDR_BITSLEN = 23
const ESP_ROM_SPIFLASH_BYTES_LEN = 24
const ESP_ROM_SPIFLASH_BUFF_BYTE_WRITE_NUM = 32
const ESP_ROM_SPIFLASH_BUFF_BYTE_READ_NUM = 16
const ESP_ROM_SPIFLASH_BUFF_BYTE_READ_BITS = 0xf

// llgo:type C
type SpiFlashFuncT func()

// llgo:type C
type SpiFlashOpT func() EspRomSpiflashResultT

// llgo:type C
type SpiFlashEraseT func(c.Uint32T) EspRomSpiflashResultT

// llgo:type C
type SpiFlashRdT func(c.Uint32T, *c.Uint32T, c.Int) EspRomSpiflashResultT

// llgo:type C
type SpiFlashWrT func(c.Uint32T, *c.Uint32T, c.Int) EspRomSpiflashResultT

// llgo:type C
type SpiFlashEwrT func(c.Uint32T, c.Pointer, c.Uint32T) EspRomSpiflashResultT

// llgo:type C
type SpiFlashWrenT func(c.Pointer) EspRomSpiflashResultT

type SpiflashLegacyFuncsT struct {
	ReadSubLen   c.Uint32T
	WriteSubLen  c.Uint32T
	Unlock       SpiFlashOpT
	EraseSector  SpiFlashEraseT
	EraseBlock   SpiFlashEraseT
	Read         SpiFlashRdT
	Write        SpiFlashWrT
	EncryptWrite SpiFlashEwrT
	CheckSus     SpiFlashFuncT
	Wren         SpiFlashWrenT
	WaitIdle     SpiFlashOpT
}

type EspRomSpiflashCommonCmdT struct {
	DataLength c.Uint8T
	ReadCmd0   c.Uint8T
	ReadCmd1   c.Uint8T
	WriteCmd   c.Uint8T
	DataMask   c.Uint16T
	Data       c.Uint16T
}
