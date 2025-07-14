package esp_timer

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspTimer struct {
	Unused [8]uint8
}
type EspTimerHandleT *EspTimer

// llgo:type C
type EspTimerCbT func(c.Pointer)
type EspTimerDispatchT c.Int

const (
	ESP_TIMER_TASK EspTimerDispatchT = 0
	ESP_TIMER_MAX  EspTimerDispatchT = 1
)

/**
 * @brief Timer configuration passed to esp_timer_create()
 */

type EspTimerCreateArgsT struct {
	Callback            EspTimerCbT
	Arg                 c.Pointer
	DispatchMethod      EspTimerDispatchT
	Name                *c.Char
	SkipUnhandledEvents bool
}

/**
 * @brief Minimal initialization of esp_timer
 *
 * @note This function is called from startup code. Applications do not need
 * to call this function before using other esp_timer APIs.
 *
 * This function can be called very early in startup process, after this call
 * only esp_timer_get_time() function can be used.
 *
 * @return
 *      - ESP_OK on success
 */
//go:linkname EspTimerEarlyInit C.esp_timer_early_init
func EspTimerEarlyInit() EspErrT

/**
 * @brief Initialize esp_timer library
 *
 * @note This function is called from startup code. Applications do not need
 * to call this function before using other esp_timer APIs.
 * Before calling this function, esp_timer_early_init() must be called by the
 * startup code.
 *
 * This function will be called from startup code on every core.
 * If Kconfig option `CONFIG_ESP_TIMER_ISR_AFFINITY` is set to `NO_AFFINITY`,
 * it allocates the timer ISR on MULTIPLE cores and
 * creates the timer task which can be run on any core.
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NO_MEM if allocation has failed
 *      - ESP_ERR_INVALID_STATE if already initialized
 *      - other errors from interrupt allocator
 */
//go:linkname EspTimerInit C.esp_timer_init
func EspTimerInit() EspErrT

/**
 * @brief De-initialize esp_timer library
 *
 * @note Normally this function should not be called from applications
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_STATE if not yet initialized
 */
//go:linkname EspTimerDeinit C.esp_timer_deinit
func EspTimerDeinit() EspErrT

/**
 * @brief Create an esp_timer instance
 *
 * @note When timer no longer needed, delete it using esp_timer_delete().
 *
 * @param create_args   Pointer to a structure with timer creation arguments.
 *                      Not saved by the library, can be allocated on the stack.
 * @param[out] out_handle  Output, pointer to esp_timer_handle_t variable that
 *                         holds the created timer handle.
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if some of the create_args are not valid
 *      - ESP_ERR_INVALID_STATE if esp_timer library is not initialized yet
 *      - ESP_ERR_NO_MEM if memory allocation fails
 */
// llgo:link (*EspTimerCreateArgsT).EspTimerCreate C.esp_timer_create
func (recv_ *EspTimerCreateArgsT) EspTimerCreate(out_handle *EspTimerHandleT) EspErrT {
	return 0
}

/**
 * @brief Start a one-shot timer
 *
 * Timer represented by `timer` should not be running when this function is called.
 *
 * @param timer timer handle created using esp_timer_create()
 * @param timeout_us timer timeout, in microseconds relative to the current moment
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the handle is invalid
 *      - ESP_ERR_INVALID_STATE if the timer is already running
 */
//go:linkname EspTimerStartOnce C.esp_timer_start_once
func EspTimerStartOnce(timer EspTimerHandleT, timeout_us c.Uint64T) EspErrT

/**
 * @brief Start a periodic timer
 *
 * Timer represented by `timer` should not be running when this function is called.
 * This function starts the timer which will trigger every `period` microseconds.
 *
 * @param timer timer handle created using esp_timer_create()
 * @param period timer period, in microseconds
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the handle is invalid
 *      - ESP_ERR_INVALID_STATE if the timer is already running
 */
//go:linkname EspTimerStartPeriodic C.esp_timer_start_periodic
func EspTimerStartPeriodic(timer EspTimerHandleT, period c.Uint64T) EspErrT

/**
 * @brief Restart a currently running timer
 *
 * Type of `timer` | Action
 * --------------- | ------
 * One-shot timer  | Restarted immediately and times out once in `timeout_us` microseconds
 * Periodic timer  | Restarted immediately with a new period of `timeout_us` microseconds
 *
 * @param timer timer handle created using esp_timer_create()
 * @param timeout_us Timeout in microseconds relative to the current time.
 *                   In case of a periodic timer, also represents the new period.
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the handle is invalid
 *      - ESP_ERR_INVALID_STATE if the timer is not running
 */
