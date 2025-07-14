package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SHATYPE c.Int

const (
	SHA1         SHATYPE = 0
	SHA2_224     SHATYPE = 1
	SHA2_256     SHATYPE = 2
	SHA_TYPE_MAX SHATYPE = 3
)

type SHAContext struct {
	Start      bool
	InHardware bool
	Type       SHATYPE
	State      [16]c.Uint32T
	Buffer     [128]c.Char
	TotalBits  [4]c.Uint32T
}
type SHACTX SHAContext
