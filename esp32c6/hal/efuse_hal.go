package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief get factory mac address
 */
//go:linkname EfuseHalGetMac C.efuse_hal_get_mac
func EfuseHalGetMac(mac *c.Uint8T)

/**
 * @brief Returns chip version
 *
 * @return Chip version in format: Major * 100 + Minor
 */
//go:linkname EfuseHalChipRevision C.efuse_hal_chip_revision
func EfuseHalChipRevision() c.Uint32T

/**
 * @brief Return block version
 *
 * @return Block version in format: Major * 100 + Minor
 */
//go:linkname EfuseHalBlkVersion C.efuse_hal_blk_version
func EfuseHalBlkVersion() c.Uint32T

/**
 * @brief Is flash encryption currently enabled in hardware?
 *
 * Flash encryption is enabled if the FLASH_CRYPT_CNT efuse has an odd number of bits set.
 *
 * @return true if flash encryption is enabled.
 */
//go:linkname EfuseHalFlashEncryptionEnabled C.efuse_hal_flash_encryption_enabled
func EfuseHalFlashEncryptionEnabled() bool

/**
 * @brief Returns the status of whether the bootloader (and OTA)
 *        will check the maximum chip version or not.
 *
 * @return true - Skip the maximum chip version check.
 */
//go:linkname EfuseHalGetDisableWaferVersionMajor C.efuse_hal_get_disable_wafer_version_major
func EfuseHalGetDisableWaferVersionMajor() bool

/**
 * @brief Returns the status of whether the app start-up (and OTA)
 *        will check the efuse block version or not.
 *
 * @return true - Skip the efuse block version check.
 */
//go:linkname EfuseHalGetDisableBlkVersionMajor C.efuse_hal_get_disable_blk_version_major
func EfuseHalGetDisableBlkVersionMajor() bool

/**
 * @brief Returns major chip version
 */
//go:linkname EfuseHalGetMajorChipVersion C.efuse_hal_get_major_chip_version
func EfuseHalGetMajorChipVersion() c.Uint32T

/**
 * @brief Returns minor chip version
 */
//go:linkname EfuseHalGetMinorChipVersion C.efuse_hal_get_minor_chip_version
func EfuseHalGetMinorChipVersion() c.Uint32T
