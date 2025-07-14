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
