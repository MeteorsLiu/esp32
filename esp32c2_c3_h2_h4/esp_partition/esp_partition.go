package esp_partition

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspFlashT struct {
	Unused [8]uint8
}
type EspPartitionMmapMemoryT c.Int

const (
	ESP_PARTITION_MMAP_DATA EspPartitionMmapMemoryT = 0
	ESP_PARTITION_MMAP_INST EspPartitionMmapMemoryT = 1
)

type EspPartitionMmapHandleT c.Uint32T
type EspPartitionTypeT c.Int

const (
	ESP_PARTITION_TYPE_APP             EspPartitionTypeT = 0
	ESP_PARTITION_TYPE_DATA            EspPartitionTypeT = 1
	ESP_PARTITION_TYPE_BOOTLOADER      EspPartitionTypeT = 2
	ESP_PARTITION_TYPE_PARTITION_TABLE EspPartitionTypeT = 3
	ESP_PARTITION_TYPE_ANY             EspPartitionTypeT = 255
)

type EspPartitionSubtypeT c.Int

const (
	ESP_PARTITION_SUBTYPE_BOOTLOADER_PRIMARY      EspPartitionSubtypeT = 0
	ESP_PARTITION_SUBTYPE_BOOTLOADER_OTA          EspPartitionSubtypeT = 1
	ESP_PARTITION_SUBTYPE_PARTITION_TABLE_PRIMARY EspPartitionSubtypeT = 0
	ESP_PARTITION_SUBTYPE_PARTITION_TABLE_OTA     EspPartitionSubtypeT = 1
	ESP_PARTITION_SUBTYPE_APP_FACTORY             EspPartitionSubtypeT = 0
	ESP_PARTITION_SUBTYPE_APP_OTA_MIN             EspPartitionSubtypeT = 16
	ESP_PARTITION_SUBTYPE_APP_OTA_0               EspPartitionSubtypeT = 16
	ESP_PARTITION_SUBTYPE_APP_OTA_1               EspPartitionSubtypeT = 17
	ESP_PARTITION_SUBTYPE_APP_OTA_2               EspPartitionSubtypeT = 18
	ESP_PARTITION_SUBTYPE_APP_OTA_3               EspPartitionSubtypeT = 19
	ESP_PARTITION_SUBTYPE_APP_OTA_4               EspPartitionSubtypeT = 20
	ESP_PARTITION_SUBTYPE_APP_OTA_5               EspPartitionSubtypeT = 21
	ESP_PARTITION_SUBTYPE_APP_OTA_6               EspPartitionSubtypeT = 22
	ESP_PARTITION_SUBTYPE_APP_OTA_7               EspPartitionSubtypeT = 23
	ESP_PARTITION_SUBTYPE_APP_OTA_8               EspPartitionSubtypeT = 24
	ESP_PARTITION_SUBTYPE_APP_OTA_9               EspPartitionSubtypeT = 25
	ESP_PARTITION_SUBTYPE_APP_OTA_10              EspPartitionSubtypeT = 26
	ESP_PARTITION_SUBTYPE_APP_OTA_11              EspPartitionSubtypeT = 27
	ESP_PARTITION_SUBTYPE_APP_OTA_12              EspPartitionSubtypeT = 28
	ESP_PARTITION_SUBTYPE_APP_OTA_13              EspPartitionSubtypeT = 29
	ESP_PARTITION_SUBTYPE_APP_OTA_14              EspPartitionSubtypeT = 30
	ESP_PARTITION_SUBTYPE_APP_OTA_15              EspPartitionSubtypeT = 31
	ESP_PARTITION_SUBTYPE_APP_OTA_MAX             EspPartitionSubtypeT = 32
	ESP_PARTITION_SUBTYPE_APP_TEST                EspPartitionSubtypeT = 32
	ESP_PARTITION_SUBTYPE_DATA_OTA                EspPartitionSubtypeT = 0
	ESP_PARTITION_SUBTYPE_DATA_PHY                EspPartitionSubtypeT = 1
	ESP_PARTITION_SUBTYPE_DATA_NVS                EspPartitionSubtypeT = 2
	ESP_PARTITION_SUBTYPE_DATA_COREDUMP           EspPartitionSubtypeT = 3
	ESP_PARTITION_SUBTYPE_DATA_NVS_KEYS           EspPartitionSubtypeT = 4
	ESP_PARTITION_SUBTYPE_DATA_EFUSE_EM           EspPartitionSubtypeT = 5
	ESP_PARTITION_SUBTYPE_DATA_UNDEFINED          EspPartitionSubtypeT = 6
	ESP_PARTITION_SUBTYPE_DATA_ESPHTTPD           EspPartitionSubtypeT = 128
	ESP_PARTITION_SUBTYPE_DATA_FAT                EspPartitionSubtypeT = 129
	ESP_PARTITION_SUBTYPE_DATA_SPIFFS             EspPartitionSubtypeT = 130
	ESP_PARTITION_SUBTYPE_DATA_LITTLEFS           EspPartitionSubtypeT = 131
	ESP_PARTITION_SUBTYPE_ANY                     EspPartitionSubtypeT = 255
)

