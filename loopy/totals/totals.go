package totals

import (
	"fmt"
	"goloopyscience/loopy/dscanner/types"
)

func Build(systems *[]*types.StarSystem, totalsMap map[string]*types.BodyCounts) map[string]*types.BodyCounts {
	for _, system := range *systems {
		for _, star := range system.Stars {
			var fullStarType = fmt.Sprintf("%s%d %s", star.StarType, star.Subclass, star.Luminosity)

			//check if star type exists in map, if not add it
			if _, ok := totalsMap[fullStarType]; !ok {
				totalsMap[fullStarType] = &types.BodyCounts{
					WaterWorlds:     0,
					EarthLikeWorlds: 0,
					AmmoniaWorlds:   0,
				}
			}

			/*for _, body := range system.Bodies {
				fmt.Println(body.PlanetClass)
			}*/
		}
	}
	return totalsMap
}
