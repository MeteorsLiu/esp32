package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SUPPORT_BTDM = 0
const SUPPORT_BTBB = 0
const SUPPORT_WIFI = 1
const SUPPORT_USB_DWCOTG = 0
const SUPPORT_COEXIST = 1
const SUPPORT_MBEDTLS = 0

/* Structure and functions for returning ROM global layout
 *
 * This is for address symbols defined in the linker script, which may change during ECOs.
 */

type EtsRomLayoutT struct {
	Dram0StackSharedMemStart   c.Pointer
	Dram0RtosReservedStart     c.Pointer
	StackSentry                c.Pointer
	Stack                      c.Pointer
	DramStartPhyrom            c.Pointer
	DramEndPhyrom              c.Pointer
	DramStartNet80211          c.Pointer
	DramEndNet80211            c.Pointer
	DataStartInterfaceNet80211 c.Pointer
	DataEndInterfaceNet80211   c.Pointer
	BssStartInterfaceNet80211  c.Pointer
	BssEndInterfaceNet80211    c.Pointer
	DramStartPp                c.Pointer
	DramEndPp                  c.Pointer
	DataStartInterfacePp       c.Pointer
	DataEndInterfacePp         c.Pointer
	BssStartInterfacePp        c.Pointer
	BssEndInterfacePp          c.Pointer
	DramStartCoexist           c.Pointer
	DramEndCoexist             c.Pointer
	DataStartInterfaceCoexist  c.Pointer
	DataEndInterfaceCoexist    c.Pointer
	BssStartInterfaceCoexist   c.Pointer
	BssEndInterfaceCoexist     c.Pointer
	DramStartUsbReservedRom    c.Pointer
	DramEndUsbReservedRom      c.Pointer
	DramStartUartRom           c.Pointer
	DramEndUartRom             c.Pointer
}
