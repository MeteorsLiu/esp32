package esp_hid

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief HIDH structure forward declaration
 */

type EspHidhDevS struct {
	Unused [8]uint8
}
type EspHidhDevT EspHidhDevS
type EspHidhEventT c.Int

const (
	ESP_HIDH_ANY_EVENT          EspHidhEventT = -1
	ESP_HIDH_OPEN_EVENT         EspHidhEventT = 0
	ESP_HIDH_BATTERY_EVENT      EspHidhEventT = 1
	ESP_HIDH_INPUT_EVENT        EspHidhEventT = 2
	ESP_HIDH_FEATURE_EVENT      EspHidhEventT = 3
	ESP_HIDH_CLOSE_EVENT        EspHidhEventT = 4
	ESP_HIDH_START_EVENT        EspHidhEventT = 5
	ESP_HIDH_STOP_EVENT         EspHidhEventT = 6
	ESP_HIDH_CONN_REQUEST_EVENT EspHidhEventT = 7
	ESP_HIDH_MAX_EVENT          EspHidhEventT = 8
)

/**
 * @brief HIDH callback parameters union
 */

type EspHidhEventDataT struct {
	Feature struct {
		Dev       *EspHidhDevT
		Usage     EspHidUsageT
		ReportId  c.Uint16T
		Length    c.Uint16T
		Data      *c.Uint8T
		MapIndex  c.Uint8T
		Status    EspErrT
		TransType EspHidTransTypeT
	}
}

type EspHidhConfigT struct {
	Callback       EspEventHandlerT
	EventStackSize c.Uint16T
	CallbackArg    c.Pointer
}

/**
 * @brief Initialize the HID Host component
 * @param config           : pointer to esp_hidh_config_t structure
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhConfigT).EspHidhInit C.esp_hidh_init
func (recv_ *EspHidhConfigT) EspHidhInit() EspErrT {
	return 0
}

/**
 * @brief De-initialize the HID Host component
 *
 * @return: ESP_OK on success
 */
//go:linkname EspHidhDeinit C.esp_hidh_deinit
func EspHidhDeinit() EspErrT

/**
 * @brief Close HID Device
 * @param dev : pointer to the device
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevClose C.esp_hidh_dev_close
func (recv_ *EspHidhDevT) EspHidhDevClose() EspErrT {
	return 0
}

/**
 * @brief Free HID Device Memory
 *        This function MUST be called when handling ESP_HIDH_CLOSE_EVENT
 *        Only then all memory used for the device will be freed.
 * @param dev : pointer to the device
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevFree C.esp_hidh_dev_free
func (recv_ *EspHidhDevT) EspHidhDevFree() EspErrT {
	return 0
}

/**
 * @brief Check if the device still exists.
 * @param dev : pointer to the device
 *
 * @return: true if exists
 */
// llgo:link (*EspHidhDevT).EspHidhDevExists C.esp_hidh_dev_exists
func (recv_ *EspHidhDevT) EspHidhDevExists() bool {
	return false
}

