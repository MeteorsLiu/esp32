package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GpioDevS struct {
	BtSelect c.Uint32T
	Out      struct {
		Val c.Uint32T
	}
	OutW1ts struct {
		Val c.Uint32T
	}
	OutW1tc struct {
		Val c.Uint32T
	}
	Reserved10 c.Uint32T
	Reserved14 c.Uint32T
	Reserved18 c.Uint32T
	SdioSelect struct {
		Val c.Uint32T
	}
	Enable struct {
		Val c.Uint32T
	}
	EnableW1ts struct {
		Val c.Uint32T
	}
	EnableW1tc struct {
		Val c.Uint32T
	}
	Reserved2c c.Uint32T
	Reserved30 c.Uint32T
	Reserved34 c.Uint32T
	Strap      struct {
		Val c.Uint32T
	}
	In struct {
		Val c.Uint32T
	}
	Reserved40 c.Uint32T
	Status     struct {
		Val c.Uint32T
	}
	StatusW1ts struct {
		Val c.Uint32T
	}
	StatusW1tc struct {
		Val c.Uint32T
	}
	Reserved50 c.Uint32T
	Reserved54 c.Uint32T
	Reserved58 c.Uint32T
	PcpuInt    struct {
		Val c.Uint32T
	}
	PcpuNmiInt struct {
		Val c.Uint32T
	}
	CpusdioInt struct {
		Val c.Uint32T
	}
	Reserved68 c.Uint32T
	Reserved6c c.Uint32T
	Reserved70 c.Uint32T
	Pin        [26]struct {
		Val c.Uint32T
	}
	ReservedDc  c.Uint32T
	ReservedE0  c.Uint32T
	ReservedE4  c.Uint32T
	ReservedE8  c.Uint32T
	ReservedEc  c.Uint32T
	ReservedF0  c.Uint32T
	ReservedF4  c.Uint32T
	ReservedF8  c.Uint32T
	ReservedFc  c.Uint32T
	Reserved100 c.Uint32T
	Reserved104 c.Uint32T
	Reserved108 c.Uint32T
	Reserved10c c.Uint32T
	Reserved110 c.Uint32T
	Reserved114 c.Uint32T
	Reserved118 c.Uint32T
	Reserved11c c.Uint32T
	Reserved120 c.Uint32T
	Reserved124 c.Uint32T
	Reserved128 c.Uint32T
	Reserved12c c.Uint32T
	Reserved130 c.Uint32T
	Reserved134 c.Uint32T
	Reserved138 c.Uint32T
	Reserved13c c.Uint32T
	Reserved140 c.Uint32T
	Reserved144 c.Uint32T
	Reserved148 c.Uint32T
	StatusNext  struct {
		Val c.Uint32T
	}
	Reserved150  c.Uint32T
	FuncInSelCfg [128]struct {
		Val c.Uint32T
	}
	Reserved354   c.Uint32T
	Reserved358   c.Uint32T
	Reserved35c   c.Uint32T
	Reserved360   c.Uint32T
	Reserved364   c.Uint32T
	Reserved368   c.Uint32T
	Reserved36c   c.Uint32T
	Reserved370   c.Uint32T
	Reserved374   c.Uint32T
	Reserved378   c.Uint32T
	Reserved37c   c.Uint32T
	Reserved380   c.Uint32T
	Reserved384   c.Uint32T
	Reserved388   c.Uint32T
	Reserved38c   c.Uint32T
	Reserved390   c.Uint32T
	Reserved394   c.Uint32T
	Reserved398   c.Uint32T
	Reserved39c   c.Uint32T
	Reserved3a0   c.Uint32T
	Reserved3a4   c.Uint32T
	Reserved3a8   c.Uint32T
	Reserved3ac   c.Uint32T
	Reserved3b0   c.Uint32T
	Reserved3b4   c.Uint32T
	Reserved3b8   c.Uint32T
	Reserved3bc   c.Uint32T
	Reserved3c0   c.Uint32T
	Reserved3c4   c.Uint32T
	Reserved3c8   c.Uint32T
	Reserved3cc   c.Uint32T
	Reserved3d0   c.Uint32T
	Reserved3d4   c.Uint32T
	Reserved3d8   c.Uint32T
	Reserved3dc   c.Uint32T
	Reserved3e0   c.Uint32T
	Reserved3e4   c.Uint32T
	Reserved3e8   c.Uint32T
	Reserved3ec   c.Uint32T
	Reserved3f0   c.Uint32T
	Reserved3f4   c.Uint32T
	Reserved3f8   c.Uint32T
	Reserved3fc   c.Uint32T
	Reserved400   c.Uint32T
	Reserved404   c.Uint32T
	Reserved408   c.Uint32T
	Reserved40c   c.Uint32T
	Reserved410   c.Uint32T
	Reserved414   c.Uint32T
	Reserved418   c.Uint32T
	Reserved41c   c.Uint32T
	Reserved420   c.Uint32T
	Reserved424   c.Uint32T
	Reserved428   c.Uint32T
	Reserved42c   c.Uint32T
	Reserved430   c.Uint32T
	Reserved434   c.Uint32T
	Reserved438   c.Uint32T
	Reserved43c   c.Uint32T
	Reserved440   c.Uint32T
	Reserved444   c.Uint32T
	Reserved448   c.Uint32T
	Reserved44c   c.Uint32T
	Reserved450   c.Uint32T
	Reserved454   c.Uint32T
	Reserved458   c.Uint32T
	Reserved45c   c.Uint32T
	Reserved460   c.Uint32T
	Reserved464   c.Uint32T
	Reserved468   c.Uint32T
	Reserved46c   c.Uint32T
	Reserved470   c.Uint32T
	Reserved474   c.Uint32T
	Reserved478   c.Uint32T
	Reserved47c   c.Uint32T
	Reserved480   c.Uint32T
	Reserved484   c.Uint32T
	Reserved488   c.Uint32T
	Reserved48c   c.Uint32T
	Reserved490   c.Uint32T
	Reserved494   c.Uint32T
	Reserved498   c.Uint32T
	Reserved49c   c.Uint32T
	Reserved4a0   c.Uint32T
	Reserved4a4   c.Uint32T
	Reserved4a8   c.Uint32T
	Reserved4ac   c.Uint32T
	Reserved4b0   c.Uint32T
	Reserved4b4   c.Uint32T
	Reserved4b8   c.Uint32T
	Reserved4bc   c.Uint32T
	Reserved4c0   c.Uint32T
	Reserved4c4   c.Uint32T
	Reserved4c8   c.Uint32T
	Reserved4cc   c.Uint32T
	Reserved4d0   c.Uint32T
	Reserved4d4   c.Uint32T
	Reserved4d8   c.Uint32T
	Reserved4dc   c.Uint32T
	Reserved4e0   c.Uint32T
	Reserved4e4   c.Uint32T
	Reserved4e8   c.Uint32T
	Reserved4ec   c.Uint32T
	Reserved4f0   c.Uint32T
	Reserved4f4   c.Uint32T
	Reserved4f8   c.Uint32T
	Reserved4fc   c.Uint32T
	Reserved500   c.Uint32T
	Reserved504   c.Uint32T
	Reserved508   c.Uint32T
	Reserved50c   c.Uint32T
	Reserved510   c.Uint32T
	Reserved514   c.Uint32T
	Reserved518   c.Uint32T
	Reserved51c   c.Uint32T
	Reserved520   c.Uint32T
	Reserved524   c.Uint32T
	Reserved528   c.Uint32T
	Reserved52c   c.Uint32T
	Reserved530   c.Uint32T
	Reserved534   c.Uint32T
	Reserved538   c.Uint32T
	Reserved53c   c.Uint32T
	Reserved540   c.Uint32T
	Reserved544   c.Uint32T
	Reserved548   c.Uint32T
	Reserved54c   c.Uint32T
	Reserved550   c.Uint32T
	FuncOutSelCfg [26]struct {
		Val c.Uint32T
	}
	Reserved5bc c.Uint32T
	Reserved5c0 c.Uint32T
	Reserved5c4 c.Uint32T
	Reserved5c8 c.Uint32T
	Reserved5cc c.Uint32T
	Reserved5d0 c.Uint32T
	Reserved5d4 c.Uint32T
	Reserved5d8 c.Uint32T
	Reserved5dc c.Uint32T
	Reserved5e0 c.Uint32T
	Reserved5e4 c.Uint32T
	Reserved5e8 c.Uint32T
	Reserved5ec c.Uint32T
	Reserved5f0 c.Uint32T
	Reserved5f4 c.Uint32T
	Reserved5f8 c.Uint32T
	Reserved5fc c.Uint32T
	Reserved600 c.Uint32T
	Reserved604 c.Uint32T
	Reserved608 c.Uint32T
	Reserved60c c.Uint32T
	Reserved610 c.Uint32T
	Reserved614 c.Uint32T
	Reserved618 c.Uint32T
	Reserved61c c.Uint32T
	Reserved620 c.Uint32T
	Reserved624 c.Uint32T
	Reserved628 c.Uint32T
	ClockGate   struct {
		Val c.Uint32T
	}
	Reserved630 c.Uint32T
	Reserved634 c.Uint32T
	Reserved638 c.Uint32T
	Reserved63c c.Uint32T
	Reserved640 c.Uint32T
	Reserved644 c.Uint32T
	Reserved648 c.Uint32T
	Reserved64c c.Uint32T
	Reserved650 c.Uint32T
	Reserved654 c.Uint32T
	Reserved658 c.Uint32T
	Reserved65c c.Uint32T
	Reserved660 c.Uint32T
	Reserved664 c.Uint32T
	Reserved668 c.Uint32T
	Reserved66c c.Uint32T
	Reserved670 c.Uint32T
	Reserved674 c.Uint32T
	Reserved678 c.Uint32T
	Reserved67c c.Uint32T
	Reserved680 c.Uint32T
	Reserved684 c.Uint32T
	Reserved688 c.Uint32T
	Reserved68c c.Uint32T
	Reserved690 c.Uint32T
	Reserved694 c.Uint32T
	Reserved698 c.Uint32T
	Reserved69c c.Uint32T
	Reserved6a0 c.Uint32T
	Reserved6a4 c.Uint32T
	Reserved6a8 c.Uint32T
	Reserved6ac c.Uint32T
	Reserved6b0 c.Uint32T
	Reserved6b4 c.Uint32T
	Reserved6b8 c.Uint32T
	Reserved6bc c.Uint32T
	Reserved6c0 c.Uint32T
	Reserved6c4 c.Uint32T
	Reserved6c8 c.Uint32T
	Reserved6cc c.Uint32T
	Reserved6d0 c.Uint32T
	Reserved6d4 c.Uint32T
	Reserved6d8 c.Uint32T
	Reserved6dc c.Uint32T
	Reserved6e0 c.Uint32T
	Reserved6e4 c.Uint32T
	Reserved6e8 c.Uint32T
	Reserved6ec c.Uint32T
	Reserved6f0 c.Uint32T
	Reserved6f4 c.Uint32T
	Reserved6f8 c.Uint32T
	Date        struct {
		Val c.Uint32T
	}
}
type GpioDevT GpioDevS
