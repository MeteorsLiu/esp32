package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PcntDevT struct {
	ConfUnit [8]struct {
		Conf0 struct {
			Val c.Uint32T
		}
		Conf1 struct {
			Val c.Uint32T
		}
		Conf2 struct {
			Val c.Uint32T
		}
	}
	CntUnit [8]struct {
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
	StatusUnit [8]struct {
		Val c.Uint32T
	}
	Ctrl struct {
		Val c.Uint32T
	}
	ReservedB4 c.Uint32T
	ReservedB8 c.Uint32T
	ReservedBc c.Uint32T
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
