package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspFlashEncModeT c.Int

const (
	ESP_FLASH_ENC_MODE_DISABLED    EspFlashEncModeT = 0
	ESP_FLASH_ENC_MODE_DEVELOPMENT EspFlashEncModeT = 1
	ESP_FLASH_ENC_MODE_RELEASE     EspFlashEncModeT = 2
)

/** @brief Is flash encryption currently enabled in hardware?
 *
 * Flash encryption is enabled if the FLASH_CRYPT_CNT efuse has an odd number of bits set.
 *
 * @return true if flash encryption is enabled.
 */
//go:linkname EspFlashEncryptionEnabled C.esp_flash_encryption_enabled
func EspFlashEncryptionEnabled() bool

/** @brief Write protect FLASH_CRYPT_CNT
 *
 * Intended to be called as a part of boot process if flash encryption
 * is enabled but secure boot is not used. This should protect against
 * serial re-flashing of an unauthorised code in absence of secure boot.
 *
 * @note On ESP32 V3 only, write protecting FLASH_CRYPT_CNT will also prevent
 * disabling UART Download Mode. If both are wanted, call
 * esp_efuse_disable_rom_download_mode() before calling this function.
 *
 */
//go:linkname EspFlashWriteProtectCryptCnt C.esp_flash_write_protect_crypt_cnt
func EspFlashWriteProtectCryptCnt()

/** @brief Return the flash encryption mode
 *
 * The API is called during boot process but can also be called by
 * application to check the current flash encryption mode of ESP32
 *
 * @return
 */
//go:linkname EspGetFlashEncryptionMode C.esp_get_flash_encryption_mode
func EspGetFlashEncryptionMode() EspFlashEncModeT

/** @brief Check the flash encryption mode during startup
 *
 * @note This function is called automatically during app startup,
 * it doesn't need to be called from the app.
 *
 * Verifies the flash encryption config during startup:
 *
 * - Correct any insecure flash encryption settings if hardware
 *   Secure Boot is enabled.
 * - Log warnings if the efuse config doesn't match the project
 *  config in any way
 */
//go:linkname EspFlashEncryptionInitChecks C.esp_flash_encryption_init_checks
func EspFlashEncryptionInitChecks()

/** @brief Returns the verification status for all physical security features of flash encryption in release mode
 *
 * If the device has flash encryption feature configured in the release mode,
 * then it is highly recommended to call this API in the application startup code.
 * This API verifies the sanity of the eFuse configuration against
 * the release (production) mode of the flash encryption feature.
 *
 * @return
 *  - True - all eFuses are configured correctly
 *  - False - not all eFuses are configured correctly.
 */
//go:linkname EspFlashEncryptionCfgVerifyReleaseMode C.esp_flash_encryption_cfg_verify_release_mode
func EspFlashEncryptionCfgVerifyReleaseMode() bool

/** @brief Switches Flash Encryption from "Development" to "Release"
 *
 * If already in "Release" mode, the function will do nothing.
 * If flash encryption efuse is not enabled yet then abort.
 * It burns:
 *  - "disable encrypt in dl mode"
 *  - set FLASH_CRYPT_CNT efuse to max
 */
//go:linkname EspFlashEncryptionSetReleaseMode C.esp_flash_encryption_set_release_mode
func EspFlashEncryptionSetReleaseMode()
