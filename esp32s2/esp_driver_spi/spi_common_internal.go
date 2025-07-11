package esp_driver_spi

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const DMA_DESC_MEM_ALIGN_SIZE = 4

type SpiDmaDescT DmaDescriptorAlign4T

// / Attributes of an SPI bus
type SpiBusAttrT struct {
	BusCfg               SpiBusConfigT
	Flags                c.Uint32T
	MaxTransferSz        c.Int
	DmaEnabled           bool
	InternalMemAlignSize c.SizeT
	Lock                 SpiBusLockHandleT
}

type SpiDmaCtxT struct {
	TxDmaChan  SpiDmaChanHandleT
	RxDmaChan  SpiDmaChanHandleT
	DmaDescNum c.Int
	DmadescTx  *SpiDmaDescT
	DmadescRx  *SpiDmaDescT
}

// llgo:type C
type SpiDestroyFuncT func(c.Pointer) EspErrT

/**
 * @brief Alloc DMA channel for SPI
 *
 * @param host_id                  SPI host ID
 * @param dma_chan                 DMA channel to be used
 * @param out_dma_ctx              Actual DMA channel context (if you choose to assign a specific DMA channel, this will be the channel you assigned before)
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_NO_MEM:        No enough memory
 *        - ESP_ERR_NOT_FOUND:     There is no available DMA channel
 */
//go:linkname SpicommonDmaChanAlloc C.spicommon_dma_chan_alloc
func SpicommonDmaChanAlloc(host_id SpiHostDeviceT, dma_chan SpiDmaChanT, out_dma_ctx **SpiDmaCtxT) EspErrT

/**
 * @brief Alloc DMA descriptors for SPI
 *
 * @param dma_ctx                  DMA context returned by `spicommon_dma_chan_alloc`
 * @param[in]  cfg_max_sz          Expected maximum transfer size, in bytes.
 * @param[out] actual_max_sz       Actual max transfer size one transaction can be, in bytes.
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_NO_MEM:        No enough memory
 */
// llgo:link (*SpiDmaCtxT).SpicommonDmaDescAlloc C.spicommon_dma_desc_alloc
func (recv_ *SpiDmaCtxT) SpicommonDmaDescAlloc(cfg_max_sz c.Int, actual_max_sz *c.Int) EspErrT {
	return 0
}

/**
 * Setupt/Configure dma descriptor link list
 *
 * @param dmadesc start of dma descriptor memory
 * @param data    start of data buffer to be configured in
 * @param len     length of data buffer, in byte
 * @param is_rx   if descriptor is for rx/receive direction
 */
// llgo:link (*SpiDmaDescT).SpicommonDmaDescSetupLink C.spicommon_dma_desc_setup_link
func (recv_ *SpiDmaDescT) SpicommonDmaDescSetupLink(data c.Pointer, len c.Int, is_rx bool) {
}

/**
 * @brief Free DMA for SPI
 *
 * @param dma_ctx  spi_dma_ctx_t struct pointer
 *
 * @return
 *        - ESP_OK: On success
 */
// llgo:link (*SpiDmaCtxT).SpicommonDmaChanFree C.spicommon_dma_chan_free
func (recv_ *SpiDmaCtxT) SpicommonDmaChanFree() EspErrT {
	return 0
}

/**
 * @brief Connect a SPI peripheral to GPIO pins
 *
 * This routine is used to connect a SPI peripheral to the IO-pads and DMA channel given in
 * the arguments. Depending on the IO-pads requested, the routing is done either using the
 * IO_mux or using the GPIO matrix.
 *
 * @param host SPI peripheral to be routed
 * @param bus_config Pointer to a spi_bus_config struct detailing the GPIO pins
 * @param flags Combination of SPICOMMON_BUSFLAG_* flags, set to ensure the pins set are capable with some functions:
 *              - ``SPICOMMON_BUSFLAG_MASTER``: Initialize I/O in master mode
 *              - ``SPICOMMON_BUSFLAG_SLAVE``: Initialize I/O in slave mode
 *              - ``SPICOMMON_BUSFLAG_IOMUX_PINS``: Pins set should match the iomux pins of the controller.
 *              - ``SPICOMMON_BUSFLAG_SCLK``, ``SPICOMMON_BUSFLAG_MISO``, ``SPICOMMON_BUSFLAG_MOSI``:
 *                  Make sure SCLK/MISO/MOSI is/are set to a valid GPIO. Also check output capability according to the mode.
 *              - ``SPICOMMON_BUSFLAG_DUAL``: Make sure both MISO and MOSI are output capable so that DIO mode is capable.
 *              - ``SPICOMMON_BUSFLAG_WPHD`` Make sure WP and HD are set to valid output GPIOs.
 *              - ``SPICOMMON_BUSFLAG_QUAD``: Combination of ``SPICOMMON_BUSFLAG_DUAL`` and ``SPICOMMON_BUSFLAG_WPHD``.
 *              - ``SPICOMMON_BUSFLAG_IO4_IO7``: Make sure spi data4 ~ spi data7 are set to valid output GPIOs.
 *              - ``SPICOMMON_BUSFLAG_OCTAL``: Combination of ``SPICOMMON_BUSFLAG_QUAL`` and ``SPICOMMON_BUSFLAG_IO4_IO7``.
 * @param[out] flags_o A SPICOMMON_BUSFLAG_* flag combination of bus abilities will be written to this address.
 *              Leave to NULL if not needed.
 *              - ``SPICOMMON_BUSFLAG_IOMUX_PINS``: The bus is connected to iomux pins.
 *              - ``SPICOMMON_BUSFLAG_SCLK``, ``SPICOMMON_BUSFLAG_MISO``, ``SPICOMMON_BUSFLAG_MOSI``: The bus has
 *                  CLK/MISO/MOSI connected.
 *              - ``SPICOMMON_BUSFLAG_DUAL``: The bus is capable with DIO mode.
 *              - ``SPICOMMON_BUSFLAG_WPHD`` The bus has WP and HD connected.
 *              - ``SPICOMMON_BUSFLAG_QUAD``: Combination of ``SPICOMMON_BUSFLAG_DUAL`` and ``SPICOMMON_BUSFLAG_WPHD``.
 *              - ``SPICOMMON_BUSFLAG_IO4_IO7``: The bus has spi data4 ~ spi data7 connected.
 *              - ``SPICOMMON_BUSFLAG_OCTAL``: Combination of ``SPICOMMON_BUSFLAG_QUAL`` and ``SPICOMMON_BUSFLAG_IO4_IO7``.
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_OK                on success
 */
