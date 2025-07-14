package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MIN_ICACHE_SIZE = 16384
const MAX_ICACHE_SIZE = 16384
const MIN_ICACHE_WAYS = 8
const MAX_ICACHE_WAYS = 8
const MAX_CACHE_WAYS = 8
const MIN_CACHE_LINE_SIZE = 32
const TAG_SIZE = 4
const MIN_ICACHE_BANK_NUM = 1
const MAX_ICACHE_BANK_NUM = 1
const CACHE_MEMORY_BANK_NUM = 1
const CACHE_MEMORY_IBANK_SIZE = 0x4000
const ESP_ROM_ERR_INVALID_ARG = 1
const MMU_SET_ADDR_ALIGNED_ERROR = 2
const MMU_SET_PASE_SIZE_ERROR = 3
const MMU_SET_VADDR_OUT_RANGE = 4
const CACHE_OP_ICACHE_Y = 1
const CACHE_OP_ICACHE_N = 0

type CacheSizeT c.Int

const (
	CACHE_SIZE_HALF CacheSizeT = 0
	CACHE_SIZE_FULL CacheSizeT = 1
)

type CacheWaysT c.Int

const (
	CACHE_4WAYS_ASSOC CacheWaysT = 0
	CACHE_8WAYS_ASSOC CacheWaysT = 1
)

type CacheLineSizeT c.Int

const (
	CACHE_LINE_SIZE_16B CacheLineSizeT = 0
	CACHE_LINE_SIZE_32B CacheLineSizeT = 1
	CACHE_LINE_SIZE_64B CacheLineSizeT = 2
)

type CacheAutoloadOrderT c.Int

const (
	CACHE_AUTOLOAD_POSITIVE CacheAutoloadOrderT = 0
	CACHE_AUTOLOAD_NEGATIVE CacheAutoloadOrderT = 1
)

type CacheAutoloadTriggerT c.Int

const (
	CACHE_AUTOLOAD_MISS_TRIGGER CacheAutoloadTriggerT = 0
	CACHE_AUTOLOAD_HIT_TRIGGER  CacheAutoloadTriggerT = 1
	CACHE_AUTOLOAD_BOTH_TRIGGER CacheAutoloadTriggerT = 2
)

type CacheFreezeModeT c.Int

const (
	CACHE_FREEZE_ACK_BUSY  CacheFreezeModeT = 0
	CACHE_FREEZE_ACK_ERROR CacheFreezeModeT = 1
)

type MmuPageModeT c.Int

const (
	MMU_PAGE_MODE_64KB    MmuPageModeT = 0
	MMU_PAGE_MODE_32KB    MmuPageModeT = 1
	MMU_PAGE_MODE_16KB    MmuPageModeT = 2
	MMU_PAGE_MODE_8KB     MmuPageModeT = 3
	MMU_PAGE_MODE_INVALID MmuPageModeT = 4
)

type CacheMode struct {
	CacheSize     c.Uint32T
	CacheLineSize c.Uint16T
	CacheWays     c.Uint8T
	Ibus          c.Uint8T
}

type IcacheTagItem struct {
	Valid    c.Uint32T
	Lock     c.Uint32T
	FifoCnt  c.Uint32T
	Tag      c.Uint32T
	Reserved c.Uint32T
}

type AutoloadConfig struct {
	Order   c.Uint8T
	Trigger c.Uint8T
	Ena0    c.Uint8T
	Ena1    c.Uint8T
	Addr0   c.Uint32T
	Size0   c.Uint32T
	Addr1   c.Uint32T
	Size1   c.Uint32T
}

type TagGroupInfo struct {
	Mode              CacheMode
	FilterAddr        c.Uint32T
	VaddrOffset       c.Uint32T
	TagAddr           [8]c.Uint32T
	CacheMemoryOffset [8]c.Uint32T
}

type LockConfig struct {
	Addr  c.Uint32T
	Size  c.Uint16T
	Group c.Uint16T
}

type CacheInternalStubTable struct {
	IcacheLineSize        c.Pointer
	IcacheAddr            c.Pointer
	DcacheAddr            c.Pointer
	InvalidateIcacheItems c.Pointer
	LockIcacheItems       c.Pointer
	UnlockIcacheItems     c.Pointer
	SuspendIcacheAutoload c.Pointer
	ResumeIcacheAutoload  c.Pointer
	FreezeIcacheEnable    c.Pointer
	FreezeIcacheDisable   c.Pointer
	OpAddr                c.Pointer
}

// llgo:type C
type CacheOpStart func()

// llgo:type C
type CacheOpEnd func()

type CacheOpCbT struct {
	Start CacheOpStart
	End   CacheOpEnd
}
