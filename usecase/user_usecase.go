package usecase

import (
	"fmt"
	"log"

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
}

func (a *userUseCase) FindAllUser(page int, size int) ([]model.User, shared_model.Paging, error) {
	// IF USER ROLE ISN'T ADMIN CANT GET ALL USER

	users, paging, err := a.userRepository.GetAll(page, size)
	if err != nil {

		return []model.User{}, shared_model.Paging{}, err
	}

	log.Println(users)
	return users, paging, nil

}

func (a *userUseCase) FindUserById(id string) (model.User, error) {
	user, err := a.userRepository.GetById(id)
	if err != nil {
		return model.User{}, err
	}

	// IF ROLE ISN'T ADMIN CANT GET USER BY ID

	// Check if user is not found
	if user.Id == "" {

		return model.User{}, fmt.Errorf(" User with ID %s not found", id)
	}

	// User successfully found By ID
	return user, nil
}

func (a *userUseCase) FindUserByEmail(email string) (model.User, error) {

	// IF ROLE ISN'T ADMIN CANT GET USER BY EMAIL

	user, err := a.userRepository.GetByEmail(email)
	if err != nil {
		// Handle error from repository
		return model.User{}, err
	}
	// Check if user not found
	if user.Email == "" {

		return model.User{}, fmt.Errorf(" User with email %s not found", email)
	}

	// User successfully found By Email

	return user, nil
}

func (a *userUseCase) CreateUser(payload model.User) (model.User, error) {

	// IF ROLE ISN'T ADMIN CANT CREATE USER

	// Validate email existence
	existingUser, err := a.userRepository.GetByEmail(payload.Email)
	if err != nil {
		return model.User{}, err
	}

	if existingUser.Email == "" {

		return model.User{}, fmt.Errorf(" Email %s is already exist", payload.Email)
	}

	// Create new user
	user, err := a.userRepository.CreateUser(payload)
	if err != nil {
		return model.User{}, err
	}

	// Create User Successfully
	log.Printf("Create User Successfully: %+v", user)
	return user, nil
}

func (a *userUseCase) UpdateUser(payload model.User) (model.User, error) {

	// IF ROLE ISN'T ADMIN CANT UPDATE USER

	// If the email is being updated, check for existence
	if payload.Email != "" {
		existingUser, err := a.userRepository.GetByEmail(payload.Email)
		if err != nil {
			return model.User{}, err
		}

		if existingUser.Email != "" {

			return model.User{}, fmt.Errorf(" Email %s is already exist", payload.Email)
		}
	}

	// Update User
	user, err := a.userRepository.Update(payload)
	if err != nil {
		return model.User{}, err
	}

	// Update User Successfully

	return user, nil
}

func (a *userUseCase) DeleteUser(id string) error {

	// IF ROLE ISN'T ADMIN CANT DELETE USER

	err := a.userRepository.Delete(id)
	if err != nil {
		return err
	}

	// Delete User Successfully

	return nil
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}
