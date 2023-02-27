package journal

/*
Define types based on Elite Dangerous Player Journal API documentation here: https://elite-journal.readthedocs.io/en/latest/Exploration/#scan
*/

type Event interface {
	EventType() string
	// add additional methods here as needed
}

// BaseScanEvent is the base type for all scan events
type BaseScanEvent struct {
	Event    string `json:"event"`
	ScanType string `json:"ScanType"`
}

// FSDJumpEvent is the type for FSDJump events
type FSDJumpEvent struct {
	BaseScanEvent
	StarSystem    string `json:"StarSystem"`
	SystemAddress uint64 `json:"SystemAddress"`
	Body          string `json:"Body"`
	BodyID        uint64 `json:"BodyID"`
	BodyType      string `json:"BodyType"`
}

// AutoScanEvent is the type for AutoScan events
type AutoScanEvent struct {
	BaseScanEvent
	BodyName              string   `json:"BodyName"`
	BodyID                uint64   `json:"BodyID"`
	Parents               []Parent `json:"Parents"`
	StarSystem            string   `json:"StarSystem"`
	SystemAddress         uint64   `json:"SystemAddress"`
	DistanceFromArrivalLS float64  `json:"DistanceFromArrivalLS"`
	StarType              string   `json:"StarType"`
	Subclass              uint     `json:"Subclass"`
	StellarMass           float64  `json:"StellarMass"`
	Radius                float64  `json:"Radius"`
	AbsoluteMagnitude     float64  `json:"AbsoluteMagnitude"`
	AgeMY                 float64  `json:"Age_MY"`
	SurfaceTemperature    float64  `json:"SurfaceTemperature"`
	Luminosity            string   `json:"Luminosity"`
	SemiMajorAxis         float64  `json:"SemiMajorAxis"`
	Eccentricity          float64  `json:"Eccentricity"`
	OrbitalInclination    float64  `json:"OrbitalInclination"`
	Periapsis             float64  `json:"Periapsis"`
	OrbitalPeriod         float64  `json:"OrbitalPeriod"`
	RotationPeriod        float64  `json:"RotationPeriod"`
	AxisTilt              float64  `json:"AxialTilt"`
	Rings                 []struct {
		Name      string  `json:"Name"`
		RingClass string  `json:"RingClass"`
		MassMT    float64 `json:"MassMT"`
		InnerRad  float64 `json:"InnerRad"`
		OuterRad  float64 `json:"OuterRad"`
	}
	WasDiscovered bool `json:"WasDiscovered"`
	WasMapped     bool `json:"WasMapped"`
}

type DetailedScanEvent struct {
	BaseScanEvent
	BodyName              string   `json:"BodyName"`
	BodyID                uint64   `json:"BodyID"`
	Parents               []Parent `json:"Parents"`
	StarSystem            string   `json:"StarSystem"`
	SystemAddress         uint64   `json:"SystemAddress"`
	DistanceFromArrivalLS float64  `json:"DistanceFromArrivalLS"`
	TidalLock             bool     `json:"TidalLock"`
	TerraformState        string   `json:"TerraformState"`
	PlanetClass           string   `json:"PlanetClass"`
	Atmosphere            string   `json:"Atmosphere"`
	AtmosphereType        string   `json:"AtmosphereType"`
	AtmosphereComposition []struct {
		Name    string  `json:"Name"`
		Percent float64 `json:"Percent"`
	}
	Volcanism          string  `json:"Volcanism"`
	MassEM             float64 `json:"MassEM"`
	Radius             float64 `json:"Radius"`
	SurfaceGravity     float64 `json:"SurfaceGravity"`
	SurfaceTemperature float64 `json:"SurfaceTemperature"`
	SurfacePressure    float64 `json:"SurfacePressure"`
	Landable           bool    `json:"Landable"`
	Materials          []struct {
		Name    string  `json:"Name"`
		Percent float64 `json:"Percent"`
	}
	Composition struct {
		Ice   float64 `json:"Ice"`
		Rock  float64 `json:"Rock"`
		Metal float64 `json:"Metal"`
	}
	SemiMajorAxis      float64 `json:"SemiMajorAxis"`
	Eccentricity       float64 `json:"Eccentricity"`
	OrbitalInclination float64 `json:"OrbitalInclination"`
	Periapsis          float64 `json:"Periapsis"`
	OrbitalPeriod      float64 `json:"OrbitalPeriod"`
	RotationPeriod     float64 `json:"RotationPeriod"`
	AxialTilt          float64 `json:"AxialTilt"`
	Rings              []struct {
		Name      string  `json:"Name"`
		RingClass string  `json:"RingClass"`
		MassMT    float64 `json:"MassMT"`
		InnerRad  float64 `json:"InnerRad"`
		OuterRad  float64 `json:"OuterRad"`
	}
	WasDiscovered bool `json:"WasDiscovered"`
	WasMapped     bool `json:"WasMapped"`
}

