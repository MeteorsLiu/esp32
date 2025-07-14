package heap

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TlsfT c.Pointer
type PoolT c.Pointer

// llgo:type C
type TlsfWalker func(c.Pointer, c.SizeT, c.Int, c.Pointer) bool
