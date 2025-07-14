package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Control and configuration registers */
/** Type of l1_icache_ctrl register
 *  L1 instruction Cache(L1-ICache) control register
 */

type CacheL1IcacheCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_ctrl register
 *  L1 data Cache(L1-Cache) control register
 */

type CacheL1CacheCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_ctrl register
 *  L2 Cache(L2-Cache) control register
 */

type CacheL2CacheCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Bypass Cache Control and configuration registers */
/** Type of l1_bypass_cache_conf register
 *  Bypass Cache configure register
 */

type CacheL1BypassCacheConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_bypass_cache_conf register
 *  Bypass Cache configure register
 */

type CacheL2BypassCacheConfRegT struct {
	Val c.Uint32T
}

/** Group: Cache Atomic Control and configuration registers */
/** Type of l1_cache_atomic_conf register
 *  L1 Cache atomic feature configure register
 */

type CacheL1CacheAtomicConfRegT struct {
	Val c.Uint32T
}

/** Group: Cache Mode Control and configuration registers */
/** Type of l1_icache_cachesize_conf register
 *  L1 instruction Cache CacheSize mode configure register
 */

type CacheL1IcacheCachesizeConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache_blocksize_conf register
 *  L1 instruction Cache BlockSize mode configure register
 */

type CacheL1IcacheBlocksizeConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_cachesize_conf register
 *  L1 data Cache CacheSize mode configure register
 */

type CacheL1CacheCachesizeConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_blocksize_conf register
 *  L1 data Cache BlockSize mode configure register
 */

type CacheL1CacheBlocksizeConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_cachesize_conf register
 *  L2 Cache CacheSize mode configure register
 */

type CacheL2CacheCachesizeConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_blocksize_conf register
 *  L2 Cache BlockSize mode configure register
 */

type CacheL2CacheBlocksizeConfRegT struct {
	Val c.Uint32T
}

/** Group: Wrap Mode Control and configuration registers */
/** Type of l1_cache_wrap_around_ctrl register
 *  Cache wrap around control register
 */

type CacheL1CacheWrapAroundCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_wrap_around_ctrl register
 *  Cache wrap around control register
 */

type CacheL2CacheWrapAroundCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Cache Tag Memory Power Control registers */
/** Type of l1_cache_tag_mem_power_ctrl register
 *  Cache tag memory power control register
 */

type CacheL1CacheTagMemPowerCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_tag_mem_power_ctrl register
 *  Cache tag memory power control register
 */

type CacheL2CacheTagMemPowerCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Cache Data Memory Power Control registers */
/** Type of l1_cache_data_mem_power_ctrl register
 *  Cache data memory power control register
 */

type CacheL1CacheDataMemPowerCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_data_mem_power_ctrl register
 *  Cache data memory power control register
 */

type CacheL2CacheDataMemPowerCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Cache Freeze Control registers */
/** Type of l1_cache_freeze_ctrl register
 *  Cache Freeze control register
 */

type CacheL1CacheFreezeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_freeze_ctrl register
 *  Cache Freeze control register
 */

type CacheL2CacheFreezeCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Cache Data Memory Access Control and Configuration registers */
/** Type of l1_cache_data_mem_acs_conf register
 *  Cache data memory access configure register
 */

type CacheL1CacheDataMemAcsConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_data_mem_acs_conf register
 *  Cache data memory access configure register
 */

type CacheL2CacheDataMemAcsConfRegT struct {
	Val c.Uint32T
}

/** Group: Cache Tag Memory Access Control and Configuration registers */
/** Type of l1_cache_tag_mem_acs_conf register
 *  Cache tag memory access configure register
 */

type CacheL1CacheTagMemAcsConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_tag_mem_acs_conf register
 *  Cache tag memory access configure register
 */

type CacheL2CacheTagMemAcsConfRegT struct {
	Val c.Uint32T
}

/** Group: Prelock Control and configuration registers */
/** Type of l1_icache0_prelock_conf register
 *  L1 instruction Cache 0 prelock configure register
 */

type CacheL1Icache0PrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_prelock_sct0_addr register
 *  L1 instruction Cache 0 prelock section0 address configure register
 */

type CacheL1Icache0PrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_prelock_sct1_addr register
 *  L1 instruction Cache 0 prelock section1 address configure register
 */

type CacheL1Icache0PrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_prelock_sct_size register
 *  L1 instruction Cache 0 prelock section size configure register
 */

type CacheL1Icache0PrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_prelock_conf register
 *  L1 instruction Cache 1 prelock configure register
 */

type CacheL1Icache1PrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_prelock_sct0_addr register
 *  L1 instruction Cache 1 prelock section0 address configure register
 */

type CacheL1Icache1PrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_prelock_sct1_addr register
 *  L1 instruction Cache 1 prelock section1 address configure register
 */

type CacheL1Icache1PrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_prelock_sct_size register
 *  L1 instruction Cache 1 prelock section size configure register
 */

type CacheL1Icache1PrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_prelock_conf register
 *  L1 instruction Cache 2 prelock configure register
 */

type CacheL1Icache2PrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_prelock_sct0_addr register
 *  L1 instruction Cache 2 prelock section0 address configure register
 */

type CacheL1Icache2PrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_prelock_sct1_addr register
 *  L1 instruction Cache 2 prelock section1 address configure register
 */

type CacheL1Icache2PrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_prelock_sct_size register
 *  L1 instruction Cache 2 prelock section size configure register
 */

type CacheL1Icache2PrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_prelock_conf register
 *  L1 instruction Cache 3 prelock configure register
 */

type CacheL1Icache3PrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_prelock_sct0_addr register
 *  L1 instruction Cache 3 prelock section0 address configure register
 */

type CacheL1Icache3PrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_prelock_sct1_addr register
 *  L1 instruction Cache 3 prelock section1 address configure register
 */

type CacheL1Icache3PrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_prelock_sct_size register
 *  L1 instruction Cache 3 prelock section size configure register
 */

type CacheL1Icache3PrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_prelock_conf register
 *  L1 Cache prelock configure register
 */

type CacheL1CachePrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_prelock_sct0_addr register
 *  L1 Cache prelock section0 address configure register
 */

type CacheL1CachePrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_dcache_prelock_sct1_addr register
 *  L1 Cache prelock section1 address configure register
 */

type CacheL1DcachePrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_dcache_prelock_sct_size register
 *  L1  Cache prelock section size configure register
 */

type CacheL1DcachePrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_prelock_conf register
 *  L2 Cache prelock configure register
 */

type CacheL2CachePrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_prelock_sct0_addr register
 *  L2 Cache prelock section0 address configure register
 */

type CacheL2CachePrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_prelock_sct1_addr register
 *  L2 Cache prelock section1 address configure register
 */

type CacheL2CachePrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_prelock_sct_size register
 *  L2 Cache prelock section size configure register
 */

type CacheL2CachePrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Group: Lock Control and configuration registers */
/** Type of lock_ctrl register
 *  Lock-class (manual lock) operation control register
 */

