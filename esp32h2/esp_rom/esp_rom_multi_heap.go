package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MultiHeapInfo struct {
	Unused [8]uint8
}
type MultiHeapHandleT *MultiHeapInfo

// llgo:type C
type MultiHeapWalkerCbT func(c.Pointer, c.SizeT, c.Int, c.Pointer) bool

/**
 * @brief Call the tlsf_walk_pool function of the heap given as parameter with
 * the walker function passed as parameter
 *
 * @param heap The heap to traverse
 * @param walker_func The walker to trigger on each block of the heap
 * @param user_data Opaque pointer to user defined data
 */
//go:linkname MultiHeapWalk C.multi_heap_walk
func MultiHeapWalk(heap MultiHeapHandleT, walker_func MultiHeapWalkerCbT, user_data c.Pointer)
