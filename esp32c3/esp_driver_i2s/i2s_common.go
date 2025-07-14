package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Group of I2S callbacks
 * @note The callbacks are all running under ISR environment
 * @note When CONFIG_I2S_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well.
 */

type I2sEventCallbacksT struct {
	OnRecv     I2sIsrCallbackT
	OnRecvQOvf I2sIsrCallbackT
	OnSent     I2sIsrCallbackT
	OnSendQOvf I2sIsrCallbackT
}

/**
 * @brief I2S controller channel configuration
 */
type I2sChanConfigT struct {
	Id                I2sPortT
	Role              I2sRoleT
	DmaDescNum        c.Uint32T
	DmaFrameNum       c.Uint32T
	AutoClearBeforeCb bool
	AllowPd           bool
	IntrPriority      c.Int
}

/**
 * @brief I2S channel information
 */

type I2sChanInfoT struct {
	Id              I2sPortT
	Role            I2sRoleT
	Dir             I2sDirT
	Mode            I2sCommModeT
	PairChan        I2sChanHandleT
	TotalDmaBufSize c.Uint32T
}

/**
 * @brief Allocate new I2S channel(s)
 * @note  The new created I2S channel handle will be REGISTERED state after it is allocated successfully.
 * @note  When the port id in channel configuration is I2S_NUM_AUTO, driver will allocate I2S port automatically
 *        on one of the I2S controller, otherwise driver will try to allocate the new channel on the selected port.
 * @note  If both tx_handle and rx_handle are not NULL, it means this I2S controller will work at full-duplex mode,
 *        the RX and TX channels will be allocated on a same I2S port in this case.
 *        Note that some configurations of TX/RX channel are shared on ESP32 and ESP32S2,
 *        so please make sure they are working at same condition and under same status(start/stop).
 *        Currently, full-duplex mode can't guarantee TX/RX channels write/read synchronously,
 *        they can only share the clock signals for now.
 * @note  If tx_handle OR rx_handle is NULL, it means this I2S controller will work at simplex mode.
 *        For ESP32 and ESP32S2, the whole I2S controller (i.e. both RX and TX channel) will be occupied,
 *        even if only one of RX or TX channel is registered.
 *        For the other targets, another channel on this controller will still available.
 *
 * @param[in]   chan_cfg    I2S controller channel configurations
 * @param[out]  ret_tx_handle   I2S channel handler used for managing the sending channel(optional)
 * @param[out]  ret_rx_handle   I2S channel handler used for managing the receiving channel(optional)
 * @return
 *      - ESP_OK    Allocate new channel(s) success
 *      - ESP_ERR_NOT_SUPPORTED The communication mode is not supported on the current chip
 *      - ESP_ERR_INVALID_ARG   NULL pointer or illegal parameter in i2s_chan_config_t
 *      - ESP_ERR_NOT_FOUND     No available I2S channel found
 */
