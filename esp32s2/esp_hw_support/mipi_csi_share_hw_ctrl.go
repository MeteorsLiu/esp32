package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MipiCsiBrgUserT c.Int

const (
	MIPI_CSI_BRG_USER_CSI     MipiCsiBrgUserT = 0
	MIPI_CSI_BRG_USER_ISP_DVP MipiCsiBrgUserT = 1
	MIPI_CSI_BRG_USER_SHARE   MipiCsiBrgUserT = 2
)
