package hal

import _ "unsafe"

/**
 * Context that should be maintained by both the driver and the HAL
 */

type WdtHalContextT struct {
	Inst WdtInstT
}
