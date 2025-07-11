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

//TODO IDF-3610, replace these with proper caps
/**
 * @brief Set ADC data invert
 * @param adc_unit ADC unit index
 * @param inv_en whether enable data invert
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname AdcSetDataInv C.adc_set_data_inv
func AdcSetDataInv(adc_unit AdcUnitT, inv_en bool) EspErrT

/**
 * @brief Set ADC source clock
 * @param clk_div ADC clock divider, ADC clock is divided from APB clock
 * @return
 *     - ESP_OK success
 */
//go:linkname AdcSetClkDiv C.adc_set_clk_div
func AdcSetClkDiv(clk_div c.Uint8T) EspErrT

/**
 * @brief Configure ADC capture width.
 *
 * @param adc_unit ADC unit index
 * @param width_bit Bit capture width for ADC unit.
 *
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname AdcSetDataWidth C.adc_set_data_width
func AdcSetDataWidth(adc_unit AdcUnitT, width_bit AdcBitsWidthT) EspErrT

/**
 * @brief Configure ADC1 to be usable by the ULP
 *
 * This function reconfigures ADC1 to be controlled by the ULP.
 * Effect of this function can be reverted using ``adc1_get_raw()`` function.
 *
 * Note that adc1_config_channel_atten, ``adc1_config_width()`` functions need
 * to be called to configure ADC1 channels, before ADC1 is used by the ULP.
 */
//go:linkname Adc1UlpEnable C.adc1_ulp_enable
func Adc1UlpEnable()

/**
 * @brief Get the GPIO number of a specific ADC2 channel.
 *
 * @param channel Channel to get the GPIO number
 *
 * @param gpio_num output buffer to hold the GPIO number
 *
 * @return
 *   - ESP_OK if success
 *   - ESP_ERR_INVALID_ARG if channel not valid
 */
// llgo:link Adc2ChannelT.Adc2PadGetIoNum C.adc2_pad_get_io_num
func (recv_ Adc2ChannelT) Adc2PadGetIoNum(gpio_num *GpioNumT) EspErrT {
	return 0
}

/**
 * @brief Configure the ADC2 channel, including setting attenuation.
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
 * @note This function also configures the input GPIO pin mux to
 *       connect it to the ADC2 channel. It must be called before calling
 *       ``adc2_get_raw()`` for this channel.
 *
 * @note For any given channel, this function must be called before the first time ``adc2_get_raw()`` is called for that channel.
 *
 * @param channel ADC2 channel to configure
 * @param atten  Attenuation level
 *
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link Adc2ChannelT.Adc2ConfigChannelAtten C.adc2_config_channel_atten
func (recv_ Adc2ChannelT) Adc2ConfigChannelAtten(atten AdcAttenT) EspErrT {
	return 0
}

/**
 * @brief Take an ADC2 reading on a single channel
 *
 * @note ESP32:
 *       When the power switch of SARADC1, SARADC2, HALL sensor and AMP sensor is turned on,
 *       the input of GPIO36 and GPIO39 will be pulled down for about 80ns.
 *       When enabling power for any of these peripherals, ignore input from GPIO36 and GPIO39.
 *       Please refer to section 3.11 of 'ECO_and_Workarounds_for_Bugs_in_ESP32' for the description of this issue.
 *       As a workaround, call sar_periph_ctrl_adc_oneshot_power_acquire() in the app. This will result in higher power consumption (by ~1mA),
 *       but will remove the glitches on GPIO36 and GPIO39.
 *
 *
 * @note ESP32:
 *       For a given channel, ``adc2_config_channel_atten()``
 *       must be called before the first time this function is called. If Wi-Fi is started via ``esp_wifi_start()``, this
 *       function will always fail with ``ESP_ERR_TIMEOUT``.
 *
 * @note ESP32-S2:
 *       ADC2 support hardware arbiter. The arbiter is to improve the use efficiency of ADC2. After the control right is robbed by the high priority,
 *       the low priority controller will read the invalid ADC2 data. Default priority: Wi-Fi > RTC > Digital;
 *
 * @param channel ADC2 channel to read
 * @param width_bit Bit capture width for ADC2
 * @param raw_out the variable to hold the output data.
 *
 * @return
 *     - ESP_OK if success
 *     - ESP_ERR_TIMEOUT ADC2 is being used by other controller and the request timed out.
 *     - ESP_ERR_INVALID_STATE The controller status is invalid. Please try again.
 */
