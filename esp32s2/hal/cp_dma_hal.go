package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief HAL context
 *
 */

type CpDmaHalContextT struct {
	Dev *CpDmaDevT
}

type CpDmaHalConfigT struct {
	Unused [8]uint8
}

/**
 * @brief Initialize HAL layer context
 *
 * @param hal HAL layer context, whose memroy should be allocated at driver layer
 * @param config configuration for the HAL layer
 */
// llgo:link (*CpDmaHalContextT).CpDmaHalInit C.cp_dma_hal_init
func (recv_ *CpDmaHalContextT) CpDmaHalInit(config *CpDmaHalConfigT) {
}

/**
 * @brief Deinitialize HAL layer context
 */
// llgo:link (*CpDmaHalContextT).CpDmaHalDeinit C.cp_dma_hal_deinit
func (recv_ *CpDmaHalContextT) CpDmaHalDeinit() {
}

/**
 * @brief Set descriptor base address
 */
// llgo:link (*CpDmaHalContextT).CpDmaHalSetDescBaseAddr C.cp_dma_hal_set_desc_base_addr
func (recv_ *CpDmaHalContextT) CpDmaHalSetDescBaseAddr(outlink_base c.IntptrT, inlink_base c.IntptrT) {
}

/**
 * @brief Start mem2mem DMA state machine
 */
// llgo:link (*CpDmaHalContextT).CpDmaHalStart C.cp_dma_hal_start
func (recv_ *CpDmaHalContextT) CpDmaHalStart() {
}

/**
 * @brief Stop mem2mem DMA state machine
 */
// llgo:link (*CpDmaHalContextT).CpDmaHalStop C.cp_dma_hal_stop
func (recv_ *CpDmaHalContextT) CpDmaHalStop() {
}

/**
 * @brief Get interrupt status word
 *
 * @return uint32_t Interrupt status
 */
// llgo:link (*CpDmaHalContextT).CpDmaHalGetIntrStatus C.cp_dma_hal_get_intr_status
func (recv_ *CpDmaHalContextT) CpDmaHalGetIntrStatus() c.Uint32T {
	return 0
}

/**
 * @brief Clear interrupt mask
 *
 * @param mask interrupt mask
 */
// llgo:link (*CpDmaHalContextT).CpDmaHalClearIntrStatus C.cp_dma_hal_clear_intr_status
func (recv_ *CpDmaHalContextT) CpDmaHalClearIntrStatus(mask c.Uint32T) {
}

/**@{*/
/**
 * @brief Give the owner of descriptors between [start_desc, end_desc] to DMA, and restart DMA HW engine
 *
 * @param hal HAL layer context
 * @param start_desc The first descriptor that carries one transaction
 * @param end_desc The last descriptor that carries one transaction
 */
// llgo:link (*CpDmaHalContextT).CpDmaHalRestartTx C.cp_dma_hal_restart_tx
func (recv_ *CpDmaHalContextT) CpDmaHalRestartTx() {
}

// llgo:link (*CpDmaHalContextT).CpDmaHalRestartRx C.cp_dma_hal_restart_rx
func (recv_ *CpDmaHalContextT) CpDmaHalRestartRx() {
}
