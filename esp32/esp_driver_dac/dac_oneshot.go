package esp_driver_dac

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DacOneshotS struct {
	Unused [8]uint8
}
type DacOneshotHandleT *DacOneshotS

/**
 * @brief DAC oneshot channel configuration
 *
 */

type DacOneshotConfigT struct {
	ChanId DacChannelT
}

/**
 * @brief Allocate a new DAC oneshot channel
 * @note  The channel will be enabled as well when the channel allocated
 *
 * @param[in]  oneshot_cfg   The configuration for the oneshot channel
 * @param[out] ret_handle    The returned oneshot channel handle
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The DAC channel has been registered already
 *      - ESP_ERR_NO_MEM        No memory for the DAC oneshot channel resources
 *      - ESP_OK                Allocate the new DAC oneshot channel success
 */
// llgo:link (*DacOneshotConfigT).DacOneshotNewChannel C.dac_oneshot_new_channel
func (recv_ *DacOneshotConfigT) DacOneshotNewChannel(ret_handle *DacOneshotHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete the DAC oneshot channel
 * @note  The channel will be disabled as well when the channel deleted
 *
 * @param[in]  handle       The DAC oneshot channel handle
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The channel has already been de-registered
 *      - ESP_OK                Delete the oneshot channel success
 */
//go:linkname DacOneshotDelChannel C.dac_oneshot_del_channel
func DacOneshotDelChannel(handle DacOneshotHandleT) EspErrT

/**
 * @brief Output the voltage
 * @note  Generally it'll take 7~11 us on ESP32 and 10~21 us on ESP32-S2
 *
 * @param[in]  handle       The DAC oneshot channel handle
 * @param[in]  digi_value   The digital value that need to be converted
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_OK                Convert the digital value success
 */
//go:linkname DacOneshotOutputVoltage C.dac_oneshot_output_voltage
func DacOneshotOutputVoltage(handle DacOneshotHandleT, digi_value c.Uint8T) EspErrT
