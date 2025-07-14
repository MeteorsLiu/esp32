package hal

import _ "unsafe"

type ParlIoDevT struct {
	Unused [8]uint8
}
type ParlioSocHandleT *ParlIoDevT

/**
 * @brief HAL context type of Parallel IO driver
 */

type ParlioHalContextT struct {
	Regs ParlioSocHandleT
}
