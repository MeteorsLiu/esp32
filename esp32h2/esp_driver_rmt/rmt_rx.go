package esp_driver_rmt

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Group of RMT RX callbacks
 * @note The callbacks are all running under ISR environment
 * @note When CONFIG_RMT_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well.
 */

type RmtRxEventCallbacksT struct {
	OnRecvDone RmtRxDoneCallbackT
}

/**
 * @brief RMT RX channel specific configuration
 */

type RmtRxChannelConfigT struct {
	GpioNum         GpioNumT
	ClkSrc          RmtClockSourceT
	ResolutionHz    c.Uint32T
	MemBlockSymbols c.SizeT
	IntrPriority    c.Int
	Flags           struct {
		InvertIn   c.Uint32T
		WithDma    c.Uint32T
		IoLoopBack c.Uint32T
		AllowPd    c.Uint32T
	}
}

/**
 * @brief RMT receive specific configuration
 */

type RmtReceiveConfigT struct {
	SignalRangeMinNs c.Uint32T
	SignalRangeMaxNs c.Uint32T
	Flags            ExtraRmtReceiveFlags
}

type ExtraRmtReceiveFlags struct {
	Unused [8]uint8
}

/**
 * @brief Create a RMT RX channel
 *
 * @param[in] config RX channel configurations
 * @param[out] ret_chan Returned generic RMT channel handle
 * @return
 *      - ESP_OK: Create RMT RX channel successfully
 *      - ESP_ERR_INVALID_ARG: Create RMT RX channel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create RMT RX channel failed because out of memory
 *      - ESP_ERR_NOT_FOUND: Create RMT RX channel failed because all RMT channels are used up and no more free one
 *      - ESP_ERR_NOT_SUPPORTED: Create RMT RX channel failed because some feature is not supported by hardware, e.g. DMA feature is not supported by hardware
 *      - ESP_FAIL: Create RMT RX channel failed because of other error
 */
// llgo:link (*RmtRxChannelConfigT).RmtNewRxChannel C.rmt_new_rx_channel
func (recv_ *RmtRxChannelConfigT) RmtNewRxChannel(ret_chan *RmtChannelHandleT) EspErrT {
	return 0
}

/**
 * @brief Initiate a receive job for RMT RX channel
 *
 * @note This function is non-blocking, it initiates a new receive job and then returns.
 *       User should check the received data from the `on_recv_done` callback that registered by `rmt_rx_register_event_callbacks()`.
 * @note This function can also be called in ISR context.
 * @note If you want this function to work even when the flash cache is disabled, please enable the `CONFIG_RMT_RECV_FUNC_IN_IRAM` option.
 *
 * @param[in] rx_channel RMT RX channel that created by `rmt_new_rx_channel()`
 * @param[in] buffer The buffer to store the received RMT symbols
 * @param[in] buffer_size size of the `buffer`, in bytes
 * @param[in] config Receive specific configurations
 * @return
 *      - ESP_OK: Initiate receive job successfully
 *      - ESP_ERR_INVALID_ARG: Initiate receive job failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Initiate receive job failed because channel is not enabled
 *      - ESP_FAIL: Initiate receive job failed because of other error
 */
//go:linkname RmtReceive C.rmt_receive
func RmtReceive(rx_channel RmtChannelHandleT, buffer c.Pointer, buffer_size c.SizeT, config *RmtReceiveConfigT) EspErrT

/**
 * @brief Set callbacks for RMT RX channel
 *
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 * @note When CONFIG_RMT_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well. The `user_data` should also reside in SRAM.
 *
 * @param[in] rx_channel RMT generic channel that created by `rmt_new_rx_channel()`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname RmtRxRegisterEventCallbacks C.rmt_rx_register_event_callbacks
func RmtRxRegisterEventCallbacks(rx_channel RmtChannelHandleT, cbs *RmtRxEventCallbacksT, user_data c.Pointer) EspErrT
