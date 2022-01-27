package authrepository

import (
	"my-edx-go/models"
	authdto "my-edx-go/modules/auth/dto"
)

func (repository *authRepository) Create(userDTO authdto.Register) (*models.User, error) {
	user := models.User{
		Email:    userDTO.Email,
		Password: userDTO.Password,
		Username: userDTO.Username,
	}

	result := *repository.db.Create(&user)

	return &user, result.Error
}
