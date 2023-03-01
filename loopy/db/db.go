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

func handleDBError(err error) error {
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
		fmt.Print("\n", err, "\n")
		return err
	}
	defer db.Close()

	_, err = db.Exec(createSystems)
	if err != nil {
		fmt.Print("\n", err, "\n")
		return err
	}

	_, err = db.Exec(createStars)
	if err != nil {
		fmt.Print("\n", err, "\n")
		return err
	}

	_, err = db.Exec(createBodies)
	if err != nil {
		fmt.Print("\n", err, "\n")
		return err
	}
	return nil
}

func InsertSystem(system *types.StarSystem) error {
	// Get the loopy database
	db, err := GetDB()
	if err != nil {
		fmt.Print("\n", err, "\n")
		return err
	}
	stmt, err := db.Prepare(insertSystemSQL)
	if err != nil {
		fmt.Print("\n", err, "\n")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(system.FSDJumpEvent.SystemAddress, system.FSDJumpEvent.StarSystem, system.FSDJumpEvent.Body, system.FSDJumpEvent.BodyID, system.FSDJumpEvent.BodyType)

	if err != nil {
		//check sql error type for unique constraint
		err = handleDBError(err)
		if err == ErrDup {
			return nil
		}
		fmt.Print("\n", err, "\n")
		return err
	}

	return nil
}

func InsertStar(star *types.Star) error {
	// Get the loopy database
	db, err := GetDB()
	if err != nil {
		fmt.Print("\n", err, "\n")
		return err
	}
	stmt, err := db.Prepare(insertStarSQL)
	if err != nil {
		fmt.Print("\n", err, "\n")
		return err
	}
	defer stmt.Close()
	//TODO: get parent star/body id from Parents array

	_, err = stmt.Exec(0, star.BodyName, star.BodyID, star.ParentsToJson(), star.SystemAddress, star.StarType, star.Subclass, star.StellarMass, star.Radius, star.AbsoluteMagnitude, star.AgeMY, star.SurfaceTemperature, star.Luminosity, star.SemiMajorAxis, star.Eccentricity, star.OrbitalInclination, star.Periapsis, star.OrbitalPeriod, star.RotationPeriod, star.AxialTilt, star.RingsToJson(), star.WasDiscovered, star.WasMapped)

	if err != nil {
		//check sql error type for unique constraint
		err = handleDBError(err)
		if err == ErrDup {
			return nil
		}
		fmt.Print("\n", err, "\n")
		return err
	}

	return nil
}

func InsertBody(body *types.Body) error {
	// Get the loopy database
	db, err := GetDB()
	if err != nil {
		fmt.Print(err)
		return err
	}
	stmt, err := db.Prepare(insertBodySQL)
	if err != nil {
		fmt.Print("\n", err, "\n")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(0, body.BodyName, body.BodyID, body.ParentsToJson(), 0, body.SystemAddress, body.TidalLock, body.TerraformState, body.PlanetClass, body.Atmosphere, body.AtmosphereType, body.AtmosphereCompositionToJson(), body.Volcanism, body.MassEM, body.Radius, body.SurfaceGravity, body.SurfaceTemperature, body.SurfacePressure, body.Landable, body.MaterialsToJson(), body.BodyCompositionToJson(), body.SemiMajorAxis, body.Eccentricity, body.OrbitalInclination, body.Periapsis, body.OrbitalPeriod, body.RotationPeriod, body.AxialTilt, body.RingsToJson(), body.WasDiscovered, body.WasMapped)
	if err != nil {
		//check sql error type for unique constraint
		err = handleDBError(err)
		if err == ErrDup {
			return nil
		}
		fmt.Print("\n", err, "\n")
		return err
	}
	return nil
}
