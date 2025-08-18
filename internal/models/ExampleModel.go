package models

import "database/sql"

type ExampleModel struct {
	DB *sql.DB
}

func (e *ExampleModel) ExampleTransaction() error {
	tx, err := e.DB.Begin()
	if err != nil {
		return err
	}
	// defer call to rollback to ensure that any errors that occur
	// or if the function fails while executing the transaction is rolled back
	defer tx.Rollback()

	// Call Exec() on the transaction, passing in your statement and any
	// parameters. It's important to notice that tx.Exec() is called on the
	// transaction object just created, NOT the connection pool. Although
	// we're using tx.Exec() here you can also use tx.Query() and tx.QueryRow() in
	// exactly the same way.
	_, err = tx.Exec("INSERT INTO ...")
	if err != nil {
		return err
	}

	// same exec with a different database query
	_, err = tx.Exec("UPDATE ....")
	if err != nil {
		return err
	}

	// if we have no error so far then we commit the transaction
	err = tx.Commit()
	return err
}
