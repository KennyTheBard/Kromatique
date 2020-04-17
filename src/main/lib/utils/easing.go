package utils

type EasingFunction func(t float64) float64

// ease in functions
func EaseInQuad(t float64) float64 {
	return t * t
}

func EaseInCubic(t float64) float64 {
	return t * EaseInQuad(t)
}

func EaseInQuart(t float64) float64 {
	return t * EaseInCubic(t)
}

// ease out functions
func EaseOutQuad(t float64) float64 {
	return 1 - EaseInQuad(1-t)
}

func EaseOutCubic(t float64) float64 {
	return 1 - EaseInCubic(1-t)
}

func EaseOutQuart(t float64) float64 {
	return 1 - EaseInQuart(1-t)
}

// ease in out functions
func EaseInOutQuad(t float64) float64 {
	if t < 0.5 {
		return EaseInQuad(2*t) / 2
	} else {
		return EaseOutQuad(2*(t-0.5)) / 2
	}
}

func EaseInOutCubic(t float64) float64 {
	if t < 0.5 {
		return EaseInCubic(2*t) / 2
	} else {
		return EaseOutCubic(2*(t-0.5)) / 2
	}
}

func EaseInOutQuart(t float64) float64 {
	if t < 0.5 {
		return EaseInQuart(2*t) / 2
	} else {
		return EaseOutQuart(2*(t-0.5)) / 2
	}
}
