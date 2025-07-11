package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MINIZ_X86_OR_X64_CPU = 0
const MINIZ_LITTLE_ENDIAN = 1
const MINIZ_USE_UNALIGNED_LOADS_AND_STORES = 0
const MINIZ_HAS_64BIT_REGISTERS = 0
const TINFL_USE_64BIT_BITBUF = 0
const MZ_DEFLATED = 8
const TINFL_LZ_DICT_SIZE = 32768
const TDEFL_LESS_MEMORY = 1

type MzUlong c.Ulong

const (
	MZ_DEFAULT_STRATEGY c.Int = 0
	MZ_FILTERED         c.Int = 1
	MZ_HUFFMAN_ONLY     c.Int = 2
	MZ_RLE              c.Int = 3
	MZ_FIXED            c.Int = 4
)

type MzUint8 c.Char
type MzInt16 int16
type MzUint16 uint16
type MzUint32 c.Uint
type MzUint c.Uint
type MzInt64 c.LongLong
type MzUint64 c.UlongLong
type MzBool c.Int

const (
	TINFL_FLAG_PARSE_ZLIB_HEADER             c.Int = 1
	TINFL_FLAG_HAS_MORE_INPUT                c.Int = 2
	TINFL_FLAG_USING_NON_WRAPPING_OUTPUT_BUF c.Int = 4
	TINFL_FLAG_COMPUTE_ADLER32               c.Int = 8
)

// llgo:type C
type TinflPutBufFuncPtr func(c.Pointer, c.Int, c.Pointer) c.Int

type TinflDecompressorTag struct {
	Unused [8]uint8
}
type TinflDecompressor TinflDecompressorTag
type TinflStatus c.Int

const (
	TINFL_STATUS_BAD_PARAM        TinflStatus = -3
	TINFL_STATUS_ADLER32_MISMATCH TinflStatus = -2
	TINFL_STATUS_FAILED           TinflStatus = -1
	TINFL_STATUS_DONE             TinflStatus = 0
	TINFL_STATUS_NEEDS_MORE_INPUT TinflStatus = 1
	TINFL_STATUS_HAS_MORE_OUTPUT  TinflStatus = 2
)
const (
	TINFL_MAX_HUFF_TABLES    c.Int = 3
	TINFL_MAX_HUFF_SYMBOLS_0 c.Int = 288
	TINFL_MAX_HUFF_SYMBOLS_1 c.Int = 32
	TINFL_MAX_HUFF_SYMBOLS_2 c.Int = 19
	TINFL_FAST_LOOKUP_BITS   c.Int = 10
	TINFL_FAST_LOOKUP_SIZE   c.Int = 1024
)

type TinflHuffTable struct {
	MCodeSize [288]MzUint8
	MLookUp   [1024]MzInt16
	MTree     [576]MzInt16
}
type TinflBitBufT MzUint32

const (
	TDEFL_HUFFMAN_ONLY       c.Int = 0
	TDEFL_DEFAULT_MAX_PROBES c.Int = 128
	TDEFL_MAX_PROBES_MASK    c.Int = 4095
)
const (
	TDEFL_WRITE_ZLIB_HEADER             c.Int = 4096
	TDEFL_COMPUTE_ADLER32               c.Int = 8192
	TDEFL_GREEDY_PARSING_FLAG           c.Int = 16384
	TDEFL_NONDETERMINISTIC_PARSING_FLAG c.Int = 32768
	TDEFL_RLE_MATCHES                   c.Int = 65536
	TDEFL_FILTER_MATCHES                c.Int = 131072
	TDEFL_FORCE_ALL_STATIC_BLOCKS       c.Int = 262144
	TDEFL_FORCE_ALL_RAW_BLOCKS          c.Int = 524288
)

// llgo:type C
type TdeflPutBufFuncPtr func(c.Pointer, c.Int, c.Pointer) MzBool

const (
	TDEFL_MAX_HUFF_TABLES    c.Int = 3
	TDEFL_MAX_HUFF_SYMBOLS_0 c.Int = 288
	TDEFL_MAX_HUFF_SYMBOLS_1 c.Int = 32
	TDEFL_MAX_HUFF_SYMBOLS_2 c.Int = 19
	TDEFL_LZ_DICT_SIZE       c.Int = 32768
	TDEFL_LZ_DICT_SIZE_MASK  c.Int = 32767
	TDEFL_MIN_MATCH_LEN      c.Int = 3
	TDEFL_MAX_MATCH_LEN      c.Int = 258
)
const (
	TDEFL_LZ_CODE_BUF_SIZE      c.Int = 24576
	TDEFL_OUT_BUF_SIZE          c.Int = 31948
	TDEFL_MAX_HUFF_SYMBOLS      c.Int = 288
	TDEFL_LZ_HASH_BITS          c.Int = 12
	TDEFL_LEVEL1_HASH_SIZE_MASK c.Int = 4095
	TDEFL_LZ_HASH_SHIFT         c.Int = 4
	TDEFL_LZ_HASH_SIZE          c.Int = 4096
)

type TdeflStatus c.Int

const (
	TDEFL_STATUS_BAD_PARAM      TdeflStatus = -2
	TDEFL_STATUS_PUT_BUF_FAILED TdeflStatus = -1
	TDEFL_STATUS_OKAY           TdeflStatus = 0
	TDEFL_STATUS_DONE           TdeflStatus = 1
)

type TdeflFlush c.Int

const (
	TDEFL_NO_FLUSH   TdeflFlush = 0
	TDEFL_SYNC_FLUSH TdeflFlush = 2
	TDEFL_FULL_FLUSH TdeflFlush = 3
	TDEFL_FINISH     TdeflFlush = 4
)

// tdefl's compression state structure.
type TdeflCompressor struct {
	MPPutBufFunc          TdeflPutBufFuncPtr
	MPPutBufUser          c.Pointer
	MFlags                MzUint
	MMaxProbes            [2]MzUint
	MGreedyParsing        c.Int
	MAdler32              MzUint
	MLookaheadPos         MzUint
	MLookaheadSize        MzUint
	MDictSize             MzUint
	MPLZCodeBuf           *MzUint8
	MPLZFlags             *MzUint8
	MPOutputBuf           *MzUint8
	MPOutputBufEnd        *MzUint8
	MNumFlagsLeft         MzUint
	MTotalLzBytes         MzUint
	MLzCodeBufDictPos     MzUint
	MBitsIn               MzUint
	MBitBuffer            MzUint
	MSavedMatchDist       MzUint
	MSavedMatchLen        MzUint
	MSavedLit             MzUint
	MOutputFlushOfs       MzUint
	MOutputFlushRemaining MzUint
	MFinished             MzUint
	MBlockIndex           MzUint
	MWantsToFinish        MzUint
	MPrevReturnStatus     TdeflStatus
	MPInBuf               c.Pointer
	MPOutBuf              c.Pointer
	MPInBufSize           *c.SizeT
	MPOutBufSize          *c.SizeT
	MFlush                TdeflFlush
	MPSrc                 *MzUint8
	MSrcBufLeft           c.SizeT
	MOutBufOfs            c.SizeT
	MDict                 [33025]MzUint8
	MHuffCount            [3][288]MzUint16
	MHuffCodes            [3][288]MzUint16
	MHuffCodeSizes        [3][288]MzUint8
	MLzCodeBuf            [24576]MzUint8
	MNext                 [32768]MzUint16
	MHash                 [4096]MzUint16
	MOutputBuf            [31948]MzUint8
}
