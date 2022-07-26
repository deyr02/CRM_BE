package userrole

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
	USER_ROLE = "user_role"
)

type Database struct {
	client *mongo.Client
}
type User_Role_Repository interface {
	GetAllUserRole(ctx context.Context) ([]*model.UserRole, error)
	GetUserRoleByID(ctx context.Context, _id string) (*model.UserRole, error)
	AddNewUserRole(ctx context.Context, input model.NewUserRole) (*model.UserRole, error)
	ModifyUserRole(ctx context.Context, _id string, input *model.NewUserRole) (*model.UserRole, error)
	DeleUserRole(ctx context.Context, _id string) (string, error)
}

func New_User_Role_repository(client *mongo.Client) User_Role_Repository {
	_client := client
	collection := _client.Database(database.DATABASE_NAME).Collection(USER_ROLE)
	_count, _ := collection.CountDocuments(nil, bson.D{})

	if _count == 0 {
		var adminRole model.UserRole
		var itsupportRole model.UserRole

		adminRole.RoleID = "1111111111111111111"
		adminRole.RoleName = AdminUserRole.RoleName
		adminRole.SystemRole = true
		adminRole.Operations = AdminUserRole.Operations

		_, err := collection.InsertOne(nil, adminRole)
		if err != nil {
			panic(err)
		}

		itsupportRole.RoleID = "1111111111111111112"
		itsupportRole.RoleName = ITSupportUserRole.RoleName
		itsupportRole.SystemRole = false
		itsupportRole.Operations = ITSupportUserRole.Operations

		_, err_1 := collection.InsertOne(nil, itsupportRole)
		if err_1 != nil {
			panic(err)
		}
	}
	return &Database{
		client: _client,
	}
}

func (db *Database) GetAllUserRole(ctx context.Context) ([]*model.UserRole, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_ROLE)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var user_role []*model.UserRole
	err_1 := cursor.All(ctx, &user_role)
	if err_1 != nil {
		return nil, err_1
	}
	return user_role, nil
}

func (db *Database) GetUserRoleByID(ctx context.Context, _id string) (*model.UserRole, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_ROLE)
	cursor := collection.FindOne(ctx, bson.D{{Key: "roleid", Value: _id}})

	var user_role *model.UserRole
	err_1 := cursor.Decode(&user_role)
	if err_1 != nil {
		return nil, &customerror.NoRecordFound{}
	}
	return user_role, nil
}

func (db *Database) AddNewUserRole(ctx context.Context, input model.NewUserRole) (*model.UserRole, error) {

	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_ROLE)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var user_roles []*model.UserRole
	err_1 := cursor.All(ctx, &user_roles)
	if err_1 != nil {
		return nil, &customerror.NoRecordFound{}
	}

	for _, element := range user_roles {
		if strings.EqualFold(element.RoleName, input.RoleName) {
			return nil, &customerror.UserRoleExist{}
		}
	}

	_user_role := &model.UserRole{
		RoleID:     database.Guid(),
		RoleName:   input.RoleName,
		SystemRole: false,
		Operations: input.Operations,
	}
	collection.InsertOne(ctx, _user_role)
	return _user_role, nil
}

func (db *Database) ModifyUserRole(ctx context.Context, _id string, input *model.NewUserRole) (*model.UserRole, error) {

	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_ROLE)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var user_roles []*model.UserRole
	err_1 := cursor.All(ctx, &user_roles)
	if err_1 != nil {
		return nil, &customerror.NoRecordFound{}
	}

	var _user_role *model.UserRole
	//checking if role exists
	for _, element := range user_roles {
		if element.RoleID != _id {
			if strings.EqualFold(element.RoleName, input.RoleName) {
				return nil, &customerror.UserRoleExist{}
			}
		}
		if element.RoleID == _id {
			_user_role = element
		}

	}
	if _user_role == nil {
		return nil, &customerror.NoRecordFound{}
	}
	//checking if trying to modify systemRole
	if _user_role.SystemRole {
		return nil, &customerror.SystemRole{}
	}

	_user_role.RoleName = input.RoleName
	_user_role.Operations = input.Operations

	collection.FindOneAndUpdate(ctx, bson.D{{Key: "roleid", Value: _id}}, bson.M{"$set": bson.M{"rolename": input.RoleName, "operations": input.Operations}})
	return _user_role, nil
}

func (db *Database) DeleUserRole(ctx context.Context, _id string) (string, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_ROLE)
	user_role, err := db.GetUserRoleByID(ctx, _id)
	if err != nil {
		return "", err
	}
	if user_role.SystemRole {
		return "", &customerror.SystemRole{}
	}
	collection.FindOneAndDelete(ctx, bson.D{{Key: "roleid", Value: _id}})
	return "Role Deleted", nil
}
