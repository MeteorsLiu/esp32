package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Update the flash id in g_rom_flashchip(global esp_rom_spiflash_chip_t structure).
 *
 * @return None
 */
//go:linkname BootloaderFlashUpdateId C.bootloader_flash_update_id
func BootloaderFlashUpdateId()

/**
 * @brief Update the flash size in g_rom_flashchip (global esp_rom_spiflash_chip_t structure).
 *
 * @param size The size to store, in bytes.
 * @return None
 */
//go:linkname BootloaderFlashUpdateSize C.bootloader_flash_update_size
func BootloaderFlashUpdateSize(size c.Uint32T)

/**
 * @brief Set the flash CS setup and hold time.
 *
 * @note CS setup time is recomemded to be 1.5T, and CS hold time is recommended to be 2.5T.
 *       cs_setup = 1, cs_setup_time = 0; cs_hold = 1, cs_hold_time = 1.
 *
 * @return None
 */
//go:linkname BootloaderFlashCsTimingConfig C.bootloader_flash_cs_timing_config
func BootloaderFlashCsTimingConfig()

/**
 * @brief Configure SPI flash clock.
 *
 * @note This function only set clock frequency for SPI0.
 *
 * @param pfhdr Pointer to App image header, from where to fetch flash settings.
 *
 * @return None
 */
// llgo:link (*EspImageHeaderT).BootloaderFlashClockConfig C.bootloader_flash_clock_config
func (recv_ *EspImageHeaderT) BootloaderFlashClockConfig() {
}

/**
 * @brief Configure SPI flash gpio, include the IO matrix and drive strength configuration.
 *
 * @param pfhdr Pointer to App image header, from where to fetch flash settings.
 *
 * @return None
 */
// llgo:link (*EspImageHeaderT).BootloaderFlashGpioConfig C.bootloader_flash_gpio_config
func (recv_ *EspImageHeaderT) BootloaderFlashGpioConfig() {
}

/**
 * @brief Configure SPI flash read dummy based on different mode and frequency.
 *
 * @param pfhdr Pointer to App image header, from where to fetch flash settings.
 *
 * @return None
 */
// llgo:link (*EspImageHeaderT).BootloaderFlashDummyConfig C.bootloader_flash_dummy_config
func (recv_ *EspImageHeaderT) BootloaderFlashDummyConfig() {
}

/**
 * @brief Return the pin number used for custom SPI flash and/or SPIRAM WP pin
 *
 * Can be determined by eFuse values in most cases, or overriden in configuration
 *
 * This value is only meaningful if the other SPI flash pins are overriden via eFuse.
 *
 * This value is only meaningful if flash is set to QIO or QOUT mode, or if
 * SPIRAM is enabled.
 *
 * @return Pin number to use, or -1 if the default should be kept
 */
//go:linkname BootloaderFlashGetWpPin C.bootloader_flash_get_wp_pin
func BootloaderFlashGetWpPin() c.Int
