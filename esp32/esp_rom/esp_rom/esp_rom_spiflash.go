package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspRomSpiflashReadModeT c.Int

const (
	ESP_ROM_SPIFLASH_QIO_MODE         EspRomSpiflashReadModeT = 0
	ESP_ROM_SPIFLASH_QOUT_MODE        EspRomSpiflashReadModeT = 1
	ESP_ROM_SPIFLASH_DIO_MODE         EspRomSpiflashReadModeT = 2
	ESP_ROM_SPIFLASH_DOUT_MODE        EspRomSpiflashReadModeT = 3
	ESP_ROM_SPIFLASH_FASTRD_MODE      EspRomSpiflashReadModeT = 4
	ESP_ROM_SPIFLASH_SLOWRD_MODE      EspRomSpiflashReadModeT = 5
	ESP_ROM_SPIFLASH_OPI_STR_MODE     EspRomSpiflashReadModeT = 6
	ESP_ROM_SPIFLASH_OPI_DTR_MODE     EspRomSpiflashReadModeT = 7
	ESP_ROM_SPIFLASH_OOUT_MODE        EspRomSpiflashReadModeT = 8
	ESP_ROM_SPIFLASH_OIO_STR_MODE     EspRomSpiflashReadModeT = 9
	ESP_ROM_SPIFLASH_OIO_DTR_MODE     EspRomSpiflashReadModeT = 10
	ESP_ROM_SPIFLASH_QPI_MODE         EspRomSpiflashReadModeT = 11
	ESP_ROM_SPIFLASH_OPI_HEX_DTR_MODE EspRomSpiflashReadModeT = 12
)

type EspRomSpiflashChipT struct {
	DeviceId   c.Uint32T
	ChipSize   c.Uint32T
	BlockSize  c.Uint32T
	SectorSize c.Uint32T
	PageSize   c.Uint32T
	StatusMask c.Uint32T
}
type EspRomSpiflashResultT c.Int

const (
	ESP_ROM_SPIFLASH_RESULT_OK      EspRomSpiflashResultT = 0
	ESP_ROM_SPIFLASH_RESULT_ERR     EspRomSpiflashResultT = 1
	ESP_ROM_SPIFLASH_RESULT_TIMEOUT EspRomSpiflashResultT = 2
)

/**
 * @brief SPI Read Flash status register. We use CMD 0x05 (RDSR).
 *    Please do not call this function in SDK.
 *
 * @param  esp_rom_spiflash_chip_t *spi : The information for Flash, which is exported from ld file.
 *
 * @param  uint32_t *status : The pointer to which to return the Flash status value.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : read OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : read error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : read timeout.
 */
// llgo:link (*EspRomSpiflashChipT).EspRomSpiflashReadStatus C.esp_rom_spiflash_read_status
func (recv_ *EspRomSpiflashChipT) EspRomSpiflashReadStatus(status *c.Uint32T) EspRomSpiflashResultT {
	return 0
}

/**
 * @brief SPI Read Flash status register bits 8-15. We use CMD 0x35 (RDSR2).
 *        Please do not call this function in SDK.
 *
 * @param  esp_rom_spiflash_chip_t *spi : The information for Flash, which is exported from ld file.
 *
 * @param  uint32_t *status : The pointer to which to return the Flash status value.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : read OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : read error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : read timeout.
 */
// llgo:link (*EspRomSpiflashChipT).EspRomSpiflashReadStatushigh C.esp_rom_spiflash_read_statushigh
func (recv_ *EspRomSpiflashChipT) EspRomSpiflashReadStatushigh(status *c.Uint32T) EspRomSpiflashResultT {
	return 0
}

/**
 * @brief Write status to Flash status register.
 *        Please do not call this function in SDK.
 *
 * @param  esp_rom_spiflash_chip_t *spi : The information for Flash, which is exported from ld file.
 *
 * @param  uint32_t status_value : Value to .
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : write OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : write error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : write timeout.
 */
// llgo:link (*EspRomSpiflashChipT).EspRomSpiflashWriteStatus C.esp_rom_spiflash_write_status
func (recv_ *EspRomSpiflashChipT) EspRomSpiflashWriteStatus(status_value c.Uint32T) EspRomSpiflashResultT {
	return 0
}

/**
 * @brief Use a command to Read Flash status register.
 *        Please do not call this function in SDK.
 *
 * @param  esp_rom_spiflash_chip_t *spi : The information for Flash, which is exported from ld file.
 *
 * @param  uint32_t*status : The pointer to which to return the Flash status value.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : read OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : read error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : read timeout.
 */
//go:linkname EspRomSpiflashReadUserCmd C.esp_rom_spiflash_read_user_cmd
func EspRomSpiflashReadUserCmd(status *c.Uint32T, cmd c.Uint8T) EspRomSpiflashResultT

/**
 * @brief Config SPI Flash read mode when init.
 *        Please do not call this function in SDK.
 *
 * @param  esp_rom_spiflash_read_mode_t mode : QIO/QOUT/DIO/DOUT/FastRD/SlowRD.
 *
 * This function does not try to set the QIO Enable bit in the status register, caller is responsible for this.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : config OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : config error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : config timeout.
 */
// llgo:link EspRomSpiflashReadModeT.EspRomSpiflashConfigReadmode C.esp_rom_spiflash_config_readmode
func (recv_ EspRomSpiflashReadModeT) EspRomSpiflashConfigReadmode() EspRomSpiflashResultT {
	return 0
}

