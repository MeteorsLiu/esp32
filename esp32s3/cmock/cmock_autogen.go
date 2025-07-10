package cmock

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const UNITY_VERSION_MAJOR = 2
const UNITY_VERSION_MINOR = 6
const UNITY_VERSION_BUILD = 0
const X__has_safe_buffers = 1
const X__DARWIN_ONLY_64_BIT_INO_T = 1
const X__DARWIN_ONLY_UNIX_CONFORMANCE = 1
const X__DARWIN_ONLY_VERS_1050 = 1
const X__DARWIN_UNIX03 = 1
const X__DARWIN_64_BIT_INO_T = 1
const X__DARWIN_VERS_1050 = 1
const X__DARWIN_NON_CANCELABLE = 0
const X__DARWIN_SUF_EXTSN = "$DARWIN_EXTSN"
const X__STDC_WANT_LIB_EXT1__ = 1
const X__DARWIN_NO_LONG_LONG = 0
const X_DARWIN_FEATURE_64_BIT_INODE = 1
const X_DARWIN_FEATURE_ONLY_64_BIT_INODE = 1
const X_DARWIN_FEATURE_ONLY_VERS_1050 = 1
const X_DARWIN_FEATURE_ONLY_UNIX_CONFORMANCE = 1
const X_DARWIN_FEATURE_UNIX_CONFORMANCE = 3
const X__has_ptrcheck = 0
const X__API_TO_BE_DEPRECATED = 100000
const X__API_TO_BE_DEPRECATED_MACOS = 100000
const X__API_TO_BE_DEPRECATED_IOS = 100000
const X__API_TO_BE_DEPRECATED_MACCATALYST = 100000
const X__API_TO_BE_DEPRECATED_WATCHOS = 100000
const X__API_TO_BE_DEPRECATED_TVOS = 100000
const X__API_TO_BE_DEPRECATED_DRIVERKIT = 100000
const X__API_TO_BE_DEPRECATED_VISIONOS = 100000
const X__MAC_10_0 = 1000
const X__MAC_10_1 = 1010
const X__MAC_10_2 = 1020
const X__MAC_10_3 = 1030
const X__MAC_10_4 = 1040
const X__MAC_10_5 = 1050
const X__MAC_10_6 = 1060
const X__MAC_10_7 = 1070
const X__MAC_10_8 = 1080
const X__MAC_10_9 = 1090
const X__MAC_10_10 = 101000
const X__MAC_10_10_2 = 101002
const X__MAC_10_10_3 = 101003
const X__MAC_10_11 = 101100
const X__MAC_10_11_2 = 101102
const X__MAC_10_11_3 = 101103
const X__MAC_10_11_4 = 101104
const X__MAC_10_12 = 101200
const X__MAC_10_12_1 = 101201
const X__MAC_10_12_2 = 101202
const X__MAC_10_12_4 = 101204
const X__MAC_10_13 = 101300
const X__MAC_10_13_1 = 101301
const X__MAC_10_13_2 = 101302
const X__MAC_10_13_4 = 101304
const X__MAC_10_14 = 101400
const X__MAC_10_14_1 = 101401
const X__MAC_10_14_4 = 101404
const X__MAC_10_14_5 = 101405
const X__MAC_10_14_6 = 101406
const X__MAC_10_15 = 101500
const X__MAC_10_15_1 = 101501
const X__MAC_10_15_4 = 101504
const X__MAC_10_16 = 101600
const X__MAC_11_0 = 110000
const X__MAC_11_1 = 110100
const X__MAC_11_3 = 110300
const X__MAC_11_4 = 110400
const X__MAC_11_5 = 110500
const X__MAC_11_6 = 110600
const X__MAC_12_0 = 120000
const X__MAC_12_1 = 120100
const X__MAC_12_2 = 120200
const X__MAC_12_3 = 120300
const X__MAC_12_4 = 120400
const X__MAC_12_5 = 120500
const X__MAC_12_6 = 120600
const X__MAC_12_7 = 120700
const X__MAC_13_0 = 130000
const X__MAC_13_1 = 130100
const X__MAC_13_2 = 130200
const X__MAC_13_3 = 130300
const X__MAC_13_4 = 130400
const X__MAC_13_5 = 130500
const X__MAC_13_6 = 130600
const X__MAC_14_0 = 140000
const X__MAC_14_1 = 140100
const X__MAC_14_2 = 140200
const X__MAC_14_3 = 140300
const X__MAC_14_4 = 140400
const X__IPHONE_2_0 = 20000
const X__IPHONE_2_1 = 20100
const X__IPHONE_2_2 = 20200
const X__IPHONE_3_0 = 30000
const X__IPHONE_3_1 = 30100
const X__IPHONE_3_2 = 30200
const X__IPHONE_4_0 = 40000
const X__IPHONE_4_1 = 40100
const X__IPHONE_4_2 = 40200
const X__IPHONE_4_3 = 40300
const X__IPHONE_5_0 = 50000
const X__IPHONE_5_1 = 50100
const X__IPHONE_6_0 = 60000
const X__IPHONE_6_1 = 60100
const X__IPHONE_7_0 = 70000
const X__IPHONE_7_1 = 70100
const X__IPHONE_8_0 = 80000
const X__IPHONE_8_1 = 80100
const X__IPHONE_8_2 = 80200
const X__IPHONE_8_3 = 80300
const X__IPHONE_8_4 = 80400
const X__IPHONE_9_0 = 90000
const X__IPHONE_9_1 = 90100
const X__IPHONE_9_2 = 90200
const X__IPHONE_9_3 = 90300
const X__IPHONE_10_0 = 100000
const X__IPHONE_10_1 = 100100
const X__IPHONE_10_2 = 100200
const X__IPHONE_10_3 = 100300
const X__IPHONE_11_0 = 110000
const X__IPHONE_11_1 = 110100
const X__IPHONE_11_2 = 110200
const X__IPHONE_11_3 = 110300
const X__IPHONE_11_4 = 110400
const X__IPHONE_12_0 = 120000
const X__IPHONE_12_1 = 120100
const X__IPHONE_12_2 = 120200
const X__IPHONE_12_3 = 120300
const X__IPHONE_12_4 = 120400
const X__IPHONE_13_0 = 130000
const X__IPHONE_13_1 = 130100
const X__IPHONE_13_2 = 130200
const X__IPHONE_13_3 = 130300
const X__IPHONE_13_4 = 130400
const X__IPHONE_13_5 = 130500
const X__IPHONE_13_6 = 130600
const X__IPHONE_13_7 = 130700
const X__IPHONE_14_0 = 140000
const X__IPHONE_14_1 = 140100
const X__IPHONE_14_2 = 140200
const X__IPHONE_14_3 = 140300
const X__IPHONE_14_5 = 140500
const X__IPHONE_14_4 = 140400
const X__IPHONE_14_6 = 140600
const X__IPHONE_14_7 = 140700
const X__IPHONE_14_8 = 140800
const X__IPHONE_15_0 = 150000
const X__IPHONE_15_1 = 150100
const X__IPHONE_15_2 = 150200
const X__IPHONE_15_3 = 150300
const X__IPHONE_15_4 = 150400
const X__IPHONE_15_5 = 150500
const X__IPHONE_15_6 = 150600
const X__IPHONE_15_7 = 150700
const X__IPHONE_15_8 = 150800
const X__IPHONE_16_0 = 160000
const X__IPHONE_16_1 = 160100
const X__IPHONE_16_2 = 160200
const X__IPHONE_16_3 = 160300
const X__IPHONE_16_4 = 160400
const X__IPHONE_16_5 = 160500
const X__IPHONE_16_6 = 160600
const X__IPHONE_16_7 = 160700
const X__IPHONE_17_0 = 170000
const X__IPHONE_17_1 = 170100
const X__IPHONE_17_2 = 170200
const X__IPHONE_17_3 = 170300
const X__IPHONE_17_4 = 170400
const X__WATCHOS_1_0 = 10000
const X__WATCHOS_2_0 = 20000
const X__WATCHOS_2_1 = 20100
const X__WATCHOS_2_2 = 20200
const X__WATCHOS_3_0 = 30000
const X__WATCHOS_3_1 = 30100
const X__WATCHOS_3_1_1 = 30101
const X__WATCHOS_3_2 = 30200
const X__WATCHOS_4_0 = 40000
const X__WATCHOS_4_1 = 40100
const X__WATCHOS_4_2 = 40200
const X__WATCHOS_4_3 = 40300
const X__WATCHOS_5_0 = 50000
const X__WATCHOS_5_1 = 50100
const X__WATCHOS_5_2 = 50200
const X__WATCHOS_5_3 = 50300
const X__WATCHOS_6_0 = 60000
const X__WATCHOS_6_1 = 60100
const X__WATCHOS_6_2 = 60200
const X__WATCHOS_7_0 = 70000
const X__WATCHOS_7_1 = 70100
const X__WATCHOS_7_2 = 70200
const X__WATCHOS_7_3 = 70300
const X__WATCHOS_7_4 = 70400
const X__WATCHOS_7_5 = 70500
const X__WATCHOS_7_6 = 70600
const X__WATCHOS_8_0 = 80000
const X__WATCHOS_8_1 = 80100
const X__WATCHOS_8_3 = 80300
const X__WATCHOS_8_4 = 80400
const X__WATCHOS_8_5 = 80500
const X__WATCHOS_8_6 = 80600
const X__WATCHOS_8_7 = 80700
const X__WATCHOS_8_8 = 80800
const X__WATCHOS_9_0 = 90000
const X__WATCHOS_9_1 = 90100
const X__WATCHOS_9_2 = 90200
const X__WATCHOS_9_3 = 90300
const X__WATCHOS_9_4 = 90400
const X__WATCHOS_9_5 = 90500
const X__WATCHOS_9_6 = 90600
const X__WATCHOS_10_0 = 100000
const X__WATCHOS_10_1 = 100100
const X__WATCHOS_10_2 = 100200
const X__WATCHOS_10_3 = 100300
const X__WATCHOS_10_4 = 100400
const X__TVOS_9_0 = 90000
const X__TVOS_9_1 = 90100
const X__TVOS_9_2 = 90200
const X__TVOS_10_0 = 100000
const X__TVOS_10_0_1 = 100001
const X__TVOS_10_1 = 100100
const X__TVOS_10_2 = 100200
const X__TVOS_11_0 = 110000
const X__TVOS_11_1 = 110100
const X__TVOS_11_2 = 110200
const X__TVOS_11_3 = 110300
const X__TVOS_11_4 = 110400
const X__TVOS_12_0 = 120000
const X__TVOS_12_1 = 120100
const X__TVOS_12_2 = 120200
const X__TVOS_12_3 = 120300
const X__TVOS_12_4 = 120400
const X__TVOS_13_0 = 130000
const X__TVOS_13_2 = 130200
const X__TVOS_13_3 = 130300
const X__TVOS_13_4 = 130400
const X__TVOS_14_0 = 140000
const X__TVOS_14_1 = 140100
const X__TVOS_14_2 = 140200
const X__TVOS_14_3 = 140300
const X__TVOS_14_5 = 140500
const X__TVOS_14_6 = 140600
const X__TVOS_14_7 = 140700
const X__TVOS_15_0 = 150000
const X__TVOS_15_1 = 150100
const X__TVOS_15_2 = 150200
const X__TVOS_15_3 = 150300
const X__TVOS_15_4 = 150400
const X__TVOS_15_5 = 150500
const X__TVOS_15_6 = 150600
const X__TVOS_16_0 = 160000
const X__TVOS_16_1 = 160100
const X__TVOS_16_2 = 160200
const X__TVOS_16_3 = 160300
const X__TVOS_16_4 = 160400
const X__TVOS_16_5 = 160500
const X__TVOS_16_6 = 160600
const X__TVOS_17_0 = 170000
const X__TVOS_17_1 = 170100
const X__TVOS_17_2 = 170200
const X__TVOS_17_3 = 170300
const X__TVOS_17_4 = 170400
const X__BRIDGEOS_2_0 = 20000
const X__BRIDGEOS_3_0 = 30000
const X__BRIDGEOS_3_1 = 30100
const X__BRIDGEOS_3_4 = 30400
const X__BRIDGEOS_4_0 = 40000
const X__BRIDGEOS_4_1 = 40100
const X__BRIDGEOS_5_0 = 50000
const X__BRIDGEOS_5_1 = 50100
const X__BRIDGEOS_5_3 = 50300
const X__BRIDGEOS_6_0 = 60000
const X__BRIDGEOS_6_2 = 60200
const X__BRIDGEOS_6_4 = 60400
const X__BRIDGEOS_6_5 = 60500
const X__BRIDGEOS_6_6 = 60600
const X__BRIDGEOS_7_0 = 70000
const X__BRIDGEOS_7_1 = 70100
const X__BRIDGEOS_7_2 = 70200
const X__BRIDGEOS_7_3 = 70300
const X__BRIDGEOS_7_4 = 70400
const X__BRIDGEOS_7_6 = 70600
const X__BRIDGEOS_8_0 = 80000
const X__BRIDGEOS_8_1 = 80100
const X__BRIDGEOS_8_2 = 80200
const X__BRIDGEOS_8_3 = 80300
const X__BRIDGEOS_8_4 = 80400
const X__DRIVERKIT_19_0 = 190000
const X__DRIVERKIT_20_0 = 200000
const X__DRIVERKIT_21_0 = 210000
const X__DRIVERKIT_22_0 = 220000
const X__DRIVERKIT_22_4 = 220400
const X__DRIVERKIT_22_5 = 220500
const X__DRIVERKIT_22_6 = 220600
const X__DRIVERKIT_23_0 = 230000
const X__DRIVERKIT_23_1 = 230100
const X__DRIVERKIT_23_2 = 230200
const X__DRIVERKIT_23_3 = 230300
const X__DRIVERKIT_23_4 = 230400
const X__VISIONOS_1_0 = 10000
const X__VISIONOS_1_1 = 10100
const X__ENABLE_LEGACY_MAC_AVAILABILITY = 1
const FP_NAN = 1
const FP_INFINITE = 2
const FP_ZERO = 3
const FP_NORMAL = 4
const FP_SUBNORMAL = 5
const FP_SUPERNORMAL = 6
const FP_FAST_FMA = 1
const FP_FAST_FMAF = 1
const FP_FAST_FMAL = 1
const MATH_ERRNO = 1
const MATH_ERREXCEPT = 2
const M_E = 2.71828182845904523536028747135266250
const M_LOG2E = 1.44269504088896340735992468100189214
const M_LOG10E = 0.434294481903251827651128918916605082
const M_LN2 = 0.693147180559945309417232121458176568
const M_LN10 = 2.30258509299404568401799145468436421
const M_PI = 3.14159265358979323846264338327950288
const M_PI_2 = 1.57079632679489661923132169163975144
const M_PI_4 = 0.785398163397448309615660845819875721
const M_1_PI = 0.318309886183790671537767526745028724
const M_2_PI = 0.636619772367581343075535053490057448
const M_2_SQRTPI = 1.12837916709551257389615890312154517
const M_SQRT2 = 1.41421356237309504880168872420969808
const M_SQRT1_2 = 0.707106781186547524400844362104849039
const X_TLOSS = 1.41484755040568800000e+16
const DOMAIN = 1
const SING = 2
const OVERFLOW = 3
const UNDERFLOW = 4
const TLOSS = 5
const PLOSS = 6
const X__WORDSIZE = 64
const X__PTHREAD_SIZE__ = 8176
const X__PTHREAD_ATTR_SIZE__ = 56
const X__PTHREAD_MUTEXATTR_SIZE__ = 8
const X__PTHREAD_MUTEX_SIZE__ = 56
const X__PTHREAD_CONDATTR_SIZE__ = 8
const X__PTHREAD_COND_SIZE__ = 40
const X__PTHREAD_ONCE_SIZE__ = 8
const X__PTHREAD_RWLOCK_SIZE__ = 192
const X__PTHREAD_RWLOCKATTR_SIZE__ = 16
const INT8_MAX = 127
const INT16_MAX = 32767
const INT32_MAX = 2147483647
const UINT8_MAX = 255
const UINT16_MAX = 65535
const X__DARWIN_CLK_TCK = 100
const MB_LEN_MAX = 6
const CHAR_BIT = 8
const SCHAR_MAX = 127
const UCHAR_MAX = 255
const CHAR_MAX = 127
const USHRT_MAX = 65535
const SHRT_MAX = 32767
const UINT_MAX = 0xffffffff
const INT_MAX = 2147483647
const LONG_BIT = 64
const WORD_BIT = 32
const CHILD_MAX = 266
const LINK_MAX = 32767
const MAX_CANON = 1024
const MAX_INPUT = 1024
const NAME_MAX = 255
const NGROUPS_MAX = 16
const OPEN_MAX = 10240
const PATH_MAX = 1024
const PIPE_BUF = 512
const BC_BASE_MAX = 99
const BC_DIM_MAX = 2048
const BC_SCALE_MAX = 99
const BC_STRING_MAX = 1000
const CHARCLASS_NAME_MAX = 14
const COLL_WEIGHTS_MAX = 2
const EQUIV_CLASS_MAX = 2
const EXPR_NEST_MAX = 32
const LINE_MAX = 2048
const RE_DUP_MAX = 255
const NZERO = 20
const X_POSIX_ARG_MAX = 4096
const X_POSIX_CHILD_MAX = 25
const X_POSIX_LINK_MAX = 8
const X_POSIX_MAX_CANON = 255
const X_POSIX_MAX_INPUT = 255
const X_POSIX_NAME_MAX = 14
const X_POSIX_NGROUPS_MAX = 8
const X_POSIX_OPEN_MAX = 20
const X_POSIX_PATH_MAX = 256
const X_POSIX_PIPE_BUF = 512
const X_POSIX_SSIZE_MAX = 32767
const X_POSIX_STREAM_MAX = 8
const X_POSIX_TZNAME_MAX = 6
const X_POSIX2_BC_BASE_MAX = 99
const X_POSIX2_BC_DIM_MAX = 2048
const X_POSIX2_BC_SCALE_MAX = 99
const X_POSIX2_BC_STRING_MAX = 1000
const X_POSIX2_EQUIV_CLASS_MAX = 2
const X_POSIX2_EXPR_NEST_MAX = 32
const X_POSIX2_LINE_MAX = 2048
const X_POSIX2_RE_DUP_MAX = 255
const X_POSIX_AIO_LISTIO_MAX = 2
const X_POSIX_AIO_MAX = 1
const X_POSIX_DELAYTIMER_MAX = 32
const X_POSIX_MQ_OPEN_MAX = 8
const X_POSIX_MQ_PRIO_MAX = 32
const X_POSIX_RTSIG_MAX = 8
const X_POSIX_SEM_NSEMS_MAX = 256
const X_POSIX_SEM_VALUE_MAX = 32767
const X_POSIX_SIGQUEUE_MAX = 32
const X_POSIX_TIMER_MAX = 32
const X_POSIX_CLOCKRES_MIN = 20000000
const X_POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
const X_POSIX_THREAD_KEYS_MAX = 128
const X_POSIX_THREAD_THREADS_MAX = 64
const PTHREAD_DESTRUCTOR_ITERATIONS = 4
const PTHREAD_KEYS_MAX = 512
const PTHREAD_STACK_MIN = 16384
const X_POSIX_HOST_NAME_MAX = 255
const X_POSIX_LOGIN_NAME_MAX = 9
const X_POSIX_SS_REPL_MAX = 4
const X_POSIX_SYMLINK_MAX = 255
const X_POSIX_SYMLOOP_MAX = 8
const X_POSIX_TRACE_EVENT_NAME_MAX = 30
const X_POSIX_TRACE_NAME_MAX = 8
const X_POSIX_TRACE_SYS_MAX = 8
const X_POSIX_TRACE_USER_EVENT_MAX = 32
const X_POSIX_TTY_NAME_MAX = 9
const X_POSIX2_CHARCLASS_NAME_MAX = 14
const X_POSIX2_COLL_WEIGHTS_MAX = 2
const PASS_MAX = 128
const NL_ARGMAX = 9
const NL_LANGMAX = 14
const NL_MSGMAX = 32767
const NL_NMAX = 1
const NL_SETMAX = 255
const NL_TEXTMAX = 2048
const X_XOPEN_IOV_MAX = 16
const IOV_MAX = 1024
const X_XOPEN_NAME_MAX = 255
const X_XOPEN_PATH_MAX = 1024
const X__noreturn_is_defined = 1
const X_FORTIFY_SOURCE = 2
const RENAME_SECLUDE = 0x00000001
const RENAME_SWAP = 0x00000002
const RENAME_EXCL = 0x00000004
const RENAME_RESERVED1 = 0x00000008
const RENAME_NOFOLLOW_ANY = 0x00000010
const SEEK_SET = 0
const SEEK_CUR = 1
const SEEK_END = 2
const SEEK_HOLE = 3
const SEEK_DATA = 4
const X__SLBF = 0x0001
const X__SNBF = 0x0002
const X__SRD = 0x0004
const X__SWR = 0x0008
const X__SRW = 0x0010
const X__SEOF = 0x0020
const X__SERR = 0x0040
const X__SMBF = 0x0080
const X__SAPP = 0x0100
const X__SSTR = 0x0200
const X__SOPT = 0x0400
const X__SNPT = 0x0800
const X__SOFF = 0x1000
const X__SMOD = 0x2000
const X__SALC = 0x4000
const X__SIGN = 0x8000
const X_IOFBF = 0
const X_IOLBF = 1
const X_IONBF = 2
const BUFSIZ = 1024
const FOPEN_MAX = 20
const FILENAME_MAX = 1024
const P_tmpdir = "/var/tmp/"
const L_tmpnam = 1024
const TMP_MAX = 308915776
const L_ctermid = 1024
const X_USE_FORTIFY_LEVEL = 2
const UNITY_DETAIL1_NAME = "Function"
const UNITY_DETAIL2_NAME = "Argument"

