package esp_pm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Power management config
 *
 * Pass a pointer to this structure as an argument to esp_pm_configure function.
 */

type EspPmConfigT struct {
	MaxFreqMhz       c.Int
	MinFreqMhz       c.Int
	LightSleepEnable bool
}
type EspPmConfigEsp32T EspPmConfigT
type EspPmConfigEsp32s2T EspPmConfigT
type EspPmConfigEsp32s3T EspPmConfigT
type EspPmConfigEsp32c3T EspPmConfigT
type EspPmConfigEsp32c2T EspPmConfigT
type EspPmConfigEsp32c6T EspPmConfigT
type EspPmLockTypeT c.Int

const (
	ESP_PM_CPU_FREQ_MAX   EspPmLockTypeT = 0
	ESP_PM_APB_FREQ_MAX   EspPmLockTypeT = 1
	ESP_PM_NO_LIGHT_SLEEP EspPmLockTypeT = 2
	ESP_PM_LOCK_MAX       EspPmLockTypeT = 3
)

/**
 * @brief Set implementation-specific power management configuration
 * @param config pointer to implementation-specific configuration structure (e.g. esp_pm_config_esp32)
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the configuration values are not correct
 *      - ESP_ERR_NOT_SUPPORTED if certain combination of values is not supported,
 *        or if CONFIG_PM_ENABLE is not enabled in sdkconfig
 */
//go:linkname EspPmConfigure C.esp_pm_configure
func EspPmConfigure(config c.Pointer) EspErrT

/**
 * @brief Get implementation-specific power management configuration
 * @param config pointer to implementation-specific configuration structure (e.g. esp_pm_config_esp32)
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the pointer is null
 */
//go:linkname EspPmGetConfiguration C.esp_pm_get_configuration
func EspPmGetConfiguration(config c.Pointer) EspErrT

type EspPmLock struct {
	Unused [8]uint8
}
type EspPmLockHandleT *EspPmLock

/**
 * @brief Initialize a lock handle for certain power management parameter
 *
 * When lock is created, initially it is not taken.
 * Call esp_pm_lock_acquire to take the lock.
 *
 * This function must not be called from an ISR.
 *
 * @param lock_type Power management constraint which the lock should control
 * @param arg argument, value depends on lock_type, see esp_pm_lock_type_t
 * @param name arbitrary string identifying the lock (e.g. "wifi" or "spi").
 *             Used by the esp_pm_dump_locks function to list existing locks.
 *             May be set to NULL. If not set to NULL, must point to a string which is valid
 *             for the lifetime of the lock.
 * @param[out] out_handle  handle returned from this function. Use this handle when calling
 *                         esp_pm_lock_delete, esp_pm_lock_acquire, esp_pm_lock_release.
 *                         Must not be NULL.
 *
 * @note If the lock_type argument is not valid, it will cause an abort.
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NO_MEM if the lock structure can not be allocated
 *      - ESP_ERR_INVALID_ARG if out_handle is NULL
 *      - ESP_ERR_NOT_SUPPORTED if CONFIG_PM_ENABLE is not enabled in sdkconfig
 */
// llgo:link EspPmLockTypeT.EspPmLockCreate C.esp_pm_lock_create
func (recv_ EspPmLockTypeT) EspPmLockCreate(arg c.Int, name *c.Char, out_handle *EspPmLockHandleT) EspErrT {
	return 0
}

/**
 * @brief Take a power management lock
 *
 * Once the lock is taken, power management algorithm will not switch to the
 * mode specified in a call to esp_pm_lock_create, or any of the lower power
 * modes (higher numeric values of 'mode').
 *
 * The lock is recursive, in the sense that if esp_pm_lock_acquire is called
 * a number of times, esp_pm_lock_release has to be called the same number of
 * times in order to release the lock.
 *
 * This function may be called from an ISR.
 *
 * This function is not thread-safe w.r.t. calls to other esp_pm_lock_*
 * functions for the same handle.
 *
 * @param handle handle obtained from esp_pm_lock_create function
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the handle is invalid
 *      - ESP_ERR_NOT_SUPPORTED if CONFIG_PM_ENABLE is not enabled in sdkconfig
 */
//go:linkname EspPmLockAcquire C.esp_pm_lock_acquire
func EspPmLockAcquire(handle EspPmLockHandleT) EspErrT

/**
 * @brief Release the lock taken using esp_pm_lock_acquire.
 *
 * Call to this functions removes power management restrictions placed when
 * taking the lock.
 *
 * Locks are recursive, so if esp_pm_lock_acquire is called a number of times,
 * esp_pm_lock_release has to be called the same number of times in order to
 * actually release the lock.
 *
 * This function may be called from an ISR.
 *
 * This function is not thread-safe w.r.t. calls to other esp_pm_lock_*
 * functions for the same handle.
 *
 * @param handle handle obtained from esp_pm_lock_create function
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the handle is invalid
 *      - ESP_ERR_INVALID_STATE if lock is not acquired
 *      - ESP_ERR_NOT_SUPPORTED if CONFIG_PM_ENABLE is not enabled in sdkconfig
 */
//go:linkname EspPmLockRelease C.esp_pm_lock_release
func EspPmLockRelease(handle EspPmLockHandleT) EspErrT

/**
 * @brief Delete a lock created using esp_pm_lock
 *
 * The lock must be released before calling this function.
 *
 * This function must not be called from an ISR.
 *
 * @param handle handle obtained from esp_pm_lock_create function
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the handle argument is NULL
 *      - ESP_ERR_INVALID_STATE if the lock is still acquired
 *      - ESP_ERR_NOT_SUPPORTED if CONFIG_PM_ENABLE is not enabled in sdkconfig
 */
//go:linkname EspPmLockDelete C.esp_pm_lock_delete
func EspPmLockDelete(handle EspPmLockHandleT) EspErrT

/**
 * Dump the list of all locks to stderr
 *
 * This function dumps debugging information about locks created using
 * esp_pm_lock_create to an output stream.
 *
 * This function must not be called from an ISR. If esp_pm_lock_acquire/release
 * are called while this function is running, inconsistent results may be
 * reported.
 *
 * @param stream stream to print information to; use stdout or stderr to print
 *               to the console; use fmemopen/open_memstream to print to a
 *               string buffer.
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NOT_SUPPORTED if CONFIG_PM_ENABLE is not enabled in sdkconfig
 */
//go:linkname EspPmDumpLocks C.esp_pm_dump_locks
func EspPmDumpLocks(stream *c.FILE) EspErrT
