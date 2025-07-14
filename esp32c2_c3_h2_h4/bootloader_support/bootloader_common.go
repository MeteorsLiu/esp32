package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspCommGpioHoldT c.Int

const (
	GPIO_LONG_HOLD  EspCommGpioHoldT = 1
	GPIO_SHORT_HOLD EspCommGpioHoldT = -1
	GPIO_NOT_HOLD   EspCommGpioHoldT = 0
)

type EspImageType c.Int

const (
	ESP_IMAGE_BOOTLOADER  EspImageType = 0
	ESP_IMAGE_APPLICATION EspImageType = 1
)

/**
 * @brief Read ota_info partition and fill array from two otadata structures.
 *
 * @param[in]   ota_info It is a pointer to the OTA data partition.
 *                       The "otadata" partition (Type = "data" and SubType = "ota")
 *                       is defined in the CSV partition table.
 * @param[out]  two_otadata Pointer to array of OTA selection structure.
 * @return      - ESP_OK: On success
 *              - ESP_ERR_NOT_FOUND: Partition table does not have otadata partition
 *              - ESP_FAIL: On failure
 */
// llgo:link (*EspPartitionPosT).BootloaderCommonReadOtadata C.bootloader_common_read_otadata
func (recv_ *EspPartitionPosT) BootloaderCommonReadOtadata(two_otadata *EspOtaSelectEntryT) EspErrT {
	return 0
}

/**
 * @brief Calculate crc for the OTA data select.
 *
 * @param[in] s The OTA data select.
 * @return    Returns crc value.
 */
// llgo:link (*EspOtaSelectEntryT).BootloaderCommonOtaSelectCrc C.bootloader_common_ota_select_crc
func (recv_ *EspOtaSelectEntryT) BootloaderCommonOtaSelectCrc() c.Uint32T {
	return 0
}

/**
 * @brief Verifies the validity of the OTA data select
 *
 * @param[in] s The OTA data select.
 * @return    Returns true on valid, false otherwise.
 */
// llgo:link (*EspOtaSelectEntryT).BootloaderCommonOtaSelectValid C.bootloader_common_ota_select_valid
func (recv_ *EspOtaSelectEntryT) BootloaderCommonOtaSelectValid() bool {
	return false
}

/**
 * @brief Returns true if OTADATA is not marked as bootable partition.
 *
 * @param[in] s The OTA data select.
 * @return    Returns true if OTADATA invalid, false otherwise.
 */
// llgo:link (*EspOtaSelectEntryT).BootloaderCommonOtaSelectInvalid C.bootloader_common_ota_select_invalid
func (recv_ *EspOtaSelectEntryT) BootloaderCommonOtaSelectInvalid() bool {
	return false
}

/**
 * @brief Check if a GPIO input is held low for a long period, short period, or not
 * at all.
 *
 * This function will configure the specified GPIO as an input with internal pull-up enabled.
 *
 * If the GPIO input is held low continuously for delay_sec period then it is a long hold.
 * If the GPIO input is held low for less period then it is a short hold.
 *
 * @param[in] num_pin Number of the GPIO input.
 * @param[in] delay_sec Input must be driven low for at least this long, continuously.
 * @return esp_comm_gpio_hold_t Type of low level hold detected, if any.
 */
//go:linkname BootloaderCommonCheckLongHoldGpio C.bootloader_common_check_long_hold_gpio
func BootloaderCommonCheckLongHoldGpio(num_pin c.Uint32T, delay_sec c.Uint32T) EspCommGpioHoldT

/**
 * @brief Check if a GPIO input is held low or high for a long period, short period, or not
 * at all.
 *
 * This function will configure the specified GPIO as an input with internal pull-up enabled.
 *
 * If the GPIO input is held at 'level' continuously for delay_sec period then it is a long hold.
 * If the GPIO input is held at 'level' for less period then it is a short hold.
 *
 * @param[in] num_pin Number of the GPIO input.
 * @param[in] delay_sec Input must be driven to 'level' for at least this long, continuously.
 * @param[in] level Input pin level to trigger on hold
 * @return esp_comm_gpio_hold_t Type of hold detected, if any.
 */
//go:linkname BootloaderCommonCheckLongHoldGpioLevel C.bootloader_common_check_long_hold_gpio_level
func BootloaderCommonCheckLongHoldGpioLevel(num_pin c.Uint32T, delay_sec c.Uint32T, level bool) EspCommGpioHoldT

/**
 * @brief Erase the partition data that is specified in the transferred list.
 *
 * @param[in] list_erase String containing a list of cleared partitions. Like this "nvs, phy". The string must be null-terminal.
 * @param[in] ota_data_erase If true then the OTA data partition will be cleared (if there is it in partition table).
 * @return    Returns true on success, false otherwise.
 */
//go:linkname BootloaderCommonErasePartTypeData C.bootloader_common_erase_part_type_data
func BootloaderCommonErasePartTypeData(list_erase *c.Char, ota_data_erase bool) bool

/**
 * @brief Determines if the list contains the label
 *
 * @param[in] list  A string of names delimited by commas or spaces. Like this "nvs, phy, data". The string must be null-terminated.
 * @param[in] label The substring that will be searched in the list.
 * @return    Returns true if the list contains the label, false otherwise.
 */
