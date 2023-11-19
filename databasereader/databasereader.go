package databasereader

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func databaseConnection () *sql.DB{
	err := godotenv.Load()
	// Handle the mysql database address connection
	databaseUserName := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	cfg := mysql.Config{
		User:                 databaseUserName,
		Passwd:               databasePassword,
		Net:                  "tcp",
		Addr:                 databaseHost,
		DBName:               databaseName,
		AllowNativePasswords: true,
	}
	DB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
	}
	DB.SetConnMaxLifetime(time.Duration(10) * time.Second)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	return DB
}


func ReadDatabaseAndDiff() {
	//Read out what is in the database and return the diff between the two
	DB := databaseConnection()
		//SQL Data Structure
	const query = `SELECT * FROM menu_item;`
	var (
		id int
		meal string
		cuisine	string
		name string
	)
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
	}
	// products := make(map[string]string)
	for rows.Next() {
		err := rows.Scan(&id, &meal, &cuisine, &name)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
		}
		fmt.Println(id, meal, cuisine, name)
	}
	
}