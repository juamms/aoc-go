package days

import (
	"strings"

	"github.com/juamms/aoc-go/2019/utils"
)

type Day6 struct{}

func buildOrbits() map[string]string {
	orbits := make(map[string]string)

	utils.ScanInputFile(6, func(data string) {
		planets := strings.Split(data, ")")

		orbits[planets[0]] = planets[1]
	})

	return orbits
}

func countOrbit(orbits map[string]string, planet string) int {
	for k, v := range orbits {
		if v == planet {
			return countOrbit(orbits, k)
		}
	}

	return 1
}

func (day Day6) Part1() interface{} {
	// orbits := buildOrbits()

	orbits := make(map[string]string)
	raw := "COM)B,B)C,C)D,D)E,E)F,B)G,G)H,D)I,E)J,J)K,K)L"

	for _, p := range strings.Split(raw, ",") {
		planets := strings.Split(p, ")")

		orbits[planets[0]] = planets[1]
	}

	orbitNumber := 0

	for planet := range orbits {
		orbitNumber += countOrbit(orbits, planet)
	}

	return orbitNumber
}

func (day Day6) Part2() interface{} {
	return 0
}
