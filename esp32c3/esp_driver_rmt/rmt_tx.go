package esp_driver_rmt

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Group of RMT TX callbacks
 * @note The callbacks are all running under ISR environment
 * @note When CONFIG_RMT_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well.
 */

type RmtTxEventCallbacksT struct {
	OnTransDone RmtTxDoneCallbackT
}

/**
 * @brief RMT TX channel specific configuration
 */

type RmtTxChannelConfigT struct {
	GpioNum         GpioNumT
	ClkSrc          RmtClockSourceT
	ResolutionHz    c.Uint32T
	MemBlockSymbols c.SizeT
	TransQueueDepth c.SizeT
	IntrPriority    c.Int
	Flags           struct {
		InvertOut  c.Uint32T
		WithDma    c.Uint32T
		IoLoopBack c.Uint32T
		IoOdMode   c.Uint32T
		AllowPd    c.Uint32T
	}
}

/**
 * @brief RMT transmit specific configuration
 */

type RmtTransmitConfigT struct {
	LoopCount c.Int
	Flags     struct {
		EotLevel         c.Uint32T
		QueueNonblocking c.Uint32T
	}
}

/**
 * @brief Synchronous manager configuration
 */

type RmtSyncManagerConfigT struct {
	TxChannelArray *RmtChannelHandleT
	ArraySize      c.SizeT
}

/**
 * @brief Create a RMT TX channel
 *
 * @param[in] config TX channel configurations
 * @param[out] ret_chan Returned generic RMT channel handle
 * @return
 *      - ESP_OK: Create RMT TX channel successfully
 *      - ESP_ERR_INVALID_ARG: Create RMT TX channel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create RMT TX channel failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create RMT TX channel failed because all RMT channels are used up and no more free one
 *      - ESP_ERR_NOT_SUPPORTED: Create RMT TX channel failed because some feature is not supported by hardware, e.g. DMA feature is not supported by hardware
 *      - ESP_FAIL: Create RMT TX channel failed because of other error
 */
// llgo:link (*RmtTxChannelConfigT).RmtNewTxChannel C.rmt_new_tx_channel
func (recv_ *RmtTxChannelConfigT) RmtNewTxChannel(ret_chan *RmtChannelHandleT) EspErrT {
	return 0
}

/**
 * @brief Transmit data by RMT TX channel
 *
 * @note This function constructs a transaction descriptor then pushes to a queue.
 *       The transaction will not start immediately if there's another one under processing.
 *       Based on the setting of `rmt_transmit_config_t::queue_nonblocking`,
 *       if there're too many transactions pending in the queue, this function can block until it has free slot,
 *       otherwise just return quickly.
 * @note The payload data to be transmitted will be encoded into RMT symbols by the specific `encoder`.
 * @note You CAN'T modify the `payload` during the transmission, it should be kept valid until the transmission is finished.
 *
 * @param[in] tx_channel RMT TX channel that created by `rmt_new_tx_channel()`
 * @param[in] encoder RMT encoder that created by various factory APIs like `rmt_new_bytes_encoder()`
 * @param[in] payload The raw data to be encoded into RMT symbols
 * @param[in] payload_bytes Size of the `payload` in bytes
 * @param[in] config Transmission specific configuration
 * @return
 *      - ESP_OK: Transmit data successfully
 *      - ESP_ERR_INVALID_ARG: Transmit data failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Transmit data failed because channel is not enabled
 *      - ESP_ERR_NOT_SUPPORTED: Transmit data failed because some feature is not supported by hardware, e.g. unsupported loop count
 *      - ESP_FAIL: Transmit data failed because of other error
 */
//go:linkname RmtTransmit C.rmt_transmit
func RmtTransmit(tx_channel RmtChannelHandleT, encoder RmtEncoderHandleT, payload c.Pointer, payload_bytes c.SizeT, config *RmtTransmitConfigT) EspErrT