type EspPartitionIteratorOpaque_ struct {
	Unused [8]uint8
}
type EspPartitionIteratorT *EspPartitionIteratorOpaque_

/**
 * @brief partition information structure
 *
 * This is not the format in flash, that format is esp_partition_info_t.
 *
 * However, this is the format used by this API.
 */

type EspPartitionT struct {
	FlashChip *EspFlashT
	Type      EspPartitionTypeT
	Subtype   EspPartitionSubtypeT
	Address   c.Uint32T
	Size      c.Uint32T
	EraseSize c.Uint32T
	Label     [17]c.Char
	Encrypted bool
	Readonly  bool
}

/**
 * @brief Find partition based on one or more parameters
 *
 * @param type Partition type, one of esp_partition_type_t values or an 8-bit unsigned integer.
 *             To find all partitions, no matter the type, use ESP_PARTITION_TYPE_ANY, and set
 *             subtype argument to ESP_PARTITION_SUBTYPE_ANY.
 * @param subtype Partition subtype, one of esp_partition_subtype_t values or an 8-bit unsigned integer.
 *                To find all partitions of given type, use ESP_PARTITION_SUBTYPE_ANY.
 * @param label (optional) Partition label. Set this value if looking
 *             for partition with a specific name. Pass NULL otherwise.
 *
 * @return iterator which can be used to enumerate all the partitions found,
 *         or NULL if no partitions were found.
 *         Iterator obtained through this function has to be released
 *         using esp_partition_iterator_release when not used any more.
 */
// llgo:link EspPartitionTypeT.EspPartitionFind C.esp_partition_find
func (recv_ EspPartitionTypeT) EspPartitionFind(subtype EspPartitionSubtypeT, label *c.Char) EspPartitionIteratorT {
	return nil
}

/**
 * @brief Find first partition based on one or more parameters
 *
 * @param type Partition type, one of esp_partition_type_t values or an 8-bit unsigned integer.
 *             To find all partitions, no matter the type, use ESP_PARTITION_TYPE_ANY, and set
 *             subtype argument to ESP_PARTITION_SUBTYPE_ANY.
 * @param subtype Partition subtype, one of esp_partition_subtype_t values or an 8-bit unsigned integer
 *                To find all partitions of given type, use ESP_PARTITION_SUBTYPE_ANY.
 * @param label (optional) Partition label. Set this value if looking
 *             for partition with a specific name. Pass NULL otherwise.
 *
 * @return pointer to esp_partition_t structure, or NULL if no partition is found.
 *         This pointer is valid for the lifetime of the application.
 */
// llgo:link EspPartitionTypeT.EspPartitionFindFirst C.esp_partition_find_first
func (recv_ EspPartitionTypeT) EspPartitionFindFirst(subtype EspPartitionSubtypeT, label *c.Char) *EspPartitionT {
	return nil
}

/**
 * @brief Get esp_partition_t structure for given partition
 *
 * @param iterator  Iterator obtained using esp_partition_find. Must be non-NULL.
 *
 * @return pointer to esp_partition_t structure. This pointer is valid for the lifetime
 *         of the application.
 */
//go:linkname EspPartitionGet C.esp_partition_get
func EspPartitionGet(iterator EspPartitionIteratorT) *EspPartitionT

