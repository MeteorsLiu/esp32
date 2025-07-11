package vfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspVfsIdT c.Int

/**
 * @brief VFS semaphore type for select()
 *
 */

type EspVfsSelectSemT struct {
	IsSemLocal bool
	Sem        c.Pointer
}

// llgo:type C
type EspVfsStatCtxOpT func(c.Pointer, *c.Char, *Stat) c.Int

// llgo:type C
type EspVfsStatOpT func(*c.Char, *Stat) c.Int

// llgo:type C
type EspVfsLinkCtxOpT func(c.Pointer, *c.Char, *c.Char) c.Int

// llgo:type C
type EspVfsLinkOpT func(*c.Char, *c.Char) c.Int

// llgo:type C
type EspVfsUnlinkCtxOpT func(c.Pointer, *c.Char) c.Int

// llgo:type C
type EspVfsUnlinkOpT func(*c.Char) c.Int

// llgo:type C
type EspVfsRenameCtxOpT func(c.Pointer, *c.Char, *c.Char) c.Int

// llgo:type C
type EspVfsRenameOpT func(*c.Char, *c.Char) c.Int

// llgo:type C
type EspVfsOpendirCtxOpT func(c.Pointer, *c.Char) *DIR

// llgo:type C
type EspVfsOpendirOpT func(*c.Char) *DIR

// llgo:type C
type EspVfsReaddirCtxOpT func(c.Pointer, *DIR) *Dirent

// llgo:type C
type EspVfsReaddirOpT func(*DIR) *Dirent

// llgo:type C
type EspVfsReaddirRCtxOpT func(c.Pointer, *DIR, *Dirent, **Dirent) c.Int

// llgo:type C
type EspVfsReaddirROpT func(*DIR, *Dirent, **Dirent) c.Int

// llgo:type C
type EspVfsTelldirCtxOpT func(c.Pointer, *DIR) c.Long

// llgo:type C
type EspVfsTelldirOpT func(*DIR) c.Long

// llgo:type C
type EspVfsSeekdirCtxOpT func(c.Pointer, *DIR, c.Long)

// llgo:type C
type EspVfsSeekdirOpT func(*DIR, c.Long)

// llgo:type C
type EspVfsClosedirCtxOpT func(c.Pointer, *DIR) c.Int

// llgo:type C
type EspVfsClosedirOpT func(*DIR) c.Int

// llgo:type C
type EspVfsMkdirCtxOpT func(c.Pointer, *c.Char, ModeT) c.Int

// llgo:type C
type EspVfsMkdirOpT func(*c.Char, ModeT) c.Int

// llgo:type C
type EspVfsRmdirCtxOpT func(c.Pointer, *c.Char) c.Int

// llgo:type C
type EspVfsRmdirOpT func(*c.Char) c.Int

// llgo:type C
type EspVfsAccessCtxOpT func(c.Pointer, *c.Char, c.Int) c.Int

// llgo:type C
type EspVfsAccessOpT func(*c.Char, c.Int) c.Int

// llgo:type C
type EspVfsTruncateCtxOpT func(c.Pointer, *c.Char, OffT) c.Int

// llgo:type C
type EspVfsTruncateOpT func(*c.Char, OffT) c.Int

// llgo:type C
type EspVfsFtruncateCtxOpT func(c.Pointer, c.Int, OffT) c.Int

// llgo:type C
type EspVfsFtruncateOpT func(c.Int, OffT) c.Int

// llgo:type C
type EspVfsUtimeCtxOpT func(c.Pointer, *c.Char, *Utimbuf) c.Int

// llgo:type C
type EspVfsUtimeOpT func(*c.Char, *Utimbuf) c.Int

/**
 * @brief Struct containing function pointers to directory related functionality.
 *
 */

type EspVfsDirOpsT struct {
	Unused [8]uint8
}

// llgo:type C
type EspVfsTcsetattrCtxOpT func(c.Pointer, c.Int, c.Int, *Termios) c.Int

// llgo:type C
type EspVfsTcsetattrOpT func(c.Int, c.Int, *Termios) c.Int

// llgo:type C
type EspVfsTcgetattrCtxOpT func(c.Pointer, c.Int, *Termios) c.Int

// llgo:type C
type EspVfsTcgetattrOpT func(c.Int, *Termios) c.Int

