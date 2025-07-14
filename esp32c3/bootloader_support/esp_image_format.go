package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_ERR_IMAGE_BASE = 0x2000
const ESP_IMAGE_HASH_LEN = 32

/* Structure to hold on-flash image metadata */

type EspImageMetadataT struct {
	StartAddr     c.Uint32T
	Image         EspImageHeaderT
	Segments      [16]EspImageSegmentHeaderT
	SegmentData   [16]c.Uint32T
	ImageLen      c.Uint32T
	ImageDigest   [32]c.Uint8T
	SecureVersion c.Uint32T
	MmuPageSize   c.Uint32T
}
type EspImageLoadModeT c.Int

const (
	ESP_IMAGE_VERIFY        EspImageLoadModeT = 0
	ESP_IMAGE_VERIFY_SILENT EspImageLoadModeT = 1
)

type RtcRetainMemT struct {
	Partition     EspPartitionPosT
	RebootCounter c.Uint16T
	Flags         struct {
		Val c.Uint8T
	}
	Reserve c.Uint8T
	Crc     c.Uint32T
}

/**
 * @brief Verify an app image.
 *
 * If encryption is enabled, data will be transparently decrypted.
 *
 * @param mode Mode of operation (verify, silent verify, or load).
 * @param part Partition to load the app from.
 * @param[inout] data Pointer to the image metadata structure which is be filled in by this function.
 *                    'start_addr' member should be set (to the start address of the image.)
 *                    Other fields will all be initialised by this function.
 *
 * Image validation checks:
 * - Magic byte.
 * - Partition smaller than 16MB.
 * - All segments & image fit in partition.
 * - 8 bit image checksum is valid.
 * - SHA-256 of image is valid (if image has this appended).
 * - (Signature) if signature verification is enabled.
 *
 * @return
 * - ESP_OK if verify or load was successful
 * - ESP_ERR_IMAGE_FLASH_FAIL if a SPI flash error occurs
 * - ESP_ERR_IMAGE_INVALID if the image appears invalid.
 * - ESP_ERR_INVALID_ARG if the partition or data pointers are invalid.
 */
// llgo:link EspImageLoadModeT.EspImageVerify C.esp_image_verify
func (recv_ EspImageLoadModeT) EspImageVerify(part *EspPartitionPosT, data *EspImageMetadataT) EspErrT {
	return 0
}

/**
 * @brief Get metadata of app
 *
 * If encryption is enabled, data will be transparently decrypted.
 *
 * @param part Partition to load the app from.
 * @param[out] metadata Pointer to the image metadata structure which is be filled in by this function.
 *                      Fields will all be initialised by this function.
 *
 * @return
 * - ESP_OK if filling of metadata was successful
 */
// llgo:link (*EspPartitionPosT).EspImageGetMetadata C.esp_image_get_metadata
func (recv_ *EspPartitionPosT) EspImageGetMetadata(metadata *EspImageMetadataT) EspErrT {
	return 0
}

/**
 * @brief Verify and load an app image (available only in space of bootloader).
 *
 * If encryption is enabled, data will be transparently decrypted.
 *
 * @param part Partition to load the app from.
 * @param[inout] data Pointer to the image metadata structure which is be filled in by this function.
 *                    'start_addr' member should be set (to the start address of the image.)
 *                    Other fields will all be initialised by this function.
 *
 * Image validation checks:
 * - Magic byte.
 * - Partition smaller than 16MB.
 * - All segments & image fit in partition.
 * - 8 bit image checksum is valid.
 * - SHA-256 of image is valid (if image has this appended).
 * - (Signature) if signature verification is enabled.
 *
 * @return
 * - ESP_OK if verify or load was successful
 * - ESP_ERR_IMAGE_FLASH_FAIL if a SPI flash error occurs
 * - ESP_ERR_IMAGE_INVALID if the image appears invalid.
 * - ESP_ERR_INVALID_ARG if the partition or data pointers are invalid.
 */
// llgo:link (*EspPartitionPosT).BootloaderLoadImage C.bootloader_load_image
func (recv_ *EspPartitionPosT) BootloaderLoadImage(data *EspImageMetadataT) EspErrT {
	return 0
}

/**
 * @brief Load an app image without verification (available only in space of bootloader).
 *
 * If encryption is enabled, data will be transparently decrypted.
 *
 * @param part Partition to load the app from.
 * @param[inout] data Pointer to the image metadata structure which is be filled in by this function.
 *                    'start_addr' member should be set (to the start address of the image.)
 *                    Other fields will all be initialised by this function.
 *
 * @return
 * - ESP_OK if verify or load was successful
 * - ESP_ERR_IMAGE_FLASH_FAIL if a SPI flash error occurs
 * - ESP_ERR_IMAGE_INVALID if the image appears invalid.
 * - ESP_ERR_INVALID_ARG if the partition or data pointers are invalid.
 */
// llgo:link (*EspPartitionPosT).BootloaderLoadImageNoVerify C.bootloader_load_image_no_verify
func (recv_ *EspPartitionPosT) BootloaderLoadImageNoVerify(data *EspImageMetadataT) EspErrT {
	return 0
}

/**
 * @brief Verify the bootloader image.
 *
 * @param[out] If result is ESP_OK and this pointer is non-NULL, it
 * will be set to the length of the bootloader image.
 *
 * @return As per esp_image_load_metadata().
 */
//go:linkname EspImageVerifyBootloader C.esp_image_verify_bootloader
func EspImageVerifyBootloader(length *c.Uint32T) EspErrT

/**
 * @brief Verify the bootloader image.
 *
 * @param[out] Metadata for the image. Only valid if result is ESP_OK.
 *
 * @return As per esp_image_load_metadata().
 */
// llgo:link (*EspImageMetadataT).EspImageVerifyBootloaderData C.esp_image_verify_bootloader_data
func (recv_ *EspImageMetadataT) EspImageVerifyBootloaderData() EspErrT {
	return 0
}

/**
 * @brief Get the flash size of the image
 *
 * @param app_flash_size The value configured in the image header
 * @return Actual size, in bytes.
 */
// llgo:link EspImageFlashSizeT.EspImageGetFlashSize C.esp_image_get_flash_size
func (recv_ EspImageFlashSizeT) EspImageGetFlashSize() c.Int {
	return 0
}

type EspImageFlashMappingT struct {
	DromAddr     c.Uint32T
	DromLoadAddr c.Uint32T
	DromSize     c.Uint32T
	IromAddr     c.Uint32T
	IromLoadAddr c.Uint32T
	IromSize     c.Uint32T
}
