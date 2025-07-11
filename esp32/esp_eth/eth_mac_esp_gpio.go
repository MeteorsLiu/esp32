package esp_eth

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:link (*EthMacMiiGpioConfigT).EmacEspGpioMatrixInitMii C.emac_esp_gpio_matrix_init_mii
func (recv_ *EthMacMiiGpioConfigT) EmacEspGpioMatrixInitMii() EspErrT {
	return 0
}

// llgo:link (*EthMacMiiGpioConfigT).EmacEspIomuxInitMii C.emac_esp_iomux_init_mii
func (recv_ *EthMacMiiGpioConfigT) EmacEspIomuxInitMii() EspErrT {
	return 0
}

// llgo:link (*EthMacRmiiGpioConfigT).EmacEspIomuxInitRmii C.emac_esp_iomux_init_rmii
func (recv_ *EthMacRmiiGpioConfigT) EmacEspIomuxInitRmii() EspErrT {
	return 0
}

//go:linkname EmacEspIomuxRmiiClkInput C.emac_esp_iomux_rmii_clk_input
func EmacEspIomuxRmiiClkInput(num c.Int) EspErrT

//go:linkname EmacEspIomuxRmiiClkOuput C.emac_esp_iomux_rmii_clk_ouput
func EmacEspIomuxRmiiClkOuput(num c.Int) EspErrT

//go:linkname EmacEspIomuxRmiiInitTxEr C.emac_esp_iomux_rmii_init_tx_er
func EmacEspIomuxRmiiInitTxEr(num c.Int) EspErrT

//go:linkname EmacEspIomuxRmiiInitRxEr C.emac_esp_iomux_rmii_init_rx_er
func EmacEspIomuxRmiiInitRxEr(num c.Int) EspErrT

//go:linkname EmacEspIomuxMiiInitTxEr C.emac_esp_iomux_mii_init_tx_er
func EmacEspIomuxMiiInitTxEr(num c.Int) EspErrT

//go:linkname EmacEspIomuxMiiInitRxEr C.emac_esp_iomux_mii_init_rx_er
func EmacEspIomuxMiiInitRxEr(num c.Int) EspErrT

// llgo:link (*EmacEspSmiGpioConfigT).EmacEspGpioInitSmi C.emac_esp_gpio_init_smi
func (recv_ *EmacEspSmiGpioConfigT) EmacEspGpioInitSmi() EspErrT {
	return 0
}

//go:linkname EmacEspGpioDeinitAll C.emac_esp_gpio_deinit_all
func EmacEspGpioDeinitAll() EspErrT