//go:linkname SpicommonBusInitializeIo C.spicommon_bus_initialize_io
func SpicommonBusInitializeIo(host SpiHostDeviceT, bus_config *SpiBusConfigT, flags c.Uint32T, flags_o *c.Uint32T) EspErrT

/**
 * @brief Free the IO used by a SPI peripheral
 *
 * @param bus_cfg Bus config struct which defines which pins to be used.
 *
 * @return
 *         - ESP_ERR_INVALID_ARG   if parameter is invalid
 *         - ESP_OK                on success
 */
// llgo:link (*SpiBusConfigT).SpicommonBusFreeIoCfg C.spicommon_bus_free_io_cfg
func (recv_ *SpiBusConfigT) SpicommonBusFreeIoCfg() EspErrT {
	return 0
}

/**
 * @brief Initialize a Chip Select pin for a specific SPI peripheral
 *
 * @param host SPI peripheral
 * @param cs_io_num GPIO pin to route
 * @param cs_num CS id to route
 * @param force_gpio_matrix If true, CS will always be routed through the GPIO matrix. If false,
 *                          if the GPIO number allows it, the routing will happen through the IO_mux.
 */
//go:linkname SpicommonCsInitialize C.spicommon_cs_initialize
func SpicommonCsInitialize(host SpiHostDeviceT, cs_io_num c.Int, cs_num c.Int, force_gpio_matrix c.Int)

/**
 * @brief Free a chip select line
 *
 * @param cs_gpio_num CS gpio num to free
 */
//go:linkname SpicommonCsFreeIo C.spicommon_cs_free_io
func SpicommonCsFreeIo(cs_gpio_num c.Int)

/**
 * @brief Check whether all pins used by a host are through IOMUX.
 *
 * @param host SPI peripheral
 *
 * @return false if any pins are through the GPIO matrix, otherwise true.
 */
//go:linkname SpicommonBusUsingIomux C.spicommon_bus_using_iomux
func SpicommonBusUsingIomux(host SpiHostDeviceT) bool

/**
 * @brief Get the IRQ source for a specific SPI host
 *
 * @param host The SPI host
 *
 * @return The hosts IRQ source
 */
//go:linkname SpicommonIrqsourceForHost C.spicommon_irqsource_for_host
func SpicommonIrqsourceForHost(host SpiHostDeviceT) c.Int

/**
 * @brief Get the IRQ source for a specific SPI DMA
 *
 * @param host The SPI host
 *
 * @return The hosts IRQ source
 */
//go:linkname SpicommonIrqdmaSourceForHost C.spicommon_irqdma_source_for_host
func SpicommonIrqdmaSourceForHost(host SpiHostDeviceT) c.Int

// llgo:type C
type DmaworkaroundCbT func(c.Pointer)

/*******************************************************************************
 * Bus attributes
 ******************************************************************************/
/**
 * @brief Set bus lock for the main bus, called by startup code.
 *
 * @param lock The lock to be used by the main SPI bus.
 */
//go:linkname SpiBusMainSetLock C.spi_bus_main_set_lock
func SpiBusMainSetLock(lock SpiBusLockHandleT)

/**
 * @brief Get the attributes of a specified SPI bus.
 *
 * @param host_id The specified host to get attribute
 * @return (Const) Pointer to the attributes
 */
//go:linkname SpiBusGetAttr C.spi_bus_get_attr
func SpiBusGetAttr(host_id SpiHostDeviceT) *SpiBusAttrT

/**
 * @brief Get the dma context of a specified SPI bus.
 *
 * @param host_id The specified host to get attribute
 * @return (Const) Pointer to the dma context
 */
//go:linkname SpiBusGetDmaCtx C.spi_bus_get_dma_ctx
func SpiBusGetDmaCtx(host_id SpiHostDeviceT) *SpiDmaCtxT

/**
 * @brief Register a function to a initialized bus to make it called when deinitializing the bus.
 *
 * @param host_id   The SPI bus to register the destructor.
 * @param f         Destructor to register
 * @param arg       The argument to call the destructor
 * @return Always ESP_OK.
 */
//go:linkname SpiBusRegisterDestroyFunc C.spi_bus_register_destroy_func
func SpiBusRegisterDestroyFunc(host_id SpiHostDeviceT, f SpiDestroyFuncT, arg c.Pointer) EspErrT
