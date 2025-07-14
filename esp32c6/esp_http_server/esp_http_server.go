package esp_http_server

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const HTTPD_200 = "200 OK"
const HTTPD_204 = "204 No Content"
const HTTPD_207 = "207 Multi-Status"
const HTTPD_400 = "400 Bad Request"
const HTTPD_404 = "404 Not Found"
const HTTPD_408 = "408 Request Timeout"
const HTTPD_500 = "500 Internal Server Error"
const HTTPD_TYPE_JSON = "application/json"
const HTTPD_TYPE_TEXT = "text/html"
const HTTPD_TYPE_OCTET = "application/octet-stream"

type EspHttpServerEventIdT c.Int

const (
	HTTP_SERVER_EVENT_ERROR        EspHttpServerEventIdT = 0
	HTTP_SERVER_EVENT_START        EspHttpServerEventIdT = 1
	HTTP_SERVER_EVENT_ON_CONNECTED EspHttpServerEventIdT = 2
	HTTP_SERVER_EVENT_ON_HEADER    EspHttpServerEventIdT = 3
	HTTP_SERVER_EVENT_HEADERS_SENT EspHttpServerEventIdT = 4
	HTTP_SERVER_EVENT_ON_DATA      EspHttpServerEventIdT = 5
	HTTP_SERVER_EVENT_SENT_DATA    EspHttpServerEventIdT = 6
	HTTP_SERVER_EVENT_DISCONNECTED EspHttpServerEventIdT = 7
	HTTP_SERVER_EVENT_STOP         EspHttpServerEventIdT = 8
)

/** Argument structure for HTTP_SERVER_EVENT_ON_DATA and HTTP_SERVER_EVENT_SENT_DATA event */

type EspHttpServerEventData struct {
	Fd      c.Int
	DataLen c.Int
}
type HttpdHandleT c.Pointer
type HttpdMethodT HttpMethod

// llgo:type C
type HttpdFreeCtxFnT func(c.Pointer)

// llgo:type C
type HttpdOpenFuncT func(HttpdHandleT, c.Int) EspErrT

// llgo:type C
type HttpdCloseFuncT func(HttpdHandleT, c.Int)

// llgo:type C
type HttpdUriMatchFuncT func(*c.Char, *c.Char, c.SizeT) bool

/**
 * @brief   HTTP Server Configuration Structure
 *
 * @note    Use HTTPD_DEFAULT_CONFIG() to initialize the configuration
 *          to a default value and then modify only those fields that are
 *          specifically determined by the use case.
 */

type HttpdConfig struct {
	TaskPriority             c.Uint
	StackSize                c.SizeT
	CoreId                   BaseTypeT
	TaskCaps                 c.Uint32T
	ServerPort               c.Uint16T
	CtrlPort                 c.Uint16T
	MaxOpenSockets           c.Uint16T
	MaxUriHandlers           c.Uint16T
	MaxRespHeaders           c.Uint16T
	BacklogConn              c.Uint16T
	LruPurgeEnable           bool
	RecvWaitTimeout          c.Uint16T
	SendWaitTimeout          c.Uint16T
	GlobalUserCtx            c.Pointer
	GlobalUserCtxFreeFn      HttpdFreeCtxFnT
	GlobalTransportCtx       c.Pointer
	GlobalTransportCtxFreeFn HttpdFreeCtxFnT
	EnableSoLinger           bool
	LingerTimeout            c.Int
	KeepAliveEnable          bool
	KeepAliveIdle            c.Int
	KeepAliveInterval        c.Int
	KeepAliveCount           c.Int
	OpenFn                   HttpdOpenFuncT
	CloseFn                  HttpdCloseFuncT
	UriMatchFn               HttpdUriMatchFuncT
}
type HttpdConfigT HttpdConfig

/**
 * @brief Starts the web server
 *
 * Create an instance of HTTP server and allocate memory/resources for it
 * depending upon the specified configuration.
 *
 * Example usage:
 * @code{c}
 *
 * //Function for starting the webserver
 * httpd_handle_t start_webserver(void)
 * {
 *      // Generate default configuration
 *      httpd_config_t config = HTTPD_DEFAULT_CONFIG();
 *
 *      // Empty handle to http_server
 *      httpd_handle_t server = NULL;
 *
 *      // Start the httpd server
 *      if (httpd_start(&server, &config) == ESP_OK) {
 *          // Register URI handlers
 *          httpd_register_uri_handler(server, &uri_get);
 *          httpd_register_uri_handler(server, &uri_post);
 *      }
 *      // If server failed to start, handle will be NULL
 *      return server;
 * }
 *
 * @endcode
 *
 * @param[in]  config   Configuration for new instance of the server
 * @param[out] handle   Handle to newly created instance of the server. NULL on error
 * @return
 *  - ESP_OK    : Instance created successfully
 *  - ESP_ERR_INVALID_ARG      : Null argument(s)
 *  - ESP_ERR_HTTPD_ALLOC_MEM  : Failed to allocate memory for instance
 *  - ESP_ERR_HTTPD_TASK       : Failed to launch server task
 */
//go:linkname HttpdStart C.httpd_start
func HttpdStart(handle *HttpdHandleT, config *HttpdConfigT) EspErrT

/**
 * @brief Stops the web server
 *
 * Deallocates memory/resources used by an HTTP server instance and
 * deletes it. Once deleted the handle can no longer be used for accessing
 * the instance.
 *
 * Example usage:
 * @code{c}
 *
 * // Function for stopping the webserver
 * void stop_webserver(httpd_handle_t server)
 * {
 *      // Ensure handle is non NULL
 *      if (server != NULL) {
 *          // Stop the httpd server
 *          httpd_stop(server);
 *      }
 * }
 *
 * @endcode
 *
 * @param[in] handle Handle to server returned by httpd_start
 * @return
 *  - ESP_OK : Server stopped successfully
 *  - ESP_ERR_INVALID_ARG : Handle argument is Null
 */
