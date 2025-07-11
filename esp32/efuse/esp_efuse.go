package efuse

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_ERR_EFUSE = 0x1600

/**
 * @brief Type definition for an eFuse field
 */

type EspEfuseDescT struct {
	EfuseBlock EspEfuseBlockT
	BitStart   c.Uint8T
	BitCount   c.Uint16T
}
type EspEfuseRomLogSchemeT c.Int

const (
	ESP_EFUSE_ROM_LOG_ALWAYS_ON    EspEfuseRomLogSchemeT = 0
	ESP_EFUSE_ROM_LOG_ON_GPIO_LOW  EspEfuseRomLogSchemeT = 1
	ESP_EFUSE_ROM_LOG_ON_GPIO_HIGH EspEfuseRomLogSchemeT = 2
	ESP_EFUSE_ROM_LOG_ALWAYS_OFF   EspEfuseRomLogSchemeT = 3
)

/**
 * @brief   Reads bits from EFUSE field and writes it into an array.
 *
 * The number of read bits will be limited to the minimum value
 * from the description of the bits in "field" structure or "dst_size_bits" required size.
 * Use "esp_efuse_get_field_size()" function to determine the length of the field.
 *
 * @note Please note that reading in the batch mode does not show uncommitted changes.
 *
 * @param[in]  field          A pointer to the structure describing the fields of efuse.
 * @param[out] dst            A pointer to array that will contain the result of reading.
 * @param[in]  dst_size_bits  The number of bits required to read.
 *                            If the requested number of bits is greater than the field,
 *                            the number will be limited to the field size.
 *
 * @return
 *    - ESP_OK: The operation was successfully completed.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 */
//go:linkname EspEfuseReadFieldBlob C.esp_efuse_read_field_blob
func EspEfuseReadFieldBlob(field **EspEfuseDescT, dst c.Pointer, dst_size_bits c.SizeT) EspErrT

/**
 * @brief Read a single bit eFuse field as a boolean value.
 *
 * @note The value must exist and must be a single bit wide. If there is any possibility of an error
 * in the provided arguments, call esp_efuse_read_field_blob() and check the returned value instead.
 *
 * @note If assertions are enabled and the parameter is invalid, execution will abort
 * @note Please note that reading in the batch mode does not show uncommitted changes.
 *
 * @param[in]  field          A pointer to the structure describing the fields of efuse.
 * @return
 *    - true: The field parameter is valid and the bit is set.
 *    - false: The bit is not set, or the parameter is invalid and assertions are disabled.
 *
 */
//go:linkname EspEfuseReadFieldBit C.esp_efuse_read_field_bit
func EspEfuseReadFieldBit(field **EspEfuseDescT) bool

/**
 * @brief   Reads bits from EFUSE field and returns number of bits programmed as "1".
 *
 * If the bits are set not sequentially, they will still be counted.
 * @note Please note that reading in the batch mode does not show uncommitted changes.
 *
 * @param[in]  field          A pointer to the structure describing the fields of efuse.
 * @param[out] out_cnt        A pointer that will contain the number of programmed as "1" bits.
 *
 * @return
 *    - ESP_OK: The operation was successfully completed.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 */
//go:linkname EspEfuseReadFieldCnt C.esp_efuse_read_field_cnt
func EspEfuseReadFieldCnt(field **EspEfuseDescT, out_cnt *c.SizeT) EspErrT

/**
 * @brief   Writes array to EFUSE field.
 *
 * The number of write bits will be limited to the minimum value
 * from the description of the bits in "field" structure or "src_size_bits" required size.
 * Use "esp_efuse_get_field_size()" function to determine the length of the field.
 * After the function is completed, the writing registers are cleared.
 * @param[in]  field          A pointer to the structure describing the fields of efuse.
 * @param[in]  src            A pointer to array that contains the data for writing.
 * @param[in]  src_size_bits  The number of bits required to write.
 *
 * @return
 *    - ESP_OK: The operation was successfully completed.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 *    - ESP_ERR_EFUSE_REPEATED_PROG: Error repeated programming of programmed bits is strictly forbidden.
 *    - ESP_ERR_CODING: Error range of data does not match the coding scheme.
 */
