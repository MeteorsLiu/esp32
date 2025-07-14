package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_IMAGE_HEADER_MAGIC = 0xE9
const ESP_IMAGE_MAX_SEGMENTS = 16

type EspChipIdT c.Int

const (
	ESP_CHIP_ID_ESP32   EspChipIdT = 0
	ESP_CHIP_ID_ESP32S2 EspChipIdT = 2
	ESP_CHIP_ID_ESP32C3 EspChipIdT = 5
	ESP_CHIP_ID_ESP32S3 EspChipIdT = 9
	ESP_CHIP_ID_ESP32C2 EspChipIdT = 12
	ESP_CHIP_ID_ESP32C6 EspChipIdT = 13
	ESP_CHIP_ID_ESP32H2 EspChipIdT = 16
	ESP_CHIP_ID_ESP32P4 EspChipIdT = 18
	ESP_CHIP_ID_ESP32C5 EspChipIdT = 23
	ESP_CHIP_ID_INVALID EspChipIdT = 65535
)

type EspImageSpiModeT c.Int

const (
	ESP_IMAGE_SPI_MODE_QIO       EspImageSpiModeT = 0
	ESP_IMAGE_SPI_MODE_QOUT      EspImageSpiModeT = 1
	ESP_IMAGE_SPI_MODE_DIO       EspImageSpiModeT = 2
	ESP_IMAGE_SPI_MODE_DOUT      EspImageSpiModeT = 3
	ESP_IMAGE_SPI_MODE_FAST_READ EspImageSpiModeT = 4
	ESP_IMAGE_SPI_MODE_SLOW_READ EspImageSpiModeT = 5
)

type EspImageSpiFreqT c.Int

const (
	ESP_IMAGE_SPI_SPEED_DIV_2 EspImageSpiFreqT = 0
	ESP_IMAGE_SPI_SPEED_DIV_3 EspImageSpiFreqT = 1
	ESP_IMAGE_SPI_SPEED_DIV_4 EspImageSpiFreqT = 2
	ESP_IMAGE_SPI_SPEED_DIV_1 EspImageSpiFreqT = 15
)

type EspImageFlashSizeT c.Int

const (
	ESP_IMAGE_FLASH_SIZE_1MB   EspImageFlashSizeT = 0
	ESP_IMAGE_FLASH_SIZE_2MB   EspImageFlashSizeT = 1
	ESP_IMAGE_FLASH_SIZE_4MB   EspImageFlashSizeT = 2
	ESP_IMAGE_FLASH_SIZE_8MB   EspImageFlashSizeT = 3
	ESP_IMAGE_FLASH_SIZE_16MB  EspImageFlashSizeT = 4
	ESP_IMAGE_FLASH_SIZE_32MB  EspImageFlashSizeT = 5
	ESP_IMAGE_FLASH_SIZE_64MB  EspImageFlashSizeT = 6
	ESP_IMAGE_FLASH_SIZE_128MB EspImageFlashSizeT = 7
	ESP_IMAGE_FLASH_SIZE_MAX   EspImageFlashSizeT = 8
)

/**
 * @brief Main header of binary image
 */

type EspImageHeaderT struct {
	Magic          c.Uint8T
	SegmentCount   c.Uint8T
	SpiMode        c.Uint8T
	SpiSpeed       c.Uint8T
	SpiSize        c.Uint8T
	EntryAddr      c.Uint32T
	WpPin          c.Uint8T
	SpiPinDrv      [3]c.Uint8T
	ChipId         EspChipIdT
	MinChipRev     c.Uint8T
	MinChipRevFull c.Uint16T
	MaxChipRevFull c.Uint16T
	Reserved       [4]c.Uint8T
	HashAppended   c.Uint8T
}

/**
 * @brief Header of binary image segment
 */

type EspImageSegmentHeaderT struct {
	LoadAddr c.Uint32T
	DataLen  c.Uint32T
}