type JmpBuf [48]c.Int
type SigjmpBuf [49]c.Int
type FloatT c.Float
type DoubleT c.Double

type X__float2 struct {
	X__sinval c.Float
	X__cosval c.Float
}

type X__double2 struct {
	X__sinval c.Double
	X__cosval c.Double
}

type Exception struct {
	Type   c.Int
	Name   *c.Char
	Arg1   c.Double
	Arg2   c.Double
	Retval c.Double
}
type PtrdiffT c.Long
type RsizeT c.Ulong
type WcharT c.Int
type MaxAlignT c.Double
type IntLeast8T c.Int8T
type IntLeast16T c.Int16T
type IntLeast32T c.Int32T
type IntLeast64T c.Int64T
type UintLeast8T c.Uint8T
type UintLeast16T c.Uint16T
type UintLeast32T c.Uint32T
type UintLeast64T c.Uint64T
type IntFast8T c.Int8T
type IntFast16T c.Int16T
type IntFast32T c.Int32T
type IntFast64T c.Int64T
type UintFast8T c.Uint8T
type UintFast16T c.Uint16T
type UintFast32T c.Uint32T
type UintFast64T c.Uint64T
type X__int8T c.Char
type X__uint8T c.Char
type X__int16T int16
type X__uint16T uint16
type X__int32T c.Int
type X__uint32T c.Uint
type X__int64T c.LongLong
type X__uint64T c.UlongLong
type X__darwinIntptrT c.Long
type X__darwinNaturalT c.Uint
type X__darwinCtRuneT c.Int

