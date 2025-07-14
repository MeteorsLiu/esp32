package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type PoisonFillPfuncT func(c.Pointer, c.SizeT, bool)

// llgo:type C
type PoisonCheckPfuncT func(c.Pointer, c.SizeT, bool, bool) bool
