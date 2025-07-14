package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief   Config parameters for protocomm HTTP server
 */

type ProtocommHttpServerConfigT struct {
	Port         c.Uint16T
	StackSize    c.SizeT
	TaskPriority c.Uint
}

/** Protocomm HTTPD Configuration Data
 */

type ProtocommHttpdConfigDataT struct {
	Config ProtocommHttpServerConfigT
}

/**
 * @brief   Config parameters for protocomm HTTP server
 */

type ProtocommHttpdConfigT struct {
	ExtHandleProvided bool
	Data              ProtocommHttpdConfigDataT
}

/**
 * @brief   Start HTTPD protocomm transport
 *
 * This API internally creates a framework to allow endpoint registration and security
 * configuration for the protocomm.
 *
 * @note    This is a singleton. ie. Protocomm can have multiple instances, but only
 *          one instance can be bound to an HTTP transport layer.
 *
 * @param[in] pc        Protocomm instance pointer obtained from protocomm_new()
 * @param[in] config    Pointer to config structure for initializing HTTP server
 *
 * @return
 *  - ESP_OK : Success
 *  - ESP_ERR_INVALID_ARG : Null arguments
 *  - ESP_ERR_NOT_SUPPORTED : Transport layer bound to another protocomm instance
 *  - ESP_ERR_INVALID_STATE : Transport layer already bound to this protocomm instance
 *  - ESP_ERR_NO_MEM : Memory allocation for server resource failed
 *  - ESP_ERR_HTTPD_* : HTTP server error on start
 */
// llgo:link (*ProtocommT).ProtocommHttpdStart C.protocomm_httpd_start
func (recv_ *ProtocommT) ProtocommHttpdStart(config *ProtocommHttpdConfigT) EspErrT {
	return 0
}

/**
 * @brief   Stop HTTPD protocomm transport
 *
 * This API cleans up the HTTPD transport protocomm and frees all the handlers registered
 * with the protocomm.
 *
 * @param[in] pc    Same protocomm instance that was passed to protocomm_httpd_start()
 *
 * @return
 *  - ESP_OK : Success
 *  - ESP_ERR_INVALID_ARG : Null / incorrect protocomm instance pointer
 */
// llgo:link (*ProtocommT).ProtocommHttpdStop C.protocomm_httpd_stop
func (recv_ *ProtocommT) ProtocommHttpdStop() EspErrT {
	return 0
}
