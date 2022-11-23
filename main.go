package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Users struct {
	name  string
	token string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/update", updateUser)

	http.ListenAndServe(":8080", nil)
}

func setConnectionString() string {

	err := godotenv.Load("./.env")
	if err != nil {
		panic("Error loading .env file")
	}

	host := os.Getenv("HOST_NAME")
	database := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER_NAME")
	password := os.Getenv("DATABASE_USER_PASSWORD")

	// Initialize connection string.
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	return connectionString

}

func createUser(w http.ResponseWriter, r *http.Request) {

	connectionString := setConnectionString()

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)

	sqlStatement, err := db.Prepare("INSERT INTO users (name, token) VALUES (?, ?);")
	res, err := sqlStatement.Exec("hogehoge", "fugafuga")
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

}

func getUser(w http.ResponseWriter, r *http.Request) {

	connectionString := setConnectionString()

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)

	// Read users table.
	rows, err := db.Query("SELECT * FROM users;")
	checkError(err)
	defer rows.Close()
	fmt.Println("Reading data:")

	//TODO: 一旦100にしている
	users := [100]Users{}

	for i := 0; rows.Next(); i++ {
		err := rows.Scan(&users[i].name, &users[i].token)
		checkError(err)
		fmt.Printf("Data row = (%s, %s)\n", users[i].name, users[i].token)
	}

	err = rows.Err()
	checkError(err)
	fmt.Println("Done.")

}

func updateUser(w http.ResponseWriter, r *http.Request) {

	connectionString := setConnectionString()

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)

	// Modify some data in table.
	rows, err := db.Exec("UPDATE users SET name = ? WHERE name = ?", "hogee", "hoge")
	checkError(err)
	rowCount, err := rows.RowsAffected()
	fmt.Printf("Updated %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")

}
