package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_ROM_MD5_DIGEST_LEN = 16

/**
 * \brief          MD5 context structure
 *
 * \warning        MD5 is considered a weak message digest and its use
 *                 constitutes a security risk. We recommend considering
 *                 stronger message digests instead.
 *
 */

type MbedtlsMd5Context struct {
	Total  [2]c.Uint32T
	State  [4]c.Uint32T
	Buffer [64]c.Char
}
type Md5ContextT MbedtlsMd5Context
