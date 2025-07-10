package espcoredump

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief  Retrieves address and size of coredump data in flash.
 *         This function is always available, even when core dump is disabled in menuconfig.
 *
 * @param  out_addr   pointer to store image address in flash.
 * @param  out_size   pointer to store image size in flash (including checksum). In bytes.
 *
 * @return ESP_OK on success, otherwise \see esp_err_t
 */
//go:linkname EspCoreDumpImageGet C.esp_core_dump_image_get
func EspCoreDumpImageGet(out_addr *c.SizeT, out_size *c.SizeT) EspErrT

/**
 * @brief  Erases coredump data in flash. esp_core_dump_image_get() will then return
 *         ESP_ERR_NOT_FOUND. Can be used after a coredump has been transmitted successfully.
 *         This function is always available, even when core dump is disabled in menuconfig.
 *
 * @return ESP_OK on success, otherwise \see esp_err_t
 */
//go:linkname EspCoreDumpImageErase C.esp_core_dump_image_erase
func EspCoreDumpImageErase() EspErrT
