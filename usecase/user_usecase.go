package usecase

import (
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
	userRepository repository.User
}

func (a *userUseCase) FindAllUser(id string) ([]model.User, error) {
	user, err := a.userRepository.GetAll(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *userUseCase) FindUserById(id string) (model.User, error) {
	user, err := a.userRepository.GetById(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *userUseCase) FindUserByEmail(email string) (model.User, error) {
	user, err := a.userRepository.GetByEmail(email)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *userUseCase) CreateUser(payload model.User) (model.User, error) {
	user, err := a.userRepository.CreateUser(payload)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *userUseCase) UpdateUser(payload model.User) (model.User, error) {
	user, err := a.userRepository.Update(payload)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *userUseCase) DeleteUser(id string) error {
	err := a.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}
