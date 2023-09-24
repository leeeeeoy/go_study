package handler

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	e.POST("/user", uh.createUser)
	e.POST("/user/sign-in", uh.signIn)
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

func (uh *UserHandler) signIn(c echo.Context) error {

	param := new(dto.SignInRequest)

	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "-1",
			"message": "잘못된 파라미터 요청입니다.",
		})
	}

	user, err := uh.userRepository.SignIn(param.Email, param.Password)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "-1",
			"message": "아이디 또는 패스워드를 확인해주세요.",
		})
	}

	claims := &dto.JwtCustomClaims{
		Name:  user,
		Admin: param.Email == "leeeeeoy@dozn.co.kr",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "-1",
			"message": "토큰 생성에 실패했습니다.",
		})
	}

	return c.JSON(http.StatusOK, &dto.SignInResponse{
		AccessToken: accessToken,
	})
}
