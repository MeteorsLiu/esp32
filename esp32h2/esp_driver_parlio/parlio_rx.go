package esp_driver_parlio

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Parallel IO RX unit configuration
 */

type ParlioRxUnitConfigT struct {
	TransQueueDepth c.SizeT
	MaxRecvSize     c.SizeT
	DataWidth       c.SizeT
	ClkSrc          ParlioClockSourceT
	ExtClkFreqHz    c.Uint32T
	ExpClkFreqHz    c.Uint32T
	ClkInGpioNum    GpioNumT
	ClkOutGpioNum   GpioNumT
	ValidGpioNum    GpioNumT
	DataGpioNums    [8]GpioNumT
	Flags           struct {
		FreeClk    c.Uint32T
		ClkGateEn  c.Uint32T
		IoLoopBack c.Uint32T
		IoNoInit   c.Uint32T
		AllowPd    c.Uint32T
	}
}

/**
 * @brief Create a Parallel IO RX unit
 *
 * @param[in]  config   Parallel IO RX unit configuration
 * @param[out] ret_unit Returned Parallel IO RX unit handle
 * @return
 *      - ESP_ERR_INVALID_ARG       Invalid arguments in the parameter list or the rx unit configuration
 *      - ESP_ERR_NOT_FOUND         No available rx unit found
 *      - ESP_ERR_NO_MEM            No enough memory for the rx unit resources
 *      - ESP_OK                    Success to allocate the rx unit
 */
