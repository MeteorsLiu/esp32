package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PmuHpDigPowerRegT struct {
	Val c.Uint32T
}

type PmuHpIcgModemRegT struct {
	Val c.Uint32T
}

type PmuHpSysCntlRegT struct {
	Val c.Uint32T
}

type PmuHpClkPowerRegT struct {
	Val c.Uint32T
}

type PmuHpBiasRegT struct {
	Val c.Uint32T
}

type PmuHpBackupRegT struct {
	Val c.Uint32T
}

type PmuHpSysclkRegT struct {
	Val c.Uint32T
}

type PmuHpRegulator0RegT struct {
	Val c.Uint32T
}

type PmuHpRegulator1RegT struct {
	Val c.Uint32T
}

type PmuHpXtalRegT struct {
	Val c.Uint32T
}

type PmuHpHwRegmapT struct {
	DigPower   PmuHpDigPowerRegT
	IcgFunc    c.Uint32T
	IcgApb     c.Uint32T
	IcgModem   PmuHpIcgModemRegT
	Syscntl    PmuHpSysCntlRegT
	ClkPower   PmuHpClkPowerRegT
	Bias       PmuHpBiasRegT
	Backup     PmuHpBackupRegT
	BackupClk  c.Uint32T
	Sysclk     PmuHpSysclkRegT
	Regulator0 PmuHpRegulator0RegT
	Regulator1 PmuHpRegulator1RegT
	Xtal       PmuHpXtalRegT
}

/** */

type PmuLpRegulator0RegT struct {
	Val c.Uint32T
}

type PmuLpRegulator1RegT struct {
	Val c.Uint32T
}

type PmuLpXtalRegT struct {
	Val c.Uint32T
}

type PmuLpDigPowerRegT struct {
	Val c.Uint32T
}

type PmuLpClkPowerRegT struct {
	Val c.Uint32T
}

type PmuLpBiasRegT struct {
	Val c.Uint32T
}

type PmuLpHwRegmapT struct {
	Regulator0 PmuLpRegulator0RegT
	Regulator1 PmuLpRegulator1RegT
	Xtal       PmuLpXtalRegT
	DigPower   PmuLpDigPowerRegT
	ClkPower   PmuLpClkPowerRegT
	Bias       PmuLpBiasRegT
}

type PmuImmHpClkPowerRegT struct {
	Val c.Uint32T
}

type PmuImmSleepSysclkRegT struct {
	Val c.Uint32T
}

type PmuImmHpFuncIcgRegT struct {
	Val c.Uint32T
}

type PmuImmHpApbIcgRegT struct {
	Val c.Uint32T
}

type PmuImmModemIcgRegT struct {
	Val c.Uint32T
}

type PmuImmLpIcgRegT struct {
	Val c.Uint32T
}

type PmuImmPadHoldAllRegT struct {
	Val c.Uint32T
}

type PmuImmI2cIsolateRegT struct {
	Val c.Uint32T
}

type PmuImmHwRegmapT struct {
	ClkPower    PmuImmHpClkPowerRegT
	SleepSysclk PmuImmSleepSysclkRegT
	HpFuncIcg   PmuImmHpFuncIcgRegT
	HpApbIcg    PmuImmHpApbIcgRegT
	ModemIcg    PmuImmModemIcgRegT
	LpIcg       PmuImmLpIcgRegT
	PadHoldAll  PmuImmPadHoldAllRegT
	I2cIso      PmuImmI2cIsolateRegT
}

type PmuPowerWaitTimer0RegT struct {
	Val c.Uint32T
}

type PmuPowerWaitTimer1RegT struct {
	Val c.Uint32T
}

type PmuPowerDomainCntlRegT struct {
	Val c.Uint32T
}

type PmuPowerMemoryCntlRegT struct {
	Val c.Uint32T
}

type PmuPowerMemoryMaskRegT struct {
	Val c.Uint32T
}

type PmuPowerHpPadRegT struct {
	Val c.Uint32T
}

type PmuPowerVddSpiCntlRegT struct {
	Val c.Uint32T
}

type PmuPowerClkWaitCntlRegT struct {
	Val c.Uint32T
}