//go:linkname HttpdStop C.httpd_stop
func HttpdStop(handle HttpdHandleT) EspErrT

/**
 * @brief HTTP Request Data Structure
 */

type HttpdReq struct {
	Handle               HttpdHandleT
	Method               c.Int
	Uri                  [513]c.Char
	ContentLen           c.SizeT
	Aux                  c.Pointer
	UserCtx              c.Pointer
	SessCtx              c.Pointer
	FreeCtx              HttpdFreeCtxFnT
	IgnoreSessCtxChanges bool
}
type HttpdReqT HttpdReq

/**
 * @brief Structure for URI handler
 */

type HttpdUri struct {
	Uri     *c.Char
	Method  HttpdMethodT
	Handler c.Pointer
	UserCtx c.Pointer
}
type HttpdUriT HttpdUri

/**
 * @brief   Registers a URI handler
 *
 * @note    URI handlers can be registered in real time as long as the
 *          server handle is valid.
 *
 * Example usage:
 * @code{c}
 *
 * esp_err_t my_uri_handler(httpd_req_t* req)
 * {
 *     // Recv , Process and Send
 *     ....
 *     ....
 *     ....
 *
 *     // Fail condition
 *     if (....) {
 *         // Return fail to close session //
 *         return ESP_FAIL;
 *     }
 *
 *     // On success
 *     return ESP_OK;
 * }
 *
 * // URI handler structure
 * httpd_uri_t my_uri {
 *     .uri      = "/my_uri/path/xyz",
 *     .method   = HTTPD_GET,
 *     .handler  = my_uri_handler,
 *     .user_ctx = NULL
 * };
 *
 * // Register handler
 * if (httpd_register_uri_handler(server_handle, &my_uri) != ESP_OK) {
 *    // If failed to register handler
 *    ....
 * }
 *
 * @endcode
 *
 * @param[in] handle      handle to HTTPD server instance
 * @param[in] uri_handler pointer to handler that needs to be registered
 *
 * @return
 *  - ESP_OK : On successfully registering the handler
 *  - ESP_ERR_INVALID_ARG : Null arguments
 *  - ESP_ERR_HTTPD_HANDLERS_FULL  : If no slots left for new handler
 *  - ESP_ERR_HTTPD_HANDLER_EXISTS : If handler with same URI and
 *                                   method is already registered
 */
//go:linkname HttpdRegisterUriHandler C.httpd_register_uri_handler
func HttpdRegisterUriHandler(handle HttpdHandleT, uri_handler *HttpdUriT) EspErrT

/**
 * @brief   Unregister a URI handler
 *
 * @param[in] handle    handle to HTTPD server instance
 * @param[in] uri       URI string
 * @param[in] method    HTTP method
 *
 * @return
 *  - ESP_OK : On successfully deregistering the handler
 *  - ESP_ERR_INVALID_ARG : Null arguments
 *  - ESP_ERR_NOT_FOUND   : Handler with specified URI and method not found
 */
//go:linkname HttpdUnregisterUriHandler C.httpd_unregister_uri_handler
func HttpdUnregisterUriHandler(handle HttpdHandleT, uri *c.Char, method HttpdMethodT) EspErrT

/**
 * @brief   Unregister all URI handlers with the specified uri string
 *
 * @param[in] handle   handle to HTTPD server instance
 * @param[in] uri      uri string specifying all handlers that need
 *                     to be deregisterd
 *
 * @return
 *  - ESP_OK : On successfully deregistering all such handlers
 *  - ESP_ERR_INVALID_ARG : Null arguments
 *  - ESP_ERR_NOT_FOUND   : No handler registered with specified uri string
 */
//go:linkname HttpdUnregisterUri C.httpd_unregister_uri
func HttpdUnregisterUri(handle HttpdHandleT, uri *c.Char) EspErrT

type HttpdErrCodeT c.Int

const (
	HTTPD_500_INTERNAL_SERVER_ERROR    HttpdErrCodeT = 0
	HTTPD_501_METHOD_NOT_IMPLEMENTED   HttpdErrCodeT = 1
	HTTPD_505_VERSION_NOT_SUPPORTED    HttpdErrCodeT = 2
	HTTPD_400_BAD_REQUEST              HttpdErrCodeT = 3
	HTTPD_401_UNAUTHORIZED             HttpdErrCodeT = 4
	HTTPD_403_FORBIDDEN                HttpdErrCodeT = 5
	HTTPD_404_NOT_FOUND                HttpdErrCodeT = 6
	HTTPD_405_METHOD_NOT_ALLOWED       HttpdErrCodeT = 7
	HTTPD_408_REQ_TIMEOUT              HttpdErrCodeT = 8
	HTTPD_411_LENGTH_REQUIRED          HttpdErrCodeT = 9
	HTTPD_413_CONTENT_TOO_LARGE        HttpdErrCodeT = 10
	HTTPD_414_URI_TOO_LONG             HttpdErrCodeT = 11
	HTTPD_431_REQ_HDR_FIELDS_TOO_LARGE HttpdErrCodeT = 12
	HTTPD_ERR_CODE_MAX                 HttpdErrCodeT = 13
)

// llgo:type C
type HttpdErrHandlerFuncT func(*HttpdReqT, HttpdErrCodeT) EspErrT

/**
 * @brief  Function for registering HTTP error handlers
 *
 * This function maps a handler function to any supported error code
 * given by `httpd_err_code_t`. See prototype `httpd_err_handler_func_t`
 * above for details.
 *
 * @param[in] handle     HTTP server handle
 * @param[in] error      Error type
 * @param[in] handler_fn User implemented handler function
 *                       (Pass NULL to unset any previously set handler)
 *
 * @return
 *  - ESP_OK : handler registered successfully
 *  - ESP_ERR_INVALID_ARG : invalid error code or server handle
 */
