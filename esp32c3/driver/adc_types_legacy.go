package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcBitsWidthT c.Int

const (
	ADC_WIDTH_BIT_12 AdcBitsWidthT = 12
	ADC_WIDTH_MAX    AdcBitsWidthT = 13
)

type Adc1ChannelT c.Int

const (
	ADC1_CHANNEL_0   Adc1ChannelT = 0
	ADC1_CHANNEL_1   Adc1ChannelT = 1
	ADC1_CHANNEL_2   Adc1ChannelT = 2
	ADC1_CHANNEL_3   Adc1ChannelT = 3
	ADC1_CHANNEL_4   Adc1ChannelT = 4
	ADC1_CHANNEL_MAX Adc1ChannelT = 5
)

type Adc2ChannelT c.Int

const (
	ADC2_CHANNEL_0   Adc2ChannelT = 0
	ADC2_CHANNEL_MAX Adc2ChannelT = 1
)

/**
 * @brief ADC DMA driver configuration
 */

type AdcDigiInitConfigS struct {
	MaxStoreBufSize c.Uint32T
	ConvNumEachIntr c.Uint32T
	Adc1ChanMask    c.Uint32T
	Adc2ChanMask    c.Uint32T
}
type AdcDigiInitConfigT AdcDigiInitConfigS

/**
 * @brief ADC digital controller settings
 */

type AdcDigiConfigurationT struct {
	ConvLimitEn  bool
	ConvLimitNum c.Uint32T
	PatternNum   c.Uint32T
	AdcPattern   *AdcDigiPatternConfigT
	SampleFreqHz c.Uint32T
	ConvMode     AdcDigiConvertModeT
	Format       AdcDigiOutputFormatT
}
