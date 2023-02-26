package EDJournalTools

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

/*
Build types based on Elite Dangerous Player Journal API documentation here: https://elite-journal.readthedocs.io/en/latest/Exploration/#scan
*/

type BaseScanEvent struct {
	Event                 string   `json:"event"`
	ScanType              string   `json:"ScanType"`
	BodyName              string   `json:"BodyName"`
	BodyId                int      `json:"BodyId"`
	Parents               []Parent `json:"Parents"`
	StarSystem            string   `json:"StarSystem"`
	SystemAddress         int64    `json:"SystemAddress"`
	DistanceFromArrivalLS float64  `json:"DistanceFromArrivalLS"`
}
type Star struct {
	BaseScanEvent
	StarType           string  `json:"StarType"`
	Subclass           int     `json:"Subclass"`
	StellarMass        float64 `json:"StellarMass"`
	Radius             float64 `json:"Radius"`
	AbsoluteMagnitude  float64 `json:"AbsoluteMagnitude"`
	AgeMY              float64 `json:"Age_MY"`
	SurfaceTemperature float64 `json:"SurfaceTemperature"`
	Luminosity         string  `json:"Luminosity"`
	SemiMajorAxis      float64 `json:"SemiMajorAxis"`
	Eccentricity       float64 `json:"Eccentricity"`
	OrbitalInclination float64 `json:"OrbitalInclination"`
	Periapsis          float64 `json:"Periapsis"`
	OrbitalPeriod      float64 `json:"OrbitalPeriod"`
	RotationPeriod     float64 `json:"RotationPeriod"`
	AxisTilt           float64 `json:"AxialTilt"`
	WasDiscovered      bool    `json:"WasDiscovered"`
	WasMapped          bool    `json:"WasMapped"`
}

type Planet struct {
	BaseScanEvent
	PlanetClass string
}

type Moon struct {
	BaseScanEvent
	Name        string
	PlanetClass string
}

type StarSystem struct {
	Name   string
	Stars  []Star
	Bodies []Planet
}

type Parent struct {
	Planet int `json:"Planet,omitempty"`
	Star   int `json:"Star,omitempty"`
	Null   int `json:"Null,omitempty"`
	Ring   int `json:"Ring,omitempty"`
}

/*
Defines all the global variables
*/

// AllStarSystems is a slice of all the star systems that have been scanned
var AllStarSystems []*StarSystem

// currentStarSystem is a pointer to the current star system being scanned
var currentStarSystem *StarSystem

/*
Custom struct methods
*/

// UnmarshalJSON is a custom unmarshaler for the BaseScanEvent struct
func (b *BaseScanEvent) UnmarshalJSON(data []byte) error {
	// Define a struct to unmarshal the JSON into without ignoring null values
	var tmp struct {
		Event                 string   `json:"event"`
		ScanType              string   `json:"ScanType"`
		BodyName              string   `json:"BodyName"`
		BodyId                int      `json:"BodyId"`
		Parents               []Parent `json:"Parents"`
		StarSystem            string   `json:"StarSystem"`
		SystemAddress         int64    `json:"SystemAddress"`
		DistanceFromArrivalLS float64  `json:"DistanceFromArrivalLS"`
	}

	// Unmarshal the JSON into the temporary struct
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	// If a key is present in the JSON object and has a non-null value, set the corresponding field in the target struct
	if tmp.Event != "" {
		b.Event = tmp.Event
	}
	if tmp.ScanType != "" {
		b.ScanType = tmp.ScanType
	}
	if tmp.BodyName != "" {
		b.BodyName = tmp.BodyName
	}
	if tmp.BodyId != 0 {
		b.BodyId = tmp.BodyId
	}
	if tmp.Parents != nil {
		b.Parents = tmp.Parents
	}
	if tmp.StarSystem != "" {
		b.StarSystem = tmp.StarSystem
	}
	if tmp.SystemAddress != 0 {
		b.SystemAddress = tmp.SystemAddress
	}
	if tmp.DistanceFromArrivalLS != 0 {
		b.DistanceFromArrivalLS = tmp.DistanceFromArrivalLS
	}

	return nil
}

