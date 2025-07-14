package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EtmRegRetentionInfoT struct {
	Module           PeriphRetentionModuleT
	RegdmaEntryArray *RegdmaEntriesConfigT
	ArraySize        c.Uint32T
}