//go:linkname HttpdRegisterErrHandler C.httpd_register_err_handler
func HttpdRegisterErrHandler(handle HttpdHandleT, error HttpdErrCodeT, handler_fn HttpdErrHandlerFuncT) EspErrT

// llgo:type C
type HttpdSendFuncT func(HttpdHandleT, c.Int, *c.Char, c.SizeT, c.Int) c.Int

// llgo:type C
type HttpdRecvFuncT func(HttpdHandleT, c.Int, *c.Char, c.SizeT, c.Int) c.Int

// llgo:type C
type HttpdPendingFuncT func(HttpdHandleT, c.Int) c.Int

/**
 * @brief   Override web server's receive function (by session FD)
 *
 * This function overrides the web server's receive function. This same function is
 * used to read HTTP request packets.
 *
 * @note    This API is supposed to be called either from the context of
 *          - an http session APIs where sockfd is a valid parameter
 *          - a URI handler where sockfd is obtained using httpd_req_to_sockfd()
 *
 * @param[in] hd        HTTPD instance handle
 * @param[in] sockfd    Session socket FD
 * @param[in] recv_func The receive function to be set for this session
 *
 * @return
 *  - ESP_OK : On successfully registering override
 *  - ESP_ERR_INVALID_ARG : Null arguments
 */
//go:linkname HttpdSessSetRecvOverride C.httpd_sess_set_recv_override
func HttpdSessSetRecvOverride(hd HttpdHandleT, sockfd c.Int, recv_func HttpdRecvFuncT) EspErrT

/**
 * @brief   Override web server's send function (by session FD)
 *
 * This function overrides the web server's send function. This same function is
 * used to send out any response to any HTTP request.
 *
 * @note    This API is supposed to be called either from the context of
 *          - an http session APIs where sockfd is a valid parameter
 *          - a URI handler where sockfd is obtained using httpd_req_to_sockfd()
 *
 * @param[in] hd        HTTPD instance handle
 * @param[in] sockfd    Session socket FD
 * @param[in] send_func The send function to be set for this session
 *
 * @return
 *  - ESP_OK : On successfully registering override
 *  - ESP_ERR_INVALID_ARG : Null arguments
 */
//go:linkname HttpdSessSetSendOverride C.httpd_sess_set_send_override
func HttpdSessSetSendOverride(hd HttpdHandleT, sockfd c.Int, send_func HttpdSendFuncT) EspErrT

/**
 * @brief   Override web server's pending function (by session FD)
 *
 * This function overrides the web server's pending function. This function is
 * used to test for pending bytes in a socket.
 *
 * @note    This API is supposed to be called either from the context of
 *          - an http session APIs where sockfd is a valid parameter
 *          - a URI handler where sockfd is obtained using httpd_req_to_sockfd()
 *
 * @param[in] hd           HTTPD instance handle
 * @param[in] sockfd       Session socket FD
 * @param[in] pending_func The receive function to be set for this session
 *
 * @return
 *  - ESP_OK : On successfully registering override
 *  - ESP_ERR_INVALID_ARG : Null arguments
 */
//go:linkname HttpdSessSetPendingOverride C.httpd_sess_set_pending_override
func HttpdSessSetPendingOverride(hd HttpdHandleT, sockfd c.Int, pending_func HttpdPendingFuncT) EspErrT

/**
 * @brief   Start an asynchronous request. This function can be called
 *          in a request handler to get a request copy that can be used on a async thread.
 *
 * @note
 * - This function is necessary in order to handle multiple requests simultaneously.
 * See examples/async_requests for example usage.
 * - You must call httpd_req_async_handler_complete() when you are done with the request.
 *
 * @param[in]   r       The request to create an async copy of
 * @param[out]  out     A newly allocated request which can be used on an async thread
 *
 * @return
 *  - ESP_OK : async request object created
 */
// llgo:link (*HttpdReqT).HttpdReqAsyncHandlerBegin C.httpd_req_async_handler_begin
func (recv_ *HttpdReqT) HttpdReqAsyncHandlerBegin(out **HttpdReqT) EspErrT {
	return 0
}

/**
 * @brief   Mark an asynchronous request as completed. This will
 *  - free the request memory
 *  - relinquish ownership of the underlying socket, so it can be reused.
 *  - allow the http server to close our socket if needed (lru_purge_enable)
 *
 * @note If async requests are not marked completed, eventually the server
 * will no longer accept incoming connections. The server will log a
 * "httpd_accept_conn: error in accept (23)" message if this happens.
 *
 * @param[in]   r   The request to mark async work as completed
 *
 * @return
 *  - ESP_OK : async request was marked completed
 */
// llgo:link (*HttpdReqT).HttpdReqAsyncHandlerComplete C.httpd_req_async_handler_complete
func (recv_ *HttpdReqT) HttpdReqAsyncHandlerComplete() EspErrT {
	return 0
}

/**
 * @brief   Get the Socket Descriptor from the HTTP request
 *
 * This API will return the socket descriptor of the session for
 * which URI handler was executed on reception of HTTP request.
 * This is useful when user wants to call functions that require
 * session socket fd, from within a URI handler, ie. :
 *      httpd_sess_get_ctx(),
 *      httpd_sess_trigger_close(),
 *      httpd_sess_update_lru_counter().
 *
 * @note    This API is supposed to be called only from the context of
 *          a URI handler where httpd_req_t* request pointer is valid.
 *
 * @param[in] r The request whose socket descriptor should be found
 *
 * @return
 *  - Socket descriptor : The socket descriptor for this request
 *  - -1 : Invalid/NULL request pointer
 */
// llgo:link (*HttpdReqT).HttpdReqToSockfd C.httpd_req_to_sockfd
func (recv_ *HttpdReqT) HttpdReqToSockfd() c.Int {
	return 0
}

