package http_parser

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

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
const X__PTHREAD_SIZE__ = 8176
const X__PTHREAD_ATTR_SIZE__ = 56
const X__PTHREAD_MUTEXATTR_SIZE__ = 8
const X__PTHREAD_MUTEX_SIZE__ = 56
const X__PTHREAD_CONDATTR_SIZE__ = 8
const X__PTHREAD_COND_SIZE__ = 40
const X__PTHREAD_ONCE_SIZE__ = 8
const X__PTHREAD_RWLOCK_SIZE__ = 192
const X__PTHREAD_RWLOCKATTR_SIZE__ = 16
const X_QUAD_HIGHWORD = 1
const X_QUAD_LOWWORD = 0
const X__DARWIN_LITTLE_ENDIAN = 1234
const X__DARWIN_BIG_ENDIAN = 4321
const X__DARWIN_PDP_ENDIAN = 3412
const X__WORDSIZE = 64
const INT8_MAX = 127
const INT16_MAX = 32767
const INT32_MAX = 2147483647
const UINT8_MAX = 255
const UINT16_MAX = 65535
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
const X__DARWIN_FD_SETSIZE = 1024
const X__DARWIN_NBBY = 8

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

type X_OSUnalignedU16 struct {
	X__val c.Uint16T
}

type X_OSUnalignedU32 struct {
	X__val c.Uint32T
}

type X_OSUnalignedU64 struct {
	X__val c.Uint64T
}
type UChar c.Char
type UShort uint16
type UInt c.Uint
type ULong c.Ulong
type Ushort uint16
type Uint c.Uint
type UQuadT UInt64T
type QuadT c.Int64T
type QaddrT *QuadT
type CaddrT *c.Char
type DaddrT c.Int32T
type DevT X__darwinDevT
type FixptT UInt32T
type BlkcntT X__darwinBlkcntT
type BlksizeT X__darwinBlksizeT
type GidT X__darwinGidT
type InAddrT X__uint32T
type InPortT X__uint16T
type InoT X__darwinInoT
type Ino64T X__darwinIno64T
type KeyT X__int32T
type ModeT X__darwinModeT
type NlinkT X__uint16T
type IdT X__darwinIdT
type PidT X__darwinPidT
type OffT X__darwinOffT
type SegszT c.Int32T
type SwblkT c.Int32T
type UidT X__darwinUidT
type ClockT X__darwinClockT
type TimeT X__darwinTimeT
type UsecondsT X__darwinUsecondsT
type SusecondsT X__darwinSusecondsT
type RsizeT X__darwinSizeT
type ErrnoT c.Int

type FdSet struct {
	FdsBits [32]X__int32T
}
type FdMask X__int32T
type PthreadAttrT X__darwinPthreadAttrT
type PthreadCondT X__darwinPthreadCondT
type PthreadCondattrT X__darwinPthreadCondattrT
type PthreadMutexT X__darwinPthreadMutexT
type PthreadMutexattrT X__darwinPthreadMutexattrT
type PthreadOnceT X__darwinPthreadOnceT
type PthreadRwlockT X__darwinPthreadRwlockT
type PthreadRwlockattrT X__darwinPthreadRwlockattrT
type PthreadT X__darwinPthreadT
type PthreadKeyT X__darwinPthreadKeyT
type FsblkcntT X__darwinFsblkcntT
type FsfilcntT X__darwinFsfilcntT
