package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type JpegDevT struct {
	Unused [8]uint8
}
type JpegSocHandleT *JpegDevT

/**
 * Context that should be maintained by both the driver and the HAL
 */

type JpegHalContextT struct {
	Dev JpegSocHandleT
}

// llgo:type C
type JpegConfigDhtTableT func(*JpegHalContextT, *c.Uint8T, *c.Uint8T, *c.Uint32T)

// llgo:type C
type JpegConfigFrameInfoT func(JpegSocHandleT, c.Uint8T, c.Uint8T, c.Uint8T, c.Uint8T)

// llgo:type C
type JpegConfigQuantizationCoefficientT func(JpegSocHandleT, *c.Uint32T)