//go:linkname BootloaderCommonLabelSearch C.bootloader_common_label_search
func BootloaderCommonLabelSearch(list *c.Char, label *c.Char) bool

/**
 * @brief Configure default SPI pin modes and drive strengths
 *
 * @param drv GPIO drive level (determined by clock frequency)
 */
//go:linkname BootloaderConfigureSpiPins C.bootloader_configure_spi_pins
func BootloaderConfigureSpiPins(drv c.Int)

/**
 * @brief Calculates a sha-256 for a given partition or returns a appended digest.
 *
 * This function can be used to return the SHA-256 digest of application, bootloader and data partitions.
 * For apps with SHA-256 appended to the app image, the result is the appended SHA-256 value for the app image content.
 * The hash is verified before returning, if app content is invalid then the function returns ESP_ERR_IMAGE_INVALID.
 * For apps without SHA-256 appended to the image, the result is the SHA-256 of all bytes in the app image.
 * For other partition types, the result is the SHA-256 of the entire partition.
 *
 * @param[in]  address      Address of partition.
 * @param[in]  size         Size of partition.
 * @param[in]  type         Type of partition. For applications the type is 0, otherwise type is data.
 * @param[out] out_sha_256  Returned SHA-256 digest for a given partition.
 *
 * @return
 *          - ESP_OK: In case of successful operation.
 *          - ESP_ERR_INVALID_ARG: The size was 0 or the sha_256 was NULL.
 *          - ESP_ERR_NO_MEM: Cannot allocate memory for sha256 operation.
 *          - ESP_ERR_IMAGE_INVALID: App partition doesn't contain a valid app image.
 *          - ESP_FAIL: An allocation error occurred.
 */
//go:linkname BootloaderCommonGetSha256OfPartition C.bootloader_common_get_sha256_of_partition
func BootloaderCommonGetSha256OfPartition(address c.Uint32T, size c.Uint32T, type_ c.Int, out_sha_256 *c.Uint8T) EspErrT

/**
 * @brief Returns the number of active otadata.
 *
 * @param[in] two_otadata Pointer on array from two otadata structures.
 *
 * @return The number of active otadata (0 or 1).
 *        - -1: If it does not have active otadata.
 */
// llgo:link (*EspOtaSelectEntryT).BootloaderCommonGetActiveOtadata C.bootloader_common_get_active_otadata
func (recv_ *EspOtaSelectEntryT) BootloaderCommonGetActiveOtadata() c.Int {
	return 0
}

/**
 * @brief Returns the number of active otadata.
 *
 * @param[in] two_otadata       Pointer on array from two otadata structures.
 * @param[in] valid_two_otadata Pointer on array from two bools. True means select.
 * @param[in] max               True - will select the maximum ota_seq number, otherwise the minimum.
 *
 * @return The number of active otadata (0 or 1).
 *        - -1: If it does not have active otadata.
 */
// llgo:link (*EspOtaSelectEntryT).BootloaderCommonSelectOtadata C.bootloader_common_select_otadata
func (recv_ *EspOtaSelectEntryT) BootloaderCommonSelectOtadata(valid_two_otadata *bool, max bool) c.Int {
	return 0
}

/**
 * @brief Get chip package
 *
 * @return Chip package number
 */
//go:linkname BootloaderCommonGetChipVerPkg C.bootloader_common_get_chip_ver_pkg
func BootloaderCommonGetChipVerPkg() c.Uint32T

/**
 * @brief Check if the image (bootloader and application) has valid chip ID and revision
 *
 * @param[in] img_hdr: image header
 * @param[in] type: image type, bootloader or application
 * @return
 *      - ESP_OK: image and chip are matched well
 *      - ESP_FAIL: image doesn't match to the chip
 */
// llgo:link (*EspImageHeaderT).BootloaderCommonCheckChipValidity C.bootloader_common_check_chip_validity
func (recv_ *EspImageHeaderT) BootloaderCommonCheckChipValidity(type_ EspImageType) EspErrT {
	return 0
}

/**
 * @brief Check the eFuse block revision
 *
 * @param[in] min_rev_full The required minimum revision of the eFuse block
 * @param[in] max_rev_full The required maximum revision of the eFuse block
 * @return
 *          - ESP_OK: The eFuse block revision is in the required range.
 *          - ESP_OK: DISABLE_BLK_VERSION_MAJOR has been set in the eFuse of the SoC. No requirements shall be checked at this time.
 *          - ESP_FAIL: The eFuse block revision of this chip does not match the requirement of the current image.
 */
//go:linkname BootloaderCommonCheckEfuseBlkValidity C.bootloader_common_check_efuse_blk_validity
func BootloaderCommonCheckEfuseBlkValidity(min_rev_full c.Uint32T, max_rev_full c.Uint32T) EspErrT

/**
 * @brief Configure VDDSDIO, call this API to rise VDDSDIO to 1.9V when VDDSDIO regulator is enabled as 1.8V mode.
 */
//go:linkname BootloaderCommonVddsdioConfigure C.bootloader_common_vddsdio_configure
func BootloaderCommonVddsdioConfigure()
