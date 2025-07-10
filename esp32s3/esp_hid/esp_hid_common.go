package esp_hid

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const HID_RM_INPUT = 0x80
const HID_RM_OUTPUT = 0x90
const HID_RM_FEATURE = 0xb0
const HID_RM_COLLECTION = 0xa0
const HID_RM_END_COLLECTION = 0xc0
const HID_RM_USAGE_PAGE = 0x04
const HID_RM_LOGICAL_MINIMUM = 0x14
const HID_RM_LOGICAL_MAXIMUM = 0x24
const HID_RM_PHYSICAL_MINIMUM = 0x34
const HID_RM_PHYSICAL_MAXIMUM = 0x44
const HID_RM_UNIT_EXPONENT = 0x54
const HID_RM_UNIT = 0x64
const HID_RM_REPORT_SIZE = 0x74
const HID_RM_REPORT_ID = 0x84
const HID_RM_REPORT_COUNT = 0x94
const HID_RM_PUSH = 0xa4
const HID_RM_POP = 0xb4
const HID_RM_USAGE = 0x08
const HID_RM_USAGE_MINIMUM = 0x18
const HID_RM_USAGE_MAXIMUM = 0x28
const HID_RM_DESIGNATOR_INDEX = 0x38
const HID_RM_DESIGNATOR_MINIMUM = 0x48
const HID_RM_DESIGNATOR_MAXIMUM = 0x58
const HID_RM_STRING_INDEX = 0x78
const HID_RM_STRING_MINIMUM = 0x88
const HID_RM_STRING_MAXIMUM = 0x98
const HID_RM_DELIMITER = 0xa8
const HID_USAGE_PAGE_GENERIC_DESKTOP = 0x01
const HID_USAGE_KEYBOARD = 0x06
const HID_USAGE_MOUSE = 0x02
const HID_USAGE_JOYSTICK = 0x04
const HID_USAGE_GAMEPAD = 0x05
const HID_USAGE_PAGE_CONSUMER_DEVICE = 0x0C
const HID_USAGE_CONSUMER_CONTROL = 0x01
const ESP_HID_COD_MIN_KEYBOARD = 0x10
const ESP_HID_COD_MIN_MOUSE = 0x20
const ESP_HID_APPEARANCE_GENERIC = 0x03C0
const ESP_HID_APPEARANCE_KEYBOARD = 0x03C1
const ESP_HID_APPEARANCE_MOUSE = 0x03C2
const ESP_HID_APPEARANCE_JOYSTICK = 0x03C3
const ESP_HID_APPEARANCE_GAMEPAD = 0x03C4
const ESP_HID_REPORT_TYPE_INPUT = 1
const ESP_HID_REPORT_TYPE_OUTPUT = 2
const ESP_HID_REPORT_TYPE_FEATURE = 3
const ESP_HID_PROTOCOL_MODE_BOOT = 0x00
const ESP_HID_PROTOCOL_MODE_REPORT = 0x01
const ESP_HID_FLAGS_REMOTE_WAKE = 0x01
const ESP_HID_FLAGS_NORMALLY_CONNECTABLE = 0x02
const ESP_HID_CONTROL_SUSPEND = 0x00
const ESP_HID_CONTROL_EXIT_SUSPEND = 0x01
const ESP_HID_CCC_NOTIFICATIONS_ENABLED = 0x01
const ESP_HID_CCC_INDICATIONS_ENABLED = 0x02
const BT_HID_DEVICE_TASK_SIZE_BT = 2048
const BT_HID_DEVICE_TASK_SIZE_BLE = 4096

type EspHidTransportT c.Int

const (
	ESP_HID_TRANSPORT_BT  EspHidTransportT = 0
	ESP_HID_TRANSPORT_BLE EspHidTransportT = 1
	ESP_HID_TRANSPORT_USB EspHidTransportT = 2
	ESP_HID_TRANSPORT_MAX EspHidTransportT = 3
)

type EspHidUsageT c.Int

const (
	ESP_HID_USAGE_GENERIC  EspHidUsageT = 0
	ESP_HID_USAGE_KEYBOARD EspHidUsageT = 1
	ESP_HID_USAGE_MOUSE    EspHidUsageT = 2
	ESP_HID_USAGE_JOYSTICK EspHidUsageT = 4
	ESP_HID_USAGE_GAMEPAD  EspHidUsageT = 8
	ESP_HID_USAGE_TABLET   EspHidUsageT = 16
	ESP_HID_USAGE_CCONTROL EspHidUsageT = 32
	ESP_HID_USAGE_VENDOR   EspHidUsageT = 64
)

type EspHidCodMinT c.Int

const (
	ESP_HID_COD_MIN_GENERIC     EspHidCodMinT = 0
	ESP_HID_COD_MIN_JOYSTICK    EspHidCodMinT = 1
	ESP_HID_COD_MIN_GAMEPAD     EspHidCodMinT = 2
	ESP_HID_COD_MIN_REMOTE      EspHidCodMinT = 3
	ESP_HID_COD_MIN_SENSOR      EspHidCodMinT = 4
	ESP_HID_COD_MIN_TABLET      EspHidCodMinT = 5
	ESP_HID_COD_MIN_CARD_READER EspHidCodMinT = 6
	ESP_HID_COD_MIN_MAX         EspHidCodMinT = 7
)

type EspHidTransTypeT c.Int

