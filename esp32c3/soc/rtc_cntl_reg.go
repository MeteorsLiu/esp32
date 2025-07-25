package soc

import _ "unsafe"

const RTC_CNTL_SW_SYS_RST_V = 0x1
const RTC_CNTL_SW_SYS_RST_S = 31
const RTC_CNTL_DG_WRAP_FORCE_NORST_V = 0x1
const RTC_CNTL_DG_WRAP_FORCE_NORST_S = 30
const RTC_CNTL_DG_WRAP_FORCE_RST_V = 0x1
const RTC_CNTL_DG_WRAP_FORCE_RST_S = 29
const RTC_CNTL_ANALOG_FORCE_NOISO_V = 0x1
const RTC_CNTL_ANALOG_FORCE_NOISO_S = 28
const RTC_CNTL_PLL_FORCE_NOISO_V = 0x1
const RTC_CNTL_PLL_FORCE_NOISO_S = 27
const RTC_CNTL_XTL_FORCE_NOISO_V = 0x1
const RTC_CNTL_XTL_FORCE_NOISO_S = 26
const RTC_CNTL_ANALOG_FORCE_ISO_V = 0x1
const RTC_CNTL_ANALOG_FORCE_ISO_S = 25
const RTC_CNTL_PLL_FORCE_ISO_V = 0x1
const RTC_CNTL_PLL_FORCE_ISO_S = 24
const RTC_CNTL_XTL_FORCE_ISO_V = 0x1
const RTC_CNTL_XTL_FORCE_ISO_S = 23
const RTC_CNTL_XTL_EXT_CTR_SEL = 0x00000007
const RTC_CNTL_XTL_EXT_CTR_SEL_V = 0x7
const RTC_CNTL_XTL_EXT_CTR_SEL_S = 20
const RTC_CNTL_XTL_EN_WAIT = 0x0000000F
const RTC_CNTL_XTL_EN_WAIT_V = 0xF
const RTC_CNTL_XTL_EN_WAIT_S = 14
const RTC_CNTL_XTL_FORCE_PU_V = 0x1
const RTC_CNTL_XTL_FORCE_PU_S = 13
const RTC_CNTL_XTL_FORCE_PD_V = 0x1
const RTC_CNTL_XTL_FORCE_PD_S = 12
const RTC_CNTL_BBPLL_FORCE_PU_V = 0x1
const RTC_CNTL_BBPLL_FORCE_PU_S = 11
const RTC_CNTL_BBPLL_FORCE_PD_V = 0x1
const RTC_CNTL_BBPLL_FORCE_PD_S = 10
const RTC_CNTL_BBPLL_I2C_FORCE_PU_V = 0x1
const RTC_CNTL_BBPLL_I2C_FORCE_PU_S = 9
const RTC_CNTL_BBPLL_I2C_FORCE_PD_V = 0x1
const RTC_CNTL_BBPLL_I2C_FORCE_PD_S = 8
const RTC_CNTL_BB_I2C_FORCE_PU_V = 0x1
const RTC_CNTL_BB_I2C_FORCE_PU_S = 7
const RTC_CNTL_BB_I2C_FORCE_PD_V = 0x1
const RTC_CNTL_BB_I2C_FORCE_PD_S = 6
const RTC_CNTL_SW_PROCPU_RST_V = 0x1
const RTC_CNTL_SW_PROCPU_RST_S = 5
const RTC_CNTL_SW_APPCPU_RST_V = 0x1
const RTC_CNTL_SW_APPCPU_RST_S = 4
const RTC_CNTL_SW_STALL_PROCPU_C0 = 0x00000003
const RTC_CNTL_SW_STALL_PROCPU_C0_V = 0x3
const RTC_CNTL_SW_STALL_PROCPU_C0_S = 2
const RTC_CNTL_SW_STALL_APPCPU_C0 = 0x00000003
const RTC_CNTL_SW_STALL_APPCPU_C0_V = 0x3
const RTC_CNTL_SW_STALL_APPCPU_C0_S = 0
const RTC_CNTL_SLP_VAL_LO = 0xFFFFFFFF
const RTC_CNTL_SLP_VAL_LO_V = 0xFFFFFFFF
const RTC_CNTL_SLP_VAL_LO_S = 0
const RTC_CNTL_MAIN_TIMER_ALARM_EN_V = 0x1
const RTC_CNTL_MAIN_TIMER_ALARM_EN_S = 16
const RTC_CNTL_SLP_VAL_HI = 0x0000FFFF
const RTC_CNTL_SLP_VAL_HI_V = 0xFFFF
const RTC_CNTL_SLP_VAL_HI_S = 0
const RTC_CNTL_TIME_UPDATE_V = 0x1
const RTC_CNTL_TIME_UPDATE_S = 31
const RTC_CNTL_TIMER_SYS_RST_V = 0x1
const RTC_CNTL_TIMER_SYS_RST_S = 29
const RTC_CNTL_TIMER_XTL_OFF_V = 0x1
const RTC_CNTL_TIMER_XTL_OFF_S = 28
const RTC_CNTL_TIMER_SYS_STALL_V = 0x1
const RTC_CNTL_TIMER_SYS_STALL_S = 27
const RTC_CNTL_TIMER_VALUE0_LOW = 0xFFFFFFFF
const RTC_CNTL_TIMER_VALUE0_LOW_V = 0xFFFFFFFF
const RTC_CNTL_TIMER_VALUE0_LOW_S = 0
const RTC_CNTL_TIMER_VALUE0_HIGH = 0x0000FFFF
const RTC_CNTL_TIMER_VALUE0_HIGH_V = 0xFFFF
const RTC_CNTL_TIMER_VALUE0_HIGH_S = 0
const RTC_CNTL_SLEEP_EN_V = 0x1
const RTC_CNTL_SLEEP_EN_S = 31
const RTC_CNTL_SLP_REJECT_V = 0x1
const RTC_CNTL_SLP_REJECT_S = 30
const RTC_CNTL_SLP_WAKEUP_V = 0x1
const RTC_CNTL_SLP_WAKEUP_S = 29
const RTC_CNTL_SDIO_ACTIVE_IND_V = 0x1
const RTC_CNTL_SDIO_ACTIVE_IND_S = 28
const RTC_CNTL_APB2RTC_BRIDGE_SEL_V = 0x1
const RTC_CNTL_APB2RTC_BRIDGE_SEL_S = 22
const RTC_CNTL_SLP_REJECT_CAUSE_CLR_V = 0x1
const RTC_CNTL_SLP_REJECT_CAUSE_CLR_S = 1
const RTC_CNTL_SW_CPU_INT_V = 0x1
const RTC_CNTL_SW_CPU_INT_S = 0
const RTC_CNTL_PLL_BUF_WAIT = 0x000000FF
const RTC_CNTL_PLL_BUF_WAIT_V = 0xFF
const RTC_CNTL_PLL_BUF_WAIT_S = 24
const RTC_CNTL_PLL_BUF_WAIT_DEFAULT = 20
const RTC_CNTL_XTL_BUF_WAIT = 0x000003FF
const RTC_CNTL_XTL_BUF_WAIT_V = 0x3FF
const RTC_CNTL_XTL_BUF_WAIT_S = 14
const RTC_CNTL_XTL_BUF_WAIT_DEFAULT = 100
const RTC_CNTL_CK8M_WAIT = 0x000000FF
const RTC_CNTL_CK8M_WAIT_V = 0xFF
const RTC_CNTL_CK8M_WAIT_S = 6
const RTC_CNTL_CPU_STALL_WAIT = 0x0000001F
const RTC_CNTL_CPU_STALL_WAIT_V = 0x1F
const RTC_CNTL_CPU_STALL_WAIT_S = 1
const RTC_CNTL_CPU_STALL_EN_V = 0x1
const RTC_CNTL_CPU_STALL_EN_S = 0
const RTC_CNTL_MIN_TIME_CK8M_OFF = 0x000000FF
const RTC_CNTL_MIN_TIME_CK8M_OFF_V = 0xFF
const RTC_CNTL_MIN_TIME_CK8M_OFF_S = 24
const RTC_CNTL_BT_POWERUP_TIMER = 0x0000007F
const RTC_CNTL_BT_POWERUP_TIMER_V = 0x7F
const RTC_CNTL_BT_POWERUP_TIMER_S = 25
const RTC_CNTL_BT_WAIT_TIMER = 0x000001FF
const RTC_CNTL_BT_WAIT_TIMER_V = 0x1FF
const RTC_CNTL_BT_WAIT_TIMER_S = 16
const RTC_CNTL_WIFI_POWERUP_TIMER = 0x0000007F
const RTC_CNTL_WIFI_POWERUP_TIMER_V = 0x7F
const RTC_CNTL_WIFI_POWERUP_TIMER_S = 9
const RTC_CNTL_WIFI_WAIT_TIMER = 0x000001FF
const RTC_CNTL_WIFI_WAIT_TIMER_V = 0x1FF
const RTC_CNTL_WIFI_WAIT_TIMER_S = 0
const RTC_CNTL_DG_WRAP_POWERUP_TIMER = 0x0000007F
const RTC_CNTL_DG_WRAP_POWERUP_TIMER_V = 0x7F
const RTC_CNTL_DG_WRAP_POWERUP_TIMER_S = 25
const RTC_CNTL_DG_WRAP_WAIT_TIMER = 0x000001FF
const RTC_CNTL_DG_WRAP_WAIT_TIMER_V = 0x1FF
const RTC_CNTL_DG_WRAP_WAIT_TIMER_S = 16
const RTC_CNTL_CPU_TOP_POWERUP_TIMER = 0x0000007F
const RTC_CNTL_CPU_TOP_POWERUP_TIMER_V = 0x7F
const RTC_CNTL_CPU_TOP_POWERUP_TIMER_S = 9
const RTC_CNTL_CPU_TOP_WAIT_TIMER = 0x000001FF
const RTC_CNTL_CPU_TOP_WAIT_TIMER_V = 0x1FF
const RTC_CNTL_CPU_TOP_WAIT_TIMER_S = 0
const RTC_CNTL_MIN_SLP_VAL = 0x000000FF
const RTC_CNTL_MIN_SLP_VAL_V = 0xFF
const RTC_CNTL_MIN_SLP_VAL_S = 8
const RTC_CNTL_DG_PERI_POWERUP_TIMER = 0x0000007F
const RTC_CNTL_DG_PERI_POWERUP_TIMER_V = 0x7F
const RTC_CNTL_DG_PERI_POWERUP_TIMER_S = 25
const RTC_CNTL_DG_PERI_WAIT_TIMER = 0x000001FF
const RTC_CNTL_DG_PERI_WAIT_TIMER_V = 0x1FF
const RTC_CNTL_DG_PERI_WAIT_TIMER_S = 16
const RTC_CNTL_PLL_I2C_PU_V = 0x1
const RTC_CNTL_PLL_I2C_PU_S = 31
const RTC_CNTL_CKGEN_I2C_PU_V = 0x1
const RTC_CNTL_CKGEN_I2C_PU_S = 30
const RTC_CNTL_RFRX_PBUS_PU_V = 0x1
const RTC_CNTL_RFRX_PBUS_PU_S = 28
const RTC_CNTL_TXRF_I2C_PU_V = 0x1
const RTC_CNTL_TXRF_I2C_PU_S = 27
const RTC_CNTL_PVTMON_PU_V = 0x1
const RTC_CNTL_PVTMON_PU_S = 26
const RTC_CNTL_BBPLL_CAL_SLP_START_V = 0x1
const RTC_CNTL_BBPLL_CAL_SLP_START_S = 25
const RTC_CNTL_PLLA_FORCE_PU_V = 0x1
const RTC_CNTL_PLLA_FORCE_PU_S = 24
const RTC_CNTL_PLLA_FORCE_PD_V = 0x1
const RTC_CNTL_PLLA_FORCE_PD_S = 23
const RTC_CNTL_SAR_I2C_PU_V = 0x1
const RTC_CNTL_SAR_I2C_PU_S = 22
const RTC_CNTL_GLITCH_RST_EN_V = 0x1
const RTC_CNTL_GLITCH_RST_EN_S = 20
const RTC_CNTL_I2C_RESET_POR_FORCE_PU_V = 0x1
const RTC_CNTL_I2C_RESET_POR_FORCE_PU_S = 19
const RTC_CNTL_I2C_RESET_POR_FORCE_PD_V = 0x1
const RTC_CNTL_I2C_RESET_POR_FORCE_PD_S = 18
const RTC_CNTL_DRESET_MASK_PROCPU_V = 0x1
const RTC_CNTL_DRESET_MASK_PROCPU_S = 25
const RTC_CNTL_DRESET_MASK_APPCPU_V = 0x1
const RTC_CNTL_DRESET_MASK_APPCPU_S = 24
const RTC_CNTL_JTAG_RESET_FLAG_CLR_APPCPU_V = 0x1
const RTC_CNTL_JTAG_RESET_FLAG_CLR_APPCPU_S = 23
const RTC_CNTL_JTAG_RESET_FLAG_CLR_PROCPU_V = 0x1
const RTC_CNTL_JTAG_RESET_FLAG_CLR_PROCPU_S = 22
const RTC_CNTL_JTAG_RESET_FLAG_APPCPU_V = 0x1
const RTC_CNTL_JTAG_RESET_FLAG_APPCPU_S = 21
const RTC_CNTL_JTAG_RESET_FLAG_PROCPU_V = 0x1
const RTC_CNTL_JTAG_RESET_FLAG_PROCPU_S = 20
const RTC_CNTL_OCD_HALT_ON_RESET_PROCPU_V = 0x1
const RTC_CNTL_OCD_HALT_ON_RESET_PROCPU_S = 19
const RTC_CNTL_OCD_HALT_ON_RESET_APPCPU_V = 0x1
const RTC_CNTL_OCD_HALT_ON_RESET_APPCPU_S = 18
const RTC_CNTL_ALL_RESET_FLAG_CLR_APPCPU_V = 0x1
const RTC_CNTL_ALL_RESET_FLAG_CLR_APPCPU_S = 17
const RTC_CNTL_ALL_RESET_FLAG_CLR_PROCPU_V = 0x1
const RTC_CNTL_ALL_RESET_FLAG_CLR_PROCPU_S = 16
const RTC_CNTL_ALL_RESET_FLAG_APPCPU_V = 0x1
const RTC_CNTL_ALL_RESET_FLAG_APPCPU_S = 15
const RTC_CNTL_ALL_RESET_FLAG_PROCPU_V = 0x1
const RTC_CNTL_ALL_RESET_FLAG_PROCPU_S = 14
const RTC_CNTL_STAT_VECTOR_SEL_PROCPU_V = 0x1
const RTC_CNTL_STAT_VECTOR_SEL_PROCPU_S = 13
const RTC_CNTL_STAT_VECTOR_SEL_APPCPU_V = 0x1
const RTC_CNTL_STAT_VECTOR_SEL_APPCPU_S = 12
const RTC_CNTL_RESET_CAUSE_APPCPU = 0x0000003F
const RTC_CNTL_RESET_CAUSE_APPCPU_V = 0x3F
const RTC_CNTL_RESET_CAUSE_APPCPU_S = 6
const RTC_CNTL_RESET_CAUSE_PROCPU = 0x0000003F
const RTC_CNTL_RESET_CAUSE_PROCPU_V = 0x3F
const RTC_CNTL_RESET_CAUSE_PROCPU_S = 0
const RTC_CNTL_WAKEUP_ENA = 0x0001FFFF
const RTC_CNTL_WAKEUP_ENA_V = 0x1FFFF
const RTC_CNTL_WAKEUP_ENA_S = 15
const RTC_CNTL_BBPLL_CAL_INT_ENA_V = 0x1
const RTC_CNTL_BBPLL_CAL_INT_ENA_S = 20
const RTC_CNTL_GLITCH_DET_INT_ENA_V = 0x1
const RTC_CNTL_GLITCH_DET_INT_ENA_S = 19
const RTC_CNTL_XTAL32K_DEAD_INT_ENA_V = 0x1
const RTC_CNTL_XTAL32K_DEAD_INT_ENA_S = 16
const RTC_CNTL_SWD_INT_ENA_V = 0x1
const RTC_CNTL_SWD_INT_ENA_S = 15
const RTC_CNTL_MAIN_TIMER_INT_ENA_V = 0x1
const RTC_CNTL_MAIN_TIMER_INT_ENA_S = 10
const RTC_CNTL_BROWN_OUT_INT_ENA_V = 0x1
const RTC_CNTL_BROWN_OUT_INT_ENA_S = 9
const RTC_CNTL_WDT_INT_ENA_V = 0x1
const RTC_CNTL_WDT_INT_ENA_S = 3
const RTC_CNTL_SLP_REJECT_INT_ENA_V = 0x1
const RTC_CNTL_SLP_REJECT_INT_ENA_S = 1
const RTC_CNTL_SLP_WAKEUP_INT_ENA_V = 0x1
const RTC_CNTL_SLP_WAKEUP_INT_ENA_S = 0
const RTC_CNTL_BBPLL_CAL_INT_RAW_V = 0x1
const RTC_CNTL_BBPLL_CAL_INT_RAW_S = 20
const RTC_CNTL_GLITCH_DET_INT_RAW_V = 0x1
const RTC_CNTL_GLITCH_DET_INT_RAW_S = 19
const RTC_CNTL_XTAL32K_DEAD_INT_RAW_V = 0x1
const RTC_CNTL_XTAL32K_DEAD_INT_RAW_S = 16
const RTC_CNTL_SWD_INT_RAW_V = 0x1
const RTC_CNTL_SWD_INT_RAW_S = 15
const RTC_CNTL_MAIN_TIMER_INT_RAW_V = 0x1
const RTC_CNTL_MAIN_TIMER_INT_RAW_S = 10
const RTC_CNTL_BROWN_OUT_INT_RAW_V = 0x1
const RTC_CNTL_BROWN_OUT_INT_RAW_S = 9
const RTC_CNTL_WDT_INT_RAW_V = 0x1
const RTC_CNTL_WDT_INT_RAW_S = 3
const RTC_CNTL_SLP_REJECT_INT_RAW_V = 0x1
const RTC_CNTL_SLP_REJECT_INT_RAW_S = 1
const RTC_CNTL_SLP_WAKEUP_INT_RAW_V = 0x1
const RTC_CNTL_SLP_WAKEUP_INT_RAW_S = 0
const RTC_CNTL_BBPLL_CAL_INT_ST_V = 0x1
const RTC_CNTL_BBPLL_CAL_INT_ST_S = 20
const RTC_CNTL_GLITCH_DET_INT_ST_V = 0x1
const RTC_CNTL_GLITCH_DET_INT_ST_S = 19
const RTC_CNTL_XTAL32K_DEAD_INT_ST_V = 0x1
const RTC_CNTL_XTAL32K_DEAD_INT_ST_S = 16
const RTC_CNTL_SWD_INT_ST_V = 0x1
const RTC_CNTL_SWD_INT_ST_S = 15
const RTC_CNTL_MAIN_TIMER_INT_ST_V = 0x1
const RTC_CNTL_MAIN_TIMER_INT_ST_S = 10
const RTC_CNTL_BROWN_OUT_INT_ST_V = 0x1
const RTC_CNTL_BROWN_OUT_INT_ST_S = 9
const RTC_CNTL_WDT_INT_ST_V = 0x1
const RTC_CNTL_WDT_INT_ST_S = 3
const RTC_CNTL_SLP_REJECT_INT_ST_V = 0x1
const RTC_CNTL_SLP_REJECT_INT_ST_S = 1
const RTC_CNTL_SLP_WAKEUP_INT_ST_V = 0x1
const RTC_CNTL_SLP_WAKEUP_INT_ST_S = 0
const RTC_CNTL_BBPLL_CAL_INT_CLR_V = 0x1
const RTC_CNTL_BBPLL_CAL_INT_CLR_S = 20
const RTC_CNTL_GLITCH_DET_INT_CLR_V = 0x1
const RTC_CNTL_GLITCH_DET_INT_CLR_S = 19
const RTC_CNTL_XTAL32K_DEAD_INT_CLR_V = 0x1
const RTC_CNTL_XTAL32K_DEAD_INT_CLR_S = 16
const RTC_CNTL_SWD_INT_CLR_V = 0x1
const RTC_CNTL_SWD_INT_CLR_S = 15
const RTC_CNTL_MAIN_TIMER_INT_CLR_V = 0x1
const RTC_CNTL_MAIN_TIMER_INT_CLR_S = 10
const RTC_CNTL_BROWN_OUT_INT_CLR_V = 0x1
const RTC_CNTL_BROWN_OUT_INT_CLR_S = 9
const RTC_CNTL_WDT_INT_CLR_V = 0x1
const RTC_CNTL_WDT_INT_CLR_S = 3
const RTC_CNTL_SLP_REJECT_INT_CLR_V = 0x1
const RTC_CNTL_SLP_REJECT_INT_CLR_S = 1
const RTC_CNTL_SLP_WAKEUP_INT_CLR_V = 0x1
const RTC_CNTL_SLP_WAKEUP_INT_CLR_S = 0
const RTC_CNTL_SCRATCH0 = 0xFFFFFFFF
const RTC_CNTL_SCRATCH0_V = 0xFFFFFFFF
const RTC_CNTL_SCRATCH0_S = 0
const RTC_CNTL_SCRATCH1 = 0xFFFFFFFF
const RTC_CNTL_SCRATCH1_V = 0xFFFFFFFF
const RTC_CNTL_SCRATCH1_S = 0
const RTC_CNTL_SCRATCH2 = 0xFFFFFFFF
const RTC_CNTL_SCRATCH2_V = 0xFFFFFFFF
const RTC_CNTL_SCRATCH2_S = 0
const RTC_CNTL_SCRATCH3 = 0xFFFFFFFF
const RTC_CNTL_SCRATCH3_V = 0xFFFFFFFF
const RTC_CNTL_SCRATCH3_S = 0
const RTC_CNTL_XTL_EXT_CTR_EN_V = 0x1
const RTC_CNTL_XTL_EXT_CTR_EN_S = 31
const RTC_CNTL_XTL_EXT_CTR_LV_V = 0x1
const RTC_CNTL_XTL_EXT_CTR_LV_S = 30
const RTC_CNTL_XTAL32K_GPIO_SEL_V = 0x1
const RTC_CNTL_XTAL32K_GPIO_SEL_S = 23
const RTC_CNTL_WDT_STATE = 0x00000007
const RTC_CNTL_WDT_STATE_V = 0x7
const RTC_CNTL_WDT_STATE_S = 20
const RTC_CNTL_DAC_XTAL_32K = 0x00000007
const RTC_CNTL_DAC_XTAL_32K_V = 0x7
const RTC_CNTL_DAC_XTAL_32K_S = 17
const RTC_CNTL_XPD_XTAL_32K_V = 0x1
const RTC_CNTL_XPD_XTAL_32K_S = 16
const RTC_CNTL_DRES_XTAL_32K = 0x00000007
const RTC_CNTL_DRES_XTAL_32K_V = 0x7
const RTC_CNTL_DRES_XTAL_32K_S = 13
const RTC_CNTL_DGM_XTAL_32K = 0x00000007
const RTC_CNTL_DGM_XTAL_32K_V = 0x7
const RTC_CNTL_DGM_XTAL_32K_S = 10
const RTC_CNTL_DBUF_XTAL_32K_V = 0x1
const RTC_CNTL_DBUF_XTAL_32K_S = 9
const RTC_CNTL_ENCKINIT_XTAL_32K_V = 0x1
const RTC_CNTL_ENCKINIT_XTAL_32K_S = 8
const RTC_CNTL_XTAL32K_XPD_FORCE_V = 0x1
const RTC_CNTL_XTAL32K_XPD_FORCE_S = 7
const RTC_CNTL_XTAL32K_AUTO_RETURN_V = 0x1
const RTC_CNTL_XTAL32K_AUTO_RETURN_S = 6
const RTC_CNTL_XTAL32K_AUTO_RESTART_V = 0x1
const RTC_CNTL_XTAL32K_AUTO_RESTART_S = 5
const RTC_CNTL_XTAL32K_AUTO_BACKUP_V = 0x1
const RTC_CNTL_XTAL32K_AUTO_BACKUP_S = 4
const RTC_CNTL_XTAL32K_EXT_CLK_FO_V = 0x1
const RTC_CNTL_XTAL32K_EXT_CLK_FO_S = 3
const RTC_CNTL_XTAL32K_WDT_RESET_V = 0x1
const RTC_CNTL_XTAL32K_WDT_RESET_S = 2
const RTC_CNTL_XTAL32K_WDT_CLK_FO_V = 0x1
const RTC_CNTL_XTAL32K_WDT_CLK_FO_S = 1
const RTC_CNTL_XTAL32K_WDT_EN_V = 0x1
const RTC_CNTL_XTAL32K_WDT_EN_S = 0
const RTC_CNTL_GPIO_WAKEUP_FILTER_V = 0x1
const RTC_CNTL_GPIO_WAKEUP_FILTER_S = 31
const RTC_CNTL_DEEP_SLP_REJECT_EN_V = 0x1
const RTC_CNTL_DEEP_SLP_REJECT_EN_S = 31
const RTC_CNTL_LIGHT_SLP_REJECT_EN_V = 0x1
const RTC_CNTL_LIGHT_SLP_REJECT_EN_S = 30
const RTC_CNTL_SLEEP_REJECT_ENA = 0x0003FFFF
const RTC_CNTL_SLEEP_REJECT_ENA_V = 0x3FFFF
const RTC_CNTL_SLEEP_REJECT_ENA_S = 12
const RTC_CNTL_CPUPERIOD_SEL = 0x00000003
const RTC_CNTL_CPUPERIOD_SEL_V = 0x3
const RTC_CNTL_CPUPERIOD_SEL_S = 30
const RTC_CNTL_CPUSEL_CONF_V = 0x1
const RTC_CNTL_CPUSEL_CONF_S = 29
const RTC_CNTL_ANA_CLK_RTC_SEL = 0x00000003
const RTC_CNTL_ANA_CLK_RTC_SEL_V = 0x3
const RTC_CNTL_ANA_CLK_RTC_SEL_S = 30
const RTC_CNTL_FAST_CLK_RTC_SEL_V = 0x1
const RTC_CNTL_FAST_CLK_RTC_SEL_S = 29
const RTC_CNTL_XTAL_GLOBAL_FORCE_NOGATING_V = 0x1
const RTC_CNTL_XTAL_GLOBAL_FORCE_NOGATING_S = 28
const RTC_CNTL_XTAL_GLOBAL_FORCE_GATING_V = 0x1
const RTC_CNTL_XTAL_GLOBAL_FORCE_GATING_S = 27
const RTC_CNTL_CK8M_FORCE_PU_V = 0x1
const RTC_CNTL_CK8M_FORCE_PU_S = 26
const RTC_CNTL_CK8M_FORCE_PD_V = 0x1
const RTC_CNTL_CK8M_FORCE_PD_S = 25
const RTC_CNTL_CK8M_DFREQ = 0x000000FF
const RTC_CNTL_CK8M_DFREQ_V = 0xFF
const RTC_CNTL_CK8M_DFREQ_S = 17
const RTC_CNTL_CK8M_FORCE_NOGATING_V = 0x1
const RTC_CNTL_CK8M_FORCE_NOGATING_S = 16
const RTC_CNTL_XTAL_FORCE_NOGATING_V = 0x1
const RTC_CNTL_XTAL_FORCE_NOGATING_S = 15
const RTC_CNTL_CK8M_DIV_SEL = 0x00000007
const RTC_CNTL_CK8M_DIV_SEL_V = 0x7
const RTC_CNTL_CK8M_DIV_SEL_S = 12
const RTC_CNTL_DIG_CLK8M_EN_V = 0x1
const RTC_CNTL_DIG_CLK8M_EN_S = 10
const RTC_CNTL_DIG_CLK8M_D256_EN_V = 0x1
const RTC_CNTL_DIG_CLK8M_D256_EN_S = 9
const RTC_CNTL_DIG_XTAL32K_EN_V = 0x1
const RTC_CNTL_DIG_XTAL32K_EN_S = 8
const RTC_CNTL_ENB_CK8M_DIV_V = 0x1
const RTC_CNTL_ENB_CK8M_DIV_S = 7
const RTC_CNTL_ENB_CK8M_V = 0x1
const RTC_CNTL_ENB_CK8M_S = 6
const RTC_CNTL_CK8M_DIV = 0x00000003
const RTC_CNTL_CK8M_DIV_V = 0x3
const RTC_CNTL_CK8M_DIV_S = 4
const RTC_CNTL_CK8M_DIV_SEL_VLD_V = 0x1
const RTC_CNTL_CK8M_DIV_SEL_VLD_S = 3
const RTC_CNTL_EFUSE_CLK_FORCE_NOGATING_V = 0x1
const RTC_CNTL_EFUSE_CLK_FORCE_NOGATING_S = 2
const RTC_CNTL_EFUSE_CLK_FORCE_GATING_V = 0x1
const RTC_CNTL_EFUSE_CLK_FORCE_GATING_S = 1
const RTC_CNTL_SLOW_CLK_NEXT_EDGE_V = 0x1
const RTC_CNTL_SLOW_CLK_NEXT_EDGE_S = 31
const RTC_CNTL_ANA_CLK_DIV = 0x000000FF
const RTC_CNTL_ANA_CLK_DIV_V = 0xFF
const RTC_CNTL_ANA_CLK_DIV_S = 23
const RTC_CNTL_ANA_CLK_DIV_VLD_V = 0x1
const RTC_CNTL_ANA_CLK_DIV_VLD_S = 22
const RTC_CNTL_XPD_SDIO_REG_V = 0x1
const RTC_CNTL_XPD_SDIO_REG_S = 31
const RTC_CNTL_DREFH_SDIO = 0x00000003
const RTC_CNTL_DREFH_SDIO_V = 0x3
const RTC_CNTL_DREFH_SDIO_S = 29
const RTC_CNTL_DREFM_SDIO = 0x00000003
const RTC_CNTL_DREFM_SDIO_V = 0x3
const RTC_CNTL_DREFM_SDIO_S = 27
const RTC_CNTL_DREFL_SDIO = 0x00000003
const RTC_CNTL_DREFL_SDIO_V = 0x3
const RTC_CNTL_DREFL_SDIO_S = 25
const RTC_CNTL_REG1P8_READY_V = 0x1
const RTC_CNTL_REG1P8_READY_S = 24
const RTC_CNTL_SDIO_TIEH_V = 0x1
const RTC_CNTL_SDIO_TIEH_S = 23
const RTC_CNTL_SDIO_FORCE_V = 0x1
const RTC_CNTL_SDIO_FORCE_S = 22
const RTC_CNTL_SDIO_PD_EN_V = 0x1
const RTC_CNTL_SDIO_PD_EN_S = 21
const RTC_CNTL_SDIO_ENCURLIM_V = 0x1
const RTC_CNTL_SDIO_ENCURLIM_S = 20
const RTC_CNTL_SDIO_MODECURLIM_V = 0x1
const RTC_CNTL_SDIO_MODECURLIM_S = 19
const RTC_CNTL_SDIO_DCURLIM = 0x00000007
const RTC_CNTL_SDIO_DCURLIM_V = 0x7
const RTC_CNTL_SDIO_DCURLIM_S = 16
const RTC_CNTL_SDIO_EN_INITI_V = 0x1
const RTC_CNTL_SDIO_EN_INITI_S = 15
const RTC_CNTL_SDIO_INITI = 0x00000003
const RTC_CNTL_SDIO_INITI_V = 0x3
const RTC_CNTL_SDIO_INITI_S = 13
const RTC_CNTL_SDIO_DCAP = 0x00000003
const RTC_CNTL_SDIO_DCAP_V = 0x3
const RTC_CNTL_SDIO_DCAP_S = 11
const RTC_CNTL_SDIO_DTHDRV = 0x00000003
const RTC_CNTL_SDIO_DTHDRV_V = 0x3
const RTC_CNTL_SDIO_DTHDRV_S = 9
const RTC_CNTL_SDIO_TIMER_TARGET = 0x000000FF
const RTC_CNTL_SDIO_TIMER_TARGET_V = 0xFF
const RTC_CNTL_SDIO_TIMER_TARGET_S = 0
const RTC_CNTL_DBG_ATTEN_MONITOR = 0x0000000F
const RTC_CNTL_DBG_ATTEN_MONITOR_V = 0xF
const RTC_CNTL_DBG_ATTEN_MONITOR_S = 22
const RTC_CNTL_DBG_ATTEN_DEEP_SLP = 0x0000000F
const RTC_CNTL_DBG_ATTEN_DEEP_SLP_V = 0xF
const RTC_CNTL_DBG_ATTEN_DEEP_SLP_S = 18
const RTC_CNTL_BIAS_SLEEP_MONITOR_V = 0x1
const RTC_CNTL_BIAS_SLEEP_MONITOR_S = 17
const RTC_CNTL_BIAS_SLEEP_DEEP_SLP_V = 0x1
const RTC_CNTL_BIAS_SLEEP_DEEP_SLP_S = 16
const RTC_CNTL_PD_CUR_MONITOR_V = 0x1
const RTC_CNTL_PD_CUR_MONITOR_S = 15
const RTC_CNTL_PD_CUR_DEEP_SLP_V = 0x1
const RTC_CNTL_PD_CUR_DEEP_SLP_S = 14
const RTC_CNTL_BIAS_BUF_MONITOR_V = 0x1
const RTC_CNTL_BIAS_BUF_MONITOR_S = 13
const RTC_CNTL_BIAS_BUF_DEEP_SLP_V = 0x1
const RTC_CNTL_BIAS_BUF_DEEP_SLP_S = 12
const RTC_CNTL_BIAS_BUF_WAKE_V = 0x1
const RTC_CNTL_BIAS_BUF_WAKE_S = 11
const RTC_CNTL_BIAS_BUF_IDLE_V = 0x1
const RTC_CNTL_BIAS_BUF_IDLE_S = 10
const RTC_CNTL_DG_VDD_DRV_B_SLP_EN_V = 0x1
const RTC_CNTL_DG_VDD_DRV_B_SLP_EN_S = 8
const RTC_CNTL_DG_VDD_DRV_B_SLP = 0x000000FF
const RTC_CNTL_DG_VDD_DRV_B_SLP_V = 0xFF
const RTC_CNTL_DG_VDD_DRV_B_SLP_S = 0
const RTC_CNTL_REGULATOR_FORCE_PU_V = 0x1
const RTC_CNTL_REGULATOR_FORCE_PU_S = 31
const RTC_CNTL_REGULATOR_FORCE_PD_V = 0x1
const RTC_CNTL_REGULATOR_FORCE_PD_S = 30
const RTC_CNTL_DBOOST_FORCE_PU_V = 0x1
const RTC_CNTL_DBOOST_FORCE_PU_S = 29
const RTC_CNTL_DBOOST_FORCE_PD_V = 0x1
const RTC_CNTL_DBOOST_FORCE_PD_S = 28
const RTC_CNTL_SCK_DCAP = 0x000000FF
const RTC_CNTL_SCK_DCAP_V = 0xFF
const RTC_CNTL_SCK_DCAP_S = 14
const RTC_CNTL_SCK_DCAP_DEFAULT = 255
const RTC_CNTL_DIG_CAL_EN_V = 0x1
const RTC_CNTL_DIG_CAL_EN_S = 7
const RTC_CNTL_PAD_FORCE_HOLD_V = 0x1
const RTC_CNTL_PAD_FORCE_HOLD_S = 21
const RTC_CNTL_DG_WRAP_PD_EN_V = 0x1
const RTC_CNTL_DG_WRAP_PD_EN_S = 31
const RTC_CNTL_WIFI_PD_EN_V = 0x1
const RTC_CNTL_WIFI_PD_EN_S = 30
const RTC_CNTL_CPU_TOP_PD_EN_V = 0x1
const RTC_CNTL_CPU_TOP_PD_EN_S = 29
const RTC_CNTL_DG_PERI_PD_EN_V = 0x1
const RTC_CNTL_DG_PERI_PD_EN_S = 28
const RTC_CNTL_BT_PD_EN_V = 0x1
const RTC_CNTL_BT_PD_EN_S = 27
const RTC_CNTL_CPU_TOP_FORCE_PU_V = 0x1
const RTC_CNTL_CPU_TOP_FORCE_PU_S = 22
const RTC_CNTL_CPU_TOP_FORCE_PD_V = 0x1
const RTC_CNTL_CPU_TOP_FORCE_PD_S = 21
const RTC_CNTL_DG_WRAP_FORCE_PU_V = 0x1
const RTC_CNTL_DG_WRAP_FORCE_PU_S = 20
const RTC_CNTL_DG_WRAP_FORCE_PD_V = 0x1
const RTC_CNTL_DG_WRAP_FORCE_PD_S = 19
const RTC_CNTL_WIFI_FORCE_PU_V = 0x1
const RTC_CNTL_WIFI_FORCE_PU_S = 18
const RTC_CNTL_WIFI_FORCE_PD_V = 0x1
const RTC_CNTL_WIFI_FORCE_PD_S = 17
const RTC_CNTL_FASTMEM_FORCE_LPU_V = 0x1
const RTC_CNTL_FASTMEM_FORCE_LPU_S = 16
const RTC_CNTL_FASTMEM_FORCE_LPD_V = 0x1
const RTC_CNTL_FASTMEM_FORCE_LPD_S = 15
const RTC_CNTL_DG_PERI_FORCE_PU_V = 0x1
const RTC_CNTL_DG_PERI_FORCE_PU_S = 14
const RTC_CNTL_DG_PERI_FORCE_PD_V = 0x1
const RTC_CNTL_DG_PERI_FORCE_PD_S = 13
const RTC_CNTL_BT_FORCE_PU_V = 0x1
const RTC_CNTL_BT_FORCE_PU_S = 12
const RTC_CNTL_BT_FORCE_PD_V = 0x1
const RTC_CNTL_BT_FORCE_PD_S = 11
const RTC_CNTL_LSLP_MEM_FORCE_PU_V = 0x1
const RTC_CNTL_LSLP_MEM_FORCE_PU_S = 4
const RTC_CNTL_LSLP_MEM_FORCE_PD_V = 0x1
const RTC_CNTL_LSLP_MEM_FORCE_PD_S = 3
const RTC_CNTL_VDD_SPI_PWR_FORCE_V = 0x1
const RTC_CNTL_VDD_SPI_PWR_FORCE_S = 2
const RTC_CNTL_VDD_SPI_PWR_DRV = 0x00000003
const RTC_CNTL_VDD_SPI_PWR_DRV_V = 0x3
const RTC_CNTL_VDD_SPI_PWR_DRV_S = 0
const RTC_CNTL_DG_WRAP_FORCE_NOISO_V = 0x1
const RTC_CNTL_DG_WRAP_FORCE_NOISO_S = 31
const RTC_CNTL_DG_WRAP_FORCE_ISO_V = 0x1
const RTC_CNTL_DG_WRAP_FORCE_ISO_S = 30
const RTC_CNTL_WIFI_FORCE_NOISO_V = 0x1
const RTC_CNTL_WIFI_FORCE_NOISO_S = 29
const RTC_CNTL_WIFI_FORCE_ISO_V = 0x1
const RTC_CNTL_WIFI_FORCE_ISO_S = 28
const RTC_CNTL_CPU_TOP_FORCE_NOISO_V = 0x1
const RTC_CNTL_CPU_TOP_FORCE_NOISO_S = 27
const RTC_CNTL_CPU_TOP_FORCE_ISO_V = 0x1
const RTC_CNTL_CPU_TOP_FORCE_ISO_S = 26
const RTC_CNTL_DG_PERI_FORCE_NOISO_V = 0x1
const RTC_CNTL_DG_PERI_FORCE_NOISO_S = 25
const RTC_CNTL_DG_PERI_FORCE_ISO_V = 0x1
const RTC_CNTL_DG_PERI_FORCE_ISO_S = 24
const RTC_CNTL_BT_FORCE_NOISO_V = 0x1
const RTC_CNTL_BT_FORCE_NOISO_S = 23
const RTC_CNTL_BT_FORCE_ISO_V = 0x1
const RTC_CNTL_BT_FORCE_ISO_S = 22
const RTC_CNTL_DG_PAD_FORCE_HOLD_V = 0x1
const RTC_CNTL_DG_PAD_FORCE_HOLD_S = 15
const RTC_CNTL_DG_PAD_FORCE_UNHOLD_V = 0x1
const RTC_CNTL_DG_PAD_FORCE_UNHOLD_S = 14
const RTC_CNTL_DG_PAD_FORCE_ISO_V = 0x1
const RTC_CNTL_DG_PAD_FORCE_ISO_S = 13
const RTC_CNTL_DG_PAD_FORCE_NOISO_V = 0x1
const RTC_CNTL_DG_PAD_FORCE_NOISO_S = 12
const RTC_CNTL_DG_PAD_AUTOHOLD_EN_V = 0x1
const RTC_CNTL_DG_PAD_AUTOHOLD_EN_S = 11
const RTC_CNTL_CLR_DG_PAD_AUTOHOLD_V = 0x1
const RTC_CNTL_CLR_DG_PAD_AUTOHOLD_S = 10
const RTC_CNTL_DG_PAD_AUTOHOLD_V = 0x1
const RTC_CNTL_DG_PAD_AUTOHOLD_S = 9
const RTC_CNTL_DIG_ISO_FORCE_ON_V = 0x1
const RTC_CNTL_DIG_ISO_FORCE_ON_S = 8
const RTC_CNTL_DIG_ISO_FORCE_OFF_V = 0x1
const RTC_CNTL_DIG_ISO_FORCE_OFF_S = 7
const RTC_CNTL_WDT_EN_V = 0x1
const RTC_CNTL_WDT_EN_S = 31
const RTC_CNTL_WDT_STG0 = 0x00000007
const RTC_CNTL_WDT_STG0_V = 0x7
const RTC_CNTL_WDT_STG0_S = 28
const RTC_CNTL_WDT_STG1 = 0x00000007
const RTC_CNTL_WDT_STG1_V = 0x7
const RTC_CNTL_WDT_STG1_S = 25
const RTC_CNTL_WDT_STG2 = 0x00000007
const RTC_CNTL_WDT_STG2_V = 0x7
const RTC_CNTL_WDT_STG2_S = 22
const RTC_CNTL_WDT_STG3 = 0x00000007
const RTC_CNTL_WDT_STG3_V = 0x7
const RTC_CNTL_WDT_STG3_S = 19
const RTC_WDT_STG_SEL_OFF = 0
const RTC_WDT_STG_SEL_INT = 1
const RTC_WDT_STG_SEL_RESET_CPU = 2
const RTC_WDT_STG_SEL_RESET_SYSTEM = 3
const RTC_WDT_STG_SEL_RESET_RTC = 4
const RTC_CNTL_WDT_CPU_RESET_LENGTH = 0x00000007
const RTC_CNTL_WDT_CPU_RESET_LENGTH_V = 0x7
const RTC_CNTL_WDT_CPU_RESET_LENGTH_S = 16
const RTC_CNTL_WDT_SYS_RESET_LENGTH = 0x00000007
const RTC_CNTL_WDT_SYS_RESET_LENGTH_V = 0x7
const RTC_CNTL_WDT_SYS_RESET_LENGTH_S = 13
const RTC_CNTL_WDT_FLASHBOOT_MOD_EN_V = 0x1
const RTC_CNTL_WDT_FLASHBOOT_MOD_EN_S = 12
const RTC_CNTL_WDT_PROCPU_RESET_EN_V = 0x1
const RTC_CNTL_WDT_PROCPU_RESET_EN_S = 11
const RTC_CNTL_WDT_APPCPU_RESET_EN_V = 0x1
const RTC_CNTL_WDT_APPCPU_RESET_EN_S = 10
const RTC_CNTL_WDT_PAUSE_IN_SLP_V = 0x1
const RTC_CNTL_WDT_PAUSE_IN_SLP_S = 9
const RTC_CNTL_WDT_CHIP_RESET_EN_V = 0x1
const RTC_CNTL_WDT_CHIP_RESET_EN_S = 8
const RTC_CNTL_WDT_CHIP_RESET_WIDTH = 0x000000FF
const RTC_CNTL_WDT_CHIP_RESET_WIDTH_V = 0xFF
const RTC_CNTL_WDT_CHIP_RESET_WIDTH_S = 0
const RTC_CNTL_WDT_STG0_HOLD = 0xFFFFFFFF
const RTC_CNTL_WDT_STG0_HOLD_V = 0xFFFFFFFF
const RTC_CNTL_WDT_STG0_HOLD_S = 0
const RTC_CNTL_WDT_STG1_HOLD = 0xFFFFFFFF
const RTC_CNTL_WDT_STG1_HOLD_V = 0xFFFFFFFF
const RTC_CNTL_WDT_STG1_HOLD_S = 0
const RTC_CNTL_WDT_STG2_HOLD = 0xFFFFFFFF
const RTC_CNTL_WDT_STG2_HOLD_V = 0xFFFFFFFF
const RTC_CNTL_WDT_STG2_HOLD_S = 0
const RTC_CNTL_WDT_STG3_HOLD = 0xFFFFFFFF
const RTC_CNTL_WDT_STG3_HOLD_V = 0xFFFFFFFF
const RTC_CNTL_WDT_STG3_HOLD_S = 0
const RTC_CNTL_WDT_FEED_V = 0x1
const RTC_CNTL_WDT_FEED_S = 31
const RTC_CNTL_WDT_WKEY = 0xFFFFFFFF
const RTC_CNTL_WDT_WKEY_V = 0xFFFFFFFF
const RTC_CNTL_WDT_WKEY_S = 0
const RTC_CNTL_SWD_AUTO_FEED_EN_V = 0x1
const RTC_CNTL_SWD_AUTO_FEED_EN_S = 31
const RTC_CNTL_SWD_DISABLE_V = 0x1
const RTC_CNTL_SWD_DISABLE_S = 30
const RTC_CNTL_SWD_FEED_V = 0x1
const RTC_CNTL_SWD_FEED_S = 29
const RTC_CNTL_SWD_RST_FLAG_CLR_V = 0x1
const RTC_CNTL_SWD_RST_FLAG_CLR_S = 28
const RTC_CNTL_SWD_SIGNAL_WIDTH = 0x000003FF
const RTC_CNTL_SWD_SIGNAL_WIDTH_V = 0x3FF
const RTC_CNTL_SWD_SIGNAL_WIDTH_S = 18
const RTC_CNTL_SWD_BYPASS_RST_V = 0x1
const RTC_CNTL_SWD_BYPASS_RST_S = 17
const RTC_CNTL_SWD_FEED_INT_V = 0x1
const RTC_CNTL_SWD_FEED_INT_S = 1
const RTC_CNTL_SWD_RESET_FLAG_V = 0x1
const RTC_CNTL_SWD_RESET_FLAG_S = 0
const RTC_CNTL_SWD_WKEY = 0xFFFFFFFF
const RTC_CNTL_SWD_WKEY_V = 0xFFFFFFFF
const RTC_CNTL_SWD_WKEY_S = 0
const RTC_CNTL_SW_STALL_PROCPU_C1 = 0x0000003F
const RTC_CNTL_SW_STALL_PROCPU_C1_V = 0x3F
const RTC_CNTL_SW_STALL_PROCPU_C1_S = 26
const RTC_CNTL_SW_STALL_APPCPU_C1 = 0x0000003F
const RTC_CNTL_SW_STALL_APPCPU_C1_V = 0x3F
const RTC_CNTL_SW_STALL_APPCPU_C1_S = 20
const RTC_CNTL_SCRATCH4 = 0xFFFFFFFF
const RTC_CNTL_SCRATCH4_V = 0xFFFFFFFF
const RTC_CNTL_SCRATCH4_S = 0
const RTC_CNTL_SCRATCH5 = 0xFFFFFFFF
const RTC_CNTL_SCRATCH5_V = 0xFFFFFFFF
const RTC_CNTL_SCRATCH5_S = 0
const RTC_CNTL_SCRATCH6 = 0xFFFFFFFF
const RTC_CNTL_SCRATCH6_V = 0xFFFFFFFF
const RTC_CNTL_SCRATCH6_S = 0
const RTC_CNTL_SCRATCH7 = 0xFFFFFFFF
const RTC_CNTL_SCRATCH7_V = 0xFFFFFFFF
const RTC_CNTL_SCRATCH7_S = 0
const RTC_CNTL_MAIN_STATE = 0x0000000F
const RTC_CNTL_MAIN_STATE_V = 0xF
const RTC_CNTL_MAIN_STATE_S = 28
const RTC_CNTL_MAIN_STATE_IN_IDLE_V = 0x1
const RTC_CNTL_MAIN_STATE_IN_IDLE_S = 27
const RTC_CNTL_MAIN_STATE_IN_SLP_V = 0x1
const RTC_CNTL_MAIN_STATE_IN_SLP_S = 26
const RTC_CNTL_MAIN_STATE_IN_WAIT_XTL_V = 0x1
const RTC_CNTL_MAIN_STATE_IN_WAIT_XTL_S = 25
const RTC_CNTL_MAIN_STATE_IN_WAIT_PLL_V = 0x1
const RTC_CNTL_MAIN_STATE_IN_WAIT_PLL_S = 24
const RTC_CNTL_MAIN_STATE_IN_WAIT_8M_V = 0x1
const RTC_CNTL_MAIN_STATE_IN_WAIT_8M_S = 23
const RTC_CNTL_IN_LOW_POWER_STATE_V = 0x1
const RTC_CNTL_IN_LOW_POWER_STATE_S = 22
const RTC_CNTL_IN_WAKEUP_STATE_V = 0x1
const RTC_CNTL_IN_WAKEUP_STATE_S = 21
const RTC_CNTL_MAIN_STATE_WAIT_END_V = 0x1
const RTC_CNTL_MAIN_STATE_WAIT_END_S = 20
const RTC_CNTL_RDY_FOR_WAKEUP_V = 0x1
const RTC_CNTL_RDY_FOR_WAKEUP_S = 19
const RTC_CNTL_MAIN_STATE_PLL_ON_V = 0x1
const RTC_CNTL_MAIN_STATE_PLL_ON_S = 18
const RTC_CNTL_MAIN_STATE_XTAL_ISO_V = 0x1
const RTC_CNTL_MAIN_STATE_XTAL_ISO_S = 17
const RTC_CNTL_COCPU_STATE_DONE_V = 0x1
const RTC_CNTL_COCPU_STATE_DONE_S = 16
const RTC_CNTL_COCPU_STATE_SLP_V = 0x1
const RTC_CNTL_COCPU_STATE_SLP_S = 15
const RTC_CNTL_COCPU_STATE_SWITCH_V = 0x1
const RTC_CNTL_COCPU_STATE_SWITCH_S = 14
const RTC_CNTL_COCPU_STATE_START_V = 0x1
const RTC_CNTL_COCPU_STATE_START_S = 13
const RTC_CNTL_TOUCH_STATE_DONE_V = 0x1
const RTC_CNTL_TOUCH_STATE_DONE_S = 12
const RTC_CNTL_TOUCH_STATE_SLP_V = 0x1
const RTC_CNTL_TOUCH_STATE_SLP_S = 11
const RTC_CNTL_TOUCH_STATE_SWITCH_V = 0x1
const RTC_CNTL_TOUCH_STATE_SWITCH_S = 10
const RTC_CNTL_TOUCH_STATE_START_V = 0x1
const RTC_CNTL_TOUCH_STATE_START_S = 9
const RTC_CNTL_XPD_DIG_V = 0x1
const RTC_CNTL_XPD_DIG_S = 8
const RTC_CNTL_DIG_ISO_V = 0x1
const RTC_CNTL_DIG_ISO_S = 7
const RTC_CNTL_XPD_WIFI_V = 0x1
const RTC_CNTL_XPD_WIFI_S = 6
const RTC_CNTL_WIFI_ISO_V = 0x1
const RTC_CNTL_WIFI_ISO_S = 5
const RTC_CNTL_XPD_RTC_PERI_V = 0x1
const RTC_CNTL_XPD_RTC_PERI_S = 4
const RTC_CNTL_PERI_ISO_V = 0x1
const RTC_CNTL_PERI_ISO_S = 3
const RTC_CNTL_XPD_DIG_DCDC_V = 0x1
const RTC_CNTL_XPD_DIG_DCDC_S = 2
const RTC_CNTL_XPD_ROM0_V = 0x1
const RTC_CNTL_XPD_ROM0_S = 0
const RTC_CNTL_LOW_POWER_DIAG1 = 0xFFFFFFFF
const RTC_CNTL_LOW_POWER_DIAG1_V = 0xFFFFFFFF
const RTC_CNTL_LOW_POWER_DIAG1_S = 0
const RTC_CNTL_GPIO_PIN5_HOLD_V = 0x1
const RTC_CNTL_GPIO_PIN5_HOLD_S = 5
const RTC_CNTL_GPIO_PIN4_HOLD_V = 0x1
const RTC_CNTL_GPIO_PIN4_HOLD_S = 4
const RTC_CNTL_GPIO_PIN3_HOLD_V = 0x1
const RTC_CNTL_GPIO_PIN3_HOLD_S = 3
const RTC_CNTL_GPIO_PIN2_HOLD_V = 0x1
const RTC_CNTL_GPIO_PIN2_HOLD_S = 2
const RTC_CNTL_GPIO_PIN1_HOLD_V = 0x1
const RTC_CNTL_GPIO_PIN1_HOLD_S = 1
const RTC_CNTL_GPIO_PIN0_HOLD_V = 0x1
const RTC_CNTL_GPIO_PIN0_HOLD_S = 0
const RTC_CNTL_DIG_PAD_HOLD = 0xFFFFFFFF
const RTC_CNTL_DIG_PAD_HOLD_V = 0xFFFFFFFF
const RTC_CNTL_DIG_PAD_HOLD_S = 0
const RTC_CNTL_BROWN_OUT_DET_V = 0x1
const RTC_CNTL_BROWN_OUT_DET_S = 31
const RTC_CNTL_BROWN_OUT_ENA_V = 0x1
const RTC_CNTL_BROWN_OUT_ENA_S = 30
const RTC_CNTL_BROWN_OUT_CNT_CLR_V = 0x1
const RTC_CNTL_BROWN_OUT_CNT_CLR_S = 29
const RTC_CNTL_BROWN_OUT_ANA_RST_EN_V = 0x1
const RTC_CNTL_BROWN_OUT_ANA_RST_EN_S = 28
const RTC_CNTL_BROWN_OUT_RST_SEL_V = 0x1
const RTC_CNTL_BROWN_OUT_RST_SEL_S = 27
const RTC_CNTL_BROWN_OUT_RST_ENA_V = 0x1
const RTC_CNTL_BROWN_OUT_RST_ENA_S = 26
const RTC_CNTL_BROWN_OUT_RST_WAIT = 0x000003FF
const RTC_CNTL_BROWN_OUT_RST_WAIT_V = 0x3FF
const RTC_CNTL_BROWN_OUT_RST_WAIT_S = 16
const RTC_CNTL_BROWN_OUT_PD_RF_ENA_V = 0x1
const RTC_CNTL_BROWN_OUT_PD_RF_ENA_S = 15
const RTC_CNTL_BROWN_OUT_CLOSE_FLASH_ENA_V = 0x1
const RTC_CNTL_BROWN_OUT_CLOSE_FLASH_ENA_S = 14
const RTC_CNTL_BROWN_OUT_INT_WAIT = 0x000003FF
const RTC_CNTL_BROWN_OUT_INT_WAIT_V = 0x3FF
const RTC_CNTL_BROWN_OUT_INT_WAIT_S = 4
const RTC_CNTL_TIMER_VALUE1_LOW = 0xFFFFFFFF
const RTC_CNTL_TIMER_VALUE1_LOW_V = 0xFFFFFFFF
const RTC_CNTL_TIMER_VALUE1_LOW_S = 0
const RTC_CNTL_TIMER_VALUE1_HIGH = 0x0000FFFF
const RTC_CNTL_TIMER_VALUE1_HIGH_V = 0xFFFF
const RTC_CNTL_TIMER_VALUE1_HIGH_S = 0
const RTC_CNTL_XTAL32K_CLK_FACTOR = 0xFFFFFFFF
const RTC_CNTL_XTAL32K_CLK_FACTOR_V = 0xFFFFFFFF
const RTC_CNTL_XTAL32K_CLK_FACTOR_S = 0
const RTC_CNTL_XTAL32K_STABLE_THRES = 0x0000000F
const RTC_CNTL_XTAL32K_STABLE_THRES_V = 0xF
const RTC_CNTL_XTAL32K_STABLE_THRES_S = 28
const RTC_CNTL_XTAL32K_WDT_TIMEOUT = 0x000000FF
const RTC_CNTL_XTAL32K_WDT_TIMEOUT_V = 0xFF
const RTC_CNTL_XTAL32K_WDT_TIMEOUT_S = 20
const RTC_CNTL_XTAL32K_RESTART_WAIT = 0x0000FFFF
const RTC_CNTL_XTAL32K_RESTART_WAIT_V = 0xFFFF
const RTC_CNTL_XTAL32K_RESTART_WAIT_S = 4
const RTC_CNTL_XTAL32K_RETURN_WAIT = 0x0000000F
const RTC_CNTL_XTAL32K_RETURN_WAIT_V = 0xF
const RTC_CNTL_XTAL32K_RETURN_WAIT_S = 0
const RTC_CNTL_IO_MUX_RESET_DISABLE_V = 0x1
const RTC_CNTL_IO_MUX_RESET_DISABLE_S = 18
const RTC_CNTL_REJECT_CAUSE = 0x0003FFFF
const RTC_CNTL_REJECT_CAUSE_V = 0x3FFFF
const RTC_CNTL_REJECT_CAUSE_S = 0
const RTC_CNTL_FORCE_DOWNLOAD_BOOT_V = 0x1
const RTC_CNTL_FORCE_DOWNLOAD_BOOT_S = 0
const RTC_CNTL_WAKEUP_CAUSE = 0x0001FFFF
const RTC_CNTL_WAKEUP_CAUSE_V = 0x1FFFF
const RTC_CNTL_WAKEUP_CAUSE_S = 0
const RTC_CNTL_ULP_CP_TIMER_SLP_CYCLE = 0x00FFFFFF
const RTC_CNTL_ULP_CP_TIMER_SLP_CYCLE_V = 0xFFFFFF
const RTC_CNTL_ULP_CP_TIMER_SLP_CYCLE_S = 8
const RTC_CNTL_BBPLL_CAL_INT_ENA_W1TS_V = 0x1
const RTC_CNTL_BBPLL_CAL_INT_ENA_W1TS_S = 20
const RTC_CNTL_GLITCH_DET_INT_ENA_W1TS_V = 0x1
const RTC_CNTL_GLITCH_DET_INT_ENA_W1TS_S = 19
const RTC_CNTL_XTAL32K_DEAD_INT_ENA_W1TS_V = 0x1
const RTC_CNTL_XTAL32K_DEAD_INT_ENA_W1TS_S = 16
const RTC_CNTL_SWD_INT_ENA_W1TS_V = 0x1
const RTC_CNTL_SWD_INT_ENA_W1TS_S = 15
const RTC_CNTL_MAIN_TIMER_INT_ENA_W1TS_V = 0x1
const RTC_CNTL_MAIN_TIMER_INT_ENA_W1TS_S = 10
const RTC_CNTL_BROWN_OUT_INT_ENA_W1TS_V = 0x1
const RTC_CNTL_BROWN_OUT_INT_ENA_W1TS_S = 9
const RTC_CNTL_WDT_INT_ENA_W1TS_V = 0x1
const RTC_CNTL_WDT_INT_ENA_W1TS_S = 3
const RTC_CNTL_SLP_REJECT_INT_ENA_W1TS_V = 0x1
const RTC_CNTL_SLP_REJECT_INT_ENA_W1TS_S = 1
const RTC_CNTL_SLP_WAKEUP_INT_ENA_W1TS_V = 0x1
const RTC_CNTL_SLP_WAKEUP_INT_ENA_W1TS_S = 0
const RTC_CNTL_BBPLL_CAL_INT_ENA_W1TC_V = 0x1
const RTC_CNTL_BBPLL_CAL_INT_ENA_W1TC_S = 20
const RTC_CNTL_GLITCH_DET_INT_ENA_W1TC_V = 0x1
const RTC_CNTL_GLITCH_DET_INT_ENA_W1TC_S = 19
const RTC_CNTL_XTAL32K_DEAD_INT_ENA_W1TC_V = 0x1
const RTC_CNTL_XTAL32K_DEAD_INT_ENA_W1TC_S = 16
const RTC_CNTL_SWD_INT_ENA_W1TC_V = 0x1
const RTC_CNTL_SWD_INT_ENA_W1TC_S = 15
const RTC_CNTL_MAIN_TIMER_INT_ENA_W1TC_V = 0x1
const RTC_CNTL_MAIN_TIMER_INT_ENA_W1TC_S = 10
const RTC_CNTL_BROWN_OUT_INT_ENA_W1TC_V = 0x1
const RTC_CNTL_BROWN_OUT_INT_ENA_W1TC_S = 9
const RTC_CNTL_WDT_INT_ENA_W1TC_V = 0x1
const RTC_CNTL_WDT_INT_ENA_W1TC_S = 3
const RTC_CNTL_SLP_REJECT_INT_ENA_W1TC_V = 0x1
const RTC_CNTL_SLP_REJECT_INT_ENA_W1TC_S = 1
const RTC_CNTL_SLP_WAKEUP_INT_ENA_W1TC_V = 0x1
const RTC_CNTL_SLP_WAKEUP_INT_ENA_W1TC_S = 0
const RTC_CNTL_RETENTION_WAIT = 0x0000001F
const RTC_CNTL_RETENTION_WAIT_V = 0x1F
const RTC_CNTL_RETENTION_WAIT_S = 27
const RTC_CNTL_RETENTION_EN_V = 0x1
const RTC_CNTL_RETENTION_EN_S = 26
const RTC_CNTL_RETENTION_CLKOFF_WAIT = 0x0000000F
const RTC_CNTL_RETENTION_CLKOFF_WAIT_V = 0xF
const RTC_CNTL_RETENTION_CLKOFF_WAIT_S = 22
const RTC_CNTL_RETENTION_DONE_WAIT = 0x00000007
const RTC_CNTL_RETENTION_DONE_WAIT_V = 0x7
const RTC_CNTL_RETENTION_DONE_WAIT_S = 19
const RTC_CNTL_RETENTION_CLK_SEL_V = 0x1
const RTC_CNTL_RETENTION_CLK_SEL_S = 18
const RTC_CNTL_FIB_SEL = 0x00000007
const RTC_CNTL_FIB_SEL_V = 0x7
const RTC_CNTL_FIB_SEL_S = 0
const RTC_CNTL_GPIO_PIN0_WAKEUP_ENABLE_V = 0x1
const RTC_CNTL_GPIO_PIN0_WAKEUP_ENABLE_S = 31
const RTC_CNTL_GPIO_PIN1_WAKEUP_ENABLE_V = 0x1
const RTC_CNTL_GPIO_PIN1_WAKEUP_ENABLE_S = 30
const RTC_CNTL_GPIO_PIN2_WAKEUP_ENABLE_V = 0x1
const RTC_CNTL_GPIO_PIN2_WAKEUP_ENABLE_S = 29
const RTC_CNTL_GPIO_PIN3_WAKEUP_ENABLE_V = 0x1
const RTC_CNTL_GPIO_PIN3_WAKEUP_ENABLE_S = 28
const RTC_CNTL_GPIO_PIN4_WAKEUP_ENABLE_V = 0x1
const RTC_CNTL_GPIO_PIN4_WAKEUP_ENABLE_S = 27
const RTC_CNTL_GPIO_PIN5_WAKEUP_ENABLE_V = 0x1
const RTC_CNTL_GPIO_PIN5_WAKEUP_ENABLE_S = 26
const RTC_CNTL_GPIO_PIN0_INT_TYPE = 0x00000007
const RTC_CNTL_GPIO_PIN0_INT_TYPE_V = 0x7
const RTC_CNTL_GPIO_PIN0_INT_TYPE_S = 23
const RTC_CNTL_GPIO_PIN1_INT_TYPE = 0x00000007
const RTC_CNTL_GPIO_PIN1_INT_TYPE_V = 0x7
const RTC_CNTL_GPIO_PIN1_INT_TYPE_S = 20
const RTC_CNTL_GPIO_PIN2_INT_TYPE = 0x00000007
const RTC_CNTL_GPIO_PIN2_INT_TYPE_V = 0x7
const RTC_CNTL_GPIO_PIN2_INT_TYPE_S = 17
const RTC_CNTL_GPIO_PIN3_INT_TYPE = 0x00000007
const RTC_CNTL_GPIO_PIN3_INT_TYPE_V = 0x7
const RTC_CNTL_GPIO_PIN3_INT_TYPE_S = 14
const RTC_CNTL_GPIO_PIN4_INT_TYPE = 0x00000007
const RTC_CNTL_GPIO_PIN4_INT_TYPE_V = 0x7
const RTC_CNTL_GPIO_PIN4_INT_TYPE_S = 11
const RTC_CNTL_GPIO_PIN5_INT_TYPE = 0x00000007
const RTC_CNTL_GPIO_PIN5_INT_TYPE_V = 0x7
const RTC_CNTL_GPIO_PIN5_INT_TYPE_S = 8
const RTC_CNTL_GPIO_PIN_CLK_GATE_V = 0x1
const RTC_CNTL_GPIO_PIN_CLK_GATE_S = 7
const RTC_CNTL_GPIO_WAKEUP_STATUS_CLR_V = 0x1
const RTC_CNTL_GPIO_WAKEUP_STATUS_CLR_S = 6
const RTC_CNTL_GPIO_WAKEUP_STATUS = 0x0000003F
const RTC_CNTL_GPIO_WAKEUP_STATUS_V = 0x3F
const RTC_CNTL_GPIO_WAKEUP_STATUS_S = 0
const RTC_CNTL_DEBUG_SEL4 = 0x0000001F
const RTC_CNTL_DEBUG_SEL4_V = 0x1F
const RTC_CNTL_DEBUG_SEL4_S = 27
const RTC_CNTL_DEBUG_SEL3 = 0x0000001F
const RTC_CNTL_DEBUG_SEL3_V = 0x1F
const RTC_CNTL_DEBUG_SEL3_S = 22
const RTC_CNTL_DEBUG_SEL2 = 0x0000001F
const RTC_CNTL_DEBUG_SEL2_V = 0x1F
const RTC_CNTL_DEBUG_SEL2_S = 17
const RTC_CNTL_DEBUG_SEL1 = 0x0000001F
const RTC_CNTL_DEBUG_SEL1_V = 0x1F
const RTC_CNTL_DEBUG_SEL1_S = 12
const RTC_CNTL_DEBUG_SEL0 = 0x0000001F
const RTC_CNTL_DEBUG_SEL0_V = 0x1F
const RTC_CNTL_DEBUG_SEL0_S = 7
const RTC_CNTL_DEBUG_BIT_SEL = 0x0000001F
const RTC_CNTL_DEBUG_BIT_SEL_V = 0x1F
const RTC_CNTL_DEBUG_BIT_SEL_S = 2
const RTC_CNTL_DEBUG_12M_NO_GATING_V = 0x1
const RTC_CNTL_DEBUG_12M_NO_GATING_S = 1
const RTC_CNTL_GPIO_PIN0_FUN_SEL = 0x0000000F
const RTC_CNTL_GPIO_PIN0_FUN_SEL_V = 0xF
const RTC_CNTL_GPIO_PIN0_FUN_SEL_S = 28
const RTC_CNTL_GPIO_PIN1_FUN_SEL = 0x0000000F
const RTC_CNTL_GPIO_PIN1_FUN_SEL_V = 0xF
const RTC_CNTL_GPIO_PIN1_FUN_SEL_S = 24
const RTC_CNTL_GPIO_PIN2_FUN_SEL = 0x0000000F
const RTC_CNTL_GPIO_PIN2_FUN_SEL_V = 0xF
const RTC_CNTL_GPIO_PIN2_FUN_SEL_S = 20
const RTC_CNTL_GPIO_PIN3_FUN_SEL = 0x0000000F
const RTC_CNTL_GPIO_PIN3_FUN_SEL_V = 0xF
const RTC_CNTL_GPIO_PIN3_FUN_SEL_S = 16
const RTC_CNTL_GPIO_PIN4_FUN_SEL = 0x0000000F
const RTC_CNTL_GPIO_PIN4_FUN_SEL_V = 0xF
const RTC_CNTL_GPIO_PIN4_FUN_SEL_S = 12
const RTC_CNTL_GPIO_PIN5_FUN_SEL = 0x0000000F
const RTC_CNTL_GPIO_PIN5_FUN_SEL_V = 0xF
const RTC_CNTL_GPIO_PIN5_FUN_SEL_S = 8
const RTC_CNTL_GPIO_PIN0_MUX_SEL_V = 0x1
const RTC_CNTL_GPIO_PIN0_MUX_SEL_S = 7
const RTC_CNTL_GPIO_PIN1_MUX_SEL_V = 0x1
const RTC_CNTL_GPIO_PIN1_MUX_SEL_S = 6
const RTC_CNTL_GPIO_PIN2_MUX_SEL_V = 0x1
const RTC_CNTL_GPIO_PIN2_MUX_SEL_S = 5
const RTC_CNTL_GPIO_PIN3_MUX_SEL_V = 0x1
const RTC_CNTL_GPIO_PIN3_MUX_SEL_S = 4
const RTC_CNTL_GPIO_PIN4_MUX_SEL_V = 0x1
const RTC_CNTL_GPIO_PIN4_MUX_SEL_S = 3
const RTC_CNTL_GPIO_PIN5_MUX_SEL_V = 0x1
const RTC_CNTL_GPIO_PIN5_MUX_SEL_S = 2
const RTC_CNTL_FORCE_XPD_SAR = 0x00000003
const RTC_CNTL_FORCE_XPD_SAR_V = 0x3
const RTC_CNTL_FORCE_XPD_SAR_S = 30
const RTC_CNTL_SAR2_PWDET_CCT = 0x00000007
const RTC_CNTL_SAR2_PWDET_CCT_V = 0x7
const RTC_CNTL_SAR2_PWDET_CCT_S = 27
const RTC_CNTL_SAR_DEBUG_SEL = 0x0000001F
const RTC_CNTL_SAR_DEBUG_SEL_V = 0x1F
const RTC_CNTL_SAR_DEBUG_SEL_S = 27
const RTC_CNTL_POWER_GLITCH_EN_V = 0x1
const RTC_CNTL_POWER_GLITCH_EN_S = 31
const RTC_CNTL_POWER_GLITCH_EFUSE_SEL_V = 0x1
const RTC_CNTL_POWER_GLITCH_EFUSE_SEL_S = 30
const RTC_CNTL_POWER_GLITCH_FORCE_PU_V = 0x1
const RTC_CNTL_POWER_GLITCH_FORCE_PU_S = 29
const RTC_CNTL_POWER_GLITCH_FORCE_PD_V = 0x1
const RTC_CNTL_POWER_GLITCH_FORCE_PD_S = 28
const RTC_CNTL_POWER_GLITCH_DSENSE = 0x00000003
const RTC_CNTL_POWER_GLITCH_DSENSE_V = 0x3
const RTC_CNTL_POWER_GLITCH_DSENSE_S = 26
const RTC_CNTL_CNTL_DATE = 0x0FFFFFFF
const RTC_CNTL_CNTL_DATE_V = 0xFFFFFFF
const RTC_CNTL_CNTL_DATE_S = 0
