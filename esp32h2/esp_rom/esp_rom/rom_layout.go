package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SUPPORT_BTDM = 0
const SUPPORT_BTBB = 0
const SUPPORT_WIFI = 0
const SUPPORT_USB_DWCOTG = 0
const SUPPORT_COEXIST = 0
const SUPPORT_MBEDTLS = 0

/* Structure and functions for returning ROM global layout
 *
 * This is for address symbols defined in the linker script, which may change during ECOs.
 */

type EtsRomLayoutT struct {
	Dram0StackSharedMemStart c.Pointer
	Dram0RtosReservedStart   c.Pointer
	StackSentry              c.Pointer
	Stack                    c.Pointer
	DramStartUsbReservedRom  c.Pointer
	DramEndUsbReservedRom    c.Pointer
	DramStartUartRom         c.Pointer
	DramEndUartRom           c.Pointer
}
