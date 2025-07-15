package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AESBITS c.Int

const (
	AES128 AESBITS = 0
	AES192 AESBITS = 1
	AES256 AESBITS = 2
)
