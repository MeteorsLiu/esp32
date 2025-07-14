package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspMspiIoT c.Int

const (
	ESP_MSPI_IO_CLK EspMspiIoT = 0
	ESP_MSPI_IO_Q   EspMspiIoT = 1
	ESP_MSPI_IO_D   EspMspiIoT = 2
	ESP_MSPI_IO_CS0 EspMspiIoT = 3
	ESP_MSPI_IO_HD  EspMspiIoT = 4
	ESP_MSPI_IO_WP  EspMspiIoT = 5
	ESP_MSPI_IO_MAX EspMspiIoT = 6
)

/**
 * @brief To setup Flash chip
 */
//go:linkname SpiFlashInitChipState C.spi_flash_init_chip_state
func SpiFlashInitChipState() EspErrT

/**
 * @brief To initislize the MSPI pins
 */
//go:linkname EspMspiPinInit C.esp_mspi_pin_init
func EspMspiPinInit()

/**
 * @brief Reserve MSPI IOs
 */
//go:linkname EspMspiPinReserve C.esp_mspi_pin_reserve
func EspMspiPinReserve()

/**
 * @brief Get the number of the GPIO corresponding to the given MSPI io
 *
 * @param[in] io  MSPI io
 *
 * @return MSPI IO number
 */
// llgo:link EspMspiIoT.EspMspiGetIo C.esp_mspi_get_io
func (recv_ EspMspiIoT) EspMspiGetIo() c.Uint8T {
	return 0
}

/**
 * @brief Set SPI1 registers to make ROM functions work
 * @note This function is used for setting SPI1 registers to the state that ROM SPI functions work
 */
//go:linkname SpiFlashSetRomRequiredRegs C.spi_flash_set_rom_required_regs
func SpiFlashSetRomRequiredRegs()

/**
 * @brief Initialize main flash
 * @param chip Pointer to main SPI flash(SPI1 CS0) chip to use..
 */
// llgo:link (*EspFlashT).EspFlashInitMain C.esp_flash_init_main
func (recv_ *EspFlashT) EspFlashInitMain() EspErrT {
	return 0
}

/**
 * @brief Judge whether need to reset flash when brownout.
 *        Set` flash_brownout_needs_reset` inside the function if really need reset.
 */
//go:linkname SpiFlashNeedsResetCheck C.spi_flash_needs_reset_check
func SpiFlashNeedsResetCheck()

/**
 * @brief Set flag to reset flash. set when erase chip or program chip
 *
 * @param bool status. True if flash is eraing. False if flash is not erasing.
 *
 * @return None.
 */
//go:linkname SpiFlashSetErasingFlag C.spi_flash_set_erasing_flag
func SpiFlashSetErasingFlag(status bool)

/**
 * @brief Judge whether need to reset flash when brownout.
 *
 * @return true if need reset, otherwise false.
 */
//go:linkname SpiFlashBrownoutNeedReset C.spi_flash_brownout_need_reset
func SpiFlashBrownoutNeedReset() bool

/**
 * @brief Check whether esp-chip supports 32bit address properly
 *
 * @return ESP_OK for supported, ESP_ERR_NOT_SUPPORTED for not supported
 */
//go:linkname EspMspi32bitAddressFlashFeatureCheck C.esp_mspi_32bit_address_flash_feature_check
func EspMspi32bitAddressFlashFeatureCheck() EspErrT

/**
 * @brief set wrap size of flash
 *
 * @param wrap_size: wrap mode support disable, 16 32, 64 byte
 *
 * @return esp_err_t : ESP_OK for successful.
 *
 */
// llgo:link SpiFlashWrapSizeT.SpiFlashWrapEnable C.spi_flash_wrap_enable
func (recv_ SpiFlashWrapSizeT) SpiFlashWrapEnable() EspErrT {
	return 0
}

/**
 * @brief Probe flash wrap method
 *
 * @return esp_err_t: ESP_OK for success
 */
//go:linkname SpiFlashWrapProbe C.spi_flash_wrap_probe
func SpiFlashWrapProbe() EspErrT

/**
 * @brief disable cache wrap
 */
//go:linkname SpiFlashWrapDisable C.spi_flash_wrap_disable
func SpiFlashWrapDisable() EspErrT

/**
 * @brief Check whether flash and esp chip supports wrap mode.
 *
 * @param wrap_size wrap size.
 * @return true: wrap support, otherwise, false.
 */
//go:linkname SpiFlashSupportWrapSize C.spi_flash_support_wrap_size
func SpiFlashSupportWrapSize(wrap_size c.Uint32T) bool

// llgo:type C
type SpiFlashGuardStartFuncT func()

// llgo:type C
type SpiFlashGuardEndFuncT func()

/**
 * Structure holding SPI flash access critical sections management functions.
 *
 * Flash API uses two types of flash access management functions:
 * 1) Functions which prepare/restore flash cache and interrupts before calling
 *    appropriate ROM functions (SPIWrite, SPIRead and SPIEraseBlock):
 *   - 'start' function should disables flash cache and non-IRAM interrupts and
 *      is invoked before the call to one of ROM function above.
 *   - 'end' function should restore state of flash cache and non-IRAM interrupts and
 *      is invoked after the call to one of ROM function above.
 *    These two functions are not recursive.
 *
 * Different versions of the guarding functions should be used depending on the context of
 * execution (with or without functional OS). In normal conditions when flash API is called
 * from task the functions use OS primitives. When there is no OS at all or when
 * it is not guaranteed that OS is functional (accessing flash from exception handler) these
 * functions cannot use OS primitives or even does not need them (multithreaded access is not possible).
 *
 * @note Structure and corresponding guard functions should not reside in flash.
 *       For example structure can be placed in DRAM and functions in IRAM sections.
 */

type SpiFlashGuardFuncsT struct {
	Start SpiFlashGuardStartFuncT
	End   SpiFlashGuardEndFuncT
}

/**
 * @brief  Sets guard functions to access flash.
 *
 * @note Pointed structure and corresponding guard functions should not reside in flash.
 *       For example structure can be placed in DRAM and functions in IRAM sections.
 *
 * @param funcs pointer to structure holding flash access guard functions.
 */
// llgo:link (*SpiFlashGuardFuncsT).SpiFlashGuardSet C.spi_flash_guard_set
func (recv_ *SpiFlashGuardFuncsT) SpiFlashGuardSet() {
}

/**
 * @brief Get the guard functions used for flash access
 *
 * @return The guard functions that were set via spi_flash_guard_set(). These functions
 * can be called if implementing custom low-level SPI flash operations.
 */
//go:linkname SpiFlashGuardGet C.spi_flash_guard_get
func SpiFlashGuardGet() *SpiFlashGuardFuncsT
