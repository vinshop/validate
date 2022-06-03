package validate

// NumberEpsilon we need this because sometime compare 2 float64 could be stupid
var NumberEpsilon = 1e-6

func SetNumberEpsilon(eps float64) {
	NumberEpsilon = eps
}
