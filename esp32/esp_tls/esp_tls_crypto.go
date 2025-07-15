package esp_tls

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Calculate sha1 sum
 * esp-tls abstraction for crypto sha1 API, calculates the sha1 sum(digest) of
 * the data provided in input which is of ilen size and returns
 * a 20 char sha1 sum
 * @param[in]   input   Input array
 * @param[in]   ilen    Length of Input array
 * @param[out]  output  calculated sha1 sum
 *
 * @return
 * mbedtls stack:-
 *              - MBEDTLS_ERR_SHA1_BAD_INPUT_DATA   on BAD INPUT.
 *              -  0 on success.
 * wolfssl stack:-
 *              - -1    on failure.
 *              -  0    on success.
 */
//go:linkname EspCryptoSha1 C.esp_crypto_sha1
func EspCryptoSha1(input *c.Char, ilen c.SizeT, output *c.Char) c.Int

/**
 * @brief Do Base64 encode of the src data
 *
 * @param[in]   dst   destination buffer
 * @param[in]   dlen  length of destination buffer
 * @param[out]  olen  number of bytes written
 * @param[in]   src   src buffer to be encoded
 * @param[in]   slen  src buffer len
 *
 * @return
 * mbedtls stack:-
 *               - MBEDTLS_ERR_BASE64_BUFFER_TOO_SMALL  if buffer is of insufficient size.
 *               -  0   if successful.
 * wolfssl stack:-
 *               - <0   on failure.
 *               -  0   if succcessful.
 */
//go:linkname EspCryptoBase64Encode C.esp_crypto_base64_encode
func EspCryptoBase64Encode(dst *c.Char, dlen c.SizeT, olen *c.SizeT, src *c.Char, slen c.SizeT) c.Int
