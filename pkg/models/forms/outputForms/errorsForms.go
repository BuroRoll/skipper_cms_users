package outputForms

import "Skipper_cms_users/pkg/models"

type ErrorResponse struct {
	Error string `json:"error"`
}

type ErrorAssignRole struct {
	Error string      `json:"error"`
	User  models.User `json:"user"`
}
