package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LpI2sSocHandleT *c.Uint32T

/**
 * Context that should be maintained by both the driver and the HAL
 */

type LpI2sHalContextT struct {
	Dev LpI2sSocHandleT
}
