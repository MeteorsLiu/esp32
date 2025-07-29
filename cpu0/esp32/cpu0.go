package esp32

import (
	_ "unsafe"

	rom "github.com/MeteorsLiu/esp32/esp32/esp_rom"
	system "github.com/MeteorsLiu/esp32/esp32/esp_system"
	freertos "github.com/MeteorsLiu/esp32/esp32/freertos"

	"github.com/MeteorsLiu/esp32/esp32/heap"
	"github.com/goplus/lib/c"
)

const (
	ESP_TASK_MAIN_STACK = system.CONFIG_ESP_MAIN_TASK_STACK_SIZE + 512
)

//go:linkname main main.main
func main()

//export esp_startup_start_app
func esp_startup_start_app() {
	system.EspIntWdtInit()
	system.EspIntWdtCpuInit()
	system.EspCrosscoreIntInit()

	rom.EspRomVprintf(c.Str("hello"), nil)
	freertos.XTaskCreatePinnedToCore(main_task, c.Str("main"), ESP_TASK_MAIN_STACK, nil, 1, nil, freertos.BaseTypeT(nil))

	freertos.XPortStartScheduler()
}

//export esp_startup_start_app_other_cores_default
func esp_startup_start_app_other_cores_default() {}

func main_task() {
	heap.HeapCapsEnableNonosStackHeaps()

	//CONFIG_ESP_TASK_WDT_TIMEOUT_S
	cfg := &system.EspTaskWdtConfigT{
		TimeoutMs: system.CONFIG_ESP_TASK_WDT_TIMEOUT_S * 1000,
	}

	cfg.IdleCoreMask |= (1 << 0)
	cfg.IdleCoreMask |= (1 << 1)

	cfg.EspTaskWdtInit()

	main()

	freertos.VTaskDelete(nil)
}
