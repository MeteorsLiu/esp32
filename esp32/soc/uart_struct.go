package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UartDevS struct {
	Fifo struct {
		Val c.Uint32T
	}
	IntRaw struct {
		Val c.Uint32T
	}
	IntSt struct {
		Val c.Uint32T
	}
	IntEna struct {
		Val c.Uint32T
	}
	IntClr struct {
		Val c.Uint32T
	}
	ClkDiv struct {
		Val c.Uint32T
	}
	AutoBaud struct {
		Val c.Uint32T
	}
	Status struct {
		Val c.Uint32T
	}
	Conf0 struct {
		Val c.Uint32T
	}
	Conf1 struct {
		Val c.Uint32T
	}
	Lowpulse struct {
		Val c.Uint32T
	}
	Highpulse struct {
		Val c.Uint32T
	}
	RxdCnt struct {
		Val c.Uint32T
	}
	FlowConf struct {
		Val c.Uint32T
	}
	SleepConf struct {
		Val c.Uint32T
	}
	SwfcConf struct {
		Val c.Uint32T
	}
	IdleConf struct {
		Val c.Uint32T
	}
	Rs485Conf struct {
		Val c.Uint32T
	}
	AtCmdPrecnt struct {
		Val c.Uint32T
	}
	AtCmdPostcnt struct {
		Val c.Uint32T
	}
	AtCmdGaptout struct {
		Val c.Uint32T
	}
	AtCmdChar struct {
		Val c.Uint32T
	}
	MemConf struct {
		Val c.Uint32T
	}
	MemTxStatus struct {
		Val c.Uint32T
	}
	MemRxStatus struct {
		Val c.Uint32T
	}
	MemCntStatus struct {
		Val c.Uint32T
	}
	Pospulse struct {
		Val c.Uint32T
	}
	Negpulse struct {
		Val c.Uint32T
	}
	Reserved70 c.Uint32T
	Reserved74 c.Uint32T
	Date       c.Uint32T
	Id         c.Uint32T
}
type UartDevT UartDevS
