package esp_system

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type EspUsbConsoleCbT func(c.Pointer)
