package cmock

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CMOCK_VERSION_MAJOR = 2
const CMOCK_VERSION_MINOR = 5
const CMOCK_VERSION_BUILD = 2

/*-------------------------------------------------------
 * Memory API
 *-------------------------------------------------------*/
//go:linkname CMockGutsMemNew C.CMock_Guts_MemNew
func CMockGutsMemNew(size c.SizeT) c.SizeT

//go:linkname CMockGutsMemChain C.CMock_Guts_MemChain
func CMockGutsMemChain(root_index c.SizeT, obj_index c.SizeT) c.SizeT

//go:linkname CMockGutsMemNext C.CMock_Guts_MemNext
func CMockGutsMemNext(previous_item_index c.SizeT) c.SizeT

//go:linkname CMockGutsMemEndOfChain C.CMock_Guts_MemEndOfChain
func CMockGutsMemEndOfChain(root_index c.SizeT) c.SizeT

//go:linkname CMockGutsGetAddressFor C.CMock_Guts_GetAddressFor
func CMockGutsGetAddressFor(index c.SizeT) c.Pointer

//go:linkname CMockGutsMemBytesCapacity C.CMock_Guts_MemBytesCapacity
func CMockGutsMemBytesCapacity() c.SizeT

//go:linkname CMockGutsMemBytesFree C.CMock_Guts_MemBytesFree
func CMockGutsMemBytesFree() c.SizeT

//go:linkname CMockGutsMemBytesUsed C.CMock_Guts_MemBytesUsed
func CMockGutsMemBytesUsed() c.SizeT

//go:linkname CMockGutsMemFreeAll C.CMock_Guts_MemFreeAll
func CMockGutsMemFreeAll()

//go:linkname CMockGutsMemFreeFinal C.CMock_Guts_MemFreeFinal
func CMockGutsMemFreeFinal()
