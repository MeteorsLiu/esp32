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

type ModemSysconTestConfRegT struct {
	Val c.Uint32T
}

type ModemSysconClkConfRegT struct {
	Val c.Uint32T
}

type ModemSysconClkConfForceOnRegT struct {
	Val c.Uint32T
}

type ModemSysconModemRstConfRegT struct {
	Val c.Uint32T
}

type ModemSysconClkConf1RegT struct {
	Val c.Uint32T
}

type ModemSysconClkConf1ForceOnRegT struct {
	Val c.Uint32T
}

type ModemSysconMemConfRegT struct {
	Val c.Uint32T
}

type ModemSysconDateRegT struct {
	Val c.Uint32T
}

type ModemSysconDevT struct {
	TestConf        ModemSysconTestConfRegT
	ClkConf         ModemSysconClkConfRegT
	ClkConfForceOn  ModemSysconClkConfForceOnRegT
	ModemRstConf    ModemSysconModemRstConfRegT
	ClkConf1        ModemSysconClkConf1RegT
	ClkConf1ForceOn ModemSysconClkConf1ForceOnRegT
	MemConf         ModemSysconMemConfRegT
	Date            ModemSysconDateRegT
}
