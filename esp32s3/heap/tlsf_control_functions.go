package heap

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TlsfConfig c.Int

const (
	ALIGN_SIZE_LOG2 TlsfConfig = 2
	ALIGN_SIZE      TlsfConfig = 4
)

/* The TLSF control structure. */

type ControlT struct {
	BlockNull        BlockHeaderT
	FlIndexCount     c.Uint
	FlIndexShift     c.Uint
	FlIndexMax       c.Uint
	SlIndexCount     c.Uint
	SlIndexCountLog2 c.Uint
	SmallBlockSize   c.Uint
	Size             c.SizeT
	FlBitmap         c.Uint
	SlBitmap         *c.Uint
	Blocks           **BlockHeaderT
}

/*!
 * @brief Weak function filling the given memory with a given fill pattern.
 *
 * @param start: pointer to the start of the memory region to fill
 * @param size: size of the memory region to fill
 * @param is_free: Indicate if the pattern to use the fill the region should be
 * an after free or after allocation pattern.
 */
//go:linkname BlockAbsorbPostHook C.block_absorb_post_hook
func BlockAbsorbPostHook(start c.Pointer, size c.SizeT, is_free bool)
