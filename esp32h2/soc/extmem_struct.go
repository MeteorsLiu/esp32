package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Control and configuration registers */
/** Type of l1_icache_ctrl register
 *  L1 instruction Cache(L1-ICache) control register
 */

type ExtmemL1IcacheCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_ctrl register
 *  L1 data Cache(L1-Cache) control register
 */

type ExtmemL1CacheCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_ctrl register
 *  L2 Cache(L2-Cache) control register
 */

type ExtmemL2CacheCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Bypass Cache Control and configuration registers */
/** Type of l1_bypass_cache_conf register
 *  Bypass Cache configure register
 */

type ExtmemL1BypassCacheConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_bypass_cache_conf register
 *  Bypass Cache configure register
 */

type ExtmemL2BypassCacheConfRegT struct {
	Val c.Uint32T
}

/** Group: Cache Atomic Control and configuration registers */
/** Type of l1_cache_atomic_conf register
 *  L1 Cache atomic feature configure register
 */

type ExtmemL1CacheAtomicConfRegT struct {
	Val c.Uint32T
}

/** Group: Cache Mode Control and configuration registers */
/** Type of l1_icache_cachesize_conf register
 *  L1 instruction Cache CacheSize mode configure register
 */

type ExtmemL1IcacheCachesizeConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache_blocksize_conf register
 *  L1 instruction Cache BlockSize mode configure register
 */

type ExtmemL1IcacheBlocksizeConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_cachesize_conf register
 *  L1 data Cache CacheSize mode configure register
 */

type ExtmemL1CacheCachesizeConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_blocksize_conf register
 *  L1 data Cache BlockSize mode configure register
 */

type ExtmemL1CacheBlocksizeConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_cachesize_conf register
 *  L2 Cache CacheSize mode configure register
 */

type ExtmemL2CacheCachesizeConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_blocksize_conf register
 *  L2 Cache BlockSize mode configure register
 */

type ExtmemL2CacheBlocksizeConfRegT struct {
	Val c.Uint32T
}

/** Group: Wrap Mode Control and configuration registers */
/** Type of l1_cache_wrap_around_ctrl register
 *  Cache wrap around control register
 */

type ExtmemL1CacheWrapAroundCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_wrap_around_ctrl register
 *  Cache wrap around control register
 */

type ExtmemL2CacheWrapAroundCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Cache Tag Memory Power Control registers */
/** Type of l1_cache_tag_mem_power_ctrl register
 *  Cache tag memory power control register
 */

type ExtmemL1CacheTagMemPowerCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_tag_mem_power_ctrl register
 *  Cache tag memory power control register
 */

type ExtmemL2CacheTagMemPowerCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Cache Data Memory Power Control registers */
/** Type of l1_cache_data_mem_power_ctrl register
 *  Cache data memory power control register
 */

type ExtmemL1CacheDataMemPowerCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_data_mem_power_ctrl register
 *  Cache data memory power control register
 */

type ExtmemL2CacheDataMemPowerCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Cache Freeze Control registers */
/** Type of l1_cache_freeze_ctrl register
 *  Cache Freeze control register
 */

type ExtmemL1CacheFreezeCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_freeze_ctrl register
 *  Cache Freeze control register
 */

type ExtmemL2CacheFreezeCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Cache Data Memory Access Control and Configuration registers */
/** Type of l1_cache_data_mem_acs_conf register
 *  Cache data memory access configure register
 */

type ExtmemL1CacheDataMemAcsConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_data_mem_acs_conf register
 *  Cache data memory access configure register
 */

type ExtmemL2CacheDataMemAcsConfRegT struct {
	Val c.Uint32T
}

/** Group: Cache Tag Memory Access Control and Configuration registers */
/** Type of l1_cache_tag_mem_acs_conf register
 *  Cache tag memory access configure register
 */

type ExtmemL1CacheTagMemAcsConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_tag_mem_acs_conf register
 *  Cache tag memory access configure register
 */

type ExtmemL2CacheTagMemAcsConfRegT struct {
	Val c.Uint32T
}

/** Group: Prelock Control and configuration registers */
/** Type of l1_icache0_prelock_conf register
 *  L1 instruction Cache 0 prelock configure register
 */

type ExtmemL1Icache0PrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_prelock_sct0_addr register
 *  L1 instruction Cache 0 prelock section0 address configure register
 */

type ExtmemL1Icache0PrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_prelock_sct1_addr register
 *  L1 instruction Cache 0 prelock section1 address configure register
 */

type ExtmemL1Icache0PrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_prelock_sct_size register
 *  L1 instruction Cache 0 prelock section size configure register
 */

type ExtmemL1Icache0PrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_prelock_conf register
 *  L1 instruction Cache 1 prelock configure register
 */

type ExtmemL1Icache1PrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_prelock_sct0_addr register
 *  L1 instruction Cache 1 prelock section0 address configure register
 */

type ExtmemL1Icache1PrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_prelock_sct1_addr register
 *  L1 instruction Cache 1 prelock section1 address configure register
 */

type ExtmemL1Icache1PrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_prelock_sct_size register
 *  L1 instruction Cache 1 prelock section size configure register
 */

type ExtmemL1Icache1PrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_prelock_conf register
 *  L1 instruction Cache 2 prelock configure register
 */

type ExtmemL1Icache2PrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_prelock_sct0_addr register
 *  L1 instruction Cache 2 prelock section0 address configure register
 */

type ExtmemL1Icache2PrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_prelock_sct1_addr register
 *  L1 instruction Cache 2 prelock section1 address configure register
 */

type ExtmemL1Icache2PrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_prelock_sct_size register
 *  L1 instruction Cache 2 prelock section size configure register
 */

type ExtmemL1Icache2PrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_prelock_conf register
 *  L1 instruction Cache 3 prelock configure register
 */

type ExtmemL1Icache3PrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_prelock_sct0_addr register
 *  L1 instruction Cache 3 prelock section0 address configure register
 */

type ExtmemL1Icache3PrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_prelock_sct1_addr register
 *  L1 instruction Cache 3 prelock section1 address configure register
 */

type ExtmemL1Icache3PrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_prelock_sct_size register
 *  L1 instruction Cache 3 prelock section size configure register
 */

type ExtmemL1Icache3PrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_prelock_conf register
 *  L1 Cache prelock configure register
 */

type ExtmemL1CachePrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_prelock_sct0_addr register
 *  L1 Cache prelock section0 address configure register
 */

type ExtmemL1CachePrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_dcache_prelock_sct1_addr register
 *  L1 Cache prelock section1 address configure register
 */

type ExtmemL1DcachePrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_dcache_prelock_sct_size register
 *  L1  Cache prelock section size configure register
 */

type ExtmemL1DcachePrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_prelock_conf register
 *  L2 Cache prelock configure register
 */

type ExtmemL2CachePrelockConfRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_prelock_sct0_addr register
 *  L2 Cache prelock section0 address configure register
 */

type ExtmemL2CachePrelockSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_prelock_sct1_addr register
 *  L2 Cache prelock section1 address configure register
 */

type ExtmemL2CachePrelockSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_prelock_sct_size register
 *  L2 Cache prelock section size configure register
 */

type ExtmemL2CachePrelockSctSizeRegT struct {
	Val c.Uint32T
}

/** Group: Lock Control and configuration registers */
/** Type of cache_lock_ctrl register
 *  Lock-class (manual lock) operation control register
 */

