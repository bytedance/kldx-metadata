package http

import (
	cConstants "github.com/bytedance/kldx-common/constants"
	cUtils "github.com/bytedance/kldx-common/utils"
	"strings"
)

type UrlPath = string

const (
	GetField  UrlPath = "/metadata/v5/namespaces/:namespace/objects/:objectApiName/fields/:fieldApiName"
	GetFields UrlPath = "/metadata/v5/namespaces/:namespace/objects/:objectApiName/describe"
)

func GetFieldUrlPath(objectApiName string, fieldApiName string) string {
	path := strings.ReplaceAll(GetField, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
	path = strings.ReplaceAll(path, cConstants.ReplaceObjectApiName, objectApiName)
	return strings.ReplaceAll(path, cConstants.ReplaceFieldApiName, fieldApiName)
}

func GetFieldsUrlPath(objectApiName string) string {
	path := strings.ReplaceAll(GetFields, cConstants.ReplaceNamespace, cUtils.GetTenant().Namespace)
	return strings.ReplaceAll(path, cConstants.ReplaceObjectApiName, objectApiName)
}
