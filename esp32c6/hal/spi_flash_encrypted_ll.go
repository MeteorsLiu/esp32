package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type FlashEncryptLlTypeT c.Int

const (
	FLASH_ENCRYPTION_MANU FlashEncryptLlTypeT = 0
	PSRAM_ENCRYPTION_MANU FlashEncryptLlTypeT = 1
)
