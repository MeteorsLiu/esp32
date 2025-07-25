package soc

import _ "unsafe"

const GPIO_EVT_CH0_RISE_EDGE = 1
const GPIO_EVT_CH1_RISE_EDGE = 2
const GPIO_EVT_CH2_RISE_EDGE = 3
const GPIO_EVT_CH3_RISE_EDGE = 4
const GPIO_EVT_CH4_RISE_EDGE = 5
const GPIO_EVT_CH5_RISE_EDGE = 6
const GPIO_EVT_CH6_RISE_EDGE = 7
const GPIO_EVT_CH7_RISE_EDGE = 8
const GPIO_EVT_CH0_FALL_EDGE = 9
const GPIO_EVT_CH1_FALL_EDGE = 10
const GPIO_EVT_CH2_FALL_EDGE = 11
const GPIO_EVT_CH3_FALL_EDGE = 12
const GPIO_EVT_CH4_FALL_EDGE = 13
const GPIO_EVT_CH5_FALL_EDGE = 14
const GPIO_EVT_CH6_FALL_EDGE = 15
const GPIO_EVT_CH7_FALL_EDGE = 16
const GPIO_EVT_CH0_ANY_EDGE = 17
const GPIO_EVT_CH1_ANY_EDGE = 18
const GPIO_EVT_CH2_ANY_EDGE = 19
const GPIO_EVT_CH3_ANY_EDGE = 20
const GPIO_EVT_CH4_ANY_EDGE = 21
const GPIO_EVT_CH5_ANY_EDGE = 22
const GPIO_EVT_CH6_ANY_EDGE = 23
const GPIO_EVT_CH7_ANY_EDGE = 24
const LEDC_EVT_DUTY_CHNG_END_CH0 = 25
const LEDC_EVT_DUTY_CHNG_END_CH1 = 26
const LEDC_EVT_DUTY_CHNG_END_CH2 = 27
const LEDC_EVT_DUTY_CHNG_END_CH3 = 28
const LEDC_EVT_DUTY_CHNG_END_CH4 = 29
const LEDC_EVT_DUTY_CHNG_END_CH5 = 30
const LEDC_EVT_OVF_CNT_PLS_CH0 = 31
const LEDC_EVT_OVF_CNT_PLS_CH1 = 32
const LEDC_EVT_OVF_CNT_PLS_CH2 = 33
const LEDC_EVT_OVF_CNT_PLS_CH3 = 34
const LEDC_EVT_OVF_CNT_PLS_CH4 = 35
const LEDC_EVT_OVF_CNT_PLS_CH5 = 36
const LEDC_EVT_TIME_OVF_TIMER0 = 37
const LEDC_EVT_TIME_OVF_TIMER1 = 38
const LEDC_EVT_TIME_OVF_TIMER2 = 39
const LEDC_EVT_TIME_OVF_TIMER3 = 40
const LEDC_EVT_TIMER0_CMP = 41
const LEDC_EVT_TIMER1_CMP = 42
const LEDC_EVT_TIMER2_CMP = 43
const LEDC_EVT_TIMER3_CMP = 44
const PCNT_EVT_CNT_EQ_THRESH = 45
const PCNT_EVT_CNT_EQ_LMT = 46
const PCNT_EVT_CNT_EQ_ZERO = 47
const TIMER0_EVT_CNT_CMP_TIMER0 = 48
const TIMER1_EVT_CNT_CMP_TIMER0 = 49
const SYSTIMER_EVT_CNT_CMP0 = 50
const SYSTIMER_EVT_CNT_CMP1 = 51
const SYSTIMER_EVT_CNT_CMP2 = 52
const RMT_EVT_TX_END = 53
const RMT_EVT_TX_LOOP = 54
const RMT_EVT_RX_END = 55
const RMT_EVT_TX_THRESH = 56
const RMT_EVT_RX_THRESH = 57
const MCPWM_EVT_TIMER0_STOP = 58
const MCPWM_EVT_TIMER1_STOP = 59
const MCPWM_EVT_TIMER2_STOP = 60
const MCPWM_EVT_TIMER0_TEZ = 61
const MCPWM_EVT_TIMER1_TEZ = 62
const MCPWM_EVT_TIMER2_TEZ = 63
const MCPWM_EVT_TIMER0_TEP = 64
const MCPWM_EVT_TIMER1_TEP = 65
const MCPWM_EVT_TIMER2_TEP = 66
const MCPWM_EVT_OP0_TEA = 67
const MCPWM_EVT_OP1_TEA = 68
const MCPWM_EVT_OP2_TEA = 69
const MCPWM_EVT_OP0_TEB = 70
const MCPWM_EVT_OP1_TEB = 71
const MCPWM_EVT_OP2_TEB = 72
const MCPWM_EVT_F0 = 73
const MCPWM_EVT_F1 = 74
const MCPWM_EVT_F2 = 75
const MCPWM_EVT_F0_CLR = 76
const MCPWM_EVT_F1_CLR = 77
const MCPWM_EVT_F2_CLR = 78
const MCPWM_EVT_TZ0_CBC = 79
const MCPWM_EVT_TZ1_CBC = 80
const MCPWM_EVT_TZ2_CBC = 81
const MCPWM_EVT_TZ0_OST = 82
const MCPWM_EVT_TZ1_OST = 83
const MCPWM_EVT_TZ2_OST = 84
const MCPWM_EVT_CAP0 = 85
const MCPWM_EVT_CAP1 = 86
const MCPWM_EVT_CAP2 = 87
const ADC_EVT_CONV_CMPLT0 = 88
const ADC_EVT_EQ_ABOVE_THRESH0 = 89
const ADC_EVT_EQ_ABOVE_THRESH1 = 90
const ADC_EVT_EQ_BELOW_THRESH0 = 91
const ADC_EVT_EQ_BELOW_THRESH1 = 92
const ADC_EVT_RESULT_DONE0 = 93
const ADC_EVT_STOPPED0 = 94
const ADC_EVT_STARTED0 = 95
const REGDMA_EVT_DONE0 = 96
const REGDMA_EVT_DONE1 = 97
const REGDMA_EVT_DONE2 = 98
const REGDMA_EVT_DONE3 = 99
const REGDMA_EVT_ERR0 = 100
const REGDMA_EVT_ERR1 = 101
const REGDMA_EVT_ERR2 = 102
const REGDMA_EVT_ERR3 = 103
const PDMA_EVT_TX_DONE = 104
const PDMA_EVT_OUT_EOF = 105
const PDMA_EVT_IN_SUC_EOF = 106
const PDMA_EVT_FULL_OR_EMPTY = 107
const PDMA_EVT_ALL_DONE = 108
const PDMA_EVT_RX_DONE = 109
const TMPSNSR_EVT_OVER_LIMIT = 110
const UART_EVT_REC_DATA_OVF0 = 111
const UART_EVT_REC_DATA_OVF1 = 112
const UART_EVT_TX_DONE0 = 113
const UART_EVT_TX_DONE1 = 114
const UART_EVT_TIMEOUT0 = 115
const UART_EVT_TIMEOUT1 = 116
const UART_EVT_ERR0 = 117
const UART_EVT_ERR1 = 118
const UART_EVT_CTS0 = 119
const UART_EVT_CTS1 = 120
const UART_EVT_TX_EMPTY0 = 121
const UART_EVT_TX_EMPTY1 = 122
const UART_EVT_AT_PATTERNS0 = 123
const UART_EVT_AT_PATTERNS1 = 124
const SPI_EVT_STOPPED = 125
const I2S_EVT_RX_DONE = 126
const I2S_EVT_TX_DONE = 127
const I2S_EVT_X_WORDS_RECEIVED = 128
const I2S_EVT_X_WORDS_SENT = 129
const I2C_EVT_TRANS_DONE = 130
const LCDCAM_EVT_TRANS_DONE = 131
const CAN_EVT_TRANS_DONE = 132
const ULP_EVT_ERR_INTR = 133
const ULP_EVT_START_INTR = 134
const RTC_EVT_TICK = 135
const RTC_EVT_OVF = 136
const RTC_EVT_CMP = 137
const GDMA_EVT_IN_DONE_CH0 = 138
const GDMA_EVT_IN_DONE_CH1 = 139
const GDMA_EVT_IN_DONE_CH2 = 140
const GDMA_EVT_IN_SUC_EOF_CH0 = 141
const GDMA_EVT_IN_SUC_EOF_CH1 = 142
const GDMA_EVT_IN_SUC_EOF_CH2 = 143
const GDMA_EVT_IN_FIFO_EMPTY_CH0 = 144
const GDMA_EVT_IN_FIFO_EMPTY_CH1 = 145
const GDMA_EVT_IN_FIFO_EMPTY_CH2 = 146
const GDMA_EVT_IN_FIFO_FULL_CH0 = 147
const GDMA_EVT_IN_FIFO_FULL_CH1 = 148
const GDMA_EVT_IN_FIFO_FULL_CH2 = 149
const GDMA_EVT_OUT_DONE_CH0 = 150
const GDMA_EVT_OUT_DONE_CH1 = 151
const GDMA_EVT_OUT_DONE_CH2 = 152
const GDMA_EVT_OUT_EOF_CH0 = 153
const GDMA_EVT_OUT_EOF_CH1 = 154
const GDMA_EVT_OUT_EOF_CH2 = 155
const GDMA_EVT_OUT_TOTAL_EOF_CH0 = 156
const GDMA_EVT_OUT_TOTAL_EOF_CH1 = 157
const GDMA_EVT_OUT_TOTAL_EOF_CH2 = 158
const GDMA_EVT_OUT_FIFO_EMPTY_CH0 = 159
const GDMA_EVT_OUT_FIFO_EMPTY_CH1 = 160
const GDMA_EVT_OUT_FIFO_EMPTY_CH2 = 161
const GDMA_EVT_OUT_FIFO_FULL_CH0 = 162
const GDMA_EVT_OUT_FIFO_FULL_CH1 = 163
const GDMA_EVT_OUT_FIFO_FULL_CH2 = 164
const PMU_EVT_SLEEP_WEEKUP = 165
const GPIO_TASK_CH0_SET = 1
const GPIO_TASK_CH1_SET = 2
const GPIO_TASK_CH2_SET = 3
const GPIO_TASK_CH3_SET = 4
const GPIO_TASK_CH4_SET = 5
const GPIO_TASK_CH5_SET = 6
const GPIO_TASK_CH6_SET = 7
const GPIO_TASK_CH7_SET = 8
const GPIO_TASK_CH0_CLEAR = 9
const GPIO_TASK_CH1_CLEAR = 10
const GPIO_TASK_CH2_CLEAR = 11
const GPIO_TASK_CH3_CLEAR = 12
const GPIO_TASK_CH4_CLEAR = 13
const GPIO_TASK_CH5_CLEAR = 14
const GPIO_TASK_CH6_CLEAR = 15
const GPIO_TASK_CH7_CLEAR = 16
const GPIO_TASK_CH0_TOGGLE = 17
const GPIO_TASK_CH1_TOGGLE = 18
const GPIO_TASK_CH2_TOGGLE = 19
const GPIO_TASK_CH3_TOGGLE = 20
const GPIO_TASK_CH4_TOGGLE = 21
const GPIO_TASK_CH5_TOGGLE = 22
const GPIO_TASK_CH6_TOGGLE = 23
const GPIO_TASK_CH7_TOGGLE = 24
const LEDC_TASK_TIMER0_RES_UPDATE = 25
const LEDC_TASK_TIMER1_RES_UPDATE = 26
const LEDC_TASK_TIMER2_RES_UPDATE = 27
const LEDC_TASK_TIMER3_RES_UPDATE = 28
const LEDC_TASK_RESERVED0 = 29
const LEDC_TASK_RESERVED1 = 30
const LEDC_TASK_DUTY_SCALE_UPDATE_CH0 = 31
const LEDC_TASK_DUTY_SCALE_UPDATE_CH1 = 32
const LEDC_TASK_DUTY_SCALE_UPDATE_CH2 = 33
const LEDC_TASK_DUTY_SCALE_UPDATE_CH3 = 34
const LEDC_TASK_DUTY_SCALE_UPDATE_CH4 = 35
const LEDC_TASK_DUTY_SCALE_UPDATE_CH5 = 36
const LEDC_TASK_TIMER0_CAP = 37
const LEDC_TASK_TIMER1_CAP = 38
const LEDC_TASK_TIMER2_CAP = 39
const LEDC_TASK_TIMER3_CAP = 40
const LEDC_TASK_SIG_OUT_DIS_CH0 = 41
const LEDC_TASK_SIG_OUT_DIS_CH1 = 42
const LEDC_TASK_SIG_OUT_DIS_CH2 = 43
const LEDC_TASK_SIG_OUT_DIS_CH3 = 44
const LEDC_TASK_SIG_OUT_DIS_CH4 = 45
const LEDC_TASK_SIG_OUT_DIS_CH5 = 46
const LEDC_TASK_OVF_CNT_RST_CH0 = 47
const LEDC_TASK_OVF_CNT_RST_CH1 = 48
const LEDC_TASK_OVF_CNT_RST_CH2 = 49
const LEDC_TASK_OVF_CNT_RST_CH3 = 50
const LEDC_TASK_OVF_CNT_RST_CH4 = 51
const LEDC_TASK_OVF_CNT_RST_CH5 = 52
const LEDC_TASK_TIMER0_RST = 53
const LEDC_TASK_TIMER1_RST = 54
const LEDC_TASK_TIMER2_RST = 55
const LEDC_TASK_TIMER3_RST = 56
const LEDC_TASK_TIMER0_RESUME = 57
const LEDC_TASK_TIMER1_RESUME = 58
const LEDC_TASK_TIMER2_RESUME = 59
const LEDC_TASK_TIMER3_RESUME = 60
const LEDC_TASK_TIMER0_PAUSE = 61
const LEDC_TASK_TIMER1_PAUSE = 62
const LEDC_TASK_TIMER2_PAUSE = 63
const LEDC_TASK_TIMER3_PAUSE = 64
const LEDC_TASK_GAMMA_RESTART_CH0 = 65
const LEDC_TASK_GAMMA_RESTART_CH1 = 66
const LEDC_TASK_GAMMA_RESTART_CH2 = 67
const LEDC_TASK_GAMMA_RESTART_CH3 = 68
const LEDC_TASK_GAMMA_RESTART_CH4 = 69
const LEDC_TASK_GAMMA_RESTART_CH5 = 70
const LEDC_TASK_GAMMA_PAUSE_CH0 = 71
const LEDC_TASK_GAMMA_PAUSE_CH1 = 72
const LEDC_TASK_GAMMA_PAUSE_CH2 = 73
const LEDC_TASK_GAMMA_PAUSE_CH3 = 74
const LEDC_TASK_GAMMA_PAUSE_CH4 = 75
const LEDC_TASK_GAMMA_PAUSE_CH5 = 76
const LEDC_TASK_GAMMA_RESUME_CH0 = 77
const LEDC_TASK_GAMMA_RESUME_CH1 = 78
const LEDC_TASK_GAMMA_RESUME_CH2 = 79
const LEDC_TASK_GAMMA_RESUME_CH3 = 80
const LEDC_TASK_GAMMA_RESUME_CH4 = 81
const LEDC_TASK_GAMMA_RESUME_CH5 = 82
const PCNT_TASK_START = 83
const PCNT_TASK_STOP = 84
const PCNT_TASK_CNT_INC = 85
const PCNT_TASK_CNT_DEC = 86
const PCNT_TASK_CNT_RST = 87
const TIMER0_TASK_CNT_START_TIMER0 = 88
const TIMER1_TASK_CNT_START_TIMER0 = 89
const TIMER0_TASK_ALARM_START_TIMER0 = 90
const TIMER1_TASK_ALARM_START_TIMER0 = 91
const TIMER0_TASK_CNT_STOP_TIMER0 = 92
const TIMER1_TASK_CNT_STOP_TIMER0 = 93
const TIMER0_TASK_CNT_RELOAD_TIMER0 = 94
const TIMER1_TASK_CNT_RELOAD_TIMER0 = 95
const TIMER0_TASK_CNT_CAP_TIMER0 = 96
const TIMER1_TASK_CNT_CAP_TIMER0 = 97
const RMT_TASK_TX_START = 98
const RMT_TASK_TX_STOP = 99
const RMT_TASK_RX_DONE = 100
const RMT_TASK_RX_START = 101
const MCPWM_TASK_CMPR0_A_UP = 102
const MCPWM_TASK_CMPR1_A_UP = 103
const MCPWM_TASK_CMPR2_A_UP = 104
const MCPWM_TASK_CMPR0_B_UP = 105
const MCPWM_TASK_CMPR1_B_UP = 106
const MCPWM_TASK_CMPR2_B_UP = 107
const MCPWM_TASK_GEN_STOP = 108
const MCPWM_TASK_TIMER0_SYN = 109
const MCPWM_TASK_TIMER1_SYN = 110
const MCPWM_TASK_TIMER2_SYN = 111
const MCPWM_TASK_TIMER0_PERIOD_UP = 112
const MCPWM_TASK_TIMER1_PERIOD_UP = 113
const MCPWM_TASK_TIMER2_PERIOD_UP = 114
const MCPWM_TASK_TZ0_OST = 115
const MCPWM_TASK_TZ1_OST = 116
const MCPWM_TASK_TZ2_OST = 117
const MCPWM_TASK_CLR0_OST = 118
const MCPWM_TASK_CLR1_OST = 119
const MCPWM_TASK_CLR2_OST = 120
const MCPWM_TASK_CAP0 = 121
const MCPWM_TASK_CAP1 = 122
const MCPWM_TASK_CAP2 = 123
const ADC_TASK_SAMPLE0 = 124
const ADC_TASK_SAMPLE1 = 125
const ADC_TASK_START0 = 126
const ADC_TASK_STOP0 = 127
const REGDMA_TASK_START0 = 128
const REGDMA_TASK_START1 = 129
const REGDMA_TASK_START2 = 130
const REGDMA_TASK_START3 = 131
const PDMA_TASK_START_TX = 132
const PDMA_TASK_START_RX = 133
const PDMA_TASK_STOP = 134
const TMPSNSR_TASK_START_SAMPLE = 135
const TMPSNSR_TASK_STOP_SAMPLE = 136
const UART_TASK_TX_START0 = 137
const UART_TASK_TX_START1 = 138
const UART_TASK_TX_STOP0 = 139
const UART_TASK_TX_STOP1 = 140
const UART_TASK_RX_START0 = 141
const UART_TASK_RX_START1 = 142
const UART_TASK_RX_STOP0 = 143
const UART_TASK_RX_STOP1 = 144
const SPI_TASK_TX_START = 145
const SPI_TASK_SLAVE_HD = 146
const SPI_TASK_STOP = 147
const I2S_TASK_START_RX = 148
const I2S_TASK_START_TX = 149
const I2S_TASK_STOP_RX = 150
const I2S_TASK_STOP_TX = 151
const I2C_TASK_START_TRANS = 152
const CAN_TASK_TRANS_START = 153
const ULP_TASK_WAKEUP_CPU = 154
const RTC_TASK_START = 155
const RTC_TASK_STOP = 156
const RTC_TASK_CLR = 157
const RTC_TASK_TRIGGERFLW = 158
const GDMA_TASK_IN_START_CH0 = 159
const GDMA_TASK_IN_START_CH1 = 160
const GDMA_TASK_IN_START_CH2 = 161
const GDMA_TASK_OUT_START_CH0 = 162
const GDMA_TASK_OUT_START_CH1 = 163
const GDMA_TASK_OUT_START_CH2 = 164
const PMU_TASK_SLEEP_REQ = 165
