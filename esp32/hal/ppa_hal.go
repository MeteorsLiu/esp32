package hal

import _ "unsafe"

type PpaDevT struct {
	Unused [8]uint8
}
type PpaSocHandleT *PpaDevT

/**
 * Context that should be maintained by both the driver and the HAL
 */

type PpaHalContextT struct {
	Dev PpaSocHandleT
}
