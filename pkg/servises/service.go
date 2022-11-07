package service

import (
	"Skipper_cms_users/pkg/models"
	"Skipper_cms_users/pkg/models/forms/inputForms"
	"Skipper_cms_users/pkg/repository"
)

type Users interface {
	GetUsers() ([]models.User, error)
	GetUser(userId uint) (models.User, error)
	GetRoles() ([]models.Role, error)
	GetUserRoles(userId uint) ([]models.Role, error)
	AddRoleToUser(userId uint, roleId []string) (models.User, error)
	CreateUser(userData inputForms.SignUpUserInput) (models.User, error)
	DeleteUserRole(userId uint, roleName string) (models.User, error)
	ChangePassword(userId uint, oldPassword string, newPassword string) error
}

type Service struct {
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repos.Users),
	}
}
