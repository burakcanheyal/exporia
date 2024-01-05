package cmd

import (
	"exporia/internal/application/handler"
	"exporia/internal/domain/dto"
	"exporia/internal/domain/service"
	"exporia/internal/middleware"
	"exporia/internal/server"
	"exporia/platform/app_log"
	"exporia/platform/postgres"
	"exporia/platform/postgres/repository"
	"exporia/platform/zap"
	"log"
)

func init() {
	zap.Logger.Info("Env dosyaları okunuyor")

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("dev.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}
func Setup() {
	zap.Logger.Info("Setup Başlatıldı")
	config := dto.Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Println(err)
	}
	db, err := postgres.InitializeDatabase(config.DBURL)
	if err != nil {
		log.Println(err)
	}

	err = app_log.InitializeAppLogDatabase(db)
	if err != nil {
		log.Println(err)
	}

	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	applicationLogRepository := app_log.NewApplicationLogRepository(db)

	appLogService := app_log.NewApplicationLogService(applicationLogRepository)
	userService := service.NewUserService(userRepository, roleRepository, appLogService)
	authenticationService := service.NewAuthentication(userRepository, config.Secret, config.Secret2, appLogService)

	authenticationMiddleware := middleware.NewMiddleware(authenticationService, userService)

	authenticationServerHandler := handler.NewAuthenticationServerHandler(authenticationService)
	profileServerHandler := handler.NewProfileServerHandler(userService)

	webServer := server.NewWebServer(
		profileServerHandler,
		authenticationServerHandler,
		authenticationMiddleware,
	)

	webServer.SetupRoot()
}