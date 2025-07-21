package esp_system

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const UNW_UNKNOWN_TARGET = 1
const UNW_ESUCCESS = 0
const UNW_EUNSPEC = 1
const UNW_EBADREG = 3
const UNW_ESTOPUNWIND = 5
const UNW_EINVAL = 8
const UNW_EBADVERSION = 9
const UNW_ENOINFO = 10

type ExecutionFrame c.Pointer
type UnwContextT ExecutionFrame
type UnwRegnumT c.Uint32T
type UnwCursorT UnwContextT
type UnwWordT c.Ulong
type UnwAddrSpaceT c.Pointer
type UnwFpregT c.Pointer
