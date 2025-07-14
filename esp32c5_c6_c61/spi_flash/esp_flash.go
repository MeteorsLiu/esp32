package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SpiFlashChipT struct {
	Unused [8]uint8
}

type EspFlashT struct {
	Unused [8]uint8
}

/** @brief Structure for describing a region of flash */

type EspFlashRegionT struct {
	Offset c.Uint32T
	Size   c.Uint32T
}

/** @brief OS-level integration hooks for accessing flash chips inside a running OS
 *
 * It's in the public header because some instances should be allocated statically in the startup
 * code. May be updated according to hardware version and new flash chip feature requirements,
 * shouldn't be treated as public API.
 *
 *  For advanced developers, you may replace some of them with your implementations at your own
 *  risk.
 */
type EspFlashOsFunctionsT struct {
	Start             c.Pointer
	End               c.Pointer
	RegionProtected   c.Pointer
	DelayUs           c.Pointer
	GetTempBuffer     c.Pointer
	ReleaseTempBuffer c.Pointer
	CheckYield        c.Pointer
	Yield             c.Pointer
	GetSystemTime     c.Pointer
	SetFlashOpStatus  c.Pointer
}

/** @brief Initialise SPI flash chip interface.
 *
 * This function must be called before any other API functions are called for this chip.
 *
 * @note Only the ``host`` and ``read_mode`` fields of the chip structure must
 *       be initialised before this function is called. Other fields may be
 *       auto-detected if left set to zero or NULL.
 *
 * @note If the chip->drv pointer is NULL, chip chip_drv will be auto-detected
 *       based on its manufacturer & product IDs. See
 *       ``esp_flash_registered_flash_drivers`` pointer for details of this process.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @return ESP_OK on success, or a flash error code if initialisation fails.
 */
// llgo:link (*EspFlashT).EspFlashInit C.esp_flash_init
func (recv_ *EspFlashT) EspFlashInit() EspErrT {
	return 0
}

/**
 * Check if appropriate chip driver is set.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 *
 * @return true if set, otherwise false.
 */
// llgo:link (*EspFlashT).EspFlashChipDriverInitialized C.esp_flash_chip_driver_initialized
func (recv_ *EspFlashT) EspFlashChipDriverInitialized() bool {
	return false
}

/** @brief Read flash ID via the common "RDID" SPI flash command.
 *
 * @param chip Pointer to identify flash chip. Must have been successfully initialised via esp_flash_init()
 * @param[out] out_id Pointer to receive ID value.
 *
 * ID is a 24-bit value. Lower 16 bits of 'id' are the chip ID, upper 8 bits are the manufacturer ID.
 *
 * @return ESP_OK on success, or a flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashReadId C.esp_flash_read_id
func (recv_ *EspFlashT) EspFlashReadId(out_id *c.Uint32T) EspErrT {
	return 0
}

/** @brief Detect flash size based on flash ID.
 *
 * @param chip Pointer to identify flash chip. Must have been successfully initialised via esp_flash_init()
 * @param[out] out_size Detected size in bytes, standing for the size in the binary image header.
 *
 * @note 1. Most flash chips use a common format for flash ID, where the lower 4 bits specify the size as a power of 2. If
 * the manufacturer doesn't follow this convention, the size may be incorrectly detected.
 *       2. The out_size returned only stands for The out_size stands for the size in the binary image header.
 *  If you want to get the real size of the chip, please call `esp_flash_get_physical_size` instead.
 *
 * @return ESP_OK on success, or a flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashGetSize C.esp_flash_get_size
func (recv_ *EspFlashT) EspFlashGetSize(out_size *c.Uint32T) EspErrT {
	return 0
}

/** @brief Detect flash size based on flash ID.
 *
 * @param chip Pointer to identify flash chip. Must have been successfully initialised via esp_flash_init()
 * @param[out] flash_size Detected size in bytes.
 *
 * @note Most flash chips use a common format for flash ID, where the lower 4 bits specify the size as a power of 2. If
 * the manufacturer doesn't follow this convention, the size may be incorrectly detected.
 *
 * @return ESP_OK on success, or a flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashGetPhysicalSize C.esp_flash_get_physical_size
func (recv_ *EspFlashT) EspFlashGetPhysicalSize(flash_size *c.Uint32T) EspErrT {
	return 0
}

/** @brief Read flash unique ID via the common "RDUID" SPI flash command.
 *
 * @note This is an optional feature, which is not supported on all flash chips. READ PROGRAMMING GUIDE FIRST!
 *
 * @param chip Pointer to identify flash chip. Must have been successfully initialised via esp_flash_init().
 * @param[out] out_id Pointer to receive unique ID value.
 *
 * ID is a 64-bit value.
 *
 * @return
 *      - ESP_OK on success, or a flash error code if operation failed.
 *      - ESP_ERR_NOT_SUPPORTED if the chip doesn't support read id.
 */
// llgo:link (*EspFlashT).EspFlashReadUniqueChipId C.esp_flash_read_unique_chip_id
func (recv_ *EspFlashT) EspFlashReadUniqueChipId(out_id *c.Uint64T) EspErrT {
	return 0
}

