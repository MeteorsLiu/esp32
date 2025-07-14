package esp_eth

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspEthMacS struct {
	Unused [8]uint8
}
type EspEthMacT EspEthMacS

/**
* @brief Configuration of Ethernet MAC object
*
 */
type EthMacConfigT struct {
	SwResetTimeoutMs c.Uint32T
	RxTaskStackSize  c.Uint32T
	RxTaskPrio       c.Uint32T
	Flags            c.Uint32T
}
