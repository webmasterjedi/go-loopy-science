package db

import (
	"database/sql"
	"errors"
	"fmt"
	"goloopyscience/loopy/dscanner/types"
	_ "modernc.org/sqlite"
)

var (
	ErrDup      = errors.New("record already exists")
	ErrNoRecord = errors.New("record not found")
)

// GetDB returns a pointer to the loopy database
func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./loopy.db")
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
		if err.Error() == "constraint failed: UNIQUE constraint failed: Systems.SystemAddress (1555)" {
			fmt.Print(ErrDup)
			return nil
		}
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
