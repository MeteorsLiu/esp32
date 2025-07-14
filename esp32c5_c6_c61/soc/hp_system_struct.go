package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Register */
/** Type of external_device_encrypt_decrypt_control register
 *  EXTERNAL DEVICE ENCRYPTION/DECRYPTION configuration register
 */

type HpSystemExternalDeviceEncryptDecryptControlRegT struct {
	Val c.Uint32T
}

/** Type of sram_usage_conf register
 *  HP memory usage configuration register
 */

type HpSystemSramUsageConfRegT struct {
	Val c.Uint32T
}

/** Type of sec_dpa_conf register
 *  HP anti-DPA security configuration register
 */

type HpSystemSecDpaConfRegT struct {
	Val c.Uint32T
}

/** Type of sdio_ctrl register
 *  SDIO Control configuration register
 */

type HpSystemSdioCtrlRegT struct {
	Val c.Uint32T
}

/** Type of retention_conf register
 *  Retention configuration register
 */

type HpSystemRetentionConfRegT struct {
	Val c.Uint32T
}

/** Type of rom_table_lock register
 *  Rom-Table lock register
 */

type HpSystemRomTableLockRegT struct {
	Val c.Uint32T
}

/** Type of rom_table register
 *  Rom-Table register
 */

type HpSystemRomTableRegT struct {
	Val c.Uint32T
}

/** Type of core_debug_runstall_conf register
 *  Core Debug runstall configure register
 */

type HpSystemCoreDebugRunstallConfRegT struct {
	Val c.Uint32T
}

/** Type of mem_test_conf register
 *  MEM_TEST configuration register
 */

type HpSystemMemTestConfRegT struct {
	Val c.Uint32T
}

/** Type of clock_gate register
 *  HP-SYSTEM clock gating configure register
 */

type HpSystemClockGateRegT struct {
	Val c.Uint32T
}

/** Group: Timeout Register */
/** Type of cpu_peri_timeout_conf register
 *  CPU_PERI_TIMEOUT configuration register
 */

type HpSystemCpuPeriTimeoutConfRegT struct {
	Val c.Uint32T
}

/** Type of cpu_peri_timeout_addr register
 *  CPU_PERI_TIMEOUT_ADDR register
 */

type HpSystemCpuPeriTimeoutAddrRegT struct {
	Val c.Uint32T
}

/** Type of cpu_peri_timeout_uid register
 *  CPU_PERI_TIMEOUT_UID register
 */

type HpSystemCpuPeriTimeoutUidRegT struct {
	Val c.Uint32T
}

/** Type of hp_peri_timeout_conf register
 *  HP_PERI_TIMEOUT configuration register
 */

type HpSystemHpPeriTimeoutConfRegT struct {
	Val c.Uint32T
}

/** Type of hp_peri_timeout_addr register
 *  HP_PERI_TIMEOUT_ADDR register
 */

type HpSystemHpPeriTimeoutAddrRegT struct {
	Val c.Uint32T
}

/** Type of hp_peri_timeout_uid register
 *  HP_PERI_TIMEOUT_UID register
 */

type HpSystemHpPeriTimeoutUidRegT struct {
	Val c.Uint32T
}

/** Type of modem_peri_timeout_conf register
 *  MODEM_PERI_TIMEOUT configuration register
 */

type HpSystemModemPeriTimeoutConfRegT struct {
	Val c.Uint32T
}

/** Type of modem_peri_timeout_addr register
 *  MODEM_PERI_TIMEOUT_ADDR register
 */

type HpSystemModemPeriTimeoutAddrRegT struct {
	Val c.Uint32T
}

/** Type of modem_peri_timeout_uid register
 *  MODEM_PERI_TIMEOUT_UID register
 */

type HpSystemModemPeriTimeoutUidRegT struct {
	Val c.Uint32T
}

/** Group: Redcy ECO Registers */
/** Type of rnd_eco register
 *  redcy eco register.
 */

type HpSystemRndEcoRegT struct {
	Val c.Uint32T
}

/** Type of rnd_eco_low register
 *  redcy eco low register.
 */

type HpSystemRndEcoLowRegT struct {
	Val c.Uint32T
}

/** Type of rnd_eco_high register
 *  redcy eco high register.
 */

type HpSystemRndEcoHighRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  Date register.
 */

type HpSystemDateRegT struct {
	Val c.Uint32T
}

type HpSystemDevT struct {
	ExternalDeviceEncryptDecryptControl HpSystemExternalDeviceEncryptDecryptControlRegT
	SramUsageConf                       HpSystemSramUsageConfRegT
	SecDpaConf                          HpSystemSecDpaConfRegT
	CpuPeriTimeoutConf                  HpSystemCpuPeriTimeoutConfRegT
	CpuPeriTimeoutAddr                  HpSystemCpuPeriTimeoutAddrRegT
	CpuPeriTimeoutUid                   HpSystemCpuPeriTimeoutUidRegT
	HpPeriTimeoutConf                   HpSystemHpPeriTimeoutConfRegT
	HpPeriTimeoutAddr                   HpSystemHpPeriTimeoutAddrRegT
	HpPeriTimeoutUid                    HpSystemHpPeriTimeoutUidRegT
	ModemPeriTimeoutConf                HpSystemModemPeriTimeoutConfRegT
	ModemPeriTimeoutAddr                HpSystemModemPeriTimeoutAddrRegT
	ModemPeriTimeoutUid                 HpSystemModemPeriTimeoutUidRegT
	SdioCtrl                            HpSystemSdioCtrlRegT
	RetentionConf                       HpSystemRetentionConfRegT
	RomTableLock                        HpSystemRomTableLockRegT
	RomTable                            HpSystemRomTableRegT
	CoreDebugRunstallConf               HpSystemCoreDebugRunstallConfRegT
	MemTestConf                         HpSystemMemTestConfRegT
	Reserved048                         [230]c.Uint32T
	RndEco                              HpSystemRndEcoRegT
	RndEcoLow                           HpSystemRndEcoLowRegT
	RndEcoHigh                          HpSystemRndEcoHighRegT
	Reserved3ec                         [3]c.Uint32T
	ClockGate                           HpSystemClockGateRegT
	Date                                HpSystemDateRegT
}
