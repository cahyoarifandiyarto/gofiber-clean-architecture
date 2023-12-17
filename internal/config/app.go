package config

import (
	"github.com/cahyoarifandiyarto/gofiber-clean-architecture/internal/delivery/http"
	"github.com/cahyoarifandiyarto/gofiber-clean-architecture/internal/delivery/http/route"
	"github.com/cahyoarifandiyarto/gofiber-clean-architecture/internal/repository"
	"github.com/cahyoarifandiyarto/gofiber-clean-architecture/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	userRepository := repository.NewUserRepository(config.Log)

	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)

	userController := http.NewUserController(userUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}
	routeConfig.Setup()
}
