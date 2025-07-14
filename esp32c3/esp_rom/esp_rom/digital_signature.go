package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ETS_DS_MAX_BITS = 3072
const ETS_DS_IV_LEN = 16

/* Encrypted ETS data. Recommended to store in flash in this format.
 */

type EtsDsDataT struct {
	RsaLength c.Uint
	Iv        [16]c.Uint8T
	C         [1200]c.Uint8T
}
type EtsDsResultT c.Int

const (
	ETS_DS_OK              EtsDsResultT = 0
	ETS_DS_INVALID_PARAM   EtsDsResultT = 1
	ETS_DS_INVALID_KEY     EtsDsResultT = 2
	ETS_DS_INVALID_PADDING EtsDsResultT = 3
	ETS_DS_INVALID_DIGEST  EtsDsResultT = 4
)

/*
Plaintext parameters used by Digital Signature.

	Not used for signing with DS peripheral, but can be encrypted
	in-device by calling ets_ds_encrypt_params()
*/
type EtsDsPDataT struct {
	Y      [96]c.Uint32T
	M      [96]c.Uint32T
	Rb     [96]c.Uint32T
	MPrime c.Uint32T
	Length c.Uint32T
}
type EtsDsKeyT c.Int

const (
	ETS_DS_KEY_HMAC EtsDsKeyT = 0
	ETS_DS_KEY_AES  EtsDsKeyT = 1
)
