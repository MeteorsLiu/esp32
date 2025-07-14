package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * ECDSA peripheral config structure
 */

type EcdsaHalConfigT struct {
	Mode        EcdsaModeT
	Curve       EcdsaCurveT
	ShaMode     EcdsaShaModeT
	EfuseKeyBlk c.Int
	UseKmKey    bool
	SignType    EcdsaSignTypeT
	LoopNumber  c.Uint16T
}
