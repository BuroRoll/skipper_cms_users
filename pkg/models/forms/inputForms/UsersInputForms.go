package inputForms

type AddUserRoleInput struct {
	UserId uint     `json:"user_id" binding:"required"`
	Roles  []string `json:"roles" binding:"required"`
}

type SignUpUserForm struct {
	Email      string   `json:"email" binding:"required"`
	Password   string   `json:"password" binding:"required"`
	FirstName  string   `json:"first_name" binding:"required"`
	SecondName string   `json:"second_name" binding:"required"`
	RolesNames []string `json:"roles_names"`
}

type DeleteUserRole struct {
	UserId   uint   `json:"user_id" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}

type UserInfoInput struct {
	UserId uint `json:"user_id"`
}
