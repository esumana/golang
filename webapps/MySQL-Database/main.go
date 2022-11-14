// Installing the go-sql-driver/mysql package
// The Go programming language comes with a handy package called `database/sql` to query all sorts of
// SQL databases.
// To install the MySQL database driver, go to your terminal of choice and run:
//
//			go get -u github.com/go-sql-driver/mysql

// Connecting to a MySQL database
// The first thing we need to check after installing all necessary packages is, if we can connect
// to our MySQL database successfully.
// To check if we can connect to our database, import the database/sql and the go-sql-driver/mysql
// package and open up a connection like so:
//			import "database/sql"
//			import _ "go-sql-driver/mysql"

// Configure the database connection (always check errors)
//			db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")

// Initialize the first connection to the database, to see if everything works correctly.
// Make sure to check the error.
//			err := db.Ping()

// Creating our first database table
// Every data entry in our database is stored in a specific table.
//			CREATE TABLE users (
//			    id INT AUTO_INCREMENT,
//			    username TEXT NOT NULL,
//			    password TEXT NOT NULL,
//			    created_at DATETIME,
//			    PRIMARY KEY (id)
//			);
// Now that we have our SQL command, we can use the database/sql package to create the table in our
// MySQL database:
//			query := `
//			CREATE TABLE users (
//				id INT AUTO_INCREMENT,
//				username TEXT NOT NULL,
//				password TEXT NOT NULL,
//				created_at DATETIME,
//				PRIMARY KEY (id)
//			);`
// Executes the SQL query in our database. Check err to ensure there was no error.
//			_, err := db.Exec(query)

// Inserting our first user
// To insert our first user into our database table, we create a SQL query like the following.
// INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)
// We can now use this SQL query in Go and insert a new row into our table:
//			import "time"
//
//			username := "johndoe"
//			password := "secret"
//			createdAt := time.Now()
//
//			// Inserts our data into the users table and returns with the result and a possible error.
//			// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
//			result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
// To grab the newly created id for your user simply get it like this:
//
//			userID, err := result.LastInsertId()

// Querying our users table
// In Go we first declare some variables to store our data in and then query a single database row like so:
//
//			var (
//			    id        int
//			    username  string
//			    password  string
//			    createdAt time.Time
//			)
//
//			// Query the database and scan the values into out variables. Don't forget to check for errors.
//			query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
//			err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)

// Querying all users
// We can use the SQL command from the example above and trim off the WHERE clause. This way, we query
// all existing users.
// In Go we first declare some variables to store our data in and then query a single database row like so:
//
//			type user struct {
//			    id        int
//			    username  string
//			    password  string
//			    createdAt time.Time
//			}
//
//			rows, err := db.Query(`SELECT id, username, password, created_at FROM users`) // check err
//			defer rows.Close()
//
//			var users []user
//			for rows.Next() {
//			    var u user
//			    err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt) // check err
//			    users = append(users, u)
//			}
//			err := rows.Err() // check err
//
// The users slice now might contain something like this:
//
//			users {
//			    user {
//			        id:        1,
//			        username:  "johndoe",
//			        password:  "secret",
//			        createdAt: time.Time{wall: 0x0, ext: 63701044325, loc: (*time.Location)(nil)},
//			    },
//			    user {
//			        id:        2,
//			        username:  "alice",
//			        password:  "bob",
//			        createdAt: time.Time{wall: 0x0, ext: 63701044622, loc: (*time.Location)(nil)},
//			    },
//			}

// Deleting a user from our table
//
//	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1) // check err
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/root?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Create a new table
		query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Insert a new user
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
