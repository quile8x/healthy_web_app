package service

import (
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
	"golang.org/x/crypto/bcrypt"
)

// UserService is a service for managing user user.
type UserService interface {
	AuthenticateByUsernameAndPassword(username string, password string) (bool, *model.User)
}

type userService struct {
	container container.Container
}

// NewUserService is constructor.
func NewUserService(container container.Container) UserService {
	return &userService{container: container}
}

// AuthenticateByUsernameAndPassword authenticates by using username and plain text password.
func (a *userService) AuthenticateByUsernameAndPassword(username string, password string) (bool, *model.User) {
	rep := a.container.GetRepository()
	logger := a.container.GetLogger()
	user := model.User{}
	result, err := user.FindByName(rep, username)
	if err != nil {
		logger.GetZapLogger().Errorf(err.Error())
		return false, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.password), []byte(password)); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
		return false, nil
	}

	return true, result
}
