package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * SPDX-FileCopyrightText: 2023 Espressif Systems (Shanghai) CO LTD
 *
 *  SPDX-License-Identifier: Apache-2.0
 */

type ModemLpconTestConfRegT struct {
	Val c.Uint32T
}

type ModemLpconCoexLpClkConfRegT struct {
	Val c.Uint32T
}

type ModemLpconClkConfRegT struct {
	Val c.Uint32T
}

type ModemLpconClkConfForceOnRegT struct {
	Val c.Uint32T
}

type ModemLpconTickConfRegT struct {
	Val c.Uint32T
}

type ModemLpconRstConfRegT struct {
	Val c.Uint32T
}

type ModemLpconMemConfRegT struct {
	Val c.Uint32T
}

type ModemLpconDateRegT struct {
	Val c.Uint32T
}

type ModemLpconDevT struct {
	TestConf       ModemLpconTestConfRegT
	CoexLpClkConf  ModemLpconCoexLpClkConfRegT
	ClkConf        ModemLpconClkConfRegT
	ClkConfForceOn ModemLpconClkConfForceOnRegT
	TickConf       ModemLpconTickConfRegT
	RstConf        ModemLpconRstConfRegT
	MemConf        ModemLpconMemConfRegT
	Date           ModemLpconDateRegT
}
