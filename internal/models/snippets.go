package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	modifiedExpires := fmt.Sprintf("%d days", expires)
	stmt := "INSERT INTO snippets(title, content, created, expires) VALUES ($1, $2, current_timestamp, current_timestamp + $3::interval) RETURNING id"
	var id int64
	// Use the DB.QueryRow method to execute the statement and scan the returned id into the id variable.
	err := m.DB.QueryRow(stmt, title, content, modifiedExpires).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > current_timestamp AND id = $1`
	snippet := &Snippet{}
	err := m.DB.QueryRow(stmt, id).Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return snippet, nil
}
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := "SELECT id, title, content, created, expires FROM snippets WHERE expires > current_timestamp ORDER BY id DESC LIMIT 10;"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// always defer close rows after possible error from the query call
	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
	// error that was encountered during the iteration. It's important to
	// call this - don't assume that a successful iteration was completed
	// over the whole resultset

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
