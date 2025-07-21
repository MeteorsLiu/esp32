package freertos

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
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
type BaseTypeT func(*c.Int) c.Int
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
	XTimeOnEntering c.Int
}
type TimeOutT XTIMEOUT

/*
 * Defines the memory ranges allocated to the task when an MPU is used.
 */

type XMEMORYREGION struct {
	PvBaseAddress   c.Pointer
	UlLengthInBytes c.Int
	UlParameters    c.Int
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
	UxPriority     c.Int
	PuxStackBuffer *c.Int
	XRegions       MemoryRegionT
}
type TaskParametersT XTASKPARAMETERS

/** Used with the uxTaskGetSystemState() function to return the state of each task
 * in the system. */

type XTASKSTATUS struct {
	XHandle              TaskHandleT
	PcTaskName           *c.Char
	XTaskNumber          c.Int
	ECurrentState        ETaskState
	UxCurrentPriority    c.Int
	UxBasePriority       c.Int
	UlRunTimeCounter     c.Int
	PxStackBase          *c.Int
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
