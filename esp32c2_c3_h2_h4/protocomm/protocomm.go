package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type ProtocommReqHandlerT func(c.Uint32T, *c.Uint8T, c.SsizeT, **c.Uint8T, *c.SsizeT, c.Pointer) EspErrT

type Protocomm struct {
	Unused [8]uint8
}
type ProtocommT Protocomm

/**
 * @brief   Create a new protocomm instance
 *
 * This API will return a new dynamically allocated protocomm instance
 * with all elements of the protocomm_t structure initialized to NULL.
 *
 * @return
 *  - protocomm_t* : On success
 *  - NULL : No memory for allocating new instance
 */
//go:linkname ProtocommNew C.protocomm_new
func ProtocommNew() *ProtocommT

/**
 * @brief   Delete a protocomm instance
 *
 * This API will deallocate a protocomm instance that was created
 * using `protocomm_new()`.
 *
 * @param[in] pc    Pointer to the protocomm instance to be deleted
 */
// llgo:link (*ProtocommT).ProtocommDelete C.protocomm_delete
func (recv_ *ProtocommT) ProtocommDelete() {
}

/**
 * @brief   Add endpoint request handler for a protocomm instance
 *
 * This API will bind an endpoint handler function to the specified
 * endpoint name, along with any private data that needs to be pass to
 * the handler at the time of call.
 *
 * @note
 *  - An endpoint must be bound to a valid protocomm instance,
 *    created using `protocomm_new()`.
 *  - This function internally calls the registered `add_endpoint()`
 *    function of the selected transport which is a member of the
 *    protocomm_t instance structure.
 *
 * @param[in] pc        Pointer to the protocomm instance
 * @param[in] ep_name   Endpoint identifier(name) string
 * @param[in] h         Endpoint handler function
 * @param[in] priv_data Pointer to private data to be passed as a
 *                      parameter to the handler function on call.
 *                      Pass NULL if not needed.
 *
 * @return
 *  - ESP_OK : Success
 *  - ESP_FAIL : Error adding endpoint / Endpoint with this name already exists
 *  - ESP_ERR_NO_MEM : Error allocating endpoint resource
 *  - ESP_ERR_INVALID_ARG : Null instance/name/handler arguments
 */
// llgo:link (*ProtocommT).ProtocommAddEndpoint C.protocomm_add_endpoint
func (recv_ *ProtocommT) ProtocommAddEndpoint(ep_name *c.Char, h ProtocommReqHandlerT, priv_data c.Pointer) EspErrT {
	return 0
}

/**
 * @brief   Remove endpoint request handler for a protocomm instance
 *
 * This API will remove a registered endpoint handler identified by
 * an endpoint name.
 *
 * @note
 *  - This function internally calls the registered `remove_endpoint()`
 *    function which is a member of the protocomm_t instance structure.
 *
 * @param[in] pc        Pointer to the protocomm instance
 * @param[in] ep_name   Endpoint identifier(name) string
 *
 * @return
 *  - ESP_OK : Success
 *  - ESP_ERR_NOT_FOUND : Endpoint with specified name doesn't exist
 *  - ESP_ERR_INVALID_ARG : Null instance/name arguments
 */
// llgo:link (*ProtocommT).ProtocommRemoveEndpoint C.protocomm_remove_endpoint
func (recv_ *ProtocommT) ProtocommRemoveEndpoint(ep_name *c.Char) EspErrT {
	return 0
}

/**
 * @brief   Allocates internal resources for new transport session
 *
 * @note
 *  - An endpoint must be bound to a valid protocomm instance,
 *    created using `protocomm_new()`.
 *
 * @param[in]  pc         Pointer to the protocomm instance
 * @param[in]  session_id Unique ID for a communication session
 *
 * @return
 *  - ESP_OK : Request handled successfully
 *  - ESP_ERR_NO_MEM : Error allocating internal resource
 *  - ESP_ERR_INVALID_ARG : Null instance/name arguments
 */
