package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SPI_FLASH_OPI_FLAG = 16

/** Definition of a common transaction. Also holds the return value. */

type SpiFlashTransT struct {
	Reserved      c.Uint8T
	MosiLen       c.Uint8T
	MisoLen       c.Uint8T
	AddressBitlen c.Uint8T
	Address       c.Uint32T
	MosiData      *c.Uint8T
	MisoData      *c.Uint8T
	Flags         c.Uint32T
	Command       c.Uint16T
	DummyBitlen   c.Uint8T
	IoMode        c.Uint32T
}
type EspFlashSpeedS c.Int

const (
	ESP_FLASH_5MHZ      EspFlashSpeedS = 5
	ESP_FLASH_10MHZ     EspFlashSpeedS = 10
	ESP_FLASH_20MHZ     EspFlashSpeedS = 20
	ESP_FLASH_26MHZ     EspFlashSpeedS = 26
	ESP_FLASH_40MHZ     EspFlashSpeedS = 40
	ESP_FLASH_80MHZ     EspFlashSpeedS = 80
	ESP_FLASH_120MHZ    EspFlashSpeedS = 120
	ESP_FLASH_SPEED_MAX EspFlashSpeedS = 121
)

type EspFlashSpeedT EspFlashSpeedS
type EspFlashIoModeT c.Int

const (
	SPI_FLASH_SLOWRD        EspFlashIoModeT = 0
	SPI_FLASH_FASTRD        EspFlashIoModeT = 1
	SPI_FLASH_DOUT          EspFlashIoModeT = 2
	SPI_FLASH_DIO           EspFlashIoModeT = 3
	SPI_FLASH_QOUT          EspFlashIoModeT = 4
	SPI_FLASH_QIO           EspFlashIoModeT = 5
	SPI_FLASH_OPI_STR       EspFlashIoModeT = 16
	SPI_FLASH_OPI_DTR       EspFlashIoModeT = 17
	SPI_FLASH_READ_MODE_MAX EspFlashIoModeT = 18
)

// / Configuration structure for the flash chip suspend feature.
type SpiFlashSusCmdConf struct {
	SusMask c.Uint32T
}

// / Structure for flash encryption operations.
type SpiFlashEncryptionT struct {
	FlashEncryptionEnable      c.Pointer
	FlashEncryptionDisable     c.Pointer
	FlashEncryptionDataPrepare c.Pointer
	FlashEncryptionDone        c.Pointer
	FlashEncryptionDestroy     c.Pointer
	FlashEncryptionCheck       c.Pointer
}

type SpiFlashHostDriverS struct {
	Unused [8]uint8
}
type SpiFlashHostDriverT SpiFlashHostDriverS

/** SPI Flash Host driver instance */

type SpiFlashHostInstT struct {
	Driver *SpiFlashHostDriverS
}
