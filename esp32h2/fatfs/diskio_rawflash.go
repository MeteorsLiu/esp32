package fatfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Register spi flash partition
 *
 * @param pdrv  drive number
 * @param part_handle  pointer to raw flash partition.
 */
//go:linkname FfDiskioRegisterRawPartition C.ff_diskio_register_raw_partition
func FfDiskioRegisterRawPartition(pdrv c.Char, part_handle *EspPartitionT) EspErrT

//go:linkname FfDiskioGetPdrvRaw C.ff_diskio_get_pdrv_raw
func FfDiskioGetPdrvRaw(part_handle *EspPartitionT) c.Char