type ExtmemCacheLockCtrlRegT struct {
	Val c.Uint32T
}

/** Type of cache_lock_map register
 *  Lock (manual lock) map configure register
 */

type ExtmemCacheLockMapRegT struct {
	Val c.Uint32T
}

/** Type of cache_lock_addr register
 *  Lock (manual lock) address configure register
 */

type ExtmemCacheLockAddrRegT struct {
	Val c.Uint32T
}

/** Type of cache_lock_size register
 *  Lock (manual lock) size configure register
 */

type ExtmemCacheLockSizeRegT struct {
	Val c.Uint32T
}

/** Group: Sync Control and configuration registers */
/** Type of cache_sync_ctrl register
 *  Sync-class operation control register
 */

type ExtmemCacheSyncCtrlRegT struct {
	Val c.Uint32T
}

/** Type of cache_sync_map register
 *  Sync map configure register
 */

type ExtmemCacheSyncMapRegT struct {
	Val c.Uint32T
}

/** Type of cache_sync_addr register
 *  Sync address configure register
 */

type ExtmemCacheSyncAddrRegT struct {
	Val c.Uint32T
}

/** Type of cache_sync_size register
 *  Sync size configure register
 */

type ExtmemCacheSyncSizeRegT struct {
	Val c.Uint32T
}

/** Group: Preload Control and configuration registers */
/** Type of l1_icache0_preload_ctrl register
 *  L1 instruction Cache 0 preload-operation control register
 */

type ExtmemL1Icache0PreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_preload_addr register
 *  L1 instruction Cache 0 preload address configure register
 */

type ExtmemL1Icache0PreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_preload_size register
 *  L1 instruction Cache 0 preload size configure register
 */

type ExtmemL1Icache0PreloadSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_preload_ctrl register
 *  L1 instruction Cache 1 preload-operation control register
 */

type ExtmemL1Icache1PreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_preload_addr register
 *  L1 instruction Cache 1 preload address configure register
 */

type ExtmemL1Icache1PreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_preload_size register
 *  L1 instruction Cache 1 preload size configure register
 */

type ExtmemL1Icache1PreloadSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_preload_ctrl register
 *  L1 instruction Cache 2 preload-operation control register
 */

type ExtmemL1Icache2PreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_preload_addr register
 *  L1 instruction Cache 2 preload address configure register
 */

type ExtmemL1Icache2PreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_preload_size register
 *  L1 instruction Cache 2 preload size configure register
 */

type ExtmemL1Icache2PreloadSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_preload_ctrl register
 *  L1 instruction Cache 3 preload-operation control register
 */

type ExtmemL1Icache3PreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_preload_addr register
 *  L1 instruction Cache 3 preload address configure register
 */

type ExtmemL1Icache3PreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_preload_size register
 *  L1 instruction Cache 3 preload size configure register
 */

type ExtmemL1Icache3PreloadSizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_preload_ctrl register
 *  L1 Cache  preload-operation control register
 */

type ExtmemL1CachePreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_dcache_preload_addr register
 *  L1 Cache  preload address configure register
 */

type ExtmemL1DcachePreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_dcache_preload_size register
 *  L1 Cache  preload size configure register
 */

type ExtmemL1DcachePreloadSizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_preload_ctrl register
 *  L2 Cache preload-operation control register
 */

type ExtmemL2CachePreloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_preload_addr register
 *  L2 Cache preload address configure register
 */

type ExtmemL2CachePreloadAddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_preload_size register
 *  L2 Cache preload size configure register
 */

type ExtmemL2CachePreloadSizeRegT struct {
	Val c.Uint32T
}

/** Group: Autoload Control and configuration registers */
/** Type of l1_icache0_autoload_ctrl register
 *  L1 instruction Cache 0 autoload-operation control register
 */

type ExtmemL1Icache0AutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_autoload_sct0_addr register
 *  L1 instruction Cache 0 autoload section 0 address configure register
 */

type ExtmemL1Icache0AutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_autoload_sct0_size register
 *  L1 instruction Cache 0 autoload section 0 size configure register
 */

type ExtmemL1Icache0AutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_autoload_sct1_addr register
 *  L1 instruction Cache 0 autoload section 1 address configure register
 */

type ExtmemL1Icache0AutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_autoload_sct1_size register
 *  L1 instruction Cache 0 autoload section 1 size configure register
 */

type ExtmemL1Icache0AutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_autoload_ctrl register
 *  L1 instruction Cache 1 autoload-operation control register
 */

type ExtmemL1Icache1AutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_autoload_sct0_addr register
 *  L1 instruction Cache 1 autoload section 0 address configure register
 */

type ExtmemL1Icache1AutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_autoload_sct0_size register
 *  L1 instruction Cache 1 autoload section 0 size configure register
 */

type ExtmemL1Icache1AutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_autoload_sct1_addr register
 *  L1 instruction Cache 1 autoload section 1 address configure register
 */

type ExtmemL1Icache1AutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_autoload_sct1_size register
 *  L1 instruction Cache 1 autoload section 1 size configure register
 */

type ExtmemL1Icache1AutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_autoload_ctrl register
 *  L1 instruction Cache 2 autoload-operation control register
 */

type ExtmemL1Icache2AutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_autoload_sct0_addr register
 *  L1 instruction Cache 2 autoload section 0 address configure register
 */

type ExtmemL1Icache2AutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_autoload_sct0_size register
 *  L1 instruction Cache 2 autoload section 0 size configure register
 */

type ExtmemL1Icache2AutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_autoload_sct1_addr register
 *  L1 instruction Cache 2 autoload section 1 address configure register
 */

type ExtmemL1Icache2AutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_autoload_sct1_size register
 *  L1 instruction Cache 2 autoload section 1 size configure register
 */

type ExtmemL1Icache2AutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_autoload_ctrl register
 *  L1 instruction Cache 3 autoload-operation control register
 */

type ExtmemL1Icache3AutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_autoload_sct0_addr register
 *  L1 instruction Cache 3 autoload section 0 address configure register
 */

type ExtmemL1Icache3AutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_autoload_sct0_size register
 *  L1 instruction Cache 3 autoload section 0 size configure register
 */

type ExtmemL1Icache3AutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_autoload_sct1_addr register
 *  L1 instruction Cache 3 autoload section 1 address configure register
 */

type ExtmemL1Icache3AutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_autoload_sct1_size register
 *  L1 instruction Cache 3 autoload section 1 size configure register
 */

type ExtmemL1Icache3AutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_ctrl register
 *  L1 Cache autoload-operation control register
 */

type ExtmemL1CacheAutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct0_addr register
 *  L1 Cache autoload section 0 address configure register
 */

type ExtmemL1CacheAutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct0_size register
 *  L1 Cache autoload section 0 size configure register
 */

type ExtmemL1CacheAutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct1_addr register
 *  L1 Cache autoload section 1 address configure register
 */

type ExtmemL1CacheAutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct1_size register
 *  L1 Cache autoload section 1 size configure register
 */

type ExtmemL1CacheAutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct2_addr register
 *  L1 Cache autoload section 2 address configure register
 */

type ExtmemL1CacheAutoloadSct2AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct2_size register
 *  L1 Cache autoload section 2 size configure register
 */

type ExtmemL1CacheAutoloadSct2SizeRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct3_addr register
 *  L1 Cache autoload section 1 address configure register
 */

