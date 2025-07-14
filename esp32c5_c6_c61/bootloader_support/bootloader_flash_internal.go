package bootloader_support

import _ "unsafe"

/**
 * @brief Initialize spi_flash in bootloader and print flash info
 *
 * @return ESP_OK on success, otherwise see esp_err_t
 */
//go:linkname BootloaderInitSpiFlash C.bootloader_init_spi_flash
func BootloaderInitSpiFlash() EspErrT
