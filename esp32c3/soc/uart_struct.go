package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * SPDX-FileCopyrightText: 2020-2024 Espressif Systems (Shanghai) CO LTD
 *
 * SPDX-License-Identifier: Apache-2.0
 */

type UartDevS struct {
	AhbFifo struct {
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
	RxFilt struct {
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
	SwfcConf0 struct {
		Val c.Uint32T
	}
	SwfcConf1 struct {
		Val c.Uint32T
	}
	TxbrkConf struct {
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
	FsmStatus struct {
		Val c.Uint32T
	}
	Pospulse struct {
		Val c.Uint32T
	}
	Negpulse struct {
		Val c.Uint32T
	}
	ClkConf struct {
		Val c.Uint32T
	}
	Date c.Uint32T
	Id   struct {
		Val c.Uint32T
	}
}
type UartDevT UartDevS
