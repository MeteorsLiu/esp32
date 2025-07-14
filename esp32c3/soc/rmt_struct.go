package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RmtDevT struct {
	DataCh [4]c.Uint32T
	TxConf [2]struct {
		Val c.Uint32T
	}
	RxConf [2]struct {
		Conf0 struct {
			Val c.Uint32T
		}
		Conf1 struct {
			Val c.Uint32T
		}
	}
	TxStatus [2]struct {
		Val c.Uint32T
	}
	RxStatus [2]struct {
		Val c.Uint32T
	}
	IntRaw struct {
		Val c.Uint32T
	}
	IntSt struct {
		Val c.Uint32T
	}
	IntEna struct {
		Val c.Uint32T
	}
	IntClr struct {
		Val c.Uint32T
	}
	TxCarrier [2]struct {
		Val c.Uint32T
	}
	RxCarrier [2]struct {
		Val c.Uint32T
	}
	TxLim [2]struct {
		Val c.Uint32T
	}
	RxLim [2]struct {
		Val c.Uint32T
	}
	SysConf struct {
		Val c.Uint32T
	}
	TxSim struct {
		Val c.Uint32T
	}
	RefCntRst struct {
		Val c.Uint32T
	}
	Reserved74 c.Uint32T
	Reserved78 c.Uint32T
	Reserved7c c.Uint32T
	Reserved80 c.Uint32T
	Reserved84 c.Uint32T
	Reserved88 c.Uint32T
	Reserved8c c.Uint32T
	Reserved90 c.Uint32T
	Reserved94 c.Uint32T
	Reserved98 c.Uint32T
	Reserved9c c.Uint32T
	ReservedA0 c.Uint32T
	ReservedA4 c.Uint32T
	ReservedA8 c.Uint32T
	ReservedAc c.Uint32T
	ReservedB0 c.Uint32T
	ReservedB4 c.Uint32T
	ReservedB8 c.Uint32T
	ReservedBc c.Uint32T
	ReservedC0 c.Uint32T
	ReservedC4 c.Uint32T
	ReservedC8 c.Uint32T
	Date       struct {
		Val c.Uint32T
	}
}
