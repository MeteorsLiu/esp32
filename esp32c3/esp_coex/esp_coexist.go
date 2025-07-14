package esp_coex

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const EXTERNAL_COEXIST_WIRE_1 = 0
const EXTERNAL_COEXIST_WIRE_2 = 1
const EXTERNAL_COEXIST_WIRE_3 = 2
const EXTERNAL_COEXIST_WIRE_4 = 3
const ESP_COEX_BLE_ST_MESH_CONFIG = 0x08
const ESP_COEX_BLE_ST_MESH_TRAFFIC = 0x10
const ESP_COEX_BLE_ST_MESH_STANDBY = 0x20
const ESP_COEX_BT_ST_A2DP_STREAMING = 0x10
const ESP_COEX_BT_ST_A2DP_PAUSED = 0x20

type EspCoexPreferT c.Int

const (
	ESP_COEX_PREFER_WIFI    EspCoexPreferT = 0
	ESP_COEX_PREFER_BT      EspCoexPreferT = 1
	ESP_COEX_PREFER_BALANCE EspCoexPreferT = 2
	ESP_COEX_PREFER_NUM     EspCoexPreferT = 3
)

type ExternalCoexWireT c.Int

const (
	EXTERN_COEX_WIRE_1   ExternalCoexWireT = 0
	EXTERN_COEX_WIRE_2   ExternalCoexWireT = 1
	EXTERN_COEX_WIRE_3   ExternalCoexWireT = 2
	EXTERN_COEX_WIRE_4   ExternalCoexWireT = 3
	EXTERN_COEX_WIRE_NUM ExternalCoexWireT = 4
)

type EspCoexStatusTypeT c.Int

const (
	ESP_COEX_ST_TYPE_WIFI EspCoexStatusTypeT = 0
	ESP_COEX_ST_TYPE_BLE  EspCoexStatusTypeT = 1
	ESP_COEX_ST_TYPE_BT   EspCoexStatusTypeT = 2
)
