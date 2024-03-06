package db

import (
	"bankassignment/models"
	"context"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

func ConnectToDB() *pg.DB {
	DB = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "admin",
		Database: "bank",
		Addr:     "localhost:5432",
	})
	if err := DB.Ping(context.Background()); err != nil {
		log.Fatalln("Could not connect to the database ")
	}

	log.Println("Database connected successfully!")

	err := createSchema(DB)
	if err != nil {
		log.Fatal(err)
	}
	return DB
}

func createSchema(DB *pg.DB) error {
	models := []interface{}{
		(*models.Bank)(nil),
		(*models.Branch)(nil),
		(*models.Account)(nil),
		(*models.Customer)(nil),
		(*models.Transaction)(nil),
		(*models.Mapping)(nil),
	}

	for _, model := range models {
		err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
		log.Println("Table created successfully!")
	}
	return nil
}
