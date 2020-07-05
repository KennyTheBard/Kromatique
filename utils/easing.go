package utils

// EasingFunction returns a value between 0 and 1
// for a given factor between 0 and 1
type EasingFunction func(t float64) float64

// EaseInLinear based on https://gist.github.com/gre/1650294
func EaseInLinear(t float64) float64 {
	return t
}

// EaseInQuad based on https://gist.github.com/gre/1650294
func EaseInQuad(t float64) float64 {
	return t * t
}

// EaseInCubic based on https://gist.github.com/gre/1650294
func EaseInCubic(t float64) float64 {
	return t * EaseInQuad(t)
}

// EaseInQuart based on https://gist.github.com/gre/1650294
func EaseInQuart(t float64) float64 {
	return t * EaseInCubic(t)
}

// EaseOutLinear based on https://gist.github.com/gre/1650294
func EaseOutLinear(t float64) float64 {
	return 1 - t
}

// EaseOutQuad based on https://gist.github.com/gre/1650294
func EaseOutQuad(t float64) float64 {
	return 1 - EaseInQuad(1-t)
}

// EaseOutCubic based on https://gist.github.com/gre/1650294
func EaseOutCubic(t float64) float64 {
	return 1 - EaseInCubic(1-t)
}

// EaseOutQuart based on https://gist.github.com/gre/1650294
func EaseOutQuart(t float64) float64 {
	return 1 - EaseInQuart(1-t)
}

// EaseInOutQuad based on https://gist.github.com/gre/1650294
func EaseInOutQuad(t float64) float64 {
	if t < 0.5 {
		return EaseInQuad(2*t) / 2
	} else {
		return EaseOutQuad(2*(t-0.5)) / 2
	}
}

// EaseInOutCubic based on https://gist.github.com/gre/1650294
func EaseInOutCubic(t float64) float64 {
	if t < 0.5 {
		return EaseInCubic(2*t) / 2
	} else {
		return EaseOutCubic(2*(t-0.5)) / 2
	}
}

// EaseInOutQuart based on https://gist.github.com/gre/1650294
func EaseInOutQuart(t float64) float64 {
	if t < 0.5 {
		return EaseInQuart(2*t) / 2
	} else {
		return EaseOutQuart(2*(t-0.5)) / 2
	}
}

// Step returns a step function with the given threshold
func Step(threshold float64) EasingFunction {
	return func(t float64) float64 {
		if t < threshold {
			return 0
		} else {
			return 1
		}
	}
}

// Spike returns a spike function with the given threshold
// for them maximum value and an EasingFunction for the 2 parts
func Spike(threshold float64, f EasingFunction) EasingFunction {
	fx1 := 1.0 / threshold
	fx2 := 1.0 / (1.0 - threshold)
	return func(t float64) float64 {
		if t < threshold {
			return f(fx1 * t)
		} else {
			return 1 - f(fx2*(t-threshold))
		}
	}
}

// MirrorSpike returns a spike function with the maximum in the middle
func MirrorSpike(f EasingFunction) EasingFunction {
	return Spike(0.5, f)
}

// Wave returns a spike function with the shape of a wave
func Wave(threshold float64, f EasingFunction) EasingFunction {
	fx1 := 1.0 / threshold
	fx2 := 1.0 / (1.0 - threshold)
	return func(t float64) float64 {
		if t < threshold {
			return f(fx1 * t)
		} else {
			return f(1 - (fx2 * (t - threshold)))
		}
	}
}
