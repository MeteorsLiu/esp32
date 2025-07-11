package esp_adc

import _ "unsafe"

type AdcIirFilterT struct {
	Unused [8]uint8
}
type AdcIirFilterHandleT *AdcIirFilterT

/**
 * @brief Filter Configuration
 */

type AdcContinuousIirFilterConfigT struct {
	Unit    AdcUnitT
	Channel AdcChannelT
	Coeff   AdcDigiIirFilterCoeffT
}
