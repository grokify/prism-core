package core

// Risk probability constants.
const (
	RiskProbabilityLow    = "low"
	RiskProbabilityMedium = "medium"
	RiskProbabilityHigh   = "high"
)

// Risk impact constants.
const (
	RiskImpactLow      = "low"
	RiskImpactMedium   = "medium"
	RiskImpactHigh     = "high"
	RiskImpactCritical = "critical"
)

// Risk status constants.
const (
	RiskStatusOpen      = "open"
	RiskStatusMitigated = "mitigated"
	RiskStatusAccepted  = "accepted"
	RiskStatusClosed    = "closed"
)

// Risk represents a risk item.
type Risk struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description"`
	Impact      string `json:"impact,omitempty"`
	Probability string `json:"probability,omitempty"`
	Mitigation  string `json:"mitigation,omitempty"`
	Status      string `json:"status,omitempty"`
	Owner       string `json:"owner,omitempty"`
}

// ValidRiskProbability checks if a risk probability value is valid.
func ValidRiskProbability(probability string) bool {
	switch probability {
	case RiskProbabilityLow, RiskProbabilityMedium, RiskProbabilityHigh, "":
		return true
	default:
		return false
	}
}

// ValidRiskImpact checks if a risk impact value is valid.
func ValidRiskImpact(impact string) bool {
	switch impact {
	case RiskImpactLow, RiskImpactMedium, RiskImpactHigh, RiskImpactCritical, "":
		return true
	default:
		return false
	}
}

// ValidRiskStatus checks if a risk status value is valid.
func ValidRiskStatus(status string) bool {
	switch status {
	case RiskStatusOpen, RiskStatusMitigated, RiskStatusAccepted, RiskStatusClosed, "":
		return true
	default:
		return false
	}
}

// RiskScore calculates a numeric score based on probability and impact.
// Returns a value from 1 (low/low) to 12 (high/critical).
func RiskScore(probability, impact string) int {
	probScore := map[string]int{
		RiskProbabilityLow:    1,
		RiskProbabilityMedium: 2,
		RiskProbabilityHigh:   3,
	}
	impactScore := map[string]int{
		RiskImpactLow:      1,
		RiskImpactMedium:   2,
		RiskImpactHigh:     3,
		RiskImpactCritical: 4,
	}

	p := probScore[probability]
	if p == 0 {
		p = 1
	}
	i := impactScore[impact]
	if i == 0 {
		i = 1
	}
	return p * i
}

// RiskSeverity returns a severity level based on risk score.
func RiskSeverity(score int) string {
	switch {
	case score >= 9:
		return "critical"
	case score >= 6:
		return "high"
	case score >= 3:
		return "medium"
	default:
		return "low"
	}
}
