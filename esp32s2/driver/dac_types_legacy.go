package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DacCwScaleT c.Int

const (
	DAC_CW_SCALE_1 DacCwScaleT = 0
	DAC_CW_SCALE_2 DacCwScaleT = 1
	DAC_CW_SCALE_4 DacCwScaleT = 2
	DAC_CW_SCALE_8 DacCwScaleT = 3
)

type DacCwPhaseT c.Int

const (
	DAC_CW_PHASE_0   DacCwPhaseT = 2
	DAC_CW_PHASE_180 DacCwPhaseT = 3
)

type DacDigiConvertModeT c.Int

const (
	DAC_CONV_NORMAL DacDigiConvertModeT = 0
	DAC_CONV_ALTER  DacDigiConvertModeT = 1
	DAC_CONV_MAX    DacDigiConvertModeT = 2
)

/**
 * @brief DAC digital controller (DMA mode) configuration parameters.
 */

type DacDigiConfigT struct {
	Mode     DacDigiConvertModeT
	Interval c.Uint32T
	DigClk   AdcDigiClkT
}

/**
 * @brief Config the cosine wave generator function in DAC module.
 */

type DacCwConfigT struct {
	EnCh   DacChannelT
	Scale  DacCwScaleT
	Phase  DacCwPhaseT
	Freq   c.Uint32T
	Offset c.Int8T
}