/**
 * @brief Wait for all pending TX transactions done
 *
 * @note This function will block forever if the pending transaction can't be finished within a limited time (e.g. an infinite loop transaction).
 *       See also `rmt_disable()` for how to terminate a working channel.
 *
 * @param[in] tx_channel RMT TX channel that created by `rmt_new_tx_channel()`
 * @param[in] timeout_ms Wait timeout, in ms. Specially, -1 means to wait forever.
 * @return
 *      - ESP_OK: Flush transactions successfully
 *      - ESP_ERR_INVALID_ARG: Flush transactions failed because of invalid argument
 *      - ESP_ERR_TIMEOUT: Flush transactions failed because of timeout
 *      - ESP_FAIL: Flush transactions failed because of other error
 */
//go:linkname RmtTxWaitAllDone C.rmt_tx_wait_all_done
func RmtTxWaitAllDone(tx_channel RmtChannelHandleT, timeout_ms c.Int) EspErrT

/**
 * @brief Set event callbacks for RMT TX channel
 *
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 * @note When CONFIG_RMT_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well. The `user_data` should also reside in SRAM.
 *
 * @param[in] tx_channel RMT generic channel that created by `rmt_new_tx_channel()`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname RmtTxRegisterEventCallbacks C.rmt_tx_register_event_callbacks
func RmtTxRegisterEventCallbacks(tx_channel RmtChannelHandleT, cbs *RmtTxEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Create a synchronization manager for multiple TX channels, so that the managed channel can start transmitting at the same time
 *
 * @note All the channels to be managed should be enabled by `rmt_enable()` before put them into sync manager.
 *
 * @param[in] config Synchronization manager configuration
 * @param[out] ret_synchro Returned synchronization manager handle
 * @return
 *      - ESP_OK: Create sync manager successfully
 *      - ESP_ERR_INVALID_ARG: Create sync manager failed because of invalid argument
 *      - ESP_ERR_NOT_SUPPORTED: Create sync manager failed because it is not supported by hardware
 *      - ESP_ERR_INVALID_STATE: Create sync manager failed because not all channels are enabled
 *      - ESP_ERR_NO_MEM: Create sync manager failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create sync manager failed because all sync controllers are used up and no more free one
 *      - ESP_FAIL: Create sync manager failed because of other error
 */
// llgo:link (*RmtSyncManagerConfigT).RmtNewSyncManager C.rmt_new_sync_manager
func (recv_ *RmtSyncManagerConfigT) RmtNewSyncManager(ret_synchro *RmtSyncManagerHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete synchronization manager
 *
 * @param[in] synchro Synchronization manager handle returned from `rmt_new_sync_manager()`
 * @return
 *      - ESP_OK: Delete the synchronization manager successfully
 *      - ESP_ERR_INVALID_ARG: Delete the synchronization manager failed because of invalid argument
 *      - ESP_FAIL: Delete the synchronization manager failed because of other error
 */
//go:linkname RmtDelSyncManager C.rmt_del_sync_manager
func RmtDelSyncManager(synchro RmtSyncManagerHandleT) EspErrT

/**
 * @brief Reset synchronization manager
 *
 * @param[in] synchro Synchronization manager handle returned from `rmt_new_sync_manager()`
 * @return
 *      - ESP_OK: Reset the synchronization manager successfully
 *      - ESP_ERR_INVALID_ARG: Reset the synchronization manager failed because of invalid argument
 *      - ESP_FAIL: Reset the synchronization manager failed because of other error
 */
//go:linkname RmtSyncReset C.rmt_sync_reset
func RmtSyncReset(synchro RmtSyncManagerHandleT) EspErrT

/**
 * @brief Switch GPIO for RMT TX channel
 *
 * @param[in] channel RMT TX channel handle
 * @param[in] gpio_num New GPIO number to be used
 * @param[in] invert_out Whether to invert the output signal
 * @return
 *      - ESP_OK: Switch GPIO successfully
 *      - ESP_ERR_INVALID_ARG: Switch GPIO failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Switch GPIO failed because channel is not disabled
 *      - ESP_FAIL: Switch GPIO failed because of other error
 */
//go:linkname RmtTxSwitchGpio C.rmt_tx_switch_gpio
func RmtTxSwitchGpio(channel RmtChannelHandleT, gpio_num GpioNumT, invert_out bool) EspErrT
