package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Configure Pulse Counter unit
 *        @note
 *        This function will disable three events: PCNT_EVT_L_LIM, PCNT_EVT_H_LIM, PCNT_EVT_ZERO.
 *
 * @param pcnt_config Pointer of Pulse Counter unit configure parameter
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver already initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link (*PcntConfigT).PcntUnitConfig C.pcnt_unit_config
func (recv_ *PcntConfigT) PcntUnitConfig() EspErrT {
	return 0
}

/**
 * @brief Get pulse counter value
 *
 * @param pcnt_unit  Pulse Counter unit number
 * @param count Pointer to accept counter value
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntGetCounterValue C.pcnt_get_counter_value
func (recv_ PcntUnitT) PcntGetCounterValue(count *c.Int16T) EspErrT {
	return 0
}

/**
 * @brief Pause PCNT counter of PCNT unit
 *
 * @param pcnt_unit PCNT unit number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntCounterPause C.pcnt_counter_pause
func (recv_ PcntUnitT) PcntCounterPause() EspErrT {
	return 0
}

/**
 * @brief Resume counting for PCNT counter
 *
 * @param pcnt_unit PCNT unit number, select from pcnt_unit_t
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntCounterResume C.pcnt_counter_resume
func (recv_ PcntUnitT) PcntCounterResume() EspErrT {
	return 0
}

/**
 * @brief Clear and reset PCNT counter value to zero
 *
 * @param  pcnt_unit PCNT unit number, select from pcnt_unit_t
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntCounterClear C.pcnt_counter_clear
func (recv_ PcntUnitT) PcntCounterClear() EspErrT {
	return 0
}

/**
 * @brief Enable PCNT interrupt for PCNT unit
 *        @note
 *        Each Pulse counter unit has five watch point events that share the same interrupt.
 *        Configure events with pcnt_event_enable() and pcnt_event_disable()
 *
 * @param pcnt_unit PCNT unit number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntIntrEnable C.pcnt_intr_enable
func (recv_ PcntUnitT) PcntIntrEnable() EspErrT {
	return 0
}

/**
 * @brief Disable PCNT interrupt for PCNT unit
 *
 * @param pcnt_unit PCNT unit number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntIntrDisable C.pcnt_intr_disable
func (recv_ PcntUnitT) PcntIntrDisable() EspErrT {
	return 0
}

/**
 * @brief Enable PCNT event of PCNT unit
 *
 * @param unit PCNT unit number
 * @param evt_type Watch point event type.
 *                 All enabled events share the same interrupt (one interrupt per pulse counter unit).
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntEventEnable C.pcnt_event_enable
func (recv_ PcntUnitT) PcntEventEnable(evt_type PcntEvtTypeT) EspErrT {
	return 0
}

/**
 * @brief Disable PCNT event of PCNT unit
 *
 * @param unit PCNT unit number
 * @param evt_type Watch point event type.
 *                 All enabled events share the same interrupt (one interrupt per pulse counter unit).
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntEventDisable C.pcnt_event_disable
func (recv_ PcntUnitT) PcntEventDisable(evt_type PcntEvtTypeT) EspErrT {
	return 0
}

/**
 * @brief Set PCNT event value of PCNT unit
 *
 * @param unit PCNT unit number
 * @param evt_type Watch point event type.
 *                 All enabled events share the same interrupt (one interrupt per pulse counter unit).
 *
 * @param value Counter value for PCNT event
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntSetEventValue C.pcnt_set_event_value
func (recv_ PcntUnitT) PcntSetEventValue(evt_type PcntEvtTypeT, value c.Int16T) EspErrT {
	return 0
}

/**
 * @brief Get PCNT event value of PCNT unit
 *
 * @param unit PCNT unit number
 * @param evt_type Watch point event type.
 *                 All enabled events share the same interrupt (one interrupt per pulse counter unit).
 * @param value Pointer to accept counter value for PCNT event
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntGetEventValue C.pcnt_get_event_value
func (recv_ PcntUnitT) PcntGetEventValue(evt_type PcntEvtTypeT, value *c.Int16T) EspErrT {
	return 0
}

/**
 * @brief Get PCNT event status of PCNT unit
 *
 * @param unit PCNT unit number
 * @param status Pointer to accept event status word
 * @return
 *      - ESP_OK Success
 *      - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *      - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntGetEventStatus C.pcnt_get_event_status
func (recv_ PcntUnitT) PcntGetEventStatus(status *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Unregister PCNT interrupt handler (registered by pcnt_isr_register), the handler is an ISR.
 *        The handler will be attached to the same CPU core that this function is running on.
 *        If the interrupt service is registered by pcnt_isr_service_install, please call pcnt_isr_service_uninstall instead
 *
 * @param handle handle to unregister the ISR service.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_NOT_FOUND Can not find the interrupt that matches the flags.
 *     - ESP_ERR_INVALID_ARG Function pointer error.
 */
//go:linkname PcntIsrUnregister C.pcnt_isr_unregister
func PcntIsrUnregister(handle PcntIsrHandleT) EspErrT

/**
 * @brief Register PCNT interrupt handler, the handler is an ISR.
 *        The handler will be attached to the same CPU core that this function is running on.
 *        Please do not use pcnt_isr_service_install if this function was called.
 *
 * @param fn Interrupt handler function.
 * @param arg Parameter for handler function
 * @param intr_alloc_flags Flags used to allocate the interrupt. One or multiple (ORred)
 *        ESP_INTR_FLAG_* values. See esp_intr_alloc.h for more info.
 * @param handle Pointer to return handle. If non-NULL, a handle for the interrupt will
 *        be returned here. Calling pcnt_isr_unregister to unregister this ISR service if needed,
 *        but only if the handle is not NULL.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_NOT_FOUND Can not find the interrupt that matches the flags.
 *     - ESP_ERR_INVALID_ARG Function pointer error.
 */
