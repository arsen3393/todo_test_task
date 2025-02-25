package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"testSkillsRock/internal/config"
)

func MustConnectDB(cfg config.DataBaseConfig) *pgx.Conn {
	dataBaseUrl := fmt.Sprintf("postgres://%s:%s@localhost:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Port, cfg.Name)
	db, err := pgx.Connect(context.Background(), dataBaseUrl)
	if err != nil {
		panic("Cannot connect to db: " + err.Error())
	}
	return db
}
