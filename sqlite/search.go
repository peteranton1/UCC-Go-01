package main

import (
	"fmt"
	"log"
	//"crawshaw.io/sqlite"
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

func main() {
	
	db, err := sql.Open("sqlite3", "../data/shakespeare.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare(`
		SELECT play, act, scene, plays.text
		FROM playsearch
		INNER JOIN plays
		ON playsearch.playrowid = plays.rowid
		WHERE playsearch.text MATCH ?;`)
	if err != nil {
		log.Fatal(err)
	}
	var play, text string
	var act, scene int
	err = stmt.QueryRow("whether tis nobler").Scan(
		&play, &act, &scene, &text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %d.%d: %q\n", play, act, scene, text)
}