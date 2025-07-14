package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
Stores a bunch of per-ledc-peripheral data.
*/
type LedcSignalConnT struct {
	SigOut0Idx c.Uint8T
}

type LedcSubRegRetentionInfoT struct {
	RegdmaEntryArray *RegdmaEntriesConfigT
	ArraySize        c.Uint32T
}

type LedcRegRetentionInfoT struct {
	Common   LedcSubRegRetentionInfoT
	Timer    [4]LedcSubRegRetentionInfoT
	Channel  [6]LedcSubRegRetentionInfoT
	ModuleId PeriphRetentionModuleT
}
