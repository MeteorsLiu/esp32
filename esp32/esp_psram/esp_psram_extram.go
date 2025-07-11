package esp_psram

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Check if the pointer is on PSRAM
 *
 * @param[in] p  The pointer to check
 *
 * @return
 *        - False: the pointer isn't on PSRAM, or PSRAM isn't initialised successfully
 *        - True:  the pointer is on PSRAM
 */
//go:linkname EspPsramCheckPtrAddr C.esp_psram_check_ptr_addr
func EspPsramCheckPtrAddr(p c.Pointer) bool

/**
 * @brief Add the initialized PSRAM to the heap allocator.
 *
 * @return
 *        - ESP_OK: On success
 *        Other error type, see `heap_caps_add_region`.
 */
//go:linkname EspPsramExtramAddToHeapAllocator C.esp_psram_extram_add_to_heap_allocator
func EspPsramExtramAddToHeapAllocator() EspErrT

/**
 * @brief Reserve a pool of internal memory for specific DMA/internal allocations
 *
 * @param size Size of reserved pool in bytes
 *
 * @return
 *          - ESP_OK:         On success
 *          - ESP_ERR_NO_MEM: When no memory available for pool
 */
//go:linkname EspPsramExtramReserveDmaPool C.esp_psram_extram_reserve_dma_pool
func EspPsramExtramReserveDmaPool(size c.SizeT) EspErrT

/**
 * @brief Memory test for PSRAM. Should be called after PSRAM is initialized and
 * (in case of a dual-core system) the app CPU is online. This test overwrites the
 * memory with crap, so do not call after e.g. the heap allocator has stored important
 * stuff in PSRAM.
 *
 * @return true on success, false on failed memory test
 */
//go:linkname EspPsramExtramTest C.esp_psram_extram_test
func EspPsramExtramTest() bool

/**
 * @brief Init .bss on psram
 */
//go:linkname EspPsramBssInit C.esp_psram_bss_init
func EspPsramBssInit()

/**
 * @brief Calculates the effective PSRAM memory that would be / is added into the heap.
 *
 * @return The size of PSRAM memory that would be / is added into the heap in bytes, or 0 if PSRAM hardware isn't successfully initialized
 * @note The function pre-calculates the effective size of the PSRAM memory that would be added into the heap after performing the XIP or
 *       ext bss and ext noinit considerations, thus, even if the function is called before esp_psram_init(), it will return the final
 *       effective size of the PSRAM memory that would have been added into the heap after esp_psram_init() is performed
 *       instead of the vanilla size of the PSRAM memory.
 *       This function is only available if CONFIG_SPIRAM_PRE_CONFIGURE_MEMORY_PROTECTION is enabled.
 */
//go:linkname EspPsramGetHeapSizeToProtect C.esp_psram_get_heap_size_to_protect
func EspPsramGetHeapSizeToProtect() c.SizeT

/**
 * @brief Force a writeback of the data in the PSRAM cache. This is to be called whenever
 * cache is disabled, because disabling cache on the ESP32 discards the data in the PSRAM
 * cache.
 *
 * This is meant for use from within the SPI flash code.
 */
//go:linkname EspPsramExtramWritebackCache C.esp_psram_extram_writeback_cache
func EspPsramExtramWritebackCache()
