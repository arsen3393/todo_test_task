package main

import (
	"testSkillsRock/internal/config"
	"testSkillsRock/pkg/api/app"
	slogger "testSkillsRock/pkg/logger"
)

func main() {

	logger := slogger.MustInitLogger()

	cfg := config.MustLoadConfig()

	app.RunApp(cfg, logger)

	// TODO MAKE SWAGGER
}
