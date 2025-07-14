package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** @brief Initialise the default SPI flash chip
 *
 * Called by OS startup code. You do not need to call this in your own applications.
 */
//go:linkname EspFlashInitDefaultChip C.esp_flash_init_default_chip
func EspFlashInitDefaultChip() EspErrT

/**
 *  Enable OS-level SPI flash protections in IDF
 *
 *  Called by OS startup code. You do not need to call this in your own applications.
 *
 * @return ESP_OK if success, otherwise failed. See return value of ``esp_flash_init_os_functions``.
 */
//go:linkname EspFlashAppInit C.esp_flash_app_init
func EspFlashAppInit() EspErrT

/**
 *  Disable (or enable) OS-level SPI flash protections in IDF
 *
 *  Called by the IDF internal code (e.g. coredump). You do not need to call this in your own applications.
 *
 * @return always ESP_OK.
 */
//go:linkname EspFlashAppDisableProtect C.esp_flash_app_disable_protect
func EspFlashAppDisableProtect(disable bool) EspErrT

/**
 *  Initialize OS-level functions for a specific chip.
 *
 * @param chip The chip to init os functions.
 * @param host_id Which SPI host to use, 1 for SPI1, 2 for SPI2 (HSPI), 3 for SPI3 (VSPI)
 * @param dev_handle SPI bus lock device handle to acquire during flash operations
 *
 * @return
 *      - ESP_OK if success
 *      - ESP_ERR_INVALID_ARG if host_id is invalid
 */
// llgo:link (*EspFlashT).EspFlashInitOsFunctions C.esp_flash_init_os_functions
func (recv_ *EspFlashT) EspFlashInitOsFunctions(host_id c.Int, dev_handle SpiBusLockDevHandleT) EspErrT {
	return 0
}

/**
 * @brief Deinitialize OS-level functions
 *
 * @param chip              The chip to deinit os functions
 * @param out_dev_handle    The SPI bus lock passed from `esp_flash_init_os_functions`. The caller should deinitialize
 *                          the lock.
 * @return always ESP_OK.
 */
// llgo:link (*EspFlashT).EspFlashDeinitOsFunctions C.esp_flash_deinit_os_functions
func (recv_ *EspFlashT) EspFlashDeinitOsFunctions(out_dev_handle *SpiBusLockDevHandleT) EspErrT {
	return 0
}

/**
 * @brief Initialize the bus lock on the SPI1 bus. Should be called if drivers (including esp_flash)
 * wants to use SPI1 bus.
 *
 * @note When using legacy spi flash API, the bus lock will not be available on SPI1 bus.
 *
 * @return esp_err_t always ESP_OK.
 */
//go:linkname EspFlashInitMainBusLock C.esp_flash_init_main_bus_lock
func EspFlashInitMainBusLock() EspErrT

/**
 *  Initialize OS-level functions for the main flash chip.
 *
 * @param chip The chip to init os functions. Only pointer to the default chip is supported now.
 *
 * @return always ESP_OK
 */
// llgo:link (*EspFlashT).EspFlashAppEnableOsFunctions C.esp_flash_app_enable_os_functions
func (recv_ *EspFlashT) EspFlashAppEnableOsFunctions() EspErrT {
	return 0
}

/**
 *  Disable OS-level functions for the main flash chip during special phases (e.g. coredump)
 *
 * @param chip The chip to init os functions. Only "esp_flash_default_chip" is supported now.
 *
 * @return always ESP_OK
 */
// llgo:link (*EspFlashT).EspFlashAppDisableOsFunctions C.esp_flash_app_disable_os_functions
func (recv_ *EspFlashT) EspFlashAppDisableOsFunctions() EspErrT {
	return 0
}

/**
 * @brief Set or clear dangerous write protection check on the flash chip.
 *
 * This function sets the runtime option to allow or disallow writing to
 * dangerous areas such as the bootloader and partition table. If
 * CONFIG_SPI_FLASH_DANGEROUS_WRITE_ALLOWED is not set, this function allows
 * the caller to toggle the protection for specific areas.
 *
 * If CONFIG_SPI_FLASH_DANGEROUS_WRITE_ALLOWED is set, there is no protection
 * check in the system, and this function does nothing.
 *
 * @param chip The flash chip on which to set the write protection. Only
 *             "esp_flash_default_chip" is supported.
 * @param protect Set to true to enable protection against writing in dangerous
 *                areas (bootloader, partition table). Set to false to disable
 *                the protection.
 * @return
 *         - ESP_OK: Successful operation.
 *         - ESP_ERR_INVALID_ARG: The chip argument is null.
 */
// llgo:link (*EspFlashT).EspFlashSetDangerousWriteProtection C.esp_flash_set_dangerous_write_protection
func (recv_ *EspFlashT) EspFlashSetDangerousWriteProtection(protect bool) EspErrT {
	return 0
}