/**
 * @brief   API to read content data from the HTTP request
 *
 * This API will read HTTP content data from the HTTP request into
 * provided buffer. Use content_len provided in httpd_req_t structure
 * to know the length of data to be fetched. If content_len is too
 * large for the buffer then user may have to make multiple calls to
 * this function, each time fetching 'buf_len' number of bytes,
 * while the pointer to content data is incremented internally by
 * the same number.
 *
 * @note
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid.
 *  - If an error is returned, the URI handler must further return an error.
 *    This will ensure that the erroneous socket is closed and cleaned up by
 *    the web server.
 *  - Presently Chunked Encoding is not supported
 *
 * @param[in] r         The request being responded to
 * @param[in] buf       Pointer to a buffer that the data will be read into
 * @param[in] buf_len   Length of the buffer
 *
 * @return
 *  - Bytes : Number of bytes read into the buffer successfully
 *  - 0     : Buffer length parameter is zero / connection closed by peer
 *  - HTTPD_SOCK_ERR_INVALID  : Invalid arguments
 *  - HTTPD_SOCK_ERR_TIMEOUT  : Timeout/interrupted while calling socket recv()
 *  - HTTPD_SOCK_ERR_FAIL     : Unrecoverable error while calling socket recv()
 */
// llgo:link (*HttpdReqT).HttpdReqRecv C.httpd_req_recv
func (recv_ *HttpdReqT) HttpdReqRecv(buf *c.Char, buf_len c.SizeT) c.Int {
	return 0
}

/**
 * @brief   Search for a field in request headers and
 *          return the string length of it's value
 *
 * @note
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid.
 *  - Once httpd_resp_send() API is called all request headers
 *    are purged, so request headers need be copied into separate
 *    buffers if they are required later.
 *
 * @param[in]  r        The request being responded to
 * @param[in]  field    The header field to be searched in the request
 *
 * @return
 *  - Length    : If field is found in the request URL
 *  - Zero      : Field not found / Invalid request / Null arguments
 */
// llgo:link (*HttpdReqT).HttpdReqGetHdrValueLen C.httpd_req_get_hdr_value_len
func (recv_ *HttpdReqT) HttpdReqGetHdrValueLen(field *c.Char) c.SizeT {
	return 0
}

/**
 * @brief   Get the value string of a field from the request headers
 *
 * @note
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid.
 *  - Once httpd_resp_send() API is called all request headers
 *    are purged, so request headers need be copied into separate
 *    buffers if they are required later.
 *  - If output size is greater than input, then the value is truncated,
 *    accompanied by truncation error as return value.
 *  - Use httpd_req_get_hdr_value_len() to know the right buffer length
 *
 * @param[in]  r        The request being responded to
 * @param[in]  field    The field to be searched in the header
 * @param[out] val      Pointer to the buffer into which the value will be copied if the field is found
 * @param[in]  val_size Size of the user buffer "val"
 *
 * @return
 *  - ESP_OK : Field found in the request header and value string copied
 *  - ESP_ERR_NOT_FOUND          : Key not found
 *  - ESP_ERR_INVALID_ARG        : Null arguments
 *  - ESP_ERR_HTTPD_INVALID_REQ  : Invalid HTTP request pointer
 *  - ESP_ERR_HTTPD_RESULT_TRUNC : Value string truncated
 */
// llgo:link (*HttpdReqT).HttpdReqGetHdrValueStr C.httpd_req_get_hdr_value_str
func (recv_ *HttpdReqT) HttpdReqGetHdrValueStr(field *c.Char, val *c.Char, val_size c.SizeT) EspErrT {
	return 0
}

/**
 * @brief   Get Query string length from the request URL
 *
 * @note    This API is supposed to be called only from the context of
 *          a URI handler where httpd_req_t* request pointer is valid
 *
 * @param[in]  r    The request being responded to
 *
 * @return
 *  - Length    : Query is found in the request URL
 *  - Zero      : Query not found / Null arguments / Invalid request / uri is empty
 */
// llgo:link (*HttpdReqT).HttpdReqGetUrlQueryLen C.httpd_req_get_url_query_len
func (recv_ *HttpdReqT) HttpdReqGetUrlQueryLen() c.SizeT {
	return 0
}

/**
 * @brief   Get Query string from the request URL
 *
 * @note
 *  - Presently, the user can fetch the full URL query string, but decoding
 *    will have to be performed by the user. Request headers can be read using
 *    httpd_req_get_hdr_value_str() to know the 'Content-Type' (eg. Content-Type:
 *    application/x-www-form-urlencoded) and then the appropriate decoding
 *    algorithm needs to be applied.
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid
 *  - If output size is greater than input, then the value is truncated,
 *    accompanied by truncation error as return value
 *  - Prior to calling this function, one can use httpd_req_get_url_query_len()
 *    to know the query string length beforehand and hence allocate the buffer
 *    of right size (usually query string length + 1 for null termination)
 *    for storing the query string
 *
 * @param[in]  r         The request being responded to
 * @param[out] buf       Pointer to the buffer into which the query string will be copied (if found)
 * @param[in]  buf_len   Length of output buffer
 *
 * @return
 *  - ESP_OK : Query is found in the request URL and copied to buffer
 *  - ESP_FAIL                   : uri is empty
 *  - ESP_ERR_NOT_FOUND          : Query not found
 *  - ESP_ERR_INVALID_ARG        : Null arguments
 *  - ESP_ERR_HTTPD_INVALID_REQ  : Invalid HTTP request pointer
 *  - ESP_ERR_HTTPD_RESULT_TRUNC : Query string truncated
 */
// llgo:link (*HttpdReqT).HttpdReqGetUrlQueryStr C.httpd_req_get_url_query_str
func (recv_ *HttpdReqT) HttpdReqGetUrlQueryStr(buf *c.Char, buf_len c.SizeT) EspErrT {
	return 0
}

