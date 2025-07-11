package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_SECURE_BOOT_DIGEST_LEN = 32
const ESP_SECURE_BOOT_KEY_DIGEST_LEN = 32
const FLASH_OFFS_SECURE_BOOT_IV_DIGEST = 0

type EspSecureBootSigSchemeT c.Int

const (
	ESP_SECURE_BOOT_V1_ECDSA EspSecureBootSigSchemeT = 0
	ESP_SECURE_BOOT_V2_RSA   EspSecureBootSigSchemeT = 2
	ESP_SECURE_BOOT_V2_ECDSA EspSecureBootSigSchemeT = 3
)

/** @brief Secure boot verification block, on-flash data format. */

type EspSecureBootSigBlockT struct {
	Version   c.Uint32T
	Signature [64]c.Uint8T
}

/** @brief Secure boot IV+digest header */

type EspSecureBootIvDigestT struct {
	Iv     [128]c.Uint8T
	Digest [64]c.Uint8T
}

/** @brief Check the secure boot V2 during startup
 *
 * @note This function is called automatically during app startup,
 * it doesn't need to be called from the app.
 *
 * Verifies the secure boot config during startup:
 *
 * - Correct any insecure secure boot settings
 */
//go:linkname EspSecureBootInitChecks C.esp_secure_boot_init_checks
func EspSecureBootInitChecks()

/** @brief Set all secure eFuse features related to secure_boot
 *
 *  @note
 *      This API needs to be called in the eFuse batch mode.
 *      i.e. A call to esp_efuse_batch_write_begin() should be made prior to calling this API to start the batch mode
 *      After the API has been executed a call to esp_efuse_batch_write_commit()/esp_efuse_batch_write_cancel()
 *      should be made accordingly.
 * @return
 *  - ESP_OK - Successfully
 */
//go:linkname EspSecureBootEnableSecureFeatures C.esp_secure_boot_enable_secure_features
func EspSecureBootEnableSecureFeatures() EspErrT

/** @brief Returns the verification status for all physical security features of secure boot in release mode
 *
 * If the device has secure boot feature configured in the release mode,
 * then it is highly recommended to call this API in the application startup code.
 * This API verifies the sanity of the eFuse configuration against
 * the release (production) mode of the secure boot feature.
 *
 * @return
 *  - True - all eFuses are configured correctly
 *  - False - not all eFuses are configured correctly.
 */
//go:linkname EspSecureBootCfgVerifyReleaseMode C.esp_secure_boot_cfg_verify_release_mode
func EspSecureBootCfgVerifyReleaseMode() bool
