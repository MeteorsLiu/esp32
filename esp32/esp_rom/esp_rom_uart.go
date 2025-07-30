package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_ROM_CDC_ACM_WORK_BUF_MIN = 128

type EspRomUartNumT c.Int

const (
	ESP_ROM_UART_0   EspRomUartNumT = 0
	ESP_ROM_UART_1   EspRomUartNumT = 1
	ESP_ROM_UART_USB EspRomUartNumT = 2
)

/**
 * @brief Wait for UART TX FIFO is empty and all data has been sent out.
 *
 * @param serial_num The serial number defined in ROM, including UART_x, USB_OTG, USB_SERIAL_JTAG..
 */
//go:linkname EspRomOutputTxWaitIdle C.esp_rom_output_tx_wait_idle
func EspRomOutputTxWaitIdle(serial_num c.Uint8T)

/**
 * @brief Transmit one character to the console channel.
 * @note This function is a wrapper over esp_rom_uart_tx_one_char, it can help handle line ending issue by replacing '\n' with '\r\n'.
 *
 * @param c Character to send
 */
//go:linkname EspRomOutputPutc C.esp_rom_output_putc
func EspRomOutputPutc(c c.Char)