type X__mbstateT struct {
	X__mbstate8 [128]c.Char
}
type X__darwinMbstateT X__mbstateT
type X__darwinPtrdiffT c.Long
type X__darwinSizeT c.Ulong
type X__darwinVaList c.Pointer
type X__darwinWcharT c.Int
type X__darwinRuneT X__darwinWcharT
type X__darwinWintT c.Int
type X__darwinClockT c.Ulong
type X__darwinSocklenT X__uint32T
type X__darwinSsizeT c.Long
type X__darwinTimeT c.Long
type X__darwinBlkcntT X__int64T
type X__darwinBlksizeT X__int32T
type X__darwinDevT X__int32T
type X__darwinFsblkcntT c.Uint
type X__darwinFsfilcntT c.Uint
type X__darwinGidT X__uint32T
type X__darwinIdT X__uint32T
type X__darwinIno64T X__uint64T
type X__darwinInoT X__darwinIno64T
type X__darwinMachPortNameT X__darwinNaturalT
type X__darwinMachPortT X__darwinMachPortNameT
type X__darwinModeT X__uint16T
type X__darwinOffT X__int64T
type X__darwinPidT X__int32T
type X__darwinSigsetT X__uint32T
type X__darwinSusecondsT X__int32T
type X__darwinUidT X__uint32T
type X__darwinUsecondsT X__uint32T
type X__darwinUuidT [16]c.Char
type X__darwinUuidStringT [37]c.Char

