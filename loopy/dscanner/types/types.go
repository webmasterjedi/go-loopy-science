package types

import (
	jsoniter "github.com/json-iterator/go"
	"log"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

/*
Define types based on Elite Dangerous Player Journal API documentation here: https://elite-journal.readthedocs.io/en/latest/Exploration/#scan
*/

type Event interface {
	EventType() string
	// add additional methods here as needed
}

type SystemBody interface {
	ParentsToJson() string
	MaterialsToJson() string
	AtmosphereCompositionToJson() string
	BodyCompositionToJson() string
	RingsToJson() string
	ToJson() string
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
	Rings              []Ring  `json:"Rings"`
	WasDiscovered      bool    `json:"WasDiscovered"`
	WasMapped          bool    `json:"WasMapped"`
}

// StarSystem is the type used to store a complete star system
type StarSystem struct {
	*FSDJumpEvent
	Stars  []*Star
	Bodies []*Body
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
	AxialTilt             float64  `json:"AxialTilt,omitempty"`
	Rings                 []Ring   `json:"Rings,omitempty"`
	WasDiscovered         bool     `json:"WasDiscovered,omitempty"`
	WasMapped             bool     `json:"WasMapped,omitempty"`
}

type Body struct {
	BodyName              string    `json:"BodyName,omitempty"`
	BodyID                uint64    `json:"BodyID,omitempty"`
	Parents               []Parent  `json:"Parents,omitempty"`
	StarSystem            string    `json:"StarSystem,omitempty"`
	SystemAddress         uint64    `json:"SystemAddress,omitempty"`
	DistanceFromArrivalLS float64   `json:"DistanceFromArrivalLS,omitempty"`
	TidalLock             bool      `json:"TidalLock,omitempty"`
	TerraformState        string    `json:"TerraformState,omitempty"`
	PlanetClass           string    `json:"PlanetClass,omitempty"`
	Atmosphere            string    `json:"Atmosphere,omitempty"`
	AtmosphereType        string    `json:"AtmosphereType,omitempty"`
	AtmosphereComposition []Percent `json:"AtmosphereComposition,omitempty"`
	Volcanism             string    `json:"Volcanism,omitempty"`
	MassEM                float64   `json:"MassEM,omitempty"`
	Radius                float64   `json:"Radius,omitempty"`
	SurfaceGravity        float64   `json:"SurfaceGravity,omitempty"`
	SurfaceTemperature    float64   `json:"SurfaceTemperature,omitempty"`
	SurfacePressure       float64   `json:"SurfacePressure,omitempty"`
	Landable              bool      `json:"Landable,omitempty"`
	Materials             []Percent `json:"Materials,omitempty"`
	BodyComposition       `json:"BodyComposition,omitempty"`
	SemiMajorAxis         float64 `json:"SemiMajorAxis,omitempty"`
	Eccentricity          float64 `json:"Eccentricity,omitempty"`
	OrbitalInclination    float64 `json:"OrbitalInclination,omitempty"`
	Periapsis             float64 `json:"Periapsis,omitempty"`
	OrbitalPeriod         float64 `json:"OrbitalPeriod,omitempty"`
	RotationPeriod        float64 `json:"RotationPeriod,omitempty"`
	AxialTilt             float64 `json:"AxialTilt,omitempty"`
	Rings                 []Ring  `json:"Rings,omitempty"`
	WasDiscovered         bool    `json:"WasDiscovered,omitempty"`
	WasMapped             bool    `json:"WasMapped,omitempty"`
}

type Parent struct {
	Planet int `json:"Planet,omitempty"`
	Star   int `json:"Star,omitempty"`
	Null   int `json:"Null,omitempty"`
	Ring   int `json:"Ring,omitempty"`
}

type Ring struct {
	Name      string  `json:"Name"`
	RingClass string  `json:"RingClass"`
	MassMT    float64 `json:"MassMT"`
	InnerRad  float64 `json:"InnerRad"`
	OuterRad  float64 `json:"OuterRad"`
}
type Percent struct {
	Name    string  `json:"Name"`
	Percent float64 `json:"Percent"`
}
type BodyComposition struct {
	Ice   float64 `json:"Ice"`
	Rock  float64 `json:"Rock"`
	Metal float64 `json:"Metal"`
}

// EventType struct methods for interface compliance
// returns the event type
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

func (s *StarSystem) AddStar(star *Star) {
	s.Stars = append(s.Stars, star)
}

func (s *StarSystem) AddBody(body *Body) {
	s.Bodies = append(s.Bodies, body)
}

func (s *Star) ParentsToJson() string {
	parents, err := json.Marshal(s.Parents)
	if err != nil {
		log.Println(err)
	}
	return string(parents)
}

func (s *Star) RingsToJson() string {
	rings, err := json.Marshal(s.Rings)
	if err != nil {
		log.Println(err)
	}
	return string(rings)
}

func (b *Body) ParentsToJson() string {
	parents, err := json.Marshal(b.Parents)
	if err != nil {
		log.Println(err)
	}
	return string(parents)
}

func (b *Body) RingsToJson() string {
	rings, err := json.Marshal(b.Rings)
	if err != nil {
		log.Println(err)
	}
	return string(rings)
}

func (b *Body) MaterialsToJson() string {
	materials, err := json.Marshal(b.Materials)
	if err != nil {
		log.Println(err)
	}
	return string(materials)
}

func (b *Body) AtmosphereCompositionToJson() string {
	atmosphereComposition, err := json.Marshal(b.AtmosphereComposition)
	if err != nil {
		log.Println(err)
	}
	return string(atmosphereComposition)
}

func (b *Body) BodyCompositionToJson() string {
	bodyComposition, err := json.Marshal(b.BodyComposition)
	if err != nil {
		log.Println(err)
	}
	return string(bodyComposition)
}

func (b *Body) ToJson() string {
	json, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	}
	return string(json)
}

// UnmarshalJSON is a custom unmarshaler for the Parent struct
func (p *Parent) UnmarshalJSON(data []byte) error {
	// Define a struct to unmarshal the JSON into without ignoring null values
	var tmp struct {
		Planet *int `json:"Planet,omitempty"`
		Star   *int `json:"Star,omitempty"`
		Null   *int `json:"Null,omitempty"`
		Ring   *int `json:"Ring,omitempty"`
	}

	// Unmarshal the JSON into the temporary struct
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	// If a key is present in the JSON object and has a non-null value, set the corresponding field in the target struct
	if tmp.Null != nil {
		p.Null = *tmp.Null
	} else {
		p.Null = -1
	}
	if tmp.Planet != nil {
		p.Planet = *tmp.Planet
	} else {
		p.Planet = -1
	}
	if tmp.Star != nil {
		p.Star = *tmp.Star
	} else {
		p.Star = -1
	}
	if tmp.Ring != nil {
		p.Ring = *tmp.Ring
	} else {
		p.Ring = -1
	}

	return nil
}
