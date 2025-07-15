package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MACSTR = "%02x:%02x:%02x:%02x:%02x:%02x"

type ETSSTATUS c.Int

const (
	ETS_OK      ETSSTATUS = 0
	ETS_FAILED  ETSSTATUS = 1
	ETS_PENDING ETSSTATUS = 2
	ETS_BUSY    ETSSTATUS = 3
	ETS_CANCEL  ETSSTATUS = 4
)

type EtsStatusT ETSSTATUS
type ETSSignal c.Uint32T
type ETSParam c.Uint32T

type ETSEventTag struct {
	Unused [8]uint8
}
type ETSEvent ETSEventTag

// llgo:type C
type ETSTask func(*ETSEvent)

// llgo:type C
type EtsIdleCbT func(c.Pointer)

/**
 * @brief  Ets_printf have two output functions： putc1 and putc2, both of which will be called if need output.
 *         To install putc1, which is defaulted installed as ets_write_char_uart in none silent boot mode, as NULL in silent mode.
 *
 * @param  void (*)(char) p: Output function to install.
 *
 * @return None
 */
//go:linkname EtsInstallPutc1 C.ets_install_putc1
func EtsInstallPutc1(p func(c.Char))

/**
 * @brief  Ets_printf have two output functions： putc1 and putc2, both of which will be called if need output.
 *         To install putc2, which is defaulted installed as NULL.
 *
 * @param  void (*)(char) p: Output function to install.
 *
 * @return None
 */
//go:linkname EtsInstallPutc2 C.ets_install_putc2
func EtsInstallPutc2(p func(c.Char))

// llgo:type C
type ETSTimerFunc func(c.Pointer)

type X_ETSTIMER_ struct {
	TimerNext   *X_ETSTIMER_
	TimerExpire c.Uint32T
	TimerPeriod c.Uint32T
	TimerFunc   *ETSTimerFunc
	TimerArg    c.Pointer
}
type ETSTimer X_ETSTIMER_

// llgo:type C
type EtsIsrT func(c.Pointer)
type STATUS c.Int

const (
	OK      STATUS = 0
	FAIL    STATUS = 1
	PENDING STATUS = 2
	BUSY    STATUS = 3
	CANCEL  STATUS = 4
)
