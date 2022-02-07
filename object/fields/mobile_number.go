package fields

type MobileNumber struct {
	FieldBase
	Required bool `json:"required"`
	Unique   bool `json:"unique"`
}
