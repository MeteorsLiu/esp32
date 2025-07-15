package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const TskKERNEL_VERSION_NUMBER = "V10.5.1"
const TskKERNEL_VERSION_MAJOR = 10
const TskKERNEL_VERSION_MINOR = 5
const TskKERNEL_VERSION_BUILD = 1

/**
 *
 * Type by which tasks are referenced.  For example, a call to xTaskCreate
 * returns (via a pointer parameter) an TaskHandle_t variable that can then
 * be used as a parameter to vTaskDelete to delete the task.
 *
 * \ingroup Tasks
 */

type TskTaskControlBlock struct {
	Unused [8]uint8
}
type TaskHandleT *TskTaskControlBlock

// llgo:type C
type TaskHookFunctionT func(c.Pointer) BaseTypeT
type ETaskState c.Int

const (
	ERunning   ETaskState = 0
	EReady     ETaskState = 1
	EBlocked   ETaskState = 2
	ESuspended ETaskState = 3
	EDeleted   ETaskState = 4
	EInvalid   ETaskState = 5
)

type ENotifyAction c.Int

const (
	ENoAction                 ENotifyAction = 0
	ESetBits                  ENotifyAction = 1
	EIncrement                ENotifyAction = 2
	ESetValueWithOverwrite    ENotifyAction = 3
	ESetValueWithoutOverwrite ENotifyAction = 4
)

/*
 * Used internally only.
 */

type XTIMEOUT struct {
	XOverflowCount  BaseTypeT
	XTimeOnEntering TickTypeT
}
type TimeOutT XTIMEOUT

/*
 * Defines the memory ranges allocated to the task when an MPU is used.
 */

type XMEMORYREGION struct {
	PvBaseAddress   c.Pointer
	UlLengthInBytes c.Uint32T
	UlParameters    c.Uint32T
}
type MemoryRegionT XMEMORYREGION

/*
 * Parameters required to create an MPU protected task.
 */

type XTASKPARAMETERS struct {
	PvTaskCode     c.Int
	PcName         *c.Char
	UsStackDepth   c.Int
	PvParameters   c.Pointer
	UxPriority     UBaseTypeT
	PuxStackBuffer *StackTypeT
	XRegions       [1]MemoryRegionT
}
type TaskParametersT XTASKPARAMETERS

/** Used with the uxTaskGetSystemState() function to return the state of each task
 * in the system. */

type XTASKSTATUS struct {
	XHandle              TaskHandleT
	PcTaskName           *c.Char
	XTaskNumber          UBaseTypeT
	ECurrentState        ETaskState
	UxCurrentPriority    UBaseTypeT
	UxBasePriority       UBaseTypeT
	UlRunTimeCounter     c.Int
	PxStackBase          *StackTypeT
	UsStackHighWaterMark c.Int
}
type TaskStatusT XTASKSTATUS
type ESleepModeStatus c.Int

const (
	EAbortSleep    ESleepModeStatus = 0
	EStandardSleep ESleepModeStatus = 1
)

/**
 *
 * INCLUDE_vTaskDelete must be defined as 1 for this function to be available.
 * See the configuration section for more information.
 *
 * Remove a task from the RTOS real time kernel's management.  The task being
 * deleted will be removed from all ready, blocked, suspended and event lists.
 *
 * NOTE:  The idle task is responsible for freeing the kernel allocated
 * memory from tasks that have been deleted.  It is therefore important that
 * the idle task is not starved of microcontroller processing time if your
 * application makes any calls to vTaskDelete ().  Memory allocated by the
 * task code is not automatically freed, and should be freed before the task
 * is deleted.
 *
 * See the demo application file death.c for sample code that utilises
 * vTaskDelete ().
 *
 * @param xTaskToDelete The handle of the task to be deleted.  Passing NULL will
 * cause the calling task to be deleted.
 *
 * Example usage:
 * @code{c}
 * void vOtherFunction( void )
 * {
 * TaskHandle_t xHandle;
 *
 *   // Create the task, storing the handle.
 *   xTaskCreate( vTaskCode, "NAME", STACK_SIZE, NULL, tskIDLE_PRIORITY, &xHandle );
 *
 *   // Use the handle to delete the task.
 *   vTaskDelete( xHandle );
 * }
 * @endcode
 * \ingroup Tasks
 */
//go:linkname VTaskDelete C.vTaskDelete
func VTaskDelete(xTaskToDelete TaskHandleT)

/**
 *
 * Delay a task for a given number of ticks.  The actual time that the
 * task remains blocked depends on the tick rate.  The constant
 * portTICK_PERIOD_MS can be used to calculate real time from the tick
 * rate - with the resolution of one tick period.
 *
 * INCLUDE_vTaskDelay must be defined as 1 for this function to be available.
 * See the configuration section for more information.
 *
 *
 * vTaskDelay() specifies a time at which the task wishes to unblock relative to
 * the time at which vTaskDelay() is called.  For example, specifying a block
 * period of 100 ticks will cause the task to unblock 100 ticks after
 * vTaskDelay() is called.  vTaskDelay() does not therefore provide a good method
 * of controlling the frequency of a periodic task as the path taken through the
 * code, as well as other task and interrupt activity, will affect the frequency
 * at which vTaskDelay() gets called and therefore the time at which the task
 * next executes.  See xTaskDelayUntil() for an alternative API function designed
 * to facilitate fixed frequency execution.  It does this by specifying an
 * absolute time (rather than a relative time) at which the calling task should
 * unblock.
 *
 * @param xTicksToDelay The amount of time, in tick periods, that
 * the calling task should block.
 *
 * Example usage:
 * @code{c}
 * void vTaskFunction( void * pvParameters )
 * {
 * // Block for 500ms.
 * const TickType_t xDelay = 500 / portTICK_PERIOD_MS;
 *
 *   for( ;; )
 *   {
 *       // Simply toggle the LED every 500ms, blocking between each toggle.
 *       vToggleLED();
 *       vTaskDelay( xDelay );
 *   }
 * }
 * @endcode
 * \ingroup TaskCtrl
 */
// llgo:link TickTypeT.VTaskDelay C.vTaskDelay
func (recv_ TickTypeT) VTaskDelay() {
}

/**
 *
 * INCLUDE_xTaskDelayUntil must be defined as 1 for this function to be available.
 * See the configuration section for more information.
 *
 * Delay a task until a specified time.  This function can be used by periodic
 * tasks to ensure a constant execution frequency.
 *
 * This function differs from vTaskDelay () in one important aspect:  vTaskDelay () will
 * cause a task to block for the specified number of ticks from the time vTaskDelay () is
 * called.  It is therefore difficult to use vTaskDelay () by itself to generate a fixed
 * execution frequency as the time between a task starting to execute and that task
 * calling vTaskDelay () may not be fixed [the task may take a different path though the
 * code between calls, or may get interrupted or preempted a different number of times
 * each time it executes].
 *
 * Whereas vTaskDelay () specifies a wake time relative to the time at which the function
 * is called, xTaskDelayUntil () specifies the absolute (exact) time at which it wishes to
 * unblock.
 *
 * The macro pdMS_TO_TICKS() can be used to calculate the number of ticks from a
 * time specified in milliseconds with a resolution of one tick period.
 *
 * @param pxPreviousWakeTime Pointer to a variable that holds the time at which the
 * task was last unblocked.  The variable must be initialised with the current time
 * prior to its first use (see the example below).  Following this the variable is
 * automatically updated within xTaskDelayUntil ().
 *
 * @param xTimeIncrement The cycle time period.  The task will be unblocked at
 * time *pxPreviousWakeTime + xTimeIncrement.  Calling xTaskDelayUntil with the
 * same xTimeIncrement parameter value will cause the task to execute with
 * a fixed interface period.
 *
 * @return Value which can be used to check whether the task was actually delayed.
 * Will be pdTRUE if the task way delayed and pdFALSE otherwise.  A task will not
 * be delayed if the next expected wake time is in the past.
 *
 * Example usage:
 * @code{c}
 * // Perform an action every 10 ticks.
 * void vTaskFunction( void * pvParameters )
 * {
 * TickType_t xLastWakeTime;
 * const TickType_t xFrequency = 10;
 * BaseType_t xWasDelayed;
 *
 *     // Initialise the xLastWakeTime variable with the current time.
 *     xLastWakeTime = xTaskGetTickCount ();
 *     for( ;; )
 *     {
 *         // Wait for the next cycle.
 *         xWasDelayed = xTaskDelayUntil( &xLastWakeTime, xFrequency );
 *
 *         // Perform action here. xWasDelayed value can be used to determine
 *         // whether a deadline was missed if the code here took too long.
 *     }
 * }
 * @endcode
 * \ingroup TaskCtrl
 */
// llgo:link (*TickTypeT).XTaskDelayUntil C.xTaskDelayUntil
func (recv_ *TickTypeT) XTaskDelayUntil(xTimeIncrement TickTypeT) BaseTypeT {
	return 0
}

