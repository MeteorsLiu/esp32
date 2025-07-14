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
 * @brief HMAC-based scheme specific configuration data
 */

type NvsSecConfigHmacT struct {
	HmacKeyId HmacKeyIdT
}

/**
 * @brief Register the HMAC-based scheme for NVS Encryption
 *
 * @param[in]  sec_scheme_cfg          Security scheme specific configuration data
 * @param[out] sec_scheme_handle_out   Security scheme specific configuration handle
 *
 * @return
 *      - ESP_OK, if `sec_scheme_handle_out` was populated successfully with the scheme configuration;
 *      - ESP_ERR_INVALID_ARG, if `scheme_cfg_hmac` is NULL;
 *      - ESP_ERR_NO_MEM, No memory for the scheme-specific handle `sec_scheme_handle_out`
 */
// llgo:link (*NvsSecConfigHmacT).NvsSecProviderRegisterHmac C.nvs_sec_provider_register_hmac
func (recv_ *NvsSecConfigHmacT) NvsSecProviderRegisterHmac(sec_scheme_handle_out **NvsSecSchemeT) EspErrT {
	return 0
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
