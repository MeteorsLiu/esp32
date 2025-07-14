package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DebugProbeSplitU16T c.Int

const (
	DEBUG_PROBE_SPLIT_LOWER16 DebugProbeSplitU16T = 0
	DEBUG_PROBE_SPLIT_UPPER16 DebugProbeSplitU16T = 1
)
