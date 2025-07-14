package heap

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TlsfT c.Pointer
type PoolT c.Pointer

/* Create/destroy a memory pool. */
//go:linkname TlsfCreate C.tlsf_create
func TlsfCreate(mem c.Pointer, max_bytes c.SizeT) TlsfT

//go:linkname TlsfCreateWithPool C.tlsf_create_with_pool
func TlsfCreateWithPool(mem c.Pointer, pool_bytes c.SizeT, max_bytes c.SizeT) TlsfT

//go:linkname TlsfDestroy C.tlsf_destroy
func TlsfDestroy(tlsf TlsfT)

//go:linkname TlsfGetPool C.tlsf_get_pool
func TlsfGetPool(tlsf TlsfT) PoolT

/* Add/remove memory pools. */
//go:linkname TlsfAddPool C.tlsf_add_pool
func TlsfAddPool(tlsf TlsfT, mem c.Pointer, bytes c.SizeT) PoolT

//go:linkname TlsfRemovePool C.tlsf_remove_pool
func TlsfRemovePool(tlsf TlsfT, pool PoolT)

/* malloc/memalign/realloc/free replacements. */
//go:linkname TlsfMalloc C.tlsf_malloc
func TlsfMalloc(tlsf TlsfT, size c.SizeT) c.Pointer

//go:linkname TlsfMemalign C.tlsf_memalign
func TlsfMemalign(tlsf TlsfT, align c.SizeT, size c.SizeT) c.Pointer

//go:linkname TlsfMemalignOffs C.tlsf_memalign_offs
func TlsfMemalignOffs(tlsf TlsfT, align c.SizeT, size c.SizeT, offset c.SizeT) c.Pointer

//go:linkname TlsfMallocAddr C.tlsf_malloc_addr
func TlsfMallocAddr(tlsf TlsfT, size c.SizeT, address c.Pointer) c.Pointer

//go:linkname TlsfRealloc C.tlsf_realloc
func TlsfRealloc(tlsf TlsfT, ptr c.Pointer, size c.SizeT) c.Pointer

//go:linkname TlsfFree C.tlsf_free
func TlsfFree(tlsf TlsfT, ptr c.Pointer)

/* Returns internal block size, not original request size */
//go:linkname TlsfBlockSize C.tlsf_block_size
func TlsfBlockSize(ptr c.Pointer) c.SizeT

/* Overheads/limits of internal structures. */
//go:linkname TlsfSize C.tlsf_size
func TlsfSize(tlsf TlsfT) c.SizeT

//go:linkname TlsfPoolOverhead C.tlsf_pool_overhead
func TlsfPoolOverhead() c.SizeT

//go:linkname TlsfAllocOverhead C.tlsf_alloc_overhead
func TlsfAllocOverhead() c.SizeT

/**
 * @brief Return the allocable size based on the size passed
 * as parameter
 *
 * @param tlsf Pointer to the tlsf structure
 * @param size The allocation size
 * @return size_t The updated allocation size
 */
//go:linkname TlsfFitSize C.tlsf_fit_size
func TlsfFitSize(tlsf TlsfT, size c.SizeT) c.SizeT

// llgo:type C
type TlsfWalker func(c.Pointer, c.SizeT, c.Int, c.Pointer) bool

//go:linkname TlsfWalkPool C.tlsf_walk_pool
func TlsfWalkPool(pool PoolT, walker TlsfWalker, user c.Pointer)

/* Returns nonzero if any internal consistency check fails. */
//go:linkname TlsfCheck C.tlsf_check
func TlsfCheck(tlsf TlsfT) c.Int

//go:linkname TlsfCheckPool C.tlsf_check_pool
func TlsfCheckPool(pool PoolT) c.Int

/**
 * @brief Find the block containing the pointer passed as parameter
 *
 * @param pool The pool into which to look for the block
 * @param ptr The pointer we want to find the containing block of
 * @return void* The pointer to the containing block if found, NULL if not.
 */
//go:linkname TlsfFindContainingBlock C.tlsf_find_containing_block
func TlsfFindContainingBlock(pool PoolT, ptr c.Pointer) c.Pointer

/**
 * @brief Weak function called on every free block of memory allowing the user to implement
 * application specific checks on the memory.
 *
 * @param start The start pointer to the memory of a block
 * @param size The size of the memory in the block
 * @param is_free Set to true when the memory belongs to a free block.
 * False if it belongs to an allocated block.
 * @return true The checks found no inconsistency in the memory
 * @return false The checks in the function highlighted an inconsistency in the memory
 */
//go:linkname TlsfCheckHook C.tlsf_check_hook
func TlsfCheckHook(start c.Pointer, size c.SizeT, is_free bool) bool
