package metaactivity

import (
	"context"
	"strings"

	"github.com/deyr02/bnzlcrm/graph/model"
	customerror "github.com/deyr02/bnzlcrm/repositories/customError"
	"github.com/deyr02/bnzlcrm/repositories/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	META_ACTIVITY = "meta_activity"
)

type Database struct {
	client *mongo.Client
}

type Meta_Activity_Repository interface {
	GetMetaActivityCollection(ctx context.Context) (*model.MetaActivityCollection, error)
	AddNewElement_Meta_Activity(ctx context.Context, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaActivityCollection, error)
	ModifyElement_Meta_Activity(ctx context.Context, _id string, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaActivityCollection, error)
	DeleteElement_Meta_Activity(ctx context.Context, _id string) (string, error)
}

func New_Meta_Activity_Repository(client *mongo.Client) Meta_Activity_Repository {
	_client := client
	collection := _client.Database(database.DATABASE_NAME).Collection(META_ACTIVITY)
	_count, _ := collection.CountDocuments(nil, bson.D{})
	if _count == 0 {
		var newcollection model.MetaUserCollection

		for _, element := range getSeedData() {
			ele := &model.CustomFieldElement{
				FieldID:        database.Guid(),
				FieldName:      element.FieldName,
				DataType:       element.DataType,
				FieldType:      element.FieldType,
				IsRequired:     element.IsRequired,
				Visibility:     element.Visibility,
				SystemFieled:   element.SystemFieled,
				MaxValue:       element.MaxValue,
				MinValue:       element.MinValue,
				DefaultValue:   element.DefaultValue,
				PossibleValues: element.PossibleValues,
				FieldOrder:     element.FieldOrder,
			}

			newcollection.Fields = append(newcollection.Fields, ele)
		}
		_, err := collection.InsertOne(nil, newcollection)
		if err != nil {
			panic(err)
		}
	}
	return &Database{
		client: _client,
	}

}

func (db *Database) GetMetaActivityCollection(ctx context.Context) (*model.MetaActivityCollection, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_ACTIVITY)
	cursor := collection.FindOne(ctx, bson.D{})
	var meta_activity *model.MetaActivityCollection
	err := cursor.Decode(&meta_activity)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	return meta_activity, nil
}

func (db *Database) AddNewElement_Meta_Activity(ctx context.Context, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaActivityCollection, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_ACTIVITY)
	cursor := collection.FindOne(ctx, bson.D{})
	var meta_activity *model.MetaActivityCollection
	err := cursor.Decode(&meta_activity)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	if newCustomFieldElement.SystemFieled {
		return nil, &customerror.NewSystemFiled{}
	}
	for _, element := range meta_activity.Fields {

		if strings.EqualFold(element.FieldName, newCustomFieldElement.FieldName) {
			return nil, &customerror.FieldExist{}
		}
	}

	customFieldElement := &model.CustomFieldElement{
		FieldID:        database.Guid(),
		FieldName:      newCustomFieldElement.FieldName,
		DataType:       newCustomFieldElement.DataType,
		FieldType:      newCustomFieldElement.FieldType,
		IsRequired:     newCustomFieldElement.IsRequired,
		Visibility:     newCustomFieldElement.Visibility,
		SystemFieled:   newCustomFieldElement.SystemFieled,
		MaxValue:       newCustomFieldElement.MaxValue,
		MinValue:       newCustomFieldElement.MinValue,
		DefaultValue:   newCustomFieldElement.DefaultValue,
		PossibleValues: newCustomFieldElement.PossibleValues,
		FieldOrder:     newCustomFieldElement.FieldOrder,
	}
	meta_activity.Fields = append(meta_activity.Fields, customFieldElement)
	collection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$set": bson.M{"fields": meta_activity.Fields}})
	return meta_activity, nil
}
func (db *Database) ModifyElement_Meta_Activity(ctx context.Context, _id string, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaActivityCollection, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_ACTIVITY)
	cursor := collection.FindOne(ctx, bson.D{})

	var meta_activity *model.MetaActivityCollection
	err := cursor.Decode(&meta_activity)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}

	//check for duplicate fieldName
	for _, element := range meta_activity.Fields {
		if element.FieldID == _id && element.SystemFieled {
			return nil, &customerror.SystemField{}
		}
		if element.FieldID != _id {
			if strings.EqualFold(element.FieldName, newCustomFieldElement.FieldName) {
				return nil, &customerror.FieldExist{}

			}
		}

	}

	for i := 0; i < len(meta_activity.Fields); i++ {
		if meta_activity.Fields[i].FieldID == _id {
			meta_activity.Fields[i] = &model.CustomFieldElement{
				FieldID:        _id,
				FieldName:      newCustomFieldElement.FieldName,
				DataType:       newCustomFieldElement.DataType,
				FieldType:      newCustomFieldElement.FieldType,
				IsRequired:     newCustomFieldElement.IsRequired,
				Visibility:     newCustomFieldElement.Visibility,
				SystemFieled:   newCustomFieldElement.SystemFieled,
				MaxValue:       newCustomFieldElement.MaxValue,
				MinValue:       newCustomFieldElement.MinValue,
				DefaultValue:   newCustomFieldElement.DefaultValue,
				PossibleValues: newCustomFieldElement.PossibleValues,
				FieldOrder:     newCustomFieldElement.FieldOrder,
			}
			break
		}
	}

	//updating field
	collection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$set": bson.M{"fields": meta_activity.Fields}})

	return meta_activity, nil
}
func (db *Database) DeleteElement_Meta_Activity(ctx context.Context, _id string) (string, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_ACTIVITY)
	cursor := collection.FindOne(ctx, bson.D{})

	var meta_activity *model.MetaActivityCollection
	err := cursor.Decode(&meta_activity)
	if err != nil {
		return "", err
	}

	//check for duplicate fieldName
	var _fields []*model.CustomFieldElement
	isFieldExist := false
	for _, element := range meta_activity.Fields {
		if element.FieldID == _id {
			isFieldExist = true
		}
		if element.FieldID == _id && element.SystemFieled {
			return "", &customerror.SystemField{}
		}
		if element.FieldID != _id {
			_fields = append(_fields, element)
		}
	}
	if !isFieldExist {
		return "", &customerror.NoRecordFound{}
	}
	meta_activity.Fields = _fields
	//updating field
	collection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$set": bson.M{"fields": meta_activity.Fields}})

	return "Field deleted", nil
}
