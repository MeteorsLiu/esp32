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
	SarTouchDenoise struct {
		Val c.Uint32T
	}
	TouchThresh [14]struct {
		Val c.Uint32T
	}
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
	SarNouse           c.Uint32T
	SarPeriClkGateConf struct {
		Val c.Uint32T
	}
	SarPeriResetConf struct {
		Val c.Uint32T
	}
	SarCocpuIntEnaW1ts struct {
		Val c.Uint32T
	}
	SarCocpuIntEnaW1tc struct {
		Val c.Uint32T
	}
	SarDebugConf struct {
		Val c.Uint32T
	}
	Reserved118 c.Uint32T
	Reserved11c c.Uint32T
	Reserved120 c.Uint32T
	Reserved124 c.Uint32T
	Reserved128 c.Uint32T
	Reserved12c c.Uint32T
	Reserved130 c.Uint32T
	Reserved134 c.Uint32T
	Reserved138 c.Uint32T
	Reserved13c c.Uint32T
	Reserved140 c.Uint32T
	Reserved144 c.Uint32T
	Reserved148 c.Uint32T
	Reserved14c c.Uint32T
	Reserved150 c.Uint32T
	Reserved154 c.Uint32T
	Reserved158 c.Uint32T
	Reserved15c c.Uint32T
	Reserved160 c.Uint32T
	Reserved164 c.Uint32T
	Reserved168 c.Uint32T
	Reserved16c c.Uint32T
	Reserved170 c.Uint32T
	Reserved174 c.Uint32T
	Reserved178 c.Uint32T
	Reserved17c c.Uint32T
	Reserved180 c.Uint32T
	Reserved184 c.Uint32T
	Reserved188 c.Uint32T
	Reserved18c c.Uint32T
	Reserved190 c.Uint32T
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
	Sardate     struct {
		Val c.Uint32T
	}
}
type SensDevT SensDevS