/**
 * @brief Send an OUTPUT report to the device
 * @param dev       : pointer to the device
 * @param map_index : index of the device report map
 * @param report_id : id of the HID OUTPUT report
 * @param data      : pointer to the data to send
 * @param length    : length of the data to send
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevOutputSet C.esp_hidh_dev_output_set
func (recv_ *EspHidhDevT) EspHidhDevOutputSet(map_index c.SizeT, report_id c.SizeT, data *c.Uint8T, length c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Send a FEATURE report to the device
 * @param dev       : pointer to the device
 * @param map_index : index of the device report map
 * @param report_id : id of the HID FEATURE report
 * @param data      : pointer to the data to send
 * @param length    : length of the data to send
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevFeatureSet C.esp_hidh_dev_feature_set
func (recv_ *EspHidhDevT) EspHidhDevFeatureSet(map_index c.SizeT, report_id c.SizeT, data *c.Uint8T, length c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Get the value a FEATURE report from the device
 * @param dev       : pointer to the device
 * @param map_index : index of the device report map
 * @param report_id : id of the HID FEATURE report
 * @param max_len   : size of the buffer that will hold the data
 * @param data      : pointer to the data buffer
 * @param length    : pointer to the value that will be set to the number of bytes received
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevFeatureGet C.esp_hidh_dev_feature_get
func (recv_ *EspHidhDevT) EspHidhDevFeatureGet(map_index c.SizeT, report_id c.SizeT, max_len c.SizeT, data *c.Uint8T, length *c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Set_Report command.
 * @note For now, this function used only for Classic Bluetooth.
 *
 * @param dev           : pointer to the device
 * @param map_index     : index of the device report map
 * @param report_id     : id of the HID FEATURE report
 * @param report_type   : report type, defines in `esp_hid_common.h`
 * @param data          : pointer to the data to send
 * @param length        : length of the data to send
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevSetReport C.esp_hidh_dev_set_report
func (recv_ *EspHidhDevT) EspHidhDevSetReport(map_index c.SizeT, report_id c.SizeT, report_type c.Int, data *c.Uint8T, length c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Get_Report command.
 * @note For now, this function used only for Classic Bluetooth.
 *
 * @param dev           : pointer to the device
 * @param map_index     : index of the device report map
 * @param report_id     : id of the HID FEATURE report
 * @param report_type   : report type, defines in `esp_hid_common.h`
 * @param max_len       : size of the buffer that will hold the data
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevGetReport C.esp_hidh_dev_get_report
func (recv_ *EspHidhDevT) EspHidhDevGetReport(map_index c.SizeT, report_id c.SizeT, report_type c.Int, max_len c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Get_Idle Command.
 * @note For now, this function used only for Classic Bluetooth.
 *
 * @param dev               : pointer to the device
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevGetIdle C.esp_hidh_dev_get_idle
func (recv_ *EspHidhDevT) EspHidhDevGetIdle() EspErrT {
	return 0
}

/**
 * @brief Set_Idle Command.
 * @note For now, this function used only for Classic Bluetooth.
 *
 * @param dev           : pointer to the device
 * @param idle_time     : idle_time
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevSetIdle C.esp_hidh_dev_set_idle
func (recv_ *EspHidhDevT) EspHidhDevSetIdle(idle_time c.Uint8T) EspErrT {
	return 0
}

/**
 * @brief Get_Protocol Command.
 * @note For now, this function used only for Classic Bluetooth.
 *
 * @param dev               : pointer to the device
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevGetProtocol C.esp_hidh_dev_get_protocol
func (recv_ *EspHidhDevT) EspHidhDevGetProtocol() EspErrT {
	return 0
}

/**
 * @brief Set_Protocol Command.
 * @note For now, this function used only for Classic Bluetooth.
 *
 * @param dev           : pointer to the device
 * @param protocol_mode : protocol_mode
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevSetProtocol C.esp_hidh_dev_set_protocol
func (recv_ *EspHidhDevT) EspHidhDevSetProtocol(protocol_mode c.Uint8T) EspErrT {
	return 0
}

/**
 * @brief Dump the properties of HID Device to UART
 * @param dev : pointer to the HID Device
 * @param fp  : pointer to the output file
 */
// llgo:link (*EspHidhDevT).EspHidhDevDump C.esp_hidh_dev_dump
func (recv_ *EspHidhDevT) EspHidhDevDump(fp *c.FILE) {
}

/**
 * @brief Get the BT Device Address of a HID Device
 * @param dev : pointer to the HID Device
 *
 * @return: pointer to the BDA byte array or NULL
 */
// llgo:link (*EspHidhDevT).EspHidhDevBdaGet C.esp_hidh_dev_bda_get
func (recv_ *EspHidhDevT) EspHidhDevBdaGet() *c.Uint8T {
	return nil
}

/**
 * @brief Get the HID Device Transport
 * @param dev : pointer to the HID Device
 *
 * @return: the transport of the connected device or ESP_HID_TRANSPORT_MAX
 */
// llgo:link (*EspHidhDevT).EspHidhDevTransportGet C.esp_hidh_dev_transport_get
func (recv_ *EspHidhDevT) EspHidhDevTransportGet() EspHidTransportT {
	return 0
}