type CacheLockCtrlRegT struct {
	Val c.Uint32T
}

/** Type of lock_map register
 *  Lock (manual lock) map configure register
 */

type CacheLockMapRegT struct {
	Val c.Uint32T
}

/** Type of lock_addr register
 *  Lock (manual lock) address configure register
 */

type CacheLockAddrRegT struct {
	Val c.Uint32T
}

/** Type of lock_size register
 *  Lock (manual lock) size configure register
 */

type CacheLockSizeRegT struct {
	Val c.Uint32T
}

/** Group: Sync Control and configuration registers */
/** Type of sync_ctrl register
 *  Sync-class operation control register
 */

type CacheSyncCtrlRegT struct {
	Val c.Uint32T
}

/** Type of sync_map register
 *  Sync map configure register
 */

type CacheSyncMapRegT struct {
	Val c.Uint32T
}

/** Type of sync_addr register
 *  Sync address configure register
 */

type CacheSyncAddrRegT struct {
	Val c.Uint32T
}

/** Type of sync_size register
 *  Sync size configure register
 */

type CacheSyncSizeRegT struct {
	Val c.Uint32T
}

/** Group: Preload Control and configuration registers */
/** Type of l1_icache0_preload_ctrl register
 *  L1 instruction Cache 0 preload-operation control register
 */

type CacheL1Icache0PreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_preload_addr register
 *  L1 instruction Cache 0 preload address configure register
 */

type CacheL1Icache0PreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_preload_size register
 *  L1 instruction Cache 0 preload size configure register
 */

type CacheL1Icache0PreloadSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_preload_ctrl register
 *  L1 instruction Cache 1 preload-operation control register
 */

type CacheL1Icache1PreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_preload_addr register
 *  L1 instruction Cache 1 preload address configure register
 */

type CacheL1Icache1PreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_preload_size register
 *  L1 instruction Cache 1 preload size configure register
 */

type CacheL1Icache1PreloadSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_preload_ctrl register
 *  L1 instruction Cache 2 preload-operation control register
 */

type CacheL1Icache2PreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_preload_addr register
 *  L1 instruction Cache 2 preload address configure register
 */

type CacheL1Icache2PreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_preload_size register
 *  L1 instruction Cache 2 preload size configure register
 */

type CacheL1Icache2PreloadSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_preload_ctrl register
 *  L1 instruction Cache 3 preload-operation control register
 */

type CacheL1Icache3PreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_preload_addr register
 *  L1 instruction Cache 3 preload address configure register
 */

type CacheL1Icache3PreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_preload_size register
 *  L1 instruction Cache 3 preload size configure register
 */

type CacheL1Icache3PreloadSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_preload_ctrl register
 *  L1 Cache  preload-operation control register
 */

type CacheL1CachePreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_dcache_preload_addr register
 *  L1 Cache  preload address configure register
 */

type CacheL1DcachePreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_dcache_preload_size register
 *  L1 Cache  preload size configure register
 */

type CacheL1DcachePreloadSizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_preload_ctrl register
 *  L2 Cache preload-operation control register
 */

type CacheL2CachePreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_preload_addr register
 *  L2 Cache preload address configure register
 */

type CacheL2CachePreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_preload_size register
 *  L2 Cache preload size configure register
 */

type CacheL2CachePreloadSizeRegT struct {
	Val c.Uint32T
}

/** Group: Autoload Control and configuration registers */
/** Type of l1_icache0_autoload_ctrl register
 *  L1 instruction Cache 0 autoload-operation control register
 */

type CacheL1Icache0AutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_autoload_sct0_addr register
 *  L1 instruction Cache 0 autoload section 0 address configure register
 */

type CacheL1Icache0AutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_autoload_sct0_size register
 *  L1 instruction Cache 0 autoload section 0 size configure register
 */

type CacheL1Icache0AutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_autoload_sct1_addr register
 *  L1 instruction Cache 0 autoload section 1 address configure register
 */

type CacheL1Icache0AutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_autoload_sct1_size register
 *  L1 instruction Cache 0 autoload section 1 size configure register
 */

type CacheL1Icache0AutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_autoload_ctrl register
 *  L1 instruction Cache 1 autoload-operation control register
 */

type CacheL1Icache1AutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_autoload_sct0_addr register
 *  L1 instruction Cache 1 autoload section 0 address configure register
 */

type CacheL1Icache1AutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_autoload_sct0_size register
 *  L1 instruction Cache 1 autoload section 0 size configure register
 */

type CacheL1Icache1AutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_autoload_sct1_addr register
 *  L1 instruction Cache 1 autoload section 1 address configure register
 */

type CacheL1Icache1AutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_autoload_sct1_size register
 *  L1 instruction Cache 1 autoload section 1 size configure register
 */

type CacheL1Icache1AutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_autoload_ctrl register
 *  L1 instruction Cache 2 autoload-operation control register
 */

type CacheL1Icache2AutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_autoload_sct0_addr register
 *  L1 instruction Cache 2 autoload section 0 address configure register
 */

type CacheL1Icache2AutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_autoload_sct0_size register
 *  L1 instruction Cache 2 autoload section 0 size configure register
 */

type CacheL1Icache2AutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_autoload_sct1_addr register
 *  L1 instruction Cache 2 autoload section 1 address configure register
 */

type CacheL1Icache2AutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_autoload_sct1_size register
 *  L1 instruction Cache 2 autoload section 1 size configure register
 */

type CacheL1Icache2AutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_autoload_ctrl register
 *  L1 instruction Cache 3 autoload-operation control register
 */

type CacheL1Icache3AutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_autoload_sct0_addr register
 *  L1 instruction Cache 3 autoload section 0 address configure register
 */

type CacheL1Icache3AutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_autoload_sct0_size register
 *  L1 instruction Cache 3 autoload section 0 size configure register
 */

type CacheL1Icache3AutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_autoload_sct1_addr register
 *  L1 instruction Cache 3 autoload section 1 address configure register
 */

type CacheL1Icache3AutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_autoload_sct1_size register
 *  L1 instruction Cache 3 autoload section 1 size configure register
 */

type CacheL1Icache3AutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_ctrl register
 *  L1 Cache autoload-operation control register
 */

type CacheL1CacheAutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct0_addr register
 *  L1 Cache autoload section 0 address configure register
 */

type CacheL1CacheAutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct0_size register
 *  L1 Cache autoload section 0 size configure register
 */

type CacheL1CacheAutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct1_addr register
 *  L1 Cache autoload section 1 address configure register
 */

type CacheL1CacheAutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct1_size register
 *  L1 Cache autoload section 1 size configure register
 */

type CacheL1CacheAutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct2_addr register
 *  L1 Cache autoload section 2 address configure register
 */

type CacheL1CacheAutoloadSct2AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct2_size register
 *  L1 Cache autoload section 2 size configure register
 */

