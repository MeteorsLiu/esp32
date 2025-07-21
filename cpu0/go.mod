module github.com/MeteorsLiu/esp32/cpu0

go 1.23.4

replace github.com/MeteorsLiu/esp32/esp32/esp_system => ../esp32/esp_system

replace github.com/MeteorsLiu/esp32/esp32/freertos => ../esp32/freertos

replace github.com/MeteorsLiu/esp32/esp32/heap => ../esp32/heap

replace github.com/MeteorsLiu/esp32/esp32/newlib => ../esp32/newlib

require (
	github.com/MeteorsLiu/esp32/esp32/esp_system v0.0.0-00010101000000-000000000000
	github.com/MeteorsLiu/esp32/esp32/freertos v0.0.0-00010101000000-000000000000
	github.com/MeteorsLiu/esp32/esp32/heap v0.0.0-00010101000000-000000000000
)

require (
	github.com/MeteorsLiu/esp32/esp32/newlib v0.0.0-00010101000000-000000000000
	github.com/goplus/lib v0.2.0
)