/**
 * @brief Move partition iterator to the next partition found
 *
 * Any copies of the iterator will be invalid after this call.
 *
 * @param iterator Iterator obtained using esp_partition_find. Must be non-NULL.
 *
 * @return NULL if no partition was found, valid esp_partition_iterator_t otherwise.
 */
//go:linkname EspPartitionNext C.esp_partition_next
func EspPartitionNext(iterator EspPartitionIteratorT) EspPartitionIteratorT

/**
 * @brief Release partition iterator
 *
 * @param iterator Iterator obtained using esp_partition_find.
 *                 The iterator is allowed to be NULL, so it is not necessary to check its value
 *                 before calling this function.
 *
 */
//go:linkname EspPartitionIteratorRelease C.esp_partition_iterator_release
func EspPartitionIteratorRelease(iterator EspPartitionIteratorT)

/**
 * @brief Verify partition data
 *
 * Given a pointer to partition data, verify this partition exists in the partition table (all fields match.)
 *
 * This function is also useful to take partition data which may be in a RAM buffer and convert it to a pointer to the
 * permanent partition data stored in flash.
 *
 * Pointers returned from this function can be compared directly to the address of any pointer returned from
 * esp_partition_get(), as a test for equality.
 *
 * @param partition Pointer to partition data to verify. Must be non-NULL. All fields of this structure must match the
 * partition table entry in flash for this function to return a successful match.
 *
 * @return
 * - If partition not found, returns NULL.
 * - If found, returns a pointer to the esp_partition_t structure in flash. This pointer is always valid for the lifetime of the application.
 */
// llgo:link (*EspPartitionT).EspPartitionVerify C.esp_partition_verify
func (recv_ *EspPartitionT) EspPartitionVerify() *EspPartitionT {
	return nil
}

/**
 * @brief Read data from the partition
 *
 * Partitions marked with an encryption flag will automatically be
 * be read and decrypted via a cache mapping.
 *
 * @param partition Pointer to partition structure obtained using
 *                  esp_partition_find_first or esp_partition_get.
 *                  Must be non-NULL.
 * @param dst Pointer to the buffer where data should be stored.
 *            Pointer must be non-NULL and buffer must be at least 'size' bytes long.
 * @param src_offset Address of the data to be read, relative to the
 *                   beginning of the partition.
 * @param size Size of data to be read, in bytes.
 *
 * @return ESP_OK, if data was read successfully;
 *         ESP_ERR_INVALID_ARG, if src_offset exceeds partition size;
 *         ESP_ERR_INVALID_SIZE, if read would go out of bounds of the partition;
 *         or one of error codes from lower-level flash driver.
 */
