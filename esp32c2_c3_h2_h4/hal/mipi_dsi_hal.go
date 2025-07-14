package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DsiHostDevT struct {
	Unused [8]uint8
}
type MipiDsiHostSocHandleT *DsiHostDevT

type DsiBrgDevT struct {
	Unused [8]uint8
}
type MipiDsiBridgeSocHandleT *DsiBrgDevT

/**
 * @brief MIPI DSI HAL driver context
 */

type MipiDsiHalContextT struct {
	Host            MipiDsiHostSocHandleT
	Bridge          MipiDsiBridgeSocHandleT
	LaneBitRateMbps c.Uint32T
	DpiClockFreqMhz c.Uint32T
}

/**
 * @brief MIPI DSI HAL driver configuration
 */

type MipiDsiHalConfigT struct {
	BusId           c.Int
	LaneBitRateMbps c.Uint32T
	NumDataLanes    c.Uint8T
}
