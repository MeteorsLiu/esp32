package esp32

import (
	"unsafe"

	_ "github.com/MeteorsLiu/esp32/esp32/newlib"

	system "github.com/MeteorsLiu/esp32/esp32/esp_system"
	freertos "github.com/MeteorsLiu/esp32/esp32/freertos"

	"github.com/MeteorsLiu/esp32/esp32/heap"
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

	fn := main_task
	freertos.XTaskCreatePinnedToCore(c.Int(uintptr(unsafe.Pointer(&fn))), c.Str("main"), ESP_TASK_MAIN_STACK, nil, 1, nil, freertos.BaseTypeT(nil))

	freertos.XPortStartScheduler()
}

func main_task() {
	heap.HeapCapsEnableNonosStackHeaps()

	//CONFIG_ESP_TASK_WDT_TIMEOUT_S
	cfg := &system.EspTaskWdtConfigT{
		TimeoutMs: system.CONFIG_ESP_TASK_WDT_TIMEOUT_S * 1000,
	}

	cfg.IdleCoreMask |= (1 << 0)
	cfg.IdleCoreMask |= (1 << 1)

	cfg.EspTaskWdtInit()

	c.Printf(c.Str("start main task"))

	main()

	freertos.VTaskDelete(nil)
}

func main() {
	c.Printf(c.Str("start main"))
}
