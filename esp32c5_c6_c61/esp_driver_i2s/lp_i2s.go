package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief LP I2S controller channel configuration
 */
type LpI2sChanConfigT struct {
	Id        c.Int
	Role      I2sRoleT
	Threshold c.SizeT
}

/**
 * @brief LP I2S event callbacks
 */

type LpI2sEvtCbsT struct {
	OnThreshMet       LpI2sCallbackT
	OnRequestNewTrans LpI2sCallbackT
}
