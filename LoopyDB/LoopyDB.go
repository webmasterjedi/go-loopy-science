package LoopyDB

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// GetLoopyDB returns a pointer to the LoopyDB database
func GetLoopyDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./Loopy.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateTables creates the tables in the LoopyDB database
func CreateTables() error {
	// Get the LoopyDB database
	db, err := GetLoopyDB()
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
	fmt.Print("Creating tables...")

	//check if the Systems table exists, if not create it
	_, err = db.Exec(`
		create table if not exists Systems
		(
			SystemAddress integer
				primary key
				unique,
			SystemName text not null
				unique,
			Body       text,
			BodyID     integer,
			BodyType   text
		);`)
	if err != nil {
		fmt.Print(err)
		return err
	}
	//check if the Stars table exists, if not create it
	_, err = db.Exec(`
		create table if not exists Stars
		(
			Id                 integer not null
				constraint Stars_pk
					primary key autoincrement,
			ParentId           integer,
			SystemAddress      integer,
			StarType           text,
			Subclass           integer,
			StellarMass        real,
			Radius             real,
			AbsoluteMagnitude  real,
			AgeMY              real,
			SurfaceTemperature real,
			Luminosity         text,
			SemiMajorAxis      real,
			Eccentricity       real,
			OrbitalInclination real,
			Periapsis          real,
			OrbitalPeriod      real,
			RotationPeriod     real,
			AxialTilt          real,
			WasDiscovered      boolean,
			WasMapped          boolean
		);`)
	if err != nil {
		fmt.Print(err)
		return err
	}

	//check if the Planets table exists, if not create it
	_, err = db.Exec(`
		create table if not exists Planets
		(
			Id                    integer not null
				constraint Planets_pk
					primary key autoincrement,
			ParentId              integer,
			StarId                integer,
			SystemAddress         integer,
			TidalLock             boolean,
			TerraformState        text,
			PlanetClass           text,
			Atmosphere            text,
			AtmosphereType        text,
			AtmosphereComposition blob,
			Volcanism             text,
			MassEM                real,
			Radius                real,
			SurfaceGravity        real,
			SurfaceTemperature    real,
			SurfacePressure       real,
			Landable              boolean,
			Materials             blob,
			Composition           blob,
			SemiMajorAxis         real,
			Eccentricity          real,
			OrbitalInclination    real,
			Periapsis             real,
			OrbitalPeriod         real,
			RotationPeriod        real,
			AxialTilt             real,
			WasDiscovered         boolean,
			WasMapped             boolean
		);`)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
