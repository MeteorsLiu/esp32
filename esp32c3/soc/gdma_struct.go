package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GdmaDevS struct {
	Intr [3]struct {
		Raw struct {
			Val c.Uint32T
		}
		St struct {
			Val c.Uint32T
		}
		Ena struct {
			Val c.Uint32T
		}
		Clr struct {
			Val c.Uint32T
		}
	}
	Reserved30 c.Uint32T
	Reserved34 c.Uint32T
	Reserved38 c.Uint32T
	Reserved3c c.Uint32T
	AhbTest    struct {
		Val c.Uint32T
	}
	MiscConf struct {
		Val c.Uint32T
	}
	Date       c.Uint32T
	Reserved4c c.Uint32T
	Reserved50 c.Uint32T
	Reserved54 c.Uint32T
	Reserved58 c.Uint32T
	Reserved5c c.Uint32T
	Reserved60 c.Uint32T
	Reserved64 c.Uint32T
	Reserved68 c.Uint32T
	Reserved6c c.Uint32T
	Channel    [3]struct {
		In struct {
			InConf0 struct {
				Val c.Uint32T
			}
			InConf1 struct {
				Val c.Uint32T
			}
			InfifoStatus struct {
				Val c.Uint32T
			}
			InPop struct {
				Val c.Uint32T
			}
			InLink struct {
				Val c.Uint32T
			}
			InState struct {
				Val c.Uint32T
			}
			InSucEofDesAddr c.Uint32T
			InErrEofDesAddr c.Uint32T
			InDscr          c.Uint32T
			InDscrBf0       c.Uint32T
			InDscrBf1       c.Uint32T
			InPri           struct {
				Val c.Uint32T
			}
			InPeriSel struct {
				Val c.Uint32T
			}
		}
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
		ReservedCc c.Uint32T
		Out        struct {
			OutConf0 struct {
				Val c.Uint32T
			}
			OutConf1 struct {
				Val c.Uint32T
			}
			OutfifoStatus struct {
				Val c.Uint32T
			}
			OutPush struct {
				Val c.Uint32T
			}
			OutLink struct {
				Val c.Uint32T
			}
			OutState struct {
				Val c.Uint32T
			}
			OutEofDesAddr    c.Uint32T
			OutEofBfrDesAddr c.Uint32T
			OutDscr          c.Uint32T
			OutDscrBf0       c.Uint32T
			OutDscrBf1       c.Uint32T
			OutPri           struct {
				Val c.Uint32T
			}
			OutPeriSel struct {
				Val c.Uint32T
			}
		}
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
	}
}
type GdmaDevT GdmaDevS
