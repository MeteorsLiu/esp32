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

/**
 * @brief Init the sdmmc hal context.
 *
 * @param hal Context of the HAL
 */
// llgo:link (*SdmmcHalContextT).SdmmcHalInit C.sdmmc_hal_init
func (recv_ *SdmmcHalContextT) SdmmcHalInit() {
}
