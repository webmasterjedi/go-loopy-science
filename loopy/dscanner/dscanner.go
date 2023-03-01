package dscanner

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"goloopyscience/loopy/db"
	"goloopyscience/loopy/dscanner/types"
	"os"
	"path/filepath"
	"strings"
)

/*
Defines all the global variables
*/
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// allSystems is a slice of all the star systems that have been scanned
var allSystems []*types.StarSystem

// currentSystem is a pointer to the current star system being scanned
var currentSystem *types.StarSystem

/*
Main package methods
*/

// Honk processes all journal files in a directory
func Honk(dir string) ([]*types.StarSystem, error) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".log" {
			err := processFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
		}
	}

	return allSystems, nil
}

// parseEvent parses the journal event and returns the proper struct based on the event types
func parseEvent(data []byte) (types.Event, error) {
	var baseScanEvent types.BaseScanEvent

	err := json.Unmarshal(data, &baseScanEvent)

	if err != nil {
		return &baseScanEvent, err
	}

	//switch to handle main event types
	switch baseScanEvent.Event {
	//FSDJump event is used to get current star system
	case "FSDJump":
		var fsd types.FSDJumpEvent
		err := json.Unmarshal(data, &fsd)
		if err != nil {
			return &baseScanEvent, err
		}

		if fsd.SystemAddress != 0 {
			fsd.BaseScanEvent = baseScanEvent
			return &fsd, nil
		}
		return &baseScanEvent, nil
	//Scan event is used to get the scan type
	case "Scan":
		switch baseScanEvent.ScanType {
		//AutoScan event is used to get the stars in the system
		case "AutoScan":
			var autoScan types.AutoScanEvent
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
			var detailed types.DetailedScanEvent
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

// processFile reads the journal file and parses the events
func processFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	entries := strings.Split(string(data), "\n")
	for _, entry := range entries {
		if entry == "" {
			continue
		}
		//parse journal entry and returns struct based on event types
		event, err := parseEvent([]byte(entry))
		if err != nil {
			return err
		}
		//if the event is an FSDJump event, set the current star system
		if event.EventType() == "FSDJumpEvent" {
			currentSystem = setStarSystem(event.(*types.FSDJumpEvent))
		}
		if currentSystem != nil {
			//use the event type to call the appropriate function
			switch event.EventType() {
			case "AutoScanEvent":
				err := addStar(event.(*types.AutoScanEvent))
				if err != nil {
					return err
				}
			case "DetailedScanEvent":
				err := addBody(event.(*types.DetailedScanEvent))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// setStarSystem checks if the star system is already in the list of star systems and returns it if it is
// If it is not in the list, it adds it to the list and returns it
func setStarSystem(entry *types.FSDJumpEvent) *types.StarSystem {
	for index := range allSystems {
		if allSystems[index].FSDJumpEvent.StarSystem == entry.StarSystem {
			return allSystems[index]
		}
	}
	starSystem := new(types.StarSystem)
	starSystem.FSDJumpEvent = entry
	//save to db
	err := db.InsertSystem(starSystem)
	if err != nil {
		fmt.Print(err)
	}
	allSystems = append(allSystems, starSystem)
	return allSystems[len(allSystems)-1]
}

// addStar adds a star to a star system
func addStar(autoScan *types.AutoScanEvent) error {
	//convert autoScan struct to Star struct
	autoScanJSON, err := json.Marshal(autoScan)
	if err != nil {
		return err
	}
	var star types.Star
	err = json.Unmarshal(autoScanJSON, &star)
	if err != nil {
		return err
	}
	//save to db
	err = db.InsertStar(&star)
	if err != nil {
		fmt.Print(err)
	}
	currentSystem.AddStar(&star)
	return nil
}

// addBody adds a body to a star system
func addBody(detailed *types.DetailedScanEvent) error {
	//convert detailed struct to Body struct
	detailedJSON, err := json.Marshal(detailed)
	if err != nil {
		return err
	}
	var body types.Body
	err = json.Unmarshal(detailedJSON, &body)
	if err != nil {
		return err
	}
	//save to db
	err = db.InsertBody(&body)
	if err != nil {
		fmt.Print(err)
	}
	currentSystem.AddBody(&body)
	return nil
}
