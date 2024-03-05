package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID         int    `json:"id"`
	Subject    string `json:"subject"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Content    string `json:"content"`
	RemoteHost string `json:"remote_host"`
	UserAgent  string `json:"user_agent"`
}

func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var post Post
			if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			_, err := db.Exec("INSERT INTO posts (subject, name, email, content, remote_host, user_agent) VALUES (?, ?, ?, ?, ?, ?)",
				post.Subject, post.Name, post.Email, post.Content, r.RemoteAddr, r.UserAgent())

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
		} else {
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
