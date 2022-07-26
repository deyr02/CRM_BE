package activityCollection

import (
	"context"
	"time"

	"github.com/deyr02/bnzlcrm/graph/model"
	customerror "github.com/deyr02/bnzlcrm/repositories/customError"
	"github.com/deyr02/bnzlcrm/repositories/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCtxKey = &contextKey{"username"}

type contextKey struct {
	name string
}

const (
	ACTIVITY_COLLECTION = "activity"
)

type Database struct {
	client *mongo.Client
}

type Activity_Repository interface {
	GetAllActivity(ctx context.Context) ([]*model.Activity, error)
	GetActivityByID(ctx context.Context, _id string) (*model.Activity, error)
	AddNewActivity(ctx context.Context, input *model.NewActivity) (*model.Activity, error)
	ModifyActivity(ctx context.Context, _id string, input *model.NewActivity) (*model.Activity, error)
	DeleteActivity(ctx context.Context, _id string) (string, error)
}

func New_Activity_Repository(client *mongo.Client) Activity_Repository {
	_client := client
	return &Database{
		client: _client,
	}
}

func (db *Database) GetAllActivity(ctx context.Context) ([]*model.Activity, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(ACTIVITY_COLLECTION)
	cursor, err_1 := collection.Find(ctx, bson.D{})
	if err_1 != nil {
		return nil, &customerror.NoRecordFound{}
	}
	var activities []*model.Activity
	err := cursor.All(ctx, &activities)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	return activities, nil

}

func (db *Database) GetActivityByID(ctx context.Context, _id string) (*model.Activity, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(ACTIVITY_COLLECTION)
	cursor := collection.FindOne(ctx, bson.D{{Key: "activityid", Value: _id}})
	var activity *model.Activity
	err := cursor.Decode(&activity)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	return activity, nil
}

func (db *Database) AddNewActivity(ctx context.Context, input *model.NewActivity) (*model.Activity, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(ACTIVITY_COLLECTION)

	//user name from context
	un, _ := ctx.Value("username").(string)
	_activityId := database.Guid()
	_time := time.Now().String()

	var _properties []*model.ElementValue
	for _, element := range input.Properties {
		var _ele = &model.ElementValue{
			Key:      element.Key,
			DataType: element.DataType,
			Value:    element.Value,
		}
		_properties = append(_properties, _ele)
	}
	activity := &model.Activity{
		ActivityID: _activityId,
		CreatedAt:  &_time,
		CreatedBy:  &un,
		Properties: _properties,
	}

	_, err := collection.InsertOne(ctx, &activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (db *Database) ModifyActivity(ctx context.Context, _id string, input *model.NewActivity) (*model.Activity, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(ACTIVITY_COLLECTION)

	//checking username already exists
	cursor := collection.FindOne(ctx, bson.D{{Key: "activityid", Value: _id}})

	var activity *model.Activity

	err := cursor.Decode(&activity)
	if err != nil {
		return nil, err
	}
	if activity == nil {
		return nil, &customerror.NoRecordFound{}
	}

	//user name from context
	un, _ := ctx.Value("username").(string)
	_time := time.Now().String()

	var _properties []*model.ElementValue
	for _, element := range input.Properties {
		var _ele = &model.ElementValue{
			Key:      element.Key,
			DataType: element.DataType,
			Value:    element.Value,
		}
		_properties = append(_properties, _ele)
	}
	newactivity := &model.Activity{
		ActivityID: activity.ActivityID,
		UpdatedAt:  &_time,
		ModifiedBy: &un,
		Properties: _properties,
	}

	collection.FindOneAndUpdate(ctx, bson.D{{Key: "activityid", Value: _id}}, bson.M{
		"$set": bson.M{
			"updatedat":  &_time,
			"modifiedby": &un,
			"properties": _properties,
		},
	})
	return newactivity, nil
}

func (db *Database) DeleteActivity(ctx context.Context, _id string) (string, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(ACTIVITY_COLLECTION)

	//checking username already exists
	cursor := collection.FindOne(ctx, bson.D{{Key: "activityid", Value: _id}})

	var activity *model.Activity
	err := cursor.Decode(&activity)

	if err != nil {
		return "", err
	}

	if activity == nil {
		return "", &customerror.NoRecordFound{}
	}

	collection.FindOneAndDelete(ctx, bson.D{{Key: "activityid", Value: _id}})
	return "activity deleted", nil
}
