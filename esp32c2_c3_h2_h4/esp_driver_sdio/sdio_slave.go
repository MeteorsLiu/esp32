package esp_driver_sdio

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SDIO_SLAVE_FLAG_HIGH_SPEED = 0

// llgo:type C
type SdioEventCbT func(c.Uint8T)

// / Configuration of SDIO slave
type SdioSlaveConfigT struct {
	Timing         SdioSlaveTimingT
	SendingMode    SdioSlaveSendingModeT
	SendQueueSize  c.Int
	RecvBufferSize c.SizeT
	EventCb        SdioEventCbT
	Flags          c.Uint32T
}
type SdioSlaveBufHandleT c.Pointer

/** Initialize the sdio slave driver
 *
 * @param config Configuration of the sdio slave driver.
 *
 * @return
 *     - ESP_ERR_NOT_FOUND if no free interrupt found.
 *     - ESP_ERR_INVALID_STATE if already initialized.
 *     - ESP_ERR_NO_MEM if fail due to memory allocation failed.
 *     - ESP_OK if success
 */
// llgo:link (*SdioSlaveConfigT).SdioSlaveInitialize C.sdio_slave_initialize
func (recv_ *SdioSlaveConfigT) SdioSlaveInitialize() EspErrT {
	return 0
}

/** De-initialize the sdio slave driver to release the resources.
 */
//go:linkname SdioSlaveDeinit C.sdio_slave_deinit
func SdioSlaveDeinit()

/** Start hardware for sending and receiving, as well as set the IOREADY1 to 1.
 *
 * @note The driver will continue sending from previous data and PKT_LEN counting, keep data received as well as start receiving from current TOKEN1 counting.
 * See ``sdio_slave_reset``.
 *
 * @return
 *  - ESP_ERR_INVALID_STATE if already started.
 *  - ESP_OK otherwise.
 */
//go:linkname SdioSlaveStart C.sdio_slave_start
func SdioSlaveStart() EspErrT

/** Stop hardware from sending and receiving, also set IOREADY1 to 0.
 *
 * @note this will not clear the data already in the driver, and also not reset the PKT_LEN and TOKEN1 counting. Call ``sdio_slave_reset`` to do that.
 */
//go:linkname SdioSlaveStop C.sdio_slave_stop
func SdioSlaveStop()

/** Clear the data still in the driver, as well as reset the PKT_LEN and TOKEN1 counting.
 *
 * @return always return ESP_OK.
 */
//go:linkname SdioSlaveReset C.sdio_slave_reset
func SdioSlaveReset() EspErrT

/*---------------------------------------------------------------------------
 *                  Receive
 *--------------------------------------------------------------------------*/
/** Register buffer used for receiving. All buffers should be registered before used, and then can be used (again) in the driver by the handle returned.
 *
 * @param start The start address of the buffer.
 *
 * @note The driver will use and only use the amount of space specified in the `recv_buffer_size` member set in the `sdio_slave_config_t`.
 *       All buffers should be larger than that. The buffer is used by the DMA, so it should be DMA capable and 32-bit aligned.
 *
 * @return The buffer handle if success, otherwise NULL.
 */
//go:linkname SdioSlaveRecvRegisterBuf C.sdio_slave_recv_register_buf
func SdioSlaveRecvRegisterBuf(start *c.Uint8T) SdioSlaveBufHandleT

/** Unregister buffer from driver, and free the space used by the descriptor pointing to the buffer.
 *
 * @param handle Handle to the buffer to release.
 *
 * @return ESP_OK if success, ESP_ERR_INVALID_ARG if the handle is NULL or the buffer is being used.
 */
//go:linkname SdioSlaveRecvUnregisterBuf C.sdio_slave_recv_unregister_buf
func SdioSlaveRecvUnregisterBuf(handle SdioSlaveBufHandleT) EspErrT

/** Load buffer to the queue waiting to receive data. The driver takes ownership of the buffer until the buffer is returned by
 *  ``sdio_slave_send_get_finished`` after the transaction is finished.
 *
 * @param handle Handle to the buffer ready to receive data.
 *
 * @return
 *     - ESP_ERR_INVALID_ARG    if invalid handle or the buffer is already in the queue. Only after the buffer is returened by
 *                              ``sdio_slave_recv`` can you load it again.
 *     - ESP_OK if success
 */
