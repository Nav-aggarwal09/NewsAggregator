package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // db/sql needs a driver
	log "github.com/sirupsen/logrus"
)

// TODO: Use this sample code to connect w signup.html

const (
	host   = "localhost"
	port   = 5432
	user   = "arnav"
	pswd   = ""
	dbname = "newsapp"
)

func sample() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pswd, dbname)

	// note- this does not create connection to db. simply validates the arguments provided
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln("Could not open database - ", err)
	}
	defer db.Close()

	// note- assumes table has already been created
	sqlStatement := `INSERT INTO users(first_name, last_name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING id`

	var userid int

	err = db.QueryRow(sqlStatement, "John", "Doe", "example@test.com", "password").Scan(&userid)
	if err != nil {
		log.Fatal("Could not ping DB - ", err)
	}
	fmt.Println("New record ID is: ", userid)
}