// llgo:link (*ParlioRxUnitConfigT).ParlioNewRxUnit C.parlio_new_rx_unit
func (recv_ *ParlioRxUnitConfigT) ParlioNewRxUnit(ret_unit *ParlioRxUnitHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete a Parallel IO RX unit
 *
 * @param[in] rx_unit   Parallel IO RX unit handle that created by `parlio_new_rx_unit`
 * @return
 *      - ESP_ERR_INVALID_ARG       rx_unit is NULL
 *      - ESP_ERR_INVALID_STATE     The rx unit is enabled, can't delete an enabled rx unit
 *      - ESP_OK                    Success to delete the rx unit
 */
//go:linkname ParlioDelRxUnit C.parlio_del_rx_unit
func ParlioDelRxUnit(rx_unit ParlioRxUnitHandleT) EspErrT

/**
 * @brief Configuration of level delimiter
 */

type ParlioRxLevelDelimiterConfigT struct {
	ValidSigLineId c.Uint32T
	SampleEdge     ParlioSampleEdgeT
	BitPackOrder   ParlioBitPackOrderT
	EofDataLen     c.Uint32T
	TimeoutTicks   c.Uint32T
	Flags          struct {
		ActiveLowEn c.Uint32T
	}
}

/**
 * @brief Create a level delimiter
 * @note This function only allocate the software resources, the hardware configurations
 *       will lazy installed while the transaction that using this delimiter start processing
 * @note The enable signal must be aligned with the valid data.
 * @note There're at most `SOC_PARLIO_RX_UNIT_MAX_DATA_WIDTH - 1` IO pins left for RXD
 *
 * @param[in]  config        Level delimiter configuration
 * @param[out] ret_delimiter Returned delimiter handle
 * @return
 *      - ESP_ERR_INVALID_ARG       Invalid arguments in the parameter list or the level delimiter configuration
 *      - ESP_ERR_NO_MEM            No enough memory for the level delimiter resources
 *      - ESP_OK                    Success to allocate the level delimiter
 */
// llgo:link (*ParlioRxLevelDelimiterConfigT).ParlioNewRxLevelDelimiter C.parlio_new_rx_level_delimiter
func (recv_ *ParlioRxLevelDelimiterConfigT) ParlioNewRxLevelDelimiter(ret_delimiter *ParlioRxDelimiterHandleT) EspErrT {
	return 0
}

/**
 * @brief Configuration of pulse delimiter
 */

type ParlioRxPulseDelimiterConfigT struct {
	ValidSigLineId c.Uint32T
	SampleEdge     ParlioSampleEdgeT
	BitPackOrder   ParlioBitPackOrderT
	EofDataLen     c.Uint32T
	TimeoutTicks   c.Uint32T
	Flags          struct {
		StartBitIncluded c.Uint32T
		EndBitIncluded   c.Uint32T
		HasEndPulse      c.Uint32T
		PulseInvert      c.Uint32T
	}
}

/**
 * @brief Create a pulse delimiter
 * @note This function only allocate the software resources, the hardware configurations
 *       will lazy installed while the transaction that using this delimiter start processing
 * @note There're at most `SOC_PARLIO_RX_UNIT_MAX_DATA_WIDTH - 1` IO pins left for RXD
 *
 * @param[in]  config        Pulse delimiter configuration
 * @param[out] ret_delimiter Returned delimiter handle
 * @return
 *      - ESP_ERR_INVALID_ARG       Invalid arguments in the parameter list or the pulse delimiter configuration
 *      - ESP_ERR_NO_MEM            No enough memory for the pulse delimiter resources
 *      - ESP_OK                    Success to allocate the pulse delimiter
 */
// llgo:link (*ParlioRxPulseDelimiterConfigT).ParlioNewRxPulseDelimiter C.parlio_new_rx_pulse_delimiter
func (recv_ *ParlioRxPulseDelimiterConfigT) ParlioNewRxPulseDelimiter(ret_delimiter *ParlioRxDelimiterHandleT) EspErrT {
	return 0
}

/**
 * @brief Configuration of soft delimiter
 */

type ParlioRxSoftDelimiterConfigT struct {
	SampleEdge   ParlioSampleEdgeT
	BitPackOrder ParlioBitPackOrderT
	EofDataLen   c.Uint32T
	TimeoutTicks c.Uint32T
}

/**
 * @brief Create a pulse delimiter
 * @note This function only allocate the software resources, the hardware configurations
 *       will lazy installed while the transaction that using this delimiter start processing
 * @param[in]  config        Soft delimiter configuration
 * @param[out] ret_delimiter Returned delimiter handle
 * @return
 *      - ESP_ERR_INVALID_ARG       Invalid arguments in the parameter list or the soft delimiter configuration
 *      - ESP_ERR_NO_MEM            No enough memory for the soft delimiter resources
 *      - ESP_OK                    Success to allocate the soft delimiter
 */
// llgo:link (*ParlioRxSoftDelimiterConfigT).ParlioNewRxSoftDelimiter C.parlio_new_rx_soft_delimiter
func (recv_ *ParlioRxSoftDelimiterConfigT) ParlioNewRxSoftDelimiter(ret_delimiter *ParlioRxDelimiterHandleT) EspErrT {
	return 0
}

/**
 * @brief Start/stop the soft delimiter
 * @note  Soft delimiter need to start or stop manually because it has no validating/enabling signal to indicate the data has started or stopped
 *
 * @param[in]  rx_unit      Parallel IO RX unit handle that created by `parlio_new_rx_unit`
 * @param[in]  delimiter    Delimiter handle
 * @param[in]  start_stop   Set true to start, set false to stop
 * @return
 *      - ESP_ERR_INVALID_ARG       Invalid arguments in the parameter list or not soft delimiter
 *      - ESP_ERR_INVALID_STATE     The rx unit not enabled
 *      - ESP_OK                    Success to start or stop the soft delimiter
 */
//go:linkname ParlioRxSoftDelimiterStartStop C.parlio_rx_soft_delimiter_start_stop
func ParlioRxSoftDelimiterStartStop(rx_unit ParlioRxUnitHandleT, delimiter ParlioRxDelimiterHandleT, start_stop bool) EspErrT

/**
 * @brief Delete the delimiter
 * @note  To delete the delimiter safely, please delete it after disable all the RX units
 *
 * @param[in]  delimiter     Delimiter handle
 * @return
 *      - ESP_ERR_INVALID_ARG       The input delimiter is NULL
 *      - ESP_ERR_INVALID_STATE     The delimiter is on receiving
 *      - ESP_OK                    Success to delete the delimiter
 */
//go:linkname ParlioDelRxDelimiter C.parlio_del_rx_delimiter
func ParlioDelRxDelimiter(delimiter ParlioRxDelimiterHandleT) EspErrT

/**
 * @brief Enable the Parallel IO RX unit
 *
 * @param[in]  rx_unit       Parallel IO RX unit handle that created by `parlio_new_rx_unit`
 * @param[in]  reset_queue   Whether to reset the receiving queue.
 *                           If set to false, the legacy receive transactions in the queue are still available,
 *                           If set to true, the legacy receive transactions in the queue are dropped.
 * @return
 *      - ESP_ERR_INVALID_ARG       The input rx_unit is NULL
 *      - ESP_ERR_INVALID_STATE     The rx unit has been enabled
 *      - ESP_OK                    Success to enable the rx unit
 */
//go:linkname ParlioRxUnitEnable C.parlio_rx_unit_enable
func ParlioRxUnitEnable(rx_unit ParlioRxUnitHandleT, reset_queue bool) EspErrT

/**
 * @brief Disable the Parallel IO RX unit
 *
 * @param[in]  rx_unit       Parallel IO RX unit handle that created by `parlio_new_rx_unit`
 * @return
 *      - ESP_ERR_INVALID_ARG       The input rx_unit is NULL
 *      - ESP_ERR_INVALID_STATE     The rx unit has been disabled
 *      - ESP_OK                    Success to disable the rx unit
 */
//go:linkname ParlioRxUnitDisable C.parlio_rx_unit_disable
func ParlioRxUnitDisable(rx_unit ParlioRxUnitHandleT) EspErrT

/**
 * @brief Configuration of a receive transaction
 */

type ParlioReceiveConfigT struct {
	Delimiter ParlioRxDelimiterHandleT
	Flags     struct {
		PartialRxEn   c.Uint32T
		IndirectMount c.Uint32T
	}
}

/**
 * @brief Receive data by Parallel IO RX unit
 * @note  This is a non-blocking and asynchronous function. To block or realize synchronous receive,
 *        please call `parlio_rx_unit_wait_all_done` after this function
 * @note  The receive transaction will start immediately when there is not other transaction on receiving,
 *        Otherwise it will be sent to the transaction queue to wait for the bus.
 *
 * @param[in]  rx_unit       Parallel IO RX unit handle that created by `parlio_new_rx_unit`
 * @param[in]  payload       The payload buffer pointer
 * @param[in]  payload_size  The size of the payload buffer, in bytes.
 * @param[in]  recv_cfg      The configuration of this receive transaction
 * @return
 *      - ESP_ERR_INVALID_ARG       Invalid arguments in the parameter list or the receive configuration
 *      - ESP_ERR_NO_MEM            No memory for the internal DMA buffer (only when parlio_receive_config_t::indirect_mount enabled)
 *      - ESP_ERR_INVALID_STATE     Transaction queue is full, failed to queue the current transaction.
 *                                  Or the internal buffer is under using by an infinite transaction, can't allocate a new one
 *      - ESP_OK                    Success to queue the current receiving transaction
 */
//go:linkname ParlioRxUnitReceive C.parlio_rx_unit_receive
func ParlioRxUnitReceive(rx_unit ParlioRxUnitHandleT, payload c.Pointer, payload_size c.SizeT, recv_cfg *ParlioReceiveConfigT) EspErrT

/**
 * @brief Wait for all pending RX transactions done
 * @note This function will block until all receiving transactions done or timeout.
 *       When timeout occurs, either the timeout limitation too short for all transactions done,
 *       or the peripheral got stuck and no more interrupts trigger (e.g., external clock stopped).
 *
 * @param[in]  rx_unit       Parallel IO RX unit handle that created by `parlio_new_rx_unit`
 * @param[in]  timeout_ms    Timeout in milliseconds, `-1` means to wait forever (software timeout)
 * @return
 *      - ESP_ERR_INVALID_ARG       The input rx_unit is NULL
 *      - ESP_ERR_TIMEOUT           Wait for all transactions done timeout
 *      - ESP_OK                    All transaction done
 */
//go:linkname ParlioRxUnitWaitAllDone C.parlio_rx_unit_wait_all_done
func ParlioRxUnitWaitAllDone(rx_unit ParlioRxUnitHandleT, timeout_ms c.Int) EspErrT

/**
 * @brief Event callback data
 */

type ParlioRxEventDataT struct {
	Delimiter ParlioRxDelimiterHandleT
	Data      c.Pointer
	RecvBytes c.SizeT
}

// llgo:type C
type ParlioRxCallbackT func(ParlioRxUnitHandleT, *ParlioRxEventDataT, c.Pointer) bool

/**
 * @brief Parallel IO RX event callbacks
 */

type ParlioRxEventCallbacksT struct {
	OnPartialReceive ParlioRxCallbackT
	OnReceiveDone    ParlioRxCallbackT
	OnTimeout        ParlioRxCallbackT
}

/**
 * @brief Register event callbacks for Parallel IO RX unit
 *
 * @param[in]  rx_unit       Parallel IO RX unit handle that created by `parlio_new_rx_unit`
 * @param[in]  cbs           Callback group, set callback to NULL to deregister the corresponding callback (callback group pointer shouldn't be NULL)
 * @param[in]  user_data     User specified data that will be transported to the callbacks
 * @return
 *      - ESP_ERR_INVALID_ARG       The input rx_unit is NULL
 *      - ESP_ERR_INVALID_STATE     The rx unit has been enabled, callback should be registered before enabling the unit
 *      - ESP_OK                    Success to register the callbacks
 */
//go:linkname ParlioRxUnitRegisterEventCallbacks C.parlio_rx_unit_register_event_callbacks
func ParlioRxUnitRegisterEventCallbacks(rx_unit ParlioRxUnitHandleT, cbs *ParlioRxEventCallbacksT, user_data c.Pointer) EspErrT
