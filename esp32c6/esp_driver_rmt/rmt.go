package esp_driver_rmt

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get the unique ID of RMT channel
 *
 * @param[in] channel RMT generic channel that created by `rmt_new_tx_channel()` or `rmt_new_rx_channel()`
 * @param[out] ret_id The unique ID of RMT channel
 * @return
 *      - ESP_OK: Get RMT channel ID successfully
 *      - ESP_ERR_INVALID_ARG: Get RMT channel ID failed because of invalid argument
 *      - ESP_FAIL: Get RMT channel ID failed because of other reasons
 */
//go:linkname RmtGetChannelId C.rmt_get_channel_id
func RmtGetChannelId(channel RmtChannelHandleT, ret_id *c.Int) EspErrT

/**
 * @brief Get the RMT channel's real clock resolution
 *
 * @param[in] channel RMT generic channel that created by `rmt_new_tx_channel()` or `rmt_new_rx_channel()`
 * @param[out] ret_resolution_hz The real clock resolution in Hz
 * @return
 *      - ESP_OK: Get RMT channel resolution successfully
 *      - ESP_ERR_INVALID_ARG: Get RMT channel resolution failed because of invalid argument
 *      - ESP_FAIL: Get RMT channel resolution failed because of other reasons
 */
//go:linkname RmtGetChannelResolution C.rmt_get_channel_resolution
func RmtGetChannelResolution(channel RmtChannelHandleT, ret_resolution_hz *c.Uint32T) EspErrT