// StarSystem is the type used to store a complete star system
type StarSystem struct {
	*FSDJumpEvent
	Stars  []Star
	Bodies []Body
}

type Star struct {
	BodyName              string   `json:"BodyName,omitempty"`
	BodyID                uint64   `json:"BodyID,omitempty"`
	Parents               []Parent `json:"Parents,omitempty"`
	StarSystem            string   `json:"StarSystem,omitempty"`
	SystemAddress         uint64   `json:"SystemAddress,omitempty"`
	DistanceFromArrivalLS float64  `json:"DistanceFromArrivalLS,omitempty"`
	StarType              string   `json:"StarType,omitempty"`
	Subclass              uint     `json:"Subclass,omitempty"`
	StellarMass           float64  `json:"StellarMass,omitempty"`
	Radius                float64  `json:"Radius,omitempty"`
	AbsoluteMagnitude     float64  `json:"AbsoluteMagnitude,omitempty"`
	AgeMY                 float64  `json:"Age_MY,omitempty"`
	SurfaceTemperature    float64  `json:"SurfaceTemperature,omitempty"`
	Luminosity            string   `json:"Luminosity,omitempty"`
	SemiMajorAxis         float64  `json:"SemiMajorAxis,omitempty"`
	Eccentricity          float64  `json:"Eccentricity,omitempty"`
	OrbitalInclination    float64  `json:"OrbitalInclination,omitempty"`
	Periapsis             float64  `json:"Periapsis,omitempty"`
	OrbitalPeriod         float64  `json:"OrbitalPeriod,omitempty"`
	RotationPeriod        float64  `json:"RotationPeriod,omitempty"`
	AxisTilt              float64  `json:"AxialTilt,omitempty"`
	Rings                 []struct {
		Name      string  `json:"Name"`
		RingClass string  `json:"RingClass"`
		MassMT    float64 `json:"MassMT"`
		InnerRad  float64 `json:"InnerRad"`
		OuterRad  float64 `json:"OuterRad"`
	}
	WasDiscovered bool `json:"WasDiscovered,omitempty"`
	WasMapped     bool `json:"WasMapped,omitempty"`
}

type Body struct {
	BodyName              string   `json:"BodyName,omitempty"`
	BodyID                uint64   `json:"BodyID,omitempty"`
	Parents               []Parent `json:"Parents,omitempty"`
	StarSystem            string   `json:"StarSystem,omitempty"`
	SystemAddress         uint64   `json:"SystemAddress,omitempty"`
	DistanceFromArrivalLS float64  `json:"DistanceFromArrivalLS,omitempty"`
	TidalLock             bool     `json:"TidalLock"`
	TerraformState        string   `json:"TerraformState"`
	PlanetClass           string   `json:"PlanetClass"`
	Atmosphere            string   `json:"Atmosphere"`
	AtmosphereType        string   `json:"AtmosphereType"`
	AtmosphereComposition []struct {
		Name    string  `json:"Name"`
		Percent float64 `json:"Percent"`
	}
	Volcanism          string  `json:"Volcanism"`
	MassEM             float64 `json:"MassEM"`
	Radius             float64 `json:"Radius"`
	SurfaceGravity     float64 `json:"SurfaceGravity"`
	SurfaceTemperature float64 `json:"SurfaceTemperature"`
	SurfacePressure    float64 `json:"SurfacePressure"`
	Landable           bool    `json:"Landable"`
	Materials          []struct {
		Name    string  `json:"Name"`
		Percent float64 `json:"Percent"`
	}
	Composition struct {
		Ice   float64 `json:"Ice"`
		Rock  float64 `json:"Rock"`
		Metal float64 `json:"Metal"`
	}
	SemiMajorAxis      float64 `json:"SemiMajorAxis"`
	Eccentricity       float64 `json:"Eccentricity"`
	OrbitalInclination float64 `json:"OrbitalInclination"`
	Periapsis          float64 `json:"Periapsis"`
	OrbitalPeriod      float64 `json:"OrbitalPeriod"`
	RotationPeriod     float64 `json:"RotationPeriod"`
	AxialTilt          float64 `json:"AxialTilt"`
	Rings              []struct {
		Name      string  `json:"Name"`
		RingClass string  `json:"RingClass"`
		MassMT    float64 `json:"MassMT"`
		InnerRad  float64 `json:"InnerRad"`
		OuterRad  float64 `json:"OuterRad"`
	}
	WasDiscovered bool `json:"WasDiscovered"`
	WasMapped     bool `json:"WasMapped"`
}

