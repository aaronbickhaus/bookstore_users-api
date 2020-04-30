package services

import (
	"github.com/aaronbickhaus/bookstore_users-api/src/domain/users"
	"github.com/aaronbickhaus/bookstore_users-api/src/utils/crypto_utils"
	"github.com/aaronbickhaus/bookstore_users-api/src/utils/date_utiils"
	"github.com/aaronbickhaus/bookstore_users-api/src/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)
type usersService struct {}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	DeleteUser(int64)  *errors.RestErr
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	SearchUsers(string) (users.Users, *errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr){
 if err := user.Validate(); err != nil {
 	return nil, err
 }

 user.DateCreated = date_utiils.GetNowDBFormat()
 user.Status = users.StatusActive
 user.Password = crypto_utils.GetMd5(user.Password)
 if err := user.Save(); err != nil {
 	return nil, err
 }

 return &user, nil
}

func  (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func  (s *usersService) DeleteUser(userId int64)  *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}

	if err := current.Get(); err != nil{
		return nil, err
	}

	if !isPartial {
		current.Email = user.Email
		current.LastName = user.LastName
		current.FirstName = user.FirstName
		current.Status = user.Status
		current.Password = user.Password
	}else{
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Password != "" {
			current.Password = user.Password
		}
		if user.Status != "" {
			current.Status = user.Status
		}
	}

	if err := current.Update(); err != nil {
		return  nil, err
	}

	return current, nil
}

func  (s *usersService) SearchUsers(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.Search(status)
}
