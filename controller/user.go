package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// UserController is a controller for managing user User.
type UserController interface {
	GetLoginStatus(c echo.Context) error
	GetLoginUser(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type UserController struct {
	context   container.Container
	service   service.UserService
	dummyUser *model.User
}

// NewUserController is constructor.
func NewUserController(container container.Container) UserController {
	return &UserController{
		context:   container,
		service:   service.NewUserService(container),
		dummyUser: model.NewUserWithPlainPassword("test", "test", 1),
	}
}

// GetLoginStatus returns the status of login.
// @Summary Get the login status.
// @Description Get the login status of current logged-in user.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {boolean} bool "The current user have already logged-in. Returns true."
// @Failure 401 {boolean} bool "The current user haven't logged-in yet. Returns false."
// @Router /auth/loginStatus [get]
func (controller *UserController) GetLoginStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}

// GetLoginUser returns the User data of logged in user.
// @Summary Get the User data of logged-in user.
// @Description Get the User data of logged-in user.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User "Success to fetch the User data. If the security function is disable, it returns the dummy data."
// @Failure 401 {boolean} bool "The current user haven't logged-in yet. Returns false."
// @Router /auth/loginUser [get]
func (controller *UserController) GetLoginUser(c echo.Context) error {
	if !controller.context.GetConfig().Extension.SecurityEnabled {
		return c.JSON(http.StatusOK, controller.dummyUser)
	}
	return c.JSON(http.StatusOK, controller.context.GetSession().GetUser())
}

// Login is the method to login using username and password by http post.
// @Summary Login using username and password.
// @Description Login using username and password.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body dto.LoginDto true "User name and Password for logged-in."
// @Success 200 {object} model.User "Success to the authentication."
// @Failure 401 {boolean} bool "Failed to the authentication."
// @Router /auth/login [post]
func (controller *UserController) Login(c echo.Context) error {
	dto := dto.NewLoginDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}

	sess := controller.context.GetSession()
	if User := sess.GetUser(); User != nil {
		return c.JSON(http.StatusOK, User)
	}

	authenticate, a := controller.service.AuthenticateByUsernameAndPassword(dto.UserName, dto.Password)
	if authenticate {
		_ = sess.SetUser(a)
		_ = sess.Save()
		return c.JSON(http.StatusOK, a)
	}
	return c.NoContent(http.StatusUnauthorized)
}

// Logout is the method to logout by http post.
// @Summary Logout.
// @Description Logout.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /auth/logout [post]
func (controller *UserController) Logout(c echo.Context) error {
	sess := controller.context.GetSession()
	_ = sess.SetUser(nil)
	_ = sess.Delete()
	return c.NoContent(http.StatusOK)
}
