package fatfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const FF_DRV_NOT_USED = 0xFF

/**
 * Structure of pointers to disk IO driver functions.
 *
 * See FatFs documentation for details about these functions
 */

type FfDiskioImplT struct {
	Init   c.Pointer
	Status c.Pointer
	Read   c.Pointer
	Write  c.Pointer
	Ioctl  c.Pointer
}

/**
 * Register or unregister diskio driver for given drive number.
 *
 * When FATFS library calls one of disk_xxx functions for driver number pdrv,
 * corresponding function in discio_impl for given pdrv will be called.
 *
 * @param pdrv drive number
 * @param discio_impl   pointer to ff_diskio_impl_t structure with diskio functions
 *                      or NULL to unregister and free previously registered drive
 */
// llgo:link BYTE.FfDiskioRegister C.ff_diskio_register
func (recv_ BYTE) FfDiskioRegister(discio_impl *FfDiskioImplT) {
}

/**
 * Get next available drive number
 *
 * @param   out_pdrv            pointer to the byte to set if successful
 *
 * @return  ESP_OK              on success
 *          ESP_ERR_NOT_FOUND   if all drives are attached
 */
// llgo:link (*BYTE).FfDiskioGetDrive C.ff_diskio_get_drive
func (recv_ *BYTE) FfDiskioGetDrive() EspErrT {
	return 0
}
