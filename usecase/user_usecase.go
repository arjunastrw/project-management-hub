package usecase

import (
	"fmt"
	"log"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/repository"
	"enigma.com/projectmanagementhub/shared/shared_model"
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
	users, paging, err := a.userRepository.GetAll(page, size)
	if err != nil {

		log.Println(err)
		return []model.User{}, shared_model.Paging{}, err
	}

	log.Println(users)
	return users, paging, nil

}

func (a *userUseCase) FindUserById(id string) (model.User, error) {
	user, err := a.userRepository.GetById(id)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get user. user id not found")
	}
	return user, nil
}

func (a *userUseCase) FindUserByEmail(email string) (model.User, error) {

	user, err := a.userRepository.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get user. user email not found")
	}
	return user, nil
}

func (a *userUseCase) CreateUser(payload model.User) (model.User, error) {

	if payload.Name == "" || payload.Email == "" || payload.Password == "" {

		return model.User{}, fmt.Errorf("failed to create user. empty field exist")
	}
	if payload.Role != "ADMIN" && payload.Role != "MANAGER" && payload.Role != "TEAM MEMBER" {

		return model.User{}, fmt.Errorf("failed to create user. invalid role. role: ('ADMIN', 'MANAGER', 'TEAM MEMBER')")
	}

	_, err := a.userRepository.GetById(payload.Id)
	if err == nil {
		return model.User{}, fmt.Errorf("failed to create user. user already exist")
	}

	_, err = a.userRepository.GetByEmail(payload.Email)
	if err == nil {
		return model.User{}, fmt.Errorf(" Email %s is already exist", payload.Email)
	}

	// Create new user
	user, err := a.userRepository.CreateUser(payload)
	if err != nil {

		log.Println(err)
		return model.User{}, err
	}

	// Create User Successfully
	log.Printf("Create User Successfully: %+v", user)
	return user, nil
}

func (a *userUseCase) UpdateUser(payload model.User) (model.User, error) {

	if payload.Name == "" || payload.Email == "" || payload.Password == "" {

		return model.User{}, fmt.Errorf("failed to update user. empty field exist")
	}
	if payload.Role != "ADMIN" && payload.Role != "MANAGER" && payload.Role != "TEAM MEMBER" {

		return model.User{}, fmt.Errorf("failed to update user. invalid role. role: ('ADMIN', 'MANAGER', 'TEAM MEMBER')")
	}

	existingUser, err := a.userRepository.GetByEmail(payload.Email)
	if err == nil && payload.Id != existingUser.Id {
		return model.User{}, fmt.Errorf("failed to update user. Email %s is already exist", payload.Email)
	}

	// Update User
	user, err := a.userRepository.Update(payload)
	if err != nil {

		log.Println(err)
		return model.User{}, err
	}

	// Update User Successfully
	log.Printf("Update User Successfully: %+v", user)
	return user, nil
}

func (a *userUseCase) DeleteUser(id string) error {

	if _, err := a.userRepository.GetById(id); err != nil {
		return fmt.Errorf("failed to delete user. user id invalid")
	}

	err := a.userRepository.Delete(id)
	if err != nil {

		log.Println(err)
		return err
	}

	// Delete User Successfully
	log.Printf("Delete User Successfully: %+v", id)
	return nil
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}
