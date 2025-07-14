package vfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_VFS_PATH_MAX = 15

/**
 * @brief VFS definition structure
 *
 * This structure should be filled with pointers to corresponding
 * FS driver functions.
 *
 * VFS component will translate all FDs so that the filesystem implementation
 * sees them starting at zero. The caller sees a global FD which is prefixed
 * with an pre-filesystem-implementation.
 *
 * Some FS implementations expect some state (e.g. pointer to some structure)
 * to be passed in as a first argument. For these implementations,
 * populate the members of this structure which have _p suffix, set
 * flags member to ESP_VFS_FLAG_CONTEXT_PTR and provide the context pointer
 * to esp_vfs_register function.
 * If the implementation doesn't use this extra argument, populate the
 * members without _p suffix and set flags member to ESP_VFS_FLAG_DEFAULT.
 *
 * If the FS driver doesn't provide some of the functions, set corresponding
 * members to NULL.
 */

type EspVfsT struct {
	Flags                    c.Int
	StartSelect              c.Pointer
	SocketSelect             c.Pointer
	StopSocketSelect         c.Pointer
	StopSocketSelectIsr      c.Pointer
	GetSocketSelectSemaphore c.Pointer
	EndSelect                c.Pointer
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
 * @param vfs  Pointer to esp_vfs_t, a structure which maps syscalls to
 *             the filesystem driver functions. VFS component doesn't
 *             assume ownership of this pointer.
 * @param ctx  If vfs->flags has ESP_VFS_FLAG_CONTEXT_PTR set, a pointer
 *             which should be passed to VFS functions. Otherwise, NULL.
 *
 * @return  ESP_OK if successful, ESP_ERR_NO_MEM if too many VFSes are
 *          registered.
 */
//go:linkname EspVfsRegister C.esp_vfs_register
func EspVfsRegister(base_path *c.Char, vfs *EspVfsT, ctx c.Pointer) EspErrT

/**
 * Special case function for registering a VFS that uses a method other than
 * open() to open new file descriptors from the interval <min_fd; max_fd).
 *
 * This is a special-purpose function intended for registering LWIP sockets to VFS.
 *
 * @param vfs Pointer to esp_vfs_t. Meaning is the same as for esp_vfs_register().
 * @param ctx Pointer to context structure. Meaning is the same as for esp_vfs_register().
 * @param min_fd The smallest file descriptor this VFS will use.
 * @param max_fd Upper boundary for file descriptors this VFS will use (the biggest file descriptor plus one).
 *
 * @return  ESP_OK if successful, ESP_ERR_NO_MEM if too many VFSes are
 *          registered, ESP_ERR_INVALID_ARG if the file descriptor boundaries
 *          are incorrect.
 */
// llgo:link (*EspVfsT).EspVfsRegisterFdRange C.esp_vfs_register_fd_range
func (recv_ *EspVfsT) EspVfsRegisterFdRange(ctx c.Pointer, min_fd c.Int, max_fd c.Int) EspErrT {
	return 0
}

/**
 * Special case function for registering a VFS that uses a method other than
 * open() to open new file descriptors. In comparison with
 * esp_vfs_register_fd_range, this function doesn't pre-registers an interval
 * of file descriptors. File descriptors can be registered later, by using
 * esp_vfs_register_fd.
 *
 * @param vfs Pointer to esp_vfs_t. Meaning is the same as for esp_vfs_register().
 * @param ctx Pointer to context structure. Meaning is the same as for esp_vfs_register().
 * @param vfs_id Here will be written the VFS ID which can be passed to
 *               esp_vfs_register_fd for registering file descriptors.
 *
 * @return  ESP_OK if successful, ESP_ERR_NO_MEM if too many VFSes are
 *          registered, ESP_ERR_INVALID_ARG if the file descriptor boundaries
 *          are incorrect.
 */
// llgo:link (*EspVfsT).EspVfsRegisterWithId C.esp_vfs_register_with_id
func (recv_ *EspVfsT) EspVfsRegisterWithId(ctx c.Pointer, vfs_id *EspVfsIdT) EspErrT {
	return 0
}

/**
 * Unregister a virtual filesystem for given path prefix
 *
 * @param base_path  file prefix previously used in esp_vfs_register call
 * @return ESP_OK if successful, ESP_ERR_INVALID_STATE if VFS for given prefix
 *         hasn't been registered
 */
//go:linkname EspVfsUnregister C.esp_vfs_unregister
func EspVfsUnregister(base_path *c.Char) EspErrT

/**
 * Unregister a virtual filesystem with the given index
 *
 * @param vfs_id  The VFS ID returned by esp_vfs_register_with_id
 * @return ESP_OK if successful, ESP_ERR_INVALID_STATE if VFS for the given index
 *         hasn't been registered
 */
// llgo:link EspVfsIdT.EspVfsUnregisterWithId C.esp_vfs_unregister_with_id
func (recv_ EspVfsIdT) EspVfsUnregisterWithId() EspErrT {
	return 0
}

/**
 * Special function for registering another file descriptor for a VFS registered
 * by esp_vfs_register_with_id. This function should only be used to register
 * permanent file descriptors (socket fd) that are not removed after being closed.
 *
 * @param vfs_id VFS identificator returned by esp_vfs_register_with_id.
 * @param fd The registered file descriptor will be written to this address.
 *
 * @return  ESP_OK if the registration is successful,
 *          ESP_ERR_NO_MEM if too many file descriptors are registered,
 *          ESP_ERR_INVALID_ARG if the arguments are incorrect.
 */
// llgo:link EspVfsIdT.EspVfsRegisterFd C.esp_vfs_register_fd
func (recv_ EspVfsIdT) EspVfsRegisterFd(fd *c.Int) EspErrT {
	return 0
}

/**
 * Special function for registering another file descriptor with given local_fd
 * for a VFS registered by esp_vfs_register_with_id.
 *
 * @param vfs_id VFS identificator returned by esp_vfs_register_with_id.
 * @param local_fd The fd in the local vfs. Passing -1 will set the local fd as the (*fd) value.
 * @param permanent Whether the fd should be treated as permannet (not removed after close())
 * @param fd The registered file descriptor will be written to this address.
 *
 * @return  ESP_OK if the registration is successful,
 *          ESP_ERR_NO_MEM if too many file descriptors are registered,
 *          ESP_ERR_INVALID_ARG if the arguments are incorrect.
 */
// llgo:link EspVfsIdT.EspVfsRegisterFdWithLocalFd C.esp_vfs_register_fd_with_local_fd
func (recv_ EspVfsIdT) EspVfsRegisterFdWithLocalFd(local_fd c.Int, permanent bool, fd *c.Int) EspErrT {
	return 0
}

/**
 * Special function for unregistering a file descriptor belonging to a VFS
 * registered by esp_vfs_register_with_id.
 *
 * @param vfs_id VFS identificator returned by esp_vfs_register_with_id.
 * @param fd File descriptor which should be unregistered.
 *
 * @return  ESP_OK if the registration is successful,
 *          ESP_ERR_INVALID_ARG if the arguments are incorrect.
 */
// llgo:link EspVfsIdT.EspVfsUnregisterFd C.esp_vfs_unregister_fd
func (recv_ EspVfsIdT) EspVfsUnregisterFd(fd c.Int) EspErrT {
	return 0
}

/**
 * These functions are to be used in newlib syscall table. They will be called by
 * newlib when it needs to use any of the syscalls.
 */
/**@{*/
//go:linkname EspVfsWrite C.esp_vfs_write
func EspVfsWrite(r *X_reent, fd c.Int, data c.Pointer, size c.SizeT) c.SsizeT

//go:linkname EspVfsLseek C.esp_vfs_lseek
func EspVfsLseek(r *X_reent, fd c.Int, size OffT, mode c.Int) OffT

//go:linkname EspVfsRead C.esp_vfs_read
func EspVfsRead(r *X_reent, fd c.Int, dst c.Pointer, size c.SizeT) c.SsizeT

//go:linkname EspVfsOpen C.esp_vfs_open
func EspVfsOpen(r *X_reent, path *c.Char, flags c.Int, mode c.Int) c.Int

//go:linkname EspVfsClose C.esp_vfs_close
func EspVfsClose(r *X_reent, fd c.Int) c.Int

//go:linkname EspVfsFstat C.esp_vfs_fstat
func EspVfsFstat(r *X_reent, fd c.Int, st *Stat) c.Int

//go:linkname EspVfsStat C.esp_vfs_stat
func EspVfsStat(r *X_reent, path *c.Char, st *Stat) c.Int

//go:linkname EspVfsLink C.esp_vfs_link
func EspVfsLink(r *X_reent, n1 *c.Char, n2 *c.Char) c.Int

//go:linkname EspVfsUnlink C.esp_vfs_unlink
func EspVfsUnlink(r *X_reent, path *c.Char) c.Int

//go:linkname EspVfsRename C.esp_vfs_rename
func EspVfsRename(r *X_reent, src *c.Char, dst *c.Char) c.Int

//go:linkname EspVfsUtime C.esp_vfs_utime
func EspVfsUtime(path *c.Char, times *Utimbuf) c.Int

/**
 * @brief Synchronous I/O multiplexing which implements the functionality of POSIX select() for VFS
 * @param nfds      Specifies the range of descriptors which should be checked.
 *                  The first nfds descriptors will be checked in each set.
 * @param readfds   If not NULL, then points to a descriptor set that on input
 *                  specifies which descriptors should be checked for being
 *                  ready to read, and on output indicates which descriptors
 *                  are ready to read.
 * @param writefds  If not NULL, then points to a descriptor set that on input
 *                  specifies which descriptors should be checked for being
 *                  ready to write, and on output indicates which descriptors
 *                  are ready to write.
 * @param errorfds  If not NULL, then points to a descriptor set that on input
 *                  specifies which descriptors should be checked for error
 *                  conditions, and on output indicates which descriptors
 *                  have error conditions.
 * @param timeout   If not NULL, then points to timeval structure which
 *                  specifies the time period after which the functions should
 *                  time-out and return. If it is NULL, then the function will
 *                  not time-out. Note that the timeout period is rounded up to
 *                  the system tick and incremented by one.
 *
 * @return      The number of descriptors set in the descriptor sets, or -1
 *              when an error (specified by errno) have occurred.
 */
//go:linkname EspVfsSelect C.esp_vfs_select
func EspVfsSelect(nfds c.Int, readfds *FdSet, writefds *FdSet, errorfds *FdSet, timeout *Timeval) c.Int

/**
 * @brief Notification from a VFS driver about a read/write/error condition
 *
 * This function is called when the VFS driver detects a read/write/error
 * condition as it was requested by the previous call to start_select.
 *
 * @param sem semaphore structure which was passed to the driver by the start_select call
 */
// llgo:link EspVfsSelectSemT.EspVfsSelectTriggered C.esp_vfs_select_triggered
func (recv_ EspVfsSelectSemT) EspVfsSelectTriggered() {
}

/**
 * @brief Notification from a VFS driver about a read/write/error condition (ISR version)
 *
 * This function is called when the VFS driver detects a read/write/error
 * condition as it was requested by the previous call to start_select.
 *
 * @param sem semaphore structure which was passed to the driver by the start_select call
 * @param woken is set to pdTRUE if the function wakes up a task with higher priority
 */
// llgo:link EspVfsSelectSemT.EspVfsSelectTriggeredIsr C.esp_vfs_select_triggered_isr
func (recv_ EspVfsSelectSemT) EspVfsSelectTriggeredIsr(woken *BaseTypeT) {
}

/**
 *
 * @brief Implements the VFS layer of POSIX pread()
 *
 * @param fd         File descriptor used for read
 * @param dst        Pointer to the buffer where the output will be written
 * @param size       Number of bytes to be read
 * @param offset     Starting offset of the read
 *
 * @return           A positive return value indicates the number of bytes read. -1 is return on failure and errno is
 *                   set accordingly.
 */
//go:linkname EspVfsPread C.esp_vfs_pread
func EspVfsPread(fd c.Int, dst c.Pointer, size c.SizeT, offset OffT) c.SsizeT

/**
 *
 * @brief Implements the VFS layer of POSIX pwrite()
 *
 * @param fd         File descriptor used for write
 * @param src        Pointer to the buffer from where the output will be read
 * @param size       Number of bytes to write
 * @param offset     Starting offset of the write
 *
 * @return           A positive return value indicates the number of bytes written. -1 is return on failure and errno is
 *                   set accordingly.
 */
//go:linkname EspVfsPwrite C.esp_vfs_pwrite
func EspVfsPwrite(fd c.Int, src c.Pointer, size c.SizeT, offset OffT) c.SsizeT

/**
*
* @brief Dump the existing VFS FDs data to FILE* fp
*
* Dump the FDs in the format:
@verbatim
        <VFS Path Prefix>-<FD seen by App>-<FD seen by driver>

   where:
    VFS Path Prefix   : file prefix used in the esp_vfs_register call
    FD seen by App    : file descriptor returned by the vfs to the application for the path prefix
    FD seen by driver : file descriptor used by the driver for the same file prefix.

@endverbatim
*
* @param fp         File descriptor where data will be dumped
*/
//go:linkname EspVfsDumpFds C.esp_vfs_dump_fds
func EspVfsDumpFds(fp *c.FILE)

/**
* @brief Dump all registered FSs to the provided FILE*
*
* Dump the FSs in the format:
@verbatim
       <index>:<VFS Path Prefix> -> <VFS entry ptr>

   where:
       index           : internal index in the table of registered FSs (the same as returned when registering fd with id)
       VFS Path Prefix : file prefix used in the esp_vfs_register call or "NULL"
       VFS entry ptr   : pointer to the esp_vfs_fs_ops_t struct used internally when resolving the calls
@endverbatim
*
* @param fp File descriptor where data will be dumped
*/
//go:linkname EspVfsDumpRegisteredPaths C.esp_vfs_dump_registered_paths
func EspVfsDumpRegisteredPaths(fp *c.FILE)
