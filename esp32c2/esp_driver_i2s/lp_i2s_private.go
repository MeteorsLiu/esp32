package esp_driver_i2s

import _ "unsafe"

/**
 * @brief Get LP I2S soc handle
 *
 * @param[in] chan  LP I2S channel handle
 *
 * @return LP I2S soc handle
 */
//go:linkname LpI2sGetSocHandle C.lp_i2s_get_soc_handle
func LpI2sGetSocHandle(chan_ LpI2sChanHandleT) LpI2sSocHandleT
