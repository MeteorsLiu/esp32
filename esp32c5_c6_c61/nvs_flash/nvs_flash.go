package nvs_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const NVS_KEY_SIZE = 32

/**
 * @brief Key for encryption and decryption
 */

type NvsSecCfgT struct {
	Eky [32]c.Uint8T
	Tky [32]c.Uint8T
}

// llgo:type C
type NvsFlashGenerateKeysT func(c.Pointer, *NvsSecCfgT) EspErrT

// llgo:type C
type NvsFlashReadCfgT func(c.Pointer, *NvsSecCfgT) EspErrT

/**
 * @brief NVS encryption: Security scheme configuration structure
 */

type NvsSecSchemeT struct {
	SchemeId        c.Int
	SchemeData      c.Pointer
	NvsFlashKeyGen  NvsFlashGenerateKeysT
	NvsFlashReadCfg NvsFlashReadCfgT
}

/**
 * @brief Initialize the default NVS partition.
 *
 * This API initialises the default NVS partition. The default NVS partition
 * is the one that is labeled "nvs" in the partition table.
 *
 * When "NVS_ENCRYPTION" is enabled in the menuconfig, this API enables
 * the NVS encryption for the default NVS partition as follows
 *      1. Read security configurations from the first NVS key
 *         partition listed in the partition table. (NVS key partition is
 *         any "data" type partition which has the subtype value set to "nvs_keys")
 *      2. If the NVS key partition obtained in the previous step is empty,
 *         generate and store new keys in that NVS key partition.
 *      3. Internally call "nvs_flash_secure_init()" with
 *         the security configurations obtained/generated in the previous steps.
 *
 * Post initialization NVS read/write APIs
 * remain the same irrespective of NVS encryption.
 *
 * @return
 *      - ESP_OK if storage was successfully initialized.
 *      - ESP_ERR_NVS_NO_FREE_PAGES if the NVS storage contains no empty pages
 *        (which may happen if NVS partition was truncated)
 *      - ESP_ERR_NOT_FOUND if no partition with label "nvs" is found in the partition table
 *      - ESP_ERR_NO_MEM in case memory could not be allocated for the internal structures
 *      - one of the error codes from the underlying flash storage driver
 *      - error codes from nvs_flash_read_security_cfg API (when "NVS_ENCRYPTION" is enabled).
 *      - error codes from nvs_flash_generate_keys API (when "NVS_ENCRYPTION" is enabled).
 *      - error codes from nvs_flash_secure_init_partition API (when "NVS_ENCRYPTION" is enabled) .
 */
//go:linkname NvsFlashInit C.nvs_flash_init
func NvsFlashInit() EspErrT

/**
 * @brief Initialize NVS flash storage for the specified partition.
 *
 * @param[in]  partition_label   Label of the partition. Must be no longer than 16 characters.
 *
 * @return
 *      - ESP_OK if storage was successfully initialized.
 *      - ESP_ERR_NVS_NO_FREE_PAGES if the NVS storage contains no empty pages
 *        (which may happen if NVS partition was truncated)
 *      - ESP_ERR_NOT_FOUND if specified partition is not found in the partition table
 *      - ESP_ERR_NO_MEM in case memory could not be allocated for the internal structures
 *      - one of the error codes from the underlying flash storage driver
 */
//go:linkname NvsFlashInitPartition C.nvs_flash_init_partition
func NvsFlashInitPartition(partition_label *c.Char) EspErrT

/**
 * @brief Initialize NVS flash storage for the partition specified by partition pointer.
 *
 * @param[in] partition pointer to a partition obtained by the ESP partition API.
 *
 * @return
 *      - ESP_OK if storage was successfully initialized
 *      - ESP_ERR_NVS_NO_FREE_PAGES if the NVS storage contains no empty pages
 *        (which may happen if NVS partition was truncated)
 *      - ESP_ERR_INVALID_ARG in case partition is NULL
 *      - ESP_ERR_NO_MEM in case memory could not be allocated for the internal structures
 *      - one of the error codes from the underlying flash storage driver
 */
//go:linkname NvsFlashInitPartitionPtr C.nvs_flash_init_partition_ptr
func NvsFlashInitPartitionPtr(partition *EspPartitionT) EspErrT

/**
 * @brief Deinitialize NVS storage for the default NVS partition
 *
 * Default NVS partition is the partition with "nvs" label in the partition table.
 *
 * @return
 *      - ESP_OK on success (storage was deinitialized)
 *      - ESP_ERR_NVS_NOT_INITIALIZED if the storage was not initialized prior to this call
 */
