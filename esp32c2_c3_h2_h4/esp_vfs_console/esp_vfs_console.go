package esp_vfs_console

import _ "unsafe"

const ESP_VFS_DEV_CONSOLE = "/dev/console"

/**
 * @brief add uart/usb_serial_jtag/usb_otg_acmcdc virtual filesystem driver
 *
 * This function is called from startup code to enable serial output
 */
//go:linkname EspVfsConsoleRegister C.esp_vfs_console_register
func EspVfsConsoleRegister() EspErrT
