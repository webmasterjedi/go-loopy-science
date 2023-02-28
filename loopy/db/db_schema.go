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
);`

// check if the Planets table exists, if not create it
var createBodies = `
create table if not exists Bodies
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
	BodyComposition           blob,
	SemiMajorAxis         real,
	Eccentricity          real,
	OrbitalInclination    real,
	Periapsis             real,
	OrbitalPeriod         real,
	RotationPeriod        real,
	AxialTilt             real,
	WasDiscovered         boolean,
	WasMapped             boolean
);`

// insert a system into the Systems table
var insertSystemSQL = `
insert into Systems
    (SystemAddress, StarSystem, Body, BodyID, BodyType)
values (?, ?, ?, ?, ?);`
