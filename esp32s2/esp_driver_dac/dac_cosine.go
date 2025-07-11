package esp_driver_dac

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DacCosineS struct {
	Unused [8]uint8
}
type DacCosineHandleT *DacCosineS

/**
 * @brief DAC cosine channel configurations
 *
 */

type DacCosineConfigT struct {
	ChanId DacChannelT
	FreqHz c.Uint32T
	ClkSrc DacCosineClkSrcT
	Atten  DacCosineAttenT
	Phase  DacCosinePhaseT
	Offset c.Int8T
	Flags  struct {
		ForceSetFreq bool
	}
}

/**
 * @brief Allocate a new DAC cosine wave channel
 * @note  Since there is only one cosine wave generator,
 *        only the first channel can set the frequency of the cosine wave.
 *        Normally, the latter one is not allowed to set a different frequency,
 *        but the it can be forced to set by setting the bit `force_set_freq` in the configuration,
 *        notice that another channel will be affected as well when the frequency is updated.
 *
 * @param[in]  cos_cfg      The configuration of cosine wave channel
 * @param[out] ret_handle   The returned cosine wave channel handle
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The DAC channel has been registered already
 *      - ESP_ERR_NO_MEM        No memory for the DAC cosine wave channel resources
 *      - ESP_OK                Allocate the new DAC cosine wave channel success
 */
// llgo:link (*DacCosineConfigT).DacCosineNewChannel C.dac_cosine_new_channel
func (recv_ *DacCosineConfigT) DacCosineNewChannel(ret_handle *DacCosineHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete the DAC cosine wave channel
 *
 * @param[in] handle        The DAC cosine wave channel handle
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The channel has already been deregistered
 *      - ESP_OK                Delete the cosine wave channel success
 */
//go:linkname DacCosineDelChannel C.dac_cosine_del_channel
func DacCosineDelChannel(handle DacCosineHandleT) EspErrT

/**
 * @brief Start outputting the cosine wave on the channel
 *
 * @param[in] handle            The DAC cosine wave channel handle
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The channel has been started already
 *      - ESP_OK                Start the cosine wave success
 */
//go:linkname DacCosineStart C.dac_cosine_start
func DacCosineStart(handle DacCosineHandleT) EspErrT

/**
 * @brief Stop outputting the cosine wave on the channel
 *
 * @param[in] handle            The DAC cosine wave channel handle
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The channel has been stopped already
 *      - ESP_OK                Stop the cosine wave success
 */
//go:linkname DacCosineStop C.dac_cosine_stop
func DacCosineStop(handle DacCosineHandleT) EspErrT
