package esp_lcd

import _ "unsafe"

const ESP_LCD_I80_BUS_WIDTH_MAX = 16

type EspLcdI80BusT struct {
	Unused [8]uint8
}
type EspLcdI80BusHandleT *EspLcdI80BusT
