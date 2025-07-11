package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief   Config parameters for protocomm console
 */

type ProtocommConsoleConfigT struct {
	StackSize    c.SizeT
	TaskPriority c.Uint
}

/**
 * @brief   Start console based protocomm transport
 *
 * @note    This is a singleton. ie. Protocomm can have multiple instances, but only
 *          one instance can be bound to a console based transport layer.
 *
 * @param[in] pc     Protocomm instance pointer obtained from protocomm_new()
 * @param[in] config Config param structure for protocomm console
 *
 * @return
 *  - ESP_OK : Success
 *  - ESP_ERR_INVALID_ARG : Null arguments
 *  - ESP_ERR_NOT_SUPPORTED : Transport layer bound to another protocomm instance
 *  - ESP_ERR_INVALID_STATE : Transport layer already bound to this protocomm instance
 *  - ESP_FAIL : Failed to start console thread
 */
// llgo:link (*ProtocommT).ProtocommConsoleStart C.protocomm_console_start
func (recv_ *ProtocommT) ProtocommConsoleStart(config *ProtocommConsoleConfigT) EspErrT {
	return 0
}

/**
 * @brief   Stop console protocomm transport
 *
 * @param[in] pc    Same protocomm instance that was passed to protocomm_console_start()
 *
 * @return
 *  - ESP_OK : Success
 *  - ESP_ERR_INVALID_ARG : Null / incorrect protocomm instance pointer
 */
// llgo:link (*ProtocommT).ProtocommConsoleStop C.protocomm_console_stop
func (recv_ *ProtocommT) ProtocommConsoleStop() EspErrT {
	return 0
}
