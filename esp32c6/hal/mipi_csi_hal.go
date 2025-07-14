package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type CsiHostDevT struct {
	Unused [8]uint8
}
type MipiCsiHostSocHandleT *CsiHostDevT

type CsiBrgDevT struct {
	Unused [8]uint8
}
type MipiCsiBridgeSocHandleT *CsiBrgDevT

/**
 * @brief MIPI CSI HAL driver context
 */

type MipiCsiHalContextT struct {
	HostDev   MipiCsiHostSocHandleT
	BridgeDev MipiCsiBridgeSocHandleT
}

/**
 * @brief MIPI CSI HAL driver configuration
 */

type MipiCsiHalConfigT struct {
	LanesNum        c.Uint8T
	FrameWidth      c.Uint32T
	FrameHeight     c.Uint32T
	InBpp           c.Uint32T
	OutBpp          c.Uint32T
	ByteSwapEn      bool
	LaneBitRateMbps c.Int
}
