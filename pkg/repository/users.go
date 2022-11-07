package repository

import (
	"Skipper_cms_users/pkg/models"
	"gorm.io/gorm"
)

type UsersPostgres struct {
	db *gorm.DB
}

func NewUsersPostgres(db *gorm.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (u UsersPostgres) GetUsers() ([]models.User, error) {
	var users []models.User
	err := u.db.Preload("Role").Find(&users)
	return users, err.Error
}

func (u UsersPostgres) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	err := u.db.Find(&roles)
	return roles, err.Error
}

func (u UsersPostgres) GetUser(userId uint) (models.User, error) {
	var user models.User
	err := u.db.Preload("Role").First(&user, userId)
	return user, err.Error
}

func (u UsersPostgres) AddRoleToUser(user models.User, role models.Role) error {
	err := u.db.Model(&user).Association("Role").Append(&role)
	return err
}

func (u UsersPostgres) CreateUser(firstName string, secondName string, email string, password string) (models.User, error) {
	var user models.User
	user = models.User{
		Email:      email,
		FirstName:  firstName,
		SecondName: secondName,
		Password:   password,
	}
	result := u.db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u UsersPostgres) DeleteUserRole(user models.User, role models.Role) error {
	return u.db.Model(&user).Association("Role").Delete(&role)
}

func (u UsersPostgres) GetRoleByName(roleName string) (models.Role, error) {
	var role models.Role
	err := u.db.First(&role, "name = ?", roleName)
	return role, err.Error
}

func (u UsersPostgres) ChangePassword(user models.User, newPassword string) error {
	err := u.db.Model(&user).Update("password", newPassword)
	return err.Error
}
