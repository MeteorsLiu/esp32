package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const PORT_OFFSET_PX_STACK = 0x30
const CORE_ID_SIZE = 0
const PortCRITICAL_NESTING_IN_TCB = 0
const PortBYTE_ALIGNMENT = 16
const PortTICK_TYPE_IS_ATOMIC = 1

type StackTypeT c.Uint8T
type UBaseTypeT c.Uint
type TickTypeT c.Uint32T

/**
 * @brief Disable interrupts in a nested manner (meant to be called from ISRs)
 *
 * @warning Only applies to current CPU.
 * @return UBaseType_t Previous interrupt level
 */
//go:linkname XPortSetInterruptMaskFromISR C.xPortSetInterruptMaskFromISR
func XPortSetInterruptMaskFromISR() UBaseTypeT

/**
 * @brief Re-enable interrupts in a nested manner (meant to be called from ISRs)
 *
 * @warning Only applies to current CPU.
 * @param prev_int_level Previous interrupt level
 */
// llgo:link UBaseTypeT.VPortClearInterruptMaskFromISR C.vPortClearInterruptMaskFromISR
func (recv_ UBaseTypeT) VPortClearInterruptMaskFromISR() {
}

/**
 * @brief Checks if the current core is in an ISR context
 *
 * - ISR context consist of Low/Mid priority ISR, or time tick ISR
 * - High priority ISRs aren't detected here, but they normally cannot call C code, so that should not be an issue anyway.
 *
 * @note [refactor-todo] Check if this should be inlined
 * @return
 *  - pdTRUE if in ISR
 *  - pdFALSE otherwise
 */
//go:linkname XPortInIsrContext C.xPortInIsrContext
func XPortInIsrContext() c.Int

/**
 * @brief Assert if in ISR context
 *
 * - Asserts on xPortInIsrContext() internally
 */
//go:linkname VPortAssertIfInISR C.vPortAssertIfInISR
func VPortAssertIfInISR()

/**
 * @brief Check if in ISR context from High priority ISRs
 *
 * - Called from High priority ISR
 * - Checks if the previous context (before high priority interrupt) was in ISR context (meaning low/med priority)
 *
 * @note [refactor-todo] Check if this should be inlined
 * @return
 *  - pdTRUE if in previous in ISR context
 *  - pdFALSE otherwise
 */
//go:linkname XPortInterruptedFromISRContext C.xPortInterruptedFromISRContext
func XPortInterruptedFromISRContext() c.Int

type PortMUXTYPE SpinlockT

/**
 * @brief Enter a critical section
 *
 * - Simply disable interrupts
 * - Can be nested
 */
//go:linkname VPortEnterCritical C.vPortEnterCritical
func VPortEnterCritical()

/**
 * @brief Exit a critical section
 *
 * - Reenables interrupts
 * - Can be nested
 */
//go:linkname VPortExitCritical C.vPortExitCritical
func VPortExitCritical()

/**
 * @brief Perform a context switch from a task
 *
 * @note [refactor-todo] The rest of ESP-IDF should call taskYield() instead
 */
//go:linkname VPortYield C.vPortYield
func VPortYield()

/**
 * @brief Perform a context switch from an ISR
 */
//go:linkname VPortYieldFromISR C.vPortYieldFromISR
func VPortYieldFromISR()

/**
 * @brief Yields the other core
 *
 * @note Added to be compatible with SMP API
 * @note [refactor-todo] Put this into private macros as its only called from task.c and is not public API
 * @param coreid ID of core to yield
 */
// llgo:link BaseTypeT.VPortYieldOtherCore C.vPortYieldOtherCore
func (recv_ BaseTypeT) VPortYieldOtherCore() {
}

/**
 * @brief Get the tick rate per second
 *
 * @note [refactor-todo] make this inline
 * @note [refactor-todo] Check if this function should be renamed (due to uint return type)
 * @return uint32_t Tick rate in Hz
 */
//go:linkname XPortGetTickRateHz C.xPortGetTickRateHz
func XPortGetTickRateHz() c.Uint32T

/**
 * @brief Set a watchpoint to watch the last 32 bytes of the stack
 *
 * Callback to set a watchpoint on the end of the stack. Called every context switch to change the stack watchpoint
 * around.
 *
 * @param pxStackStart Pointer to the start of the stack
 */
//go:linkname VPortSetStackWatchpoint C.vPortSetStackWatchpoint
func VPortSetStackWatchpoint(pxStackStart c.Pointer)

/**
 * @brief TCB cleanup hook
 *
 * The portCLEAN_UP_TCB() macro is called in prvDeleteTCB() right before a
 * deleted task's memory is freed. We map that macro to this internal function
 * so that IDF FreeRTOS ports can inject some task pre-deletion operations.
 *
 * @note We can't use vPortCleanUpTCB() due to API compatibility issues. See
 * CONFIG_FREERTOS_ENABLE_STATIC_TASK_CLEAN_UP. Todo: IDF-8097
 */
//go:linkname VPortTCBPreDeleteHook C.vPortTCBPreDeleteHook
func VPortTCBPreDeleteHook(pxTCB c.Pointer)

/**
 * @brief Checks if a given piece of memory can be used to store a FreeRTOS list
 *
 * - Defined in heap_idf.c
 *
 * @param ptr Pointer to memory
 * @return true Memory can be used to store a List
 * @return false Otherwise
 */
//go:linkname XPortCheckValidListMem C.xPortCheckValidListMem
func XPortCheckValidListMem(ptr c.Pointer) bool

/**
 * @brief Checks if a given piece of memory can be used to store a task's TCB
 *
 * - Defined in heap_idf.c
 *
 * @param ptr Pointer to memory
 * @return true Memory can be used to store a TCB
 * @return false Otherwise
 */
//go:linkname XPortCheckValidTCBMem C.xPortCheckValidTCBMem
func XPortCheckValidTCBMem(ptr c.Pointer) bool

/**
 * @brief Checks if a given piece of memory can be used to store a task's stack
 *
 * - Defined in heap_idf.c
 *
 * @param ptr Pointer to memory
 * @return true Memory can be used to store a task stack
 * @return false Otherwise
 */
//go:linkname XPortcheckValidStackMem C.xPortcheckValidStackMem
func XPortcheckValidStackMem(ptr c.Pointer) bool
