package esp_system

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief esp_xt_wdt configuration struct
 *
 */

type EspXtWdtConfigT struct {
	Timeout             c.Uint8T
	AutoBackupClkEnable bool
}

// llgo:type C
type EspXtCallbackT func(c.Pointer)
