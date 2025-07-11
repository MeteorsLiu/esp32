package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type CryptoDmaModeT c.Int

const (
	CRYPTO_DMA_AES CryptoDmaModeT = 0
	CRYPTO_DMA_SHA CryptoDmaModeT = 1
)
