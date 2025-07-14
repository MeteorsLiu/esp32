package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*---------------------------------------------------------------
            Deprecated API
---------------------------------------------------------------*/
/*---------------------------------------------------------------
                    ADC Single Read Setting
---------------------------------------------------------------*/
/**
 * @brief Get the GPIO number of a specific ADC1 channel.
 *
 * @param channel Channel to get the GPIO number
 * @param gpio_num output buffer to hold the GPIO number
 *
 * @return
 *   - ESP_OK if success
 *   - ESP_ERR_INVALID_ARG if channel not valid
 */
// llgo:link Adc1ChannelT.Adc1PadGetIoNum C.adc1_pad_get_io_num
func (recv_ Adc1ChannelT) Adc1PadGetIoNum(gpio_num *GpioNumT) EspErrT {
	return 0
}

/**
 * @brief Set the attenuation of a particular channel on ADC1, and configure its associated GPIO pin mux.
 *
 * The default ADC voltage is for attenuation 0 dB and listed in the table below.
 * By setting higher attenuation it is possible to read higher voltages.
 *
 * Due to ADC characteristics, most accurate results are obtained within the "suggested range"
 * shown in the following table.
 *
 *     +----------+-------------+-----------------+
 *     |          | attenuation | suggested range |
 *     |    SoC   |     (dB)    |      (mV)       |
 *     +==========+=============+=================+
 *     |          |       0     |    100 ~  950   |
 *     |          +-------------+-----------------+
 *     |          |       2.5   |    100 ~ 1250   |
 *     |   ESP32  +-------------+-----------------+
 *     |          |       6     |    150 ~ 1750   |
 *     |          +-------------+-----------------+
 *     |          |      11     |    150 ~ 2450   |
 *     +----------+-------------+-----------------+
 *     |          |       0     |      0 ~  750   |
 *     |          +-------------+-----------------+
 *     |          |       2.5   |      0 ~ 1050   |
 *     | ESP32-S2 +-------------+-----------------+
 *     |          |       6     |      0 ~ 1300   |
 *     |          +-------------+-----------------+
 *     |          |      11     |      0 ~ 2500   |
 *     +----------+-------------+-----------------+
 *
 * For maximum accuracy, use the ADC calibration APIs and measure voltages within these recommended ranges.
 *
 * @note For any given channel, this function must be called before the first time ``adc1_get_raw()`` is called for that channel.
 *
 * @note This function can be called multiple times to configure multiple
 *       ADC channels simultaneously. You may call ``adc1_get_raw()`` only after configuring a channel.
 *
 * @param channel ADC1 channel to configure
 * @param atten  Attenuation level
 *
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link Adc1ChannelT.Adc1ConfigChannelAtten C.adc1_config_channel_atten
func (recv_ Adc1ChannelT) Adc1ConfigChannelAtten(atten AdcAttenT) EspErrT {
	return 0
}

/**
 * @brief Configure ADC1 capture width, meanwhile enable output invert for ADC1.
 *        The configuration is for all channels of ADC1
 * @param width_bit Bit capture width for ADC1
 *
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link AdcBitsWidthT.Adc1ConfigWidth C.adc1_config_width
func (recv_ AdcBitsWidthT) Adc1ConfigWidth() EspErrT {
	return 0
}

/**
 * @brief Take an ADC1 reading from a single channel.
 * @note ESP32:
 *       When the power switch of SARADC1, SARADC2, HALL sensor and AMP sensor is turned on,
 *       the input of GPIO36 and GPIO39 will be pulled down for about 80ns.
 *       When enabling power for any of these peripherals, ignore input from GPIO36 and GPIO39.
 *       Please refer to section 3.11 of 'ECO_and_Workarounds_for_Bugs_in_ESP32' for the description of this issue.
 *       As a workaround, call sar_periph_ctrl_adc_oneshot_power_acquire() in the app. This will result in higher power consumption (by ~1mA),
 *       but will remove the glitches on GPIO36 and GPIO39.
 *
 * @note Call ``adc1_config_width()`` before the first time this
 *       function is called.
 *
 * @note For any given channel, adc1_config_channel_atten(channel)
 *       must be called before the first time this function is called. Configuring
 *       a new channel does not prevent a previously configured channel from being read.
 *
 * @param  channel ADC1 channel to read
 *
 * @return
 *     - -1: Parameter error
 *     -  Other: ADC1 channel reading.
 */
// llgo:link Adc1ChannelT.Adc1GetRaw C.adc1_get_raw
func (recv_ Adc1ChannelT) Adc1GetRaw() c.Int {
	return 0
}