// llgo:link (*EspPartitionT).EspPartitionRead C.esp_partition_read
func (recv_ *EspPartitionT) EspPartitionRead(src_offset c.SizeT, dst c.Pointer, size c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Write data to the partition
 *
 * Before writing data to flash, corresponding region of flash needs to be erased.
 * This can be done using esp_partition_erase_range function.
 *
 * Partitions marked with an encryption flag will automatically be
 * written via the esp_flash_write_encrypted() function. If writing to
 * an encrypted partition, all write offsets and lengths must be
 * multiples of 16 bytes. See the esp_flash_write_encrypted() function
 * for more details. Unencrypted partitions do not have this
 * restriction.
 *
 * @param partition Pointer to partition structure obtained using
 *                  esp_partition_find_first or esp_partition_get.
 *                  Must be non-NULL.
 * @param dst_offset Address where the data should be written, relative to the
 *                   beginning of the partition.
 * @param src Pointer to the source buffer.  Pointer must be non-NULL and
 *            buffer must be at least 'size' bytes long.
 * @param size Size of data to be written, in bytes.
 *
 * @note Prior to writing to flash memory, make sure it has been erased with
 *       esp_partition_erase_range call.
 *
 * @return ESP_OK, if data was written successfully;
 *         ESP_ERR_INVALID_ARG, if dst_offset exceeds partition size;
 *         ESP_ERR_INVALID_SIZE, if write would go out of bounds of the partition;
 *         ESP_ERR_NOT_ALLOWED, if partition is read-only;
 *         or one of error codes from lower-level flash driver.
 */
// llgo:link (*EspPartitionT).EspPartitionWrite C.esp_partition_write
func (recv_ *EspPartitionT) EspPartitionWrite(dst_offset c.SizeT, src c.Pointer, size c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Read data from the partition without any transformation/decryption.
 *
 * @note This function is essentially the same as \c esp_partition_read() above.
 *       It just never decrypts data but returns it as is.
 *
 * @param partition Pointer to partition structure obtained using
 *                  esp_partition_find_first or esp_partition_get.
 *                  Must be non-NULL.
 * @param dst Pointer to the buffer where data should be stored.
 *            Pointer must be non-NULL and buffer must be at least 'size' bytes long.
 * @param src_offset Address of the data to be read, relative to the
 *                   beginning of the partition.
 * @param size Size of data to be read, in bytes.
 *
 * @return ESP_OK, if data was read successfully;
 *         ESP_ERR_INVALID_ARG, if src_offset exceeds partition size;
 *         ESP_ERR_INVALID_SIZE, if read would go out of bounds of the partition;
 *         or one of error codes from lower-level flash driver.
 */
// llgo:link (*EspPartitionT).EspPartitionReadRaw C.esp_partition_read_raw
func (recv_ *EspPartitionT) EspPartitionReadRaw(src_offset c.SizeT, dst c.Pointer, size c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Write data to the partition without any transformation/encryption.
 *
 * @note This function is essentially the same as \c esp_partition_write() above.
 *       It just never encrypts data but writes it as is.
 *
 * Before writing data to flash, corresponding region of flash needs to be erased.
 * This can be done using esp_partition_erase_range function.
 *
 * @param partition Pointer to partition structure obtained using
 *                  esp_partition_find_first or esp_partition_get.
 *                  Must be non-NULL.
 * @param dst_offset Address where the data should be written, relative to the
 *                   beginning of the partition.
 * @param src Pointer to the source buffer.  Pointer must be non-NULL and
 *            buffer must be at least 'size' bytes long.
 * @param size Size of data to be written, in bytes.
 *
 * @note Prior to writing to flash memory, make sure it has been erased with
 *       esp_partition_erase_range call.
 *
 * @return ESP_OK, if data was written successfully;
 *         ESP_ERR_INVALID_ARG, if dst_offset exceeds partition size;
 *         ESP_ERR_INVALID_SIZE, if write would go out of bounds of the partition;
 *         ESP_ERR_NOT_ALLOWED, if partition is read-only;
 *         or one of the error codes from lower-level flash driver.
 */
// llgo:link (*EspPartitionT).EspPartitionWriteRaw C.esp_partition_write_raw
func (recv_ *EspPartitionT) EspPartitionWriteRaw(dst_offset c.SizeT, src c.Pointer, size c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Erase part of the partition
 *
 * @param partition Pointer to partition structure obtained using
 *                  esp_partition_find_first or esp_partition_get.
 *                  Must be non-NULL.
 * @param offset Offset from the beginning of partition where erase operation
 *               should start. Must be aligned to partition->erase_size.
 * @param size Size of the range which should be erased, in bytes.
 *                   Must be divisible by partition->erase_size.
 *
 * @return ESP_OK, if the range was erased successfully;
 *         ESP_ERR_INVALID_ARG, if iterator or dst are NULL;
 *         ESP_ERR_INVALID_SIZE, if erase would go out of bounds of the partition;
 *         ESP_ERR_NOT_ALLOWED, if partition is read-only;
 *         or one of error codes from lower-level flash driver.
 */
// llgo:link (*EspPartitionT).EspPartitionEraseRange C.esp_partition_erase_range
func (recv_ *EspPartitionT) EspPartitionEraseRange(offset c.SizeT, size c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Configure MMU to map partition into data memory
 *
 * Unlike spi_flash_mmap function, which requires a 64kB aligned base address,
 * this function doesn't impose such a requirement.
 * If offset results in a flash address which is not aligned to 64kB boundary,
 * address will be rounded to the lower 64kB boundary, so that mapped region
 * includes requested range.
 * Pointer returned via out_ptr argument will be adjusted to point to the
 * requested offset (not necessarily to the beginning of mmap-ed region).
 *
 * To release mapped memory, pass handle returned via out_handle argument to
 * esp_partition_munmap function.
 *
 * @param partition Pointer to partition structure obtained using
 *                  esp_partition_find_first or esp_partition_get.
 *                  Must be non-NULL.
 * @param offset Offset from the beginning of partition where mapping should start.
 * @param size Size of the area to be mapped.
 * @param memory  Memory space where the region should be mapped
 * @param out_ptr  Output, pointer to the mapped memory region
 * @param out_handle  Output, handle which should be used for esp_partition_munmap call
 *
 * @return ESP_OK, if successful
 */
// llgo:link (*EspPartitionT).EspPartitionMmap C.esp_partition_mmap
func (recv_ *EspPartitionT) EspPartitionMmap(offset c.SizeT, size c.SizeT, memory EspPartitionMmapMemoryT, out_ptr *c.Pointer, out_handle *EspPartitionMmapHandleT) EspErrT {
	return 0
}

/**
 * @brief Release region previously obtained using esp_partition_mmap
 *
 * @note Calling this function will not necessarily unmap memory region.
 *       Region will only be unmapped when there are no other handles which
 *       reference this region. In case of partially overlapping regions
 *       it is possible that memory will be unmapped partially.
 *
 * @param handle  Handle obtained from spi_flash_mmap
 */
// llgo:link EspPartitionMmapHandleT.EspPartitionMunmap C.esp_partition_munmap
func (recv_ EspPartitionMmapHandleT) EspPartitionMunmap() {
}

/**
 * @brief Get SHA-256 digest for required partition.
 *
 * For apps with SHA-256 appended to the app image, the result is the appended SHA-256 value for the app image content.
 * The hash is verified before returning, if app content is invalid then the function returns ESP_ERR_IMAGE_INVALID.
 * For apps without SHA-256 appended to the image, the result is the SHA-256 of all bytes in the app image.
 * For other partition types, the result is the SHA-256 of the entire partition.
 *
 * @param[in]  partition    Pointer to info for partition containing app or data. (fields: address, size and type, are required to be filled).
 * @param[out] sha_256      Returned SHA-256 digest for a given partition.
 *
 * @return
 *          - ESP_OK: In case of successful operation.
 *          - ESP_ERR_INVALID_ARG: The size was 0 or the sha_256 was NULL.
 *          - ESP_ERR_NO_MEM: Cannot allocate memory for sha256 operation.
 *          - ESP_ERR_IMAGE_INVALID: App partition doesn't contain a valid app image.
 *          - ESP_FAIL: An allocation error occurred.
 */
// llgo:link (*EspPartitionT).EspPartitionGetSha256 C.esp_partition_get_sha256
func (recv_ *EspPartitionT) EspPartitionGetSha256(sha_256 *c.Uint8T) EspErrT {
	return 0
}

/**
 * @brief Check for the identity of two partitions by SHA-256 digest.
 *
 * @param[in] partition_1 Pointer to info for partition 1 containing app or data. (fields: address, size and type, are required to be filled).
 * @param[in] partition_2 Pointer to info for partition 2 containing app or data. (fields: address, size and type, are required to be filled).
 *
 * @return
 *         - True:  In case of the two firmware is equal.
 *         - False: Otherwise
 */
// llgo:link (*EspPartitionT).EspPartitionCheckIdentity C.esp_partition_check_identity
func (recv_ *EspPartitionT) EspPartitionCheckIdentity(partition_2 *EspPartitionT) bool {
	return false
}

/**
 * @brief Register a partition on an external flash chip
 *
 * This API allows designating certain areas of external flash chips (identified by the esp_flash_t structure)
 * as partitions. This allows using them with components which access SPI flash through the esp_partition API.
 *
 * @param flash_chip  Pointer to the structure identifying the flash chip. If NULL then the internal flash chip is used (esp_flash_default_chip).
 * @param offset  Address in bytes, where the partition starts
 * @param size  Size of the partition in bytes
 * @param label  Partition name
 * @param type  One of the partition types (ESP_PARTITION_TYPE_*), or an integer. Note that applications can not be booted from external flash
 *              chips, so using ESP_PARTITION_TYPE_APP is not supported.
 * @param subtype  One of the partition subtypes (ESP_PARTITION_SUBTYPE_*), or an integer.
 * @param[out] out_partition  Output, if non-NULL, receives the pointer to the resulting esp_partition_t structure
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NO_MEM if memory allocation has failed
 *      - ESP_ERR_INVALID_ARG if the new partition overlaps another partition on the same flash chip
 *      - ESP_ERR_INVALID_SIZE if the partition doesn't fit into the flash chip size
 */
// llgo:link (*EspFlashT).EspPartitionRegisterExternal C.esp_partition_register_external
func (recv_ *EspFlashT) EspPartitionRegisterExternal(offset c.SizeT, size c.SizeT, label *c.Char, type_ EspPartitionTypeT, subtype EspPartitionSubtypeT, out_partition **EspPartitionT) EspErrT {
	return 0
}

/**
 * @brief Deregister the partition previously registered using esp_partition_register_external
 * @param partition  pointer to the partition structure obtained from esp_partition_register_external,
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NOT_FOUND if the partition pointer is not found
 *      - ESP_ERR_INVALID_ARG if the partition comes from the partition table
 *      - ESP_ERR_INVALID_ARG if the partition was not registered using
 *        esp_partition_register_external function.
 */
// llgo:link (*EspPartitionT).EspPartitionDeregisterExternal C.esp_partition_deregister_external
func (recv_ *EspPartitionT) EspPartitionDeregisterExternal() EspErrT {
	return 0
}

/**
 * @brief Unload partitions and free space allocated by them
 */
//go:linkname EspPartitionUnloadAll C.esp_partition_unload_all
func EspPartitionUnloadAll()

/**
 * @brief Get the main flash sector size
 * @return
 *      - SPI_FLASH_SEC_SIZE - For esp32xx target
 *      - ESP_PARTITION_EMULATED_SECTOR_SIZE - For linux target
 */
//go:linkname EspPartitionGetMainFlashSectorSize C.esp_partition_get_main_flash_sector_size
func EspPartitionGetMainFlashSectorSize() c.Uint32T

/**
 * @brief Copy data from a source partition at a specific offset to a destination partition at a specific offset.
 *
 * The destination offset must be aligned to the flash sector size (SPI_FLASH_SEC_SIZE = 0x1000).
 * If "size" is SIZE_MAX, the entire destination partition (from dest_offset onward) will be erased,
 * and the function will copy all of the source partition starting from src_offset into the destination.
 * The function ensures that the destination partition is erased on sector boundaries (erase size is aligned up SPI_FLASH_SEC_SIZE).
 *
 * This function does the following:
 * - erases the destination partition from dest_offset to the specified size (or the whole partition if "size" == SIZE_MAX),
 * - maps data from the source partition in chunks,
 * - writes the source data into the destination partition in corresponding chunks.
 *
 * @param dest_part   Pointer to a destination partition.
 * @param dest_offset Offset in the destination partition where the data should be written (must be aligned to SPI_FLASH_SEC_SIZE = 0x1000).
 * @param src_part    Pointer to a source partition (must be located on internal flash).
 * @param src_offset  Offset in the source partition where the data should be read from.
 * @param size        Number of bytes to copy from the source partition to the destination partition. If "size" is SIZE_MAX,
 *                    the function copies from src_offset to the end of the source partition and erases
 *                    the entire destination partition (from dest_offset onward).
 *
 * @return ESP_OK, if the source partition was copied successfully to the destination partition;
 *         ESP_ERR_INVALID_ARG, if src_part or dest_part are incorrect, or if dest_offset is not sector aligned;
 *         ESP_ERR_INVALID_SIZE, if the copy would go out of bounds of the source or destination partition;
 *         ESP_ERR_NOT_ALLOWED, if the destination partition is read-only;
 *         or one of the error codes from the lower-level flash driver.
 */
// llgo:link (*EspPartitionT).EspPartitionCopy C.esp_partition_copy
func (recv_ *EspPartitionT) EspPartitionCopy(dest_offset c.Uint32T, src_part *EspPartitionT, src_offset c.Uint32T, size c.SizeT) EspErrT {
	return 0
}
