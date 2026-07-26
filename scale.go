package core

// SCALE aspect constants for the SCALE framework.
// SCALE measures how engineering best practices are Standardized, Consumed,
// Automated, Leveraged, and made Effective across an organization.
const (
	SCALEAspectStandards     = "standards"
	SCALEAspectConsumption   = "consumption"
	SCALEAspectAutomation    = "automation"
	SCALEAspectLeverage      = "leverage"
	SCALEAspectEffectiveness = "effectiveness"
)

// SCALE consumption kind constants.
const (
	SCALEConsumptionKindAdoption    = "adoption"
	SCALEConsumptionKindConformance = "conformance"
)

// AllSCALEAspects returns all SCALE aspect constants in canonical order.
func AllSCALEAspects() []string {
	return []string{
		SCALEAspectStandards,
		SCALEAspectConsumption,
		SCALEAspectAutomation,
		SCALEAspectLeverage,
		SCALEAspectEffectiveness,
	}
}

// ValidSCALEAspect checks if an aspect value is valid.
func ValidSCALEAspect(aspect string) bool {
	switch aspect {
	case SCALEAspectStandards, SCALEAspectConsumption, SCALEAspectAutomation,
		SCALEAspectLeverage, SCALEAspectEffectiveness:
		return true
	default:
		return false
	}
}

// SCALEAspectLetter returns the single-letter abbreviation for an aspect.
func SCALEAspectLetter(aspect string) string {
	letters := map[string]string{
		SCALEAspectStandards:     "S",
		SCALEAspectConsumption:   "C",
		SCALEAspectAutomation:    "A",
		SCALEAspectLeverage:      "L",
		SCALEAspectEffectiveness: "E",
	}
	if letter, ok := letters[aspect]; ok {
		return letter
	}
	return "?"
}

// SCALEAspectDisplayName returns a human-readable name for an aspect.
func SCALEAspectDisplayName(aspect string) string {
	names := map[string]string{
		SCALEAspectStandards:     "Standards",
		SCALEAspectConsumption:   "Consumption",
		SCALEAspectAutomation:    "Automation",
		SCALEAspectLeverage:      "Leverage",
		SCALEAspectEffectiveness: "Effectiveness",
	}
	if name, ok := names[aspect]; ok {
		return name
	}
	return aspect
}

// SCALEAspectDescription returns the question each aspect answers.
func SCALEAspectDescription(aspect string) string {
	descriptions := map[string]string{
		SCALEAspectStandards:     "Do executable engineering standards exist?",
		SCALEAspectConsumption:   "Are teams using them (adoption) and using them correctly (conformance)?",
		SCALEAspectAutomation:    "How much is enforced automatically?",
		SCALEAspectLeverage:      "What reuse and engineering capacity do they create?",
		SCALEAspectEffectiveness: "Did engineering and business outcomes improve?",
	}
	if desc, ok := descriptions[aspect]; ok {
		return desc
	}
	return ""
}

// AllSCALEConsumptionKinds returns all consumption kind constants.
func AllSCALEConsumptionKinds() []string {
	return []string{
		SCALEConsumptionKindAdoption,
		SCALEConsumptionKindConformance,
	}
}

// ValidSCALEConsumptionKind checks if a consumption kind value is valid.
func ValidSCALEConsumptionKind(kind string) bool {
	switch kind {
	case SCALEConsumptionKindAdoption, SCALEConsumptionKindConformance, "":
		return true
	default:
		return false
	}
}
