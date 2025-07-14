package esp_driver_parlio

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Parallel IO TX unit configuration
 */

type ParlioTxUnitConfigT struct {
	ClkSrc            ParlioClockSourceT
	ClkInGpioNum      GpioNumT
	InputClkSrcFreqHz c.Uint32T
	OutputClkFreqHz   c.Uint32T
	DataWidth         c.SizeT
	DataGpioNums      [8]GpioNumT
	ClkOutGpioNum     GpioNumT
	ValidGpioNum      GpioNumT
	TransQueueDepth   c.SizeT
	MaxTransferSize   c.SizeT
	DmaBurstSize      c.SizeT
	SampleEdge        ParlioSampleEdgeT
	BitPackOrder      ParlioBitPackOrderT
	Flags             struct {
		ClkGateEn  c.Uint32T
		IoLoopBack c.Uint32T
		AllowPd    c.Uint32T
	}
}

/**
 * @brief Create a Parallel IO TX unit
 *
 * @param[in] config Parallel IO TX unit configuration
 * @param[out] ret_unit Returned Parallel IO TX unit handle
 * @return
 *      - ESP_OK: Create Parallel IO TX unit successfully
 *      - ESP_ERR_INVALID_ARG: Create Parallel IO TX unit failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create Parallel IO TX unit failed because of out of memory
 *      - ESP_ERR_NOT_FOUND: Create Parallel IO TX unit failed because all TX units are used up and no more free one
 *      - ESP_ERR_NOT_SUPPORTED: Create Parallel IO TX unit failed because some feature is not supported by hardware, e.g. clock gating
 *      - ESP_FAIL: Create Parallel IO TX unit failed because of other error
 */
// llgo:link (*ParlioTxUnitConfigT).ParlioNewTxUnit C.parlio_new_tx_unit
func (recv_ *ParlioTxUnitConfigT) ParlioNewTxUnit(ret_unit *ParlioTxUnitHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete a Parallel IO TX unit
 *
 * @param[in] unit Parallel IO TX unit that created by `parlio_new_tx_unit`
 * @return
 *      - ESP_OK: Delete Parallel IO TX unit successfully
 *      - ESP_ERR_INVALID_ARG: Delete Parallel IO TX unit failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Delete Parallel IO TX unit failed because it is still in working
 *      - ESP_FAIL: Delete Parallel IO TX unit failed because of other error
 */
//go:linkname ParlioDelTxUnit C.parlio_del_tx_unit
func ParlioDelTxUnit(unit ParlioTxUnitHandleT) EspErrT

/**
 * @brief Enable the Parallel IO TX unit
 *
 * @note This function will transit the driver state from init to enable
 * @note This function will acquire a PM lock that might be installed during channel allocation
 * @note If there're transaction pending in the queue, this function will pick up the first one and start the transfer
 *
 * @param[in] unit Parallel IO TX unit that created by `parlio_new_tx_unit`
 * @return
 *      - ESP_OK: Enable Parallel IO TX unit successfully
 *      - ESP_ERR_INVALID_ARG: Enable Parallel IO TX unit failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Enable Parallel IO TX unit failed because it is already enabled
 *      - ESP_FAIL: Enable Parallel IO TX unit failed because of other error
 */
//go:linkname ParlioTxUnitEnable C.parlio_tx_unit_enable
func ParlioTxUnitEnable(unit ParlioTxUnitHandleT) EspErrT

/**
 * @brief Disable the Parallel IO TX unit
 *
 * @note This function will transit the driver state from enable to init
 * @note This function will release the PM lock that might be installed during channel allocation
 * @note If one transaction is undergoing, this function will terminate it immediately
 *
 * @param[in] unit Parallel IO TX unit that created by `parlio_new_tx_unit`
 * @return
 *      - ESP_OK: Disable Parallel IO TX unit successfully
 *      - ESP_ERR_INVALID_ARG: Disable Parallel IO TX unit failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Disable Parallel IO TX unit failed because it's not enabled yet
 *      - ESP_FAIL: Disable Parallel IO TX unit failed because of other error
 */
//go:linkname ParlioTxUnitDisable C.parlio_tx_unit_disable
func ParlioTxUnitDisable(unit ParlioTxUnitHandleT) EspErrT

/**
 * @brief Type of Parallel IO TX done event data
 */

type ParlioTxDoneEventDataT struct {
	Unused [8]uint8
}

// llgo:type C
type ParlioTxDoneCallbackT func(ParlioTxUnitHandleT, *ParlioTxDoneEventDataT, c.Pointer) bool

/**
 * @brief Group of Parallel IO TX callbacks
 * @note The callbacks are all running under ISR environment
 * @note When CONFIG_PARLIO_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well.
 */

type ParlioTxEventCallbacksT struct {
	OnTransDone ParlioTxDoneCallbackT
}

/**
 * @brief Set event callbacks for Parallel IO TX unit
 *
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 * @note When CONFIG_PARLIO_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well. The `user_data` should also reside in SRAM.
 *
 * @param[in] tx_unit Parallel IO TX unit that created by `parlio_new_tx_unit`
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname ParlioTxUnitRegisterEventCallbacks C.parlio_tx_unit_register_event_callbacks
func ParlioTxUnitRegisterEventCallbacks(tx_unit ParlioTxUnitHandleT, cbs *ParlioTxEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Parallel IO transmit configuration
 */

type ParlioTransmitConfigT struct {
	IdleValue c.Uint32T
	Flags     struct {
		QueueNonblocking c.Uint32T
	}
}

/**
 * @brief Transmit data on by Parallel IO TX unit
 *
 * @note After the function returns, it doesn't mean the transaction is finished. This function only constructs a transaction structure and push into a queue.
 *
 * @param[in] tx_unit Parallel IO TX unit that created by `parlio_new_tx_unit`
 * @param[in] payload Pointer to the data to be transmitted
 * @param[in] payload_bits Length of the data to be transmitted, in bits
 * @param[in] config Transmit configuration
 * @return
 *      - ESP_OK: Transmit data successfully
 *      - ESP_ERR_INVALID_ARG: Transmit data failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Transmit data failed because the Parallel IO TX unit is not enabled
 *      - ESP_FAIL: Transmit data failed because of other error
 */
//go:linkname ParlioTxUnitTransmit C.parlio_tx_unit_transmit
func ParlioTxUnitTransmit(tx_unit ParlioTxUnitHandleT, payload c.Pointer, payload_bits c.SizeT, config *ParlioTransmitConfigT) EspErrT

/**
 * @brief Wait for all pending TX transactions done
 *
 * @param[in] tx_unit Parallel IO TX unit that created by `parlio_new_tx_unit`
 * @param[in] timeout_ms Timeout in milliseconds, `-1` means to wait forever
 * @return
 *      - ESP_OK: All pending TX transactions is finished and recycled
 *      - ESP_ERR_INVALID_ARG: Wait for all pending TX transactions done failed because of invalid argument
 *      - ESP_ERR_TIMEOUT: Wait for all pending TX transactions done timeout
 *      - ESP_FAIL: Wait for all pending TX transactions done failed because of other error
 */
//go:linkname ParlioTxUnitWaitAllDone C.parlio_tx_unit_wait_all_done
func ParlioTxUnitWaitAllDone(tx_unit ParlioTxUnitHandleT, timeout_ms c.Int) EspErrT