/**
 * @brief Erase whole flash chip.
 *        Please do not call this function in SDK.
 *
 * @param  None
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Erase OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Erase error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Erase timeout.
 */
//go:linkname EspRomSpiflashEraseChip C.esp_rom_spiflash_erase_chip
func EspRomSpiflashEraseChip() EspRomSpiflashResultT

/**
 * @brief Erase a 64KB block of flash
 *        Uses SPI flash command D8H.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t block_num : Which block to erase.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Erase OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Erase error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Erase timeout.
 */
//go:linkname EspRomSpiflashEraseBlock C.esp_rom_spiflash_erase_block
func EspRomSpiflashEraseBlock(block_num c.Uint32T) EspRomSpiflashResultT

/**
 * @brief Erase a sector of flash.
 *        Uses SPI flash command 20H.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t sector_num : Which sector to erase.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Erase OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Erase error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Erase timeout.
 */
//go:linkname EspRomSpiflashEraseSector C.esp_rom_spiflash_erase_sector
func EspRomSpiflashEraseSector(sector_num c.Uint32T) EspRomSpiflashResultT

/**
 * @brief Erase some sectors.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t start_addr : Start addr to erase, should be sector aligned.
 *
 * @param  uint32_t area_len : Length to erase, should be sector aligned.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Erase OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Erase error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Erase timeout.
 */
//go:linkname EspRomSpiflashEraseArea C.esp_rom_spiflash_erase_area
func EspRomSpiflashEraseArea(start_addr c.Uint32T, area_len c.Uint32T) EspRomSpiflashResultT

/**
 * @brief Write Data to Flash, you should Erase it yourself if need.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t dest_addr : Address to write, should be 4 bytes aligned.
 *
 * @param  const uint32_t *src : The pointer to data which is to write.
 *
 * @param  uint32_t len : Length to write, should be 4 bytes aligned.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Write OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Write error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Write timeout.
 */
//go:linkname EspRomSpiflashWrite C.esp_rom_spiflash_write
func EspRomSpiflashWrite(dest_addr c.Uint32T, src *c.Uint32T, len c.Int32T) EspRomSpiflashResultT

/**
 * @brief Read Data from Flash, you should Erase it yourself if need.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t src_addr : Address to read, should be 4 bytes aligned.
 *
 * @param  uint32_t *dest : The buf to read the data.
 *
 * @param  uint32_t len : Length to read, should be 4 bytes aligned.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Read OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Read error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Read timeout.
 */
//go:linkname EspRomSpiflashRead C.esp_rom_spiflash_read
func EspRomSpiflashRead(src_addr c.Uint32T, dest *c.Uint32T, len c.Int32T) EspRomSpiflashResultT

/**
 * @brief SPI1 go into encrypto mode.
 *        Please do not call this function in SDK.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname EspRomSpiflashWriteEncryptedEnable C.esp_rom_spiflash_write_encrypted_enable
func EspRomSpiflashWriteEncryptedEnable()

/**
 * @brief Prepare 32 Bytes data to encrpto writing, you should Erase it yourself if need.
 *        Please do not call this function in SDK.
 *
 * @param  uint32_t flash_addr : Address to write, should be 32 bytes aligned.
 *
 * @param  uint32_t *data : The pointer to data which is to write.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Prepare OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Prepare error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Prepare timeout.
 */
//go:linkname EspRomSpiflashPrepareEncryptedData C.esp_rom_spiflash_prepare_encrypted_data
func EspRomSpiflashPrepareEncryptedData(flash_addr c.Uint32T, data *c.Uint32T) EspRomSpiflashResultT

/**
 * @brief SPI1 go out of encrypto mode.
 *        Please do not call this function in SDK.
 *
 * @param  None
 *
 * @return None
 */
//go:linkname EspRomSpiflashWriteEncryptedDisable C.esp_rom_spiflash_write_encrypted_disable
func EspRomSpiflashWriteEncryptedDisable()

/**
 * @brief Write data to flash with transparent encryption.
 * @note Sectors to be written should already be erased.
 *
 * @note Please do not call this function in SDK.
 *
 * @param  uint32_t flash_addr : Address to write, should be 32 byte aligned.
 *
 * @param  uint32_t *data : The pointer to data to write. Note, this pointer must
 *                          be 32 bit aligned and the content of the data will be
 *                          modified by the encryption function.
 *
 * @param  uint32_t len : Length to write, should be 32 bytes aligned.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Data written successfully.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Encryption write error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Encrypto write timeout.
 */
//go:linkname EspRomSpiflashWriteEncrypted C.esp_rom_spiflash_write_encrypted
func EspRomSpiflashWriteEncrypted(flash_addr c.Uint32T, data *c.Uint32T, len c.Uint32T) EspRomSpiflashResultT

/** @brief Wait until SPI flash write operation is complete
 *
 * @note Please do not call this function in SDK.
 *
 * Reads the Write In Progress bit of the SPI flash status register,
 * repeats until this bit is zero (indicating write complete).
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Write is complete
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Error while reading status.
 */
// llgo:link (*EspRomSpiflashChipT).EspRomSpiflashWaitIdle C.esp_rom_spiflash_wait_idle
func (recv_ *EspRomSpiflashChipT) EspRomSpiflashWaitIdle() EspRomSpiflashResultT {
	return 0
}

/**
 * @brief Clear WEL bit unconditionally.
 *
 * @return always ESP_ROM_SPIFLASH_RESULT_OK
 */
//go:linkname EspRomSpiflashWriteDisable C.esp_rom_spiflash_write_disable
func EspRomSpiflashWriteDisable() EspRomSpiflashResultT
