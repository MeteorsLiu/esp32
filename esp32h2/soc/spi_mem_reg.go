package soc

import _ "unsafe"

const SPI_MEM_FLASH_READ_V = 0x1
const SPI_MEM_FLASH_READ_S = 31
const SPI_MEM_FLASH_WREN_V = 0x1
const SPI_MEM_FLASH_WREN_S = 30
const SPI_MEM_FLASH_WRDI_V = 0x1
const SPI_MEM_FLASH_WRDI_S = 29
const SPI_MEM_FLASH_RDID_V = 0x1
const SPI_MEM_FLASH_RDID_S = 28
const SPI_MEM_FLASH_RDSR_V = 0x1
const SPI_MEM_FLASH_RDSR_S = 27
const SPI_MEM_FLASH_WRSR_V = 0x1
const SPI_MEM_FLASH_WRSR_S = 26
const SPI_MEM_FLASH_PP_V = 0x1
const SPI_MEM_FLASH_PP_S = 25
const SPI_MEM_FLASH_SE_V = 0x1
const SPI_MEM_FLASH_SE_S = 24
const SPI_MEM_FLASH_BE_V = 0x1
const SPI_MEM_FLASH_BE_S = 23
const SPI_MEM_FLASH_CE_V = 0x1
const SPI_MEM_FLASH_CE_S = 22
const SPI_MEM_FLASH_DP_V = 0x1
const SPI_MEM_FLASH_DP_S = 21
const SPI_MEM_FLASH_RES_V = 0x1
const SPI_MEM_FLASH_RES_S = 20
const SPI_MEM_FLASH_HPM_V = 0x1
const SPI_MEM_FLASH_HPM_S = 19
const SPI_MEM_USR_V = 0x1
const SPI_MEM_USR_S = 18
const SPI_MEM_FLASH_PE_V = 0x1
const SPI_MEM_FLASH_PE_S = 17
const SPI_MEM_SLV_ST = 0x0000000F
const SPI_MEM_SLV_ST_V = 0xF
const SPI_MEM_SLV_ST_S = 4
const SPI_MEM_MST_ST = 0x0000000F
const SPI_MEM_MST_ST_V = 0xF
const SPI_MEM_MST_ST_S = 0
const SPI_MEM_USR_ADDR_VALUE = 0xFFFFFFFF
const SPI_MEM_USR_ADDR_VALUE_V = 0xFFFFFFFF
const SPI_MEM_USR_ADDR_VALUE_S = 0
const SPI_MEM_DATA_IE_ALWAYS_ON_V = 0x1
const SPI_MEM_DATA_IE_ALWAYS_ON_S = 31
const SPI_MEM_DQS_IE_ALWAYS_ON_V = 0x1
const SPI_MEM_DQS_IE_ALWAYS_ON_S = 30
const SPI_MEM_FREAD_QIO_V = 0x1
const SPI_MEM_FREAD_QIO_S = 24
const SPI_MEM_FREAD_DIO_V = 0x1
const SPI_MEM_FREAD_DIO_S = 23
const SPI_MEM_WRSR_2B_V = 0x1
const SPI_MEM_WRSR_2B_S = 22
const SPI_MEM_WP_REG_V = 0x1
const SPI_MEM_WP_REG_S = 21
const SPI_MEM_FREAD_QUAD_V = 0x1
const SPI_MEM_FREAD_QUAD_S = 20
const SPI_MEM_D_POL_V = 0x1
const SPI_MEM_D_POL_S = 19
const SPI_MEM_Q_POL_V = 0x1
const SPI_MEM_Q_POL_S = 18
const SPI_MEM_RESANDRES_V = 0x1
const SPI_MEM_RESANDRES_S = 15
const SPI_MEM_FREAD_DUAL_V = 0x1
const SPI_MEM_FREAD_DUAL_S = 14
const SPI_MEM_FASTRD_MODE_V = 0x1
const SPI_MEM_FASTRD_MODE_S = 13
const SPI_MEM_TX_CRC_EN_V = 0x1
const SPI_MEM_TX_CRC_EN_S = 11
const SPI_MEM_FCS_CRC_EN_V = 0x1
const SPI_MEM_FCS_CRC_EN_S = 10
const SPI_MEM_FCMD_OCT_V = 0x1
const SPI_MEM_FCMD_OCT_S = 9
const SPI_MEM_FCMD_QUAD_V = 0x1
const SPI_MEM_FCMD_QUAD_S = 8
const SPI_MEM_FADDR_OCT_V = 0x1
const SPI_MEM_FADDR_OCT_S = 6
const SPI_MEM_FDIN_OCT_V = 0x1
const SPI_MEM_FDIN_OCT_S = 5
const SPI_MEM_FDOUT_OCT_V = 0x1
const SPI_MEM_FDOUT_OCT_S = 4
const SPI_MEM_FDUMMY_WOUT_V = 0x1
const SPI_MEM_FDUMMY_WOUT_S = 3
const SPI_MEM_FDUMMY_RIN_V = 0x1
const SPI_MEM_FDUMMY_RIN_S = 2
const SPI_MEM_WDUMMY_ALWAYS_OUT_V = 0x1
const SPI_MEM_WDUMMY_ALWAYS_OUT_S = 1
const SPI_MEM_WDUMMY_DQS_ALWAYS_OUT_V = 0x1
const SPI_MEM_WDUMMY_DQS_ALWAYS_OUT_S = 0
const SPI_MEM_TXFIFO_RST_V = 0x1
const SPI_MEM_TXFIFO_RST_S = 31
const SPI_MEM_RXFIFO_RST_V = 0x1
const SPI_MEM_RXFIFO_RST_S = 30
const SPI_MEM_FAST_WRITE_EN_V = 0x1
const SPI_MEM_FAST_WRITE_EN_S = 29
const SPI_MEM_DUAL_RAM_EN_V = 0x1
const SPI_MEM_DUAL_RAM_EN_S = 28
const SPI_MEM_RAM0_EN_V = 0x1
const SPI_MEM_RAM0_EN_S = 27
const SPI_MEM_AW_SPLICE_EN_V = 0x1
const SPI_MEM_AW_SPLICE_EN_S = 26
const SPI_MEM_AR_SPLICE_EN_V = 0x1
const SPI_MEM_AR_SPLICE_EN_S = 25
const SPI_MEM_RRESP_ECC_ERR_EN_V = 0x1
const SPI_MEM_RRESP_ECC_ERR_EN_S = 24
const SPI_MEM_AXI_RDATA_BACK_FAST_V = 0x1
const SPI_MEM_AXI_RDATA_BACK_FAST_S = 23
const SPI_MEM_AW_SIZE0_1_SUPPORT_EN_V = 0x1
const SPI_MEM_AW_SIZE0_1_SUPPORT_EN_S = 22
const SPI_MEM_AR_SIZE0_1_SUPPORT_EN_V = 0x1
const SPI_MEM_AR_SIZE0_1_SUPPORT_EN_S = 21
const SPI_MEM_CS_HOLD_DLY_RES = 0x000003FF
const SPI_MEM_CS_HOLD_DLY_RES_V = 0x3FF
const SPI_MEM_CS_HOLD_DLY_RES_S = 2
const SPI_MEM_CLK_MODE = 0x00000003
const SPI_MEM_CLK_MODE_V = 0x3
const SPI_MEM_CLK_MODE_S = 0
const SPI_MEM_SYNC_RESET_V = 0x1
const SPI_MEM_SYNC_RESET_S = 31
const SPI_MEM_CS_HOLD_DELAY = 0x0000003F
const SPI_MEM_CS_HOLD_DELAY_V = 0x3F
const SPI_MEM_CS_HOLD_DELAY_S = 25
const SPI_MEM_SPLIT_TRANS_EN_V = 0x1
const SPI_MEM_SPLIT_TRANS_EN_S = 24
const SPI_MEM_ECC_16TO18_BYTE_EN_V = 0x1
const SPI_MEM_ECC_16TO18_BYTE_EN_S = 14
const SPI_MEM_ECC_SKIP_PAGE_CORNER_V = 0x1
const SPI_MEM_ECC_SKIP_PAGE_CORNER_S = 13
const SPI_MEM_ECC_CS_HOLD_TIME = 0x00000007
const SPI_MEM_ECC_CS_HOLD_TIME_V = 0x7
const SPI_MEM_ECC_CS_HOLD_TIME_S = 10
const SPI_MEM_CS_HOLD_TIME = 0x0000001F
const SPI_MEM_CS_HOLD_TIME_V = 0x1F
const SPI_MEM_CS_HOLD_TIME_S = 5
const SPI_MEM_CS_SETUP_TIME = 0x0000001F
const SPI_MEM_CS_SETUP_TIME_V = 0x1F
const SPI_MEM_CS_SETUP_TIME_S = 0
const SPI_MEM_CLK_EQU_SYSCLK_V = 0x1
const SPI_MEM_CLK_EQU_SYSCLK_S = 31
const SPI_MEM_CLKCNT_N = 0x000000FF
const SPI_MEM_CLKCNT_N_V = 0xFF
const SPI_MEM_CLKCNT_N_S = 16
const SPI_MEM_CLKCNT_H = 0x000000FF
const SPI_MEM_CLKCNT_H_V = 0xFF
const SPI_MEM_CLKCNT_H_S = 8
const SPI_MEM_CLKCNT_L = 0x000000FF
const SPI_MEM_CLKCNT_L_V = 0xFF
const SPI_MEM_CLKCNT_L_S = 0
const SPI_MEM_USR_COMMAND_V = 0x1
const SPI_MEM_USR_COMMAND_S = 31
const SPI_MEM_USR_ADDR_V = 0x1
const SPI_MEM_USR_ADDR_S = 30
const SPI_MEM_USR_DUMMY_V = 0x1
const SPI_MEM_USR_DUMMY_S = 29
const SPI_MEM_USR_MISO_V = 0x1
const SPI_MEM_USR_MISO_S = 28
const SPI_MEM_USR_MOSI_V = 0x1
const SPI_MEM_USR_MOSI_S = 27
const SPI_MEM_USR_DUMMY_IDLE_V = 0x1
const SPI_MEM_USR_DUMMY_IDLE_S = 26
const SPI_MEM_USR_MOSI_HIGHPART_V = 0x1
const SPI_MEM_USR_MOSI_HIGHPART_S = 25
const SPI_MEM_USR_MISO_HIGHPART_V = 0x1
const SPI_MEM_USR_MISO_HIGHPART_S = 24
const SPI_MEM_FWRITE_QIO_V = 0x1
const SPI_MEM_FWRITE_QIO_S = 15
const SPI_MEM_FWRITE_DIO_V = 0x1
const SPI_MEM_FWRITE_DIO_S = 14
const SPI_MEM_FWRITE_QUAD_V = 0x1
const SPI_MEM_FWRITE_QUAD_S = 13
const SPI_MEM_FWRITE_DUAL_V = 0x1
const SPI_MEM_FWRITE_DUAL_S = 12
const SPI_MEM_CK_OUT_EDGE_V = 0x1
const SPI_MEM_CK_OUT_EDGE_S = 9
const SPI_MEM_CS_SETUP_V = 0x1
const SPI_MEM_CS_SETUP_S = 7
const SPI_MEM_CS_HOLD_V = 0x1
const SPI_MEM_CS_HOLD_S = 6
const SPI_MEM_USR_ADDR_BITLEN = 0x0000003F
const SPI_MEM_USR_ADDR_BITLEN_V = 0x3F
const SPI_MEM_USR_ADDR_BITLEN_S = 26
const SPI_MEM_USR_DBYTELEN = 0x00000007
const SPI_MEM_USR_DBYTELEN_V = 0x7
const SPI_MEM_USR_DBYTELEN_S = 6
const SPI_MEM_USR_DUMMY_CYCLELEN = 0x0000003F
const SPI_MEM_USR_DUMMY_CYCLELEN_V = 0x3F
const SPI_MEM_USR_DUMMY_CYCLELEN_S = 0
const SPI_MEM_USR_COMMAND_BITLEN = 0x0000000F
const SPI_MEM_USR_COMMAND_BITLEN_V = 0xF
const SPI_MEM_USR_COMMAND_BITLEN_S = 28
const SPI_MEM_USR_COMMAND_VALUE = 0x0000FFFF
const SPI_MEM_USR_COMMAND_VALUE_V = 0xFFFF
const SPI_MEM_USR_COMMAND_VALUE_S = 0
const SPI_MEM_USR_MOSI_DBITLEN = 0x000003FF
const SPI_MEM_USR_MOSI_DBITLEN_V = 0x3FF
const SPI_MEM_USR_MOSI_DBITLEN_S = 0
const SPI_MEM_USR_MISO_DBITLEN = 0x000003FF
const SPI_MEM_USR_MISO_DBITLEN_V = 0x3FF
const SPI_MEM_USR_MISO_DBITLEN_S = 0
const SPI_MEM_WB_MODE = 0x000000FF
const SPI_MEM_WB_MODE_V = 0xFF
const SPI_MEM_WB_MODE_S = 16
const SPI_MEM_STATUS = 0x0000FFFF
const SPI_MEM_STATUS_V = 0xFFFF
const SPI_MEM_STATUS_S = 0
const SPI_MEM_CS_KEEP_ACTIVE_V = 0x1
const SPI_MEM_CS_KEEP_ACTIVE_S = 10
const SPI_MEM_CK_IDLE_EDGE_V = 0x1
const SPI_MEM_CK_IDLE_EDGE_S = 9
const SPI_MEM_SSUB_PIN_V = 0x1
const SPI_MEM_SSUB_PIN_S = 8
const SPI_MEM_FSUB_PIN_V = 0x1
const SPI_MEM_FSUB_PIN_S = 7
const SPI_MEM_CS1_DIS_V = 0x1
const SPI_MEM_CS1_DIS_S = 1
const SPI_MEM_CS0_DIS_V = 0x1
const SPI_MEM_CS0_DIS_S = 0
const SPI_MEM_TX_CRC_DATA = 0xFFFFFFFF
const SPI_MEM_TX_CRC_DATA_V = 0xFFFFFFFF
const SPI_MEM_TX_CRC_DATA_S = 0
const SPI_MEM_CLOSE_AXI_INF_EN_V = 0x1
const SPI_MEM_CLOSE_AXI_INF_EN_S = 31
const SPI_MEM_SAME_AW_AR_ADDR_CHK_EN_V = 0x1
const SPI_MEM_SAME_AW_AR_ADDR_CHK_EN_S = 30
const SPI_MEM_FADDR_QUAD_V = 0x1
const SPI_MEM_FADDR_QUAD_S = 8
const SPI_MEM_FDOUT_QUAD_V = 0x1
const SPI_MEM_FDOUT_QUAD_S = 7
const SPI_MEM_FDIN_QUAD_V = 0x1
const SPI_MEM_FDIN_QUAD_S = 6
const SPI_MEM_FADDR_DUAL_V = 0x1
const SPI_MEM_FADDR_DUAL_S = 5
const SPI_MEM_FDOUT_DUAL_V = 0x1
const SPI_MEM_FDOUT_DUAL_S = 4
const SPI_MEM_FDIN_DUAL_V = 0x1
const SPI_MEM_FDIN_DUAL_S = 3
const SPI_MEM_CACHE_FLASH_USR_CMD_V = 0x1
const SPI_MEM_CACHE_FLASH_USR_CMD_S = 2
const SPI_MEM_CACHE_USR_ADDR_4BYTE_V = 0x1
const SPI_MEM_CACHE_USR_ADDR_4BYTE_S = 1
const SPI_MEM_AXI_REQ_EN_V = 0x1
const SPI_MEM_AXI_REQ_EN_S = 0
const SPI_MEM_SRAM_WDUMMY_CYCLELEN = 0x0000003F
const SPI_MEM_SRAM_WDUMMY_CYCLELEN_V = 0x3F
const SPI_MEM_SRAM_WDUMMY_CYCLELEN_S = 22
const SPI_MEM_SRAM_OCT_V = 0x1
const SPI_MEM_SRAM_OCT_S = 21
const SPI_MEM_CACHE_SRAM_USR_WCMD_V = 0x1
const SPI_MEM_CACHE_SRAM_USR_WCMD_S = 20
const SPI_MEM_SRAM_ADDR_BITLEN = 0x0000003F
const SPI_MEM_SRAM_ADDR_BITLEN_V = 0x3F
const SPI_MEM_SRAM_ADDR_BITLEN_S = 14
const SPI_MEM_SRAM_RDUMMY_CYCLELEN = 0x0000003F
const SPI_MEM_SRAM_RDUMMY_CYCLELEN_V = 0x3F
const SPI_MEM_SRAM_RDUMMY_CYCLELEN_S = 6
const SPI_MEM_CACHE_SRAM_USR_RCMD_V = 0x1
const SPI_MEM_CACHE_SRAM_USR_RCMD_S = 5
const SPI_MEM_USR_RD_SRAM_DUMMY_V = 0x1
const SPI_MEM_USR_RD_SRAM_DUMMY_S = 4
const SPI_MEM_USR_WR_SRAM_DUMMY_V = 0x1
const SPI_MEM_USR_WR_SRAM_DUMMY_S = 3
const SPI_MEM_USR_SRAM_QIO_V = 0x1
const SPI_MEM_USR_SRAM_QIO_S = 2
const SPI_MEM_USR_SRAM_DIO_V = 0x1
const SPI_MEM_USR_SRAM_DIO_S = 1
const SPI_MEM_CACHE_USR_SADDR_4BYTE_V = 0x1
const SPI_MEM_CACHE_USR_SADDR_4BYTE_S = 0
const SPI_MEM_SMEM_DATA_IE_ALWAYS_ON_V = 0x1
const SPI_MEM_SMEM_DATA_IE_ALWAYS_ON_S = 31
const SPI_MEM_SMEM_DQS_IE_ALWAYS_ON_V = 0x1
const SPI_MEM_SMEM_DQS_IE_ALWAYS_ON_S = 30
const SPI_MEM_SMEM_WDUMMY_ALWAYS_OUT_V = 0x1
const SPI_MEM_SMEM_WDUMMY_ALWAYS_OUT_S = 25
const SPI_MEM_SMEM_WDUMMY_DQS_ALWAYS_OUT_V = 0x1
const SPI_MEM_SMEM_WDUMMY_DQS_ALWAYS_OUT_S = 24
const SPI_MEM_SDUMMY_WOUT_V = 0x1
const SPI_MEM_SDUMMY_WOUT_S = 23
const SPI_MEM_SDUMMY_RIN_V = 0x1
const SPI_MEM_SDUMMY_RIN_S = 22
const SPI_MEM_SCMD_OCT_V = 0x1
const SPI_MEM_SCMD_OCT_S = 21
const SPI_MEM_SADDR_OCT_V = 0x1
const SPI_MEM_SADDR_OCT_S = 20
const SPI_MEM_SDOUT_OCT_V = 0x1
const SPI_MEM_SDOUT_OCT_S = 19
const SPI_MEM_SDIN_OCT_V = 0x1
const SPI_MEM_SDIN_OCT_S = 18
const SPI_MEM_SCMD_QUAD_V = 0x1
const SPI_MEM_SCMD_QUAD_S = 17
const SPI_MEM_SADDR_QUAD_V = 0x1
const SPI_MEM_SADDR_QUAD_S = 16
const SPI_MEM_SDOUT_QUAD_V = 0x1
const SPI_MEM_SDOUT_QUAD_S = 15
const SPI_MEM_SDIN_QUAD_V = 0x1
const SPI_MEM_SDIN_QUAD_S = 14
const SPI_MEM_SADDR_DUAL_V = 0x1
const SPI_MEM_SADDR_DUAL_S = 12
const SPI_MEM_SDOUT_DUAL_V = 0x1
const SPI_MEM_SDOUT_DUAL_S = 11
const SPI_MEM_SDIN_DUAL_V = 0x1
const SPI_MEM_SDIN_DUAL_S = 10
const SPI_MEM_SWB_MODE = 0x000000FF
const SPI_MEM_SWB_MODE_V = 0xFF
const SPI_MEM_SWB_MODE_S = 2
const SPI_MEM_SCLK_MODE = 0x00000003
const SPI_MEM_SCLK_MODE_V = 0x3
const SPI_MEM_SCLK_MODE_S = 0
const SPI_MEM_CACHE_SRAM_USR_RD_CMD_BITLEN = 0x0000000F
const SPI_MEM_CACHE_SRAM_USR_RD_CMD_BITLEN_V = 0xF
const SPI_MEM_CACHE_SRAM_USR_RD_CMD_BITLEN_S = 28
const SPI_MEM_CACHE_SRAM_USR_RD_CMD_VALUE = 0x0000FFFF
const SPI_MEM_CACHE_SRAM_USR_RD_CMD_VALUE_V = 0xFFFF
const SPI_MEM_CACHE_SRAM_USR_RD_CMD_VALUE_S = 0
const SPI_MEM_CACHE_SRAM_USR_WR_CMD_BITLEN = 0x0000000F
const SPI_MEM_CACHE_SRAM_USR_WR_CMD_BITLEN_V = 0xF
const SPI_MEM_CACHE_SRAM_USR_WR_CMD_BITLEN_S = 28
const SPI_MEM_CACHE_SRAM_USR_WR_CMD_VALUE = 0x0000FFFF
const SPI_MEM_CACHE_SRAM_USR_WR_CMD_VALUE_V = 0xFFFF
const SPI_MEM_CACHE_SRAM_USR_WR_CMD_VALUE_S = 0
const SPI_MEM_SCLK_EQU_SYSCLK_V = 0x1
const SPI_MEM_SCLK_EQU_SYSCLK_S = 31
const SPI_MEM_SCLKCNT_N = 0x000000FF
const SPI_MEM_SCLKCNT_N_V = 0xFF
const SPI_MEM_SCLKCNT_N_S = 16
const SPI_MEM_SCLKCNT_H = 0x000000FF
const SPI_MEM_SCLKCNT_H_V = 0xFF
const SPI_MEM_SCLKCNT_H_S = 8
const SPI_MEM_SCLKCNT_L = 0x000000FF
const SPI_MEM_SCLKCNT_L_V = 0xFF
const SPI_MEM_SCLKCNT_L_S = 0
const SPI_MEM_LOCK_DELAY_TIME = 0x0000001F
const SPI_MEM_LOCK_DELAY_TIME_V = 0x1F
const SPI_MEM_LOCK_DELAY_TIME_S = 7
const SPI_MEM_BUF0 = 0xFFFFFFFF
const SPI_MEM_BUF0_V = 0xFFFFFFFF
const SPI_MEM_BUF0_S = 0
const SPI_MEM_BUF1 = 0xFFFFFFFF
const SPI_MEM_BUF1_V = 0xFFFFFFFF
const SPI_MEM_BUF1_S = 0
const SPI_MEM_BUF2 = 0xFFFFFFFF
const SPI_MEM_BUF2_V = 0xFFFFFFFF
const SPI_MEM_BUF2_S = 0
const SPI_MEM_BUF3 = 0xFFFFFFFF
const SPI_MEM_BUF3_V = 0xFFFFFFFF
const SPI_MEM_BUF3_S = 0
const SPI_MEM_BUF4 = 0xFFFFFFFF
const SPI_MEM_BUF4_V = 0xFFFFFFFF
const SPI_MEM_BUF4_S = 0
const SPI_MEM_BUF5 = 0xFFFFFFFF
const SPI_MEM_BUF5_V = 0xFFFFFFFF
const SPI_MEM_BUF5_S = 0
const SPI_MEM_BUF6 = 0xFFFFFFFF
const SPI_MEM_BUF6_V = 0xFFFFFFFF
const SPI_MEM_BUF6_S = 0
const SPI_MEM_BUF7 = 0xFFFFFFFF
const SPI_MEM_BUF7_V = 0xFFFFFFFF
const SPI_MEM_BUF7_S = 0
const SPI_MEM_BUF8 = 0xFFFFFFFF
const SPI_MEM_BUF8_V = 0xFFFFFFFF
const SPI_MEM_BUF8_S = 0
const SPI_MEM_BUF9 = 0xFFFFFFFF
const SPI_MEM_BUF9_V = 0xFFFFFFFF
const SPI_MEM_BUF9_S = 0
const SPI_MEM_BUF10 = 0xFFFFFFFF
const SPI_MEM_BUF10_V = 0xFFFFFFFF
const SPI_MEM_BUF10_S = 0
const SPI_MEM_BUF11 = 0xFFFFFFFF
const SPI_MEM_BUF11_V = 0xFFFFFFFF
const SPI_MEM_BUF11_S = 0
const SPI_MEM_BUF12 = 0xFFFFFFFF
const SPI_MEM_BUF12_V = 0xFFFFFFFF
const SPI_MEM_BUF12_S = 0
const SPI_MEM_BUF13 = 0xFFFFFFFF
const SPI_MEM_BUF13_V = 0xFFFFFFFF
const SPI_MEM_BUF13_S = 0
const SPI_MEM_BUF14 = 0xFFFFFFFF
const SPI_MEM_BUF14_V = 0xFFFFFFFF
const SPI_MEM_BUF14_S = 0
const SPI_MEM_BUF15 = 0xFFFFFFFF
const SPI_MEM_BUF15_V = 0xFFFFFFFF
const SPI_MEM_BUF15_S = 0
const SPI_MEM_WAITI_CMD = 0x0000FFFF
const SPI_MEM_WAITI_CMD_V = 0xFFFF
const SPI_MEM_WAITI_CMD_S = 16
const SPI_MEM_WAITI_DUMMY_CYCLELEN = 0x0000003F
const SPI_MEM_WAITI_DUMMY_CYCLELEN_V = 0x3F
const SPI_MEM_WAITI_DUMMY_CYCLELEN_S = 10
const SPI_MEM_WAITI_CMD_2B_V = 0x1
const SPI_MEM_WAITI_CMD_2B_S = 9
const SPI_MEM_WAITI_ADDR_CYCLELEN = 0x00000003
const SPI_MEM_WAITI_ADDR_CYCLELEN_V = 0x3
const SPI_MEM_WAITI_ADDR_CYCLELEN_S = 3
const SPI_MEM_WAITI_ADDR_EN_V = 0x1
const SPI_MEM_WAITI_ADDR_EN_S = 2
const SPI_MEM_WAITI_DUMMY_V = 0x1
const SPI_MEM_WAITI_DUMMY_S = 1
const SPI_MEM_WAITI_EN_V = 0x1
const SPI_MEM_WAITI_EN_S = 0
const SPI_MEM_SUS_TIMEOUT_CNT = 0x0000007F
const SPI_MEM_SUS_TIMEOUT_CNT_V = 0x7F
const SPI_MEM_SUS_TIMEOUT_CNT_S = 25
const SPI_MEM_PES_END_EN_V = 0x1
const SPI_MEM_PES_END_EN_S = 24
const SPI_MEM_PER_END_EN_V = 0x1
const SPI_MEM_PER_END_EN_S = 23
const SPI_MEM_FMEM_RD_SUS_2B_V = 0x1
const SPI_MEM_FMEM_RD_SUS_2B_S = 22
const SPI_MEM_PESR_END_MSK = 0x0000FFFF
const SPI_MEM_PESR_END_MSK_V = 0xFFFF
const SPI_MEM_PESR_END_MSK_S = 6
const SPI_MEM_FLASH_PES_EN_V = 0x1
const SPI_MEM_FLASH_PES_EN_S = 5
const SPI_MEM_PES_PER_EN_V = 0x1
const SPI_MEM_PES_PER_EN_S = 4
const SPI_MEM_FLASH_PES_WAIT_EN_V = 0x1
const SPI_MEM_FLASH_PES_WAIT_EN_S = 3
const SPI_MEM_FLASH_PER_WAIT_EN_V = 0x1
const SPI_MEM_FLASH_PER_WAIT_EN_S = 2
const SPI_MEM_FLASH_PES_V = 0x1
const SPI_MEM_FLASH_PES_S = 1
const SPI_MEM_FLASH_PER_V = 0x1
const SPI_MEM_FLASH_PER_S = 0
const SPI_MEM_WAIT_PESR_COMMAND = 0x0000FFFF
const SPI_MEM_WAIT_PESR_COMMAND_V = 0xFFFF
const SPI_MEM_WAIT_PESR_COMMAND_S = 16
const SPI_MEM_FLASH_PES_COMMAND = 0x0000FFFF
const SPI_MEM_FLASH_PES_COMMAND_V = 0xFFFF
const SPI_MEM_FLASH_PES_COMMAND_S = 0
const SPI_MEM_FLASH_PER_COMMAND = 0x0000FFFF
const SPI_MEM_FLASH_PER_COMMAND_V = 0xFFFF
const SPI_MEM_FLASH_PER_COMMAND_S = 16
const SPI_MEM_FLASH_PESR_CMD_2B_V = 0x1
const SPI_MEM_FLASH_PESR_CMD_2B_S = 15
const SPI_MEM_SPI0_LOCK_EN_V = 0x1
const SPI_MEM_SPI0_LOCK_EN_S = 7
const SPI_MEM_FLASH_PES_DLY_128_V = 0x1
const SPI_MEM_FLASH_PES_DLY_128_S = 6
const SPI_MEM_FLASH_PER_DLY_128_V = 0x1
const SPI_MEM_FLASH_PER_DLY_128_S = 5
const SPI_MEM_FLASH_DP_DLY_128_V = 0x1
const SPI_MEM_FLASH_DP_DLY_128_S = 4
const SPI_MEM_FLASH_RES_DLY_128_V = 0x1
const SPI_MEM_FLASH_RES_DLY_128_S = 3
const SPI_MEM_FLASH_HPM_DLY_128_V = 0x1
const SPI_MEM_FLASH_HPM_DLY_128_S = 2
const SPI_MEM_WAIT_PESR_CMD_2B_V = 0x1
const SPI_MEM_WAIT_PESR_CMD_2B_S = 1
const SPI_MEM_FLASH_SUS_V = 0x1
const SPI_MEM_FLASH_SUS_S = 0
const SPI_MEM_BROWN_OUT_INT_ENA_V = 0x1
const SPI_MEM_BROWN_OUT_INT_ENA_S = 10
const SPI_MEM_AXI_WADDR_ERR_INT__ENA_V = 0x1
const SPI_MEM_AXI_WADDR_ERR_INT__ENA_S = 9
const SPI_MEM_AXI_WR_FLASH_ERR_INT_ENA_V = 0x1
const SPI_MEM_AXI_WR_FLASH_ERR_INT_ENA_S = 8
const SPI_MEM_AXI_RADDR_ERR_INT_ENA_V = 0x1
const SPI_MEM_AXI_RADDR_ERR_INT_ENA_S = 7
const SPI_MEM_PMS_REJECT_INT_ENA_V = 0x1
const SPI_MEM_PMS_REJECT_INT_ENA_S = 6
const SPI_MEM_ECC_ERR_INT_ENA_V = 0x1
const SPI_MEM_ECC_ERR_INT_ENA_S = 5
const SPI_MEM_MST_ST_END_INT_ENA_V = 0x1
const SPI_MEM_MST_ST_END_INT_ENA_S = 4
const SPI_MEM_SLV_ST_END_INT_ENA_V = 0x1
const SPI_MEM_SLV_ST_END_INT_ENA_S = 3
const SPI_MEM_WPE_END_INT_ENA_V = 0x1
const SPI_MEM_WPE_END_INT_ENA_S = 2
const SPI_MEM_PES_END_INT_ENA_V = 0x1
const SPI_MEM_PES_END_INT_ENA_S = 1
const SPI_MEM_PER_END_INT_ENA_V = 0x1
const SPI_MEM_PER_END_INT_ENA_S = 0
const SPI_MEM_BROWN_OUT_INT_CLR_V = 0x1
const SPI_MEM_BROWN_OUT_INT_CLR_S = 10
const SPI_MEM_AXI_WADDR_ERR_INT_CLR_V = 0x1
const SPI_MEM_AXI_WADDR_ERR_INT_CLR_S = 9
const SPI_MEM_AXI_WR_FLASH_ERR_INT_CLR_V = 0x1
const SPI_MEM_AXI_WR_FLASH_ERR_INT_CLR_S = 8
const SPI_MEM_AXI_RADDR_ERR_INT_CLR_V = 0x1
const SPI_MEM_AXI_RADDR_ERR_INT_CLR_S = 7
const SPI_MEM_PMS_REJECT_INT_CLR_V = 0x1
const SPI_MEM_PMS_REJECT_INT_CLR_S = 6
const SPI_MEM_ECC_ERR_INT_CLR_V = 0x1
const SPI_MEM_ECC_ERR_INT_CLR_S = 5
const SPI_MEM_MST_ST_END_INT_CLR_V = 0x1
const SPI_MEM_MST_ST_END_INT_CLR_S = 4
const SPI_MEM_SLV_ST_END_INT_CLR_V = 0x1
const SPI_MEM_SLV_ST_END_INT_CLR_S = 3
const SPI_MEM_WPE_END_INT_CLR_V = 0x1
const SPI_MEM_WPE_END_INT_CLR_S = 2
const SPI_MEM_PES_END_INT_CLR_V = 0x1
const SPI_MEM_PES_END_INT_CLR_S = 1
const SPI_MEM_PER_END_INT_CLR_V = 0x1
const SPI_MEM_PER_END_INT_CLR_S = 0
const SPI_MEM_BROWN_OUT_INT_RAW_V = 0x1
const SPI_MEM_BROWN_OUT_INT_RAW_S = 10
const SPI_MEM_AXI_WADDR_ERR_INT_RAW_V = 0x1
const SPI_MEM_AXI_WADDR_ERR_INT_RAW_S = 9
const SPI_MEM_AXI_WR_FLASH_ERR_INT_RAW_V = 0x1
const SPI_MEM_AXI_WR_FLASH_ERR_INT_RAW_S = 8
const SPI_MEM_AXI_RADDR_ERR_INT_RAW_V = 0x1
const SPI_MEM_AXI_RADDR_ERR_INT_RAW_S = 7
const SPI_MEM_PMS_REJECT_INT_RAW_V = 0x1
const SPI_MEM_PMS_REJECT_INT_RAW_S = 6
const SPI_MEM_ECC_ERR_INT_RAW_V = 0x1
const SPI_MEM_ECC_ERR_INT_RAW_S = 5
const SPI_MEM_MST_ST_END_INT_RAW_V = 0x1
const SPI_MEM_MST_ST_END_INT_RAW_S = 4
const SPI_MEM_SLV_ST_END_INT_RAW_V = 0x1
const SPI_MEM_SLV_ST_END_INT_RAW_S = 3
const SPI_MEM_WPE_END_INT_RAW_V = 0x1
const SPI_MEM_WPE_END_INT_RAW_S = 2
const SPI_MEM_PES_END_INT_RAW_V = 0x1
const SPI_MEM_PES_END_INT_RAW_S = 1
const SPI_MEM_PER_END_INT_RAW_V = 0x1
const SPI_MEM_PER_END_INT_RAW_S = 0
const SPI_MEM_BROWN_OUT_INT_ST_V = 0x1
const SPI_MEM_BROWN_OUT_INT_ST_S = 10
const SPI_MEM_AXI_WADDR_ERR_INT_ST_V = 0x1
const SPI_MEM_AXI_WADDR_ERR_INT_ST_S = 9
const SPI_MEM_AXI_WR_FLASH_ERR_INT_ST_V = 0x1
const SPI_MEM_AXI_WR_FLASH_ERR_INT_ST_S = 8
const SPI_MEM_AXI_RADDR_ERR_INT_ST_V = 0x1
const SPI_MEM_AXI_RADDR_ERR_INT_ST_S = 7
const SPI_MEM_PMS_REJECT_INT_ST_V = 0x1
const SPI_MEM_PMS_REJECT_INT_ST_S = 6
const SPI_MEM_ECC_ERR_INT_ST_V = 0x1
const SPI_MEM_ECC_ERR_INT_ST_S = 5
const SPI_MEM_MST_ST_END_INT_ST_V = 0x1
const SPI_MEM_MST_ST_END_INT_ST_S = 4
const SPI_MEM_SLV_ST_END_INT_ST_V = 0x1
const SPI_MEM_SLV_ST_END_INT_ST_S = 3
const SPI_MEM_WPE_END_INT_ST_V = 0x1
const SPI_MEM_WPE_END_INT_ST_S = 2
const SPI_MEM_PES_END_INT_ST_V = 0x1
const SPI_MEM_PES_END_INT_ST_S = 1
const SPI_MEM_PER_END_INT_ST_V = 0x1
const SPI_MEM_PER_END_INT_ST_S = 0
const SPI_MEM_FMEM_HYPERBUS_CA_V = 0x1
const SPI_MEM_FMEM_HYPERBUS_CA_S = 30
const SPI_MEM_FMEM_OCTA_RAM_ADDR_V = 0x1
const SPI_MEM_FMEM_OCTA_RAM_ADDR_S = 29
const SPI_MEM_FMEM_CLK_DIFF_INV_V = 0x1
const SPI_MEM_FMEM_CLK_DIFF_INV_S = 28
const SPI_MEM_FMEM_HYPERBUS_DUMMY_2X_V = 0x1
const SPI_MEM_FMEM_HYPERBUS_DUMMY_2X_S = 27
const SPI_MEM_FMEM_DQS_CA_IN_V = 0x1
const SPI_MEM_FMEM_DQS_CA_IN_S = 26
const SPI_MEM_FMEM_CLK_DIFF_EN_V = 0x1
const SPI_MEM_FMEM_CLK_DIFF_EN_S = 24
const SPI_MEM_FMEM_DDR_DQS_LOOP_V = 0x1
const SPI_MEM_FMEM_DDR_DQS_LOOP_S = 21
const SPI_MEM_FMEM_USR_DDR_DQS_THD = 0x0000007F
const SPI_MEM_FMEM_USR_DDR_DQS_THD_V = 0x7F
const SPI_MEM_FMEM_USR_DDR_DQS_THD_S = 14
const SPI_MEM_FMEM_RX_DDR_MSK_EN_V = 0x1
const SPI_MEM_FMEM_RX_DDR_MSK_EN_S = 13
const SPI_MEM_FMEM_TX_DDR_MSK_EN_V = 0x1
const SPI_MEM_FMEM_TX_DDR_MSK_EN_S = 12
const SPI_MEM_FMEM_OUTMINBYTELEN = 0x0000007F
const SPI_MEM_FMEM_OUTMINBYTELEN_V = 0x7F
const SPI_MEM_FMEM_OUTMINBYTELEN_S = 5
const SPI_MEM_FMEM_DDR_CMD_DIS_V = 0x1
const SPI_MEM_FMEM_DDR_CMD_DIS_S = 4
const SPI_MEM_FMEM_DDR_WDAT_SWP_V = 0x1
const SPI_MEM_FMEM_DDR_WDAT_SWP_S = 3
const SPI_MEM_FMEM_DDR_RDAT_SWP_V = 0x1
const SPI_MEM_FMEM_DDR_RDAT_SWP_S = 2
const SPI_MEM_FMEM_VAR_DUMMY_V = 0x1
const SPI_MEM_FMEM_VAR_DUMMY_S = 1
const SPI_MEM_FMEM_DDR_EN_V = 0x1
const SPI_MEM_FMEM_DDR_EN_S = 0
const SPI_MEM_SMEM_HYPERBUS_CA_V = 0x1
const SPI_MEM_SMEM_HYPERBUS_CA_S = 30
const SPI_MEM_SMEM_OCTA_RAM_ADDR_V = 0x1
const SPI_MEM_SMEM_OCTA_RAM_ADDR_S = 29
const SPI_MEM_SMEM_CLK_DIFF_INV_V = 0x1
const SPI_MEM_SMEM_CLK_DIFF_INV_S = 28
const SPI_MEM_SMEM_HYPERBUS_DUMMY_2X_V = 0x1
const SPI_MEM_SMEM_HYPERBUS_DUMMY_2X_S = 27
const SPI_MEM_SMEM_DQS_CA_IN_V = 0x1
const SPI_MEM_SMEM_DQS_CA_IN_S = 26
const SPI_MEM_SMEM_CLK_DIFF_EN_V = 0x1
const SPI_MEM_SMEM_CLK_DIFF_EN_S = 24
const SPI_MEM_SMEM_DDR_DQS_LOOP_V = 0x1
const SPI_MEM_SMEM_DDR_DQS_LOOP_S = 21
const SPI_MEM_SMEM_USR_DDR_DQS_THD = 0x0000007F
const SPI_MEM_SMEM_USR_DDR_DQS_THD_V = 0x7F
const SPI_MEM_SMEM_USR_DDR_DQS_THD_S = 14
const SPI_MEM_SMEM_RX_DDR_MSK_EN_V = 0x1
const SPI_MEM_SMEM_RX_DDR_MSK_EN_S = 13
const SPI_MEM_SMEM_TX_DDR_MSK_EN_V = 0x1
const SPI_MEM_SMEM_TX_DDR_MSK_EN_S = 12
const SPI_MEM_SMEM_OUTMINBYTELEN = 0x0000007F
const SPI_MEM_SMEM_OUTMINBYTELEN_V = 0x7F
const SPI_MEM_SMEM_OUTMINBYTELEN_S = 5
const SPI_MEM_SMEM_DDR_CMD_DIS_V = 0x1
const SPI_MEM_SMEM_DDR_CMD_DIS_S = 4
const SPI_MEM_SMEM_DDR_WDAT_SWP_V = 0x1
const SPI_MEM_SMEM_DDR_WDAT_SWP_S = 3
const SPI_MEM_SMEM_DDR_RDAT_SWP_V = 0x1
const SPI_MEM_SMEM_DDR_RDAT_SWP_S = 2
const SPI_MEM_SMEM_VAR_DUMMY_V = 0x1
const SPI_MEM_SMEM_VAR_DUMMY_S = 1
const SPI_MEM_SMEM_DDR_EN_V = 0x1
const SPI_MEM_SMEM_DDR_EN_S = 0
const SPI_MEM_FMEM_PMS0_ECC_V = 0x1
const SPI_MEM_FMEM_PMS0_ECC_S = 2
const SPI_MEM_FMEM_PMS0_WR_ATTR_V = 0x1
const SPI_MEM_FMEM_PMS0_WR_ATTR_S = 1
const SPI_MEM_FMEM_PMS0_RD_ATTR_V = 0x1
const SPI_MEM_FMEM_PMS0_RD_ATTR_S = 0
const SPI_MEM_FMEM_PMS1_ECC_V = 0x1
const SPI_MEM_FMEM_PMS1_ECC_S = 2
const SPI_MEM_FMEM_PMS1_WR_ATTR_V = 0x1
const SPI_MEM_FMEM_PMS1_WR_ATTR_S = 1
const SPI_MEM_FMEM_PMS1_RD_ATTR_V = 0x1
const SPI_MEM_FMEM_PMS1_RD_ATTR_S = 0
const SPI_MEM_FMEM_PMS2_ECC_V = 0x1
const SPI_MEM_FMEM_PMS2_ECC_S = 2
const SPI_MEM_FMEM_PMS2_WR_ATTR_V = 0x1
const SPI_MEM_FMEM_PMS2_WR_ATTR_S = 1
const SPI_MEM_FMEM_PMS2_RD_ATTR_V = 0x1
const SPI_MEM_FMEM_PMS2_RD_ATTR_S = 0
const SPI_MEM_FMEM_PMS3_ECC_V = 0x1
const SPI_MEM_FMEM_PMS3_ECC_S = 2
const SPI_MEM_FMEM_PMS3_WR_ATTR_V = 0x1
const SPI_MEM_FMEM_PMS3_WR_ATTR_S = 1
const SPI_MEM_FMEM_PMS3_RD_ATTR_V = 0x1
const SPI_MEM_FMEM_PMS3_RD_ATTR_S = 0
const SPI_MEM_FMEM_PMS0_ADDR_S = 0x03FFFFFF
const SPI_MEM_FMEM_PMS0_ADDR_S_V = 0x3FFFFFF
const SPI_MEM_FMEM_PMS0_ADDR_S_S = 0
const SPI_MEM_FMEM_PMS1_ADDR_S = 0x03FFFFFF
const SPI_MEM_FMEM_PMS1_ADDR_S_V = 0x3FFFFFF
const SPI_MEM_FMEM_PMS1_ADDR_S_S = 0
const SPI_MEM_FMEM_PMS2_ADDR_S = 0x03FFFFFF
const SPI_MEM_FMEM_PMS2_ADDR_S_V = 0x3FFFFFF
const SPI_MEM_FMEM_PMS2_ADDR_S_S = 0
const SPI_MEM_FMEM_PMS3_ADDR_S = 0x03FFFFFF
const SPI_MEM_FMEM_PMS3_ADDR_S_V = 0x3FFFFFF
const SPI_MEM_FMEM_PMS3_ADDR_S_S = 0
const SPI_MEM_FMEM_PMS0_SIZE = 0x00003FFF
const SPI_MEM_FMEM_PMS0_SIZE_V = 0x3FFF
const SPI_MEM_FMEM_PMS0_SIZE_S = 0
const SPI_MEM_FMEM_PMS1_SIZE = 0x00003FFF
const SPI_MEM_FMEM_PMS1_SIZE_V = 0x3FFF
const SPI_MEM_FMEM_PMS1_SIZE_S = 0
const SPI_MEM_FMEM_PMS2_SIZE = 0x00003FFF
const SPI_MEM_FMEM_PMS2_SIZE_V = 0x3FFF
const SPI_MEM_FMEM_PMS2_SIZE_S = 0
const SPI_MEM_FMEM_PMS3_SIZE = 0x00003FFF
const SPI_MEM_FMEM_PMS3_SIZE_V = 0x3FFF
const SPI_MEM_FMEM_PMS3_SIZE_S = 0
const SPI_MEM_SMEM_PMS0_ECC_V = 0x1
const SPI_MEM_SMEM_PMS0_ECC_S = 2
const SPI_MEM_SMEM_PMS0_WR_ATTR_V = 0x1
const SPI_MEM_SMEM_PMS0_WR_ATTR_S = 1
const SPI_MEM_SMEM_PMS0_RD_ATTR_V = 0x1
const SPI_MEM_SMEM_PMS0_RD_ATTR_S = 0
const SPI_MEM_SMEM_PMS1_ECC_V = 0x1
const SPI_MEM_SMEM_PMS1_ECC_S = 2
const SPI_MEM_SMEM_PMS1_WR_ATTR_V = 0x1
const SPI_MEM_SMEM_PMS1_WR_ATTR_S = 1
const SPI_MEM_SMEM_PMS1_RD_ATTR_V = 0x1
const SPI_MEM_SMEM_PMS1_RD_ATTR_S = 0
const SPI_MEM_SMEM_PMS2_ECC_V = 0x1
const SPI_MEM_SMEM_PMS2_ECC_S = 2
const SPI_MEM_SMEM_PMS2_WR_ATTR_V = 0x1
const SPI_MEM_SMEM_PMS2_WR_ATTR_S = 1
const SPI_MEM_SMEM_PMS2_RD_ATTR_V = 0x1
const SPI_MEM_SMEM_PMS2_RD_ATTR_S = 0
const SPI_MEM_SMEM_PMS3_ECC_V = 0x1
const SPI_MEM_SMEM_PMS3_ECC_S = 2
const SPI_MEM_SMEM_PMS3_WR_ATTR_V = 0x1
const SPI_MEM_SMEM_PMS3_WR_ATTR_S = 1
const SPI_MEM_SMEM_PMS3_RD_ATTR_V = 0x1
const SPI_MEM_SMEM_PMS3_RD_ATTR_S = 0
const SPI_MEM_SMEM_PMS0_ADDR_S = 0x03FFFFFF
const SPI_MEM_SMEM_PMS0_ADDR_S_V = 0x3FFFFFF
const SPI_MEM_SMEM_PMS0_ADDR_S_S = 0
const SPI_MEM_SMEM_PMS1_ADDR_S = 0x03FFFFFF
const SPI_MEM_SMEM_PMS1_ADDR_S_V = 0x3FFFFFF
const SPI_MEM_SMEM_PMS1_ADDR_S_S = 0
const SPI_MEM_SMEM_PMS2_ADDR_S = 0x03FFFFFF
const SPI_MEM_SMEM_PMS2_ADDR_S_V = 0x3FFFFFF
const SPI_MEM_SMEM_PMS2_ADDR_S_S = 0
const SPI_MEM_SMEM_PMS3_ADDR_S = 0x03FFFFFF
const SPI_MEM_SMEM_PMS3_ADDR_S_V = 0x3FFFFFF
const SPI_MEM_SMEM_PMS3_ADDR_S_S = 0
const SPI_MEM_SMEM_PMS0_SIZE = 0x00003FFF
const SPI_MEM_SMEM_PMS0_SIZE_V = 0x3FFF
const SPI_MEM_SMEM_PMS0_SIZE_S = 0
const SPI_MEM_SMEM_PMS1_SIZE = 0x00003FFF
const SPI_MEM_SMEM_PMS1_SIZE_V = 0x3FFF
const SPI_MEM_SMEM_PMS1_SIZE_S = 0
const SPI_MEM_SMEM_PMS2_SIZE = 0x00003FFF
const SPI_MEM_SMEM_PMS2_SIZE_V = 0x3FFF
const SPI_MEM_SMEM_PMS2_SIZE_S = 0
const SPI_MEM_SMEM_PMS3_SIZE = 0x00003FFF
const SPI_MEM_SMEM_PMS3_SIZE_V = 0x3FFF
const SPI_MEM_SMEM_PMS3_SIZE_S = 0
const SPI_MEM_PMS_IVD_V = 0x1
const SPI_MEM_PMS_IVD_S = 31
const SPI_MEM_PMS_MULTI_HIT_V = 0x1
const SPI_MEM_PMS_MULTI_HIT_S = 30
const SPI_MEM_PMS_ST_V = 0x1
const SPI_MEM_PMS_ST_S = 29
const SPI_MEM_PMS_LD_V = 0x1
const SPI_MEM_PMS_LD_S = 28
const SPI_MEM_PM_EN_V = 0x1
const SPI_MEM_PM_EN_S = 26
const SPI_MEM_REJECT_ADDR = 0x03FFFFFF
const SPI_MEM_REJECT_ADDR_V = 0x3FFFFFF
const SPI_MEM_REJECT_ADDR_S = 0
const SPI_MEM_ECC_ERR_BITS = 0x0000007F
const SPI_MEM_ECC_ERR_BITS_V = 0x7F
const SPI_MEM_ECC_ERR_BITS_S = 25
const SPI_MEM_ECC_CONTINUE_RECORD_ERR_EN_V = 0x1
const SPI_MEM_ECC_CONTINUE_RECORD_ERR_EN_S = 24
const SPI_MEM_USR_ECC_ADDR_EN_V = 0x1
const SPI_MEM_USR_ECC_ADDR_EN_S = 21
const SPI_MEM_FMEM_ECC_ADDR_EN_V = 0x1
const SPI_MEM_FMEM_ECC_ADDR_EN_S = 20
const SPI_MEM_FMEM_PAGE_SIZE = 0x00000003
const SPI_MEM_FMEM_PAGE_SIZE_V = 0x3
const SPI_MEM_FMEM_PAGE_SIZE_S = 18
const SPI_MEM_FMEM_ECC_ERR_INT_EN_V = 0x1
const SPI_MEM_FMEM_ECC_ERR_INT_EN_S = 17
const SPI_MEM_FMEM_ECC_ERR_INT_NUM = 0x0000003F
const SPI_MEM_FMEM_ECC_ERR_INT_NUM_V = 0x3F
const SPI_MEM_FMEM_ECC_ERR_INT_NUM_S = 11
const SPI_MEM_ECC_ERR_CNT = 0x0000003F
const SPI_MEM_ECC_ERR_CNT_V = 0x3F
const SPI_MEM_ECC_ERR_CNT_S = 26
const SPI_MEM_ECC_ERR_ADDR = 0x03FFFFFF
const SPI_MEM_ECC_ERR_ADDR_V = 0x3FFFFFF
const SPI_MEM_ECC_ERR_ADDR_S = 0
const SPI_MEM_ALL_AXI_TRANS_AFIFO_EMPTY_V = 0x1
const SPI_MEM_ALL_AXI_TRANS_AFIFO_EMPTY_S = 31
const SPI_MEM_WBLEN_AFIFO_REMPTY_V = 0x1
const SPI_MEM_WBLEN_AFIFO_REMPTY_S = 30
const SPI_MEM_WDATA_AFIFO_REMPTY_V = 0x1
const SPI_MEM_WDATA_AFIFO_REMPTY_S = 29
const SPI_MEM_RADDR_AFIFO_REMPTY_V = 0x1
const SPI_MEM_RADDR_AFIFO_REMPTY_S = 28
const SPI_MEM_RDATA_AFIFO_REMPTY_V = 0x1
const SPI_MEM_RDATA_AFIFO_REMPTY_S = 27
const SPI_MEM_ALL_FIFO_EMPTY_V = 0x1
const SPI_MEM_ALL_FIFO_EMPTY_S = 26
const SPI_MEM_AXI_ERR_ADDR = 0x03FFFFFF
const SPI_MEM_AXI_ERR_ADDR_V = 0x3FFFFFF
const SPI_MEM_AXI_ERR_ADDR_S = 0
const SPI_MEM_SMEM_ECC_ADDR_EN_V = 0x1
const SPI_MEM_SMEM_ECC_ADDR_EN_S = 20
const SPI_MEM_SMEM_PAGE_SIZE = 0x00000003
const SPI_MEM_SMEM_PAGE_SIZE_V = 0x3
const SPI_MEM_SMEM_PAGE_SIZE_S = 18
const SPI_MEM_SMEM_ECC_ERR_INT_EN_V = 0x1
const SPI_MEM_SMEM_ECC_ERR_INT_EN_S = 17
const SPI_MEM_TIMING_CALI_UPDATE_V = 0x1
const SPI_MEM_TIMING_CALI_UPDATE_S = 6
const SPI_MEM_DLL_TIMING_CALI_V = 0x1
const SPI_MEM_DLL_TIMING_CALI_S = 5
const SPI_MEM_EXTRA_DUMMY_CYCLELEN = 0x00000007
const SPI_MEM_EXTRA_DUMMY_CYCLELEN_V = 0x7
const SPI_MEM_EXTRA_DUMMY_CYCLELEN_S = 2
const SPI_MEM_TIMING_CALI_V = 0x1
const SPI_MEM_TIMING_CALI_S = 1
const SPI_MEM_TIMING_CLK_ENA_V = 0x1
const SPI_MEM_TIMING_CLK_ENA_S = 0
const SPI_MEM_DINS_MODE = 0x00000007
const SPI_MEM_DINS_MODE_V = 0x7
const SPI_MEM_DINS_MODE_S = 24
const SPI_MEM_DIN7_MODE = 0x00000007
const SPI_MEM_DIN7_MODE_V = 0x7
const SPI_MEM_DIN7_MODE_S = 21
const SPI_MEM_DIN6_MODE = 0x00000007
const SPI_MEM_DIN6_MODE_V = 0x7
const SPI_MEM_DIN6_MODE_S = 18
const SPI_MEM_DIN5_MODE = 0x00000007
const SPI_MEM_DIN5_MODE_V = 0x7
const SPI_MEM_DIN5_MODE_S = 15
const SPI_MEM_DIN4_MODE = 0x00000007
const SPI_MEM_DIN4_MODE_V = 0x7
const SPI_MEM_DIN4_MODE_S = 12
const SPI_MEM_DIN3_MODE = 0x00000007
const SPI_MEM_DIN3_MODE_V = 0x7
const SPI_MEM_DIN3_MODE_S = 9
const SPI_MEM_DIN2_MODE = 0x00000007
const SPI_MEM_DIN2_MODE_V = 0x7
const SPI_MEM_DIN2_MODE_S = 6
const SPI_MEM_DIN1_MODE = 0x00000007
const SPI_MEM_DIN1_MODE_V = 0x7
const SPI_MEM_DIN1_MODE_S = 3
const SPI_MEM_DIN0_MODE = 0x00000007
const SPI_MEM_DIN0_MODE_V = 0x7
const SPI_MEM_DIN0_MODE_S = 0
const SPI_MEM_DINS_NUM = 0x00000003
const SPI_MEM_DINS_NUM_V = 0x3
const SPI_MEM_DINS_NUM_S = 16
const SPI_MEM_DIN7_NUM = 0x00000003
const SPI_MEM_DIN7_NUM_V = 0x3
const SPI_MEM_DIN7_NUM_S = 14
const SPI_MEM_DIN6_NUM = 0x00000003
const SPI_MEM_DIN6_NUM_V = 0x3
const SPI_MEM_DIN6_NUM_S = 12
const SPI_MEM_DIN5_NUM = 0x00000003
const SPI_MEM_DIN5_NUM_V = 0x3
const SPI_MEM_DIN5_NUM_S = 10
const SPI_MEM_DIN4_NUM = 0x00000003
const SPI_MEM_DIN4_NUM_V = 0x3
const SPI_MEM_DIN4_NUM_S = 8
const SPI_MEM_DIN3_NUM = 0x00000003
const SPI_MEM_DIN3_NUM_V = 0x3
const SPI_MEM_DIN3_NUM_S = 6
const SPI_MEM_DIN2_NUM = 0x00000003
const SPI_MEM_DIN2_NUM_V = 0x3
const SPI_MEM_DIN2_NUM_S = 4
const SPI_MEM_DIN1_NUM = 0x00000003
const SPI_MEM_DIN1_NUM_V = 0x3
const SPI_MEM_DIN1_NUM_S = 2
const SPI_MEM_DIN0_NUM = 0x00000003
const SPI_MEM_DIN0_NUM_V = 0x3
const SPI_MEM_DIN0_NUM_S = 0
const SPI_MEM_DOUTS_MODE_V = 0x1
const SPI_MEM_DOUTS_MODE_S = 8
const SPI_MEM_DOUT7_MODE_V = 0x1
const SPI_MEM_DOUT7_MODE_S = 7
const SPI_MEM_DOUT6_MODE_V = 0x1
const SPI_MEM_DOUT6_MODE_S = 6
const SPI_MEM_DOUT5_MODE_V = 0x1
const SPI_MEM_DOUT5_MODE_S = 5
const SPI_MEM_DOUT4_MODE_V = 0x1
const SPI_MEM_DOUT4_MODE_S = 4
const SPI_MEM_DOUT3_MODE_V = 0x1
const SPI_MEM_DOUT3_MODE_S = 3
const SPI_MEM_DOUT2_MODE_V = 0x1
const SPI_MEM_DOUT2_MODE_S = 2
const SPI_MEM_DOUT1_MODE_V = 0x1
const SPI_MEM_DOUT1_MODE_S = 1
const SPI_MEM_DOUT0_MODE_V = 0x1
const SPI_MEM_DOUT0_MODE_S = 0
const SPI_MEM_SMEM_DLL_TIMING_CALI_V = 0x1
const SPI_MEM_SMEM_DLL_TIMING_CALI_S = 5
const SPI_MEM_SMEM_EXTRA_DUMMY_CYCLELEN = 0x00000007
const SPI_MEM_SMEM_EXTRA_DUMMY_CYCLELEN_V = 0x7
const SPI_MEM_SMEM_EXTRA_DUMMY_CYCLELEN_S = 2
const SPI_MEM_SMEM_TIMING_CALI_V = 0x1
const SPI_MEM_SMEM_TIMING_CALI_S = 1
const SPI_MEM_SMEM_TIMING_CLK_ENA_V = 0x1
const SPI_MEM_SMEM_TIMING_CLK_ENA_S = 0
const SPI_MEM_SMEM_DINS_MODE = 0x00000007
const SPI_MEM_SMEM_DINS_MODE_V = 0x7
const SPI_MEM_SMEM_DINS_MODE_S = 24
const SPI_MEM_SMEM_DIN7_MODE = 0x00000007
const SPI_MEM_SMEM_DIN7_MODE_V = 0x7
const SPI_MEM_SMEM_DIN7_MODE_S = 21
const SPI_MEM_SMEM_DIN6_MODE = 0x00000007
const SPI_MEM_SMEM_DIN6_MODE_V = 0x7
const SPI_MEM_SMEM_DIN6_MODE_S = 18
const SPI_MEM_SMEM_DIN5_MODE = 0x00000007
const SPI_MEM_SMEM_DIN5_MODE_V = 0x7
const SPI_MEM_SMEM_DIN5_MODE_S = 15
const SPI_MEM_SMEM_DIN4_MODE = 0x00000007
const SPI_MEM_SMEM_DIN4_MODE_V = 0x7
const SPI_MEM_SMEM_DIN4_MODE_S = 12
const SPI_MEM_SMEM_DIN3_MODE = 0x00000007
const SPI_MEM_SMEM_DIN3_MODE_V = 0x7
const SPI_MEM_SMEM_DIN3_MODE_S = 9
const SPI_MEM_SMEM_DIN2_MODE = 0x00000007
const SPI_MEM_SMEM_DIN2_MODE_V = 0x7
const SPI_MEM_SMEM_DIN2_MODE_S = 6
const SPI_MEM_SMEM_DIN1_MODE = 0x00000007
const SPI_MEM_SMEM_DIN1_MODE_V = 0x7
const SPI_MEM_SMEM_DIN1_MODE_S = 3
const SPI_MEM_SMEM_DIN0_MODE = 0x00000007
const SPI_MEM_SMEM_DIN0_MODE_V = 0x7
const SPI_MEM_SMEM_DIN0_MODE_S = 0
const SPI_MEM_SMEM_DINS_NUM = 0x00000003
const SPI_MEM_SMEM_DINS_NUM_V = 0x3
const SPI_MEM_SMEM_DINS_NUM_S = 16
const SPI_MEM_SMEM_DIN7_NUM = 0x00000003
const SPI_MEM_SMEM_DIN7_NUM_V = 0x3
const SPI_MEM_SMEM_DIN7_NUM_S = 14
const SPI_MEM_SMEM_DIN6_NUM = 0x00000003
const SPI_MEM_SMEM_DIN6_NUM_V = 0x3
const SPI_MEM_SMEM_DIN6_NUM_S = 12
const SPI_MEM_SMEM_DIN5_NUM = 0x00000003
const SPI_MEM_SMEM_DIN5_NUM_V = 0x3
const SPI_MEM_SMEM_DIN5_NUM_S = 10
const SPI_MEM_SMEM_DIN4_NUM = 0x00000003
const SPI_MEM_SMEM_DIN4_NUM_V = 0x3
const SPI_MEM_SMEM_DIN4_NUM_S = 8
const SPI_MEM_SMEM_DIN3_NUM = 0x00000003
const SPI_MEM_SMEM_DIN3_NUM_V = 0x3
const SPI_MEM_SMEM_DIN3_NUM_S = 6
const SPI_MEM_SMEM_DIN2_NUM = 0x00000003
const SPI_MEM_SMEM_DIN2_NUM_V = 0x3
const SPI_MEM_SMEM_DIN2_NUM_S = 4
const SPI_MEM_SMEM_DIN1_NUM = 0x00000003
const SPI_MEM_SMEM_DIN1_NUM_V = 0x3
const SPI_MEM_SMEM_DIN1_NUM_S = 2
const SPI_MEM_SMEM_DIN0_NUM = 0x00000003
const SPI_MEM_SMEM_DIN0_NUM_V = 0x3
const SPI_MEM_SMEM_DIN0_NUM_S = 0
const SPI_MEM_SMEM_DOUTS_MODE_V = 0x1
const SPI_MEM_SMEM_DOUTS_MODE_S = 8
const SPI_MEM_SMEM_DOUT7_MODE_V = 0x1
const SPI_MEM_SMEM_DOUT7_MODE_S = 7
const SPI_MEM_SMEM_DOUT6_MODE_V = 0x1
const SPI_MEM_SMEM_DOUT6_MODE_S = 6
const SPI_MEM_SMEM_DOUT5_MODE_V = 0x1
const SPI_MEM_SMEM_DOUT5_MODE_S = 5
const SPI_MEM_SMEM_DOUT4_MODE_V = 0x1
const SPI_MEM_SMEM_DOUT4_MODE_S = 4
const SPI_MEM_SMEM_DOUT3_MODE_V = 0x1
const SPI_MEM_SMEM_DOUT3_MODE_S = 3
const SPI_MEM_SMEM_DOUT2_MODE_V = 0x1
const SPI_MEM_SMEM_DOUT2_MODE_S = 2
const SPI_MEM_SMEM_DOUT1_MODE_V = 0x1
const SPI_MEM_SMEM_DOUT1_MODE_S = 1
const SPI_MEM_SMEM_DOUT0_MODE_V = 0x1
const SPI_MEM_SMEM_DOUT0_MODE_S = 0
const SPI_MEM_SMEM_SPLIT_TRANS_EN_V = 0x1
const SPI_MEM_SMEM_SPLIT_TRANS_EN_S = 31
const SPI_MEM_SMEM_CS_HOLD_DELAY = 0x0000003F
const SPI_MEM_SMEM_CS_HOLD_DELAY_V = 0x3F
const SPI_MEM_SMEM_CS_HOLD_DELAY_S = 25
const SPI_MEM_SMEM_ECC_16TO18_BYTE_EN_V = 0x1
const SPI_MEM_SMEM_ECC_16TO18_BYTE_EN_S = 16
const SPI_MEM_SMEM_ECC_SKIP_PAGE_CORNER_V = 0x1
const SPI_MEM_SMEM_ECC_SKIP_PAGE_CORNER_S = 15
const SPI_MEM_SMEM_ECC_CS_HOLD_TIME = 0x00000007
const SPI_MEM_SMEM_ECC_CS_HOLD_TIME_V = 0x7
const SPI_MEM_SMEM_ECC_CS_HOLD_TIME_S = 12
const SPI_MEM_SMEM_CS_HOLD_TIME = 0x0000001F
const SPI_MEM_SMEM_CS_HOLD_TIME_V = 0x1F
const SPI_MEM_SMEM_CS_HOLD_TIME_S = 7
const SPI_MEM_SMEM_CS_SETUP_TIME = 0x0000001F
const SPI_MEM_SMEM_CS_SETUP_TIME_V = 0x1F
const SPI_MEM_SMEM_CS_SETUP_TIME_S = 2
const SPI_MEM_SMEM_CS_HOLD_V = 0x1
const SPI_MEM_SMEM_CS_HOLD_S = 1
const SPI_MEM_SMEM_CS_SETUP_V = 0x1
const SPI_MEM_SMEM_CS_SETUP_S = 0
const SPI_MEM_CLK_EN_V = 0x1
const SPI_MEM_CLK_EN_S = 0
const SPI_MEM_MMU_ITEM_CONTENT = 0xFFFFFFFF
const SPI_MEM_MMU_ITEM_CONTENT_V = 0xFFFFFFFF
const SPI_MEM_MMU_ITEM_CONTENT_S = 0
const SPI_MEM_MMU_ITEM_INDEX = 0xFFFFFFFF
const SPI_MEM_MMU_ITEM_INDEX_V = 0xFFFFFFFF
const SPI_MEM_MMU_ITEM_INDEX_S = 0
const SPI_MEM_RDN_RESULT_V = 0x1
const SPI_MEM_RDN_RESULT_S = 31
const SPI_MEM_RDN_ENA_V = 0x1
const SPI_MEM_RDN_ENA_S = 30
const SPI_MEM_AUX_CTRL = 0x00003FFF
const SPI_MEM_AUX_CTRL_V = 0x3FFF
const SPI_MEM_AUX_CTRL_S = 16
const SPI_MEM_MMU_PAGE_SIZE = 0x00000003
const SPI_MEM_MMU_PAGE_SIZE_V = 0x3
const SPI_MEM_MMU_PAGE_SIZE_S = 3
const SPI_MEM_MMU_MEM_FORCE_PU_V = 0x1
const SPI_MEM_MMU_MEM_FORCE_PU_S = 2
const SPI_MEM_MMU_MEM_FORCE_PD_V = 0x1
const SPI_MEM_MMU_MEM_FORCE_PD_S = 1
const SPI_MEM_MMU_MEM_FORCE_ON_V = 0x1
const SPI_MEM_MMU_MEM_FORCE_ON_S = 0
const SPI_MEM_REGISTERRND_ECO_HIGH = 0xFFFFFFFF
const SPI_MEM_REGISTERRND_ECO_HIGH_V = 0xFFFFFFFF
const SPI_MEM_REGISTERRND_ECO_HIGH_S = 0
const SPI_MEM_REGISTERRND_ECO_LOW = 0xFFFFFFFF
const SPI_MEM_REGISTERRND_ECO_LOW_V = 0xFFFFFFFF
const SPI_MEM_REGISTERRND_ECO_LOW_S = 0
const SPI_MEM_DATE = 0x0FFFFFFF
const SPI_MEM_DATE_V = 0xFFFFFFF
const SPI_MEM_DATE_S = 0
