package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SendStateT c.Int

const (
	STATE_IDLE                SendStateT = 1
	STATE_WAIT_FOR_START      SendStateT = 2
	STATE_SENDING             SendStateT = 3
	STATE_GETTING_RESULT      SendStateT = 4
	STATE_GETTING_UNSENT_DESC SendStateT = 5
)

type SdioRingbufT struct {
	Data     *c.Uint8T
	Size     c.SizeT
	WritePtr *c.Uint8T
	ReadPtr  *c.Uint8T
	FreePtr  *c.Uint8T
}

// / DMA descriptor with extra fields
type SdioSlaveHalSendDescS struct {
	DmaDesc SdioSlaveLlDescT
	PktLen  c.Uint32T
	Arg     c.Pointer
}
type SdioSlaveHalSendDescT SdioSlaveHalSendDescS
type SdioSlaveHalRecvDescT SdioSlaveLlDescT

type RecvStailqHeadS struct {
	StqhFirst *SdioSlaveLlDescS
	StqhLast  **SdioSlaveLlDescS
}
type SdioSlaveHalRecvStailqT RecvStailqHeadS

/** HAL context structure. Call `sdio_slave_hal_init` to initialize it and
 * configure required members before actually use the HAL.
 */

type SdioSlaveContextT struct {
	SendingMode    SdioSlaveSendingModeT
	Timing         SdioSlaveTimingT
	SendQueueSize  c.Int
	RecvBufferSize c.SizeT
	SendDescQueue  SdioRingbufT
	SendState      SendStateT
	TailPktLen     c.Uint32T
	InFlightHead   *SdioSlaveHalSendDescT
	InFlightEnd    *SdioSlaveHalSendDescT
	InFlightNext   *SdioSlaveHalSendDescT
	ReturnedDesc   *SdioSlaveHalSendDescT
	RecvLinkList   SdioSlaveHalRecvStailqT
	RecvCurRet     *SdioSlaveHalRecvDescT
}

/**
 * Initialize the HAL, should provide buffers to the context and configure the
 * members before this funciton is called.
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalInit C.sdio_slave_hal_init
func (recv_ *SdioSlaveContextT) SdioSlaveHalInit() {
}

/**
 * Initialize the SDIO slave peripheral hardware.
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalHwInit C.sdio_slave_hal_hw_init
func (recv_ *SdioSlaveContextT) SdioSlaveHalHwInit() {
}

/**
 * Set the IO ready for host to read.
 *
 * @param hal Context of the HAL layer.
 * @param ready true to tell the host the slave is ready, otherwise false.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSetIoready C.sdio_slave_hal_set_ioready
func (recv_ *SdioSlaveContextT) SdioSlaveHalSetIoready(ready bool) {
}

/**
 * The hardware sending DMA starts. If there is existing data, send them.
 *
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSendStart C.sdio_slave_hal_send_start
func (recv_ *SdioSlaveContextT) SdioSlaveHalSendStart() EspErrT {
	return 0
}

/**
 * Stops hardware sending DMA.
 *
 * @note The data in the queue, as well as the counter are not touched.
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSendStop C.sdio_slave_hal_send_stop
func (recv_ *SdioSlaveContextT) SdioSlaveHalSendStop() {
}

/**
 * Put some data into the sending queue.
 *
 * @note The caller should keeps the buffer, until the `arg` is returned by
 *       `sdio_slave_hal_send_get_next_finished_arg`.
 * @note The caller should count to ensure there is enough space in the queue.
 *       The initial queue size is sizeof(sendbuf.data)/sizeof(sdio_slave_hal_send_desc_t)-1,
 *       Will decrease by one when this function successfully returns.
 *       Released only by `sdio_slave_hal_send_get_next_finished_arg` or
 *       `sdio_slave_hal_send_flush_next_buffer`.
 *
 * @note The HAL is not thread-safe. The caller should use a spinlock to ensure
 *       the `sdio_slave_hal_send_queue` and ... are not called at the same time.
 *
 * @param hal Context of the HAL layer.
 * @param addr Address of data in the memory to send.
 * @param len Length of data to send.
 * @param arg Argument indicating this sending.
 * @return Always ESP_OK.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSendQueue C.sdio_slave_hal_send_queue
func (recv_ *SdioSlaveContextT) SdioSlaveHalSendQueue(addr *c.Uint8T, len c.SizeT, arg c.Pointer) EspErrT {
	return 0
}

/**
 * The ISR should call this, to handle the SW invoking event.
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSendHandleIsrInvoke C.sdio_slave_hal_send_handle_isr_invoke
func (recv_ *SdioSlaveContextT) SdioSlaveHalSendHandleIsrInvoke() {
}

/**
 * Check whether there is no in-flight transactions, and send new packet if there
 * is new packets queued.
 *
 * @param hal Context of the HAL layer.
 * @return
 *  - ESP_OK: The DMA starts to send a new packet.
 *  - ESP_ERR_NOT_FOUND: No packet waiting to be sent.
 *  - ESP_ERR_INVALID_STATE: There is packet in-flight.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSendNewPacketIfExist C.sdio_slave_hal_send_new_packet_if_exist
func (recv_ *SdioSlaveContextT) SdioSlaveHalSendNewPacketIfExist() EspErrT {
	return 0
}

/**
 * Check whether the sending EOF has happened and clear the interrupt.
 *
 * Call `sdio_slave_hal_send_get_next_finished_arg` recursively to retrieve arguments of finished
 * buffers.
 *
 * @param hal Context of the HAL layer.
 * @return true if happened, otherwise false.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSendEofHappened C.sdio_slave_hal_send_eof_happened
func (recv_ *SdioSlaveContextT) SdioSlaveHalSendEofHappened() bool {
	return false
}

/**
 * Get the arguments of finished packets. Call recursively until all finished
 * arguments are all retrieved.
 *
 * @param hal Context of the HAL layer.
 * @param out_arg Output argument of the finished buffer.
 * @param out_returned_cnt Released queue size to be queued again.
 * @return
 *  - ESP_OK: if one argument retrieved.
 *  - ESP_ERR_NOT_FOUND: All the arguments of the finished buffers are retrieved.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSendGetNextFinishedArg C.sdio_slave_hal_send_get_next_finished_arg
func (recv_ *SdioSlaveContextT) SdioSlaveHalSendGetNextFinishedArg(out_arg *c.Pointer, out_returned_cnt *c.Uint32T) EspErrT {
	return 0
}

/**
 * Flush one buffer in the queue, no matter sent, canceled or not sent yet.
 *
 * Call recursively to clear the whole queue before deinitialization.
 *
 * @note Only call when the DMA is stopped!
 * @param hal Context of the HAL layer.
 * @param out_arg Argument indiciating the buffer to send
 * @param out_return_cnt Space in the queue released after this descriptor is flushed.
 * @return
 *  - ESP_ERR_INVALID_STATE: This function call be called only when the DMA is stopped.
 *  - ESP_ERR_NOT_FOUND: if no buffer in the queue
 *  - ESP_OK: if a buffer is successfully flushed and returned.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSendFlushNextBuffer C.sdio_slave_hal_send_flush_next_buffer
func (recv_ *SdioSlaveContextT) SdioSlaveHalSendFlushNextBuffer(out_arg *c.Pointer, out_return_cnt *c.Uint32T) EspErrT {
	return 0
}

/**
 * Walk through all the unsent buffers and reset the counter to the accumulated length of them. The data will be kept.
 *
 * @note Only call when the DMA is stopped!
 * @param hal Context of the HAL layer.
 * @return
 *  - ESP_ERR_INVALID_STATE: this function call be called only when the DMA is stopped
 *  - ESP_OK: if success
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSendResetCounter C.sdio_slave_hal_send_reset_counter
func (recv_ *SdioSlaveContextT) SdioSlaveHalSendResetCounter() EspErrT {
	return 0
}

/*---------------------------------------------------------------------------
 *                  Receive
 *--------------------------------------------------------------------------*/