type CacheL1CacheAutoloadSct2SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct3_addr register
 *  L1 Cache autoload section 1 address configure register
 */

type CacheL1CacheAutoloadSct3AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct3_size register
 *  L1 Cache autoload section 1 size configure register
 */

type CacheL1CacheAutoloadSct3SizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_ctrl register
 *  L2 Cache autoload-operation control register
 */

type CacheL2CacheAutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct0_addr register
 *  L2 Cache autoload section 0 address configure register
 */

type CacheL2CacheAutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct0_size register
 *  L2 Cache autoload section 0 size configure register
 */

type CacheL2CacheAutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct1_addr register
 *  L2 Cache autoload section 1 address configure register
 */

type CacheL2CacheAutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct1_size register
 *  L2 Cache autoload section 1 size configure register
 */

type CacheL2CacheAutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct2_addr register
 *  L2 Cache autoload section 2 address configure register
 */

type CacheL2CacheAutoloadSct2AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct2_size register
 *  L2 Cache autoload section 2 size configure register
 */

type CacheL2CacheAutoloadSct2SizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct3_addr register
 *  L2 Cache autoload section 3 address configure register
 */

type CacheL2CacheAutoloadSct3AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct3_size register
 *  L2 Cache autoload section 3 size configure register
 */

type CacheL2CacheAutoloadSct3SizeRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of l1_cache_acs_cnt_int_ena register
 *  Cache Access Counter Interrupt enable register
 */

type CacheL1CacheAcsCntIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_cnt_int_clr register
 *  Cache Access Counter Interrupt clear register
 */

type CacheL1CacheAcsCntIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_cnt_int_raw register
 *  Cache Access Counter Interrupt raw register
 */

type CacheL1CacheAcsCntIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_cnt_int_st register
 *  Cache Access Counter Interrupt status register
 */

type CacheL1CacheAcsCntIntStRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_fail_int_ena register
 *  Cache Access Fail Interrupt enable register
 */

type CacheL1CacheAcsFailIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_fail_int_clr register
 *  L1-Cache Access Fail Interrupt clear register
 */

type CacheL1CacheAcsFailIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_fail_int_raw register
 *  Cache Access Fail Interrupt raw register
 */

type CacheL1CacheAcsFailIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_fail_int_st register
 *  Cache Access Fail Interrupt status register
 */

type CacheL1CacheAcsFailIntStRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_sync_preload_int_ena register
 *  L1-Cache Access Fail Interrupt enable register
 */

type CacheL1CacheSyncPreloadIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_sync_preload_int_clr register
 *  Sync Preload operation Interrupt clear register
 */

type CacheL1CacheSyncPreloadIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_sync_preload_int_raw register
 *  Sync Preload operation Interrupt raw register
 */

type CacheL1CacheSyncPreloadIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_sync_preload_int_st register
 *  L1-Cache Access Fail Interrupt status register
 */

type CacheL1CacheSyncPreloadIntStRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_cnt_int_ena register
 *  Cache Access Counter Interrupt enable register
 */

type CacheL2CacheAcsCntIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_cnt_int_clr register
 *  Cache Access Counter Interrupt clear register
 */

type CacheL2CacheAcsCntIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_cnt_int_raw register
 *  Cache Access Counter Interrupt raw register
 */

type CacheL2CacheAcsCntIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_cnt_int_st register
 *  Cache Access Counter Interrupt status register
 */

type CacheL2CacheAcsCntIntStRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_int_ena register
 *  Cache Access Fail Interrupt enable register
 */

type CacheL2CacheAcsFailIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_int_clr register
 *  L1-Cache Access Fail Interrupt clear register
 */

type CacheL2CacheAcsFailIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_int_raw register
 *  Cache Access Fail Interrupt raw register
 */

type CacheL2CacheAcsFailIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_int_st register
 *  Cache Access Fail Interrupt status register
 */

type CacheL2CacheAcsFailIntStRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_preload_int_ena register
 *  L1-Cache Access Fail Interrupt enable register
 */

type CacheL2CacheSyncPreloadIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_preload_int_clr register
 *  Sync Preload operation Interrupt clear register
 */

type CacheL2CacheSyncPreloadIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_preload_int_raw register
 *  Sync Preload operation Interrupt raw register
 */

type CacheL2CacheSyncPreloadIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_preload_int_st register
 *  L1-Cache Access Fail Interrupt status register
 */

type CacheL2CacheSyncPreloadIntStRegT struct {
	Val c.Uint32T
}

/** Group: Access Statistics registers */
/** Type of l1_cache_acs_cnt_ctrl register
 *  Cache Access Counter enable and clear register
 */

type CacheL1CacheAcsCntCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus0_acs_hit_cnt register
 *  L1-ICache bus0 Hit-Access Counter register
 */

type CacheL1Ibus0AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus0_acs_miss_cnt register
 *  L1-ICache bus0 Miss-Access Counter register
 */

type CacheL1Ibus0AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus0_acs_conflict_cnt register
 *  L1-ICache bus0 Conflict-Access Counter register
 */

type CacheL1Ibus0AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus0_acs_nxtlvl_cnt register
 *  L1-ICache bus0 Next-Level-Access Counter register
 */

type CacheL1Ibus0AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus1_acs_hit_cnt register
 *  L1-ICache bus1 Hit-Access Counter register
 */

type CacheL1Ibus1AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus1_acs_miss_cnt register
 *  L1-ICache bus1 Miss-Access Counter register
 */

type CacheL1Ibus1AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus1_acs_conflict_cnt register
 *  L1-ICache bus1 Conflict-Access Counter register
 */

type CacheL1Ibus1AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus1_acs_nxtlvl_cnt register
 *  L1-ICache bus1 Next-Level-Access Counter register
 */

type CacheL1Ibus1AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus2_acs_hit_cnt register
 *  L1-ICache bus2 Hit-Access Counter register
 */

type CacheL1Ibus2AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus2_acs_miss_cnt register
 *  L1-ICache bus2 Miss-Access Counter register
 */

type CacheL1Ibus2AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus2_acs_conflict_cnt register
 *  L1-ICache bus2 Conflict-Access Counter register
 */

type CacheL1Ibus2AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus2_acs_nxtlvl_cnt register
 *  L1-ICache bus2 Next-Level-Access Counter register
 */

type CacheL1Ibus2AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus3_acs_hit_cnt register
 *  L1-ICache bus3 Hit-Access Counter register
 */

type CacheL1Ibus3AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus3_acs_miss_cnt register
 *  L1-ICache bus3 Miss-Access Counter register
 */

type CacheL1Ibus3AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus3_acs_conflict_cnt register
 *  L1-ICache bus3 Conflict-Access Counter register
 */

type CacheL1Ibus3AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus3_acs_nxtlvl_cnt register
 *  L1-ICache bus3 Next-Level-Access Counter register
 */

type CacheL1Ibus3AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus0_acs_hit_cnt register
 *  L1-Cache bus0 Hit-Access Counter register
 */

