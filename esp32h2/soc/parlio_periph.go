package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ParlioSignalConnT struct {
	Groups [1]struct {
		TxUnits [1]struct {
			DataSigs  [8]c.Int
			ClkOutSig c.Int
			ClkInSig  c.Int
		}
		RxUnits [1]struct {
			DataSigs  [8]c.Int
			ClkOutSig c.Int
			ClkInSig  c.Int
		}
		TxIrqId c.Int
		RxIrqId c.Int
		Module  PeriphModuleT
	}
}

type ParlioRegRetentionInfoT struct {
	RetentionModule  PeriphRetentionModuleT
	RegdmaEntryArray *RegdmaEntriesConfigT
	ArraySize        c.Uint32T
}
