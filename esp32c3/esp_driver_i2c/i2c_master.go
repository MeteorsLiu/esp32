package esp_driver_i2c

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief I2C master bus specific configurations
 */

type I2cMasterBusConfigT struct {
	I2cPort         I2cPortNumT
	SdaIoNum        GpioNumT
	SclIoNum        GpioNumT
	GlitchIgnoreCnt c.Uint8T
	IntrPriority    c.Int
	TransQueueDepth c.SizeT
	Flags           struct {
		EnableInternalPullup c.Uint32T
		AllowPd              c.Uint32T
	}
}

/**
 * @brief I2C device configuration
 */

type I2cDeviceConfigT struct {
	DevAddrLength I2cAddrBitLenT
	DeviceAddress c.Uint16T
	SclSpeedHz    c.Uint32T
	SclWaitUs     c.Uint32T
	Flags         struct {
		DisableAckCheck c.Uint32T
	}
}

/**
 * @brief Structure representing an I2C operation job
 *
 * This structure is used to define individual I2C operations (write or read)
 * within a sequence of I2C master transactions.
 */

type I2cOperationJobT struct {
	Command I2cMasterCommandT
}

/**
 * @brief I2C master transmit buffer information structure
 */

type I2cMasterTransmitMultiBufferInfoT struct {
	WriteBuffer *c.Uint8T
	BufferSize  c.SizeT
}

/**
 * @brief Group of I2C master callbacks, can be used to get status during transaction or doing other small things. But take care potential concurrency issues.
 * @note The callbacks are all running under ISR context
 * @note When CONFIG_I2C_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well.
 */

type I2cMasterEventCallbacksT struct {
	OnTransDone I2cMasterCallbackT
}

/**
 * @brief Allocate an I2C master bus
 *
 * @param[in] bus_config I2C master bus configuration.
 * @param[out] ret_bus_handle I2C bus handle
 * @return
 *      - ESP_OK: I2C master bus initialized successfully.
 *      - ESP_ERR_INVALID_ARG: I2C bus initialization failed because of invalid argument.
 *      - ESP_ERR_NO_MEM: Create I2C bus failed because of out of memory.
 *      - ESP_ERR_NOT_FOUND: No more free bus.
 */
// llgo:link (*I2cMasterBusConfigT).I2cNewMasterBus C.i2c_new_master_bus
func (recv_ *I2cMasterBusConfigT) I2cNewMasterBus(ret_bus_handle *I2cMasterBusHandleT) EspErrT {
	return 0
}

/**
 * @brief Add I2C master BUS device.
 *
 * @param[in] bus_handle I2C bus handle.
 * @param[in] dev_config device config.
 * @param[out] ret_handle device handle.
 * @return
 *      - ESP_OK: Create I2C master device successfully.
 *      - ESP_ERR_INVALID_ARG: I2C bus initialization failed because of invalid argument.
 *      - ESP_ERR_NO_MEM: Create I2C bus failed because of out of memory.
 */
//go:linkname I2cMasterBusAddDevice C.i2c_master_bus_add_device
func I2cMasterBusAddDevice(bus_handle I2cMasterBusHandleT, dev_config *I2cDeviceConfigT, ret_handle *I2cMasterDevHandleT) EspErrT

/**
 * @brief Deinitialize the I2C master bus and delete the handle.
 *
 * @param[in] bus_handle I2C bus handle.
 * @return
 *      - ESP_OK: Delete I2C bus success, otherwise, failed.
 *      - Otherwise: Some module delete failed.
 */
//go:linkname I2cDelMasterBus C.i2c_del_master_bus
func I2cDelMasterBus(bus_handle I2cMasterBusHandleT) EspErrT

/**
 * @brief I2C master bus delete device
 *
 * @param handle i2c device handle
 * @return
 *      - ESP_OK: If device is successfully deleted.
 */
//go:linkname I2cMasterBusRmDevice C.i2c_master_bus_rm_device
func I2cMasterBusRmDevice(handle I2cMasterDevHandleT) EspErrT

