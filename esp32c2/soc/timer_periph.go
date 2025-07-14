package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TimerGroupSignalConnT struct {
	Groups [1]struct {
		Module     PeriphModuleT
		TimerIrqId [1]c.Int
	}
}
