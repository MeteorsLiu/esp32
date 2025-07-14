package esp_event

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspEventBaseT *c.Char
type EspEventLoopHandleT c.Pointer

// llgo:type C
type EspEventHandlerT func(c.Pointer, EspEventBaseT, c.Int32T, c.Pointer)
type EspEventHandlerInstanceT c.Pointer
