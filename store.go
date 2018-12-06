package main

// we'll need the sql go library to play with the db.
import (
	"database/sql"
)

// methods of our store. if something goes wrong on
// either of them, an error is returned. interfaces are 
// named collections of method signatures. any type 
// that defines 'CreateNote' or 'GetNotes' method is 
// said to satisfy the Store interface.
type Store interface {
	CreateNote(note *Note) error
	GetNotes() ([]*Note, error)
}

// since there's no implements keyword in Go, we'll need to
// create a 'dbStore' type that satisfies the Store interface
// it takes in the sql DB connection object, which represents 
// the db connection.
type dbStore struct {
	db *sql.DB
}

// 'Note' is a simple struct with 'title', 'note', 'handle' & 'urlhandle' attributes
// 'CreateNote' has a receiver type of dbStore
func (store *dbStore) CreateNote(note *Note) error {

	// the first underscore means that we don't care about what's returned from
	// this insert query. we just want to know if it was inserted correctly,
	// and the error will be populated if it wasn't
	_, err := store.db.Query("INSERT INTO notes(title, note, handle, urlhandle) VALUES ($1, $2, $3, $4)", note.Title, note.Note, note.Handle, note.Urlhandle)
	return err
}

func (store *dbStore) GetNotes() ([]*Note, error) {
	// query the database for all notes, and return the result to the
	// `rows` object
	rows, err := store.db.Query("SELECT title, note, handle, urlhandle from notes")
	// return in case of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// create the data structure that is returned from the function.
	// by default, this will be an empty array of notes
	notes := []*Note{}
	for rows.Next() {
		// for each row returned by the table, create a pointer to a note,
		note := &Note{}
		// populate the 'Title', 'Note', 'Handle', and 'Urlhandle' attributes of the note,
		// and return in case of an error
		if err := rows.Scan(&note.Title, &note.Note, &note.Handle, &note.Urlhandle); err != nil {
			return nil, err
		}
		// finally, append the result to the returned array, and repeat for
		// the next row
		notes = append(notes, note)
	}
	return notes, nil
}

// the store variable is a package level variable that will be available for
// use throughout our application code
var store Store

/*
we'll need to call the InitStore method to initialize the store. this will
typically be done at the beginning of our application (in this case, when the server starts up)
*/ 
func InitStore(s Store) {
	store = s
}
