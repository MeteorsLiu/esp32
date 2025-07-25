package soc

import _ "unsafe"

const INTERRUPT_CORE0_MAC_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_MAC_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_MAC_INTR_MAP_S = 0
const INTERRUPT_CORE0_MAC_NMI_MAP = 0x0000001F
const INTERRUPT_CORE0_MAC_NMI_MAP_V = 0x1F
const INTERRUPT_CORE0_MAC_NMI_MAP_S = 0
const INTERRUPT_CORE0_PWR_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_PWR_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_PWR_INTR_MAP_S = 0
const INTERRUPT_CORE0_BB_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_BB_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_BB_INT_MAP_S = 0
const INTERRUPT_CORE0_BT_MAC_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_BT_MAC_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_BT_MAC_INT_MAP_S = 0
const INTERRUPT_CORE0_BT_BB_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_BT_BB_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_BT_BB_INT_MAP_S = 0
const INTERRUPT_CORE0_BT_BB_NMI_MAP = 0x0000001F
const INTERRUPT_CORE0_BT_BB_NMI_MAP_V = 0x1F
const INTERRUPT_CORE0_BT_BB_NMI_MAP_S = 0
const INTERRUPT_CORE0_RWBT_IRQ_MAP = 0x0000001F
const INTERRUPT_CORE0_RWBT_IRQ_MAP_V = 0x1F
const INTERRUPT_CORE0_RWBT_IRQ_MAP_S = 0
const INTERRUPT_CORE0_RWBLE_IRQ_MAP = 0x0000001F
const INTERRUPT_CORE0_RWBLE_IRQ_MAP_V = 0x1F
const INTERRUPT_CORE0_RWBLE_IRQ_MAP_S = 0
const INTERRUPT_CORE0_RWBT_NMI_MAP = 0x0000001F
const INTERRUPT_CORE0_RWBT_NMI_MAP_V = 0x1F
const INTERRUPT_CORE0_RWBT_NMI_MAP_S = 0
const INTERRUPT_CORE0_RWBLE_NMI_MAP = 0x0000001F
const INTERRUPT_CORE0_RWBLE_NMI_MAP_V = 0x1F
const INTERRUPT_CORE0_RWBLE_NMI_MAP_S = 0
const INTERRUPT_CORE0_I2C_MST_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_I2C_MST_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_I2C_MST_INT_MAP_S = 0
const INTERRUPT_CORE0_SLC0_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_SLC0_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_SLC0_INTR_MAP_S = 0
const INTERRUPT_CORE0_SLC1_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_SLC1_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_SLC1_INTR_MAP_S = 0
const INTERRUPT_CORE0_APB_CTRL_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_APB_CTRL_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_APB_CTRL_INTR_MAP_S = 0
const INTERRUPT_CORE0_UHCI0_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_UHCI0_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_UHCI0_INTR_MAP_S = 0
const INTERRUPT_CORE0_GPIO_INTERRUPT_PRO_MAP = 0x0000001F
const INTERRUPT_CORE0_GPIO_INTERRUPT_PRO_MAP_V = 0x1F
const INTERRUPT_CORE0_GPIO_INTERRUPT_PRO_MAP_S = 0
const INTERRUPT_CORE0_GPIO_INTERRUPT_PRO_NMI_MAP = 0x0000001F
const INTERRUPT_CORE0_GPIO_INTERRUPT_PRO_NMI_MAP_V = 0x1F
const INTERRUPT_CORE0_GPIO_INTERRUPT_PRO_NMI_MAP_S = 0
const INTERRUPT_CORE0_SPI_INTR_1_MAP = 0x0000001F
const INTERRUPT_CORE0_SPI_INTR_1_MAP_V = 0x1F
const INTERRUPT_CORE0_SPI_INTR_1_MAP_S = 0
const INTERRUPT_CORE0_SPI_INTR_2_MAP = 0x0000001F
const INTERRUPT_CORE0_SPI_INTR_2_MAP_V = 0x1F
const INTERRUPT_CORE0_SPI_INTR_2_MAP_S = 0
const INTERRUPT_CORE0_I2S_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_I2S_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_I2S_INT_MAP_S = 0
const INTERRUPT_CORE0_UART_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_UART_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_UART_INTR_MAP_S = 0
const INTERRUPT_CORE0_UART1_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_UART1_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_UART1_INTR_MAP_S = 0
const INTERRUPT_CORE0_LEDC_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_LEDC_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_LEDC_INT_MAP_S = 0
const INTERRUPT_CORE0_EFUSE_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_EFUSE_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_EFUSE_INT_MAP_S = 0
const INTERRUPT_CORE0_CAN_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_CAN_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_CAN_INT_MAP_S = 0
const INTERRUPT_CORE0_USB_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_USB_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_USB_INTR_MAP_S = 0
const INTERRUPT_CORE0_RTC_CORE_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_RTC_CORE_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_RTC_CORE_INTR_MAP_S = 0
const INTERRUPT_CORE0_RMT_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_RMT_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_RMT_INTR_MAP_S = 0
const INTERRUPT_CORE0_I2C_EXT0_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_I2C_EXT0_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_I2C_EXT0_INTR_MAP_S = 0
const INTERRUPT_CORE0_TIMER_INT1_MAP = 0x0000001F
const INTERRUPT_CORE0_TIMER_INT1_MAP_V = 0x1F
const INTERRUPT_CORE0_TIMER_INT1_MAP_S = 0
const INTERRUPT_CORE0_TIMER_INT2_MAP = 0x0000001F
const INTERRUPT_CORE0_TIMER_INT2_MAP_V = 0x1F
const INTERRUPT_CORE0_TIMER_INT2_MAP_S = 0
const INTERRUPT_CORE0_TG_T0_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_TG_T0_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_TG_T0_INT_MAP_S = 0
const INTERRUPT_CORE0_TG_WDT_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_TG_WDT_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_TG_WDT_INT_MAP_S = 0
const INTERRUPT_CORE0_TG1_T0_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_TG1_T0_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_TG1_T0_INT_MAP_S = 0
const INTERRUPT_CORE0_TG1_WDT_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_TG1_WDT_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_TG1_WDT_INT_MAP_S = 0
const INTERRUPT_CORE0_CACHE_IA_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_CACHE_IA_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_CACHE_IA_INT_MAP_S = 0
const INTERRUPT_CORE0_SYSTIMER_TARGET0_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_SYSTIMER_TARGET0_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_SYSTIMER_TARGET0_INT_MAP_S = 0
const INTERRUPT_CORE0_SYSTIMER_TARGET1_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_SYSTIMER_TARGET1_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_SYSTIMER_TARGET1_INT_MAP_S = 0
const INTERRUPT_CORE0_SYSTIMER_TARGET2_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_SYSTIMER_TARGET2_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_SYSTIMER_TARGET2_INT_MAP_S = 0
const INTERRUPT_CORE0_SPI_MEM_REJECT_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_SPI_MEM_REJECT_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_SPI_MEM_REJECT_INTR_MAP_S = 0
const INTERRUPT_CORE0_ICACHE_PRELOAD_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_ICACHE_PRELOAD_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_ICACHE_PRELOAD_INT_MAP_S = 0
const INTERRUPT_CORE0_ICACHE_SYNC_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_ICACHE_SYNC_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_ICACHE_SYNC_INT_MAP_S = 0
const INTERRUPT_CORE0_APB_ADC_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_APB_ADC_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_APB_ADC_INT_MAP_S = 0
const INTERRUPT_CORE0_DMA_CH0_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_DMA_CH0_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_DMA_CH0_INT_MAP_S = 0
const INTERRUPT_CORE0_DMA_CH1_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_DMA_CH1_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_DMA_CH1_INT_MAP_S = 0
const INTERRUPT_CORE0_DMA_CH2_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_DMA_CH2_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_DMA_CH2_INT_MAP_S = 0
const INTERRUPT_CORE0_RSA_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_RSA_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_RSA_INT_MAP_S = 0
const INTERRUPT_CORE0_AES_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_AES_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_AES_INT_MAP_S = 0
const INTERRUPT_CORE0_SHA_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_SHA_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_SHA_INT_MAP_S = 0
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_0_MAP = 0x0000001F
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_0_MAP_V = 0x1F
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_0_MAP_S = 0
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_1_MAP = 0x0000001F
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_1_MAP_V = 0x1F
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_1_MAP_S = 0
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_2_MAP = 0x0000001F
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_2_MAP_V = 0x1F
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_2_MAP_S = 0
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_3_MAP = 0x0000001F
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_3_MAP_V = 0x1F
const INTERRUPT_CORE0_CPU_INTR_FROM_CPU_3_MAP_S = 0
const INTERRUPT_CORE0_ASSIST_DEBUG_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_ASSIST_DEBUG_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_ASSIST_DEBUG_INTR_MAP_S = 0
const INTERRUPT_CORE0_DMA_APBPERI_PMS_MONITOR_VIOLATE_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_DMA_APBPERI_PMS_MONITOR_VIOLATE_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_DMA_APBPERI_PMS_MONITOR_VIOLATE_INTR_MAP_S = 0
const INTERRUPT_CORE0_CORE_0_IRAM0_PMS_MONITOR_VIOLATE_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_CORE_0_IRAM0_PMS_MONITOR_VIOLATE_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_CORE_0_IRAM0_PMS_MONITOR_VIOLATE_INTR_MAP_S = 0
const INTERRUPT_CORE0_CORE_0_DRAM0_PMS_MONITOR_VIOLATE_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_CORE_0_DRAM0_PMS_MONITOR_VIOLATE_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_CORE_0_DRAM0_PMS_MONITOR_VIOLATE_INTR_MAP_S = 0
const INTERRUPT_CORE0_CORE_0_PIF_PMS_MONITOR_VIOLATE_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_CORE_0_PIF_PMS_MONITOR_VIOLATE_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_CORE_0_PIF_PMS_MONITOR_VIOLATE_INTR_MAP_S = 0
const INTERRUPT_CORE0_CORE_0_PIF_PMS_MONITOR_VIOLATE_SIZE_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_CORE_0_PIF_PMS_MONITOR_VIOLATE_SIZE_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_CORE_0_PIF_PMS_MONITOR_VIOLATE_SIZE_INTR_MAP_S = 0
const INTERRUPT_CORE0_BACKUP_PMS_VIOLATE_INTR_MAP = 0x0000001F
const INTERRUPT_CORE0_BACKUP_PMS_VIOLATE_INTR_MAP_V = 0x1F
const INTERRUPT_CORE0_BACKUP_PMS_VIOLATE_INTR_MAP_S = 0
const INTERRUPT_CORE0_CACHE_CORE0_ACS_INT_MAP = 0x0000001F
const INTERRUPT_CORE0_CACHE_CORE0_ACS_INT_MAP_V = 0x1F
const INTERRUPT_CORE0_CACHE_CORE0_ACS_INT_MAP_S = 0
const INTERRUPT_CORE0_INTR_STATUS_0 = 0xFFFFFFFF
const INTERRUPT_CORE0_INTR_STATUS_0_V = 0xFFFFFFFF
const INTERRUPT_CORE0_INTR_STATUS_0_S = 0
const INTERRUPT_CORE0_INTR_STATUS_1 = 0xFFFFFFFF
const INTERRUPT_CORE0_INTR_STATUS_1_V = 0xFFFFFFFF
const INTERRUPT_CORE0_INTR_STATUS_1_S = 0
const INTERRUPT_CORE0_CLK_EN_V = 0x1
const INTERRUPT_CORE0_CLK_EN_S = 0
const INTERRUPT_CORE0_CPU_INT_ENABLE = 0xFFFFFFFF
const INTERRUPT_CORE0_CPU_INT_ENABLE_V = 0xFFFFFFFF
const INTERRUPT_CORE0_CPU_INT_ENABLE_S = 0
const INTERRUPT_CORE0_CPU_INT_TYPE = 0xFFFFFFFF
const INTERRUPT_CORE0_CPU_INT_TYPE_V = 0xFFFFFFFF
const INTERRUPT_CORE0_CPU_INT_TYPE_S = 0
const INTERRUPT_CORE0_CPU_INT_CLEAR = 0xFFFFFFFF
const INTERRUPT_CORE0_CPU_INT_CLEAR_V = 0xFFFFFFFF
const INTERRUPT_CORE0_CPU_INT_CLEAR_S = 0
const INTERRUPT_CORE0_CPU_INT_EIP_STATUS = 0xFFFFFFFF
const INTERRUPT_CORE0_CPU_INT_EIP_STATUS_V = 0xFFFFFFFF
const INTERRUPT_CORE0_CPU_INT_EIP_STATUS_S = 0
const INTERRUPT_CORE0_CPU_PRI_0_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_0_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_0_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_1_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_1_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_1_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_2_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_2_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_2_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_3_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_3_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_3_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_4_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_4_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_4_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_5_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_5_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_5_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_6_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_6_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_6_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_7_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_7_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_7_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_8_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_8_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_8_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_9_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_9_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_9_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_10_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_10_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_10_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_11_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_11_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_11_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_12_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_12_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_12_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_13_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_13_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_13_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_14_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_14_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_14_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_15_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_15_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_15_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_16_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_16_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_16_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_17_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_17_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_17_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_18_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_18_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_18_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_19_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_19_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_19_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_20_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_20_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_20_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_21_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_21_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_21_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_22_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_22_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_22_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_23_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_23_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_23_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_24_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_24_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_24_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_25_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_25_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_25_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_26_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_26_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_26_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_27_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_27_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_27_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_28_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_28_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_28_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_29_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_29_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_29_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_30_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_30_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_30_MAP_S = 0
const INTERRUPT_CORE0_CPU_PRI_31_MAP = 0x0000000F
const INTERRUPT_CORE0_CPU_PRI_31_MAP_V = 0xF
const INTERRUPT_CORE0_CPU_PRI_31_MAP_S = 0
const INTERRUPT_CORE0_CPU_INT_THRESH = 0x0000000F
const INTERRUPT_CORE0_CPU_INT_THRESH_V = 0xF
const INTERRUPT_CORE0_CPU_INT_THRESH_S = 0
const INTERRUPT_CORE0_INTERRUPT_DATE = 0x0FFFFFFF
const INTERRUPT_CORE0_INTERRUPT_DATE_V = 0xFFFFFFF
const INTERRUPT_CORE0_INTERRUPT_DATE_S = 0
