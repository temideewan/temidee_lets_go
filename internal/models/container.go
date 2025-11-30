package models

import "database/sql"

type ContainerModel struct {
	Snippets SnippetModel
	Users    UserModel
}

func NewContainerModel(db *sql.DB) ContainerModel {

	return ContainerModel{
		Snippets: SnippetModel{DB: db},
		Users:    UserModel{DB: db},
	}
}