/**
 *
 * INCLUDE_xTaskAbortDelay must be defined as 1 in FreeRTOSConfig.h for this
 * function to be available.
 *
 * A task will enter the Blocked state when it is waiting for an event.  The
 * event it is waiting for can be a temporal event (waiting for a time), such
 * as when vTaskDelay() is called, or an event on an object, such as when
 * xQueueReceive() or ulTaskNotifyTake() is called.  If the handle of a task
 * that is in the Blocked state is used in a call to xTaskAbortDelay() then the
 * task will leave the Blocked state, and return from whichever function call
 * placed the task into the Blocked state.
 *
 * There is no 'FromISR' version of this function as an interrupt would need to
 * know which object a task was blocked on in order to know which actions to
 * take.  For example, if the task was blocked on a queue the interrupt handler
 * would then need to know if the queue was locked.
 *
 * @param xTask The handle of the task to remove from the Blocked state.
 *
 * @return If the task referenced by xTask was not in the Blocked state then
 * pdFAIL is returned.  Otherwise pdPASS is returned.
 *
 * \ingroup TaskCtrl
 */
//go:linkname XTaskAbortDelay C.xTaskAbortDelay
func XTaskAbortDelay(xTask TaskHandleT) BaseTypeT

/**
 *
 * INCLUDE_uxTaskPriorityGet must be defined as 1 for this function to be available.
 * See the configuration section for more information.
 *
 * Obtain the priority of any task.
 *
 * @param xTask Handle of the task to be queried.  Passing a NULL
 * handle results in the priority of the calling task being returned.
 *
 * @return The priority of xTask.
 *
 * Example usage:
 * @code{c}
 * void vAFunction( void )
 * {
 * TaskHandle_t xHandle;
 *
 *   // Create a task, storing the handle.
 *   xTaskCreate( vTaskCode, "NAME", STACK_SIZE, NULL, tskIDLE_PRIORITY, &xHandle );
 *
 *   // ...
 *
 *   // Use the handle to obtain the priority of the created task.
 *   // It was created with tskIDLE_PRIORITY, but may have changed
 *   // it itself.
 *   if( uxTaskPriorityGet( xHandle ) != tskIDLE_PRIORITY )
 *   {
 *       // The task has changed it's priority.
 *   }
 *
 *   // ...
 *
 *   // Is our priority higher than the created task?
 *   if( uxTaskPriorityGet( xHandle ) < uxTaskPriorityGet( NULL ) )
 *   {
 *       // Our priority (obtained using NULL handle) is higher.
 *   }
 * }
 * @endcode
 * \ingroup TaskCtrl
 */
//go:linkname UxTaskPriorityGet C.uxTaskPriorityGet
func UxTaskPriorityGet(xTask TaskHandleT) UBaseTypeT

/**
 *
 * A version of uxTaskPriorityGet() that can be used from an ISR.
 */
//go:linkname UxTaskPriorityGetFromISR C.uxTaskPriorityGetFromISR
func UxTaskPriorityGetFromISR(xTask TaskHandleT) UBaseTypeT

/**
 *
 * INCLUDE_eTaskGetState must be defined as 1 for this function to be available.
 * See the configuration section for more information.
 *
 * Obtain the state of any task.  States are encoded by the eTaskState
 * enumerated type.
 *
 * @param xTask Handle of the task to be queried.
 *
 * @return The state of xTask at the time the function was called.  Note the
 * state of the task might change between the function being called, and the
 * functions return value being tested by the calling task.
 */
//go:linkname ETaskGetState C.eTaskGetState
func ETaskGetState(xTask TaskHandleT) ETaskState

/**
 *
 * INCLUDE_vTaskPrioritySet must be defined as 1 for this function to be available.
 * See the configuration section for more information.
 *
 * Set the priority of any task.
 *
 * A context switch will occur before the function returns if the priority
 * being set is higher than the currently executing task.
 *
 * @param xTask Handle to the task for which the priority is being set.
 * Passing a NULL handle results in the priority of the calling task being set.
 *
 * @param uxNewPriority The priority to which the task will be set.
 *
 * Example usage:
 * @code{c}
 * void vAFunction( void )
 * {
 * TaskHandle_t xHandle;
 *
 *   // Create a task, storing the handle.
 *   xTaskCreate( vTaskCode, "NAME", STACK_SIZE, NULL, tskIDLE_PRIORITY, &xHandle );
 *
 *   // ...
 *
 *   // Use the handle to raise the priority of the created task.
 *   vTaskPrioritySet( xHandle, tskIDLE_PRIORITY + 1 );
 *
 *   // ...
 *
 *   // Use a NULL handle to raise our priority to the same value.
 *   vTaskPrioritySet( NULL, tskIDLE_PRIORITY + 1 );
 * }
 * @endcode
 * \ingroup TaskCtrl
 */
//go:linkname VTaskPrioritySet C.vTaskPrioritySet
func VTaskPrioritySet(xTask TaskHandleT, uxNewPriority UBaseTypeT)

/**
 *
 * INCLUDE_vTaskSuspend must be defined as 1 for this function to be available.
 * See the configuration section for more information.
 *
 * Suspend any task.  When suspended a task will never get any microcontroller
 * processing time, no matter what its priority.
 *
 * Calls to vTaskSuspend are not accumulative -
 * i.e. calling vTaskSuspend () twice on the same task still only requires one
 * call to vTaskResume () to ready the suspended task.
 *
 * @param xTaskToSuspend Handle to the task being suspended.  Passing a NULL
 * handle will cause the calling task to be suspended.
 *
 * Example usage:
 * @code{c}
 * void vAFunction( void )
 * {
 * TaskHandle_t xHandle;
 *
 *   // Create a task, storing the handle.
 *   xTaskCreate( vTaskCode, "NAME", STACK_SIZE, NULL, tskIDLE_PRIORITY, &xHandle );
 *
 *   // ...
 *
 *   // Use the handle to suspend the created task.
 *   vTaskSuspend( xHandle );
 *
 *   // ...
 *
 *   // The created task will not run during this period, unless
 *   // another task calls vTaskResume( xHandle ).
 *
 *   //...
 *
 *
 *   // Suspend ourselves.
 *   vTaskSuspend( NULL );
 *
 *   // We cannot get here unless another task calls vTaskResume
 *   // with our handle as the parameter.
 * }
 * @endcode
 * \ingroup TaskCtrl
 */
//go:linkname VTaskSuspend C.vTaskSuspend
func VTaskSuspend(xTaskToSuspend TaskHandleT)

/**
 *
 * INCLUDE_vTaskSuspend must be defined as 1 for this function to be available.
 * See the configuration section for more information.
 *
 * Resumes a suspended task.
 *
 * A task that has been suspended by one or more calls to vTaskSuspend ()
 * will be made available for running again by a single call to
 * vTaskResume ().
 *
 * @param xTaskToResume Handle to the task being readied.
 *
 * Example usage:
 * @code{c}
 * void vAFunction( void )
 * {
 * TaskHandle_t xHandle;
 *
 *   // Create a task, storing the handle.
 *   xTaskCreate( vTaskCode, "NAME", STACK_SIZE, NULL, tskIDLE_PRIORITY, &xHandle );
 *
 *   // ...
 *
 *   // Use the handle to suspend the created task.
 *   vTaskSuspend( xHandle );
 *
 *   // ...
 *
 *   // The created task will not run during this period, unless
 *   // another task calls vTaskResume( xHandle ).
 *
 *   //...
 *
 *
 *   // Resume the suspended task ourselves.
 *   vTaskResume( xHandle );
 *
 *   // The created task will once again get microcontroller processing
 *   // time in accordance with its priority within the system.
 * }
 * @endcode
 * \ingroup TaskCtrl
 */
//go:linkname VTaskResume C.vTaskResume
func VTaskResume(xTaskToResume TaskHandleT)

/**
 *
 * INCLUDE_xTaskResumeFromISR must be defined as 1 for this function to be
 * available.  See the configuration section for more information.
 *
 * An implementation of vTaskResume() that can be called from within an ISR.
 *
 * A task that has been suspended by one or more calls to vTaskSuspend ()
 * will be made available for running again by a single call to
 * xTaskResumeFromISR ().
 *
 * xTaskResumeFromISR() should not be used to synchronise a task with an
 * interrupt if there is a chance that the interrupt could arrive prior to the
 * task being suspended - as this can lead to interrupts being missed. Use of a
 * semaphore as a synchronisation mechanism would avoid this eventuality.
 *
 * @param xTaskToResume Handle to the task being readied.
 *
 * @return pdTRUE if resuming the task should result in a context switch,
 * otherwise pdFALSE. This is used by the ISR to determine if a context switch
 * may be required following the ISR.
 *
 * \ingroup TaskCtrl
 */
//go:linkname XTaskResumeFromISR C.xTaskResumeFromISR
func XTaskResumeFromISR(xTaskToResume TaskHandleT) BaseTypeT

/**
 *
 * Starts the real time kernel tick processing.  After calling the kernel
 * has control over which tasks are executed and when.
 *
 * See the demo application file main.c for an example of creating
 * tasks and starting the kernel.
 *
 * Example usage:
 * @code{c}
 * void vAFunction( void )
 * {
 *   // Create at least one task before starting the kernel.
 *   xTaskCreate( vTaskCode, "NAME", STACK_SIZE, NULL, tskIDLE_PRIORITY, NULL );
 *
 *   // Start the real time kernel with preemption.
 *   vTaskStartScheduler ();
 *
 *   // Will not get here unless a task calls vTaskEndScheduler ()
 * }
 * @endcode
 *
 * \ingroup SchedulerControl
 */
