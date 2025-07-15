package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LdoRegulatorChannelT struct {
	Unused [8]uint8
}
type EspLdoChannelHandleT *LdoRegulatorChannelT

/**
 * @brief LDO channel configurations
 */

type EspLdoChannelConfigT struct {
	ChanId    c.Int
	VoltageMv c.Int
	Flags     LdoExtraFlags
}

type LdoExtraFlags struct {
	Unused [8]uint8
}
