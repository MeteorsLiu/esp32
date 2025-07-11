package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SysconDevS struct {
	ClkConf struct {
		Val c.Uint32T
	}
	XtalTickConf struct {
		Val c.Uint32T
	}
	PllTickConf struct {
		Val c.Uint32T
	}
	Ck8mTickConf struct {
		Val c.Uint32T
	}
	SaradcCtrl struct {
		Val c.Uint32T
	}
	SaradcCtrl2 struct {
		Val c.Uint32T
	}
	SaradcFsm struct {
		Val c.Uint32T
	}
	SaradcSar1PattTab [4]c.Uint32T
	SaradcSar2PattTab [4]c.Uint32T
	ApllTickConf      struct {
		Val c.Uint32T
	}
	Reserved40 c.Uint32T
	Reserved44 c.Uint32T
	Reserved48 c.Uint32T
	Reserved4c c.Uint32T
	Reserved50 c.Uint32T
	Reserved54 c.Uint32T
	Reserved58 c.Uint32T
	Reserved5c c.Uint32T
	Reserved60 c.Uint32T
	Reserved64 c.Uint32T
	Reserved68 c.Uint32T
	Reserved6c c.Uint32T
	Reserved70 c.Uint32T
	Reserved74 c.Uint32T
	Reserved78 c.Uint32T
	Date       c.Uint32T
}
type SysconDevT SysconDevS
