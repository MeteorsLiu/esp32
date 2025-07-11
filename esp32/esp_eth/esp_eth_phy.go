package esp_eth

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EthPhyAutonegCmdT c.Int

const (
	ESP_ETH_PHY_AUTONEGO_RESTART EthPhyAutonegCmdT = 0
	ESP_ETH_PHY_AUTONEGO_EN      EthPhyAutonegCmdT = 1
	ESP_ETH_PHY_AUTONEGO_DIS     EthPhyAutonegCmdT = 2
	ESP_ETH_PHY_AUTONEGO_G_STAT  EthPhyAutonegCmdT = 3
)

type EspEthPhyS struct {
	Unused [8]uint8
}
type EspEthPhyT EspEthPhyS

/**
* @brief Ethernet PHY configuration
*
 */
type EthPhyConfigT struct {
	PhyAddr             c.Int32T
	ResetTimeoutMs      c.Uint32T
	AutonegoTimeoutMs   c.Uint32T
	ResetGpioNum        c.Int
	HwResetAssertTimeUs c.Int32T
	PostHwResetDelayMs  c.Int32T
}

/**
* @brief Create a PHY instance of generic chip which conforms with IEEE 802.3
*
* @note Default reset timing configuration is set conservatively( @c DEFAULT_PHY_RESET_ASSERTION_TIME_US ).
*       If you need faster response and your chip supports it, configure it via @c config parameter.
*
* @warning While basic functionality should always work, some specific features might be limited,
*          even if the PHY meets IEEE 802.3 standard. A typical example is loopback functionality,
*          where certain PHYs may require setting a specific speed mode to operate correctly.
*
* @param[in] config configuration of PHY
* @return
*      - instance: create PHY instance successfully
*      - NULL: create PHY instance failed because some error occurred
 */
// llgo:link (*EthPhyConfigT).EspEthPhyNewGeneric C.esp_eth_phy_new_generic
func (recv_ *EthPhyConfigT) EspEthPhyNewGeneric() *EspEthPhyT {
	return nil
}

/**
* @brief Create a PHY instance of IP101
*
* @param[in] config: configuration of PHY
*
* @return
*      - instance: create PHY instance successfully
*      - NULL: create PHY instance failed because some error occurred
 */
// llgo:link (*EthPhyConfigT).EspEthPhyNewIp101 C.esp_eth_phy_new_ip101
func (recv_ *EthPhyConfigT) EspEthPhyNewIp101() *EspEthPhyT {
	return nil
}

/**
* @brief Create a PHY instance of RTL8201
*
* @param[in] config: configuration of PHY
*
* @return
*      - instance: create PHY instance successfully
*      - NULL: create PHY instance failed because some error occurred
 */
// llgo:link (*EthPhyConfigT).EspEthPhyNewRtl8201 C.esp_eth_phy_new_rtl8201
func (recv_ *EthPhyConfigT) EspEthPhyNewRtl8201() *EspEthPhyT {
	return nil
}

/**
* @brief Create a PHY instance of LAN87xx
*
* @param[in] config: configuration of PHY
*
* @return
*      - instance: create PHY instance successfully
*      - NULL: create PHY instance failed because some error occurred
 */
// llgo:link (*EthPhyConfigT).EspEthPhyNewLan87xx C.esp_eth_phy_new_lan87xx
func (recv_ *EthPhyConfigT) EspEthPhyNewLan87xx() *EspEthPhyT {
	return nil
}

/**
* @brief Create a PHY instance of DP83848
*
* @param[in] config: configuration of PHY
*
* @return
*      - instance: create PHY instance successfully
*      - NULL: create PHY instance failed because some error occurred
 */
// llgo:link (*EthPhyConfigT).EspEthPhyNewDp83848 C.esp_eth_phy_new_dp83848
func (recv_ *EthPhyConfigT) EspEthPhyNewDp83848() *EspEthPhyT {
	return nil
}

/**
* @brief Create a PHY instance of KSZ80xx
*
* The phy model from the KSZ80xx series is detected automatically. If the driver
* is unable to detect a supported model, \c NULL is returned.
*
* Currently, the following models are supported:
* KSZ8001, KSZ8021, KSZ8031, KSZ8041, KSZ8051, KSZ8061, KSZ8081, KSZ8091
*
* @param[in] config: configuration of PHY
*
* @return
*      - instance: create PHY instance successfully
*      - NULL: create PHY instance failed because some error occurred
 */
// llgo:link (*EthPhyConfigT).EspEthPhyNewKsz80xx C.esp_eth_phy_new_ksz80xx
func (recv_ *EthPhyConfigT) EspEthPhyNewKsz80xx() *EspEthPhyT {
	return nil
}
