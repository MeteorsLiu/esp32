package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GpioSdDevT struct {
	Channel [4]struct {
		Val c.Uint32T
	}
	Reserved10 c.Uint32T
	Reserved14 c.Uint32T
	Reserved18 c.Uint32T
	Reserved1c c.Uint32T
	Cg         struct {
		Val c.Uint32T
	}
	Misc struct {
		Val c.Uint32T
	}
	Version struct {
		Val c.Uint32T
	}
}