type PmuPowerHwRegmapT struct {
	WaitTimer0 PmuPowerWaitTimer0RegT
	WaitTimer1 PmuPowerWaitTimer1RegT
	HpPd       [5]PmuPowerDomainCntlRegT
	LpPeri     PmuPowerDomainCntlRegT
	MemCntl    PmuPowerMemoryCntlRegT
	MemMask    PmuPowerMemoryMaskRegT
	HpPad      PmuPowerHpPadRegT
	VddSpi     PmuPowerVddSpiCntlRegT
	ClkWait    PmuPowerClkWaitCntlRegT
}

type PmuSlpWakeupCntl0RegT struct {
	Val c.Uint32T
}

type PmuSlpWakeupCntl1RegT struct {
	Val c.Uint32T
}

type PmuSlpWakeupCntl3RegT struct {
	Val c.Uint32T
}

type PmuSlpWakeupCntl4RegT struct {
	Val c.Uint32T
}

type PmuSlpWakeupCntl5RegT struct {
	Val c.Uint32T
}

type PmuSlpWakeupCntl6RegT struct {
	Val c.Uint32T
}

type PmuSlpWakeupCntl7RegT struct {
	Val c.Uint32T
}

type PmuWakeupHwRegmapT struct {
	Cntl0   PmuSlpWakeupCntl0RegT
	Cntl1   PmuSlpWakeupCntl1RegT
	Cntl2   c.Uint32T
	Cntl3   PmuSlpWakeupCntl3RegT
	Cntl4   PmuSlpWakeupCntl4RegT
	Cntl5   PmuSlpWakeupCntl5RegT
	Cntl6   PmuSlpWakeupCntl6RegT
	Cntl7   PmuSlpWakeupCntl7RegT
	Status0 c.Uint32T
	Status1 c.Uint32T
}

type PmuHpClkPoweronRegT struct {
	Val c.Uint32T
}

type PmuHpClkCntlRegT struct {
	Val c.Uint32T
}

type PmuPorStatusRegT struct {
	Val c.Uint32T
}

type PmuRfPwcRegT struct {
	Val c.Uint32T
}

type PmuVddbatCfgRegT struct {
	Val c.Uint32T
}

type PmuBackupCfgRegT struct {
	Val c.Uint32T
}

type PmuHpIntrRegT struct {
	Val c.Uint32T
}

type PmuHpExtHwRegmapT struct {
	ClkPoweron PmuHpClkPoweronRegT
	ClkCntl    PmuHpClkCntlRegT
	PorStatus  PmuPorStatusRegT
	RfPwc      PmuRfPwcRegT
	VddbatCfg  PmuVddbatCfgRegT
	BackupCfg  PmuBackupCfgRegT
	IntRaw     PmuHpIntrRegT
	IntSt      PmuHpIntrRegT
	IntEna     PmuHpIntrRegT
	IntClr     PmuHpIntrRegT
}

type PmuLpIntrRegT struct {
	Val c.Uint32T
}

type PmuLpCpuPwr0RegT struct {
	Val c.Uint32T
}

type PmuLpCpuPwr1RegT struct {
	Val c.Uint32T
}

type PmuLpExtHwRegmapT struct {
	IntRaw PmuLpIntrRegT
	IntSt  PmuLpIntrRegT
	IntEna PmuLpIntrRegT
	IntClr PmuLpIntrRegT
	Pwr0   PmuLpCpuPwr0RegT
	Pwr1   PmuLpCpuPwr1RegT
}

type PmuHpLpHwRegmapT struct {
	Common struct {
	}
}

type PmuDevT struct {
	HpSys       [3]PmuHpHwRegmapT
	LpSys       [2]PmuLpHwRegmapT
	Imm         PmuImmHwRegmapT
	Power       PmuPowerHwRegmapT
	Wakeup      PmuWakeupHwRegmapT
	HpExt       PmuHpExtHwRegmapT
	LpExt       PmuLpExtHwRegmapT
	HpLpCpuComm struct {
		Val c.Uint32T
	}
	HpRegulatorCfg struct {
		Val c.Uint32T
	}
	MainState struct {
		Val c.Uint32T
	}
	PwrState struct {
		Val c.Uint32T
	}
	ClkState0 struct {
		Val c.Uint32T
	}
	ClkState1    c.Uint32T
	ClkState2    c.Uint32T
	VddSpiStatus struct {
		Val c.Uint32T
	}
	Reserved [149]c.Uint32T
	Date     struct {
		Val c.Uint32T
	}
}
