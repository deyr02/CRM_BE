package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/deyr02/bnzlcrm/graph/generated"
	"github.com/deyr02/bnzlcrm/graph/model"
	"github.com/deyr02/bnzlcrm/repositories/activity/activityCollection"
	metaactivity "github.com/deyr02/bnzlcrm/repositories/activity/metaActivity"
	"github.com/deyr02/bnzlcrm/repositories/database"
	"github.com/deyr02/bnzlcrm/repositories/user/userCollection"
	user "github.com/deyr02/bnzlcrm/repositories/user/userMeta"
	userrole "github.com/deyr02/bnzlcrm/repositories/userRole"
)

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///---------------------------------------- Repositories --------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
var client = database.CreateConnection()
var userRepository user.User_Meta_Repository = user.New_User_Meta_repository(client)
var UserRoleRepository userrole.User_Role_Repository = userrole.New_User_Role_repository(client)
var userCollectionRepository userCollection.User_Repository = userCollection.New_User_Repository(client)
var metaActivityRepository metaactivity.Meta_Activity_Repository = metaactivity.New_Meta_Activity_Repository(client)
var activityRepository activityCollection.Activity_Repository = activityCollection.New_Activity_Repository(client)

///
///
///
///
///
///
///
///
///
///
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///----------------------------------------- Mutations ----------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------

///--------------------------------------------------------------------------------------------------------
///------------------------------------Meta User Collection------------------------------------------------
///--------------------------------------------------------------------------------------------------------

// AddNewElementMetaUser is the resolver for the AddNewElement_Meta_User field.
func (r *mutationResolver) AddNewElementMetaUser(ctx context.Context, input *model.NewCustomFieldElement) (*model.MetaUserCollection, error) {
	return userRepository.AddNewElement_Meta_User(ctx, input)
}

// ModifyElementMetaUser is the resolver for the ModifyElement_Meta_user field.
func (r *mutationResolver) ModifyElementMetaUser(ctx context.Context, id *string, input *model.NewCustomFieldElement) (*model.MetaUserCollection, error) {
	return userRepository.ModifyElement_Meta_User(ctx, *id, input)
}

