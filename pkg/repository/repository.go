package repository

import (
	"Skipper_cms_users/pkg/models"
	"gorm.io/gorm"
)

type Users interface {
	GetUsers() ([]models.User, error)
	GetRoles() ([]models.Role, error)
	GetUser(userId uint) (models.User, error)
	GetRoleByName(roleName string) (models.Role, error)
	AddRoleToUser(user models.User, role models.Role) error
	CreateUser(firstName string, secondName string, email string, password string) (models.User, error)
	DeleteUserRole(user models.User, role models.Role) error
	ChangePassword(user models.User, newPassword string) error
}

type Repository struct {
	Users
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users: NewUsersPostgres(db),
	}
}
