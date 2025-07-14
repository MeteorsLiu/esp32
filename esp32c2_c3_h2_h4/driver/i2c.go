package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief I2C initialization parameters
 */

type I2cConfigT struct {
	Mode        I2cModeT
	SdaIoNum    c.Int
	SclIoNum    c.Int
	SdaPullupEn bool
	SclPullupEn bool
	ClkFlags    c.Uint32T
}
type I2cCmdHandleT c.Pointer

/**
 * @brief Install an I2C driver
 * @note  Not all Espressif chips can support slave mode (e.g. ESP32C2)
 *
 * @param i2c_num I2C port number
 * @param mode I2C mode (either master or slave).
 * @param slv_rx_buf_len Receiving buffer size. Only slave mode will use this value, it is ignored in master mode.
 * @param slv_tx_buf_len Sending buffer size. Only slave mode will use this value, it is ignored in master mode.
 * @param intr_alloc_flags Flags used to allocate the interrupt. One or multiple (ORred) ESP_INTR_FLAG_* values.
 *                         See esp_intr_alloc.h for more info.
 *        @note
 *        In master mode, if the cache is likely to be disabled(such as write flash) and the slave is time-sensitive,
 *        `ESP_INTR_FLAG_IRAM` is suggested to be used. In this case, please use the memory allocated from internal RAM in i2c read and write function,
 *        because we can not access the psram(if psram is enabled) in interrupt handle function when cache is disabled.
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_FAIL Driver installation error
 */
//go:linkname I2cDriverInstall C.i2c_driver_install
func I2cDriverInstall(i2c_num I2cPortT, mode I2cModeT, slv_rx_buf_len c.SizeT, slv_tx_buf_len c.SizeT, intr_alloc_flags c.Int) EspErrT

/**
 * @brief Delete I2C driver
 *
 * @note This function does not guarantee thread safety.
 *       Please make sure that no thread will continuously hold semaphores before calling the delete function.
 *
 * @param i2c_num I2C port to delete
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cDriverDelete C.i2c_driver_delete
func I2cDriverDelete(i2c_num I2cPortT) EspErrT

/**
 * @brief Configure an I2C bus with the given configuration.
 *
 * @param i2c_num I2C port to configure
 * @param i2c_conf Pointer to the I2C configuration
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cParamConfig C.i2c_param_config
func I2cParamConfig(i2c_num I2cPortT, i2c_conf *I2cConfigT) EspErrT

/**
 * @brief reset I2C tx hardware fifo
 *
 * @param i2c_num I2C port number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cResetTxFifo C.i2c_reset_tx_fifo
func I2cResetTxFifo(i2c_num I2cPortT) EspErrT

/**
 * @brief reset I2C rx fifo
 *
 * @param i2c_num I2C port number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cResetRxFifo C.i2c_reset_rx_fifo
func I2cResetRxFifo(i2c_num I2cPortT) EspErrT

/**
 * @brief Configure GPIO pins for I2C SCK and SDA signals.
 *
 * @param i2c_num I2C port number
 * @param sda_io_num GPIO number for I2C SDA signal
 * @param scl_io_num GPIO number for I2C SCL signal
 * @param sda_pullup_en Enable the internal pullup for SDA pin
 * @param scl_pullup_en Enable the internal pullup for SCL pin
 * @param mode I2C mode
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cSetPin C.i2c_set_pin
func I2cSetPin(i2c_num I2cPortT, sda_io_num c.Int, scl_io_num c.Int, sda_pullup_en bool, scl_pullup_en bool, mode I2cModeT) EspErrT

/**
 * @brief Perform a write to a device connected to a particular I2C port.
 *        This function is a wrapper to `i2c_master_start()`, `i2c_master_write()`, `i2c_master_read()`, etc...
 *        It shall only be called in I2C master mode.
 *
 * @param i2c_num I2C port number to perform the transfer on
 * @param device_address I2C device's 7-bit address
 * @param write_buffer Bytes to send on the bus
 * @param write_size Size, in bytes, of the write buffer
 * @param ticks_to_wait Maximum ticks to wait before issuing a timeout.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_FAIL Sending command error, slave hasn't ACK the transfer.
 *     - ESP_ERR_INVALID_STATE I2C driver not installed or not in master mode.
 *     - ESP_ERR_TIMEOUT Operation timeout because the bus is busy.
 */
