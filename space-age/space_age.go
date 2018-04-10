// Package space_age provides a way to calculate how old someone is on different planets
package space

// Age returns a person's age in years on a given planet (p) given how many seconds (t) they have been alive

func Age(t float64, p string) float64 {
	/*- Earth: orbital period 365.25 Earth days, or 31557600 seconds
	  - Mercury: orbital period 0.2408467 Earth years
	  - Venus: orbital period 0.61519726 Earth years
	  - Mars: orbital period 1.8808158 Earth years
	  - Jupiter: orbital period 11.862615 Earth years
	  - Saturn: orbital period 29.447498 Earth years
	  - Uranus: orbital period 84.016846 Earth years
	  - Neptune: orbital period 164.79132 Earth years*/
	age := t / 31557600
	switch p {
	case "Mercury":
		return age / 0.2408467
	case "Venus":
		return age / 0.61519726
	case "Mars":
		return age / 1.8808158
	case "Jupiter":
		return age / 11.862615
	case "Saturn":
		return age / 29.447498
	case "Uranus":
		return age / 84.016846
	case "Neptune":
		return age / 164.79132
	}

	return age

}