// UnmarshalJSON is a custom unmarshaler for the Parent struct
func (p *Parent) UnmarshalJSON(data []byte) error {
	// Define a struct to unmarshal the JSON into without ignoring null values
	var tmp struct {
		Planet *int `json:"Planet"`
		Star   *int `json:"Star"`
		Null   *int `json:"Null"`
		Ring   *int `json:"Ring"`
	}

	// Unmarshal the JSON into the temporary struct
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	// If a key is present in the JSON object and has a non-null value, set the corresponding field in the target struct
	if tmp.Null != nil {
		p.Null = *tmp.Null
	}
	if tmp.Planet != nil {
		p.Planet = *tmp.Planet
	}
	if tmp.Star != nil {
		p.Star = *tmp.Star
	}
	if tmp.Ring != nil {
		p.Ring = *tmp.Ring
	}

	return nil
}

// UnmarshalJSON is a custom unmarshaler for the Star struct
func (s *Star) UnmarshalJSON(data []byte) error {
	// Define a struct to unmarshal the JSON into without ignoring null values
	var tmp struct {
		BaseScanEvent
		StarType           string  `json:"StarType"`
		Subclass           int     `json:"Subclass"`
		StellarMass        float64 `json:"StellarMass"`
		Radius             float64 `json:"Radius"`
		AbsoluteMagnitude  float64 `json:"AbsoluteMagnitude"`
		AgeMY              float64 `json:"Age_MY"`
		SurfaceTemperature float64 `json:"SurfaceTemperature"`
		Luminosity         string  `json:"Luminosity"`
		SemiMajorAxis      float64 `json:"SemiMajorAxis"`
		Eccentricity       float64 `json:"Eccentricity"`
		OrbitalInclination float64 `json:"OrbitalInclination"`
		Periapsis          float64 `json:"Periapsis"`
		OrbitalPeriod      float64 `json:"OrbitalPeriod"`
		RotationPeriod     float64 `json:"RotationPeriod"`
		AxisTilt           float64 `json:"AxialTilt"`
		WasDiscovered      bool    `json:"WasDiscovered"`
		WasMapped          bool    `json:"WasMapped"`
	}

	// Unmarshal the JSON into the temporary struct
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	// If a key is present in the JSON object and has a non-null value, set the corresponding field in the target struct
	if tmp.StarType != "" {
		s.StarType = tmp.StarType
	}
	if tmp.Subclass != 0 {
		s.Subclass = tmp.Subclass
	}
	if tmp.StellarMass != 0 {
		s.StellarMass = tmp.StellarMass
	}
	if tmp.Radius != 0 {
		s.Radius = tmp.Radius
	}
	if tmp.AbsoluteMagnitude != 0 {
		s.AbsoluteMagnitude = tmp.AbsoluteMagnitude
	}
	if tmp.AgeMY != 0 {
		s.AgeMY = tmp.AgeMY
	}
	if tmp.SurfaceTemperature != 0 {
		s.SurfaceTemperature = tmp.SurfaceTemperature
	}
	if tmp.Luminosity != "" {
		s.Luminosity = tmp.Luminosity
	}
	if tmp.SemiMajorAxis != 0 {
		s.SemiMajorAxis = tmp.SemiMajorAxis
	}
	if tmp.Eccentricity != 0 {
		s.Eccentricity = tmp.Eccentricity
	}
	if tmp.OrbitalInclination != 0 {
		s.OrbitalInclination = tmp.OrbitalInclination
	}
	if tmp.Periapsis != 0 {
		s.Periapsis = tmp.Periapsis
	}
	if tmp.OrbitalPeriod != 0 {
		s.OrbitalPeriod = tmp.OrbitalPeriod
	}
	if tmp.RotationPeriod != 0 {
		s.RotationPeriod = tmp.RotationPeriod
	}
	if tmp.AxisTilt != 0 {
		s.AxisTilt = tmp.AxisTilt
	}
	if tmp.WasDiscovered {
		s.WasDiscovered = tmp.WasDiscovered
	}
	if tmp.WasMapped {
		s.WasMapped = tmp.WasMapped
	}
	return nil
}

