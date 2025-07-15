package esp_hw_support

import _ "unsafe"

type DebugProbeUnitT struct {
	Unused [8]uint8
}
type DebugProbeUnitHandleT *DebugProbeUnitT

type DebugProbeChannelT struct {
	Unused [8]uint8
}
type DebugProbeChannelHandleT *DebugProbeChannelT

/**
 * @brief Configuration for a debug probe unit
 */

type DebugProbeUnitConfigT struct {
	ProbeOutGpioNums [16]GpioNumT
}

/**
 * @brief Configuration for a debug probe channel
 */

type DebugProbeChannelConfigT struct {
	TargetModule DebugProbeTargetT
}
