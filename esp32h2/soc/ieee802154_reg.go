package soc

import _ "unsafe"

const IEEE802154_OPCODE = 0x000000FF
const IEEE802154_OPCODE_S = 0
const IEEE802154_MAC_INF3_ENABLE_S = 31
const IEEE802154_MAC_INF2_ENABLE_S = 30
const IEEE802154_MAC_INF1_ENABLE_S = 29
const IEEE802154_MAC_INF0_ENABLE_S = 28
const IEEE802154_RX_DONE_TRIGGER_IDLE_S = 27
const IEEE802154_FORCE_RX_ENB_S = 26
const IEEE802154_NO_RSS_TRK_ENB_S = 25
const IEEE802154_BIT_ORDER_S = 24
const IEEE802154_COEX_ARB_DELAY = 0x000000FF
const IEEE802154_COEX_ARB_DELAY_S = 16
const IEEE802154_FILTER_ENHANCE_S = 14
const IEEE802154_AUTOPEND_ENHANCE_S = 12
const IEEE802154_DIS_FRAME_VERSION_RSV_FILTER_S = 11
const IEEE802154_PROMISCUOUS_MODE_S = 7
const IEEE802154_PAN_COORDINATOR_S = 6
const IEEE802154_DIS_IFS_CONTROL_S = 5
const IEEE802154_HW_AUTO_ACK_RX_EN_S = 3
const HW_ENHANCE_ACK_TX_EN_S = 1
const IEEE802154_HW_AUTO_ACK_TX_EN_S = 0
const IEEE802154_MAC_INF0_SHORT_ADDR = 0x0000FFFF
const IEEE802154_MAC_INF0_SHORT_ADDR_S = 0
const IEEE802154_MAC_INF0_PAN_ID = 0x0000FFFF
const IEEE802154_MAC_INF0_PAN_ID_S = 0
const IEEE802154_MAC_INF0_EXTEND_ADDR0 = 0xFFFFFFFF
const IEEE802154_MAC_INF0_EXTEND_ADDR0_S = 0
const IEEE802154_MAC_INF0_EXTEND_ADDR1 = 0xFFFFFFFF
const IEEE802154_MAC_INF0_EXTEND_ADDR1_S = 0
const IEEE802154_MAC_INF1_SHORT_ADDR = 0x0000FFFF
const IEEE802154_MAC_INF1_SHORT_ADDR_S = 0
const IEEE802154_MAC_INF1_PAN_ID = 0x0000FFFF
const IEEE802154_MAC_INF1_PAN_ID_S = 0
const IEEE802154_MAC_INF1_EXTEND_ADDR0 = 0xFFFFFFFF
const IEEE802154_MAC_INF1_EXTEND_ADDR0_S = 0
const IEEE802154_MAC_INF1_EXTEND_ADDR1 = 0xFFFFFFFF
const IEEE802154_MAC_INF1_EXTEND_ADDR1_S = 0
const IEEE802154_MAC_INF2_SHORT_ADDR = 0x0000FFFF
const IEEE802154_MAC_INF2_SHORT_ADDR_S = 0
const IEEE802154_MAC_INF2_PAN_ID = 0x0000FFFF
const IEEE802154_MAC_INF2_PAN_ID_S = 0
const IEEE802154_MAC_INF2_EXTEND_ADDR0 = 0xFFFFFFFF
const IEEE802154_MAC_INF2_EXTEND_ADDR0_S = 0
const IEEE802154_MAC_INF2_EXTEND_ADDR1 = 0xFFFFFFFF
const IEEE802154_MAC_INF2_EXTEND_ADDR1_S = 0
const IEEE802154_MAC_INF3_SHORT_ADDR = 0x0000FFFF
const IEEE802154_MAC_INF3_SHORT_ADDR_S = 0
const IEEE802154_MAC_INF3_PAN_ID = 0x0000FFFF
const IEEE802154_MAC_INF3_PAN_ID_S = 0
const IEEE802154_MAC_INF3_EXTEND_ADDR0 = 0xFFFFFFFF
const IEEE802154_MAC_INF3_EXTEND_ADDR0_S = 0
const IEEE802154_MAC_INF3_EXTEND_ADDR1 = 0xFFFFFFFF
const IEEE802154_MAC_INF3_EXTEND_ADDR1_S = 0
const IEEE802154_HOP = 0x0000007F
const IEEE802154_HOP_S = 0
const IEEE802154_TX_POWER = 0x0000001F
const IEEE802154_TX_POWER_S = 0
const IEEE802154_ED_SCAN_WAIT_DLY = 0x0000000F
const IEEE802154_ED_SCAN_WAIT_DLY_S = 24
const IEEE802154_ED_SCAN_DURATION = 0x00FFFFFF
const IEEE802154_ED_SCAN_DURATION_S = 0
const IEEE802154_CCA_BUSY_S = 24
const IEEE802154_ED_RSS = 0x000000FF
const IEEE802154_ED_RSS_S = 16
const IEEE802154_CCA_MODE = 0x00000003
const IEEE802154_CCA_MODE_S = 14
const IEEE802154_DIS_ED_POWER_SEL_S = 13
const IEEE802154_ED_SAMPLE_MODE = 0x00000003
const IEEE802154_ED_SAMPLE_MODE_S = 11
const IEEE802154_CCA_ED_THRESHOLD = 0x000000FF
const IEEE802154_CCA_ED_THRESHOLD_S = 0
const IEEE802154_LIFS = 0x000003FF
const IEEE802154_LIFS_S = 16
const IEEE802154_SIFS = 0x000000FF
const IEEE802154_SIFS_S = 0
const IEEE802154_ACK_TIMEOUT = 0x0000FFFF
const IEEE802154_ACK_TIMEOUT_S = 0
const IEEE802154_EVENT_EN = 0x00001FFF
const IEEE802154_EVENT_EN_S = 0
const IEEE802154_EVENT_STATUS = 0x00001FFF
const IEEE802154_EVENT_STATUS_S = 0
const IEEE802154_RX_ABORT_INTR_CTRL = 0x7FFFFFFF
const IEEE802154_RX_ABORT_INTR_CTRL_S = 0
const IEEE802154_ACK_TX_ACK_TIMEOUT = 0x0000FFFF
const IEEE802154_ACK_TX_ACK_TIMEOUT_S = 16
const IEEE802154_ACK_FRAME_PENDING_EN_S = 0
const IEEE802154_CLOSE_RF_SEL_S = 8
const IEEE802154_COEX_ACK_PTI = 0x0000000F
const IEEE802154_COEX_ACK_PTI_S = 4
const IEEE802154_COEX_PTI = 0x0000000F
const IEEE802154_COEX_PTI_S = 0
const IEEE802154_TX_ABORT_INTERRUPT_CONTROL = 0x7FFFFFFF
const IEEE802154_TX_ABORT_INTERRUPT_CONTROL_S = 0
const IEEE802154_TX_ENH_ACK_GENERATE_DONE_NOTIFY = 0xFFFFFFFF
const IEEE802154_TX_ENH_ACK_GENERATE_DONE_NOTIFY_S = 0
const IEEE802154_SFD_MATCH_S = 21
const IEEE802154_PREAMBLE_MATCH_S = 20
const IEEE802154_RX_STATE = 0x00000007
const IEEE802154_RX_STATE_S = 16
const IEEE802154_RX_ABORT_STATUS = 0x0000001F
const IEEE802154_RX_ABORT_STATUS_S = 4
const IEEE802154_FILTER_FAIL_STATUS = 0x0000000F
const IEEE802154_FILTER_FAIL_STATUS_S = 0
const IEEE802154_TX_SEC_ERROR_CODE = 0x0000000F
const IEEE802154_TX_SEC_ERROR_CODE_S = 16
const IEEE802154_TX_ABORT_STATUS = 0x0000001F
const IEEE802154_TX_ABORT_STATUS_S = 4
const IEEE802154_TX_STATE = 0x0000000F
const IEEE802154_TX_STATE_S = 0
const IEEE802154_RF_CTRL_STATE = 0x0000000F
const IEEE802154_RF_CTRL_STATE_S = 16
const IEEE802154_ED_TRIGGER_TX_PROC_S = 11
const IEEE802154_ED_PROC_S = 10
const IEEE802154_RX_PROC_S = 9
const IEEE802154_TX_PROC_S = 8
const IEEE802154_TXRX_STATE = 0x0000000F
const IEEE802154_TXRX_STATE_S = 0
const IEEE802154_TX_CCM_SCHEDULE_STATUS = 0x7FFFFFFF
const IEEE802154_TX_CCM_SCHEDULE_STATUS_S = 0
const IEEE802154_RX_LENGTH = 0x0000007F
const IEEE802154_RX_LENGTH_S = 0
const IEEE802154_TIMER0_THRESHOLD = 0xFFFFFFFF
const IEEE802154_TIMER0_THRESHOLD_S = 0
const IEEE802154_TIMER0_VALUE = 0xFFFFFFFF
const IEEE802154_TIMER0_VALUE_S = 0
const IEEE802154_TIMER1_THRESHOLD = 0xFFFFFFFF
const IEEE802154_TIMER1_THRESHOLD_S = 0
const IEEE802154_TIMER1_VALUE = 0xFFFFFFFF
const IEEE802154_TIMER1_VALUE_S = 0
const IEEE802154_CLK_COUNT_MATCH_VAL = 0x0000FFFF
const IEEE802154_CLK_COUNT_MATCH_VAL_S = 0
const IEEE802154_CLK_625US_CNT = 0x0000FFFF
const IEEE802154_CLK_625US_CNT_S = 0
const IEEE802154_IFS_COUNTER_EN_S = 16
const IEEE802154_IFS_COUNTER = 0x000003FF
const IEEE802154_IFS_COUNTER_S = 0
const IEEE802154_SFD_WAIT_SYMBOL_NUM = 0x0000000F
const IEEE802154_SFD_WAIT_SYMBOL_NUM_S = 0
const IEEE802154_RX_PATH_DELAY = 0x0000003F
const IEEE802154_RX_PATH_DELAY_S = 16
const IEEE802154_TX_PATH_DELAY = 0x0000003F
const IEEE802154_TX_PATH_DELAY_S = 0
const IEEE802154_BB_CLK_FREQ_MINUS_1 = 0x0000001F
const IEEE802154_BB_CLK_FREQ_MINUS_1_S = 0
const IEEE802154_TXDMA_ADDR = 0xFFFFFFFF
const IEEE802154_TXDMA_ADDR_S = 0
const IEEE802154_TXDMA_FETCH_BYTE_CNT = 0x0000007F
const IEEE802154_TXDMA_FETCH_BYTE_CNT_S = 24
const IEEE802154_TXDMA_STATE = 0x0000001F
const IEEE802154_TXDMA_STATE_S = 16
const IEEE802154_TXDMA_FILL_ENTRY = 0x00000007
const IEEE802154_TXDMA_FILL_ENTRY_S = 4
const IEEE802154_TXDMA_WATER_LEVEL = 0x00000007
const IEEE802154_TXDMA_WATER_LEVEL_S = 0
const IEEE802154_TXDMA_ERR = 0x0000000F
const IEEE802154_TXDMA_ERR_S = 0
const IEEE802154_RXDMA_ADDR = 0xFFFFFFFF
const IEEE802154_RXDMA_ADDR_S = 0
const IEEE802154_RXDMA_APPEND_FREQ_OFFSET_S = 25
const IEEE802154_RXDMA_APPEND_LQI_OFFSET_S = 24
const IEEE802154_RXDMA_STATE = 0x0000001F
const IEEE802154_RXDMA_STATE_S = 16
const IEEE802154_RXDMA_WATER_LEVEL = 0x00000007
const IEEE802154_RXDMA_WATER_LEVEL_S = 0
const IEEE802154_RXDMA_ERR = 0x0000000F
const IEEE802154_RXDMA_ERR_S = 0
const IEEE802154_DMA_GCK_CFG_S = 0
const IEEE802154_DMA_DUMMY_DATA = 0xFFFFFFFF
const IEEE802154_PAON_DELAY = 0x000003FF
const IEEE802154_PAON_DELAY_S = 0
const IEEE802154_TXON_DELAY = 0x000003FF
const IEEE802154_TXON_DELAY_S = 0
const IEEE802154_TXEN_STOP_DLY = 0x0000003F
const IEEE802154_TXEN_STOP_DLY_S = 0
const IEEE802154_TXOFF_DELAY = 0x0000003F
const IEEE802154_TXOFF_DELAY_S = 0
const IEEE802154_RXON_DELAY = 0x000007FF
const IEEE802154_RXON_DELAY_S = 0
const IEEE802154_TXRX_SWITCH_DELAY = 0x000003FF
const IEEE802154_TXRX_SWITCH_DELAY_S = 0
const IEEE802154_CONT_RX_DELAY = 0x0000003F
const IEEE802154_CONT_RX_DELAY_S = 0
const IEEE802154_TX_DCDC_UP_S = 31
const IEEE802154_DCDC_CTRL_EN_S = 16
const IEEE802154_DCDC_DOWN_DELAY = 0x000000FF
const IEEE802154_DCDC_DOWN_DELAY_S = 8
const IEEE802154_DCDC_PRE_UP_DELAY = 0x000000FF
const IEEE802154_DCDC_PRE_UP_DELAY_S = 0
const IEEE802154_DEBUG_TRIGGER_DUMP_EN_S = 31
const IEEE802154_DEBUG_STATE_MATCH_DUMP_EN_S = 30
const IEEE802154_DEBUG_TRIGGER_PULSE_SELECT = 0x00000007
const IEEE802154_DEBUG_TRIGGER_PULSE_SELECT_S = 24
const IEEE802154_DEBUG_TRIGGER_STATE_MATCH_VALUE = 0x0000001F
const IEEE802154_DEBUG_TRIGGER_STATE_MATCH_VALUE_S = 16
const IEEE802154_DEBUG_SER_DEBUG_SEL = 0x0000000F
const IEEE802154_DEBUG_SER_DEBUG_SEL_S = 12
const IEEE802154_DEBUG_TRIGGER_STATE_SELECT = 0x0000000F
const IEEE802154_DEBUG_TRIGGER_STATE_SELECT_S = 8
const IEEE802154_DEBUG_SIGNAL_SEL = 0x00000007
const IEEE802154_DEBUG_SIGNAL_SEL_S = 0
const IEEE802154_SEC_PAYLOAD_OFFSET = 0x0000007F
const IEEE802154_SEC_PAYLOAD_OFFSET_S = 8
const IEEE802154_SEC_EN_S = 0
const IEEE802154_SEC_EXTEND_ADDRESS0 = 0xFFFFFFFF
const IEEE802154_SEC_EXTEND_ADDRESS0_S = 0
const IEEE802154_SEC_EXTEND_ADDRESS1 = 0xFFFFFFFF
const IEEE802154_SEC_EXTEND_ADDRESS1_S = 0
const IEEE802154_SEC_KEY0 = 0xFFFFFFFF
const IEEE802154_SEC_KEY0_S = 0
const IEEE802154_SEC_KEY1 = 0xFFFFFFFF
const IEEE802154_SEC_KEY1_S = 0
const IEEE802154_SEC_KEY2 = 0xFFFFFFFF
const IEEE802154_SEC_KEY2_S = 0
const IEEE802154_SEC_KEY3 = 0xFFFFFFFF
const IEEE802154_SEC_KEY3_S = 0
const IEEE802154_SFD_TIMEOUT_CNT = 0x0000FFFF
const IEEE802154_SFD_TIMEOUT_CNT_S = 0
const IEEE802154_CRC_ERROR_CNT = 0x0000FFFF
const IEEE802154_CRC_ERROR_CNT_S = 0
const IEEE802154_ED_ABORT_CNT = 0x0000FFFF
const IEEE802154_ED_ABORT_CNT_S = 0
const IEEE802154_CCA_FAIL_CNT = 0x0000FFFF
const IEEE802154_CCA_FAIL_CNT_S = 0
const IEEE802154_RX_FILTER_FAIL_CNT = 0x0000FFFF
const IEEE802154_RX_FILTER_FAIL_CNT_S = 0
const IEEE802154_NO_RSS_DETECT_CNT = 0x0000FFFF
const IEEE802154_NO_RSS_DETECT_CNT_S = 0
const IEEE802154_RX_ABORT_COEX_CNT = 0x0000FFFF
const IEEE802154_RX_ABORT_COEX_CNT_S = 0
const IEEE802154_RX_RESTART_CNT = 0x0000FFFF
const IEEE802154_RX_RESTART_CNT_S = 0
const IEEE802154_TX_ACK_ABORT_COEX_CNT = 0x0000FFFF
const IEEE802154_TX_ACK_ABORT_COEX_CNT_S = 0
const IEEE802154_ED_SCAN_COEX_CNT = 0x0000FFFF
const IEEE802154_ED_SCAN_COEX_CNT_S = 0
const IEEE802154_RX_ACK_ABORT_COEX_CNT = 0x0000FFFF
const IEEE802154_RX_ACK_ABORT_COEX_CNT_S = 0
const IEEE802154_RX_ACK_TIMEOUT_CNT = 0x0000FFFF
const IEEE802154_RX_ACK_TIMEOUT_CNT_S = 0
const IEEE802154_TX_BREAK_COEX_CNT = 0x0000FFFF
const IEEE802154_TX_BREAK_COEX_CNT_S = 0
const IEEE802154_TX_SECURITY_ERROR_CNT = 0x0000FFFF
const IEEE802154_TX_SECURITY_ERROR_CNT_S = 0
const IEEE802154_CCA_BUSY_CNT = 0x0000FFFF
const IEEE802154_CCA_BUSY_CNT_S = 0
const IEEE802154_SFD_TIMEOUT_CNT_CLEAR_S = 14
const IEEE802154_CRC_ERROR_CNT_CLEAR_S = 13
const IEEE802154_RX_FILTER_FAIL_CNT_CLEAR_S = 12
const IEEE802154_NO_RSS_DETECT_CNT_CLEAR_S = 11
const IEEE802154_RX_ABORT_COEX_CNT_CLEAR_S = 10
const IEEE802154_RX_ACK_ABORT_COEX_CNT_CLEAR_S = 9
const IEEE802154_RX_RESTART_CNT_CLEAR_S = 8
const IEEE802154_RX_ACK_TIMEOUT_CNT_CLEAR_S = 7
const IEEE802154_TX_ACK_ABORT_COEX_CNT_CLEAR_S = 6
const IEEE802154_TX_BREAK_COEX_CNT_CLEAR_S = 5
const IEEE802154_TX_SECURITY_ERROR_CNT_CLEAR_S = 4
const IEEE802154_ED_ABORT_CNT_CLEAR_S = 3
const IEEE802154_CCA_FAIL_CNT_CLEAR_S = 2
const IEEE802154_CCA_BUSY_CNT_CLEAR_S = 1
const IEEE802154_ED_SCAN_COEX_CNT_CLEAR_S = 0
const DEBUG_FIELD3_SEL = 0x0000001F
const DEBUG_FIELD3_SEL_S = 24
const DEBUG_FIELD2_SEL = 0x0000001F
const DEBUG_FIELD2_SEL_S = 16
const DEBUG_FIELD1_SEL = 0x0000001F
const DEBUG_FIELD1_SEL_S = 8
const DEBUG_FIELD0_SEL = 0x0000001F
const DEBUG_FIELD0_SEL_S = 0
const DEBUG_FIELD7_SEL = 0x0000001F
const DEBUG_FIELD7_SEL_S = 24
const DEBUG_FIELD6_SEL = 0x0000001F
const DEBUG_FIELD6_SEL_S = 16
const DEBUG_FIELD5_SEL = 0x0000001F
const DEBUG_FIELD5_SEL_S = 8
const DEBUG_FIELD4_SEL = 0x0000001F
const DEBUG_FIELD4_SEL_S = 0
const IEEE802154_MAC_DATE = 0xFFFFFFFF
const IEEE802154_MAC_DATE_S = 0
const IEEE802154_MAC_DATE_VERSION = 0x220907
const ETM_REG_BASE = 0x600A4000
const ETM_CH_OFFSET = 0x08
const ETM_EVENT_TIMER1_OVERFLOW = 58
const ETM_EVENT_TIMER0_OVERFLOW = 59
const ETM_TASK_ED_TRIG_TX = 64
const ETM_TASK_RX_START = 65
const ETM_TASK_TX_START = 68
