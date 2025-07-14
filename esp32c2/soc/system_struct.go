package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: peripheral clock gating and reset registers */
/** Type of cpu_peri_clk_en register
 *  cpu_peripheral clock gating register
 */

type SystemCpuPeriClkEnRegT struct {
	Val c.Uint32T
}

/** Type of cpu_peri_rst_en register
 *  cpu_peripheral reset register
 */

type SystemCpuPeriRstEnRegT struct {
	Val c.Uint32T
}

/** Type of perip_clk_en0 register
 *  peripheral clock gating register
 */

type SystemPeripClkEn0RegT struct {
	Val c.Uint32T
}

/** Type of perip_clk_en1 register
 *  peripheral clock gating register
 */

type SystemPeripClkEn1RegT struct {
	Val c.Uint32T
}

/** Type of perip_rst_en0 register
 *  reserved
 */

type SystemPeripRstEn0RegT struct {
	Val c.Uint32T
}

/** Type of perip_rst_en1 register
 *  peripheral reset register
 */

type SystemPeripRstEn1RegT struct {
	Val c.Uint32T
}

/** Type of edma_ctrl register
 *  edma clcok and reset register
 */

type SystemEdmaCtrlRegT struct {
	Val c.Uint32T
}

/** Type of cache_control register
 *  cache control register
 */

type SystemCacheControlRegT struct {
	Val c.Uint32T
}

/** Group: clock config register */
/** Type of cpu_per_conf register
 *  cpu clock config register
 */

type SystemCpuPerConfRegT struct {
	Val c.Uint32T
}

/** Type of bt_lpck_div_int register
 *  clock config register
 */

type SystemBtLpckDivIntRegT struct {
	Val c.Uint32T
}

/** Type of bt_lpck_div_frac register
 *  low power clock configuration register
 */

type SystemBtLpckDivFracRegT struct {
	Val c.Uint32T
}

/** Type of sysclk_conf register
 *  system clock config register
 */

type SystemSysclkConfRegT struct {
	Val c.Uint32T
}

/** Group: Low-power management register */
/** Type of mem_pd_mask register
 *  memory power down mask register
 */

type SystemMemPdMaskRegT struct {
	Val c.Uint32T
}

/** Type of rtc_fastmem_config register
 *  fast memory config register
 */

type SystemRtcFastmemConfigRegT struct {
	Val c.Uint32T
}

/** Type of rtc_fastmem_crc register
 *  reserved
 */

type SystemRtcFastmemCrcRegT struct {
	Val c.Uint32T
}

/** Group: interrupt register */
/** Type of cpu_intr_from_cpu_0 register
 *  interrupt generate register
 */

type SystemCpuIntrFromCpu0RegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_1 register
 *  interrupt generate register
 */

type SystemCpuIntrFromCpu1RegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_2 register
 *  interrupt generate register
 */

type SystemCpuIntrFromCpu2RegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_3 register
 *  interrupt generate register
 */

type SystemCpuIntrFromCpu3RegT struct {
	Val c.Uint32T
}

/** Group: system and memory register */
/** Type of rsa_pd_ctrl register
 *  rsa memory power control register
 */

type SystemRsaPdCtrlRegT struct {
	Val c.Uint32T
}

/** Type of external_device_encrypt_decrypt_control register
 *  SYSTEM_EXTERNAL_DEVICE_ENCRYPT_DECRYPT_CONTROL_REG
 */

type SystemExternalDeviceEncryptDecryptControlRegT struct {
	Val c.Uint32T
}

/** Group: eco register */
/** Type of redundant_eco_ctrl register
 *  eco register
 */

type SystemRedundantEcoCtrlRegT struct {
	Val c.Uint32T
}

/** Group: clock gating register */
/** Type of clock_gate register
 *  clock gating register
 */

type SystemClockGateRegT struct {
	Val c.Uint32T
}

/** Group: pvt register */
/** Type of mem_pvt register
 *  mem pvt register
 */

type SystemMemPvtRegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_lvt_conf register
 *  mem pvt register
 */

type SystemCombPvtLvtConfRegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_nvt_conf register
 *  mem pvt register
 */

type SystemCombPvtNvtConfRegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_hvt_conf register
 *  mem pvt register
 */

type SystemCombPvtHvtConfRegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_lvt_site0 register
 *  mem pvt register
 */

type SystemCombPvtErrLvtSite0RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_nvt_site0 register
 *  mem pvt register
 */

