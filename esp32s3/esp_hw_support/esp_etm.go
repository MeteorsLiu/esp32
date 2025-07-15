package esp_hw_support

import _ "unsafe"

type EspEtmChannelT struct {
	Unused [8]uint8
}
type EspEtmChannelHandleT *EspEtmChannelT
type EspEtmEventHandleT *EspEtmEventT
type EspEtmTaskHandleT *EspEtmTaskT

/**
 * @brief ETM channel configuration
 */

type EspEtmChannelConfigT struct {
	Flags EtmChanFlags
}

type EtmChanFlags struct {
	Unused [8]uint8
}