type CacheL1Bus0AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus0_acs_miss_cnt register
 *  L1-Cache bus0 Miss-Access Counter register
 */

type CacheL1Bus0AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus0_acs_conflict_cnt register
 *  L1-Cache bus0 Conflict-Access Counter register
 */

type CacheL1Bus0AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus0_acs_nxtlvl_cnt register
 *  L1-Cache bus0 Next-Level-Access Counter register
 */

type CacheL1Bus0AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus1_acs_hit_cnt register
 *  L1-Cache bus1 Hit-Access Counter register
 */

type CacheL1Bus1AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus1_acs_miss_cnt register
 *  L1-Cache bus1 Miss-Access Counter register
 */

type CacheL1Bus1AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus1_acs_conflict_cnt register
 *  L1-Cache bus1 Conflict-Access Counter register
 */

type CacheL1Bus1AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus1_acs_nxtlvl_cnt register
 *  L1-Cache bus1 Next-Level-Access Counter register
 */

type CacheL1Bus1AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus2_acs_hit_cnt register
 *  L1-DCache bus2 Hit-Access Counter register
 */

type CacheL1Dbus2AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus2_acs_miss_cnt register
 *  L1-DCache bus2 Miss-Access Counter register
 */

type CacheL1Dbus2AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus2_acs_conflict_cnt register
 *  L1-DCache bus2 Conflict-Access Counter register
 */

type CacheL1Dbus2AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus2_acs_nxtlvl_cnt register
 *  L1-DCache bus2 Next-Level-Access Counter register
 */

type CacheL1Dbus2AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus3_acs_hit_cnt register
 *  L1-DCache bus3 Hit-Access Counter register
 */

type CacheL1Dbus3AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus3_acs_miss_cnt register
 *  L1-DCache bus3 Miss-Access Counter register
 */

type CacheL1Dbus3AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus3_acs_conflict_cnt register
 *  L1-DCache bus3 Conflict-Access Counter register
 */

type CacheL1Dbus3AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus3_acs_nxtlvl_cnt register
 *  L1-DCache bus3 Next-Level-Access Counter register
 */

type CacheL1Dbus3AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_cnt_ctrl register
 *  Cache Access Counter enable and clear register
 */

type CacheL2CacheAcsCntCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus0_acs_hit_cnt register
 *  L2-Cache bus0 Hit-Access Counter register
 */

type CacheL2Ibus0AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus0_acs_miss_cnt register
 *  L2-Cache bus0 Miss-Access Counter register
 */

type CacheL2Ibus0AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus0_acs_conflict_cnt register
 *  L2-Cache bus0 Conflict-Access Counter register
 */

type CacheL2Ibus0AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus0_acs_nxtlvl_cnt register
 *  L2-Cache bus0 Next-Level-Access Counter register
 */

type CacheL2Ibus0AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus1_acs_hit_cnt register
 *  L2-Cache bus1 Hit-Access Counter register
 */

type CacheL2Ibus1AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus1_acs_miss_cnt register
 *  L2-Cache bus1 Miss-Access Counter register
 */

type CacheL2Ibus1AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus1_acs_conflict_cnt register
 *  L2-Cache bus1 Conflict-Access Counter register
 */

type CacheL2Ibus1AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus1_acs_nxtlvl_cnt register
 *  L2-Cache bus1 Next-Level-Access Counter register
 */

type CacheL2Ibus1AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus2_acs_hit_cnt register
 *  L2-Cache bus2 Hit-Access Counter register
 */

type CacheL2Ibus2AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus2_acs_miss_cnt register
 *  L2-Cache bus2 Miss-Access Counter register
 */

type CacheL2Ibus2AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus2_acs_conflict_cnt register
 *  L2-Cache bus2 Conflict-Access Counter register
 */

type CacheL2Ibus2AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus2_acs_nxtlvl_cnt register
 *  L2-Cache bus2 Next-Level-Access Counter register
 */

type CacheL2Ibus2AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus3_acs_hit_cnt register
 *  L2-Cache bus3 Hit-Access Counter register
 */

type CacheL2Ibus3AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus3_acs_miss_cnt register
 *  L2-Cache bus3 Miss-Access Counter register
 */

type CacheL2Ibus3AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus3_acs_conflict_cnt register
 *  L2-Cache bus3 Conflict-Access Counter register
 */

type CacheL2Ibus3AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus3_acs_nxtlvl_cnt register
 *  L2-Cache bus3 Next-Level-Access Counter register
 */

type CacheL2Ibus3AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus0_acs_hit_cnt register
 *  L2-Cache bus0 Hit-Access Counter register
 */

type CacheL2Dbus0AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus0_acs_miss_cnt register
 *  L2-Cache bus0 Miss-Access Counter register
 */

type CacheL2Dbus0AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus0_acs_conflict_cnt register
 *  L2-Cache bus0 Conflict-Access Counter register
 */

type CacheL2Dbus0AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus0_acs_nxtlvl_cnt register
 *  L2-Cache bus0 Next-Level-Access Counter register
 */

type CacheL2Dbus0AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus1_acs_hit_cnt register
 *  L2-Cache bus1 Hit-Access Counter register
 */

type CacheL2Dbus1AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus1_acs_miss_cnt register
 *  L2-Cache bus1 Miss-Access Counter register
 */

type CacheL2Dbus1AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus1_acs_conflict_cnt register
 *  L2-Cache bus1 Conflict-Access Counter register
 */

type CacheL2Dbus1AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus1_acs_nxtlvl_cnt register
 *  L2-Cache bus1 Next-Level-Access Counter register
 */

type CacheL2Dbus1AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus2_acs_hit_cnt register
 *  L2-Cache bus2 Hit-Access Counter register
 */

type CacheL2Dbus2AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus2_acs_miss_cnt register
 *  L2-Cache bus2 Miss-Access Counter register
 */

type CacheL2Dbus2AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus2_acs_conflict_cnt register
 *  L2-Cache bus2 Conflict-Access Counter register
 */

type CacheL2Dbus2AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus2_acs_nxtlvl_cnt register
 *  L2-Cache bus2 Next-Level-Access Counter register
 */

type CacheL2Dbus2AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus3_acs_hit_cnt register
 *  L2-Cache bus3 Hit-Access Counter register
 */

type CacheL2Dbus3AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus3_acs_miss_cnt register
 *  L2-Cache bus3 Miss-Access Counter register
 */

type CacheL2Dbus3AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus3_acs_conflict_cnt register
 *  L2-Cache bus3 Conflict-Access Counter register
 */

type CacheL2Dbus3AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus3_acs_nxtlvl_cnt register
 *  L2-Cache bus3 Next-Level-Access Counter register
 */

type CacheL2Dbus3AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Group: Access Fail Debug registers */
/** Type of l1_icache0_acs_fail_id_attr register
 *  L1-ICache0 Access Fail ID/attribution information register
 */

type CacheL1Icache0AcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_acs_fail_addr register
 *  L1-ICache0 Access Fail Address information register
 */

type CacheL1Icache0AcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_acs_fail_id_attr register
 *  L1-ICache0 Access Fail ID/attribution information register
 */

