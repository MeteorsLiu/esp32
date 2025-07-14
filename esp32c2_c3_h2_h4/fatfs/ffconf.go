package fatfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const FFCONF_DEF = 80286
const FF_FS_READONLY = 0
const FF_FS_MINIMIZE = 0
const FF_USE_FIND = 0
const FF_USE_MKFS = 1
const FF_USE_EXPAND = 1
const FF_USE_CHMOD = 1
const FF_USE_FORWARD = 0
const FF_USE_STRFUNC = 0
const FF_PRINT_LLI = 0
const FF_PRINT_FLOAT = 0
const FF_STRF_ENCODE = 3
const FF_USE_LFN = 0
const FF_LFN_UNICODE = 0
const FF_LFN_BUF = 255
const FF_SFN_BUF = 12
const FF_FS_RPATH = 0
const FF_STR_VOLUME_ID = 0
const FF_MULTI_PARTITION = 1
const FF_SS_SDCARD = 512
const FF_LBA64 = 0
const FF_MIN_GPT = 0x10000000
const FF_USE_TRIM = 1
const FF_FS_EXFAT = 0
const FF_FS_NORTC = 0
const FF_NORTC_MON = 1
const FF_NORTC_MDAY = 1
const FF_NORTC_YEAR = 2022
const FF_FS_REENTRANT = 1

/* Some memory allocation functions are declared here in addition to ff.h, so that
   they can be used also by external code when LFN feature is disabled.
*/
//go:linkname FfMemalloc C.ff_memalloc
func FfMemalloc(msize c.Uint) c.Pointer

//go:linkname FfMemfree C.ff_memfree
func FfMemfree(c.Pointer)
