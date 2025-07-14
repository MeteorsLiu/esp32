package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EcdsaLlParamT c.Int

const (
	ECDSA_PARAM_R   EcdsaLlParamT = 0
	ECDSA_PARAM_S   EcdsaLlParamT = 1
	ECDSA_PARAM_Z   EcdsaLlParamT = 2
	ECDSA_PARAM_QAX EcdsaLlParamT = 3
	ECDSA_PARAM_QAY EcdsaLlParamT = 4
)

type EcdsaLlIntrTypeT c.Int

const (
	ECDSA_INT_PREP_DONE   EcdsaLlIntrTypeT = 0
	ECDSA_INT_SHA_RELEASE EcdsaLlIntrTypeT = 1
)

type EcdsaLlStageT c.Int

const (
	ECDSA_STAGE_START_CALC EcdsaLlStageT = 0
	ECDSA_STAGE_LOAD_DONE  EcdsaLlStageT = 1
	ECDSA_STAGE_GET_DONE   EcdsaLlStageT = 2
)

type EcdsaLlStateT c.Int

const (
	ECDSA_STATE_IDLE EcdsaLlStateT = 0
	ECDSA_STATE_LOAD EcdsaLlStateT = 1
	ECDSA_STATE_GET  EcdsaLlStateT = 2
	ECDSA_STATE_BUSY EcdsaLlStateT = 3
)

type EcdsaLlShaTypeT c.Int

const (
	ECDSA_SHA_224 EcdsaLlShaTypeT = 0
	ECDSA_SHA_256 EcdsaLlShaTypeT = 1
)

type EcdsaLlShaModeT c.Int

const (
	ECDSA_MODE_SHA_START    EcdsaLlShaModeT = 0
	ECDSA_MODE_SHA_CONTINUE EcdsaLlShaModeT = 1
)
