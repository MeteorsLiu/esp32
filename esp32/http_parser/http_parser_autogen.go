package http_parser

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const X__NEWLIB_H__ = 1
const X_NEWLIB_VERSION_H__ = 1
const X_NEWLIB_VERSION = "4.3.0"
const X__NEWLIB__ = 4
const X__NEWLIB_MINOR__ = 3
const X__NEWLIB_PATCHLEVEL__ = 0
const X_ATEXIT_DYNAMIC_ALLOC = 1
const X_FSEEK_OPTIMIZATION = 1
const X_FVWRITE_IN_STREAMIO = 1
const X_HAVE_INITFINI_ARRAY = 1
const X_HAVE_LONG_DOUBLE = 1
const X_ICONV_ENABLED = 1
const X_LDBL_EQ_DBL = 1
const X_MB_LEN_MAX = 1
const X_NANO_MALLOC = 1
const X_REENT_CHECK_VERIFY = 1
const X_RETARGETABLE_LOCKING = 1
const X_UNBUF_STREAM_OPT = 1
const X_WANT_IO_C99_FORMATS = 1
const X_WANT_IO_LONG_LONG = 1
const X_WANT_IO_POS_ARGS = 1
const X_WANT_REENT_BACKWARD_BINARY_COMPAT = 1
const X_WANT_REENT_SMALL = 1
const X_WANT_USE_GDTOA = 1
const X__OBSOLETE_MATH_DEFAULT = 1
const X_ATFILE_SOURCE = 1
const X_DEFAULT_SOURCE = 1
const X_ISOC99_SOURCE = 1
const X_ISOC11_SOURCE = 1
const X_POSIX_SOURCE = 1
const X_XOPEN_SOURCE = 700
const X_XOPEN_SOURCE_EXTENDED = 1
const X__ATFILE_VISIBLE = 1
const X__BSD_VISIBLE = 1
const X__GNU_VISIBLE = 1
const X__ISO_C_VISIBLE = 2011
const X__LARGEFILE_VISIBLE = 1
const X__MISC_VISIBLE = 1
const X__POSIX_VISIBLE = 200809
const X__SVID_VISIBLE = 1
const X__XSI_VISIBLE = 700
const X__SSP_FORTIFY_LEVEL = 0
const X_POSIX_THREADS = 1
const X_POSIX_TIMEOUTS = 1
const X_POSIX_TIMERS = 1
const X_UNIX98_THREAD_MUTEX_ATTRIBUTES = 1
const X__BUFSIZ__ = 128
const X__RAND_MAX = 0x7fffffff
const X__have_longlong64 = 1
const X__have_long32 = 1
const X___int8_t_defined = 1
const X___int16_t_defined = 1
const X___int32_t_defined = 1
const X___int64_t_defined = 1
const X___int_least8_t_defined = 1
const X___int_least16_t_defined = 1
const X___int_least32_t_defined = 1
const X___int_least64_t_defined = 1
const X__GNUCLIKE_ASM = 3
const X__GNUCLIKE___TYPEOF = 1
const X__GNUCLIKE___SECTION = 1
const X__GNUCLIKE_CTOR_SECTION_HANDLING = 1
const X__GNUCLIKE_BUILTIN_CONSTANT_P = 1
const X__GNUCLIKE_BUILTIN_VARARGS = 1
const X__GNUCLIKE_BUILTIN_STDARG = 1
const X__GNUCLIKE_BUILTIN_VAALIST = 1
const X__GNUC_VA_LIST_COMPATIBILITY = 1
const X__GNUCLIKE_BUILTIN_NEXT_ARG = 1
const X__GNUCLIKE_BUILTIN_MEMCPY = 1
const X__CC_SUPPORTS_INLINE = 1
const X__CC_SUPPORTS___INLINE = 1
const X__CC_SUPPORTS___INLINE__ = 1
const X__CC_SUPPORTS___FUNC__ = 1
const X__CC_SUPPORTS_WARNING = 1
const X__CC_SUPPORTS_VARADIC_XXX = 1
const X__CC_SUPPORTS_DYNAMIC_ARRAY_INIT = 1
const X__BIT_TYPES_DEFINED__ = 1
const X__int8_t_defined = 1
const X__int16_t_defined = 1
const X__int32_t_defined = 1
const X__int64_t_defined = 1
const X_LITTLE_ENDIAN = 1234
const X_BIG_ENDIAN = 4321
const X_PDP_ENDIAN = 3412
const X_QUAD_HIGHWORD = 1
const X_QUAD_LOWWORD = 0
const FD_SETSIZE = 64
const SCHED_OTHER = 0
const SCHED_FIFO = 1
const SCHED_RR = 2
const PTHREAD_SCOPE_PROCESS = 0
const PTHREAD_SCOPE_SYSTEM = 1
const PTHREAD_INHERIT_SCHED = 1
const PTHREAD_EXPLICIT_SCHED = 2
const PTHREAD_CREATE_DETACHED = 0
const PTHREAD_CREATE_JOINABLE = 1
const PTHREAD_MUTEX_NORMAL = 0
const PTHREAD_MUTEX_RECURSIVE = 1
const PTHREAD_MUTEX_ERRORCHECK = 2
const PTHREAD_MUTEX_DEFAULT = 3
const X__INT8 = "hh"
const X__INT16 = "h"
const X__INT64 = "ll"
const X__FAST8 = "hh"
const X__FAST16 = "h"
const X__FAST64 = "ll"
const X__LEAST8 = "hh"
const X__LEAST16 = "h"
const X__LEAST64 = "ll"
const X__int_least8_t_defined = 1
const X__int_least16_t_defined = 1
const X__int_least32_t_defined = 1
const X__int_least64_t_defined = 1
const X__int_fast8_t_defined = 1
const X__int_fast16_t_defined = 1
const X__int_fast32_t_defined = 1
const X__int_fast64_t_defined = 1

