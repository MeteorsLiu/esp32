package heap

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CONFIG_HEAP_TRACING_STACK_DEPTH = 0

type HeapTraceModeT c.Int

const (
	HEAP_TRACE_ALL   HeapTraceModeT = 0
	HEAP_TRACE_LEAKS HeapTraceModeT = 1
)

/**
 * @brief Trace record data type. Stores information about an allocated region of memory.
 */

type HeapTraceRecordT struct {
	Ccount    c.Uint32T
	Address   c.Pointer
	Size      c.SizeT
	AllocedBy [0]c.Pointer
	FreedBy   [0]c.Pointer
}

/**
 * @brief Stores information about the result of a heap trace.
 */

type HeapTraceSummaryT struct {
	Mode             HeapTraceModeT
	TotalAllocations c.SizeT
	TotalFrees       c.SizeT
	Count            c.SizeT
	Capacity         c.SizeT
	HighWaterMark    c.SizeT
	HasOverflowed    c.SizeT
}