//go:linkname SdioSlaveRecvLoadBuf C.sdio_slave_recv_load_buf
func SdioSlaveRecvLoadBuf(handle SdioSlaveBufHandleT) EspErrT

/** Get buffer of received data if exist with packet information. The driver returns the ownership of the buffer to the app.
 *
 * When you see return value is ``ESP_ERR_NOT_FINISHED``, you should call this API iteratively until the return value is ``ESP_OK``.
 * All the continuous buffers returned with ``ESP_ERR_NOT_FINISHED``, together with the last buffer returned with ``ESP_OK``, belong to one packet from the host.
 *
 * You can call simpler ``sdio_slave_recv`` instead, if the host never send data longer than the Receiving buffer size,
 * or you don't care about the packet boundary (e.g. the data is only a byte stream).
 *
 * @param handle_ret Handle of the buffer holding received data. Use this handle in ``sdio_slave_recv_load_buf()`` to receive in the same buffer again.
 * @param wait Time to wait before data received.
 *
 * @note Call ``sdio_slave_load_buf`` with the handle to re-load the buffer onto the link list, and receive with the same buffer again.
 *       The address and length of the buffer got here is the same as got from `sdio_slave_get_buffer`.
 *
 * @return
 *     - ESP_ERR_INVALID_ARG    if handle_ret is NULL
 *     - ESP_ERR_TIMEOUT        if timeout before receiving new data
 *     - ESP_ERR_NOT_FINISHED   if returned buffer is not the end of a packet from the host, should call this API again until the end of a packet
 *     - ESP_OK if success
 */
//go:linkname SdioSlaveRecvPacket C.sdio_slave_recv_packet
func SdioSlaveRecvPacket(handle_ret *SdioSlaveBufHandleT, wait TickTypeT) EspErrT

/** Get received data if exist. The driver returns the ownership of the buffer to the app.
 *
 * @param handle_ret Handle to the buffer holding received data. Use this handle in ``sdio_slave_recv_load_buf`` to receive in the same buffer again.
 * @param[out] out_addr Output of the start address, set to NULL if not needed.
 * @param[out] out_len Actual length of the data in the buffer, set to NULL if not needed.
 * @param wait Time to wait before data received.
 *
 * @note Call ``sdio_slave_load_buf`` with the handle to re-load the buffer onto the link list, and receive with the same buffer again.
 *       The address and length of the buffer got here is the same as got from `sdio_slave_get_buffer`.
 *
 * @return
 *     - ESP_ERR_INVALID_ARG    if handle_ret is NULL
 *     - ESP_ERR_TIMEOUT        if timeout before receiving new data
 *     - ESP_OK if success
 */
//go:linkname SdioSlaveRecv C.sdio_slave_recv
func SdioSlaveRecv(handle_ret *SdioSlaveBufHandleT, out_addr **c.Uint8T, out_len *c.SizeT, wait TickTypeT) EspErrT

/** Retrieve the buffer corresponding to a handle.
 *
 * @param handle Handle to get the buffer.
 * @param len_o Output of buffer length
 *
 * @return buffer address if success, otherwise NULL.
 */
//go:linkname SdioSlaveRecvGetBuf C.sdio_slave_recv_get_buf
func SdioSlaveRecvGetBuf(handle SdioSlaveBufHandleT, len_o *c.SizeT) *c.Uint8T

/*---------------------------------------------------------------------------
 *                  Send
 *--------------------------------------------------------------------------*/
/** Put a new sending transfer into the send queue. The driver takes ownership of the buffer until the buffer is returned by
 *  ``sdio_slave_send_get_finished`` after the transaction is finished.
 *
 * @param addr Address for data to be sent. The buffer should be DMA capable and 32-bit aligned.
 * @param len Length of the data, should not be longer than 4092 bytes (may support longer in the future).
 * @param arg Argument to returned in ``sdio_slave_send_get_finished``. The argument can be used to indicate which transaction is done,
 *            or as a parameter for a callback. Set to NULL if not needed.
 * @param wait Time to wait if the buffer is full.
 *
 * @return
 *     - ESP_ERR_INVALID_ARG if the length is not greater than 0.
 *     - ESP_ERR_TIMEOUT if the queue is still full until timeout.
 *     - ESP_OK if success.
 */
