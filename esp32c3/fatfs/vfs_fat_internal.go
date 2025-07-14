package fatfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type VfsFatXCtxFlagsT c.Int

const FORMATTED_DURING_LAST_MOUNT VfsFatXCtxFlagsT = 1

type VfsFatSpiflashCtxT struct {
	Partition   *EspPartitionT
	ByLabel     bool
	Pdrv        BYTE
	Fs          *FATFS
	Wlhandle    WlHandleT
	MountConfig EspVfsFatMountConfigT
	Flags       VfsFatXCtxFlagsT
}

type VfsFatSdCtxT struct {
	Pdrv        BYTE
	MountConfig EspVfsFatMountConfigT
	Fs          *FATFS
	Card        *SdmmcCardT
	BasePath    *c.Char
	Flags       VfsFatXCtxFlagsT
}

//go:linkname GetVfsFatSpiflashCtx C.get_vfs_fat_spiflash_ctx
func GetVfsFatSpiflashCtx(wlhandle WlHandleT) *VfsFatSpiflashCtxT

//go:linkname GetVfsFatGetSdCtx C.get_vfs_fat_get_sd_ctx
func GetVfsFatGetSdCtx(card *SdmmcCardT) *VfsFatSdCtxT
