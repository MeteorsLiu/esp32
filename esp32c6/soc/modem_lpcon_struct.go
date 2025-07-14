package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * SPDX-FileCopyrightText: 2022 Espressif Systems (Shanghai) CO LTD
 *
 *  SPDX-License-Identifier: Apache-2.0
 */

type ModemLpconTestConfRegT struct {
	Val c.Uint32T
}

type ModemLpconLpTimerConfRegT struct {
	Val c.Uint32T
}

type ModemLpconCoexLpClkConfRegT struct {
	Val c.Uint32T
}

type ModemLpconWifiLpClkConfRegT struct {
	Val c.Uint32T
}

type ModemLpconI2cMstClkConfRegT struct {
	Val c.Uint32T
}

type ModemLpconModem32kClkConfRegT struct {
	Val c.Uint32T
}

type ModemLpconClkConfRegT struct {
	Val c.Uint32T
}

type ModemLpconClkConfForceOnRegT struct {
	Val c.Uint32T
}

type ModemLpconClkConfPowerStRegT struct {
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
	TestConf        ModemLpconTestConfRegT
	LpTimerConf     ModemLpconLpTimerConfRegT
	CoexLpClkConf   ModemLpconCoexLpClkConfRegT
	WifiLpClkConf   ModemLpconWifiLpClkConfRegT
	I2cMstClkConf   ModemLpconI2cMstClkConfRegT
	Modem32kClkConf ModemLpconModem32kClkConfRegT
	ClkConf         ModemLpconClkConfRegT
	ClkConfForceOn  ModemLpconClkConfForceOnRegT
	ClkConfPowerSt  ModemLpconClkConfPowerStRegT
	RstConf         ModemLpconRstConfRegT
	MemConf         ModemLpconMemConfRegT
	Date            ModemLpconDateRegT
}
