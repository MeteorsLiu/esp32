package esp_mm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Initialise the MMU MMAP driver
 *
 * This is called once in the IDF startup code. Don't call it in applications
 */
//go:linkname EspMmuMapInit C.esp_mmu_map_init
func EspMmuMapInit()

/**
 * @brief Reserve a consecutive external virtual memory block, with given capabilities and size
 *
 * @note  This private API shall be only called internally during startup stage. DO NOT call
 *        this API in your applications
 *
 * @param[in] size      Size, in bytes, the amount of memory to find
 * @param[in] caps      Bitwise OR of `mmu_mem_caps_t` flags indicating the memory block capability
 * @param[in] target    Target memory type. See `mmu_target_t`
 * @param[out] out_ptr  Pointer to start address of the memory block that is reserved
 *
 * @return
 *        - ESP_OK:              On success
 *        - ESP_ERR_INVALID_ARG: Invalid arguments, could be wrong caps makeup, or null pointer
 *        - ESP_ERR_NOT_FOUND:   Didn't find enough memory with give caps
 */
//go:linkname EspMmuMapReserveBlockWithCaps C.esp_mmu_map_reserve_block_with_caps
func EspMmuMapReserveBlockWithCaps(size c.SizeT, caps MmuMemCapsT, target MmuTargetT, out_ptr *c.Pointer) EspErrT

/*
 * @brief Dump all mapped blocks
 *
 * @return
 *        - ESP_OK
 */
//go:linkname EspMmuMapDumpMappedBlocksPrivate C.esp_mmu_map_dump_mapped_blocks_private
func EspMmuMapDumpMappedBlocksPrivate() EspErrT
