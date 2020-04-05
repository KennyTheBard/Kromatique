package utils

type EasingFunction func(t float64) float64

func Linear(t float64) float64 {
	return t
}

func EaseInQuad(t float64) float64 {
	return t * t
}

func EaseOutQuad(t float64) float64 {
	return t * (2 - t)
}

func EaseInOutQuad(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	} else {
		return -1 + (4-2*t)*t
	}
}

func EaseInCubic(t float64) float64 {
	return t * t * t
}

func EaseOutCubic(t float64) float64 {
	return t*(t-1)*(t-1) + 1
}

func EaseInOutCubic(t float64) float64 {
	if t < 0.5 {
		return 4 * t * t * t
	} else {
		return (t-1)*(2*t-2)*(2*t-2) + 1
	}
}

func EaseInQuart(t float64) float64 {
	return t * t * t * t
}

func EaseOutQuart(t float64) float64 {
	return 1 - t*(t-1)*(t-1)*(t-1)
}

func EaseInOutQuart(t float64) float64 {
	if t < 0.5 {
		return 8 * t * t * t * t
	} else {
		return 1 - 8*t*(t-1)*(t-1)*(t-1)
	}
}

func EaseInQuint(t float64) float64 {
	return t * t * t * t * t
}

func EaseOutQuint(t float64) float64 {
	return 1 + t*(t-1)*(t-1)*(t-1)*(t-1)
}

func EaseInOutQuint(t float64) float64 {
	if t < 0.5 {
		return 16 * t * t * t * t * t
	} else {
		return 1 + 16*t*(t-1)*(t-1)*(t-1)*(t-1)
	}
}
