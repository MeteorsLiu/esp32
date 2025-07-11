package esp_driver_uart

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief UART configuration parameters for uart_param_config function
 */

type UartConfigT struct {
	BaudRate         c.Int
	DataBits         UartWordLengthT
	Parity           UartParityT
	StopBits         UartStopBitsT
	FlowCtrl         UartHwFlowcontrolT
	RxFlowCtrlThresh c.Uint8T
	Flags            struct {
		AllowPd           c.Uint32T
		BackupBeforeSleep c.Uint32T
	}
}

/**
 * @brief UART interrupt configuration parameters for uart_intr_config function
 */

type UartIntrConfigT struct {
	IntrEnableMask        c.Uint32T
	RxTimeoutThresh       c.Uint8T
	TxfifoEmptyIntrThresh c.Uint8T
	RxfifoFullThresh      c.Uint8T
}
type UartEventTypeT c.Int

const (
	UART_DATA        UartEventTypeT = 0
	UART_BREAK       UartEventTypeT = 1
	UART_BUFFER_FULL UartEventTypeT = 2
	UART_FIFO_OVF    UartEventTypeT = 3
	UART_FRAME_ERR   UartEventTypeT = 4
	UART_PARITY_ERR  UartEventTypeT = 5
	UART_DATA_BREAK  UartEventTypeT = 6
	UART_PATTERN_DET UartEventTypeT = 7
	UART_EVENT_MAX   UartEventTypeT = 8
)

/**
 * @brief Event structure used in UART event queue
 */

type UartEventT struct {
	Type        UartEventTypeT
	Size        c.SizeT
	TimeoutFlag bool
}
type UartIsrHandleT IntrHandleT

/**
 * @brief Install UART driver and set the UART to the default configuration.
 *
 * UART ISR handler will be attached to the same CPU core that this function is running on.
 *
 * @note  Rx_buffer_size should be greater than UART_HW_FIFO_LEN(uart_num). Tx_buffer_size should be either zero or greater than UART_HW_FIFO_LEN(uart_num).
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param rx_buffer_size UART RX ring buffer size.
 * @param tx_buffer_size UART TX ring buffer size.
 *        If set to zero, driver will not use TX buffer, TX function will block task until all data have been sent out.
 * @param queue_size UART event queue size/depth.
 * @param uart_queue UART event queue handle (out param). On success, a new queue handle is written here to provide
 *        access to UART events. If set to NULL, driver will not use an event queue.
 * @param intr_alloc_flags Flags used to allocate the interrupt. One or multiple (ORred)
 *        ESP_INTR_FLAG_* values. See esp_intr_alloc.h for more info. Do not set ESP_INTR_FLAG_IRAM here
 *        (the driver's ISR handler is not located in IRAM)
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartDriverInstall C.uart_driver_install
func UartDriverInstall(uart_num UartPortT, rx_buffer_size c.Int, tx_buffer_size c.Int, queue_size c.Int, uart_queue *QueueHandleT, intr_alloc_flags c.Int) EspErrT

/**
 * @brief Uninstall UART driver.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartDriverDelete C.uart_driver_delete
func UartDriverDelete(uart_num UartPortT) EspErrT

/**
 * @brief Checks whether the driver is installed or not
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 *
 * @return
 *     - true  driver is installed
 *     - false driver is not installed
 */
//go:linkname UartIsDriverInstalled C.uart_is_driver_installed
func UartIsDriverInstalled(uart_num UartPortT) bool

/**
 * @brief Set UART data bits.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param data_bit UART data bits
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartSetWordLength C.uart_set_word_length
func UartSetWordLength(uart_num UartPortT, data_bit UartWordLengthT) EspErrT

/**
 * @brief Get the UART data bit configuration.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param data_bit Pointer to accept value of UART data bits.
 *
 * @return
 *     - ESP_FAIL  Parameter error
 *     - ESP_OK    Success, result will be put in (*data_bit)
 */
//go:linkname UartGetWordLength C.uart_get_word_length
func UartGetWordLength(uart_num UartPortT, data_bit *UartWordLengthT) EspErrT

/**
 * @brief Set UART stop bits.
 *
 * @param uart_num  UART port number, the max port number is (UART_NUM_MAX -1).
 * @param stop_bits  UART stop bits
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Fail
 */
//go:linkname UartSetStopBits C.uart_set_stop_bits
func UartSetStopBits(uart_num UartPortT, stop_bits UartStopBitsT) EspErrT

