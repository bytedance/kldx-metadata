package fields

type ReferenceField struct {
	FieldBase
	GuideFieldApiName string `json:"guideFieldApiName"`
	FieldApiName      string `json:"fieldApiName"`
}
