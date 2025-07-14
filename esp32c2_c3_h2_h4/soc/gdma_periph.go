package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GdmaSignalConnT struct {
	Groups [1]struct {
		Module PeriphModuleT
		Pairs  [3]struct {
			RxIrqId c.Int
			TxIrqId c.Int
		}
	}
}
