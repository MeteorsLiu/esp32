package esp_mm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Adjust size, alignment and caps of a memory allocation request to the specific
 *        hardware requirements, dependent on where the memory gets allocated.
 *
 * @note Note that heap_caps_base.c has its own definition for this function in order not to depend
 *       on this component.
 *
 * @param[in,out] p_alignment Pointer to alignment requirements. This may be modified upwards if the
 *                            hardware has stricter alignment requirements.
 * @param[in,out] p_size      Pointer to size of memory to be allocated. This may be modified upwards
 *                            if e.g. the memory needs to be aligned to a cache line.
 * @param[in,out] p_caps      Pointer to memory requirements. This may be adjusted if the memory
 *                            requirements need modification for the heap caps allocator to work
 *                            properly.
 */
//go:linkname EspHeapAdjustAlignmentToHw C.esp_heap_adjust_alignment_to_hw
func EspHeapAdjustAlignmentToHw(p_alignment *c.SizeT, p_size *c.SizeT, p_caps *c.Uint32T)
