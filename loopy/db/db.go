package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"goloopyscience/loopy/dscanner/types"
)

var (
	ErrDup      = errors.New("record already exists")
	ErrNoRecord = errors.New("record not found")
)

func handleSQLiteError(err error) error {
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
			return ErrDup
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		return ErrNoRecord
	}
	return err
}

// GetDB returns a pointer to the loopy database
func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./loopy.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateTables creates the tables in the loopy database
func CreateTables() error {
	// Get the loopy database
	db, err := GetDB()
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	_, err = db.Exec(createSystems)
	if err != nil {
		fmt.Print(err)
		return err
	}

	_, err = db.Exec(createStars)
	if err != nil {
		fmt.Print(err)
		return err
	}

	_, err = db.Exec(createBodies)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func InsertSystem(system *types.StarSystem) error {
	// Get the loopy database
	db, err := GetDB()
	if err != nil {
		fmt.Print(err)
		return err
	}
	stmt, err := db.Prepare(insertSystemSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(system.FSDJumpEvent.SystemAddress, system.FSDJumpEvent.StarSystem, system.FSDJumpEvent.Body, system.FSDJumpEvent.BodyID, system.FSDJumpEvent.BodyType)

	if err != nil {
		//check sql error type for unique constraint
		err = handleSQLiteError(err)
		if err == ErrDup {
			return nil
		}
		fmt.Print(err)
		return err
	}

	newSystemId, err := res.LastInsertId()
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Print(newSystemId)

	/*for _, star := range system.Stars {

	}*/

	return nil
}
