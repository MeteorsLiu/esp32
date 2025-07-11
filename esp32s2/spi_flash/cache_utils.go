package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// Init mutex protecting access to spi_flash_* APIs
//
//go:linkname SpiFlashInitLock C.spi_flash_init_lock
func SpiFlashInitLock()

// Take mutex protecting access to spi_flash_* APIs
//
//go:linkname SpiFlashOpLock C.spi_flash_op_lock
func SpiFlashOpLock()

// Release said mutex
//
//go:linkname SpiFlashOpUnlock C.spi_flash_op_unlock
func SpiFlashOpUnlock()

// Suspend the scheduler on both CPUs, disable cache.
// Contrary to its name this doesn't do anything with interrupts, yet.
// Interrupt disabling capability will be added once we implement
// interrupt allocation API.
//
//go:linkname SpiFlashDisableInterruptsCachesAndOtherCpu C.spi_flash_disable_interrupts_caches_and_other_cpu
func SpiFlashDisableInterruptsCachesAndOtherCpu()

// Enable cache, enable interrupts (to be added in future), resume scheduler
//
//go:linkname SpiFlashEnableInterruptsCachesAndOtherCpu C.spi_flash_enable_interrupts_caches_and_other_cpu
func SpiFlashEnableInterruptsCachesAndOtherCpu()

// Disables non-IRAM interrupt handlers on current CPU and caches on both CPUs.
// This function is implied to be called when other CPU is not running or running code from IRAM.
//
//go:linkname SpiFlashDisableInterruptsCachesAndOtherCpuNoOs C.spi_flash_disable_interrupts_caches_and_other_cpu_no_os
func SpiFlashDisableInterruptsCachesAndOtherCpuNoOs()

// Enable cache, enable interrupts on current CPU.
// This function is implied to be called when other CPU is not running or running code from IRAM.
//
//go:linkname SpiFlashEnableInterruptsCachesNoOs C.spi_flash_enable_interrupts_caches_no_os
func SpiFlashEnableInterruptsCachesNoOs()

// Mark the pages containing a flash region as having been
// erased or written to. This means the flash cache needs
// to be evicted before these pages can be flash_mmap()ed again,
// as they may contain stale data
//
// Only call this while holding spi_flash_op_lock()
// Returns true if cache was flushed, false otherwise
//
//go:linkname SpiFlashCheckAndFlushCache C.spi_flash_check_and_flush_cache
func SpiFlashCheckAndFlushCache(start_addr c.SizeT, length c.SizeT) bool

// config instrcutin cache size and cache block size by menuconfig
//
//go:linkname EspConfigInstructionCacheMode C.esp_config_instruction_cache_mode
func EspConfigInstructionCacheMode()

// config data cache size and cache block size by menuconfig
//
//go:linkname EspConfigDataCacheMode C.esp_config_data_cache_mode
func EspConfigDataCacheMode()

/**
 * @brief  enable cache wrap mode for instruction cache and data cache
 * @param  icache_wrap_enable   enable cache wrap mode for i cache
 * @param  dcache_wrap_enable   enable cache wrap mode for d cache
 * @return ESP_OK on success, ESP_FAIL otherwise
 */
//go:linkname EspEnableCacheWrap C.esp_enable_cache_wrap
func EspEnableCacheWrap(icache_wrap_enable bool, dcache_wrap_enable bool) EspErrT

/** @brief Check at runtime if flash cache is enabled on both CPUs
 *
 * @return true if both CPUs have flash cache enabled, false otherwise.
 */
//go:linkname SpiFlashCacheEnabled C.spi_flash_cache_enabled
func SpiFlashCacheEnabled() bool

/**
 * @brief Re-enable cache for the core defined as cpuid parameter.
 *
 * @param cpuid the core number to enable instruction cache for
 */
//go:linkname SpiFlashEnableCache C.spi_flash_enable_cache
func SpiFlashEnableCache(cpuid c.Uint32T)

/**
 * @brief Suspend the Cache access to external memory, will disable branch predictor if supported.
 *
 * @param cpuid       the core number to enable the cache for, meaning less on shared cache.
 * @param saved_state Cache status hold by hal (Used only on ROM impl. in idf, this param unused)
 */
//go:linkname SpiFlashDisableCache C.spi_flash_disable_cache
func SpiFlashDisableCache(cpuid c.Uint32T, saved_state *c.Uint32T)

/**
 * @brief Resume the Cache access to external memory, will enable branch predictor if supported.
 *
 * @param cpuid       the core number to enable the cache for, meaning less on shared cache.
 * @param saved_state Cache status hold by hal (Used only on ROM impl. in idf, this param unused)
 */
//go:linkname SpiFlashRestoreCache C.spi_flash_restore_cache
func SpiFlashRestoreCache(cpuid c.Uint32T, saved_state c.Uint32T)
