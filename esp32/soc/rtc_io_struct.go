package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RtcIoDevS struct {
	Out struct {
		Val c.Uint32T
	}
	OutW1ts struct {
		Val c.Uint32T
	}
	OutW1tc struct {
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
	Status struct {
		Val c.Uint32T
	}
	StatusW1ts struct {
		Val c.Uint32T
	}
	StatusW1tc struct {
		Val c.Uint32T
	}
	InVal struct {
		Val c.Uint32T
	}
	Pin [18]struct {
		Val c.Uint32T
	}
	DebugSel struct {
		Val c.Uint32T
	}
	DigPadHold c.Uint32T
	HallSens   struct {
		Val c.Uint32T
	}
	SensorPads struct {
		Val c.Uint32T
	}
	AdcPad struct {
		Val c.Uint32T
	}
	PadDac [2]struct {
		Val c.Uint32T
	}
	Xtal32kPad struct {
		Val c.Uint32T
	}
	TouchCfg struct {
		Val c.Uint32T
	}
	TouchPad [10]struct {
		Val c.Uint32T
	}
	ExtWakeup0 struct {
		Val c.Uint32T
	}
	XtlExtCtr struct {
		Val c.Uint32T
	}
	SarI2cIo struct {
		Val c.Uint32T
	}
	Date struct {
		Val c.Uint32T
	}
}
type RtcIoDevT RtcIoDevS
