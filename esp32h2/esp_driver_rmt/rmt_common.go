package esp_driver_rmt

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief RMT carrier wave configuration (for either modulation or demodulation)
 */

type RmtCarrierConfigT struct {
	FrequencyHz c.Uint32T
	DutyCycle   c.Float
	Flags       struct {
		PolarityActiveLow c.Uint32T
		AlwaysOn          c.Uint32T
	}
}

/**
 * @brief Delete an RMT channel
 *
 * @param[in] channel RMT generic channel that created by `rmt_new_tx_channel()` or `rmt_new_rx_channel()`
 * @return
 *      - ESP_OK: Delete RMT channel successfully
 *      - ESP_ERR_INVALID_ARG: Delete RMT channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Delete RMT channel failed because it is still in working
 *      - ESP_FAIL: Delete RMT channel failed because of other error
 */
//go:linkname RmtDelChannel C.rmt_del_channel
func RmtDelChannel(channel RmtChannelHandleT) EspErrT

/**
 * @brief Apply modulation feature for TX channel or demodulation feature for RX channel
 *
 * @param[in] channel RMT generic channel that created by `rmt_new_tx_channel()` or `rmt_new_rx_channel()`
 * @param[in] config Carrier configuration. Specially, a NULL config means to disable the carrier modulation or demodulation feature
 * @return
 *      - ESP_OK: Apply carrier configuration successfully
 *      - ESP_ERR_INVALID_ARG: Apply carrier configuration failed because of invalid argument
 *      - ESP_FAIL: Apply carrier configuration failed because of other error
 */
//go:linkname RmtApplyCarrier C.rmt_apply_carrier
func RmtApplyCarrier(channel RmtChannelHandleT, config *RmtCarrierConfigT) EspErrT

/**
 * @brief Enable the RMT channel
 *
 * @note This function will acquire a PM lock that might be installed during channel allocation
 *
 * @param[in] channel RMT generic channel that created by `rmt_new_tx_channel()` or `rmt_new_rx_channel()`
 * @return
 *      - ESP_OK: Enable RMT channel successfully
 *      - ESP_ERR_INVALID_ARG: Enable RMT channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Enable RMT channel failed because it's enabled already
 *      - ESP_FAIL: Enable RMT channel failed because of other error
 */
//go:linkname RmtEnable C.rmt_enable
func RmtEnable(channel RmtChannelHandleT) EspErrT

/**
 * @brief Disable the RMT channel
 *
 * @note This function will release a PM lock that might be installed during channel allocation
 *
 * @param[in] channel RMT generic channel that created by `rmt_new_tx_channel()` or `rmt_new_rx_channel()`
 * @return
 *      - ESP_OK: Disable RMT channel successfully
 *      - ESP_ERR_INVALID_ARG: Disable RMT channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Disable RMT channel failed because it's not enabled yet
 *      - ESP_FAIL: Disable RMT channel failed because of other error
 */
//go:linkname RmtDisable C.rmt_disable
func RmtDisable(channel RmtChannelHandleT) EspErrT