//go:linkname EspEfuseWriteFieldBlob C.esp_efuse_write_field_blob
func EspEfuseWriteFieldBlob(field **EspEfuseDescT, src c.Pointer, src_size_bits c.SizeT) EspErrT

/**
 * @brief   Writes a required count of bits as "1" to EFUSE field.
 *
 * If there are no free bits in the field to set the required number of bits to "1",
 * ESP_ERR_EFUSE_CNT_IS_FULL error is returned, the field will not be partially recorded.
 * After the function is completed, the writing registers are cleared.
 * @param[in]  field          A pointer to the structure describing the fields of efuse.
 * @param[in]  cnt            Required number of programmed as "1" bits.
 *
 * @return
 *    - ESP_OK: The operation was successfully completed.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 *    - ESP_ERR_EFUSE_CNT_IS_FULL: Not all requested cnt bits is set.
 */
//go:linkname EspEfuseWriteFieldCnt C.esp_efuse_write_field_cnt
func EspEfuseWriteFieldCnt(field **EspEfuseDescT, cnt c.SizeT) EspErrT

/**
 * @brief Write a single bit eFuse field to 1
 *
 * For use with eFuse fields that are a single bit. This function will write the bit to value 1 if
 * it is not already set, or does nothing if the bit is already set.
 *
 * This is equivalent to calling esp_efuse_write_field_cnt() with the cnt parameter equal to 1,
 * except that it will return ESP_OK if the field is already set to 1.
 *
 * @param[in] field Pointer to the structure describing the efuse field.
 *
 * @return
 * - ESP_OK: The operation was successfully completed, or the bit was already set to value 1.
 * - ESP_ERR_INVALID_ARG: Error in the passed arugments, including if the efuse field is not 1 bit wide.
 */
//go:linkname EspEfuseWriteFieldBit C.esp_efuse_write_field_bit
func EspEfuseWriteFieldBit(field **EspEfuseDescT) EspErrT

/**
 * @brief   Sets a write protection for the whole block.
 *
 * After that, it is impossible to write to this block.
 * The write protection does not apply to block 0.
 * @param[in]  blk          Block number of eFuse. (EFUSE_BLK1, EFUSE_BLK2 and EFUSE_BLK3)
 *
 * @return
 *    - ESP_OK: The operation was successfully completed.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 *    - ESP_ERR_EFUSE_CNT_IS_FULL: Not all requested cnt bits is set.
 *    - ESP_ERR_NOT_SUPPORTED: The block does not support this command.
 */
// llgo:link EspEfuseBlockT.EspEfuseSetWriteProtect C.esp_efuse_set_write_protect
func (recv_ EspEfuseBlockT) EspEfuseSetWriteProtect() EspErrT {
	return 0
}

/**
 * @brief   Sets a read protection for the whole block.
 *
 * After that, it is impossible to read from this block.
 * The read protection does not apply to block 0.
 * @param[in]  blk          Block number of eFuse. (EFUSE_BLK1, EFUSE_BLK2 and EFUSE_BLK3)
 *
 * @return
 *    - ESP_OK: The operation was successfully completed.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 *    - ESP_ERR_EFUSE_CNT_IS_FULL: Not all requested cnt bits is set.
 *    - ESP_ERR_NOT_SUPPORTED: The block does not support this command.
 */
// llgo:link EspEfuseBlockT.EspEfuseSetReadProtect C.esp_efuse_set_read_protect
func (recv_ EspEfuseBlockT) EspEfuseSetReadProtect() EspErrT {
	return 0
}

/**
 * @brief   Returns the number of bits used by field.
 *
 * @param[in]  field          A pointer to the structure describing the fields of efuse.
 *
 * @return Returns the number of bits used by field.
 */
//go:linkname EspEfuseGetFieldSize C.esp_efuse_get_field_size
func EspEfuseGetFieldSize(field **EspEfuseDescT) c.Int

/**
 * @brief   Returns value of efuse register.
 *
 * This is a thread-safe implementation.
 * Example: EFUSE_BLK2_RDATA3_REG where (blk=2, num_reg=3)
 * @note Please note that reading in the batch mode does not show uncommitted changes.
 *
 * @param[in]  blk     Block number of eFuse.
 * @param[in]  num_reg The register number in the block.
 *
 * @return Value of register
 */
