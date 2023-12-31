package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vasapolrittideah/accord/apps/api/features/auth/handler"
	"github.com/vasapolrittideah/accord/apps/api/features/auth/middleware"
	"github.com/vasapolrittideah/accord/apps/api/features/auth/repository"
	"github.com/vasapolrittideah/accord/apps/api/features/auth/usecase"
	"github.com/vasapolrittideah/accord/apps/api/internal/config"
	"github.com/vasapolrittideah/accord/apps/api/internal/healthcheck"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Server struct {
	conf    *config.Config
	version string
}

func NewServer(version string) *Server {
	conf, err := config.New()
	if err != nil {
		log.Panicln("Failed to load application configuration: ", err)
	}

	return &Server{
		conf:    conf,
		version: version,
	}
}

func (s *Server) Run() {
	app := fiber.New()
	app.Use(
		recover.New(),
		logger.New(logger.Config{
			TimeFormat: time.RFC1123Z,
			TimeZone:   "Asia/Bangkok",
		}),
		cors.New(cors.Config{
			AllowOrigins: "http://*, https://*",
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: strings.Join([]string{
				fiber.MethodGet,
				fiber.MethodPut,
				fiber.MethodPost,
				fiber.MethodDelete,
			}, ","),
		}),
	)

	healthcheck.RegisterHandler(app, s.version)

	router := app.Group("/api/v1")

	authService := usecase.NewAuthUseCase(repository.NewRepository(s.conf.DB), s.conf)
	authMiddleware := middleware.NewAuthMiddleware(authService)
	handler.RegisterHandlers(router, s.conf, authService, authMiddleware)

	go func() {
		if err := app.Listen(":" + s.conf.ServerPort); err != nil {
			log.Fatalf("Failed to listen and serve application: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	<-quit
}
