package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SpiMemDevS struct {
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
	RdStatus struct {
		Val c.Uint32T
	}
	Reserved30 c.Uint32T
	Misc       struct {
		Val c.Uint32T
	}
	TxCrc      c.Uint32T
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
	SramClk struct {
		Val c.Uint32T
	}
	Fsm struct {
		Val c.Uint32T
	}
	DataBuf        [16]c.Uint32T
	FlashWaitiCtrl struct {
		Val c.Uint32T
	}
	FlashSusCtrl struct {
		Val c.Uint32T
	}
	FlashSusCmd struct {
		Val c.Uint32T
	}
	SusStatus struct {
		Val c.Uint32T
	}
	ReservedA8 c.Uint32T
	ReservedAc c.Uint32T
	ReservedB0 c.Uint32T
	ReservedB4 c.Uint32T
	ReservedB8 c.Uint32T
	ReservedBc c.Uint32T
	IntEna     struct {
		Val c.Uint32T
	}
	IntClr struct {
		Val c.Uint32T
	}
	IntRaw struct {
		Val c.Uint32T
	}
	IntSt struct {
		Val c.Uint32T
	}
	ReservedD0 c.Uint32T
	Ddr        struct {
		Val c.Uint32T
	}
	SpiSmemDdr struct {
		Val c.Uint32T
	}
	ReservedDc      c.Uint32T
	ReservedE0      c.Uint32T
	ReservedE4      c.Uint32T
	ReservedE8      c.Uint32T
	ReservedEc      c.Uint32T
	ReservedF0      c.Uint32T
	ReservedF4      c.Uint32T
	ReservedF8      c.Uint32T
	ReservedFc      c.Uint32T
	SpiFmemPms0Attr struct {
		Val c.Uint32T
	}
	SpiFmemPms1Attr struct {
		Val c.Uint32T
	}
	SpiFmemPms2Attr struct {
		Val c.Uint32T
	}
	SpiFmemPms3Attr struct {
		Val c.Uint32T
	}
	SpiFmemPms0Addr struct {
		Val c.Uint32T
	}
	SpiFmemPms1Addr struct {
		Val c.Uint32T
	}
	SpiFmemPms2Addr struct {
		Val c.Uint32T
	}
	SpiFmemPms3Addr struct {
		Val c.Uint32T
	}
	SpiFmemPms0Size struct {
		Val c.Uint32T
	}
	SpiFmemPms1Size struct {
		Val c.Uint32T
	}
	SpiFmemPms2Size struct {
		Val c.Uint32T
	}
	SpiFmemPms3Size struct {
		Val c.Uint32T
	}
	SpiSmemPms0Attr struct {
		Val c.Uint32T
	}
	SpiSmemPms1Attr struct {
		Val c.Uint32T
	}
	SpiSmemPms2Attr struct {
		Val c.Uint32T
	}
	SpiSmemPms3Attr struct {
		Val c.Uint32T
	}
	SpiSmemPms0Addr struct {
		Val c.Uint32T
	}
	SpiSmemPms1Addr struct {
		Val c.Uint32T
	}
	SpiSmemPms2Addr struct {
		Val c.Uint32T
	}
	SpiSmemPms3Addr struct {
		Val c.Uint32T
	}
	SpiSmemPms0Size struct {
		Val c.Uint32T
	}
	SpiSmemPms1Size struct {
		Val c.Uint32T
	}
	SpiSmemPms2Size struct {
		Val c.Uint32T
	}
	SpiSmemPms3Size struct {
		Val c.Uint32T
	}
	Reserved160 c.Uint32T
	PmsReject   struct {
		Val c.Uint32T
	}
	EccCtrl struct {
		Val c.Uint32T
	}
	EccErrAddr struct {
		Val c.Uint32T
	}
	AxiErrAddr struct {
		Val c.Uint32T
	}
	SpiSmemEccCtrl struct {
		Val c.Uint32T
	}
	Reserved178 c.Uint32T
	Reserved17c c.Uint32T
	TimingCali  struct {
		Val c.Uint32T
	}
	DinMode struct {
		Val c.Uint32T
	}
	DinNum struct {
		Val c.Uint32T
	}
	DoutMode struct {
		Val c.Uint32T
	}
	SpiSmemTimingCali struct {
		Val c.Uint32T
	}
	SpiSmemDinMode struct {
		Val c.Uint32T
	}
	SpiSmemDinNum struct {
		Val c.Uint32T
	}
	SpiSmemDoutMode struct {
		Val c.Uint32T
	}
	SpiSmemAc struct {
		Val c.Uint32T
	}
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
	Reserved1f8 c.Uint32T
	Reserved1fc c.Uint32T
	ClockGate   struct {
		Val c.Uint32T
	}
	Reserved204  c.Uint32T
	Reserved208  c.Uint32T
	Reserved20c  c.Uint32T
	Reserved210  c.Uint32T
	Reserved214  c.Uint32T
	Reserved218  c.Uint32T
	Reserved21c  c.Uint32T
	Reserved220  c.Uint32T
	Reserved224  c.Uint32T
	Reserved228  c.Uint32T
	Reserved22c  c.Uint32T
	Reserved230  c.Uint32T
	Reserved234  c.Uint32T
	Reserved238  c.Uint32T
	Reserved23c  c.Uint32T
	Reserved240  c.Uint32T
	Reserved244  c.Uint32T
	Reserved248  c.Uint32T
	Reserved24c  c.Uint32T
	Reserved250  c.Uint32T
	Reserved254  c.Uint32T
	Reserved258  c.Uint32T
	Reserved25c  c.Uint32T
	Reserved260  c.Uint32T
	Reserved264  c.Uint32T
	Reserved268  c.Uint32T
	Reserved26c  c.Uint32T
	Reserved270  c.Uint32T
	Reserved274  c.Uint32T
	Reserved278  c.Uint32T
	Reserved27c  c.Uint32T
	Reserved280  c.Uint32T
	Reserved284  c.Uint32T
	Reserved288  c.Uint32T
	Reserved28c  c.Uint32T
	Reserved290  c.Uint32T
	Reserved294  c.Uint32T
	Reserved298  c.Uint32T
	Reserved29c  c.Uint32T
	Reserved2a0  c.Uint32T
	Reserved2a4  c.Uint32T
	Reserved2a8  c.Uint32T
	Reserved2ac  c.Uint32T
	Reserved2b0  c.Uint32T
	Reserved2b4  c.Uint32T
	Reserved2b8  c.Uint32T
	Reserved2bc  c.Uint32T
	Reserved2c0  c.Uint32T
	Reserved2c4  c.Uint32T
	Reserved2c8  c.Uint32T
	Reserved2cc  c.Uint32T
	Reserved2d0  c.Uint32T
	Reserved2d4  c.Uint32T
	Reserved2d8  c.Uint32T
	Reserved2dc  c.Uint32T
	Reserved2e0  c.Uint32T
	Reserved2e4  c.Uint32T
	Reserved2e8  c.Uint32T
	Reserved2ec  c.Uint32T
	Reserved2f0  c.Uint32T
	Reserved2f4  c.Uint32T
	Reserved2f8  c.Uint32T
	Reserved2fc  c.Uint32T
	XtsPlainBase c.Uint32T
	Reserved304  c.Uint32T
	Reserved308  c.Uint32T
	Reserved30c  c.Uint32T
	Reserved310  c.Uint32T
	Reserved314  c.Uint32T
	Reserved318  c.Uint32T
	Reserved31c  c.Uint32T
	Reserved320  c.Uint32T
	Reserved324  c.Uint32T
	Reserved328  c.Uint32T
	Reserved32c  c.Uint32T
	Reserved330  c.Uint32T
	Reserved334  c.Uint32T
	Reserved338  c.Uint32T
	Reserved33c  c.Uint32T
	XtsLinesize  struct {
		Val c.Uint32T
	}
	XtsDestination struct {
		Val c.Uint32T
	}
	XtsPhysicalAddress struct {
		Val c.Uint32T
	}
	XtsTrigger struct {
		Val c.Uint32T
	}
	XtsRelease struct {
		Val c.Uint32T
	}
	XtsDestroy struct {
		Val c.Uint32T
	}
	XtsState struct {
		Val c.Uint32T
	}
	XtsDate struct {
		Val c.Uint32T
	}
	Reserved360    c.Uint32T
	Reserved364    c.Uint32T
	Reserved368    c.Uint32T
	Reserved36c    c.Uint32T
	Reserved370    c.Uint32T
	Reserved374    c.Uint32T
	Reserved378    c.Uint32T
	MmuItemContent c.Uint32T
	MmuItemIndex   c.Uint32T
	MmuPowerCtrl   struct {
		Val c.Uint32T
	}
	DpaCtrl struct {
		Val c.Uint32T
	}
	XtsPseudoRoundConf struct {
		Val c.Uint32T
	}
	Reserved390           c.Uint32T
	Reserved394           c.Uint32T
	Reserved398           c.Uint32T
	Reserved39c           c.Uint32T
	Reserved3a0           c.Uint32T
	Reserved3a4           c.Uint32T
	Reserved3a8           c.Uint32T
	Reserved3ac           c.Uint32T
	Reserved3b0           c.Uint32T
	Reserved3b4           c.Uint32T
	Reserved3b8           c.Uint32T
	Reserved3bc           c.Uint32T
	Reserved3c0           c.Uint32T
	Reserved3c4           c.Uint32T
	Reserved3c8           c.Uint32T
	Reserved3cc           c.Uint32T
	Reserved3d0           c.Uint32T
	Reserved3d4           c.Uint32T
	Reserved3d8           c.Uint32T
	Reserved3dc           c.Uint32T
	Reserved3e0           c.Uint32T
	Reserved3e4           c.Uint32T
	Reserved3e8           c.Uint32T
	Reserved3ec           c.Uint32T
	SpiMemisterrndEcoHigh c.Uint32T
	SpiMemisterrndEcoLow  c.Uint32T
	Reserved3f8           c.Uint32T
	Date                  struct {
		Val c.Uint32T
	}
}
type SpiMemDevT SpiMemDevS
