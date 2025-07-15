package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_DS_IV_BIT_LEN = 128
const ESP_DS_SIGNATURE_MD_BIT_LEN = 256
const ESP_DS_SIGNATURE_M_PRIME_BIT_LEN = 32
const ESP_DS_SIGNATURE_L_BIT_LEN = 32
const ESP_DS_SIGNATURE_PADDING_BIT_LEN = 64

type EspDsContext struct {
	Unused [8]uint8
}
type EspDsContextT EspDsContext
type EspDigitalSignatureLengthT c.Int

const (
	ESP_DS_RSA_1024 EspDigitalSignatureLengthT = 31
	ESP_DS_RSA_2048 EspDigitalSignatureLengthT = 63
	ESP_DS_RSA_3072 EspDigitalSignatureLengthT = 95
	ESP_DS_RSA_4096 EspDigitalSignatureLengthT = 127
)

/**
 * Encrypted private key data. Recommended to store in flash in this format.
 *
 * @note This struct has to match to one from the ROM code! This documentation is mostly taken from there.
 */

type EspDigitalSignatureData struct {
	RsaLength EspDigitalSignatureLengthT
	Iv        [4]c.Uint32T
	C         [1200]c.Uint8T
}
type EspDsDataT EspDigitalSignatureData

/**
 * Plaintext parameters used by Digital Signature.
 *
 * This is only used for encrypting the RSA parameters by calling esp_ds_encrypt_params().
 * Afterwards, the result can be stored in flash or in other persistent memory.
 * The encryption is a prerequisite step before any signature operation can be done.
 *
 * @note
 * Y, M, Rb, & M_Prime must all be in little endian format.
 */

type EspDsPDataT struct {
	Y      [96]c.Uint32T
	M      [96]c.Uint32T
	Rb     [96]c.Uint32T
	MPrime c.Uint32T
	Length c.Uint32T
}
