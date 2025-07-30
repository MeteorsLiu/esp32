package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MD5Context struct {
	Buf  [4]c.Uint32T
	Bits [2]c.Uint32T
	In   [64]c.Uint8T
}
