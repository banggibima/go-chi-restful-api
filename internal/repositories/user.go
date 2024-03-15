package repositories

import (
	"github.com/banggibima/go-chi-restful-api/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetAll() ([]entities.User, error) {
	var users []entities.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(id uint) (entities.User, error) {
	var user entities.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Create(user entities.User) (entities.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Update(id uint, user entities.User) (entities.User, error) {
	existingUser, err := r.GetByID(id)
	if err != nil {
		return entities.User{}, err
	}

	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := r.DB.Save(&existingUser).Error; err != nil {
		return entities.User{}, err
	}
	return existingUser, nil
}

func (r *UserRepository) Delete(id uint) error {
	var user entities.User
	if err := r.DB.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}