// llgo:type C
type EspVfsTcdrainCtxOpT func(c.Pointer, c.Int) c.Int

// llgo:type C
type EspVfsTcdrainOpT func(c.Int) c.Int

// llgo:type C
type EspVfsTcflushCtxOpT func(c.Pointer, c.Int, c.Int) c.Int

// llgo:type C
type EspVfsTcflushOpT func(c.Int, c.Int) c.Int

// llgo:type C
type EspVfsTcflowCtxOpT func(c.Pointer, c.Int, c.Int) c.Int

// llgo:type C
type EspVfsTcflowOpT func(c.Int, c.Int) c.Int

// llgo:type C
type EspVfsTcgetsidCtxOpT func(c.Pointer, c.Int) PidT

// llgo:type C
type EspVfsTcgetsidOpT func(c.Int) PidT

// llgo:type C
type EspVfsTcsendbreakCtxOpT func(c.Pointer, c.Int, c.Int) c.Int

// llgo:type C
type EspVfsTcsendbreakOpT func(c.Int, c.Int) c.Int

/**
 * @brief Struct containing function pointers to termios related functionality.
 *
 */

type EspVfsTermiosOpsT struct {
	Unused [8]uint8
}

// llgo:type C
type EspVfsStartSelectOpT func(c.Int, *FdSet, *FdSet, *FdSet, EspVfsSelectSemT, *c.Pointer) EspErrT

// llgo:type C
type EspVfsSocketSelectOpT func(c.Int, *FdSet, *FdSet, *FdSet, *Timeval) c.Int

// llgo:type C
type EspVfsStopSocketSelectOpT func(c.Pointer)

// llgo:type C
type EspVfsStopSocketSelectIsrOpT func(c.Pointer, *BaseTypeT)

// llgo:type C
type EspVfsGetSocketSelectSemaphoreOpT func() c.Pointer

// llgo:type C
type EspVfsEndSelectOpT func(c.Pointer) EspErrT

/**
 * @brief Struct containing function pointers to select related functionality.
 *
 */

type EspVfsSelectOpsT struct {
	StartSelect              EspVfsStartSelectOpT
	SocketSelect             EspVfsSocketSelectOpT
	StopSocketSelect         EspVfsStopSocketSelectOpT
	StopSocketSelectIsr      EspVfsStopSocketSelectIsrOpT
	GetSocketSelectSemaphore EspVfsGetSocketSelectSemaphoreOpT
	EndSelect                EspVfsEndSelectOpT
}

// llgo:type C
type EspVfsWriteCtxOpT func(c.Pointer, c.Int, c.Pointer, c.SizeT) c.SsizeT

// llgo:type C
type EspVfsWriteOpT func(c.Int, c.Pointer, c.SizeT) c.SsizeT

// llgo:type C
type EspVfsLseekCtxOpT func(c.Pointer, c.Int, OffT, c.Int) OffT

// llgo:type C
type EspVfsLseekOpT func(c.Int, OffT, c.Int) OffT

// llgo:type C
type EspVfsReadCtxOpT func(c.Pointer, c.Int, c.Pointer, c.SizeT) c.SsizeT

// llgo:type C
type EspVfsReadOpT func(c.Int, c.Pointer, c.SizeT) c.SsizeT

// llgo:type C
type EspVfsPreadCtxOpT func(c.Pointer, c.Int, c.Pointer, c.SizeT, OffT) c.SsizeT

// llgo:type C
type EspVfsPreadOpT func(c.Int, c.Pointer, c.SizeT, OffT) c.SsizeT

// llgo:type C
type EspVfsPwriteCtxOpT func(c.Pointer, c.Int, c.Pointer, c.SizeT, OffT) c.SsizeT

// llgo:type C
type EspVfsPwriteOpT func(c.Int, c.Pointer, c.SizeT, OffT) c.SsizeT

// llgo:type C
type EspVfsOpenCtxOpT func(c.Pointer, *c.Char, c.Int, c.Int) c.Int

// llgo:type C
type EspVfsOpenOpT func(*c.Char, c.Int, c.Int) c.Int

// llgo:type C
type EspVfsCloseCtxOpT func(c.Pointer, c.Int) c.Int

// llgo:type C
type EspVfsCloseOpT func(c.Int) c.Int