type ExtmemL1CacheAutoloadSct3AddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_autoload_sct3_size register
 *  L1 Cache autoload section 1 size configure register
 */

type ExtmemL1CacheAutoloadSct3SizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_ctrl register
 *  L2 Cache autoload-operation control register
 */

type ExtmemL2CacheAutoloadCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct0_addr register
 *  L2 Cache autoload section 0 address configure register
 */

type ExtmemL2CacheAutoloadSct0AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct0_size register
 *  L2 Cache autoload section 0 size configure register
 */

type ExtmemL2CacheAutoloadSct0SizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct1_addr register
 *  L2 Cache autoload section 1 address configure register
 */

type ExtmemL2CacheAutoloadSct1AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct1_size register
 *  L2 Cache autoload section 1 size configure register
 */

type ExtmemL2CacheAutoloadSct1SizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct2_addr register
 *  L2 Cache autoload section 2 address configure register
 */

type ExtmemL2CacheAutoloadSct2AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct2_size register
 *  L2 Cache autoload section 2 size configure register
 */

type ExtmemL2CacheAutoloadSct2SizeRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct3_addr register
 *  L2 Cache autoload section 3 address configure register
 */

type ExtmemL2CacheAutoloadSct3AddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_sct3_size register
 *  L2 Cache autoload section 3 size configure register
 */

type ExtmemL2CacheAutoloadSct3SizeRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of l1_cache_acs_cnt_int_ena register
 *  Cache Access Counter Interrupt enable register
 */

type ExtmemL1CacheAcsCntIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_cnt_int_clr register
 *  Cache Access Counter Interrupt clear register
 */

type ExtmemL1CacheAcsCntIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_cnt_int_raw register
 *  Cache Access Counter Interrupt raw register
 */

type ExtmemL1CacheAcsCntIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_cnt_int_st register
 *  Cache Access Counter Interrupt status register
 */

type ExtmemL1CacheAcsCntIntStRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_fail_int_ena register
 *  Cache Access Fail Interrupt enable register
 */

type ExtmemL1CacheAcsFailIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_fail_int_clr register
 *  L1-Cache Access Fail Interrupt clear register
 */

type ExtmemL1CacheAcsFailIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_fail_int_raw register
 *  Cache Access Fail Interrupt raw register
 */

type ExtmemL1CacheAcsFailIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_fail_int_st register
 *  Cache Access Fail Interrupt status register
 */

type ExtmemL1CacheAcsFailIntStRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_sync_preload_int_ena register
 *  L1-Cache Access Fail Interrupt enable register
 */

type ExtmemL1CacheSyncPreloadIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_sync_preload_int_clr register
 *  Sync Preload operation Interrupt clear register
 */

type ExtmemL1CacheSyncPreloadIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_sync_preload_int_raw register
 *  Sync Preload operation Interrupt raw register
 */

type ExtmemL1CacheSyncPreloadIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_sync_preload_int_st register
 *  L1-Cache Access Fail Interrupt status register
 */

type ExtmemL1CacheSyncPreloadIntStRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_cnt_int_ena register
 *  Cache Access Counter Interrupt enable register
 */

type ExtmemL2CacheAcsCntIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_cnt_int_clr register
 *  Cache Access Counter Interrupt clear register
 */

type ExtmemL2CacheAcsCntIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_cnt_int_raw register
 *  Cache Access Counter Interrupt raw register
 */

type ExtmemL2CacheAcsCntIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_cnt_int_st register
 *  Cache Access Counter Interrupt status register
 */

type ExtmemL2CacheAcsCntIntStRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_int_ena register
 *  Cache Access Fail Interrupt enable register
 */

type ExtmemL2CacheAcsFailIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_int_clr register
 *  L1-Cache Access Fail Interrupt clear register
 */

type ExtmemL2CacheAcsFailIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_int_raw register
 *  Cache Access Fail Interrupt raw register
 */

type ExtmemL2CacheAcsFailIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_int_st register
 *  Cache Access Fail Interrupt status register
 */

type ExtmemL2CacheAcsFailIntStRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_preload_int_ena register
 *  L1-Cache Access Fail Interrupt enable register
 */

type ExtmemL2CacheSyncPreloadIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_preload_int_clr register
 *  Sync Preload operation Interrupt clear register
 */

type ExtmemL2CacheSyncPreloadIntClrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_preload_int_raw register
 *  Sync Preload operation Interrupt raw register
 */

type ExtmemL2CacheSyncPreloadIntRawRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_preload_int_st register
 *  L1-Cache Access Fail Interrupt status register
 */

type ExtmemL2CacheSyncPreloadIntStRegT struct {
	Val c.Uint32T
}

/** Group: Access Statistics registers */
/** Type of l1_cache_acs_cnt_ctrl register
 *  Cache Access Counter enable and clear register
 */

type ExtmemL1CacheAcsCntCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus0_acs_hit_cnt register
 *  L1-ICache bus0 Hit-Access Counter register
 */

type ExtmemL1Ibus0AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus0_acs_miss_cnt register
 *  L1-ICache bus0 Miss-Access Counter register
 */

type ExtmemL1Ibus0AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus0_acs_conflict_cnt register
 *  L1-ICache bus0 Conflict-Access Counter register
 */

type ExtmemL1Ibus0AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus0_acs_nxtlvl_cnt register
 *  L1-ICache bus0 Next-Level-Access Counter register
 */

type ExtmemL1Ibus0AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus1_acs_hit_cnt register
 *  L1-ICache bus1 Hit-Access Counter register
 */

type ExtmemL1Ibus1AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus1_acs_miss_cnt register
 *  L1-ICache bus1 Miss-Access Counter register
 */

type ExtmemL1Ibus1AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus1_acs_conflict_cnt register
 *  L1-ICache bus1 Conflict-Access Counter register
 */

type ExtmemL1Ibus1AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus1_acs_nxtlvl_cnt register
 *  L1-ICache bus1 Next-Level-Access Counter register
 */

type ExtmemL1Ibus1AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus2_acs_hit_cnt register
 *  L1-ICache bus2 Hit-Access Counter register
 */

type ExtmemL1Ibus2AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus2_acs_miss_cnt register
 *  L1-ICache bus2 Miss-Access Counter register
 */

type ExtmemL1Ibus2AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus2_acs_conflict_cnt register
 *  L1-ICache bus2 Conflict-Access Counter register
 */

type ExtmemL1Ibus2AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus2_acs_nxtlvl_cnt register
 *  L1-ICache bus2 Next-Level-Access Counter register
 */

type ExtmemL1Ibus2AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus3_acs_hit_cnt register
 *  L1-ICache bus3 Hit-Access Counter register
 */

type ExtmemL1Ibus3AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus3_acs_miss_cnt register
 *  L1-ICache bus3 Miss-Access Counter register
 */

type ExtmemL1Ibus3AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus3_acs_conflict_cnt register
 *  L1-ICache bus3 Conflict-Access Counter register
 */

type ExtmemL1Ibus3AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_ibus3_acs_nxtlvl_cnt register
 *  L1-ICache bus3 Next-Level-Access Counter register
 */

type ExtmemL1Ibus3AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus0_acs_hit_cnt register
 *  L1-Cache bus0 Hit-Access Counter register
 */

type ExtmemL1Bus0AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus0_acs_miss_cnt register
 *  L1-Cache bus0 Miss-Access Counter register
 */

