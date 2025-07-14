package esp_mm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Memory sync between Cache and storage memory
 *
 *
 * For cache-to-memory (C2M) direction:
 * - For cache writeback supported chips (you can refer to SOC_CACHE_WRITEBACK_SUPPORTED in soc_caps.h)
 *   - This API will do a writeback to synchronise between cache and storage memory
 *   - With ESP_CACHE_MSYNC_FLAG_INVALIDATE, this API will also invalidate the values that just written
 *   - Note: although ESP32 is with PSRAM, but cache writeback isn't supported, so this API will do nothing on ESP32
 * - For other chips, this API will do nothing. The out-of-sync should be already dealt by the SDK
 *
 * For memory-to-cache (M2C) direction:
 * - This API will by default do an invalidation
 *
 * This API is cache-safe and thread-safe
 *
 * @note If you don't set direction (ESP_CACHE_MSYNC_FLAG_DIR_x flags), this API is by default C2M direction
 * @note If you don't set type (ESP_CACHE_MSYNC_FLAG_TYPE_x flags), this API is by default doing msync for data
 * @note You should not call this during any Flash operations (e.g. esp_flash APIs, nvs and some other APIs that are based on esp_flash APIs)
 * @note If XIP_From_PSRAM is enabled (by enabling both CONFIG_SPIRAM_FETCH_INSTRUCTIONS and CONFIG_SPIRAM_RODATA), you can call this API during Flash operations
 *
 * @param[in] addr   Starting address to do the msync
 * @param[in] size   Size to do the msync
 * @param[in] flags  Flags, see `ESP_CACHE_MSYNC_FLAG_x`
 *
 * @return
 *        - ESP_OK:
 *                  - Successful msync
 *                  - For C2M direction, if this chip doesn't support cache writeback, if the input addr is a cache supported one, this API will return ESP_OK
 *        - ESP_ERR_INVALID_ARG:   Invalid argument, not cache supported addr, see printed logs
 */
//go:linkname EspCacheMsync C.esp_cache_msync
func EspCacheMsync(addr c.Pointer, size c.SizeT, flags c.Int) EspErrT
