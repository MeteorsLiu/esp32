package fatfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Register spi flash partition
 *
 * @param pdrv  drive number
 * @param flash_handle  handle of the wear levelling partition.
 */
//go:linkname FfDiskioRegisterWlPartition C.ff_diskio_register_wl_partition
func FfDiskioRegisterWlPartition(pdrv c.Char, flash_handle WlHandleT) EspErrT

//go:linkname FfDiskioGetPdrvWl C.ff_diskio_get_pdrv_wl
func FfDiskioGetPdrvWl(flash_handle WlHandleT) c.Char

//go:linkname FfDiskioClearPdrvWl C.ff_diskio_clear_pdrv_wl
func FfDiskioClearPdrvWl(flash_handle WlHandleT)
