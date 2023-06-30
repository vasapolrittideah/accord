package repository

import (
	"github.com/google/uuid"
	"github.com/vasapolrittideah/accord/apps/api/models"
	"gorm.io/gorm"
)

//go:generate mockery --name UserRepository --filename repository_mock.go
type UserRepository interface {
	GetAllUsers() (users []models.User, err error)
	GetUser(id uuid.UUID) (user models.User, err error)
	GetByEmail(email string) (user models.User, err error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(id uuid.UUID, newData models.User) (user models.User, err error)
	DeleteUser(id uuid.UUID) (user models.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return userRepository{db}
}

func (r userRepository) GetAllUsers() (users []models.User, err error) {
	return users, r.db.Find(&users).Error
}

func (r userRepository) GetUser(id uuid.UUID) (user models.User, err error) {
	return user, r.db.First(&user, id).Error
}

func (r userRepository) GetByEmail(email string) (user models.User, err error) {
	return user, r.db.First(&user, "email=?", email).Error
}

func (r userRepository) CreateUser(user models.User) (models.User, error) {
	return user, r.db.Create(&user).Error
}

func (r userRepository) UpdateUser(id uuid.UUID, newData models.User) (user models.User, err error) {
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, r.db.Model(&user).Updates(&newData).Error
}

func (r userRepository) DeleteUser(id uuid.UUID) (user models.User, err error) {
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, r.db.Delete(&user).Error
}
