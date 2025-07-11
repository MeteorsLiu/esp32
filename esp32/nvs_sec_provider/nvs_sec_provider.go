package nvs_sec_provider

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_ERR_NVS_SEC_BASE = 0xF000

type NvsSecSchemeIdT c.Int

const (
	NVS_SEC_SCHEME_FLASH_ENC NvsSecSchemeIdT = 0
	NVS_SEC_SCHEME_HMAC      NvsSecSchemeIdT = 1
	NVS_SEC_SCHEME_MAX       NvsSecSchemeIdT = 2
)

/**
 * @brief Flash encryption-based scheme specific configuration data
 */

type NvsSecConfigFlashEncT struct {
	NvsKeysPart *EspPartitionT
}

/**
 * @brief Deregister the NVS encryption scheme registered with the given handle
 *
 * @param[in] sec_scheme_handle  Security scheme specific configuration handle

 * @return
 *      - ESP_OK, if the scheme registered with `sec_scheme_handle` was deregistered successfully
 *      - ESP_ERR_INVALID_ARG, if `sec_scheme_handle` is NULL;
 */
//go:linkname NvsSecProviderDeregister C.nvs_sec_provider_deregister
func NvsSecProviderDeregister(sec_scheme_handle *NvsSecSchemeT) EspErrT
