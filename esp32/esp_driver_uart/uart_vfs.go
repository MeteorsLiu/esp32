package esp_driver_uart

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Add /dev/uart virtual filesystem driver
 *
 * This function is called from startup code to enable serial output
 */
//go:linkname UartVfsDevRegister C.uart_vfs_dev_register
func UartVfsDevRegister()

/**
 * @brief Set the line endings expected to be received on specified UART
 *
 * This specifies the conversion between line endings received on UART and
 * newlines ('\n', LF) passed into stdin:
 *
 * - ESP_LINE_ENDINGS_CRLF: convert CRLF to LF
 * - ESP_LINE_ENDINGS_CR: convert CR to LF
 * - ESP_LINE_ENDINGS_LF: no modification
 *
 * @note this function is not thread safe w.r.t. reading from UART
 *
 * @param uart_num the UART number
 * @param mode line endings to send to UART
 *
 * @return  0 if succeeded, or -1
 *              when an error (specified by errno) have occurred.
 */
//go:linkname UartVfsDevPortSetRxLineEndings C.uart_vfs_dev_port_set_rx_line_endings
func UartVfsDevPortSetRxLineEndings(uart_num c.Int, mode EspLineEndingsT) c.Int

/**
 * @brief Set the line endings to sent to specified UART
 *
 * This specifies the conversion between newlines ('\n', LF) on stdout and line
 * endings sent over UART:
 *
 * - ESP_LINE_ENDINGS_CRLF: convert LF to CRLF
 * - ESP_LINE_ENDINGS_CR: convert LF to CR
 * - ESP_LINE_ENDINGS_LF: no modification
 *
 * @note this function is not thread safe w.r.t. writing to UART
 *
 * @param uart_num the UART number
 * @param mode line endings to send to UART
 *
 * @return  0 if succeeded, or -1
 *              when an error (specified by errno) have occurred.
 */
//go:linkname UartVfsDevPortSetTxLineEndings C.uart_vfs_dev_port_set_tx_line_endings
func UartVfsDevPortSetTxLineEndings(uart_num c.Int, mode EspLineEndingsT) c.Int

/**
 * @brief set VFS to use simple functions for reading and writing UART
 *
 * Read is non-blocking, write is busy waiting until TX FIFO has enough space.
 * These functions are used by default.
 *
 * @param uart_num UART peripheral number
 */
//go:linkname UartVfsDevUseNonblocking C.uart_vfs_dev_use_nonblocking
func UartVfsDevUseNonblocking(uart_num c.Int)

/**
 * @brief set VFS to use UART driver for reading and writing
 *
 * @note Application must configure UART driver before calling these functions
 * With these functions, read and write are blocking and interrupt-driven.
 *
 * @param uart_num UART peripheral number
 */
//go:linkname UartVfsDevUseDriver C.uart_vfs_dev_use_driver
func UartVfsDevUseDriver(uart_num c.Int)

/**
 * @brief get pointer of uart vfs.
 *
 * This function is called in vfs_console in order to get the vfs implementation
 * of uart.
 *
 * @return pointer to structure esp_vfs_t
 */
//go:linkname EspVfsUartGetVfs C.esp_vfs_uart_get_vfs
func EspVfsUartGetVfs() *EspVfsFsOpsT
