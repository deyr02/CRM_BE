package metaactivity

import "github.com/deyr02/bnzlcrm/graph/model"

var max = 32
var min = 8

var field_1 = &model.NewCustomFieldElement{
	FieldName:      "ActivityName",
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

func getPossibleActivityType() []*string {
	var _possibleActivity []*string
	activity1 := "Phone Call"
	activity2 := "Team Meeting"
	activity3 := "Webinar"
	activity4 := "IT Support"

	_possibleActivity = append(_possibleActivity, &activity1)
	_possibleActivity = append(_possibleActivity, &activity2)
	_possibleActivity = append(_possibleActivity, &activity3)
	_possibleActivity = append(_possibleActivity, &activity4)
	return _possibleActivity

}

var field_2 = &model.NewCustomFieldElement{
	FieldName:      "ActiviyType",
	DataType:       "String",
	FieldType:      model.FieldTypeDropDownList,
	IsRequired:     true,
	Visibility:     true,
	SystemFieled:   false,
	MaxValue:       &max2,
	MinValue:       &min2,
	DefaultValue:   "",
	PossibleValues: getPossibleActivityType(),
	FieldOrder:     1,
}

var field_3 = &model.NewCustomFieldElement{
	FieldName:      "Start",
	DataType:       model.DataTypeDate,
	FieldType:      model.FieldTypeDatePicker,
	IsRequired:     true,
	Visibility:     true,
	SystemFieled:   false,
	MaxValue:       &max2,
	MinValue:       &min2,
	DefaultValue:   "",
	PossibleValues: nil,
	FieldOrder:     2,
}

var field_4 = &model.NewCustomFieldElement{
	FieldName:      "Finish",
	DataType:       model.DataTypeDate,
	FieldType:      model.FieldTypeDatePicker,
	IsRequired:     true,
	Visibility:     true,
	SystemFieled:   false,
	MaxValue:       &max2,
	MinValue:       &min2,
	DefaultValue:   "",
	PossibleValues: nil,
	FieldOrder:     3,
}

var seedData []*model.NewCustomFieldElement

func getSeedData() []*model.NewCustomFieldElement {
	seedData = append(seedData, field_1)
	seedData = append(seedData, field_2)
	seedData = append(seedData, field_3)
	seedData = append(seedData, field_4)
	return seedData
}
