package vfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief add virtual filesystem semihosting driver
 *
 * @param base_path VFS path to mount host directory
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if esp_vfs_semihost_register was already called for specified VFS path
 *      - ESP_ERR_NO_MEM if there are no slots to register new mount point
 */
//go:linkname EspVfsSemihostRegister C.esp_vfs_semihost_register
func EspVfsSemihostRegister(base_path *c.Char) EspErrT

/**
 * @brief Un-register semihosting driver from VFS
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if semihosting driver is not registered in VFS at that path
 */
//go:linkname EspVfsSemihostUnregister C.esp_vfs_semihost_unregister
func EspVfsSemihostUnregister(base_path *c.Char) EspErrT
