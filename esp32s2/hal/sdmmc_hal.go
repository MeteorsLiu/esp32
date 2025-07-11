package hal

import _ "unsafe"

type SdmmcDevT struct {
	Unused [8]uint8
}
type SdmmcSocHandleT *SdmmcDevT

/**
 * @brief Context of the HAL
 */

type SdmmcHalContextT struct {
	Dev SdmmcSocHandleT
}
