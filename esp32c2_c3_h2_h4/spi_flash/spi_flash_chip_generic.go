package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Generic probe function
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param flash_id expected manufacture id.
 *
 * @return ESP_OK if the id read from chip->drv_read_id matches (always).
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericProbe C.spi_flash_chip_generic_probe
func (recv_ *EspFlashT) SpiFlashChipGenericProbe(flash_id c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Generic reset function
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 *
 * @return ESP_OK if sending success, or error code passed from ``common_command`` or ``wait_idle`` functions of host driver.
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericReset C.spi_flash_chip_generic_reset
func (recv_ *EspFlashT) SpiFlashChipGenericReset() EspErrT {
	return 0
}

/**
 * @brief Generic size detection function
 *
 * Tries to detect the size of chip by using the lower 4 bits of the chip->drv->read_id result = N, and assuming size is 2 ^ N.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param size Output of the detected size
 *
 * @return
 *      - ESP_OK if success
 *      - ESP_ERR_FLASH_UNSUPPORTED_CHIP if the manufacturer id is not correct, which may means an error in the reading
 *      - or other error passed from the ``read_id`` function of host driver
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericDetectSize C.spi_flash_chip_generic_detect_size
func (recv_ *EspFlashT) SpiFlashChipGenericDetectSize(size *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Erase chip by using the generic erase chip command.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 *
 * @return
 *      - ESP_OK if success
 *      - ESP_ERR_NOT_SUPPORTED if the chip is not able to perform the operation. This is indicated by WREN = 1 after the command is sent.
 *      - or other error passed from the ``set_write_protect``, ``wait_idle`` or ``erase_chip`` function of host driver
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericEraseChip C.spi_flash_chip_generic_erase_chip
func (recv_ *EspFlashT) SpiFlashChipGenericEraseChip() EspErrT {
	return 0
}

/**
 * @brief Erase sector by using the generic sector erase command.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param start_address Start address of the sector to erase
 *
 * @return
 *      - ESP_OK if success
 *      - ESP_ERR_NOT_SUPPORTED if the chip is not able to perform the operation. This is indicated by WREN = 1 after the command is sent.
 *      - or other error passed from the ``set_write_protect``, ``wait_idle`` or ``erase_sector`` function of host driver
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericEraseSector C.spi_flash_chip_generic_erase_sector
func (recv_ *EspFlashT) SpiFlashChipGenericEraseSector(start_address c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Erase block by the generic 64KB block erase command
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param start_address Start address of the block to erase
 *
 * @return
 *      - ESP_OK if success
 *      - ESP_ERR_NOT_SUPPORTED if the chip is not able to perform the operation. This is indicated by WREN = 1 after the command is sent.
 *      - or other error passed from the ``set_write_protect``, ``wait_idle`` or ``erase_block`` function of host driver
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericEraseBlock C.spi_flash_chip_generic_erase_block
func (recv_ *EspFlashT) SpiFlashChipGenericEraseBlock(start_address c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Read from flash by using a read command that matches the programmed
 * read mode.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param buffer Buffer to hold the data read from flash
 * @param address Start address of the data on the flash
 * @param length Length to read
 *
 * @return always ESP_OK currently
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericRead C.spi_flash_chip_generic_read
func (recv_ *EspFlashT) SpiFlashChipGenericRead(buffer c.Pointer, address c.Uint32T, length c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Perform a page program using the page program command.
 *
 * @note Length of each call should not excced the limitation in
 * ``chip->host->max_write_bytes``. This function is called in
 * ``spi_flash_chip_generic_write`` recursively until the whole page is
 * programmed. Strongly suggest to call ``spi_flash_chip_generic_write``
 * instead.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param buffer Buffer holding the data to program
 * @param address Start address to write to flash
 * @param length Length to write, no longer than ``chip->host->max_write_bytes``.
 *
 * @return
 *      - ESP_OK if success
 *      - ESP_ERR_NOT_SUPPORTED if the chip is not able to perform the operation. This is indicated by WREN = 1 after the command is sent.
 *      - or other error passed from the ``wait_idle`` or ``program_page`` function of host driver
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericPageProgram C.spi_flash_chip_generic_page_program
func (recv_ *EspFlashT) SpiFlashChipGenericPageProgram(buffer c.Pointer, address c.Uint32T, length c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Perform a generic write. Split the write buffer into page program
 * operations, and call chip->chip_drv->page-program() for each.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param buffer Buffer holding the data to program
 * @param address Start address to write to flash
 * @param length Length to write
 *
 * @return
 *      - ESP_OK if success
 *      - or other error passed from the ``wait_idle`` or ``program_page`` function of host driver
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericWrite C.spi_flash_chip_generic_write
func (recv_ *EspFlashT) SpiFlashChipGenericWrite(buffer c.Pointer, address c.Uint32T, length c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Perform a write using on-chip flash encryption. Not implemented yet.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param buffer Buffer holding the data to program
 * @param address Start address to write to flash
 * @param length Length to write
 *
 * @return always ESP_ERR_FLASH_UNSUPPORTED_HOST.
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericWriteEncrypted C.spi_flash_chip_generic_write_encrypted
func (recv_ *EspFlashT) SpiFlashChipGenericWriteEncrypted(buffer c.Pointer, address c.Uint32T, length c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Send the write enable or write disable command and verify the expected bit (1) in
 * the status register is set.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param write_protect true to enable write protection, false to send write enable.
 *
 * @return
 *      - ESP_OK if success
 *      - or other error passed from the ``wait_idle``, ``read_status`` or
 *        ``set_write_protect`` function of host driver
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericSetWriteProtect C.spi_flash_chip_generic_set_write_protect
func (recv_ *EspFlashT) SpiFlashChipGenericSetWriteProtect(write_protect bool) EspErrT {
	return 0
}

/**
 * @brief Check whether WEL (write enable latch) bit is set in the Status Register read from RDSR.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param out_write_protect Output of whether the write protect is set.
 *
 * @return
 *      - ESP_OK if success
 *      - or other error passed from the ``read_status`` function of host driver
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericGetWriteProtect C.spi_flash_chip_generic_get_write_protect
func (recv_ *EspFlashT) SpiFlashChipGenericGetWriteProtect(out_write_protect *bool) EspErrT {
	return 0
}

/**
 * @brief Send commands to read one of the reg of the chip
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param reg_id     Type of the register to read
 * @param out_reg    Output of the register value
 * @return esp_err_t Error code passed from the ``read_status`` function of host driver.
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericReadReg C.spi_flash_chip_generic_read_reg
func (recv_ *EspFlashT) SpiFlashChipGenericReadReg(reg_id SpiFlashRegisterT, out_reg *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Read flash status via the RDSR command and wait for bit 0 (write in
 * progress bit) to be cleared.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param timeout_us Time to wait before timeout, in us.
 *
 * @return
 *      - ESP_OK if success
 *      - ESP_ERR_TIMEOUT if not idle before timeout
 *      - or other error passed from the ``wait_idle`` or ``read_status`` function of host driver
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericWaitIdle C.spi_flash_chip_generic_wait_idle
func (recv_ *EspFlashT) SpiFlashChipGenericWaitIdle(timeout_us c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Set the specified SPI read mode according to the data in the chip
 *        context. Set quad enable status register bit if needed.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 *
 * @return
 *      - ESP_OK if success
*      - ESP_ERR_TIMEOUT if not idle before timeout
 *      - or other error passed from the ``set_write_protect`` or ``common_command`` function of host driver
*/
// llgo:link (*EspFlashT).SpiFlashChipGenericSetIoMode C.spi_flash_chip_generic_set_io_mode
func (recv_ *EspFlashT) SpiFlashChipGenericSetIoMode() EspErrT {
	return 0
}

/**
 * Get whether the Quad Enable (QE) is set.
 *
* @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
* @param out_quad_mode Pointer to store the output mode.
*          - SPI_FLASH_QOUT: QE is enabled
*          - otherwise: QE is disabled
*
* @return
*      - ESP_OK if success
*      - or other error passed from the ``common_command`` function of host driver
*/
// llgo:link (*EspFlashT).SpiFlashChipGenericGetIoMode C.spi_flash_chip_generic_get_io_mode
func (recv_ *EspFlashT) SpiFlashChipGenericGetIoMode(out_quad_mode *EspFlashIoModeT) EspErrT {
	return 0
}

/**
 * @brief Read the chip unique ID.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param flash_unique_id Pointer to store output unique id.
 *
 * @return
 *      - ESP_OK if success
 *      - ESP_ERR_NOT_SUPPORTED if the chip doesn't support read id.
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericReadUniqueId C.spi_flash_chip_generic_read_unique_id
func (recv_ *EspFlashT) SpiFlashChipGenericReadUniqueId(flash_unique_id *c.Uint64T) EspErrT {
	return 0
}

// llgo:type C
type EspFlashRdsrFuncT func(*EspFlashT, *c.Uint32T) EspErrT

/**
 * Use RDSR2 (35H) to read bit 15-8 of the SR, and RDSR (05H) to read bit 7-0.
 *
 * @param chip Pointer to SPI flash chip to use.
 * @param out_sr Pointer to buffer to hold the status register, 16 bits.
 *
 * @return ESP_OK if success, otherwise error code passed from the
 *         `common_command` function of the host driver.
 */
// llgo:link (*EspFlashT).SpiFlashCommonReadStatus16bRdsrRdsr2 C.spi_flash_common_read_status_16b_rdsr_rdsr2
func (recv_ *EspFlashT) SpiFlashCommonReadStatus16bRdsrRdsr2(out_sr *c.Uint32T) EspErrT {
	return 0
}

/**
 * Use RDSR2 (35H) to read bit 15-8 of the SR.
 *
 * @param chip Pointer to SPI flash chip to use.
 * @param out_sr Pointer to buffer to hold the status register, 8 bits.
 *
 * @return ESP_OK if success, otherwise error code passed from the
 *         `common_command` function of the host driver.
 */
// llgo:link (*EspFlashT).SpiFlashCommonReadStatus8bRdsr2 C.spi_flash_common_read_status_8b_rdsr2
func (recv_ *EspFlashT) SpiFlashCommonReadStatus8bRdsr2(out_sr *c.Uint32T) EspErrT {
	return 0
}

/**
 * Use RDSR (05H) to read bit 7-0 of the SR.
 *
 * @param chip Pointer to SPI flash chip to use.
 * @param out_sr Pointer to buffer to hold the status register, 8 bits.
 *
 * @return ESP_OK if success, otherwise error code passed from the
 *         `common_command` function of the host driver.
 */
// llgo:link (*EspFlashT).SpiFlashCommonReadStatus8bRdsr C.spi_flash_common_read_status_8b_rdsr
func (recv_ *EspFlashT) SpiFlashCommonReadStatus8bRdsr(out_sr *c.Uint32T) EspErrT {
	return 0
}

// llgo:type C
type EspFlashWrsrFuncT func(*EspFlashT, c.Uint32T) EspErrT

/**
 * Use WRSR (01H) to write bit 7-0 of the SR.
 *
 * @param chip Pointer to SPI flash chip to use.
 * @param sr Value of the status register to write, 8 bits.
 *
 * @return ESP_OK if success, otherwise error code passed from the
 *         `common_command` function of the host driver.
 */
// llgo:link (*EspFlashT).SpiFlashCommonWriteStatus8bWrsr C.spi_flash_common_write_status_8b_wrsr
func (recv_ *EspFlashT) SpiFlashCommonWriteStatus8bWrsr(sr c.Uint32T) EspErrT {
	return 0
}

/**
 * Use WRSR (01H) to write bit 15-0 of the SR.
 *
 * @param chip Pointer to SPI flash chip to use.
 * @param sr Value of the status register to write, 16 bits.
 *
 * @return ESP_OK if success, otherwise error code passed from the
 *         `common_command` function of the host driver.
 */
// llgo:link (*EspFlashT).SpiFlashCommonWriteStatus16bWrsr C.spi_flash_common_write_status_16b_wrsr
func (recv_ *EspFlashT) SpiFlashCommonWriteStatus16bWrsr(sr c.Uint32T) EspErrT {
	return 0
}

/**
 * Use WRSR2 (31H) to write bit 15-8 of the SR.
 *
 * @param chip Pointer to SPI flash chip to use.
 * @param sr Value of the status register to write, 8 bits.
 *
 * @return ESP_OK if success, otherwise error code passed from the
 *         `common_command` function of the host driver.
 */
// llgo:link (*EspFlashT).SpiFlashCommonWriteStatus8bWrsr2 C.spi_flash_common_write_status_8b_wrsr2
func (recv_ *EspFlashT) SpiFlashCommonWriteStatus8bWrsr2(sr c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Utility function for set_read_mode chip_drv function. If required,
 * set and check the QE bit in the flash chip to enable the QIO/QOUT mode.
 *
 * Most chip QE enable follows a common pattern, though commands to read/write
 * the status register may be different, as well as the position of QE bit.
 *
 * Registers to actually do Quad transtions and command to be sent in reading
 * should also be configured via
 * spi_flash_chip_generic_config_host_io_mode().
 *
 * Note that the bit length and qe position of wrsr_func, rdsr_func and
 * qe_sr_bit should be consistent.
 *
 * @param chip Pointer to SPI flash chip to use.
 * @param wrsr_func Function pointer for writing the status register
 * @param rdsr_func Function pointer for reading the status register
 * @param qe_sr_bit status with the qe bit only.
 *
 * @return always ESP_OK (currently).
 */
// llgo:link (*EspFlashT).SpiFlashCommonSetIoMode C.spi_flash_common_set_io_mode
func (recv_ *EspFlashT) SpiFlashCommonSetIoMode(wrsr_func EspFlashWrsrFuncT, rdsr_func EspFlashRdsrFuncT, qe_sr_bit c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Configure the host registers to use the specified read mode set in
 *        the ``chip->read_mode``.
 *
 * Usually called in chip_drv read() functions before actual reading
 * transactions. Also prepare the command to be sent in read functions.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param flags Special rules to configure io mode, (i.e. Whether 32 bit commands will be used (Currently only W25Q256 and GD25Q256 are supported))
 *
 * @return
 *      - ESP_OK if success
 *      - ESP_ERR_FLASH_NOT_INITIALISED if chip not initialized properly
 *      - or other error passed from the ``configure_host_mode`` function of host driver
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericConfigHostIoMode C.spi_flash_chip_generic_config_host_io_mode
func (recv_ *EspFlashT) SpiFlashChipGenericConfigHostIoMode(flags c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Handle explicit yield requests
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @param wip  Write (erase) in progress, `true` if this function is called during waiting idle of a erase/write command; else `false`.
 * @return ESP_OK if success, otherwise failed.
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericYield C.spi_flash_chip_generic_yield
func (recv_ *EspFlashT) SpiFlashChipGenericYield(wip c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Setup for flash suspend command configuration.
 *
 * @param chip Pointer to SPI flash chip to use. If NULL, esp_flash_default_chip is substituted.
 * @return ESP_OK
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericSuspendCmdConf C.spi_flash_chip_generic_suspend_cmd_conf
func (recv_ *EspFlashT) SpiFlashChipGenericSuspendCmdConf() EspErrT {
	return 0
}

/**
 *
 * @brief Read the chip unique ID unsupported function.
 *
 * @param chip Pointer to SPI flash chip to use.
 * @param flash_unique_id Pointer to store output unique id (Although this function is an unsupported function, but the parameter should be kept for the consistence of the function pointer).
 * @return Always ESP_ERR_NOT_SUPPORTED.
 */
// llgo:link (*EspFlashT).SpiFlashChipGenericReadUniqueIdNone C.spi_flash_chip_generic_read_unique_id_none
func (recv_ *EspFlashT) SpiFlashChipGenericReadUniqueIdNone(flash_unique_id *c.Uint64T) EspErrT {
	return 0
}