type CacheL1Icache1AcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_acs_fail_addr register
 *  L1-ICache0 Access Fail Address information register
 */

type CacheL1Icache1AcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_acs_fail_id_attr register
 *  L1-ICache0 Access Fail ID/attribution information register
 */

type CacheL1Icache2AcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_acs_fail_addr register
 *  L1-ICache0 Access Fail Address information register
 */

type CacheL1Icache2AcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_acs_fail_id_attr register
 *  L1-ICache0 Access Fail ID/attribution information register
 */

type CacheL1Icache3AcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_acs_fail_addr register
 *  L1-ICache0 Access Fail Address information register
 */

type CacheL1Icache3AcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_fail_id_attr register
 *  L1-Cache Access Fail ID/attribution information register
 */

type CacheL1CacheAcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l1_dcache_acs_fail_addr register
 *  L1-Cache Access Fail Address information register
 */

type CacheL1DcacheAcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_id_attr register
 *  L2-Cache Access Fail ID/attribution information register
 */

type CacheL2CacheAcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_addr register
 *  L2-Cache Access Fail Address information register
 */

type CacheL2CacheAcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Group: Operation Exception registers */
/** Type of l1_cache_sync_preload_exception register
 *  Cache Sync/Preload Operation exception register
 */

type CacheL1CacheSyncPreloadExceptionRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_preload_exception register
 *  Cache Sync/Preload Operation exception register
 */

type CacheL2CacheSyncPreloadExceptionRegT struct {
	Val c.Uint32T
}

/** Group: Sync Reset control and configuration registers */
/** Type of l1_cache_sync_rst_ctrl register
 *  Cache Sync Reset control register
 */

type CacheL1CacheSyncRstCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_rst_ctrl register
 *  Cache Sync Reset control register
 */

type CacheL2CacheSyncRstCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Preload Reset control and configuration registers */
/** Type of l1_cache_preload_rst_ctrl register
 *  Cache Preload Reset control register
 */

type CacheL1CachePreloadRstCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_preload_rst_ctrl register
 *  Cache Preload Reset control register
 */

type CacheL2CachePreloadRstCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Autoload buffer clear control and configuration registers */
/** Type of l1_cache_autoload_buf_clr_ctrl register
 *  Cache Autoload buffer clear control register
 */

type CacheL1CacheAutoloadBufClrCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_buf_clr_ctrl register
 *  Cache Autoload buffer clear control register
 */

type CacheL2CacheAutoloadBufClrCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Unallocate request buffer clear registers */
/** Type of l1_unallocate_buffer_clear register
 *  Unallocate request buffer clear registers
 */

type CacheL1UnallocateBufferClearRegT struct {
	Val c.Uint32T
}

/** Type of l2_unallocate_buffer_clear register
 *  Unallocate request buffer clear registers
 */

type CacheL2UnallocateBufferClearRegT struct {
	Val c.Uint32T
}

/** Group: Tag and Data Memory Access Control and configuration register */
/** Type of l1_cache_object_ctrl register
 *  Cache Tag and Data memory Object control register
 */

type CacheL1CacheObjectCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_way_object register
 *  Cache Tag and Data memory way register
 */

type CacheL1CacheWayObjectRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_vaddr register
 *  Cache Vaddr register
 */

type CacheL1CacheVaddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_debug_bus register
 *  Cache Tag/data memory content register
 */

type CacheL1CacheDebugBusRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_object_ctrl register
 *  Cache Tag and Data memory Object control register
 */

type CacheL2CacheObjectCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_way_object register
 *  Cache Tag and Data memory way register
 */

type CacheL2CacheWayObjectRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_vaddr register
 *  Cache Vaddr register
 */

type CacheL2CacheVaddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_debug_bus register
 *  Cache Tag/data memory content register
 */

type CacheL2CacheDebugBusRegT struct {
	Val c.Uint32T
}

/** Group: Split L1 and L2 registers */
/** Type of level_split0 register
 *  USED TO SPLIT L1 CACHE AND L2 CACHE
 */

type CacheLevelSplit0RegT struct {
	Val c.Uint32T
}

/** Type of level_split1 register
 *  USED TO SPLIT L1 CACHE AND L2 CACHE
 */

type CacheLevelSplit1RegT struct {
	Val c.Uint32T
}

/** Group: L2 cache access attribute control register */
/** Type of l2_cache_access_attr_ctrl register
 *  L1 Cache access Attribute propagation control register
 */

type CacheL2CacheAccessAttrCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Clock Gate Control and configuration register */
/** Type of clock_gate register
 *  Clock gate control register
 */

type CacheClockGateRegT struct {
	Val c.Uint32T
}

/** Group: Redundancy  register (Prepare for ECO) */
/** Type of redundancy_sig0 register
 *  Cache redundancy signal 0 register
 */

type CacheRedundancySig0RegT struct {
	Val c.Uint32T
}

/** Type of redundancy_sig1 register
 *  Cache redundancy signal 1 register
 */

type CacheRedundancySig1RegT struct {
	Val c.Uint32T
}

/** Type of redundancy_sig2 register
 *  Cache redundancy signal 2 register
 */

type CacheRedundancySig2RegT struct {
	Val c.Uint32T
}

/** Type of redundancy_sig3 register
 *  Cache redundancy signal 3 register
 */

type CacheRedundancySig3RegT struct {
	Val c.Uint32T
}

/** Type of redundancy_sig4 register
 *  Cache redundancy signal 0 register
 */

type CacheRedundancySig4RegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version control register
 */

type CacheDateRegT struct {
	Val c.Uint32T
}

