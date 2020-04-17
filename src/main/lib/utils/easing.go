package utils

type EasingFunction func(t float64) float64

// ease in functions
func EaseInLinear(t float64) float64 {
	return t
}

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
func EaseOutLinear(t float64) float64 {
	return 1 - t
}

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

// composite functions
func Step(x float64) EasingFunction {
	return func(t float64) float64 {
		if t < x {
			return 0
		} else {
			return 1
		}
	}
}

func Spike(x float64, f EasingFunction) EasingFunction {
	fx1 := 1.0 / x
	fx2 := 1.0 / (1.0 - x)
	return func(t float64) float64 {
		if t < x {
			return f(fx1 * t)
		} else {
			return 1 - f(fx2*(t-x))
		}
	}
}

func MirrorSpike(f EasingFunction) EasingFunction {
	return Spike(0.5, f)
}

func Wave(x float64, f EasingFunction) EasingFunction {
	fx1 := 1.0 / x
	fx2 := 1.0 / (1.0 - x)
	return func(t float64) float64 {
		if t < x {
			return f(fx1 * t)
		} else {
			return f(1 - (fx2 * (t - x)))
		}
	}
}
