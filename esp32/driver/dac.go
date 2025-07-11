package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get the GPIO number of a specific DAC channel.
 *
 * @param channel Channel to get the gpio number
 * @param gpio_num output buffer to hold the gpio number
 * @return
 *   - ESP_OK if success
 */
//go:linkname DacPadGetIoNum C.dac_pad_get_io_num
func DacPadGetIoNum(channel DacChannelT, gpio_num *GpioNumT) EspErrT

/**
 * @brief Set DAC output voltage.
 *        DAC output is 8-bit. Maximum (255) corresponds to VDD3P3_RTC.
 *
 * @note Need to configure DAC pad before calling this function.
 *       DAC channel 0 is attached to GPIO25, DAC channel 1 is attached to GPIO26
 * @param channel DAC channel
 * @param dac_value DAC output value
 *
 * @return
 *     - ESP_OK success
 */
//go:linkname DacOutputVoltage C.dac_output_voltage
func DacOutputVoltage(channel DacChannelT, dac_value c.Uint8T) EspErrT

/**
 * @brief DAC pad output enable
 *
 * @param channel DAC channel
 * @note DAC channel 0 is attached to GPIO25, DAC channel 1 is attached to GPIO26
 *       I2S left channel will be mapped to DAC channel 1
 *       I2S right channel will be mapped to DAC channel 0
 */
//go:linkname DacOutputEnable C.dac_output_enable
func DacOutputEnable(channel DacChannelT) EspErrT

/**
 * @brief DAC pad output disable
 *
 * @param channel DAC channel
 * @note DAC channel 0 is attached to GPIO25, DAC channel 1 is attached to GPIO26
 * @return
 *     - ESP_OK success
 */
//go:linkname DacOutputDisable C.dac_output_disable
func DacOutputDisable(channel DacChannelT) EspErrT

/**
 * @brief Enable cosine wave generator output.
 *
 * @return
 *     - ESP_OK success
 */
//go:linkname DacCwGeneratorEnable C.dac_cw_generator_enable
func DacCwGeneratorEnable() EspErrT

/**
 * @brief Disable cosine wave generator output.
 *
 * @return
 *     - ESP_OK success
 */
//go:linkname DacCwGeneratorDisable C.dac_cw_generator_disable
func DacCwGeneratorDisable() EspErrT

/**
 * @brief Config the cosine wave generator function in DAC module.
 *
 * @param cw Configuration.
 * @return
 *     - ESP_OK success
 *     - ESP_ERR_INVALID_ARG The parameter is NULL.
 */
// llgo:link (*DacCwConfigT).DacCwGeneratorConfig C.dac_cw_generator_config
func (recv_ *DacCwConfigT) DacCwGeneratorConfig() EspErrT {
	return 0
}

/**
 * @brief Enable DAC output data from I2S
 *
 * @return
 *     - ESP_OK success
 */
//go:linkname DacI2sEnable C.dac_i2s_enable
func DacI2sEnable() EspErrT

/**
 * @brief Disable DAC output data from I2S
 *
 * @return
 *     - ESP_OK success
 */
//go:linkname DacI2sDisable C.dac_i2s_disable
func DacI2sDisable() EspErrT
