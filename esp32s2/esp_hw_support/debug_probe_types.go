package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const DEBUG_PROBE_MAX_OUTPUT_WIDTH = 16

type DebugProbeTargetT c.Int