//go:linkname VTaskStartScheduler C.vTaskStartScheduler
func VTaskStartScheduler()

/**
 *
 * NOTE:  At the time of writing only the x86 real mode port, which runs on a PC
 * in place of DOS, implements this function.
 *
 * Stops the real time kernel tick.  All created tasks will be automatically
 * deleted and multitasking (either preemptive or cooperative) will
 * stop.  Execution then resumes from the point where vTaskStartScheduler ()
 * was called, as if vTaskStartScheduler () had just returned.
 *
 * See the demo application file main. c in the demo/PC directory for an
 * example that uses vTaskEndScheduler ().
 *
 * vTaskEndScheduler () requires an exit function to be defined within the
 * portable layer (see vPortEndScheduler () in port. c for the PC port).  This
 * performs hardware specific operations such as stopping the kernel tick.
 *
 * vTaskEndScheduler () will cause all of the resources allocated by the
 * kernel to be freed - but will not free resources allocated by application
 * tasks.
 *
 * Example usage:
 * @code{c}
 * void vTaskCode( void * pvParameters )
 * {
 *   for( ;; )
 *   {
 *       // Task code goes here.
 *
 *       // At some point we want to end the real time kernel processing
 *       // so call ...
 *       vTaskEndScheduler ();
 *   }
 * }
 *
 * void vAFunction( void )
 * {
 *   // Create at least one task before starting the kernel.
 *   xTaskCreate( vTaskCode, "NAME", STACK_SIZE, NULL, tskIDLE_PRIORITY, NULL );
 *
 *   // Start the real time kernel with preemption.
 *   vTaskStartScheduler ();
 *
 *   // Will only get here when the vTaskCode () task has called
 *   // vTaskEndScheduler ().  When we get here we are back to single task
 *   // execution.
 * }
 * @endcode
 *
 * \ingroup SchedulerControl
 */
//go:linkname VTaskEndScheduler C.vTaskEndScheduler
func VTaskEndScheduler()

/**
 *
 * Suspends the scheduler without disabling interrupts.  Context switches will
 * not occur while the scheduler is suspended.
 *
 * After calling vTaskSuspendAll () the calling task will continue to execute
 * without risk of being swapped out until a call to xTaskResumeAll () has been
 * made.
 *
 * API functions that have the potential to cause a context switch (for example,
 * xTaskDelayUntil(), xQueueSend(), etc.) must not be called while the scheduler
 * is suspended.
 *
 * Example usage:
 * @code{c}
 * void vTask1( void * pvParameters )
 * {
 *   for( ;; )
 *   {
 *       // Task code goes here.
 *
 *       // ...
 *
 *       // At some point the task wants to perform a long operation during
 *       // which it does not want to get swapped out.  It cannot use
 *       // taskENTER_CRITICAL ()/taskEXIT_CRITICAL () as the length of the
 *       // operation may cause interrupts to be missed - including the
 *       // ticks.
 *
 *       // Prevent the real time kernel swapping out the task.
 *       vTaskSuspendAll ();
 *
 *       // Perform the operation here.  There is no need to use critical
 *       // sections as we have all the microcontroller processing time.
 *       // During this time interrupts will still operate and the kernel
 *       // tick count will be maintained.
 *
 *       // ...
 *
 *       // The operation is complete.  Restart the kernel.
 *       xTaskResumeAll ();
 *   }
 * }
 * @endcode
 * \ingroup SchedulerControl
 */
//go:linkname VTaskSuspendAll C.vTaskSuspendAll
func VTaskSuspendAll()

/**
 *
 * Resumes scheduler activity after it was suspended by a call to
 * vTaskSuspendAll().
 *
 * xTaskResumeAll() only resumes the scheduler.  It does not unsuspend tasks
 * that were previously suspended by a call to vTaskSuspend().
 *
 * @return If resuming the scheduler caused a context switch then pdTRUE is
 *         returned, otherwise pdFALSE is returned.
 *
 * Example usage:
 * @code{c}
 * void vTask1( void * pvParameters )
 * {
 *   for( ;; )
 *   {
 *       // Task code goes here.
 *
 *       // ...
 *
 *       // At some point the task wants to perform a long operation during
 *       // which it does not want to get swapped out.  It cannot use
 *       // taskENTER_CRITICAL ()/taskEXIT_CRITICAL () as the length of the
 *       // operation may cause interrupts to be missed - including the
 *       // ticks.
 *
 *       // Prevent the real time kernel swapping out the task.
 *       vTaskSuspendAll ();
 *
 *       // Perform the operation here.  There is no need to use critical
 *       // sections as we have all the microcontroller processing time.
 *       // During this time interrupts will still operate and the real
 *       // time kernel tick count will be maintained.
 *
 *       // ...
 *
 *       // The operation is complete.  Restart the kernel.  We want to force
 *       // a context switch - but there is no point if resuming the scheduler
 *       // caused a context switch already.
 *       if( !xTaskResumeAll () )
 *       {
 *            taskYIELD ();
 *       }
 *   }
 * }
 * @endcode
 * \ingroup SchedulerControl
 */
//go:linkname XTaskResumeAll C.xTaskResumeAll
func XTaskResumeAll() BaseTypeT

/**
 *
 * @return The count of ticks since vTaskStartScheduler was called.
 *
 * \ingroup TaskUtils
 */
//go:linkname XTaskGetTickCount C.xTaskGetTickCount
func XTaskGetTickCount() TickTypeT

/**
 *
 * @return The count of ticks since vTaskStartScheduler was called.
 *
 * This is a version of xTaskGetTickCount() that is safe to be called from an
 * ISR - provided that TickType_t is the natural word size of the
 * microcontroller being used or interrupt nesting is either not supported or
 * not being used.
 *
 * \ingroup TaskUtils
 */
//go:linkname XTaskGetTickCountFromISR C.xTaskGetTickCountFromISR
func XTaskGetTickCountFromISR() TickTypeT

/**
 *
 * @return The number of tasks that the real time kernel is currently managing.
 * This includes all ready, blocked and suspended tasks.  A task that
 * has been deleted but not yet freed by the idle task will also be
 * included in the count.
 *
 * \ingroup TaskUtils
 */
//go:linkname UxTaskGetNumberOfTasks C.uxTaskGetNumberOfTasks
func UxTaskGetNumberOfTasks() UBaseTypeT

/**
 *
 * @return The text (human readable) name of the task referenced by the handle
 * xTaskToQuery.  A task can query its own name by either passing in its own
 * handle, or by setting xTaskToQuery to NULL.
 *
 * \ingroup TaskUtils
 */
//go:linkname PcTaskGetName C.pcTaskGetName
func PcTaskGetName(xTaskToQuery TaskHandleT) *c.Char

/**
 *
 * NOTE:  This function takes a relatively long time to complete and should be
 * used sparingly.
 *
 * @return The handle of the task that has the human readable name pcNameToQuery.
 * NULL is returned if no matching name is found.  INCLUDE_xTaskGetHandle
 * must be set to 1 in FreeRTOSConfig.h for pcTaskGetHandle() to be available.
 *
 * \ingroup TaskUtils
 */
//go:linkname XTaskGetHandle C.xTaskGetHandle
func XTaskGetHandle(pcNameToQuery *c.Char) TaskHandleT

/**
 *
 * INCLUDE_uxTaskGetStackHighWaterMark must be set to 1 in FreeRTOSConfig.h for
 * this function to be available.
 *
 * Returns the high water mark of the stack associated with xTask.  That is,
 * the minimum free stack space there has been (in words, so on a 32 bit machine
 * a value of 1 means 4 bytes) since the task started.  The smaller the returned
 * number the closer the task has come to overflowing its stack.
 *
 * uxTaskGetStackHighWaterMark() and uxTaskGetStackHighWaterMark2() are the
 * same except for their return type.  Using configSTACK_DEPTH_TYPE allows the
 * user to determine the return type.  It gets around the problem of the value
 * overflowing on 8-bit types without breaking backward compatibility for
 * applications that expect an 8-bit return type.
 *
 * @param xTask Handle of the task associated with the stack to be checked.
 * Set xTask to NULL to check the stack of the calling task.
 *
 * @return The smallest amount of free stack space there has been (in words, so
 * actual spaces on the stack rather than bytes) since the task referenced by
 * xTask was created.
 */
//go:linkname UxTaskGetStackHighWaterMark C.uxTaskGetStackHighWaterMark
func UxTaskGetStackHighWaterMark(xTask TaskHandleT) UBaseTypeT

