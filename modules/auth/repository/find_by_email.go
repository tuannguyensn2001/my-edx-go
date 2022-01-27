package authrepository

import (
	"my-edx-go/models"
)

func (repository *authRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	result := repository.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
