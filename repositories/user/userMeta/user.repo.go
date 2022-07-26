package user

import (
	"context"
	"log"
	"strings"

	"github.com/deyr02/bnzlcrm/graph/model"
	customerror "github.com/deyr02/bnzlcrm/repositories/customError"
	"github.com/deyr02/bnzlcrm/repositories/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	META_USER = "meta_user"
)

type Database struct {
	client *mongo.Client
}
type User_Meta_Repository interface {
	GetUser_MetaCollection(ctx context.Context) (*model.MetaUserCollection, error)
	AddNewElement_Meta_User(ctx context.Context, newCustomFieldlement *model.NewCustomFieldElement) (*model.MetaUserCollection, error)
	ModifyElement_Meta_User(ctx context.Context, _id string, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaUserCollection, error)
	DeleteElement_Meta_User(ctx context.Context, _id string) (*model.MetaUserCollection, error)
}

func New_User_Meta_repository(client *mongo.Client) User_Meta_Repository {
	_client := client
	collection := _client.Database(database.DATABASE_NAME).Collection(META_USER)
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

func (db *Database) GetUser_MetaCollection(ctx context.Context) (*model.MetaUserCollection, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_USER)
	cursor := collection.FindOne(ctx, bson.D{})

	var meta_user *model.MetaUserCollection
	err := cursor.Decode(&meta_user)
	if err != nil {
		return nil, err
	}
	return meta_user, nil
}
func (db *Database) AddNewElement_Meta_User(ctx context.Context, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaUserCollection, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_USER)
	cursor := collection.FindOne(ctx, bson.D{})
	var meta_user *model.MetaUserCollection
	err := cursor.Decode(&meta_user)
	if err != nil {
		log.Fatal(err)
	}
	if newCustomFieldElement.SystemFieled {
		return nil, &customerror.NewSystemFiled{}
	}
	for _, element := range meta_user.Fields {

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
	meta_user.Fields = append(meta_user.Fields, customFieldElement)

	collection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$set": bson.M{"fields": meta_user.Fields}})
	return meta_user, nil
}

func (db *Database) ModifyElement_Meta_User(ctx context.Context, _id string, newCustomFieldElement *model.NewCustomFieldElement) (*model.MetaUserCollection, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_USER)
	cursor := collection.FindOne(ctx, bson.D{})

	var meta_user *model.MetaUserCollection
	err := cursor.Decode(&meta_user)
	if err != nil {
		return nil, err
	}

	//check for duplicate fieldName
	for _, element := range meta_user.Fields {
		if element.FieldID == _id && element.SystemFieled {
			return nil, &customerror.SystemField{}
		}
		if element.FieldID != _id {
			if strings.EqualFold(element.FieldName, newCustomFieldElement.FieldName) {
				return nil, &customerror.FieldExist{}

			}
		}

	}

	for i := 0; i < len(meta_user.Fields); i++ {
		if meta_user.Fields[i].FieldID == _id {
			meta_user.Fields[i] = &model.CustomFieldElement{
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
	collection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$set": bson.M{"fields": meta_user.Fields}})

	return meta_user, nil
}
func (db *Database) DeleteElement_Meta_User(ctx context.Context, _id string) (*model.MetaUserCollection, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_USER)
	cursor := collection.FindOne(ctx, bson.D{})

	var meta_user *model.MetaUserCollection
	err := cursor.Decode(&meta_user)
	if err != nil {
		return nil, err
	}

	//check for duplicate fieldName
	var _fields []*model.CustomFieldElement
	isFieldExist := false
	for _, element := range meta_user.Fields {
		if element.FieldID == _id {
			isFieldExist = true
		}
		if element.FieldID == _id && element.SystemFieled {
			return nil, &customerror.SystemField{}
		}
		if element.FieldID != _id {
			_fields = append(_fields, element)
		}
	}
	if !isFieldExist {
		return nil, &customerror.NoRecordFound{}
	}
	meta_user.Fields = _fields
	//updating field
	collection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$set": bson.M{"fields": meta_user.Fields}})

	return meta_user, nil
}