/**
 *
 * INCLUDE_uxTaskGetStackHighWaterMark2 must be set to 1 in FreeRTOSConfig.h for
 * this function to be available.
 *
 * Returns the high water mark of the stack associated with xTask.  That is,
 * the minimum free stack space there has been (in words, so on a 32 bit machine
 * a value of 1 means 4 bytes) since the task started.  The smaller the returned
 * number the closer the task has come to overflowing its stack.
 *
 * uxTaskGetStackHighWaterMark() and uxTaskGetStackHighWaterMark2() are the
 * same except for their return type.  Using configSTACK_DEPTH_TYPE allows the
 * user to determine the return type.  It gets around the problem of the value
 * overflowing on 8-bit types without breaking backward compatibility for
 * applications that expect an 8-bit return type.
 *
 * @param xTask Handle of the task associated with the stack to be checked.
 * Set xTask to NULL to check the stack of the calling task.
 *
 * @return The smallest amount of free stack space there has been (in words, so
 * actual spaces on the stack rather than bytes) since the task referenced by
 * xTask was created.
 */
//go:linkname UxTaskGetStackHighWaterMark2 C.uxTaskGetStackHighWaterMark2
func UxTaskGetStackHighWaterMark2(xTask TaskHandleT) c.Int

/**
 * xTaskGetIdleTaskHandle() is only available if
 * INCLUDE_xTaskGetIdleTaskHandle is set to 1 in FreeRTOSConfig.h.
 *
 * Simply returns the handle of the idle task of the current core.  It is not
 * valid to call xTaskGetIdleTaskHandle() before the scheduler has been started.
 */
//go:linkname XTaskGetIdleTaskHandle C.xTaskGetIdleTaskHandle
func XTaskGetIdleTaskHandle() TaskHandleT

/**
 *
 * See https://www.FreeRTOS.org/RTOS-task-notifications.html for details.
 *
 * configUSE_TASK_NOTIFICATIONS must be undefined or defined as 1 for these
 * functions to be available.
 *
 * Sends a direct to task notification to a task, with an optional value and
 * action.
 *
 * Each task has a private array of "notification values" (or 'notifications'),
 * each of which is a 32-bit unsigned integer (uint32_t).  The constant
 * configTASK_NOTIFICATION_ARRAY_ENTRIES sets the number of indexes in the
 * array, and (for backward compatibility) defaults to 1 if left undefined.
 * Prior to FreeRTOS V10.4.0 there was only one notification value per task.
 *
 * Events can be sent to a task using an intermediary object.  Examples of such
 * objects are queues, semaphores, mutexes and event groups.  Task notifications
 * are a method of sending an event directly to a task without the need for such
 * an intermediary object.
 *
 * A notification sent to a task can optionally perform an action, such as
 * update, overwrite or increment one of the task's notification values.  In
 * that way task notifications can be used to send data to a task, or be used as
 * light weight and fast binary or counting semaphores.
 *
 * A task can use xTaskNotifyWaitIndexed() or ulTaskNotifyTakeIndexed() to
 * [optionally] block to wait for a notification to be pending.  The task does
 * not consume any CPU time while it is in the Blocked state.
 *
 * A notification sent to a task will remain pending until it is cleared by the
 * task calling xTaskNotifyWaitIndexed() or ulTaskNotifyTakeIndexed() (or their
 * un-indexed equivalents).  If the task was already in the Blocked state to
 * wait for a notification when the notification arrives then the task will
 * automatically be removed from the Blocked state (unblocked) and the
 * notification cleared.
 *
 * **NOTE** Each notification within the array operates independently - a task
 * can only block on one notification within the array at a time and will not be
 * unblocked by a notification sent to any other array index.
 *
 * Backward compatibility information:
 * Prior to FreeRTOS V10.4.0 each task had a single "notification value", and
 * all task notification API functions operated on that value. Replacing the
 * single notification value with an array of notification values necessitated a
 * new set of API functions that could address specific notifications within the
 * array.  xTaskNotify() is the original API function, and remains backward
 * compatible by always operating on the notification value at index 0 in the
 * array. Calling xTaskNotify() is equivalent to calling xTaskNotifyIndexed()
 * with the uxIndexToNotify parameter set to 0.
 *
 * @param xTaskToNotify The handle of the task being notified.  The handle to a
 * task can be returned from the xTaskCreate() API function used to create the
 * task, and the handle of the currently running task can be obtained by calling
 * xTaskGetCurrentTaskHandle().
 *
 * @param uxIndexToNotify The index within the target task's array of
 * notification values to which the notification is to be sent.  uxIndexToNotify
 * must be less than configTASK_NOTIFICATION_ARRAY_ENTRIES.  xTaskNotify() does
 * not have this parameter and always sends notifications to index 0.
 *
 * @param ulValue Data that can be sent with the notification.  How the data is
 * used depends on the value of the eAction parameter.
 *
 * @param eAction Specifies how the notification updates the task's notification
 * value, if at all.  Valid values for eAction are as follows:
 *
 * eSetBits -
 * The target notification value is bitwise ORed with ulValue.
 * xTaskNotifyIndexed() always returns pdPASS in this case.
 *
 * eIncrement -
 * The target notification value is incremented.  ulValue is not used and
 * xTaskNotifyIndexed() always returns pdPASS in this case.
 *
 * eSetValueWithOverwrite -
 * The target notification value is set to the value of ulValue, even if the
 * task being notified had not yet processed the previous notification at the
 * same array index (the task already had a notification pending at that index).
 * xTaskNotifyIndexed() always returns pdPASS in this case.
 *
 * eSetValueWithoutOverwrite -
 * If the task being notified did not already have a notification pending at the
 * same array index then the target notification value is set to ulValue and
 * xTaskNotifyIndexed() will return pdPASS.  If the task being notified already
 * had a notification pending at the same array index then no action is
 * performed and pdFAIL is returned.
 *
 * eNoAction -
 * The task receives a notification at the specified array index without the
 * notification value at that index being updated.  ulValue is not used and
 * xTaskNotifyIndexed() always returns pdPASS in this case.
 *
 * pulPreviousNotificationValue -
 * Can be used to pass out the subject task's notification value before any
 * bits are modified by the notify function.
 *
 * @return Dependent on the value of eAction.  See the description of the
 * eAction parameter.
 *
 * \ingroup TaskNotifications
 */
/** @cond !DOC_EXCLUDE_HEADER_SECTION */
//go:linkname XTaskGenericNotify C.xTaskGenericNotify
func XTaskGenericNotify(xTaskToNotify TaskHandleT, uxIndexToNotify UBaseTypeT, ulValue c.Uint32T, eAction ENotifyAction, pulPreviousNotificationValue *c.Uint32T) BaseTypeT

/**
 *
 * See https://www.FreeRTOS.org/RTOS-task-notifications.html for details.
 *
 * configUSE_TASK_NOTIFICATIONS must be undefined or defined as 1 for these
 * functions to be available.
 *
 * A version of xTaskNotifyIndexed() that can be used from an interrupt service
 * routine (ISR).
 *
 * Each task has a private array of "notification values" (or 'notifications'),
 * each of which is a 32-bit unsigned integer (uint32_t).  The constant
 * configTASK_NOTIFICATION_ARRAY_ENTRIES sets the number of indexes in the
 * array, and (for backward compatibility) defaults to 1 if left undefined.
 * Prior to FreeRTOS V10.4.0 there was only one notification value per task.
 *
 * Events can be sent to a task using an intermediary object.  Examples of such
 * objects are queues, semaphores, mutexes and event groups.  Task notifications
 * are a method of sending an event directly to a task without the need for such
 * an intermediary object.
 *
 * A notification sent to a task can optionally perform an action, such as
 * update, overwrite or increment one of the task's notification values.  In
 * that way task notifications can be used to send data to a task, or be used as
 * light weight and fast binary or counting semaphores.
 *
 * A task can use xTaskNotifyWaitIndexed() to [optionally] block to wait for a
 * notification to be pending, or ulTaskNotifyTakeIndexed() to [optionally] block
 * to wait for a notification value to have a non-zero value.  The task does
 * not consume any CPU time while it is in the Blocked state.
 *
 * A notification sent to a task will remain pending until it is cleared by the
 * task calling xTaskNotifyWaitIndexed() or ulTaskNotifyTakeIndexed() (or their
 * un-indexed equivalents).  If the task was already in the Blocked state to
 * wait for a notification when the notification arrives then the task will
 * automatically be removed from the Blocked state (unblocked) and the
 * notification cleared.
 *
 * **NOTE** Each notification within the array operates independently - a task
 * can only block on one notification within the array at a time and will not be
 * unblocked by a notification sent to any other array index.
 *
 * Backward compatibility information:
 * Prior to FreeRTOS V10.4.0 each task had a single "notification value", and
 * all task notification API functions operated on that value. Replacing the
 * single notification value with an array of notification values necessitated a
 * new set of API functions that could address specific notifications within the
 * array.  xTaskNotifyFromISR() is the original API function, and remains
 * backward compatible by always operating on the notification value at index 0
 * within the array. Calling xTaskNotifyFromISR() is equivalent to calling
 * xTaskNotifyIndexedFromISR() with the uxIndexToNotify parameter set to 0.
 *
 * @param uxIndexToNotify The index within the target task's array of
 * notification values to which the notification is to be sent.  uxIndexToNotify
 * must be less than configTASK_NOTIFICATION_ARRAY_ENTRIES.  xTaskNotifyFromISR()
 * does not have this parameter and always sends notifications to index 0.
 *
 * @param xTaskToNotify The handle of the task being notified.  The handle to a
 * task can be returned from the xTaskCreate() API function used to create the
 * task, and the handle of the currently running task can be obtained by calling
 * xTaskGetCurrentTaskHandle().
 *
 * @param ulValue Data that can be sent with the notification.  How the data is
 * used depends on the value of the eAction parameter.
 *
 * @param eAction Specifies how the notification updates the task's notification
 * value, if at all.  Valid values for eAction are as follows:
 *
 * eSetBits -
 * The task's notification value is bitwise ORed with ulValue.  xTaskNotify()
 * always returns pdPASS in this case.
 *
 * eIncrement -
 * The task's notification value is incremented.  ulValue is not used and
 * xTaskNotify() always returns pdPASS in this case.
 *
 * eSetValueWithOverwrite -
 * The task's notification value is set to the value of ulValue, even if the
 * task being notified had not yet processed the previous notification (the
 * task already had a notification pending).  xTaskNotify() always returns
 * pdPASS in this case.
 *
 * eSetValueWithoutOverwrite -
 * If the task being notified did not already have a notification pending then
 * the task's notification value is set to ulValue and xTaskNotify() will
 * return pdPASS.  If the task being notified already had a notification
 * pending then no action is performed and pdFAIL is returned.
 *
 * eNoAction -
 * The task receives a notification without its notification value being
 * updated.  ulValue is not used and xTaskNotify() always returns pdPASS in
 * this case.
 *
 * @param pxHigherPriorityTaskWoken  xTaskNotifyFromISR() will set
 * *pxHigherPriorityTaskWoken to pdTRUE if sending the notification caused the
 * task to which the notification was sent to leave the Blocked state, and the
 * unblocked task has a priority higher than the currently running task.  If
 * xTaskNotifyFromISR() sets this value to pdTRUE then a context switch should
 * be requested before the interrupt is exited.  How a context switch is
 * requested from an ISR is dependent on the port - see the documentation page
 * for the port in use.
 *
 * @return Dependent on the value of eAction.  See the description of the
 * eAction parameter.
 *
 * \ingroup TaskNotifications
 */
