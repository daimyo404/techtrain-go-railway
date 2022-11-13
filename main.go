package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// res, err = sqlStatement.Exec("orange", 154)
	// checkError(err)
	// rowCount, err = res.RowsAffected()
	// fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	// res, err = sqlStatement.Exec("apple", 100)
	// checkError(err)
	// rowCount, err = res.RowsAffected()
	// fmt.Printf("Inserted %d row(s) of data.\n", rowCount)
	// fmt.Println("Done.")

	http.HandleFunc("/user/create", createUser)
	// http.HandleFunc("/user/create", getUser)
	// http.HandleFunc("/user/create", updateUser)

	http.ListenAndServe(":8080", nil)
}

func createUser(w http.ResponseWriter, r *http.Request) {

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

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)

	// Insert some data into table.
	sqlStatement, err := db.Prepare("INSERT INTO users (name, token) VALUES (?, ?);")
	res, err := sqlStatement.Exec("hoge", "fuga")
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

}

// func getUser(w http.ResponseWriter, r *http.Request) {

// }

// func updateUser(w http.ResponseWriter, r *http.Request) {

// }
