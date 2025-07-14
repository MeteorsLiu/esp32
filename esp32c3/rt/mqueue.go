package rt

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MqdT c.Pointer

/**
 * @brief Message queue attributes.
 */

type MqAttr struct {
	MqFlags   c.Long
	MqMaxmsg  c.Long
	MqMsgsize c.Long
	MqCurmsgs c.Long
}

/**
 * @brief Close a message queue.
 *
 * Please refer to http://pubs.opengroup.org/onlinepubs/9699919799/functions/mq_close.html for more details.
 */
//go:linkname MqClose C.mq_close
func MqClose(mqdes MqdT) c.Int

/**
 * @brief Get message queue attributes.
 *
 * Please refer to http://pubs.opengroup.org/onlinepubs/9699919799/functions/mq_getattr.html for more details.
 */
//go:linkname MqGetattr C.mq_getattr
func MqGetattr(mqdes MqdT, mqstat *MqAttr) c.Int

/**
 * @brief Open a message queue.
 *
 * Please refer to http://pubs.opengroup.org/onlinepubs/9699919799/functions/mq_open.html for more details.
 *
 * @note Supported name pattern: leading &lt;slash&gt; character in name is always required;
 * the maximum length (excluding null-terminator) of the name argument can be NAME_MAX + 2.
 * @note mode argument is not supported.
 * @note Supported oflags: O_RDWR, O_CREAT, O_EXCL, and O_NONBLOCK.
 */
//go:linkname MqOpen C.mq_open
func MqOpen(name *c.Char, oflag c.Int, __llgo_va_list ...interface{}) MqdT

/**
 * @brief Receive a message from a message queue.
 *
 * Please refer to http://pubs.opengroup.org/onlinepubs/9699919799/functions/mq_receive.html for more details.
 *
 * @note msg_prio argument is not supported. Messages are not checked for corruption.
 */
//go:linkname MqReceive C.mq_receive
func MqReceive(mqdes MqdT, msg_ptr *c.Char, msg_len c.SizeT, msg_prio *c.Uint) c.SsizeT

/**
 * @brief Send a message to a message queue.
 *
 * Please refer to http://pubs.opengroup.org/onlinepubs/9699919799/functions/mq_send.html for more details.
 *
 * @note msg_prio argument is not supported.
 */
//go:linkname MqSend C.mq_send
func MqSend(mqdes MqdT, msg_ptr *c.Char, msg_len c.SizeT, msg_prio c.Uint) c.Int

/**
 * @brief Receive a message from a message queue with timeout.
 *
 * Please refer to http://pubs.opengroup.org/onlinepubs/9699919799/functions/mq_timedreceive.html for more details.
 *
 * @note msg_prio argument is not supported. Messages are not checked for corruption.
 */
//go:linkname MqTimedreceive C.mq_timedreceive
func MqTimedreceive(mqdes MqdT, msg_ptr *c.Char, msg_len c.SizeT, msg_prio *c.Uint, abstime *Timespec) c.SsizeT

/**
 * @brief Send a message to a message queue with timeout.
 *
 * Please refer to http://pubs.opengroup.org/onlinepubs/9699919799/functions/mq_timedsend.html for more details.
 *
 * @note msg_prio argument is not supported.
 */
//go:linkname MqTimedsend C.mq_timedsend
func MqTimedsend(mqdes MqdT, msg_ptr *c.Char, msg_len c.SizeT, msg_prio c.Uint, abstime *Timespec) c.Int

/**
 * @brief Remove a message queue.
 *
 * Please refer to http://pubs.opengroup.org/onlinepubs/9699919799/functions/mq_unlink.html for more details.
 */
//go:linkname MqUnlink C.mq_unlink
func MqUnlink(name *c.Char) c.Int

/* Added by Espressif - specified but not implemented functions, return ENOSYS */
//go:linkname MqNotify C.mq_notify
func MqNotify(MqdT, *Sigevent) c.Int

//go:linkname MqSetattr C.mq_setattr
func MqSetattr(MqdT, *MqAttr, *MqAttr) c.Int
