package esp_driver_mcpwm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type McpwmTimerT struct {
	Unused [8]uint8
}
type McpwmTimerHandleT *McpwmTimerT

type McpwmOperT struct {
	Unused [8]uint8
}
type McpwmOperHandleT *McpwmOperT

type McpwmCmprT struct {
	Unused [8]uint8
}
type McpwmCmprHandleT *McpwmCmprT

type McpwmGenT struct {
	Unused [8]uint8
}
type McpwmGenHandleT *McpwmGenT

type McpwmFaultT struct {
	Unused [8]uint8
}
type McpwmFaultHandleT *McpwmFaultT

type McpwmSyncT struct {
	Unused [8]uint8
}
type McpwmSyncHandleT *McpwmSyncT

type McpwmCapTimerT struct {
	Unused [8]uint8
}
type McpwmCapTimerHandleT *McpwmCapTimerT

type McpwmCapChannelT struct {
	Unused [8]uint8
}
type McpwmCapChannelHandleT *McpwmCapChannelT

/**
 * @brief MCPWM timer event data
 */

type McpwmTimerEventDataT struct {
	CountValue c.Uint32T
	Direction  McpwmTimerDirectionT
}

// llgo:type C
type McpwmTimerEventCbT func(McpwmTimerHandleT, *McpwmTimerEventDataT, c.Pointer) bool

/**
 * @brief MCPWM brake event data
 */

type McpwmBrakeEventDataT struct {
	Unused [8]uint8
}

// llgo:type C
type McpwmBrakeEventCbT func(McpwmOperHandleT, *McpwmBrakeEventDataT, c.Pointer) bool

/**
 * @brief MCPWM fault event data
 */

type McpwmFaultEventDataT struct {
	Unused [8]uint8
}

// llgo:type C
type McpwmFaultEventCbT func(McpwmFaultHandleT, *McpwmFaultEventDataT, c.Pointer) bool

/**
 * @brief MCPWM compare event data
 */

type McpwmCompareEventDataT struct {
	CompareTicks c.Uint32T
	Direction    McpwmTimerDirectionT
}

// llgo:type C
type McpwmCompareEventCbT func(McpwmCmprHandleT, *McpwmCompareEventDataT, c.Pointer) bool

/**
 * @brief MCPWM capture event data
 */

type McpwmCaptureEventDataT struct {
	CapValue c.Uint32T
	CapEdge  McpwmCaptureEdgeT
}

// llgo:type C
type McpwmCaptureEventCbT func(McpwmCapChannelHandleT, *McpwmCaptureEventDataT, c.Pointer) bool
