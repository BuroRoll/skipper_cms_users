package handlers

import (
	"Skipper_cms_users/pkg/models/forms/inputForms"
	"Skipper_cms_users/pkg/models/forms/outputForms"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Description Получение списка всех пользователей
// @Tags 		Users
// @Security BearerAuth
// @Accept 		json
// @Produce 	json
// @Success 	200 		{object} 	[]models.User
// @Failure     500         {object}  	outputForms.ErrorResponse
// @Router /users [get]
func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.services.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, outputForms.ErrorResponse{
			Error: "Ошибка получения данных",
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Description Получение списка всех ролей
// @Tags 		Users
// @Security BearerAuth
// @Accept 		json
// @Produce 	json
// @Success 	200 		{object} 	[]models.Role
// @Failure     500         {object}  	outputForms.ErrorResponse
// @Router /roles [get]
func (h *Handler) getRoles(c *gin.Context) {
	roles, err := h.services.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, outputForms.ErrorResponse{
			Error: "Ошибка получения данных",
		})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// @Description Добавление роли сотруднику
// @Tags 		Users
// @Security BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		request 	body 		inputForms.AddUserRoleInput	true 	"query params"
// @Success 	200 		{object} 	outputForms.SuccessResponse
// @Failure     500         {object}  	outputForms.ErrorResponse
// @Router /users/add-role [put]
func (h *Handler) addRoleToUser(c *gin.Context) {
	var params inputForms.AddUserRoleInput
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, outputForms.ErrorResponse{Error: "Неверная форма запроса"})
	}
	user, err := h.services.AddRoleToUser(params.UserId, params.Roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, outputForms.ErrorResponse{
			Error: "Ошибка добавления данных",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Description Регистрация нового пользователя
// @Tags 		Users
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		request 	body 		inputForms.SignUpUserForm 	true 	"query params"
// @Success 	200 		{object} 	models.User
// @Failure     400         {object}  	outputForms.ErrorResponse
// @Failure     206         {object}  	outputForms.ErrorAssignRole
// @Failure     500         {object}  	outputForms.ErrorResponse
// @Router /users/register [post]
func (h *Handler) registerUser(c *gin.Context) {
	var input inputForms.SignUpUserForm
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, outputForms.ErrorResponse{
			Error: "Неверная форма регистрации",
		})
		return
	}
	user, err := h.services.CreateUser(input)
	if err != nil && user.ID <= 0 {
		c.JSON(http.StatusInternalServerError, outputForms.ErrorResponse{
			Error: "Ошибка создание профиля",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusPartialContent, outputForms.ErrorAssignRole{
			Error: "Ошибка добавления ролей пользователю",
			User:  user,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Description Удаление роли пользователя
// @Tags 		Users
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		request 	body 		inputForms.DeleteUserRole 	true 	"query params"
// @Success 	200 		{object} 	models.User
// @Failure     400         {object}  	outputForms.ErrorResponse
// @Failure     500         {object}  	outputForms.ErrorResponse
// @Router /users/delete-role [delete]
func (h *Handler) deleteUserRole(c *gin.Context) {
	var input inputForms.DeleteUserRole
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, outputForms.ErrorResponse{
			Error: "Неверная форма для удаления роли",
		})
		return
	}
	user, err := h.services.DeleteUserRole(input.UserId, input.RoleName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, outputForms.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Description Получение данных о пользователе, запрос без параметров вернёт данные о текущем пользователе
// @Tags 		Users
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param        id   		query     	int  	false  "UserId"
// @Success 	200 		{object} 	models.User
// @Failure     400         {object}  	outputForms.ErrorResponse
// @Failure     500         {object}  	outputForms.ErrorResponse
// @Router /users/info [get]
func (h *Handler) getUserInfo(c *gin.Context) {
	userIdInput, _ := strconv.ParseUint(c.Request.URL.Query().Get("id"), 10, 64)
	userId := uint(userIdInput)
	if userId == 0 {
		userId = c.GetUint(userCtx)
	}
	user, err := h.services.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, outputForms.ErrorResponse{Error: "Пользователь не найден"})
		return
	}
	c.JSON(http.StatusOK, user)
}
