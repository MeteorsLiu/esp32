package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief LP I2S controller channel configuration
 */
type LpI2sChanConfigT struct {
	Id        c.Int
	Role      I2sRoleT
	Threshold c.SizeT
}

/**
 * @brief LP I2S event callbacks
 */

type LpI2sEvtCbsT struct {
	OnThreshMet       LpI2sCallbackT
	OnRequestNewTrans LpI2sCallbackT
}

/**
 * @brief Allocate new LP I2S channel(s)
 *
 * @param[in]   chan_cfg        LP I2S controller channel configurations
 * @param[out]  ret_tx_handle   LP I2S channel handler used for managing the sending channel(optional), this one is not supported and is kept here for future-proof.
 * @param[out]  ret_rx_handle   LP I2S channel handler used for managing the receiving channel(optional)
 * @return
 *      - ESP_OK                 Allocate new channel(s) success
 *      - ESP_ERR_NOT_SUPPORTED  The communication mode is not supported on the current chip
 *      - ESP_ERR_INVALID_ARG    NULL pointer or illegal parameter in lp_i2s_chan_config_t
 *      - ESP_ERR_NOT_FOUND      No available LP I2S channel found
 */
// llgo:link (*LpI2sChanConfigT).LpI2sNewChannel C.lp_i2s_new_channel
func (recv_ *LpI2sChanConfigT) LpI2sNewChannel(ret_tx_handle *LpI2sChanHandleT, ret_rx_handle *LpI2sChanHandleT) EspErrT {
	return 0
}

/**
 * @brief Register LP I2S event callbacks
 *
 * @param[in] chan       LP I2S channel handle
 * @param[in] cbs        Callbacks
 * @param[in] user_data  User data
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname LpI2sRegisterEventCallbacks C.lp_i2s_register_event_callbacks
func LpI2sRegisterEventCallbacks(chan_ LpI2sChanHandleT, cbs *LpI2sEvtCbsT, user_data c.Pointer) EspErrT

/**
 * @brief Enable LP I2S driver
 *
 * @param[in] chan     LP I2S channel handle
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname LpI2sChannelEnable C.lp_i2s_channel_enable
func LpI2sChannelEnable(chan_ LpI2sChanHandleT) EspErrT

/**
 * @brief Read LP I2S received data
 *
 * @param[in] chan        LP I2S channel handle
 * @param[in] trans       LP I2S transaction
 * @param[in] timeout_ms  Timeout in ms, set to `LP_I2S_MAX_DELAY` to wait until read is done
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state, e.g. `on_request_new_trans` callback is registered
 */
//go:linkname LpI2sChannelRead C.lp_i2s_channel_read
func LpI2sChannelRead(chan_ LpI2sChanHandleT, trans *LpI2sTransT, timeout_ms c.Uint32T) EspErrT

/**
 * @brief Read LP I2S received data until certain bytes
 *
 * @param[in] chan        LP I2S channel handle
 * @param[in] trans       LP I2S transaction
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state, e.g. `on_request_new_trans` callback is registered
 */
//go:linkname LpI2sChannelReadUntilBytes C.lp_i2s_channel_read_until_bytes
func LpI2sChannelReadUntilBytes(chan_ LpI2sChanHandleT, trans *LpI2sTransT) EspErrT

/**
 * @brief Disable LP I2S driver
 *
 * @param[in] chan       LP I2S channel handle
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname LpI2sChannelDisable C.lp_i2s_channel_disable
func LpI2sChannelDisable(chan_ LpI2sChanHandleT) EspErrT

/**
 * @brief Delete the LP I2S channel
 *
 * @param[in]   chan        LP I2S channel handler
 *
 * @return
 *      - ESP_OK                Delete successfully
 *      - ESP_ERR_INVALID_ARG   NULL pointer
 */
//go:linkname LpI2sDelChannel C.lp_i2s_del_channel
func LpI2sDelChannel(chan_ LpI2sChanHandleT) EspErrT
