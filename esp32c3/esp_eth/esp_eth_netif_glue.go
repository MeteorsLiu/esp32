package esp_eth

import _ "unsafe"

type EspEthNetifGlueT struct {
	Unused [8]uint8
}
type EspEthNetifGlueHandleT *EspEthNetifGlueT

/**
 * @brief Create a netif glue for Ethernet driver
 * @note netif glue is used to attach io driver to TCP/IP netif
 *
 * @param eth_hdl Ethernet driver handle
 * @return glue object, which inherits esp_netif_driver_base_t
 */
//go:linkname EspEthNewNetifGlue C.esp_eth_new_netif_glue
func EspEthNewNetifGlue(eth_hdl EspEthHandleT) EspEthNetifGlueHandleT

/**
 * @brief Delete netif glue of Ethernet driver
 *
 * @param eth_netif_glue netif glue
 * @return -ESP_OK: delete netif glue successfully
 */
//go:linkname EspEthDelNetifGlue C.esp_eth_del_netif_glue
func EspEthDelNetifGlue(eth_netif_glue EspEthNetifGlueHandleT) EspErrT