//go:linkname EspTimerRestart C.esp_timer_restart
func EspTimerRestart(timer EspTimerHandleT, timeout_us c.Uint64T) EspErrT

/**
 * @brief Stop a running timer
 *
 * This function stops the timer previously started using esp_timer_start_once()
 * or esp_timer_start_periodic().
 *
 * @param timer timer handle created using esp_timer_create()
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_STATE if the timer is not running
 */
//go:linkname EspTimerStop C.esp_timer_stop
func EspTimerStop(timer EspTimerHandleT) EspErrT

/**
 * @brief Delete an esp_timer instance
 *
 * The timer must be stopped before deleting. A one-shot timer which has expired
 * does not need to be stopped.
 *
 * @param timer timer handle created using esp_timer_create()
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_STATE if the timer is running
 */
//go:linkname EspTimerDelete C.esp_timer_delete
func EspTimerDelete(timer EspTimerHandleT) EspErrT

/**
 * @brief Get time in microseconds since boot
 * @return Number of microseconds since the initialization of ESP Timer
 */
//go:linkname EspTimerGetTime C.esp_timer_get_time
func EspTimerGetTime() c.Int64T

/**
 * @brief Get the timestamp of the next expected timeout
 * @return Timestamp of the nearest timer event, in microseconds.
 *         The timebase is the same as for the values returned by esp_timer_get_time().
 */
//go:linkname EspTimerGetNextAlarm C.esp_timer_get_next_alarm
func EspTimerGetNextAlarm() c.Int64T

/**
 * @brief Get the timestamp of the next expected timeout excluding those timers
 *        that should not interrupt light sleep (such timers have
 *        ::esp_timer_create_args_t::skip_unhandled_events enabled)
 * @return Timestamp of the nearest timer event, in microseconds.
 *         The timebase is the same as for the values returned by esp_timer_get_time().
 */
//go:linkname EspTimerGetNextAlarmForWakeUp C.esp_timer_get_next_alarm_for_wake_up
func EspTimerGetNextAlarmForWakeUp() c.Int64T

/**
 * @brief Get the period of a timer
 *
 * This function fetches the timeout period of a timer.
 * For a one-shot timer, the timeout period will be 0.
 *
 * @param timer timer handle created using esp_timer_create()
 * @param period memory to store the timer period value in microseconds
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the arguments are invalid
 */
//go:linkname EspTimerGetPeriod C.esp_timer_get_period
func EspTimerGetPeriod(timer EspTimerHandleT, period *c.Uint64T) EspErrT

/**
 * @brief Get the expiry time of a one-shot timer
 *
 * This function fetches the expiry time of a one-shot timer.
 *
 * @note Passing the timer handle of a periodic timer will result in an error.
 *
 * @param timer timer handle created using esp_timer_create()
 * @param expiry memory to store the timeout value in microseconds
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the arguments are invalid
 *      - ESP_ERR_NOT_SUPPORTED if the timer type is periodic
 */
//go:linkname EspTimerGetExpiryTime C.esp_timer_get_expiry_time
func EspTimerGetExpiryTime(timer EspTimerHandleT, expiry *c.Uint64T) EspErrT

/**
 * @brief Dump the list of timers to a stream
 *
 * By default, this function prints the list of active (running) timers. The output format is:
 *
 * | Name | Period | Alarm |
 *
 * - Name — timer pointer
 * - Period — period of timer in microseconds, or 0 for one-shot timer
 * - Alarm - time of the next alarm in microseconds since boot, or 0 if the timer is not started
 *
 * To print the list of all created timers, enable Kconfig option `CONFIG_ESP_TIMER_PROFILING`.
 * In this case, the output format is:
 *
 * | Name | Period | Alarm | Times_armed | Times_trigg | Times_skip | Cb_exec_time |
 *
 * - Name — timer name
 * - Period — same as above
 * - Alarm — same as above
 * - Times_armed — number of times the timer was armed via esp_timer_start_X
 * - Times_triggered - number of times the callback was triggered
 * - Times_skipped - number of times the callback was skipped
 * - Callback_exec_time - total time taken by callback to execute, across all calls
 *
 * @param stream stream (such as stdout) to which to dump the information
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NO_MEM if can not allocate temporary buffer for the output
 */
//go:linkname EspTimerDump C.esp_timer_dump
func EspTimerDump(stream *c.FILE) EspErrT

/**
 * @brief Returns status of a timer, active or not
 *
 * This function is used to identify if the timer is still active (running) or not.
 *
 * @param timer timer handle created using esp_timer_create()
 * @return
 *      - 1 if timer is still active (running)
 *      - 0 if timer is not active
 */
//go:linkname EspTimerIsActive C.esp_timer_is_active
func EspTimerIsActive(timer EspTimerHandleT) bool
