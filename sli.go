package core

// SLI type constants.
const (
	SLITypeAvailability = "availability"
	SLITypeLatency      = "latency"
	SLITypeThroughput   = "throughput"
	SLITypeErrorRate    = "error_rate"
	SLITypeFreshness    = "freshness"
	SLITypeCorrectness  = "correctness"
	SLITypeCoverage     = "coverage"
	SLITypeDurability   = "durability"
)

// SLI comparison direction constants.
const (
	SLIDirectionHigherIsBetter = "higher_is_better"
	SLIDirectionLowerIsBetter  = "lower_is_better"
)

// SLI aggregation method constants.
const (
	SLIAggregationAverage    = "average"
	SLIAggregationPercentile = "percentile"
	SLIAggregationSum        = "sum"
	SLIAggregationMin        = "min"
	SLIAggregationMax        = "max"
	SLIAggregationCount      = "count"
)

// SLITypeDef defines an SLI type with metadata.
type SLITypeDef struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Direction   string `json:"direction"`
	Unit        string `json:"unit,omitempty"`
}

// AllSLITypes returns all SLI type constants.
func AllSLITypes() []string {
	return []string{
		SLITypeAvailability,
		SLITypeLatency,
		SLITypeThroughput,
		SLITypeErrorRate,
		SLITypeFreshness,
		SLITypeCorrectness,
		SLITypeCoverage,
		SLITypeDurability,
	}
}

// ValidSLIType checks if an SLI type value is valid.
func ValidSLIType(sliType string) bool {
	switch sliType {
	case SLITypeAvailability, SLITypeLatency, SLITypeThroughput, SLITypeErrorRate,
		SLITypeFreshness, SLITypeCorrectness, SLITypeCoverage, SLITypeDurability, "":
		return true
	default:
		return false
	}
}

// SLITypeDirection returns the default comparison direction for an SLI type.
func SLITypeDirection(sliType string) string {
	switch sliType {
	case SLITypeAvailability, SLITypeThroughput, SLITypeFreshness,
		SLITypeCorrectness, SLITypeCoverage, SLITypeDurability:
		return SLIDirectionHigherIsBetter
	case SLITypeLatency, SLITypeErrorRate:
		return SLIDirectionLowerIsBetter
	default:
		return SLIDirectionHigherIsBetter
	}
}

// SLITypeDisplayName returns a human-readable name for an SLI type.
func SLITypeDisplayName(sliType string) string {
	names := map[string]string{
		SLITypeAvailability: "Availability",
		SLITypeLatency:      "Latency",
		SLITypeThroughput:   "Throughput",
		SLITypeErrorRate:    "Error Rate",
		SLITypeFreshness:    "Freshness",
		SLITypeCorrectness:  "Correctness",
		SLITypeCoverage:     "Coverage",
		SLITypeDurability:   "Durability",
	}
	if name, ok := names[sliType]; ok {
		return name
	}
	return sliType
}

// DefaultSLITypes returns default SLI type definitions.
func DefaultSLITypes() []SLITypeDef {
	return []SLITypeDef{
		{Type: SLITypeAvailability, Name: "Availability", Description: "Service uptime and reachability", Direction: SLIDirectionHigherIsBetter, Unit: "percent"},
		{Type: SLITypeLatency, Name: "Latency", Description: "Request/response time", Direction: SLIDirectionLowerIsBetter, Unit: "ms"},
		{Type: SLITypeThroughput, Name: "Throughput", Description: "Requests or transactions per unit time", Direction: SLIDirectionHigherIsBetter, Unit: "req/s"},
		{Type: SLITypeErrorRate, Name: "Error Rate", Description: "Percentage of failed requests", Direction: SLIDirectionLowerIsBetter, Unit: "percent"},
		{Type: SLITypeFreshness, Name: "Freshness", Description: "Data age or staleness", Direction: SLIDirectionHigherIsBetter, Unit: "seconds"},
		{Type: SLITypeCorrectness, Name: "Correctness", Description: "Accuracy of results", Direction: SLIDirectionHigherIsBetter, Unit: "percent"},
		{Type: SLITypeCoverage, Name: "Coverage", Description: "Proportion of valid responses", Direction: SLIDirectionHigherIsBetter, Unit: "percent"},
		{Type: SLITypeDurability, Name: "Durability", Description: "Data persistence and reliability", Direction: SLIDirectionHigherIsBetter, Unit: "percent"},
	}
}
