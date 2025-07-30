package main

import (
	_ "unsafe"

	gpio "github.com/MeteorsLiu/esp32/esp32/esp_driver_gpio"
	rom "github.com/MeteorsLiu/esp32/esp32/esp_rom"
	system "github.com/MeteorsLiu/esp32/esp32/esp_system"
	freertos "github.com/MeteorsLiu/esp32/esp32/freertos"

	"github.com/MeteorsLiu/esp32/esp32/heap"
	"github.com/MeteorsLiu/esp32/esp32/log"
	"github.com/goplus/lib/c"
)

const (
	ESP_TASK_MAIN_STACK = system.CONFIG_ESP_MAIN_TASK_STACK_SIZE + 512
)

//export esp_startup_start_app
func esp_startup_start_app() {
	system.EspIntWdtInit()
	system.EspIntWdtCpuInit()
	system.EspCrosscoreIntInit()

	rom.EspRomVprintf(c.Str("hello\n"), nil)

	freertos.XTaskCreatePinnedToCore(main_task, c.Str("main"), ESP_TASK_MAIN_STACK, nil, 1, nil, 0)

	rom.EspRomVprintf(c.Str("start scheduler\n"), nil)

	freertos.VTaskStartScheduler()

}

//export esp_startup_start_app_other_cores_default
func esp_startup_start_app_other_cores_default() {
	rom.EspRomVprintf(c.Str("hello core1"), nil)

	system.EspIntWdtCpuInit()
	system.EspCrosscoreIntInit()
	freertos.XPortStartScheduler()
}

func main_task(args *c.Void) {
	rom.EspRomVprintf(c.Str("enter main task\n"), nil)

	callMain()

}

func callMain() {

	rom.EspRomVprintf(c.Str("start to init heap\n"), nil)

	heap.HeapCapsEnableNonosStackHeaps()

	//CONFIG_ESP_TASK_WDT_TIMEOUT_S
	cfg := &system.EspTaskWdtConfigT{
		TimeoutMs:    system.CONFIG_ESP_TASK_WDT_TIMEOUT_S * 1000,
		TriggerPanic: true,
		IdleCoreMask: 0,
	}
	cfg.IdleCoreMask |= (1 << 0)

	cfg.EspTaskWdtInit()

	rom.EspRomVprintf(c.Str("call main"), nil)

	main()

	freertos.VTaskDelete(nil)
}

func main() {
	tag := c.Str("llgo-esp32")
	btnIO := gpio.GpioNumT(gpio.GPIO_NUM_34)
	gpio.GpioSetDirection(btnIO, gpio.GPIO_MODE_INPUT)
	log.EspLogWrite(log.ESP_LOG_INFO, tag, c.Str("hello llgo-esp32!!!!!\n"))
	for {
		level := gpio.GpioGetLevel(btnIO)
		if level == 1 {
			log.EspLogWrite(log.ESP_LOG_INFO, tag, c.Str("No Press!\n"))
		} else {
			log.EspLogWrite(log.ESP_LOG_INFO, tag, c.Str("Pressed!\n"))
		}
	}
}