/**
 * Start the receiving DMA.
 *
 * @note If there are already some buffers loaded, will receive from them first.
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalRecvStart C.sdio_slave_hal_recv_start
func (recv_ *SdioSlaveContextT) SdioSlaveHalRecvStart() {
}

/**
 * Stop the receiving DMA.
 *
 * @note Data and the counter will not be touched. You can still call
 *       `sdio_slave_hal_recv_has_next_item` to get the received buffer.
 *       And unused buffers loaded to the HAL will still be in the `loaded`
 *       state in the HAL, until returned by `sdio_slave_hal_recv_unload_desc`.
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalRecvStop C.sdio_slave_hal_recv_stop
func (recv_ *SdioSlaveContextT) SdioSlaveHalRecvStop() {
}

/**
 * Associate the buffer to the descriptor given. The descriptor may also be initialized with some
 * other data.
 *
 * @param hal Context of the HAL layer.
 * @param desc Descriptor to associate with the buffer
 * @param start Start address of the buffer
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalRecvInitDesc C.sdio_slave_hal_recv_init_desc
func (recv_ *SdioSlaveContextT) SdioSlaveHalRecvInitDesc(desc *SdioSlaveHalRecvDescT, start *c.Uint8T) {
}

/**
 * Load the buffer to the HAL to be used to receive data.
 *
 * @note Loaded buffers will be returned to the upper layer only when:
 *       1. Returned by `sdio_slave_hal_recv_has_next_item` when receiving to that buffer successfully
 *       done.
 *       2. Returned by `sdio_slave_hal_recv_unload_desc` unconditionally.
 * @param hal Context of the HAL layer.
 * @param desc Descriptor to load to the HAL to receive.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalLoadBuf C.sdio_slave_hal_load_buf
func (recv_ *SdioSlaveContextT) SdioSlaveHalLoadBuf(desc *SdioSlaveHalRecvDescT) {
}

/**
 * Check and clear the interrupt indicating a buffer has finished receiving.
 *
 * @param hal Context of the HAL layer.
 * @return true if interrupt triggered, otherwise false.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalRecvDone C.sdio_slave_hal_recv_done
func (recv_ *SdioSlaveContextT) SdioSlaveHalRecvDone() bool {
	return false
}

/**
 * Call this function recursively to check whether there is any buffer that has
 * finished receiving.
 *
 * Will walk through the linked list to find a newer finished buffer. For each successful return,
 * it means there is one finished buffer. You can one by `sdio_slave_hal_recv_unload_desc`. You can
 * also call `sdio_slave_hal_recv_has_next_item` several times continuously before you call the
 * `sdio_slave_hal_recv_unload_desc` for the same times.
 *
 * @param hal Context of the HAL layer.
 * @return true if there is
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalRecvHasNextItem C.sdio_slave_hal_recv_has_next_item
func (recv_ *SdioSlaveContextT) SdioSlaveHalRecvHasNextItem() bool {
	return false
}

/**
 * Unconditionally remove and return the first descriptor loaded to the HAL.
 *
 * Unless during de-initialization, `sdio_slave_hal_recv_has_next_item` should have succeed for the
 * same times as this function is called, to ensure the returned descriptor has finished its
 * receiving job.
 *
 * @param hal Context of the HAL layer.
 * @return The removed descriptor, NULL means the linked-list is empty.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalRecvUnloadDesc C.sdio_slave_hal_recv_unload_desc
func (recv_ *SdioSlaveContextT) SdioSlaveHalRecvUnloadDesc() *SdioSlaveHalRecvDescT {
	return nil
}

/**
 * Walk through all the unused buffers and reset the counter to the number of
 * them.
 *
 * @note Only call when the DMA is stopped!
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalRecvResetCounter C.sdio_slave_hal_recv_reset_counter
func (recv_ *SdioSlaveContextT) SdioSlaveHalRecvResetCounter() {
}

/**
 * Walk through all the used buffers, clear the finished flag and appended them
 * back to the end of the unused list, waiting to receive then.
 *
 * @note You will lose all the received data in the buffer.
 * @note Only call when the DMA is stopped!
 * @param hal Context of the HAL layer.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalRecvFlushOneBuffer C.sdio_slave_hal_recv_flush_one_buffer
func (recv_ *SdioSlaveContextT) SdioSlaveHalRecvFlushOneBuffer() {
}

/**
 * Enable some of the interrupts for the host.
 *
 * @note May have concurrency issue wit the host or other tasks, suggest only use it during
 *       initialization.
 * @param hal Context of the HAL layer.
 * @param mask Bitwise mask for the interrupts to enable.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalHostintSetEna C.sdio_slave_hal_hostint_set_ena
func (recv_ *SdioSlaveContextT) SdioSlaveHalHostintSetEna(mask *SdioSlaveHostintT) {
}

/**
 * Get the enabled interrupts.
 *
 * @param hal Context of the HAL layer.
 * @param out_int_mask Output of the enabled interrupts
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalHostintGetEna C.sdio_slave_hal_hostint_get_ena
func (recv_ *SdioSlaveContextT) SdioSlaveHalHostintGetEna(out_int_mask *SdioSlaveHostintT) {
}

/**
 * Send general purpose interrupt (slave send to host).
 * @param hal Context of the HAL layer.
 * @param mask Interrupts to send, only `SDIO_SLAVE_HOSTINT_BIT*` are allowed.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalHostintSend C.sdio_slave_hal_hostint_send
func (recv_ *SdioSlaveContextT) SdioSlaveHalHostintSend(mask *SdioSlaveHostintT) {
}

/**
 * Cleared the specified interrupts for the host.
 *
 * @param hal Context of the HAL layer.
 * @param mask Interrupts to clear.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalHostintClear C.sdio_slave_hal_hostint_clear
func (recv_ *SdioSlaveContextT) SdioSlaveHalHostintClear(mask *SdioSlaveHostintT) {
}

/**
 * Fetch the interrupt (host send to slave) status bits and clear all of them.
 * @param hal Context of the HAL layer.
 * @param out_int_mask Output interrupt status
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalSlvintFetchClear C.sdio_slave_hal_slvint_fetch_clear
func (recv_ *SdioSlaveContextT) SdioSlaveHalSlvintFetchClear(out_int_mask *SdioSlaveLlSlvintT) {
}

/**
 * Get the value of a shared general purpose register.
 *
 * @param hal Context of the HAL layer.
 * @param pos Position of the register, 4 bytes share a word. 0-63 except 24-27.
 * @return The register value.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalHostGetReg C.sdio_slave_hal_host_get_reg
func (recv_ *SdioSlaveContextT) SdioSlaveHalHostGetReg(pos c.Int) c.Uint8T {
	return 0
}

/**
 * Set the value of shared general purpose register.
 *
 * @param hal Context of the HAL layer.
 * @param pos Position of the register, 4 bytes share a word. 0-63 except 24-27.
 * @param reg Value to set.
 */
// llgo:link (*SdioSlaveContextT).SdioSlaveHalHostSetReg C.sdio_slave_hal_host_set_reg
func (recv_ *SdioSlaveContextT) SdioSlaveHalHostSetReg(pos c.Int, reg c.Uint8T) {
}