type ExtmemL1Bus0AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus0_acs_conflict_cnt register
 *  L1-Cache bus0 Conflict-Access Counter register
 */

type ExtmemL1Bus0AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus0_acs_nxtlvl_cnt register
 *  L1-Cache bus0 Next-Level-Access Counter register
 */

type ExtmemL1Bus0AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus1_acs_hit_cnt register
 *  L1-Cache bus1 Hit-Access Counter register
 */

type ExtmemL1Bus1AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus1_acs_miss_cnt register
 *  L1-Cache bus1 Miss-Access Counter register
 */

type ExtmemL1Bus1AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus1_acs_conflict_cnt register
 *  L1-Cache bus1 Conflict-Access Counter register
 */

type ExtmemL1Bus1AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_bus1_acs_nxtlvl_cnt register
 *  L1-Cache bus1 Next-Level-Access Counter register
 */

type ExtmemL1Bus1AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus2_acs_hit_cnt register
 *  L1-DCache bus2 Hit-Access Counter register
 */

type ExtmemL1Dbus2AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus2_acs_miss_cnt register
 *  L1-DCache bus2 Miss-Access Counter register
 */

type ExtmemL1Dbus2AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus2_acs_conflict_cnt register
 *  L1-DCache bus2 Conflict-Access Counter register
 */

type ExtmemL1Dbus2AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus2_acs_nxtlvl_cnt register
 *  L1-DCache bus2 Next-Level-Access Counter register
 */

type ExtmemL1Dbus2AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus3_acs_hit_cnt register
 *  L1-DCache bus3 Hit-Access Counter register
 */

type ExtmemL1Dbus3AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus3_acs_miss_cnt register
 *  L1-DCache bus3 Miss-Access Counter register
 */

type ExtmemL1Dbus3AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus3_acs_conflict_cnt register
 *  L1-DCache bus3 Conflict-Access Counter register
 */

type ExtmemL1Dbus3AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l1_dbus3_acs_nxtlvl_cnt register
 *  L1-DCache bus3 Next-Level-Access Counter register
 */

type ExtmemL1Dbus3AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_cnt_ctrl register
 *  Cache Access Counter enable and clear register
 */

type ExtmemL2CacheAcsCntCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus0_acs_hit_cnt register
 *  L2-Cache bus0 Hit-Access Counter register
 */

type ExtmemL2Ibus0AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus0_acs_miss_cnt register
 *  L2-Cache bus0 Miss-Access Counter register
 */

type ExtmemL2Ibus0AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus0_acs_conflict_cnt register
 *  L2-Cache bus0 Conflict-Access Counter register
 */

type ExtmemL2Ibus0AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus0_acs_nxtlvl_cnt register
 *  L2-Cache bus0 Next-Level-Access Counter register
 */

type ExtmemL2Ibus0AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus1_acs_hit_cnt register
 *  L2-Cache bus1 Hit-Access Counter register
 */

type ExtmemL2Ibus1AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus1_acs_miss_cnt register
 *  L2-Cache bus1 Miss-Access Counter register
 */

type ExtmemL2Ibus1AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus1_acs_conflict_cnt register
 *  L2-Cache bus1 Conflict-Access Counter register
 */

type ExtmemL2Ibus1AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus1_acs_nxtlvl_cnt register
 *  L2-Cache bus1 Next-Level-Access Counter register
 */

type ExtmemL2Ibus1AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus2_acs_hit_cnt register
 *  L2-Cache bus2 Hit-Access Counter register
 */

type ExtmemL2Ibus2AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus2_acs_miss_cnt register
 *  L2-Cache bus2 Miss-Access Counter register
 */

type ExtmemL2Ibus2AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus2_acs_conflict_cnt register
 *  L2-Cache bus2 Conflict-Access Counter register
 */

type ExtmemL2Ibus2AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus2_acs_nxtlvl_cnt register
 *  L2-Cache bus2 Next-Level-Access Counter register
 */

type ExtmemL2Ibus2AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus3_acs_hit_cnt register
 *  L2-Cache bus3 Hit-Access Counter register
 */

type ExtmemL2Ibus3AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus3_acs_miss_cnt register
 *  L2-Cache bus3 Miss-Access Counter register
 */

type ExtmemL2Ibus3AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus3_acs_conflict_cnt register
 *  L2-Cache bus3 Conflict-Access Counter register
 */

type ExtmemL2Ibus3AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_ibus3_acs_nxtlvl_cnt register
 *  L2-Cache bus3 Next-Level-Access Counter register
 */

type ExtmemL2Ibus3AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus0_acs_hit_cnt register
 *  L2-Cache bus0 Hit-Access Counter register
 */

type ExtmemL2Dbus0AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus0_acs_miss_cnt register
 *  L2-Cache bus0 Miss-Access Counter register
 */

type ExtmemL2Dbus0AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus0_acs_conflict_cnt register
 *  L2-Cache bus0 Conflict-Access Counter register
 */

type ExtmemL2Dbus0AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus0_acs_nxtlvl_cnt register
 *  L2-Cache bus0 Next-Level-Access Counter register
 */

type ExtmemL2Dbus0AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus1_acs_hit_cnt register
 *  L2-Cache bus1 Hit-Access Counter register
 */

type ExtmemL2Dbus1AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus1_acs_miss_cnt register
 *  L2-Cache bus1 Miss-Access Counter register
 */

type ExtmemL2Dbus1AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus1_acs_conflict_cnt register
 *  L2-Cache bus1 Conflict-Access Counter register
 */

type ExtmemL2Dbus1AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus1_acs_nxtlvl_cnt register
 *  L2-Cache bus1 Next-Level-Access Counter register
 */

type ExtmemL2Dbus1AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus2_acs_hit_cnt register
 *  L2-Cache bus2 Hit-Access Counter register
 */

type ExtmemL2Dbus2AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus2_acs_miss_cnt register
 *  L2-Cache bus2 Miss-Access Counter register
 */

type ExtmemL2Dbus2AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus2_acs_conflict_cnt register
 *  L2-Cache bus2 Conflict-Access Counter register
 */

type ExtmemL2Dbus2AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus2_acs_nxtlvl_cnt register
 *  L2-Cache bus2 Next-Level-Access Counter register
 */

type ExtmemL2Dbus2AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus3_acs_hit_cnt register
 *  L2-Cache bus3 Hit-Access Counter register
 */

type ExtmemL2Dbus3AcsHitCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus3_acs_miss_cnt register
 *  L2-Cache bus3 Miss-Access Counter register
 */

type ExtmemL2Dbus3AcsMissCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus3_acs_conflict_cnt register
 *  L2-Cache bus3 Conflict-Access Counter register
 */

type ExtmemL2Dbus3AcsConflictCntRegT struct {
	Val c.Uint32T
}

/** Type of l2_dbus3_acs_nxtlvl_cnt register
 *  L2-Cache bus3 Next-Level-Access Counter register
 */

type ExtmemL2Dbus3AcsNxtlvlCntRegT struct {
	Val c.Uint32T
}

/** Group: Access Fail Debug registers */
/** Type of l1_icache0_acs_fail_id_attr register
 *  L1-ICache0 Access Fail ID/attribution information register
 */

type ExtmemL1Icache0AcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache0_acs_fail_addr register
 *  L1-ICache0 Access Fail Address information register
 */

type ExtmemL1Icache0AcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_acs_fail_id_attr register
 *  L1-ICache0 Access Fail ID/attribution information register
 */

