package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Register */
/** Type of external_device_encrypt_decrypt_control register
 *  EXTERNAL DEVICE ENCRYPTION/DECRYPTION configuration register
 */

type HpSysExternalDeviceEncryptDecryptControlRegT struct {
	Val c.Uint32T
}

/** Type of sram_usage_conf register
 *  HP memory usage configuration register
 */

type HpSysSramUsageConfRegT struct {
	Val c.Uint32T
}

/** Type of sec_dpa_conf register
 *  HP anti-DPA security configuration register
 */

type HpSysSecDpaConfRegT struct {
	Val c.Uint32T
}

/** Type of rom_table_lock register
 *  Rom-Table lock register
 */

type HpSysRomTableLockRegT struct {
	Val c.Uint32T
}

/** Type of rom_table register
 *  Rom-Table register
 */

type HpSysRomTableRegT struct {
	Val c.Uint32T
}

/** Type of mem_test_conf register
 *  MEM_TEST configuration register
 */

type HpSysMemTestConfRegT struct {
	Val c.Uint32T
}

/** Type of clock_gate register
 *  HP-SYSTEM clock gating configure register
 */

type HpSysClockGateRegT struct {
	Val c.Uint32T
}

/** Group: Timeout Register */
/** Type of cpu_peri_timeout_conf register
 *  CPU_PERI_TIMEOUT configuration register
 */

type HpSysCpuPeriTimeoutConfRegT struct {
	Val c.Uint32T
}

/** Type of cpu_peri_timeout_addr register
 *  CPU_PERI_TIMEOUT_ADDR register
 */

type HpSysCpuPeriTimeoutAddrRegT struct {
	Val c.Uint32T
}

/** Type of cpu_peri_timeout_uid register
 *  CPU_PERI_TIMEOUT_UID register
 */

type HpSysCpuPeriTimeoutUidRegT struct {
	Val c.Uint32T
}

/** Type of hp_peri_timeout_conf register
 *  HP_PERI_TIMEOUT configuration register
 */

type HpSysHpPeriTimeoutConfRegT struct {
	Val c.Uint32T
}

/** Type of hp_peri_timeout_addr register
 *  HP_PERI_TIMEOUT_ADDR register
 */

type HpSysHpPeriTimeoutAddrRegT struct {
	Val c.Uint32T
}

/** Type of hp_peri_timeout_uid register
 *  HP_PERI_TIMEOUT_UID register
 */

type HpSysHpPeriTimeoutUidRegT struct {
	Val c.Uint32T
}

/** Group: Redcy ECO Registers */
/** Type of rnd_eco register
 *  redcy eco register.
 */

type HpSysRndEcoRegT struct {
	Val c.Uint32T
}

/** Type of rnd_eco_low register
 *  redcy eco low register.
 */

type HpSysRndEcoLowRegT struct {
	Val c.Uint32T
}

/** Type of rnd_eco_high register
 *  redcy eco high register.
 */

type HpSysRndEcoHighRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  Date register.
 */

type HpSysDateRegT struct {
	Val c.Uint32T
}

type HpSysDevT struct {
	ExternalDeviceEncryptDecryptControl HpSysExternalDeviceEncryptDecryptControlRegT
	SramUsageConf                       HpSysSramUsageConfRegT
	SecDpaConf                          HpSysSecDpaConfRegT
	CpuPeriTimeoutConf                  HpSysCpuPeriTimeoutConfRegT
	CpuPeriTimeoutAddr                  HpSysCpuPeriTimeoutAddrRegT
	CpuPeriTimeoutUid                   HpSysCpuPeriTimeoutUidRegT
	HpPeriTimeoutConf                   HpSysHpPeriTimeoutConfRegT
	HpPeriTimeoutAddr                   HpSysHpPeriTimeoutAddrRegT
	HpPeriTimeoutUid                    HpSysHpPeriTimeoutUidRegT
	RomTableLock                        HpSysRomTableLockRegT
	RomTable                            HpSysRomTableRegT
	MemTestConf                         HpSysMemTestConfRegT
	Reserved030                         [236]c.Uint32T
	RndEco                              HpSysRndEcoRegT
	RndEcoLow                           HpSysRndEcoLowRegT
	RndEcoHigh                          HpSysRndEcoHighRegT
	Reserved3ec                         [3]c.Uint32T
	ClockGate                           HpSysClockGateRegT
	Date                                HpSysDateRegT
}
