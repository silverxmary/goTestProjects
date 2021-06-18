package database

import (
	"log"

	"git.manomano.tech/mariellys.soto/go-rest/app/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostDB interface {
	Open() error
	Close() error
	CreatePost(p *models.Post) error
	GetPosts() ([]*models.Post, error)
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	schema := "postgres"
	pg, err := sqlx.Open(schema, pgConnStr)
	if err != nil {
		return err
	}
	log.Println("Connected to Database!")

	pg.MustExec(createTable)

	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
