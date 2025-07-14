package esp_hid

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspHiddEventT c.Int

const (
	ESP_HIDD_ANY_EVENT           EspHiddEventT = -1
	ESP_HIDD_START_EVENT         EspHiddEventT = 0
	ESP_HIDD_CONNECT_EVENT       EspHiddEventT = 1
	ESP_HIDD_PROTOCOL_MODE_EVENT EspHiddEventT = 2
	ESP_HIDD_CONTROL_EVENT       EspHiddEventT = 3
	ESP_HIDD_OUTPUT_EVENT        EspHiddEventT = 4
	ESP_HIDD_FEATURE_EVENT       EspHiddEventT = 5
	ESP_HIDD_DISCONNECT_EVENT    EspHiddEventT = 6
	ESP_HIDD_STOP_EVENT          EspHiddEventT = 7
	ESP_HIDD_MAX_EVENT           EspHiddEventT = 8
)

/**
 * @brief HIDD structure forward declaration
 */

type EspHiddDevS struct {
	Unused [8]uint8
}
type EspHiddDevT EspHiddDevS

/**
 * @brief HIDD callback parameters union
 */

type EspHiddEventDataT struct {
	Output struct {
		Dev      *EspHiddDevT
		Usage    EspHidUsageT
		ReportId c.Uint16T
		Length   c.Uint16T
		Data     *c.Uint8T
		MapIndex c.Uint8T
	}
}

/**
 * @brief Init HID Device
 * @param       config   : configuration for the device
 * @param       transport: protocol that the device will use (ESP_HID_TRANSPORT_BLE/BT/USB)
 * @param       callback : function to call when events for this device are generated.
 *                         Can be NULL, but will miss the START event.
 * @param[out]  dev      : location to return the pointer to the device structure
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHidDeviceConfigT).EspHiddDevInit C.esp_hidd_dev_init
func (recv_ *EspHidDeviceConfigT) EspHiddDevInit(transport EspHidTransportT, callback EspEventHandlerT, dev **EspHiddDevT) EspErrT {
	return 0
}

/**
 * @brief Deinit HID Device
 * @param dev : pointer to the device to deinit
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHiddDevT).EspHiddDevDeinit C.esp_hidd_dev_deinit
func (recv_ *EspHiddDevT) EspHiddDevDeinit() EspErrT {
	return 0
}

/**
 * @brief Get the HID Device Transport
 * @param dev : pointer to the HID Device
 *
 * @return: the transport of the connected device or ESP_HID_TRANSPORT_MAX
 */
// llgo:link (*EspHiddDevT).EspHiddDevTransportGet C.esp_hidd_dev_transport_get
func (recv_ *EspHiddDevT) EspHiddDevTransportGet() EspHidTransportT {
	return 0
}

/**
 * @brief Check if HID Device is connected
 * @param dev : pointer to the device
 *
 * @return: true if the device is connected
 */
// llgo:link (*EspHiddDevT).EspHiddDevConnected C.esp_hidd_dev_connected
func (recv_ *EspHiddDevT) EspHiddDevConnected() bool {
	return false
}

/**
 * @brief Set the battery level reported by the HID Device
 * @param dev   : pointer to the device
 * @param level : battery level (0-100)
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHiddDevT).EspHiddDevBatterySet C.esp_hidd_dev_battery_set
func (recv_ *EspHiddDevT) EspHiddDevBatterySet(level c.Uint8T) EspErrT {
	return 0
}

/**
 * @brief Send an INPUT report to the host
 * @param dev       : pointer to the device
 * @param map_index : index of the device report map in the init config
 * @param report_id : id of the HID INPUT report
 * @param data      : pointer to the data to send
 * @param length    : length of the data to send
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHiddDevT).EspHiddDevInputSet C.esp_hidd_dev_input_set
func (recv_ *EspHiddDevT) EspHiddDevInputSet(map_index c.SizeT, report_id c.SizeT, data *c.Uint8T, length c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Send a FEATURE report to the host
 * @param dev       : pointer to the device
 * @param map_index : index of the device report map in the init config
 * @param report_id : id of the HID FEATURE report
 * @param data      : pointer to the data to send
 * @param length    : length of the data to send
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHiddDevT).EspHiddDevFeatureSet C.esp_hidd_dev_feature_set
func (recv_ *EspHiddDevT) EspHiddDevFeatureSet(map_index c.SizeT, report_id c.SizeT, data *c.Uint8T, length c.SizeT) EspErrT {
	return 0
}

/**
 * @brief Register function to listen for device events
 * @param dev       : pointer to the device
 * @param callback  : event handler function
 * @param event     : event to listen for (ESP_HIDD_ANY_EVENT for all)
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHiddDevT).EspHiddDevEventHandlerRegister C.esp_hidd_dev_event_handler_register
func (recv_ *EspHiddDevT) EspHiddDevEventHandlerRegister(callback EspEventHandlerT, event EspHiddEventT) EspErrT {
	return 0
}

/**
 * @brief Unregister function that is listening for device events
 * @param dev       : pointer to the device
 * @param callback  : event handler function
 * @param event     : event that is listening for (ESP_HIDD_ANY_EVENT for all)
 *
 * @return: ESP_OK on success
 */
// llgo:link (*EspHiddDevT).EspHiddDevEventHandlerUnregister C.esp_hidd_dev_event_handler_unregister
func (recv_ *EspHiddDevT) EspHiddDevEventHandlerUnregister(callback EspEventHandlerT, event EspHiddEventT) EspErrT {
	return 0
}