type ExtmemL1Icache1AcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache1_acs_fail_addr register
 *  L1-ICache0 Access Fail Address information register
 */

type ExtmemL1Icache1AcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_acs_fail_id_attr register
 *  L1-ICache0 Access Fail ID/attribution information register
 */

type ExtmemL1Icache2AcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache2_acs_fail_addr register
 *  L1-ICache0 Access Fail Address information register
 */

type ExtmemL1Icache2AcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_acs_fail_id_attr register
 *  L1-ICache0 Access Fail ID/attribution information register
 */

type ExtmemL1Icache3AcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l1_icache3_acs_fail_addr register
 *  L1-ICache0 Access Fail Address information register
 */

type ExtmemL1Icache3AcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_acs_fail_id_attr register
 *  L1-Cache Access Fail ID/attribution information register
 */

type ExtmemL1CacheAcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l1_dcache_acs_fail_addr register
 *  L1-Cache Access Fail Address information register
 */

type ExtmemL1DcacheAcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_id_attr register
 *  L2-Cache Access Fail ID/attribution information register
 */

type ExtmemL2CacheAcsFailIdAttrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_acs_fail_addr register
 *  L2-Cache Access Fail Address information register
 */

type ExtmemL2CacheAcsFailAddrRegT struct {
	Val c.Uint32T
}

/** Group: Operation Exception registers */
/** Type of l1_cache_sync_preload_exception register
 *  Cache Sync/Preload Operation exception register
 */

type ExtmemL1CacheSyncPreloadExceptionRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_preload_exception register
 *  Cache Sync/Preload Operation exception register
 */

type ExtmemL2CacheSyncPreloadExceptionRegT struct {
	Val c.Uint32T
}

/** Group: Sync Reset control and configuration registers */
/** Type of l1_cache_sync_rst_ctrl register
 *  Cache Sync Reset control register
 */

type ExtmemL1CacheSyncRstCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_sync_rst_ctrl register
 *  Cache Sync Reset control register
 */

type ExtmemL2CacheSyncRstCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Preload Reset control and configuration registers */
/** Type of l1_cache_preload_rst_ctrl register
 *  Cache Preload Reset control register
 */

type ExtmemL1CachePreloadRstCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_preload_rst_ctrl register
 *  Cache Preload Reset control register
 */

type ExtmemL2CachePreloadRstCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Autoload buffer clear control and configuration registers */
/** Type of l1_cache_autoload_buf_clr_ctrl register
 *  Cache Autoload buffer clear control register
 */

type ExtmemL1CacheAutoloadBufClrCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_autoload_buf_clr_ctrl register
 *  Cache Autoload buffer clear control register
 */

type ExtmemL2CacheAutoloadBufClrCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Unallocate request buffer clear registers */
/** Type of l1_unallocate_buffer_clear register
 *  Unallocate request buffer clear registers
 */

type ExtmemL1UnallocateBufferClearRegT struct {
	Val c.Uint32T
}

/** Type of l2_unallocate_buffer_clear register
 *  Unallocate request buffer clear registers
 */

type ExtmemL2UnallocateBufferClearRegT struct {
	Val c.Uint32T
}

/** Group: Tag and Data Memory Access Control and configuration register */
/** Type of l1_cache_object_ctrl register
 *  Cache Tag and Data memory Object control register
 */

type ExtmemL1CacheObjectCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_way_object register
 *  Cache Tag and Data memory way register
 */

type ExtmemL1CacheWayObjectRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_vaddr register
 *  Cache Vaddr register
 */

type ExtmemL1CacheVaddrRegT struct {
	Val c.Uint32T
}

/** Type of l1_cache_debug_bus register
 *  Cache Tag/data memory content register
 */

type ExtmemL1CacheDebugBusRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_object_ctrl register
 *  Cache Tag and Data memory Object control register
 */

type ExtmemL2CacheObjectCtrlRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_way_object register
 *  Cache Tag and Data memory way register
 */

type ExtmemL2CacheWayObjectRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_vaddr register
 *  Cache Vaddr register
 */

type ExtmemL2CacheVaddrRegT struct {
	Val c.Uint32T
}

/** Type of l2_cache_debug_bus register
 *  Cache Tag/data memory content register
 */

type ExtmemL2CacheDebugBusRegT struct {
	Val c.Uint32T
}

/** Group: Split L1 and L2 registers */
/** Type of level_split0 register
 *  USED TO SPLIT L1 CACHE AND L2 CACHE
 */

type ExtmemLevelSplit0RegT struct {
	Val c.Uint32T
}

/** Type of level_split1 register
 *  USED TO SPLIT L1 CACHE AND L2 CACHE
 */

type ExtmemLevelSplit1RegT struct {
	Val c.Uint32T
}

/** Group: L2 cache access attribute control register */
/** Type of l2_cache_access_attr_ctrl register
 *  L1 Cache access Attribute propagation control register
 */

type ExtmemL2CacheAccessAttrCtrlRegT struct {
	Val c.Uint32T
}

/** Group: Clock Gate Control and configuration register */
/** Type of clock_gate register
 *  Clock gate control register
 */

type ExtmemClockGateRegT struct {
	Val c.Uint32T
}

/** Group: Redundancy  register (Prepare for ECO) */
/** Type of redundancy_sig0 register
 *  Cache redundancy signal 0 register
 */

type ExtmemRedundancySig0RegT struct {
	Val c.Uint32T
}

/** Type of redundancy_sig1 register
 *  Cache redundancy signal 1 register
 */

type ExtmemRedundancySig1RegT struct {
	Val c.Uint32T
}

/** Type of redundancy_sig2 register
 *  Cache redundancy signal 2 register
 */

type ExtmemRedundancySig2RegT struct {
	Val c.Uint32T
}

/** Type of redundancy_sig3 register
 *  Cache redundancy signal 3 register
 */

type ExtmemRedundancySig3RegT struct {
	Val c.Uint32T
}

/** Type of redundancy_sig4 register
 *  Cache redundancy signal 0 register
 */

type ExtmemRedundancySig4RegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version control register
 */

type ExtmemDateRegT struct {
	Val c.Uint32T
}

