package fatfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Enable/disable SD card status checking
 *
 * @param pdrv   drive number
 * @param enable mock ff_sdmmc_status function (return 0)
 */
// llgo:link BYTE.FfSdmmcSetDiskStatusCheck C.ff_sdmmc_set_disk_status_check
func (recv_ BYTE) FfSdmmcSetDiskStatusCheck(enable bool) {
}

/**
 * Register SD/MMC diskio driver
 *
 * @param pdrv  drive number
 * @param card  pointer to sdmmc_card_t structure describing a card; card should be initialized before calling f_mount.
 */
//go:linkname FfDiskioRegisterSdmmc C.ff_diskio_register_sdmmc
func FfDiskioRegisterSdmmc(pdrv c.Char, card *SdmmcCardT)

/**
 * @brief Get the driver number corresponding to a card
 *
 * @param card The card to get its driver
 * @return Driver number to the card
 */
//go:linkname FfDiskioGetPdrvCard C.ff_diskio_get_pdrv_card
func FfDiskioGetPdrvCard(card *SdmmcCardT) BYTE
