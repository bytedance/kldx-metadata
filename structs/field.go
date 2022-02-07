package structs

import "encoding/json"

type TypeSetting struct {
	Name     string      `json:"name" binding:"required"`
	Settings interface{} `json:"settings" binding:"required"`
}

type Field struct {
	ID         int64        `json:"id"`
	TenantID   int64        `json:"tenant_id"`
	ObjectID   int64        `json:"object_id"`
	ApiName    string       `json:"api_name"`
	APIAlias   string       `json:"api_alias,omitempty"`
	Label      I18ns        `json:"label"`
	Required   bool         `json:"required"`
	UniqueType int          `json:"unique_type"`
	Type       *TypeSetting `json:"type"`
}

type NestedBooleanSetting struct {
	DefaultValue bool `json:"default_value"`

	Description4True  I18ns `json:"description_if_true"`
	Description4False I18ns `json:"description_if_false"`
}

type RelatedToGlobalOption struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	ApiName string `json:"apiName"`
}

type (
	Options []Option
	Option  struct {
		ApiName  string `json:"id"` // 输出标记 id
		Active   int64  `json:"active"`
		APIAlias string `json:"api_alias,omitempty"`                    // rpc 调用时传递的 ApiAlias
		APIName  string `json:"api_name,omitempty"` // 将 ApiAlias 输出
		// display
		ColorID     string      `json:"color_id"`
		Label       I18ns `json:"name"`
		Description I18ns `json:"description"`
	}
)

type NestedEnumSetting struct {
	IsArray               bool                   `json:"is_array"`
	OptionType            string                 `json:"option_type"`
	Options               *Options               `json:"options"`
	RelatedToGlobalOption RelatedToGlobalOption `json:"related_to_global_option"`
}

//type NestedEnumSetting struct {
//	// property
//	IsArray bool `json:"is_array"`
//
//	// resolve
//	OptionType            string                `json:"option_type"`
//	RelatedToConst        string                `json:"related_to_const"`
//	RelatedToGlobalOption RelatedToGlobalOption `json:"related_to_global_option"`
//	Options               Options               `json:"options"`
//	DefaultValue          []string              `json:"default_value"`
//
//	// display
//	Placeholder I18ns `json:"placeholder"`
//}

type NestedFloatSetting struct {
	// property
	DisplayAsPercentage bool `json:"display_as_percentage"`

	// display
	DecimalPlaces int   `json:"decimal_places"`
	Placeholder   I18ns `json:"placeholder"`
}

type NestedEncryptNumberSetting struct {
	// property
	DisplayAsPercentage bool `json:"display_as_percentage"`

	// display
	DecimalPlaces int `json:"decimal_places"`
}

type NestedTextSetting struct {
	// resolve
	MaxLength int64 `json:"max_length"`

	// display
	Multiline   bool  `json:"multiline"`
	Placeholder I18ns `json:"placeholder"`

	// TODO: regexp_expression 未返回
}

type NestedMultilingualSetting struct {
	// resolve
	MaxLength int64 `json:"max_length"`

	// display
	Multiline   bool  `json:"multiline"`
	Placeholder I18ns `json:"placeholder"`
}

type LookupResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"is_deleted"`
}

type DisplayFields []LookupResponse

type Sort struct {
	FieldApiName string `json:"field"`
	Type         string `json:"type"`
	Direction    string `json:"direction"`
}

type Sorts []*Sort

type NestedLookupSetting struct {
	// property
	IsArray                  bool   `json:"is_array,omitempty"`
	IsHierarchy              bool   `json:"is_hierarchy,omitempty"`
	TargetIsHierarchy        bool   `json:"target_is_hierarchy"`
	ReferencedObjectID       int64  `json:"referenced_object_id"`
	ReferencedObjectAPIName  string `json:"referenced_object_api_name,omitempty"`
	ReferencedObjectAPIAlias string `json:"referenced_object_api_alias,omitempty"`
	RelationID               int64  `json:"relation_id,omitempty"`

	// resolve
	DisplayAsTree bool          `json:"display_as_tree,omitempty"`
	DisplayFields DisplayFields `json:"display_fields,omitempty"`
	DisplayOrder  Sorts         `json:"display_order,omitempty"`

	// display
	DisplayStyle          string `json:"display_style,omitempty"`
	ReferencedObjectLabel I18ns  `json:"referenced_object_label,omitempty"`
}

//type NestedLookupSetting struct {
//	// property
//	IsArray                  bool  `json:"is_array"`
//	IsHierarchy              bool  `json:"is_hierarchy"`
//	TargetIsHierarchy        bool   `json:"target_is_hierarchy"`
//	ReferencedObjectID       int64  `json:"referenced_object_id"`
//	ReferencedObjectApiName  string `json:"referenced_object_api_name"`
//	ReferencedObjectApiAlias string `json:"referenced_object_api_alias,omitempty"`
//	RelationID               int64  `json:"relation_id"`
//
//	// resolve
//	DisplayAsTree bool          `json:"display_as_tree"`
//	DisplayFields DisplayFields `json:"display_fields"`
//	DisplayOrder  Sorts         `json:"display_order"`
//
//	// display
//	DisplayStyle         string `json:"display_style"`
//	ReferencedObjectName string `json:"referenced_object_name"`
//	Placeholder          I18ns  `json:"placeholder"`
//
//	ReferencedObjectLabel I18ns `json:"referenced_object_label"`
//}

type NestedBackLookupSetting struct {
	// property
	IsArray                bool   `json:"is_array"`
	IsInherit              bool   `json:"is_inherit"`
	RelatedToObjectID      int64  `json:"related_to_object_id"`
	RelatedToObjectApiName string `json:"related_to_object_api_name"`
	RelatedToFieldID       int64  `json:"related_to_field_id"`
	RelatedToFieldApiName  string `json:"related_to_field_api_name"`

	// display
	RelatedToObjectName string `json:"related_to_object_name"`
	RelatedToFieldName  string `json:"related_to_field_name"`
}

type AvatarImages map[string]string

type AvatarModel struct {
	Source  string        `json:"source"`
	Image   *AvatarImages `json:"image"`
	Color   string        `json:"color"`
	ColorID string        `json:"color_id" mapstructure:"color_id"`
	Content interface{}   `json:"content"`
}

type LookupWithAvatar struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	I18nName  I18ns        `json:"i18n_name,omitempty"`
	Avatar    *AvatarModel `json:"avatar,omitempty"`
	IsDeleted bool         `json:"is_deleted"`
}

type LazyUserLookup struct {
	value LookupWithAvatar

	loader func() (interface{}, error)
}

type FieldDefinition struct {
	ID             int64          `json:"id" yaml:"id"`
	TenantID       int64          `json:"tenant_id"`
	Column         int64          `json:"column" yaml:"column"`
	Label          I18ns          `json:"label"`
	ApiName        string         `json:"api_name"`
	ApiAlias       string         `json:"api_alias"`
	Namespace      string         `json:"namespace"`
	Desc           I18ns          `json:"description"`
	ObjectID       int64          `json:"object_id"`
	ObjectApiName  string         `json:"object_api_name"`
	ObjectApiAlias string         `json:"object_api_alias"`
	Type           *TypeSetting   `json:"type"`
	Required       bool           `json:"required"`
	UniqueType     int            `json:"unique_type"`
	Domain         int64          `json:"domain"`
	CreatedAt      int64          `json:"created_at" `
	CreatedBy      LazyUserLookup `json:"created_by" `
	UpdatedAt      int64          `json:"updated_at"`
	UpdatedBy      LazyUserLookup `json:"updated_by"`

	//字段是否有查询权限，0无查看权限，1有查看权限，nil时默认有查看权限
	FieldPermission int64 `json:"field_permission,omitempty" mapstructure:"field_permission"`
}

type RefFieldTableDescribe struct {
	ColumnName string `json:"column_name"`

	ObjectTblName     string `json:"object_tbl_name"`
	ObjectTblI18nName string `json:"object_tbl_i18n_name"`
}

type NestedReferenceFieldSetting struct {
	// property
	CurrentLookupFieldID          int64  `json:"current_lookup_field_id"`
	CurrentLookupFieldApiName     string `json:"current_lookup_field_api_name,omitempty"`
	TargetReferenceFieldID        int64  `json:"target_reference_field_id"`
	TargetReferenceFieldApiName   string `json:"target_reference_field_api_name,omitempty"`
	TargetReferencedObjectID      int64  `json:"target_reference_object_id,omitempty"`
	TargetReferencedObjectApiName string `json:"target_referenced_object_api_name,omitempty"`

	// display
	TargetReferenceField *Field `json:"target_reference_field,omitempty"`

	// 支持analytics接入专表专用提供
	TargetRefFieldTblDesc *RefFieldTableDescribe `json:"target_ref_field_tbl_desc,omitempty"`
}

//type NestedReferenceFieldSetting struct {
//	// property
//	CurrentLookupFieldID          int64  `json:"current_lookup_field_id"`
//	CurrentLookupFieldApiName     string `json:"current_lookup_field_api_name"`
//	TargetReferenceFieldID        int64  `json:"target_reference_field_id"`
//	TargetReferenceFieldApiName   string `json:"target_reference_field_api_name"`
//	TargetReferencedObjectID      int64  `json:"target_reference_object_id"`
//	TargetReferencedObjectApiName string `json:"target_referenced_object_api_name"`
//
//	// display
//	CurrentLookupFieldName     string           `json:"current_lookup_field_name"`
//	TargetReferenceFieldName   string           `json:"target_reference_field_name"`
//	TargetReferencedObjectName string           `json:"target_reference_object_name"`
//	TargetReferenceField       *FieldDefinition `json:"target_reference_field"`
//
//	TargetRefFieldTblDesc *RefFieldTableDescribe `json:"target_ref_field_tbl_desc"`
//}

type NestedEmailSetting struct {
	// display
	Placeholder I18ns `json:"placeholder"`
}

type NestedPhoneSetting struct {
	// display
	Placeholder I18ns `json:"placeholder"`
}

type NestedDatetimeSetting struct {
	// display
	Placeholder I18ns `json:"placeholder"`
}

type NestedDateSetting struct {
	// display
	Placeholder I18ns `json:"placeholder"`
}

type NestedAutoNumberSetting struct {
	// resolve
	GenerationMethod    string `json:"generation_method"`
	Digits              int64  `json:"digits"`
	Prefix              string `json:"prefix"`
	Suffix              string `json:"suffix"`
	GenerationSettingID int64  `json:"generation_setting_id"`
	// display
	Placeholder I18ns `json:"placeholder"`
}

type NestedAvatarSetting struct {
	// display
	DisplayStyle string `json:"display_style"`
}

type NestedAttachmentSetting struct {
	// property
	Multiple bool `json:"multiple"`

	// resolve
	AnyType   bool     `json:"any_type,omitempty"`
	MimeTypes []string `json:"mime_types"`
}

type ExpressionField struct {
	Type     string          `json:"type"`
	Settings json.RawMessage `json:"settings"`

	DisplayName        []string `json:"displayName"`
	FieldsDisplayNames []string `json:"fieldsDisplayNames"`
	DisplayLabel       []string `json:"displayLabel,omitempty"`

	settingVal interface{}
	useCache   bool
}

type Expression struct {
	Index    int64           `json:"index"`
	Left     ExpressionField `json:"left"`
	Operator string          `json:"operator"`
	Right    ExpressionField `json:"right"`
}

type Criterion struct {
	Conditions []Expression `json:"conditions"`
	Logic      string       `json:"logic"`

	// 不对外暴露
	//ApiName string `json:"api_name"`
	//BizType string `json:"biz_type"`

	//UseType string `json:"use_type"`
}

type ExtractFilterRecord struct {
	Type  string `json:"type"`
	Index int64  `json:"index"`
}

type NestedCompositeSetting struct {
	// property
	Multiple         bool   `json:"multiple"`
	RelatedToID      int64  `json:"related_to_id,omitempty"`
	RelatedToApiName string `json:"related_to_api_name,omitempty"`
	IsExtract        bool   `json:"is_extract,omitempty"`
	RelationID       int64  `json:"relation_id,omitempty"`

	// ExtractCompositeSetting
	Filter                *Criterion           `json:"filter,omitempty"`
	Sort                  *Sorts               `json:"sort,omitempty"`
	FilterRecord          *ExtractFilterRecord `json:"filter_record,omitempty"`
	RelatedToFieldId      int64                `json:"related_to_field_id,omitempty"`
	RelatedToFieldApiName string               `json:"related_to_field_api_name,omitempty"`

	// display
	RelatedToLabel *I18ns `json:"related_to_label,omitempty"`

	// not in cache
	RelatedToFields []*Field `json:"related_to_fields,omitempty"`
}

//type NestedCompositeSetting struct {
//	// property
//	Multiple         bool   `json:"multiple"`
//	RelatedToID      int64  `json:"related_to_id"`
//	RelatedToApiName string `json:"related_to_api_name"`
//	IsExtract        bool  `json:"is_extract,omitempty"`
//	RelationID       int64  `json:"relation_id"`
//
//	// ExtractCompositeSetting
//	Filter                *Criterion          `json:"filter,omitempty"`
//	Sort                  *Sorts              `json:"sort,omitempty"`
//	FilterRecord          ExtractFilterRecord `json:"filter_record,omitempty"`
//	RelatedToFieldId      int64               `json:"related_to_field_id,omitempty"`
//	RelatedToFieldApiName string              `json:"related_to_field_api_name,omitempty"`
//
//	// display
//	RelatedToLabel I18ns `json:"related_to_label"`
//
//	// not in cache
//	RelatedToFields []FieldDefinition `json:"related_to_fields"`
//
//	RelatedToMultiFields [][]FieldDefinition `json:"-"`
//}

type NestedRichTextSetting struct {
	// resolve
	MaxLength int `json:"max_length"`
}

type LookupWithApiName struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"is_deleted"`
	ApiName   string `json:"api_name"`
}

type NestedConditionSetting struct {
	// property
	TargetObject LookupWithApiName `json:"target_object"`

	LeftFilterFields  interface{} `json:"left_filter_fields"`
	RightFilterFields interface{} `json:"right_filter_fields"`
}

type NestedFormulaSetting struct {
	// property
	MaxLength            int64  `json:"max_length"`
	ReturnType           string `json:"return_type"`
	FormulaVersion       int64  `json:"formula_version"`
	ProcessMode          string `json:"process_mode"`
	HasRefEncryptedField bool   `json:"has_ref_encrypted_field"`

	// resolve
	RelatedToConst string            `json:"ref_to_const"`
	RefRecords     I18ns             `json:"ref_record_ids"`
	RefFieldType   map[string]string `json:"ref_field_type"`
	Formulas       I18ns             `json:"formulas"`
	RefApiNames    I18ns             `json:"ref_api_names"`

	// display
	ReturnFieldMeta map[string]interface{} `json:"return_field_meta"`
	BridgeApiName   []interface{}          `json:"bridge_api_names"`
}

type ObjFields struct {
	Fields []*Field `json:"fields"`
}
