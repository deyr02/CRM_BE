package auth

import (
	"context"

	"github.com/deyr02/bnzlcrm/graph/model"
	customerror "github.com/deyr02/bnzlcrm/repositories/customError"
	"github.com/deyr02/bnzlcrm/repositories/database"
	userrole "github.com/deyr02/bnzlcrm/repositories/userRole"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USER_COLLECTION = "user"
)

type Database struct {
	client *mongo.Client
}
type UserAuthRepository interface {
	GetUserByUserName(username string) (*model.User, error)
	IsUserAuthorized(roleid string, operation *model.Operation) bool
	IsAdmin(roleId string) bool
}

func NewUserAuthRepository() UserAuthRepository {
	_client := database.CreateConnection()
	return &Database{
		client: _client,
	}
}

func (db *Database) GetUserByUserName(username string) (*model.User, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_COLLECTION)
	cursor := collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}})
	var user *model.User
	err := cursor.Decode(&user)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	return user, nil
}

func (db *Database) IsUserAuthorized(roleid string, operation *model.Operation) bool {
	collection := db.client.Database(database.DATABASE_NAME).Collection(userrole.USER_ROLE)
	cursor := collection.FindOne(context.TODO(), bson.D{{Key: "roleid", Value: roleid}})
	var userrole *model.UserRole
	err := cursor.Decode(&userrole)
	if err != nil {
		return false
	}
	if userrole == nil {
		return false
	}
	if userrole.RoleName == "Admin" {
		return true
	} else {
		for _, ele := range userrole.Operations {
			if ele == *operation {
				return true
			}
		}
		return false
	}
}

func (db *Database) IsAdmin(roleid string) bool {
	collection := db.client.Database(database.DATABASE_NAME).Collection(userrole.USER_ROLE)
	cursor := collection.FindOne(context.TODO(), bson.D{{Key: "roleid", Value: roleid}})
	var userrole *model.UserRole
	err := cursor.Decode(&userrole)
	if err != nil {
		return false
	}
	if userrole == nil {
		return false
	}
	if userrole.RoleName == "Admin" {
		return true
	}
	return false
}
