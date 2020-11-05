package errs

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	headerTokenField = "Authorization"
)

var (
	BadParamInBodyErr = errors.New("BAD_PARAM_IN_BODY")
	InternalServerErr = errors.New("INTERNAL_SERVER_ERROR")
)

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(
		[]byte(fmt.Sprintf(`{"code": %d, "error": "%s"}`, http.StatusBadRequest, err)),
	)
}

func InternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(
		[]byte(fmt.Sprintf(`{"code": %d, "error": "%s"}`, http.StatusInternalServerError, InternalServerErr)),
	)
}
