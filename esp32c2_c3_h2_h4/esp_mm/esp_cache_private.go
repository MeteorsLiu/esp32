package esp_mm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Suspend external memory cache
 *
 * @note CPU branch predictor should be disabled before calling this API
 * @note This API is only for internal usage, no thread safety guaranteed
 * @note This API is Non-Nestable and Non-Reentrant
 * @note Call to this API must be followed by a `esp_cache_resume_ext_mem_cache`
 */
//go:linkname EspCacheSuspendExtMemCache C.esp_cache_suspend_ext_mem_cache
func EspCacheSuspendExtMemCache()

/**
 * @brief Resume external memory cache
 *
 * @note This API is only for internal usage, no thread safety guaranteed
 * @note This API is Non-Nestable and Non-Reentrant
 * @note This API must be called after a `esp_cache_suspend_ext_mem_cache`
 */
//go:linkname EspCacheResumeExtMemCache C.esp_cache_resume_ext_mem_cache
func EspCacheResumeExtMemCache()

/**
 * @brief Helper function for malloc a cache aligned data memory buffer
 *
 * @note Now only support 'MALLOC_CAP_INTERNAL', 'MALLOC_CAP_DMA' and 'MALLOC_CAP_SPIRAM'
 *
 * @param[in]  size             Size in bytes, the amount of memory to allocate
 * @param[in]  heap_caps        Flags, see `MALLOC_CAP_x`
 * @param[out] out_ptr          A pointer to the memory allocated successfully
 * @param[out] actual_size      Actual size for allocation in bytes, when the size you specified doesn't meet the cache alignment requirements, this value might be bigger than the size you specified. Set null if you don't care this value.
 *
 * @deprecated This function is deprecated and will be removed in the future.
 *             Use 'heap_caps_malloc' with MALLOC_CAP_CACHE_ALIGNED caps instead
 *
 * @return
 *        - ESP_OK:
 *        - ESP_ERR_INVALID_ARG: Invalid argument
 *        - ESP_ERR_NO_MEM:      No enough memory for allocation
 */
//go:linkname EspCacheAlignedMalloc C.esp_cache_aligned_malloc
func EspCacheAlignedMalloc(size c.SizeT, heap_caps c.Uint32T, out_ptr *c.Pointer, actual_size *c.SizeT) EspErrT

/**
 * @brief Helper function for malloc a cache aligned data memory buffer as preference in decreasing order.
 *
 * @note Now only support 'MALLOC_CAP_INTERNAL', 'MALLOC_CAP_DMA' and 'MALLOC_CAP_SPIRAM'
 *
 * @param[in]  size         Size in bytes, the amount of memory to allocate
 * @param[out] out_ptr      A pointer to the memory allocated successfully
 * @param[out] actual_size  Actual size for allocation in bytes, when the size you specified doesn't meet the cache alignment requirements, this value might be bigger than the size you specified. Set null if you don't care this value.
 * @param[in]  flag_nums    Number of variable parameters
 * @param[in]  spread param The spread params are bitwise OR of Flags, see `MALLOC_CAP_x`. This API prefers to allocate memory with the first parameter. If failed, allocate memory with
 *                          the next parameter. It will try in this order until allocating a chunk of memory successfully
 *                          or fail to allocate memories with any of the parameters.
 *
 * @deprecated This function is deprecated and will be removed in the future.
 *             Use 'heap_caps_malloc_prefer' with MALLOC_CAP_CACHE_ALIGNED caps instead
 *
 * @return
 *        - ESP_OK:
 *        - ESP_ERR_INVALID_ARG: Invalid argument
 *        - ESP_ERR_NO_MEM:      No enough memory for allocation
 */
//go:linkname EspCacheAlignedMallocPrefer C.esp_cache_aligned_malloc_prefer
func EspCacheAlignedMallocPrefer(size c.SizeT, out_ptr *c.Pointer, actual_size *c.SizeT, flag_nums c.SizeT, __llgo_va_list ...interface{}) EspErrT

