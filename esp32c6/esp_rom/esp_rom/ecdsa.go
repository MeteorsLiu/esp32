package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ETS_DIGEST_LEN = 32

type ECDSACURVE c.Int

const (
	ECDSA_CURVE_P192 ECDSACURVE = 1
	ECDSA_CURVE_P256 ECDSACURVE = 2
)
