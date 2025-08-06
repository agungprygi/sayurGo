package adapter

import (
	"net/http"
	"strings"
	"user-service/config"
	"user-service/internal/adapter/handler/response"

	"github.com/labstack/echo/v4"
)

type MiddlewareAdapterInterface interface {
	CheckToken() echo.MiddlewareFunc
}

type middlewareAdapter struct {
	cfg *config.Config
}

func (m *middlewareAdapter) CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			resErr := response.DefaultResponse{}
			redisCoon := config.NewRedisClient(m.cfg)
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				resErr.Message = "Missing or invalid token"
				resErr.Data = nil
				return c.JSON(http.StatusUnauthorized, resErr)
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			getSession, err := redisCoon.HGetAll(c.Request().Context(), tokenString).Result()
			if err != nil || getSession["logged_in"] == "false" {
				resErr.Message = "Missing or Invalid token"
				resErr.Data = nil
				return c.JSON(http.StatusUnauthorized, resErr)
			}
			c.Set("session", getSession)
			return next(c)
		}
	}
}

func NewMiddlewareAdapter(cfg *config.Config) MiddlewareAdapterInterface {
	return &middlewareAdapter{cfg: cfg}
}
