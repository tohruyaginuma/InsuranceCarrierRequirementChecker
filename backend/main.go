package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/registry"
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/route"
)

const (
	port      = "8080"
	clientURL = "http://localhost:5173"
)

func setLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
}

func setEcho() *echo.Echo {
	e := echo.New()

	e.Debug = true

	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{clientURL},
			AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Authorization"},
			AllowCredentials: true,
		},
	))

	return e
}

func main() {
	setLogger()
	e := setEcho()

	registry := registry.NewRegistry()
	route.SetRoute(e, registry)

	slog.Info("app starting")

	e.Logger.Fatal(e.Start(":" + port))
}
