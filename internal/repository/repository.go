package repository

import (
	"Tour/internal/requests"
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var mu sync.Mutex

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
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(25)

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		logrus.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Поднять миграции
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		logrus.Fatalf("Failed to apply migrations: %v", err)
	}
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
		"file://../migrations",
		"postgres", driver)
	if err != nil {
		logrus.Fatalf("Failed to create migrate instance: %v", err)
	}
	// Поднять миграции
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		logrus.Fatalf("Failed to apply migrations: %v", err)
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

func (r *Repository) SelectUser(user requests.RegistrationRequest) error {

	query := "SELECT id FROM users WHERE name = $1"
	var id int
	err := r.db.QueryRow(query, user.Name).Scan(&id)

	return err

}

func (r *Repository) InsertUser(user requests.RegistrationRequest) error {

	query := "INSERT INTO users(name,password) VALUES($1,$2) "

	_, err := r.db.Exec(query, user.Name, user.Password)

	return err

}
