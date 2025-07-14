package esp_local_ctrl

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief   Property description data structure, which is to be populated
 *          and passed to the `esp_local_ctrl_add_property()` function
 *
 * Once a property is added, its structure is available for read-only access
 * inside `get_prop_values()` and `set_prop_values()` handlers.
 */

type EspLocalCtrlProp struct {
	Name      *c.Char
	Type      c.Uint32T
	Size      c.SizeT
	Flags     c.Uint32T
	Ctx       c.Pointer
	CtxFreeFn c.Pointer
}
type EspLocalCtrlPropT EspLocalCtrlProp

/**
 * @brief   Property value data structure. This gets passed to the
 *          `get_prop_values()` and `set_prop_values()` handlers for
 *          the purpose of retrieving or setting the present value
 *          of a property.
 */

type EspLocalCtrlPropVal struct {
	Data   c.Pointer
	Size   c.SizeT
	FreeFn c.Pointer
}
type EspLocalCtrlPropValT EspLocalCtrlPropVal

/**
 * @brief   Handlers for receiving and responding to local
 *          control commands for getting and setting properties.
 */

type EspLocalCtrlHandlers struct {
	GetPropValues c.Pointer
	SetPropValues c.Pointer
	UsrCtx        c.Pointer
	UsrCtxFreeFn  c.Pointer
}
type EspLocalCtrlHandlersT EspLocalCtrlHandlers

type EspLocalCtrlTransport struct {
	Unused [8]uint8
}
type EspLocalCtrlTransportT EspLocalCtrlTransport

/**
 * @brief   Function for obtaining HTTPD transport mode
 */
//go:linkname EspLocalCtrlGetTransportHttpd C.esp_local_ctrl_get_transport_httpd
func EspLocalCtrlGetTransportHttpd() *EspLocalCtrlTransportT

type ProtocommBleConfig struct {
	Unused [8]uint8
}
type EspLocalCtrlTransportConfigBleT ProtocommBleConfig
type EspLocalCtrlTransportConfigHttpdT HttpdConfig

/**
 * @brief   Transport mode (BLE / HTTPD) configuration
 */

type EspLocalCtrlTransportConfigT struct {
	Ble *EspLocalCtrlTransportConfigBleT
}
type EspLocalCtrlProtoSec c.Int

const (
	PROTOCOM_SEC0       EspLocalCtrlProtoSec = 0
	PROTOCOM_SEC1       EspLocalCtrlProtoSec = 1
	PROTOCOM_SEC2       EspLocalCtrlProtoSec = 2
	PROTOCOM_SEC_CUSTOM EspLocalCtrlProtoSec = 3
)

type EspLocalCtrlProtoSecT EspLocalCtrlProtoSec
type EspLocalCtrlSecurity1ParamsT ProtocommSecurity1ParamsT
type EspLocalCtrlSecurity2ParamsT ProtocommSecurity2ParamsT

/**
 * Protocom security configs
 */

type EspLocalCtrlProtoSecCfg struct {
	Version      EspLocalCtrlProtoSecT
	CustomHandle c.Pointer
}
type EspLocalCtrlProtoSecCfgT EspLocalCtrlProtoSecCfg

/**
 * @brief   Configuration structure to pass to `esp_local_ctrl_start()`
 */

type EspLocalCtrlConfig struct {
	Transport       *EspLocalCtrlTransportT
	TransportConfig EspLocalCtrlTransportConfigT
	ProtoSec        EspLocalCtrlProtoSecCfgT
	Handlers        EspLocalCtrlHandlersT
	MaxProperties   c.SizeT
}
type EspLocalCtrlConfigT EspLocalCtrlConfig

/**
 * @brief   Start local control service
 *
 * @param[in] config    Pointer to configuration structure
 *
 * @return
 *  - ESP_OK      : Success
 *  - ESP_FAIL    : Failure
 */
// llgo:link (*EspLocalCtrlConfigT).EspLocalCtrlStart C.esp_local_ctrl_start
func (recv_ *EspLocalCtrlConfigT) EspLocalCtrlStart() EspErrT {
	return 0
}

/**
 * @brief   Stop local control service
 */
//go:linkname EspLocalCtrlStop C.esp_local_ctrl_stop
func EspLocalCtrlStop() EspErrT

/**
 * @brief   Add a new property
 *
 * This adds a new property and allocates internal resources for it.
 * The total number of properties that could be added is limited by
 * configuration option `max_properties`
 *
 * @param[in] prop    Property description structure
 *
 * @return
 *  - ESP_OK      : Success
 *  - ESP_FAIL    : Failure
 */
// llgo:link (*EspLocalCtrlPropT).EspLocalCtrlAddProperty C.esp_local_ctrl_add_property
func (recv_ *EspLocalCtrlPropT) EspLocalCtrlAddProperty() EspErrT {
	return 0
}

/**
 * @brief   Remove a property
 *
 * This finds a property by name, and releases the internal resources
 * which are associated with it.
 *
 * @param[in] name    Name of the property to remove
 *
 * @return
 *  - ESP_OK      : Success
 *  - ESP_ERR_NOT_FOUND : Failure
 */
//go:linkname EspLocalCtrlRemoveProperty C.esp_local_ctrl_remove_property
func EspLocalCtrlRemoveProperty(name *c.Char) EspErrT

/**
 * @brief   Get property description structure by name
 *
 * This API may be used to get a property's context structure
 * `esp_local_ctrl_prop_t` when its name is known
 *
 * @param[in] name    Name of the property to find
 *
 * @return
 *  - Pointer to property
 *  - NULL if not found
 */
//go:linkname EspLocalCtrlGetProperty C.esp_local_ctrl_get_property
func EspLocalCtrlGetProperty(name *c.Char) *EspLocalCtrlPropT

/**
 * @brief   Register protocomm handler for a custom endpoint
 *
 * This API can be called by the application to register a protocomm handler
 * for an endpoint after the local control service has started.
 *
 * @note In case of BLE transport the names and uuids of all custom
 * endpoints must be provided beforehand as a part of the `protocomm_ble_config_t`
 * structure set in `esp_local_ctrl_config_t`, and passed to `esp_local_ctrl_start()`.
 *
 * @param[in] ep_name   Name of the endpoint
 * @param[in] handler   Endpoint handler function
 * @param[in] user_ctx  User data
 *
 * @return
 *  - ESP_OK      : Success
 *  - ESP_FAIL    : Failure
 */
//go:linkname EspLocalCtrlSetHandler C.esp_local_ctrl_set_handler
func EspLocalCtrlSetHandler(ep_name *c.Char, handler ProtocommReqHandlerT, user_ctx c.Pointer) EspErrT
