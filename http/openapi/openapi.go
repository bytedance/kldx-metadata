package openapi

import (
	"encoding/json"
	cExceptions "github.com/bytedance/kldx-common/exceptions"
	cHttp "github.com/bytedance/kldx-common/http"
	"github.com/bytedance/kldx-metadata/http"
	"github.com/bytedance/kldx-metadata/parser"
	"github.com/bytedance/kldx-metadata/structs"
	"github.com/tidwall/gjson"
)

type StatusCode = string

const (
	FileDownload StatusCode = ""
	Success      StatusCode = "0"

	InternalError  StatusCode = "k_ec_000001"
	NoTenantID     StatusCode = "k_ec_000002"
	NoUserID       StatusCode = "k_ec_000003"
	UnknownError   StatusCode = "k_ec_000004"
	OpUnknownError StatusCode = "k_op_ec_00001"
	SystemBusy     StatusCode = "k_op_ec_20001"
	SystemError    StatusCode = "k_op_ec_20002"
	RateLimitError StatusCode = "k_op_ec_20003"
	TokenExpire    StatusCode = "k_ident_013000"
	IllegalToken   StatusCode = "k_ident_013001"
	MissingToken   StatusCode = "k_op_ec_10205"
)

func errorWrapper(body []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, cExceptions.ErrorWrap(err)
	}

	code := gjson.GetBytes(body, "code").String()
	msg := gjson.GetBytes(body, "msg").String()
	switch code {
	case FileDownload:
		return body, nil
	case Success:
		data := gjson.GetBytes(body, "data")
		if data.Type == gjson.String {
			return []byte(data.Str), nil
		}
		return []byte(data.Raw), nil
	case InternalError, NoTenantID, NoUserID, UnknownError,
		OpUnknownError, SystemBusy, SystemError, RateLimitError,
		TokenExpire, IllegalToken, MissingToken:
		return nil, cExceptions.InternalError("request openapi failed, code: %s, msg: %s", code, msg)
	default:
		return nil, cExceptions.InvalidParamError("request openapi failed, code: %s, msg: %s", code, msg)
	}
}

func GetField(objectApiName, filedApiName string) (interface{}, error) {
	data, err := errorWrapper(http.GetOpenapiClient().Get(http.GetFieldUrlPath(objectApiName, filedApiName), nil, cHttp.AppTokenMiddleware, cHttp.TenantAndUserMiddleware))
	if err != nil {
		return nil, err
	}
	field := &structs.Field{}
	if err := json.Unmarshal(data, &field); err != nil {
		return nil, err
	}
	return parser.ParseField(field)
}

func GetFields(objectApiName string) (map[string]interface{}, error) {
	data, err := errorWrapper(http.GetOpenapiClient().Get(http.GetFieldsUrlPath(objectApiName), nil, cHttp.AppTokenMiddleware, cHttp.TenantAndUserMiddleware))
	if err != nil {
		return nil, err
	}
	obj := &structs.ObjFields{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}
	return parser.ParseFields(obj.Fields)
}
