package userCollection

import (
	"context"

	"time"

	"github.com/deyr02/bnzlcrm/graph/model"
	"github.com/deyr02/bnzlcrm/jwt"
	customerror "github.com/deyr02/bnzlcrm/repositories/customError"
	"github.com/deyr02/bnzlcrm/repositories/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	USER_COLLECTION = "user"
)

type Database struct {
	client *mongo.Client
}

type User_Repository interface {
	GetAllUser(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, _id string) (*model.User, error)
	AddNewUser(ctx context.Context, input *model.NewUser) (*model.User, error)
	ModifyUser(ctx context.Context, _id string, input *model.NewUser) (*model.User, error)
	DeleteUser(ctx context.Context, _id string) (string, error)
	Login(ctx context.Context, input *model.Login) (*model.UserDto, error)
}

func New_User_Repository(client *mongo.Client) User_Repository {
	_client := client
	collection := _client.Database(database.DATABASE_NAME).Collection(USER_COLLECTION)
	_count, _ := collection.CountDocuments(context.TODO(), bson.D{})

	if _count == 0 {
		_hashPassword, _ := HashPassword(adminUser.Password)
		_adminId := database.Guid()
		_time := time.Now().String()
		var _properties []*model.ElementValue

		for _, element := range adminUser.Properties {
			var _ele = &model.ElementValue{
				Key:      element.Key,
				DataType: element.DataType,
				Value:    element.Value,
			}
			_properties = append(_properties, _ele)
		}

		adminuser := &model.User{
			UserID:     _adminId,
			RoleID:     "1111111111111111111",
			UserName:   adminUser.UserName,
			Email:      adminUser.Email,
			Password:   _hashPassword,
			CreatedBy:  &_adminId,
			ModifiedBy: nil,
			IsActive:   true,
			SystemUser: true,
			CreatedAt:  &_time,
			Properties: _properties,
		}
		collection.InsertOne(context.TODO(), &adminuser)
	}

	return &Database{
		client: _client,
	}
}

func (db *Database) GetAllUser(ctx context.Context) ([]*model.User, error) {

	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_COLLECTION)
	cursor, err_1 := collection.Find(ctx, bson.D{})

	if err_1 != nil {
		return nil, err_1
	}
	var users []*model.User
	err := cursor.All(ctx, &users)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	return users, nil

}
func (db *Database) GetUserByID(ctx context.Context, _id string) (*model.User, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_COLLECTION)
	cursor := collection.FindOne(ctx, bson.D{{Key: "userid", Value: _id}})
	var user *model.User
	err := cursor.Decode(&user)
	if err != nil {
		return nil, &customerror.NoRecordFound{}
	}
	return user, nil
}
func (db *Database) AddNewUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_COLLECTION)

	//checking username already exists
	cursor := collection.FindOne(ctx, bson.D{{Key: "username", Value: input.UserName}})
	var user *model.User
	cursor.Decode(&user)
	if user != nil {
		return nil, &customerror.UserNameTaken{}
	}
	//checking EmailAddress
	_cursor_1 := collection.FindOne(ctx, bson.D{{Key: "email", Value: input.Email}})
	var user_1 *model.User
	_cursor_1.Decode(&user_1)
	if user_1 != nil {
		return nil, &customerror.EmailAddressTaken{}
	}

	//user name from context
	un, _ := ctx.Value("username").(string)
	_hashPassword, _ := HashPassword(input.Password)
	_userId := database.Guid()
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

	user = &model.User{
		UserID:     _userId,
		RoleID:     input.RoleID,
		UserName:   input.UserName,
		Email:      input.Email,
		Password:   _hashPassword,
		CreatedBy:  &un,
		ModifiedBy: nil,
		IsActive:   true,
		SystemUser: false,
		CreatedAt:  &_time,
		Properties: _properties,
	}

	_, err := collection.InsertOne(ctx, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//it will not modify password
func (db *Database) ModifyUser(ctx context.Context, _id string, input *model.NewUser) (*model.User, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_COLLECTION)

	//checking username already exists
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var users []*model.User
	cursor.All(ctx, &users)
	var _user *model.User
	for _, element := range users {

		if element.UserID != _id {
			if element.UserName == input.UserName {
				return nil, &customerror.UserNameTaken{}
			}
			if element.Email == input.Email {
				return nil, &customerror.EmailAddressTaken{}
			}
		}
		if element.UserID == _id {
			_user = element
		}
	}

	//user not exist
	if _user == nil {
		return nil, &customerror.NoRecordFound{}
	}

	// user as systemUser
	if _user.SystemUser {
		return nil, &customerror.SystemUser{}
	}

	un, _ := ctx.Value("username").(string)
	_time := time.Now().String()
	var properties []*model.ElementValue
	for _, element := range input.Properties {
		var _ele = &model.ElementValue{
			Key:      element.Key,
			DataType: element.DataType,
			Value:    element.Value,
		}
		properties = append(properties, _ele)
	}

	_user.RoleID = input.RoleID
	_user.Email = input.Email
	_user.UserName = input.UserName
	_user.UpdatedAt = &_time
	_user.ModifiedBy = &un
	_user.Properties = properties

	collection.FindOneAndUpdate(ctx, bson.D{{Key: "userid", Value: _id}}, bson.M{
		"$set": bson.M{
			"roleid":     input.RoleID,
			"email":      input.Email,
			"username":   input.UserName,
			"updatedat":  &_time,
			"modifiedby": &un,
			"properties": properties,
		},
	})
	return _user, nil
}

func (db *Database) DeleteUser(ctx context.Context, _id string) (string, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_COLLECTION)

	//checking username already exists
	cursor := collection.FindOne(ctx, bson.D{{Key: "userid", Value: _id}})

	var user *model.User
	cursor.Decode(&user)

	if user == nil {
		return "", &customerror.NoRecordFound{}
	}

	if user.SystemUser {
		return "", &customerror.SystemUser{}
	}

	collection.FindOneAndDelete(ctx, bson.D{{Key: "userid", Value: _id}})
	return "User deleted", nil
}

func (db *Database) Login(ctx context.Context, input *model.Login) (*model.UserDto, error) {
	collection := db.client.Database(database.DATABASE_NAME).Collection(USER_COLLECTION)
	//checking username already exists
	cursor := collection.FindOne(ctx, bson.D{{Key: "username", Value: input.UserName}})

	var user *model.User
	err := cursor.Decode(&user)

	if err != nil {
		return nil, &customerror.WrongUsernameOrPasswordError{}
	}

	if user == nil {
		return nil, &customerror.WrongUsernameOrPasswordError{}
	}
	checkPassword := CheckPasswordHash(input.Password, user.Password)

	if !checkPassword {
		return nil, &customerror.WrongUsernameOrPasswordError{}
	}
	tokenServiceDto := &model.TokenServiceDto{
		UserName:   user.UserName,
		RoleID:     user.RoleID,
		ExpiryDate: time.Now().Add(time.Hour * 24).String(),
	}
	token, erro := jwt.GenerateToken(tokenServiceDto)

	if erro != nil {
		return nil, erro
	}
	userDto := &model.UserDto{
		UserName: user.UserName,
		Token:    token,
		Expiry:   tokenServiceDto.ExpiryDate,
	}

	return userDto, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
