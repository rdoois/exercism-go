package resistorcolor

// Colors should return the list of all colors.
func Colors() []string {
	return []string{
		"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for k, v := range Colors() {
		if v == color {
			return k
		}
	}
	return 0
}