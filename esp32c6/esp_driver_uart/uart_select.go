package esp_driver_uart

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UartSelectNotifT c.Int

const (
	UART_SELECT_READ_NOTIF  UartSelectNotifT = 0
	UART_SELECT_WRITE_NOTIF UartSelectNotifT = 1
	UART_SELECT_ERROR_NOTIF UartSelectNotifT = 2
)

// llgo:type C
type UartSelectNotifCallbackT func(UartPortT, UartSelectNotifT, *BaseTypeT)

/**
 * @brief Set notification callback function for select() events
 * @param uart_num UART port number
 * @param uart_select_notif_callback callback function
 */
//go:linkname UartSetSelectNotifCallback C.uart_set_select_notif_callback
func UartSetSelectNotifCallback(uart_num UartPortT, uart_select_notif_callback UartSelectNotifCallbackT)

/**
 * @brief Get mutex guarding select() notifications
 */
//go:linkname UartGetSelectlock C.uart_get_selectlock
func UartGetSelectlock() *PortMUXTYPE
