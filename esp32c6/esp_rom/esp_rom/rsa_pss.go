package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ETS_SIG_LEN = 384

type EtsRsaPubkeyT struct {
	N     [384]c.Uint8T
	E     c.Uint32T
	Rinv  [384]c.Uint8T
	Mdash c.Uint32T
}
