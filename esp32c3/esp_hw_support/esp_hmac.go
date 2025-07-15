package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type HmacKeyIdT c.Int

const (
	HMAC_KEY0    HmacKeyIdT = 0
	HMAC_KEY1    HmacKeyIdT = 1
	HMAC_KEY2    HmacKeyIdT = 2
	HMAC_KEY3    HmacKeyIdT = 3
	HMAC_KEY4    HmacKeyIdT = 4
	HMAC_KEY5    HmacKeyIdT = 5
	HMAC_KEY_MAX HmacKeyIdT = 6
)
