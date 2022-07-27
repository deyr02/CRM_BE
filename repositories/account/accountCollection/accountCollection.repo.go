package accountcollection

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
	ACCOUNT_COLLECTION = "account"
)

type Database struct {
	client *mongo.Client
}

type Account_Repository interface {
	GetAllAccount(ctx context.Context) ([]*model.Account, error)
	GetAccountByID(ctx context.Context, _id string) (*model.Account, error)
	AddNewAccount(ctx context.Context, input *model.NewAccount) (*model.Account, error)
	ModifyAccount(ctx context.Context, _id string, input *model.NewAccount) (*model.Account, error)
	Deleteaccount(ctx context.Context, _id string) (string, error)
}

func New_Account_Repository(client *mongo.Client) Account_Repository {
	_client := client
	return &Database{
		client: _client,
	}
}

func (db *Database) GetAllAccount(ctx context.Context) ([]*model.Account, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(ACCOUNT_COLLECTION)
	cursor, err_1 := collection.Find(ctx, bson.D{})
	if err_1 != nil {
		return nil, &customerror.NoRecordFound{}
	}
	var accounts []*model.Account
	err := cursor.All(ctx, &accounts)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	return accounts, nil
}

func (db *Database) GetAccountByID(ctx context.Context, _id string) (*model.Account, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(ACCOUNT_COLLECTION)
	cursor := collection.FindOne(ctx, bson.D{{Key: "accountid", Value: _id}})
	var account *model.Account
	err := cursor.Decode(&account)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	return account, nil
}
func (db *Database) AddNewAccount(ctx context.Context, input *model.NewAccount) (*model.Account, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(ACCOUNT_COLLECTION)

	//user name from context
	un, _ := ctx.Value("username").(string)
	_accountID := database.Guid()
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
	account := &model.Account{
		AccountID:   _accountID,
		AccountName: input.AccountName,
		CreatedAt:   &_time,
		CreatedBy:   &un,
		Properties:  _properties,
	}

	_, err := collection.InsertOne(ctx, &account)
	if err != nil {
		return nil, err
	}
	return account, nil
}
func (db *Database) ModifyAccount(ctx context.Context, _id string, input *model.NewAccount) (*model.Account, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(ACCOUNT_COLLECTION)

	//checking username already exists
	cursor := collection.FindOne(ctx, bson.D{{Key: "accountid", Value: _id}})

	var account *model.Account

	err := cursor.Decode(&account)
	if err != nil {
		return nil, err
	}
	if account == nil {
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
	newactivity := &model.Account{
		AccountID:   account.AccountID,
		AccountName: input.AccountName,
		UpdatedAt:   &_time,
		ModifiedBy:  &un,
		Properties:  _properties,
	}

	collection.FindOneAndUpdate(ctx, bson.D{{Key: "accountid", Value: _id}}, bson.M{
		"$set": bson.M{
			"accountname": input.AccountName,
			"updatedat":   &_time,
			"modifiedby":  &un,
			"properties":  _properties,
		},
	})
	return newactivity, nil
}
func (db *Database) Deleteaccount(ctx context.Context, _id string) (string, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(ACCOUNT_COLLECTION)

	//checking username already exists
	cursor := collection.FindOne(ctx, bson.D{{Key: "accountid", Value: _id}})

	var account *model.Account
	err := cursor.Decode(&account)

	if err != nil {
		return "", err
	}

	if account == nil {
		return "", &customerror.NoRecordFound{}
	}

	collection.FindOneAndDelete(ctx, bson.D{{Key: "accountid", Value: _id}})
	return "Account deleted", nil
}
