package object

import (
	"github.com/bytedance/kldx-common/utils"
	"github.com/bytedance/kldx-metadata/http/openapi"
)

type Object struct {
	ObjectApiName string
}

func (o Object) GetField(fieldApiName string, field interface{}) error {
	f, err := openapi.GetField(o.ObjectApiName, fieldApiName)
	if err != nil {
		return err
	}
	return utils.Decode(f, &field)
}

func (o Object) GetFields(fields interface{}) error {
	fs, err := openapi.GetFields(o.ObjectApiName)
	if err != nil {
		return err
	}
	return utils.Decode(fs, fields)
}
