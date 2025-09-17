package postgres

import (
	"github.com/HericVirgilio/api-go/internal/domain"
	"gorm.io/gorm"
)

type userPostgresRepository struct{
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) domain.UserRepository{
	return &userPostgresRepository{db}
}

func (r *userPostgresRepository) Create(user *domain.User) error{
	return r.db.Create(user).Error
}

func (r *userPostgresRepository) FindAll() ([]domain.User, error){
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userPostgresRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userPostgresRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userPostgresRepository) Delete(id uint) error {
	// Usamos o Soft Delete do GORM
	return r.db.Delete(&domain.User{}, id).Error
}