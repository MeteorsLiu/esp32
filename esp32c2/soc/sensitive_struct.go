package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Registers */
/** Type of rom_table_lock register
 *  register description
 */

type SensitiveRomTableLockRegT struct {
	Val c.Uint32T
}

/** Type of rom_table register
 *  register description
 */

type SensitiveRomTableRegT struct {
	Val c.Uint32T
}

/** Type of apb_peripheral_access_0 register
 *  register description
 */

type SensitiveApbPeripheralAccess0RegT struct {
	Val c.Uint32T
}

/** Type of apb_peripheral_access_1 register
 *  register description
 */

type SensitiveApbPeripheralAccess1RegT struct {
	Val c.Uint32T
}

/** Type of internal_sram_usage_0 register
 *  register description
 */

type SensitiveInternalSramUsage0RegT struct {
	Val c.Uint32T
}

/** Type of internal_sram_usage_1 register
 *  register description
 */

type SensitiveInternalSramUsage1RegT struct {
	Val c.Uint32T
}

/** Type of internal_sram_usage_3 register
 *  register description
 */

type SensitiveInternalSramUsage3RegT struct {
	Val c.Uint32T
}

/** Type of cache_tag_access_0 register
 *  register description
 */

type SensitiveCacheTagAccess0RegT struct {
	Val c.Uint32T
}

/** Type of cache_tag_access_1 register
 *  register description
 */

type SensitiveCacheTagAccess1RegT struct {
	Val c.Uint32T
}

/** Type of cache_mmu_access_0 register
 *  register description
 */

type SensitiveCacheMmuAccess0RegT struct {
	Val c.Uint32T
}

/** Type of cache_mmu_access_1 register
 *  register description
 */

type SensitiveCacheMmuAccess1RegT struct {
	Val c.Uint32T
}

/** Type of pif_access_monitor_0 register
 *  register description
 */

type SensitivePifAccessMonitor0RegT struct {
	Val c.Uint32T
}

/** Type of pif_access_monitor_1 register
 *  register description
 */

type SensitivePifAccessMonitor1RegT struct {
	Val c.Uint32T
}

/** Type of pif_access_monitor_2 register
 *  register description
 */

type SensitivePifAccessMonitor2RegT struct {
	Val c.Uint32T
}

/** Type of pif_access_monitor_3 register
 *  register description
 */

type SensitivePifAccessMonitor3RegT struct {
	Val c.Uint32T
}

/** Type of xts_aes_key_update register
 *  register description
 */

type SensitiveXtsAesKeyUpdateRegT struct {
	Val c.Uint32T
}

/** Type of clock_gate register
 *  register description
 */

type SensitiveClockGateRegT struct {
	Val c.Uint32T
}

/** Type of sensitive_reg_date register
 *  register description
 */

type SensitiveSensitiveRegDateRegT struct {
	Val c.Uint32T
}

type SensitiveDevT struct {
	RomTableLock         SensitiveRomTableLockRegT
	RomTable             SensitiveRomTableRegT
	ApbPeripheralAccess0 SensitiveApbPeripheralAccess0RegT
	ApbPeripheralAccess1 SensitiveApbPeripheralAccess1RegT
	InternalSramUsage0   SensitiveInternalSramUsage0RegT
	InternalSramUsage1   SensitiveInternalSramUsage1RegT
	InternalSramUsage3   SensitiveInternalSramUsage3RegT
	CacheTagAccess0      SensitiveCacheTagAccess0RegT
	CacheTagAccess1      SensitiveCacheTagAccess1RegT
	CacheMmuAccess0      SensitiveCacheMmuAccess0RegT
	CacheMmuAccess1      SensitiveCacheMmuAccess1RegT
	PifAccessMonitor0    SensitivePifAccessMonitor0RegT
	PifAccessMonitor1    SensitivePifAccessMonitor1RegT
	PifAccessMonitor2    SensitivePifAccessMonitor2RegT
	PifAccessMonitor3    SensitivePifAccessMonitor3RegT
	XtsAesKeyUpdate      SensitiveXtsAesKeyUpdateRegT
	ClockGate            SensitiveClockGateRegT
	Reserved044          [1006]c.Uint32T
	SensitiveRegDate     SensitiveSensitiveRegDateRegT
}
