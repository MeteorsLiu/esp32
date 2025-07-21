package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const PortUSING_MPU_WRAPPERS = 0
const PortNUM_CONFIGURABLE_REGIONS = 1
const PortHAS_STACK_OVERFLOW_CHECKING = 0
const ConfigSTACK_ALLOCATION_FROM_SEPARATE_HEAP = 0

// llgo:link (*StackTypeT).PxPortInitialiseStack C.pxPortInitialiseStack
func (recv_ *StackTypeT) PxPortInitialiseStack(pxCode c.Int, pvParameters c.Pointer) *StackTypeT {
	return nil
}

/* Used by heap_5.c to define the start address and size of each memory region
 * that together comprise the total FreeRTOS heap space. */

type HeapRegion struct {
	PucStartAddress *c.Uint8T
	XSizeInBytes    c.SizeT
}
type HeapRegionT HeapRegion

/* Used to pass information about the heap out of vPortGetHeapStats(). */

type XHeapStats struct {
	XAvailableHeapSpaceInBytes      c.SizeT
	XSizeOfLargestFreeBlockInBytes  c.SizeT
	XSizeOfSmallestFreeBlockInBytes c.SizeT
	XNumberOfFreeBlocks             c.SizeT
	XMinimumEverFreeBytesRemaining  c.SizeT
	XNumberOfSuccessfulAllocations  c.SizeT
	XNumberOfSuccessfulFrees        c.SizeT
}
type HeapStatsT XHeapStats

/*
 * Map to the memory management routines required for the port.
 */
//go:linkname PvPortMalloc C.pvPortMalloc
func PvPortMalloc(xSize c.SizeT) c.Pointer

//go:linkname VPortFree C.vPortFree
func VPortFree(pv c.Pointer)

/*
 * Setup the hardware ready for the scheduler to take control.  This generally
 * sets up a tick interrupt and sets timers for the correct tick frequency.
 */
//go:linkname XPortStartScheduler C.xPortStartScheduler
func XPortStartScheduler() c.Int

/*
 * Undo any hardware/ISR setup that was performed by xPortStartScheduler() so
 * the hardware is left in its original condition after the scheduler stops
 * executing.
 */
//go:linkname VPortEndScheduler C.vPortEndScheduler
func VPortEndScheduler()
