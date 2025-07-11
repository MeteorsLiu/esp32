package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
ESP32 ROM code contains implementations of some of C library functions.
Whenever a function in ROM needs to use a syscall, it calls a pointer to the corresponding syscall
implementation defined in the following struct.

The table itself, by default, is not allocated in RAM. There are two pointers, `syscall_table_ptr_pro` and
`syscall_table_ptr_app`, which can be set to point to the locations of syscall tables of CPU 0 (aka PRO CPU)
and CPU 1 (aka APP CPU). Location of these pointers in .bss segment of ROM code is defined in linker script.

So, before using any of the C library functions (except for pure functions and memcpy/memset functions),
application must allocate syscall table structure for each CPU being used, and populate it with pointers
to actual implementations of corresponding syscalls.
*/
type SyscallStubTable struct {
	X__getreent               c.Pointer
	X_mallocR                 c.Pointer
	X_freeR                   c.Pointer
	X_reallocR                c.Pointer
	X_callocR                 c.Pointer
	X_abort                   c.Pointer
	X_systemR                 c.Pointer
	X_renameR                 c.Pointer
	X_timesR                  c.Pointer
	X_gettimeofdayR           c.Pointer
	X_raiseR                  c.Pointer
	X_unlinkR                 c.Pointer
	X_linkR                   c.Pointer
	X_statR                   c.Pointer
	X_fstatR                  c.Pointer
	X_sbrkR                   c.Pointer
	X_getpidR                 c.Pointer
	X_killR                   c.Pointer
	X_exitR                   c.Pointer
	X_closeR                  c.Pointer
	X_openR                   c.Pointer
	X_writeR                  c.Pointer
	X_lseekR                  c.Pointer
	X_readR                   c.Pointer
	X_lockInit                c.Pointer
	X_lockInitRecursive       c.Pointer
	X_lockClose               c.Pointer
	X_lockCloseRecursive      c.Pointer
	X_lockAcquire             c.Pointer
	X_lockAcquireRecursive    c.Pointer
	X_lockTryAcquire          c.Pointer
	X_lockTryAcquireRecursive c.Pointer
	X_lockRelease             c.Pointer
	X_lockReleaseRecursive    c.Pointer
	X_printfFloat             c.Pointer
	X_scanfFloat              c.Pointer
}
