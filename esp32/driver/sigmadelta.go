package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Configure Sigma-delta channel
 *
 * @param  config Pointer of Sigma-delta channel configuration struct
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE sigmadelta driver already initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link (*SigmadeltaConfigT).SigmadeltaConfig C.sigmadelta_config
func (recv_ *SigmadeltaConfigT) SigmadeltaConfig() EspErrT {
	return 0
}

/**
 * @brief Set Sigma-delta channel duty.
 *
 *        This function is used to set Sigma-delta channel duty,
 *        If you add a capacitor between the output pin and ground,
 *        the average output voltage will be Vdc = VDDIO / 256 * duty + VDDIO/2,
 *        where VDDIO is the power supply voltage.
 *
 * @param channel Sigma-delta channel number
 * @param duty Sigma-delta duty of one channel, the value ranges from -128 to 127, recommended range is -90 ~ 90.
 *             The waveform is more like a random one in this range.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE sigmadelta driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link SigmadeltaChannelT.SigmadeltaSetDuty C.sigmadelta_set_duty
func (recv_ SigmadeltaChannelT) SigmadeltaSetDuty(duty c.Int8T) EspErrT {
	return 0
}

/**
 * @brief Set Sigma-delta channel's clock pre-scale value.
 *        The source clock is APP_CLK, 80MHz. The clock frequency of the sigma-delta channel is APP_CLK / pre_scale
 *
 * @param channel Sigma-delta channel number
 * @param prescale The divider of source clock, ranges from 0 to 255
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE sigmadelta driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link SigmadeltaChannelT.SigmadeltaSetPrescale C.sigmadelta_set_prescale
func (recv_ SigmadeltaChannelT) SigmadeltaSetPrescale(prescale c.Uint8T) EspErrT {
	return 0
}

/**
 * @brief Set Sigma-delta signal output pin
 *
 * @param channel Sigma-delta channel number
 * @param gpio_num GPIO number of output pin.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE sigmadelta driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link SigmadeltaChannelT.SigmadeltaSetPin C.sigmadelta_set_pin
func (recv_ SigmadeltaChannelT) SigmadeltaSetPin(gpio_num GpioNumT) EspErrT {
	return 0
}
