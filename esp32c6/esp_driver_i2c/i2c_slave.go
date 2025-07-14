package esp_driver_i2c

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief I2C slave specific configurations
 */

type I2cSlaveConfigT struct {
	I2cPort      I2cPortNumT
	SdaIoNum     GpioNumT
	SclIoNum     GpioNumT
	ClkSource    I2cClockSourceT
	SendBufDepth c.Uint32T
	SlaveAddr    c.Uint16T
	AddrBitLen   I2cAddrBitLenT
	IntrPriority c.Int
	Flags        struct {
		StretchEn      c.Uint32T
		BroadcastEn    c.Uint32T
		AccessRamEn    c.Uint32T
		SlaveUnmatchEn c.Uint32T
		AllowPd        c.Uint32T
	}
}

/**
 * @brief Group of I2C slave callbacks (e.g. get i2c slave stretch cause). But take care of potential concurrency issues.
 * @note The callbacks are all running under ISR context
 * @note When CONFIG_I2C_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well.
 */

type I2cSlaveEventCallbacksT struct {
	OnStretchOccur I2cSlaveStretchCallbackT
	OnRecvDone     I2cSlaveReceivedCallbackT
}

/**
 * @brief Read bytes from I2C internal buffer. Start a job to receive I2C data.
 *
 * @note This function is non-blocking, it initiates a new receive job and then returns.
 *       User should check the received data from the `on_recv_done` callback that registered by `i2c_slave_register_event_callbacks()`.
 *
 * @param[in] i2c_slave I2C slave device handle that created by `i2c_new_slave_device`.
 * @param[out] data Buffer to store data from I2C fifo. Should be valid until `on_recv_done` is triggered.
 * @param[in] buffer_size Buffer size of data that provided by users.
 * @return
 *      - ESP_OK: I2C slave receive success.
 *      - ESP_ERR_INVALID_ARG: I2C slave receive parameter invalid.
 *      - ESP_ERR_NOT_SUPPORTED: This function should be work in fifo mode, but I2C_SLAVE_NONFIFO mode is configured
 */
//go:linkname I2cSlaveReceive C.i2c_slave_receive
func I2cSlaveReceive(i2c_slave I2cSlaveDevHandleT, data *c.Uint8T, buffer_size c.SizeT) EspErrT

/**
 * @brief Write bytes to internal ringbuffer of the I2C slave data. When the TX fifo empty, the ISR will
 *        fill the hardware FIFO with the internal ringbuffer's data.
 *
 * @note If you connect this slave device to some master device, the data transaction direction is from slave
 *       device to master device.
 *
 * @param[in] i2c_slave I2C slave device handle that created by `i2c_new_slave_device`.
 * @param[in] data Buffer to write to slave fifo, can pickup by master. Can be freed after this function returns. Equal or larger than `size`.
 * @param[in] size In bytes, of `data` buffer.
 * @param[in] xfer_timeout_ms Wait timeout, in ms. Note: -1 means wait forever.
 * @return
 *      - ESP_OK: I2C slave transmit success.
 *      - ESP_ERR_INVALID_ARG: I2C slave transmit parameter invalid.
 *      - ESP_ERR_TIMEOUT: Operation timeout(larger than xfer_timeout_ms) because the device is busy or hardware crash.
 *      - ESP_ERR_NOT_SUPPORTED: This function should be work in fifo mode, but I2C_SLAVE_NONFIFO mode is configured
 */
//go:linkname I2cSlaveTransmit C.i2c_slave_transmit
func I2cSlaveTransmit(i2c_slave I2cSlaveDevHandleT, data *c.Uint8T, size c.Int, xfer_timeout_ms c.Int) EspErrT

/**
 * @brief Read bytes from I2C internal ram. This can be only used when `access_ram_en` in configuration structure set to true.
 *
 * @param[in] i2c_slave I2C slave device handle that created by `i2c_new_slave_device`.
 * @param[in] ram_address The offset of RAM (Cannot larger than I2C RAM memory)
 * @param[out] data Buffer to store data read from I2C ram.
 * @param[in] receive_size Received size from RAM.
 * @return
 *      - ESP_OK: I2C slave transmit success.
 *      - ESP_ERR_INVALID_ARG: I2C slave transmit parameter invalid.
 *      - ESP_ERR_NOT_SUPPORTED: This function should be work in non-fifo mode, but I2C_SLAVE_FIFO mode is configured
 */
//go:linkname I2cSlaveReadRam C.i2c_slave_read_ram
func I2cSlaveReadRam(i2c_slave I2cSlaveDevHandleT, ram_address c.Uint8T, data *c.Uint8T, receive_size c.SizeT) EspErrT

/**
 * @brief Write bytes to I2C internal ram. This can be only used when `access_ram_en` in configuration structure set to true.
 *
 * @param[in] i2c_slave I2C slave device handle that created by `i2c_new_slave_device`.
 * @param[in] ram_address The offset of RAM (Cannot larger than I2C RAM memory)
 * @param[in] data Buffer to fill.
 * @param[in] size Received size from RAM.
 * @return
 *      - ESP_OK: I2C slave transmit success.
 *      - ESP_ERR_INVALID_ARG: I2C slave transmit parameter invalid.
 *      - ESP_ERR_INVALID_SIZE: Write size is larger than
 *      - ESP_ERR_NOT_SUPPORTED: This function should be work in non-fifo mode, but I2C_SLAVE_FIFO mode is configured
 */
//go:linkname I2cSlaveWriteRam C.i2c_slave_write_ram
func I2cSlaveWriteRam(i2c_slave I2cSlaveDevHandleT, ram_address c.Uint8T, data *c.Uint8T, size c.SizeT) EspErrT

/**
 * @brief Initialize an I2C slave device
 *
 * @param[in] slave_config I2C slave device configurations
 * @param[out] ret_handle Return a generic I2C device handle
 * @return
 *      - ESP_OK: I2C slave device initialized successfully
 *      - ESP_ERR_INVALID_ARG: I2C device initialization failed because of invalid argument.
 *      - ESP_ERR_NO_MEM: Create I2C device failed because of out of memory.
 */
// llgo:link (*I2cSlaveConfigT).I2cNewSlaveDevice C.i2c_new_slave_device
func (recv_ *I2cSlaveConfigT) I2cNewSlaveDevice(ret_handle *I2cSlaveDevHandleT) EspErrT {
	return 0
}

/**
 * @brief Set I2C slave event callbacks for I2C slave channel.
 *
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `cbs` structure to NULL.
 * @note When CONFIG_I2C_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well. The `user_data` should also reside in SRAM.
 *
 * @param[in] i2c_slave I2C slave device handle that created by `i2c_new_slave_device`.
 * @param[in] cbs Group of callback functions
 * @param[in] user_data User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK: Set I2C transaction callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set I2C transaction callbacks failed because of invalid argument
 *      - ESP_FAIL: Set I2C transaction callbacks failed because of other error
 */
//go:linkname I2cSlaveRegisterEventCallbacks C.i2c_slave_register_event_callbacks
func I2cSlaveRegisterEventCallbacks(i2c_slave I2cSlaveDevHandleT, cbs *I2cSlaveEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Deinitialize the I2C slave device
 *
 * @param[in] i2c_slave I2C slave device handle that created by `i2c_new_slave_device`.
 * @return
 *      - ESP_OK: Delete I2C device successfully.
 *      - ESP_ERR_INVALID_ARG: I2C device initialization failed because of invalid argument.
 */
//go:linkname I2cDelSlaveDevice C.i2c_del_slave_device
func I2cDelSlaveDevice(i2c_slave I2cSlaveDevHandleT) EspErrT