/**
 * @brief Get the UART stop bit configuration.
 *
 * @param uart_num  UART port number, the max port number is (UART_NUM_MAX -1).
 * @param stop_bits  Pointer to accept value of UART stop bits.
 *
 * @return
 *     - ESP_FAIL Parameter error
 *     - ESP_OK   Success, result will be put in (*stop_bit)
 */
//go:linkname UartGetStopBits C.uart_get_stop_bits
func UartGetStopBits(uart_num UartPortT, stop_bits *UartStopBitsT) EspErrT

/**
 * @brief Set UART parity mode.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param parity_mode the enum of uart parity configuration
 *
 * @return
 *     - ESP_FAIL  Parameter error
 *     - ESP_OK    Success
 */
//go:linkname UartSetParity C.uart_set_parity
func UartSetParity(uart_num UartPortT, parity_mode UartParityT) EspErrT

/**
 * @brief Get the UART parity mode configuration.
 *
 * @param uart_num  UART port number, the max port number is (UART_NUM_MAX -1).
 * @param parity_mode Pointer to accept value of UART parity mode.
 *
 * @return
 *     - ESP_FAIL  Parameter error
 *     - ESP_OK    Success, result will be put in (*parity_mode)
 *
 */
//go:linkname UartGetParity C.uart_get_parity
func UartGetParity(uart_num UartPortT, parity_mode *UartParityT) EspErrT

/**
 * @brief Get the frequency of a clock source for the HP UART port
 *
 * @param sclk Clock source
 * @param[out] out_freq_hz Output of frequency, in Hz
 *
 * @return
 *  - ESP_ERR_INVALID_ARG: if the clock source is not supported
 *  - otherwise ESP_OK
 */
//go:linkname UartGetSclkFreq C.uart_get_sclk_freq
func UartGetSclkFreq(sclk UartSclkT, out_freq_hz *c.Uint32T) EspErrT

/**
 * @brief Set desired UART baud rate.
 *
 * Note that the actual baud rate set could have a slight deviation from the user-configured value due to rounding error.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param baudrate UART baud rate.
 *
 * @return
 *     - ESP_FAIL Parameter error, such as baud rate unachievable
 *     - ESP_OK   Success
 */
//go:linkname UartSetBaudrate C.uart_set_baudrate
func UartSetBaudrate(uart_num UartPortT, baudrate c.Uint32T) EspErrT

/**
 * @brief Get the actual UART baud rate.
 *
 * It returns the real UART rate set in the hardware. It could have a slight deviation from the user-configured baud rate.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param baudrate Pointer to accept value of UART baud rate
 *
 * @return
 *     - ESP_FAIL Parameter error
 *     - ESP_OK   Success, result will be put in (*baudrate)
 *
 */
//go:linkname UartGetBaudrate C.uart_get_baudrate
func UartGetBaudrate(uart_num UartPortT, baudrate *c.Uint32T) EspErrT

/**
 * @brief Set UART line inverse mode
 *
 * @param uart_num  UART port number, the max port number is (UART_NUM_MAX -1).
 * @param inverse_mask Choose the wires that need to be inverted. Using the ORred mask of `uart_signal_inv_t`
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartSetLineInverse C.uart_set_line_inverse
func UartSetLineInverse(uart_num UartPortT, inverse_mask c.Uint32T) EspErrT

/**
 * @brief Set hardware flow control.
 *
 * @param uart_num   UART port number, the max port number is (UART_NUM_MAX -1).
 * @param flow_ctrl Hardware flow control mode
 * @param rx_thresh Threshold of Hardware RX flow control (0 ~ UART_HW_FIFO_LEN(uart_num)).
 *        Only when UART_HW_FLOWCTRL_RTS is set, will the rx_thresh value be set.
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartSetHwFlowCtrl C.uart_set_hw_flow_ctrl
func UartSetHwFlowCtrl(uart_num UartPortT, flow_ctrl UartHwFlowcontrolT, rx_thresh c.Uint8T) EspErrT

/**
 * @brief Set software flow control.
 *
 * @param uart_num   UART port number, the max port number is (UART_NUM_MAX -1)
 * @param enable     switch on or off
 * @param rx_thresh_xon  low water mark
 * @param rx_thresh_xoff high water mark
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartSetSwFlowCtrl C.uart_set_sw_flow_ctrl
func UartSetSwFlowCtrl(uart_num UartPortT, enable bool, rx_thresh_xon c.Uint8T, rx_thresh_xoff c.Uint8T) EspErrT

/**
 * @brief Get the UART hardware flow control configuration.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param flow_ctrl Option for different flow control mode.
 *
 * @return
 *     - ESP_FAIL Parameter error
 *     - ESP_OK   Success, result will be put in (*flow_ctrl)
 */