type Moon struct {
	BodyName              string   `json:"BodyName,omitempty"`
	BodyID                uint64   `json:"BodyID,omitempty"`
	Parents               []Parent `json:"Parents,omitempty"`
	StarSystem            string   `json:"StarSystem,omitempty"`
	SystemAddress         uint64   `json:"SystemAddress,omitempty"`
	DistanceFromArrivalLS float64  `json:"DistanceFromArrivalLS,omitempty"`
	TidalLock             bool     `json:"TidalLock"`
	TerraformState        string   `json:"TerraformState"`
	PlanetClass           string   `json:"PlanetClass"`
	Atmosphere            string   `json:"Atmosphere"`
	AtmosphereType        string   `json:"AtmosphereType"`
	AtmosphereComposition []struct {
		Name    string  `json:"Name"`
		Percent float64 `json:"Percent"`
	}
	Volcanism          string  `json:"Volcanism"`
	MassEM             float64 `json:"MassEM"`
	Radius             float64 `json:"Radius"`
	SurfaceGravity     float64 `json:"SurfaceGravity"`
	SurfaceTemperature float64 `json:"SurfaceTemperature"`
	SurfacePressure    float64 `json:"SurfacePressure"`
	Landable           bool    `json:"Landable"`
	Materials          []struct {
		Name    string  `json:"Name"`
		Percent float64 `json:"Percent"`
	}
	Composition struct {
		Ice   float64 `json:"Ice"`
		Rock  float64 `json:"Rock"`
		Metal float64 `json:"Metal"`
	}
	SemiMajorAxis      float64 `json:"SemiMajorAxis"`
	Eccentricity       float64 `json:"Eccentricity"`
	OrbitalInclination float64 `json:"OrbitalInclination"`
	Periapsis          float64 `json:"Periapsis"`
	OrbitalPeriod      float64 `json:"OrbitalPeriod"`
	RotationPeriod     float64 `json:"RotationPeriod"`
	AxialTilt          float64 `json:"AxialTilt"`
	Rings              []struct {
		Name      string  `json:"Name"`
		RingClass string  `json:"RingClass"`
		MassMT    float64 `json:"MassMT"`
		InnerRad  float64 `json:"InnerRad"`
		OuterRad  float64 `json:"OuterRad"`
	}
	WasDiscovered bool `json:"WasDiscovered"`
	WasMapped     bool `json:"WasMapped"`
}

type Parent struct {
	Planet int `json:"Planet,omitempty"`
	Star   int `json:"Star,omitempty"`
	Null   int `json:"Null,omitempty"`
	Ring   int `json:"Ring,omitempty"`
}

// struct methods for interface compliance
func (bse *BaseScanEvent) EventType() string {
	return "BaseScanEvent"
}

func (fsd *FSDJumpEvent) EventType() string {
	return "FSDJumpEvent"
}

func (ase *AutoScanEvent) EventType() string {
	return "AutoScanEvent"
}

func (dse *DetailedScanEvent) EventType() string {
	return "DetailedScanEvent"
}

func (bse *BaseScanEvent) SkipEvent() bool {
	if bse.ScanType != "Detailed" && bse.ScanType != "AutoScan" {
		return false
	}
	return true
}

func (s *StarSystem) AddStar(star Star) {
	s.Stars = append(s.Stars, star)
}

func (s *StarSystem) AddBody(body Body) {
	s.Bodies = append(s.Bodies, body)
}
