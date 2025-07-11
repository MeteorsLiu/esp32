package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const FLASH_SECTOR_SIZE = 0x1000
const FLASH_BLOCK_SIZE = 0x10000
const CMD_RDID = 0x9F
const CMD_WRSR = 0x01
const CMD_WRSR2 = 0x31
const CMD_WRSR3 = 0x11
const CMD_WREN = 0x06
const CMD_WRENVSR = 0x50
const CMD_WRDI = 0x04
const CMD_RDSR = 0x05
const CMD_RDSR2 = 0x35
const CMD_RDSR3 = 0x15
const CMD_OTPEN = 0x3A
const CMD_RDSFDP = 0x5A
const CMD_RESUME = 0x7A
const CMD_RESETEN = 0x66
const CMD_RESET = 0x99
const CMD_FASTRD_QIO_4B = 0xEC
const CMD_FASTRD_QUAD_4B = 0x6C
const CMD_FASTRD_DIO_4B = 0xBC
const CMD_FASTRD_DUAL_4B = 0x3C
const CMD_FASTRD_4B = 0x0C
const CMD_SLOWRD_4B = 0x13

/**
 * @brief Get number of free pages
 *
 * @return Number of free pages
 */
//go:linkname BootloaderMmapGetFreePages C.bootloader_mmap_get_free_pages
func BootloaderMmapGetFreePages() c.Uint32T

/**
 * @brief Map a region of flash to data memory
 *
 * @important In bootloader code, only one region can be bootloader_mmaped at once. The previous region must be bootloader_munmapped before another region is mapped.
 *
 * @important In app code, these functions are not thread safe.
 *
 * Call bootloader_munmap once for each successful call to bootloader_mmap.
 *
 * In esp-idf app, this function maps directly to spi_flash_mmap.
 *
 * @param offset - Starting flash offset to map to memory.
 * @param length - Length of data to map.
 *
 * @return Pointer to mapped data memory (at src_addr), or NULL
 * if an allocation error occurred.
 */
//go:linkname BootloaderMmap C.bootloader_mmap
func BootloaderMmap(src_addr c.Uint32T, size c.Uint32T) c.Pointer

/**
 * @brief Unmap a previously mapped region of flash
 *
 * Call bootloader_munmap once for each successful call to bootloader_mmap.
 */
//go:linkname BootloaderMunmap C.bootloader_munmap
func BootloaderMunmap(mapping c.Pointer)

/**
 * @brief  Read data from Flash.
 *
 *
 * @note All of src, dest and size have to be 4-byte aligned.
 *
 * @param  src   source address of the data in Flash.
 * @param  dest  pointer to the destination buffer
 * @param  size  length of data
 * @param  allow_decrypt If true and flash encryption is enabled, data on flash
 *         will be decrypted transparently as part of the read.
 *
 * @return ESP_OK on success, ESP_ERR_FLASH_OP_FAIL on SPI failure,
 * ESP_ERR_FLASH_OP_TIMEOUT on SPI timeout.
 */
//go:linkname BootloaderFlashRead C.bootloader_flash_read
func BootloaderFlashRead(src_addr c.SizeT, dest c.Pointer, size c.SizeT, allow_decrypt bool) EspErrT

/**
 * @brief  Write data to Flash.
 *
 * @note All of dest_addr, src and size have to be 4-byte aligned. If write_encrypted is set, dest_addr and size must be 32-byte aligned.
 *
 * Note: In bootloader, when write_encrypted == true, the src buffer is encrypted in place.
 *
 * @param  dest_addr Destination address to write in Flash.
 * @param  src Pointer to the data to write to flash
 * @param  size Length of data in bytes.
 * @param  write_encrypted If true, data will be written encrypted on flash.
 *
 * @return ESP_OK on success, ESP_ERR_FLASH_OP_FAIL on SPI failure,
 * ESP_ERR_FLASH_OP_TIMEOUT on SPI timeout.
 */
//go:linkname BootloaderFlashWrite C.bootloader_flash_write
func BootloaderFlashWrite(dest_addr c.SizeT, src c.Pointer, size c.SizeT, write_encrypted bool) EspErrT

/**
 * @brief  Erase the Flash sector.
 *
 * @param  sector  Sector number, the count starts at sector 0, 4KB per sector.
 *
 * @return esp_err_t
 */
//go:linkname BootloaderFlashEraseSector C.bootloader_flash_erase_sector
func BootloaderFlashEraseSector(sector c.SizeT) EspErrT

/**
 * @brief  Erase the Flash range.
 *
 * @param  start_addr start address of flash offset
 * @param  size       sector aligned size to be erased
 *
 * @return esp_err_t
 */
//go:linkname BootloaderFlashEraseRange C.bootloader_flash_erase_range
func BootloaderFlashEraseRange(start_addr c.Uint32T, size c.Uint32T) EspErrT

/**
 * @brief Execute a user command on the flash
 *
 * @param command The command value to execute.
 * @param mosi_data MOSI data to send
 * @param mosi_len Length of MOSI data, in bits
 * @param miso_len Length of MISO data to receive, in bits
 * @return Received MISO data
 */
//go:linkname BootloaderExecuteFlashCommand C.bootloader_execute_flash_command
func BootloaderExecuteFlashCommand(command c.Uint8T, mosi_data c.Uint32T, mosi_len c.Uint8T, miso_len c.Uint8T) c.Uint32T

/**
 * @brief Read the SFDP of the flash
 *
 * @param sfdp_addr Address of the parameter to read
 * @param miso_byte_num Bytes to read
 * @return The read SFDP, little endian, 4 bytes at most
 */
//go:linkname BootloaderFlashReadSfdp C.bootloader_flash_read_sfdp
func BootloaderFlashReadSfdp(sfdp_addr c.Uint32T, miso_byte_num c.Uint) c.Uint32T

/**
 * @brief Enable the flash write protect (WEL bit).
 */
//go:linkname BootloaderEnableWp C.bootloader_enable_wp
func BootloaderEnableWp()

/**
 * @brief Once this function is called,
 * any on-going internal operations will be terminated and the device will return to its default power-on
 * state and lose all the current volatile settings, such as Volatile Status Register bits, Write Enable Latch
 * (WEL) status, Program/Erase Suspend status, etc.
 */
//go:linkname BootloaderSpiFlashReset C.bootloader_spi_flash_reset
func BootloaderSpiFlashReset()