// llgo:link Adc2ChannelT.Adc2GetRaw C.adc2_get_raw
func (recv_ Adc2ChannelT) Adc2GetRaw(width_bit AdcBitsWidthT, raw_out *c.Int) EspErrT {
	return 0
}

/**
 *  @brief Output ADC1 or ADC2's reference voltage to ``adc2_channe_t``'s IO.
 *
 *  This function routes the internal reference voltage of ADCn to one of
 *  ADC2's channels. This reference voltage can then be manually measured
 *  for calibration purposes.
 *
 *  @note  ESP32 only supports output of ADC2's internal reference voltage.
 *  @param[in]  adc_unit ADC unit index
 *  @param[in]  gpio     GPIO number (Only ADC2's channels IO are supported)
 *
 *  @return
 *                  - ESP_OK: v_ref successfully routed to selected GPIO
 *                  - ESP_ERR_INVALID_ARG: Unsupported GPIO
 */
//go:linkname AdcVrefToGpio C.adc_vref_to_gpio
func AdcVrefToGpio(adc_unit AdcUnitT, gpio GpioNumT) EspErrT

/**
 * @brief Initialize the Digital ADC.
 *
 * @param init_config Pointer to Digital ADC initialization config. Refer to ``adc_digi_init_config_t``.
 *
 * @return
 *         - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *         - ESP_ERR_NOT_FOUND     No free interrupt found with the specified flags
 *         - ESP_ERR_NO_MEM        If out of memory
 *         - ESP_OK                On success
 */
// llgo:link (*AdcDigiInitConfigT).AdcDigiInitialize C.adc_digi_initialize
func (recv_ *AdcDigiInitConfigT) AdcDigiInitialize() EspErrT {
	return 0
}

/**
 * @brief Read bytes from Digital ADC through DMA.
 *
 * @param[out] buf                 Buffer to read from ADC.
 * @param[in]  length_max          Expected length of data read from the ADC.
 * @param[out] out_length          Real length of data read from the ADC via this API.
 * @param[in]  timeout_ms          Time to wait for data via this API, in millisecond.
 *
 * @return
 *         - ESP_ERR_INVALID_STATE Driver state is invalid. Usually it means the ADC sampling rate is faster than the task processing rate.
 *         - ESP_ERR_TIMEOUT       Operation timed out
 *         - ESP_OK                On success
 */
//go:linkname AdcDigiReadBytes C.adc_digi_read_bytes
func AdcDigiReadBytes(buf *c.Uint8T, length_max c.Uint32T, out_length *c.Uint32T, timeout_ms c.Uint32T) EspErrT

/**
 * @brief Start the Digital ADC and DMA peripherals. After this, the hardware starts working.
 *
 * @return
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 *         - ESP_OK                On success
 */
//go:linkname AdcDigiStart C.adc_digi_start
func AdcDigiStart() EspErrT

/**
 * @brief Stop the Digital ADC and DMA peripherals. After this, the hardware stops working.
 *
 * @return
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 *         - ESP_OK                On success
 */
//go:linkname AdcDigiStop C.adc_digi_stop
func AdcDigiStop() EspErrT

/**
 * @brief Deinitialize the Digital ADC.
 *
 * @return
 *         - ESP_ERR_INVALID_STATE Driver state is invalid.
 *         - ESP_OK                On success
 */
//go:linkname AdcDigiDeinitialize C.adc_digi_deinitialize
func AdcDigiDeinitialize() EspErrT

/**
 * @brief Setting the digital controller.
 *
 * @param config Pointer to digital controller parameter. Refer to ``adc_digi_config_t``.
 *
 * @return
 *      - ESP_ERR_INVALID_STATE Driver state is invalid.
 *      - ESP_ERR_INVALID_ARG   If the combination of arguments is invalid.
 *      - ESP_OK                On success
 */
// llgo:link (*AdcDigiConfigurationT).AdcDigiControllerConfigure C.adc_digi_controller_configure
func (recv_ *AdcDigiConfigurationT) AdcDigiControllerConfigure() EspErrT {
	return 0
}
