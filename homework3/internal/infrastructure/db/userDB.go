package db

import (
	"database/sql"
	"fmt"
	"homework3/internal/domain/entity"

	_ "github.com/lib/pq"
)

func getDB(dbHost, dbPort, dbUser, dbPass, dbName string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("Postgress error: ", err.Error())
	}

	return db, nil

}

type PostgresUserRepository struct {
	db *sql.DB
}

func GetRepository(dbHost, dbPort, dbUser, dbPass, dbName string) (*PostgresUserRepository, error) {
	rep := new(PostgresUserRepository)
	db, err := getDB(dbHost, dbPort, dbUser, dbPass, dbName)
	if (err != nil) {
		return nil, err
	}
	rep.db = db
	return rep, nil
}

func (r *PostgresUserRepository) GetUser(phoneNumber string) (*entity.User, error) {
	query := "SELECT name, ruble, kopeck FROM users WHERE phoneNumber = $1"
	row := r.db.QueryRow(query, phoneNumber)

	var (
		name   string
		ruble  int
		kopeck int8
	)

	err := row.Scan(&name, &ruble, &kopeck)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with phone number %s not found", phoneNumber)
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	user, err := entity.GetUser(name, ruble, kopeck, phoneNumber)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) SaveUser(user *entity.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := "INSERT INTO users (name, ruble, kopeck, phoneNumber) VALUES ($1, $2, $3, $4)"
	_, err = tx.Exec(query,
		user.GetName(),
		user.GetRuble(),
		user.GetKopeck(),
		user.GetPhoneNumber(),
	)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}