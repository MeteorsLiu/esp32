package soc

import _ "unsafe"

const UHCI_UART_RX_BRK_EOF_EN_V = 0x1
const UHCI_UART_RX_BRK_EOF_EN_S = 23
const UHCI_CLK_EN_V = 0x1
const UHCI_CLK_EN_S = 22
const UHCI_ENCODE_CRC_EN_V = 0x1
const UHCI_ENCODE_CRC_EN_S = 21
const UHCI_LEN_EOF_EN_V = 0x1
const UHCI_LEN_EOF_EN_S = 20
const UHCI_UART_IDLE_EOF_EN_V = 0x1
const UHCI_UART_IDLE_EOF_EN_S = 19
const UHCI_CRC_REC_EN_V = 0x1
const UHCI_CRC_REC_EN_S = 18
const UHCI_HEAD_EN_V = 0x1
const UHCI_HEAD_EN_S = 17
const UHCI_SEPER_EN_V = 0x1
const UHCI_SEPER_EN_S = 16
const UHCI_MEM_TRANS_EN_V = 0x1
const UHCI_MEM_TRANS_EN_S = 15
const UHCI_OUT_DATA_BURST_EN_V = 0x1
const UHCI_OUT_DATA_BURST_EN_S = 14
const UHCI_INDSCR_BURST_EN_V = 0x1
const UHCI_INDSCR_BURST_EN_S = 13
const UHCI_OUTDSCR_BURST_EN_V = 0x1
const UHCI_OUTDSCR_BURST_EN_S = 12
const UHCI_UART1_CE_V = 0x1
const UHCI_UART1_CE_S = 10
const UHCI_UART0_CE_V = 0x1
const UHCI_UART0_CE_S = 9
const UHCI_OUT_EOF_MODE_V = 0x1
const UHCI_OUT_EOF_MODE_S = 8
const UHCI_OUT_NO_RESTART_CLR_V = 0x1
const UHCI_OUT_NO_RESTART_CLR_S = 7
const UHCI_OUT_AUTO_WRBACK_V = 0x1
const UHCI_OUT_AUTO_WRBACK_S = 6
const UHCI_OUT_LOOP_TEST_V = 0x1
const UHCI_OUT_LOOP_TEST_S = 5
const UHCI_IN_LOOP_TEST_V = 0x1
const UHCI_IN_LOOP_TEST_S = 4
const UHCI_AHBM_RST_V = 0x1
const UHCI_AHBM_RST_S = 3
const UHCI_AHBM_FIFO_RST_V = 0x1
const UHCI_AHBM_FIFO_RST_S = 2
const UHCI_OUT_RST_V = 0x1
const UHCI_OUT_RST_S = 1
const UHCI_IN_RST_V = 0x1
const UHCI_IN_RST_S = 0
const UHCI_DMA_INFIFO_FULL_WM_INT_RAW_V = 0x1
const UHCI_DMA_INFIFO_FULL_WM_INT_RAW_S = 16
const UHCI_SEND_A_Q_INT_RAW_V = 0x1
const UHCI_SEND_A_Q_INT_RAW_S = 15
const UHCI_SEND_S_Q_INT_RAW_V = 0x1
const UHCI_SEND_S_Q_INT_RAW_S = 14
const UHCI_OUT_TOTAL_EOF_INT_RAW_V = 0x1
const UHCI_OUT_TOTAL_EOF_INT_RAW_S = 13
const UHCI_OUTLINK_EOF_ERR_INT_RAW_V = 0x1
const UHCI_OUTLINK_EOF_ERR_INT_RAW_S = 12
const UHCI_IN_DSCR_EMPTY_INT_RAW_V = 0x1
const UHCI_IN_DSCR_EMPTY_INT_RAW_S = 11
const UHCI_OUT_DSCR_ERR_INT_RAW_V = 0x1
const UHCI_OUT_DSCR_ERR_INT_RAW_S = 10
const UHCI_IN_DSCR_ERR_INT_RAW_V = 0x1
const UHCI_IN_DSCR_ERR_INT_RAW_S = 9
const UHCI_OUT_EOF_INT_RAW_V = 0x1
const UHCI_OUT_EOF_INT_RAW_S = 8
const UHCI_OUT_DONE_INT_RAW_V = 0x1
const UHCI_OUT_DONE_INT_RAW_S = 7
const UHCI_IN_ERR_EOF_INT_RAW_V = 0x1
const UHCI_IN_ERR_EOF_INT_RAW_S = 6
const UHCI_IN_SUC_EOF_INT_RAW_V = 0x1
const UHCI_IN_SUC_EOF_INT_RAW_S = 5
const UHCI_IN_DONE_INT_RAW_V = 0x1
const UHCI_IN_DONE_INT_RAW_S = 4
const UHCI_TX_HUNG_INT_RAW_V = 0x1
const UHCI_TX_HUNG_INT_RAW_S = 3
const UHCI_RX_HUNG_INT_RAW_V = 0x1
const UHCI_RX_HUNG_INT_RAW_S = 2
const UHCI_TX_START_INT_RAW_V = 0x1
const UHCI_TX_START_INT_RAW_S = 1
const UHCI_RX_START_INT_RAW_V = 0x1
const UHCI_RX_START_INT_RAW_S = 0
const UHCI_DMA_INFIFO_FULL_WM_INT_ST_V = 0x1
const UHCI_DMA_INFIFO_FULL_WM_INT_ST_S = 16
const UHCI_SEND_A_Q_INT_ST_V = 0x1
const UHCI_SEND_A_Q_INT_ST_S = 15
const UHCI_SEND_S_Q_INT_ST_V = 0x1
const UHCI_SEND_S_Q_INT_ST_S = 14
const UHCI_OUT_TOTAL_EOF_INT_ST_V = 0x1
const UHCI_OUT_TOTAL_EOF_INT_ST_S = 13
const UHCI_OUTLINK_EOF_ERR_INT_ST_V = 0x1
const UHCI_OUTLINK_EOF_ERR_INT_ST_S = 12
const UHCI_IN_DSCR_EMPTY_INT_ST_V = 0x1
const UHCI_IN_DSCR_EMPTY_INT_ST_S = 11
const UHCI_OUT_DSCR_ERR_INT_ST_V = 0x1
const UHCI_OUT_DSCR_ERR_INT_ST_S = 10
const UHCI_IN_DSCR_ERR_INT_ST_V = 0x1
const UHCI_IN_DSCR_ERR_INT_ST_S = 9
const UHCI_OUT_EOF_INT_ST_V = 0x1
const UHCI_OUT_EOF_INT_ST_S = 8
const UHCI_OUT_DONE_INT_ST_V = 0x1
const UHCI_OUT_DONE_INT_ST_S = 7
const UHCI_IN_ERR_EOF_INT_ST_V = 0x1
const UHCI_IN_ERR_EOF_INT_ST_S = 6
const UHCI_IN_SUC_EOF_INT_ST_V = 0x1
const UHCI_IN_SUC_EOF_INT_ST_S = 5
const UHCI_IN_DONE_INT_ST_V = 0x1
const UHCI_IN_DONE_INT_ST_S = 4
const UHCI_TX_HUNG_INT_ST_V = 0x1
const UHCI_TX_HUNG_INT_ST_S = 3
const UHCI_RX_HUNG_INT_ST_V = 0x1
const UHCI_RX_HUNG_INT_ST_S = 2
const UHCI_TX_START_INT_ST_V = 0x1
const UHCI_TX_START_INT_ST_S = 1
const UHCI_RX_START_INT_ST_V = 0x1
const UHCI_RX_START_INT_ST_S = 0
const UHCI_DMA_INFIFO_FULL_WM_INT_ENA_V = 0x1
const UHCI_DMA_INFIFO_FULL_WM_INT_ENA_S = 16
const UHCI_SEND_A_Q_INT_ENA_V = 0x1
const UHCI_SEND_A_Q_INT_ENA_S = 15
const UHCI_SEND_S_Q_INT_ENA_V = 0x1
const UHCI_SEND_S_Q_INT_ENA_S = 14
const UHCI_OUT_TOTAL_EOF_INT_ENA_V = 0x1
const UHCI_OUT_TOTAL_EOF_INT_ENA_S = 13
const UHCI_OUTLINK_EOF_ERR_INT_ENA_V = 0x1
const UHCI_OUTLINK_EOF_ERR_INT_ENA_S = 12
const UHCI_IN_DSCR_EMPTY_INT_ENA_V = 0x1
const UHCI_IN_DSCR_EMPTY_INT_ENA_S = 11
const UHCI_OUT_DSCR_ERR_INT_ENA_V = 0x1
const UHCI_OUT_DSCR_ERR_INT_ENA_S = 10
const UHCI_IN_DSCR_ERR_INT_ENA_V = 0x1
const UHCI_IN_DSCR_ERR_INT_ENA_S = 9
const UHCI_OUT_EOF_INT_ENA_V = 0x1
const UHCI_OUT_EOF_INT_ENA_S = 8
const UHCI_OUT_DONE_INT_ENA_V = 0x1
const UHCI_OUT_DONE_INT_ENA_S = 7
const UHCI_IN_ERR_EOF_INT_ENA_V = 0x1
const UHCI_IN_ERR_EOF_INT_ENA_S = 6
const UHCI_IN_SUC_EOF_INT_ENA_V = 0x1
const UHCI_IN_SUC_EOF_INT_ENA_S = 5
const UHCI_IN_DONE_INT_ENA_V = 0x1
const UHCI_IN_DONE_INT_ENA_S = 4
const UHCI_TX_HUNG_INT_ENA_V = 0x1
const UHCI_TX_HUNG_INT_ENA_S = 3
const UHCI_RX_HUNG_INT_ENA_V = 0x1
const UHCI_RX_HUNG_INT_ENA_S = 2
const UHCI_TX_START_INT_ENA_V = 0x1
const UHCI_TX_START_INT_ENA_S = 1
const UHCI_RX_START_INT_ENA_V = 0x1
const UHCI_RX_START_INT_ENA_S = 0
const UHCI_DMA_INFIFO_FULL_WM_INT_CLR_V = 0x1
const UHCI_DMA_INFIFO_FULL_WM_INT_CLR_S = 16
const UHCI_SEND_A_Q_INT_CLR_V = 0x1
const UHCI_SEND_A_Q_INT_CLR_S = 15
const UHCI_SEND_S_Q_INT_CLR_V = 0x1
const UHCI_SEND_S_Q_INT_CLR_S = 14
const UHCI_OUT_TOTAL_EOF_INT_CLR_V = 0x1
const UHCI_OUT_TOTAL_EOF_INT_CLR_S = 13
const UHCI_OUTLINK_EOF_ERR_INT_CLR_V = 0x1
const UHCI_OUTLINK_EOF_ERR_INT_CLR_S = 12
const UHCI_IN_DSCR_EMPTY_INT_CLR_V = 0x1
const UHCI_IN_DSCR_EMPTY_INT_CLR_S = 11
const UHCI_OUT_DSCR_ERR_INT_CLR_V = 0x1
const UHCI_OUT_DSCR_ERR_INT_CLR_S = 10
const UHCI_IN_DSCR_ERR_INT_CLR_V = 0x1
const UHCI_IN_DSCR_ERR_INT_CLR_S = 9
const UHCI_OUT_EOF_INT_CLR_V = 0x1
const UHCI_OUT_EOF_INT_CLR_S = 8
const UHCI_OUT_DONE_INT_CLR_V = 0x1
const UHCI_OUT_DONE_INT_CLR_S = 7
const UHCI_IN_ERR_EOF_INT_CLR_V = 0x1
const UHCI_IN_ERR_EOF_INT_CLR_S = 6
const UHCI_IN_SUC_EOF_INT_CLR_V = 0x1
const UHCI_IN_SUC_EOF_INT_CLR_S = 5
const UHCI_IN_DONE_INT_CLR_V = 0x1
const UHCI_IN_DONE_INT_CLR_S = 4
const UHCI_TX_HUNG_INT_CLR_V = 0x1
const UHCI_TX_HUNG_INT_CLR_S = 3
const UHCI_RX_HUNG_INT_CLR_V = 0x1
const UHCI_RX_HUNG_INT_CLR_S = 2
const UHCI_TX_START_INT_CLR_V = 0x1
const UHCI_TX_START_INT_CLR_S = 1
const UHCI_RX_START_INT_CLR_V = 0x1
const UHCI_RX_START_INT_CLR_S = 0
const UHCI_OUT_EMPTY_V = 0x1
const UHCI_OUT_EMPTY_S = 1
const UHCI_OUT_FULL_V = 0x1
const UHCI_OUT_FULL_S = 0
const UHCI_OUTFIFO_PUSH_V = 0x1
const UHCI_OUTFIFO_PUSH_S = 16
const UHCI_OUTFIFO_WDATA = 0x000001FF
const UHCI_OUTFIFO_WDATA_V = 0x1FF
const UHCI_OUTFIFO_WDATA_S = 0
const UHCI_RX_ERR_CAUSE = 0x00000007
const UHCI_RX_ERR_CAUSE_V = 0x7
const UHCI_RX_ERR_CAUSE_S = 4
const UHCI_IN_EMPTY_V = 0x1
const UHCI_IN_EMPTY_S = 1
const UHCI_IN_FULL_V = 0x1
const UHCI_IN_FULL_S = 0
const UHCI_INFIFO_POP_V = 0x1
const UHCI_INFIFO_POP_S = 16
const UHCI_INFIFO_RDATA = 0x00000FFF
const UHCI_INFIFO_RDATA_V = 0xFFF
const UHCI_INFIFO_RDATA_S = 0
const UHCI_OUTLINK_PARK_V = 0x1
const UHCI_OUTLINK_PARK_S = 31
const UHCI_OUTLINK_RESTART_V = 0x1
const UHCI_OUTLINK_RESTART_S = 30
const UHCI_OUTLINK_START_V = 0x1
const UHCI_OUTLINK_START_S = 29
const UHCI_OUTLINK_STOP_V = 0x1
const UHCI_OUTLINK_STOP_S = 28
const UHCI_OUTLINK_ADDR = 0x000FFFFF
const UHCI_OUTLINK_ADDR_V = 0xFFFFF
const UHCI_OUTLINK_ADDR_S = 0
const UHCI_INLINK_PARK_V = 0x1
const UHCI_INLINK_PARK_S = 31
const UHCI_INLINK_RESTART_V = 0x1
const UHCI_INLINK_RESTART_S = 30
const UHCI_INLINK_START_V = 0x1
const UHCI_INLINK_START_S = 29
const UHCI_INLINK_STOP_V = 0x1
const UHCI_INLINK_STOP_S = 28
const UHCI_INLINK_AUTO_RET_V = 0x1
const UHCI_INLINK_AUTO_RET_S = 20
const UHCI_INLINK_ADDR = 0x000FFFFF
const UHCI_INLINK_ADDR_V = 0xFFFFF
const UHCI_INLINK_ADDR_S = 0
const UHCI_DMA_INFIFO_FULL_THRS = 0x00000FFF
const UHCI_DMA_INFIFO_FULL_THRS_V = 0xFFF
const UHCI_DMA_INFIFO_FULL_THRS_S = 9
const UHCI_SW_START_V = 0x1
const UHCI_SW_START_S = 8
const UHCI_WAIT_SW_START_V = 0x1
const UHCI_WAIT_SW_START_S = 7
const UHCI_CHECK_OWNER_V = 0x1
const UHCI_CHECK_OWNER_S = 6
const UHCI_TX_ACK_NUM_RE_V = 0x1
const UHCI_TX_ACK_NUM_RE_S = 5
const UHCI_TX_CHECK_SUM_RE_V = 0x1
const UHCI_TX_CHECK_SUM_RE_S = 4
const UHCI_SAVE_HEAD_V = 0x1
const UHCI_SAVE_HEAD_S = 3
const UHCI_CRC_DISABLE_V = 0x1
const UHCI_CRC_DISABLE_S = 2
const UHCI_CHECK_SEQ_EN_V = 0x1
const UHCI_CHECK_SEQ_EN_S = 1
const UHCI_CHECK_SUM_EN_V = 0x1
const UHCI_CHECK_SUM_EN_S = 0
const UHCI_DECODE_STATE = 0x00000007
const UHCI_DECODE_STATE_V = 0x7
const UHCI_DECODE_STATE_S = 28
const UHCI_INFIFO_CNT_DEBUG = 0x0000001F
const UHCI_INFIFO_CNT_DEBUG_V = 0x1F
const UHCI_INFIFO_CNT_DEBUG_S = 23
const UHCI_IN_STATE = 0x00000007
const UHCI_IN_STATE_V = 0x7
const UHCI_IN_STATE_S = 20
const UHCI_IN_DSCR_STATE = 0x00000003
const UHCI_IN_DSCR_STATE_V = 0x3
const UHCI_IN_DSCR_STATE_S = 18
const UHCI_INLINK_DSCR_ADDR = 0x0003FFFF
const UHCI_INLINK_DSCR_ADDR_V = 0x3FFFF
const UHCI_INLINK_DSCR_ADDR_S = 0
const UHCI_ENCODE_STATE = 0x00000007
const UHCI_ENCODE_STATE_V = 0x7
const UHCI_ENCODE_STATE_S = 28
const UHCI_OUTFIFO_CNT = 0x0000001F
const UHCI_OUTFIFO_CNT_V = 0x1F
const UHCI_OUTFIFO_CNT_S = 23
const UHCI_OUT_STATE = 0x00000007
const UHCI_OUT_STATE_V = 0x7
const UHCI_OUT_STATE_S = 20
const UHCI_OUT_DSCR_STATE = 0x00000003
const UHCI_OUT_DSCR_STATE_V = 0x3
const UHCI_OUT_DSCR_STATE_S = 18
const UHCI_OUTLINK_DSCR_ADDR = 0x0003FFFF
const UHCI_OUTLINK_DSCR_ADDR_V = 0x3FFFF
const UHCI_OUTLINK_DSCR_ADDR_S = 0
const UHCI_OUT_EOF_DES_ADDR = 0xFFFFFFFF
const UHCI_OUT_EOF_DES_ADDR_V = 0xFFFFFFFF
const UHCI_OUT_EOF_DES_ADDR_S = 0
const UHCI_IN_SUC_EOF_DES_ADDR = 0xFFFFFFFF
const UHCI_IN_SUC_EOF_DES_ADDR_V = 0xFFFFFFFF
const UHCI_IN_SUC_EOF_DES_ADDR_S = 0
const UHCI_IN_ERR_EOF_DES_ADDR = 0xFFFFFFFF
const UHCI_IN_ERR_EOF_DES_ADDR_V = 0xFFFFFFFF
const UHCI_IN_ERR_EOF_DES_ADDR_S = 0
const UHCI_OUT_EOF_BFR_DES_ADDR = 0xFFFFFFFF
const UHCI_OUT_EOF_BFR_DES_ADDR_V = 0xFFFFFFFF
const UHCI_OUT_EOF_BFR_DES_ADDR_S = 0
const UHCI_AHB_TESTADDR = 0x00000003
const UHCI_AHB_TESTADDR_V = 0x3
const UHCI_AHB_TESTADDR_S = 4
const UHCI_AHB_TESTMODE = 0x00000007
const UHCI_AHB_TESTMODE_V = 0x7
const UHCI_AHB_TESTMODE_S = 0
const UHCI_INLINK_DSCR = 0xFFFFFFFF
const UHCI_INLINK_DSCR_V = 0xFFFFFFFF
const UHCI_INLINK_DSCR_S = 0
const UHCI_INLINK_DSCR_BF0 = 0xFFFFFFFF
const UHCI_INLINK_DSCR_BF0_V = 0xFFFFFFFF
const UHCI_INLINK_DSCR_BF0_S = 0
const UHCI_INLINK_DSCR_BF1 = 0xFFFFFFFF
const UHCI_INLINK_DSCR_BF1_V = 0xFFFFFFFF
const UHCI_INLINK_DSCR_BF1_S = 0
const UHCI_OUTLINK_DSCR = 0xFFFFFFFF
const UHCI_OUTLINK_DSCR_V = 0xFFFFFFFF
const UHCI_OUTLINK_DSCR_S = 0
const UHCI_OUTLINK_DSCR_BF0 = 0xFFFFFFFF
const UHCI_OUTLINK_DSCR_BF0_V = 0xFFFFFFFF
const UHCI_OUTLINK_DSCR_BF0_S = 0
const UHCI_OUTLINK_DSCR_BF1 = 0xFFFFFFFF
const UHCI_OUTLINK_DSCR_BF1_V = 0xFFFFFFFF
const UHCI_OUTLINK_DSCR_BF1_S = 0
const UHCI_RX_13_ESC_EN_V = 0x1
const UHCI_RX_13_ESC_EN_S = 7
const UHCI_RX_11_ESC_EN_V = 0x1
const UHCI_RX_11_ESC_EN_S = 6
const UHCI_RX_DB_ESC_EN_V = 0x1
const UHCI_RX_DB_ESC_EN_S = 5
const UHCI_RX_C0_ESC_EN_V = 0x1
const UHCI_RX_C0_ESC_EN_S = 4
const UHCI_TX_13_ESC_EN_V = 0x1
const UHCI_TX_13_ESC_EN_S = 3
const UHCI_TX_11_ESC_EN_V = 0x1
const UHCI_TX_11_ESC_EN_S = 2
const UHCI_TX_DB_ESC_EN_V = 0x1
const UHCI_TX_DB_ESC_EN_S = 1
const UHCI_TX_C0_ESC_EN_V = 0x1
const UHCI_TX_C0_ESC_EN_S = 0
const UHCI_RXFIFO_TIMEOUT_ENA_V = 0x1
const UHCI_RXFIFO_TIMEOUT_ENA_S = 23
const UHCI_RXFIFO_TIMEOUT_SHIFT = 0x00000007
const UHCI_RXFIFO_TIMEOUT_SHIFT_V = 0x7
const UHCI_RXFIFO_TIMEOUT_SHIFT_S = 20
const UHCI_RXFIFO_TIMEOUT = 0x000000FF
const UHCI_RXFIFO_TIMEOUT_V = 0xFF
const UHCI_RXFIFO_TIMEOUT_S = 12
const UHCI_TXFIFO_TIMEOUT_ENA_V = 0x1
const UHCI_TXFIFO_TIMEOUT_ENA_S = 11
const UHCI_TXFIFO_TIMEOUT_SHIFT = 0x00000007
const UHCI_TXFIFO_TIMEOUT_SHIFT_V = 0x7
const UHCI_TXFIFO_TIMEOUT_SHIFT_S = 8
const UHCI_TXFIFO_TIMEOUT = 0x000000FF
const UHCI_TXFIFO_TIMEOUT_V = 0xFF
const UHCI_TXFIFO_TIMEOUT_S = 0
const UHCI_RX_HEAD = 0xFFFFFFFF
const UHCI_RX_HEAD_V = 0xFFFFFFFF
const UHCI_RX_HEAD_S = 0
const UHCI_ALWAYS_SEND_EN_V = 0x1
const UHCI_ALWAYS_SEND_EN_S = 7
const UHCI_ALWAYS_SEND_NUM = 0x00000007
const UHCI_ALWAYS_SEND_NUM_V = 0x7
const UHCI_ALWAYS_SEND_NUM_S = 4
const UHCI_SINGLE_SEND_EN_V = 0x1
const UHCI_SINGLE_SEND_EN_S = 3
const UHCI_SINGLE_SEND_NUM = 0x00000007
const UHCI_SINGLE_SEND_NUM_V = 0x7
const UHCI_SINGLE_SEND_NUM_S = 0
const UHCI_SEND_Q0_WORD0 = 0xFFFFFFFF
const UHCI_SEND_Q0_WORD0_V = 0xFFFFFFFF
const UHCI_SEND_Q0_WORD0_S = 0
const UHCI_SEND_Q0_WORD1 = 0xFFFFFFFF
const UHCI_SEND_Q0_WORD1_V = 0xFFFFFFFF
const UHCI_SEND_Q0_WORD1_S = 0
const UHCI_SEND_Q1_WORD0 = 0xFFFFFFFF
const UHCI_SEND_Q1_WORD0_V = 0xFFFFFFFF
const UHCI_SEND_Q1_WORD0_S = 0
const UHCI_SEND_Q1_WORD1 = 0xFFFFFFFF
const UHCI_SEND_Q1_WORD1_V = 0xFFFFFFFF
const UHCI_SEND_Q1_WORD1_S = 0
const UHCI_SEND_Q2_WORD0 = 0xFFFFFFFF
const UHCI_SEND_Q2_WORD0_V = 0xFFFFFFFF
const UHCI_SEND_Q2_WORD0_S = 0
const UHCI_SEND_Q2_WORD1 = 0xFFFFFFFF
const UHCI_SEND_Q2_WORD1_V = 0xFFFFFFFF
const UHCI_SEND_Q2_WORD1_S = 0
const UHCI_SEND_Q3_WORD0 = 0xFFFFFFFF
const UHCI_SEND_Q3_WORD0_V = 0xFFFFFFFF
const UHCI_SEND_Q3_WORD0_S = 0
const UHCI_SEND_Q3_WORD1 = 0xFFFFFFFF
const UHCI_SEND_Q3_WORD1_V = 0xFFFFFFFF
const UHCI_SEND_Q3_WORD1_S = 0
const UHCI_SEND_Q4_WORD0 = 0xFFFFFFFF
const UHCI_SEND_Q4_WORD0_V = 0xFFFFFFFF
const UHCI_SEND_Q4_WORD0_S = 0
const UHCI_SEND_Q4_WORD1 = 0xFFFFFFFF
const UHCI_SEND_Q4_WORD1_V = 0xFFFFFFFF
const UHCI_SEND_Q4_WORD1_S = 0
const UHCI_SEND_Q5_WORD0 = 0xFFFFFFFF
const UHCI_SEND_Q5_WORD0_V = 0xFFFFFFFF
const UHCI_SEND_Q5_WORD0_S = 0
const UHCI_SEND_Q5_WORD1 = 0xFFFFFFFF
const UHCI_SEND_Q5_WORD1_V = 0xFFFFFFFF
const UHCI_SEND_Q5_WORD1_S = 0
const UHCI_SEND_Q6_WORD0 = 0xFFFFFFFF
const UHCI_SEND_Q6_WORD0_V = 0xFFFFFFFF
const UHCI_SEND_Q6_WORD0_S = 0
const UHCI_SEND_Q6_WORD1 = 0xFFFFFFFF
const UHCI_SEND_Q6_WORD1_V = 0xFFFFFFFF
const UHCI_SEND_Q6_WORD1_S = 0
const UHCI_SEPER_ESC_CHAR1 = 0x000000FF
const UHCI_SEPER_ESC_CHAR1_V = 0xFF
const UHCI_SEPER_ESC_CHAR1_S = 16
const UHCI_SEPER_ESC_CHAR0 = 0x000000FF
const UHCI_SEPER_ESC_CHAR0_V = 0xFF
const UHCI_SEPER_ESC_CHAR0_S = 8
const UHCI_SEPER_CHAR = 0x000000FF
const UHCI_SEPER_CHAR_V = 0xFF
const UHCI_SEPER_CHAR_S = 0
const UHCI_ESC_SEQ0_CHAR1 = 0x000000FF
const UHCI_ESC_SEQ0_CHAR1_V = 0xFF
const UHCI_ESC_SEQ0_CHAR1_S = 16
const UHCI_ESC_SEQ0_CHAR0 = 0x000000FF
const UHCI_ESC_SEQ0_CHAR0_V = 0xFF
const UHCI_ESC_SEQ0_CHAR0_S = 8
const UHCI_ESC_SEQ0 = 0x000000FF
const UHCI_ESC_SEQ0_V = 0xFF
const UHCI_ESC_SEQ0_S = 0
const UHCI_ESC_SEQ1_CHAR1 = 0x000000FF
const UHCI_ESC_SEQ1_CHAR1_V = 0xFF
const UHCI_ESC_SEQ1_CHAR1_S = 16
const UHCI_ESC_SEQ1_CHAR0 = 0x000000FF
const UHCI_ESC_SEQ1_CHAR0_V = 0xFF
const UHCI_ESC_SEQ1_CHAR0_S = 8
const UHCI_ESC_SEQ1 = 0x000000FF
const UHCI_ESC_SEQ1_V = 0xFF
const UHCI_ESC_SEQ1_S = 0
const UHCI_ESC_SEQ2_CHAR1 = 0x000000FF
const UHCI_ESC_SEQ2_CHAR1_V = 0xFF
const UHCI_ESC_SEQ2_CHAR1_S = 16
const UHCI_ESC_SEQ2_CHAR0 = 0x000000FF
const UHCI_ESC_SEQ2_CHAR0_V = 0xFF
const UHCI_ESC_SEQ2_CHAR0_S = 8
const UHCI_ESC_SEQ2 = 0x000000FF
const UHCI_ESC_SEQ2_V = 0xFF
const UHCI_ESC_SEQ2_S = 0
const UHCI_PKT_THRS = 0x00001FFF
const UHCI_PKT_THRS_V = 0x1FFF
const UHCI_PKT_THRS_S = 0
const UHCI_DATE = 0xFFFFFFFF
const UHCI_DATE_V = 0xFFFFFFFF
const UHCI_DATE_S = 0
