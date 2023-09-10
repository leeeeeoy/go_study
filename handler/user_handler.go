package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leeeeeoy/go_study/dto"
	"github.com/leeeeeoy/go_study/repository"
)

type UserHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

func (uh *UserHandler) InitUserHandler(e *echo.Echo) {
	e.GET("/user/:email", uh.getUserByEmail)
	e.POST("/user", uh.createUser)
}

func (uh *UserHandler) createUser(c echo.Context) error {
	param := new(dto.UserRequest)

	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "-1",
			"message": "잘못된 파라미터 요청입니다.",
		})
	}

	res, err := uh.userRepository.CreateUser(param)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "-1",
			"message": "중복된 이메일입니다.",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (uh *UserHandler) getUserByEmail(c echo.Context) error {
	param := c.Param("email")

	if param == "" {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "-1",
			"message": "잘못된 파라미터 요청입니다.",
		})
	}

	res, err := uh.userRepository.GetUserByEmail(param)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "-1",
			"message": "해당 유저가 존재하지 않습니다.",
		})
	}

	return c.JSON(http.StatusOK, res)
}