// llgo:link (*ProtocommT).ProtocommOpenSession C.protocomm_open_session
func (recv_ *ProtocommT) ProtocommOpenSession(session_id c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief   Frees internal resources used by a transport session
 *
 * @note
 *  - An endpoint must be bound to a valid protocomm instance,
 *    created using `protocomm_new()`.
 *
 * @param[in]  pc         Pointer to the protocomm instance
 * @param[in]  session_id Unique ID for a communication session
 *
 * @return
 *  - ESP_OK : Request handled successfully
 *  - ESP_ERR_INVALID_ARG : Null instance/name arguments
 */
// llgo:link (*ProtocommT).ProtocommCloseSession C.protocomm_close_session
func (recv_ *ProtocommT) ProtocommCloseSession(session_id c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief   Calls the registered handler of an endpoint session
 *          for processing incoming data and generating the response
 *
 * @note
 *  - An endpoint must be bound to a valid protocomm instance,
 *    created using `protocomm_new()`.
 *  - Resulting output buffer must be deallocated by the caller.
 *
 * @param[in]  pc         Pointer to the protocomm instance
 * @param[in]  ep_name    Endpoint identifier(name) string
 * @param[in]  session_id Unique ID for a communication session
 * @param[in]  inbuf      Input buffer contains input request data which is to be
 *                        processed by the registered handler
 * @param[in]  inlen      Length of the input buffer
 * @param[out] outbuf     Pointer to internally allocated output buffer,
 *                        where the resulting response data output from
 *                        the registered handler is to be stored
 * @param[out] outlen     Buffer length of the allocated output buffer
 *
 * @return
 *  - ESP_OK : Request handled successfully
 *  - ESP_FAIL : Internal error in execution of registered handler
 *  - ESP_ERR_NO_MEM : Error allocating internal resource
 *  - ESP_ERR_NOT_FOUND : Endpoint with specified name doesn't exist
 *  - ESP_ERR_INVALID_ARG : Null instance/name arguments
 */
// llgo:link (*ProtocommT).ProtocommReqHandle C.protocomm_req_handle
func (recv_ *ProtocommT) ProtocommReqHandle(ep_name *c.Char, session_id c.Uint32T, inbuf *c.Uint8T, inlen c.SsizeT, outbuf **c.Uint8T, outlen *c.SsizeT) EspErrT {
	return 0
}

/**
 * @brief   Add endpoint security for a protocomm instance
 *
 * This API will bind a security session establisher to the specified
 * endpoint name, along with any proof of possession that may be required
 * for authenticating a session client.
 *
 * @note
 *  - An endpoint must be bound to a valid protocomm instance,
 *    created using `protocomm_new()`.
 *  - The choice of security can be any `protocomm_security_t` instance.
 *    Choices `protocomm_security0` and `protocomm_security1` and `protocomm_security2` are readily available.
 *
 * @param[in] pc            Pointer to the protocomm instance
 * @param[in] ep_name       Endpoint identifier(name) string
 * @param[in] sec           Pointer to endpoint security instance
 * @param[in] sec_params    Pointer to security params (NULL if not needed)
 *                          The pointer should contain the security params struct
 *                          of appropriate security version.
 *                          For protocomm security version 1 and 2
 *                          sec_params should contain pointer to struct of type
 *                          protocomm_security1_params_t and protocmm_security2_params_t respectively.
 *                          The contents of this pointer must be valid till the security session
 *                          has been running and is not closed.
 * @return
 *  - ESP_OK : Success
 *  - ESP_FAIL : Error adding endpoint / Endpoint with this name already exists
 *  - ESP_ERR_INVALID_STATE : Security endpoint already set
 *  - ESP_ERR_NO_MEM : Error allocating endpoint resource
 *  - ESP_ERR_INVALID_ARG : Null instance/name/handler arguments
 */
// llgo:link (*ProtocommT).ProtocommSetSecurity C.protocomm_set_security
func (recv_ *ProtocommT) ProtocommSetSecurity(ep_name *c.Char, sec *ProtocommSecurityT, sec_params c.Pointer) EspErrT {
	return 0
}

/**
 * @brief   Remove endpoint security for a protocomm instance
 *
 * This API will remove a registered security endpoint identified by
 * an endpoint name.
 *
 * @param[in] pc        Pointer to the protocomm instance
 * @param[in] ep_name   Endpoint identifier(name) string
 *
 * @return
 *  - ESP_OK : Success
 *  - ESP_ERR_NOT_FOUND : Endpoint with specified name doesn't exist
 *  - ESP_ERR_INVALID_ARG : Null instance/name arguments
 */
// llgo:link (*ProtocommT).ProtocommUnsetSecurity C.protocomm_unset_security
func (recv_ *ProtocommT) ProtocommUnsetSecurity(ep_name *c.Char) EspErrT {
	return 0
}

/**
 * @brief   Set endpoint for version verification
 *
 * This API can be used for setting an application specific protocol
 * version which can be verified by clients through the endpoint.
 *
 * @note
 *  - An endpoint must be bound to a valid protocomm instance,
 *    created using `protocomm_new()`.

 * @param[in] pc        Pointer to the protocomm instance
 * @param[in] ep_name   Endpoint identifier(name) string
 * @param[in] version   Version identifier(name) string
 *
 * @return
 *  - ESP_OK : Success
 *  - ESP_FAIL : Error adding endpoint / Endpoint with this name already exists
 *  - ESP_ERR_INVALID_STATE : Version endpoint already set
 *  - ESP_ERR_NO_MEM : Error allocating endpoint resource
 *  - ESP_ERR_INVALID_ARG : Null instance/name/handler arguments
 */
// llgo:link (*ProtocommT).ProtocommSetVersion C.protocomm_set_version
func (recv_ *ProtocommT) ProtocommSetVersion(ep_name *c.Char, version *c.Char) EspErrT {
	return 0
}

/**
 * @brief   Remove version verification endpoint from a protocomm instance
 *
 * This API will remove a registered version endpoint identified by
 * an endpoint name.
 *
 * @param[in] pc        Pointer to the protocomm instance
 * @param[in] ep_name   Endpoint identifier(name) string
 *
 * @return
 *  - ESP_OK : Success
 *  - ESP_ERR_NOT_FOUND : Endpoint with specified name doesn't exist
 *  - ESP_ERR_INVALID_ARG : Null instance/name arguments
 */
// llgo:link (*ProtocommT).ProtocommUnsetVersion C.protocomm_unset_version
func (recv_ *ProtocommT) ProtocommUnsetVersion(ep_name *c.Char) EspErrT {
	return 0
}

/**
 * @brief   Get the security version of the protocomm instance
 *
 * This API will return the security version of the protocomm instance.
 *
 * @param[in] pc        Pointer to the protocomm instance
 * @param[out] sec_ver  Pointer to the security version
 * @param[out] sec_patch_ver  Pointer to the security patch version
 *
 * @return
 *  - ESP_OK : Success
 *  - ESP_ERR_INVALID_ARG : Null instance/name arguments
 */
// llgo:link (*ProtocommT).ProtocommGetSecVersion C.protocomm_get_sec_version
func (recv_ *ProtocommT) ProtocommGetSecVersion(sec_ver *c.Int, sec_patch_ver *c.Uint8T) EspErrT {
	return 0
}
