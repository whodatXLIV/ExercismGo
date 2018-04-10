// Package space provides a way to calculate how old someone is on different planets
package space

type Planet string

// Age returns a person's age in years on a given planet (p) given how many seconds (t) they have been alive
func Age(t float64, p Planet) float64 {
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

	return age // If Planet is Earth

}
