package usecases

import (
	"strconv"

	"github.com/banggibima/go-chi-restful-api/internal/entities"
)

type UserRepository interface {
	GetAll() ([]entities.User, error)
	GetByID(id uint) (entities.User, error)
	Create(user entities.User) (entities.User, error)
	Update(id uint, user entities.User) (entities.User, error)
	Delete(id uint) error
}

type UserUseCase struct {
	UserRepository UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	return &UserUseCase{UserRepository: repo}
}

func (u *UserUseCase) GetUsers() ([]entities.User, error) {
	return u.UserRepository.GetAll()
}

func (u *UserUseCase) GetUserByID(id string) (entities.User, error) {
	convertedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return entities.User{}, err
	}

	return u.UserRepository.GetByID(uint(convertedID))
}

func (u *UserUseCase) CreateUser(user entities.User) (entities.User, error) {
	return u.UserRepository.Create(user)
}

func (u *UserUseCase) UpdateUser(id string, user entities.User) (entities.User, error) {
	convertedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return entities.User{}, err
	}

	return u.UserRepository.Update(uint(convertedID), user)
}

func (u *UserUseCase) DeleteUser(id string) error {
	convertedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	return u.UserRepository.Delete(uint(convertedID))
}