//go:linkname NvsFlashDeinit C.nvs_flash_deinit
func NvsFlashDeinit() EspErrT

/**
 * @brief Deinitialize NVS storage for the given NVS partition
 *
 * @param[in]  partition_label   Label of the partition
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NVS_NOT_INITIALIZED if the storage for given partition was not
 *        initialized prior to this call
 */
//go:linkname NvsFlashDeinitPartition C.nvs_flash_deinit_partition
func NvsFlashDeinitPartition(partition_label *c.Char) EspErrT

/**
 * @brief Erase the default NVS partition
 *
 * Erases all contents of the default NVS partition (one with label "nvs").
 *
 * @note If the partition is initialized, this function first de-initializes it. Afterwards, the partition has to
 *       be initialized again to be used.
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NOT_FOUND if there is no NVS partition labeled "nvs" in the
 *        partition table
 *      - different error in case de-initialization fails (shouldn't happen)
 */
//go:linkname NvsFlashErase C.nvs_flash_erase
func NvsFlashErase() EspErrT

/**
 * @brief Erase specified NVS partition
 *
 * Erase all content of a specified NVS partition
 *
 * @note If the partition is initialized, this function first de-initializes it. Afterwards, the partition has to
 *       be initialized again to be used.
 *
 * @param[in]  part_name    Name (label) of the partition which should be erased
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NOT_FOUND if there is no NVS partition with the specified name
 *        in the partition table
 *      - different error in case de-initialization fails (shouldn't happen)
 */
//go:linkname NvsFlashErasePartition C.nvs_flash_erase_partition
func NvsFlashErasePartition(part_name *c.Char) EspErrT

/**
 * @brief Erase custom partition.
 *
 * Erase all content of specified custom partition.
 *
 * @note
 *  If the partition is initialized, this function first de-initializes it.
 *  Afterwards, the partition has to be initialized again to be used.
 *
 * @param[in] partition pointer to a partition obtained by the ESP partition API.
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NOT_FOUND if there is no partition with the specified
 *        parameters in the partition table
 *      - ESP_ERR_INVALID_ARG in case partition is NULL
 *      - one of the error codes from the underlying flash storage driver
 */
//go:linkname NvsFlashErasePartitionPtr C.nvs_flash_erase_partition_ptr
func NvsFlashErasePartitionPtr(partition *EspPartitionT) EspErrT

/**
 * @brief Initialize the default NVS partition.
 *
 * This API initialises the default NVS partition. The default NVS partition
 * is the one that is labeled "nvs" in the partition table.
 *
 * @param[in]  cfg Security configuration (keys) to be used for NVS encryption/decryption.
 *                              If cfg is NULL, no encryption is used.
 *
 * @return
 *      - ESP_OK if storage has been initialized successfully.
 *      - ESP_ERR_NVS_NO_FREE_PAGES if the NVS storage contains no empty pages
 *        (which may happen if NVS partition was truncated)
 *      - ESP_ERR_NOT_FOUND if no partition with label "nvs" is found in the partition table
 *      - ESP_ERR_NO_MEM in case memory could not be allocated for the internal structures
 *      - one of the error codes from the underlying flash storage driver
 */
// llgo:link (*NvsSecCfgT).NvsFlashSecureInit C.nvs_flash_secure_init
func (recv_ *NvsSecCfgT) NvsFlashSecureInit() EspErrT {
	return 0
}

/**
 * @brief Initialize NVS flash storage for the specified partition.
 *
 * @param[in]  partition_label   Label of the partition. Note that internally, a reference to
 *                               passed value is kept and it should be accessible for future operations
 *
 * @param[in]  cfg Security configuration (keys) to be used for NVS encryption/decryption.
 *                              If cfg is null, no encryption/decryption is used.
 * @return
 *      - ESP_OK if storage has been initialized successfully.
 *      - ESP_ERR_NVS_NO_FREE_PAGES if the NVS storage contains no empty pages
 *        (which may happen if NVS partition was truncated)
 *      - ESP_ERR_NOT_FOUND if specified partition is not found in the partition table
 *      - ESP_ERR_NO_MEM in case memory could not be allocated for the internal structures
 *      - one of the error codes from the underlying flash storage driver
 */
//go:linkname NvsFlashSecureInitPartition C.nvs_flash_secure_init_partition
func NvsFlashSecureInitPartition(partition_label *c.Char, cfg *NvsSecCfgT) EspErrT

