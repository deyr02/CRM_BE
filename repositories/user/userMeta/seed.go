package user

import "github.com/deyr02/bnzlcrm/graph/model"

var max = 32
var min = 8

var field_1 = &model.NewCustomFieldElement{
	FieldName:      "UserName",
	DataType:       "String",
	FieldType:      "TextBox",
	IsRequired:     true,
	Visibility:     true,
	SystemFieled:   true,
	MaxValue:       &max,
	MinValue:       &min,
	DefaultValue:   "",
	PossibleValues: nil,
	FieldOrder:     0,
}

var max2 = 255
var min2 = 6

var field_2 = &model.NewCustomFieldElement{
	FieldName:      "Email",
	DataType:       "String",
	FieldType:      "TextBox",
	IsRequired:     true,
	Visibility:     true,
	SystemFieled:   true,
	MaxValue:       &max2,
	MinValue:       &min2,
	DefaultValue:   "",
	PossibleValues: nil,
	FieldOrder:     1,
}

var field_3 = &model.NewCustomFieldElement{
	FieldName:      "Password",
	DataType:       "String",
	FieldType:      "TextBox",
	IsRequired:     true,
	Visibility:     true,
	SystemFieled:   true,
	MaxValue:       &max2,
	MinValue:       &min2,
	DefaultValue:   "",
	PossibleValues: nil,
	FieldOrder:     2,
}

var seedData []*model.NewCustomFieldElement

func getSeedData() []*model.NewCustomFieldElement {
	seedData = append(seedData, field_1)
	seedData = append(seedData, field_2)
	seedData = append(seedData, field_3)
	return seedData
}
