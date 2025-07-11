package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SpiDevS struct {
	Cmd struct {
		Val c.Uint32T
	}
	Addr c.Uint32T
	Ctrl struct {
		Val c.Uint32T
	}
	Ctrl1 struct {
		Val c.Uint32T
	}
	RdStatus struct {
		Val c.Uint32T
	}
	Ctrl2 struct {
		Val c.Uint32T
	}
	Clock struct {
		Val c.Uint32T
	}
	User struct {
		Val c.Uint32T
	}
	User1 struct {
		Val c.Uint32T
	}
	User2 struct {
		Val c.Uint32T
	}
	MosiDlen struct {
		Val c.Uint32T
	}
	MisoDlen struct {
		Val c.Uint32T
	}
	SlvWrStatus c.Uint32T
	Pin         struct {
		Val c.Uint32T
	}
	Slave struct {
		Val c.Uint32T
	}
	Slave1 struct {
		Val c.Uint32T
	}
	Slave2 struct {
		Val c.Uint32T
	}
	Slave3 struct {
		Val c.Uint32T
	}
	SlvWrbufDlen struct {
		Val c.Uint32T
	}
	SlvRdbufDlen struct {
		Val c.Uint32T
	}
	CacheFctrl struct {
		Val c.Uint32T
	}
	CacheSctrl struct {
		Val c.Uint32T
	}
	SramCmd struct {
		Val c.Uint32T
	}
	SramDrdCmd struct {
		Val c.Uint32T
	}
	SramDwrCmd struct {
		Val c.Uint32T
	}
	SlvRdBit struct {
		Val c.Uint32T
	}
	Reserved68 c.Uint32T
	Reserved6c c.Uint32T
	Reserved70 c.Uint32T
	Reserved74 c.Uint32T
	Reserved78 c.Uint32T
	Reserved7c c.Uint32T
	DataBuf    [16]c.Uint32T
	TxCrc      c.Uint32T
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
	Ext0       struct {
		Val c.Uint32T
	}
	Ext1 struct {
		Val c.Uint32T
	}
	Ext2 struct {
		Val c.Uint32T
	}
	Ext3 struct {
		Val c.Uint32T
	}
	DmaConf struct {
		Val c.Uint32T
	}
	DmaOutLink struct {
		Val c.Uint32T
	}
	DmaInLink struct {
		Val c.Uint32T
	}
	DmaStatus struct {
		Val c.Uint32T
	}
	DmaIntEna struct {
		Val c.Uint32T
	}
	DmaIntRaw struct {
		Val c.Uint32T
	}
	DmaIntSt struct {
		Val c.Uint32T
	}
	DmaIntClr struct {
		Val c.Uint32T
	}
	DmaInErrEofDesAddr  c.Uint32T
	DmaInSucEofDesAddr  c.Uint32T
	DmaInlinkDscr       c.Uint32T
	DmaInlinkDscrBf0    c.Uint32T
	DmaInlinkDscrBf1    c.Uint32T
	DmaOutEofBfrDesAddr c.Uint32T
	DmaOutEofDesAddr    c.Uint32T
	DmaOutlinkDscr      c.Uint32T
	DmaOutlinkDscrBf0   c.Uint32T
	DmaOutlinkDscrBf1   c.Uint32T
	DmaRxStatus         c.Uint32T
	DmaTxStatus         c.Uint32T
	Reserved150         [171]c.Uint32T
	Date                struct {
		Val c.Uint32T
	}
}
type SpiDevT SpiDevS
