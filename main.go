package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Event struct {
	name string
}

type Subscriber struct {
	name string
	fn   func(wg *sync.WaitGroup, e Event)
}

const DB_NAME = "./bus.db"

type subscribers []Subscriber

// TODO: Add sqlite3 
// TODO: Add subscriber to the insert into db action which will emit the
// events, something along those lines, the idea is to use the pub/sub pattern
// here and try not to block on posting to the api 
// TODO: Cleanup, we have two main libs -> pub/sub and database, so I think
// it's time to tidy up before things get too messy
func main() {
	e := Event{}
	subscribers := []Subscriber{}

	db := setupDatabase()
	defer db.Close()

	sub1 := Subscriber{"sub1", func(wg *sync.WaitGroup, e Event) {
		defer wg.Done()
		d, _ := time.ParseDuration("1s")
		time.Sleep(d)
		fmt.Printf("Inside sub 1\n")
	}}
	subscribers = append(subscribers, sub1)

	sub2 := Subscriber{"sub2", func(wg *sync.WaitGroup, e Event) {
		defer wg.Done()
		d, _ := time.ParseDuration("3s")
		time.Sleep(d)
		fmt.Printf("Inside sub 2\n")
	}}
	subscribers = append(subscribers, sub2)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /", postPublish(db))

	publish(subscribers, e)

  err := http.ListenAndServe(":6969", mux)
  if err != nil {
		log.Printf("%q\n", err)
  }
}

// Publish functions
func publish(s []Subscriber, e Event) {
	var wg sync.WaitGroup
	wg.Add(len(s))

	for _, sub := range s {
		fmt.Printf("%s\n", sub.name)
		go sub.fn(&wg, e)
	}

	wg.Wait()
}

func postPublish(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
  return func (w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")

    insertBy(db, name, "Custom Event")
    // publish(subscribers, e)
    dumpDatabase(db)

		// log.Printf("%q\n", r.Body)
    io.WriteString(w, "This is my website!\n")
  }
}

// Database functions
func setupDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
  CREATE TABLE IF NOT EXISTS events(
    id integer not null primary key autoincrement, 
    name text,
    type text
  );
  `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}

	return db
}

func insertBy(db *sql.DB, name, t_type string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into events(name, type) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, t_type)
	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
}

func dumpDatabase(db *sql.DB) {
	rows, err := db.Query("select id, name, type from events")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var type_t string

		err = rows.Scan(&id, &name, &type_t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, type_t)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
