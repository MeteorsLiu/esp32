package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Type by which software timers are referenced.  For example, a call to
 * xTimerCreate() returns an TimerHandle_t variable that can then be used to
 * reference the subject timer in calls to other software timer API functions
 * (for example, xTimerStart(), xTimerReset(), etc.).
 */

type TmrTimerControl struct {
	Unused [8]uint8
}
type TimerHandleT *TmrTimerControl

// llgo:type C
type TimerCallbackFunctionT func(TimerHandleT)

// llgo:type C
type PendedFunctionT func(c.Pointer, c.Uint32T)

/**
 *
 * Returns the ID assigned to the timer.
 *
 * IDs are assigned to timers using the pvTimerID parameter of the call to
 * xTimerCreated() that was used to create the timer, and by calling the
 * vTimerSetTimerID() API function.
 *
 * If the same callback function is assigned to multiple timers then the timer
 * ID can be used as time specific (timer local) storage.
 *
 * @param xTimer The timer being queried.
 *
 * @return The ID assigned to the timer being queried.
 *
 * Example usage:
 *
 * See the xTimerCreate() API function example usage scenario.
 */
//go:linkname PvTimerGetTimerID C.pvTimerGetTimerID
func PvTimerGetTimerID(xTimer TimerHandleT) c.Pointer

/**
 *
 * Sets the ID assigned to the timer.
 *
 * IDs are assigned to timers using the pvTimerID parameter of the call to
 * xTimerCreated() that was used to create the timer.
 *
 * If the same callback function is assigned to multiple timers then the timer
 * ID can be used as time specific (timer local) storage.
 *
 * @param xTimer The timer being updated.
 *
 * @param pvNewID The ID to assign to the timer.
 *
 * Example usage:
 *
 * See the xTimerCreate() API function example usage scenario.
 */
//go:linkname VTimerSetTimerID C.vTimerSetTimerID
func VTimerSetTimerID(xTimer TimerHandleT, pvNewID c.Pointer)

/**
 *
 * Queries a timer to see if it is active or dormant.
 *
 * A timer will be dormant if:
 *     1) It has been created but not started, or
 *     2) It is an expired one-shot timer that has not been restarted.
 *
 * Timers are created in the dormant state.  The xTimerStart(), xTimerReset(),
 * xTimerStartFromISR(), xTimerResetFromISR(), xTimerChangePeriod() and
 * xTimerChangePeriodFromISR() API functions can all be used to transition a timer into the
 * active state.
 *
 * @param xTimer The timer being queried.
 *
 * @return pdFALSE will be returned if the timer is dormant.  A value other than
 * pdFALSE will be returned if the timer is active.
 *
 * Example usage:
 * @verbatim
 * // This function assumes xTimer has already been created.
 * void vAFunction( TimerHandle_t xTimer )
 * {
 *     if( xTimerIsTimerActive( xTimer ) != pdFALSE ) // or more simply and equivalently "if( xTimerIsTimerActive( xTimer ) )"
 *     {
 *         // xTimer is active, do something.
 *     }
 *     else
 *     {
 *         // xTimer is not active, do something else.
 *     }
 * }
 * @endverbatim
 */
//go:linkname XTimerIsTimerActive C.xTimerIsTimerActive
func XTimerIsTimerActive(xTimer TimerHandleT) BaseTypeT

/**
 *
 * Simply returns the handle of the timer service/daemon task.  It it not valid
 * to call xTimerGetTimerDaemonTaskHandle() before the scheduler has been started.
 */
//go:linkname XTimerGetTimerDaemonTaskHandle C.xTimerGetTimerDaemonTaskHandle
func XTimerGetTimerDaemonTaskHandle() TaskHandleT