/**
 * @brief Perform a write transaction on the I2C bus.
 *        The transaction will be undergoing until it finishes or it reaches
 *        the timeout provided.
 *
 * @note If a callback was registered with `i2c_master_register_event_callbacks`, the transaction will be asynchronous, and thus, this function will return directly, without blocking.
 *       You will get finish information from callback. Besides, data buffer should always be completely prepared when callback is registered, otherwise, the data will get corrupt.
 *
 * @param[in] i2c_dev I2C master device handle that created by `i2c_master_bus_add_device`.
 * @param[in] write_buffer Data bytes to send on the I2C bus.
 * @param[in] write_size Size, in bytes, of the write buffer.
 * @param[in] xfer_timeout_ms Wait timeout, in ms. Note: -1 means wait forever.
 * @return
 *      - ESP_OK: I2C master transmit success
 *      - ESP_ERR_INVALID_ARG: I2C master transmit parameter invalid.
 *      - ESP_ERR_TIMEOUT: Operation timeout(larger than xfer_timeout_ms) because the bus is busy or hardware crash.
 */
//go:linkname I2cMasterTransmit C.i2c_master_transmit
func I2cMasterTransmit(i2c_dev I2cMasterDevHandleT, write_buffer *c.Uint8T, write_size c.SizeT, xfer_timeout_ms c.Int) EspErrT

/**
 * @brief Transmit multiple buffers of data over an I2C bus.
 *
 * This function transmits multiple buffers of data over an I2C bus using the specified I2C master device handle.
 * It takes in an array of buffer information structures along with the size of the array and a transfer timeout value in milliseconds.
 *
 * @param i2c_dev I2C master device handle that created by `i2c_master_bus_add_device`.
 * @param buffer_info_array Pointer to buffer information array.
 * @param array_size size of buffer information array.
 * @param xfer_timeout_ms Wait timeout, in ms. Note: -1 means wait forever.
 *
 * @return
 *      - ESP_OK: I2C master transmit success
 *      - ESP_ERR_INVALID_ARG: I2C master transmit parameter invalid.
 *      - ESP_ERR_TIMEOUT: Operation timeout(larger than xfer_timeout_ms) because the bus is busy or hardware crash.
 */
//go:linkname I2cMasterMultiBufferTransmit C.i2c_master_multi_buffer_transmit
func I2cMasterMultiBufferTransmit(i2c_dev I2cMasterDevHandleT, buffer_info_array *I2cMasterTransmitMultiBufferInfoT, array_size c.SizeT, xfer_timeout_ms c.Int) EspErrT

/**
 * @brief Perform a write-read transaction on the I2C bus.
 *        The transaction will be undergoing until it finishes or it reaches
 *        the timeout provided.
 *
 * @note If a callback was registered with `i2c_master_register_event_callbacks`, the transaction will be asynchronous, and thus, this function will return directly, without blocking.
 *       You will get finish information from callback. Besides, data buffer should always be completely prepared when callback is registered, otherwise, the data will get corrupt.
 *
 * @param[in] i2c_dev I2C master device handle that created by `i2c_master_bus_add_device`.
 * @param[in] write_buffer Data bytes to send on the I2C bus.
 * @param[in] write_size Size, in bytes, of the write buffer.
 * @param[out] read_buffer Data bytes received from i2c bus.
 * @param[in] read_size Size, in bytes, of the read buffer.
 * @param[in] xfer_timeout_ms Wait timeout, in ms. Note: -1 means wait forever.
 * @return
 *      - ESP_OK: I2C master transmit-receive success
 *      - ESP_ERR_INVALID_ARG: I2C master transmit parameter invalid.
 *      - ESP_ERR_TIMEOUT: Operation timeout(larger than xfer_timeout_ms) because the bus is busy or hardware crash.
 */
