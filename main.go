package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/"+r.URL.Path[1:])
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("sqlite3", ":memory:?cache=shared")
		defer db.Close()

		if err != nil {
			panic(err)
		}

		password, err := generateRandomString(64)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec("CREATE TABLE users (id INT PRIMARY KEY, username TEXT, password TEXT)")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec("INSERT INTO users (id, username, password) VALUES (1, ?, ?)", "admin", password)
		if err != nil {
			panic(err)
		}

		query := "SELECT id FROM users WHERE username = '" + r.FormValue("username") + "' AND password = '" + r.FormValue("password") + "';"

		newQuery := strings.Split(query, "--")[0] // either the mysql3 library or go's database/sql doesn't like comments, so we just get rid of them in go.

		rows, err := db.Query(newQuery)
		if err != nil {
			fmt.Fprintf(w, "An error occured executing query:\n"+query+"\n\n"+err.Error())
			return
		}

		defer rows.Close()

		if rows.Next() {
			fmt.Fprintf(w, "Welcome to the site. The flag is: hackDalton{s4n4t1z3_y0ur_1nputs_6iDbwDO6ms}")
			return
		}

		fmt.Fprintf(w, "Invalid login")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)[:s], err
}
