package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2cDevS struct {
	SclLowPeriod struct {
		Val c.Uint32T
	}
	Ctr struct {
		Val c.Uint32T
	}
	Sr struct {
		Val c.Uint32T
	}
	Timeout struct {
		Val c.Uint32T
	}
	SlaveAddr struct {
		Val c.Uint32T
	}
	FifoSt struct {
		Val c.Uint32T
	}
	FifoConf struct {
		Val c.Uint32T
	}
	FifoData struct {
		Val c.Uint32T
	}
	IntRaw struct {
		Val c.Uint32T
	}
	IntClr struct {
		Val c.Uint32T
	}
	IntEna struct {
		Val c.Uint32T
	}
	IntStatus struct {
		Val c.Uint32T
	}
	SdaHold struct {
		Val c.Uint32T
	}
	SdaSample struct {
		Val c.Uint32T
	}
	SclHighPeriod struct {
		Val c.Uint32T
	}
	Reserved3c   c.Uint32T
	SclStartHold struct {
		Val c.Uint32T
	}
	SclRstartSetup struct {
		Val c.Uint32T
	}
	SclStopHold struct {
		Val c.Uint32T
	}
	SclStopSetup struct {
		Val c.Uint32T
	}
	FilterCfg struct {
		Val c.Uint32T
	}
	ClkConf struct {
		Val c.Uint32T
	}
	Command [8]struct {
		Val c.Uint32T
	}
	SclStTimeOut struct {
		Val c.Uint32T
	}
	SclMainStTimeOut struct {
		Val c.Uint32T
	}
	SclSpConf struct {
		Val c.Uint32T
	}
	SclStretchConf struct {
		Val c.Uint32T
	}
	Reserved88 c.Uint32T
	Reserved8c c.Uint32T
	Reserved90 c.Uint32T
	Reserved94 c.Uint32T
	Reserved98 c.Uint32T
	Reserved9c c.Uint32T
	ReservedA0 c.Uint32T
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
	Date       c.Uint32T
	ReservedFc c.Uint32T
	TxfifoMem  [32]c.Uint32T
	RxfifoMem  [32]c.Uint32T
}
type I2cDevT I2cDevS
