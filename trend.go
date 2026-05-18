package core

// Trend direction constants.
const (
	TrendUp      = "up"
	TrendDown    = "down"
	TrendFlat    = "flat"
	TrendUnknown = "unknown"
)

// TrendInfo represents a trend with direction and magnitude.
type TrendInfo struct {
	Direction  string  `json:"direction"`
	Magnitude  float64 `json:"magnitude,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
}

// ValidTrend checks if a trend direction value is valid.
func ValidTrend(direction string) bool {
	switch direction {
	case TrendUp, TrendDown, TrendFlat, TrendUnknown, "":
		return true
	default:
		return false
	}
}

// TrendFromDelta returns a trend direction based on a delta value.
func TrendFromDelta(delta float64) string {
	switch {
	case delta > 0:
		return TrendUp
	case delta < 0:
		return TrendDown
	default:
		return TrendFlat
	}
}

// TrendIcon returns an emoji icon for a trend direction.
func TrendIcon(direction string) string {
	switch direction {
	case TrendUp:
		return "↑"
	case TrendDown:
		return "↓"
	case TrendFlat:
		return "→"
	default:
		return "?"
	}
}

// NewTrendInfo creates a TrendInfo from a delta value.
func NewTrendInfo(delta, previous float64) TrendInfo {
	info := TrendInfo{
		Direction: TrendFromDelta(delta),
		Magnitude: delta,
	}
	if previous != 0 {
		info.Percentage = (delta / previous) * 100
	}
	return info
}