//go:linkname UartGetHwFlowCtrl C.uart_get_hw_flow_ctrl
func UartGetHwFlowCtrl(uart_num UartPortT, flow_ctrl *UartHwFlowcontrolT) EspErrT

/**
 * @brief Clear UART interrupt status
 *
 * @param uart_num  UART port number, the max port number is (UART_NUM_MAX -1).
 * @param clr_mask  Bit mask of the interrupt status to be cleared.
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartClearIntrStatus C.uart_clear_intr_status
func UartClearIntrStatus(uart_num UartPortT, clr_mask c.Uint32T) EspErrT

/**
 * @brief Set UART interrupt enable
 *
 * @param uart_num     UART port number, the max port number is (UART_NUM_MAX -1).
 * @param enable_mask  Bit mask of the enable bits.
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartEnableIntrMask C.uart_enable_intr_mask
func UartEnableIntrMask(uart_num UartPortT, enable_mask c.Uint32T) EspErrT

/**
 * @brief Clear UART interrupt enable bits
 *
 * @param uart_num      UART port number, the max port number is (UART_NUM_MAX -1).
 * @param disable_mask  Bit mask of the disable bits.
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartDisableIntrMask C.uart_disable_intr_mask
func UartDisableIntrMask(uart_num UartPortT, disable_mask c.Uint32T) EspErrT

/**
 * @brief Enable UART RX interrupt (RX_FULL & RX_TIMEOUT INTERRUPT)
 *
 * @param uart_num  UART port number, the max port number is (UART_NUM_MAX -1).
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartEnableRxIntr C.uart_enable_rx_intr
func UartEnableRxIntr(uart_num UartPortT) EspErrT

/**
 * @brief Disable UART RX interrupt (RX_FULL & RX_TIMEOUT INTERRUPT)
 *
 * @param uart_num  UART port number, the max port number is (UART_NUM_MAX -1).
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartDisableRxIntr C.uart_disable_rx_intr
func UartDisableRxIntr(uart_num UartPortT) EspErrT

/**
 * @brief Disable UART TX interrupt (TX_FULL & TX_TIMEOUT INTERRUPT)
 *
 * @param uart_num  UART port number
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartDisableTxIntr C.uart_disable_tx_intr
func UartDisableTxIntr(uart_num UartPortT) EspErrT

/**
 * @brief Enable UART TX interrupt (TX_FULL & TX_TIMEOUT INTERRUPT)
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param enable  1: enable; 0: disable
 * @param thresh  Threshold of TX interrupt, 0 ~ UART_HW_FIFO_LEN(uart_num)
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartEnableTxIntr C.uart_enable_tx_intr
func UartEnableTxIntr(uart_num UartPortT, enable c.Int, thresh c.Int) EspErrT

/**
 * @brief Assign signals of a UART peripheral to GPIO pins
 *
 * @note If the GPIO number configured for a UART signal matches one of the
 *       IOMUX signals for that GPIO, the signal will be connected directly
 *       via the IOMUX. Otherwise the GPIO and signal will be connected via
 *       the GPIO Matrix. For example, if on an ESP32 the call
 *       `uart_set_pin(0, 1, 3, -1, -1)` is performed, as GPIO1 is UART0's
 *       default TX pin and GPIO3 is UART0's default RX pin, both will be
 *       connected to respectively U0TXD and U0RXD through the IOMUX, totally
 *       bypassing the GPIO matrix.
 *       The check is performed on a per-pin basis. Thus, it is possible to have
 *       RX pin binded to a GPIO through the GPIO matrix, whereas TX is binded
 *       to its GPIO through the IOMUX.
 *
 * @note It is possible to configure TX and RX to share the same IO (single wire mode),
 *       but please be aware of output conflict, which could damage the pad.
 *       Apply open-drain and pull-up to the pad ahead of time as a protection,
 *       or the upper layer protocol must guarantee no output from two ends at the same time.
 *
 * @param uart_num   UART port number, the max port number is (UART_NUM_MAX -1).
 * @param tx_io_num  UART TX pin GPIO number.
 * @param rx_io_num  UART RX pin GPIO number.
 * @param rts_io_num UART RTS pin GPIO number.
 * @param cts_io_num UART CTS pin GPIO number.
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartSetPin C.uart_set_pin
func UartSetPin(uart_num UartPortT, tx_io_num c.Int, rx_io_num c.Int, rts_io_num c.Int, cts_io_num c.Int) EspErrT

/**
 * @brief Manually set the UART RTS pin level.
 * @note  UART must be configured with hardware flow control disabled.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param level    1: RTS output low (active); 0: RTS output high (block)
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartSetRts C.uart_set_rts
func UartSetRts(uart_num UartPortT, level c.Int) EspErrT

/**
 * @brief Manually set the UART DTR pin level.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param level    1: DTR output low; 0: DTR output high
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartSetDtr C.uart_set_dtr
func UartSetDtr(uart_num UartPortT, level c.Int) EspErrT

/**
 * @brief Set UART idle interval after tx FIFO is empty
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param idle_num idle interval after tx FIFO is empty(unit: the time it takes to send one bit
 *        under current baudrate)
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartSetTxIdleNum C.uart_set_tx_idle_num
func UartSetTxIdleNum(uart_num UartPortT, idle_num c.Uint16T) EspErrT

/**
 * @brief Set UART configuration parameters.
 *
 * @param uart_num    UART port number, the max port number is (UART_NUM_MAX -1).
 * @param uart_config UART parameter settings
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error, such as baud rate unachievable
 */