/** @cond !DOC_EXCLUDE_HEADER_SECTION */
//go:linkname XTaskGenericNotifyFromISR C.xTaskGenericNotifyFromISR
func XTaskGenericNotifyFromISR(xTaskToNotify TaskHandleT, uxIndexToNotify UBaseTypeT, ulValue c.Uint32T, eAction ENotifyAction, pulPreviousNotificationValue *c.Uint32T, pxHigherPriorityTaskWoken *BaseTypeT) BaseTypeT

/**
 *
 * Waits for a direct to task notification to be pending at a given index within
 * an array of direct to task notifications.
 *
 * See https://www.FreeRTOS.org/RTOS-task-notifications.html for details.
 *
 * configUSE_TASK_NOTIFICATIONS must be undefined or defined as 1 for this
 * function to be available.
 *
 * Each task has a private array of "notification values" (or 'notifications'),
 * each of which is a 32-bit unsigned integer (uint32_t).  The constant
 * configTASK_NOTIFICATION_ARRAY_ENTRIES sets the number of indexes in the
 * array, and (for backward compatibility) defaults to 1 if left undefined.
 * Prior to FreeRTOS V10.4.0 there was only one notification value per task.
 *
 * Events can be sent to a task using an intermediary object.  Examples of such
 * objects are queues, semaphores, mutexes and event groups.  Task notifications
 * are a method of sending an event directly to a task without the need for such
 * an intermediary object.
 *
 * A notification sent to a task can optionally perform an action, such as
 * update, overwrite or increment one of the task's notification values.  In
 * that way task notifications can be used to send data to a task, or be used as
 * light weight and fast binary or counting semaphores.
 *
 * A notification sent to a task will remain pending until it is cleared by the
 * task calling xTaskNotifyWaitIndexed() or ulTaskNotifyTakeIndexed() (or their
 * un-indexed equivalents).  If the task was already in the Blocked state to
 * wait for a notification when the notification arrives then the task will
 * automatically be removed from the Blocked state (unblocked) and the
 * notification cleared.
 *
 * A task can use xTaskNotifyWaitIndexed() to [optionally] block to wait for a
 * notification to be pending, or ulTaskNotifyTakeIndexed() to [optionally] block
 * to wait for a notification value to have a non-zero value.  The task does
 * not consume any CPU time while it is in the Blocked state.
 *
 * **NOTE** Each notification within the array operates independently - a task
 * can only block on one notification within the array at a time and will not be
 * unblocked by a notification sent to any other array index.
 *
 * Backward compatibility information:
 * Prior to FreeRTOS V10.4.0 each task had a single "notification value", and
 * all task notification API functions operated on that value. Replacing the
 * single notification value with an array of notification values necessitated a
 * new set of API functions that could address specific notifications within the
 * array.  xTaskNotifyWait() is the original API function, and remains backward
 * compatible by always operating on the notification value at index 0 in the
 * array. Calling xTaskNotifyWait() is equivalent to calling
 * xTaskNotifyWaitIndexed() with the uxIndexToWaitOn parameter set to 0.
 *
 * @param uxIndexToWaitOn The index within the calling task's array of
 * notification values on which the calling task will wait for a notification to
 * be received.  uxIndexToWaitOn must be less than
 * configTASK_NOTIFICATION_ARRAY_ENTRIES.  xTaskNotifyWait() does
 * not have this parameter and always waits for notifications on index 0.
 *
 * @param ulBitsToClearOnEntry Bits that are set in ulBitsToClearOnEntry value
 * will be cleared in the calling task's notification value before the task
 * checks to see if any notifications are pending, and optionally blocks if no
 * notifications are pending.  Setting ulBitsToClearOnEntry to ULONG_MAX (if
 * limits.h is included) or 0xffffffffUL (if limits.h is not included) will have
 * the effect of resetting the task's notification value to 0.  Setting
 * ulBitsToClearOnEntry to 0 will leave the task's notification value unchanged.
 *
 * @param ulBitsToClearOnExit If a notification is pending or received before
 * the calling task exits the xTaskNotifyWait() function then the task's
 * notification value (see the xTaskNotify() API function) is passed out using
 * the pulNotificationValue parameter.  Then any bits that are set in
 * ulBitsToClearOnExit will be cleared in the task's notification value (note
 * *pulNotificationValue is set before any bits are cleared).  Setting
 * ulBitsToClearOnExit to ULONG_MAX (if limits.h is included) or 0xffffffffUL
 * (if limits.h is not included) will have the effect of resetting the task's
 * notification value to 0 before the function exits.  Setting
 * ulBitsToClearOnExit to 0 will leave the task's notification value unchanged
 * when the function exits (in which case the value passed out in
 * pulNotificationValue will match the task's notification value).
 *
 * @param pulNotificationValue Used to pass the task's notification value out
 * of the function.  Note the value passed out will not be effected by the
 * clearing of any bits caused by ulBitsToClearOnExit being non-zero.
 *
 * @param xTicksToWait The maximum amount of time that the task should wait in
 * the Blocked state for a notification to be received, should a notification
 * not already be pending when xTaskNotifyWait() was called.  The task
 * will not consume any processing time while it is in the Blocked state.  This
 * is specified in kernel ticks, the macro pdMS_TO_TICKS( value_in_ms ) can be
 * used to convert a time specified in milliseconds to a time specified in
 * ticks.
 *
 * @return If a notification was received (including notifications that were
 * already pending when xTaskNotifyWait was called) then pdPASS is
 * returned.  Otherwise pdFAIL is returned.
 *
 * \ingroup TaskNotifications
 */
// llgo:link UBaseTypeT.XTaskGenericNotifyWait C.xTaskGenericNotifyWait
func (recv_ UBaseTypeT) XTaskGenericNotifyWait(ulBitsToClearOnEntry c.Uint32T, ulBitsToClearOnExit c.Uint32T, pulNotificationValue *c.Uint32T, xTicksToWait TickTypeT) BaseTypeT {
	return 0
}

