package app

import (
	"context"
	"log"
	"news-portal/config"
	"news-portal/lib/auth"
	"news-portal/lib/middleware"
	"news-portal/lib/pagination"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"os"
)

func RunServer() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("[RunServer-1] Failed to load config")
	}


	_, err = cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal("[RunServer-2] Failed to connect to database")
	}

	log.Println("[RunServer-3] Server successfully started with database connection")

	cdfR2, err := cfg.LoadCloudflareR2Config()
	if err != nil {
		log.Fatal("[RunServer-4] Failed to load Cloudflare R2 config")
	}
	_ = s3.NewFromConfig(cdfR2)

	_ = auth.NewJwt(cfg)
	_ = middleware.NewMiddleware(cfg)

	_ = pagination.NewPagination()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] %{ip} %{status} - %{latency} %{method} %{path}\n",
	}))
	
	_ = app.Group("/api")

	go func() {
		if cfg.AppPort == "" {
			cfg.AppPort = os.Getenv("APP_PORT")
		}

		ERR := app.Listen(":" + cfg.AppPort)
		if ERR != nil {
			log.Fatal("[RunServer-5] Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	
	<-quit

	log.Println("[RunServer-6] Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.ShutdownWithContext(ctx)

}
