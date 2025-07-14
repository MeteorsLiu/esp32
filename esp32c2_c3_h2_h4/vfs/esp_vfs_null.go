package vfs

import _ "unsafe"

/**
 * @brief Register filesystem for /dev/null
 *
 * @return ESP_OK on success; any other value indicates an error
 */
//go:linkname EspVfsNullRegister C.esp_vfs_null_register
func EspVfsNullRegister() EspErrT
