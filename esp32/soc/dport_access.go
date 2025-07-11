package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Read value from register, SMP-safe version.
 *
 * This method uses the pre-reading of the APB register before reading the register of the DPORT.
 * This implementation is useful for reading DORT registers for single reading without stall other CPU.
 * There is disable/enable interrupt.
 *
 * @note Should be placed in IRAM.
 *
 * @param reg Register address
 * @return Value
 */
//go:linkname EspDportAccessRegRead C.esp_dport_access_reg_read
func EspDportAccessRegRead(reg c.Uint32T) c.Uint32T

/**
 * @brief Read value from register, NOT SMP-safe version.
 *
 * This method uses the pre-reading of the APB register before reading the register of the DPORT.
 * There is not disable/enable interrupt.
 * The difference from DPORT_REG_READ() is that the user himself must disable interrupts while DPORT reading.
 * This implementation is useful for reading DORT registers in loop without stall other CPU. Note the usage example.
 * The recommended way to read registers sequentially without stall other CPU
 * is to use the method esp_dport_read_buffer(buff_out, address, num_words). It allows you to read registers in the buffer.
 *
 * @note Should be placed in IRAM.
 *
 * \code{c}
 * // This example shows how to use it.
 * { // Use curly brackets to limit the visibility of variables in macros DPORT_INTERRUPT_DISABLE/RESTORE.
 *     DPORT_INTERRUPT_DISABLE(); // Disable interrupt only on current CPU.
 *     for (i = 0; i < max; ++i) {
 *        array[i] = esp_dport_access_sequence_reg_read(Address + i * 4); // reading DPORT registers
 *     }
 *     DPORT_INTERRUPT_RESTORE(); // restore the previous interrupt level
 * }
 * \endcode
 *
 * @param reg Register address
 * @return Value
 */
//go:linkname EspDportAccessSequenceRegRead C.esp_dport_access_sequence_reg_read
func EspDportAccessSequenceRegRead(reg c.Uint32T) c.Uint32T

/**
 * @brief Read a sequence of DPORT registers to the buffer, SMP-safe version.
 *
 * This implementation uses a method of the pre-reading of the APB register
 * before reading the register of the DPORT, without stall other CPU.
 * There is disable/enable interrupt.
 *
 * @note Should be placed in IRAM.
 *
 * @param[out] buff_out  Contains the read data.
 * @param[in]  address   Initial address for reading registers.
 * @param[in]  num_words The number of words.
 */
//go:linkname EspDportAccessReadBuffer C.esp_dport_access_read_buffer
func EspDportAccessReadBuffer(buff_out *c.Uint32T, address c.Uint32T, num_words c.Uint32T)
