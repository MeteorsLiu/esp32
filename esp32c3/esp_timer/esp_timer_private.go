package esp_timer

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Set esp_timer time to a certain value
 *
 * Called from light sleep code to synchronize esp_timer time with RTC time.
 *
 * @param new_us  the value to be set to esp_timer time, in microseconds
 */
//go:linkname EspTimerPrivateSet C.esp_timer_private_set
func EspTimerPrivateSet(new_us c.Uint64T)

/**
 * @brief Adjust current esp_timer time by a certain value
 *
 * @param time_diff_us  adjustment to apply to esp_timer time, in microseconds
 */
//go:linkname EspTimerPrivateAdvance C.esp_timer_private_advance
func EspTimerPrivateAdvance(time_diff_us c.Int64T)

/**
 * @brief obtain internal critical section used in the esp_timer implementation
 * This can be used when a sequence of calls to esp_timer has to be made,
 * and it is necessary that the state of the timer is consistent between
 * the calls. Should be treated in the same way as a spinlock.
 * Call esp_timer_private_unlock to release the lock
 */
//go:linkname EspTimerPrivateLock C.esp_timer_private_lock
func EspTimerPrivateLock()

/**
 * @brief counterpart of esp_timer_lock
 */
//go:linkname EspTimerPrivateUnlock C.esp_timer_private_unlock
func EspTimerPrivateUnlock()
