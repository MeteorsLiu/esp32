package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Registers */
/** Type of wifi_mac_intr_map register
 *  register description
 */

type InterruptMatrixWifiMacIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of wifi_mac_nmi_map register
 *  register description
 */

type InterruptMatrixWifiMacNmiMapRegT struct {
	Val c.Uint32T
}

/** Type of wifi_pwr_intr_map register
 *  register description
 */

type InterruptMatrixWifiPwrIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of wifi_bb_intr_map register
 *  register description
 */

type InterruptMatrixWifiBbIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of bt_mac_intr_map register
 *  register description
 */

type InterruptMatrixBtMacIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of bt_bb_intr_map register
 *  register description
 */

type InterruptMatrixBtBbIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of bt_bb_nmi_map register
 *  register description
 */

type InterruptMatrixBtBbNmiMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_timer_intr_map register
 *  register description
 */

type InterruptMatrixLpTimerIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of coex_intr_map register
 *  register description
 */

type InterruptMatrixCoexIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of ble_timer_intr_map register
 *  register description
 */

type InterruptMatrixBleTimerIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of ble_sec_intr_map register
 *  register description
 */

type InterruptMatrixBleSecIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of i2c_mst_intr_map register
 *  register description
 */

type InterruptMatrixI2cMstIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of zb_mac_intr_map register
 *  register description
 */

type InterruptMatrixZbMacIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of pmu_intr_map register
 *  register description
 */

type InterruptMatrixPmuIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of efuse_intr_map register
 *  register description
 */

type InterruptMatrixEfuseIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_rtc_timer_intr_map register
 *  register description
 */

type InterruptMatrixLpRtcTimerIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_uart_intr_map register
 *  register description
 */

type InterruptMatrixLpUartIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_i2c_intr_map register
 *  register description
 */

type InterruptMatrixLpI2cIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_wdt_intr_map register
 *  register description
 */

type InterruptMatrixLpWdtIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_peri_timeout_intr_map register
 *  register description
 */

type InterruptMatrixLpPeriTimeoutIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_apm_m0_intr_map register
 *  register description
 */

type InterruptMatrixLpApmM0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_apm_m1_intr_map register
 *  register description
 */

type InterruptMatrixLpApmM1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_0_map register
 *  register description
 */

type InterruptMatrixCpuIntrFromCpu0MapRegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_1_map register
 *  register description
 */

type InterruptMatrixCpuIntrFromCpu1MapRegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_2_map register
 *  register description
 */

type InterruptMatrixCpuIntrFromCpu2MapRegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_3_map register
 *  register description
 */

type InterruptMatrixCpuIntrFromCpu3MapRegT struct {
	Val c.Uint32T
}

/** Type of assist_debug_intr_map register
 *  register description
 */

type InterruptMatrixAssistDebugIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of trace_intr_map register
 *  register description
 */

type InterruptMatrixTraceIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of cache_intr_map register
 *  register description
 */

type InterruptMatrixCacheIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of cpu_peri_timeout_intr_map register
 *  register description
 */

type InterruptMatrixCpuPeriTimeoutIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of gpio_interrupt_pro_map register
 *  register description
 */

type InterruptMatrixGpioInterruptProMapRegT struct {
	Val c.Uint32T
}

/** Type of gpio_interrupt_pro_nmi_map register
 *  register description
 */

type InterruptMatrixGpioInterruptProNmiMapRegT struct {
	Val c.Uint32T
}

/** Type of pau_intr_map register
 *  register description
 */

type InterruptMatrixPauIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of hp_peri_timeout_intr_map register
 *  register description
 */

type InterruptMatrixHpPeriTimeoutIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of modem_peri_timeout_intr_map register
 *  register description
 */

type InterruptMatrixModemPeriTimeoutIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of hp_apm_m0_intr_map register
 *  register description
 */

type InterruptMatrixHpApmM0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of hp_apm_m1_intr_map register
 *  register description
 */

