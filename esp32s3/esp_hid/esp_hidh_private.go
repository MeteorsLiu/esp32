package esp_hid

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief HIDH device report data
 */

type EspHidhDevReportS struct {
	Next         *EspHidhDevReportS
	MapIndex     c.Uint8T
	ReportId     c.Uint8T
	ReportType   c.Uint8T
	ProtocolMode c.Uint8T
	ValueLen     c.SizeT
	Usage        EspHidUsageT
	Handle       c.Uint16T
	CccHandle    c.Uint16T
	Permissions  c.Uint8T
}
type EspHidhDevReportT EspHidhDevReportS

/**
 * @brief Get HIDH event loop
 *
 * Transport layers will post events into the loop
 *
 * @return esp_event_loop_handle_t Handle to HIDH event loop
 */
//go:linkname EspHidhGetEventLoop C.esp_hidh_get_event_loop
func EspHidhGetEventLoop() EspEventLoopHandleT

/**
 * @brief Allocate HIDH device
 *
 * The resources can be freed by esp_hidh_dev_free_inner()
 *
 * @return esp_hidh_dev_t* Pointer to newly allocated HIDH device
 */
//go:linkname EspHidhDevMalloc C.esp_hidh_dev_malloc
func EspHidhDevMalloc() *EspHidhDevT

/**
 * @brief Free HIDH device
 *
 * @param[in] dev Device handle obtained from esp_hidh_dev_malloc()
 * @return
 *     - ESP_OK:   Success
 *     - ESP_FAIL: Parameter is NULL or it is not a valid HIDH device
 */
// llgo:link (*EspHidhDevT).EspHidhDevFreeInner C.esp_hidh_dev_free_inner
func (recv_ *EspHidhDevT) EspHidhDevFreeInner() EspErrT {
	return 0
}

// llgo:link (*EspHidhDevT).EspHidhDevGetReportByIdTypeProto C.esp_hidh_dev_get_report_by_id_type_proto
func (recv_ *EspHidhDevT) EspHidhDevGetReportByIdTypeProto(map_index c.SizeT, report_id c.SizeT, report_type c.Int, protocol_mode c.Uint8T) *EspHidhDevReportT {
	return nil
}

// llgo:link (*EspHidhDevT).EspHidhDevGetReportByIdAndType C.esp_hidh_dev_get_report_by_id_and_type
func (recv_ *EspHidhDevT) EspHidhDevGetReportByIdAndType(map_index c.SizeT, report_id c.SizeT, report_type c.Int) *EspHidhDevReportT {
	return nil
}

// llgo:link (*EspHidhDevT).EspHidhDevGetInputReportByLenAndProto C.esp_hidh_dev_get_input_report_by_len_and_proto
func (recv_ *EspHidhDevT) EspHidhDevGetInputReportByLenAndProto(len c.SizeT, protocol_mode c.Int) *EspHidhDevReportT {
	return nil
}

// llgo:link (*EspHidhDevT).EspHidhDevGetInputReportByIdAndProto C.esp_hidh_dev_get_input_report_by_id_and_proto
func (recv_ *EspHidhDevT) EspHidhDevGetInputReportByIdAndProto(report_id c.SizeT, protocol_mode c.Int) *EspHidhDevReportT {
	return nil
}

// llgo:link (*EspHidhDevT).EspHidhDevGetInputReportByProtoAndData C.esp_hidh_dev_get_input_report_by_proto_and_data
func (recv_ *EspHidhDevT) EspHidhDevGetInputReportByProtoAndData(protocol_mode c.Int, len c.SizeT, data *c.Uint8T, has_report_id *bool) *EspHidhDevReportT {
	return nil
}

// llgo:link (*EspHidhDevT).EspHidhDevGetReportByHandle C.esp_hidh_dev_get_report_by_handle
func (recv_ *EspHidhDevT) EspHidhDevGetReportByHandle(handle c.Uint16T) *EspHidhDevReportT {
	return nil
}

//go:linkname EspHidhPreprocessEventHandler C.esp_hidh_preprocess_event_handler
func EspHidhPreprocessEventHandler(event_handler_arg c.Pointer, event_base EspEventBaseT, event_id c.Int32T, event_data c.Pointer)

//go:linkname EspHidhPostprocessEventHandler C.esp_hidh_postprocess_event_handler
func EspHidhPostprocessEventHandler(event_handler_arg c.Pointer, event_base EspEventBaseT, event_id c.Int32T, event_data c.Pointer)

// llgo:link (*EspHidhDevT).EspHidhDevLock C.esp_hidh_dev_lock
func (recv_ *EspHidhDevT) EspHidhDevLock() {
}

// llgo:link (*EspHidhDevT).EspHidhDevUnlock C.esp_hidh_dev_unlock
func (recv_ *EspHidhDevT) EspHidhDevUnlock() {
}

// llgo:link (*EspHidhDevT).EspHidhDevWait C.esp_hidh_dev_wait
func (recv_ *EspHidhDevT) EspHidhDevWait() {
}

// llgo:link (*EspHidhDevT).EspHidhDevSend C.esp_hidh_dev_send
func (recv_ *EspHidhDevT) EspHidhDevSend() {
}
