package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LpVadT c.Uint32T

type VadUnitCtxT struct {
	Unused [8]uint8
}
type VadUnitHandleT *VadUnitCtxT

/**
 * @brief LP VAD configurations
 */

type LpVadConfigT struct {
	InitFrameNum           c.Int
	MinEnergyThresh        c.Int
	SkipBandEnergyThresh   bool
	SpeakActivityThresh    c.Int
	NonSpeakActivityThresh c.Int
	MinSpeakActivityThresh c.Int
	MaxSpeakActivityThresh c.Int
}

/**
 * @brief LP VAD Init Configurations
 */

type LpVadInitConfigT struct {
	LpI2sChan LpI2sChanHandleT
	VadConfig LpVadConfigT
}
