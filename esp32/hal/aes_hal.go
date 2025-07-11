package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Sets the key used for AES encryption/decryption
 *
 * @param key pointer to the key
 * @param key_bytes number of bytes in key
 * @param mode key mode, 0 : decrypt, 1: encrypt
 *
 * @return uint8_t number of key bytes written to hardware, used for fault injection check
 */
//go:linkname AesHalSetkey C.aes_hal_setkey
func AesHalSetkey(key *c.Uint8T, key_bytes c.SizeT, mode c.Int) c.Uint8T

/**
 * @brief encrypts/decrypts a single block
 *
 * @param input_block input block, size of AES_BLOCK_BYTES
 * @param output_block output block, size of AES_BLOCK_BYTES
 */
//go:linkname AesHalTransformBlock C.aes_hal_transform_block
func AesHalTransformBlock(input_block c.Pointer, output_block c.Pointer)
