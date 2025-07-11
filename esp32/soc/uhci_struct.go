package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UhciDevS struct {
	Conf0 struct {
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
	DmaOutStatus struct {
		Val c.Uint32T
	}
	DmaOutPush struct {
		Val c.Uint32T
	}
	DmaInStatus struct {
		Val c.Uint32T
	}
	DmaInPop struct {
		Val c.Uint32T
	}
	DmaOutLink struct {
		Val c.Uint32T
	}
	DmaInLink struct {
		Val c.Uint32T
	}
	Conf1 struct {
		Val c.Uint32T
	}
	State0              c.Uint32T
	State1              c.Uint32T
	DmaOutEofDesAddr    c.Uint32T
	DmaInSucEofDesAddr  c.Uint32T
	DmaInErrEofDesAddr  c.Uint32T
	DmaOutEofBfrDesAddr c.Uint32T
	AhbTest             struct {
		Val c.Uint32T
	}
	DmaInDscr     c.Uint32T
	DmaInDscrBf0  c.Uint32T
	DmaInDscrBf1  c.Uint32T
	DmaOutDscr    c.Uint32T
	DmaOutDscrBf0 c.Uint32T
	DmaOutDscrBf1 c.Uint32T
	EscapeConf    struct {
		Val c.Uint32T
	}
	HungConf struct {
		Val c.Uint32T
	}
	AckNum    c.Uint32T
	RxHead    c.Uint32T
	QuickSent struct {
		Val c.Uint32T
	}
	QData [7]struct {
		WData [2]c.Uint32T
	}
	EscConf0 struct {
		Val c.Uint32T
	}
	EscConf1 struct {
		Val c.Uint32T
	}
	EscConf2 struct {
		Val c.Uint32T
	}
	EscConf3 struct {
		Val c.Uint32T
	}
	PktThres struct {
		Val c.Uint32T
	}
	ReservedC4 c.Uint32T
	ReservedC8 c.Uint32T
	ReservedCc c.Uint32T
	ReservedD0 c.Uint32T
	ReservedD4 c.Uint32T
	ReservedD8 c.Uint32T
	ReservedDc c.Uint32T
	ReservedE0 c.Uint32T
	ReservedE4 c.Uint32T
	ReservedE8 c.Uint32T
	ReservedEc c.Uint32T
	ReservedF0 c.Uint32T
	ReservedF4 c.Uint32T
	ReservedF8 c.Uint32T
	Date       c.Uint32T
}
type UhciDevT UhciDevS
