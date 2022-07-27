package metaaccount

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
	META_ACCOUNT = "meta_account"
)

type Database struct {
	client *mongo.Client
}

type Meta_Account_Repository interface {
	GetMetaAccountCollection(ctx context.Context) (*model.MetaAccountCollection, error)
	AddNewElement_Meta_Account(ctx context.Context, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaAccountCollection, error)
	ModifyElement_Meta_Account(ctx context.Context, _id string, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaAccountCollection, error)
	DeleteElement_Meta_Account(ctx context.Context, _id string) (string, error)
}

func New_Meta_Activity_Repository(client *mongo.Client) Meta_Account_Repository {
	_client := client
	collection := _client.Database(database.DATABASE_NAME).Collection(META_ACCOUNT)
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

func (db *Database) GetMetaAccountCollection(ctx context.Context) (*model.MetaAccountCollection, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_ACCOUNT)
	cursor := collection.FindOne(ctx, bson.D{})
	var meta_activity *model.MetaAccountCollection
	err := cursor.Decode(&meta_activity)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	return meta_activity, nil
}
func (db *Database) AddNewElement_Meta_Account(ctx context.Context, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaAccountCollection, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_ACCOUNT)
	cursor := collection.FindOne(ctx, bson.D{})
	var meta_account *model.MetaAccountCollection
	err := cursor.Decode(&meta_account)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	if newCustomFieldElement.SystemFieled {
		return nil, &customerror.NewSystemFiled{}
	}
	for _, element := range meta_account.Fields {

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
	meta_account.Fields = append(meta_account.Fields, customFieldElement)
	collection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$set": bson.M{"fields": meta_account.Fields}})
	return meta_account, nil
}
func (db *Database) ModifyElement_Meta_Account(ctx context.Context, _id string, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaAccountCollection, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_ACCOUNT)
	cursor := collection.FindOne(ctx, bson.D{})

	var meta_account *model.MetaAccountCollection
	err := cursor.Decode(&meta_account)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}

	//check for duplicate fieldName
	for _, element := range meta_account.Fields {
		if element.FieldID == _id && element.SystemFieled {
			return nil, &customerror.SystemField{}
		}
		if element.FieldID != _id {
			if strings.EqualFold(element.FieldName, newCustomFieldElement.FieldName) {
				return nil, &customerror.FieldExist{}

			}
		}
	}

	for i := 0; i < len(meta_account.Fields); i++ {
		if meta_account.Fields[i].FieldID == _id {
			meta_account.Fields[i] = &model.CustomFieldElement{
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
	collection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$set": bson.M{"fields": meta_account.Fields}})

	return meta_account, nil
}

func (db *Database) DeleteElement_Meta_Account(ctx context.Context, _id string) (string, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_ACCOUNT)
	cursor := collection.FindOne(ctx, bson.D{})

	var meta_account *model.MetaAccountCollection
	err := cursor.Decode(&meta_account)
	if err != nil {
		return "", err
	}

	//check for duplicate fieldName
	var _fields []*model.CustomFieldElement
	isFieldExist := false
	for _, element := range meta_account.Fields {
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
	meta_account.Fields = _fields
	//updating field
	collection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$set": bson.M{"fields": meta_account.Fields}})

	return "Field deleted", nil
}
