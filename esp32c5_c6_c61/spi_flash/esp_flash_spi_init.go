package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// / Configurations for the SPI Flash to init
type EspFlashSpiDeviceConfigT struct {
	HostId       SpiHostDeviceT
	CsIoNum      c.Int
	IoMode       EspFlashIoModeT
	Speed        EspFlashSpeedS
	InputDelayNs c.Int
	CsId         c.Int
	FreqMhz      c.Int
}

/**
 *  Add a SPI Flash device onto the SPI bus.
 *
 * The bus should be already initialized by ``spi_bus_initialization``.
 *
 * @param out_chip Pointer to hold the initialized chip.
 * @param config Configuration of the chips to initialize.
 *
 * @return
 *      - ESP_ERR_INVALID_ARG: out_chip is NULL, or some field in the config is invalid.
 *      - ESP_ERR_NO_MEM: failed to allocate memory for the chip structures.
 *      - ESP_OK: success.
 */
//go:linkname SpiBusAddFlashDevice C.spi_bus_add_flash_device
func SpiBusAddFlashDevice(out_chip **EspFlashT, config *EspFlashSpiDeviceConfigT) EspErrT

/**
 *  Remove a SPI Flash device from the SPI bus.
 *
 * @param chip The flash device to remove.
 *
 * @return
 *      - ESP_ERR_INVALID_ARG: The chip is invalid.
 *      - ESP_OK: success.
 */
// llgo:link (*EspFlashT).SpiBusRemoveFlashDevice C.spi_bus_remove_flash_device
func (recv_ *EspFlashT) SpiBusRemoveFlashDevice() EspErrT {
	return 0
}
