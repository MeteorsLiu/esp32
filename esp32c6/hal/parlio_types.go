package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ParlioSampleEdgeT c.Int

const (
	PARLIO_SAMPLE_EDGE_NEG ParlioSampleEdgeT = 0
	PARLIO_SAMPLE_EDGE_POS ParlioSampleEdgeT = 1
)

type ParlioBitPackOrderT c.Int

const (
	PARLIO_BIT_PACK_ORDER_LSB ParlioBitPackOrderT = 0
	PARLIO_BIT_PACK_ORDER_MSB ParlioBitPackOrderT = 1
)

type ParlioClockSourceT SocPeriphParlioClkSrcT
