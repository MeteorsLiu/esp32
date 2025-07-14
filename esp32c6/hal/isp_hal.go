package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief BF configurations
 */

type IspHalBfCfgT struct {
	PaddingMode                    IspBfEdgePaddingModeT
	PaddingData                    c.Uint8T
	BfTemplate                     [0][0]c.Uint8T
	DenoisingLevel                 c.Uint8T
	PaddingLineTailValidStartPixel c.Uint8T
	PaddingLineTailValidEndPixel   c.Uint8T
}

/**
 * @brief Demosaic configurations
 */

type IspHalDemosaicCfgT struct {
	GradRatio                      IspDemosaicGradRatioT
	PaddingMode                    IspDemosaicEdgePaddingModeT
	PaddingData                    c.Uint8T
	PaddingLineTailValidStartPixel c.Uint8T
	PaddingLineTailValidEndPixel   c.Uint8T
}

/**
 * @brief Sharpen configurations
 */

type IspHalSharpenCfgT struct {
	HFreqCoeff                     IspSharpenHFreqCoeffT
	MFreqCoeff                     IspSharpenMFreqCoeff
	HThresh                        c.Uint8T
	LThresh                        c.Uint8T
	PaddingMode                    IspSharpenEdgePaddingModeT
	PaddingData                    c.Uint8T
	SharpenTemplate                [0][0]c.Uint8T
	PaddingLineTailValidStartPixel c.Uint8T
	PaddingLineTailValidEndPixel   c.Uint8T
}

/**
 * @brief Context that should be maintained by both the driver and the HAL
 */

type IspHalContextT struct {
	Hw    c.Pointer
	BfCfg IspHalBfCfgT
}

/**
 * @brief Color configurations
 */

type IspHalColorCfgT struct {
	ColorContrast   IspColorContrastT
	ColorSaturation IspColorSaturationT
	ColorHue        c.Uint32T
	ColorBrightness c.Int
}
