package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspCryptoDpaSecLevelT c.Int

const (
	ESP_CRYPTO_DPA_SEC_LEVEL_OFF    EspCryptoDpaSecLevelT = 0
	ESP_CRYPTO_DPA_SEC_LEVEL_LOW    EspCryptoDpaSecLevelT = 1
	ESP_CRYPTO_DPA_SEC_LEVEL_MIDDLE EspCryptoDpaSecLevelT = 2
	ESP_CRYPTO_DPA_SEC_LEVEL_HIGH   EspCryptoDpaSecLevelT = 3
)
