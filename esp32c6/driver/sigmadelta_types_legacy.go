package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SigmadeltaPortT c.Int

const (
	SIGMADELTA_PORT_0   SigmadeltaPortT = 0
	SIGMADELTA_PORT_MAX SigmadeltaPortT = 1
)

type SigmadeltaChannelT c.Int

const (
	SIGMADELTA_CHANNEL_0   SigmadeltaChannelT = 0
	SIGMADELTA_CHANNEL_1   SigmadeltaChannelT = 1
	SIGMADELTA_CHANNEL_2   SigmadeltaChannelT = 2
	SIGMADELTA_CHANNEL_3   SigmadeltaChannelT = 3
	SIGMADELTA_CHANNEL_MAX SigmadeltaChannelT = 4
)

/**
 * @brief Sigma-delta configure struct
 */

type SigmadeltaConfigT struct {
	Channel            SigmadeltaChannelT
	SigmadeltaDuty     c.Int8T
	SigmadeltaPrescale c.Uint8T
	SigmadeltaGpio     GpioNumT
}
