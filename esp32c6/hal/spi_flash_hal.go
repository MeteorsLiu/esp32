package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SPI_FLASH_HAL_MAX_WRITE_BYTES = 64
const SPI_FLASH_HAL_MAX_READ_BYTES = 64

/**
 * Generic driver context structure for all chips using the SPI peripheral.
 * Include this into the HEAD of the driver data for other driver
 * implementations that also use the SPI peripheral.
 */

type SpiFlashHalContextT struct {
	Inst         SpiFlashHostInstT
	Spi          *SpiDevT
	CsNum        c.Int
	ClockConf    SpiFlashLlClockRegT
	BaseIoMode   EspFlashIoModeT
	Flags        c.Uint32T
	SusCfg       SpiFlashSusCmdConf
	SlicerFlags  c.Uint32T
	FreqMhz      c.Int
	TsusVal      c.Uint8T
	AutoWaitiPes bool
}

// / This struct provide MSPI Flash necessary timing related config, should be consistent with that in union in `spi_flash_hal_config_t`.
type SpiFlashHalTimingConfigT struct {
	ExtraDummy  c.Uint32T
	CsHold      c.Uint32T
	CsSetup     c.Uint8T
	ClockConfig SpiFlashLlClockRegT
}

// / Configuration structure for the SPI driver.
type SpiFlashHalConfigT struct {
	Iomux             bool
	InputDelayNs      c.Int
	Speed             EspFlashSpeedS
	HostId            SpiHostDeviceT
	CsNum             c.Int
	AutoSusEn         bool
	OctalModeEn       bool
	UsingTimingTuning bool
	DefaultIoMode     EspFlashIoModeT
	FreqMhz           c.Int
	ClockSrcFreq      c.Int
	TsusVal           c.Uint8T
	AutoWaitiPes      bool
}

/**
 * Configure SPI flash hal settings.
 *
 * @param data Buffer to hold configured data, the buffer should be in DRAM to be available when cache disabled
 * @param cfg Configurations to set
 *
 * @return
 *      - ESP_OK: success
 *      - ESP_ERR_INVALID_ARG: the data buffer is not in the DRAM.
 */
// llgo:link (*SpiFlashHalContextT).SpiFlashHalInit C.spi_flash_hal_init
func (recv_ *SpiFlashHalContextT) SpiFlashHalInit(cfg *SpiFlashHalConfigT) EspErrT {
	return 0
}

