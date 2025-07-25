package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type McpwmSignalConnT struct {
	Groups [1]struct {
		Module    PeriphModuleT
		IrqId     c.Int
		Operators [3]struct {
			Generators [2]struct {
				PwmSig c.Uint32T
			}
		}
		GpioFaults [3]struct {
			FaultSig c.Uint32T
		}
		Captures [3]struct {
			CapSig c.Uint32T
		}
		GpioSynchros [3]struct {
			SyncSig c.Uint32T
		}
	}
}

type McpwmRegRetentionInfoT struct {
	RegdmaEntryArray *RegdmaEntriesConfigT
	ArraySize        c.Uint32T
	RetentionModule  PeriphRetentionModuleT
}
