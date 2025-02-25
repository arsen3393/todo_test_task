package models

import "github.com/jackc/pgx/v5"

type DBconnection struct {
	db *pgx.Conn
}

func (db *DBconnection) SetDB(pointerDB *pgx.Conn) {
	db.db = pointerDB
}

type Model interface {
	Init(db *DBconnection)
}

func (db *DBconnection) GetDB() *pgx.Conn {
	return db.db
}
