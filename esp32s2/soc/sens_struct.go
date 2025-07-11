package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SensDevS struct {
	SarReader1Ctrl struct {
		Val c.Uint32T
	}
	SarReader1Status c.Uint32T
	SarMeas1Ctrl1    struct {
		Val c.Uint32T
	}
	SarMeas1Ctrl2 struct {
		Val c.Uint32T
	}
	SarMeas1Mux struct {
		Val c.Uint32T
	}
	SarAtten1   c.Uint32T
	SarAmpCtrl1 struct {
		Val c.Uint32T
	}
	SarAmpCtrl2 struct {
		Val c.Uint32T
	}
	SarAmpCtrl3 struct {
		Val c.Uint32T
	}
	SarReader2Ctrl struct {
		Val c.Uint32T
	}
	SarReader2Status c.Uint32T
	SarMeas2Ctrl1    struct {
		Val c.Uint32T
	}
	SarMeas2Ctrl2 struct {
		Val c.Uint32T
	}
	SarMeas2Mux struct {
		Val c.Uint32T
	}
	SarAtten2      c.Uint32T
	SarPowerXpdSar struct {
		Val c.Uint32T
	}
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
	SarTctrl2 struct {
		Val c.Uint32T
	}
	SarI2cCtrl struct {
		Val c.Uint32T
	}
	SarTouchConf struct {
		Val c.Uint32T
	}
	TouchThresh [14]struct {
		Val c.Uint32T
	}
	Reserved98    c.Uint32T
	Reserved9c    c.Uint32T
	ReservedA0    c.Uint32T
	ReservedA4    c.Uint32T
	ReservedA8    c.Uint32T
	ReservedAc    c.Uint32T
	ReservedB0    c.Uint32T
	ReservedB4    c.Uint32T
	ReservedB8    c.Uint32T
	ReservedBc    c.Uint32T
	ReservedC0    c.Uint32T
	ReservedC4    c.Uint32T
	ReservedC8    c.Uint32T
	ReservedCc    c.Uint32T
	ReservedD0    c.Uint32T
	SarTouchChnSt struct {
		Val c.Uint32T
	}
	SarTouchStatus0 struct {
		Val c.Uint32T
	}
	SarTouchStatus [14]struct {
		Val c.Uint32T
	}
	SarTouchSlpStatus struct {
		Val c.Uint32T
	}
	SarTouchApprStatus struct {
		Val c.Uint32T
	}
	SarDacCtrl1 struct {
		Val c.Uint32T
	}
	SarDacCtrl2 struct {
		Val c.Uint32T
	}
	SarCocpuState struct {
		Val c.Uint32T
	}
	SarCocpuIntRaw struct {
		Val c.Uint32T
	}
	SarCocpuIntEna struct {
		Val c.Uint32T
	}
	SarCocpuIntSt struct {
		Val c.Uint32T
	}
	SarCocpuIntClr struct {
		Val c.Uint32T
	}
	SarCocpuDebug struct {
		Val c.Uint32T
	}
	SarHallCtrl struct {
		Val c.Uint32T
	}
	SarNouse     c.Uint32T
	SarIoMuxConf struct {
		Val c.Uint32T
	}
	Sardate struct {
		Val c.Uint32T
	}
}
type SensDevT SensDevS
