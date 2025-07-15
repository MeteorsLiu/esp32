package esp_system

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * @brief   Structure used for backtracing
 *
 * This structure stores the backtrace information of a particular stack frame
 * (i.e. the PC and SP). This structure is used iteratively with the
 * esp_cpu_get_next_backtrace_frame() function to traverse each frame within a
 * single stack. The next_pc represents the PC of the current frame's caller, thus
 * a next_pc of 0 indicates that the current frame is the last frame on the stack.
 *
 * @note    Call esp_backtrace_get_start() to obtain initialization values for
 *          this structure
 */

type EspBacktraceFrameT struct {
	Pc       c.Uint32T
	Sp       c.Uint32T
	NextPc   c.Uint32T
	ExcFrame c.Pointer
}

/**
 * @brief Print the backtrace of the current stack
 *
 * @param depth The maximum number of stack frames to print (should be > 0)
 *
 * @note On RISC-V targets printing backtrace at run-time is only available if
 *       CONFIG_ESP_SYSTEM_USE_EH_FRAME is selected. Otherwise we simply print
 *       a register dump. Function assumes it is called in a context where the
 *       calling task will not migrate to another core, e.g. interrupts disabled/panic handler.
 *
 * @return
 *      - ESP_OK    Backtrace successfully printed to completion or to depth limit
 *      - ESP_FAIL  Backtrace is corrupted
 */
//go:linkname EspBacktracePrint C.esp_backtrace_print
func EspBacktracePrint(depth c.Int) EspErrT