//go:linkname I2cMasterTransmitReceive C.i2c_master_transmit_receive
func I2cMasterTransmitReceive(i2c_dev I2cMasterDevHandleT, write_buffer *c.Uint8T, write_size c.SizeT, read_buffer *c.Uint8T, read_size c.SizeT, xfer_timeout_ms c.Int) EspErrT

/**
 * @brief Perform a read transaction on the I2C bus.
 *        The transaction will be undergoing until it finishes or it reaches
 *        the timeout provided.
 *
 * @note If a callback was registered with `i2c_master_register_event_callbacks`, the transaction will be asynchronous, and thus, this function will return directly, without blocking.
 *       You will get finish information from callback. Besides, data buffer should always be completely prepared when callback is registered, otherwise, the data will get corrupt.
 *
 * @param[in] i2c_dev I2C master device handle that created by `i2c_master_bus_add_device`.
 * @param[out] read_buffer Data bytes received from i2c bus.
 * @param[in] read_size Size, in bytes, of the read buffer.
 * @param[in] xfer_timeout_ms Wait timeout, in ms. Note: -1 means wait forever.
 * @return
 *      - ESP_OK: I2C master receive success
 *      - ESP_ERR_INVALID_ARG: I2C master receive parameter invalid.
 *      - ESP_ERR_TIMEOUT: Operation timeout(larger than xfer_timeout_ms) because the bus is busy or hardware crash.
 */
//go:linkname I2cMasterReceive C.i2c_master_receive
func I2cMasterReceive(i2c_dev I2cMasterDevHandleT, read_buffer *c.Uint8T, read_size c.SizeT, xfer_timeout_ms c.Int) EspErrT

/**
 * @brief Probe I2C address, if address is correct and ACK is received, this function will return ESP_OK.
 *
 * @param[in] bus_handle I2C master device handle that created by `i2c_master_bus_add_device`.
 * @param[in] address I2C device address that you want to probe.
 * @param[in] xfer_timeout_ms Wait timeout, in ms. Note: -1 means wait forever (Not recommended in this function).
 *
 * @attention Pull-ups must be connected to the SCL and SDA pins when this function is called. If you get `ESP_ERR_TIMEOUT
 * while `xfer_timeout_ms` was parsed correctly, you should check the pull-up resistors. If you do not have proper resistors nearby.
 * `flags.enable_internal_pullup` is also acceptable.
 *
 * @note The principle of this function is to sent device address with a write command. If the device on your I2C bus, there would be an ACK signal and function
 * returns `ESP_OK`. If the device is not on your I2C bus, there would be a NACK signal and function returns `ESP_ERR_NOT_FOUND`. `ESP_ERR_TIMEOUT` is not an expected
 * failure, which indicated that the i2c probe not works properly, usually caused by pull-up resistors not be connected properly. Suggestion check data on SDA/SCL line
 * to see whether there is ACK/NACK signal is on line when i2c probe function fails.
 *
 * @note There are lots of I2C devices all over the world, we assume that not all I2C device support the behavior like `device_address+nack/ack`.
 * So, if the on line data is strange and no ack/nack got respond. Please check the device datasheet.
 *
 * @return
 *      - ESP_OK: I2C device probe successfully
 *      - ESP_ERR_NOT_FOUND: I2C probe failed, doesn't find the device with specific address you gave.
 *      - ESP_ERR_TIMEOUT: Operation timeout(larger than xfer_timeout_ms) because the bus is busy or hardware crash.
 */
//go:linkname I2cMasterProbe C.i2c_master_probe
func I2cMasterProbe(bus_handle I2cMasterBusHandleT, address c.Uint16T, xfer_timeout_ms c.Int) EspErrT

