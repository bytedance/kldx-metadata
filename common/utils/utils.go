package utils

import (
	"encoding/json"
	"github.com/bytedance/kldx-common/constants"
	"github.com/bytedance/kldx-common/utils"
	"net/http"
)

func Decode(input interface{}, output interface{}) error {
	bytes, err := json.Marshal(input)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &output); err != nil {
		return err
	}
	return nil
}

func TenantAndUserMiddleware(req *http.Request) error {
	if req == nil || req.Header == nil {
		return nil
	}
	req.Header.Add(constants.HttpHeaderKey_Tenant, utils.GetTenant().Name)
	req.Header.Add(constants.HttpHeaderKey_User, "-1")
	return nil
}