type X__darwinPthreadHandlerRec struct {
	X__routine c.Pointer
	X__arg     c.Pointer
	X__next    *X__darwinPthreadHandlerRec
}

type X_opaquePthreadAttrT struct {
	X__sig    c.Long
	X__opaque [56]c.Char
}

type X_opaquePthreadCondT struct {
	X__sig    c.Long
	X__opaque [40]c.Char
}

type X_opaquePthreadCondattrT struct {
	X__sig    c.Long
	X__opaque [8]c.Char
}

type X_opaquePthreadMutexT struct {
	X__sig    c.Long
	X__opaque [56]c.Char
}

type X_opaquePthreadMutexattrT struct {
	X__sig    c.Long
	X__opaque [8]c.Char
}

type X_opaquePthreadOnceT struct {
	X__sig    c.Long
	X__opaque [8]c.Char
}

type X_opaquePthreadRwlockT struct {
	X__sig    c.Long
	X__opaque [192]c.Char
}

type X_opaquePthreadRwlockattrT struct {
	X__sig    c.Long
	X__opaque [16]c.Char
}

type X_opaquePthreadT struct {
	X__sig          c.Long
	X__cleanupStack *X__darwinPthreadHandlerRec
	X__opaque       [8176]c.Char
}
type X__darwinPthreadAttrT X_opaquePthreadAttrT
type X__darwinPthreadCondT X_opaquePthreadCondT
type X__darwinPthreadCondattrT X_opaquePthreadCondattrT
type X__darwinPthreadKeyT c.Ulong
type X__darwinPthreadMutexT X_opaquePthreadMutexT
type X__darwinPthreadMutexattrT X_opaquePthreadMutexattrT
type X__darwinPthreadOnceT X_opaquePthreadOnceT
type X__darwinPthreadRwlockT X_opaquePthreadRwlockT
type X__darwinPthreadRwlockattrT X_opaquePthreadRwlockattrT
type X__darwinPthreadT *X_opaquePthreadT
type UInt8T c.Char
type UInt16T uint16
type UInt32T c.Uint
type UInt64T c.UlongLong
type RegisterT c.Int64T
type UserAddrT UInt64T
type UserSizeT UInt64T
type UserSsizeT c.Int64T
type UserLongT c.Int64T
type UserUlongT UInt64T
type UserTimeT c.Int64T
type UserOffT c.Int64T
type SyscallArgT UInt64T
type UNITYUINT8 c.Char
type UNITYUINT16 uint16
type UNITYUINT32 c.Uint
type UNITYINT8 c.Char
type UNITYINT16 int16
type UNITYINT32 c.Int
type UNITYUINT64 c.Ulong
type UNITYINT64 c.Long
type UNITYUINT UNITYUINT64
type UNITYINT UNITYINT64
type UNITYFLOAT c.Float
type UNITYDOUBLE UNITYFLOAT
type X__darwinNlItem c.Int
type X__darwinWctransT c.Int
type X__darwinWctypeT X__uint32T
type FposT X__darwinOffT

