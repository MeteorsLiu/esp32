package esp_adc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcContinuousCtxT struct {
	Unused [8]uint8
}
type AdcContinuousHandleT *AdcContinuousCtxT

/**
 * @brief ADC continuous mode driver initial configurations
 */

type AdcContinuousHandleCfgT struct {
	MaxStoreBufSize c.Uint32T
	ConvFrameSize   c.Uint32T
	Flags           struct {
		FlushPool c.Uint32T
	}
}

/**
 * @brief ADC continuous mode driver configurations
 */

type AdcContinuousConfigT struct {
	PatternNum   c.Uint32T
	AdcPattern   *AdcDigiPatternConfigT
	SampleFreqHz c.Uint32T
	ConvMode     AdcDigiConvertModeT
	Format       AdcDigiOutputFormatT
}

/**
 * @brief Event data structure
 * @note The `conv_frame_buffer` is maintained by the driver itself, so never free this piece of memory.
 */

type AdcContinuousEvtDataT struct {
	ConvFrameBuffer *c.Uint8T
	Size            c.Uint32T
}

// llgo:type C
type AdcContinuousCallbackT func(AdcContinuousHandleT, *AdcContinuousEvtDataT, c.Pointer) bool

/**
 * @brief Group of ADC continuous mode callbacks
 *
 * @note These callbacks are all running in an ISR environment.
 * @note When CONFIG_ADC_CONTINUOUS_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables should be in internal RAM as well.
 */

type AdcContinuousEvtCbsT struct {
	OnConvDone AdcContinuousCallbackT
	OnPoolOvf  AdcContinuousCallbackT
}
