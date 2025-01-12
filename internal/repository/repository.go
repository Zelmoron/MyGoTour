package repository

import (
	"Tour/internal/requests"
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func Connect() *sql.DB {
	user := os.Getenv("name")           //user Postgres
	password := os.Getenv("dbpassword") //password Postgres
	dbname := os.Getenv("dbname")       //name of the database
	host := os.Getenv("dbhost")         //host
	port := os.Getenv("dbport")         //database`s port

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		logrus.Fatalf("Failed to connect to database: %v", err)
		return nil
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logrus.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		logrus.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Поднять миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.Fatalf("Failed to apply migrations: %v", err)
	}
	logrus.Info("Succes migrations")

	// if err := m.Down(); err != nil && err != migrate.ErrNoChange {
	// 	logrus.Fatalf("Failed to delete migrations: %v", err)
	// }

	return db
}

func TestConnect() *sql.DB {
	user := "testuser"     //user Postgres
	password := "testpass" //password Postgres
	dbname := "testdb"     //name of the database
	host := "localhost"    //host
	port := "5433"         //database`s port

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		logrus.Fatalf("Failed to connect to database: %v", err)
		return nil
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logrus.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		logrus.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Поднять миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.Fatalf("Failed to apply migrations: %v", err)
	}
	logrus.Info("Succes migrations")

	// if err := m.Down(); err != nil && err != migrate.ErrNoChange {
	// 	logrus.Fatalf("Failed to delete migrations: %v", err)
	// }

	return db
}

func (r *Repository) SelectUser(ctx context.Context, ch1 chan<- bool, errChan chan bool, user requests.RegistrationRequest) {

	if user.Name == "err" {
		errChan <- false

	}

	select {
	case <-ctx.Done():
		log.Info("SelectUser прерван по контексту")
		return
	default:
		if user.Name == "A" {
			log.Info("Запись ch1 - false")
			ch1 <- false
		} else {
			log.Info("Запись ch1 - true")
			ch1 <- true
		}
	}

}

func (r *Repository) InsertUser(ctx context.Context, ch1 <-chan bool, ch2 chan<- bool, errChan chan bool, user requests.RegistrationRequest) {

	select {
	case <-ctx.Done():
		log.Info("InsertUser прерван по контексту")
		return
	case flag := <-ch1:
		if flag {
			log.Info("Запись ch2 - true")
			ch2 <- true
		} else {
			log.Info("Запись ch2 - false")
			ch2 <- false
		}
	}
}
