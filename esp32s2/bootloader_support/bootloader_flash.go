package bootloader_support

import _ "unsafe"

/**
 * @brief Startup flow recommended by XMC. Call at startup before any erase/write operation.
 *
 * @return ESP_OK When startup successfully, otherwise ESP_FAIL (indiciating you should reboot before erase/write).
 */
//go:linkname BootloaderFlashXmcStartup C.bootloader_flash_xmc_startup
func BootloaderFlashXmcStartup() EspErrT

/**
 * @brief Reset the flash chip (66H + 99H).
 *
 * @return ESP_OK if success, otherwise ESP_FAIL.
 */
//go:linkname BootloaderFlashResetChip C.bootloader_flash_reset_chip
func BootloaderFlashResetChip() EspErrT

/**
 * @brief Check if octal flash mode is enabled in eFuse
 *
 * @return True if flash is in octal mode, false else
 */
//go:linkname BootloaderFlashIsOctalModeEnabled C.bootloader_flash_is_octal_mode_enabled
func BootloaderFlashIsOctalModeEnabled() bool

/**
 * @brief Get the spi flash working mode.
 *
 * @return The mode of flash working mode, see `esp_rom_spiflash_read_mode_t`
 */
//go:linkname BootloaderFlashGetSpiMode C.bootloader_flash_get_spi_mode
func BootloaderFlashGetSpiMode() EspRomSpiflashReadModeT