/**
 *
 *
 * Used from application interrupt service routines to defer the execution of a
 * function to the RTOS daemon task (the timer service task, hence this function
 * is implemented in timers.c and is prefixed with 'Timer').
 *
 * Ideally an interrupt service routine (ISR) is kept as short as possible, but
 * sometimes an ISR either has a lot of processing to do, or needs to perform
 * processing that is not deterministic.  In these cases
 * xTimerPendFunctionCallFromISR() can be used to defer processing of a function
 * to the RTOS daemon task.
 *
 * A mechanism is provided that allows the interrupt to return directly to the
 * task that will subsequently execute the pended callback function.  This
 * allows the callback function to execute contiguously in time with the
 * interrupt - just as if the callback had executed in the interrupt itself.
 *
 * @param xFunctionToPend The function to execute from the timer service/
 * daemon task.  The function must conform to the PendedFunction_t
 * prototype.
 *
 * @param pvParameter1 The value of the callback function's first parameter.
 * The parameter has a void * type to allow it to be used to pass any type.
 * For example, unsigned longs can be cast to a void *, or the void * can be
 * used to point to a structure.
 *
 * @param ulParameter2 The value of the callback function's second parameter.
 *
 * @param pxHigherPriorityTaskWoken As mentioned above, calling this function
 * will result in a message being sent to the timer daemon task.  If the
 * priority of the timer daemon task (which is set using
 * configTIMER_TASK_PRIORITY in FreeRTOSConfig.h) is higher than the priority of
 * the currently running task (the task the interrupt interrupted) then
 * *pxHigherPriorityTaskWoken will be set to pdTRUE within
 * xTimerPendFunctionCallFromISR(), indicating that a context switch should be
 * requested before the interrupt exits.  For that reason
 * *pxHigherPriorityTaskWoken must be initialised to pdFALSE.  See the
 * example code below.
 *
 * @return pdPASS is returned if the message was successfully sent to the
 * timer daemon task, otherwise pdFALSE is returned.
 *
 * Example usage:
 * @verbatim
 *
 *  // The callback function that will execute in the context of the daemon task.
 *  // Note callback functions must all use this same prototype.
 *  void vProcessInterface( void *pvParameter1, uint32_t ulParameter2 )
 *  {
 *      BaseType_t xInterfaceToService;
 *
 *      // The interface that requires servicing is passed in the second
 *      // parameter.  The first parameter is not used in this case.
 *      xInterfaceToService = ( BaseType_t ) ulParameter2;
 *
 *      // ...Perform the processing here...
 *  }
 *
 *  // An ISR that receives data packets from multiple interfaces
 *  void vAnISR( void )
 *  {
 *      BaseType_t xInterfaceToService, xHigherPriorityTaskWoken;
 *
 *      // Query the hardware to determine which interface needs processing.
 *      xInterfaceToService = prvCheckInterfaces();
 *
 *      // The actual processing is to be deferred to a task.  Request the
 *      // vProcessInterface() callback function is executed, passing in the
 *      // number of the interface that needs processing.  The interface to
 *      // service is passed in the second parameter.  The first parameter is
 *      // not used in this case.
 *      xHigherPriorityTaskWoken = pdFALSE;
 *      xTimerPendFunctionCallFromISR( vProcessInterface, NULL, ( uint32_t ) xInterfaceToService, &xHigherPriorityTaskWoken );
 *
 *      // If xHigherPriorityTaskWoken is now set to pdTRUE then a context
 *      // switch should be requested.  The macro used is port specific and will
 *      // be either portYIELD_FROM_ISR() or portEND_SWITCHING_ISR() - refer to
 *      // the documentation page for the port being used.
 *      portYIELD_FROM_ISR( xHigherPriorityTaskWoken );
 *
 *  }
 * @endverbatim
 */
//go:linkname XTimerPendFunctionCallFromISR C.xTimerPendFunctionCallFromISR
func XTimerPendFunctionCallFromISR(xFunctionToPend PendedFunctionT, pvParameter1 c.Pointer, ulParameter2 c.Uint32T, pxHigherPriorityTaskWoken *BaseTypeT) BaseTypeT

/**
 *
 *
 * Used to defer the execution of a function to the RTOS daemon task (the timer
 * service task, hence this function is implemented in timers.c and is prefixed
 * with 'Timer').
 *
 * @param xFunctionToPend The function to execute from the timer service/
 * daemon task.  The function must conform to the PendedFunction_t
 * prototype.
 *
 * @param pvParameter1 The value of the callback function's first parameter.
 * The parameter has a void * type to allow it to be used to pass any type.
 * For example, unsigned longs can be cast to a void *, or the void * can be
 * used to point to a structure.
 *
 * @param ulParameter2 The value of the callback function's second parameter.
 *
 * @param xTicksToWait Calling this function will result in a message being
 * sent to the timer daemon task on a queue.  xTicksToWait is the amount of
 * time the calling task should remain in the Blocked state (so not using any
 * processing time) for space to become available on the timer queue if the
 * queue is found to be full.
 *
 * @return pdPASS is returned if the message was successfully sent to the
 * timer daemon task, otherwise pdFALSE is returned.
 *
 */
