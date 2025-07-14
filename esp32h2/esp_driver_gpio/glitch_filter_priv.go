package esp_driver_gpio

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const GLITCH_FILTER_PM_LOCK_NAME_LEN_MAX = 16

type GlitchFilterFsmT c.Int

const (
	GLITCH_FILTER_FSM_INIT   GlitchFilterFsmT = 0
	GLITCH_FILTER_FSM_ENABLE GlitchFilterFsmT = 1
)
