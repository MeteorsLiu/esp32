package soc

import _ "unsafe"

const SPI_USR_V = 0x1
const SPI_USR_S = 24
const SPI_CONF_BITLEN = 0x007FFFFF
const SPI_CONF_BITLEN_V = 0x7FFFFF
const SPI_CONF_BITLEN_S = 0
const SPI_USR_ADDR_VALUE = 0xFFFFFFFF
const SPI_USR_ADDR_VALUE_V = 0xFFFFFFFF
const SPI_USR_ADDR_VALUE_S = 0
const SPI_WR_BIT_ORDER_V = 0x1
const SPI_WR_BIT_ORDER_S = 26
const SPI_RD_BIT_ORDER_V = 0x1
const SPI_RD_BIT_ORDER_S = 25
const SPI_WP_REG_V = 0x1
const SPI_WP_REG_S = 21
const SPI_D_POL_V = 0x1
const SPI_D_POL_S = 19
const SPI_Q_POL_V = 0x1
const SPI_Q_POL_S = 18
const SPI_FREAD_OCT_V = 0x1
const SPI_FREAD_OCT_S = 16
const SPI_FREAD_QUAD_V = 0x1
const SPI_FREAD_QUAD_S = 15
const SPI_FREAD_DUAL_V = 0x1
const SPI_FREAD_DUAL_S = 14
const SPI_FCMD_OCT_V = 0x1
const SPI_FCMD_OCT_S = 10
const SPI_FCMD_QUAD_V = 0x1
const SPI_FCMD_QUAD_S = 9
const SPI_FCMD_DUAL_V = 0x1
const SPI_FCMD_DUAL_S = 8
const SPI_FADDR_OCT_V = 0x1
const SPI_FADDR_OCT_S = 7
const SPI_FADDR_QUAD_V = 0x1
const SPI_FADDR_QUAD_S = 6
const SPI_FADDR_DUAL_V = 0x1
const SPI_FADDR_DUAL_S = 5
const SPI_DUMMY_OUT_V = 0x1
const SPI_DUMMY_OUT_S = 3
const SPI_EXT_HOLD_EN_V = 0x1
const SPI_EXT_HOLD_EN_S = 2
const SPI_CS_HOLD_DELAY = 0x0000003F
const SPI_CS_HOLD_DELAY_V = 0x3F
const SPI_CS_HOLD_DELAY_S = 14
const SPI_W16_17_WR_ENA_V = 0x1
const SPI_W16_17_WR_ENA_S = 4
const SPI_RSCK_DATA_OUT_V = 0x1
const SPI_RSCK_DATA_OUT_S = 3
const SPI_CLK_MODE_13_V = 0x1
const SPI_CLK_MODE_13_S = 2
const SPI_CLK_MODE = 0x00000003
const SPI_CLK_MODE_V = 0x3
const SPI_CLK_MODE_S = 0
const SPI_CS_DELAY_NUM = 0x00000003
const SPI_CS_DELAY_NUM_V = 0x3
const SPI_CS_DELAY_NUM_S = 29
const SPI_CS_DELAY_MODE = 0x00000007
const SPI_CS_DELAY_MODE_V = 0x7
const SPI_CS_DELAY_MODE_S = 26
const SPI_CS_HOLD_TIME = 0x00001FFF
const SPI_CS_HOLD_TIME_V = 0x1FFF
const SPI_CS_HOLD_TIME_S = 13
const SPI_CS_SETUP_TIME = 0x00001FFF
const SPI_CS_SETUP_TIME_V = 0x1FFF
const SPI_CS_SETUP_TIME_S = 0
const SPI_CLK_EQU_SYSCLK_V = 0x1
const SPI_CLK_EQU_SYSCLK_S = 31
const SPI_CLKDIV_PRE = 0x00001FFF
const SPI_CLKDIV_PRE_V = 0x1FFF
const SPI_CLKDIV_PRE_S = 18
const SPI_CLKCNT_N = 0x0000003F
const SPI_CLKCNT_N_V = 0x3F
const SPI_CLKCNT_N_S = 12
const SPI_CLKCNT_H = 0x0000003F
const SPI_CLKCNT_H_V = 0x3F
const SPI_CLKCNT_H_S = 6
const SPI_CLKCNT_L = 0x0000003F
const SPI_CLKCNT_L_V = 0x3F
const SPI_CLKCNT_L_S = 0
const SPI_USR_COMMAND_V = 0x1
const SPI_USR_COMMAND_S = 31
const SPI_USR_ADDR_V = 0x1
const SPI_USR_ADDR_S = 30
const SPI_USR_DUMMY_V = 0x1
const SPI_USR_DUMMY_S = 29
const SPI_USR_MISO_V = 0x1
const SPI_USR_MISO_S = 28
const SPI_USR_MOSI_V = 0x1
const SPI_USR_MOSI_S = 27
const SPI_USR_DUMMY_IDLE_V = 0x1
const SPI_USR_DUMMY_IDLE_S = 26
const SPI_USR_MOSI_HIGHPART_V = 0x1
const SPI_USR_MOSI_HIGHPART_S = 25
const SPI_USR_MISO_HIGHPART_V = 0x1
const SPI_USR_MISO_HIGHPART_S = 24
const SPI_USR_PREP_HOLD_V = 0x1
const SPI_USR_PREP_HOLD_S = 23
const SPI_USR_CMD_HOLD_V = 0x1
const SPI_USR_CMD_HOLD_S = 22
const SPI_USR_ADDR_HOLD_V = 0x1
const SPI_USR_ADDR_HOLD_S = 21
const SPI_USR_DUMMY_HOLD_V = 0x1
const SPI_USR_DUMMY_HOLD_S = 20
const SPI_USR_DIN_HOLD_V = 0x1
const SPI_USR_DIN_HOLD_S = 19
const SPI_USR_DOUT_HOLD_V = 0x1
const SPI_USR_DOUT_HOLD_S = 18
const SPI_USR_HOLD_POL_V = 0x1
const SPI_USR_HOLD_POL_S = 17
const SPI_SIO_V = 0x1
const SPI_SIO_S = 16
const SPI_USR_CONF_NXT_V = 0x1
const SPI_USR_CONF_NXT_S = 15
const SPI_FWRITE_OCT_V = 0x1
const SPI_FWRITE_OCT_S = 14
const SPI_FWRITE_QUAD_V = 0x1
const SPI_FWRITE_QUAD_S = 13
const SPI_FWRITE_DUAL_V = 0x1
const SPI_FWRITE_DUAL_S = 12
const SPI_WR_BYTE_ORDER_V = 0x1
const SPI_WR_BYTE_ORDER_S = 11
const SPI_RD_BYTE_ORDER_V = 0x1
const SPI_RD_BYTE_ORDER_S = 10
const SPI_CK_OUT_EDGE_V = 0x1
const SPI_CK_OUT_EDGE_S = 9
const SPI_RSCK_I_EDGE_V = 0x1
const SPI_RSCK_I_EDGE_S = 8
const SPI_CS_SETUP_V = 0x1
const SPI_CS_SETUP_S = 7
const SPI_CS_HOLD_V = 0x1
const SPI_CS_HOLD_S = 6
const SPI_TSCK_I_EDGE_V = 0x1
const SPI_TSCK_I_EDGE_S = 5
const SPI_OPI_MODE_V = 0x1
const SPI_OPI_MODE_S = 4
const SPI_QPI_MODE_V = 0x1
const SPI_QPI_MODE_S = 3
const SPI_DOUTDIN_V = 0x1
const SPI_DOUTDIN_S = 0
const SPI_USR_ADDR_BITLEN = 0x0000001F
const SPI_USR_ADDR_BITLEN_V = 0x1F
const SPI_USR_ADDR_BITLEN_S = 27
const SPI_USR_DUMMY_CYCLELEN = 0x000000FF
const SPI_USR_DUMMY_CYCLELEN_V = 0xFF
const SPI_USR_DUMMY_CYCLELEN_S = 0
const SPI_USR_COMMAND_BITLEN = 0x0000000F
const SPI_USR_COMMAND_BITLEN_V = 0xF
const SPI_USR_COMMAND_BITLEN_S = 28
const SPI_USR_COMMAND_VALUE = 0x0000FFFF
const SPI_USR_COMMAND_VALUE_V = 0xFFFF
const SPI_USR_COMMAND_VALUE_S = 0
const SPI_USR_MOSI_DBITLEN = 0x007FFFFF
const SPI_USR_MOSI_DBITLEN_V = 0x7FFFFF
const SPI_USR_MOSI_DBITLEN_S = 0
const SPI_USR_MISO_DBITLEN = 0x007FFFFF
const SPI_USR_MISO_DBITLEN_V = 0x7FFFFF
const SPI_USR_MISO_DBITLEN_S = 0
const SPI_QUAD_DIN_PIN_SWAP_V = 0x1
const SPI_QUAD_DIN_PIN_SWAP_S = 31
const SPI_CS_KEEP_ACTIVE_V = 0x1
const SPI_CS_KEEP_ACTIVE_S = 30
const SPI_CK_IDLE_EDGE_V = 0x1
const SPI_CK_IDLE_EDGE_S = 29
const SPI_CD_IDLE_EDGE_V = 0x1
const SPI_CD_IDLE_EDGE_S = 26
const SPI_CD_CMD_SET_V = 0x1
const SPI_CD_CMD_SET_S = 25
const SPI_DQS_IDLE_EDGE_V = 0x1
const SPI_DQS_IDLE_EDGE_S = 24
const SPI_SLAVE_CS_POL_V = 0x1
const SPI_SLAVE_CS_POL_S = 23
const SPI_CD_ADDR_SET_V = 0x1
const SPI_CD_ADDR_SET_S = 22
const SPI_CD_DUMMY_SET_V = 0x1
const SPI_CD_DUMMY_SET_S = 21
const SPI_CD_DATA_SET_V = 0x1
const SPI_CD_DATA_SET_S = 20
const SPI_CMD_DTR_EN_V = 0x1
const SPI_CMD_DTR_EN_S = 19
const SPI_ADDR_DTR_EN_V = 0x1
const SPI_ADDR_DTR_EN_S = 18
const SPI_DATA_DTR_EN_V = 0x1
const SPI_DATA_DTR_EN_S = 17
const SPI_CLK_DATA_DTR_EN_V = 0x1
const SPI_CLK_DATA_DTR_EN_S = 16
const SPI_MASTER_CS_POL = 0x0000003F
const SPI_MASTER_CS_POL_V = 0x3F
const SPI_MASTER_CS_POL_S = 7
const SPI_CK_DIS_V = 0x1
const SPI_CK_DIS_S = 6
const SPI_CS5_DIS_V = 0x1
const SPI_CS5_DIS_S = 5
const SPI_CS4_DIS_V = 0x1
const SPI_CS4_DIS_S = 4
const SPI_CS3_DIS_V = 0x1
const SPI_CS3_DIS_S = 3
const SPI_CS2_DIS_V = 0x1
const SPI_CS2_DIS_S = 2
const SPI_CS1_DIS_V = 0x1
const SPI_CS1_DIS_S = 1
const SPI_CS0_DIS_V = 0x1
const SPI_CS0_DIS_S = 0
const SPI_SOFT_RESET_V = 0x1
const SPI_SOFT_RESET_S = 31
const SPI_SLAVE_MODE_V = 0x1
const SPI_SLAVE_MODE_S = 30
const SPI_TRANS_DONE_AUTO_CLR_EN_V = 0x1
const SPI_TRANS_DONE_AUTO_CLR_EN_S = 29
const SPI_TRANS_CNT = 0x0000000F
const SPI_TRANS_CNT_V = 0xF
const SPI_TRANS_CNT_S = 23
const SPI_SEG_MAGIC_ERR_INT_EN_V = 0x1
const SPI_SEG_MAGIC_ERR_INT_EN_S = 11
const SPI_INT_DMA_SEG_TRANS_EN_V = 0x1
const SPI_INT_DMA_SEG_TRANS_EN_S = 10
const SPI_INT_TRANS_DONE_EN_V = 0x1
const SPI_INT_TRANS_DONE_EN_S = 9
const SPI_INT_WR_DMA_DONE_EN_V = 0x1
const SPI_INT_WR_DMA_DONE_EN_S = 8
const SPI_INT_RD_DMA_DONE_EN_V = 0x1
const SPI_INT_RD_DMA_DONE_EN_S = 7
const SPI_INT_WR_BUF_DONE_EN_V = 0x1
const SPI_INT_WR_BUF_DONE_EN_S = 6
const SPI_INT_RD_BUF_DONE_EN_V = 0x1
const SPI_INT_RD_BUF_DONE_EN_S = 5
const SPI_TRANS_DONE_V = 0x1
const SPI_TRANS_DONE_S = 4
const SPI_SLV_LAST_ADDR = 0x000000FF
const SPI_SLV_LAST_ADDR_V = 0xFF
const SPI_SLV_LAST_ADDR_S = 24
const SPI_SLV_LAST_COMMAND = 0x000000FF
const SPI_SLV_LAST_COMMAND_V = 0xFF
const SPI_SLV_LAST_COMMAND_S = 16
const SPI_SLV_WR_DMA_DONE_V = 0x1
const SPI_SLV_WR_DMA_DONE_S = 15
const SPI_SLV_CMD_ERR_V = 0x1
const SPI_SLV_CMD_ERR_S = 14
const SPI_SLV_ADDR_ERR_V = 0x1
const SPI_SLV_ADDR_ERR_S = 13
const SPI_SLV_NO_QPI_EN_V = 0x1
const SPI_SLV_NO_QPI_EN_S = 12
const SPI_SLV_CMD_ERR_CLR_V = 0x1
const SPI_SLV_CMD_ERR_CLR_S = 11
const SPI_SLV_ADDR_ERR_CLR_V = 0x1
const SPI_SLV_ADDR_ERR_CLR_S = 10
const SPI_CONF_BASE_BITLEN = 0x0000007F
const SPI_CONF_BASE_BITLEN_V = 0x7F
const SPI_CONF_BASE_BITLEN_S = 25
const SPI_SLV_WR_BUF_DONE_V = 0x1
const SPI_SLV_WR_BUF_DONE_S = 24
const SPI_SEG_MAGIC_ERR_V = 0x1
const SPI_SEG_MAGIC_ERR_S = 25
const SPI_SLV_RD_BUF_DONE_V = 0x1
const SPI_SLV_RD_BUF_DONE_S = 24
const SPI_SLV_DMA_RD_BYTELEN = 0x000FFFFF
const SPI_SLV_DMA_RD_BYTELEN_V = 0xFFFFF
const SPI_SLV_DMA_RD_BYTELEN_S = 0
const SPI_USR_CONF_V = 0x1
const SPI_USR_CONF_S = 31
const SPI_SLV_RD_DMA_DONE_V = 0x1
const SPI_SLV_RD_DMA_DONE_S = 30
const SPI_DMA_SEG_MAGIC_VALUE = 0x0000000F
const SPI_DMA_SEG_MAGIC_VALUE_V = 0xF
const SPI_DMA_SEG_MAGIC_VALUE_S = 24
const SPI_SLV_WRBUF_BYTELEN_EN_V = 0x1
const SPI_SLV_WRBUF_BYTELEN_EN_S = 23
const SPI_SLV_RDBUF_BYTELEN_EN_V = 0x1
const SPI_SLV_RDBUF_BYTELEN_EN_S = 22
const SPI_SLV_WRDMA_BYTELEN_EN_V = 0x1
const SPI_SLV_WRDMA_BYTELEN_EN_S = 21
const SPI_SLV_RDDMA_BYTELEN_EN_V = 0x1
const SPI_SLV_RDDMA_BYTELEN_EN_S = 20
const SPI_SLV_DATA_BYTELEN = 0x000FFFFF
const SPI_SLV_DATA_BYTELEN_V = 0xFFFFF
const SPI_SLV_DATA_BYTELEN_S = 0
const SPI_MST_DMA_RD_BYTELEN = 0x000FFFFF
const SPI_MST_DMA_RD_BYTELEN_V = 0xFFFFF
const SPI_MST_DMA_RD_BYTELEN_S = 12
const SPI_ST = 0x0000000F
const SPI_ST_V = 0xF
const SPI_ST_S = 0
const SPI_DMA_SEG_TRANS_DONE_V = 0x1
const SPI_DMA_SEG_TRANS_DONE_S = 7
const SPI_HOLD_OUT_TIME = 0x00000007
const SPI_HOLD_OUT_TIME_V = 0x7
const SPI_HOLD_OUT_TIME_S = 4
const SPI_HOLD_OUT_EN_V = 0x1
const SPI_HOLD_OUT_EN_S = 3
const SPI_HOLD_VAL_REG_V = 0x1
const SPI_HOLD_VAL_REG_S = 2
const SPI_INT_HOLD_ENA = 0x00000003
const SPI_INT_HOLD_ENA_V = 0x3
const SPI_INT_HOLD_ENA_S = 0
const SPI_DMA_SEG_TRANS_CLR_V = 0x1
const SPI_DMA_SEG_TRANS_CLR_S = 28
const SPI_EXT_MEM_BK_SIZE = 0x00000003
const SPI_EXT_MEM_BK_SIZE_V = 0x3
const SPI_EXT_MEM_BK_SIZE_S = 26
const SPI_DMA_OUTFIFO_EMPTY_CLR_V = 0x1
const SPI_DMA_OUTFIFO_EMPTY_CLR_S = 23
const SPI_DMA_INFIFO_FULL_CLR_V = 0x1
const SPI_DMA_INFIFO_FULL_CLR_S = 22
const SPI_RX_EOF_EN_V = 0x1
const SPI_RX_EOF_EN_S = 21
const SPI_SLV_TX_SEG_TRANS_CLR_EN_V = 0x1
const SPI_SLV_TX_SEG_TRANS_CLR_EN_S = 20
const SPI_SLV_RX_SEG_TRANS_CLR_EN_V = 0x1
const SPI_SLV_RX_SEG_TRANS_CLR_EN_S = 19
const SPI_DMA_SLV_SEG_TRANS_EN_V = 0x1
const SPI_DMA_SLV_SEG_TRANS_EN_S = 18
const SPI_SLV_LAST_SEG_POP_CLR_V = 0x1
const SPI_SLV_LAST_SEG_POP_CLR_S = 17
const SPI_DMA_CONTINUE_V = 0x1
const SPI_DMA_CONTINUE_S = 16
const SPI_DMA_TX_STOP_V = 0x1
const SPI_DMA_TX_STOP_S = 15
const SPI_DMA_RX_STOP_V = 0x1
const SPI_DMA_RX_STOP_S = 14
const SPI_MEM_TRANS_EN_V = 0x1
const SPI_MEM_TRANS_EN_S = 13
const SPI_OUT_DATA_BURST_EN_V = 0x1
const SPI_OUT_DATA_BURST_EN_S = 12
const SPI_INDSCR_BURST_EN_V = 0x1
const SPI_INDSCR_BURST_EN_S = 11
const SPI_OUTDSCR_BURST_EN_V = 0x1
const SPI_OUTDSCR_BURST_EN_S = 10
const SPI_OUT_EOF_MODE_V = 0x1
const SPI_OUT_EOF_MODE_S = 9
const SPI_OUT_AUTO_WRBACK_V = 0x1
const SPI_OUT_AUTO_WRBACK_S = 8
const SPI_OUT_LOOP_TEST_V = 0x1
const SPI_OUT_LOOP_TEST_S = 7
const SPI_IN_LOOP_TEST_V = 0x1
const SPI_IN_LOOP_TEST_S = 6
const SPI_AHBM_RST_V = 0x1
const SPI_AHBM_RST_S = 5
const SPI_AHBM_FIFO_RST_V = 0x1
const SPI_AHBM_FIFO_RST_S = 4
const SPI_OUT_RST_V = 0x1
const SPI_OUT_RST_S = 3
const SPI_IN_RST_V = 0x1
const SPI_IN_RST_S = 2
const SPI_DMA_TX_ENA_V = 0x1
const SPI_DMA_TX_ENA_S = 31
const SPI_OUTLINK_RESTART_V = 0x1
const SPI_OUTLINK_RESTART_S = 30
const SPI_OUTLINK_START_V = 0x1
const SPI_OUTLINK_START_S = 29
const SPI_OUTLINK_STOP_V = 0x1
const SPI_OUTLINK_STOP_S = 28
const SPI_OUTLINK_ADDR = 0x000FFFFF
const SPI_OUTLINK_ADDR_V = 0xFFFFF
const SPI_OUTLINK_ADDR_S = 0
const SPI_DMA_RX_ENA_V = 0x1
const SPI_DMA_RX_ENA_S = 31
const SPI_INLINK_RESTART_V = 0x1
const SPI_INLINK_RESTART_S = 30
const SPI_INLINK_START_V = 0x1
const SPI_INLINK_START_S = 29
const SPI_INLINK_STOP_V = 0x1
const SPI_INLINK_STOP_S = 28
const SPI_INLINK_AUTO_RET_V = 0x1
const SPI_INLINK_AUTO_RET_S = 20
const SPI_INLINK_ADDR = 0x000FFFFF
const SPI_INLINK_ADDR_V = 0xFFFFF
const SPI_INLINK_ADDR_S = 0
const SPI_SLV_CMDA_INT_ENA_V = 0x1
const SPI_SLV_CMDA_INT_ENA_S = 15
const SPI_SLV_CMD9_INT_ENA_V = 0x1
const SPI_SLV_CMD9_INT_ENA_S = 14
const SPI_SLV_CMD8_INT_ENA_V = 0x1
const SPI_SLV_CMD8_INT_ENA_S = 13
const SPI_SLV_CMD7_INT_ENA_V = 0x1
const SPI_SLV_CMD7_INT_ENA_S = 12
const SPI_SLV_CMD6_INT_ENA_V = 0x1
const SPI_SLV_CMD6_INT_ENA_S = 11
const SPI_OUTFIFO_EMPTY_ERR_INT_ENA_V = 0x1
const SPI_OUTFIFO_EMPTY_ERR_INT_ENA_S = 10
const SPI_INFIFO_FULL_ERR_INT_ENA_V = 0x1
const SPI_INFIFO_FULL_ERR_INT_ENA_S = 9
const SPI_OUT_TOTAL_EOF_INT_ENA_V = 0x1
const SPI_OUT_TOTAL_EOF_INT_ENA_S = 8
const SPI_OUT_EOF_INT_ENA_V = 0x1
const SPI_OUT_EOF_INT_ENA_S = 7
const SPI_OUT_DONE_INT_ENA_V = 0x1
const SPI_OUT_DONE_INT_ENA_S = 6
const SPI_IN_SUC_EOF_INT_ENA_V = 0x1
const SPI_IN_SUC_EOF_INT_ENA_S = 5
const SPI_IN_ERR_EOF_INT_ENA_V = 0x1
const SPI_IN_ERR_EOF_INT_ENA_S = 4
const SPI_IN_DONE_INT_ENA_V = 0x1
const SPI_IN_DONE_INT_ENA_S = 3
const SPI_INLINK_DSCR_ERROR_INT_ENA_V = 0x1
const SPI_INLINK_DSCR_ERROR_INT_ENA_S = 2
const SPI_OUTLINK_DSCR_ERROR_INT_ENA_V = 0x1
const SPI_OUTLINK_DSCR_ERROR_INT_ENA_S = 1
const SPI_INLINK_DSCR_EMPTY_INT_ENA_V = 0x1
const SPI_INLINK_DSCR_EMPTY_INT_ENA_S = 0
const SPI_SLV_CMDA_INT_RAW_V = 0x1
const SPI_SLV_CMDA_INT_RAW_S = 15
const SPI_SLV_CMD9_INT_RAW_V = 0x1
const SPI_SLV_CMD9_INT_RAW_S = 14
const SPI_SLV_CMD8_INT_RAW_V = 0x1
const SPI_SLV_CMD8_INT_RAW_S = 13
const SPI_SLV_CMD7_INT_RAW_V = 0x1
const SPI_SLV_CMD7_INT_RAW_S = 12
const SPI_SLV_CMD6_INT_RAW_V = 0x1
const SPI_SLV_CMD6_INT_RAW_S = 11
const SPI_OUTFIFO_EMPTY_ERR_INT_RAW_V = 0x1
const SPI_OUTFIFO_EMPTY_ERR_INT_RAW_S = 10
const SPI_INFIFO_FULL_ERR_INT_RAW_V = 0x1
const SPI_INFIFO_FULL_ERR_INT_RAW_S = 9
const SPI_OUT_TOTAL_EOF_INT_RAW_V = 0x1
const SPI_OUT_TOTAL_EOF_INT_RAW_S = 8
const SPI_OUT_EOF_INT_RAW_V = 0x1
const SPI_OUT_EOF_INT_RAW_S = 7
const SPI_OUT_DONE_INT_RAW_V = 0x1
const SPI_OUT_DONE_INT_RAW_S = 6
const SPI_IN_SUC_EOF_INT_RAW_V = 0x1
const SPI_IN_SUC_EOF_INT_RAW_S = 5
const SPI_IN_ERR_EOF_INT_RAW_V = 0x1
const SPI_IN_ERR_EOF_INT_RAW_S = 4
const SPI_IN_DONE_INT_RAW_V = 0x1
const SPI_IN_DONE_INT_RAW_S = 3
const SPI_INLINK_DSCR_ERROR_INT_RAW_V = 0x1
const SPI_INLINK_DSCR_ERROR_INT_RAW_S = 2
const SPI_OUTLINK_DSCR_ERROR_INT_RAW_V = 0x1
const SPI_OUTLINK_DSCR_ERROR_INT_RAW_S = 1
const SPI_INLINK_DSCR_EMPTY_INT_RAW_V = 0x1
const SPI_INLINK_DSCR_EMPTY_INT_RAW_S = 0
const SPI_SLV_CMDA_INT_ST_V = 0x1
const SPI_SLV_CMDA_INT_ST_S = 15
const SPI_SLV_CMD9_INT_ST_V = 0x1
const SPI_SLV_CMD9_INT_ST_S = 14
const SPI_SLV_CMD8_INT_ST_V = 0x1
const SPI_SLV_CMD8_INT_ST_S = 13
const SPI_SLV_CMD7_INT_ST_V = 0x1
const SPI_SLV_CMD7_INT_ST_S = 12
const SPI_SLV_CMD6_INT_ST_V = 0x1
const SPI_SLV_CMD6_INT_ST_S = 11
const SPI_OUTFIFO_EMPTY_ERR_INT_ST_V = 0x1
const SPI_OUTFIFO_EMPTY_ERR_INT_ST_S = 10
const SPI_INFIFO_FULL_ERR_INT_ST_V = 0x1
const SPI_INFIFO_FULL_ERR_INT_ST_S = 9
const SPI_OUT_TOTAL_EOF_INT_ST_V = 0x1
const SPI_OUT_TOTAL_EOF_INT_ST_S = 8
const SPI_OUT_EOF_INT_ST_V = 0x1
const SPI_OUT_EOF_INT_ST_S = 7
const SPI_OUT_DONE_INT_ST_V = 0x1
const SPI_OUT_DONE_INT_ST_S = 6
const SPI_IN_SUC_EOF_INT_ST_V = 0x1
const SPI_IN_SUC_EOF_INT_ST_S = 5
const SPI_IN_ERR_EOF_INT_ST_V = 0x1
const SPI_IN_ERR_EOF_INT_ST_S = 4
const SPI_IN_DONE_INT_ST_V = 0x1
const SPI_IN_DONE_INT_ST_S = 3
const SPI_INLINK_DSCR_ERROR_INT_ST_V = 0x1
const SPI_INLINK_DSCR_ERROR_INT_ST_S = 2
const SPI_OUTLINK_DSCR_ERROR_INT_ST_V = 0x1
const SPI_OUTLINK_DSCR_ERROR_INT_ST_S = 1
const SPI_INLINK_DSCR_EMPTY_INT_ST_V = 0x1
const SPI_INLINK_DSCR_EMPTY_INT_ST_S = 0
const SPI_SLV_CMDA_INT_CLR_V = 0x1
const SPI_SLV_CMDA_INT_CLR_S = 15
const SPI_SLV_CMD9_INT_CLR_V = 0x1
const SPI_SLV_CMD9_INT_CLR_S = 14
const SPI_SLV_CMD8_INT_CLR_V = 0x1
const SPI_SLV_CMD8_INT_CLR_S = 13
const SPI_SLV_CMD7_INT_CLR_V = 0x1
const SPI_SLV_CMD7_INT_CLR_S = 12
const SPI_SLV_CMD6_INT_CLR_V = 0x1
const SPI_SLV_CMD6_INT_CLR_S = 11
const SPI_OUTFIFO_EMPTY_ERR_INT_CLR_V = 0x1
const SPI_OUTFIFO_EMPTY_ERR_INT_CLR_S = 10
const SPI_INFIFO_FULL_ERR_INT_CLR_V = 0x1
const SPI_INFIFO_FULL_ERR_INT_CLR_S = 9
const SPI_OUT_TOTAL_EOF_INT_CLR_V = 0x1
const SPI_OUT_TOTAL_EOF_INT_CLR_S = 8
const SPI_OUT_EOF_INT_CLR_V = 0x1
const SPI_OUT_EOF_INT_CLR_S = 7
const SPI_OUT_DONE_INT_CLR_V = 0x1
const SPI_OUT_DONE_INT_CLR_S = 6
const SPI_IN_SUC_EOF_INT_CLR_V = 0x1
const SPI_IN_SUC_EOF_INT_CLR_S = 5
const SPI_IN_ERR_EOF_INT_CLR_V = 0x1
const SPI_IN_ERR_EOF_INT_CLR_S = 4
const SPI_IN_DONE_INT_CLR_V = 0x1
const SPI_IN_DONE_INT_CLR_S = 3
const SPI_INLINK_DSCR_ERROR_INT_CLR_V = 0x1
const SPI_INLINK_DSCR_ERROR_INT_CLR_S = 2
const SPI_OUTLINK_DSCR_ERROR_INT_CLR_V = 0x1
const SPI_OUTLINK_DSCR_ERROR_INT_CLR_S = 1
const SPI_INLINK_DSCR_EMPTY_INT_CLR_V = 0x1
const SPI_INLINK_DSCR_EMPTY_INT_CLR_S = 0
const SPI_DMA_IN_ERR_EOF_DES_ADDR = 0xFFFFFFFF
const SPI_DMA_IN_ERR_EOF_DES_ADDR_V = 0xFFFFFFFF
const SPI_DMA_IN_ERR_EOF_DES_ADDR_S = 0
const SPI_DMA_IN_SUC_EOF_DES_ADDR = 0xFFFFFFFF
const SPI_DMA_IN_SUC_EOF_DES_ADDR_V = 0xFFFFFFFF
const SPI_DMA_IN_SUC_EOF_DES_ADDR_S = 0
const SPI_DMA_INLINK_DSCR = 0xFFFFFFFF
const SPI_DMA_INLINK_DSCR_V = 0xFFFFFFFF
const SPI_DMA_INLINK_DSCR_S = 0
const SPI_DMA_INLINK_DSCR_BF0 = 0xFFFFFFFF
const SPI_DMA_INLINK_DSCR_BF0_V = 0xFFFFFFFF
const SPI_DMA_INLINK_DSCR_BF0_S = 0
const SPI_DMA_INLINK_DSCR_BF1 = 0xFFFFFFFF
const SPI_DMA_INLINK_DSCR_BF1_V = 0xFFFFFFFF
const SPI_DMA_INLINK_DSCR_BF1_S = 0
const SPI_DMA_OUT_EOF_BFR_DES_ADDR = 0xFFFFFFFF
const SPI_DMA_OUT_EOF_BFR_DES_ADDR_V = 0xFFFFFFFF
const SPI_DMA_OUT_EOF_BFR_DES_ADDR_S = 0
const SPI_DMA_OUT_EOF_DES_ADDR = 0xFFFFFFFF
const SPI_DMA_OUT_EOF_DES_ADDR_V = 0xFFFFFFFF
const SPI_DMA_OUT_EOF_DES_ADDR_S = 0
const SPI_DMA_OUTLINK_DSCR = 0xFFFFFFFF
const SPI_DMA_OUTLINK_DSCR_V = 0xFFFFFFFF
const SPI_DMA_OUTLINK_DSCR_S = 0
const SPI_DMA_OUTLINK_DSCR_BF0 = 0xFFFFFFFF
const SPI_DMA_OUTLINK_DSCR_BF0_V = 0xFFFFFFFF
const SPI_DMA_OUTLINK_DSCR_BF0_S = 0
const SPI_DMA_OUTLINK_DSCR_BF1 = 0xFFFFFFFF
const SPI_DMA_OUTLINK_DSCR_BF1_V = 0xFFFFFFFF
const SPI_DMA_OUTLINK_DSCR_BF1_S = 0
const SPI_DMA_OUTFIFO_EMPTY_V = 0x1
const SPI_DMA_OUTFIFO_EMPTY_S = 31
const SPI_DMA_OUTFIFO_FULL_V = 0x1
const SPI_DMA_OUTFIFO_FULL_S = 30
const SPI_DMA_OUTFIFO_CNT = 0x0000007F
const SPI_DMA_OUTFIFO_CNT_V = 0x7F
const SPI_DMA_OUTFIFO_CNT_S = 23
const SPI_DMA_OUT_STATE = 0x00000007
const SPI_DMA_OUT_STATE_V = 0x7
const SPI_DMA_OUT_STATE_S = 20
const SPI_DMA_OUTDSCR_STATE = 0x00000003
const SPI_DMA_OUTDSCR_STATE_V = 0x3
const SPI_DMA_OUTDSCR_STATE_S = 18
const SPI_DMA_OUTDSCR_ADDR = 0x0003FFFF
const SPI_DMA_OUTDSCR_ADDR_V = 0x3FFFF
const SPI_DMA_OUTDSCR_ADDR_S = 0
const SPI_DMA_INFIFO_EMPTY_V = 0x1
const SPI_DMA_INFIFO_EMPTY_S = 31
const SPI_DMA_INFIFO_FULL_V = 0x1
const SPI_DMA_INFIFO_FULL_S = 30
const SPI_DMA_INFIFO_CNT = 0x0000007F
const SPI_DMA_INFIFO_CNT_V = 0x7F
const SPI_DMA_INFIFO_CNT_S = 23
const SPI_DMA_IN_STATE = 0x00000007
const SPI_DMA_IN_STATE_V = 0x7
const SPI_DMA_IN_STATE_S = 20
const SPI_DMA_INDSCR_STATE = 0x00000003
const SPI_DMA_INDSCR_STATE_V = 0x3
const SPI_DMA_INDSCR_STATE_S = 18
const SPI_DMA_INDSCR_ADDR = 0x0003FFFF
const SPI_DMA_INDSCR_ADDR_V = 0x3FFFF
const SPI_DMA_INDSCR_ADDR_S = 0
const SPI_BUF0 = 0xFFFFFFFF
const SPI_BUF0_V = 0xFFFFFFFF
const SPI_BUF0_S = 0
const SPI_BUF1 = 0xFFFFFFFF
const SPI_BUF1_V = 0xFFFFFFFF
const SPI_BUF1_S = 0
const SPI_BUF2 = 0xFFFFFFFF
const SPI_BUF2_V = 0xFFFFFFFF
const SPI_BUF2_S = 0
const SPI_BUF3 = 0xFFFFFFFF
const SPI_BUF3_V = 0xFFFFFFFF
const SPI_BUF3_S = 0
const SPI_BUF4 = 0xFFFFFFFF
const SPI_BUF4_V = 0xFFFFFFFF
const SPI_BUF4_S = 0
const SPI_BUF5 = 0xFFFFFFFF
const SPI_BUF5_V = 0xFFFFFFFF
const SPI_BUF5_S = 0
const SPI_BUF6 = 0xFFFFFFFF
const SPI_BUF6_V = 0xFFFFFFFF
const SPI_BUF6_S = 0
const SPI_BUF7 = 0xFFFFFFFF
const SPI_BUF7_V = 0xFFFFFFFF
const SPI_BUF7_S = 0
const SPI_BUF8 = 0xFFFFFFFF
const SPI_BUF8_V = 0xFFFFFFFF
const SPI_BUF8_S = 0
const SPI_BUF9 = 0xFFFFFFFF
const SPI_BUF9_V = 0xFFFFFFFF
const SPI_BUF9_S = 0
const SPI_BUF10 = 0xFFFFFFFF
const SPI_BUF10_V = 0xFFFFFFFF
const SPI_BUF10_S = 0
const SPI_BUF11 = 0xFFFFFFFF
const SPI_BUF11_V = 0xFFFFFFFF
const SPI_BUF11_S = 0
const SPI_BUF12 = 0xFFFFFFFF
const SPI_BUF12_V = 0xFFFFFFFF
const SPI_BUF12_S = 0
const SPI_BUF13 = 0xFFFFFFFF
const SPI_BUF13_V = 0xFFFFFFFF
const SPI_BUF13_S = 0
const SPI_BUF14 = 0xFFFFFFFF
const SPI_BUF14_V = 0xFFFFFFFF
const SPI_BUF14_S = 0
const SPI_BUF15 = 0xFFFFFFFF
const SPI_BUF15_V = 0xFFFFFFFF
const SPI_BUF15_S = 0
const SPI_BUF16 = 0xFFFFFFFF
const SPI_BUF16_V = 0xFFFFFFFF
const SPI_BUF16_S = 0
const SPI_BUF17 = 0xFFFFFFFF
const SPI_BUF17_V = 0xFFFFFFFF
const SPI_BUF17_S = 0
const SPI_TIMING_CLK_ENA_V = 0x1
const SPI_TIMING_CLK_ENA_S = 24
const SPI_DIN7_MODE = 0x00000007
const SPI_DIN7_MODE_V = 0x7
const SPI_DIN7_MODE_S = 21
const SPI_DIN6_MODE = 0x00000007
const SPI_DIN6_MODE_V = 0x7
const SPI_DIN6_MODE_S = 18
const SPI_DIN5_MODE = 0x00000007
const SPI_DIN5_MODE_V = 0x7
const SPI_DIN5_MODE_S = 15
const SPI_DIN4_MODE = 0x00000007
const SPI_DIN4_MODE_V = 0x7
const SPI_DIN4_MODE_S = 12
const SPI_DIN3_MODE = 0x00000007
const SPI_DIN3_MODE_V = 0x7
const SPI_DIN3_MODE_S = 9
const SPI_DIN2_MODE = 0x00000007
const SPI_DIN2_MODE_V = 0x7
const SPI_DIN2_MODE_S = 6
const SPI_DIN1_MODE = 0x00000007
const SPI_DIN1_MODE_V = 0x7
const SPI_DIN1_MODE_S = 3
const SPI_DIN0_MODE = 0x00000007
const SPI_DIN0_MODE_V = 0x7
const SPI_DIN0_MODE_S = 0
const SPI_DIN7_NUM = 0x00000003
const SPI_DIN7_NUM_V = 0x3
const SPI_DIN7_NUM_S = 14
const SPI_DIN6_NUM = 0x00000003
const SPI_DIN6_NUM_V = 0x3
const SPI_DIN6_NUM_S = 12
const SPI_DIN5_NUM = 0x00000003
const SPI_DIN5_NUM_V = 0x3
const SPI_DIN5_NUM_S = 10
const SPI_DIN4_NUM = 0x00000003
const SPI_DIN4_NUM_V = 0x3
const SPI_DIN4_NUM_S = 8
const SPI_DIN3_NUM = 0x00000003
const SPI_DIN3_NUM_V = 0x3
const SPI_DIN3_NUM_S = 6
const SPI_DIN2_NUM = 0x00000003
const SPI_DIN2_NUM_V = 0x3
const SPI_DIN2_NUM_S = 4
const SPI_DIN1_NUM = 0x00000003
const SPI_DIN1_NUM_V = 0x3
const SPI_DIN1_NUM_S = 2
const SPI_DIN0_NUM = 0x00000003
const SPI_DIN0_NUM_V = 0x3
const SPI_DIN0_NUM_S = 0
const SPI_DOUT7_MODE = 0x00000007
const SPI_DOUT7_MODE_V = 0x7
const SPI_DOUT7_MODE_S = 21
const SPI_DOUT6_MODE = 0x00000007
const SPI_DOUT6_MODE_V = 0x7
const SPI_DOUT6_MODE_S = 18
const SPI_DOUT5_MODE = 0x00000007
const SPI_DOUT5_MODE_V = 0x7
const SPI_DOUT5_MODE_S = 15
const SPI_DOUT4_MODE = 0x00000007
const SPI_DOUT4_MODE_V = 0x7
const SPI_DOUT4_MODE_S = 12
const SPI_DOUT3_MODE = 0x00000007
const SPI_DOUT3_MODE_V = 0x7
const SPI_DOUT3_MODE_S = 9
const SPI_DOUT2_MODE = 0x00000007
const SPI_DOUT2_MODE_V = 0x7
const SPI_DOUT2_MODE_S = 6
const SPI_DOUT1_MODE = 0x00000007
const SPI_DOUT1_MODE_V = 0x7
const SPI_DOUT1_MODE_S = 3
const SPI_DOUT0_MODE = 0x00000007
const SPI_DOUT0_MODE_V = 0x7
const SPI_DOUT0_MODE_S = 0
const SPI_DOUT7_NUM = 0x00000003
const SPI_DOUT7_NUM_V = 0x3
const SPI_DOUT7_NUM_S = 14
const SPI_DOUT6_NUM = 0x00000003
const SPI_DOUT6_NUM_V = 0x3
const SPI_DOUT6_NUM_S = 12
const SPI_DOUT5_NUM = 0x00000003
const SPI_DOUT5_NUM_V = 0x3
const SPI_DOUT5_NUM_S = 10
const SPI_DOUT4_NUM = 0x00000003
const SPI_DOUT4_NUM_V = 0x3
const SPI_DOUT4_NUM_S = 8
const SPI_DOUT3_NUM = 0x00000003
const SPI_DOUT3_NUM_V = 0x3
const SPI_DOUT3_NUM_S = 6
const SPI_DOUT2_NUM = 0x00000003
const SPI_DOUT2_NUM_V = 0x3
const SPI_DOUT2_NUM_S = 4
const SPI_DOUT1_NUM = 0x00000003
const SPI_DOUT1_NUM_V = 0x3
const SPI_DOUT1_NUM_S = 2
const SPI_DOUT0_NUM = 0x00000003
const SPI_DOUT0_NUM_V = 0x3
const SPI_DOUT0_NUM_S = 0
const SPI_LCD_MODE_EN_V = 0x1
const SPI_LCD_MODE_EN_S = 31
const SPI_LCD_VT_HEIGHT = 0x000003FF
const SPI_LCD_VT_HEIGHT_V = 0x3FF
const SPI_LCD_VT_HEIGHT_S = 21
const SPI_LCD_VA_HEIGHT = 0x000003FF
const SPI_LCD_VA_HEIGHT_V = 0x3FF
const SPI_LCD_VA_HEIGHT_S = 11
const SPI_LCD_HB_FRONT = 0x000007FF
const SPI_LCD_HB_FRONT_V = 0x7FF
const SPI_LCD_HB_FRONT_S = 0
const SPI_LCD_HT_WIDTH = 0x00000FFF
const SPI_LCD_HT_WIDTH_V = 0xFFF
const SPI_LCD_HT_WIDTH_S = 20
const SPI_LCD_HA_WIDTH = 0x00000FFF
const SPI_LCD_HA_WIDTH_V = 0xFFF
const SPI_LCD_HA_WIDTH_S = 8
const SPI_LCD_VB_FRONT = 0x000000FF
const SPI_LCD_VB_FRONT_V = 0xFF
const SPI_LCD_VB_FRONT_S = 0
const SPI_LCD_HSYNC_POSITION = 0x000000FF
const SPI_LCD_HSYNC_POSITION_V = 0xFF
const SPI_LCD_HSYNC_POSITION_S = 24
const SPI_HSYNC_IDLE_POL_V = 0x1
const SPI_HSYNC_IDLE_POL_S = 23
const SPI_LCD_HSYNC_WIDTH = 0x0000007F
const SPI_LCD_HSYNC_WIDTH_V = 0x7F
const SPI_LCD_HSYNC_WIDTH_S = 16
const SPI_VSYNC_IDLE_POL_V = 0x1
const SPI_VSYNC_IDLE_POL_S = 7
const SPI_LCD_VSYNC_WIDTH = 0x0000007F
const SPI_LCD_VSYNC_WIDTH_V = 0x7F
const SPI_LCD_VSYNC_WIDTH_S = 0
const SPI_HS_BLANK_EN_V = 0x1
const SPI_HS_BLANK_EN_S = 16
const SPI_DE_IDLE_POL_V = 0x1
const SPI_DE_IDLE_POL_S = 15
const SPI_D_VSYNC_MODE = 0x00000007
const SPI_D_VSYNC_MODE_V = 0x7
const SPI_D_VSYNC_MODE_S = 12
const SPI_D_HSYNC_MODE = 0x00000007
const SPI_D_HSYNC_MODE_V = 0x7
const SPI_D_HSYNC_MODE_S = 9
const SPI_D_DE_MODE = 0x00000007
const SPI_D_DE_MODE_V = 0x7
const SPI_D_DE_MODE_S = 6
const SPI_D_CD_MODE = 0x00000007
const SPI_D_CD_MODE_V = 0x7
const SPI_D_CD_MODE_S = 3
const SPI_D_DQS_MODE = 0x00000007
const SPI_D_DQS_MODE_V = 0x7
const SPI_D_DQS_MODE_S = 0
const SPI_D_VSYNC_NUM = 0x00000003
const SPI_D_VSYNC_NUM_V = 0x3
const SPI_D_VSYNC_NUM_S = 8
const SPI_D_HSYNC_NUM = 0x00000003
const SPI_D_HSYNC_NUM_V = 0x3
const SPI_D_HSYNC_NUM_S = 6
const SPI_D_DE_NUM = 0x00000003
const SPI_D_DE_NUM_V = 0x3
const SPI_D_DE_NUM_S = 4
const SPI_D_CD_NUM = 0x00000003
const SPI_D_CD_NUM_V = 0x3
const SPI_D_CD_NUM_S = 2
const SPI_D_DQS_NUM = 0x00000003
const SPI_D_DQS_NUM_V = 0x3
const SPI_D_DQS_NUM_S = 0
const SPI_DATE = 0x0FFFFFFF
const SPI_DATE_V = 0xFFFFFFF
const SPI_DATE_S = 0
