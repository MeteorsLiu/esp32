package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
ESP32-C3 ROM code contains implementations of some of C library functions.
Whenever a function in ROM needs to use a syscall, it calls a pointer to the corresponding syscall
implementation defined in the following struct.

The table itself, by default, is not allocated in RAM. A global pointer syscall_table_ptr is used to
set the address

So, before using any of the C library functions (except for pure functions and memcpy/memset functions),
application must allocate syscall table structure for each CPU being used, and populate it with pointers
to actual implementations of corresponding syscalls.
*/
type SyscallStubTable struct {
	X__getreent                       c.Pointer
	X_mallocR                         c.Pointer
	X_freeR                           c.Pointer
	X_reallocR                        c.Pointer
	X_callocR                         c.Pointer
	X_abort                           c.Pointer
	X_systemR                         c.Pointer
	X_renameR                         c.Pointer
	X_timesR                          c.Pointer
	X_gettimeofdayR                   c.Pointer
	X_raiseR                          c.Pointer
	X_unlinkR                         c.Pointer
	X_linkR                           c.Pointer
	X_statR                           c.Pointer
	X_fstatR                          c.Pointer
	X_sbrkR                           c.Pointer
	X_getpidR                         c.Pointer
	X_killR                           c.Pointer
	X_exitR                           c.Pointer
	X_closeR                          c.Pointer
	X_openR                           c.Pointer
	X_writeR                          c.Pointer
	X_lseekR                          c.Pointer
	X_readR                           c.Pointer
	X_retargetLockInit                c.Pointer
	X_retargetLockInitRecursive       c.Pointer
	X_retargetLockClose               c.Pointer
	X_retargetLockCloseRecursive      c.Pointer
	X_retargetLockAcquire             c.Pointer
	X_retargetLockAcquireRecursive    c.Pointer
	X_retargetLockTryAcquire          c.Pointer
	X_retargetLockTryAcquireRecursive c.Pointer
	X_retargetLockRelease             c.Pointer
	X_retargetLockReleaseRecursive    c.Pointer
	X_printfFloat                     c.Pointer
	X_scanfFloat                      c.Pointer
	X__assertFunc                     c.Pointer
	X__sinit                          c.Pointer
	X_cleanupR                        c.Pointer
}
