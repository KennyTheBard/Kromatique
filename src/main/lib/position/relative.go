package position

const (
	Start  = iota
	Center = iota
	End    = iota
)

func getRelative(start, end, mode int) int {
	switch mode {
	case Start:
		return start
	case End:
		return end
	case Center:
		return (start + end) / 2
	default:
		return 0
	}
}