type InterruptMatrixHpApmM1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of hp_apm_m2_intr_map register
 *  register description
 */

type InterruptMatrixHpApmM2IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of hp_apm_m3_intr_map register
 *  register description
 */

type InterruptMatrixHpApmM3IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_apm0_intr_map register
 *  register description
 */

type InterruptMatrixLpApm0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of mspi_intr_map register
 *  register description
 */

type InterruptMatrixMspiIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of i2s_intr_map register
 *  register description
 */

type InterruptMatrixI2sIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of uhci0_intr_map register
 *  register description
 */

type InterruptMatrixUhci0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of uart0_intr_map register
 *  register description
 */

type InterruptMatrixUart0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of uart1_intr_map register
 *  register description
 */

type InterruptMatrixUart1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of ledc_intr_map register
 *  register description
 */

type InterruptMatrixLedcIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of can0_intr_map register
 *  register description
 */

type InterruptMatrixCan0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of can1_intr_map register
 *  register description
 */

type InterruptMatrixCan1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of usb_intr_map register
 *  register description
 */

type InterruptMatrixUsbIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of rmt_intr_map register
 *  register description
 */

type InterruptMatrixRmtIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of i2c_ext0_intr_map register
 *  register description
 */

type InterruptMatrixI2cExt0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of tg0_t0_intr_map register
 *  register description
 */

type InterruptMatrixTg0T0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of tg0_t1_intr_map register
 *  register description
 */

type InterruptMatrixTg0T1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of tg0_wdt_intr_map register
 *  register description
 */

type InterruptMatrixTg0WdtIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of tg1_t0_intr_map register
 *  register description
 */

type InterruptMatrixTg1T0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of tg1_t1_intr_map register
 *  register description
 */

type InterruptMatrixTg1T1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of tg1_wdt_intr_map register
 *  register description
 */

type InterruptMatrixTg1WdtIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of systimer_target0_intr_map register
 *  register description
 */

type InterruptMatrixSystimerTarget0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of systimer_target1_intr_map register
 *  register description
 */

type InterruptMatrixSystimerTarget1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of systimer_target2_intr_map register
 *  register description
 */

type InterruptMatrixSystimerTarget2IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of apb_adc_intr_map register
 *  register description
 */

type InterruptMatrixApbAdcIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of pwm_intr_map register
 *  register description
 */

type InterruptMatrixPwmIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of pcnt_intr_map register
 *  register description
 */

type InterruptMatrixPcntIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of parl_io_intr_map register
 *  register description
 */

type InterruptMatrixParlIoIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of slc0_intr_map register
 *  register description
 */

type InterruptMatrixSlc0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of slc1_intr_map register
 *  register description
 */

type InterruptMatrixSlc1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_in_ch0_intr_map register
 *  register description
 */

type InterruptMatrixDmaInCh0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_in_ch1_intr_map register
 *  register description
 */

type InterruptMatrixDmaInCh1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_in_ch2_intr_map register
 *  register description
 */

type InterruptMatrixDmaInCh2IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_out_ch0_intr_map register
 *  register description
 */

type InterruptMatrixDmaOutCh0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_out_ch1_intr_map register
 *  register description
 */

type InterruptMatrixDmaOutCh1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_out_ch2_intr_map register
 *  register description
 */

type InterruptMatrixDmaOutCh2IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of gpspi2_intr_map register
 *  register description
 */

type InterruptMatrixGpspi2IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of aes_intr_map register
 *  register description
 */

type InterruptMatrixAesIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of sha_intr_map register
 *  register description
 */

type InterruptMatrixShaIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of rsa_intr_map register
 *  register description
 */

type InterruptMatrixRsaIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of ecc_intr_map register
 *  register description
 */

type InterruptMatrixEccIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of int_status_reg_0 register
 *  register description
 */

type InterruptMatrixIntStatusReg0RegT struct {
	Val c.Uint32T
}

