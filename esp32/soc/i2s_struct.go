package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2sDevS struct {
	FifoWr c.Uint32T
	FifoRd c.Uint32T
	Conf   struct {
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
	Timing struct {
		Val c.Uint32T
	}
	FifoConf struct {
		Val c.Uint32T
	}
	RxEofNum       c.Uint32T
	ConfSingleData c.Uint32T
	ConfChan       struct {
		Val c.Uint32T
	}
	OutLink struct {
		Val c.Uint32T
	}
	InLink struct {
		Val c.Uint32T
	}
	OutEofDesAddr    c.Uint32T
	InEofDesAddr     c.Uint32T
	OutEofBfrDesAddr c.Uint32T
	AhbTest          struct {
		Val c.Uint32T
	}
	InLinkDscr     c.Uint32T
	InLinkDscrBf0  c.Uint32T
	InLinkDscrBf1  c.Uint32T
	OutLinkDscr    c.Uint32T
	OutLinkDscrBf0 c.Uint32T
	OutLinkDscrBf1 c.Uint32T
	LcConf         struct {
		Val c.Uint32T
	}
	OutFifoPush struct {
		Val c.Uint32T
	}
	InFifoPop struct {
		Val c.Uint32T
	}
	LcState0   c.Uint32T
	LcState1   c.Uint32T
	LcHungConf struct {
		Val c.Uint32T
	}
	Reserved78 c.Uint32T
	Reserved7c c.Uint32T
	CvsdConf0  struct {
		Val c.Uint32T
	}
	CvsdConf1 struct {
		Val c.Uint32T
	}
	CvsdConf2 struct {
		Val c.Uint32T
	}
	PlcConf0 struct {
		Val c.Uint32T
	}
	PlcConf1 struct {
		Val c.Uint32T
	}
	PlcConf2 struct {
		Val c.Uint32T
	}
	EscoConf0 struct {
		Val c.Uint32T
	}
	ScoConf0 struct {
		Val c.Uint32T
	}
	Conf1 struct {
		Val c.Uint32T
	}
	PdConf struct {
		Val c.Uint32T
	}
	Conf2 struct {
		Val c.Uint32T
	}
	ClkmConf struct {
		Val c.Uint32T
	}
	SampleRateConf struct {
		Val c.Uint32T
	}
	PdmConf struct {
		Val c.Uint32T
	}
	PdmFreqConf struct {
		Val c.Uint32T
	}
	State struct {
		Val c.Uint32T
	}
	ReservedC0 c.Uint32T
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
type I2sDevT I2sDevS
