package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Registers */
/** Type of pmu_intr_map register
 *  register description
 */

type IntmtxCore0PmuIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of efuse_intr_map register
 *  register description
 */

type IntmtxCore0EfuseIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_rtc_timer_intr_map register
 *  register description
 */

type IntmtxCore0LpRtcTimerIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_ble_timer_intr_map register
 *  register description
 */

type IntmtxCore0LpBleTimerIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_wdt_intr_map register
 *  register description
 */

type IntmtxCore0LpWdtIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_peri_timeout_intr_map register
 *  register description
 */

type IntmtxCore0LpPeriTimeoutIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of lp_apm_m0_intr_map register
 *  register description
 */

type IntmtxCore0LpApmM0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_0_map register
 *  register description
 */

type IntmtxCore0CpuIntrFromCpu0MapRegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_1_map register
 *  register description
 */

type IntmtxCore0CpuIntrFromCpu1MapRegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_2_map register
 *  register description
 */

type IntmtxCore0CpuIntrFromCpu2MapRegT struct {
	Val c.Uint32T
}

/** Type of cpu_intr_from_cpu_3_map register
 *  register description
 */

type IntmtxCore0CpuIntrFromCpu3MapRegT struct {
	Val c.Uint32T
}

/** Type of assist_debug_intr_map register
 *  register description
 */

type IntmtxCore0AssistDebugIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of trace_intr_map register
 *  register description
 */

type IntmtxCore0TraceIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of cache_intr_map register
 *  register description
 */

type IntmtxCore0CacheIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of cpu_peri_timeout_intr_map register
 *  register description
 */

type IntmtxCore0CpuPeriTimeoutIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of bt_mac_intr_map register
 *  register description
 */

type IntmtxCore0BtMacIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of bt_bb_intr_map register
 *  register description
 */

type IntmtxCore0BtBbIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of bt_bb_nmi_map register
 *  register description
 */

type IntmtxCore0BtBbNmiMapRegT struct {
	Val c.Uint32T
}

/** Type of coex_intr_map register
 *  register description
 */

type IntmtxCore0CoexIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of ble_timer_intr_map register
 *  register description
 */

type IntmtxCore0BleTimerIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of ble_sec_intr_map register
 *  register description
 */

type IntmtxCore0BleSecIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of zb_mac_intr_map register
 *  register description
 */

type IntmtxCore0ZbMacIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of gpio_interrupt_pro_map register
 *  register description
 */

type IntmtxCore0GpioInterruptProMapRegT struct {
	Val c.Uint32T
}

/** Type of gpio_interrupt_pro_nmi_map register
 *  register description
 */

type IntmtxCore0GpioInterruptProNmiMapRegT struct {
	Val c.Uint32T
}

/** Type of pau_intr_map register
 *  register description
 */

type IntmtxCore0PauIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of hp_peri_timeout_intr_map register
 *  register description
 */

type IntmtxCore0HpPeriTimeoutIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of hp_apm_m0_intr_map register
 *  register description
 */

type IntmtxCore0HpApmM0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of hp_apm_m1_intr_map register
 *  register description
 */

type IntmtxCore0HpApmM1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of hp_apm_m2_intr_map register
 *  register description
 */

type IntmtxCore0HpApmM2IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of hp_apm_m3_intr_map register
 *  register description
 */

type IntmtxCore0HpApmM3IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of mspi_intr_map register
 *  register description
 */

type IntmtxCore0MspiIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of i2s_intr_map register
 *  register description
 */

type IntmtxCore0I2sIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of uhci0_intr_map register
 *  register description
 */

type IntmtxCore0Uhci0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of uart0_intr_map register
 *  register description
 */

type IntmtxCore0Uart0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of uart1_intr_map register
 *  register description
 */

type IntmtxCore0Uart1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of ledc_intr_map register
 *  register description
 */

type IntmtxCore0LedcIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of can0_intr_map register
 *  register description
 */

type IntmtxCore0Can0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of usb_intr_map register
 *  register description
 */

type IntmtxCore0UsbIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of rmt_intr_map register
 *  register description
 */

type IntmtxCore0RmtIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of i2c_ext0_intr_map register
 *  register description
 */

type IntmtxCore0I2cExt0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of i2c_ext1_intr_map register
 *  register description
 */

type IntmtxCore0I2cExt1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of tg0_t0_intr_map register
 *  register description
 */

type IntmtxCore0Tg0T0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of tg0_wdt_intr_map register
 *  register description
 */

type IntmtxCore0Tg0WdtIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of tg1_t0_intr_map register
 *  register description
 */

