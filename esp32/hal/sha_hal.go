package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Hashes a single message block
 *
 * @param sha_type          SHA algorithm to hash with
 * @param data_block        Input message to be hashed
 * @param block_word_len    Length of the input message
 * @param first_block       Is this the first block in a message or a continuation?
 */
// llgo:link EspShaType.ShaHalHashBlock C.sha_hal_hash_block
func (recv_ EspShaType) ShaHalHashBlock(data_block c.Pointer, block_word_len c.SizeT, first_block bool) {
}

/**
 * @brief Polls and waits until the SHA engine is idle
 *
 */
//go:linkname ShaHalWaitIdle C.sha_hal_wait_idle
func ShaHalWaitIdle()

/**
 * @brief Reads the current message digest from the SHA engine
 *
 * @param sha_type SHA algorithm used
 * @param digest_state Output buffer to which to read message digest to
 */
// llgo:link EspShaType.ShaHalReadDigest C.sha_hal_read_digest
func (recv_ EspShaType) ShaHalReadDigest(digest_state c.Pointer) {
}
