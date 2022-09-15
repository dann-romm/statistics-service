package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"statistics-service/internal/entity"
	"statistics-service/internal/service"
)

type authRoutes struct {
	s service.Auth
}

func newAuthRoutes(g *echo.Group, s service.Auth) {
	r := &authRoutes{s: s}

	g.POST("/sign-up", r.signUp)
	g.POST("/sign-in", r.signIn)
}

// регистрация пользователя
func (r *authRoutes) signUp(c echo.Context) error {
	var input entity.User

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	id, err := r.s.CreateUser(c.Request().Context(), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// аутентификация пользователя
func (r *authRoutes) signIn(c echo.Context) error {
	username, password, ok := c.Request().BasicAuth()
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return nil
	}

	token, err := r.s.GenerateToken(c.Request().Context(), username, password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