/**
 * @brief Execute a series of pre-defined I2C operations.
 *
 * This function processes a list of I2C operations, such as start, write, read, and stop,
 * according to the user-defined `i2c_operation_job_t` array. It performs these operations
 * sequentially on the specified I2C master device.
 *
 * @param[in] i2c_dev           Handle to the I2C master device.
 * @param[in] i2c_operation     Pointer to an array of user-defined I2C operation jobs.
 *                              Each job specifies a command and associated parameters.
 * @param[in] operation_list_num The number of operations in the `i2c_operation` array.
 * @param[in] xfer_timeout_ms   Timeout for the transaction, in milliseconds.
 *
 * @return
 *  - ESP_OK: Transaction completed successfully.
 *  - ESP_ERR_INVALID_ARG: One or more arguments are invalid.
 *  - ESP_ERR_TIMEOUT: Transaction timed out.
 *  - ESP_FAIL: Other error during transaction.
 *
 * @note The `ack_value` field in the READ operation must be set to `I2C_NACK_VAL` if the next
 *       operation is a STOP command.
 */
//go:linkname I2cMasterExecuteDefinedOperations C.i2c_master_execute_defined_operations
func I2cMasterExecuteDefinedOperations(i2c_dev I2cMasterDevHandleT, i2c_operation *I2cOperationJobT, operation_list_num c.SizeT, xfer_timeout_ms c.Int) EspErrT

/**
 * @brief Register I2C transaction callbacks for a master device
 *
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 * @note When CONFIG_I2C_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well. The `user_data` should also reside in SRAM.
 * @note If the callback is used for helping asynchronous transaction. On the same bus, only one device can be used for performing asynchronous operation.
 *
 * @param[in] i2c_dev I2C master device handle that created by `i2c_master_bus_add_device`.
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set I2C transaction callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set I2C transaction callbacks failed because of invalid argument
 *      - ESP_FAIL: Set I2C transaction callbacks failed because of other error
 */
//go:linkname I2cMasterRegisterEventCallbacks C.i2c_master_register_event_callbacks
func I2cMasterRegisterEventCallbacks(i2c_dev I2cMasterDevHandleT, cbs *I2cMasterEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Reset the I2C master bus.
 *
 * @param bus_handle I2C bus handle.
 * @return
 *      - ESP_OK: Reset succeed.
 *      - ESP_ERR_INVALID_ARG: I2C master bus handle is not initialized.
 *      - Otherwise: Reset failed.
 */
//go:linkname I2cMasterBusReset C.i2c_master_bus_reset
func I2cMasterBusReset(bus_handle I2cMasterBusHandleT) EspErrT

/**
 * @brief Wait for all pending I2C transactions done
 *
 * @param[in] bus_handle I2C bus handle
 * @param[in] timeout_ms Wait timeout, in ms. Specially, -1 means to wait forever.
 * @return
 *      - ESP_OK: Flush transactions successfully
 *      - ESP_ERR_INVALID_ARG: Flush transactions failed because of invalid argument
 *      - ESP_ERR_TIMEOUT: Flush transactions failed because of timeout
 *      - ESP_FAIL: Flush transactions failed because of other error
 */
//go:linkname I2cMasterBusWaitAllDone C.i2c_master_bus_wait_all_done
func I2cMasterBusWaitAllDone(bus_handle I2cMasterBusHandleT, timeout_ms c.Int) EspErrT

/**
 * @brief Retrieves the I2C master bus handle for a specified I2C port number.
 *
 * This function retrieves the I2C master bus handle for the
 * given I2C port number. Please make sure the handle has already been initialized, and this
 * function would simply returns the existing handle. Note that the returned handle still can't be used concurrently
 *
 * @param port_num I2C port number for which the handle is to be retrieved.
 * @param ret_handle Pointer to a variable where the retrieved handle will be stored.
 * @return
 *     - ESP_OK: Success. The handle is retrieved successfully.
 *     - ESP_ERR_INVALID_ARG: Invalid argument, such as invalid port number
 *     - ESP_ERR_INVALID_STATE: Invalid state, such as the I2C port is not initialized.
 */
// llgo:link I2cPortNumT.I2cMasterGetBusHandle C.i2c_master_get_bus_handle
func (recv_ I2cPortNumT) I2cMasterGetBusHandle(ret_handle *I2cMasterBusHandleT) EspErrT {
	return 0
}