/** Type of int_status_reg_1 register
 *  register description
 */

type InterruptMatrixIntStatusReg1RegT struct {
	Val c.Uint32T
}

/** Type of int_status_reg_2 register
 *  register description
 */

type InterruptMatrixIntStatusReg2RegT struct {
	Val c.Uint32T
}

/** Type of clock_gate register
 *  register description
 */

type InterruptMatrixClockGateRegT struct {
	Val c.Uint32T
}

/** Type of interrupt_reg_date register
 *  register description
 */

type InterruptMatrixInterruptRegDateRegT struct {
	Val c.Uint32T
}

type InterruptMatrixDevT struct {
	WifiMacIntrMap          InterruptMatrixWifiMacIntrMapRegT
	WifiMacNmiMap           InterruptMatrixWifiMacNmiMapRegT
	WifiPwrIntrMap          InterruptMatrixWifiPwrIntrMapRegT
	WifiBbIntrMap           InterruptMatrixWifiBbIntrMapRegT
	BtMacIntrMap            InterruptMatrixBtMacIntrMapRegT
	BtBbIntrMap             InterruptMatrixBtBbIntrMapRegT
	BtBbNmiMap              InterruptMatrixBtBbNmiMapRegT
	LpTimerIntrMap          InterruptMatrixLpTimerIntrMapRegT
	CoexIntrMap             InterruptMatrixCoexIntrMapRegT
	BleTimerIntrMap         InterruptMatrixBleTimerIntrMapRegT
	BleSecIntrMap           InterruptMatrixBleSecIntrMapRegT
	I2cMstIntrMap           InterruptMatrixI2cMstIntrMapRegT
	ZbMacIntrMap            InterruptMatrixZbMacIntrMapRegT
	PmuIntrMap              InterruptMatrixPmuIntrMapRegT
	EfuseIntrMap            InterruptMatrixEfuseIntrMapRegT
	LpRtcTimerIntrMap       InterruptMatrixLpRtcTimerIntrMapRegT
	LpUartIntrMap           InterruptMatrixLpUartIntrMapRegT
	LpI2cIntrMap            InterruptMatrixLpI2cIntrMapRegT
	LpWdtIntrMap            InterruptMatrixLpWdtIntrMapRegT
	LpPeriTimeoutIntrMap    InterruptMatrixLpPeriTimeoutIntrMapRegT
	LpApmM0IntrMap          InterruptMatrixLpApmM0IntrMapRegT
	LpApmM1IntrMap          InterruptMatrixLpApmM1IntrMapRegT
	CpuIntrFromCpu0Map      InterruptMatrixCpuIntrFromCpu0MapRegT
	CpuIntrFromCpu1Map      InterruptMatrixCpuIntrFromCpu1MapRegT
	CpuIntrFromCpu2Map      InterruptMatrixCpuIntrFromCpu2MapRegT
	CpuIntrFromCpu3Map      InterruptMatrixCpuIntrFromCpu3MapRegT
	AssistDebugIntrMap      InterruptMatrixAssistDebugIntrMapRegT
	TraceIntrMap            InterruptMatrixTraceIntrMapRegT
	CacheIntrMap            InterruptMatrixCacheIntrMapRegT
	CpuPeriTimeoutIntrMap   InterruptMatrixCpuPeriTimeoutIntrMapRegT
	GpioInterruptProMap     InterruptMatrixGpioInterruptProMapRegT
	GpioInterruptProNmiMap  InterruptMatrixGpioInterruptProNmiMapRegT
	PauIntrMap              InterruptMatrixPauIntrMapRegT
	HpPeriTimeoutIntrMap    InterruptMatrixHpPeriTimeoutIntrMapRegT
	ModemPeriTimeoutIntrMap InterruptMatrixModemPeriTimeoutIntrMapRegT
	HpApmM0IntrMap          InterruptMatrixHpApmM0IntrMapRegT
	HpApmM1IntrMap          InterruptMatrixHpApmM1IntrMapRegT
	HpApmM2IntrMap          InterruptMatrixHpApmM2IntrMapRegT
	HpApmM3IntrMap          InterruptMatrixHpApmM3IntrMapRegT
	LpApm0IntrMap           InterruptMatrixLpApm0IntrMapRegT
	MspiIntrMap             InterruptMatrixMspiIntrMapRegT
	I2sIntrMap              InterruptMatrixI2sIntrMapRegT
	Uhci0IntrMap            InterruptMatrixUhci0IntrMapRegT
	Uart0IntrMap            InterruptMatrixUart0IntrMapRegT
	Uart1IntrMap            InterruptMatrixUart1IntrMapRegT
	LedcIntrMap             InterruptMatrixLedcIntrMapRegT
	Can0IntrMap             InterruptMatrixCan0IntrMapRegT
	Can1IntrMap             InterruptMatrixCan1IntrMapRegT
	UsbIntrMap              InterruptMatrixUsbIntrMapRegT
	RmtIntrMap              InterruptMatrixRmtIntrMapRegT
	I2cExt0IntrMap          InterruptMatrixI2cExt0IntrMapRegT
	Tg0T0IntrMap            InterruptMatrixTg0T0IntrMapRegT
	Tg0T1IntrMap            InterruptMatrixTg0T1IntrMapRegT
	Tg0WdtIntrMap           InterruptMatrixTg0WdtIntrMapRegT
	Tg1T0IntrMap            InterruptMatrixTg1T0IntrMapRegT
	Tg1T1IntrMap            InterruptMatrixTg1T1IntrMapRegT
	Tg1WdtIntrMap           InterruptMatrixTg1WdtIntrMapRegT
	SystimerTarget0IntrMap  InterruptMatrixSystimerTarget0IntrMapRegT
	SystimerTarget1IntrMap  InterruptMatrixSystimerTarget1IntrMapRegT
	SystimerTarget2IntrMap  InterruptMatrixSystimerTarget2IntrMapRegT
	ApbAdcIntrMap           InterruptMatrixApbAdcIntrMapRegT
	PwmIntrMap              InterruptMatrixPwmIntrMapRegT
	PcntIntrMap             InterruptMatrixPcntIntrMapRegT
	ParlIoIntrMap           InterruptMatrixParlIoIntrMapRegT
	Slc0IntrMap             InterruptMatrixSlc0IntrMapRegT
	Slc1IntrMap             InterruptMatrixSlc1IntrMapRegT
	DmaInCh0IntrMap         InterruptMatrixDmaInCh0IntrMapRegT
	DmaInCh1IntrMap         InterruptMatrixDmaInCh1IntrMapRegT
	DmaInCh2IntrMap         InterruptMatrixDmaInCh2IntrMapRegT
	DmaOutCh0IntrMap        InterruptMatrixDmaOutCh0IntrMapRegT
	DmaOutCh1IntrMap        InterruptMatrixDmaOutCh1IntrMapRegT
	DmaOutCh2IntrMap        InterruptMatrixDmaOutCh2IntrMapRegT
	Gpspi2IntrMap           InterruptMatrixGpspi2IntrMapRegT
	AesIntrMap              InterruptMatrixAesIntrMapRegT
	ShaIntrMap              InterruptMatrixShaIntrMapRegT
	RsaIntrMap              InterruptMatrixRsaIntrMapRegT
	EccIntrMap              InterruptMatrixEccIntrMapRegT
	IntStatusReg0           InterruptMatrixIntStatusReg0RegT
	IntStatusReg1           InterruptMatrixIntStatusReg1RegT
	IntStatusReg2           InterruptMatrixIntStatusReg2RegT
	ClockGate               InterruptMatrixClockGateRegT
	Reserved144             [430]c.Uint32T
	InterruptRegDate        InterruptMatrixInterruptRegDateRegT
}
