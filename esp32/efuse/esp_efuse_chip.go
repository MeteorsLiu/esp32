package efuse

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspEfuseBlockT c.Int

const (
	EFUSE_BLK0              EspEfuseBlockT = 0
	EFUSE_BLK1              EspEfuseBlockT = 1
	EFUSE_BLK_KEY0          EspEfuseBlockT = 1
	EFUSE_BLK_ENCRYPT_FLASH EspEfuseBlockT = 1
	EFUSE_BLK2              EspEfuseBlockT = 2
	EFUSE_BLK_KEY1          EspEfuseBlockT = 2
	EFUSE_BLK_SECURE_BOOT   EspEfuseBlockT = 2
	EFUSE_BLK3              EspEfuseBlockT = 3
	EFUSE_BLK_KEY2          EspEfuseBlockT = 3
	EFUSE_BLK_KEY_MAX       EspEfuseBlockT = 4
	EFUSE_BLK_MAX           EspEfuseBlockT = 4
)

type EspEfuseCodingSchemeT c.Int

const (
	EFUSE_CODING_SCHEME_NONE   EspEfuseCodingSchemeT = 0
	EFUSE_CODING_SCHEME_3_4    EspEfuseCodingSchemeT = 1
	EFUSE_CODING_SCHEME_REPEAT EspEfuseCodingSchemeT = 2
)

type EspEfusePurposeT c.Int

const (
	ESP_EFUSE_KEY_PURPOSE_USER             EspEfusePurposeT = 0
	ESP_EFUSE_KEY_PURPOSE_SYSTEM           EspEfusePurposeT = 1
	ESP_EFUSE_KEY_PURPOSE_FLASH_ENCRYPTION EspEfusePurposeT = 2
	ESP_EFUSE_KEY_PURPOSE_SECURE_BOOT_V2   EspEfusePurposeT = 3
	ESP_EFUSE_KEY_PURPOSE_MAX              EspEfusePurposeT = 4
)
