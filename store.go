package main

import (
	"database/sql"
)

type Store interface {
	CreateNote(note *Note) error
	GetNotes() ([]*Note, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateNote(note *Note) error {
	_, err := store.db.Query("INSERT INTO notes(title, note, handle, urlhandle) VALUES ($1, $2, $3, $4)", note.Title, note.Note, note.Handle, note.Urlhandle)
	return err
}

func (store *dbStore) GetNotes() ([]*Note, error) {
	rows, err := store.db.Query("SELECT title, note, handle, urlhandle from notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []*Note{}
	for rows.Next() {
		note := &Note{}
		if err := rows.Scan(&note.Title, &note.Note, &note.Handle, &note.Urlhandle); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

var store Store

func InitStore(s Store) {
	store = s
}
