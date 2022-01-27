package authrepository

import "my-edx-go/models"

func (repository *authRepository) FindById() models.User {
	return models.User{
		Email:    "devpro2001@gmail.com",
		Password: "java2001",
	}
}
