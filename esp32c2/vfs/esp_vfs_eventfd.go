package vfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Eventfd vfs initialization settings
 */

type EspVfsEventfdConfigT struct {
	MaxFds c.SizeT
}

/**
 * @brief  Registers the event vfs.
 *
 * @return  ESP_OK if successful, ESP_ERR_NO_MEM if too many VFSes are
 *          registered.
 */
// llgo:link (*EspVfsEventfdConfigT).EspVfsEventfdRegister C.esp_vfs_eventfd_register
func (recv_ *EspVfsEventfdConfigT) EspVfsEventfdRegister() EspErrT {
	return 0
}

/**
 * @brief  Unregisters the event vfs.
 *
 * @return ESP_OK if successful, ESP_ERR_INVALID_STATE if VFS for given prefix
 *         hasn't been registered
 */
//go:linkname EspVfsEventfdUnregister C.esp_vfs_eventfd_unregister
func EspVfsEventfdUnregister() EspErrT

/*
 * @brief Creates an event file descriptor.
 *
 * The behavior of read, write and select is the same as man(2) eventfd with
 * EFD_SEMAPHORE **NOT** specified. A new flag EFD_SUPPORT_ISR has been added.
 * This flag is required to write to event fds in interrupt handlers. Accessing
 * the control blocks of event fds with EFD_SUPPORT_ISR will cause interrupts to
 * be temporarily blocked (e.g. during read, write and beginning and ending of
 * the * select).
 *
 * @return The file descriptor if successful, -1 if error happens.
 */
//go:linkname Eventfd C.eventfd
func Eventfd(initval c.Uint, flags c.Int) c.Int