/*
Main package methods
*/

// ParseJournalEvent parses the journal event and returns the proper struct based on the event types
func ParseJournalEvent(data []byte) (any, error) {
	var baseScanEvent BaseScanEvent

	err := json.Unmarshal(data, &baseScanEvent)

	if err != nil {
		return baseScanEvent, err
	}

	//switch to handle main event types
	switch baseScanEvent.Event {
	//FSDJump event is used to get current star system
	case "FSDJump":
		return baseScanEvent, nil
	//Scan event is used to get the scan type
	case "Scan":
		switch baseScanEvent.ScanType {
		//AutoScan event is used to get the stars in the system
		case "AutoScan":
			var star Star
			err := json.Unmarshal(data, &star)
			if err != nil {
				return baseScanEvent, err
			}
			//spew.Dump(star)
			if star.StarType != "" {
				fmt.Print("Star struct\n")
				star.BaseScanEvent = baseScanEvent
				return star, nil
			}
			return baseScanEvent, nil

		//DetailedScan event is used to get the planets in the system
		case "DetailedScan":
			var planet Planet
			err := json.Unmarshal(data, &planet)
			if err != nil {
				return baseScanEvent, err
			}
			fmt.Print(planet.PlanetClass + "\n")
			if planet.PlanetClass != "" {
				planet.BaseScanEvent = baseScanEvent
				return planet, nil
			}
		}
	}

	return baseScanEvent, nil
}

// ReadJournalFile reads the journal file and returns the data
func ReadJournalFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// ProcessJournalFile reads the journal file and parses the events
func ProcessJournalFile(path string) error {
	data, err := ReadJournalFile(path)
	if err != nil {
		return err
	}

	entries := strings.Split(string(data), "\n")
	for _, entry := range entries {
		if entry == "" {
			continue
		}
		//parse journal entry and returns struct based on event types
		journalEntry, err := ParseJournalEvent([]byte(entry))
		if err != nil {
			return err
		}

		if baseScanEvent, ok := journalEntry.(*BaseScanEvent); ok {
			//set current star system
			currentStarSystem = SetCurrentStarSystem(baseScanEvent)
			continue
		}
		if star, ok := journalEntry.(*Star); ok {
			//add star to current star system
			AddStarToStarSystem(star)
			continue
		}
		if planet, ok := journalEntry.(*Planet); ok {
			//add planet to current star system
			AddBodyToStarSystem(planet)
			continue
		}
	}
	return nil
}

// SetCurrentStarSystem checks if the star system is already in the list of star systems and returns it if it is
// If it is not in the list, it adds it to the list and returns it
func SetCurrentStarSystem(entry *BaseScanEvent) *StarSystem {

	for index := range AllStarSystems {
		if AllStarSystems[index].Name == entry.StarSystem {

			return AllStarSystems[index]
		}
	}

	starSystem := new(StarSystem)
	starSystem.Name = entry.StarSystem

	AllStarSystems = append(AllStarSystems, starSystem)

	return AllStarSystems[len(AllStarSystems)-1]

}

// AddBodyToStarSystem adds a body to a star system
func AddBodyToStarSystem(planet *Planet) {
	fmt.Print("Adding planet to star system:\n")
	spew.Dump(planet)
	currentStarSystem.Bodies = append(currentStarSystem.Bodies, *planet)
}

// AddStarToStarSystem adds a star to a star system
func AddStarToStarSystem(star *Star) {
	fmt.Print("Adding star to star system:\n")
	spew.Dump(star)
	currentStarSystem.Stars = append(currentStarSystem.Stars, *star)
}

// ProcessJournalDirectory processes all journal files in a directory
func ProcessJournalDirectory(dir string) error {

	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".log" {
			err := ProcessJournalFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