/** @brief Erase flash chip contents
 *
 * @param chip Pointer to identify flash chip. Must have been successfully initialised via esp_flash_init()
 *
 *
 * @return
 *      - ESP_OK on success,
 *      - ESP_ERR_NOT_SUPPORTED if the chip is not able to perform the operation. This is indicated by WREN = 1 after the command is sent.
 *      - ESP_ERR_NOT_ALLOWED if a read-only partition is present.
 *      - Other flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashEraseChip C.esp_flash_erase_chip
func (recv_ *EspFlashT) EspFlashEraseChip() EspErrT {
	return 0
}

/** @brief Erase a region of the flash chip
 *
 * @param chip Pointer to identify flash chip. If NULL, esp_flash_default_chip is substituted. Must have been successfully initialised via esp_flash_init()
 * @param start Address to start erasing flash. Must be sector aligned.
 * @param len Length of region to erase. Must also be sector aligned.
 *
 * Sector size is specified in chip->drv->sector_size field (typically 4096 bytes.) ESP_ERR_INVALID_ARG will be
 * returned if the start & length are not a multiple of this size.
 *
 * Erase is performed using block (multi-sector) erases where possible (block size is specified in
 * chip->drv->block_erase_size field, typically 65536 bytes). Remaining sectors are erased using individual sector erase
 * commands.
 *
 * @return
 *      - ESP_OK on success,
 *      - ESP_ERR_NOT_SUPPORTED if the chip is not able to perform the operation. This is indicated by WREN = 1 after the command is sent.
 *      - ESP_ERR_NOT_ALLOWED if the address range (start -- start + len) overlaps with a read-only partition address space
 *      - Other flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashEraseRegion C.esp_flash_erase_region
func (recv_ *EspFlashT) EspFlashEraseRegion(start c.Uint32T, len c.Uint32T) EspErrT {
	return 0
}

/** @brief Read if the entire chip is write protected
 *
 * @param chip Pointer to identify flash chip. If NULL, esp_flash_default_chip is substituted. Must have been successfully initialised via esp_flash_init()
 * @param[out] write_protected Pointer to boolean, set to the value of the write protect flag.
 *
 * @note A correct result for this flag depends on the SPI flash chip model and chip_drv in use (via the 'chip->drv'
 * field).
 *
 * @return ESP_OK on success, or a flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashGetChipWriteProtect C.esp_flash_get_chip_write_protect
func (recv_ *EspFlashT) EspFlashGetChipWriteProtect(write_protected *bool) EspErrT {
	return 0
}

/** @brief Set write protection for the SPI flash chip
 *
 * @param chip Pointer to identify flash chip. If NULL, esp_flash_default_chip is substituted. Must have been successfully initialised via esp_flash_init()
 * @param write_protect Boolean value for the write protect flag
 *
 * @note Correct behaviour of this function depends on the SPI flash chip model and chip_drv in use (via the 'chip->drv'
 * field).
 *
 * Some SPI flash chips may require a power cycle before write protect status can be cleared. Otherwise,
 * write protection can be removed via a follow-up call to this function.
 *
 * @return ESP_OK on success, or a flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashSetChipWriteProtect C.esp_flash_set_chip_write_protect
func (recv_ *EspFlashT) EspFlashSetChipWriteProtect(write_protect bool) EspErrT {
	return 0
}

/** @brief Read the list of individually protectable regions of this SPI flash chip.
 *
 * @param chip Pointer to identify flash chip. Must have been successfully initialised via esp_flash_init()
 * @param[out] out_regions Pointer to receive a pointer to the array of protectable regions of the chip.
 * @param[out] out_num_regions Pointer to an integer receiving the count of protectable regions in the array returned in 'regions'.
 *
 * @note Correct behaviour of this function depends on the SPI flash chip model and chip_drv in use (via the 'chip->drv'
 * field).
 *
 * @return ESP_OK on success, or a flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashGetProtectableRegions C.esp_flash_get_protectable_regions
func (recv_ *EspFlashT) EspFlashGetProtectableRegions(out_regions **EspFlashRegionT, out_num_regions *c.Uint32T) EspErrT {
	return 0
}

/** @brief Detect if a region of the SPI flash chip is protected
 *
 * @param chip Pointer to identify flash chip. Must have been successfully initialised via esp_flash_init()
 * @param region Pointer to a struct describing a protected region. This must match one of the regions returned from esp_flash_get_protectable_regions(...).
 * @param[out] out_protected Pointer to a flag which is set based on the protected status for this region.
 *
 * @note It is possible for this result to be false and write operations to still fail, if protection is enabled for the entire chip.
 *
 * @note Correct behaviour of this function depends on the SPI flash chip model and chip_drv in use (via the 'chip->drv'
 * field).
 *
 * @return ESP_OK on success, or a flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashGetProtectedRegion C.esp_flash_get_protected_region
func (recv_ *EspFlashT) EspFlashGetProtectedRegion(region *EspFlashRegionT, out_protected *bool) EspErrT {
	return 0
}

/** @brief Update the protected status for a region of the SPI flash chip
 *
 * @param chip Pointer to identify flash chip. Must have been successfully initialised via esp_flash_init()
 * @param region Pointer to a struct describing a protected region. This must match one of the regions returned from esp_flash_get_protectable_regions(...).
 * @param protect Write protection flag to set.
 *
 * @note It is possible for the region protection flag to be cleared and write operations to still fail, if protection is enabled for the entire chip.
 *
 * @note Correct behaviour of this function depends on the SPI flash chip model and chip_drv in use (via the 'chip->drv'
 * field).
 *
 * @return ESP_OK on success, or a flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashSetProtectedRegion C.esp_flash_set_protected_region
func (recv_ *EspFlashT) EspFlashSetProtectedRegion(region *EspFlashRegionT, protect bool) EspErrT {
	return 0
}

/** @brief Read data from the SPI flash chip
 *
 * @param chip Pointer to identify flash chip. If NULL, esp_flash_default_chip is substituted. Must have been successfully initialised via esp_flash_init()
 * @param buffer Pointer to a buffer where the data will be read. To get better performance, this should be in the DRAM and word aligned.
 * @param address Address on flash to read from. Must be less than chip->size field.
 * @param length Length (in bytes) of data to read.
 *
 * There are no alignment constraints on buffer, address or length.
 *
 * @note If on-chip flash encryption is used, this function returns raw (ie encrypted) data. Use the flash cache
 * to transparently decrypt data.
 *
 * @return
 *      - ESP_OK: success
 *      - ESP_ERR_NO_MEM: Buffer is in external PSRAM which cannot be concurrently accessed, and a temporary internal buffer could not be allocated.
 *      - or a flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashRead C.esp_flash_read
func (recv_ *EspFlashT) EspFlashRead(buffer c.Pointer, address c.Uint32T, length c.Uint32T) EspErrT {
	return 0
}

/** @brief Write data to the SPI flash chip
 *
 * @param chip Pointer to identify flash chip. If NULL, esp_flash_default_chip is substituted. Must have been successfully initialised via esp_flash_init()
 * @param address Address on flash to write to. Must be previously erased (SPI NOR flash can only write bits 1->0).
 * @param buffer Pointer to a buffer with the data to write. To get better performance, this should be in the DRAM and word aligned.
 * @param length Length (in bytes) of data to write.
 *
 * There are no alignment constraints on buffer, address or length.
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_FAIL, bad write, this will be detected only when CONFIG_SPI_FLASH_VERIFY_WRITE is enabled
 *      - ESP_ERR_NOT_SUPPORTED if the chip is not able to perform the operation. This is indicated by WREN = 1 after the command is sent.
 *      - ESP_ERR_NOT_ALLOWED if the address range (address -- address + length) overlaps with a read-only partition address space
 *      - Other flash error code if operation failed.
 */
