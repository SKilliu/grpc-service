package errs

import (
	"net/http"

	"github.com/pkg/errors"
)

type ErrResp struct {
	Message string `json:"message" example:"INTERNAL_SERVER_ERROR"`
	Code    int64  `json:"code" example:"500"`
} //@name ErrorResp

func (e ErrResp) ToError() error {
	return errors.New(e.Message)
}

var (
	BadParamInBodyErr = ErrResp{"BAD_PARAM_IN_BODY", http.StatusBadRequest}
	InternalServerErr = ErrResp{"INTERNAL_SERVER_ERROR", http.StatusInternalServerError}
)
