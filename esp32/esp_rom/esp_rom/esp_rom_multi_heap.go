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