/**
 * @brief Generate and store NVS keys in the provided esp partition
 *
 * @param[in]  partition Pointer to partition structure obtained using
 *                       esp_partition_find_first or esp_partition_get.
 *                       Must be non-NULL.
 * @param[out] cfg       Pointer to nvs security configuration structure.
 *                       Pointer must be non-NULL.
 *                       Generated keys will be populated in this structure.
 *
 *
 * @return
 *      - ESP_OK, if cfg was read successfully;
 *      - ESP_ERR_INVALID_ARG, if partition or cfg is NULL;
 *      - or error codes from esp_partition_write/erase APIs.
 */
//go:linkname NvsFlashGenerateKeys C.nvs_flash_generate_keys
func NvsFlashGenerateKeys(partition *EspPartitionT, cfg *NvsSecCfgT) EspErrT

/**
 * @brief Read NVS security configuration from a partition.
 *
 * @param[in]  partition Pointer to partition structure obtained using
 *                       esp_partition_find_first or esp_partition_get.
 *                       Must be non-NULL.
 * @param[out] cfg       Pointer to nvs security configuration structure.
 *                       Pointer must be non-NULL.
 *
 * @note  Provided partition is assumed to be marked 'encrypted'.
 *
 * @return
 *      - ESP_OK, if cfg was read successfully;
 *      - ESP_ERR_INVALID_ARG, if partition or cfg is NULL
 *      - ESP_ERR_NVS_KEYS_NOT_INITIALIZED, if the partition is not yet written with keys.
 *      - ESP_ERR_NVS_CORRUPT_KEY_PART, if the partition containing keys is found to be corrupt
 *      - or error codes from esp_partition_read API.
 */
//go:linkname NvsFlashReadSecurityCfg C.nvs_flash_read_security_cfg
func NvsFlashReadSecurityCfg(partition *EspPartitionT, cfg *NvsSecCfgT) EspErrT

/**
 * @brief Registers the given security scheme for NVS encryption
 *        The scheme registered with sec_scheme_id by this API be used as
 *        the default security scheme for the "nvs" partition.
 *        Users will have to call this API explicitly in their application.
 *
 * @param[in] scheme_cfg  Pointer to the security scheme configuration structure
 *                        that the user (or the nvs_key_provider) wants to register.
 *
 * @return
 *      - ESP_OK, if security scheme registration succeeds;
 *      - ESP_ERR_INVALID_ARG, if scheme_cfg is NULL;
 *      - ESP_FAIL, if security scheme registration fails
 */
// llgo:link (*NvsSecSchemeT).NvsFlashRegisterSecurityScheme C.nvs_flash_register_security_scheme
func (recv_ *NvsSecSchemeT) NvsFlashRegisterSecurityScheme() EspErrT {
	return 0
}

/**
 * @brief   Fetch the configuration structure for the default active
 *          security scheme for NVS encryption
 *
 * @return  Pointer to the default active security scheme configuration
 *          (NULL if no scheme is registered yet i.e. active)
 */
//go:linkname NvsFlashGetDefaultSecurityScheme C.nvs_flash_get_default_security_scheme
func NvsFlashGetDefaultSecurityScheme() *NvsSecSchemeT

/**
 * @brief Generate (and store) the NVS keys using the specified key-protection scheme
 *
 * @param[in] scheme_cfg   Security scheme specific configuration
 *
 * @param[out] cfg         Security configuration (encryption keys)
 *
 * @return
 *      - ESP_OK, if cfg was populated successfully with generated encryption keys;
 *      - ESP_ERR_INVALID_ARG, if scheme_cfg or cfg is NULL;
 *      - ESP_FAIL, if the key generation process fails
 */
// llgo:link (*NvsSecSchemeT).NvsFlashGenerateKeysV2 C.nvs_flash_generate_keys_v2
func (recv_ *NvsSecSchemeT) NvsFlashGenerateKeysV2(cfg *NvsSecCfgT) EspErrT {
	return 0
}

/**
 * @brief Read NVS security configuration set by the specified security scheme
 *
 * @param[in] scheme_cfg   Security scheme specific configuration
 *
 * @param[out] cfg         Security configuration (encryption keys)
 *
 * @return
 *      - ESP_OK, if cfg was read successfully;
 *      - ESP_ERR_INVALID_ARG, if scheme_cfg or cfg is NULL;
 *      - ESP_FAIL, if the key reading process fails
 */
// llgo:link (*NvsSecSchemeT).NvsFlashReadSecurityCfgV2 C.nvs_flash_read_security_cfg_v2
func (recv_ *NvsSecSchemeT) NvsFlashReadSecurityCfgV2(cfg *NvsSecCfgT) EspErrT {
	return 0
}
