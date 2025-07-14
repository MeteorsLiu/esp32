package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CRC_SIGN_BLOCK_LEN = 1196
const SIG_BLOCK_PADDING = 4096
const ETS_SECURE_BOOT_V2_SIGNATURE_MAGIC = 0xE7
const SECURE_BOOT_NUM_BLOCKS = 1
const MAX_KEY_DIGESTS = 1

type EtsSecureBootSigBlock struct {
	Unused [8]uint8
}
type EtsSecureBootSigBlockT EtsSecureBootSigBlock

type EtsSecureBootSignature struct {
	Unused [8]uint8
}
type EtsSecureBootSignatureT EtsSecureBootSignature

type EtsSecureBootKeyDigests struct {
	Unused [8]uint8
}
type EtsSecureBootKeyDigestsT EtsSecureBootKeyDigests
type EtsSecureBootStatusT c.Int

const (
	SB_SUCCESS EtsSecureBootStatusT = 978999973
	SB_FAILED  EtsSecureBootStatusT = 1966311518
)