// DeleteElementMetaUser is the resolver for the DeleteElement_Meta_user field.
func (r *mutationResolver) DeleteElementMetaUser(ctx context.Context, id *string) (*model.MetaUserCollection, error) {
	return userRepository.DeleteElement_Meta_User(ctx, *id)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------Meta Activity Collection------------------------------------------------
///--------------------------------------------------------------------------------------------------------

// AddNewElementMetaActivity is the resolver for the AddNewElement_Meta_Activity field.
func (r *mutationResolver) AddNewElementMetaActivity(ctx context.Context, input *model.NewCustomFieldElement) (*model.MetaActivityCollection, error) {
	return metaActivityRepository.AddNewElement_Meta_Activity(ctx, input)
}

// ModifyElementMetaActivity is the resolver for the ModifyElement_Meta_Activity field.
func (r *mutationResolver) ModifyElementMetaActivity(ctx context.Context, id *string, input *model.NewCustomFieldElement) (*model.MetaActivityCollection, error) {
	return metaActivityRepository.ModifyElement_Meta_Activity(ctx, *id, input)
}

// DeleteElementMetaActivity is the resolver for the DeleteElement_Meta_Activity field.
func (r *mutationResolver) DeleteElementMetaActivity(ctx context.Context, id *string) (string, error) {
	return metaActivityRepository.DeleteElement_Meta_Activity(ctx, *id)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------Meta Account Collection------------------------------------------------
///--------------------------------------------------------------------------------------------------------
// AddNewElementMetaAccount is the resolver for the AddNewElement_Meta_Account field.
func (r *mutationResolver) AddNewElementMetaAccount(ctx context.Context, input *model.NewCustomFieldElement) (*model.MetaAccountCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

// ModifyElementMetaAccount is the resolver for the ModifyElement_Meta_Account field.
func (r *mutationResolver) ModifyElementMetaAccount(ctx context.Context, id *string, input *model.NewCustomFieldElement) (*model.MetaAccountCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteElementMetaAccount is the resolver for the DeleteElement_Meta_Account field.
func (r *mutationResolver) DeleteElementMetaAccount(ctx context.Context, id *string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

///
///
///
///
///
///
///--------------------------------------------------------------------------------------------------------
///--------------------------------------- User Collection ------------------------------------------------
///--------------------------------------------------------------------------------------------------------

// AddNewUser is the resolver for the AddNewUser field.
func (r *mutationResolver) AddNewUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	return userCollectionRepository.AddNewUser(ctx, input)
}

// DeleteUser is the resolver for the DeleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id *string) (string, error) {
	return userCollectionRepository.DeleteUser(ctx, *id)
}

// ModifyUser is the resolver for the ModifyUser field.
func (r *mutationResolver) ModifyUser(ctx context.Context, id *string, input *model.NewUser) (*model.User, error) {
	return userCollectionRepository.ModifyUser(ctx, *id, input)
}

// Login is the resolver for the Login field.
func (r *mutationResolver) Login(ctx context.Context, input *model.Login) (*model.UserDto, error) {
	return userCollectionRepository.Login(ctx, input)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ User Role collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------

// AddNewUserRole is the resolver for the AddNewUserRole field.
func (r *mutationResolver) AddNewUserRole(ctx context.Context, input *model.NewUserRole) (*model.UserRole, error) {
	return UserRoleRepository.AddNewUserRole(ctx, *input)
}

// ModifyUserRole is the resolver for the ModifyUserRole field.
func (r *mutationResolver) ModifyUserRole(ctx context.Context, id *string, input *model.NewUserRole) (*model.UserRole, error) {
	return UserRoleRepository.ModifyUserRole(ctx, *id, input)
}

// DeleUserRole is the resolver for the DeleUserRole field.
func (r *mutationResolver) DeleUserRole(ctx context.Context, id *string) (string, error) {
	return UserRoleRepository.DeleUserRole(ctx, *id)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ Activity collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------
// AddNewActivity is the resolver for the AddNewActivity field.
func (r *mutationResolver) AddNewActivity(ctx context.Context, input *model.NewActivity) (*model.Activity, error) {
	return activityRepository.AddNewActivity(ctx, input)
}

// ModifyActivity is the resolver for the ModifyActivity field.
func (r *mutationResolver) ModifyActivity(ctx context.Context, id *string, input *model.NewActivity) (*model.Activity, error) {
	return activityRepository.ModifyActivity(ctx, *id, input)
}

// DeleteActivity is the resolver for the DeleteActivity field.
func (r *mutationResolver) DeleteActivity(ctx context.Context, id *string) (string, error) {
	return activityRepository.DeleteActivity(ctx, *id)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ Account collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------

// AddNewAccount is the resolver for the AddNewAccount field.
func (r *mutationResolver) AddNewAccount(ctx context.Context, input *model.NewActivity) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

// ModifyAccount is the resolver for the ModifyAccount field.
func (r *mutationResolver) ModifyAccount(ctx context.Context, id *string, input *model.NewActivity) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteAccount is the resolver for the DeleteAccount field.
func (r *mutationResolver) DeleteAccount(ctx context.Context, id *string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

///
///
///
///
///
///
///
///
///
///
///
///
///
///
///
///
///
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///------------------------------------------ Queries -----------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------

///--------------------------------------------------------------------------------------------------------
///------------------------------------ User Meta collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------

// GetUserMetaCollection is the resolver for the GetUserMetaCollection field.
func (r *queryResolver) GetUserMetaCollection(ctx context.Context) (*model.MetaUserCollection, error) {
	return userRepository.GetUser_MetaCollection(ctx)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ meta Activity collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------
// GetMetaActivityCollection is the resolver for the GetMetaActivityCollection field.
func (r *queryResolver) GetMetaActivityCollection(ctx context.Context) (*model.MetaActivityCollection, error) {
	return metaActivityRepository.GetMetaActivityCollection(ctx)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ Meta Account collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------
// GetMetaAccountCollection is the resolver for the GetMetaAccountCollection field.
func (r *queryResolver) GetMetaAccountCollection(ctx context.Context) (*model.MetaAccountCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ User Role collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------
// GetAllUserRole is the resolver for the GetAllUserRole field.
func (r *queryResolver) GetAllUserRole(ctx context.Context) ([]*model.UserRole, error) {
	return UserRoleRepository.GetAllUserRole(ctx)
}

// GetUserRoleByID is the resolver for the GetUserRoleByID field.
func (r *queryResolver) GetUserRoleByID(ctx context.Context, id *string) (*model.UserRole, error) {
	return UserRoleRepository.GetUserRoleByID(ctx, *id)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ Usercollection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------
// GetAllUser is the resolver for the GetAllUser field.
func (r *queryResolver) GetAllUser(ctx context.Context) ([]*model.User, error) {
	return userCollectionRepository.GetAllUser(ctx)
}

// GetUserByID is the resolver for the GetUserByID field.
func (r *queryResolver) GetUserByID(ctx context.Context, id *string) (*model.User, error) {
	return userCollectionRepository.GetUserByID(ctx, *id)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ Activity Collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------
// GetAllActivity is the resolver for the GetAllActivity field.
func (r *queryResolver) GetAllActivity(ctx context.Context) ([]*model.Activity, error) {
	return activityRepository.GetAllActivity(ctx)
}

// GetActivityByID is the resolver for the GetActivityByID field.
func (r *queryResolver) GetActivityByID(ctx context.Context, id *string) (*model.Activity, error) {
	return activityRepository.GetActivityByID(ctx, *id)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ Account Collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------
// GetAllAccount is the resolver for the GetAllAccount field.
func (r *queryResolver) GetAllAccount(ctx context.Context) ([]*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetAccountByID is the resolver for the GetAccountByID field.
func (r *queryResolver) GetAccountByID(ctx context.Context, id *string) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
