package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EccLlParamT c.Int

const (
	ECC_PARAM_PX EccLlParamT = 0
	ECC_PARAM_PY EccLlParamT = 1
	ECC_PARAM_K  EccLlParamT = 2
)
