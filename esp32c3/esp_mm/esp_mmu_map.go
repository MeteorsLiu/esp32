package esp_mm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspPaddrT c.Uint32T

/**
 * @brief Map a physical memory block to external virtual address block, with given capabilities.
 *
 * @param[in]  paddr_start  Start address of the physical memory block
 * @param[in]  size         Size to be mapped. Size will be rounded up by to the nearest multiple of MMU page size
 * @param[in]  target       Physical memory target you're going to map to, see `mmu_target_t`
 * @param[in]  caps         Memory capabilities, see `mmu_mem_caps_t`
 * @param[in]  flags        Mmap flags
 * @param[out] out_ptr      Start address of the mapped virtual memory
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Invalid argument, see printed logs
 *        - ESP_ERR_NOT_SUPPORTED: Only on ESP32, PSRAM is not a supported physical memory target
 *        - ESP_ERR_NOT_FOUND:     No enough size free block to use
 *        - ESP_ERR_NO_MEM:        Out of memory, this API will allocate some heap memory for internal usage
 *        - ESP_ERR_INVALID_STATE: Paddr is mapped already, this API will return corresponding `vaddr_start + new_block_offset` as per the previously mapped block.
 *                                 Only to-be-mapped paddr block is totally enclosed by a previously mapped block will lead to this error. (Identical scenario will behave similarly)
 *                                 new_block_start               new_block_end
 *                                              |-------- New Block --------|
 *                                      |--------------- Block ---------------|
 *                                 block_start                              block_end
 *
 */
// llgo:link EspPaddrT.EspMmuMap C.esp_mmu_map
func (recv_ EspPaddrT) EspMmuMap(size c.SizeT, target MmuTargetT, caps MmuMemCapsT, flags c.Int, out_ptr *c.Pointer) EspErrT {
	return 0
}

/**
 * @brief Unmap a previously mapped virtual memory block
 *
 * @param[in] ptr  Start address of the virtual memory
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Null pointer
 *        - ESP_ERR_NOT_FOUND:     Vaddr is not in external memory, or it's not mapped yet
 */
//go:linkname EspMmuUnmap C.esp_mmu_unmap
func EspMmuUnmap(ptr c.Pointer) EspErrT

/**
 * @brief Get largest consecutive free external virtual memory block size, with given capabilities and given physical target
 *
 * @param[in] caps      Bitwise OR of MMU_MEM_CAP_* flags indicating the memory block
 * @param[in] target    Physical memory target you're going to map to, see `mmu_target_t`.
 * @param[out] out_len  Largest free block length, in bytes.
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG: Invalid arguments, could be null pointer
 */
//go:linkname EspMmuMapGetMaxConsecutiveFreeBlockSize C.esp_mmu_map_get_max_consecutive_free_block_size
func EspMmuMapGetMaxConsecutiveFreeBlockSize(caps MmuMemCapsT, target MmuTargetT, out_len *c.SizeT) EspErrT

/**
 * Dump all the previously mapped blocks
 *
 * @note This API shall not be called from an ISR.
 * @note This API does not guarantee thread safety
 *
 * @param stream stream to print information to; use stdout or stderr to print
 *               to the console; use fmemopen/open_memstream to print to a
 *               string buffer.
 * @return
 *        - ESP_OK
 */
//go:linkname EspMmuMapDumpMappedBlocks C.esp_mmu_map_dump_mapped_blocks
func EspMmuMapDumpMappedBlocks(stream *c.FILE) EspErrT

/**
 * @brief Convert virtual address to physical address
 *
 * @param[in]  vaddr       Virtual address
 * @param[out] out_paddr   Physical address
 * @param[out] out_target  Physical memory target, see `mmu_target_t`
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:  Null pointer, or vaddr is not within external memory
 *        - ESP_ERR_NOT_FOUND:    Vaddr is not mapped yet
 */
//go:linkname EspMmuVaddrToPaddr C.esp_mmu_vaddr_to_paddr
func EspMmuVaddrToPaddr(vaddr c.Pointer, out_paddr *EspPaddrT, out_target *MmuTargetT) EspErrT

/**
 * @brief Convert physical address to virtual address
 *
 * @param[in]  paddr      Physical address
 * @param[in]  target     Physical memory target, see `mmu_target_t`
 * @param[in]  type       Virtual address type, could be either instruction or data
 * @param[out] out_vaddr  Virtual address
 *
 * @return
 *        - ESP_OK
 *        - ESP_ERR_INVALID_ARG:  Null pointer
 *        - ESP_ERR_NOT_FOUND:    Paddr is not mapped yet
 */
// llgo:link EspPaddrT.EspMmuPaddrToVaddr C.esp_mmu_paddr_to_vaddr
func (recv_ EspPaddrT) EspMmuPaddrToVaddr(target MmuTargetT, type_ MmuVaddrT, out_vaddr *c.Pointer) EspErrT {
	return 0
}

/**
 * @brief If the physical address is mapped, this API will provide the capabilities of the virtual address where the physical address is mapped to.
 *
 * @note: Only return value is ESP_OK(which means physically address is successfully mapped), then caps you get make sense.
 * @note This API only check one page (see CONFIG_MMU_PAGE_SIZE), starting from the `paddr`
 *
 * @param[in]  paddr     Physical address
 * @param[out] out_caps  Bitwise OR of MMU_MEM_CAP_* flags indicating the capabilities of a virtual address where the physical address is mapped to.
 * @return
 *      - ESP_OK: Physical address successfully mapped.
 *      - ESP_ERR_INVALID_ARG: Null pointer
 *      - ESP_ERR_NOT_FOUND: Physical address is not mapped successfully.
 */
// llgo:link EspPaddrT.EspMmuPaddrFindCaps C.esp_mmu_paddr_find_caps
func (recv_ EspPaddrT) EspMmuPaddrFindCaps(out_caps *MmuMemCapsT) EspErrT {
	return 0
}
