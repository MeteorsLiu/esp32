package esp_adc

import _ "unsafe"

/**
 * @brief Initialize DMA interrupt event
 *
 * @param[in] adc_ctx ADC continuous mode driver handle
 *
 * @return
 *      - ESP_OK:                On success
 */
//go:linkname AdcDmaIntrEventInit C.adc_dma_intr_event_init
func AdcDmaIntrEventInit(adc_ctx *AdcContinuousCtxT) EspErrT

/**
 * @brief Initialize DMA on ADC continuous mode
 *
 * @param[in] adc_dma ADC DMA context
 *
 * @return
 *      - ESP_OK:                On success
 */
//go:linkname AdcDmaInit C.adc_dma_init
func AdcDmaInit(adc_dma *AdcDmaT) EspErrT

/**
 * @brief Deinitialize DMA on ADC continuous mode
 *
 * @param[in] adc_dma ADC DMA context
 *
 * @return
 *      - ESP_OK:                On success
 */
//go:linkname AdcDmaDeinit C.adc_dma_deinit
func AdcDmaDeinit(adc_dma AdcDmaT) EspErrT

/**
 * @brief Start DMA on ADC continuous mode
 *
 * @param[in] adc_dma ADC DMA context
 * @param[in] addr    ADC DMA descriptor
 *
 * @return
 *      - ESP_OK:                On success
 */
//go:linkname AdcDmaStart C.adc_dma_start
func AdcDmaStart(adc_dma AdcDmaT, addr *DmaDescriptorT) EspErrT

/**
 * @brief Stop DMA on ADC continuous mode
 *
 * @param[in] adc_dma ADC DMA context
 *
 * @return
 *      - ESP_OK:                On success
 */
//go:linkname AdcDmaStop C.adc_dma_stop
func AdcDmaStop(adc_dma AdcDmaT) EspErrT

/**
 * @brief Reset DMA on ADC continuous mode
 *
 * @param[in] adc_dma ADC DMA context
 *
 * @return
 *      - ESP_OK:                On success
 */
//go:linkname AdcDmaReset C.adc_dma_reset
func AdcDmaReset(adc_dma AdcDmaT) EspErrT
