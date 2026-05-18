package core

// Maturity level constants (M1-M5).
const (
	MaturityLevel1 = 1 // Reactive/Ad-hoc
	MaturityLevel2 = 2 // Basic
	MaturityLevel3 = 3 // Defined
	MaturityLevel4 = 4 // Managed
	MaturityLevel5 = 5 // Optimizing
)

// Maturity level name constants.
const (
	MaturityNameReactive   = "Reactive"
	MaturityNameBasic      = "Basic"
	MaturityNameDefined    = "Defined"
	MaturityNameManaged    = "Managed"
	MaturityNameOptimizing = "Optimizing"
)

// MaturityLevelDef defines a maturity level with metadata.
type MaturityLevelDef struct {
	Level       int     `json:"level"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
}

// MaturityLevelName returns the name for a maturity level.
func MaturityLevelName(level int) string {
	switch level {
	case MaturityLevel1:
		return MaturityNameReactive
	case MaturityLevel2:
		return MaturityNameBasic
	case MaturityLevel3:
		return MaturityNameDefined
	case MaturityLevel4:
		return MaturityNameManaged
	case MaturityLevel5:
		return MaturityNameOptimizing
	default:
		return "Unknown"
	}
}

// MaturityLevelDescription returns a description for a maturity level.
func MaturityLevelDescription(level int) string {
	switch level {
	case MaturityLevel1:
		return "Ad-hoc processes, firefighting mode, inconsistent outcomes"
	case MaturityLevel2:
		return "Basic controls, some documentation, individual-dependent"
	case MaturityLevel3:
		return "Standardized processes, repeatable across teams"
	case MaturityLevel4:
		return "Data-driven, measured and controlled, integrated"
	case MaturityLevel5:
		return "Continuous improvement, automated optimization, adaptive"
	default:
		return "Unknown maturity level"
	}
}

// ValidMaturityLevel checks if a maturity level is valid (1-5).
func ValidMaturityLevel(level int) bool {
	return level >= MaturityLevel1 && level <= MaturityLevel5
}

// MaturityLevelShortName returns short name like "M1", "M2", etc.
func MaturityLevelShortName(level int) string {
	if level >= 1 && level <= 5 {
		return "M" + string(rune('0'+level))
	}
	return "M?"
}

// DefaultMaturityLevels returns the standard 5-level maturity definitions.
func DefaultMaturityLevels() []MaturityLevelDef {
	return []MaturityLevelDef{
		{
			Level:       MaturityLevel1,
			Name:        MaturityNameReactive,
			Description: "Ad-hoc processes, firefighting mode, inconsistent outcomes",
			Weight:      0.2,
		},
		{
			Level:       MaturityLevel2,
			Name:        MaturityNameBasic,
			Description: "Basic controls, some documentation, individual-dependent",
			Weight:      0.4,
		},
		{
			Level:       MaturityLevel3,
			Name:        MaturityNameDefined,
			Description: "Standardized processes, repeatable across teams",
			Weight:      0.6,
		},
		{
			Level:       MaturityLevel4,
			Name:        MaturityNameManaged,
			Description: "Data-driven, measured and controlled, integrated",
			Weight:      0.8,
		},
		{
			Level:       MaturityLevel5,
			Name:        MaturityNameOptimizing,
			Description: "Continuous improvement, automated optimization, adaptive",
			Weight:      1.0,
		},
	}
}

// MaturityGap calculates the gap between current and target maturity levels.
func MaturityGap(current, target int) int {
	if target > current {
		return target - current
	}
	return 0
}

// MaturityProgress calculates progress as a percentage (0.0 to 1.0).
func MaturityProgress(current, target int) float64 {
	if target <= 1 {
		return 1.0
	}
	if current >= target {
		return 1.0
	}
	return float64(current-1) / float64(target-1)
}
