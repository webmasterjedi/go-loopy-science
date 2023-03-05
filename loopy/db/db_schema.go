package db

// create the Systems table if it doesn't exist
var createSystems = `
create table if not exists Systems
(
	SystemAddress integer
		primary key
		unique,
	StarSystem text not null
		unique,
	Body       text,
	BodyID     integer,
	BodyType   text
);`

// create the Stars table if it doesn't exist
var createStars = `
create table if not exists Stars
(
	ID                 integer not null
		constraint Stars_pk
			primary key autoincrement,
	ParentID           integer,
	ParentType         text,
	BodyName           text unique,
	BodyID             integer,
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
	Rings              blob,
	WasDiscovered      boolean,
	WasMapped          boolean
);`

// check if the Planets table exists, if not create it
var createBodies = `
create table if not exists Bodies
(
	ID                    integer not null
		constraint Bodies_pk
			primary key autoincrement,
	ParentID              integer,
	ParentType            text,
	BodyName           text not null unique,
	BodyID             integer,
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
	BodyComposition           blob,
	SemiMajorAxis         real,
	Eccentricity          real,
	OrbitalInclination    real,
	Periapsis             real,
	OrbitalPeriod         real,
	RotationPeriod        real,
	AxialTilt             real,
	Rings                 blob,
	WasDiscovered         boolean,
	WasMapped             boolean
);`

// insert a system into the Systems table
var insertSystemSQL = `
insert into Systems
    (SystemAddress, StarSystem, Body, BodyID, BodyType)
values (?, ?, ?, ?, ?);`

// insert a star into the Stars table
var insertStarSQL = `
insert into Stars
	(ParentID, ParentType, BodyName, BodyID, SystemAddress, StarType, Subclass, StellarMass, Radius, AbsoluteMagnitude, AgeMY, SurfaceTemperature, Luminosity, SemiMajorAxis, Eccentricity, OrbitalInclination, Periapsis, OrbitalPeriod, RotationPeriod, AxialTilt, Rings, WasDiscovered, WasMapped)
values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

// insert a body into the Bodies table
var insertBodySQL = `
insert into Bodies
	(ParentID, ParentType, BodyName, BodyID, SystemAddress, TidalLock, TerraformState, PlanetClass, Atmosphere, AtmosphereType, AtmosphereComposition, Volcanism, MassEM, Radius, SurfaceGravity, SurfaceTemperature, SurfacePressure, Landable, Materials, BodyComposition, SemiMajorAxis, Eccentricity, OrbitalInclination, Periapsis, OrbitalPeriod, RotationPeriod, AxialTilt, Rings, WasDiscovered, WasMapped)
values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
