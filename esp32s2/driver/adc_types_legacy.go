package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AdcBitsWidthT c.Int

const (
	ADC_WIDTH_BIT_13 AdcBitsWidthT = 13
	ADC_WIDTH_MAX    AdcBitsWidthT = 14
)

type Adc1ChannelT c.Int

const (
	ADC1_CHANNEL_0   Adc1ChannelT = 0
	ADC1_CHANNEL_1   Adc1ChannelT = 1
	ADC1_CHANNEL_2   Adc1ChannelT = 2
	ADC1_CHANNEL_3   Adc1ChannelT = 3
	ADC1_CHANNEL_4   Adc1ChannelT = 4
	ADC1_CHANNEL_5   Adc1ChannelT = 5
	ADC1_CHANNEL_6   Adc1ChannelT = 6
	ADC1_CHANNEL_7   Adc1ChannelT = 7
	ADC1_CHANNEL_8   Adc1ChannelT = 8
	ADC1_CHANNEL_9   Adc1ChannelT = 9
	ADC1_CHANNEL_MAX Adc1ChannelT = 10
)

type Adc2ChannelT c.Int

const (
	ADC2_CHANNEL_0   Adc2ChannelT = 0
	ADC2_CHANNEL_1   Adc2ChannelT = 1
	ADC2_CHANNEL_2   Adc2ChannelT = 2
	ADC2_CHANNEL_3   Adc2ChannelT = 3
	ADC2_CHANNEL_4   Adc2ChannelT = 4
	ADC2_CHANNEL_5   Adc2ChannelT = 5
	ADC2_CHANNEL_6   Adc2ChannelT = 6
	ADC2_CHANNEL_7   Adc2ChannelT = 7
	ADC2_CHANNEL_8   Adc2ChannelT = 8
	ADC2_CHANNEL_9   Adc2ChannelT = 9
	ADC2_CHANNEL_MAX Adc2ChannelT = 10
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