type X__int8T c.Char
type X__uint8T c.Char
type X__int16T int16
type X__uint16T uint16
type X__int32T c.Int
type X__uint32T c.Uint
type X__int64T c.LongLong
type X__uint64T c.UlongLong
type X__intLeast8T c.Char
type X__uintLeast8T c.Char
type X__intLeast16T int16
type X__uintLeast16T uint16
type X__intLeast32T c.Int
type X__uintLeast32T c.Uint
type X__intLeast64T c.LongLong
type X__uintLeast64T c.UlongLong
type X__intmaxT c.LongLong
type X__uintmaxT c.UlongLong
type X__intptrT c.Int
type X__uintptrT c.Uint
type PtrdiffT c.Int
type WcharT c.Int

type MaxAlignT struct {
	X__clangMaxAlignNonce1 c.LongLong
	X__clangMaxAlignNonce2 c.Double
}
type UInt8T X__uint8T
type UInt16T X__uint16T
type UInt32T X__uint32T
type UInt64T X__uint64T
type RegisterT X__intptrT
type WintT c.Uint
type X__blkcntT c.Long
type X__blksizeT c.Long
type X__fsblkcntT X__uint64T
type X__fsfilcntT X__uint32T
type X_offT c.Long
type X__pidT c.Int
type X__devT int16
type X__uidT uint16
type X__gidT uint16
type X__idT X__uint32T
type X__inoT uint16
type X__modeT X__uint32T
type X_off64T c.LongLong
type X__offT X_offT
type X__loffT X_off64T
type X__keyT c.Long
type X_fposT c.Long
type X__sizeT c.Uint
type X_ssizeT c.Int
type X__ssizeT X_ssizeT

type X_mbstateT struct {
	X__count c.Int
	X__value struct {
		X__wch WintT
	}
}
type X_iconvT c.Pointer
type X__clockT c.Ulong
type X__timeT X__intLeast64T
type X__clockidT c.Ulong
type X__daddrT c.Long
type X__timerT c.Ulong
type X__saFamilyT X__uint8T
type X__socklenT X__uint32T
type X__nlItem c.Int
type X__nlinkT uint16
type X__susecondsT c.Long
type X__usecondsT c.Ulong
type X__vaList c.Pointer
type X__sigsetT c.Ulong
type SusecondsT X__susecondsT
type TimeT X__intLeast64T

type Timeval struct {
	TvSec  TimeT
	TvUsec SusecondsT
}

type Timespec struct {
	TvSec  TimeT
	TvNsec c.Long
}

type Itimerspec struct {
	ItInterval Timespec
	ItValue    Timespec
}
type SigsetT X__sigsetT
type X__fdMask c.Ulong
type FdMask X__fdMask

type FdSet struct {
	X__fdsBits [1]X__fdMask
}
type InAddrT X__uint32T
type InPortT X__uint16T
type URegisterT X__uintptrT
type UChar c.Char
type UShort uint16
type UInt c.Uint
type ULong c.Ulong
type Ushort uint16
type Uint c.Uint
type Ulong c.Ulong
type BlkcntT X__blkcntT
type BlksizeT X__blksizeT
type ClockT c.Ulong
type DaddrT X__daddrT
type CaddrT *c.Char
type FsblkcntT X__fsblkcntT
type FsfilcntT X__fsfilcntT
type IdT X__idT
type InoT X__inoT
type OffT X__offT
type DevT X__devT
type UidT X__uidT
type GidT X__gidT
type PidT X__pidT
type KeyT X__keyT
type ModeT X__modeT
type NlinkT X__nlinkT
type ClockidT X__clockidT
type TimerT X__timerT
type UsecondsT X__usecondsT
type SbintimeT X__int64T

type SchedParam struct {
	SchedPriority c.Int
}
type PthreadT X__uint32T

type PthreadAttrT struct {
	IsInitialized   c.Int
	Stackaddr       c.Pointer
	Stacksize       c.Int
	Contentionscope c.Int
	Inheritsched    c.Int
	Schedpolicy     c.Int
	Schedparam      SchedParam
	Detachstate     c.Int
}
type PthreadMutexT X__uint32T

type PthreadMutexattrT struct {
	IsInitialized c.Int
	Type          c.Int
	Recursive     c.Int
}
type PthreadCondT X__uint32T

type PthreadCondattrT struct {
	IsInitialized c.Int
	Clock         ClockT
}
type PthreadKeyT X__uint32T

type PthreadOnceT struct {
	IsInitialized c.Int
	InitExecuted  c.Int
}
type PthreadRwlockT X__uint32T

type PthreadRwlockattrT struct {
	IsInitialized c.Int
}
type IntLeast8T X__intLeast8T
type UintLeast8T X__uintLeast8T
type IntLeast16T X__intLeast16T
type UintLeast16T X__uintLeast16T
type IntLeast32T X__intLeast32T
type UintLeast32T X__uintLeast32T
type IntLeast64T X__intLeast64T
type UintLeast64T X__uintLeast64T
type IntFast8T c.Char
type UintFast8T c.Char
type IntFast16T int16
type UintFast16T uint16
type IntFast32T c.Int
type UintFast32T c.Uint
type IntFast64T c.LongLong
type UintFast64T c.UlongLong