type IntmtxCore0Tg1T0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of tg1_wdt_intr_map register
 *  register description
 */

type IntmtxCore0Tg1WdtIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of systimer_target0_intr_map register
 *  register description
 */

type IntmtxCore0SystimerTarget0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of systimer_target1_intr_map register
 *  register description
 */

type IntmtxCore0SystimerTarget1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of systimer_target2_intr_map register
 *  register description
 */

type IntmtxCore0SystimerTarget2IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of apb_adc_intr_map register
 *  register description
 */

type IntmtxCore0ApbAdcIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of pwm_intr_map register
 *  register description
 */

type IntmtxCore0PwmIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of pcnt_intr_map register
 *  register description
 */

type IntmtxCore0PcntIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of parl_io_tx_intr_map register
 *  register description
 */

type IntmtxCore0ParlIoTxIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of parl_io_rx_intr_map register
 *  register description
 */

type IntmtxCore0ParlIoRxIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_in_ch0_intr_map register
 *  register description
 */

type IntmtxCore0DmaInCh0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_in_ch1_intr_map register
 *  register description
 */

type IntmtxCore0DmaInCh1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_in_ch2_intr_map register
 *  register description
 */

type IntmtxCore0DmaInCh2IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_out_ch0_intr_map register
 *  register description
 */

type IntmtxCore0DmaOutCh0IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_out_ch1_intr_map register
 *  register description
 */

type IntmtxCore0DmaOutCh1IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of dma_out_ch2_intr_map register
 *  register description
 */

type IntmtxCore0DmaOutCh2IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of gpspi2_intr_map register
 *  register description
 */

type IntmtxCore0Gpspi2IntrMapRegT struct {
	Val c.Uint32T
}

/** Type of aes_intr_map register
 *  register description
 */

type IntmtxCore0AesIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of sha_intr_map register
 *  register description
 */

type IntmtxCore0ShaIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of rsa_intr_map register
 *  register description
 */

type IntmtxCore0RsaIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of ecc_intr_map register
 *  register description
 */

type IntmtxCore0EccIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of ecdsa_intr_map register
 *  register description
 */

type IntmtxCore0EcdsaIntrMapRegT struct {
	Val c.Uint32T
}

/** Type of int_status_reg_0 register
 *  register description
 */

type IntmtxCore0IntStatusReg0RegT struct {
	Val c.Uint32T
}

/** Type of int_status_reg_1 register
 *  register description
 */

type IntmtxCore0IntStatusReg1RegT struct {
	Val c.Uint32T
}

/** Type of int_status_reg_2 register
 *  register description
 */

type IntmtxCore0IntStatusReg2RegT struct {
	Val c.Uint32T
}

/** Type of clock_gate register
 *  register description
 */

type IntmtxCore0ClockGateRegT struct {
	Val c.Uint32T
}

/** Type of interrupt_reg_date register
 *  register description
 */

type IntmtxCore0InterruptRegDateRegT struct {
	Val c.Uint32T
}