/**
 * @brief Helper function for calloc a cache aligned data memory buffer
 *
 * @note Now only support 'MALLOC_CAP_INTERNAL', 'MALLOC_CAP_DMA' and 'MALLOC_CAP_SPIRAM'
 *
 * @param[in]  n                Number of continuing chunks of memory to allocate
 * @param[in]  size             Size of one chunk, in bytes
 * @param[in]  heap_caps        Flags, see `MALLOC_CAP_x`
 * @param[out] out_ptr          A pointer to the memory allocated successfully
 * @param[out] actual_size      Actual size for allocation in bytes, when the size you specified doesn't meet the cache alignment requirements, this value might be bigger than the size you specified. Set null if you don't care this value.
 *
 * @deprecated This function is deprecated and will be removed in the future.
 *             Use 'heap_caps_calloc' with MALLOC_CAP_CACHE_ALIGNED caps instead
 *
 * @return
 *        - ESP_OK:
 *        - ESP_ERR_INVALID_ARG: Invalid argument
 *        - ESP_ERR_NO_MEM:      No enough memory for allocation
 */
//go:linkname EspCacheAlignedCalloc C.esp_cache_aligned_calloc
func EspCacheAlignedCalloc(n c.SizeT, size c.SizeT, heap_caps c.Uint32T, out_ptr *c.Pointer, actual_size *c.SizeT) EspErrT

/**
 * @brief Helper function for calloc a cache aligned data memory buffer as preference in decreasing order.
 *
 * @note Now only support 'MALLOC_CAP_INTERNAL', 'MALLOC_CAP_DMA' and 'MALLOC_CAP_SPIRAM'
 *
 * @param[in]  n            Number of continuing chunks of memory to allocate
 * @param[in]  size         Size in bytes, the amount of memory to allocate
 * @param[out] out_ptr      A pointer to the memory allocated successfully
 * @param[out] actual_size  Actual size for allocation in bytes, when the size you specified doesn't meet the cache alignment requirements, this value might be bigger than the size you specified. Set null if you don't care this value.
 * @param[in]  flag_nums    Number of variable parameters
 * @param[in]  spread param The spread params are bitwise OR of Flags, see `MALLOC_CAP_x`. This API prefers to allocate memory with the first parameter. If failed, allocate memory with
 *                          the next parameter. It will try in this order until allocating a chunk of memory successfully
 *                          or fail to allocate memories with any of the parameters.
 *
 * @deprecated This function is deprecated and will be removed in the future.
 *             Use 'heap_caps_calloc_prefer' with MALLOC_CAP_CACHE_ALIGNED caps instead
 *
 * @return
 *        - ESP_OK:
 *        - ESP_ERR_INVALID_ARG: Invalid argument
 *        - ESP_ERR_NO_MEM:      No enough memory for allocation
 */
//go:linkname EspCacheAlignedCallocPrefer C.esp_cache_aligned_calloc_prefer
func EspCacheAlignedCallocPrefer(n c.SizeT, size c.SizeT, out_ptr *c.Pointer, actual_size *c.SizeT, flag_nums c.SizeT, __llgo_va_list ...interface{}) EspErrT

/**
 * @brief Get Cache alignment requirement for data
 *
 * @note Now only support 'MALLOC_CAP_INTERNAL', 'MALLOC_CAP_DMA' and 'MALLOC_CAP_SPIRAM'
 *
 * @param[in]  heap_caps        Flags, see `MALLOC_CAP_x`
 * @param[out] out_alignment    Alignment
 *
 * @return
 *        - ESP_OK:
 *        - ESP_ERR_INVALID_ARG: Invalid argument
 */
//go:linkname EspCacheGetAlignment C.esp_cache_get_alignment
func EspCacheGetAlignment(heap_caps c.Uint32T, out_alignment *c.SizeT) EspErrT
