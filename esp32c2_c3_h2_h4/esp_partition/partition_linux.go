package esp_partition

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_PARTITION_EMULATED_SECTOR_SIZE = 0x1000
const ESP_PARTITION_DEFAULT_EMULATED_FLASH_SIZE = 0x400000
const ESP_PARTITION_FAIL_AFTER_MODE_ERASE = 0x01
const ESP_PARTITION_FAIL_AFTER_MODE_WRITE = 0x02
const ESP_PARTITION_FAIL_AFTER_MODE_BOTH = 0x03

type EspPartitionFileMmapCtrlT struct {
	FlashFileName     [1024]c.Char
	FlashFileSize     c.SizeT
	PartitionFileName [1024]c.Char
	RemoveDump        bool
}
