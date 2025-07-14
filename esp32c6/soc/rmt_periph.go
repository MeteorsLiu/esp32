package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RmtSignalConnT struct {
	Groups [1]struct {
		Irq      c.Int
		Channels [4]struct {
		}
	}
}

type RmtRegRetentionInfoT struct {
	Module           PeriphRetentionModuleT
	RegdmaEntryArray *RegdmaEntriesConfigT
	ArraySize        c.Uint32T
}