type X__sbuf struct {
	X_base *c.Char
	X_size c.Int
}

type X__sFILEX struct {
	Unused [8]uint8
}

type X__sFILE struct {
	X_p       *c.Char
	X_r       c.Int
	X_w       c.Int
	X_flags   int16
	X_file    int16
	X_bf      X__sbuf
	X_lbfsize c.Int
	X_cookie  c.Pointer
	X_close   c.Pointer
	X_read    c.Pointer
	X_seek    c.Pointer
	X_write   c.Pointer
	X_ub      X__sbuf
	X_extra   *X__sFILEX
	X_ur      c.Int
	X_ubuf    [3]c.Char
	X_nbuf    [1]c.Char
	X_lb      X__sbuf
	X_blksize c.Int
	X_offset  FposT
}
type OffT X__darwinOffT

// llgo:type C
type UnityTestFunction func()
type UNITYDISPLAYSTYLET c.Int

const (
	UNITY_DISPLAY_STYLE_INT     UNITYDISPLAYSTYLET = 20
	UNITY_DISPLAY_STYLE_INT8    UNITYDISPLAYSTYLET = 17
	UNITY_DISPLAY_STYLE_INT16   UNITYDISPLAYSTYLET = 18
	UNITY_DISPLAY_STYLE_INT32   UNITYDISPLAYSTYLET = 20
	UNITY_DISPLAY_STYLE_INT64   UNITYDISPLAYSTYLET = 24
	UNITY_DISPLAY_STYLE_UINT    UNITYDISPLAYSTYLET = 36
	UNITY_DISPLAY_STYLE_UINT8   UNITYDISPLAYSTYLET = 33
	UNITY_DISPLAY_STYLE_UINT16  UNITYDISPLAYSTYLET = 34
	UNITY_DISPLAY_STYLE_UINT32  UNITYDISPLAYSTYLET = 36
	UNITY_DISPLAY_STYLE_UINT64  UNITYDISPLAYSTYLET = 40
	UNITY_DISPLAY_STYLE_HEX8    UNITYDISPLAYSTYLET = 65
	UNITY_DISPLAY_STYLE_HEX16   UNITYDISPLAYSTYLET = 66
	UNITY_DISPLAY_STYLE_HEX32   UNITYDISPLAYSTYLET = 68
	UNITY_DISPLAY_STYLE_HEX64   UNITYDISPLAYSTYLET = 72
	UNITY_DISPLAY_STYLE_CHAR    UNITYDISPLAYSTYLET = 145
	UNITY_DISPLAY_STYLE_UNKNOWN UNITYDISPLAYSTYLET = 146
)

