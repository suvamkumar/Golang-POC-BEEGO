package services

import (
	// "crud_with_gin_gonic/internal/users/domain/users"
	// "crud_with_gin_gonic/internal/users/utils/date_utils"
	// "crud_with_gin_gonic/internal/users/utils/errors"

	"crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/models"
	"crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/utils/date_utils"
	"crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/utils/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	//UserService available to others to call the method
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(models.User) (*models.User, *errors.RestErr)
	UpdateUser(models.User, string) (*models.User, *errors.RestErr)
	GetUser(id string) (*models.User, *errors.RestErr)
	DeleteUser(id string) *errors.RestErr
	GetAllUser() ([]models.User, *errors.RestErr)
}

//CreateUser ...
func (s *userService) CreateUser(user models.User) (*models.User, *errors.RestErr) {
	user.Status = models.StatusActive
	user.ID = primitive.NewObjectID()
	user.CreateDate = date_utils.GetNowDBFormat()
	if err := user.Insert(); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUser ...
func (s *userService) GetUser(id string) (*models.User, *errors.RestErr) {
	user := models.User{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}
	user.ID = objectID
	if err := user.GetUser(id); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetAllUser ...
func (s *userService) GetAllUser() ([]models.User, *errors.RestErr) {
	user := models.User{}
	return user.GetAllUser()
}

//UpdateUser ...
func (s *userService) UpdateUser(user models.User, id string) (*models.User, *errors.RestErr) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	user.ID = objectID
	if err := user.UpdateUser(); err != nil {
		return nil, err
	}
	return &user, nil
}

//DeleteUser ...
func (s *userService) DeleteUser(id string) *errors.RestErr {
	var user models.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	user.ID = objectID
	if err := user.DeleteUser(); err != nil {
		return err
	}
	return nil
}
