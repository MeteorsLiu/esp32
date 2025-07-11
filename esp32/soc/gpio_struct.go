package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GpioDevS struct {
	BtSelect c.Uint32T
	Out      c.Uint32T
	OutW1ts  c.Uint32T
	OutW1tc  c.Uint32T
	Out1     struct {
		Val c.Uint32T
	}
	Out1W1ts struct {
		Val c.Uint32T
	}
	Out1W1tc struct {
		Val c.Uint32T
	}
	SdioSelect struct {
		Val c.Uint32T
	}
	Enable     c.Uint32T
	EnableW1ts c.Uint32T
	EnableW1tc c.Uint32T
	Enable1    struct {
		Val c.Uint32T
	}
	Enable1W1ts struct {
		Val c.Uint32T
	}
	Enable1W1tc struct {
		Val c.Uint32T
	}
	Strap struct {
		Val c.Uint32T
	}
	In  c.Uint32T
	In1 struct {
		Val c.Uint32T
	}
	Status     c.Uint32T
	StatusW1ts c.Uint32T
	StatusW1tc c.Uint32T
	Status1    struct {
		Val c.Uint32T
	}
	Status1W1ts struct {
		Val c.Uint32T
	}
	Status1W1tc struct {
		Val c.Uint32T
	}
	Reserved5c c.Uint32T
	AcpuInt    c.Uint32T
	AcpuNmiInt c.Uint32T
	PcpuInt    c.Uint32T
	PcpuNmiInt c.Uint32T
	CpusdioInt c.Uint32T
	AcpuInt1   struct {
		Val c.Uint32T
	}
	AcpuNmiInt1 struct {
		Val c.Uint32T
	}
	PcpuInt1 struct {
		Val c.Uint32T
	}
	PcpuNmiInt1 struct {
		Val c.Uint32T
	}
	CpusdioInt1 struct {
		Val c.Uint32T
	}
	Pin [40]struct {
		Val c.Uint32T
	}
	CaliConf struct {
		Val c.Uint32T
	}
	CaliData struct {
		Val c.Uint32T
	}
	FuncInSelCfg [256]struct {
		Val c.Uint32T
	}
	FuncOutSelCfg [40]struct {
		Val c.Uint32T
	}
}
type GpioDevT GpioDevS