//go:linkname SdioSlaveSendQueue C.sdio_slave_send_queue
func SdioSlaveSendQueue(addr *c.Uint8T, len c.SizeT, arg c.Pointer, wait TickTypeT) EspErrT

/** Return the ownership of a finished transaction.
 * @param out_arg Argument of the finished transaction. Set to NULL if unused.
 * @param wait Time to wait if there's no finished sending transaction.
 *
 * @return ESP_ERR_TIMEOUT if no transaction finished, or ESP_OK if succeed.
 */
//go:linkname SdioSlaveSendGetFinished C.sdio_slave_send_get_finished
func SdioSlaveSendGetFinished(out_arg *c.Pointer, wait TickTypeT) EspErrT

/** Start a new sending transfer, and wait for it (blocked) to be finished.
 *
 * @param addr Start address of the buffer to send
 * @param len Length of buffer to send.
 *
 * @return
 *     - ESP_ERR_INVALID_ARG if the length of descriptor is not greater than 0.
 *     - ESP_ERR_TIMEOUT if the queue is full or host do not start a transfer before timeout.
 *     - ESP_OK if success.
 */
//go:linkname SdioSlaveTransmit C.sdio_slave_transmit
func SdioSlaveTransmit(addr *c.Uint8T, len c.SizeT) EspErrT

/*---------------------------------------------------------------------------
 *                  Host
 *--------------------------------------------------------------------------*/
/** Read the spi slave register shared with host.
 *
 * @param pos register address, 0-27 or 32-63.
 *
 * @note register 28 to 31 are reserved for interrupt vector.
 *
 * @return value of the register.
 */
//go:linkname SdioSlaveReadReg C.sdio_slave_read_reg
func SdioSlaveReadReg(pos c.Int) c.Uint8T

/** Write the spi slave register shared with host.
 *
 * @param pos register address, 0-11, 14-15, 18-19, 24-27 and 32-63, other address are reserved.
 * @param reg the value to write.
 *
 * @note register 29 and 31 are used for interrupt vector.
 *
 * @return ESP_ERR_INVALID_ARG if address wrong, otherwise ESP_OK.
 */
//go:linkname SdioSlaveWriteReg C.sdio_slave_write_reg
func SdioSlaveWriteReg(pos c.Int, reg c.Uint8T) EspErrT

/** Get the interrupt enable for host.
 *
 * @return the interrupt mask.
 */
//go:linkname SdioSlaveGetHostIntena C.sdio_slave_get_host_intena
func SdioSlaveGetHostIntena() SdioSlaveHostintT

/** Set the interrupt enable for host.
 *
 * @param mask Enable mask for host interrupt.
 */
//go:linkname SdioSlaveSetHostIntena C.sdio_slave_set_host_intena
func SdioSlaveSetHostIntena(mask SdioSlaveHostintT)

/** Interrupt the host by general purpose interrupt.
 *
 * @param pos Interrupt num, 0-7.
 *
 * @return
 *     - ESP_ERR_INVALID_ARG if interrupt num error
 *     - ESP_OK otherwise
 */
//go:linkname SdioSlaveSendHostInt C.sdio_slave_send_host_int
func SdioSlaveSendHostInt(pos c.Uint8T) EspErrT

/** Clear general purpose interrupt to host.
 *
 * @param mask Interrupt bits to clear, by bit mask.
 */
//go:linkname SdioSlaveClearHostInt C.sdio_slave_clear_host_int
func SdioSlaveClearHostInt(mask SdioSlaveHostintT)

/** Wait for general purpose interrupt from host.
 *
 * @param pos Interrupt source number to wait for.
 * is set.
 * @param wait Time to wait before interrupt triggered.
 *
 * @note this clears the interrupt at the same time.
 *
 * @return ESP_OK if success, ESP_ERR_TIMEOUT if timeout.
 */
//go:linkname SdioSlaveWaitInt C.sdio_slave_wait_int
func SdioSlaveWaitInt(pos c.Int, wait TickTypeT) EspErrT
