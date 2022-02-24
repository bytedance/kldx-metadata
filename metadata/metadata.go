package metadata

import (
	"github.com/bytedance/kldx-metadata/object"
)

// IMetadata 元数据读写接口
type IMetadata interface {
	// Object 操作指定对象的元数据信息
	Object(objectApiName string) IObject
}

// IObject 对象读写接口
type IObject interface {
	// GetFields 读取对象的字段元数据信息列表
	GetFields(fields interface{}) error
	// GetField 读取指定字段的元数据信息
	GetField(fieldApiName string, field interface{}) error
}

type MetaData struct{}

func (m MetaData) Object(objectApiName string) IObject {
	return object.Object{ObjectApiName: objectApiName}
}
