package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LcdI2sSignalConnT struct {
	Buses [2]struct {
		Module   PeriphModuleT
		IrqId    c.Int
		DataSigs [24]c.Int
		WrSig    c.Int
	}
}
