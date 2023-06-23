package repository

import (
	"github.com/google/uuid"
	"github.com/vasapolrittideah/accord/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllUsers() (users []models.User, err error)
	GetUser(id uuid.UUID) (user models.User, err error)
	GetByEmail(email string) (user models.User, err error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(id uuid.UUID, newData models.User) (user models.User, err error)
	DeleteUser(id uuid.UUID) (user models.User, err error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetAllUsers() (users []models.User, err error) {
	return users, r.db.Find(&users).Error
}

func (r repository) GetUser(id uuid.UUID) (user models.User, err error) {
	return user, r.db.First(&user, id).Error
}

func (r repository) GetByEmail(email string) (user models.User, err error) {
	return user, r.db.First(&user, "email=?", email).Error
}

func (r repository) CreateUser(user models.User) (models.User, error) {
	return user, r.db.Create(&user).Error
}

func (r repository) UpdateUser(id uuid.UUID, newData models.User) (user models.User, err error) {
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, r.db.Model(&user).Updates(&newData).Error
}

func (r repository) DeleteUser(id uuid.UUID) (user models.User, err error) {
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, r.db.Delete(&user).Error
}
