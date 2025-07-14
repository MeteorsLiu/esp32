package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_PARTITION_MAGIC = 0x50AA
const ESP_PARTITION_MAGIC_MD5 = 0xEBEB
const PART_TYPE_APP = 0x00
const PART_SUBTYPE_FACTORY = 0x00
const PART_SUBTYPE_OTA_FLAG = 0x10
const PART_SUBTYPE_OTA_MASK = 0x0f
const PART_SUBTYPE_TEST = 0x20
const PART_TYPE_DATA = 0x01
const PART_SUBTYPE_DATA_OTA = 0x00
const PART_SUBTYPE_DATA_RF = 0x01
const PART_SUBTYPE_DATA_WIFI = 0x02
const PART_SUBTYPE_DATA_NVS_KEYS = 0x04
const PART_SUBTYPE_DATA_EFUSE_EM = 0x05
const PART_TYPE_BOOTLOADER = 0x02
const PART_SUBTYPE_BOOTLOADER_PRIMARY = 0x00
const PART_SUBTYPE_BOOTLOADER_OTA = 0x01
const PART_TYPE_PARTITION_TABLE = 0x03
const PART_SUBTYPE_PARTITION_TABLE_PRIMARY = 0x00
const PART_SUBTYPE_PARTITION_TABLE_OTA = 0x01
const PART_TYPE_END = 0xff
const PART_SUBTYPE_END = 0xff
const ESP_PARTITION_MD5_OFFSET = 16
const ESP_BOOTLOADER_DIGEST_OFFSET = 0x0
const ESP_PARTITION_TABLE_MAX_LEN = 0xC00

type EspOtaImgStatesT c.Int

const (
	ESP_OTA_IMG_NEW            EspOtaImgStatesT = 0
	ESP_OTA_IMG_PENDING_VERIFY EspOtaImgStatesT = 1
	ESP_OTA_IMG_VALID          EspOtaImgStatesT = 2
	ESP_OTA_IMG_INVALID        EspOtaImgStatesT = 3
	ESP_OTA_IMG_ABORTED        EspOtaImgStatesT = 4
	ESP_OTA_IMG_UNDEFINED      EspOtaImgStatesT = -1
)

/*
OTA selection structure (two copies in the OTA data partition.)

	Size of 32 bytes is friendly to flash encryption
*/
type EspOtaSelectEntryT struct {
	OtaSeq   c.Uint32T
	SeqLabel [20]c.Uint8T
	OtaState c.Uint32T
	Crc      c.Uint32T
}

type EspPartitionPosT struct {
	Offset c.Uint32T
	Size   c.Uint32T
}

/* Structure which describes the layout of partition table entry.
 * See docs/partition_tables.rst for more information about individual fields.
 */

type EspPartitionInfoT struct {
	Magic   c.Uint16T
	Type    c.Uint8T
	Subtype c.Uint8T
	Pos     EspPartitionPosT
	Label   [16]c.Uint8T
	Flags   c.Uint32T
}

/* @brief Verify the partition table
 *
 * @param partition_table Pointer to at least ESP_PARTITION_TABLE_MAX_ENTRIES of potential partition table data. (ESP_PARTITION_TABLE_MAX_LEN bytes.)
 * @param log_errors Log errors if the partition table is invalid.
 * @param num_partitions If result is ESP_OK, num_partitions is updated with total number of partitions (not including terminating entry).
 *
 * @return ESP_OK on success, ESP_ERR_INVALID_STATE if partition table is not valid.
 */
// llgo:link (*EspPartitionInfoT).EspPartitionTableVerify C.esp_partition_table_verify
func (recv_ *EspPartitionInfoT) EspPartitionTableVerify(log_errors bool, num_partitions *c.Int) EspErrT {
	return 0
}
