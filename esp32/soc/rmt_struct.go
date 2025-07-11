package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RmtDevT struct {
	DataCh [8]c.Uint32T
	ConfCh [8]struct {
		Conf0 struct {
			Val c.Uint32T
		}
		Conf1 struct {
			Val c.Uint32T
		}
	}
	StatusCh     [8]c.Uint32T
	ApbMemAddrCh [8]c.Uint32T
	IntRaw       struct {
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
	CarrierDutyCh [8]struct {
		Val c.Uint32T
	}
	TxLimCh [8]struct {
		Val c.Uint32T
	}
	ApbConf struct {
		Val c.Uint32T
	}
	ReservedF4 c.Uint32T
	ReservedF8 c.Uint32T
	Date       c.Uint32T
}
