package hal

import _ "unsafe"

type SocEtmDevT struct {
	Unused [8]uint8
}
type EtmSocHandleT *SocEtmDevT

/**
 * @brief HAL context type of ETM driver
 */

type EtmHalContextT struct {
	Regs EtmSocHandleT
}
