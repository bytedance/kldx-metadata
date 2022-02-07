package fields

type Number struct {
	FieldBase
	Required            bool `json:"required"`
	Unique              bool `json:"unique"`
	DisplayAsPercentage bool `json:"displayAsPercentage"`
	DecimalPlacesNumber int  `json:"decimalPlacesNumber"`
}
