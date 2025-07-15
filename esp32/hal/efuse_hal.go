package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief get rated frequency in MHz
 */
//go:linkname EfuseHalGetRatedFreqMhz C.efuse_hal_get_rated_freq_mhz
func EfuseHalGetRatedFreqMhz() c.Uint32T

/**
 * @brief set eFuse timings
 *
 * @param apb_freq_mhz APB frequency in MHz
 */
//go:linkname EfuseHalSetTiming C.efuse_hal_set_timing
func EfuseHalSetTiming(apb_freq_mhz c.Uint32T)

/**
 * @brief trigger eFuse read operation
 */
//go:linkname EfuseHalRead C.efuse_hal_read
func EfuseHalRead()

/**
 * @brief clear registers for programming eFuses
 */
//go:linkname EfuseHalClearProgramRegisters C.efuse_hal_clear_program_registers
func EfuseHalClearProgramRegisters()

/**
 * @brief burn eFuses written in programming registers (all blocks at once)
 *
 * @param block not used
 */
//go:linkname EfuseHalProgram C.efuse_hal_program
func EfuseHalProgram(block c.Uint32T)

/**
 * @brief Checks coding error in a block
 *
 * @param block Index of efuse block
 *
 * @return True  - block has an error.
 *         False - no error.
 */
//go:linkname EfuseHalIsCodingErrorInBlock C.efuse_hal_is_coding_error_in_block
func EfuseHalIsCodingErrorInBlock(block c.Uint) bool