//go:linkname PcntIsrRegister C.pcnt_isr_register
func PcntIsrRegister(fn func(c.Pointer), arg c.Pointer, intr_alloc_flags c.Int, handle *PcntIsrHandleT) EspErrT

/**
 * @brief Configure PCNT pulse signal input pin and control input pin
 *
 * @param unit PCNT unit number
 * @param channel PCNT channel number
 * @param pulse_io Pulse signal input GPIO
 * @param ctrl_io Control signal input GPIO
 *
 * @note  Set the signal input to PCNT_PIN_NOT_USED if unused.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntSetPin C.pcnt_set_pin
func (recv_ PcntUnitT) PcntSetPin(channel PcntChannelT, pulse_io c.Int, ctrl_io c.Int) EspErrT {
	return 0
}

/**
 * @brief Enable PCNT input filter
 *
 * @param unit PCNT unit number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntFilterEnable C.pcnt_filter_enable
func (recv_ PcntUnitT) PcntFilterEnable() EspErrT {
	return 0
}

/**
 * @brief Disable PCNT input filter
 *
 * @param unit PCNT unit number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntFilterDisable C.pcnt_filter_disable
func (recv_ PcntUnitT) PcntFilterDisable() EspErrT {
	return 0
}

/**
 * @brief Set PCNT filter value
 *
 * @param unit PCNT unit number
 * @param filter_val PCNT signal filter value, counter in APB_CLK cycles.
 *        Any pulses lasting shorter than this will be ignored when the filter is enabled.
 *        @note
 *        filter_val is a 10-bit value, so the maximum filter_val should be limited to 1023.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntSetFilterValue C.pcnt_set_filter_value
func (recv_ PcntUnitT) PcntSetFilterValue(filter_val c.Uint16T) EspErrT {
	return 0
}

/**
 * @brief Get PCNT filter value
 *
 * @param unit PCNT unit number
 * @param filter_val Pointer to accept PCNT filter value.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntGetFilterValue C.pcnt_get_filter_value
func (recv_ PcntUnitT) PcntGetFilterValue(filter_val *c.Uint16T) EspErrT {
	return 0
}

/**
 * @brief Set PCNT counter mode
 *
 * @param unit PCNT unit number
 * @param channel PCNT channel number
 * @param pos_mode Counter mode when detecting positive edge
 * @param neg_mode Counter mode when detecting negative edge
 * @param hctrl_mode Counter mode when control signal is high level
 * @param lctrl_mode Counter mode when control signal is low level
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntSetMode C.pcnt_set_mode
func (recv_ PcntUnitT) PcntSetMode(channel PcntChannelT, pos_mode PcntCountModeT, neg_mode PcntCountModeT, hctrl_mode PcntCtrlModeT, lctrl_mode PcntCtrlModeT) EspErrT {
	return 0
}

/**
 * @brief Add ISR handler for specified unit.
 *
 * Call this function after using pcnt_isr_service_install() to
 * install the PCNT driver's ISR handler service.
 *
 * The ISR handlers do not need to be declared with IRAM_ATTR,
 * unless you pass the ESP_INTR_FLAG_IRAM flag when allocating the
 * ISR in pcnt_isr_service_install().
 *
 * This ISR handler will be called from an ISR. So there is a stack
 * size limit (configurable as "ISR stack size" in menuconfig). This
 * limit is smaller compared to a global PCNT interrupt handler due
 * to the additional level of indirection.
 *
 * @param unit PCNT unit number
 * @param isr_handler Interrupt handler function.
 * @param args Parameter for handler function
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntIsrHandlerAdd C.pcnt_isr_handler_add
func (recv_ PcntUnitT) PcntIsrHandlerAdd(isr_handler func(c.Pointer), args c.Pointer) EspErrT {
	return 0
}

/**
 * @brief Install PCNT ISR service.
 * @note We can manage different interrupt service for each unit.
 *       This function will use the default ISR handle service, Calling pcnt_isr_service_uninstall to
 *       uninstall the default service if needed. Please do not use pcnt_isr_register if this function was called.
 *
 * @param intr_alloc_flags Flags used to allocate the interrupt. One or multiple (ORred)
 *        ESP_INTR_FLAG_* values. See esp_intr_alloc.h for more info.
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_NO_MEM No memory to install this service
 *     - ESP_ERR_INVALID_STATE ISR service already installed
 */
//go:linkname PcntIsrServiceInstall C.pcnt_isr_service_install
func PcntIsrServiceInstall(intr_alloc_flags c.Int) EspErrT

/**
 * @brief Uninstall PCNT ISR service, freeing related resources.
 */
//go:linkname PcntIsrServiceUninstall C.pcnt_isr_service_uninstall
func PcntIsrServiceUninstall()

/**
 * @brief Delete ISR handler for specified unit.
 *
 * @param unit PCNT unit number
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE pcnt driver has not been initialized
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
// llgo:link PcntUnitT.PcntIsrHandlerRemove C.pcnt_isr_handler_remove
func (recv_ PcntUnitT) PcntIsrHandlerRemove() EspErrT {
	return 0
}
