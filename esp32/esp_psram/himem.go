package esp_psram

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_HIMEM_MAPFLAG_RO = 1

type EspHimemRamdataT struct {
	Unused [8]uint8
}
type EspHimemHandleT *EspHimemRamdataT

type EspHimemRangedataT struct {
	Unused [8]uint8
}
type EspHimemRangehandleT *EspHimemRangedataT

/**
 * @brief Allocate a block in high memory
 *
 * @param size Size of the to-be-allocated block, in bytes. Note that this needs to be
 *             a multiple of the external RAM mmu block size (32K).
 * @param[out] handle_out Handle to be returned
 * @returns - ESP_OK if succesful
 *          - ESP_ERR_NO_MEM if out of memory
 *          - ESP_ERR_INVALID_SIZE if size is not a multiple of 32K
 */
//go:linkname EspHimemAlloc C.esp_himem_alloc
func EspHimemAlloc(size c.SizeT, handle_out *EspHimemHandleT) EspErrT

/**
 * @brief Allocate a memory region to map blocks into
 *
 * This allocates a contiguous CPU memory region that can be used to map blocks
 * of physical memory into.
 *
 * @param size Size of the range to be allocated. Note this needs to be a multiple of
 *             the external RAM mmu block size (32K).
 * @param[out] handle_out Handle to be returned
 * @returns - ESP_OK if succesful
 *          - ESP_ERR_NO_MEM if out of memory or address space
 *          - ESP_ERR_INVALID_SIZE if size is not a multiple of 32K
 */
//go:linkname EspHimemAllocMapRange C.esp_himem_alloc_map_range
func EspHimemAllocMapRange(size c.SizeT, handle_out *EspHimemRangehandleT) EspErrT

/**
 * @brief Map a block of high memory into the CPUs address space
 *
 * This effectively makes the block available for read/write operations.
 *
 * @note The region to be mapped needs to have offsets and sizes that are aligned to the
 *       SPI RAM MMU block size (32K)
 *
 * @param handle Handle to the block of memory, as given by esp_himem_alloc
 * @param range Range handle to map the memory in
 * @param ram_offset Offset into the block of physical memory of the block to map
 * @param range_offset Offset into the address range where the block will be mapped
 * @param len Length of region to map
 * @param flags One of ESP_HIMEM_MAPFLAG_*
 * @param[out] out_ptr Pointer to variable to store resulting memory pointer in
 * @returns - ESP_OK if the memory could be mapped
 *          - ESP_ERR_INVALID_ARG if offset, range or len aren't MMU-block-aligned (32K)
 *          - ESP_ERR_INVALID_SIZE if the offsets/lengths don't fit in the allocated memory or range
 *          - ESP_ERR_INVALID_STATE if a block in the selected ram offset/length is already mapped, or
 *                                  if a block in the selected range offset/length already has a mapping.
 */
//go:linkname EspHimemMap C.esp_himem_map
func EspHimemMap(handle EspHimemHandleT, range_ EspHimemRangehandleT, ram_offset c.SizeT, range_offset c.SizeT, len c.SizeT, flags c.Int, out_ptr *c.Pointer) EspErrT

/**
 * @brief Free a block of physical memory
 *
 * This clears out the associated handle making the memory available for re-allocation again.
 * This will only succeed if none of the memory blocks currently have a mapping.
 *
 * @param handle Handle to the block of memory, as given by esp_himem_alloc
 * @returns - ESP_OK if the memory is succesfully freed
 *          - ESP_ERR_INVALID_ARG if the handle still is (partially) mapped
 */
//go:linkname EspHimemFree C.esp_himem_free
func EspHimemFree(handle EspHimemHandleT) EspErrT

/**
 * @brief Free a mapping range
 *
 * This clears out the associated handle making the range available for re-allocation again.
 * This will only succeed if none of the range blocks currently are used for a mapping.
 *
 * @param handle Handle to the range block, as given by esp_himem_alloc_map_range
 * @returns - ESP_OK if the memory is succesfully freed
 *          - ESP_ERR_INVALID_ARG if the handle still is (partially) mapped to
 */
//go:linkname EspHimemFreeMapRange C.esp_himem_free_map_range
func EspHimemFreeMapRange(handle EspHimemRangehandleT) EspErrT

/**
 * @brief Unmap a region
 *
 * @param range Range handle
 * @param ptr Pointer returned by esp_himem_map
 * @param len Length of the block to be unmapped. Must be aligned to the SPI RAM MMU blocksize (32K)
 * @returns - ESP_OK if the memory is succesfully unmapped,
 *          - ESP_ERR_INVALID_ARG if ptr or len are invalid.
 */
//go:linkname EspHimemUnmap C.esp_himem_unmap
func EspHimemUnmap(range_ EspHimemRangehandleT, ptr c.Pointer, len c.SizeT) EspErrT

/**
 * @brief Get total amount of memory under control of himem API
 *
 * @returns Amount of memory, in bytes
 */
//go:linkname EspHimemGetPhysSize C.esp_himem_get_phys_size
func EspHimemGetPhysSize() c.SizeT

/**
 * @brief Get free amount of memory under control of himem API
 *
 * @returns Amount of free memory, in bytes
 */
//go:linkname EspHimemGetFreeSize C.esp_himem_get_free_size
func EspHimemGetFreeSize() c.SizeT

/**
 * @brief Get amount of SPI memory address space needed for bankswitching
 *
 * @note This is also weakly defined in esp32/spiram.c and returns 0 there, so
 *       if no other function in this file is used, no memory is reserved.
 *
 * @returns Amount of reserved area, in bytes
 */
//go:linkname EspHimemReservedAreaSize C.esp_himem_reserved_area_size
func EspHimemReservedAreaSize() c.SizeT
