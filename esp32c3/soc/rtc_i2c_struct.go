package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * SPDX-FileCopyrightText: 2020-2022 Espressif Systems (Shanghai) CO LTD
 *
 * SPDX-License-Identifier: Apache-2.0
 */

type RtcI2cDevS struct {
	SclLow struct {
		Val c.Uint32T
	}
	Ctrl struct {
		Val c.Uint32T
	}
	Status struct {
		Val c.Uint32T
	}
	Timeout struct {
		Val c.Uint32T
	}
	SlaveAddr struct {
		Val c.Uint32T
	}
	SclHigh struct {
		Val c.Uint32T
	}
	SdaDuty struct {
		Val c.Uint32T
	}
	SclStartPeriod struct {
		Val c.Uint32T
	}
	SclStopPeriod struct {
		Val c.Uint32T
	}
	IntClr struct {
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
	FifoData struct {
		Val c.Uint32T
	}
	Command [16]struct {
		Val c.Uint32T
	}
	Reserved78 c.Uint32T
	Reserved7c c.Uint32T
	Reserved80 c.Uint32T
	Reserved84 c.Uint32T
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
	ReservedF8 c.Uint32T
	Date       struct {
		Val c.Uint32T
	}
}
type RtcI2cDevT RtcI2cDevS
