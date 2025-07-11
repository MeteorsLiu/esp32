package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspAesStateT c.Int

const (
	ESP_AES_STATE_BUSY EspAesStateT = 0
	ESP_AES_STATE_IDLE EspAesStateT = 1
)
