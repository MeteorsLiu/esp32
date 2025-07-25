package soc

import _ "unsafe"

const CACHE_L1_CACHE_SHUT_BUS1_V = 0x1
const CACHE_L1_CACHE_SHUT_BUS1_S = 1
const CACHE_L1_CACHE_SHUT_BUS0_V = 0x1
const CACHE_L1_CACHE_SHUT_BUS0_S = 0
const CACHE_L1_CACHE_WRAP_V = 0x1
const CACHE_L1_CACHE_WRAP_S = 4
const CACHE_L1_CACHE_TAG_MEM_FORCE_PU_V = 0x1
const CACHE_L1_CACHE_TAG_MEM_FORCE_PU_S = 18
const CACHE_L1_CACHE_TAG_MEM_FORCE_PD_V = 0x1
const CACHE_L1_CACHE_TAG_MEM_FORCE_PD_S = 17
const CACHE_L1_CACHE_TAG_MEM_FORCE_ON_V = 0x1
const CACHE_L1_CACHE_TAG_MEM_FORCE_ON_S = 16
const CACHE_L1_CACHE_DATA_MEM_FORCE_PU_V = 0x1
const CACHE_L1_CACHE_DATA_MEM_FORCE_PU_S = 18
const CACHE_L1_CACHE_DATA_MEM_FORCE_PD_V = 0x1
const CACHE_L1_CACHE_DATA_MEM_FORCE_PD_S = 17
const CACHE_L1_CACHE_DATA_MEM_FORCE_ON_V = 0x1
const CACHE_L1_CACHE_DATA_MEM_FORCE_ON_S = 16
const CACHE_L1_CACHE_FREEZE_DONE_V = 0x1
const CACHE_L1_CACHE_FREEZE_DONE_S = 18
const CACHE_L1_CACHE_FREEZE_MODE_V = 0x1
const CACHE_L1_CACHE_FREEZE_MODE_S = 17
const CACHE_L1_CACHE_FREEZE_EN_V = 0x1
const CACHE_L1_CACHE_FREEZE_EN_S = 16
const CACHE_L1_CACHE_DATA_MEM_WR_EN_V = 0x1
const CACHE_L1_CACHE_DATA_MEM_WR_EN_S = 17
const CACHE_L1_CACHE_DATA_MEM_RD_EN_V = 0x1
const CACHE_L1_CACHE_DATA_MEM_RD_EN_S = 16
const CACHE_L1_CACHE_TAG_MEM_WR_EN_V = 0x1
const CACHE_L1_CACHE_TAG_MEM_WR_EN_S = 17
const CACHE_L1_CACHE_TAG_MEM_RD_EN_V = 0x1
const CACHE_L1_CACHE_TAG_MEM_RD_EN_S = 16
const CACHE_L1_CACHE_PRELOCK_RGID = 0x0000000F
const CACHE_L1_CACHE_PRELOCK_RGID_V = 0xF
const CACHE_L1_CACHE_PRELOCK_RGID_S = 2
const CACHE_L1_CACHE_PRELOCK_SCT1_EN_V = 0x1
const CACHE_L1_CACHE_PRELOCK_SCT1_EN_S = 1
const CACHE_L1_CACHE_PRELOCK_SCT0_EN_V = 0x1
const CACHE_L1_CACHE_PRELOCK_SCT0_EN_S = 0
const CACHE_L1_CACHE_PRELOCK_SCT0_ADDR = 0xFFFFFFFF
const CACHE_L1_CACHE_PRELOCK_SCT0_ADDR_V = 0xFFFFFFFF
const CACHE_L1_CACHE_PRELOCK_SCT0_ADDR_S = 0
const CACHE_L1_CACHE_PRELOCK_SCT1_ADDR = 0xFFFFFFFF
const CACHE_L1_CACHE_PRELOCK_SCT1_ADDR_V = 0xFFFFFFFF
const CACHE_L1_CACHE_PRELOCK_SCT1_ADDR_S = 0
const CACHE_L1_CACHE_PRELOCK_SCT1_SIZE = 0x00003FFF
const CACHE_L1_CACHE_PRELOCK_SCT1_SIZE_V = 0x3FFF
const CACHE_L1_CACHE_PRELOCK_SCT1_SIZE_S = 16
const CACHE_L1_CACHE_PRELOCK_SCT0_SIZE = 0x00003FFF
const CACHE_L1_CACHE_PRELOCK_SCT0_SIZE_V = 0x3FFF
const CACHE_L1_CACHE_PRELOCK_SCT0_SIZE_S = 0
const CACHE_LOCK_DONE_V = 0x1
const CACHE_LOCK_DONE_S = 2
const CACHE_UNLOCK_ENA_V = 0x1
const CACHE_UNLOCK_ENA_S = 1
const CACHE_LOCK_ENA_V = 0x1
const CACHE_LOCK_ENA_S = 0
const CACHE_LOCK_MAP = 0x0000003F
const CACHE_LOCK_MAP_V = 0x3F
const CACHE_LOCK_MAP_S = 0
const CACHE_LOCK_ADDR = 0xFFFFFFFF
const CACHE_LOCK_ADDR_V = 0xFFFFFFFF
const CACHE_LOCK_ADDR_S = 0
const CACHE_LOCK_SIZE = 0x0000FFFF
const CACHE_LOCK_SIZE_V = 0xFFFF
const CACHE_LOCK_SIZE_S = 0
const CACHE_SYNC_DONE_V = 0x1
const CACHE_SYNC_DONE_S = 4
const CACHE_WRITEBACK_INVALIDATE_ENA_V = 0x1
const CACHE_WRITEBACK_INVALIDATE_ENA_S = 3
const CACHE_WRITEBACK_ENA_V = 0x1
const CACHE_WRITEBACK_ENA_S = 2
const CACHE_CLEAN_ENA_V = 0x1
const CACHE_CLEAN_ENA_S = 1
const CACHE_INVALIDATE_ENA_V = 0x1
const CACHE_INVALIDATE_ENA_S = 0
const CACHE_SYNC_MAP = 0x0000003F
const CACHE_SYNC_MAP_V = 0x3F
const CACHE_SYNC_MAP_S = 0
const CACHE_SYNC_ADDR = 0xFFFFFFFF
const CACHE_SYNC_ADDR_V = 0xFFFFFFFF
const CACHE_SYNC_ADDR_S = 0
const CACHE_SYNC_SIZE = 0x00FFFFFF
const CACHE_SYNC_SIZE_V = 0xFFFFFF
const CACHE_SYNC_SIZE_S = 0
const CACHE_L1_CACHE_PRELOAD_RGID = 0x0000000F
const CACHE_L1_CACHE_PRELOAD_RGID_V = 0xF
const CACHE_L1_CACHE_PRELOAD_RGID_S = 3
const CACHE_L1_CACHE_PRELOAD_ORDER_V = 0x1
const CACHE_L1_CACHE_PRELOAD_ORDER_S = 2
const CACHE_L1_CACHE_PRELOAD_DONE_V = 0x1
const CACHE_L1_CACHE_PRELOAD_DONE_S = 1
const CACHE_L1_CACHE_PRELOAD_ENA_V = 0x1
const CACHE_L1_CACHE_PRELOAD_ENA_S = 0
const CACHE_L1_CACHE_PRELOAD_ADDR = 0xFFFFFFFF
const CACHE_L1_CACHE_PRELOAD_ADDR_V = 0xFFFFFFFF
const CACHE_L1_CACHE_PRELOAD_ADDR_S = 0
const CACHE_L1_CACHE_PRELOAD_SIZE = 0x00003FFF
const CACHE_L1_CACHE_PRELOAD_SIZE_V = 0x3FFF
const CACHE_L1_CACHE_PRELOAD_SIZE_S = 0
const CACHE_L1_CACHE_AUTOLOAD_SCT1_ENA_V = 0x1
const CACHE_L1_CACHE_AUTOLOAD_SCT1_ENA_S = 9
const CACHE_L1_CACHE_AUTOLOAD_SCT0_ENA_V = 0x1
const CACHE_L1_CACHE_AUTOLOAD_SCT0_ENA_S = 8
const CACHE_L1_CACHE_AUTOLOAD_TRIGGER_MODE = 0x00000003
const CACHE_L1_CACHE_AUTOLOAD_TRIGGER_MODE_V = 0x3
const CACHE_L1_CACHE_AUTOLOAD_TRIGGER_MODE_S = 3
const CACHE_L1_CACHE_AUTOLOAD_ORDER_V = 0x1
const CACHE_L1_CACHE_AUTOLOAD_ORDER_S = 2
const CACHE_L1_CACHE_AUTOLOAD_DONE_V = 0x1
const CACHE_L1_CACHE_AUTOLOAD_DONE_S = 1
const CACHE_L1_CACHE_AUTOLOAD_ENA_V = 0x1
const CACHE_L1_CACHE_AUTOLOAD_ENA_S = 0
const CACHE_L1_CACHE_AUTOLOAD_SCT0_ADDR = 0xFFFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT0_ADDR_V = 0xFFFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT0_ADDR_S = 0
const CACHE_L1_CACHE_AUTOLOAD_SCT0_SIZE = 0x00FFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT0_SIZE_V = 0xFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT0_SIZE_S = 0
const CACHE_L1_CACHE_AUTOLOAD_SCT1_ADDR = 0xFFFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT1_ADDR_V = 0xFFFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT1_ADDR_S = 0
const CACHE_L1_CACHE_AUTOLOAD_SCT1_SIZE = 0x00FFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT1_SIZE_V = 0xFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT1_SIZE_S = 0
const CACHE_L1_CACHE_AUTOLOAD_SCT2_ADDR = 0xFFFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT2_ADDR_V = 0xFFFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT2_ADDR_S = 0
const CACHE_L1_CACHE_AUTOLOAD_SCT2_SIZE = 0x00FFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT2_SIZE_V = 0xFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT2_SIZE_S = 0
const CACHE_L1_CACHE_AUTOLOAD_SCT3_ADDR = 0xFFFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT3_ADDR_V = 0xFFFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT3_ADDR_S = 0
const CACHE_L1_CACHE_AUTOLOAD_SCT3_SIZE = 0x00FFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT3_SIZE_V = 0xFFFFFF
const CACHE_L1_CACHE_AUTOLOAD_SCT3_SIZE_S = 0
const CACHE_L1_BUS1_OVF_INT_ENA_V = 0x1
const CACHE_L1_BUS1_OVF_INT_ENA_S = 5
const CACHE_L1_BUS0_OVF_INT_ENA_V = 0x1
const CACHE_L1_BUS0_OVF_INT_ENA_S = 4
const CACHE_L1_BUS1_OVF_INT_CLR_V = 0x1
const CACHE_L1_BUS1_OVF_INT_CLR_S = 5
const CACHE_L1_BUS0_OVF_INT_CLR_V = 0x1
const CACHE_L1_BUS0_OVF_INT_CLR_S = 4
const CACHE_L1_BUS1_OVF_INT_RAW_V = 0x1
const CACHE_L1_BUS1_OVF_INT_RAW_S = 5
const CACHE_L1_BUS0_OVF_INT_RAW_V = 0x1
const CACHE_L1_BUS0_OVF_INT_RAW_S = 4
const CACHE_L1_BUS1_OVF_INT_ST_V = 0x1
const CACHE_L1_BUS1_OVF_INT_ST_S = 5
const CACHE_L1_BUS0_OVF_INT_ST_V = 0x1
const CACHE_L1_BUS0_OVF_INT_ST_S = 4
const CACHE_L1_CACHE_FAIL_INT_ENA_V = 0x1
const CACHE_L1_CACHE_FAIL_INT_ENA_S = 4
const CACHE_L1_CACHE_FAIL_INT_CLR_V = 0x1
const CACHE_L1_CACHE_FAIL_INT_CLR_S = 4
const CACHE_L1_CACHE_FAIL_INT_RAW_V = 0x1
const CACHE_L1_CACHE_FAIL_INT_RAW_S = 4
const CACHE_L1_CACHE_FAIL_INT_ST_V = 0x1
const CACHE_L1_CACHE_FAIL_INT_ST_S = 4
const CACHE_L1_BUS1_CNT_CLR_V = 0x1
const CACHE_L1_BUS1_CNT_CLR_S = 21
const CACHE_L1_BUS0_CNT_CLR_V = 0x1
const CACHE_L1_BUS0_CNT_CLR_S = 20
const CACHE_L1_BUS1_CNT_ENA_V = 0x1
const CACHE_L1_BUS1_CNT_ENA_S = 5
const CACHE_L1_BUS0_CNT_ENA_V = 0x1
const CACHE_L1_BUS0_CNT_ENA_S = 4
const CACHE_L1_BUS0_HIT_CNT = 0xFFFFFFFF
const CACHE_L1_BUS0_HIT_CNT_V = 0xFFFFFFFF
const CACHE_L1_BUS0_HIT_CNT_S = 0
const CACHE_L1_BUS0_MISS_CNT = 0xFFFFFFFF
const CACHE_L1_BUS0_MISS_CNT_V = 0xFFFFFFFF
const CACHE_L1_BUS0_MISS_CNT_S = 0
const CACHE_L1_BUS0_CONFLICT_CNT = 0xFFFFFFFF
const CACHE_L1_BUS0_CONFLICT_CNT_V = 0xFFFFFFFF
const CACHE_L1_BUS0_CONFLICT_CNT_S = 0
const CACHE_L1_BUS0_NXTLVL_CNT = 0xFFFFFFFF
const CACHE_L1_BUS0_NXTLVL_CNT_V = 0xFFFFFFFF
const CACHE_L1_BUS0_NXTLVL_CNT_S = 0
const CACHE_L1_BUS1_HIT_CNT = 0xFFFFFFFF
const CACHE_L1_BUS1_HIT_CNT_V = 0xFFFFFFFF
const CACHE_L1_BUS1_HIT_CNT_S = 0
const CACHE_L1_BUS1_MISS_CNT = 0xFFFFFFFF
const CACHE_L1_BUS1_MISS_CNT_V = 0xFFFFFFFF
const CACHE_L1_BUS1_MISS_CNT_S = 0
const CACHE_L1_BUS1_CONFLICT_CNT = 0xFFFFFFFF
const CACHE_L1_BUS1_CONFLICT_CNT_V = 0xFFFFFFFF
const CACHE_L1_BUS1_CONFLICT_CNT_S = 0
const CACHE_L1_BUS1_NXTLVL_CNT = 0xFFFFFFFF
const CACHE_L1_BUS1_NXTLVL_CNT_V = 0xFFFFFFFF
const CACHE_L1_BUS1_NXTLVL_CNT_S = 0
const CACHE_L1_CACHE_FAIL_ATTR = 0x0000FFFF
const CACHE_L1_CACHE_FAIL_ATTR_V = 0xFFFF
const CACHE_L1_CACHE_FAIL_ATTR_S = 16
const CACHE_L1_CACHE_FAIL_ID = 0x0000FFFF
const CACHE_L1_CACHE_FAIL_ID_V = 0xFFFF
const CACHE_L1_CACHE_FAIL_ID_S = 0
const CACHE_L1_CACHE_FAIL_ADDR = 0xFFFFFFFF
const CACHE_L1_CACHE_FAIL_ADDR_V = 0xFFFFFFFF
const CACHE_L1_CACHE_FAIL_ADDR_S = 0
const CACHE_SYNC_ERR_INT_ENA_V = 0x1
const CACHE_SYNC_ERR_INT_ENA_S = 13
const CACHE_L1_CACHE_PLD_ERR_INT_ENA_V = 0x1
const CACHE_L1_CACHE_PLD_ERR_INT_ENA_S = 11
const CACHE_SYNC_DONE_INT_ENA_V = 0x1
const CACHE_SYNC_DONE_INT_ENA_S = 6
const CACHE_L1_CACHE_PLD_DONE_INT_ENA_V = 0x1
const CACHE_L1_CACHE_PLD_DONE_INT_ENA_S = 4
const CACHE_SYNC_ERR_INT_CLR_V = 0x1
const CACHE_SYNC_ERR_INT_CLR_S = 13
const CACHE_L1_CACHE_PLD_ERR_INT_CLR_V = 0x1
const CACHE_L1_CACHE_PLD_ERR_INT_CLR_S = 11
const CACHE_SYNC_DONE_INT_CLR_V = 0x1
const CACHE_SYNC_DONE_INT_CLR_S = 6
const CACHE_L1_CACHE_PLD_DONE_INT_CLR_V = 0x1
const CACHE_L1_CACHE_PLD_DONE_INT_CLR_S = 4
const CACHE_SYNC_ERR_INT_RAW_V = 0x1
const CACHE_SYNC_ERR_INT_RAW_S = 13
const CACHE_L1_CACHE_PLD_ERR_INT_RAW_V = 0x1
const CACHE_L1_CACHE_PLD_ERR_INT_RAW_S = 11
const CACHE_SYNC_DONE_INT_RAW_V = 0x1
const CACHE_SYNC_DONE_INT_RAW_S = 6
const CACHE_L1_CACHE_PLD_DONE_INT_RAW_V = 0x1
const CACHE_L1_CACHE_PLD_DONE_INT_RAW_S = 4
const CACHE_SYNC_ERR_INT_ST_V = 0x1
const CACHE_SYNC_ERR_INT_ST_S = 13
const CACHE_L1_CACHE_PLD_ERR_INT_ST_V = 0x1
const CACHE_L1_CACHE_PLD_ERR_INT_ST_S = 11
const CACHE_SYNC_DONE_INT_ST_V = 0x1
const CACHE_SYNC_DONE_INT_ST_S = 6
const CACHE_L1_CACHE_PLD_DONE_INT_ST_V = 0x1
const CACHE_L1_CACHE_PLD_DONE_INT_ST_S = 4
const CACHE_SYNC_ERR_CODE = 0x00000003
const CACHE_SYNC_ERR_CODE_V = 0x3
const CACHE_SYNC_ERR_CODE_S = 12
const CACHE_L1_CACHE_PLD_ERR_CODE = 0x00000003
const CACHE_L1_CACHE_PLD_ERR_CODE_V = 0x3
const CACHE_L1_CACHE_PLD_ERR_CODE_S = 8
const CACHE_L1_CACHE_SYNC_RST_V = 0x1
const CACHE_L1_CACHE_SYNC_RST_S = 4
const CACHE_L1_CACHE_PLD_RST_V = 0x1
const CACHE_L1_CACHE_PLD_RST_S = 4
const CACHE_L1_CACHE_ALD_BUF_CLR_V = 0x1
const CACHE_L1_CACHE_ALD_BUF_CLR_S = 4
const CACHE_L1_CACHE_UNALLOC_CLR_V = 0x1
const CACHE_L1_CACHE_UNALLOC_CLR_S = 4
const CACHE_L1_CACHE_MEM_OBJECT_V = 0x1
const CACHE_L1_CACHE_MEM_OBJECT_S = 10
const CACHE_L1_CACHE_TAG_OBJECT_V = 0x1
const CACHE_L1_CACHE_TAG_OBJECT_S = 4
const CACHE_L1_CACHE_WAY_OBJECT = 0x00000007
const CACHE_L1_CACHE_WAY_OBJECT_V = 0x7
const CACHE_L1_CACHE_WAY_OBJECT_S = 0
const CACHE_L1_CACHE_VADDR = 0xFFFFFFFF
const CACHE_L1_CACHE_VADDR_V = 0xFFFFFFFF
const CACHE_L1_CACHE_VADDR_S = 0
const CACHE_L1_CACHE_DEBUG_BUS = 0xFFFFFFFF
const CACHE_L1_CACHE_DEBUG_BUS_V = 0xFFFFFFFF
const CACHE_L1_CACHE_DEBUG_BUS_S = 0
const CACHE_CLK_EN_V = 0x1
const CACHE_CLK_EN_S = 0
const CACHE_REDCY_SIG0 = 0xFFFFFFFF
const CACHE_REDCY_SIG0_V = 0xFFFFFFFF
const CACHE_REDCY_SIG0_S = 0
const CACHE_REDCY_SIG1 = 0xFFFFFFFF
const CACHE_REDCY_SIG1_V = 0xFFFFFFFF
const CACHE_REDCY_SIG1_S = 0
const CACHE_REDCY_SIG2 = 0xFFFFFFFF
const CACHE_REDCY_SIG2_V = 0xFFFFFFFF
const CACHE_REDCY_SIG2_S = 0
const CACHE_REDCY_SIG3 = 0xFFFFFFFF
const CACHE_REDCY_SIG3_V = 0xFFFFFFFF
const CACHE_REDCY_SIG3_S = 0
const CACHE_REDCY_SIG4 = 0x0000000F
const CACHE_REDCY_SIG4_V = 0xF
const CACHE_REDCY_SIG4_S = 0
const CACHE_DATE = 0x0FFFFFFF
const CACHE_DATE_V = 0xFFFFFFF
const CACHE_DATE_S = 0
