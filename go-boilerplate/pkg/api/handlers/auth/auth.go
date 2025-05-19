package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	// TODO: ambil data user dari request dan simpan ke DB
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Register endpoint hit",
	})
}

func Login(c echo.Context) error {
	// TODO: verifikasi email/password, lalu generate JWT
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login endpoint hit",
	})
}
