package esp_eth

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Custom SPI Driver Configuration.
 * This structure declares configuration and callback functions to access Ethernet SPI module via
 * user's custom SPI driver.
 *
 */

type EthSpiCustomDriverConfigT struct {
	Config c.Pointer
	Init   c.Pointer
	Deinit c.Pointer
	Read   c.Pointer
	Write  c.Pointer
}
