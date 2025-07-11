package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SHAContext struct {
	Start          bool
	TotalInputBits [4]c.Uint32T
}
type SHACTX SHAContext
type SHATYPE c.Int

const (
	SHA1        SHATYPE = 0
	SHA2_256    SHATYPE = 1
	SHA2_384    SHATYPE = 2
	SHA2_512    SHATYPE = 3
	SHA_INVALID SHATYPE = -1
)
