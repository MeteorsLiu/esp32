package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SlcDevS struct {
	Conf0 struct {
		Val c.Uint32T
	}
	Slc0IntRaw struct {
		Val c.Uint32T
	}
	Slc0IntSt struct {
		Val c.Uint32T
	}
	Slc0IntEna struct {
		Val c.Uint32T
	}
	Slc0IntClr struct {
		Val c.Uint32T
	}
	Slc1IntRaw struct {
		Val c.Uint32T
	}
	Slc1IntSt struct {
		Val c.Uint32T
	}
	Slc1IntEna struct {
		Val c.Uint32T
	}
	Slc1IntClr struct {
		Val c.Uint32T
	}
	RxStatus struct {
		Val c.Uint32T
	}
	Slc0RxfifoPush struct {
		Val c.Uint32T
	}
	Slc1RxfifoPush struct {
		Val c.Uint32T
	}
	TxStatus struct {
		Val c.Uint32T
	}
	Slc0TxfifoPop struct {
		Val c.Uint32T
	}
	Slc1TxfifoPop struct {
		Val c.Uint32T
	}
	Slc0RxLink struct {
		Val c.Uint32T
	}
	Slc0TxLink struct {
		Val c.Uint32T
	}
	Slc1RxLink struct {
		Val c.Uint32T
	}
	Slc1TxLink struct {
		Val c.Uint32T
	}
	IntvecTohost struct {
		Val c.Uint32T
	}
	Slc0Token0 struct {
		Val c.Uint32T
	}
	Slc0Token1 struct {
		Val c.Uint32T
	}
	Slc1Token0 struct {
		Val c.Uint32T
	}
	Slc1Token1 struct {
		Val c.Uint32T
	}
	Conf1 struct {
		Val c.Uint32T
	}
	Slc0State0 c.Uint32T
	Slc0State1 c.Uint32T
	Slc1State0 c.Uint32T
	Slc1State1 c.Uint32T
	BridgeConf struct {
		Val c.Uint32T
	}
	Slc0ToEofDesAddr    c.Uint32T
	Slc0TxEofDesAddr    c.Uint32T
	Slc0ToEofBfrDesAddr c.Uint32T
	Slc1ToEofDesAddr    c.Uint32T
	Slc1TxEofDesAddr    c.Uint32T
	Slc1ToEofBfrDesAddr c.Uint32T
	AhbTest             struct {
		Val c.Uint32T
	}
	SdioSt struct {
		Val c.Uint32T
	}
	RxDscrConf struct {
		Val c.Uint32T
	}
	Slc0TxlinkDscr      c.Uint32T
	Slc0TxlinkDscrBf0   c.Uint32T
	Slc0TxlinkDscrBf1   c.Uint32T
	Slc0RxlinkDscr      c.Uint32T
	Slc0RxlinkDscrBf0   c.Uint32T
	Slc0RxlinkDscrBf1   c.Uint32T
	Slc1TxlinkDscr      c.Uint32T
	Slc1TxlinkDscrBf0   c.Uint32T
	Slc1TxlinkDscrBf1   c.Uint32T
	Slc1RxlinkDscr      c.Uint32T
	Slc1RxlinkDscrBf0   c.Uint32T
	Slc1RxlinkDscrBf1   c.Uint32T
	Slc0TxErreofDesAddr c.Uint32T
	Slc1TxErreofDesAddr c.Uint32T
	TokenLat            struct {
		Val c.Uint32T
	}
	TxDscrConf struct {
		Val c.Uint32T
	}
	CmdInfor0   c.Uint32T
	CmdInfor1   c.Uint32T
	Slc0LenConf struct {
		Val c.Uint32T
	}
	Slc0Length struct {
		Val c.Uint32T
	}
	Slc0TxpktHDscr  c.Uint32T
	Slc0TxpktEDscr  c.Uint32T
	Slc0RxpktHDscr  c.Uint32T
	Slc0RxpktEDscr  c.Uint32T
	Slc0TxpktuHDscr c.Uint32T
	Slc0TxpktuEDscr c.Uint32T
	Slc0RxpktuHDscr c.Uint32T
	Slc0RxpktuEDscr c.Uint32T
	Reserved10c     c.Uint32T
	Reserved110     c.Uint32T
	SeqPosition     struct {
		Val c.Uint32T
	}
	Slc0DscrRecConf struct {
		Val c.Uint32T
	}
	SdioCrcSt0 struct {
		Val c.Uint32T
	}
	SdioCrcSt1 struct {
		Val c.Uint32T
	}
	Slc0EofStartDes  c.Uint32T
	Slc0PushDscrAddr c.Uint32T
	Slc0DoneDscrAddr c.Uint32T
	Slc0SubStartDes  c.Uint32T
	Slc0DscrCnt      struct {
		Val c.Uint32T
	}
	Slc0LenLimConf struct {
		Val c.Uint32T
	}
	Slc0IntSt1 struct {
		Val c.Uint32T
	}
	Slc0IntEna1 struct {
		Val c.Uint32T
	}
	Slc1IntSt1 struct {
		Val c.Uint32T
	}
	Slc1IntEna1 struct {
		Val c.Uint32T
	}
	Reserved14c c.Uint32T
	Reserved150 c.Uint32T
	Reserved154 c.Uint32T
	Reserved158 c.Uint32T
	Reserved15c c.Uint32T
	Reserved160 c.Uint32T
	Reserved164 c.Uint32T
	Reserved168 c.Uint32T
	Reserved16c c.Uint32T
	Reserved170 c.Uint32T
	Reserved174 c.Uint32T
	Reserved178 c.Uint32T
	Reserved17c c.Uint32T
	Reserved180 c.Uint32T
	Reserved184 c.Uint32T
	Reserved188 c.Uint32T
	Reserved18c c.Uint32T
	Reserved190 c.Uint32T
	Reserved194 c.Uint32T
	Reserved198 c.Uint32T
	Reserved19c c.Uint32T
	Reserved1a0 c.Uint32T
	Reserved1a4 c.Uint32T
	Reserved1a8 c.Uint32T
	Reserved1ac c.Uint32T
	Reserved1b0 c.Uint32T
	Reserved1b4 c.Uint32T
	Reserved1b8 c.Uint32T
	Reserved1bc c.Uint32T
	Reserved1c0 c.Uint32T
	Reserved1c4 c.Uint32T
	Reserved1c8 c.Uint32T
	Reserved1cc c.Uint32T
	Reserved1d0 c.Uint32T
	Reserved1d4 c.Uint32T
	Reserved1d8 c.Uint32T
	Reserved1dc c.Uint32T
	Reserved1e0 c.Uint32T
	Reserved1e4 c.Uint32T
	Reserved1e8 c.Uint32T
	Reserved1ec c.Uint32T
	Reserved1f0 c.Uint32T
	Reserved1f4 c.Uint32T
	Date        c.Uint32T
	Id          c.Uint32T
}
type SlcDevT SlcDevS
