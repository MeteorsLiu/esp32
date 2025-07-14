package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Enable the flash encryption
 */
//go:linkname SpiFlashEncryptionHalEnable C.spi_flash_encryption_hal_enable
func SpiFlashEncryptionHalEnable()

/**
 * @brief Disable the flash encryption
 */
//go:linkname SpiFlashEncryptionHalDisable C.spi_flash_encryption_hal_disable
func SpiFlashEncryptionHalDisable()

/**
 * Prepare flash encryption before operation.
 *
 * @param address The destination address in flash for the write operation.
 * @param buffer Data for programming
 * @param size Size to program.
 *
 * @note address and buffer must be 8-word aligned.
 */
//go:linkname SpiFlashEncryptionHalPrepare C.spi_flash_encryption_hal_prepare
func SpiFlashEncryptionHalPrepare(address c.Uint32T, buffer *c.Uint32T, size c.Uint32T)

/**
 * @brief flash data encryption operation is done.
 */
//go:linkname SpiFlashEncryptionHalDone C.spi_flash_encryption_hal_done
func SpiFlashEncryptionHalDone()

/**
 * Destroy encrypted result
 */
//go:linkname SpiFlashEncryptionHalDestroy C.spi_flash_encryption_hal_destroy
func SpiFlashEncryptionHalDestroy()

/**
 * Check if is qualified to encrypt the buffer
 *
 * @param address the address of written flash partition.
 * @param length Buffer size.
 */
//go:linkname SpiFlashEncryptionHalCheck C.spi_flash_encryption_hal_check
func SpiFlashEncryptionHalCheck(address c.Uint32T, length c.Uint32T) bool
