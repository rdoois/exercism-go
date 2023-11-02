package triangle

type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "Nat" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

func KindFromSides(a, b, c float64) Kind {
	if a == 0 || b == 0 || c == 0 {
		return NaT
	}

	if a+b < c || b+c < a || a+c < b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	}

	return Sca

}