// llgo:link (*EspFlashT).EspFlashWrite C.esp_flash_write
func (recv_ *EspFlashT) EspFlashWrite(buffer c.Pointer, address c.Uint32T, length c.Uint32T) EspErrT {
	return 0
}

/** @brief Encrypted and write data to the SPI flash chip using on-chip hardware flash encryption
 *
 * @param chip Pointer to identify flash chip. Must be NULL (the main flash chip). For other chips, encrypted write is not supported.
 * @param address Address on flash to write to. 16 byte aligned. Must be previously erased (SPI NOR flash can only write bits 1->0).
 * @param buffer Pointer to a buffer with the data to write.
 * @param length Length (in bytes) of data to write. 16 byte aligned.
 *
 * @note Both address & length must be 16 byte aligned, as this is the encryption block size
 *
 * @return
 *  - ESP_OK: on success
 *  - ESP_FAIL: bad write, this will be detected only when CONFIG_SPI_FLASH_VERIFY_WRITE is enabled
 *  - ESP_ERR_NOT_SUPPORTED: encrypted write not supported for this chip.
 *  - ESP_ERR_INVALID_ARG: Either the address, buffer or length is invalid.
 *  - ESP_ERR_NOT_ALLOWED if the address range (address -- address + length) overlaps with a read-only partition address space
 */
// llgo:link (*EspFlashT).EspFlashWriteEncrypted C.esp_flash_write_encrypted
func (recv_ *EspFlashT) EspFlashWriteEncrypted(address c.Uint32T, buffer c.Pointer, length c.Uint32T) EspErrT {
	return 0
}

/** @brief Read and decrypt data from the SPI flash chip using on-chip hardware flash encryption
 *
 * @param chip Pointer to identify flash chip. Must be NULL (the main flash chip). For other chips, encrypted read is not supported.
 * @param address Address on flash to read from.
 * @param out_buffer Pointer to a buffer for the data to read to.
 * @param length Length (in bytes) of data to read.
 *
 * @return
 *  - ESP_OK: on success
 *  - ESP_ERR_NOT_SUPPORTED: encrypted read not supported for this chip.
 */
// llgo:link (*EspFlashT).EspFlashReadEncrypted C.esp_flash_read_encrypted
func (recv_ *EspFlashT) EspFlashReadEncrypted(address c.Uint32T, out_buffer c.Pointer, length c.Uint32T) EspErrT {
	return 0
}
