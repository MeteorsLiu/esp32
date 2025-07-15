package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get the current core ID of a particular task
 *
 * Helper function to get the core ID of a particular task. If the task is
 * pinned to a particular core, the core ID is returned. If the task is not
 * pinned to a particular core, tskNO_AFFINITY is returned.
 *
 * If CONFIG_FREERTOS_UNICORE is enabled, this function simply returns 0.
 *
 * [refactor-todo] See if this needs to be deprecated (IDF-8145)(IDF-8164)
 *
 * @note If CONFIG_FREERTOS_SMP is enabled, please call vTaskCoreAffinityGet()
 * instead.
 * @note In IDF FreerTOS when configNUMBER_OF_CORES == 1, this function will
 * always return 0,
 * @param xTask The task to query
 * @return The task's core ID or tskNO_AFFINITY
 */
//go:linkname XTaskGetCoreID C.xTaskGetCoreID
func XTaskGetCoreID(xTask TaskHandleT) BaseTypeT

/**
 * Returns the start of the stack associated with xTask.
 *
 * Returns the lowest stack memory address, regardless of whether the stack
 * grows up or down.
 *
 * [refactor-todo] Change return type to StackType_t (IDF-8158)
 *
 * @param xTask Handle of the task associated with the stack returned.
 * Set xTask to NULL to return the stack of the calling task.
 *
 * @return A pointer to the start of the stack.
 */
//go:linkname PxTaskGetStackStart C.pxTaskGetStackStart
func PxTaskGetStackStart(xTask TaskHandleT) *c.Uint8T

// llgo:type C
type TlsDeleteCallbackFunctionT func(c.Int, c.Pointer)

/**
 * Set local storage pointer and deletion callback.
 *
 * Each task contains an array of pointers that is dimensioned by the
 * configNUM_THREAD_LOCAL_STORAGE_POINTERS setting in FreeRTOSConfig.h. The
 * kernel does not use the pointers itself, so the application writer can use
 * the pointers for any purpose they wish.
 *
 * Local storage pointers set for a task can reference dynamically allocated
 * resources. This function is similar to vTaskSetThreadLocalStoragePointer, but
 * provides a way to release these resources when the task gets deleted. For
 * each pointer, a callback function can be set. This function will be called
 * when task is deleted, with the local storage pointer index and value as
 * arguments.
 *
 * @param xTaskToSet  Task to set thread local storage pointer for
 * @param xIndex The index of the pointer to set, from 0 to
 * configNUM_THREAD_LOCAL_STORAGE_POINTERS - 1.
 * @param pvValue Pointer value to set.
 * @param pvDelCallback  Function to call to dispose of the local storage
 * pointer when the task is deleted.
 */
//go:linkname VTaskSetThreadLocalStoragePointerAndDelCallback C.vTaskSetThreadLocalStoragePointerAndDelCallback
func VTaskSetThreadLocalStoragePointerAndDelCallback(xTaskToSet TaskHandleT, xIndex BaseTypeT, pvValue c.Pointer, pvDelCallback TlsDeleteCallbackFunctionT)

// llgo:link UBaseTypeT.XSemaphoreCreateGenericWithCaps C.xSemaphoreCreateGenericWithCaps
func (recv_ UBaseTypeT) XSemaphoreCreateGenericWithCaps(uxInitialCount UBaseTypeT, ucQueueType c.Uint8T, uxMemoryCaps UBaseTypeT) SemaphoreHandleT {
	return nil
}

//go:linkname XStreamBufferGenericCreateWithCaps C.xStreamBufferGenericCreateWithCaps
func XStreamBufferGenericCreateWithCaps(xBufferSizeBytes c.SizeT, xTriggerLevelBytes c.SizeT, xIsMessageBuffer BaseTypeT, uxMemoryCaps UBaseTypeT) StreamBufferHandleT

//go:linkname VStreamBufferGenericDeleteWithCaps C.vStreamBufferGenericDeleteWithCaps
func VStreamBufferGenericDeleteWithCaps(xStreamBuffer StreamBufferHandleT, xIsMessageBuffer BaseTypeT)