/**
 * @brief   Helper function to get a URL query tag from a query
 *          string of the type param1=val1&param2=val2
 *
 * @note
 *  - The components of URL query string (keys and values) are not URLdecoded.
 *    The user must check for 'Content-Type' field in the request headers and
 *    then depending upon the specified encoding (URLencoded or otherwise) apply
 *    the appropriate decoding algorithm.
 *  - If actual value size is greater than val_size, then the value is truncated,
 *    accompanied by truncation error as return value.
 *
 * @param[in]  qry       Pointer to query string
 * @param[in]  key       The key to be searched in the query string
 * @param[out] val       Pointer to the buffer into which the value will be copied if the key is found
 * @param[in]  val_size  Size of the user buffer "val"
 *
 * @return
 *  - ESP_OK : Key is found in the URL query string and copied to buffer
 *  - ESP_ERR_NOT_FOUND          : Key not found
 *  - ESP_ERR_INVALID_ARG        : Null arguments
 *  - ESP_ERR_HTTPD_RESULT_TRUNC : Value string truncated
 */
//go:linkname HttpdQueryKeyValue C.httpd_query_key_value
func HttpdQueryKeyValue(qry *c.Char, key *c.Char, val *c.Char, val_size c.SizeT) EspErrT

/**
 * @brief   Get the value string of a cookie value from the "Cookie" request headers by cookie name.
 *
 * @param[in]       req             Pointer to the HTTP request
 * @param[in]       cookie_name     The cookie name to be searched in the request
 * @param[out]      val             Pointer to the buffer into which the value of cookie will be copied if the cookie is found
 * @param[inout]    val_size        Pointer to size of the user buffer "val". This variable will contain cookie length if
 *                                  ESP_OK is returned and required buffer length in case ESP_ERR_HTTPD_RESULT_TRUNC is returned.
 *
 * @return
 *  - ESP_OK : Key is found in the cookie string and copied to buffer
 *  - ESP_ERR_NOT_FOUND          : Key not found
 *  - ESP_ERR_INVALID_ARG        : Null arguments
 *  - ESP_ERR_HTTPD_RESULT_TRUNC : Value string truncated
 *  - ESP_ERR_NO_MEM             : Memory allocation failure
 */