//go:linkname XTimerPendFunctionCall C.xTimerPendFunctionCall
func XTimerPendFunctionCall(xFunctionToPend PendedFunctionT, pvParameter1 c.Pointer, ulParameter2 c.Uint32T, xTicksToWait TickTypeT) BaseTypeT

/**
 *
 * Returns the name that was assigned to a timer when the timer was created.
 *
 * @param xTimer The handle of the timer being queried.
 *
 * @return The name assigned to the timer specified by the xTimer parameter.
 */
//go:linkname PcTimerGetName C.pcTimerGetName
func PcTimerGetName(xTimer TimerHandleT) *c.Char

/**
 *
 * Updates a timer to be either an auto-reload timer, in which case the timer
 * automatically resets itself each time it expires, or a one-shot timer, in
 * which case the timer will only expire once unless it is manually restarted.
 *
 * @param xTimer The handle of the timer being updated.
 *
 * @param xAutoReload If xAutoReload is set to pdTRUE then the timer will
 * expire repeatedly with a frequency set by the timer's period (see the
 * xTimerPeriodInTicks parameter of the xTimerCreate() API function).  If
 * xAutoReload is set to pdFALSE then the timer will be a one-shot timer and
 * enter the dormant state after it expires.
 */
//go:linkname VTimerSetReloadMode C.vTimerSetReloadMode
func VTimerSetReloadMode(xTimer TimerHandleT, xAutoReload BaseTypeT)

/**
 *
 * Queries a timer to determine if it is an auto-reload timer, in which case the timer
 * automatically resets itself each time it expires, or a one-shot timer, in
 * which case the timer will only expire once unless it is manually restarted.
 *
 * @param xTimer The handle of the timer being queried.
 *
 * @return If the timer is an auto-reload timer then pdTRUE is returned, otherwise
 * pdFALSE is returned.
 */
//go:linkname XTimerGetReloadMode C.xTimerGetReloadMode
func XTimerGetReloadMode(xTimer TimerHandleT) BaseTypeT

/**
 *
 * Queries a timer to determine if it is an auto-reload timer, in which case the timer
 * automatically resets itself each time it expires, or a one-shot timer, in
 * which case the timer will only expire once unless it is manually restarted.
 *
 * @param xTimer The handle of the timer being queried.
 *
 * @return If the timer is an auto-reload timer then pdTRUE is returned, otherwise
 * pdFALSE is returned.
 */
//go:linkname UxTimerGetReloadMode C.uxTimerGetReloadMode
func UxTimerGetReloadMode(xTimer TimerHandleT) UBaseTypeT

/**
 *
 * Returns the period of a timer.
 *
 * @param xTimer The handle of the timer being queried.
 *
 * @return The period of the timer in ticks.
 */
//go:linkname XTimerGetPeriod C.xTimerGetPeriod
func XTimerGetPeriod(xTimer TimerHandleT) TickTypeT

/**
 *
 * Returns the time in ticks at which the timer will expire.  If this is less
 * than the current tick count then the expiry time has overflowed from the
 * current time.
 *
 * @param xTimer The handle of the timer being queried.
 *
 * @return If the timer is running then the time in ticks at which the timer
 * will next expire is returned.  If the timer is not running then the return
 * value is undefined.
 */
//go:linkname XTimerGetExpiryTime C.xTimerGetExpiryTime
func XTimerGetExpiryTime(xTimer TimerHandleT) TickTypeT

/*
 * Functions beyond this part are not part of the public API and are intended
 * for use by the kernel only.
 */
//go:linkname XTimerCreateTimerTask C.xTimerCreateTimerTask
func XTimerCreateTimerTask() BaseTypeT

//go:linkname XTimerGenericCommand C.xTimerGenericCommand
func XTimerGenericCommand(xTimer TimerHandleT, xCommandID BaseTypeT, xOptionalValue TickTypeT, pxHigherPriorityTaskWoken *BaseTypeT, xTicksToWait TickTypeT) BaseTypeT
