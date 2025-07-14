package esp_adc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*---------------------------------------------------------------
            ADC IOs
---------------------------------------------------------------*/
/**
 * @brief Get ADC channel from the given GPIO number
 *
 * @param[in]  io_num     GPIO number
 * @param[out] unit_id    ADC unit
 * @param[out] channel    ADC channel
 *
 * @return
 *        - ESP_OK:              On success
 *        - ESP_ERR_INVALID_ARG: Invalid argument
 *        - ESP_ERR_NOT_FOUND:   The IO is not a valid ADC pad
 */
//go:linkname AdcIoToChannel C.adc_io_to_channel
func AdcIoToChannel(io_num c.Int, unit_id *AdcUnitT, channel *AdcChannelT) EspErrT

/**
 * @brief Get GPIO number from the given ADC channel
 *
 * @param[in]  unit_id    ADC unit
 * @param[in]  channel    ADC channel
 * @param[out] io_num     GPIO number
 *
 * @param
 *       - ESP_OK:              On success
 *       - ESP_ERR_INVALID_ARG: Invalid argument
 */
//go:linkname AdcChannelToIo C.adc_channel_to_io
func AdcChannelToIo(unit_id AdcUnitT, channel AdcChannelT, io_num *c.Int) EspErrT

type AdcOneshotUnitCtxT struct {
	Unused [8]uint8
}
type AdcOneshotUnitHandleT *AdcOneshotUnitCtxT

/**
 * @brief ISR version to get one ADC conversion raw result
 *
 * @note This API only provide atomic register settings, without hardware resources protection. When other drivers are using
 *       SAR-ADCs, calling this API may get wrong ADC result.
 * @note This API can be called in an ISR context.
 * @note Strongly suggest using this function when there's no concurrent hardware usage to the ADC. You can refer to ADC Oneshot
 *       Programming Guide to know ADC Hardware Limitations
 *
 * @param[in] handle    ADC handle
 * @param[in] chan      ADC channel
 * @param[out] out_raw  ADC conversion raw result
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid arguments
 *        - ESP_ERR_INVALID_STATE: Invalid state, the ADC result is invalid
 */
//go:linkname AdcOneshotReadIsr C.adc_oneshot_read_isr
func AdcOneshotReadIsr(handle AdcOneshotUnitHandleT, chan_ AdcChannelT, out_raw *c.Int) EspErrT
