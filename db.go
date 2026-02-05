package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	ID       int64  `db:"id"       json:"id"`
	Name     string `db:"name"     json:"name"`
	Email    string `db:"email"    json:"email"`
	Password string `db:"password" json:"password"`
}
type Student struct {
	ID   int64   `db:"id"   json:"id"`
	Name string  `db:"name" json:"name"`
	Mark float32 `db:"mark" json:"mark"`
}

func GainAccessToDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbPassword := os.Getenv("DBPASS")
	if dbPassword == "" {
		return nil, err
	}

	dbUser := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")

	databaseSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", databaseSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTables(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS users
	(
	 	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL ,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	);`)
	if err != nil {

		log.Println(err)
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS students (
	 	id BIGINT UNSIGNED PRIMARY KEY,
		name VARCHAR(50) NOT NULL ,
		mark DECIMAL(5,2) NOT NULL); `)
	if err != nil {

		log.Println(err)
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
	}
	return nil
}

func createUser(db *sql.DB, name string, email string, password string) error {
	insertionResult, err := db.Exec(`
	INSERT INTO users ( name, email, password) VALUES (?, ?, ?  );
	`, name, email, password)
	if err != nil {
		return err
	}
	rows, err := insertionResult.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		log.Printf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}

func InsertStudent(db *sql.DB, id int64, name string, mark float32) error {
	insertionResult, err := db.Exec(`
	INSERT INTO students (id, name, mark) VALUES (?,?, ? );
	`, id, name, mark)
	if err != nil {
		return err
	}
	rows, err := insertionResult.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		log.Printf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}

func UpdateStudent(db *sql.DB, id int64, name string, mark float32) error {
	insertionResult, err := db.Exec(`
	UPDATE students SET  name = ?, mark = ?
	WHERE id = ?
	`, name, mark, id)
	if err != nil {
		return err
	}
	rows, err := insertionResult.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		log.Printf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}

func GetAllStudents(db *sql.DB) ([]Student, error) {
	rows, err := db.Query(`select * from students `)
	if err != nil {
		return []Student{}, err
	}

	defer rows.Close()

	var students []Student

	for rows.Next() {
		var std Student
		if err := rows.Scan(&std.ID, &std.Name, &std.Mark); err != nil {
			log.Printf("\ncan't copy value to user struct : %v", err)
		}
		students = append(students, std)
	}
	if err = rows.Err(); err != nil {
		return []Student{}, err
	}

	return students, nil
}
