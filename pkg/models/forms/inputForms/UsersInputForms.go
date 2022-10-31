package inputForms

type AddUserRoleForm struct {
	UserId uint `json:"user_id" binding:"required"`
	RoleId uint `json:"role_id" binding:"required"`
}

type SignUpUserForm struct {
	Phone      string `json:"phone" binding:"required"`
	Password   string `json:"password" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	SecondName string `json:"second_name" binding:"required"`
	RolesIds   []uint `json:"roles_ids"`
}

type DeleteUserRole struct {
	UserId uint `json:"user_id" binding:"required"`
	RoleId uint `json:"role_id" binding:"required"`
}
