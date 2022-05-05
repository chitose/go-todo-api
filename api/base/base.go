package base

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chitose/todo-api/schema/model"

	"github.com/chitose/todo-api/schema/dao"
)

type ContextInitializerFunc func(r *http.Request) (ctx context.Context)

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// PagedResults results for pages GetAll results.
type PagedResults struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"page_size"`
	Data         interface{} `json:"data"`
	TotalRecords int         `json:"total_records"`
}

type RequestValidatorFunc func(ctx context.Context, r *http.Request, table string, action model.Action) error

var ContextInitializer ContextInitializerFunc
var RequestValidator RequestValidatorFunc

func ValidateRequest(ctx context.Context, r *http.Request, table string, action model.Action) error {
	if RequestValidator != nil {
		return RequestValidator(ctx, r, table, action)
	}

	return nil
}

func WriteJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func ReturnError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	status := 0
	switch err {
	case dao.ErrNotFound:
		status = http.StatusBadRequest
	case dao.ErrUnableToMarshalJSON:
		status = http.StatusBadRequest
	case dao.ErrUpdateFailed:
		status = http.StatusBadRequest
	case dao.ErrInsertFailed:
		status = http.StatusBadRequest
	case dao.ErrDeleteFailed:
		status = http.StatusBadRequest
	case dao.ErrBadParams:
		status = http.StatusBadRequest
	default:
		status = http.StatusBadRequest
	}
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	SendJSON(w, r, er.Code, er)
}

func InitializeContext(r *http.Request) (ctx context.Context) {
	if ContextInitializer != nil {
		ctx = ContextInitializer(r)
	} else {
		ctx = r.Context()
	}
	return ctx
}

func ReadInt(r *http.Request, param string, defaultValue int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return defaultValue, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func ParseUint8(ps map[string]string, key string) (uint8, error) {
	idStr := ps[key]
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return uint8(id), err
	}
	return uint8(id), err
}

func ParseUint16(ps map[string]string, key string) (uint16, error) {
	idStr := ps[key]
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return uint16(id), err
	}
	return uint16(id), err
}

func ParseUint32(ps map[string]string, key string) (uint32, error) {
	idStr := ps[key]
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return uint32(id), err
	}
	return uint32(id), err
}

func ParseUint64(ps map[string]string, key string) (uint64, error) {
	idStr := ps[key]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return uint64(id), err
	}
	return uint64(id), err
}

func ParseInt(ps map[string]string, key string) (int, error) {
	idStr := ps[key]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(id), err
}

func ParseInt8(ps map[string]string, key string) (int8, error) {
	idStr := ps[key]
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return -1, err
	}
	return int8(id), err
}

func ParseInt16(ps map[string]string, key string) (int16, error) {
	idStr := ps[key]
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return -1, err
	}
	return int16(id), err
}

func ParseInt32(ps map[string]string, key string) (int32, error) {
	idStr := ps[key]
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(id), err
}

func ParseInt64(ps map[string]string, key string) (int64, error) {
	idStr := ps[key]
	id, err := strconv.ParseInt(idStr, 10, 54)
	if err != nil {
		return -1, err
	}
	return id, err

}
func ParseString(ps map[string]string, key string) (string, error) {
	idStr := ps[key]
	return idStr, nil
}

func ParseUUID(ps map[string]string, key string) (string, error) {
	idStr := ps[key]
	return idStr, nil
}
