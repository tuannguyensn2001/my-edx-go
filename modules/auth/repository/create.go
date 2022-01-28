package authrepository

import (
	"my-edx-go/models"
	authdto "my-edx-go/modules/auth/dto"
	hash_provider "my-edx-go/providers/hash"
)

func (repository *authRepository) Create(userDTO authdto.Register) (*models.User, error) {

	hashPassword, err := hash_provider.Hash(userDTO.Password)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:    userDTO.Email,
		Password: hashPassword,
		Username: userDTO.Username,
	}

	result := *repository.db.Create(&user)

	return &user, result.Error
}