// llgo:type C
type EspVfsFstatCtxOpT func(c.Pointer, c.Int, *Stat) c.Int

// llgo:type C
type EspVfsFstatOpT func(c.Int, *Stat) c.Int

// llgo:type C
type EspVfsFcntlCtxOpT func(c.Pointer, c.Int, c.Int, c.Int) c.Int

// llgo:type C
type EspVfsFcntlOpT func(c.Int, c.Int, c.Int) c.Int

// llgo:type C
type EspVfsIoctlCtxOpT func(c.Pointer, c.Int, c.Int, c.VaList) c.Int

// llgo:type C
type EspVfsIoctlOpT func(c.Int, c.Int, c.VaList) c.Int

// llgo:type C
type EspVfsFsyncCtxOpT func(c.Pointer, c.Int) c.Int

// llgo:type C
type EspVfsFsyncOpT func(c.Int) c.Int

/**
 * @brief Main struct of the minified vfs API, containing basic function pointers as well as pointers to the other subcomponents.
 *
 */

type EspVfsFsOpsT struct {
	Dir     *EspVfsDirOpsT
	Termios *EspVfsTermiosOpsT
	Select  *EspVfsSelectOpsT
}

/**
 * Register a virtual filesystem for given path prefix.
 *
 * @param base_path  file path prefix associated with the filesystem.
 *                   Must be a zero-terminated C string, may be empty.
 *                   If not empty, must be up to ESP_VFS_PATH_MAX
 *                   characters long, and at least 2 characters long.
 *                   Name must start with a "/" and must not end with "/".
 *                   For example, "/data" or "/dev/spi" are valid.
 *                   These VFSes would then be called to handle file paths such as
 *                   "/data/myfile.txt" or "/dev/spi/0".
 *                   In the special case of an empty base_path, a "fallback"
 *                   VFS is registered. Such VFS will handle paths which are not
 *                   matched by any other registered VFS.
 * @param vfs  Pointer to esp_vfs_fs_ops_t, a structure which maps syscalls to
 *             the filesystem driver functions. VFS component does not assume ownership of this struct, but see flags for more info
 *
 * @param flags Set of binary flags controlling how the registered FS should be treated
 *             - ESP_VFS_FLAG_STATIC - if this flag is specified VFS assumes the provided esp_vfs_fs_ops_t and all its subcomponents are statically allocated,
 *                                     if it is not enabled a deep copy of the provided struct will be created, which will be managed by the VFS component
 *             - ESP_VFS_FLAG_CONTEXT_PTR - If set, the VFS will use the context-aware versions of the filesystem operation functions (suffixed with `_p`) in `esp_vfs_fs_ops_t` and its subcomponents.
 *                                          The `ctx` parameter will be passed as the context argument when these functions are invoked.
 *
 * @param ctx  Context pointer for fs operation functions, see the ESP_VFS_FLAG_CONTEXT_PTR.
 *             Should be `NULL` if not used.
 *
 * @return  ESP_OK if successful, ESP_ERR_NO_MEM if too many FSes are
 *          registered.
 */
//go:linkname EspVfsRegisterFs C.esp_vfs_register_fs
func EspVfsRegisterFs(base_path *c.Char, vfs *EspVfsFsOpsT, flags c.Int, ctx c.Pointer) EspErrT

/**
 * Analog of esp_vfs_register_with_id which accepts esp_vfs_fs_ops_t instead.
 *
 */
// llgo:link (*EspVfsFsOpsT).EspVfsRegisterFsWithId C.esp_vfs_register_fs_with_id
func (recv_ *EspVfsFsOpsT) EspVfsRegisterFsWithId(flags c.Int, ctx c.Pointer, id *EspVfsIdT) EspErrT {
	return 0
}

/**
 * Alias for esp_vfs_unregister for naming consistency
 */
//go:linkname EspVfsUnregisterFs C.esp_vfs_unregister_fs
func EspVfsUnregisterFs(base_path *c.Char) EspErrT

/**
 * Alias for esp_vfs_unregister_with_id for naming consistency
 */
// llgo:link EspVfsIdT.EspVfsUnregisterFsWithId C.esp_vfs_unregister_fs_with_id
func (recv_ EspVfsIdT) EspVfsUnregisterFsWithId() EspErrT {
	return 0
}
