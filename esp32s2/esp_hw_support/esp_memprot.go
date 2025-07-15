package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
* @brief Basic PMS interrupt source info
 */
type EspMempIntrSourceT struct {
	MemType EspMprotMemT
	Core    c.Int
}