/**
 *
 * A version of xTaskNotifyGiveIndexed() that can be called from an interrupt
 * service routine (ISR).
 *
 * See https://www.FreeRTOS.org/RTOS-task-notifications.html for more details.
 *
 * configUSE_TASK_NOTIFICATIONS must be undefined or defined as 1 for this macro
 * to be available.
 *
 * Each task has a private array of "notification values" (or 'notifications'),
 * each of which is a 32-bit unsigned integer (uint32_t).  The constant
 * configTASK_NOTIFICATION_ARRAY_ENTRIES sets the number of indexes in the
 * array, and (for backward compatibility) defaults to 1 if left undefined.
 * Prior to FreeRTOS V10.4.0 there was only one notification value per task.
 *
 * Events can be sent to a task using an intermediary object.  Examples of such
 * objects are queues, semaphores, mutexes and event groups.  Task notifications
 * are a method of sending an event directly to a task without the need for such
 * an intermediary object.
 *
 * A notification sent to a task can optionally perform an action, such as
 * update, overwrite or increment one of the task's notification values.  In
 * that way task notifications can be used to send data to a task, or be used as
 * light weight and fast binary or counting semaphores.
 *
 * vTaskNotifyGiveIndexedFromISR() is intended for use when task notifications
 * are used as light weight and faster binary or counting semaphore equivalents.
 * Actual FreeRTOS semaphores are given from an ISR using the
 * xSemaphoreGiveFromISR() API function, the equivalent action that instead uses
 * a task notification is vTaskNotifyGiveIndexedFromISR().
 *
 * When task notifications are being used as a binary or counting semaphore
 * equivalent then the task being notified should wait for the notification
 * using the ulTaskNotifyTakeIndexed() API function rather than the
 * xTaskNotifyWaitIndexed() API function.
 *
 * **NOTE** Each notification within the array operates independently - a task
 * can only block on one notification within the array at a time and will not be
 * unblocked by a notification sent to any other array index.
 *
 * Backward compatibility information:
 * Prior to FreeRTOS V10.4.0 each task had a single "notification value", and
 * all task notification API functions operated on that value. Replacing the
 * single notification value with an array of notification values necessitated a
 * new set of API functions that could address specific notifications within the
 * array.  xTaskNotifyFromISR() is the original API function, and remains
 * backward compatible by always operating on the notification value at index 0
 * within the array. Calling xTaskNotifyGiveFromISR() is equivalent to calling
 * xTaskNotifyGiveIndexedFromISR() with the uxIndexToNotify parameter set to 0.
 *
 * @param xTaskToNotify The handle of the task being notified.  The handle to a
 * task can be returned from the xTaskCreate() API function used to create the
 * task, and the handle of the currently running task can be obtained by calling
 * xTaskGetCurrentTaskHandle().
 *
 * @param uxIndexToNotify The index within the target task's array of
 * notification values to which the notification is to be sent.  uxIndexToNotify
 * must be less than configTASK_NOTIFICATION_ARRAY_ENTRIES.
 * xTaskNotifyGiveFromISR() does not have this parameter and always sends
 * notifications to index 0.
 *
 * @param pxHigherPriorityTaskWoken  vTaskNotifyGiveFromISR() will set
 * *pxHigherPriorityTaskWoken to pdTRUE if sending the notification caused the
 * task to which the notification was sent to leave the Blocked state, and the
 * unblocked task has a priority higher than the currently running task.  If
 * vTaskNotifyGiveFromISR() sets this value to pdTRUE then a context switch
 * should be requested before the interrupt is exited.  How a context switch is
 * requested from an ISR is dependent on the port - see the documentation page
 * for the port in use.
 *
 * \ingroup TaskNotifications
 */
//go:linkname VTaskGenericNotifyGiveFromISR C.vTaskGenericNotifyGiveFromISR
func VTaskGenericNotifyGiveFromISR(xTaskToNotify TaskHandleT, uxIndexToNotify UBaseTypeT, pxHigherPriorityTaskWoken *BaseTypeT)

/**
 *
 * Waits for a direct to task notification on a particular index in the calling
 * task's notification array in a manner similar to taking a counting semaphore.
 *
 * See https://www.FreeRTOS.org/RTOS-task-notifications.html for details.
 *
 * configUSE_TASK_NOTIFICATIONS must be undefined or defined as 1 for this
 * function to be available.
 *
 * Each task has a private array of "notification values" (or 'notifications'),
 * each of which is a 32-bit unsigned integer (uint32_t).  The constant
 * configTASK_NOTIFICATION_ARRAY_ENTRIES sets the number of indexes in the
 * array, and (for backward compatibility) defaults to 1 if left undefined.
 * Prior to FreeRTOS V10.4.0 there was only one notification value per task.
 *
 * Events can be sent to a task using an intermediary object.  Examples of such
 * objects are queues, semaphores, mutexes and event groups.  Task notifications
 * are a method of sending an event directly to a task without the need for such
 * an intermediary object.
 *
 * A notification sent to a task can optionally perform an action, such as
 * update, overwrite or increment one of the task's notification values.  In
 * that way task notifications can be used to send data to a task, or be used as
 * light weight and fast binary or counting semaphores.
 *
 * ulTaskNotifyTakeIndexed() is intended for use when a task notification is
 * used as a faster and lighter weight binary or counting semaphore alternative.
 * Actual FreeRTOS semaphores are taken using the xSemaphoreTake() API function,
 * the equivalent action that instead uses a task notification is
 * ulTaskNotifyTakeIndexed().
 *
 * When a task is using its notification value as a binary or counting semaphore
 * other tasks should send notifications to it using the xTaskNotifyGiveIndexed()
 * macro, or xTaskNotifyIndex() function with the eAction parameter set to
 * eIncrement.
 *
 * ulTaskNotifyTakeIndexed() can either clear the task's notification value at
 * the array index specified by the uxIndexToWaitOn parameter to zero on exit,
 * in which case the notification value acts like a binary semaphore, or
 * decrement the notification value on exit, in which case the notification
 * value acts like a counting semaphore.
 *
 * A task can use ulTaskNotifyTakeIndexed() to [optionally] block to wait for
 * a notification.  The task does not consume any CPU time while it is in the
 * Blocked state.
 *
 * Where as xTaskNotifyWaitIndexed() will return when a notification is pending,
 * ulTaskNotifyTakeIndexed() will return when the task's notification value is
 * not zero.
 *
 * **NOTE** Each notification within the array operates independently - a task
 * can only block on one notification within the array at a time and will not be
 * unblocked by a notification sent to any other array index.
 *
 * Backward compatibility information:
 * Prior to FreeRTOS V10.4.0 each task had a single "notification value", and
 * all task notification API functions operated on that value. Replacing the
 * single notification value with an array of notification values necessitated a
 * new set of API functions that could address specific notifications within the
 * array.  ulTaskNotifyTake() is the original API function, and remains backward
 * compatible by always operating on the notification value at index 0 in the
 * array. Calling ulTaskNotifyTake() is equivalent to calling
 * ulTaskNotifyTakeIndexed() with the uxIndexToWaitOn parameter set to 0.
 *
 * @param uxIndexToWaitOn The index within the calling task's array of
 * notification values on which the calling task will wait for a notification to
 * be non-zero.  uxIndexToWaitOn must be less than
 * configTASK_NOTIFICATION_ARRAY_ENTRIES.  xTaskNotifyTake() does
 * not have this parameter and always waits for notifications on index 0.
 *
 * @param xClearCountOnExit if xClearCountOnExit is pdFALSE then the task's
 * notification value is decremented when the function exits.  In this way the
 * notification value acts like a counting semaphore.  If xClearCountOnExit is
 * not pdFALSE then the task's notification value is cleared to zero when the
 * function exits.  In this way the notification value acts like a binary
 * semaphore.
 *
 * @param xTicksToWait The maximum amount of time that the task should wait in
 * the Blocked state for the task's notification value to be greater than zero,
 * should the count not already be greater than zero when
 * ulTaskNotifyTake() was called.  The task will not consume any processing
 * time while it is in the Blocked state.  This is specified in kernel ticks,
 * the macro pdMS_TO_TICKS( value_in_ms ) can be used to convert a time
 * specified in milliseconds to a time specified in ticks.
 *
 * @return The task's notification count before it is either cleared to zero or
 * decremented (see the xClearCountOnExit parameter).
 *
 * \ingroup TaskNotifications
 */
/** @cond !DOC_EXCLUDE_HEADER_SECTION */
// llgo:link UBaseTypeT.UlTaskGenericNotifyTake C.ulTaskGenericNotifyTake
func (recv_ UBaseTypeT) UlTaskGenericNotifyTake(xClearCountOnExit BaseTypeT, xTicksToWait TickTypeT) c.Uint32T {
	return 0
}

/**
 *
 * See https://www.FreeRTOS.org/RTOS-task-notifications.html for details.
 *
 * configUSE_TASK_NOTIFICATIONS must be undefined or defined as 1 for these
 * functions to be available.
 *
 * Each task has a private array of "notification values" (or 'notifications'),
 * each of which is a 32-bit unsigned integer (uint32_t).  The constant
 * configTASK_NOTIFICATION_ARRAY_ENTRIES sets the number of indexes in the
 * array, and (for backward compatibility) defaults to 1 if left undefined.
 * Prior to FreeRTOS V10.4.0 there was only one notification value per task.
 *
 * If a notification is sent to an index within the array of notifications then
 * the notification at that index is said to be 'pending' until it is read or
 * explicitly cleared by the receiving task.  xTaskNotifyStateClearIndexed()
 * is the function that clears a pending notification without reading the
 * notification value.  The notification value at the same array index is not
 * altered.  Set xTask to NULL to clear the notification state of the calling
 * task.
 *
 * Backward compatibility information:
 * Prior to FreeRTOS V10.4.0 each task had a single "notification value", and
 * all task notification API functions operated on that value. Replacing the
 * single notification value with an array of notification values necessitated a
 * new set of API functions that could address specific notifications within the
 * array.  xTaskNotifyStateClear() is the original API function, and remains
 * backward compatible by always operating on the notification value at index 0
 * within the array. Calling xTaskNotifyStateClear() is equivalent to calling
 * xTaskNotifyStateClearIndexed() with the uxIndexToNotify parameter set to 0.
 *
 * @param xTask The handle of the RTOS task that will have a notification state
 * cleared.  Set xTask to NULL to clear a notification state in the calling
 * task.  To obtain a task's handle create the task using xTaskCreate() and
 * make use of the pxCreatedTask parameter, or create the task using
 * xTaskCreateStatic() and store the returned value, or use the task's name in
 * a call to xTaskGetHandle().
 *
 * @param uxIndexToClear The index within the target task's array of
 * notification values to act upon.  For example, setting uxIndexToClear to 1
 * will clear the state of the notification at index 1 within the array.
 * uxIndexToClear must be less than configTASK_NOTIFICATION_ARRAY_ENTRIES.
 * ulTaskNotifyStateClear() does not have this parameter and always acts on the
 * notification at index 0.
 *
 * @return pdTRUE if the task's notification state was set to
 * eNotWaitingNotification, otherwise pdFALSE.
 *
 * \ingroup TaskNotifications
 */
