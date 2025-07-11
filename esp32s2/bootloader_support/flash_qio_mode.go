package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** @brief Enable Quad I/O mode in bootloader (if configured)
 *
 * Queries attached SPI flash ID and sends correct SPI flash
 * commands to enable QIO or QOUT mode, then enables this mode.
 */
//go:linkname BootloaderEnableQioMode C.bootloader_enable_qio_mode
func BootloaderEnableQioMode()

/**
* @brief Read flash ID by sending 0x9F command
* @return flash raw ID
*     mfg_id = (ID >> 16) & 0xFF;
      flash_id = ID & 0xffff;
*/
//go:linkname BootloaderReadFlashId C.bootloader_read_flash_id
func BootloaderReadFlashId() c.Uint32T