type UNITYCOMPARISONT c.Int

const (
	UNITY_WITHIN           UNITYCOMPARISONT = 0
	UNITY_EQUAL_TO         UNITYCOMPARISONT = 1
	UNITY_GREATER_THAN     UNITYCOMPARISONT = 2
	UNITY_GREATER_OR_EQUAL UNITYCOMPARISONT = 3
	UNITY_SMALLER_THAN     UNITYCOMPARISONT = 4
	UNITY_SMALLER_OR_EQUAL UNITYCOMPARISONT = 5
	UNITY_NOT_EQUAL        UNITYCOMPARISONT = 0
	UNITY_UNKNOWN          UNITYCOMPARISONT = 1
)

type UNITYFLOATTRAIT c.Int

const (
	UNITY_FLOAT_IS_NOT_INF     UNITYFLOATTRAIT = 0
	UNITY_FLOAT_IS_INF         UNITYFLOATTRAIT = 1
	UNITY_FLOAT_IS_NOT_NEG_INF UNITYFLOATTRAIT = 2
	UNITY_FLOAT_IS_NEG_INF     UNITYFLOATTRAIT = 3
	UNITY_FLOAT_IS_NOT_NAN     UNITYFLOATTRAIT = 4
	UNITY_FLOAT_IS_NAN         UNITYFLOATTRAIT = 5
	UNITY_FLOAT_IS_NOT_DET     UNITYFLOATTRAIT = 6
	UNITY_FLOAT_IS_DET         UNITYFLOATTRAIT = 7
	UNITY_FLOAT_INVALID_TRAIT  UNITYFLOATTRAIT = 8
)

type UNITYFLOATTRAITT UNITYFLOATTRAIT
type UNITYFLAGST c.Int

const (
	UNITY_ARRAY_TO_VAL   UNITYFLAGST = 0
	UNITY_ARRAY_TO_ARRAY UNITYFLAGST = 1
	UNITY_ARRAY_UNKNOWN  UNITYFLAGST = 2
)

type UNITYSTORAGET struct {
	TestFile              *c.Char
	CurrentTestName       *c.Char
	CurrentDetail1        *c.Char
	CurrentDetail2        *c.Char
	CurrentTestLineNumber UNITYUINT
	NumberOfTests         UNITYUINT
	TestFailures          UNITYUINT
	TestIgnores           UNITYUINT
	CurrentTestFailed     UNITYUINT
	CurrentTestIgnored    UNITYUINT
	AbortFrame            JmpBuf
}
