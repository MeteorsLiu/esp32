package efuse

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspEfuseBlockT c.Int

const (
	EFUSE_BLK0               EspEfuseBlockT = 0
	EFUSE_BLK1               EspEfuseBlockT = 1
	EFUSE_BLK2               EspEfuseBlockT = 2
	EFUSE_BLK_SYS_DATA_PART1 EspEfuseBlockT = 2
	EFUSE_BLK3               EspEfuseBlockT = 3
	EFUSE_BLK_USER_DATA      EspEfuseBlockT = 3
	EFUSE_BLK4               EspEfuseBlockT = 4
	EFUSE_BLK_KEY0           EspEfuseBlockT = 4
	EFUSE_BLK5               EspEfuseBlockT = 5
	EFUSE_BLK_KEY1           EspEfuseBlockT = 5
	EFUSE_BLK6               EspEfuseBlockT = 6
	EFUSE_BLK_KEY2           EspEfuseBlockT = 6
	EFUSE_BLK7               EspEfuseBlockT = 7
	EFUSE_BLK_KEY3           EspEfuseBlockT = 7
	EFUSE_BLK8               EspEfuseBlockT = 8
	EFUSE_BLK_KEY4           EspEfuseBlockT = 8
	EFUSE_BLK9               EspEfuseBlockT = 9
	EFUSE_BLK_KEY5           EspEfuseBlockT = 9
	EFUSE_BLK_KEY_MAX        EspEfuseBlockT = 10
	EFUSE_BLK10              EspEfuseBlockT = 10
	EFUSE_BLK_SYS_DATA_PART2 EspEfuseBlockT = 10
	EFUSE_BLK_MAX            EspEfuseBlockT = 11
)

type EspEfuseCodingSchemeT c.Int

const (
	EFUSE_CODING_SCHEME_NONE EspEfuseCodingSchemeT = 0
	EFUSE_CODING_SCHEME_RS   EspEfuseCodingSchemeT = 3
)

type EspEfusePurposeT c.Int

const (
	ESP_EFUSE_KEY_PURPOSE_USER                        EspEfusePurposeT = 0
	ESP_EFUSE_KEY_PURPOSE_RESERVED                    EspEfusePurposeT = 1
	ESP_EFUSE_KEY_PURPOSE_XTS_AES_256_KEY_1           EspEfusePurposeT = 2
	ESP_EFUSE_KEY_PURPOSE_XTS_AES_256_KEY_2           EspEfusePurposeT = 3
	ESP_EFUSE_KEY_PURPOSE_XTS_AES_128_KEY             EspEfusePurposeT = 4
	ESP_EFUSE_KEY_PURPOSE_HMAC_DOWN_ALL               EspEfusePurposeT = 5
	ESP_EFUSE_KEY_PURPOSE_HMAC_DOWN_JTAG              EspEfusePurposeT = 6
	ESP_EFUSE_KEY_PURPOSE_HMAC_DOWN_DIGITAL_SIGNATURE EspEfusePurposeT = 7
	ESP_EFUSE_KEY_PURPOSE_HMAC_UP                     EspEfusePurposeT = 8
	ESP_EFUSE_KEY_PURPOSE_SECURE_BOOT_DIGEST0         EspEfusePurposeT = 9
	ESP_EFUSE_KEY_PURPOSE_SECURE_BOOT_DIGEST1         EspEfusePurposeT = 10
	ESP_EFUSE_KEY_PURPOSE_SECURE_BOOT_DIGEST2         EspEfusePurposeT = 11
	ESP_EFUSE_KEY_PURPOSE_MAX                         EspEfusePurposeT = 12
)
