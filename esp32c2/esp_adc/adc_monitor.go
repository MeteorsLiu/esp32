package esp_adc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcMonitorT struct {
	Unused [8]uint8
}
type AdcMonitorHandleT *AdcMonitorT

/**
 * @brief ADC digital controller (DMA mode) monitor configuration.
 */

type AdcMonitorConfigT struct {
	AdcUnit    AdcUnitT
	Channel    AdcChannelT
	HThreshold c.Int32T
	LThreshold c.Int32T
}

/**
 * @brief Type of adc monitor event data
 */

type AdcMonitorEvtData struct {
	Unused [8]uint8
}
type AdcMonitorEvtDataT AdcMonitorEvtData

// llgo:type C
type AdcMonitorEvtCbT func(AdcMonitorHandleT, *AdcMonitorEvtDataT, c.Pointer) bool

/**
 * @brief Struct type of many different adc_monitor evt callbacks.
 */

type AdcMonitorEvtCbsT struct {
	OnOverHighThresh AdcMonitorEvtCbT
	OnBelowLowThresh AdcMonitorEvtCbT
}
