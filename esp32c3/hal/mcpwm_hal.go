package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type McpwmDevT struct {
	Unused [8]uint8
}
type McpwmSocHandleT *McpwmDevT

/**
 * @brief HAL layer configuration
 */

type McpwmHalInitConfigT struct {
	GroupId c.Int
}

/**
 * Context that should be maintained by both the driver and the HAL
 */

type McpwmHalContextT struct {
	Dev McpwmSocHandleT
}