//go:linkname UartParamConfig C.uart_param_config
func UartParamConfig(uart_num UartPortT, uart_config *UartConfigT) EspErrT

/**
 * @brief Configure UART interrupts.
 *
 * @param uart_num  UART port number, the max port number is (UART_NUM_MAX -1).
 * @param intr_conf UART interrupt settings
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartIntrConfig C.uart_intr_config
func UartIntrConfig(uart_num UartPortT, intr_conf *UartIntrConfigT) EspErrT

/**
 * @brief Wait until UART TX FIFO is empty.
 *
 * @param uart_num      UART port number, the max port number is (UART_NUM_MAX -1).
 * @param ticks_to_wait Timeout, count in RTOS ticks
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_FAIL Parameter error
 *     - ESP_ERR_TIMEOUT  Timeout
 */
//go:linkname UartWaitTxDone C.uart_wait_tx_done
func UartWaitTxDone(uart_num UartPortT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Send data to the UART port from a given buffer and length.
 *
 * This function will not wait for enough space in TX FIFO. It will just fill the available TX FIFO and return when the FIFO is full.
 * @note This function should only be used when UART TX buffer is not enabled.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param buffer data buffer address
 * @param len    data length to send
 *
 * @return
 *     - (-1)  Parameter error
 *     - OTHERS (>=0) The number of bytes pushed to the TX FIFO
 */
//go:linkname UartTxChars C.uart_tx_chars
func UartTxChars(uart_num UartPortT, buffer *c.Char, len c.Uint32T) c.Int

/**
 * @brief Send data to the UART port from a given buffer and length,
 *
 * If the UART driver's parameter 'tx_buffer_size' is set to zero:
 * This function will not return until all the data have been sent out, or at least pushed into TX FIFO.
 *
 * Otherwise, if the 'tx_buffer_size' > 0, this function will return after copying all the data to tx ring buffer,
 * UART ISR will then move data from the ring buffer to TX FIFO gradually.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param src   data buffer address
 * @param size  data length to send
 *
 * @return
 *     - (-1) Parameter error
 *     - OTHERS (>=0) The number of bytes pushed to the TX FIFO
 */
//go:linkname UartWriteBytes C.uart_write_bytes
func UartWriteBytes(uart_num UartPortT, src c.Pointer, size c.SizeT) c.Int

/**
 * @brief Send data to the UART port from a given buffer and length,
 *
 * If the UART driver's parameter 'tx_buffer_size' is set to zero:
 * This function will not return until all the data and the break signal have been sent out.
 * After all data is sent out, send a break signal.
 *
 * Otherwise, if the 'tx_buffer_size' > 0, this function will return after copying all the data to tx ring buffer,
 * UART ISR will then move data from the ring buffer to TX FIFO gradually.
 * After all data sent out, send a break signal.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param src   data buffer address
 * @param size  data length to send
 * @param brk_len break signal duration(unit: the time it takes to send one bit at current baudrate)
 *
 * @return
 *     - (-1) Parameter error
 *     - OTHERS (>=0) The number of bytes pushed to the TX FIFO
 */
//go:linkname UartWriteBytesWithBreak C.uart_write_bytes_with_break
func UartWriteBytesWithBreak(uart_num UartPortT, src c.Pointer, size c.SizeT, brk_len c.Int) c.Int

/**
 * @brief UART read bytes from UART buffer
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param buf     pointer to the buffer.
 * @param length  data length
 * @param ticks_to_wait sTimeout, count in RTOS ticks
 *
 * @return
 *     - (-1) Error
 *     - OTHERS (>=0) The number of bytes read from UART buffer
 */
//go:linkname UartReadBytes C.uart_read_bytes
func UartReadBytes(uart_num UartPortT, buf c.Pointer, length c.Uint32T, ticks_to_wait TickTypeT) c.Int

/**
 * @brief Alias of uart_flush_input.
 *        UART ring buffer flush. This will discard all data in the UART RX buffer.
 * @note  Instead of waiting the data sent out, this function will clear UART rx buffer.
 *        In order to send all the data in tx FIFO, we can use uart_wait_tx_done function.
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartFlush C.uart_flush
func UartFlush(uart_num UartPortT) EspErrT

/**
 * @brief Clear input buffer, discard all the data is in the ring-buffer.
 * @note  In order to send all the data in tx FIFO, we can use uart_wait_tx_done function.
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartFlushInput C.uart_flush_input
func UartFlushInput(uart_num UartPortT) EspErrT

/**
 * @brief   UART get RX ring buffer cached data length
 *
 * @param   uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param   size Pointer of size_t to accept cached data length
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartGetBufferedDataLen C.uart_get_buffered_data_len
func UartGetBufferedDataLen(uart_num UartPortT, size *c.SizeT) EspErrT

/**
 * @brief   UART get TX ring buffer free space size
 *
 * @param   uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param   size Pointer of size_t to accept the free space size
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname UartGetTxBufferFreeSize C.uart_get_tx_buffer_free_size
func UartGetTxBufferFreeSize(uart_num UartPortT, size *c.SizeT) EspErrT

/**
 * @brief   UART disable pattern detect function.
 *          Designed for applications like 'AT commands'.
 *          When the hardware detects a series of one same character, the interrupt will be triggered.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartDisablePatternDetIntr C.uart_disable_pattern_det_intr
func UartDisablePatternDetIntr(uart_num UartPortT) EspErrT

/**
 * @brief UART enable pattern detect function.
 *        Designed for applications like 'AT commands'.
 *        When the hardware detect a series of one same character, the interrupt will be triggered.
 *
 * @param uart_num UART port number.
 * @param pattern_chr character of the pattern.
 * @param chr_num number of the character, 8bit value.
 * @param chr_tout timeout of the interval between each pattern characters, 16bit value, unit is the baud-rate cycle you configured.
 *        When the duration is more than this value, it will not take this data as at_cmd char.
 * @param post_idle idle time after the last pattern character, 16bit value, unit is the baud-rate cycle you configured.
 *        When the duration is less than this value, it will not take the previous data as the last at_cmd char
 * @param pre_idle idle time before the first pattern character, 16bit value, unit is the baud-rate cycle you configured.
 *        When the duration is less than this value, it will not take this data as the first at_cmd char.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_FAIL Parameter error
 */
//go:linkname UartEnablePatternDetBaudIntr C.uart_enable_pattern_det_baud_intr
func UartEnablePatternDetBaudIntr(uart_num UartPortT, pattern_chr c.Char, chr_num c.Uint8T, chr_tout c.Int, post_idle c.Int, pre_idle c.Int) EspErrT

/**
 * @brief Return the nearest detected pattern position in buffer.
 *        The positions of the detected pattern are saved in a queue,
 *        this function will dequeue the first pattern position and move the pointer to next pattern position.
 * @note  If the RX buffer is full and flow control is not enabled,
 *        the detected pattern may not be found in the rx buffer due to overflow.
 *
 *        The following APIs will modify the pattern position info:
 *        uart_flush_input, uart_read_bytes, uart_driver_delete, uart_pop_pattern_pos
 *        It is the application's responsibility to ensure atomic access to the pattern queue and the rx data buffer
 *        when using pattern detect feature.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @return
 *     - (-1) No pattern found for current index or parameter error
 *     - others the pattern position in rx buffer.
 */
//go:linkname UartPatternPopPos C.uart_pattern_pop_pos
func UartPatternPopPos(uart_num UartPortT) c.Int

/**
 * @brief Return the nearest detected pattern position in buffer.
 *        The positions of the detected pattern are saved in a queue,
 *        This function do nothing to the queue.
 * @note  If the RX buffer is full and flow control is not enabled,
 *        the detected pattern may not be found in the rx buffer due to overflow.
 *
 *        The following APIs will modify the pattern position info:
 *        uart_flush_input, uart_read_bytes, uart_driver_delete, uart_pop_pattern_pos
 *        It is the application's responsibility to ensure atomic access to the pattern queue and the rx data buffer
 *        when using pattern detect feature.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @return
 *     - (-1) No pattern found for current index or parameter error
 *     - others the pattern position in rx buffer.
 */
//go:linkname UartPatternGetPos C.uart_pattern_get_pos
func UartPatternGetPos(uart_num UartPortT) c.Int

/**
 * @brief Allocate a new memory with the given length to save record the detected pattern position in rx buffer.
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1).
 * @param queue_length Max queue length for the detected pattern.
 *        If the queue length is not large enough, some pattern positions might be lost.
 *        Set this value to the maximum number of patterns that could be saved in data buffer at the same time.
 * @return
 *     - ESP_ERR_NO_MEM No enough memory
 *     - ESP_ERR_INVALID_STATE Driver not installed
 *     - ESP_FAIL Parameter error
 *     - ESP_OK Success
 */
//go:linkname UartPatternQueueReset C.uart_pattern_queue_reset
func UartPatternQueueReset(uart_num UartPortT, queue_length c.Int) EspErrT

/**
 * @brief UART set communication mode
 *
 * @note  This function must be executed after uart_driver_install(), when the driver object is initialized.
 * @param uart_num     Uart number to configure, the max port number is (UART_NUM_MAX -1).
 * @param mode UART    UART mode to set
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname UartSetMode C.uart_set_mode
func UartSetMode(uart_num UartPortT, mode UartModeT) EspErrT

/**
 * @brief Set uart threshold value for RX fifo full
 * @note If application is using higher baudrate and it is observed that bytes
 *       in hardware RX fifo are overwritten then this threshold can be reduced
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1)
 * @param threshold Threshold value above which RX fifo full interrupt is generated
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Driver is not installed
 */
//go:linkname UartSetRxFullThreshold C.uart_set_rx_full_threshold
func UartSetRxFullThreshold(uart_num UartPortT, threshold c.Int) EspErrT

/**
 * @brief Set uart threshold values for TX fifo empty
 *
 * @param uart_num UART port number, the max port number is (UART_NUM_MAX -1)
 * @param threshold Threshold value below which TX fifo empty interrupt is generated
 *
 * @return
 *     - ESP_OK   Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Driver is not installed
 */
//go:linkname UartSetTxEmptyThreshold C.uart_set_tx_empty_threshold
func UartSetTxEmptyThreshold(uart_num UartPortT, threshold c.Int) EspErrT

/**
 * @brief UART set threshold timeout for TOUT feature
 *
 * @param uart_num     Uart number to configure, the max port number is (UART_NUM_MAX -1).
 * @param tout_thresh  This parameter defines timeout threshold in uart symbol periods. The maximum value of threshold is 126.
 *        tout_thresh = 1, defines TOUT interrupt timeout equal to transmission time of one symbol (~11 bit) on current baudrate.
 *        If the time is expired the UART_RXFIFO_TOUT_INT interrupt is triggered. If tout_thresh == 0,
 *        the TOUT feature is disabled.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE Driver is not installed
 */
//go:linkname UartSetRxTimeout C.uart_set_rx_timeout
func UartSetRxTimeout(uart_num UartPortT, tout_thresh c.Uint8T) EspErrT

/**
 * @brief Returns collision detection flag for RS485 mode
 *        Function returns the collision detection flag into variable pointed by collision_flag.
 *        *collision_flag = true, if collision detected else it is equal to false.
 *        This function should be executed when actual transmission is completed (after uart_write_bytes()).
 *
 * @param uart_num  Uart number to configure the max port number is (UART_NUM_MAX -1).
 * @param collision_flag Pointer to variable of type bool to return collision flag.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname UartGetCollisionFlag C.uart_get_collision_flag
func UartGetCollisionFlag(uart_num UartPortT, collision_flag *bool) EspErrT

/**
 * @brief Set the number of RX pin signal edges for light sleep wakeup
 *
 * UART can be used to wake up the system from light sleep. This feature works
 * by counting the number of positive edges on RX pin and comparing the count to
 * the threshold. When the count exceeds the threshold, system is woken up from
 * light sleep. This function allows setting the threshold value.
 *
 * Stop bit and parity bits (if enabled) also contribute to the number of edges.
 * For example, letter 'a' with ASCII code 97 is encoded as 0100001101 on the wire
 * (with 8n1 configuration), start and stop bits included. This sequence has 3
 * positive edges (transitions from 0 to 1). Therefore, to wake up the system
 * when 'a' is sent, set wakeup_threshold=3.
 *
 * The character that triggers wakeup is not received by UART (i.e. it can not
 * be obtained from UART FIFO). Depending on the baud rate, a few characters
 * after that will also not be received. Note that when the chip enters and exits
 * light sleep mode, APB frequency will be changing. To ensure that UART has
 * correct Baud rate all the time, it is necessary to select a source clock which has
 * a fixed frequency and remains active during sleep. For the supported clock sources
 * of the chips, please refer to `uart_sclk_t` or `soc_periph_uart_clk_src_legacy_t`
 *
 * @note in ESP32, the wakeup signal can only be input via IO_MUX (i.e.
 *       GPIO3 should be configured as function_1 to wake up UART0,
 *       GPIO9 should be configured as function_5 to wake up UART1), UART2
 *       does not support light sleep wakeup feature.
 *
 * @param uart_num  UART number, the max port number is (UART_NUM_MAX -1).
 * @param wakeup_threshold  number of RX edges for light sleep wakeup, value is 3 .. 0x3ff.
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if uart_num is incorrect or wakeup_threshold is
 *        outside of [3, 0x3ff] range.
 */
//go:linkname UartSetWakeupThreshold C.uart_set_wakeup_threshold
func UartSetWakeupThreshold(uart_num UartPortT, wakeup_threshold c.Int) EspErrT

/**
 * @brief Get the number of RX pin signal edges for light sleep wakeup.
 *
 * See description of uart_set_wakeup_threshold for the explanation of UART
 * wakeup feature.
 *
 * @param uart_num  UART number, the max port number is (UART_NUM_MAX -1).
 * @param[out] out_wakeup_threshold  output, set to the current value of wakeup
 *                                   threshold for the given UART.
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if out_wakeup_threshold is NULL
 */
//go:linkname UartGetWakeupThreshold C.uart_get_wakeup_threshold
func UartGetWakeupThreshold(uart_num UartPortT, out_wakeup_threshold *c.Int) EspErrT

/**
 * @brief Wait until UART tx memory empty and the last char send ok (polling mode).
 *
 * @param uart_num UART number
 *
 * * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG Parameter error
 *      - ESP_FAIL Driver not installed
 */
//go:linkname UartWaitTxIdlePolling C.uart_wait_tx_idle_polling
func UartWaitTxIdlePolling(uart_num UartPortT) EspErrT

/**
 * @brief Configure TX signal loop back to RX module, just for the test usage.
 *
 * @param uart_num UART number
 * @param loop_back_en Set true to enable the loop back function, else set it false.
 *
 * * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG Parameter error
 *      - ESP_FAIL Driver not installed
 */
//go:linkname UartSetLoopBack C.uart_set_loop_back
func UartSetLoopBack(uart_num UartPortT, loop_back_en bool) EspErrT

/**
 * @brief Configure behavior of UART RX timeout interrupt.
 *
 * When always_rx_timeout is true, timeout interrupt is triggered even if FIFO is full.
 * This function can cause extra timeout interrupts triggered only to send the timeout event.
 * Call this function only if you want to ensure timeout interrupt will always happen after a byte stream.
 *
 * @param uart_num UART number
 * @param always_rx_timeout_en Set to false enable the default behavior of timeout interrupt,
 *                             set it to true to always trigger timeout interrupt.
 *
 */
//go:linkname UartSetAlwaysRxTimeout C.uart_set_always_rx_timeout
func UartSetAlwaysRxTimeout(uart_num UartPortT, always_rx_timeout_en bool)
