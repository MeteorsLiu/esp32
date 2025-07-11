package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EmacMacDevS struct {
	Gmacconfig struct {
		Val c.Uint32T
	}
	Gmacff struct {
		Val c.Uint32T
	}
	Reserved1008 c.Uint32T
	Reserved100c c.Uint32T
	Emacgmiiaddr struct {
		Val c.Uint32T
	}
	Emacmiidata struct {
		Val c.Uint32T
	}
	Gmacfc struct {
		Val c.Uint32T
	}
	Reserved101c c.Uint32T
	Reserved1020 c.Uint32T
	Emacdebug    struct {
		Val c.Uint32T
	}
	PmtRwuffr c.Uint32T
	PmtCsr    struct {
		Val c.Uint32T
	}
	GmaclpiCrs struct {
		Val c.Uint32T
	}
	Gmaclpitimerscontrol struct {
		Val c.Uint32T
	}
	Emacints struct {
		Val c.Uint32T
	}
	Emacintmask struct {
		Val c.Uint32T
	}
	Emacaddr0high struct {
		Val c.Uint32T
	}
	Emacaddr0low  c.Uint32T
	Emacaddr1high struct {
		Val c.Uint32T
	}
	Emacaddr1low  c.Uint32T
	Emacaddr2high struct {
		Val c.Uint32T
	}
	Emacaddr2low  c.Uint32T
	Emacaddr3high struct {
		Val c.Uint32T
	}
	Emacaddr3low  c.Uint32T
	Emacaddr4high struct {
		Val c.Uint32T
	}
	Emacaddr4low  c.Uint32T
	Emacaddr5high struct {
		Val c.Uint32T
	}
	Emacaddr5low  c.Uint32T
	Emacaddr6high struct {
		Val c.Uint32T
	}
	Emacaddr6low  c.Uint32T
	Emacaddr7high struct {
		Val c.Uint32T
	}
	Emacaddr7low c.Uint32T
	Reserved1080 c.Uint32T
	Reserved1084 c.Uint32T
	Reserved1088 c.Uint32T
	Reserved108c c.Uint32T
	Reserved1090 c.Uint32T
	Reserved1094 c.Uint32T
	Reserved1098 c.Uint32T
	Reserved109c c.Uint32T
	Reserved10a0 c.Uint32T
	Reserved10a4 c.Uint32T
	Reserved10a8 c.Uint32T
	Reserved10ac c.Uint32T
	Reserved10b0 c.Uint32T
	Reserved10b4 c.Uint32T
	Reserved10b8 c.Uint32T
	Reserved10bc c.Uint32T
	Reserved10c0 c.Uint32T
	Reserved10c4 c.Uint32T
	Reserved10c8 c.Uint32T
	Reserved10cc c.Uint32T
	Reserved10d0 c.Uint32T
	Reserved10d4 c.Uint32T
	Emaccstatus  struct {
		Val c.Uint32T
	}
	Emacwdogto struct {
		Val c.Uint32T
	}
}
type EmacMacDevT EmacMacDevS
