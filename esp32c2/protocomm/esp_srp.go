package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspNgTypeT c.Int

const ESP_NG_3072 EspNgTypeT = 0

type EspSrpHandle struct {
	Unused [8]uint8
}
type EspSrpHandleT EspSrpHandle

/**
 * @brief Initialize srp context for given NG type
 *
 * @param ng NG type given by `esp_ng_type_t`
 * @return   esp_srp_handle_t* srp handle
 *
 * @note    the handle gets freed with `esp_srp_free`
 */
// llgo:link EspNgTypeT.EspSrpInit C.esp_srp_init
func (recv_ EspNgTypeT) EspSrpInit() *EspSrpHandleT {
	return nil
}

/**
 * @brief free esp_srp_context
 *
 * @param hd handle to be free
 */
// llgo:link (*EspSrpHandleT).EspSrpFree C.esp_srp_free
func (recv_ *EspSrpHandleT) EspSrpFree() {
}

/**
 * @brief   Returns B (pub key) and salt. [Step2.b]
 *
 * @param hd    esp_srp handle
 * @param username  Username not expected NULL terminated
 * @param username_len  Username length
 * @param pass      Password not expected to be NULL terminated
 * @param pass_len  Pasword length
 * @param salt_len  Salt length
 * @param bytes_B   Public Key returned
 * @param len_B     Length of the public key
 * @param bytes_salt    Salt bytes generated
 * @return esp_err_t    ESP_OK on success, appropriate error otherwise
 *
 * @note    *bytes_B MUST NOT BE FREED BY THE CALLER
 * @note    *bytes_salt MUST NOT BE FREE BY THE CALLER
 */
// llgo:link (*EspSrpHandleT).EspSrpSrvPubkey C.esp_srp_srv_pubkey
func (recv_ *EspSrpHandleT) EspSrpSrvPubkey(username *c.Char, username_len c.Int, pass *c.Char, pass_len c.Int, salt_len c.Int, bytes_B **c.Char, len_B *c.Int, bytes_salt **c.Char) EspErrT {
	return 0
}

/**
 * @brief   Generate salt-verifier pair, given username, password and salt length
 *
 * @param[in] username      username
 * @param[in] username_len  length of the username
 * @param[in] pass          password
 * @param[in] pass_len      length of the password
 * @param[out] bytes_salt   generated salt on successful generation, or NULL
 * @param[in] salt_len      salt length
 * @param[out] verifier     generated verifier on successful generation, or NULL
 * @param[out] verifier_len length of the generated verifier
 * @return esp_err_t        ESP_OK on success, appropriate error otherwise
 *
 * @note if API has returned ESP_OK, salt and verifier generated need to be freed by caller
 * @note Usually, username and password are not saved on the device. Rather salt and verifier are
 *      generated outside the device and are embedded.
 *      this covenience API can be used to generate salt and verifier on the fly for development use case.
 *      OR for devices which intentionally want to generate different password each time and can send it
 *      to the client securely. e.g., a device has a display and it shows the pin
 */
//go:linkname EspSrpGenSaltVerifier C.esp_srp_gen_salt_verifier
func EspSrpGenSaltVerifier(username *c.Char, username_len c.Int, pass *c.Char, pass_len c.Int, bytes_salt **c.Char, salt_len c.Int, verifier **c.Char, verifier_len *c.Int) EspErrT

/**
 * @brief       Set the Salt and Verifier pre-generated for a given password.
 * This should be used only if the actual password is not available.
 * The public key can then be generated using esp_srp_srv_pubkey_from_salt_verifier()
 * and not esp_srp_srv_pubkey()
 *
 * @param hd            esp_srp_handle
 * @param salt          pre-generated salt bytes
 * @param salt_len      length of the salt bytes
 * @param verifier      pre-generated verifier
 * @param verifier_len  length of the verifier bytes
 * @return esp_err_t    ESP_OK on success, appropriate error otherwise
 */
// llgo:link (*EspSrpHandleT).EspSrpSetSaltVerifier C.esp_srp_set_salt_verifier
func (recv_ *EspSrpHandleT) EspSrpSetSaltVerifier(salt *c.Char, salt_len c.Int, verifier *c.Char, verifier_len c.Int) EspErrT {
	return 0
}

/**
 * @brief   Returns B (pub key)[Step2.b] when the salt and verifier are set using esp_srp_set_salt_verifier()
 *
 * @param hd        esp_srp handle
 * @param bytes_B   Key returned to the called
 * @param len_B     Length of the key returned
 * @return esp_err_t ESP_OK on success, appropriate error otherwise
 *
 * @note    *bytes_B MUST NOT BE FREED BY THE CALLER
 */
// llgo:link (*EspSrpHandleT).EspSrpSrvPubkeyFromSaltVerifier C.esp_srp_srv_pubkey_from_salt_verifier
func (recv_ *EspSrpHandleT) EspSrpSrvPubkeyFromSaltVerifier(bytes_B **c.Char, len_B *c.Int) EspErrT {
	return 0
}

/**
 * @brief   Get session key in `*bytes_key` given by len in `*len_key`. [Step2.c].
 *
 * This calculated session key is used for further communication given the proofs are
 * exchanged/authenticated with `esp_srp_exchange_proofs`
 *
 * @param hd        esp_srp handle
 * @param bytes_A   Private Key
 * @param len_A     Private Key length
 * @param bytes_key Key returned to the caller
 * @param len_key   length of the key in *bytes_key
 * @return esp_err_t ESP_OK on success, appropriate error otherwise
 *
 * @note    *bytes_key MUST NOT BE FREED BY THE CALLER
 */
// llgo:link (*EspSrpHandleT).EspSrpGetSessionKey C.esp_srp_get_session_key
func (recv_ *EspSrpHandleT) EspSrpGetSessionKey(bytes_A *c.Char, len_A c.Int, bytes_key **c.Char, len_key *c.Uint16T) EspErrT {
	return 0
}

/**
 * @brief Complete the authentication. If this step fails, the session_key exchanged should not be used
 *
 * This is the final authentication step in SRP algorithm [Step4.1, Step4.b, Step4.c]
 *
 * @param hd                esp_srp handle
 * @param username          Username not expected NULL terminated
 * @param username_len      Username length
 * @param bytes_user_proof  param in
 * @param bytes_host_proof  parameter out (should be SHA512_DIGEST_LENGTH) bytes in size
 * @return esp_err_t    ESP_OK if user's proof is ok and subsequently bytes_host_proof is populated with our own proof.
 */
// llgo:link (*EspSrpHandleT).EspSrpExchangeProofs C.esp_srp_exchange_proofs
func (recv_ *EspSrpHandleT) EspSrpExchangeProofs(username *c.Char, username_len c.Uint16T, bytes_user_proof *c.Char, bytes_host_proof *c.Char) EspErrT {
	return 0
}