// llgo:link (*HttpdReqT).HttpdReqGetCookieVal C.httpd_req_get_cookie_val
func (recv_ *HttpdReqT) HttpdReqGetCookieVal(cookie_name *c.Char, val *c.Char, val_size *c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Test if a URI matches the given wildcard template.
 *
 * Template may end with '?' to make the previous character optional (typically a slash),
 * '*' for a wildcard match, and '?*' to make the previous character optional, and if present,
 * allow anything to follow.
 *
 * Example:
 *   - * matches everything
 *   - /api/? matches /api and /api/
 *   - /api/\* (sans the backslash) matches /api/ and /api/status, but not /api or /ap
 *   - /api/?* or /api/\*?  (sans the backslash) matches /api/, /api/status, and also /api, but not /apix or /ap
 *
 * The special characters '?' and '*' anywhere else in the template will be taken literally.
 *
 * @param[in] uri_template   URI template (pattern)
 * @param[in] uri_to_match   URI to be matched
 * @param[in] match_upto     how many characters of the URI buffer to test
 *                          (there may be trailing query string etc.)
 *
 * @return true if a match was found
 */
//go:linkname HttpdUriMatchWildcard C.httpd_uri_match_wildcard
func HttpdUriMatchWildcard(uri_template *c.Char, uri_to_match *c.Char, match_upto c.SizeT) bool

/**
 * @brief   API to send a complete HTTP response.
 *
 * This API will send the data as an HTTP response to the request.
 * This assumes that you have the entire response ready in a single
 * buffer. If you wish to send response in incremental chunks use
 * httpd_resp_send_chunk() instead.
 *
 * If no status code and content-type were set, by default this
 * will send 200 OK status code and content type as text/html.
 * You may call the following functions before this API to configure
 * the response headers :
 *      httpd_resp_set_status() - for setting the HTTP status string,
 *      httpd_resp_set_type()   - for setting the Content Type,
 *      httpd_resp_set_hdr()    - for appending any additional field
 *                                value entries in the response header
 *
 * @note
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid.
 *  - Once this API is called, the request has been responded to.
 *  - No additional data can then be sent for the request.
 *  - Once this API is called, all request headers are purged, so
 *    request headers need be copied into separate buffers if
 *    they are required later.
 *
 * @param[in] r         The request being responded to
 * @param[in] buf       Buffer from where the content is to be fetched
 * @param[in] buf_len   Length of the buffer, HTTPD_RESP_USE_STRLEN to use strlen()
 *
 * @return
 *  - ESP_OK : On successfully sending the response packet
 *  - ESP_ERR_INVALID_ARG : Null request pointer
 *  - ESP_ERR_HTTPD_RESP_HDR    : Essential headers are too large for internal buffer
 *  - ESP_ERR_HTTPD_RESP_SEND   : Error in raw send
 *  - ESP_ERR_HTTPD_INVALID_REQ : Invalid request
 */
// llgo:link (*HttpdReqT).HttpdRespSend C.httpd_resp_send
func (recv_ *HttpdReqT) HttpdRespSend(buf *c.Char, buf_len c.SsizeT) EspErrT {
	return 0
}

/**
 * @brief   API to send one HTTP chunk
 *
 * This API will send the data as an HTTP response to the
 * request. This API will use chunked-encoding and send the response
 * in the form of chunks. If you have the entire response contained in
 * a single buffer, please use httpd_resp_send() instead.
 *
 * If no status code and content-type were set, by default this will
 * send 200 OK status code and content type as text/html. You may
 * call the following functions before this API to configure the
 * response headers
 *      httpd_resp_set_status() - for setting the HTTP status string,
 *      httpd_resp_set_type()   - for setting the Content Type,
 *      httpd_resp_set_hdr()    - for appending any additional field
 *                                value entries in the response header
 *
 * @note
 * - This API is supposed to be called only from the context of
 *   a URI handler where httpd_req_t* request pointer is valid.
 * - When you are finished sending all your chunks, you must call
 *   this function with buf_len as 0.
 * - Once this API is called, all request headers are purged, so
 *   request headers need be copied into separate buffers if they
 *   are required later.
 *
 * @param[in] r         The request being responded to
 * @param[in] buf       Pointer to a buffer that stores the data
 * @param[in] buf_len   Length of the buffer, HTTPD_RESP_USE_STRLEN to use strlen()
 *
 * @return
 *  - ESP_OK : On successfully sending the response packet chunk
 *  - ESP_ERR_INVALID_ARG : Null request pointer
 *  - ESP_ERR_HTTPD_RESP_HDR    : Essential headers are too large for internal buffer
 *  - ESP_ERR_HTTPD_RESP_SEND   : Error in raw send
 *  - ESP_ERR_HTTPD_INVALID_REQ : Invalid request pointer
 */
// llgo:link (*HttpdReqT).HttpdRespSendChunk C.httpd_resp_send_chunk
func (recv_ *HttpdReqT) HttpdRespSendChunk(buf *c.Char, buf_len c.SsizeT) EspErrT {
	return 0
}

/**
 * @brief   API to set the HTTP status code
 *
 * This API sets the status of the HTTP response to the value specified.
 * By default, the '200 OK' response is sent as the response.
 *
 * @note
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid.
 *  - This API only sets the status to this value. The status isn't
 *    sent out until any of the send APIs is executed.
 *  - Make sure that the lifetime of the status string is valid till
 *    send function is called.
 *
 * @param[in] r         The request being responded to
 * @param[in] status    The HTTP status code of this response
 *
 * @return
 *  - ESP_OK : On success
 *  - ESP_ERR_INVALID_ARG : Null arguments
 *  - ESP_ERR_HTTPD_INVALID_REQ : Invalid request pointer
 */
// llgo:link (*HttpdReqT).HttpdRespSetStatus C.httpd_resp_set_status
func (recv_ *HttpdReqT) HttpdRespSetStatus(status *c.Char) EspErrT {
	return 0
}

/**
 * @brief   API to set the HTTP content type
 *
 * This API sets the 'Content Type' field of the response.
 * The default content type is 'text/html'.
 *
 * @note
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid.
 *  - This API only sets the content type to this value. The type
 *    isn't sent out until any of the send APIs is executed.
 *  - Make sure that the lifetime of the type string is valid till
 *    send function is called.
 *
 * @param[in] r     The request being responded to
 * @param[in] type  The Content Type of the response
 *
 * @return
 *  - ESP_OK   : On success
 *  - ESP_ERR_INVALID_ARG : Null arguments
 *  - ESP_ERR_HTTPD_INVALID_REQ : Invalid request pointer
 */
// llgo:link (*HttpdReqT).HttpdRespSetType C.httpd_resp_set_type
func (recv_ *HttpdReqT) HttpdRespSetType(type_ *c.Char) EspErrT {
	return 0
}

/**
 * @brief   API to append any additional headers
 *
 * This API sets any additional header fields that need to be sent in the response.
 *
 * @note
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid.
 *  - The header isn't sent out until any of the send APIs is executed.
 *  - The maximum allowed number of additional headers is limited to
 *    value of max_resp_headers in config structure.
 *  - Make sure that the lifetime of the field value strings are valid till
 *    send function is called.
 *
 * @param[in] r     The request being responded to
 * @param[in] field The field name of the HTTP header
 * @param[in] value The value of this HTTP header
 *
 * @return
 *  - ESP_OK : On successfully appending new header
 *  - ESP_ERR_INVALID_ARG : Null arguments
 *  - ESP_ERR_HTTPD_RESP_HDR    : Total additional headers exceed max allowed
 *  - ESP_ERR_HTTPD_INVALID_REQ : Invalid request pointer
 */
// llgo:link (*HttpdReqT).HttpdRespSetHdr C.httpd_resp_set_hdr
func (recv_ *HttpdReqT) HttpdRespSetHdr(field *c.Char, value *c.Char) EspErrT {
	return 0
}

/**
 * @brief   For sending out error code in response to HTTP request.
 *
 * @note
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid.
 *  - Once this API is called, all request headers are purged, so
 *    request headers need be copied into separate buffers if
 *    they are required later.
 *  - If you wish to send additional data in the body of the
 *    response, please use the lower-level functions directly.
 *
 * @param[in] req     Pointer to the HTTP request for which the response needs to be sent
 * @param[in] error   Error type to send
 * @param[in] msg     Error message string (pass NULL for default message)
 *
 * @return
 *  - ESP_OK : On successfully sending the response packet
 *  - ESP_ERR_INVALID_ARG : Null arguments
 *  - ESP_ERR_HTTPD_RESP_SEND   : Error in raw send
 *  - ESP_ERR_HTTPD_INVALID_REQ : Invalid request pointer
 */
// llgo:link (*HttpdReqT).HttpdRespSendErr C.httpd_resp_send_err
func (recv_ *HttpdReqT) HttpdRespSendErr(error HttpdErrCodeT, msg *c.Char) EspErrT {
	return 0
}

/**
 * @brief   For sending out custom error code in response to HTTP request.
 *
 * @note
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid.
 *  - Once this API is called, all request headers are purged, so
 *    request headers need be copied into separate buffers if
 *    they are required later.
 *  - If you wish to send additional data in the body of the
 *    response, please use the lower-level functions directly.
 *
 * @param[in] req     Pointer to the HTTP request for which the response needs to be sent
 * @param[in] status  Error status to send
 * @param[in] msg     Error message string
 *
 * @return
 *  - ESP_OK : On successfully sending the response packet
 *  - ESP_ERR_INVALID_ARG : Null arguments
 *  - ESP_ERR_HTTPD_RESP_SEND   : Error in raw send
 *  - ESP_ERR_HTTPD_INVALID_REQ : Invalid request pointer
 */
// llgo:link (*HttpdReqT).HttpdRespSendCustomErr C.httpd_resp_send_custom_err
func (recv_ *HttpdReqT) HttpdRespSendCustomErr(status *c.Char, msg *c.Char) EspErrT {
	return 0
}

/**
 * @brief   Raw HTTP send
 *
 * Call this API if you wish to construct your custom response packet.
 * When using this, all essential header, eg. HTTP version, Status Code,
 * Content Type and Length, Encoding, etc. will have to be constructed
 * manually, and HTTP delimiters (CRLF) will need to be placed correctly
 * for separating sub-sections of the HTTP response packet.
 *
 * If the send override function is set, this API will end up
 * calling that function eventually to send data out.
 *
 * @note
 *  - This API is supposed to be called only from the context of
 *    a URI handler where httpd_req_t* request pointer is valid.
 *  - Unless the response has the correct HTTP structure (which the
 *    user must now ensure) it is not guaranteed that it will be
 *    recognized by the client. For most cases, you wouldn't have
 *    to call this API, but you would rather use either of :
 *          httpd_resp_send(),
 *          httpd_resp_send_chunk()
 *
 * @param[in] r         The request being responded to
 * @param[in] buf       Buffer from where the fully constructed packet is to be read
 * @param[in] buf_len   Length of the buffer
 *
 * @return
 *  - Bytes : Number of bytes that were sent successfully
 *  - HTTPD_SOCK_ERR_INVALID  : Invalid arguments
 *  - HTTPD_SOCK_ERR_TIMEOUT  : Timeout/interrupted while calling socket send()
 *  - HTTPD_SOCK_ERR_FAIL     : Unrecoverable error while calling socket send()
 */
// llgo:link (*HttpdReqT).HttpdSend C.httpd_send
func (recv_ *HttpdReqT) HttpdSend(buf *c.Char, buf_len c.SizeT) c.Int {
	return 0
}

/**
 * A low level API to send data on a given socket
 *
 * @note This API is not recommended to be used in any request handler.
 * Use this only for advanced use cases, wherein some asynchronous
 * data is to be sent over a socket.
 *
 * This internally calls the default send function, or the function registered by
 * httpd_sess_set_send_override().
 *
 * @param[in] hd        server instance
 * @param[in] sockfd    session socket file descriptor
 * @param[in] buf       buffer with bytes to send
 * @param[in] buf_len   data size
 * @param[in] flags     flags for the send() function
 * @return
 *  - Bytes : The number of bytes sent successfully
 *  - HTTPD_SOCK_ERR_INVALID  : Invalid arguments
 *  - HTTPD_SOCK_ERR_TIMEOUT  : Timeout/interrupted while calling socket send()
 *  - HTTPD_SOCK_ERR_FAIL     : Unrecoverable error while calling socket send()
 */
//go:linkname HttpdSocketSend C.httpd_socket_send
func HttpdSocketSend(hd HttpdHandleT, sockfd c.Int, buf *c.Char, buf_len c.SizeT, flags c.Int) c.Int

/**
 * A low level API to receive data from a given socket
 *
 * @note This API is not recommended to be used in any request handler.
 * Use this only for advanced use cases, wherein some asynchronous
 * communication is required.
 *
 * This internally calls the default recv function, or the function registered by
 * httpd_sess_set_recv_override().
 *
 * @param[in] hd        server instance
 * @param[in] sockfd    session socket file descriptor
 * @param[in] buf       buffer with bytes to send
 * @param[in] buf_len   data size
 * @param[in] flags     flags for the send() function
 * @return
 *  - Bytes : The number of bytes received successfully
 *  - 0     : Buffer length parameter is zero / connection closed by peer
 *  - HTTPD_SOCK_ERR_INVALID  : Invalid arguments
 *  - HTTPD_SOCK_ERR_TIMEOUT  : Timeout/interrupted while calling socket recv()
 *  - HTTPD_SOCK_ERR_FAIL     : Unrecoverable error while calling socket recv()
 */
//go:linkname HttpdSocketRecv C.httpd_socket_recv
func HttpdSocketRecv(hd HttpdHandleT, sockfd c.Int, buf *c.Char, buf_len c.SizeT, flags c.Int) c.Int

/**
 * @brief   Get session context from socket descriptor
 *
 * Typically if a session context is created, it is available to URI handlers
 * through the httpd_req_t structure. But, there are cases where the web
 * server's send/receive functions may require the context (for example, for
 * accessing keying information etc). Since the send/receive function only have
 * the socket descriptor at their disposal, this API provides them with a way to
 * retrieve the session context.
 *
 * @param[in] handle    Handle to server returned by httpd_start
 * @param[in] sockfd    The socket descriptor for which the context should be extracted.
 *
 * @return
 *  - void* : Pointer to the context associated with this session
 *  - NULL  : Empty context / Invalid handle / Invalid socket fd
 */
//go:linkname HttpdSessGetCtx C.httpd_sess_get_ctx
func HttpdSessGetCtx(handle HttpdHandleT, sockfd c.Int) c.Pointer

/**
 * @brief   Set session context by socket descriptor
 *
 * @param[in] handle    Handle to server returned by httpd_start
 * @param[in] sockfd    The socket descriptor for which the context should be extracted.
 * @param[in] ctx       Context object to assign to the session
 * @param[in] free_fn   Function that should be called to free the context
 */
//go:linkname HttpdSessSetCtx C.httpd_sess_set_ctx
func HttpdSessSetCtx(handle HttpdHandleT, sockfd c.Int, ctx c.Pointer, free_fn HttpdFreeCtxFnT)

/**
 * @brief   Get session 'transport' context by socket descriptor
 * @see     httpd_sess_get_ctx()
 *
 * This context is used by the send/receive functions, for example to manage SSL context.
 *
 * @param[in] handle    Handle to server returned by httpd_start
 * @param[in] sockfd    The socket descriptor for which the context should be extracted.
 * @return
 *  - void* : Pointer to the transport context associated with this session
 *  - NULL  : Empty context / Invalid handle / Invalid socket fd
 */
//go:linkname HttpdSessGetTransportCtx C.httpd_sess_get_transport_ctx
func HttpdSessGetTransportCtx(handle HttpdHandleT, sockfd c.Int) c.Pointer

/**
 * @brief   Set session 'transport' context by socket descriptor
 * @see     httpd_sess_set_ctx()
 *
 * @param[in] handle    Handle to server returned by httpd_start
 * @param[in] sockfd    The socket descriptor for which the context should be extracted.
 * @param[in] ctx       Transport context object to assign to the session
 * @param[in] free_fn   Function that should be called to free the transport context
 */
//go:linkname HttpdSessSetTransportCtx C.httpd_sess_set_transport_ctx
func HttpdSessSetTransportCtx(handle HttpdHandleT, sockfd c.Int, ctx c.Pointer, free_fn HttpdFreeCtxFnT)

/**
 * @brief   Get HTTPD global user context (it was set in the server config struct)
 *
 * @param[in] handle    Handle to server returned by httpd_start
 * @return global user context
 */
//go:linkname HttpdGetGlobalUserCtx C.httpd_get_global_user_ctx
func HttpdGetGlobalUserCtx(handle HttpdHandleT) c.Pointer

/**
 * @brief   Get HTTPD global transport context (it was set in the server config struct)
 *
 * @param[in] handle    Handle to server returned by httpd_start
 * @return global transport context
 */
//go:linkname HttpdGetGlobalTransportCtx C.httpd_get_global_transport_ctx
func HttpdGetGlobalTransportCtx(handle HttpdHandleT) c.Pointer

/**
 * @brief   Trigger an httpd session close externally
 *
 * @note    Calling this API is only required in special circumstances wherein
 *          some application requires to close an httpd client session asynchronously.
 *
 * @param[in] handle    Handle to server returned by httpd_start
 * @param[in] sockfd    The socket descriptor of the session to be closed
 *
 * @return
 *  - ESP_OK    : On successfully initiating closure
 *  - ESP_FAIL  : Failure to queue work
 *  - ESP_ERR_NOT_FOUND   : Socket fd not found
 *  - ESP_ERR_INVALID_ARG : Null arguments
 */
//go:linkname HttpdSessTriggerClose C.httpd_sess_trigger_close
func HttpdSessTriggerClose(handle HttpdHandleT, sockfd c.Int) EspErrT

/**
 * @brief   Update LRU counter for a given socket
 *
 * LRU Counters are internally associated with each session to monitor
 * how recently a session exchanged traffic. When LRU purge is enabled,
 * if a client is requesting for connection but maximum number of
 * sockets/sessions is reached, then the session having the earliest
 * LRU counter is closed automatically.
 *
 * Updating the LRU counter manually prevents the socket from being purged
 * due to the Least Recently Used (LRU) logic, even though it might not
 * have received traffic for some time. This is useful when all open
 * sockets/session are frequently exchanging traffic but the user specifically
 * wants one of the sessions to be kept open, irrespective of when it last
 * exchanged a packet.
 *
 * @note    Calling this API is only necessary if the LRU Purge Enable option
 *          is enabled.
 *
 * @param[in] handle    Handle to server returned by httpd_start
 * @param[in] sockfd    The socket descriptor of the session for which LRU counter
 *                      is to be updated
 *
 * @return
 *  - ESP_OK : Socket found and LRU counter updated
 *  - ESP_ERR_NOT_FOUND   : Socket not found
 *  - ESP_ERR_INVALID_ARG : Null arguments
 */
//go:linkname HttpdSessUpdateLruCounter C.httpd_sess_update_lru_counter
func HttpdSessUpdateLruCounter(handle HttpdHandleT, sockfd c.Int) EspErrT

/**
 * @brief   Returns list of current socket descriptors of active sessions
 *
 * @param[in] handle    Handle to server returned by httpd_start
 * @param[in,out] fds   In: Size of provided client_fds array
 *                      Out: Number of valid client fds returned in client_fds,
 * @param[out] client_fds  Array of client fds
 *
 * @note Size of provided array has to be equal or greater then maximum number of opened
 *       sockets, configured upon initialization with max_open_sockets field in
 *       httpd_config_t structure.
 *
 * @return
 *  - ESP_OK              : Successfully retrieved session list
 *  - ESP_ERR_INVALID_ARG : Wrong arguments or list is longer than provided array
 */
//go:linkname HttpdGetClientList C.httpd_get_client_list
func HttpdGetClientList(handle HttpdHandleT, fds *c.SizeT, client_fds *c.Int) EspErrT

// llgo:type C
type HttpdWorkFnT func(c.Pointer)

/**
 * @brief   Queue execution of a function in HTTPD's context
 *
 * This API queues a work function for asynchronous execution
 *
 * @note    Some protocols require that the web server generate some asynchronous data
 *          and send it to the persistently opened connection. This facility is for use
 *          by such protocols.
 *
 * @param[in] handle    Handle to server returned by httpd_start
 * @param[in] work      Pointer to the function to be executed in the HTTPD's context
 * @param[in] arg       Pointer to the arguments that should be passed to this function
 *
 * @return
 *  - ESP_OK   : On successfully queueing the work
 *  - ESP_FAIL : Failure in ctrl socket
 *  - ESP_ERR_INVALID_ARG : Null arguments
 */
//go:linkname HttpdQueueWork C.httpd_queue_work
func HttpdQueueWork(handle HttpdHandleT, work HttpdWorkFnT, arg c.Pointer) EspErrT