//go:linkname XTaskGenericNotifyStateClear C.xTaskGenericNotifyStateClear
func XTaskGenericNotifyStateClear(xTask TaskHandleT, uxIndexToClear UBaseTypeT) BaseTypeT

/**
 *
 * See https://www.FreeRTOS.org/RTOS-task-notifications.html for details.
 *
 * configUSE_TASK_NOTIFICATIONS must be undefined or defined as 1 for these
 * functions to be available.
 *
 * Each task has a private array of "notification values" (or 'notifications'),
 * each of which is a 32-bit unsigned integer (uint32_t).  The constant
 * configTASK_NOTIFICATION_ARRAY_ENTRIES sets the number of indexes in the
 * array, and (for backward compatibility) defaults to 1 if left undefined.
 * Prior to FreeRTOS V10.4.0 there was only one notification value per task.
 *
 * ulTaskNotifyValueClearIndexed() clears the bits specified by the
 * ulBitsToClear bit mask in the notification value at array index uxIndexToClear
 * of the task referenced by xTask.
 *
 * Backward compatibility information:
 * Prior to FreeRTOS V10.4.0 each task had a single "notification value", and
 * all task notification API functions operated on that value. Replacing the
 * single notification value with an array of notification values necessitated a
 * new set of API functions that could address specific notifications within the
 * array.  ulTaskNotifyValueClear() is the original API function, and remains
 * backward compatible by always operating on the notification value at index 0
 * within the array. Calling ulTaskNotifyValueClear() is equivalent to calling
 * ulTaskNotifyValueClearIndexed() with the uxIndexToClear parameter set to 0.
 *
 * @param xTask The handle of the RTOS task that will have bits in one of its
 * notification values cleared. Set xTask to NULL to clear bits in a
 * notification value of the calling task.  To obtain a task's handle create the
 * task using xTaskCreate() and make use of the pxCreatedTask parameter, or
 * create the task using xTaskCreateStatic() and store the returned value, or
 * use the task's name in a call to xTaskGetHandle().
 *
 * @param uxIndexToClear The index within the target task's array of
 * notification values in which to clear the bits.  uxIndexToClear
 * must be less than configTASK_NOTIFICATION_ARRAY_ENTRIES.
 * ulTaskNotifyValueClear() does not have this parameter and always clears bits
 * in the notification value at index 0.
 *
 * @param ulBitsToClear Bit mask of the bits to clear in the notification value of
 * xTask. Set a bit to 1 to clear the corresponding bits in the task's notification
 * value. Set ulBitsToClear to 0xffffffff (UINT_MAX on 32-bit architectures) to clear
 * the notification value to 0.  Set ulBitsToClear to 0 to query the task's
 * notification value without clearing any bits.
 *
 *
 * @return The value of the target task's notification value before the bits
 * specified by ulBitsToClear were cleared.
 * \ingroup TaskNotifications
 */
//go:linkname UlTaskGenericNotifyValueClear C.ulTaskGenericNotifyValueClear
func UlTaskGenericNotifyValueClear(xTask TaskHandleT, uxIndexToClear UBaseTypeT, ulBitsToClear c.Uint32T) c.Uint32T

/**
 *
 * Capture the current time for future use with xTaskCheckForTimeOut().
 *
 * @param pxTimeOut Pointer to a timeout object into which the current time
 * is to be captured.  The captured time includes the tick count and the number
 * of times the tick count has overflowed since the system first booted.
 * \ingroup TaskCtrl
 */
// llgo:link (*TimeOutT).VTaskSetTimeOutState C.vTaskSetTimeOutState
func (recv_ *TimeOutT) VTaskSetTimeOutState() {
}

/**
 *
 * Determines if pxTicksToWait ticks has passed since a time was captured
 * using a call to vTaskSetTimeOutState().  The captured time includes the tick
 * count and the number of times the tick count has overflowed.
 *
 * @param pxTimeOut The time status as captured previously using
 * vTaskSetTimeOutState. If the timeout has not yet occurred, it is updated
 * to reflect the current time status.
 * @param pxTicksToWait The number of ticks to check for timeout i.e. if
 * pxTicksToWait ticks have passed since pxTimeOut was last updated (either by
 * vTaskSetTimeOutState() or xTaskCheckForTimeOut()), the timeout has occurred.
 * If the timeout has not occurred, pxTicksToWait is updated to reflect the
 * number of remaining ticks.
 *
 * @return If timeout has occurred, pdTRUE is returned. Otherwise pdFALSE is
 * returned and pxTicksToWait is updated to reflect the number of remaining
 * ticks.
 *
 * @see https://www.FreeRTOS.org/xTaskCheckForTimeOut.html
 *
 * Example Usage:
 * @code{c}
 *  // Driver library function used to receive uxWantedBytes from an Rx buffer
 *  // that is filled by a UART interrupt. If there are not enough bytes in the
 *  // Rx buffer then the task enters the Blocked state until it is notified that
 *  // more data has been placed into the buffer. If there is still not enough
 *  // data then the task re-enters the Blocked state, and xTaskCheckForTimeOut()
 *  // is used to re-calculate the Block time to ensure the total amount of time
 *  // spent in the Blocked state does not exceed MAX_TIME_TO_WAIT. This
 *  // continues until either the buffer contains at least uxWantedBytes bytes,
 *  // or the total amount of time spent in the Blocked state reaches
 *  // MAX_TIME_TO_WAIT - at which point the task reads however many bytes are
 *  // available up to a maximum of uxWantedBytes.
 *
 *  size_t xUART_Receive( uint8_t *pucBuffer, size_t uxWantedBytes )
 *  {
 *  size_t uxReceived = 0;
 *  TickType_t xTicksToWait = MAX_TIME_TO_WAIT;
 *  TimeOut_t xTimeOut;
 *
 *      // Initialize xTimeOut.  This records the time at which this function
 *      // was entered.
 *      vTaskSetTimeOutState( &xTimeOut );
 *
 *      // Loop until the buffer contains the wanted number of bytes, or a
 *      // timeout occurs.
 *      while( UART_bytes_in_rx_buffer( pxUARTInstance ) < uxWantedBytes )
 *      {
 *          // The buffer didn't contain enough data so this task is going to
 *          // enter the Blocked state. Adjusting xTicksToWait to account for
 *          // any time that has been spent in the Blocked state within this
 *          // function so far to ensure the total amount of time spent in the
 *          // Blocked state does not exceed MAX_TIME_TO_WAIT.
 *          if( xTaskCheckForTimeOut( &xTimeOut, &xTicksToWait ) != pdFALSE )
 *          {
 *              //Timed out before the wanted number of bytes were available,
 *              // exit the loop.
 *              break;
 *          }
 *
 *          // Wait for a maximum of xTicksToWait ticks to be notified that the
 *          // receive interrupt has placed more data into the buffer.
 *          ulTaskNotifyTake( pdTRUE, xTicksToWait );
 *      }
 *
 *      // Attempt to read uxWantedBytes from the receive buffer into pucBuffer.
 *      // The actual number of bytes read (which might be less than
 *      // uxWantedBytes) is returned.
 *      uxReceived = UART_read_from_receive_buffer( pxUARTInstance,
 *                                                  pucBuffer,
 *                                                  uxWantedBytes );
 *
 *      return uxReceived;
 *  }
 * @endcode
 * \ingroup TaskCtrl
 */
// llgo:link (*TimeOutT).XTaskCheckForTimeOut C.xTaskCheckForTimeOut
func (recv_ *TimeOutT) XTaskCheckForTimeOut(pxTicksToWait *TickTypeT) BaseTypeT {
	return 0
}

/**
 *
 * This function corrects the tick count value after the application code has held
 * interrupts disabled for an extended period resulting in tick interrupts having
 * been missed.
 *
 * This function is similar to vTaskStepTick(), however, unlike
 * vTaskStepTick(), xTaskCatchUpTicks() may move the tick count forward past a
 * time at which a task should be removed from the blocked state.  That means
 * tasks may have to be removed from the blocked state as the tick count is
 * moved.
 *
 * @param xTicksToCatchUp The number of tick interrupts that have been missed due to
 * interrupts being disabled.  Its value is not computed automatically, so must be
 * computed by the application writer.
 *
 * @return pdTRUE if moving the tick count forward resulted in a task leaving the
 * blocked state and a context switch being performed.  Otherwise pdFALSE.
 *
 * \ingroup TaskCtrl
 */
