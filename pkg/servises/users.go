package service

import (
	"Skipper_cms_users/pkg/models"
	"Skipper_cms_users/pkg/models/forms/inputForms"
	"Skipper_cms_users/pkg/repository"
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
)

type UsersService struct {
	repo repository.Users
}

const salt = "14hjqrhj1231qw124617ajfha1123ssfqa3ssjs190"

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (u UsersService) GetUsers() ([]models.User, error) {
	return u.repo.GetUsers()
}

func (u UsersService) GetRoles() ([]models.Role, error) {
	return u.repo.GetRoles()
}

func (u UsersService) GetUserRoles(userId uint) ([]models.Role, error) {
	user, err := u.repo.GetUser(userId)
	return user.Role, err
}
func (u UsersService) AddRoleToUser(userId uint, roleId uint) error {
	user, err := u.repo.GetUser(userId)
	role, err := u.repo.GetRole(roleId)
	err = u.repo.AddRoleToUser(user, role)
	if err != nil {
		return err
	}
	return nil
}

func (u UsersService) CreateUser(userData inputForms.SignUpUserForm) (models.User, error) {
	user, err := u.repo.CreateUser(userData.FirstName, userData.SecondName, userData.Phone, generatePasswordHash(userData.Password))
	if err != nil {
		return models.User{}, err
	}
	for _, id := range userData.RolesIds {
		role, err := u.repo.GetRole(id)
		if err != nil {
			return user, err
		}
		err = u.repo.AddRoleToUser(user, role)
		if err != nil {
			return user, err
		}
	}
	user, err = u.repo.GetUser(user.ID)
	return user, nil
}

func (u UsersService) DeleteUserRole(userId uint, roleId uint) (models.User, error) {
	user, err := u.repo.GetUser(userId)
	role, err := u.repo.GetRole(roleId)
	if err != nil {
		return user, errors.New("Ошибка получения данных о пользователе или роли")
	}
	for _, value := range user.Role {
		if value.Name == "super_admin" {
			return user, errors.New("Невозможно удалить роли суперадмина")
		}
	}
	err = u.repo.DeleteUserRole(user, role)
	user, err = u.repo.GetUser(userId)
	return user, err
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