const (
	ESP_HID_TRANS_HANDSHAKE    EspHidTransTypeT = 0
	ESP_HID_TRANS_CONTROL      EspHidTransTypeT = 1
	ESP_HID_TRANS_GET_REPORT   EspHidTransTypeT = 4
	ESP_HID_TRANS_SET_REPORT   EspHidTransTypeT = 5
	ESP_HID_TRANS_GET_PROTOCOL EspHidTransTypeT = 6
	ESP_HID_TRANS_SET_PROTOCOL EspHidTransTypeT = 7
	ESP_HID_TRANS_GET_IDLE     EspHidTransTypeT = 8
	ESP_HID_TRANS_SET_IDLE     EspHidTransTypeT = 9
	ESP_HID_TRANS_DATA         EspHidTransTypeT = 10
	ESP_HID_TRANS_DATAC        EspHidTransTypeT = 11
	ESP_HID_TRANS_MAX          EspHidTransTypeT = 12
)

/**
 * @brief HID report item structure
 */

type EspHidReportItemT struct {
	MapIndex     c.Uint8T
	ReportId     c.Uint8T
	ReportType   c.Uint8T
	ProtocolMode c.Uint8T
	Usage        EspHidUsageT
	ValueLen     c.Uint16T
}

/**
 * @brief HID parsed report map structure
 */

type EspHidReportMapT struct {
	Usage      EspHidUsageT
	Appearance c.Uint16T
	ReportsLen c.Uint8T
	Reports    *EspHidReportItemT
}

/**
 * @brief HID raw report map structure
 */

type EspHidRawReportMapT struct {
	Data *c.Uint8T
	Len  c.Uint16T
}

/**
 * @brief HID device config structure
 */

type EspHidDeviceConfigT struct {
	VendorId         c.Uint16T
	ProductId        c.Uint16T
	Version          c.Uint16T
	DeviceName       *c.Char
	ManufacturerName *c.Char
	SerialNumber     *c.Char
	ReportMaps       *EspHidRawReportMapT
	ReportMapsLen    c.Uint8T
}

/**
 * @brief HID device address
 */

type EspHidAddressT struct {
	Bda [6]c.Uint8T
}

/*
 * @brief Parse RAW HID report map
 *        It is a responsibility of the user to free the parsed report map,
 *        when it's no longer needed. Use esp_hid_free_report_map
 * @param hid_rm      : pointer to the hid report map data
 * @param hid_rm_len  : length to the hid report map data
 *
 * @return: pointer to the parsed report map
 */
//go:linkname EspHidParseReportMap C.esp_hid_parse_report_map
func EspHidParseReportMap(hid_rm *c.Uint8T, hid_rm_len c.SizeT) *EspHidReportMapT

/*
 * @brief Free parsed HID report map
 * @param map      : pointer to the parsed hid report map
 */
// llgo:link (*EspHidReportMapT).EspHidFreeReportMap C.esp_hid_free_report_map
func (recv_ *EspHidReportMapT) EspHidFreeReportMap() {
}

/**
 * @brief Calculate the HID Device usage type from the BLE Appearance
 * @param appearance : BLE Appearance value
 *
 * @return: the hid usage type
 */
//go:linkname EspHidUsageFromAppearance C.esp_hid_usage_from_appearance
func EspHidUsageFromAppearance(appearance c.Uint16T) EspHidUsageT

/**
 * @brief Calculate the HID Device usage type from the BT CoD
 * @param cod : BT CoD value
 *
 * @return: the hid usage type
 */
//go:linkname EspHidUsageFromCod C.esp_hid_usage_from_cod
func EspHidUsageFromCod(cod c.Uint32T) EspHidUsageT

/**
 * @brief Convert device usage type to string
 * @param usage : The HID usage type to convert
 *
 * @return: a pointer to the string or NULL
 */
// llgo:link EspHidUsageT.EspHidUsageStr C.esp_hid_usage_str
func (recv_ EspHidUsageT) EspHidUsageStr() *c.Char {
	return nil
}

/**
 * @brief Convert HID protocol mode to string
 * @param protocol_mode : The HID protocol mode to convert
 *                        BOOT/REPORT
 *
 * @return: a pointer to the string or NULL
 */
//go:linkname EspHidProtocolModeStr C.esp_hid_protocol_mode_str
func EspHidProtocolModeStr(protocol_mode c.Uint8T) *c.Char

/**
 * @brief Convert HID report type to string
 * @param report_type : The HID report type to convert
 *                      INPUT/OUTPUT/FEATURE
 *
 * @return: a pointer to the string or NULL
 */
//go:linkname EspHidReportTypeStr C.esp_hid_report_type_str
func EspHidReportTypeStr(report_type c.Uint8T) *c.Char

/**
 * @brief Convert BT CoD major to string
 * @param cod_major : The CoD major value to convert
 *
 * @return: a pointer to the string or NULL
 */
//go:linkname EspHidCodMajorStr C.esp_hid_cod_major_str
func EspHidCodMajorStr(cod_major c.Uint8T) *c.Char

/**
 * @brief Print BT CoD minor value
 * @param cod_min : The CoD minor value to print
 * @param fp      : pointer to the output file
 */
//go:linkname EspHidCodMinorPrint C.esp_hid_cod_minor_print
func EspHidCodMinorPrint(cod_min c.Uint8T, fp *c.FILE)

/**
 * @brief Convert BLE disconnect reason to string
 * @param reason : The value of the reason
 *
 * @return: a pointer to the string or NULL
 */
// llgo:link EspHidTransportT.EspHidDisconnectReasonStr C.esp_hid_disconnect_reason_str
func (recv_ EspHidTransportT) EspHidDisconnectReasonStr(reason c.Int) *c.Char {
	return nil
}
