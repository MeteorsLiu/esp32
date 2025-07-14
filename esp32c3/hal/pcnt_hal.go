package hal

import _ "unsafe"

type PcntDevT struct {
	Unused [8]uint8
}
type PcntSocHandleT *PcntDevT

/**
 * Context that should be maintained by both the driver and the HAL
 */

type PcntHalContextT struct {
	Dev PcntSocHandleT
}
