package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * SPDX-FileCopyrightText: 2019-2021 Espressif Systems (Shanghai) CO LTD
 *
 * SPDX-License-Identifier: Apache-2.0
 */

type LedcDevS struct {
	ChannelGroup [1]struct {
		Channel [6]struct {
			Conf0 struct {
				Val c.Uint32T
			}
			Hpoint struct {
				Val c.Uint32T
			}
			Duty struct {
				Val c.Uint32T
			}
			Conf1 struct {
				Val c.Uint32T
			}
			DutyRd struct {
				Val c.Uint32T
			}
		}
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
	TimerGroup [1]struct {
		Timer [4]struct {
			Conf struct {
				Val c.Uint32T
			}
			Value struct {
				Val c.Uint32T
			}
		}
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
	Conf struct {
		Val c.Uint32T
	}
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
type LedcDevT LedcDevS