type ExtmemDevS struct {
	L1IcacheCtrl                ExtmemL1IcacheCtrlRegT
	L1CacheCtrl                 ExtmemL1CacheCtrlRegT
	L1BypassCacheConf           ExtmemL1BypassCacheConfRegT
	L1CacheAtomicConf           ExtmemL1CacheAtomicConfRegT
	L1IcacheCachesizeConf       ExtmemL1IcacheCachesizeConfRegT
	L1IcacheBlocksizeConf       ExtmemL1IcacheBlocksizeConfRegT
	L1CacheCachesizeConf        ExtmemL1CacheCachesizeConfRegT
	L1CacheBlocksizeConf        ExtmemL1CacheBlocksizeConfRegT
	L1CacheWrapAroundCtrl       ExtmemL1CacheWrapAroundCtrlRegT
	L1CacheTagMemPowerCtrl      ExtmemL1CacheTagMemPowerCtrlRegT
	L1CacheDataMemPowerCtrl     ExtmemL1CacheDataMemPowerCtrlRegT
	L1CacheFreezeCtrl           ExtmemL1CacheFreezeCtrlRegT
	L1CacheDataMemAcsConf       ExtmemL1CacheDataMemAcsConfRegT
	L1CacheTagMemAcsConf        ExtmemL1CacheTagMemAcsConfRegT
	L1Icache0PrelockConf        ExtmemL1Icache0PrelockConfRegT
	L1Icache0PrelockSct0Addr    ExtmemL1Icache0PrelockSct0AddrRegT
	L1Icache0PrelockSct1Addr    ExtmemL1Icache0PrelockSct1AddrRegT
	L1Icache0PrelockSctSize     ExtmemL1Icache0PrelockSctSizeRegT
	L1Icache1PrelockConf        ExtmemL1Icache1PrelockConfRegT
	L1Icache1PrelockSct0Addr    ExtmemL1Icache1PrelockSct0AddrRegT
	L1Icache1PrelockSct1Addr    ExtmemL1Icache1PrelockSct1AddrRegT
	L1Icache1PrelockSctSize     ExtmemL1Icache1PrelockSctSizeRegT
	L1Icache2PrelockConf        ExtmemL1Icache2PrelockConfRegT
	L1Icache2PrelockSct0Addr    ExtmemL1Icache2PrelockSct0AddrRegT
	L1Icache2PrelockSct1Addr    ExtmemL1Icache2PrelockSct1AddrRegT
	L1Icache2PrelockSctSize     ExtmemL1Icache2PrelockSctSizeRegT
	L1Icache3PrelockConf        ExtmemL1Icache3PrelockConfRegT
	L1Icache3PrelockSct0Addr    ExtmemL1Icache3PrelockSct0AddrRegT
	L1Icache3PrelockSct1Addr    ExtmemL1Icache3PrelockSct1AddrRegT
	L1Icache3PrelockSctSize     ExtmemL1Icache3PrelockSctSizeRegT
	L1CachePrelockConf          ExtmemL1CachePrelockConfRegT
	L1CachePrelockSct0Addr      ExtmemL1CachePrelockSct0AddrRegT
	L1DcachePrelockSct1Addr     ExtmemL1DcachePrelockSct1AddrRegT
	L1DcachePrelockSctSize      ExtmemL1DcachePrelockSctSizeRegT
	CacheLockCtrl               ExtmemCacheLockCtrlRegT
	CacheLockMap                ExtmemCacheLockMapRegT
	CacheLockAddr               ExtmemCacheLockAddrRegT
	CacheLockSize               ExtmemCacheLockSizeRegT
	CacheSyncCtrl               ExtmemCacheSyncCtrlRegT
	CacheSyncMap                ExtmemCacheSyncMapRegT
	CacheSyncAddr               ExtmemCacheSyncAddrRegT
	CacheSyncSize               ExtmemCacheSyncSizeRegT
	L1Icache0PreloadCtrl        ExtmemL1Icache0PreloadCtrlRegT
	L1Icache0PreloadAddr        ExtmemL1Icache0PreloadAddrRegT
	L1Icache0PreloadSize        ExtmemL1Icache0PreloadSizeRegT
	L1Icache1PreloadCtrl        ExtmemL1Icache1PreloadCtrlRegT
	L1Icache1PreloadAddr        ExtmemL1Icache1PreloadAddrRegT
	L1Icache1PreloadSize        ExtmemL1Icache1PreloadSizeRegT
	L1Icache2PreloadCtrl        ExtmemL1Icache2PreloadCtrlRegT
	L1Icache2PreloadAddr        ExtmemL1Icache2PreloadAddrRegT
	L1Icache2PreloadSize        ExtmemL1Icache2PreloadSizeRegT
	L1Icache3PreloadCtrl        ExtmemL1Icache3PreloadCtrlRegT
	L1Icache3PreloadAddr        ExtmemL1Icache3PreloadAddrRegT
	L1Icache3PreloadSize        ExtmemL1Icache3PreloadSizeRegT
	L1CachePreloadCtrl          ExtmemL1CachePreloadCtrlRegT
	L1DcachePreloadAddr         ExtmemL1DcachePreloadAddrRegT
	L1DcachePreloadSize         ExtmemL1DcachePreloadSizeRegT
	L1Icache0AutoloadCtrl       ExtmemL1Icache0AutoloadCtrlRegT
	L1Icache0AutoloadSct0Addr   ExtmemL1Icache0AutoloadSct0AddrRegT
	L1Icache0AutoloadSct0Size   ExtmemL1Icache0AutoloadSct0SizeRegT
	L1Icache0AutoloadSct1Addr   ExtmemL1Icache0AutoloadSct1AddrRegT
	L1Icache0AutoloadSct1Size   ExtmemL1Icache0AutoloadSct1SizeRegT
	L1Icache1AutoloadCtrl       ExtmemL1Icache1AutoloadCtrlRegT
	L1Icache1AutoloadSct0Addr   ExtmemL1Icache1AutoloadSct0AddrRegT
	L1Icache1AutoloadSct0Size   ExtmemL1Icache1AutoloadSct0SizeRegT
	L1Icache1AutoloadSct1Addr   ExtmemL1Icache1AutoloadSct1AddrRegT
	L1Icache1AutoloadSct1Size   ExtmemL1Icache1AutoloadSct1SizeRegT
	L1Icache2AutoloadCtrl       ExtmemL1Icache2AutoloadCtrlRegT
	L1Icache2AutoloadSct0Addr   ExtmemL1Icache2AutoloadSct0AddrRegT
	L1Icache2AutoloadSct0Size   ExtmemL1Icache2AutoloadSct0SizeRegT
	L1Icache2AutoloadSct1Addr   ExtmemL1Icache2AutoloadSct1AddrRegT
	L1Icache2AutoloadSct1Size   ExtmemL1Icache2AutoloadSct1SizeRegT
	L1Icache3AutoloadCtrl       ExtmemL1Icache3AutoloadCtrlRegT
	L1Icache3AutoloadSct0Addr   ExtmemL1Icache3AutoloadSct0AddrRegT
	L1Icache3AutoloadSct0Size   ExtmemL1Icache3AutoloadSct0SizeRegT
	L1Icache3AutoloadSct1Addr   ExtmemL1Icache3AutoloadSct1AddrRegT
	L1Icache3AutoloadSct1Size   ExtmemL1Icache3AutoloadSct1SizeRegT
	L1CacheAutoloadCtrl         ExtmemL1CacheAutoloadCtrlRegT
	L1CacheAutoloadSct0Addr     ExtmemL1CacheAutoloadSct0AddrRegT
	L1CacheAutoloadSct0Size     ExtmemL1CacheAutoloadSct0SizeRegT
	L1CacheAutoloadSct1Addr     ExtmemL1CacheAutoloadSct1AddrRegT
	L1CacheAutoloadSct1Size     ExtmemL1CacheAutoloadSct1SizeRegT
	L1CacheAutoloadSct2Addr     ExtmemL1CacheAutoloadSct2AddrRegT
	L1CacheAutoloadSct2Size     ExtmemL1CacheAutoloadSct2SizeRegT
	L1CacheAutoloadSct3Addr     ExtmemL1CacheAutoloadSct3AddrRegT
	L1CacheAutoloadSct3Size     ExtmemL1CacheAutoloadSct3SizeRegT
	L1CacheAcsCntIntEna         ExtmemL1CacheAcsCntIntEnaRegT
	L1CacheAcsCntIntClr         ExtmemL1CacheAcsCntIntClrRegT
	L1CacheAcsCntIntRaw         ExtmemL1CacheAcsCntIntRawRegT
	L1CacheAcsCntIntSt          ExtmemL1CacheAcsCntIntStRegT
	L1CacheAcsFailIntEna        ExtmemL1CacheAcsFailIntEnaRegT
	L1CacheAcsFailIntClr        ExtmemL1CacheAcsFailIntClrRegT
	L1CacheAcsFailIntRaw        ExtmemL1CacheAcsFailIntRawRegT
	L1CacheAcsFailIntSt         ExtmemL1CacheAcsFailIntStRegT
	L1CacheAcsCntCtrl           ExtmemL1CacheAcsCntCtrlRegT
	L1Ibus0AcsHitCnt            ExtmemL1Ibus0AcsHitCntRegT
	L1Ibus0AcsMissCnt           ExtmemL1Ibus0AcsMissCntRegT
	L1Ibus0AcsConflictCnt       ExtmemL1Ibus0AcsConflictCntRegT
	L1Ibus0AcsNxtlvlCnt         ExtmemL1Ibus0AcsNxtlvlCntRegT
	L1Ibus1AcsHitCnt            ExtmemL1Ibus1AcsHitCntRegT
	L1Ibus1AcsMissCnt           ExtmemL1Ibus1AcsMissCntRegT
	L1Ibus1AcsConflictCnt       ExtmemL1Ibus1AcsConflictCntRegT
	L1Ibus1AcsNxtlvlCnt         ExtmemL1Ibus1AcsNxtlvlCntRegT
	L1Ibus2AcsHitCnt            ExtmemL1Ibus2AcsHitCntRegT
	L1Ibus2AcsMissCnt           ExtmemL1Ibus2AcsMissCntRegT
	L1Ibus2AcsConflictCnt       ExtmemL1Ibus2AcsConflictCntRegT
	L1Ibus2AcsNxtlvlCnt         ExtmemL1Ibus2AcsNxtlvlCntRegT
	L1Ibus3AcsHitCnt            ExtmemL1Ibus3AcsHitCntRegT
	L1Ibus3AcsMissCnt           ExtmemL1Ibus3AcsMissCntRegT
	L1Ibus3AcsConflictCnt       ExtmemL1Ibus3AcsConflictCntRegT
	L1Ibus3AcsNxtlvlCnt         ExtmemL1Ibus3AcsNxtlvlCntRegT
	L1Bus0AcsHitCnt             ExtmemL1Bus0AcsHitCntRegT
	L1Bus0AcsMissCnt            ExtmemL1Bus0AcsMissCntRegT
	L1Bus0AcsConflictCnt        ExtmemL1Bus0AcsConflictCntRegT
	L1Bus0AcsNxtlvlCnt          ExtmemL1Bus0AcsNxtlvlCntRegT
	L1Bus1AcsHitCnt             ExtmemL1Bus1AcsHitCntRegT
	L1Bus1AcsMissCnt            ExtmemL1Bus1AcsMissCntRegT
	L1Bus1AcsConflictCnt        ExtmemL1Bus1AcsConflictCntRegT
	L1Bus1AcsNxtlvlCnt          ExtmemL1Bus1AcsNxtlvlCntRegT
	L1Dbus2AcsHitCnt            ExtmemL1Dbus2AcsHitCntRegT
	L1Dbus2AcsMissCnt           ExtmemL1Dbus2AcsMissCntRegT
	L1Dbus2AcsConflictCnt       ExtmemL1Dbus2AcsConflictCntRegT
	L1Dbus2AcsNxtlvlCnt         ExtmemL1Dbus2AcsNxtlvlCntRegT
	L1Dbus3AcsHitCnt            ExtmemL1Dbus3AcsHitCntRegT
	L1Dbus3AcsMissCnt           ExtmemL1Dbus3AcsMissCntRegT
	L1Dbus3AcsConflictCnt       ExtmemL1Dbus3AcsConflictCntRegT
	L1Dbus3AcsNxtlvlCnt         ExtmemL1Dbus3AcsNxtlvlCntRegT
	L1Icache0AcsFailIdAttr      ExtmemL1Icache0AcsFailIdAttrRegT
	L1Icache0AcsFailAddr        ExtmemL1Icache0AcsFailAddrRegT
	L1Icache1AcsFailIdAttr      ExtmemL1Icache1AcsFailIdAttrRegT
	L1Icache1AcsFailAddr        ExtmemL1Icache1AcsFailAddrRegT
	L1Icache2AcsFailIdAttr      ExtmemL1Icache2AcsFailIdAttrRegT
	L1Icache2AcsFailAddr        ExtmemL1Icache2AcsFailAddrRegT
	L1Icache3AcsFailIdAttr      ExtmemL1Icache3AcsFailIdAttrRegT
	L1Icache3AcsFailAddr        ExtmemL1Icache3AcsFailAddrRegT
	L1CacheAcsFailIdAttr        ExtmemL1CacheAcsFailIdAttrRegT
	L1DcacheAcsFailAddr         ExtmemL1DcacheAcsFailAddrRegT
	L1CacheSyncPreloadIntEna    ExtmemL1CacheSyncPreloadIntEnaRegT
	L1CacheSyncPreloadIntClr    ExtmemL1CacheSyncPreloadIntClrRegT
	L1CacheSyncPreloadIntRaw    ExtmemL1CacheSyncPreloadIntRawRegT
	L1CacheSyncPreloadIntSt     ExtmemL1CacheSyncPreloadIntStRegT
	L1CacheSyncPreloadException ExtmemL1CacheSyncPreloadExceptionRegT
	L1CacheSyncRstCtrl          ExtmemL1CacheSyncRstCtrlRegT
	L1CachePreloadRstCtrl       ExtmemL1CachePreloadRstCtrlRegT
	L1CacheAutoloadBufClrCtrl   ExtmemL1CacheAutoloadBufClrCtrlRegT
	L1UnallocateBufferClear     ExtmemL1UnallocateBufferClearRegT
	L1CacheObjectCtrl           ExtmemL1CacheObjectCtrlRegT
	L1CacheWayObject            ExtmemL1CacheWayObjectRegT
	L1CacheVaddr                ExtmemL1CacheVaddrRegT
	L1CacheDebugBus             ExtmemL1CacheDebugBusRegT
	LevelSplit0                 ExtmemLevelSplit0RegT
	L2CacheCtrl                 ExtmemL2CacheCtrlRegT
	L2BypassCacheConf           ExtmemL2BypassCacheConfRegT
	L2CacheCachesizeConf        ExtmemL2CacheCachesizeConfRegT
	L2CacheBlocksizeConf        ExtmemL2CacheBlocksizeConfRegT
	L2CacheWrapAroundCtrl       ExtmemL2CacheWrapAroundCtrlRegT
	L2CacheTagMemPowerCtrl      ExtmemL2CacheTagMemPowerCtrlRegT
	L2CacheDataMemPowerCtrl     ExtmemL2CacheDataMemPowerCtrlRegT
	L2CacheFreezeCtrl           ExtmemL2CacheFreezeCtrlRegT
	L2CacheDataMemAcsConf       ExtmemL2CacheDataMemAcsConfRegT
	L2CacheTagMemAcsConf        ExtmemL2CacheTagMemAcsConfRegT
	L2CachePrelockConf          ExtmemL2CachePrelockConfRegT
	L2CachePrelockSct0Addr      ExtmemL2CachePrelockSct0AddrRegT
	L2CachePrelockSct1Addr      ExtmemL2CachePrelockSct1AddrRegT
	L2CachePrelockSctSize       ExtmemL2CachePrelockSctSizeRegT
	L2CachePreloadCtrl          ExtmemL2CachePreloadCtrlRegT
	L2CachePreloadAddr          ExtmemL2CachePreloadAddrRegT
	L2CachePreloadSize          ExtmemL2CachePreloadSizeRegT
	L2CacheAutoloadCtrl         ExtmemL2CacheAutoloadCtrlRegT
	L2CacheAutoloadSct0Addr     ExtmemL2CacheAutoloadSct0AddrRegT
	L2CacheAutoloadSct0Size     ExtmemL2CacheAutoloadSct0SizeRegT
	L2CacheAutoloadSct1Addr     ExtmemL2CacheAutoloadSct1AddrRegT
	L2CacheAutoloadSct1Size     ExtmemL2CacheAutoloadSct1SizeRegT
	L2CacheAutoloadSct2Addr     ExtmemL2CacheAutoloadSct2AddrRegT
	L2CacheAutoloadSct2Size     ExtmemL2CacheAutoloadSct2SizeRegT
	L2CacheAutoloadSct3Addr     ExtmemL2CacheAutoloadSct3AddrRegT
	L2CacheAutoloadSct3Size     ExtmemL2CacheAutoloadSct3SizeRegT
	L2CacheAcsCntIntEna         ExtmemL2CacheAcsCntIntEnaRegT
	L2CacheAcsCntIntClr         ExtmemL2CacheAcsCntIntClrRegT
	L2CacheAcsCntIntRaw         ExtmemL2CacheAcsCntIntRawRegT
	L2CacheAcsCntIntSt          ExtmemL2CacheAcsCntIntStRegT
	L2CacheAcsFailIntEna        ExtmemL2CacheAcsFailIntEnaRegT
	L2CacheAcsFailIntClr        ExtmemL2CacheAcsFailIntClrRegT
	L2CacheAcsFailIntRaw        ExtmemL2CacheAcsFailIntRawRegT
	L2CacheAcsFailIntSt         ExtmemL2CacheAcsFailIntStRegT
	L2CacheAcsCntCtrl           ExtmemL2CacheAcsCntCtrlRegT
	L2Ibus0AcsHitCnt            ExtmemL2Ibus0AcsHitCntRegT
	L2Ibus0AcsMissCnt           ExtmemL2Ibus0AcsMissCntRegT
	L2Ibus0AcsConflictCnt       ExtmemL2Ibus0AcsConflictCntRegT
	L2Ibus0AcsNxtlvlCnt         ExtmemL2Ibus0AcsNxtlvlCntRegT
	L2Ibus1AcsHitCnt            ExtmemL2Ibus1AcsHitCntRegT
	L2Ibus1AcsMissCnt           ExtmemL2Ibus1AcsMissCntRegT
	L2Ibus1AcsConflictCnt       ExtmemL2Ibus1AcsConflictCntRegT
	L2Ibus1AcsNxtlvlCnt         ExtmemL2Ibus1AcsNxtlvlCntRegT
	L2Ibus2AcsHitCnt            ExtmemL2Ibus2AcsHitCntRegT
	L2Ibus2AcsMissCnt           ExtmemL2Ibus2AcsMissCntRegT
	L2Ibus2AcsConflictCnt       ExtmemL2Ibus2AcsConflictCntRegT
	L2Ibus2AcsNxtlvlCnt         ExtmemL2Ibus2AcsNxtlvlCntRegT
	L2Ibus3AcsHitCnt            ExtmemL2Ibus3AcsHitCntRegT
	L2Ibus3AcsMissCnt           ExtmemL2Ibus3AcsMissCntRegT
	L2Ibus3AcsConflictCnt       ExtmemL2Ibus3AcsConflictCntRegT
	L2Ibus3AcsNxtlvlCnt         ExtmemL2Ibus3AcsNxtlvlCntRegT
	L2Dbus0AcsHitCnt            ExtmemL2Dbus0AcsHitCntRegT
	L2Dbus0AcsMissCnt           ExtmemL2Dbus0AcsMissCntRegT
	L2Dbus0AcsConflictCnt       ExtmemL2Dbus0AcsConflictCntRegT
	L2Dbus0AcsNxtlvlCnt         ExtmemL2Dbus0AcsNxtlvlCntRegT
	L2Dbus1AcsHitCnt            ExtmemL2Dbus1AcsHitCntRegT
	L2Dbus1AcsMissCnt           ExtmemL2Dbus1AcsMissCntRegT
	L2Dbus1AcsConflictCnt       ExtmemL2Dbus1AcsConflictCntRegT
	L2Dbus1AcsNxtlvlCnt         ExtmemL2Dbus1AcsNxtlvlCntRegT
	L2Dbus2AcsHitCnt            ExtmemL2Dbus2AcsHitCntRegT
	L2Dbus2AcsMissCnt           ExtmemL2Dbus2AcsMissCntRegT
	L2Dbus2AcsConflictCnt       ExtmemL2Dbus2AcsConflictCntRegT
	L2Dbus2AcsNxtlvlCnt         ExtmemL2Dbus2AcsNxtlvlCntRegT
	L2Dbus3AcsHitCnt            ExtmemL2Dbus3AcsHitCntRegT
	L2Dbus3AcsMissCnt           ExtmemL2Dbus3AcsMissCntRegT
	L2Dbus3AcsConflictCnt       ExtmemL2Dbus3AcsConflictCntRegT
	L2Dbus3AcsNxtlvlCnt         ExtmemL2Dbus3AcsNxtlvlCntRegT
	L2CacheAcsFailIdAttr        ExtmemL2CacheAcsFailIdAttrRegT
	L2CacheAcsFailAddr          ExtmemL2CacheAcsFailAddrRegT
	L2CacheSyncPreloadIntEna    ExtmemL2CacheSyncPreloadIntEnaRegT
	L2CacheSyncPreloadIntClr    ExtmemL2CacheSyncPreloadIntClrRegT
	L2CacheSyncPreloadIntRaw    ExtmemL2CacheSyncPreloadIntRawRegT
	L2CacheSyncPreloadIntSt     ExtmemL2CacheSyncPreloadIntStRegT
	L2CacheSyncPreloadException ExtmemL2CacheSyncPreloadExceptionRegT
	L2CacheSyncRstCtrl          ExtmemL2CacheSyncRstCtrlRegT
	L2CachePreloadRstCtrl       ExtmemL2CachePreloadRstCtrlRegT
	L2CacheAutoloadBufClrCtrl   ExtmemL2CacheAutoloadBufClrCtrlRegT
	L2UnallocateBufferClear     ExtmemL2UnallocateBufferClearRegT
	L2CacheAccessAttrCtrl       ExtmemL2CacheAccessAttrCtrlRegT
	L2CacheObjectCtrl           ExtmemL2CacheObjectCtrlRegT
	L2CacheWayObject            ExtmemL2CacheWayObjectRegT
	L2CacheVaddr                ExtmemL2CacheVaddrRegT
	L2CacheDebugBus             ExtmemL2CacheDebugBusRegT
	LevelSplit1                 ExtmemLevelSplit1RegT
	ClockGate                   ExtmemClockGateRegT
	RedundancySig0              ExtmemRedundancySig0RegT
	RedundancySig1              ExtmemRedundancySig1RegT
	RedundancySig2              ExtmemRedundancySig2RegT
	RedundancySig3              ExtmemRedundancySig3RegT
	RedundancySig4              ExtmemRedundancySig4RegT
	Reserved3c4                 [14]c.Uint32T
	Date                        ExtmemDateRegT
}
type ExtmemDevT ExtmemDevS