/**
 * @brief Get the HID Device Cofiguration
 * @param dev : pointer to the HID Device
 *
 * @return: pointer to the config structure or NULL
 */
// llgo:link (*EspHidhDevT).EspHidhDevConfigGet C.esp_hidh_dev_config_get
func (recv_ *EspHidhDevT) EspHidhDevConfigGet() *EspHidDeviceConfigT {
	return nil
}

/**
 * @brief Get the name of a HID Device
 * @param dev : pointer to the HID Device
 *
 * @return: pointer to the character array or NULL
 */
// llgo:link (*EspHidhDevT).EspHidhDevNameGet C.esp_hidh_dev_name_get
func (recv_ *EspHidhDevT) EspHidhDevNameGet() *c.Char {
	return nil
}

/**
 * @brief Get the manufacturer of a HID Device
 * @param dev : pointer to the HID Device
 *
 * @return: pointer to the character array
 */
// llgo:link (*EspHidhDevT).EspHidhDevManufacturerGet C.esp_hidh_dev_manufacturer_get
func (recv_ *EspHidhDevT) EspHidhDevManufacturerGet() *c.Char {
	return nil
}

/**
 * @brief Get the serial number of a HID Device
 * @param dev : pointer to the HID Device
 *
 * @return: pointer to the character array or NULL
 */
// llgo:link (*EspHidhDevT).EspHidhDevSerialGet C.esp_hidh_dev_serial_get
func (recv_ *EspHidhDevT) EspHidhDevSerialGet() *c.Char {
	return nil
}

/**
 * @brief Get the VID of a HID Device
 * @param dev : pointer to the HID Device
 *
 * @return: the VID value
 */
// llgo:link (*EspHidhDevT).EspHidhDevVendorIdGet C.esp_hidh_dev_vendor_id_get
func (recv_ *EspHidhDevT) EspHidhDevVendorIdGet() c.Uint16T {
	return 0
}

/**
 * @brief Get the PID of a HID Device
 * @param dev : pointer to the HID Device
 *
 * @return: the PID value
 */
// llgo:link (*EspHidhDevT).EspHidhDevProductIdGet C.esp_hidh_dev_product_id_get
func (recv_ *EspHidhDevT) EspHidhDevProductIdGet() c.Uint16T {
	return 0
}

/**
 * @brief Get the version HID Device
 * @param dev : pointer to the HID Device
 *
 * @return: the version value
 */
// llgo:link (*EspHidhDevT).EspHidhDevVersionGet C.esp_hidh_dev_version_get
func (recv_ *EspHidhDevT) EspHidhDevVersionGet() c.Uint16T {
	return 0
}

/**
 * @brief Get the calculated HID Device usage type
 * @param dev : pointer to the HID Device
 *
 * @return: the hid usage type
 */
// llgo:link (*EspHidhDevT).EspHidhDevUsageGet C.esp_hidh_dev_usage_get
func (recv_ *EspHidhDevT) EspHidhDevUsageGet() EspHidUsageT {
	return 0
}

/**
 * @brief Get an array of all reports found on a device
 * @param dev           : pointer to the device
 * @param num_reports   : pointer to the value that will be set to the number of reports
 * @param reports       : location to set to the pointer of the reports array
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevReportsGet C.esp_hidh_dev_reports_get
func (recv_ *EspHidhDevT) EspHidhDevReportsGet(num_reports *c.SizeT, reports **EspHidReportItemT) EspErrT {
	return 0
}

/**
 * @brief Get an array of the report maps found on a device
 * @param dev        : pointer to the device
 * @param num_maps   : pointer to the value that will be set to the number of report maps found
 * @param maps       : location to set to the pointer of the report maps array
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidhDevT).EspHidhDevReportMapsGet C.esp_hidh_dev_report_maps_get
func (recv_ *EspHidhDevT) EspHidhDevReportMapsGet(num_maps *c.SizeT, maps **EspHidRawReportMapT) EspErrT {
	return 0
}
