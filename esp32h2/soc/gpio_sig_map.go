package soc

import _ "unsafe"

const EXT_ADC_START_IDX = 0
const LEDC_LS_SIG_OUT0_IDX = 0
const MODEM_DIAG0_IDX = 0
const LEDC_LS_SIG_OUT1_IDX = 1
const MODEM_DIAG1_IDX = 1
const LEDC_LS_SIG_OUT2_IDX = 2
const MODEM_DIAG2_IDX = 2
const LEDC_LS_SIG_OUT3_IDX = 3
const MODEM_DIAG3_IDX = 3
const LEDC_LS_SIG_OUT4_IDX = 4
const MODEM_DIAG4_IDX = 4
const LEDC_LS_SIG_OUT5_IDX = 5
const MODEM_DIAG5_IDX = 5
const U0RXD_IN_IDX = 6
const U0TXD_OUT_IDX = 6
const U0CTS_IN_IDX = 7
const U0RTS_OUT_IDX = 7
const U0DSR_IN_IDX = 8
const U0DTR_OUT_IDX = 8
const U1RXD_IN_IDX = 9
const U1TXD_OUT_IDX = 9
const U1CTS_IN_IDX = 10
const U1RTS_OUT_IDX = 10
const MODEM_DIAG6_IDX = 10
const U1DSR_IN_IDX = 11
const U1DTR_OUT_IDX = 11
const I2S_MCLK_IN_IDX = 12
const I2S_MCLK_OUT_IDX = 12
const I2SO_BCK_IN_IDX = 13
const I2SO_BCK_OUT_IDX = 13
const I2SO_WS_IN_IDX = 14
const I2SO_WS_OUT_IDX = 14
const I2SI_SD_IN_IDX = 15
const I2SO_SD_OUT_IDX = 15
const I2SI_BCK_IN_IDX = 16
const I2SI_BCK_OUT_IDX = 16
const I2SI_WS_IN_IDX = 17
const I2SI_WS_OUT_IDX = 17
const I2SO_SD1_OUT_IDX = 18
const USB_JTAG_TDO_BRIDGE_IDX = 19
const USB_JTAG_TRST_IDX = 19
const CPU_TESTBUS0_IDX = 20
const CPU_TESTBUS1_IDX = 21
const CPU_TESTBUS2_IDX = 22
const CPU_TESTBUS3_IDX = 23
const CPU_TESTBUS4_IDX = 24
const CPU_TESTBUS5_IDX = 25
const CPU_TESTBUS6_IDX = 26
const CPU_TESTBUS7_IDX = 27
const CPU_GPIO_IN0_IDX = 28
const CPU_GPIO_OUT0_IDX = 28
const CPU_GPIO_IN1_IDX = 29
const CPU_GPIO_OUT1_IDX = 29
const CPU_GPIO_IN2_IDX = 30
const CPU_GPIO_OUT2_IDX = 30
const CPU_GPIO_IN3_IDX = 31
const CPU_GPIO_OUT3_IDX = 31
const CPU_GPIO_IN4_IDX = 32
const CPU_GPIO_OUT4_IDX = 32
const CPU_GPIO_IN5_IDX = 33
const CPU_GPIO_OUT5_IDX = 33
const CPU_GPIO_IN6_IDX = 34
const CPU_GPIO_OUT6_IDX = 34
const CPU_GPIO_IN7_IDX = 35
const CPU_GPIO_OUT7_IDX = 35
const USB_JTAG_TCK_IDX = 36
const USB_JTAG_TMS_IDX = 37
const USB_JTAG_TDI_IDX = 38
const USB_JTAG_TDO_IDX = 39
const USB_EXTPHY_VP_IDX = 40
const USB_EXTPHY_OEN_IDX = 40
const USB_EXTPHY_VM_IDX = 41
const USB_EXTPHY_SPEED_IDX = 41
const USB_EXTPHY_RCV_IDX = 42
const USB_EXTPHY_VPO_IDX = 42
const USB_EXTPHY_VMO_IDX = 43
const USB_EXTPHY_SUSPND_IDX = 44
const I2CEXT0_SCL_IN_IDX = 45
const I2CEXT0_SCL_OUT_IDX = 45
const I2CEXT0_SDA_IN_IDX = 46
const I2CEXT0_SDA_OUT_IDX = 46
const PARL_RX_DATA0_IDX = 47
const PARL_TX_DATA0_IDX = 47
const PARL_RX_DATA1_IDX = 48
const PARL_TX_DATA1_IDX = 48
const PARL_RX_DATA2_IDX = 49
const PARL_TX_DATA2_IDX = 49
const PARL_RX_DATA3_IDX = 50
const PARL_TX_DATA3_IDX = 50
const PARL_RX_DATA4_IDX = 51
const PARL_TX_DATA4_IDX = 51
const PARL_RX_DATA5_IDX = 52
const PARL_TX_DATA5_IDX = 52
const PARL_RX_DATA6_IDX = 53
const PARL_TX_DATA6_IDX = 53
const PARL_RX_DATA7_IDX = 54
const PARL_TX_DATA7_IDX = 54
const I2CEXT1_SCL_IN_IDX = 55
const I2CEXT1_SCL_OUT_IDX = 55
const I2CEXT1_SDA_IN_IDX = 56
const I2CEXT1_SDA_OUT_IDX = 56
const CTE_ANT0_IDX = 57
const CTE_ANT1_IDX = 58
const CTE_ANT2_IDX = 59
const CTE_ANT3_IDX = 60
const CTE_ANT4_IDX = 61
const CTE_ANT5_IDX = 62
const FSPICLK_IN_IDX = 63
const FSPICLK_OUT_IDX = 63
const FSPIQ_IN_IDX = 64
const FSPIQ_OUT_IDX = 64
const FSPID_IN_IDX = 65
const FSPID_OUT_IDX = 65
const FSPIHD_IN_IDX = 66
const FSPIHD_OUT_IDX = 66
const FSPIWP_IN_IDX = 67
const FSPIWP_OUT_IDX = 67
const FSPICS0_IN_IDX = 68
const FSPICS0_OUT_IDX = 68
const MODEM_DIAG7_IDX = 68
const PARL_RX_CLK_IN_IDX = 69
const PARL_RX_CLK_OUT_IDX = 69
const PARL_TX_CLK_IN_IDX = 70
const PARL_TX_CLK_OUT_IDX = 70
const RMT_SIG_IN0_IDX = 71
const RMT_SIG_OUT0_IDX = 71
const MODEM_DIAG8_IDX = 71
const RMT_SIG_IN1_IDX = 72
const RMT_SIG_OUT1_IDX = 72
const MODEM_DIAG9_IDX = 72
const TWAI_RX_IDX = 73
const TWAI_TX_IDX = 73
const MODEM_DIAG10_IDX = 73
const TWAI_BUS_OFF_ON_IDX = 74
const MODEM_DIAG11_IDX = 74
const TWAI_CLKOUT_IDX = 75
const MODEM_DIAG12_IDX = 75
const TWAI_STANDBY_IDX = 76
const MODEM_DIAG13_IDX = 76
const CTE_ANT6_IDX = 77
const CTE_ANT7_IDX = 78
const CTE_ANT8_IDX = 79
const CTE_ANT9_IDX = 80
const EXTERN_PRIORITY_I_IDX = 81
const EXTERN_PRIORITY_O_IDX = 81
const EXTERN_ACTIVE_I_IDX = 82
const EXTERN_ACTIVE_O_IDX = 82
const GPIO_SD0_OUT_IDX = 83
const GPIO_SD1_OUT_IDX = 84
const GPIO_SD2_OUT_IDX = 85
const GPIO_SD3_OUT_IDX = 86
const PWM0_SYNC0_IN_IDX = 87
const PWM0_OUT0A_IDX = 87
const MODEM_DIAG14_IDX = 87
const PWM0_SYNC1_IN_IDX = 88
const PWM0_OUT0B_IDX = 88
const MODEM_DIAG15_IDX = 88
const PWM0_SYNC2_IN_IDX = 89
const PWM0_OUT1A_IDX = 89
const MODEM_DIAG16_IDX = 89
const PWM0_F0_IN_IDX = 90
const PWM0_OUT1B_IDX = 90
const MODEM_DIAG17_IDX = 90
const PWM0_F1_IN_IDX = 91
const PWM0_OUT2A_IDX = 91
const MODEM_DIAG18_IDX = 91
const PWM0_F2_IN_IDX = 92
const PWM0_OUT2B_IDX = 92
const MODEM_DIAG19_IDX = 92
const PWM0_CAP0_IN_IDX = 93
const ANT_SEL0_IDX = 93
const PWM0_CAP1_IN_IDX = 94
const ANT_SEL1_IDX = 94
const PWM0_CAP2_IN_IDX = 95
const ANT_SEL2_IDX = 95
const ANT_SEL3_IDX = 96
const SIG_IN_FUNC_97_IDX = 97
const SIG_IN_FUNC97_IDX = 97
const SIG_IN_FUNC_98_IDX = 98
const SIG_IN_FUNC98_IDX = 98
const SIG_IN_FUNC_99_IDX = 99
const SIG_IN_FUNC99_IDX = 99
const SIG_IN_FUNC_100_IDX = 100
const SIG_IN_FUNC100_IDX = 100
const PCNT_SIG_CH0_IN0_IDX = 101
const FSPICS1_OUT_IDX = 101
const MODEM_DIAG20_IDX = 101
const PCNT_SIG_CH1_IN0_IDX = 102
const FSPICS2_OUT_IDX = 102
const MODEM_DIAG21_IDX = 102
const PCNT_CTRL_CH0_IN0_IDX = 103
const FSPICS3_OUT_IDX = 103
const MODEM_DIAG22_IDX = 103
const PCNT_CTRL_CH1_IN0_IDX = 104
const FSPICS4_OUT_IDX = 104
const MODEM_DIAG23_IDX = 104
const PCNT_SIG_CH0_IN1_IDX = 105
const FSPICS5_OUT_IDX = 105
const MODEM_DIAG24_IDX = 105
const PCNT_SIG_CH1_IN1_IDX = 106
const CTE_ANT10_IDX = 106
const PCNT_CTRL_CH0_IN1_IDX = 107
const CTE_ANT11_IDX = 107
const PCNT_CTRL_CH1_IN1_IDX = 108
const CTE_ANT12_IDX = 108
const PCNT_SIG_CH0_IN2_IDX = 109
const CTE_ANT13_IDX = 109
const PCNT_SIG_CH1_IN2_IDX = 110
const CTE_ANT14_IDX = 110
const PCNT_CTRL_CH0_IN2_IDX = 111
const CTE_ANT15_IDX = 111
const PCNT_CTRL_CH1_IN2_IDX = 112
const MODEM_DIAG25_IDX = 112
const PCNT_SIG_CH0_IN3_IDX = 113
const MODEM_DIAG26_IDX = 113
const PCNT_SIG_CH1_IN3_IDX = 114
const SPICLK_OUT_IDX = 114
const PCNT_CTRL_CH0_IN3_IDX = 115
const SPICS0_OUT_IDX = 115
const MODEM_DIAG27_IDX = 115
const PCNT_CTRL_CH1_IN3_IDX = 116
const SPICS1_OUT_IDX = 116
const MODEM_DIAG28_IDX = 116
const GPIO_EVENT_MATRIX_IN0_IDX = 117
const GPIO_TASK_MATRIX_OUT0_IDX = 117
const GPIO_EVENT_MATRIX_IN1_IDX = 118
const GPIO_TASK_MATRIX_OUT1_IDX = 118
const GPIO_EVENT_MATRIX_IN2_IDX = 119
const GPIO_TASK_MATRIX_OUT2_IDX = 119
const GPIO_EVENT_MATRIX_IN3_IDX = 120
const GPIO_TASK_MATRIX_OUT3_IDX = 120
const SPIQ_IN_IDX = 121
const SPIQ_OUT_IDX = 121
const SPID_IN_IDX = 122
const SPID_OUT_IDX = 122
const SPIHD_IN_IDX = 123
const SPIHD_OUT_IDX = 123
const SPIWP_IN_IDX = 124
const SPIWP_OUT_IDX = 124
const CLK_OUT_OUT1_IDX = 125
const MODEM_DIAG29_IDX = 125
const CLK_OUT_OUT2_IDX = 126
const MODEM_DIAG30_IDX = 126
const CLK_OUT_OUT3_IDX = 127
const MODEM_DIAG31_IDX = 127
const SIG_GPIO_OUT_IDX = 128
const GPIO_MAP_DATE_IDX = 0x2201120
