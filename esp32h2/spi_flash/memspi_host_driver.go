package spi_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MemspiHostConfigT SpiFlashHalConfigT
type MemspiHostInstT SpiFlashHalContextT

/**
 * Initialize the memory SPI host.
 *
 * @param host Pointer to the host structure.
 * @param cfg Pointer to configuration structure
 *
 * @return always return ESP_OK
 */
// llgo:link (*MemspiHostInstT).MemspiHostInitPointers C.memspi_host_init_pointers
func (recv_ *MemspiHostInstT) MemspiHostInitPointers(cfg *MemspiHostConfigT) EspErrT {
	return 0
}

/**
 * @brief Read the Status Register read from RDSR (05h).
 *
 * High speed implementation of RDID through memspi interface relying on the
 * ``common_command``.
 *
 * @param host The driver context.
 * @param id Output of the read ID from the slave.
 *
 * @return
 *  - ESP_OK: if success
 *  - ESP_ERR_FLASH_NO_RESPONSE: if no response from chip
 *  - or other cases from ``spi_hal_common_command``
 */
//go:linkname MemspiHostReadIdHs C.memspi_host_read_id_hs
func MemspiHostReadIdHs(host *SpiFlashHostInstT, id *c.Uint32T) EspErrT

/**
 * High speed implementation of RDSR through memspi interface relying on the
 * ``common_command``.
 *
 * @param host The driver context.
 * @param id Output of the read ID from the slave.
 *
 * @return
 *  - ESP_OK: if success
 *  - or other cases from ``spi_hal_common_command``
 */
//go:linkname MemspiHostReadStatusHs C.memspi_host_read_status_hs
func MemspiHostReadStatusHs(host *SpiFlashHostInstT, out_sr *c.Uint8T) EspErrT

/**
 * Flush the cache (if needed) after the contents are modified.
 *
 * @param host The driver context.
 * @param addr Start address of the modified region
 * @param size Size of the region modified.
 *
 * @return always ESP_OK.
 */
//go:linkname MemspiHostFlushCache C.memspi_host_flush_cache
func MemspiHostFlushCache(host *SpiFlashHostInstT, addr c.Uint32T, size c.Uint32T) EspErrT

/**
 *  Erase contents of entire chip.
 *
 * @param host The driver context.
 */
//go:linkname MemspiHostEraseChip C.memspi_host_erase_chip
func MemspiHostEraseChip(host *SpiFlashHostInstT)

/**
 *  Erase a sector starting from a given address. For 24bit address only.
 *
 * @param host The driver context.
 * @param start_address Starting address of the sector.
 */
//go:linkname MemspiHostEraseSector C.memspi_host_erase_sector
func MemspiHostEraseSector(host *SpiFlashHostInstT, start_address c.Uint32T)

/**
 *  Erase a block starting from a given address. For 24bit address only.
 *
 * @param host The driver context.
 * @param start_address Starting address of the block.
 */
//go:linkname MemspiHostEraseBlock C.memspi_host_erase_block
func MemspiHostEraseBlock(host *SpiFlashHostInstT, start_address c.Uint32T)

/**
 * Program a page with contents of a buffer. For 24bit address only.
 *
 * @param host The driver context.
 * @param buffer Buffer which contains the data to be flashed.
 * @param address Starting address of where to flash the data.
 * @param length The number of bytes to flash.
 */
//go:linkname MemspiHostProgramPage C.memspi_host_program_page
func MemspiHostProgramPage(host *SpiFlashHostInstT, buffer c.Pointer, address c.Uint32T, length c.Uint32T)

/**
 * Set ability to write to chip.
 *
 * @param host The driver context.
 * @param wp Enable or disable write protect (true - enable, false - disable).
 */
//go:linkname MemspiHostSetWriteProtect C.memspi_host_set_write_protect
func MemspiHostSetWriteProtect(host *SpiFlashHostInstT, wp bool) EspErrT

/**
 * Read data to buffer.
 *
 * @param host The driver context.
 * @param buffer Buffer which contains the data to be read.
 * @param address Starting address of where to read the data.
 * @param length The number of bytes to read.
 */
//go:linkname MemspiHostRead C.memspi_host_read
func MemspiHostRead(host *SpiFlashHostInstT, buffer c.Pointer, address c.Uint32T, read_len c.Uint32T) EspErrT

/**
 * @brief Slicer for read data used in non-encrypted regions. This slicer does nothing but
 *        limit the length to the maximum size the host supports.
 *
 * @param address Flash address to read
 * @param len Length to read
 * @param align_address Output of the address to read, should be equal to the input `address`
 * @param page_size Physical SPI flash page size
 *
 * @return Length that can actually be read in one `read` call in `spi_flash_host_driver_t`.
 */
//go:linkname MemspiHostReadDataSlicer C.memspi_host_read_data_slicer
func MemspiHostReadDataSlicer(host *SpiFlashHostInstT, address c.Uint32T, len c.Uint32T, align_address *c.Uint32T, page_size c.Uint32T) c.Int

/**
 * @brief Slicer for write data used in non-encrypted regions. This slicer limit the length to the
 *        maximum size the host supports, and truncate if the write data lie accross the page boundary
 *        (256 bytes)
 *
 * @param address Flash address to write
 * @param len Length to write
 * @param align_address Output of the address to write, should be equal to the input `address`
 * @param page_size Physical SPI flash page size
 *
 * @return Length that can actually be written in one `program_page` call in `spi_flash_host_driver_t`.
 */
//go:linkname MemspiHostWriteDataSlicer C.memspi_host_write_data_slicer
func MemspiHostWriteDataSlicer(host *SpiFlashHostInstT, address c.Uint32T, len c.Uint32T, align_address *c.Uint32T, page_size c.Uint32T) c.Int
