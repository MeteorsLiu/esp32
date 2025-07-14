package esp_gdbstub

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

//go:linkname EspGdbstubPanicHandler C.esp_gdbstub_panic_handler
func EspGdbstubPanicHandler(frame c.Pointer)
