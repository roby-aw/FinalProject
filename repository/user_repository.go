package repository

import (
	"errors"
	"final-project/dto"
	"final-project/entity"
	"final-project/interfaces"

	"gorm.io/gorm"
)

type userrepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userrepository{
		DB: db,
	}
}

func (r *userrepository) Register(register *dto.Register) (entity.User, error) {
	user := entity.User{
		Email:    register.Email,
		Password: register.Password,
		Username: register.Username,
		Age:      register.Age,
	}

	result := r.DB.Model(&user).Create(&user)

	if result.RowsAffected < 1 {
		return user, result.Error
	}
	return user, nil
}

func (r *userrepository) Login(email, password string) (dto.Login, error) {
	user := entity.User{
		Email:    email,
		Password: password,
	}

	login := dto.Login{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}

	result := r.DB.Model(&user).Where("email = ?", email).First(&login)

	if result.RowsAffected < 1 {
		return login, result.Error
	}

	return login, nil
}

func (r *userrepository) UpdateUser(id uint, user *dto.UpdateUser) (dto.UpdateUser, error) {
	userUpdate := entity.User{
		Email:    user.Email,
		Username: user.Username,
	}

	result := r.DB.Model(&userUpdate).Where("id = ?", id).Updates(&userUpdate)

	if result.RowsAffected < 1 {
		return *user, errors.New("user not found")
	}

	return *user, nil
}

func (r *userrepository) DeleteUser(id uint) error {
	user := entity.User{
		ID: id,
	}

	result := r.DB.Model(&user).Where("id = ?", id).Delete(&user)

	if result.RowsAffected < 1 {
		return errors.New("user not found")
	}

	return nil
}