type CacheDevT struct {
	L1IcacheCtrl                CacheL1IcacheCtrlRegT
	L1CacheCtrl                 CacheL1CacheCtrlRegT
	L1BypassCacheConf           CacheL1BypassCacheConfRegT
	L1CacheAtomicConf           CacheL1CacheAtomicConfRegT
	L1IcacheCachesizeConf       CacheL1IcacheCachesizeConfRegT
	L1IcacheBlocksizeConf       CacheL1IcacheBlocksizeConfRegT
	L1CacheCachesizeConf        CacheL1CacheCachesizeConfRegT
	L1CacheBlocksizeConf        CacheL1CacheBlocksizeConfRegT
	L1CacheWrapAroundCtrl       CacheL1CacheWrapAroundCtrlRegT
	L1CacheTagMemPowerCtrl      CacheL1CacheTagMemPowerCtrlRegT
	L1CacheDataMemPowerCtrl     CacheL1CacheDataMemPowerCtrlRegT
	L1CacheFreezeCtrl           CacheL1CacheFreezeCtrlRegT
	L1CacheDataMemAcsConf       CacheL1CacheDataMemAcsConfRegT
	L1CacheTagMemAcsConf        CacheL1CacheTagMemAcsConfRegT
	L1Icache0PrelockConf        CacheL1Icache0PrelockConfRegT
	L1Icache0PrelockSct0Addr    CacheL1Icache0PrelockSct0AddrRegT
	L1Icache0PrelockSct1Addr    CacheL1Icache0PrelockSct1AddrRegT
	L1Icache0PrelockSctSize     CacheL1Icache0PrelockSctSizeRegT
	L1Icache1PrelockConf        CacheL1Icache1PrelockConfRegT
	L1Icache1PrelockSct0Addr    CacheL1Icache1PrelockSct0AddrRegT
	L1Icache1PrelockSct1Addr    CacheL1Icache1PrelockSct1AddrRegT
	L1Icache1PrelockSctSize     CacheL1Icache1PrelockSctSizeRegT
	L1Icache2PrelockConf        CacheL1Icache2PrelockConfRegT
	L1Icache2PrelockSct0Addr    CacheL1Icache2PrelockSct0AddrRegT
	L1Icache2PrelockSct1Addr    CacheL1Icache2PrelockSct1AddrRegT
	L1Icache2PrelockSctSize     CacheL1Icache2PrelockSctSizeRegT
	L1Icache3PrelockConf        CacheL1Icache3PrelockConfRegT
	L1Icache3PrelockSct0Addr    CacheL1Icache3PrelockSct0AddrRegT
	L1Icache3PrelockSct1Addr    CacheL1Icache3PrelockSct1AddrRegT
	L1Icache3PrelockSctSize     CacheL1Icache3PrelockSctSizeRegT
	L1CachePrelockConf          CacheL1CachePrelockConfRegT
	L1CachePrelockSct0Addr      CacheL1CachePrelockSct0AddrRegT
	L1DcachePrelockSct1Addr     CacheL1DcachePrelockSct1AddrRegT
	L1DcachePrelockSctSize      CacheL1DcachePrelockSctSizeRegT
	LockCtrl                    CacheLockCtrlRegT
	LockMap                     CacheLockMapRegT
	LockAddr                    CacheLockAddrRegT
	LockSize                    CacheLockSizeRegT
	SyncCtrl                    CacheSyncCtrlRegT
	SyncMap                     CacheSyncMapRegT
	SyncAddr                    CacheSyncAddrRegT
	SyncSize                    CacheSyncSizeRegT
	L1Icache0PreloadCtrl        CacheL1Icache0PreloadCtrlRegT
	L1Icache0PreloadAddr        CacheL1Icache0PreloadAddrRegT
	L1Icache0PreloadSize        CacheL1Icache0PreloadSizeRegT
	L1Icache1PreloadCtrl        CacheL1Icache1PreloadCtrlRegT
	L1Icache1PreloadAddr        CacheL1Icache1PreloadAddrRegT
	L1Icache1PreloadSize        CacheL1Icache1PreloadSizeRegT
	L1Icache2PreloadCtrl        CacheL1Icache2PreloadCtrlRegT
	L1Icache2PreloadAddr        CacheL1Icache2PreloadAddrRegT
	L1Icache2PreloadSize        CacheL1Icache2PreloadSizeRegT
	L1Icache3PreloadCtrl        CacheL1Icache3PreloadCtrlRegT
	L1Icache3PreloadAddr        CacheL1Icache3PreloadAddrRegT
	L1Icache3PreloadSize        CacheL1Icache3PreloadSizeRegT
	L1CachePreloadCtrl          CacheL1CachePreloadCtrlRegT
	L1DcachePreloadAddr         CacheL1DcachePreloadAddrRegT
	L1DcachePreloadSize         CacheL1DcachePreloadSizeRegT
	L1Icache0AutoloadCtrl       CacheL1Icache0AutoloadCtrlRegT
	L1Icache0AutoloadSct0Addr   CacheL1Icache0AutoloadSct0AddrRegT
	L1Icache0AutoloadSct0Size   CacheL1Icache0AutoloadSct0SizeRegT
	L1Icache0AutoloadSct1Addr   CacheL1Icache0AutoloadSct1AddrRegT
	L1Icache0AutoloadSct1Size   CacheL1Icache0AutoloadSct1SizeRegT
	L1Icache1AutoloadCtrl       CacheL1Icache1AutoloadCtrlRegT
	L1Icache1AutoloadSct0Addr   CacheL1Icache1AutoloadSct0AddrRegT
	L1Icache1AutoloadSct0Size   CacheL1Icache1AutoloadSct0SizeRegT
	L1Icache1AutoloadSct1Addr   CacheL1Icache1AutoloadSct1AddrRegT
	L1Icache1AutoloadSct1Size   CacheL1Icache1AutoloadSct1SizeRegT
	L1Icache2AutoloadCtrl       CacheL1Icache2AutoloadCtrlRegT
	L1Icache2AutoloadSct0Addr   CacheL1Icache2AutoloadSct0AddrRegT
	L1Icache2AutoloadSct0Size   CacheL1Icache2AutoloadSct0SizeRegT
	L1Icache2AutoloadSct1Addr   CacheL1Icache2AutoloadSct1AddrRegT
	L1Icache2AutoloadSct1Size   CacheL1Icache2AutoloadSct1SizeRegT
	L1Icache3AutoloadCtrl       CacheL1Icache3AutoloadCtrlRegT
	L1Icache3AutoloadSct0Addr   CacheL1Icache3AutoloadSct0AddrRegT
	L1Icache3AutoloadSct0Size   CacheL1Icache3AutoloadSct0SizeRegT
	L1Icache3AutoloadSct1Addr   CacheL1Icache3AutoloadSct1AddrRegT
	L1Icache3AutoloadSct1Size   CacheL1Icache3AutoloadSct1SizeRegT
	L1CacheAutoloadCtrl         CacheL1CacheAutoloadCtrlRegT
	L1CacheAutoloadSct0Addr     CacheL1CacheAutoloadSct0AddrRegT
	L1CacheAutoloadSct0Size     CacheL1CacheAutoloadSct0SizeRegT
	L1CacheAutoloadSct1Addr     CacheL1CacheAutoloadSct1AddrRegT
	L1CacheAutoloadSct1Size     CacheL1CacheAutoloadSct1SizeRegT
	L1CacheAutoloadSct2Addr     CacheL1CacheAutoloadSct2AddrRegT
	L1CacheAutoloadSct2Size     CacheL1CacheAutoloadSct2SizeRegT
	L1CacheAutoloadSct3Addr     CacheL1CacheAutoloadSct3AddrRegT
	L1CacheAutoloadSct3Size     CacheL1CacheAutoloadSct3SizeRegT
	L1CacheAcsCntIntEna         CacheL1CacheAcsCntIntEnaRegT
	L1CacheAcsCntIntClr         CacheL1CacheAcsCntIntClrRegT
	L1CacheAcsCntIntRaw         CacheL1CacheAcsCntIntRawRegT
	L1CacheAcsCntIntSt          CacheL1CacheAcsCntIntStRegT
	L1CacheAcsFailIntEna        CacheL1CacheAcsFailIntEnaRegT
	L1CacheAcsFailIntClr        CacheL1CacheAcsFailIntClrRegT
	L1CacheAcsFailIntRaw        CacheL1CacheAcsFailIntRawRegT
	L1CacheAcsFailIntSt         CacheL1CacheAcsFailIntStRegT
	L1CacheAcsCntCtrl           CacheL1CacheAcsCntCtrlRegT
	L1Ibus0AcsHitCnt            CacheL1Ibus0AcsHitCntRegT
	L1Ibus0AcsMissCnt           CacheL1Ibus0AcsMissCntRegT
	L1Ibus0AcsConflictCnt       CacheL1Ibus0AcsConflictCntRegT
	L1Ibus0AcsNxtlvlCnt         CacheL1Ibus0AcsNxtlvlCntRegT
	L1Ibus1AcsHitCnt            CacheL1Ibus1AcsHitCntRegT
	L1Ibus1AcsMissCnt           CacheL1Ibus1AcsMissCntRegT
	L1Ibus1AcsConflictCnt       CacheL1Ibus1AcsConflictCntRegT
	L1Ibus1AcsNxtlvlCnt         CacheL1Ibus1AcsNxtlvlCntRegT
	L1Ibus2AcsHitCnt            CacheL1Ibus2AcsHitCntRegT
	L1Ibus2AcsMissCnt           CacheL1Ibus2AcsMissCntRegT
	L1Ibus2AcsConflictCnt       CacheL1Ibus2AcsConflictCntRegT
	L1Ibus2AcsNxtlvlCnt         CacheL1Ibus2AcsNxtlvlCntRegT
	L1Ibus3AcsHitCnt            CacheL1Ibus3AcsHitCntRegT
	L1Ibus3AcsMissCnt           CacheL1Ibus3AcsMissCntRegT
	L1Ibus3AcsConflictCnt       CacheL1Ibus3AcsConflictCntRegT
	L1Ibus3AcsNxtlvlCnt         CacheL1Ibus3AcsNxtlvlCntRegT
	L1Bus0AcsHitCnt             CacheL1Bus0AcsHitCntRegT
	L1Bus0AcsMissCnt            CacheL1Bus0AcsMissCntRegT
	L1Bus0AcsConflictCnt        CacheL1Bus0AcsConflictCntRegT
	L1Bus0AcsNxtlvlCnt          CacheL1Bus0AcsNxtlvlCntRegT
	L1Bus1AcsHitCnt             CacheL1Bus1AcsHitCntRegT
	L1Bus1AcsMissCnt            CacheL1Bus1AcsMissCntRegT
	L1Bus1AcsConflictCnt        CacheL1Bus1AcsConflictCntRegT
	L1Bus1AcsNxtlvlCnt          CacheL1Bus1AcsNxtlvlCntRegT
	L1Dbus2AcsHitCnt            CacheL1Dbus2AcsHitCntRegT
	L1Dbus2AcsMissCnt           CacheL1Dbus2AcsMissCntRegT
	L1Dbus2AcsConflictCnt       CacheL1Dbus2AcsConflictCntRegT
	L1Dbus2AcsNxtlvlCnt         CacheL1Dbus2AcsNxtlvlCntRegT
	L1Dbus3AcsHitCnt            CacheL1Dbus3AcsHitCntRegT
	L1Dbus3AcsMissCnt           CacheL1Dbus3AcsMissCntRegT
	L1Dbus3AcsConflictCnt       CacheL1Dbus3AcsConflictCntRegT
	L1Dbus3AcsNxtlvlCnt         CacheL1Dbus3AcsNxtlvlCntRegT
	L1Icache0AcsFailIdAttr      CacheL1Icache0AcsFailIdAttrRegT
	L1Icache0AcsFailAddr        CacheL1Icache0AcsFailAddrRegT
	L1Icache1AcsFailIdAttr      CacheL1Icache1AcsFailIdAttrRegT
	L1Icache1AcsFailAddr        CacheL1Icache1AcsFailAddrRegT
	L1Icache2AcsFailIdAttr      CacheL1Icache2AcsFailIdAttrRegT
	L1Icache2AcsFailAddr        CacheL1Icache2AcsFailAddrRegT
	L1Icache3AcsFailIdAttr      CacheL1Icache3AcsFailIdAttrRegT
	L1Icache3AcsFailAddr        CacheL1Icache3AcsFailAddrRegT
	L1CacheAcsFailIdAttr        CacheL1CacheAcsFailIdAttrRegT
	L1DcacheAcsFailAddr         CacheL1DcacheAcsFailAddrRegT
	L1CacheSyncPreloadIntEna    CacheL1CacheSyncPreloadIntEnaRegT
	L1CacheSyncPreloadIntClr    CacheL1CacheSyncPreloadIntClrRegT
	L1CacheSyncPreloadIntRaw    CacheL1CacheSyncPreloadIntRawRegT
	L1CacheSyncPreloadIntSt     CacheL1CacheSyncPreloadIntStRegT
	L1CacheSyncPreloadException CacheL1CacheSyncPreloadExceptionRegT
	L1CacheSyncRstCtrl          CacheL1CacheSyncRstCtrlRegT
	L1CachePreloadRstCtrl       CacheL1CachePreloadRstCtrlRegT
	L1CacheAutoloadBufClrCtrl   CacheL1CacheAutoloadBufClrCtrlRegT
	L1UnallocateBufferClear     CacheL1UnallocateBufferClearRegT
	L1CacheObjectCtrl           CacheL1CacheObjectCtrlRegT
	L1CacheWayObject            CacheL1CacheWayObjectRegT
	L1CacheVaddr                CacheL1CacheVaddrRegT
	L1CacheDebugBus             CacheL1CacheDebugBusRegT
	LevelSplit0                 CacheLevelSplit0RegT
	L2CacheCtrl                 CacheL2CacheCtrlRegT
	L2BypassCacheConf           CacheL2BypassCacheConfRegT
	L2CacheCachesizeConf        CacheL2CacheCachesizeConfRegT
	L2CacheBlocksizeConf        CacheL2CacheBlocksizeConfRegT
	L2CacheWrapAroundCtrl       CacheL2CacheWrapAroundCtrlRegT
	L2CacheTagMemPowerCtrl      CacheL2CacheTagMemPowerCtrlRegT
	L2CacheDataMemPowerCtrl     CacheL2CacheDataMemPowerCtrlRegT
	L2CacheFreezeCtrl           CacheL2CacheFreezeCtrlRegT
	L2CacheDataMemAcsConf       CacheL2CacheDataMemAcsConfRegT
	L2CacheTagMemAcsConf        CacheL2CacheTagMemAcsConfRegT
	L2CachePrelockConf          CacheL2CachePrelockConfRegT
	L2CachePrelockSct0Addr      CacheL2CachePrelockSct0AddrRegT
	L2CachePrelockSct1Addr      CacheL2CachePrelockSct1AddrRegT
	L2CachePrelockSctSize       CacheL2CachePrelockSctSizeRegT
	L2CachePreloadCtrl          CacheL2CachePreloadCtrlRegT
	L2CachePreloadAddr          CacheL2CachePreloadAddrRegT
	L2CachePreloadSize          CacheL2CachePreloadSizeRegT
	L2CacheAutoloadCtrl         CacheL2CacheAutoloadCtrlRegT
	L2CacheAutoloadSct0Addr     CacheL2CacheAutoloadSct0AddrRegT
	L2CacheAutoloadSct0Size     CacheL2CacheAutoloadSct0SizeRegT
	L2CacheAutoloadSct1Addr     CacheL2CacheAutoloadSct1AddrRegT
	L2CacheAutoloadSct1Size     CacheL2CacheAutoloadSct1SizeRegT
	L2CacheAutoloadSct2Addr     CacheL2CacheAutoloadSct2AddrRegT
	L2CacheAutoloadSct2Size     CacheL2CacheAutoloadSct2SizeRegT
	L2CacheAutoloadSct3Addr     CacheL2CacheAutoloadSct3AddrRegT
	L2CacheAutoloadSct3Size     CacheL2CacheAutoloadSct3SizeRegT
	L2CacheAcsCntIntEna         CacheL2CacheAcsCntIntEnaRegT
	L2CacheAcsCntIntClr         CacheL2CacheAcsCntIntClrRegT
	L2CacheAcsCntIntRaw         CacheL2CacheAcsCntIntRawRegT
	L2CacheAcsCntIntSt          CacheL2CacheAcsCntIntStRegT
	L2CacheAcsFailIntEna        CacheL2CacheAcsFailIntEnaRegT
	L2CacheAcsFailIntClr        CacheL2CacheAcsFailIntClrRegT
	L2CacheAcsFailIntRaw        CacheL2CacheAcsFailIntRawRegT
	L2CacheAcsFailIntSt         CacheL2CacheAcsFailIntStRegT
	L2CacheAcsCntCtrl           CacheL2CacheAcsCntCtrlRegT
	L2Ibus0AcsHitCnt            CacheL2Ibus0AcsHitCntRegT
	L2Ibus0AcsMissCnt           CacheL2Ibus0AcsMissCntRegT
	L2Ibus0AcsConflictCnt       CacheL2Ibus0AcsConflictCntRegT
	L2Ibus0AcsNxtlvlCnt         CacheL2Ibus0AcsNxtlvlCntRegT
	L2Ibus1AcsHitCnt            CacheL2Ibus1AcsHitCntRegT
	L2Ibus1AcsMissCnt           CacheL2Ibus1AcsMissCntRegT
	L2Ibus1AcsConflictCnt       CacheL2Ibus1AcsConflictCntRegT
	L2Ibus1AcsNxtlvlCnt         CacheL2Ibus1AcsNxtlvlCntRegT
	L2Ibus2AcsHitCnt            CacheL2Ibus2AcsHitCntRegT
	L2Ibus2AcsMissCnt           CacheL2Ibus2AcsMissCntRegT
	L2Ibus2AcsConflictCnt       CacheL2Ibus2AcsConflictCntRegT
	L2Ibus2AcsNxtlvlCnt         CacheL2Ibus2AcsNxtlvlCntRegT
	L2Ibus3AcsHitCnt            CacheL2Ibus3AcsHitCntRegT
	L2Ibus3AcsMissCnt           CacheL2Ibus3AcsMissCntRegT
	L2Ibus3AcsConflictCnt       CacheL2Ibus3AcsConflictCntRegT
	L2Ibus3AcsNxtlvlCnt         CacheL2Ibus3AcsNxtlvlCntRegT
	L2Dbus0AcsHitCnt            CacheL2Dbus0AcsHitCntRegT
	L2Dbus0AcsMissCnt           CacheL2Dbus0AcsMissCntRegT
	L2Dbus0AcsConflictCnt       CacheL2Dbus0AcsConflictCntRegT
	L2Dbus0AcsNxtlvlCnt         CacheL2Dbus0AcsNxtlvlCntRegT
	L2Dbus1AcsHitCnt            CacheL2Dbus1AcsHitCntRegT
	L2Dbus1AcsMissCnt           CacheL2Dbus1AcsMissCntRegT
	L2Dbus1AcsConflictCnt       CacheL2Dbus1AcsConflictCntRegT
	L2Dbus1AcsNxtlvlCnt         CacheL2Dbus1AcsNxtlvlCntRegT
	L2Dbus2AcsHitCnt            CacheL2Dbus2AcsHitCntRegT
	L2Dbus2AcsMissCnt           CacheL2Dbus2AcsMissCntRegT
	L2Dbus2AcsConflictCnt       CacheL2Dbus2AcsConflictCntRegT
	L2Dbus2AcsNxtlvlCnt         CacheL2Dbus2AcsNxtlvlCntRegT
	L2Dbus3AcsHitCnt            CacheL2Dbus3AcsHitCntRegT
	L2Dbus3AcsMissCnt           CacheL2Dbus3AcsMissCntRegT
	L2Dbus3AcsConflictCnt       CacheL2Dbus3AcsConflictCntRegT
	L2Dbus3AcsNxtlvlCnt         CacheL2Dbus3AcsNxtlvlCntRegT
	L2CacheAcsFailIdAttr        CacheL2CacheAcsFailIdAttrRegT
	L2CacheAcsFailAddr          CacheL2CacheAcsFailAddrRegT
	L2CacheSyncPreloadIntEna    CacheL2CacheSyncPreloadIntEnaRegT
	L2CacheSyncPreloadIntClr    CacheL2CacheSyncPreloadIntClrRegT
	L2CacheSyncPreloadIntRaw    CacheL2CacheSyncPreloadIntRawRegT
	L2CacheSyncPreloadIntSt     CacheL2CacheSyncPreloadIntStRegT
	L2CacheSyncPreloadException CacheL2CacheSyncPreloadExceptionRegT
	L2CacheSyncRstCtrl          CacheL2CacheSyncRstCtrlRegT
	L2CachePreloadRstCtrl       CacheL2CachePreloadRstCtrlRegT
	L2CacheAutoloadBufClrCtrl   CacheL2CacheAutoloadBufClrCtrlRegT
	L2UnallocateBufferClear     CacheL2UnallocateBufferClearRegT
	L2CacheAccessAttrCtrl       CacheL2CacheAccessAttrCtrlRegT
	L2CacheObjectCtrl           CacheL2CacheObjectCtrlRegT
	L2CacheWayObject            CacheL2CacheWayObjectRegT
	L2CacheVaddr                CacheL2CacheVaddrRegT
	L2CacheDebugBus             CacheL2CacheDebugBusRegT
	LevelSplit1                 CacheLevelSplit1RegT
	ClockGate                   CacheClockGateRegT
	RedundancySig0              CacheRedundancySig0RegT
	RedundancySig1              CacheRedundancySig1RegT
	RedundancySig2              CacheRedundancySig2RegT
	RedundancySig3              CacheRedundancySig3RegT
	RedundancySig4              CacheRedundancySig4RegT
	Reserved3c4                 [14]c.Uint32T
	Date                        CacheDateRegT
}
