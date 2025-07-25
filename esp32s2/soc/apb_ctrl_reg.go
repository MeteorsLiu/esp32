package soc

import _ "unsafe"

const APB_CTRL_SOC_CLK_SEL = 0x00000003
const APB_CTRL_SOC_CLK_SEL_V = 0x3
const APB_CTRL_SOC_CLK_SEL_S = 14
const APB_CTRL_SOC_CLK_SEL_XTL = 0
const APB_CTRL_SOC_CLK_SEL_PLL = 1
const APB_CTRL_SOC_CLK_SEL_8M = 2
const APB_CTRL_SOC_CLK_SEL_APLL = 3
const APB_CTRL_RST_TICK_CNT_V = 0x1
const APB_CTRL_RST_TICK_CNT_S = 12
const APB_CTRL_CLK_EN_V = 0x1
const APB_CTRL_CLK_EN_S = 11
const APB_CTRL_CLK_320M_EN_V = 0x1
const APB_CTRL_CLK_320M_EN_S = 10
const APB_CTRL_PRE_DIV_CNT = 0x000003FF
const APB_CTRL_PRE_DIV_CNT_V = 0x3FF
const APB_CTRL_PRE_DIV_CNT_S = 0
const APB_CTRL_TICK_ENABLE_V = 0x1
const APB_CTRL_TICK_ENABLE_S = 16
const APB_CTRL_CK8M_TICK_NUM = 0x000000FF
const APB_CTRL_CK8M_TICK_NUM_V = 0xFF
const APB_CTRL_CK8M_TICK_NUM_S = 8
const APB_CTRL_XTAL_TICK_NUM = 0x000000FF
const APB_CTRL_XTAL_TICK_NUM_V = 0xFF
const APB_CTRL_XTAL_TICK_NUM_S = 0
const APB_CTRL_CLK_XTAL_OEN_V = 0x1
const APB_CTRL_CLK_XTAL_OEN_S = 10
const APB_CTRL_CLK40X_BB_OEN_V = 0x1
const APB_CTRL_CLK40X_BB_OEN_S = 9
const APB_CTRL_CLK_DAC_CPU_OEN_V = 0x1
const APB_CTRL_CLK_DAC_CPU_OEN_S = 8
const APB_CTRL_CLK_ADC_INF_OEN_V = 0x1
const APB_CTRL_CLK_ADC_INF_OEN_S = 7
const APB_CTRL_CLK_320M_OEN_V = 0x1
const APB_CTRL_CLK_320M_OEN_S = 6
const APB_CTRL_CLK160_OEN_V = 0x1
const APB_CTRL_CLK160_OEN_S = 5
const APB_CTRL_CLK80_OEN_V = 0x1
const APB_CTRL_CLK80_OEN_S = 4
const APB_CTRL_CLK_BB_OEN_V = 0x1
const APB_CTRL_CLK_BB_OEN_S = 3
const APB_CTRL_CLK44_OEN_V = 0x1
const APB_CTRL_CLK44_OEN_S = 2
const APB_CTRL_CLK22_OEN_V = 0x1
const APB_CTRL_CLK22_OEN_S = 1
const APB_CTRL_CLK20_OEN_V = 0x1
const APB_CTRL_CLK20_OEN_S = 0
const APB_CTRL_PERI_IO_SWAP = 0x000000FF
const APB_CTRL_PERI_IO_SWAP_V = 0xFF
const APB_CTRL_PERI_IO_SWAP_S = 0
const APB_CTRL_EXT_MEM_PMS_LOCK_V = 0x1
const APB_CTRL_EXT_MEM_PMS_LOCK_S = 0
const APB_CTRL_FLASH_ACE0_ATTR = 0x00000007
const APB_CTRL_FLASH_ACE0_ATTR_V = 0x7
const APB_CTRL_FLASH_ACE0_ATTR_S = 0
const APB_CTRL_FLASH_ACE1_ATTR = 0x00000007
const APB_CTRL_FLASH_ACE1_ATTR_V = 0x7
const APB_CTRL_FLASH_ACE1_ATTR_S = 0
const APB_CTRL_FLASH_ACE2_ATTR = 0x00000007
const APB_CTRL_FLASH_ACE2_ATTR_V = 0x7
const APB_CTRL_FLASH_ACE2_ATTR_S = 0
const APB_CTRL_FLASH_ACE3_ATTR = 0x00000007
const APB_CTRL_FLASH_ACE3_ATTR_V = 0x7
const APB_CTRL_FLASH_ACE3_ATTR_S = 0
const APB_CTRL_FLASH_ACE0_ADDR_S = 0xFFFFFFFF
const APB_CTRL_FLASH_ACE0_ADDR_S_V = 0xFFFFFFFF
const APB_CTRL_FLASH_ACE0_ADDR_S_S = 0
const APB_CTRL_FLASH_ACE1_ADDR_S = 0xFFFFFFFF
const APB_CTRL_FLASH_ACE1_ADDR_S_V = 0xFFFFFFFF
const APB_CTRL_FLASH_ACE1_ADDR_S_S = 0
const APB_CTRL_FLASH_ACE2_ADDR_S = 0xFFFFFFFF
const APB_CTRL_FLASH_ACE2_ADDR_S_V = 0xFFFFFFFF
const APB_CTRL_FLASH_ACE2_ADDR_S_S = 0
const APB_CTRL_FLASH_ACE3_ADDR_S = 0xFFFFFFFF
const APB_CTRL_FLASH_ACE3_ADDR_S_V = 0xFFFFFFFF
const APB_CTRL_FLASH_ACE3_ADDR_S_S = 0
const APB_CTRL_FLASH_ACE0_SIZE = 0x0000FFFF
const APB_CTRL_FLASH_ACE0_SIZE_V = 0xFFFF
const APB_CTRL_FLASH_ACE0_SIZE_S = 0
const APB_CTRL_FLASH_ACE1_SIZE = 0x0000FFFF
const APB_CTRL_FLASH_ACE1_SIZE_V = 0xFFFF
const APB_CTRL_FLASH_ACE1_SIZE_S = 0
const APB_CTRL_FLASH_ACE2_SIZE = 0x0000FFFF
const APB_CTRL_FLASH_ACE2_SIZE_V = 0xFFFF
const APB_CTRL_FLASH_ACE2_SIZE_S = 0
const APB_CTRL_FLASH_ACE3_SIZE = 0x0000FFFF
const APB_CTRL_FLASH_ACE3_SIZE_V = 0xFFFF
const APB_CTRL_FLASH_ACE3_SIZE_S = 0
const APB_CTRL_SRAM_ACE0_ATTR = 0x00000007
const APB_CTRL_SRAM_ACE0_ATTR_V = 0x7
const APB_CTRL_SRAM_ACE0_ATTR_S = 0
const APB_CTRL_SRAM_ACE1_ATTR = 0x00000007
const APB_CTRL_SRAM_ACE1_ATTR_V = 0x7
const APB_CTRL_SRAM_ACE1_ATTR_S = 0
const APB_CTRL_SRAM_ACE2_ATTR = 0x00000007
const APB_CTRL_SRAM_ACE2_ATTR_V = 0x7
const APB_CTRL_SRAM_ACE2_ATTR_S = 0
const APB_CTRL_SRAM_ACE3_ATTR = 0x00000007
const APB_CTRL_SRAM_ACE3_ATTR_V = 0x7
const APB_CTRL_SRAM_ACE3_ATTR_S = 0
const APB_CTRL_SRAM_ACE0_ADDR_S = 0xFFFFFFFF
const APB_CTRL_SRAM_ACE0_ADDR_S_V = 0xFFFFFFFF
const APB_CTRL_SRAM_ACE0_ADDR_S_S = 0
const APB_CTRL_SRAM_ACE1_ADDR_S = 0xFFFFFFFF
const APB_CTRL_SRAM_ACE1_ADDR_S_V = 0xFFFFFFFF
const APB_CTRL_SRAM_ACE1_ADDR_S_S = 0
const APB_CTRL_SRAM_ACE2_ADDR_S = 0xFFFFFFFF
const APB_CTRL_SRAM_ACE2_ADDR_S_V = 0xFFFFFFFF
const APB_CTRL_SRAM_ACE2_ADDR_S_S = 0
const APB_CTRL_SRAM_ACE3_ADDR_S = 0xFFFFFFFF
const APB_CTRL_SRAM_ACE3_ADDR_S_V = 0xFFFFFFFF
const APB_CTRL_SRAM_ACE3_ADDR_S_S = 0
const APB_CTRL_SRAM_ACE0_SIZE = 0x0000FFFF
const APB_CTRL_SRAM_ACE0_SIZE_V = 0xFFFF
const APB_CTRL_SRAM_ACE0_SIZE_S = 0
const APB_CTRL_SRAM_ACE1_SIZE = 0x0000FFFF
const APB_CTRL_SRAM_ACE1_SIZE_V = 0xFFFF
const APB_CTRL_SRAM_ACE1_SIZE_S = 0
const APB_CTRL_SRAM_ACE2_SIZE = 0x0000FFFF
const APB_CTRL_SRAM_ACE2_SIZE_V = 0xFFFF
const APB_CTRL_SRAM_ACE2_SIZE_S = 0
const APB_CTRL_SRAM_ACE3_SIZE = 0x0000FFFF
const APB_CTRL_SRAM_ACE3_SIZE_V = 0xFFFF
const APB_CTRL_SRAM_ACE3_SIZE_S = 0
const APB_CTRL_SPI_MEM_REJECT_CDE = 0x0000001F
const APB_CTRL_SPI_MEM_REJECT_CDE_V = 0x1F
const APB_CTRL_SPI_MEM_REJECT_CDE_S = 2
const APB_CTRL_SPI_MEM_REJECT_CLR_V = 0x1
const APB_CTRL_SPI_MEM_REJECT_CLR_S = 1
const APB_CTRL_SPI_MEM_REJECT_INT_V = 0x1
const APB_CTRL_SPI_MEM_REJECT_INT_S = 0
const APB_CTRL_SPI_MEM_REJECT_ADDR = 0xFFFFFFFF
const APB_CTRL_SPI_MEM_REJECT_ADDR_V = 0xFFFFFFFF
const APB_CTRL_SPI_MEM_REJECT_ADDR_S = 0
const APB_CTRL_SDIO_WIN_ACCESS_EN_V = 0x1
const APB_CTRL_SDIO_WIN_ACCESS_EN_S = 0
const APB_CTRL_REDCY_ANDOR_V = 0x1
const APB_CTRL_REDCY_ANDOR_S = 31
const APB_CTRL_REDCY_SIG0 = 0x7FFFFFFF
const APB_CTRL_REDCY_SIG0_V = 0x7FFFFFFF
const APB_CTRL_REDCY_SIG0_S = 0
const APB_CTRL_REDCY_NANDOR_V = 0x1
const APB_CTRL_REDCY_NANDOR_S = 31
const APB_CTRL_REDCY_SIG1 = 0x7FFFFFFF
const APB_CTRL_REDCY_SIG1_V = 0x7FFFFFFF
const APB_CTRL_REDCY_SIG1_S = 0
const APB_CTRL_WIFI_BB_CFG = 0xFFFFFFFF
const APB_CTRL_WIFI_BB_CFG_V = 0xFFFFFFFF
const APB_CTRL_WIFI_BB_CFG_S = 0
const APB_CTRL_WIFI_BB_CFG_2 = 0xFFFFFFFF
const APB_CTRL_WIFI_BB_CFG_2_V = 0xFFFFFFFF
const APB_CTRL_WIFI_BB_CFG_2_S = 0
const APB_CTRL_WIFI_CLK_EN = 0xFFFFFFFF
const APB_CTRL_WIFI_CLK_EN_V = 0xFFFFFFFF
const APB_CTRL_WIFI_CLK_EN_S = 0
const APB_CTRL_WIFI_RST = 0xFFFFFFFF
const APB_CTRL_WIFI_RST_V = 0xFFFFFFFF
const APB_CTRL_WIFI_RST_S = 0
const APB_CTRL_DC_MEM_FORCE_PD_V = 0x1
const APB_CTRL_DC_MEM_FORCE_PD_S = 5
const APB_CTRL_DC_MEM_FORCE_PU_V = 0x1
const APB_CTRL_DC_MEM_FORCE_PU_S = 4
const APB_CTRL_PBUS_MEM_FORCE_PD_V = 0x1
const APB_CTRL_PBUS_MEM_FORCE_PD_S = 3
const APB_CTRL_PBUS_MEM_FORCE_PU_V = 0x1
const APB_CTRL_PBUS_MEM_FORCE_PU_S = 2
const APB_CTRL_AGC_MEM_FORCE_PD_V = 0x1
const APB_CTRL_AGC_MEM_FORCE_PD_S = 1
const APB_CTRL_AGC_MEM_FORCE_PU_V = 0x1
const APB_CTRL_AGC_MEM_FORCE_PU_S = 0
const APB_CTRL_DATE = 0xFFFFFFFF
const APB_CTRL_DATE_V = 0xFFFFFFFF
const APB_CTRL_DATE_S = 0
