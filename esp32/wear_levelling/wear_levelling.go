package wear_levelling

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type WlHandleT c.Int32T

/**
* @brief Mount WL for defined partition
*
* @param partition that will be used for access
* @param out_handle handle of the WL instance
*
* @return
*       - ESP_OK, if the WL allocation is successful;
*       - ESP_ERR_INVALID_ARG, if the arguments for WL configuration are not valid;
*       - ESP_ERR_NO_MEM, if the WL allocation fails because of insufficient memory;
 */
//go:linkname WlMount C.wl_mount
func WlMount(partition *EspPartitionT, out_handle *WlHandleT) EspErrT

/**
* @brief Unmount WL for defined partition
*
* @param handle WL partition handle
*
* @return
*       - ESP_OK, if the operation is successful;
*       - or one of error codes from lower-level flash driver.
 */
// llgo:link WlHandleT.WlUnmount C.wl_unmount
func (recv_ WlHandleT) WlUnmount() EspErrT {
	return 0
}

/**
* @brief Erase part of the WL storage
*
* @param handle WL handle that are related to the partition
* @param start_addr Address from where erase operation should start. Must be aligned
*                   to the result of function wl_sector_size(...).
* @param size Size of the range which should be erased, in bytes.
*                   Must be divisible by the result of function wl_sector_size(...)..
*
* @return
*       - ESP_OK, if the given range was erased successfully;
*       - ESP_ERR_INVALID_ARG, if iterator or dst are NULL;
*       - ESP_ERR_INVALID_SIZE, if erase would go out of bounds of the partition;
*       - or one of error codes from lower-level flash driver.
 */
// llgo:link WlHandleT.WlEraseRange C.wl_erase_range
func (recv_ WlHandleT) WlEraseRange(start_addr c.SizeT, size c.SizeT) EspErrT {
	return 0
}

/**
* @brief Write data to the WL storage
*
* Before writing data to flash, corresponding region of flash needs to be erased.
* This can be done using wl_erase_range function.
*
* @param handle WL handle corresponding to the WL partition
* @param dest_addr Address where the data should be written, relative to the
*                  beginning of the partition.
* @param src Pointer to the source buffer. Pointer must be non-NULL and
*            buffer must be at least 'size' bytes long.
* @param size Size of data to be written, in bytes.
*
* @note Prior to writing to WL storage, make sure it has been erased with
*       wl_erase_range call.
*
* @return
*       - ESP_OK, if data was written successfully;
*       - ESP_ERR_INVALID_ARG, if dst_offset exceeds partition size;
*       - ESP_ERR_INVALID_SIZE, if write would go out of bounds of the partition;
*       - or one of error codes from lower-level flash driver.
 */
// llgo:link WlHandleT.WlWrite C.wl_write
func (recv_ WlHandleT) WlWrite(dest_addr c.SizeT, src c.Pointer, size c.SizeT) EspErrT {
	return 0
}

/**
* @brief Read data from the WL storage
*
* @param handle WL module instance that was initialized before
* @param dest Pointer to the buffer where data should be stored.
*             The Pointer must be non-NULL and the buffer must be at least 'size' bytes long.
* @param src_addr Address of the data to be read, relative to the
*                 beginning of the partition.
* @param size Size of data to be read, in bytes.
*
* @return
*       - ESP_OK, if data was read successfully;
*       - ESP_ERR_INVALID_ARG, if src_offset exceeds partition size;
*       - ESP_ERR_INVALID_SIZE, if read would go out of bounds of the partition;
*       - or one of error codes from lower-level flash driver.
 */
// llgo:link WlHandleT.WlRead C.wl_read
func (recv_ WlHandleT) WlRead(src_addr c.SizeT, dest c.Pointer, size c.SizeT) EspErrT {
	return 0
}

/**
* @brief Get the actual flash size in use for the WL storage partition
*
* @param handle WL module handle that was initialized before
* @return usable size, in bytes
 */
// llgo:link WlHandleT.WlSize C.wl_size
func (recv_ WlHandleT) WlSize() c.SizeT {
	return 0
}

/**
* @brief Get sector size of the WL instance
*
* @param handle WL module handle that was initialized before
* @return sector size, in bytes
 */
// llgo:link WlHandleT.WlSectorSize C.wl_sector_size
func (recv_ WlHandleT) WlSectorSize() c.SizeT {
	return 0
}
