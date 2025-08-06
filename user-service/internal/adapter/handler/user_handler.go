package handler

import (
	"net/http"
	"user-service/internal/adapter/handler/request"
	"user-service/internal/adapter/handler/response"
	"user-service/internal/core/domain/entity"
	"user-service/internal/core/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type UserHandler interface {
	SignIn(ctx echo.Context) error
}

type userHandler struct {
	userService service.UserServiceInterface
}

var err error

func (u *userHandler) SignIn(c echo.Context) error {
	var (
		req       = request.SignInRequest{}
		res       = response.DefaultResponse{}
		signInRes = response.SignInResponse{}
		ctx       = c.Request().Context()
	)

	if err := c.Bind(&req); err != nil {
		log.Errorf("[UserHandler-1] SignIn: %v", err)
		res.Message = err.Error()
		res.Data = nil
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	if err := c.Validate(req); err != nil {
		log.Errorf("[UserHandler-2] SignIn: %v", err)
		res.Message = err.Error()
		res.Data = nil
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	reqEntity := entity.UserEntity{
		Email:    req.Email,
		Password: req.Password,
	}

	user, token, err := u.userService.SignIn(ctx, reqEntity)
	if err != nil {
		if err.Error() == "404" {
			log.Errorf("[UserHandler-3] SignIn: %s", "User not found")
			res.Message = "User not found"
			res.Data = nil
			return c.JSON(http.StatusNotFound, res)
		}
		log.Errorf("[UserHandler-4] SignIn: %v", err)
		res.Message = err.Error()
		res.Data = nil
		return c.JSON(http.StatusInternalServerError, res)
	}

	signInRes = response.SignInResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Phone:       user.Phone,
		Role:        user.RoleName,
		Lat:         user.Lat,
		Lng:         user.Lng,
		AccessToken: token,
	}

	res.Message = "Sign in successfully"
	res.Data = signInRes
	return c.JSON(http.StatusOK, res)
}

func NewUserHandler(e *echo.Echo, userService service.UserServiceInterface) UserHandler {
	userHandler := &userHandler{userService: userService}

	e.Use(middleware.Recover())

	e.POST("/signin", userHandler.SignIn)
	return userHandler
}
