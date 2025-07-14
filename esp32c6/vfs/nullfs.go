package vfs

import _ "unsafe"

/**
 * @brief Get VFS structure for /dev/null
 *
 * @return VFS structure for /dev/null
 */
//go:linkname EspVfsNullGetVfs C.esp_vfs_null_get_vfs
func EspVfsNullGetVfs() *EspVfsFsOpsT
