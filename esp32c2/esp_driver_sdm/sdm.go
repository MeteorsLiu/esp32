package esp_driver_sdm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SdmChannelT struct {
	Unused [8]uint8
}
type SdmChannelHandleT *SdmChannelT

/**
 * @brief Sigma Delta channel configuration
 */

type SdmConfigT struct {
	GpioNum      c.Int
	ClkSrc       SdmClockSourceT
	SampleRateHz c.Uint32T
	Flags        struct {
		InvertOut  c.Uint32T
		IoLoopBack c.Uint32T
	}
}

/**
 * @brief Create a new Sigma Delta channel
 *
 * @param[in] config SDM configuration
 * @param[out] ret_chan Returned SDM channel handle
 * @return
 *      - ESP_OK: Create SDM channel successfully
 *      - ESP_ERR_INVALID_ARG: Create SDM channel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create SDM channel failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create SDM channel failed because all channels are used up and no more free one
 *      - ESP_FAIL: Create SDM channel failed because of other error
 */
// llgo:link (*SdmConfigT).SdmNewChannel C.sdm_new_channel
func (recv_ *SdmConfigT) SdmNewChannel(ret_chan *SdmChannelHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete the Sigma Delta channel
 *
 * @param[in] chan SDM channel created by `sdm_new_channel`
 * @return
 *      - ESP_OK: Delete the SDM channel successfully
 *      - ESP_ERR_INVALID_ARG: Delete the SDM channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Delete the SDM channel failed because the channel is not in init state
 *      - ESP_FAIL: Delete the SDM channel failed because of other error
 */
//go:linkname SdmDelChannel C.sdm_del_channel
func SdmDelChannel(chan_ SdmChannelHandleT) EspErrT

/**
 * @brief Enable the Sigma Delta channel
 *
 * @note This function will transit the channel state from init to enable.
 * @note This function will acquire a PM lock, if a specific source clock (e.g. APB) is selected in the `sdm_config_t`, while `CONFIG_PM_ENABLE` is enabled.
 *
 * @param[in] chan SDM channel created by `sdm_new_channel`
 * @return
 *      - ESP_OK: Enable SDM channel successfully
 *      - ESP_ERR_INVALID_ARG: Enable SDM channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Enable SDM channel failed because the channel is already enabled
 *      - ESP_FAIL: Enable SDM channel failed because of other error
 */
//go:linkname SdmChannelEnable C.sdm_channel_enable
func SdmChannelEnable(chan_ SdmChannelHandleT) EspErrT

/**
 * @brief Disable the Sigma Delta channel
 *
 * @note This function will do the opposite work to the `sdm_channel_enable()`
 *
 * @param[in] chan SDM channel created by `sdm_new_channel`
 * @return
 *      - ESP_OK: Disable SDM channel successfully
 *      - ESP_ERR_INVALID_ARG: Disable SDM channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Disable SDM channel failed because the channel is not enabled yet
 *      - ESP_FAIL: Disable SDM channel failed because of other error
 */
//go:linkname SdmChannelDisable C.sdm_channel_disable
func SdmChannelDisable(chan_ SdmChannelHandleT) EspErrT

/**
 * @brief Set the pulse density of the PDM output signal.
 *
 * @note The raw output signal requires a low-pass filter to restore it into analog voltage,
*        the restored analog output voltage could be Vout = VDD_IO / 256 * density + VDD_IO / 2
 * @note This function is allowed to run within ISR context
 * @note This function will be placed into IRAM if `CONFIG_SDM_CTRL_FUNC_IN_IRAM` is on,
 *       so that it's allowed to be executed when Cache is disabled
 *
 * @param[in] chan SDM channel created by `sdm_new_channel`
 * @param[in] density Quantized pulse density of the PDM output signal, ranges from -128 to 127.
 *                    But the range of [-90, 90] can provide a better randomness.
 * @return
 *      - ESP_OK: Set pulse density successfully
 *      - ESP_ERR_INVALID_ARG: Set pulse density failed because of invalid argument
 *      - ESP_FAIL: Set pulse density failed because of other error
*/
//go:linkname SdmChannelSetPulseDensity C.sdm_channel_set_pulse_density
func SdmChannelSetPulseDensity(chan_ SdmChannelHandleT, density c.Int8T) EspErrT

/**
 * @brief The alias function of `sdm_channel_set_pulse_density`, it decides the pulse density of the output signal
 *
 * @note  `sdm_channel_set_pulse_density` has a more appropriate name compare this
 *        alias function, suggest to turn to `sdm_channel_set_pulse_density` instead
 *
 * @param[in] chan SDM channel created by `sdm_new_channel`
 * @param[in] duty Actually it's the quantized pulse density of the PDM output signal
 *
 * @return
 *      - ESP_OK: Set duty cycle successfully
 *      - ESP_ERR_INVALID_ARG: Set duty cycle failed because of invalid argument
 *      - ESP_FAIL: Set duty cycle failed because of other error
 */
//go:linkname SdmChannelSetDuty C.sdm_channel_set_duty
func SdmChannelSetDuty(chan_ SdmChannelHandleT, duty c.Int8T) EspErrT
