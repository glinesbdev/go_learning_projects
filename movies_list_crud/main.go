package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var DATABASE = "db/movies.db"

type Movie struct {
	ID       int    `json:"-"`
	Title    string `json:"title"`
	Isbn     string `json:"isbn"`
	Director *Director
}

type Director struct {
	ID        int    `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getAllMovies(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sql := `
		select movies.title, movies.isbn,
			   directors.first_name, directors.last_name
		from movies
		join directors on movies.id = directors.movie_id;
		`
		result, err := db.Query(sql)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		defer result.Close()

		var movies []Movie

		for result.Next() {
			movie := Movie{Director: &Director{}}
			if err := result.Scan(&movie.Title, &movie.Isbn, &movie.Director.FirstName, &movie.Director.LastName); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "%v", err)
				return
			}

			movies = append(movies, movie)
		}

		json.NewEncoder(w).Encode(movies)
	}
}

func getMovie(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		sql := "select movies.title, movies.isbn from movies where id = ?;"
		result := db.QueryRow(sql, id)

		var movie Movie
		result.Scan(&movie.Title, &movie.Isbn)

		json.NewEncoder(w).Encode(movie)
	}
}

func createMovie(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		catch_err_func := func(err error) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "%v", err)
				return
			}
		}

		tx, err := db.BeginTx(r.Context(), nil)
		catch_err_func(err)
		defer tx.Rollback()

		var movie Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)

		movie_sql := "insert into movies (title, isbn) values (?, ?) returning id;"
		movie_id_result, err := tx.Exec(movie_sql, movie.Title, movie.Isbn)
		catch_err_func(err)

		movie_id, err := movie_id_result.LastInsertId()
		catch_err_func(err)

		director_sql := "insert into directors (first_name, last_name, movie_id) values (?, ?, ?)"
		_, err = tx.Exec(director_sql, movie.Director.FirstName, movie.Director.LastName, movie_id)
		catch_err_func(err)
		catch_err_func(tx.Commit())

		json.NewEncoder(w).Encode(movie)
	}
}

func updateMovie(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var movie Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)

		id := mux.Vars(r)["id"]
		sql := `
		update movies
		set title = ?, isbn = ?
		where id = ?	
		`

		_, err := db.Exec(sql, movie.Title, movie.Isbn, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%v", err)
		}

		json.NewEncoder(w).Encode(movie)
	}
}

func deleteMovie(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		sql := `
		delete from movies where id = ?;
		delete from directors where movie_id = ?;
		`
		_, err := db.Exec(sql, id, id)

		if err != nil {
			fmt.Fprintf(w, "%v", err)
			return
		}
	}
}

func reigsterRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/movies", getAllMovies(db)).Methods(http.MethodGet)
	r.HandleFunc("/movies", createMovie(db)).Methods(http.MethodPost)
	r.HandleFunc("/movies/{id}", getMovie(db)).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}", deleteMovie(db)).Methods(http.MethodDelete)
	r.HandleFunc("/movies/{id}", updateMovie(db)).Methods(http.MethodPatch, http.MethodPut)
}

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", DATABASE)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	return db
}

func main() {
	db := connectDB()
	defer db.Close()

	r := mux.NewRouter()
	reigsterRoutes(r, db)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
