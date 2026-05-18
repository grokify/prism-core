package core

// Status constants for document lifecycle.
const (
	StatusDraft      = "draft"
	StatusInReview   = "in_review"
	StatusApproved   = "approved"
	StatusActive     = "active"
	StatusCompleted  = "completed"
	StatusDeprecated = "deprecated"
	StatusArchived   = "archived"
)

// HealthStatus constants for metric health indicators.
const (
	HealthGreen  = "green"
	HealthYellow = "yellow"
	HealthRed    = "red"
)

// Priority constants for prioritization.
const (
	PriorityCritical = "critical"
	PriorityHigh     = "high"
	PriorityMedium   = "medium"
	PriorityLow      = "low"
)

// MoSCoW prioritization constants.
const (
	MoSCoWMust   = "must"
	MoSCoWShould = "should"
	MoSCoWCould  = "could"
	MoSCoWWont   = "wont"
)

// CapabilityStatus constants for capability lifecycle.
const (
	CapabilityStatusPlanned     = "planned"
	CapabilityStatusInProgress  = "in_progress"
	CapabilityStatusImplemented = "implemented"
	CapabilityStatusOperational = "operational"
	CapabilityStatusDeprecated  = "deprecated"
)

// InitiativeStatus constants for initiative tracking.
const (
	InitiativeStatusPlanned    = "planned"
	InitiativeStatusInProgress = "in_progress"
	InitiativeStatusCompleted  = "completed"
	InitiativeStatusCancelled  = "cancelled"
	InitiativeStatusOnHold     = "on_hold"
)

// ValidPriority checks if a priority value is valid.
func ValidPriority(priority string) bool {
	switch priority {
	case PriorityCritical, PriorityHigh, PriorityMedium, PriorityLow, "":
		return true
	default:
		return false
	}
}

// ValidMoSCoW checks if a MoSCoW value is valid.
func ValidMoSCoW(moscow string) bool {
	switch moscow {
	case MoSCoWMust, MoSCoWShould, MoSCoWCould, MoSCoWWont, "":
		return true
	default:
		return false
	}
}

// ValidHealthStatus checks if a health status value is valid.
func ValidHealthStatus(status string) bool {
	switch status {
	case HealthGreen, HealthYellow, HealthRed, "":
		return true
	default:
		return false
	}
}

// PriorityWeight returns a numeric weight for sorting priorities.
// Higher weight = higher priority.
func PriorityWeight(priority string) int {
	switch priority {
	case PriorityCritical:
		return 4
	case PriorityHigh:
		return 3
	case PriorityMedium:
		return 2
	case PriorityLow:
		return 1
	default:
		return 0
	}
}

// MoSCoWWeight returns a numeric weight for sorting MoSCoW priorities.
func MoSCoWWeight(moscow string) int {
	switch moscow {
	case MoSCoWMust:
		return 4
	case MoSCoWShould:
		return 3
	case MoSCoWCould:
		return 2
	case MoSCoWWont:
		return 1
	default:
		return 0
	}
}
