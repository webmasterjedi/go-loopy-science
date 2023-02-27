package journal

import (
	"os"
	"path/filepath"
	"strings"

	jsoniter "github.com/json-iterator/go"
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
func ParseJournalEvent(data []byte) (Event, error) {
	var baseScanEvent BaseScanEvent

	err := json.Unmarshal(data, &baseScanEvent)

	if err != nil {
		return &baseScanEvent, err
	}

	//switch to handle main event types
	switch baseScanEvent.Event {
	//FSDJump event is used to get current star system
	case "FSDJump":
		return &baseScanEvent, nil
	//Scan event is used to get the scan type
	case "Scan":
		switch baseScanEvent.ScanType {
		//AutoScan event is used to get the stars in the system
		case "AutoScan":
			var autoScan AutoScanEvent
			err := json.Unmarshal(data, &autoScan)
			if err != nil {
				return &baseScanEvent, err
			}

			if autoScan.StarType != "" {
				autoScan.BaseScanEvent = baseScanEvent
				return &autoScan, nil
			}
			return &baseScanEvent, nil

		//DetailedScan event is used to get the planets in the system
		case "Detailed":
			var detailed DetailedScanEvent
			err := json.Unmarshal(data, &detailed)
			if err != nil {
				return &baseScanEvent, err
			}
			if detailed.PlanetClass != "" {
				detailed.BaseScanEvent = baseScanEvent
				return &detailed, nil
			}
		}
	}

	return &baseScanEvent, nil
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
		event, err := ParseJournalEvent([]byte(entry))
		if err != nil {
			return err
		}

		//use the event type to call the appropriate function
		switch event.EventType() {
		case "FSDJumpEvent":
			currentStarSystem = SetCurrentStarSystem(event.(*FSDJumpEvent))
		case "AutoScanEvent":
			AddStarToStarSystem(event.(*AutoScanEvent))
		case "DetailedScanEvent":
			AddBodyToStarSystem(event.(*DetailedScanEvent))
		}
	}
	return nil
}

// SetCurrentStarSystem checks if the star system is already in the list of star systems and returns it if it is
// If it is not in the list, it adds it to the list and returns it
func SetCurrentStarSystem(entry *FSDJumpEvent) *StarSystem {

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

// AddStarToStarSystem adds a star to a star system
func AddStarToStarSystem(autoScan *AutoScanEvent) error {
	//convert autoScan struct to Star struct

	//copy all fields from autoScan to star
	autoScanJSON, err := json.Marshal(autoScan)
	if err != nil {
		return err
	}
	var star Star
	err = json.Unmarshal(autoScanJSON, &star)
	if err != nil {
		return err
	}

	currentStarSystem.Stars = append(currentStarSystem.Stars, star)
	return nil
}

// AddBodyToStarSystem adds a body to a star system
func AddBodyToStarSystem(detailed *DetailedScanEvent) error {
	//convert detailed struct to Body struct
	detailedJSON, err := json.Marshal(detailed)
	if err != nil {
		return err
	}
	var body Body
	err = json.Unmarshal(detailedJSON, &body)
	if err != nil {
		return err
	}
	currentStarSystem.Bodies = append(currentStarSystem.Bodies, body)
	return nil
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
