package totals

import (
	"fmt"
	"goloopyscience/loopy/dscanner/types"
)

func Build(systems *[]*types.StarSystem, totalsMap map[string]*types.BodyCounts) map[string]*types.BodyCounts {
	for _, system := range *systems {
		//fmt.Println(system.SystemName)
		totalsMap = processStars(*system, totalsMap)

	}
	return totalsMap
}

func processStars(system types.StarSystem, totalsMap map[string]*types.BodyCounts) map[string]*types.BodyCounts {
	for _, star := range system.Stars {

		//check if star type exists in map, if not add it
		if _, ok := totalsMap[star.GetFullStarType()]; !ok {
			totalsMap[star.GetFullStarType()] = &types.BodyCounts{
				WaterWorlds:     0,
				EarthlikeBodies: 0,
				AmmoniaWorlds:   0,
			}
		}

		totalsMap = processBodies(&system.Bodies, *star, totalsMap)
	}
	return totalsMap
}

func processBodies(bodies *[]*types.Body, star types.Star, totalsMap map[string]*types.BodyCounts) map[string]*types.BodyCounts {
	for _, body := range *bodies {
		if body.Counted {
			continue
		}
		if body.IsValuable() && body.ParentID == star.ParentID {
			fmt.Println(body.BodyName, body.PlanetClass, body.ParentID, star.ParentID)
			switch body.PlanetClass {
			case "Water world":
				totalsMap[star.GetFullStarType()].WaterWorlds++
			case "Earthlike body":
				totalsMap[star.GetFullStarType()].EarthlikeBodies++
			case "Ammonia world":
				totalsMap[star.GetFullStarType()].AmmoniaWorlds++
			}
			body.Counted = true
		}

	}
	return totalsMap
}
