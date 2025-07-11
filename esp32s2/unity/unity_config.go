package unity

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

//go:linkname UnityFlush C.unity_flush
func UnityFlush()

//go:linkname UnityPutc C.unity_putc
func UnityPutc(c c.Int)

//go:linkname UnityGets C.unity_gets
func UnityGets(dst *c.Char, len c.SizeT)

//go:linkname UnityExecTimeStart C.unity_exec_time_start
func UnityExecTimeStart()

//go:linkname UnityExecTimeStop C.unity_exec_time_stop
func UnityExecTimeStop()

//go:linkname UnityExecTimeGetMs C.unity_exec_time_get_ms
func UnityExecTimeGetMs() c.Uint32T
