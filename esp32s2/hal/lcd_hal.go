package hal

import _ "unsafe"

type LcdCamDevT struct {
	Unused [8]uint8
}
type LcdSocHandleT *LcdCamDevT

/**
 * @brief LCD HAL layer context
 */

type LcdHalContextT struct {
	Dev LcdSocHandleT
}
