package usecase

import (
	"fmt"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/repository"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/sirupsen/logrus"
)

type UserUseCase interface {
	FindAllUser(page int, size int) ([]model.User, shared_model.Paging, error)
	FindUserById(id string) (model.User, error)
	FindUserByEmail(email string) (model.User, error)
	CreateUser(payload model.User) (model.User, error)
	UpdateUser(payload model.User) (model.User, error)
	DeleteUser(id string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
	logger         *logrus.Logger
}

func (a *userUseCase) FindAllUser(page int, size int) ([]model.User, shared_model.Paging, error) {

	// Validate Role
	//if user.Role != "ADMIN" {
	//	return []model.User{}, shared_model.Paging{}, fmt.Errorf("Can't Access this Resource!")
	//}

	users, paging, err := a.userRepository.GetAll(page, size)
	if err != nil {
		return []model.User{}, shared_model.Paging{}, err
	}

	a.logger.Infof("Get All User Successfully")
	return users, paging, nil
}

func (a *userUseCase) FindUserById(id string) (model.User, error) {
	user, err := a.userRepository.GetById(id)
	if err != nil {
		return model.User{}, err
	}

	// Check if user is not found
	if user.Id == "" {
		a.logger.Warnf("User with ID %s not found", id)
		return model.User{}, fmt.Errorf(" User with ID %s not found", id)
	}

	// User successfully found By ID
	a.logger.Infof("User with ID %s found successfully", id)

	return user, nil
}

func (a *userUseCase) FindUserByEmail(email string) (model.User, error) {
	// Validate Role
	//if user.Role != "ADMIN" {
	//	return model.User{}, fmt.Errorf("Can't Access this Resource!")
	//}

	user, err := a.userRepository.GetByEmail(email)
	if err != nil {
		// Handle error from repository
		return model.User{}, err
	}
	// Check if user not found
	if user.Email == "" {
		a.logger.Warnf("User with email %s not found", email)
		return model.User{}, fmt.Errorf(" User with email %s not found", email)
	}

	// User successfully found By Email
	a.logger.Infof("User with Email %s found successfully", email)
	return user, nil
}

func (a *userUseCase) CreateUser(payload model.User) (model.User, error) {

	//// Validate Role
	//if payload.Role != "ADMIN" {
	//	return model.User{}, fmt.Errorf("Can't Access this Resource!")
	//}

	// Validate email existence
	existingUser, err := a.userRepository.GetByEmail(payload.Email)
	if err != nil {
		return model.User{}, err
	}

	if existingUser.Email != "" {
		a.logger.Warnf("Email %s is already exist", payload.Email)
		return model.User{}, fmt.Errorf(" Email %s is already exist", payload.Email)
	}

	// Create new user
	user, err := a.userRepository.CreateUser(payload)
	if err != nil {
		return model.User{}, err
	}

	// Create User Successfully
	a.logger.Infof("Create User Successfully: %+v", user)
	return user, nil
}

func (a *userUseCase) UpdateUser(payload model.User) (model.User, error) {
	//// Validate Role
	//if currentUser.Role != "ADMIN" {
	//	return model.User{}, fmt.Errorf("Unauthorized access: Only admin can perform this operation")
	//}

	// If the email is being updated, check for existence
	if payload.Email != "" {
		existingUser, err := a.userRepository.GetByEmail(payload.Email)
		if err != nil {
			return model.User{}, err
		}

		if existingUser.Email != "" {
			a.logger.Warnf("Email %s is already exist", payload.Email)
			return model.User{}, fmt.Errorf(" Email %s is already exist", payload.Email)
		}
	}

	// Update User
	user, err := a.userRepository.Update(payload)
	if err != nil {
		return model.User{}, err
	}

	// Update User Successfully
	a.logger.Infof("Update User Successfully")
	return user, nil
}

func (a *userUseCase) DeleteUser(id string) error {

	//// Validate Role
	//if user.Role != "ADMIN" {
	//	return fmt.Errorf("Can't Access this Resource!")
	//}

	err := a.userRepository.Delete(id)
	if err != nil {
		return err
	}

	// Delete User Successfully
	a.logger.Infof("Delete User %s Successfully", id)
	return nil
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
		logger:         logrus.New(),
	}
}
