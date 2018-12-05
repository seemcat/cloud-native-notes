package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Note struct {
	Title     string `json:"title"`
	Note      string `json:"note"`
	Handle    string `json:"handle"`
	Urlhandle string `json:"urlhandle"`
}

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := store.GetNotes()
	noteListBytes, err := json.Marshal(notes)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(noteListBytes)
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	note := Note{}
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	note.Title = r.Form.Get("title")
	note.Note = r.Form.Get("note")
	note.Handle = r.Form.Get("handle")
	note.Urlhandle = r.Form.Get("urlhandle")

	err = store.CreateNote(&note)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)
}
