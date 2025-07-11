package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
* @brief Ethernet DMA TX Descriptor
*
 */
type EthDmaTxDescriptorT struct {
	TDES0 struct {
		Value c.Uint32T
	}
	TDES1 struct {
		Value c.Uint32T
	}
	Buffer1Addr         c.Uint32T
	Buffer2NextDescAddr c.Uint32T
	Reserved1           c.Uint32T
	Reserved2           c.Uint32T
	TimeStampLow        c.Uint32T
	TimeStampHigh       c.Uint32T
}

/**
* @brief Ethernet DMA RX Descriptor
*
 */
type EthDmaRxDescriptorT struct {
	RDES0 struct {
		Value c.Uint32T
	}
	RDES1 struct {
		Value c.Uint32T
	}
	Buffer1Addr         c.Uint32T
	Buffer2NextDescAddr c.Uint32T
	ExtendedStatus      struct {
		Value c.Uint32T
	}
	Reserved      c.Uint32T
	TimeStampLow  c.Uint32T
	TimeStampHigh c.Uint32T
}
type EmacMacSocRegsT *EmacMacDevS
type EmacDmaSocRegsT *EmacDmaDevS
type EmacExtSocRegsT *EmacExtDevS

type EmacHalContextT struct {
	MacRegs EmacMacSocRegsT
	DmaRegs EmacDmaSocRegsT
	ExtRegs EmacExtSocRegsT
}

/**
 * @brief EMAC related configuration
 */

type EmacHalDmaConfigT struct {
	DmaBurstLen EthMacDmaBurstLenT
}

// llgo:link (*EmacHalContextT).EmacHalInit C.emac_hal_init
func (recv_ *EmacHalContextT) EmacHalInit() {
}

// llgo:link (*EmacHalContextT).EmacHalSetCsrClockRange C.emac_hal_set_csr_clock_range
func (recv_ *EmacHalContextT) EmacHalSetCsrClockRange(freq c.Int) {
}

// llgo:link (*EmacHalContextT).EmacHalInitMacDefault C.emac_hal_init_mac_default
func (recv_ *EmacHalContextT) EmacHalInitMacDefault() {
}

// llgo:link (*EmacHalContextT).EmacHalInitDmaDefault C.emac_hal_init_dma_default
func (recv_ *EmacHalContextT) EmacHalInitDmaDefault(hal_config *EmacHalDmaConfigT) {
}

// llgo:link (*EmacHalContextT).EmacHalSetPhyCmd C.emac_hal_set_phy_cmd
func (recv_ *EmacHalContextT) EmacHalSetPhyCmd(phy_addr c.Uint32T, phy_reg c.Uint32T, write bool) {
}

// llgo:link (*EmacHalContextT).EmacHalSetAddress C.emac_hal_set_address
func (recv_ *EmacHalContextT) EmacHalSetAddress(mac_addr *c.Uint8T) {
}

/**
 * @brief Starts EMAC Transmission & Reception
 *
 * @param hal EMAC HAL context infostructure
 */
// llgo:link (*EmacHalContextT).EmacHalStart C.emac_hal_start
func (recv_ *EmacHalContextT) EmacHalStart() {
}

/**
* @brief Stops EMAC Transmission & Reception
*
* @param hal EMAC HAL context infostructure
* @return
*     - ESP_OK: succeed
 *    - ESP_ERR_INVALID_STATE: previous frame transmission/reception is not completed. When this error occurs,
 *      wait and repeat the EMAC stop again.
*/
// llgo:link (*EmacHalContextT).EmacHalStop C.emac_hal_stop
func (recv_ *EmacHalContextT) EmacHalStop() EspErrT {
	return 0
}

// llgo:link (*EmacHalContextT).EmacHalEnableFlowCtrl C.emac_hal_enable_flow_ctrl
func (recv_ *EmacHalContextT) EmacHalEnableFlowCtrl(enable bool) {
}

// llgo:link (*EmacHalContextT).EmacHalSetRxTxDescAddr C.emac_hal_set_rx_tx_desc_addr
func (recv_ *EmacHalContextT) EmacHalSetRxTxDescAddr(rx_desc *EthDmaRxDescriptorT, tx_desc *EthDmaTxDescriptorT) {
}
