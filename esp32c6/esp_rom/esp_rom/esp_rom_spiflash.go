package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspRomSpiflashReadModeT c.Int

const (
	ESP_ROM_SPIFLASH_QIO_MODE         EspRomSpiflashReadModeT = 0
	ESP_ROM_SPIFLASH_QOUT_MODE        EspRomSpiflashReadModeT = 1
	ESP_ROM_SPIFLASH_DIO_MODE         EspRomSpiflashReadModeT = 2
	ESP_ROM_SPIFLASH_DOUT_MODE        EspRomSpiflashReadModeT = 3
	ESP_ROM_SPIFLASH_FASTRD_MODE      EspRomSpiflashReadModeT = 4
	ESP_ROM_SPIFLASH_SLOWRD_MODE      EspRomSpiflashReadModeT = 5
	ESP_ROM_SPIFLASH_OPI_STR_MODE     EspRomSpiflashReadModeT = 6
	ESP_ROM_SPIFLASH_OPI_DTR_MODE     EspRomSpiflashReadModeT = 7
	ESP_ROM_SPIFLASH_OOUT_MODE        EspRomSpiflashReadModeT = 8
	ESP_ROM_SPIFLASH_OIO_STR_MODE     EspRomSpiflashReadModeT = 9
	ESP_ROM_SPIFLASH_OIO_DTR_MODE     EspRomSpiflashReadModeT = 10
	ESP_ROM_SPIFLASH_QPI_MODE         EspRomSpiflashReadModeT = 11
	ESP_ROM_SPIFLASH_OPI_HEX_DTR_MODE EspRomSpiflashReadModeT = 12
)

type EspRomSpiflashChipT struct {
	DeviceId   c.Uint32T
	ChipSize   c.Uint32T
	BlockSize  c.Uint32T
	SectorSize c.Uint32T
	PageSize   c.Uint32T
	StatusMask c.Uint32T
}
type EspRomSpiflashResultT c.Int

const (
	ESP_ROM_SPIFLASH_RESULT_OK      EspRomSpiflashResultT = 0
	ESP_ROM_SPIFLASH_RESULT_ERR     EspRomSpiflashResultT = 1
	ESP_ROM_SPIFLASH_RESULT_TIMEOUT EspRomSpiflashResultT = 2
)

/* Flash data defined in ROM*/

type EspRomSpiflashLegacyDataT struct {
	Chip         EspRomSpiflashChipT
	DummyLenPlus [3]c.Uint8T
	SigMatrix    c.Uint8T
}
