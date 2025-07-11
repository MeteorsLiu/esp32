package esp_driver_usb_serial_jtag

import _ "unsafe"

/**
 * @brief get pointer of usb_serial_jtag vfs.
 *
 * This function is called in vfs_console in order to get the vfs implementation
 * of usb_serial_jtag.
 *
 * @return pointer to structure esp_vfs_t
 */
//go:linkname EspVfsUsbSerialJtagGetVfs C.esp_vfs_usb_serial_jtag_get_vfs
func EspVfsUsbSerialJtagGetVfs() *EspVfsFsOpsT

/**
 * @brief add /dev/usbserjtag virtual filesystem driver
 *
 * This function is called from startup code to enable console output
 */
//go:linkname UsbSerialJtagVfsRegister C.usb_serial_jtag_vfs_register
func UsbSerialJtagVfsRegister() EspErrT

/**
 * @brief Set the line endings expected to be received
 *
 * This specifies the conversion between line endings received and
 * newlines ('\n', LF) passed into stdin:
 *
 * - ESP_LINE_ENDINGS_CRLF: convert CRLF to LF
 * - ESP_LINE_ENDINGS_CR: convert CR to LF
 * - ESP_LINE_ENDINGS_LF: no modification
 *
 * @note this function is not thread safe w.r.t. reading
 *
 * @param mode line endings expected
 */
//go:linkname UsbSerialJtagVfsSetRxLineEndings C.usb_serial_jtag_vfs_set_rx_line_endings
func UsbSerialJtagVfsSetRxLineEndings(mode EspLineEndingsT)

/**
 * @brief Set the line endings to sent
 *
 * This specifies the conversion between newlines ('\n', LF) on stdout and line
 * endings sent:
 *
 * - ESP_LINE_ENDINGS_CRLF: convert LF to CRLF
 * - ESP_LINE_ENDINGS_CR: convert LF to CR
 * - ESP_LINE_ENDINGS_LF: no modification
 *
 * @note this function is not thread safe w.r.t. writing
 *
 * @param mode line endings to send
 */
//go:linkname UsbSerialJtagVfsSetTxLineEndings C.usb_serial_jtag_vfs_set_tx_line_endings
func UsbSerialJtagVfsSetTxLineEndings(mode EspLineEndingsT)

/**
 * @brief set VFS to use USB-SERIAL-JTAG driver for reading and writing
 * @note application must configure USB-SERIAL-JTAG driver before calling these functions
 * With these functions, read and write are blocking and interrupt-driven.
 */
//go:linkname UsbSerialJtagVfsUseDriver C.usb_serial_jtag_vfs_use_driver
func UsbSerialJtagVfsUseDriver()

/**
 * @brief set VFS to use simple functions for reading and writing UART
 * Read is non-blocking, write is busy waiting until TX FIFO has enough space.
 * These functions are used by default.
 */
//go:linkname UsbSerialJtagVfsUseNonblocking C.usb_serial_jtag_vfs_use_nonblocking
func UsbSerialJtagVfsUseNonblocking()
