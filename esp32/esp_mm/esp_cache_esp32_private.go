package esp_mm

import _ "unsafe"

type CacheDriverS struct {
	Unused [8]uint8
}
type CacheDriverT CacheDriverS

/**
 * @brief Register cache writeback
 *
 * @param[in] func    Cache driver
 */
// llgo:link (*CacheDriverT).CacheRegisterWriteback C.cache_register_writeback
func (recv_ *CacheDriverT) CacheRegisterWriteback() {
}

/**
 * @brief Cache sync
 *
 * @note This API only do cache sync, but doesn't guarantee concurrent access to cache
 * @note Do not use in your application
 */
//go:linkname CacheSync C.cache_sync
func CacheSync()
