package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LedcDevS struct {
	ChannelGroup [2]struct {
		Channel [8]struct {
			Conf0 struct {
				Val c.Uint32T
			}
			Hpoint struct {
				Val c.Uint32T
			}
			Duty struct {
				Val c.Uint32T
			}
			Conf1 struct {
				Val c.Uint32T
			}
			DutyRd struct {
				Val c.Uint32T
			}
		}
	}
	TimerGroup [2]struct {
		Timer [4]struct {
			Conf struct {
				Val c.Uint32T
			}
			Value struct {
				Val c.Uint32T
			}
		}
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
	Conf struct {
		Val c.Uint32T
	}
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
	Reserved1f8 c.Uint32T
	Date        c.Uint32T
}
type LedcDevT LedcDevS
