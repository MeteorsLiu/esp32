package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_ROM_MD5_DIGEST_LEN = 16
/**
 * @brief Type defined for MD5 context
 *
 */

type Struct (unnamed at /Users/haolan/esp/esp-idf/components/espRom/include/espRomMd5.h:78:16) struct {
	Buf  [4]c.Uint32T
	Bits [2]c.Uint32T
	In   [64]c.Uint8T
}
type Md5ContextT struct {
	Buf  [4]c.Uint32T
	Bits [2]c.Uint32T
	In   [64]c.Uint8T
}
