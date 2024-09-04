package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/lahnasti/clientServerApp/server/models"
)

type DBstorage struct {
	conn *pgx.Conn
}

func NewDB(conn *pgx.Conn) (DBstorage, error) {
	return DBstorage{
		conn: conn,
	}, nil
}

func (db DBstorage) RegisterUser(user models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `INSERT INTO users(login, password) VALUES ($1, $2) RETURNING uid`
	var uid int
	err := db.conn.QueryRow(ctx, query, user.Login, user.Password).Scan(&uid)
	if err != nil {
		return -1, err
	}
	return uid, nil
}