//go:linkname I2cMasterWriteToDevice C.i2c_master_write_to_device
func I2cMasterWriteToDevice(i2c_num I2cPortT, device_address c.Uint8T, write_buffer *c.Uint8T, write_size c.SizeT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Perform a read to a device connected to a particular I2C port.
 *        This function is a wrapper to `i2c_master_start()`, `i2c_master_write()`, `i2c_master_read()`, etc...
 *        It shall only be called in I2C master mode.
 *
 * @param i2c_num I2C port number to perform the transfer on
 * @param device_address I2C device's 7-bit address
 * @param read_buffer Buffer to store the bytes received on the bus
 * @param read_size Size, in bytes, of the read buffer
 * @param ticks_to_wait Maximum ticks to wait before issuing a timeout.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_FAIL Sending command error, slave hasn't ACK the transfer.
 *     - ESP_ERR_INVALID_STATE I2C driver not installed or not in master mode.
 *     - ESP_ERR_TIMEOUT Operation timeout because the bus is busy.
 */
//go:linkname I2cMasterReadFromDevice C.i2c_master_read_from_device
func I2cMasterReadFromDevice(i2c_num I2cPortT, device_address c.Uint8T, read_buffer *c.Uint8T, read_size c.SizeT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Perform a write followed by a read to a device on the I2C bus.
 *        A repeated start signal is used between the `write` and `read`, thus, the bus is
 *        not released until the two transactions are finished.
 *        This function is a wrapper to `i2c_master_start()`, `i2c_master_write()`, `i2c_master_read()`, etc...
 *        It shall only be called in I2C master mode.
 *
 * @param i2c_num I2C port number to perform the transfer on
 * @param device_address I2C device's 7-bit address
 * @param write_buffer Bytes to send on the bus
 * @param write_size Size, in bytes, of the write buffer
 * @param read_buffer Buffer to store the bytes received on the bus
 * @param read_size Size, in bytes, of the read buffer
 * @param ticks_to_wait Maximum ticks to wait before issuing a timeout.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_FAIL Sending command error, slave hasn't ACK the transfer.
 *     - ESP_ERR_INVALID_STATE I2C driver not installed or not in master mode.
 *     - ESP_ERR_TIMEOUT Operation timeout because the bus is busy.
 */
//go:linkname I2cMasterWriteReadDevice C.i2c_master_write_read_device
func I2cMasterWriteReadDevice(i2c_num I2cPortT, device_address c.Uint8T, write_buffer *c.Uint8T, write_size c.SizeT, read_buffer *c.Uint8T, read_size c.SizeT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Create and initialize an I2C commands list with a given buffer.
 *        All the allocations for data or signals (START, STOP, ACK, ...) will be
 *        performed within this buffer.
 *        This buffer must be valid during the whole transaction.
 *        After finishing the I2C transactions, it is required to call `i2c_cmd_link_delete_static()`.
 *
 * @note It is **highly** advised to not allocate this buffer on the stack. The size of the data
 *       used underneath may increase in the future, resulting in a possible stack overflow as the macro
 *       `I2C_LINK_RECOMMENDED_SIZE` would also return a bigger value.
 *       A better option is to use a buffer allocated statically or dynamically (with `malloc`).
 *
 * @param buffer Buffer to use for commands allocations
 * @param size Size in bytes of the buffer
 *
 * @return Handle to the I2C command link or NULL if the buffer provided is too small, please
 *         use `I2C_LINK_RECOMMENDED_SIZE` macro to get the recommended size for the buffer.
 */
//go:linkname I2cCmdLinkCreateStatic C.i2c_cmd_link_create_static
func I2cCmdLinkCreateStatic(buffer *c.Uint8T, size c.Uint32T) I2cCmdHandleT

/**
 * @brief Create and initialize an I2C commands list with a given buffer.
 *        After finishing the I2C transactions, it is required to call `i2c_cmd_link_delete()`
 *        to release and return the resources.
 *        The required bytes will be dynamically allocated.
 *
 * @return Handle to the I2C command link or NULL in case of insufficient dynamic memory.
 */
//go:linkname I2cCmdLinkCreate C.i2c_cmd_link_create
func I2cCmdLinkCreate() I2cCmdHandleT

/**
 * @brief Free the I2C commands list allocated statically with `i2c_cmd_link_create_static`.
 *
 * @param cmd_handle I2C commands list allocated statically. This handle should be created thanks to
 *                   `i2c_cmd_link_create_static()` function
 */
//go:linkname I2cCmdLinkDeleteStatic C.i2c_cmd_link_delete_static
func I2cCmdLinkDeleteStatic(cmd_handle I2cCmdHandleT)

/**
 * @brief Free the I2C commands list
 *
 * @param cmd_handle I2C commands list. This handle should be created thanks to
 *                   `i2c_cmd_link_create()` function
 */
//go:linkname I2cCmdLinkDelete C.i2c_cmd_link_delete
func I2cCmdLinkDelete(cmd_handle I2cCmdHandleT)

/**
 * @brief Queue a "START signal" to the given commands list.
 *        This function shall only be called in I2C master mode.
 *        Call `i2c_master_cmd_begin()` to send all the queued commands.
 *
 * @param cmd_handle I2C commands list
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_NO_MEM The static buffer used to create `cmd_handler` is too small
 *     - ESP_FAIL No more memory left on the heap
 */
//go:linkname I2cMasterStart C.i2c_master_start
func I2cMasterStart(cmd_handle I2cCmdHandleT) EspErrT

/**
 * @brief Queue a "write byte" command to the commands list.
 *        A single byte will be sent on the I2C port. This function shall only be
 *        called in I2C master mode.
 *        Call `i2c_master_cmd_begin()` to send all queued commands
 *
 * @param cmd_handle I2C commands list
 * @param data Byte to send on the port
 * @param ack_en Enable ACK signal
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_NO_MEM The static buffer used to create `cmd_handler` is too small
 *     - ESP_FAIL No more memory left on the heap
 */
//go:linkname I2cMasterWriteByte C.i2c_master_write_byte
func I2cMasterWriteByte(cmd_handle I2cCmdHandleT, data c.Uint8T, ack_en bool) EspErrT

/**
 * @brief Queue a "write (multiple) bytes" command to the commands list.
 *        This function shall only be called in I2C master mode.
 *        Call `i2c_master_cmd_begin()` to send all queued commands
 *
 * @param cmd_handle I2C commands list
 * @param data Bytes to send. This buffer shall remain **valid** until the transaction is finished.
 *             If the PSRAM is enabled and `intr_flag` is set to `ESP_INTR_FLAG_IRAM`,
 *             `data` should be allocated from internal RAM.
 * @param data_len Length, in bytes, of the data buffer
 * @param ack_en Enable ACK signal
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_NO_MEM The static buffer used to create `cmd_handler` is too small
 *     - ESP_FAIL No more memory left on the heap
 */
//go:linkname I2cMasterWrite C.i2c_master_write
func I2cMasterWrite(cmd_handle I2cCmdHandleT, data *c.Uint8T, data_len c.SizeT, ack_en bool) EspErrT

/**
 * @brief Queue a "read byte" command to the commands list.
 *        A single byte will be read on the I2C bus. This function shall only be
 *        called in I2C master mode.
 *        Call `i2c_master_cmd_begin()` to send all queued commands
 *
 * @param cmd_handle I2C commands list
 * @param data Pointer where the received byte will the stored. This buffer shall remain **valid**
 *             until the transaction is finished.
 * @param ack ACK signal
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_NO_MEM The static buffer used to create `cmd_handler` is too small
 *     - ESP_FAIL No more memory left on the heap
 */
//go:linkname I2cMasterReadByte C.i2c_master_read_byte
func I2cMasterReadByte(cmd_handle I2cCmdHandleT, data *c.Uint8T, ack I2cAckTypeT) EspErrT

/**
 * @brief Queue a "read (multiple) bytes" command to the commands list.
 *        Multiple bytes will be read on the I2C bus. This function shall only be
 *        called in I2C master mode.
 *        Call `i2c_master_cmd_begin()` to send all queued commands
 *
 * @param cmd_handle I2C commands list
 * @param data Pointer where the received bytes will the stored. This buffer shall remain **valid**
 *             until the transaction is finished.
 * @param data_len Size, in bytes, of the `data` buffer
 * @param ack ACK signal
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_NO_MEM The static buffer used to create `cmd_handler` is too small
 *     - ESP_FAIL No more memory left on the heap
 */
//go:linkname I2cMasterRead C.i2c_master_read
func I2cMasterRead(cmd_handle I2cCmdHandleT, data *c.Uint8T, data_len c.SizeT, ack I2cAckTypeT) EspErrT

/**
 * @brief Queue a "STOP signal" to the given commands list.
 *        This function shall only be called in I2C master mode.
 *        Call `i2c_master_cmd_begin()` to send all the queued commands.
 *
 * @param cmd_handle I2C commands list
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_NO_MEM The static buffer used to create `cmd_handler` is too small
 *     - ESP_FAIL No more memory left on the heap
 */
//go:linkname I2cMasterStop C.i2c_master_stop
func I2cMasterStop(cmd_handle I2cCmdHandleT) EspErrT

/**
 * @brief Send all the queued commands on the I2C bus, in master mode.
 *        The task will be blocked until all the commands have been sent out.
 *        The I2C port is protected by mutex, so this function is thread-safe.
 *        This function shall only be called in I2C master mode.
 *
 * @param i2c_num I2C port number
 * @param cmd_handle I2C commands list
 * @param ticks_to_wait Maximum ticks to wait before issuing a timeout.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_FAIL Sending command error, slave hasn't ACK the transfer.
 *     - ESP_ERR_INVALID_STATE I2C driver not installed or not in master mode.
 *     - ESP_ERR_TIMEOUT Operation timeout because the bus is busy.
 */
//go:linkname I2cMasterCmdBegin C.i2c_master_cmd_begin
func I2cMasterCmdBegin(i2c_num I2cPortT, cmd_handle I2cCmdHandleT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Write bytes to internal ringbuffer of the I2C slave data. When the TX fifo empty, the ISR will
 *        fill the hardware FIFO with the internal ringbuffer's data.
 *        @note This function shall only be called in I2C slave mode.
 *
 * @param i2c_num I2C port number
 * @param data Bytes to write into internal buffer
 * @param size Size, in bytes, of `data` buffer
 * @param ticks_to_wait Maximum ticks to wait.
 *
 * @return
 *     - ESP_FAIL (-1) Parameter error
 *     - Other (>=0) The number of data bytes pushed to the I2C slave buffer.
 */
//go:linkname I2cSlaveWriteBuffer C.i2c_slave_write_buffer
func I2cSlaveWriteBuffer(i2c_num I2cPortT, data *c.Uint8T, size c.Int, ticks_to_wait TickTypeT) c.Int

/**
 * @brief Read bytes from I2C internal buffer. When the I2C bus receives data, the ISR will copy them
 *        from the hardware RX FIFO to the internal ringbuffer.
 *        Calling this function will then copy bytes from the internal ringbuffer to the `data` user buffer.
 *        @note This function shall only be called in I2C slave mode.
 *
 * @param i2c_num I2C port number
 * @param data Buffer to fill with ringbuffer's bytes
 * @param max_size Maximum bytes to read
 * @param ticks_to_wait Maximum waiting ticks
 *
 * @return
 *     - ESP_FAIL(-1) Parameter error
 *     - Others(>=0) The number of data bytes read from I2C slave buffer.
 */
//go:linkname I2cSlaveReadBuffer C.i2c_slave_read_buffer
func I2cSlaveReadBuffer(i2c_num I2cPortT, data *c.Uint8T, max_size c.SizeT, ticks_to_wait TickTypeT) c.Int

/**
 * @brief Set I2C master clock period
 *
 * @param i2c_num I2C port number
 * @param high_period Clock cycle number during SCL is high level, high_period is a 14 bit value
 * @param low_period Clock cycle number during SCL is low level, low_period is a 14 bit value
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cSetPeriod C.i2c_set_period
func I2cSetPeriod(i2c_num I2cPortT, high_period c.Int, low_period c.Int) EspErrT

/**
 * @brief Get I2C master clock period
 *
 * @param i2c_num I2C port number
 * @param high_period pointer to get clock cycle number during SCL is high level, will get a 14 bit value
 * @param low_period pointer to get clock cycle number during SCL is low level, will get a 14 bit value
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cGetPeriod C.i2c_get_period
func I2cGetPeriod(i2c_num I2cPortT, high_period *c.Int, low_period *c.Int) EspErrT

/**
 * @brief Enable hardware filter on I2C bus
 *        Sometimes the I2C bus is disturbed by high frequency noise(about 20ns), or the rising edge of
 *        the SCL clock is very slow, these may cause the master state machine to break.
 *        Enable hardware filter can filter out high frequency interference and make the master more stable.
 *        @note Enable filter will slow down the SCL clock.
 *
 * @param i2c_num I2C port number to filter
 * @param cyc_num the APB cycles need to be filtered (0<= cyc_num <=7).
 *        When the period of a pulse is less than cyc_num * APB_cycle, the I2C controller will ignore this pulse.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cFilterEnable C.i2c_filter_enable
func I2cFilterEnable(i2c_num I2cPortT, cyc_num c.Uint8T) EspErrT

/**
 * @brief Disable filter on I2C bus
 *
 * @param i2c_num I2C port number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cFilterDisable C.i2c_filter_disable
func I2cFilterDisable(i2c_num I2cPortT) EspErrT

/**
 * @brief set I2C master start signal timing
 *
 * @param i2c_num I2C port number
 * @param setup_time clock number between the falling-edge of SDA and rising-edge of SCL for start mark, it's a 10-bit value.
 * @param hold_time clock num between the falling-edge of SDA and falling-edge of SCL for start mark, it's a 10-bit value.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cSetStartTiming C.i2c_set_start_timing
func I2cSetStartTiming(i2c_num I2cPortT, setup_time c.Int, hold_time c.Int) EspErrT

/**
 * @brief get I2C master start signal timing
 *
 * @param i2c_num I2C port number
 * @param setup_time pointer to get setup time
 * @param hold_time pointer to get hold time
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cGetStartTiming C.i2c_get_start_timing
func I2cGetStartTiming(i2c_num I2cPortT, setup_time *c.Int, hold_time *c.Int) EspErrT

/**
 * @brief set I2C master stop signal timing
 *
 * @param i2c_num I2C port number
 * @param setup_time clock num between the rising-edge of SCL and the rising-edge of SDA, it's a 10-bit value.
 * @param hold_time clock number after the STOP bit's rising-edge, it's a 14-bit value.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cSetStopTiming C.i2c_set_stop_timing
func I2cSetStopTiming(i2c_num I2cPortT, setup_time c.Int, hold_time c.Int) EspErrT

/**
 * @brief get I2C master stop signal timing
 *
 * @param i2c_num I2C port number
 * @param setup_time pointer to get setup time.
 * @param hold_time pointer to get hold time.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cGetStopTiming C.i2c_get_stop_timing
func I2cGetStopTiming(i2c_num I2cPortT, setup_time *c.Int, hold_time *c.Int) EspErrT

/**
 * @brief set I2C data signal timing
 *
 * @param i2c_num I2C port number
 * @param sample_time clock number I2C used to sample data on SDA after the rising-edge of SCL, it's a 10-bit value
 * @param hold_time clock number I2C used to hold the data after the falling-edge of SCL, it's a 10-bit value
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cSetDataTiming C.i2c_set_data_timing
func I2cSetDataTiming(i2c_num I2cPortT, sample_time c.Int, hold_time c.Int) EspErrT

/**
 * @brief get I2C data signal timing
 *
 * @param i2c_num I2C port number
 * @param sample_time pointer to get sample time
 * @param hold_time pointer to get hold time
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cGetDataTiming C.i2c_get_data_timing
func I2cGetDataTiming(i2c_num I2cPortT, sample_time *c.Int, hold_time *c.Int) EspErrT

/**
 * @brief set I2C timeout value
 * @param i2c_num I2C port number
 * @param timeout timeout value for I2C bus (unit: APB 80Mhz clock cycle)
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cSetTimeout C.i2c_set_timeout
func I2cSetTimeout(i2c_num I2cPortT, timeout c.Int) EspErrT

/**
 * @brief get I2C timeout value
 * @param i2c_num I2C port number
 * @param timeout pointer to get timeout value
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cGetTimeout C.i2c_get_timeout
func I2cGetTimeout(i2c_num I2cPortT, timeout *c.Int) EspErrT

/**
 * @brief set I2C data transfer mode
 *
 * @param i2c_num I2C port number
 * @param tx_trans_mode I2C sending data mode
 * @param rx_trans_mode I2C receiving data mode
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cSetDataMode C.i2c_set_data_mode
func I2cSetDataMode(i2c_num I2cPortT, tx_trans_mode I2cTransModeT, rx_trans_mode I2cTransModeT) EspErrT

/**
 * @brief get I2C data transfer mode
 *
 * @param i2c_num I2C port number
 * @param tx_trans_mode pointer to get I2C sending data mode
 * @param rx_trans_mode pointer to get I2C receiving data mode
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2cGetDataMode C.i2c_get_data_mode
func I2cGetDataMode(i2c_num I2cPortT, tx_trans_mode *I2cTransModeT, rx_trans_mode *I2cTransModeT) EspErrT
