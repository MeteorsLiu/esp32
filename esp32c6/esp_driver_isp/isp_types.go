package esp_driver_isp

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief ISP unsigned integer range type
 * @note  Whether the edge value are included depends on the variable itself
 */

type IspU32RangeT struct {
	Min c.Uint32T
	Max c.Uint32T
}

/**
 * @brief ISP float range type
 * @note  Whether the edge value are included depends on the variable itself
 */

type IspFloatRangeT struct {
	Min c.Float
	Max c.Float
}

/**
 * @brief ISP AF result
 */

type IspAfResultT struct {
	Definition [0]c.Int
	Luminance  [0]c.Int
}

/**
 * @brief ISP AWB result
 */

type IspAwbStatResultT struct {
	WhitePatchNum c.Uint32T
	SumR          c.Uint32T
	SumG          c.Uint32T
	SumB          c.Uint32T
}

/**
 * @brief ISP AE result
 */

type IspAeResultT struct {
	Luminance [0][0]c.Int
}

type IspProcessorT struct {
	Unused [8]uint8
}
type IspProcHandleT *IspProcessorT

type IspAfControllerT struct {
	Unused [8]uint8
}
type IspAfCtlrT *IspAfControllerT

type IspAwbControllerT struct {
	Unused [8]uint8
}
type IspAwbCtlrT *IspAwbControllerT

type IspAeControllerT struct {
	Unused [8]uint8
}
type IspAeCtlrT *IspAeControllerT

type IspHistControllerT struct {
	Unused [8]uint8
}
type IspHistCtlrT *IspHistControllerT

/*---------------------------------------------
                Event Callbacks
----------------------------------------------*/
/**
 * @brief Event data structure
 */

type EspIspSharpenEvtDataT struct {
	HighFreqPixelMax c.Uint8T
}

// llgo:type C
type EspIspSharpenCallbackT func(IspProcHandleT, *EspIspSharpenEvtDataT, c.Pointer) bool

/**
 * @brief Group of ISP event callbacks
 *
 * @note These callbacks are all running in an ISR environment.
 * @note When CONFIG_ISP_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       Involved variables should be in internal RAM as well.
 */

type EspIspEvtCbsT struct {
	OnSharpenFrameDone EspIspSharpenCallbackT
}
