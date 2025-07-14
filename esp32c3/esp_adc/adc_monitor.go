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

/**
 * @brief Allocate an ADC continuous mode monitor (and configure it into an initial state)
 *
 * @param[in]  handle           ADC continuous mode driver handle
 * @param[in]  monitor_cfg      ADC monitor config struct
 * @param[out] ret_handle       Handle of allocated monitor
 *
 * @return
 *       - ESP_OK:                  On success
 *       - ESP_ERR_INVALID_ARG:     Invalid argument
 *       - ESP_ERR_INVALID_STATE：  Install monitor when ADC converter is running
 *       - ESP_ERR_NOT_FOUND：      No appropriate monitor Or no free monitor
 *       - ESP_ERR_NOT_SUPPORTED:   Threshold configuration not supported
 *       - ESP_ERR_NO_MEM:          Free memory not enough
 */
//go:linkname AdcNewContinuousMonitor C.adc_new_continuous_monitor
func AdcNewContinuousMonitor(handle AdcContinuousHandleT, monitor_cfg *AdcMonitorConfigT, ret_handle *AdcMonitorHandleT) EspErrT

/**
 * @brief Register/Unregister threshold interrupt callbacks for allocated monitor.
 *        Passing `cbs` contain the NULL `over_high/below_low` will unregister relative callbacks.
 *
 * @param[in]  monitor_handle       Monitor handle
 * @param[in]  cbs                  Pointer to a adc_monitor_evt_cbs_t struct
 * @param[in]  user_data            User data, which will be delivered to the callback functions directly
 *
 * @return
 *       - ESP_OK:                  On success
 *       - ESP_ERR_INVALID_STATE:   To register cbs when monitor is running, or cbs has been installed
 *       - ESP_ERR_NOT_SUPPORTED:   Register not supported callbacks to esp32s2
 *       - ESP_ERR_INVALID_ARG:     Invalid argument
 */
//go:linkname AdcContinuousMonitorRegisterEventCallbacks C.adc_continuous_monitor_register_event_callbacks
func AdcContinuousMonitorRegisterEventCallbacks(monitor_handle AdcMonitorHandleT, cbs *AdcMonitorEvtCbsT, user_data c.Pointer) EspErrT

/**
 * @brief Enable an ADC continuous mode monitor.
 *
 * @param[in]  monitor_handle       Monitor handle
 *
 * @return
 *       - ESP_OK:                  On success
 *       - ESP_ERR_INVALID_STATE:   Monitor has enabled, no need to enable again
 *       - ESP_ERR_INVALID_ARG:     Invalid argument
 */
//go:linkname AdcContinuousMonitorEnable C.adc_continuous_monitor_enable
func AdcContinuousMonitorEnable(monitor_handle AdcMonitorHandleT) EspErrT

/**
 * @brief Disable an ADC continuous mode monitor.
 *
 * @param[in]  monitor_handle       Monitor handle
 *
 * @return
 *       - ESP_OK:                  On success
 *       - ESP_ERR_INVALID_STATE:   Monitor not enabled, no need to disable
 *       - ESP_ERR_INVALID_ARG:     Invalid argument
 */
//go:linkname AdcContinuousMonitorDisable C.adc_continuous_monitor_disable
func AdcContinuousMonitorDisable(monitor_handle AdcMonitorHandleT) EspErrT

/**
 * @brief Free an ADC continuous mode monitor.
 *
 * @param[in]  monitor_handle       Monitor handle
 *
 * @return
 *       - ESP_OK:                  On success
 *       - ESP_ERR_INVALID_STATE:   Monitor is in enabled state, should call `adc_continuous_monitor_disable` first
 *       - ESP_ERR_NOT_FOUND:       Monitor haven't been used
 *       - ESP_ERR_INVALID_ARG:     Invalid argument
 */
//go:linkname AdcDelContinuousMonitor C.adc_del_continuous_monitor
func AdcDelContinuousMonitor(monitor_handle AdcMonitorHandleT) EspErrT