type SystemCombPvtErrNvtSite0RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_hvt_site0 register
 *  mem pvt register
 */

type SystemCombPvtErrHvtSite0RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_lvt_site1 register
 *  mem pvt register
 */

type SystemCombPvtErrLvtSite1RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_nvt_site1 register
 *  mem pvt register
 */

type SystemCombPvtErrNvtSite1RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_hvt_site1 register
 *  mem pvt register
 */

type SystemCombPvtErrHvtSite1RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_lvt_site2 register
 *  mem pvt register
 */

type SystemCombPvtErrLvtSite2RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_nvt_site2 register
 *  mem pvt register
 */

type SystemCombPvtErrNvtSite2RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_hvt_site2 register
 *  mem pvt register
 */

type SystemCombPvtErrHvtSite2RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_lvt_site3 register
 *  mem pvt register
 */

type SystemCombPvtErrLvtSite3RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_nvt_site3 register
 *  mem pvt register
 */

type SystemCombPvtErrNvtSite3RegT struct {
	Val c.Uint32T
}

/** Type of comb_pvt_err_hvt_site3 register
 *  mem pvt register
 */

type SystemCombPvtErrHvtSite3RegT struct {
	Val c.Uint32T
}

/** Group: date register */
/** Type of reg_date register
 *  Version register
 */

type SystemRegDateRegT struct {
	Val c.Uint32T
}

type SystemDevT struct {
	CpuPeriClkEn                        SystemCpuPeriClkEnRegT
	CpuPeriRstEn                        SystemCpuPeriRstEnRegT
	CpuPerConf                          SystemCpuPerConfRegT
	MemPdMask                           SystemMemPdMaskRegT
	PeripClkEn0                         SystemPeripClkEn0RegT
	PeripClkEn1                         SystemPeripClkEn1RegT
	PeripRstEn0                         SystemPeripRstEn0RegT
	PeripRstEn1                         SystemPeripRstEn1RegT
	BtLpckDivInt                        SystemBtLpckDivIntRegT
	BtLpckDivFrac                       SystemBtLpckDivFracRegT
	CpuIntrFromCpu0                     SystemCpuIntrFromCpu0RegT
	CpuIntrFromCpu1                     SystemCpuIntrFromCpu1RegT
	CpuIntrFromCpu2                     SystemCpuIntrFromCpu2RegT
	CpuIntrFromCpu3                     SystemCpuIntrFromCpu3RegT
	RsaPdCtrl                           SystemRsaPdCtrlRegT
	EdmaCtrl                            SystemEdmaCtrlRegT
	CacheControl                        SystemCacheControlRegT
	ExternalDeviceEncryptDecryptControl SystemExternalDeviceEncryptDecryptControlRegT
	RtcFastmemConfig                    SystemRtcFastmemConfigRegT
	RtcFastmemCrc                       SystemRtcFastmemCrcRegT
	RedundantEcoCtrl                    SystemRedundantEcoCtrlRegT
	ClockGate                           SystemClockGateRegT
	SysclkConf                          SystemSysclkConfRegT
	MemPvt                              SystemMemPvtRegT
	CombPvtLvtConf                      SystemCombPvtLvtConfRegT
	CombPvtNvtConf                      SystemCombPvtNvtConfRegT
	CombPvtHvtConf                      SystemCombPvtHvtConfRegT
	CombPvtErrLvtSite0                  SystemCombPvtErrLvtSite0RegT
	CombPvtErrNvtSite0                  SystemCombPvtErrNvtSite0RegT
	CombPvtErrHvtSite0                  SystemCombPvtErrHvtSite0RegT
	CombPvtErrLvtSite1                  SystemCombPvtErrLvtSite1RegT
	CombPvtErrNvtSite1                  SystemCombPvtErrNvtSite1RegT
	CombPvtErrHvtSite1                  SystemCombPvtErrHvtSite1RegT
	CombPvtErrLvtSite2                  SystemCombPvtErrLvtSite2RegT
	CombPvtErrNvtSite2                  SystemCombPvtErrNvtSite2RegT
	CombPvtErrHvtSite2                  SystemCombPvtErrHvtSite2RegT
	CombPvtErrLvtSite3                  SystemCombPvtErrLvtSite3RegT
	CombPvtErrNvtSite3                  SystemCombPvtErrNvtSite3RegT
	CombPvtErrHvtSite3                  SystemCombPvtErrHvtSite3RegT
	Reserved09c                         [984]c.Uint32T
	RegDate                             SystemRegDateRegT
}
