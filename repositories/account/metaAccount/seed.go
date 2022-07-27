package metaaccount

import "github.com/deyr02/bnzlcrm/graph/model"

var max = 32
var min = 8

var field_1 = &model.NewCustomFieldElement{
	FieldName:      "AccountName",
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

func getPossibleAccountType() []*string {
	var _possibleAccountType []*string
	account1 := "Customer"
	account2 := "Competitor"
	account3 := "Supplier"

	_possibleAccountType = append(_possibleAccountType, &account1)
	_possibleAccountType = append(_possibleAccountType, &account2)
	_possibleAccountType = append(_possibleAccountType, &account3)
	return _possibleAccountType

}

var field_2 = &model.NewCustomFieldElement{
	FieldName:      "AccountType",
	DataType:       "String",
	FieldType:      model.FieldTypeDropDownList,
	IsRequired:     true,
	Visibility:     true,
	SystemFieled:   true,
	MaxValue:       &max2,
	MinValue:       &min2,
	DefaultValue:   "",
	PossibleValues: getPossibleAccountType(),
	FieldOrder:     1,
}

var seedData []*model.NewCustomFieldElement

func getSeedData() []*model.NewCustomFieldElement {
	seedData = append(seedData, field_1)
	seedData = append(seedData, field_2)
	return seedData
}
