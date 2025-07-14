package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type IspFsmT c.Int

const (
	ISP_FSM_INIT       IspFsmT = 0
	ISP_FSM_ENABLE     IspFsmT = 1
	ISP_FSM_START      IspFsmT = 2
	ISP_FSM_ONESHOT    IspFsmT = 3
	ISP_FSM_CONTINUOUS IspFsmT = 4
)

type IspSubmoduleT c.Int

const (
	ISP_SUBMODULE_AF      IspSubmoduleT = 0
	ISP_SUBMODULE_AE      IspSubmoduleT = 1
	ISP_SUBMODULE_AWB     IspSubmoduleT = 2
	ISP_SUBMODULE_SHARPEN IspSubmoduleT = 3
	ISP_SUBMODULE_HIST    IspSubmoduleT = 4
)

/*---------------------------------------------------------------
                      INTR
---------------------------------------------------------------*/
//go:linkname EspIspRegisterIsr C.esp_isp_register_isr
func EspIspRegisterIsr(proc IspProcHandleT, submodule IspSubmoduleT) EspErrT

//go:linkname EspIspDeregisterIsr C.esp_isp_deregister_isr
func EspIspDeregisterIsr(proc IspProcHandleT, submodule IspSubmoduleT) EspErrT

//go:linkname EspIspAfIsr C.esp_isp_af_isr
func EspIspAfIsr(proc IspProcHandleT, af_events c.Uint32T) bool

//go:linkname EspIspAeIsr C.esp_isp_ae_isr
func EspIspAeIsr(proc IspProcHandleT, ae_events c.Uint32T) bool

//go:linkname EspIspAwbIsr C.esp_isp_awb_isr
func EspIspAwbIsr(proc IspProcHandleT, awb_events c.Uint32T) bool

//go:linkname EspIspSharpenIsr C.esp_isp_sharpen_isr
func EspIspSharpenIsr(proc IspProcHandleT, sharp_events c.Uint32T) bool

//go:linkname EspIspHistIsr C.esp_isp_hist_isr
func EspIspHistIsr(proc IspProcHandleT, hist_events c.Uint32T) bool