// llgo:link (*I2sChanConfigT).I2sNewChannel C.i2s_new_channel
func (recv_ *I2sChanConfigT) I2sNewChannel(ret_tx_handle *I2sChanHandleT, ret_rx_handle *I2sChanHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete the I2S channel
 * @note  Only allowed to be called when the I2S channel is at REGISTERED or READY state (i.e., it should stop before deleting it).
 * @note  Resource will be free automatically if all channels in one port are deleted
 *
 * @param[in]   handle      I2S channel handler
 *      - ESP_OK    Delete successfully
 *      - ESP_ERR_INVALID_ARG   NULL pointer
 */
//go:linkname I2sDelChannel C.i2s_del_channel
func I2sDelChannel(handle I2sChanHandleT) EspErrT

/**
 * @brief Get I2S channel information
 *
 * @param[in]   handle      I2S channel handler
 * @param[out]  chan_info   I2S channel basic information
 * @return
 *      - ESP_OK    Get I2S channel information success
 *      - ESP_ERR_NOT_FOUND     The input handle doesn't match any registered I2S channels, it may not an I2S channel handle or not available any more
 *      - ESP_ERR_INVALID_ARG   The input handle or chan_info pointer is NULL
 */
//go:linkname I2sChannelGetInfo C.i2s_channel_get_info
func I2sChannelGetInfo(handle I2sChanHandleT, chan_info *I2sChanInfoT) EspErrT

/**
 * @brief Enable the I2S channel
 * @note  Only allowed to be called when the channel state is READY, (i.e., channel has been initialized, but not started)
 *        the channel will enter RUNNING state once it is enabled successfully.
 * @note  Enable the channel can start the I2S communication on hardware. It will start outputting BCLK and WS signal.
 *        For MCLK signal, it will start to output when initialization is finished
 *
 * @param[in]   handle      I2S channel handler
 *      - ESP_OK    Start successfully
 *      - ESP_ERR_INVALID_ARG   NULL pointer
 *      - ESP_ERR_INVALID_STATE This channel has not initialized or already started
 */
//go:linkname I2sChannelEnable C.i2s_channel_enable
func I2sChannelEnable(handle I2sChanHandleT) EspErrT

/**
 * @brief Disable the I2S channel
 * @note  Only allowed to be called when the channel state is RUNNING, (i.e., channel has been started)
 *        the channel will enter READY state once it is disabled successfully.
 * @note  Disable the channel can stop the I2S communication on hardware. It will stop BCLK and WS signal but not MCLK signal
 *
 * @param[in]   handle      I2S channel handler
 * @return
 *      - ESP_OK    Stop successfully
 *      - ESP_ERR_INVALID_ARG   NULL pointer
 *      - ESP_ERR_INVALID_STATE This channel has not stated
 */
//go:linkname I2sChannelDisable C.i2s_channel_disable
func I2sChannelDisable(handle I2sChanHandleT) EspErrT

/**
 * @brief Preload the data into TX DMA buffer
 * @note  Only allowed to be called when the channel state is READY, (i.e., channel has been initialized, but not started)
 * @note  As the initial DMA buffer has no data inside, it will transmit the empty buffer after enabled the channel,
 *        this function is used to preload the data into the DMA buffer, so that the valid data can be transmitted immediately
 *        after the channel is enabled.
 * @note  This function can be called multiple times before enabling the channel, the buffer that loaded later will be concatenated
 *        behind the former loaded buffer. But when all the DMA buffers have been loaded, no more data can be preload then, please
 *        check the `bytes_loaded` parameter to see how many bytes are loaded successfully, when the `bytes_loaded` is smaller than
 *        the `size`, it means the DMA buffers are full.
 *
 * @param[in]   tx_handle   I2S TX channel handler
 * @param[in]   src         The pointer of the source buffer to be loaded
 * @param[in]   size        The source buffer size
 * @param[out]  bytes_loaded    The bytes that successfully been loaded into the TX DMA buffer
 * @return
 *      - ESP_OK    Load data successful
 *      - ESP_FAIL  Failed to push the message queue
 *      - ESP_ERR_INVALID_ARG   NULL pointer or not TX direction
 *      - ESP_ERR_INVALID_STATE This channel has not stated
 */
//go:linkname I2sChannelPreloadData C.i2s_channel_preload_data
func I2sChannelPreloadData(tx_handle I2sChanHandleT, src c.Pointer, size c.SizeT, bytes_loaded *c.SizeT) EspErrT

/**
 * @brief I2S write data
 * @note  Only allowed to be called when the channel state is RUNNING, (i.e., TX channel has been started and is not writing now)
 *        but the RUNNING only stands for the software state, it doesn't mean there is no the signal transporting on line.
 *
 * @param[in]   handle      I2S channel handler
 * @param[in]   src         The pointer of sent data buffer
 * @param[in]   size        Max data buffer length
 * @param[out]  bytes_written   Byte number that actually be sent, can be NULL if not needed
 * @param[in]   timeout_ms      Max block time
 * @return
 *      - ESP_OK    Write successfully
 *      - ESP_ERR_INVALID_ARG   NULL pointer or this handle is not TX handle
 *      - ESP_ERR_TIMEOUT       Writing timeout, no writing event received from ISR within ticks_to_wait
 *      - ESP_ERR_INVALID_STATE I2S is not ready to write
 */
//go:linkname I2sChannelWrite C.i2s_channel_write
func I2sChannelWrite(handle I2sChanHandleT, src c.Pointer, size c.SizeT, bytes_written *c.SizeT, timeout_ms c.Uint32T) EspErrT

/**
 * @brief I2S read data
 * @note  Only allowed to be called when the channel state is RUNNING
 *        but the RUNNING only stands for the software state, it doesn't mean there is no the signal transporting on line.
 *
 * @param[in]   handle      I2S channel handler
 * @param[in]   dest        The pointer of receiving data buffer
 * @param[in]   size        Max data buffer length
 * @param[out]  bytes_read      Byte number that actually be read, can be NULL if not needed
 * @param[in]   timeout_ms      Max block time
 * @return
 *      - ESP_OK    Read successfully
 *      - ESP_ERR_INVALID_ARG   NULL pointer or this handle is not RX handle
 *      - ESP_ERR_TIMEOUT       Reading timeout, no reading event received from ISR within ticks_to_wait
 *      - ESP_ERR_INVALID_STATE I2S is not ready to read
 */
//go:linkname I2sChannelRead C.i2s_channel_read
func I2sChannelRead(handle I2sChanHandleT, dest c.Pointer, size c.SizeT, bytes_read *c.SizeT, timeout_ms c.Uint32T) EspErrT

/**
 * @brief Set event callbacks for I2S channel
 *
 * @note Only allowed to be called when the channel state is REGISTERED / READY, (i.e., before channel starts)
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `callbacks` structure to NULL.
 * @note When CONFIG_I2S_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well. The `user_data` should also reside in SRAM or internal RAM as well.
 *
 * @param[in] handle        I2S channel handler
 * @param[in] callbacks     Group of callback functions
 * @param[in] user_data     User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK                Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG   Set event callbacks failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE Set event callbacks failed because the current channel state is not REGISTERED or READY
 */
//go:linkname I2sChannelRegisterEventCallback C.i2s_channel_register_event_callback
func I2sChannelRegisterEventCallback(handle I2sChanHandleT, callbacks *I2sEventCallbacksT, user_data c.Pointer) EspErrT
