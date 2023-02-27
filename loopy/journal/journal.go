package journal

import (
	jsoniter "github.com/json-iterator/go"
	"os"
	"path/filepath"
	"strings"
)

/*
Defines all the global variables
*/
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// AllStarSystems is a slice of all the star systems that have been scanned
var AllStarSystems []*StarSystem

// currentStarSystem is a pointer to the current star system being scanned
var currentStarSystem *StarSystem

/*
Custom struct methods
*/

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

			if star.StarType != "" {
				star.BaseScanEvent = baseScanEvent
				return star, nil
			}
			return baseScanEvent, nil

		//DetailedScan event is used to get the planets in the system
		case "Detailed":
			var planet Planet
			err := json.Unmarshal(data, &planet)
			if err != nil {
				return baseScanEvent, err
			}
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

		if fsdJumpEvent, ok := journalEntry.(FSDJumpEvent); ok {
			//set current star system
			currentStarSystem = SetCurrentStarSystem(fsdJumpEvent)
			continue
		}
		if star, ok := journalEntry.(Star); ok {
			//add star to current star system
			AddStarToStarSystem(&star)
			continue
		}
		if planet, ok := journalEntry.(Planet); ok {
			//add planet to current star system
			AddBodyToStarSystem(&planet)
			continue
		}
	}
	return nil
}

// SetCurrentStarSystem checks if the star system is already in the list of star systems and returns it if it is
// If it is not in the list, it adds it to the list and returns it
func SetCurrentStarSystem(entry FSDJumpEvent) *StarSystem {

	for index := range AllStarSystems {
		if AllStarSystems[index].FSDJumpEvent.StarSystem == entry.StarSystem {

			return AllStarSystems[index]
		}
	}

	starSystem := new(StarSystem)
	starSystem.FSDJumpEvent = entry
	AllStarSystems = append(AllStarSystems, starSystem)
	return AllStarSystems[len(AllStarSystems)-1]
}

// AddBodyToStarSystem adds a body to a star system
func AddBodyToStarSystem(planet *Planet) {
	currentStarSystem.Bodies = append(currentStarSystem.Bodies, *planet)
}

// AddStarToStarSystem adds a star to a star system
func AddStarToStarSystem(star *Star) {
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
