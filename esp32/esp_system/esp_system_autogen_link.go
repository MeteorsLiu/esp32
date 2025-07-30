package esp_system

// spi_flash
// bootloader_support
// mbedtls
// soc
// esp_hw_support
// efuse
// esp_partition
// partition_table
// esptool_py
// app_update
// esp_app_format
// esp_security
// esp_timer
// esp_driver_gpio
// esp_pm
// esp_mm
// heap
// esp_bootloader_format
import (
	_ "github.com/MeteorsLiu/esp32/esp32/esp_gdbstub"
	_ "github.com/MeteorsLiu/esp32/esp32/hal"
	_ "github.com/MeteorsLiu/esp32/esp32/log"
	_ "github.com/MeteorsLiu/esp32/esp32/pthread"
	_ "github.com/MeteorsLiu/esp32/esp32/soc"
	_ "github.com/MeteorsLiu/esp32/esp32/spi_flash"
	_ "github.com/MeteorsLiu/esp32/esp32/xtensa"

	_ "github.com/MeteorsLiu/esp32/esp32/app_update"
	_ "github.com/MeteorsLiu/esp32/esp32/bootloader_support"
	_ "github.com/MeteorsLiu/esp32/esp32/efuse"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_app_format"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_bootloader_format"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_common"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_driver_gpio"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_hw_support"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_mm"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_partition"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_pm"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_rom"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_security"
	_ "github.com/MeteorsLiu/esp32/esp32/esp_timer"
	_ "github.com/MeteorsLiu/esp32/esp32/heap"
	_ "github.com/MeteorsLiu/esp32/esp32/newlib"
	_ "github.com/MeteorsLiu/esp32/esp32/vfs"

	_ "github.com/goplus/lib/c"
)

const LLGoPackage string = "link: -L/Users/haolan/esp/esp-idf/examples/get-started/hello_world/build/esp-idf/esp_system -lesp_system;"
