package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2sBitsPerSampleT c.Int

const (
	I2S_BITS_PER_SAMPLE_8BIT  I2sBitsPerSampleT = 8
	I2S_BITS_PER_SAMPLE_16BIT I2sBitsPerSampleT = 16
	I2S_BITS_PER_SAMPLE_24BIT I2sBitsPerSampleT = 24
	I2S_BITS_PER_SAMPLE_32BIT I2sBitsPerSampleT = 32
)

type I2sBitsPerChanT c.Int

const (
	I2S_BITS_PER_CHAN_DEFAULT I2sBitsPerChanT = 0
	I2S_BITS_PER_CHAN_8BIT    I2sBitsPerChanT = 8
	I2S_BITS_PER_CHAN_16BIT   I2sBitsPerChanT = 16
	I2S_BITS_PER_CHAN_24BIT   I2sBitsPerChanT = 24
	I2S_BITS_PER_CHAN_32BIT   I2sBitsPerChanT = 32
)

type I2sChannelT c.Int

const (
	I2S_CHANNEL_MONO   I2sChannelT = 1
	I2S_CHANNEL_STEREO I2sChannelT = 2
)

type I2sCommFormatT c.Int

const (
	I2S_COMM_FORMAT_STAND_I2S       I2sCommFormatT = 1
	I2S_COMM_FORMAT_STAND_MSB       I2sCommFormatT = 2
	I2S_COMM_FORMAT_STAND_PCM_SHORT I2sCommFormatT = 4
	I2S_COMM_FORMAT_STAND_PCM_LONG  I2sCommFormatT = 12
	I2S_COMM_FORMAT_STAND_MAX       I2sCommFormatT = 13
	I2S_COMM_FORMAT_I2S             I2sCommFormatT = 1
	I2S_COMM_FORMAT_I2S_MSB         I2sCommFormatT = 1
	I2S_COMM_FORMAT_I2S_LSB         I2sCommFormatT = 2
	I2S_COMM_FORMAT_PCM             I2sCommFormatT = 4
	I2S_COMM_FORMAT_PCM_SHORT       I2sCommFormatT = 4
	I2S_COMM_FORMAT_PCM_LONG        I2sCommFormatT = 8
)

type I2sChannelFmtT c.Int

const (
	I2S_CHANNEL_FMT_RIGHT_LEFT I2sChannelFmtT = 0
	I2S_CHANNEL_FMT_ALL_RIGHT  I2sChannelFmtT = 1
	I2S_CHANNEL_FMT_ALL_LEFT   I2sChannelFmtT = 2
	I2S_CHANNEL_FMT_ONLY_RIGHT I2sChannelFmtT = 3
	I2S_CHANNEL_FMT_ONLY_LEFT  I2sChannelFmtT = 4
)

type I2sModeT c.Int

const (
	I2S_MODE_MASTER I2sModeT = 1
	I2S_MODE_SLAVE  I2sModeT = 2
	I2S_MODE_TX     I2sModeT = 4
	I2S_MODE_RX     I2sModeT = 8
	I2S_MODE_PDM    I2sModeT = 64
)

type I2sEventTypeT c.Int

const (
	I2S_EVENT_DMA_ERROR I2sEventTypeT = 0
	I2S_EVENT_TX_DONE   I2sEventTypeT = 1
	I2S_EVENT_RX_DONE   I2sEventTypeT = 2
	I2S_EVENT_TX_Q_OVF  I2sEventTypeT = 3
	I2S_EVENT_RX_Q_OVF  I2sEventTypeT = 4
)

/**
 * @brief Event structure used in I2S event queue
 */

type I2sEventT struct {
	Type I2sEventTypeT
	Size c.SizeT
}

/**
 * @brief I2S GPIO pins configuration
 */

type I2sPinConfigT struct {
	MckIoNum   c.Int
	BckIoNum   c.Int
	WsIoNum    c.Int
	DataOutNum c.Int
	DataInNum  c.Int
}

/**
 * @brief I2S driver configuration parameters
 *
 */

type I2sDriverConfigT struct {
	Mode                I2sModeT
	SampleRate          c.Uint32T
	BitsPerSample       I2sBitsPerSampleT
	ChannelFormat       I2sChannelFmtT
	CommunicationFormat I2sCommFormatT
	IntrAllocFlags      c.Int
	UseApll             bool
	TxDescAutoClear     bool
	FixedMclk           c.Int
	MclkMultiple        I2sMclkMultipleT
	BitsPerChan         I2sBitsPerChanT
}
type I2sConfigT I2sDriverConfigT
