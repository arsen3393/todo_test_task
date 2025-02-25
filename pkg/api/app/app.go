package app

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"testSkillsRock/internal/config"
	"testSkillsRock/pkg/api/handler"
	"testSkillsRock/pkg/database"
	"testSkillsRock/pkg/models"
)

// @title  TODO API
// @version 1.0
// @description Тестовое задание SkillsRock API для simpletodo приложения
// @BasePath /

func RunApp(cfg *config.Config, logger *slog.Logger) {
	app := fiber.New()

	// MiddleWare Logger

	// Init DB
	db := database.MustConnectDB(cfg.DBConfig)
	connectDB := models.DBconnection{}
	connectDB.SetDB(db)
	logger.Info("Connected to database")

	// Models
	taskModel := models.NewTaskModel(&connectDB)

	//Create Middlewares for add Models in context
	app.Use(func(c fiber.Ctx) error {
		c.Locals("TaskModel", taskModel)
		return c.Next()
	})

	// Routes
	app.Get("/tasks", handler.GetTasks)
	app.Post("/tasks", handler.CreateTask)
	app.Delete("/tasks/:id", handler.DeleteTask)
	app.Put("tasks/:id", handler.UpdateTask)

	// Swagger
	//docs.SwaggerInfo.BasePath = "/api"
	//app.Get("/swagger/*", swagger.HandlerDefault)

	// Остановка CTRL+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-stop
		logger.Info("Shutting down...")
		log.Fatal(app.Shutdown())
	}()

	// Listener
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)))

}
