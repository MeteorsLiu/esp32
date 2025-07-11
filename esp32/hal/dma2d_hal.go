package hal

import _ "unsafe"

type Dma2dDevT struct {
	Unused [8]uint8
}
type Dma2dSocHandleT *Dma2dDevT

/**
 * Context that should be maintained by both the driver and the HAL
 */

type Dma2dHalContextT struct {
	Dev Dma2dSocHandleT
}
