package esp_driver_dac

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DacContinuousChannelModeT c.Int

const (
	DAC_CHANNEL_MODE_SIMUL DacContinuousChannelModeT = 0
	DAC_CHANNEL_MODE_ALTER DacContinuousChannelModeT = 1
)

type DacContinuousDigiClkSrcT SocPeriphDacDigiClkSrcT
type DacCosineClkSrcT SocPeriphDacCosineClkSrcT