// llgo:link TickTypeT.XTaskCatchUpTicks C.xTaskCatchUpTicks
func (recv_ TickTypeT) XTaskCatchUpTicks() BaseTypeT {
	return 0
}

/*
 * THIS FUNCTION MUST NOT BE USED FROM APPLICATION CODE.  IT IS ONLY
 * INTENDED FOR USE WHEN IMPLEMENTING A PORT OF THE SCHEDULER AND IS
 * AN INTERFACE WHICH IS FOR THE EXCLUSIVE USE OF THE SCHEDULER.
 *
 * Called from the real time kernel tick (either preemptive or cooperative),
 * this increments the tick count and checks if any tasks that are blocked
 * for a finite period required removing from a blocked list and placing on
 * a ready list.  If a non-zero value is returned then a context switch is
 * required because either:
 *   + A task was removed from a blocked list because its timeout had expired,
 *     or
 *   + Time slicing is in use and there is a task of equal priority to the
 *     currently running task.
 *
 * Note: If configNUMBER_OF_CORES > 1, this function must only be called by
 * core 0. Other cores should call xTaskIncrementTickOtherCores() instead.
 */
//go:linkname XTaskIncrementTick C.xTaskIncrementTick
func XTaskIncrementTick() BaseTypeT

/*
 * THIS FUNCTION MUST NOT BE USED FROM APPLICATION CODE.  IT IS AN
 * INTERFACE WHICH IS FOR THE EXCLUSIVE USE OF THE SCHEDULER.
 *
 * THIS FUNCTION MUST BE CALLED WITH INTERRUPTS DISABLED.
 *
 * Removes the calling task from the ready list and places it both
 * on the list of tasks waiting for a particular event, and the
 * list of delayed tasks.  The task will be removed from both lists
 * and replaced on the ready list should either the event occur (and
 * there be no higher priority tasks waiting on the same event) or
 * the delay period expires.
 *
 * The 'unordered' version replaces the event list item value with the
 * xItemValue value, and inserts the list item at the end of the list.
 *
 * The 'ordered' version uses the existing event list item value (which is the
 * owning task's priority) to insert the list item into the event list in task
 * priority order.
 *
 * @param pxEventList The list containing tasks that are blocked waiting
 * for the event to occur.
 *
 * @param xItemValue The item value to use for the event list item when the
 * event list is not ordered by task priority.
 *
 * @param xTicksToWait The maximum amount of time that the task should wait
 * for the event to occur.  This is specified in kernel ticks, the constant
 * portTICK_PERIOD_MS can be used to convert kernel ticks into a real time
 * period.
 */
// llgo:link (*ListT).VTaskPlaceOnEventList C.vTaskPlaceOnEventList
func (recv_ *ListT) VTaskPlaceOnEventList(xTicksToWait TickTypeT) {
}

// llgo:link (*ListT).VTaskPlaceOnUnorderedEventList C.vTaskPlaceOnUnorderedEventList
func (recv_ *ListT) VTaskPlaceOnUnorderedEventList(xItemValue TickTypeT, xTicksToWait TickTypeT) {
}

/*
 * THIS FUNCTION MUST NOT BE USED FROM APPLICATION CODE.  IT IS AN
 * INTERFACE WHICH IS FOR THE EXCLUSIVE USE OF THE SCHEDULER.
 *
 * THIS FUNCTION MUST BE CALLED WITH INTERRUPTS DISABLED.
 *
 * This function performs nearly the same function as vTaskPlaceOnEventList().
 * The difference being that this function does not permit tasks to block
 * indefinitely, whereas vTaskPlaceOnEventList() does.
 *
 */
// llgo:link (*ListT).VTaskPlaceOnEventListRestricted C.vTaskPlaceOnEventListRestricted
func (recv_ *ListT) VTaskPlaceOnEventListRestricted(xTicksToWait TickTypeT, xWaitIndefinitely BaseTypeT) {
}

/*
 * THIS FUNCTION MUST NOT BE USED FROM APPLICATION CODE.  IT IS AN
 * INTERFACE WHICH IS FOR THE EXCLUSIVE USE OF THE SCHEDULER.
 *
 * THIS FUNCTION MUST BE CALLED WITH INTERRUPTS DISABLED.
 *
 * Removes a task from both the specified event list and the list of blocked
 * tasks, and places it on a ready queue.
 *
 * xTaskRemoveFromEventList()/vTaskRemoveFromUnorderedEventList() will be called
 * if either an event occurs to unblock a task, or the block timeout period
 * expires.
 *
 * xTaskRemoveFromEventList() is used when the event list is in task priority
 * order.  It removes the list item from the head of the event list as that will
 * have the highest priority owning task of all the tasks on the event list.
 * vTaskRemoveFromUnorderedEventList() is used when the event list is not
 * ordered and the event list items hold something other than the owning tasks
 * priority.  In this case the event list item value is updated to the value
 * passed in the xItemValue parameter.
 *
 * @return pdTRUE if the task being removed has a higher priority than the task
 * making the call, otherwise pdFALSE.
 */
// llgo:link (*ListT).XTaskRemoveFromEventList C.xTaskRemoveFromEventList
func (recv_ *ListT) XTaskRemoveFromEventList() BaseTypeT {
	return 0
}

// llgo:link (*ListItemT).VTaskRemoveFromUnorderedEventList C.vTaskRemoveFromUnorderedEventList
func (recv_ *ListItemT) VTaskRemoveFromUnorderedEventList(xItemValue TickTypeT) {
}

/*
 * THIS FUNCTION MUST NOT BE USED FROM APPLICATION CODE.  IT IS ONLY
 * INTENDED FOR USE WHEN IMPLEMENTING A PORT OF THE SCHEDULER AND IS
 * AN INTERFACE WHICH IS FOR THE EXCLUSIVE USE OF THE SCHEDULER.
 *
 * Sets the pointer to the current TCB to the TCB of the highest priority task
 * that is ready to run.
 */
//go:linkname VTaskSwitchContext C.vTaskSwitchContext
func VTaskSwitchContext() c.Int

/*
 * THESE FUNCTIONS MUST NOT BE USED FROM APPLICATION CODE.  THEY ARE USED BY
 * THE EVENT BITS MODULE.
 */
//go:linkname UxTaskResetEventItemValue C.uxTaskResetEventItemValue
func UxTaskResetEventItemValue() TickTypeT

/*
 * Return the handle of the calling task.
 */
//go:linkname XTaskGetCurrentTaskHandle C.xTaskGetCurrentTaskHandle
func XTaskGetCurrentTaskHandle() TaskHandleT

/*
 * Shortcut used by the queue implementation to prevent unnecessary call to
 * taskYIELD();
 */
//go:linkname VTaskMissedYield C.vTaskMissedYield
func VTaskMissedYield()

/*
 * Returns the scheduler state as taskSCHEDULER_RUNNING,
 * taskSCHEDULER_NOT_STARTED or taskSCHEDULER_SUSPENDED.
 */
//go:linkname XTaskGetSchedulerState C.xTaskGetSchedulerState
func XTaskGetSchedulerState() BaseTypeT

/*
 * Raises the priority of the mutex holder to that of the calling task should
 * the mutex holder have a priority less than the calling task.
 */
//go:linkname XTaskPriorityInherit C.xTaskPriorityInherit
func XTaskPriorityInherit(pxMutexHolder TaskHandleT) BaseTypeT

/*
 * Set the priority of a task back to its proper priority in the case that it
 * inherited a higher priority while it was holding a semaphore.
 */
//go:linkname XTaskPriorityDisinherit C.xTaskPriorityDisinherit
func XTaskPriorityDisinherit(pxMutexHolder TaskHandleT) BaseTypeT

/*
 * If a higher priority task attempting to obtain a mutex caused a lower
 * priority task to inherit the higher priority task's priority - but the higher
 * priority task then timed out without obtaining the mutex, then the lower
 * priority task will disinherit the priority again - but only down as far as
 * the highest priority task that is still waiting for the mutex (if there were
 * more than one task waiting for the mutex).
 */
//go:linkname VTaskPriorityDisinheritAfterTimeout C.vTaskPriorityDisinheritAfterTimeout
func VTaskPriorityDisinheritAfterTimeout(pxMutexHolder TaskHandleT, uxHighestPriorityWaitingTask UBaseTypeT)

/*
 * For internal use only.  Increment the mutex held count when a mutex is
 * taken and return the handle of the task that has taken the mutex.
 */
//go:linkname PvTaskIncrementMutexHeldCount C.pvTaskIncrementMutexHeldCount
func PvTaskIncrementMutexHeldCount() TaskHandleT

/*
 * For internal use only.  Same as vTaskSetTimeOutState(), but without a critical
 * section.
 */
// llgo:link (*TimeOutT).VTaskInternalSetTimeOutState C.vTaskInternalSetTimeOutState
func (recv_ *TimeOutT) VTaskInternalSetTimeOutState() {
}
