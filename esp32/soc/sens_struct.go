package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SensDevS struct {
	SarReadCtrl struct {
		Val c.Uint32T
	}
	SarReadStatus1 c.Uint32T
	SarMeasWait1   struct {
		Val c.Uint32T
	}
	SarMeasWait2 struct {
		Val c.Uint32T
	}
	SarMeasCtrl struct {
		Val c.Uint32T
	}
	SarReadStatus2 c.Uint32T
	UlpCpSleepCyc0 c.Uint32T
	UlpCpSleepCyc1 c.Uint32T
	UlpCpSleepCyc2 c.Uint32T
	UlpCpSleepCyc3 c.Uint32T
	UlpCpSleepCyc4 c.Uint32T
	SarStartForce  struct {
		Val c.Uint32T
	}
	SarMemWrCtrl struct {
		Val c.Uint32T
	}
	SarAtten1     c.Uint32T
	SarAtten2     c.Uint32T
	SarSlaveAddr1 struct {
		Val c.Uint32T
	}
	SarSlaveAddr2 struct {
		Val c.Uint32T
	}
	SarSlaveAddr3 struct {
		Val c.Uint32T
	}
	SarSlaveAddr4 struct {
		Val c.Uint32T
	}
	SarTctrl struct {
		Val c.Uint32T
	}
	SarI2cCtrl struct {
		Val c.Uint32T
	}
	SarMeasStart1 struct {
		Val c.Uint32T
	}
	SarTouchCtrl1 struct {
		Val c.Uint32T
	}
	TouchThresh [5]struct {
		Val c.Uint32T
	}
	TouchMeas [5]struct {
		Val c.Uint32T
	}
	SarTouchCtrl2 struct {
		Val c.Uint32T
	}
	Reserved88     c.Uint32T
	SarTouchEnable struct {
		Val c.Uint32T
	}
	SarReadCtrl2 struct {
		Val c.Uint32T
	}
	SarMeasStart2 struct {
		Val c.Uint32T
	}
	SarDacCtrl1 struct {
		Val c.Uint32T
	}
	SarDacCtrl2 struct {
		Val c.Uint32T
	}
	SarMeasCtrl2 struct {
		Val c.Uint32T
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
	SarNouse   c.Uint32T
	Sardate    struct {
		Val c.Uint32T
	}
}
type SensDevT SensDevS
