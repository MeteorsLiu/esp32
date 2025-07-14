package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * SPDX-FileCopyrightText: 2022-2023 Espressif Systems (Shanghai) CO LTD
 *
 *  SPDX-License-Identifier: Apache-2.0
 */

type LpTimerTargetRegT struct {
	Lo struct {
		Val c.Uint32T
	}
	Hi struct {
		Val c.Uint32T
	}
}

type LpTimerUpdateRegT struct {
	Val c.Uint32T
}

type LpTimerCounterRegT struct {
	Lo struct {
		Val c.Uint32T
	}
	Hi struct {
		Val c.Uint32T
	}
}

type LpTimerOverflowRegT struct {
	Val c.Uint32T
}

type LpTimerIntrRegT struct {
	Val c.Uint32T
}

type LpTimerLpIntrRegT struct {
	Val c.Uint32T
}

type LpTimerDateClkenRegT struct {
	Val c.Uint32T
}

type LpTimerDevT struct {
	Target    [2]LpTimerTargetRegT
	Update    LpTimerUpdateRegT
	Counter   [2]LpTimerCounterRegT
	Overflow  LpTimerOverflowRegT
	IntRaw    LpTimerIntrRegT
	IntSt     LpTimerIntrRegT
	IntEn     LpTimerIntrRegT
	IntClr    LpTimerIntrRegT
	LpIntRaw  LpTimerLpIntrRegT
	LpIntSt   LpTimerLpIntrRegT
	LpIntEn   LpTimerLpIntrRegT
	LpIntClr  LpTimerLpIntrRegT
	Reserved  [237]c.Uint32T
	DateClken LpTimerDateClkenRegT
}
