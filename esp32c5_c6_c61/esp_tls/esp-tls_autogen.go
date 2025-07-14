package esp_tls

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
)

type PtrdiffT c.Int
type WcharT c.Int

type MaxAlignT struct {
	X__clangMaxAlignNonce1 c.LongLong
	X__clangMaxAlignNonce2 c.Double
}