// llgo:link EspEfuseBlockT.EspEfuseReadReg C.esp_efuse_read_reg
func (recv_ EspEfuseBlockT) EspEfuseReadReg(num_reg c.Uint) c.Uint32T {
	return 0
}

/**
 * @brief   Write value to efuse register.
 *
 * Apply a coding scheme if necessary.
 * This is a thread-safe implementation.
 * Example: EFUSE_BLK3_WDATA0_REG where (blk=3, num_reg=0)
 * @param[in]  blk     Block number of eFuse.
 * @param[in]  num_reg The register number in the block.
 * @param[in]  val     Value to write.
 *
 * @return
 *      - ESP_OK: The operation was successfully completed.
 *      - ESP_ERR_EFUSE_REPEATED_PROG: Error repeated programming of programmed bits is strictly forbidden.
 */
// llgo:link EspEfuseBlockT.EspEfuseWriteReg C.esp_efuse_write_reg
func (recv_ EspEfuseBlockT) EspEfuseWriteReg(num_reg c.Uint, val c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief   Return efuse coding scheme for blocks.
 *
 * @note The coding scheme is applicable only to 1, 2 and 3 blocks. For 0 block, the coding scheme is always ``NONE``.
 *
 * @param[in]  blk     Block number of eFuse.
 * @return Return efuse coding scheme for blocks
 */
// llgo:link EspEfuseBlockT.EspEfuseGetCodingScheme C.esp_efuse_get_coding_scheme
func (recv_ EspEfuseBlockT) EspEfuseGetCodingScheme() EspEfuseCodingSchemeT {
	return 0
}

/**
 * @brief   Read key to efuse block starting at the offset and the required size.
 *
 * @note Please note that reading in the batch mode does not show uncommitted changes.
 *
 * @param[in]  blk             Block number of eFuse.
 * @param[in]  dst_key         A pointer to array that will contain the result of reading.
 * @param[in]  offset_in_bits  Start bit in block.
 * @param[in]  size_bits       The number of bits required to read.
 *
 * @return
 *    - ESP_OK: The operation was successfully completed.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 *    - ESP_ERR_CODING: Error range of data does not match the coding scheme.
 */
// llgo:link EspEfuseBlockT.EspEfuseReadBlock C.esp_efuse_read_block
func (recv_ EspEfuseBlockT) EspEfuseReadBlock(dst_key c.Pointer, offset_in_bits c.SizeT, size_bits c.SizeT) EspErrT {
	return 0
}

/**
 * @brief   Write key to efuse block starting at the offset and the required size.
 *
 * @param[in]  blk             Block number of eFuse.
 * @param[in]  src_key         A pointer to array that contains the key for writing.
 * @param[in]  offset_in_bits  Start bit in block.
 * @param[in]  size_bits       The number of bits required to write.
 *
 * @return
 *    - ESP_OK: The operation was successfully completed.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 *    - ESP_ERR_CODING: Error range of data does not match the coding scheme.
 *    - ESP_ERR_EFUSE_REPEATED_PROG: Error repeated programming of programmed bits
 */
// llgo:link EspEfuseBlockT.EspEfuseWriteBlock C.esp_efuse_write_block
func (recv_ EspEfuseBlockT) EspEfuseWriteBlock(src_key c.Pointer, offset_in_bits c.SizeT, size_bits c.SizeT) EspErrT {
	return 0
}

/**
 * @brief   Returns chip package from efuse
 *
 * @return chip package
 */
//go:linkname EspEfuseGetPkgVer C.esp_efuse_get_pkg_ver
func EspEfuseGetPkgVer() c.Uint32T

/**
 *  @brief Reset efuse write registers
 *
 * Efuse write registers are written to zero, to negate
 * any changes that have been staged here.
 *
 * @note This function is not threadsafe, if calling code updates
 * efuse values from multiple tasks then this is caller's
 * responsibility to serialise.
 */
//go:linkname EspEfuseReset C.esp_efuse_reset
func EspEfuseReset()

/**
 *  @brief Disable BASIC ROM Console via efuse
 *
 * By default, if booting from flash fails the ESP32 will boot a
 * BASIC console in ROM.
 *
 * Call this function (from bootloader or app) to permanently disable the console on this chip.
 *
 */
//go:linkname EspEfuseDisableBasicRomConsole C.esp_efuse_disable_basic_rom_console
func EspEfuseDisableBasicRomConsole()

/**
 *  @brief Disable ROM Download Mode via eFuse
 *
 * Permanently disables the ROM Download Mode feature. Once disabled, if the SoC is booted with
 * strapping pins set for ROM Download Mode then an error is printed instead.
 *
 * @note Not all SoCs support this option. An error will be returned if called on an ESP32
 * with a silicon revision lower than 3, as these revisions do not support this option.
 *
 * @note If ROM Download Mode is already disabled, this function does nothing and returns success.
 *
 * @return
 * - ESP_OK If the eFuse was successfully burned, or had already been burned.
 * - ESP_ERR_NOT_SUPPORTED (ESP32 only) This SoC is not capable of disabling UART download mode
 * - ESP_ERR_INVALID_STATE (ESP32 only) This eFuse is write protected and cannot be written
 */
//go:linkname EspEfuseDisableRomDownloadMode C.esp_efuse_disable_rom_download_mode
func EspEfuseDisableRomDownloadMode() EspErrT

/**
 * @brief Set boot ROM log scheme via eFuse
 *
 * @note By default, the boot ROM will always print to console. This API can be called to set the log scheme only once per chip,
 *       once the value is changed from the default it can't be changed again.
 *
 * @param log_scheme Supported ROM log scheme
 * @return
 *      - ESP_OK If the eFuse was successfully burned, or had already been burned.
 *      - ESP_ERR_NOT_SUPPORTED (ESP32 only) This SoC is not capable of setting ROM log scheme
 *      - ESP_ERR_INVALID_STATE This eFuse is write protected or has been burned already
 */
// llgo:link EspEfuseRomLogSchemeT.EspEfuseSetRomLogScheme C.esp_efuse_set_rom_log_scheme
func (recv_ EspEfuseRomLogSchemeT) EspEfuseSetRomLogScheme() EspErrT {
	return 0
}

/**
 *  @brief Return secure_version from efuse field.
 * @return Secure version from efuse field
 */
//go:linkname EspEfuseReadSecureVersion C.esp_efuse_read_secure_version
func EspEfuseReadSecureVersion() c.Uint32T

/**
 *  @brief Check secure_version from app and secure_version and from efuse field.
 *
 * @param secure_version Secure version from app.
 * @return
 *          - True: If version of app is equal or more then secure_version from efuse.
 */
//go:linkname EspEfuseCheckSecureVersion C.esp_efuse_check_secure_version
func EspEfuseCheckSecureVersion(secure_version c.Uint32T) bool

/**
 *  @brief Write efuse field by secure_version value.
 *
 * Update the secure_version value is available if the coding scheme is None.
 * Note: Do not use this function in your applications. This function is called as part of the other API.
 *
 * @param[in] secure_version Secure version from app.
 * @return
 *          - ESP_OK: Successful.
 *          - ESP_FAIL: secure version of app cannot be set to efuse field.
 *          - ESP_ERR_NOT_SUPPORTED: Anti rollback is not supported with the 3/4 and Repeat coding scheme.
 */
//go:linkname EspEfuseUpdateSecureVersion C.esp_efuse_update_secure_version
func EspEfuseUpdateSecureVersion(secure_version c.Uint32T) EspErrT

/**
 *  @brief Set the batch mode of writing fields.
 *
 * This mode allows you to write the fields in the batch mode when need to burn several efuses at one time.
 * To enable batch mode call begin() then perform as usually the necessary operations
 * read and write and at the end call commit() to actually burn all written efuses.
 * The batch mode can be used nested. The commit will be done by the last commit() function.
 * The number of begin() functions should be equal to the number of commit() functions.
 *
 * @note Please note that reading in the batch mode does not show uncommitted changes.
 *
 * Note: If batch mode is enabled by the first task, at this time the second task cannot write/read efuses.
 * The second task will wait for the first task to complete the batch operation.
 *
 * \code{c}
 * // Example of using the batch writing mode.
 *
 * // set the batch writing mode
 * esp_efuse_batch_write_begin();
 *
 * // use any writing functions as usual
 * esp_efuse_write_field_blob(ESP_EFUSE_...);
 * esp_efuse_write_field_cnt(ESP_EFUSE_...);
 * esp_efuse_set_write_protect(EFUSE_BLKx);
 * esp_efuse_write_reg(EFUSE_BLKx, ...);
 * esp_efuse_write_block(EFUSE_BLKx, ...);
 * esp_efuse_write(ESP_EFUSE_1, 3);  // ESP_EFUSE_1 == 1, here we write a new value = 3. The changes will be burn by the commit() function.
 * esp_efuse_read_...(ESP_EFUSE_1);  // this function returns ESP_EFUSE_1 == 1 because uncommitted changes are not readable, it will be available only after commit.
 * ...
 *
 * // esp_efuse_batch_write APIs can be called recursively.
 * esp_efuse_batch_write_begin();
 * esp_efuse_set_write_protect(EFUSE_BLKx);
 * esp_efuse_batch_write_commit(); // the burn will be skipped here, it will be done in the last commit().
 *
 * ...
 *
 * // Write all of these fields to the efuse registers
 * esp_efuse_batch_write_commit();
 * esp_efuse_read_...(ESP_EFUSE_1);  // this function returns ESP_EFUSE_1 == 3.
 *
 * \endcode
 *
 * @return
 *          - ESP_OK: Successful.
 */
//go:linkname EspEfuseBatchWriteBegin C.esp_efuse_batch_write_begin
func EspEfuseBatchWriteBegin() EspErrT

/**
 *  @brief Reset the batch mode of writing fields.
 *
 * It will reset the batch writing mode and any written changes.
 *
 * @return
 *          - ESP_OK: Successful.
 *          - ESP_ERR_INVALID_STATE: Tha batch mode was not set.
 */
//go:linkname EspEfuseBatchWriteCancel C.esp_efuse_batch_write_cancel
func EspEfuseBatchWriteCancel() EspErrT

/**
 *  @brief Writes all prepared data for the batch mode.
 *
 * Must be called to ensure changes are written to the efuse registers.
 * After this the batch writing mode will be reset.
 *
 * @return
 *          - ESP_OK: Successful.
 *          - ESP_ERR_INVALID_STATE: The deferred writing mode was not set.
 */
//go:linkname EspEfuseBatchWriteCommit C.esp_efuse_batch_write_commit
func EspEfuseBatchWriteCommit() EspErrT

/**
 *  @brief Checks that the given block is empty.
 *
 * @return
 *          - True: The block is empty.
 *          - False: The block is not empty or was an error.
 */
// llgo:link EspEfuseBlockT.EspEfuseBlockIsEmpty C.esp_efuse_block_is_empty
func (recv_ EspEfuseBlockT) EspEfuseBlockIsEmpty() bool {
	return false
}

/**
 * @brief Returns a read protection for the key block.
 *
 * @param[in] block A key block in the range EFUSE_BLK_KEY0..EFUSE_BLK_KEY_MAX
 *
 * @return True: The key block is read protected
 *         False: The key block is readable.
 */
// llgo:link EspEfuseBlockT.EspEfuseGetKeyDisRead C.esp_efuse_get_key_dis_read
func (recv_ EspEfuseBlockT) EspEfuseGetKeyDisRead() bool {
	return false
}

/**
 * @brief Sets a read protection for the key block.
 *
 * @param[in] block A key block in the range EFUSE_BLK_KEY0..EFUSE_BLK_KEY_MAX
 *
 * @return
 *    - ESP_OK: Successful.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 *    - ESP_ERR_EFUSE_REPEATED_PROG: Error repeated programming of programmed bits is strictly forbidden.
 *    - ESP_ERR_CODING: Error range of data does not match the coding scheme.
 */
// llgo:link EspEfuseBlockT.EspEfuseSetKeyDisRead C.esp_efuse_set_key_dis_read
func (recv_ EspEfuseBlockT) EspEfuseSetKeyDisRead() EspErrT {
	return 0
}

/**
 * @brief Returns a write protection for the key block.
 *
 * @param[in] block A key block in the range EFUSE_BLK_KEY0..EFUSE_BLK_KEY_MAX
 *
 * @return True: The key block is write protected
 *         False: The key block is writeable.
 */
// llgo:link EspEfuseBlockT.EspEfuseGetKeyDisWrite C.esp_efuse_get_key_dis_write
func (recv_ EspEfuseBlockT) EspEfuseGetKeyDisWrite() bool {
	return false
}

/**
 * @brief Sets a write protection for the key block.
 *
 * @param[in] block A key block in the range EFUSE_BLK_KEY0..EFUSE_BLK_KEY_MAX
 *
 * @return
 *    - ESP_OK: Successful.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 *    - ESP_ERR_EFUSE_REPEATED_PROG: Error repeated programming of programmed bits is strictly forbidden.
 *    - ESP_ERR_CODING: Error range of data does not match the coding scheme.
 */
// llgo:link EspEfuseBlockT.EspEfuseSetKeyDisWrite C.esp_efuse_set_key_dis_write
func (recv_ EspEfuseBlockT) EspEfuseSetKeyDisWrite() EspErrT {
	return 0
}

/**
 * @brief Returns true if the key block is unused, false otherwise.
 *
 * An unused key block is all zero content, not read or write protected,
 * and has purpose 0 (ESP_EFUSE_KEY_PURPOSE_USER)
 *
 * @param block key block to check.
 *
 * @return
 *         - True if key block is unused,
 *         - False if key block is used or the specified block index is not a key block.
 */
// llgo:link EspEfuseBlockT.EspEfuseKeyBlockUnused C.esp_efuse_key_block_unused
func (recv_ EspEfuseBlockT) EspEfuseKeyBlockUnused() bool {
	return false
}

/**
 * @brief Find a key block with the particular purpose set.
 *
 * @param[in] purpose Purpose to search for.
 * @param[out] block Pointer in the range EFUSE_BLK_KEY0..EFUSE_BLK_KEY_MAX which will be set to the key block if found.
 *                   Can be NULL, if only need to test the key block exists.
 *
 * @return
 *         - True: If found,
 *         - False: If not found (value at block pointer is unchanged).
 */
// llgo:link EspEfusePurposeT.EspEfuseFindPurpose C.esp_efuse_find_purpose
func (recv_ EspEfusePurposeT) EspEfuseFindPurpose(block *EspEfuseBlockT) bool {
	return false
}

/**
 * @brief Returns a write protection of the key purpose field for an efuse key block.
 *
 * @param[in] block A key block in the range EFUSE_BLK_KEY0..EFUSE_BLK_KEY_MAX
 *
 * @note For ESP32: no keypurpose, it returns always True.
 *
 * @return True: The key purpose is write protected.
 *         False: The key purpose is writeable.
 */
// llgo:link EspEfuseBlockT.EspEfuseGetKeypurposeDisWrite C.esp_efuse_get_keypurpose_dis_write
func (recv_ EspEfuseBlockT) EspEfuseGetKeypurposeDisWrite() bool {
	return false
}

/**
 * @brief Returns the current purpose set for an efuse key block.
 *
 * @param[in] block A key block in the range EFUSE_BLK_KEY0..EFUSE_BLK_KEY_MAX
 *
 * @return
 *         - Value: If Successful, it returns the value of the purpose related to the given key block.
 *         - ESP_EFUSE_KEY_PURPOSE_MAX: Otherwise.
 */
// llgo:link EspEfuseBlockT.EspEfuseGetKeyPurpose C.esp_efuse_get_key_purpose
func (recv_ EspEfuseBlockT) EspEfuseGetKeyPurpose() EspEfusePurposeT {
	return 0
}

/**
 * @brief Program a block of key data to an efuse block
 *
 * The burn of a key, protection bits, and a purpose happens in batch mode.
 *
 * @note This API also enables the read protection efuse bit for certain key blocks like XTS-AES, HMAC, ECDSA etc.
 * This ensures that the key is only accessible to hardware peripheral.
 *
 * @note For SoC's with capability `SOC_EFUSE_ECDSA_USE_HARDWARE_K` (e.g., ESP32-H2), this API writes an additional
 * efuse bit for ECDSA key purpose to enforce hardware TRNG generated k mode in the peripheral.
 *
 * @param[in] block Block to read purpose for. Must be in range EFUSE_BLK_KEY0 to EFUSE_BLK_KEY_MAX. Key block must be unused (esp_efuse_key_block_unused).
 * @param[in] purpose Purpose to set for this key. Purpose must be already unset.
 * @param[in] key Pointer to data to write.
 * @param[in] key_size_bytes Bytes length of data to write.
 *
 * @return
 *    - ESP_OK: Successful.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 *    - ESP_ERR_INVALID_STATE: Error in efuses state, unused block not found.
 *    - ESP_ERR_EFUSE_REPEATED_PROG: Error repeated programming of programmed bits is strictly forbidden.
 *    - ESP_ERR_CODING: Error range of data does not match the coding scheme.
 */
// llgo:link EspEfuseBlockT.EspEfuseWriteKey C.esp_efuse_write_key
func (recv_ EspEfuseBlockT) EspEfuseWriteKey(purpose EspEfusePurposeT, key c.Pointer, key_size_bytes c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Program keys to unused efuse blocks
 *
 * The burn of keys, protection bits, and purposes happens in batch mode.
 *
 * @note This API also enables the read protection efuse bit for certain key blocks like XTS-AES, HMAC, ECDSA etc.
 * This ensures that the key is only accessible to hardware peripheral.
 *
 * @note For SoC's with capability `SOC_EFUSE_ECDSA_USE_HARDWARE_K` (e.g., ESP32-H2), this API writes an additional
 * efuse bit for ECDSA key purpose to enforce hardware TRNG generated k mode in the peripheral.
 *
 * @param[in] purposes Array of purposes (purpose[number_of_keys]).
 * @param[in] keys Array of keys (uint8_t keys[number_of_keys][32]). Each key is 32 bytes long.
 * @param[in] number_of_keys The number of keys to write (up to 6 keys).
 *
 * @return
 *    - ESP_OK: Successful.
 *    - ESP_ERR_INVALID_ARG: Error in the passed arguments.
 *    - ESP_ERR_INVALID_STATE: Error in efuses state, unused block not found.
 *    - ESP_ERR_NOT_ENOUGH_UNUSED_KEY_BLOCKS: Error not enough unused key blocks available
 *    - ESP_ERR_EFUSE_REPEATED_PROG: Error repeated programming of programmed bits is strictly forbidden.
 *    - ESP_ERR_CODING: Error range of data does not match the coding scheme.
 */
//go:linkname EspEfuseWriteKeys C.esp_efuse_write_keys
func EspEfuseWriteKeys(purposes *EspEfusePurposeT, keys **c.Uint8T, number_of_keys c.Uint) EspErrT

/**
 * @brief   Checks eFuse errors in BLOCK0.
 *
 * @note Refers to ESP32-C3 only.
 *
 * It does a BLOCK0 check if eFuse EFUSE_ERR_RST_ENABLE is set.
 * If BLOCK0 has an error, it prints the error and returns ESP_FAIL, which should be treated as esp_restart.
 *
 * @return
 *         - ESP_OK: No errors in BLOCK0.
 *         - ESP_FAIL: Error in BLOCK0 requiring reboot.
 */
//go:linkname EspEfuseCheckErrors C.esp_efuse_check_errors
func EspEfuseCheckErrors() EspErrT

/**
 * @brief   Destroys the data in the given efuse block, if possible.
 *
 * Data destruction occurs through the following steps:
 * 1) Destroy data in the block:
 *    - If write protection is inactive for the block, then unset bits are burned.
 *    - If write protection is active, the block remains unaltered.
 * 2) Set read protection for the block if possible (check write-protection for RD_DIS).
 *    In this case, data becomes inaccessible, and the software reads it as all zeros.
 * If write protection is enabled and read protection can not be set,
 * data in the block remains readable (returns an error).
 *
 * Do not use the batch mode with this function as it does the burning itself!
 *
 * @param[in] block A key block in the range EFUSE_BLK_KEY0..EFUSE_BLK_KEY_MAX
 *
 * @return
 *    - ESP_OK: Successful.
 *    - ESP_FAIL: Data remained readable because the block is write-protected
 *                and read protection can not be set.
 */
// llgo:link EspEfuseBlockT.EspEfuseDestroyBlock C.esp_efuse_destroy_block
func (recv_ EspEfuseBlockT) EspEfuseDestroyBlock() EspErrT {
	return 0
}