/**
 * Configure the device-related register before transactions.
 *
 * @param host The driver context.
 *
 * @return always return ESP_OK.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalDeviceConfig C.spi_flash_hal_device_config
func (recv_ *SpiFlashHostInstT) SpiFlashHalDeviceConfig() EspErrT {
	return 0
}

/**
 * Send an user-defined spi transaction to the device.
 *
 * @note This is usually used when the memspi interface doesn't support some
 *      particular commands. Since this function supports timing compensation, it is
 *      also used to receive some data when the frequency is high.
 *
 * @param host The driver context.
 * @param trans The transaction to send, also holds the received data.
 *
 * @return always return ESP_OK.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalCommonCommand C.spi_flash_hal_common_command
func (recv_ *SpiFlashHostInstT) SpiFlashHalCommonCommand(trans *SpiFlashTransT) EspErrT {
	return 0
}

/**
 * Erase whole flash chip by using the erase chip (C7h) command.
 *
 * @param host The driver context.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalEraseChip C.spi_flash_hal_erase_chip
func (recv_ *SpiFlashHostInstT) SpiFlashHalEraseChip() {
}

/**
 * Erase a specific sector by its start address through the sector erase (20h)
 * command. For 24bit address only.
 *
 * @param host The driver context.
 * @param start_address Start address of the sector to erase.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalEraseSector C.spi_flash_hal_erase_sector
func (recv_ *SpiFlashHostInstT) SpiFlashHalEraseSector(start_address c.Uint32T) {
}

/**
 * Erase a specific 64KB block by its start address through the 64KB block
 * erase (D8h) command. For 24bit address only.
 *
 * @param host The driver context.
 * @param start_address Start address of the block to erase.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalEraseBlock C.spi_flash_hal_erase_block
func (recv_ *SpiFlashHostInstT) SpiFlashHalEraseBlock(start_address c.Uint32T) {
}

/**
 * Program a page of the flash using the page program (02h) command. For 24bit address only.
 *
 * @param host The driver context.
 * @param address Address of the page to program
 * @param buffer Data to program
 * @param length Size of the buffer in bytes, no larger than ``SPI_FLASH_HAL_MAX_WRITE_BYTES`` (64) bytes.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalProgramPage C.spi_flash_hal_program_page
func (recv_ *SpiFlashHostInstT) SpiFlashHalProgramPage(buffer c.Pointer, address c.Uint32T, length c.Uint32T) {
}

/**
 * Read from the flash. Call ``spi_flash_hal_configure_host_read_mode`` to
 * configure the read command before calling this function.
 *
 * @param host The driver context.
 * @param buffer Buffer to store the read data
 * @param address Address to read
 * @param length Length to read, no larger than ``SPI_FLASH_HAL_MAX_READ_BYTES`` (64) bytes.
 *
 * @return always return ESP_OK.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalRead C.spi_flash_hal_read
func (recv_ *SpiFlashHostInstT) SpiFlashHalRead(buffer c.Pointer, address c.Uint32T, read_len c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Send the write enable (06h) or write disable (04h) command to the flash chip.
 *
 * @param driver The driver context.
 * @param wp true to enable the write protection, otherwise false.
 *
 * @return always return ESP_OK.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalSetWriteProtect C.spi_flash_hal_set_write_protect
func (recv_ *SpiFlashHostInstT) SpiFlashHalSetWriteProtect(wp bool) EspErrT {
	return 0
}

/**
 * Check whether the SPI host is idle and can perform other operations.
 *
 * @param host The driver context.
 *
 * @return 0:busy, 1:idle, 2:suspended.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalCheckStatus C.spi_flash_hal_check_status
func (recv_ *SpiFlashHostInstT) SpiFlashHalCheckStatus() c.Uint32T {
	return 0
}

/**
 * @brief Configure the SPI host hardware registers for the specified io mode.
 *
 *  Note that calling this configures SPI host registers, so if running any
 *  other commands as part of set_io_mode() then these must be run before
 *  calling this function.
 *
 *  The command value, address length and dummy cycles are configured according
 *  to the format of read commands:
 *
 *  - command: 8 bits, value set.
 *  - address: 24 bits
 *  - dummy: cycles to compensate the input delay
 *  - out & in data: 0 bits.
 *
 *  The following commands still need to:
 *
 *  - Read data: set address value and data (length and contents), no need
 *    to touch command and dummy phases.
 *  - Common read: set command value, address value (or length to 0 if not used)
 *  - Common write: set command value, address value (or length to 0 if not
 *    used), disable dummy phase, and set output data.
 *
 * @param host The driver context
 * @param io_mode The HW read mode to use
 * @param addr_bitlen Length of the address phase, in bits
 * @param dummy_cyclelen_base Base cycles of the dummy phase, some extra dummy cycles may be appended to compensate the timing.
 * @param command  Actual reading command to send to flash chip on the bus.
 *
 * @return always return ESP_OK.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalConfigureHostIoMode C.spi_flash_hal_configure_host_io_mode
func (recv_ *SpiFlashHostInstT) SpiFlashHalConfigureHostIoMode(command c.Uint32T, addr_bitlen c.Uint32T, dummy_cyclelen_base c.Int, io_mode EspFlashIoModeT) EspErrT {
	return 0
}

/**
 * Poll until the last operation is done.
 *
 * @param host The driver context.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalPollCmdDone C.spi_flash_hal_poll_cmd_done
func (recv_ *SpiFlashHostInstT) SpiFlashHalPollCmdDone() {
}

/**
 * Check whether the given buffer can be used as the write buffer directly. If 'chip' is connected to the main SPI bus, we can only write directly from
 * regions that are accessible ith cache disabled. *
 *
 * @param host The driver context
 * @param p The buffer holding data to send.
 *
 * @return True if the buffer can be used to send data, otherwise false.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalSupportsDirectWrite C.spi_flash_hal_supports_direct_write
func (recv_ *SpiFlashHostInstT) SpiFlashHalSupportsDirectWrite(p c.Pointer) bool {
	return false
}

/**
 * Check whether the given buffer can be used as the read buffer directly. If 'chip' is connected to the main SPI bus, we can only read directly from
 * regions that are accessible ith cache disabled. *
 *
 * @param host The driver context
 * @param p The buffer to hold the received data.
 *
 * @return True if the buffer can be used to receive data, otherwise false.
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalSupportsDirectRead C.spi_flash_hal_supports_direct_read
func (recv_ *SpiFlashHostInstT) SpiFlashHalSupportsDirectRead(p c.Pointer) bool {
	return false
}

/**
 * @brief Resume flash chip status from suspend.
 *
 * @param host The driver context.
 *
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalResume C.spi_flash_hal_resume
func (recv_ *SpiFlashHostInstT) SpiFlashHalResume() {
}

/**
 * @brief Set the flash into suspend status manually.
 *
 * @param host The driver context.
 *
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalSuspend C.spi_flash_hal_suspend
func (recv_ *SpiFlashHostInstT) SpiFlashHalSuspend() {
}

/**
 * To setup for reading flash suspend status register
 *
 * @param host The driver context.
 * @param sus_conf Flash chip suspend feature configuration, mainly for command config, may vary from chip to chip.
 *
 * @return Always ESP_OK
 */
// llgo:link (*SpiFlashHostInstT).SpiFlashHalSetupReadSuspend C.spi_flash_hal_setup_read_suspend
func (recv_ *SpiFlashHostInstT) SpiFlashHalSetupReadSuspend(sus_conf *SpiFlashSusCmdConf) EspErrT {
	return 0
}
