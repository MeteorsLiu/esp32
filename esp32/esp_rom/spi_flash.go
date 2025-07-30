package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SPI0_R_QIO_DUMMY_CYCLELEN = 3
const SPI0_R_QIO_ADDR_BITSLEN = 31
const SPI0_R_FAST_DUMMY_CYCLELEN = 7
const SPI0_R_DIO_DUMMY_CYCLELEN = 1
const SPI0_R_DIO_ADDR_BITSLEN = 27
const SPI0_R_FAST_ADDR_BITSLEN = 23
const SPI0_R_SIO_ADDR_BITSLEN = 23
const SPI1_R_QIO_DUMMY_CYCLELEN = 3
const SPI1_R_QIO_ADDR_BITSLEN = 31
const SPI1_R_FAST_DUMMY_CYCLELEN = 7
const SPI1_R_DIO_DUMMY_CYCLELEN = 3
const SPI1_R_DIO_ADDR_BITSLEN = 31
const SPI1_R_FAST_ADDR_BITSLEN = 23
const SPI1_R_SIO_ADDR_BITSLEN = 23
const ESP_ROM_SPIFLASH_W_SIO_ADDR_BITSLEN = 23
const ESP_ROM_SPIFLASH_BYTES_LEN = 24
const ESP_ROM_SPIFLASH_BUFF_BYTE_WRITE_NUM = 32
const ESP_ROM_SPIFLASH_BUFF_BYTE_READ_NUM = 64
const ESP_ROM_SPIFLASH_BUFF_BYTE_READ_BITS = 0x3f
const ESP_ROM_SPIFLASH_DUMMY_LEN_PLUS_20M = 0
const ESP_ROM_SPIFLASH_DUMMY_LEN_PLUS_26M = 0
const ESP_ROM_SPIFLASH_DUMMY_LEN_PLUS_40M = 1
const ESP_ROM_SPIFLASH_DUMMY_LEN_PLUS_80M = 2

type EspRomSpiflashCommonCmdT struct {
	DataLength c.Uint8T
	ReadCmd0   c.Uint8T
	ReadCmd1   c.Uint8T
	WriteCmd   c.Uint8T
	DataMask   c.Uint16T
	Data       c.Uint16T
}

/**
 * @brief Clear all SR bits except QE bit.
 *        Please do not call this function in SDK.
 *
 * @param  None.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Unlock OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Unlock error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Unlock timeout.
 */
//go:linkname EspRomSpiflashClearBp C.esp_rom_spiflash_clear_bp
func EspRomSpiflashClearBp() EspRomSpiflashResultT

/**
 * @brief Clear all SR bits except QE bit.
 *        Please do not call this function in SDK.
 *
 * @param  None.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Unlock OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Unlock error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Unlock timeout.
 */
//go:linkname EspRomSpiflashUnlock C.esp_rom_spiflash_unlock
func EspRomSpiflashUnlock() EspRomSpiflashResultT

/**
 * @brief SPI flash set BP0 to BP2.(Only valid when WRSR+2Bytes)
 *        Please do not call this function in SDK.
 *
 * @param  None.
 *
 * @return ESP_ROM_SPIFLASH_RESULT_OK : Lock OK.
 *         ESP_ROM_SPIFLASH_RESULT_ERR : Lock error.
 *         ESP_ROM_SPIFLASH_RESULT_TIMEOUT : Lock timeout.
 */
//go:linkname EspRomSpiflashLock C.esp_rom_spiflash_lock
func EspRomSpiflashLock() EspRomSpiflashResultT
