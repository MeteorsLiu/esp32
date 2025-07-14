package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const AES_BLOCK_SIZE = 16

type AESTYPE c.Int

const (
	AES_ENC AESTYPE = 0
	AES_DEC AESTYPE = 1
)

type AESBITS c.Int

const (
	AES128 AESBITS = 0
	AES192 AESBITS = 1
	AES256 AESBITS = 2
)