type InterruptMatrixDevT struct {
	PmuIntrMap             IntmtxCore0PmuIntrMapRegT
	EfuseIntrMap           IntmtxCore0EfuseIntrMapRegT
	LpRtcTimerIntrMap      IntmtxCore0LpRtcTimerIntrMapRegT
	LpBleTimerIntrMap      IntmtxCore0LpBleTimerIntrMapRegT
	LpWdtIntrMap           IntmtxCore0LpWdtIntrMapRegT
	LpPeriTimeoutIntrMap   IntmtxCore0LpPeriTimeoutIntrMapRegT
	LpApmM0IntrMap         IntmtxCore0LpApmM0IntrMapRegT
	CpuIntrFromCpu0Map     IntmtxCore0CpuIntrFromCpu0MapRegT
	CpuIntrFromCpu1Map     IntmtxCore0CpuIntrFromCpu1MapRegT
	CpuIntrFromCpu2Map     IntmtxCore0CpuIntrFromCpu2MapRegT
	CpuIntrFromCpu3Map     IntmtxCore0CpuIntrFromCpu3MapRegT
	AssistDebugIntrMap     IntmtxCore0AssistDebugIntrMapRegT
	TraceIntrMap           IntmtxCore0TraceIntrMapRegT
	CacheIntrMap           IntmtxCore0CacheIntrMapRegT
	CpuPeriTimeoutIntrMap  IntmtxCore0CpuPeriTimeoutIntrMapRegT
	BtMacIntrMap           IntmtxCore0BtMacIntrMapRegT
	BtBbIntrMap            IntmtxCore0BtBbIntrMapRegT
	BtBbNmiMap             IntmtxCore0BtBbNmiMapRegT
	CoexIntrMap            IntmtxCore0CoexIntrMapRegT
	BleTimerIntrMap        IntmtxCore0BleTimerIntrMapRegT
	BleSecIntrMap          IntmtxCore0BleSecIntrMapRegT
	ZbMacIntrMap           IntmtxCore0ZbMacIntrMapRegT
	GpioInterruptProMap    IntmtxCore0GpioInterruptProMapRegT
	GpioInterruptProNmiMap IntmtxCore0GpioInterruptProNmiMapRegT
	PauIntrMap             IntmtxCore0PauIntrMapRegT
	HpPeriTimeoutIntrMap   IntmtxCore0HpPeriTimeoutIntrMapRegT
	HpApmM0IntrMap         IntmtxCore0HpApmM0IntrMapRegT
	HpApmM1IntrMap         IntmtxCore0HpApmM1IntrMapRegT
	HpApmM2IntrMap         IntmtxCore0HpApmM2IntrMapRegT
	HpApmM3IntrMap         IntmtxCore0HpApmM3IntrMapRegT
	MspiIntrMap            IntmtxCore0MspiIntrMapRegT
	I2sIntrMap             IntmtxCore0I2sIntrMapRegT
	Uhci0IntrMap           IntmtxCore0Uhci0IntrMapRegT
	Uart0IntrMap           IntmtxCore0Uart0IntrMapRegT
	Uart1IntrMap           IntmtxCore0Uart1IntrMapRegT
	LedcIntrMap            IntmtxCore0LedcIntrMapRegT
	Can0IntrMap            IntmtxCore0Can0IntrMapRegT
	UsbIntrMap             IntmtxCore0UsbIntrMapRegT
	RmtIntrMap             IntmtxCore0RmtIntrMapRegT
	I2cExt0IntrMap         IntmtxCore0I2cExt0IntrMapRegT
	I2cExt1IntrMap         IntmtxCore0I2cExt1IntrMapRegT
	Tg0T0IntrMap           IntmtxCore0Tg0T0IntrMapRegT
	Tg0WdtIntrMap          IntmtxCore0Tg0WdtIntrMapRegT
	Tg1T0IntrMap           IntmtxCore0Tg1T0IntrMapRegT
	Tg1WdtIntrMap          IntmtxCore0Tg1WdtIntrMapRegT
	SystimerTarget0IntrMap IntmtxCore0SystimerTarget0IntrMapRegT
	SystimerTarget1IntrMap IntmtxCore0SystimerTarget1IntrMapRegT
	SystimerTarget2IntrMap IntmtxCore0SystimerTarget2IntrMapRegT
	ApbAdcIntrMap          IntmtxCore0ApbAdcIntrMapRegT
	PwmIntrMap             IntmtxCore0PwmIntrMapRegT
	PcntIntrMap            IntmtxCore0PcntIntrMapRegT
	ParlIoTxIntrMap        IntmtxCore0ParlIoTxIntrMapRegT
	ParlIoRxIntrMap        IntmtxCore0ParlIoRxIntrMapRegT
	DmaInCh0IntrMap        IntmtxCore0DmaInCh0IntrMapRegT
	DmaInCh1IntrMap        IntmtxCore0DmaInCh1IntrMapRegT
	DmaInCh2IntrMap        IntmtxCore0DmaInCh2IntrMapRegT
	DmaOutCh0IntrMap       IntmtxCore0DmaOutCh0IntrMapRegT
	DmaOutCh1IntrMap       IntmtxCore0DmaOutCh1IntrMapRegT
	DmaOutCh2IntrMap       IntmtxCore0DmaOutCh2IntrMapRegT
	Gpspi2IntrMap          IntmtxCore0Gpspi2IntrMapRegT
	AesIntrMap             IntmtxCore0AesIntrMapRegT
	ShaIntrMap             IntmtxCore0ShaIntrMapRegT
	RsaIntrMap             IntmtxCore0RsaIntrMapRegT
	EccIntrMap             IntmtxCore0EccIntrMapRegT
	EcdsaIntrMap           IntmtxCore0EcdsaIntrMapRegT
	IntStatusReg0          IntmtxCore0IntStatusReg0RegT
	IntStatusReg1          IntmtxCore0IntStatusReg1RegT
	IntStatusReg2          IntmtxCore0IntStatusReg2RegT
	ClockGate              IntmtxCore0ClockGateRegT
	Reserved114            [442]c.Uint32T
	InterruptRegDate       IntmtxCore0InterruptRegDateRegT
}
